namespace go rpc

service AiInvestRPCService {
    TestResponse Test(1: TestRequest req)
}

struct TestRequest {
    1: string message
}

struct TestResponse {
    1: string message
}