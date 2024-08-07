package initialize

import (
	"fmt"
	retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	proto "public/api/qvbilam/public/v1"
	userProto "public/api/qvbilam/user/v1"
	"public/global"
	"time"
)

type dialConfig struct {
	host string
	port int64
}

type serverClientConfig struct {
	userDialConfig *dialConfig
	fileDialConfig *dialConfig
}

func InitServer() {
	s := serverClientConfig{
		userDialConfig: &dialConfig{
			host: global.ServerConfig.UserServerConfig.Host,
			port: global.ServerConfig.UserServerConfig.Port,
		},
		fileDialConfig: &dialConfig{
			host: global.ServerConfig.PublicServerConfig.Host,
			port: global.ServerConfig.PublicServerConfig.Port,
		},
	}

	s.initFileServer()
	s.initUserServer()
}

func clientOption() []retry.CallOption {
	opts := []retry.CallOption{
		retry.WithBackoff(retry.BackoffLinear(100 * time.Millisecond)), // 重试间隔
		retry.WithMax(3), // 最大重试次数
		retry.WithPerRetryTimeout(1 * time.Second),                                 // 请求超时时间
		retry.WithCodes(codes.NotFound, codes.DeadlineExceeded, codes.Unavailable), // 指定返回码重试
	}
	return opts
}

func (s *serverClientConfig) initFileServer() {
	opts := clientOption()
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", s.fileDialConfig.host, s.fileDialConfig.port),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(retry.UnaryClientInterceptor(opts...)))
	if err != nil {
		zap.S().Fatalf("%s dial error: %s", global.ServerConfig.PublicServerConfig.Name, err)
	}

	videoClient := proto.NewVideoClient(conn)
	fileClient := proto.NewFileClient(conn)
	smsClient := proto.NewSmsClient(conn)

	global.FileVideoServerClient = videoClient
	global.FileServerClient = fileClient
	global.SmsServerClient = smsClient
}

func (s *serverClientConfig) initUserServer() {
	opts := clientOption()

	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", s.userDialConfig.host, s.userDialConfig.port),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(retry.UnaryClientInterceptor(opts...)))
	if err != nil {
		zap.S().Fatalf("%s dial error: %s", global.ServerConfig.UserServerConfig.Name, err)
	}

	userClient := userProto.NewUserClient(conn)
	global.UserServerClient = userClient
	//u, err := global.UserServerClient.Detail(context.Background(), &userProto.GetUserRequest{Id: 1})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(u)
}
