package conf

import (
	"fmt"
	"os"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	TokenConfig *Token
	configMutex sync.RWMutex // 添加读写锁保证线程安全
)

type Token struct {
	LongBridge struct {
		AppKey      string `mapstructure:"app_key"`
		AppSecret   string `mapstructure:"app_secret"`
		AccessToken string `mapstructure:"access_token"`
		Region      string `mapstructure:"region"`
	} `mapstructure:"long_bridge"`
}

func InitConfig() {
	// 初始化 Viper
	v := viper.New()
	v.SetConfigName("token")    // 配置文件名称（无扩展名）
	v.SetConfigType("yaml")     // 格式
	v.AddConfigPath("../conf/") // 搜索路径

	// 首次加载配置
	if err := loadConfig(v); err != nil {
		panic(err)
	}

	// 设置监听
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("检测到配置文件变更:", e.Name)
		if err := loadConfig(v); err != nil {
			fmt.Printf("配置重载失败: %v\n", err)
		} else {
			fmt.Println("配置已热更新")
		}
	})
}

// 封装加载配置的公共逻辑
func loadConfig(v *viper.Viper) error {
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("配置文件读取失败: %w", err)
	}

	newConfig := &Token{}
	if err := v.Unmarshal(newConfig); err != nil {
		return fmt.Errorf("配置解析失败: %w", err)
	}

	// 使用写锁更新配置
	configMutex.Lock()
	defer configMutex.Unlock()
	TokenConfig = newConfig

	os.Setenv("LONGPORT_APP_KEY", newConfig.LongBridge.AppKey)
	os.Setenv("LONGPORT_APP_SECRET", newConfig.LongBridge.AppSecret)
	os.Setenv("LONGPORT_ACCESS_TOKEN", newConfig.LongBridge.AccessToken)
	os.Setenv("LONGPORT_REGION", newConfig.LongBridge.Region)

	return nil
}

// GetConfigToken 安全获取配置的封装方法
func GetConfigToken() *Token {
	configMutex.RLock()
	defer configMutex.RUnlock()
	return TokenConfig
}
