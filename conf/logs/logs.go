package logs

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const (
	INFO  = "info"
	ERROR = "error"
	DEBUG = "debug"
)

var (
	logDir          string
	retainDays      int
	currentDate     string
	fileHandles     = make(map[string]*os.File)
	loggers         = make(map[string]*log.Logger)
	mu              sync.Mutex
	cleanupTicker   *time.Ticker
	stopCleanupChan chan struct{}
)

func InitLogger() error {
	mu.Lock()
	defer mu.Unlock()

	logDir = "./logs"
	retainDays = 3
	currentDate = time.Now().Format("2006-01-02")

	// 创建日志目录
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return fmt.Errorf("failed to create log directory: %v", err)
	}

	// 启动定时清理任务
	if cleanupTicker == nil {
		cleanupTicker = time.NewTicker(24 * time.Hour)
		stopCleanupChan = make(chan struct{})
		go startCleanupTask()
	}

	return nil
}

func Close() {
	mu.Lock()
	defer mu.Unlock()

	// 关闭所有文件句柄
	for level, f := range fileHandles {
		f.Close()
		delete(fileHandles, level)
	}

	// 停止清理任务
	if cleanupTicker != nil {
		cleanupTicker.Stop()
		close(stopCleanupChan)
		cleanupTicker = nil
	}
}

func startCleanupTask() {
	for {
		select {
		case <-cleanupTicker.C:
			cleanupOldLogs()
		case <-stopCleanupChan:
			return
		}
	}
}

func cleanupOldLogs() {
	cutoffDate := time.Now().AddDate(0, 0, -retainDays)

	filepath.Walk(logDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if info.IsDir() && info.Name() != filepath.Base(logDir) {
			dirDate, err := time.Parse("2006-01-02", info.Name())
			if err == nil && dirDate.Before(cutoffDate) {
				os.RemoveAll(path)
			}
		}
		return nil
	})
}

func getLogger(level string) (*log.Logger, error) {
	mu.Lock()
	defer mu.Unlock()

	// 检查日期是否变化
	today := time.Now().Format("2006-01-02")
	if today != currentDate {
		// 关闭旧文件句柄
		for _, f := range fileHandles {
			f.Close()
		}
		fileHandles = make(map[string]*os.File)
		loggers = make(map[string]*log.Logger)
		currentDate = today
	}

	// 如果已存在直接返回
	if logger, exists := loggers[level]; exists {
		return logger, nil
	}

	// 创建日期目录
	dateDir := filepath.Join(logDir, currentDate)
	if err := os.MkdirAll(dateDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create date directory: %v", err)
	}

	// 创建日志文件
	filePath := filepath.Join(dateDir, level+".log")
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to create log file: %v", err)
	}

	// 创建Logger实例
	logger := log.New(f, "["+level+"] ", log.LstdFlags|log.Lmicroseconds)
	fileHandles[level] = f
	loggers[level] = logger

	return logger, nil
}

func logMessage(level string, format string, v ...interface{}) {
	if logDir == "" {
		return
	}

	logger, err := getLogger(level)
	if err != nil {
		fmt.Printf("Failed to get logger: %v\n", err)
		return
	}

	msg := fmt.Sprintf(format, v...)
	logger.Println(msg)
}

func Info(format string, v ...interface{}) {
	logMessage(INFO, format, v...)
}

func Error(format string, v ...interface{}) {
	logMessage(ERROR, format, v...)
}

func Debug(format string, v ...interface{}) {
	logMessage(DEBUG, format, v...)
}
