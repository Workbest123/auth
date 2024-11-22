package auth
import (
    "context"

    "google.golang.org/grpc"

    // Сгенерированный код
    ssov1 "github.com/GolangLessons/protos/gen/go/sso"
)

type serverAPI struct {
    ssov1.UnimplementedAuthServer // Хитрая штука, о ней ниже
}

func Register(gRPCServer *grpc.Server) {  
    ssov1.RegisterAuthServer(gRPCServer, &serverAPI{})  
}

func (s *serverAPI) Login(
    ctx context.Context,
    in *ssov1.LoginRequest,
) (*ssov1.LoginResponse, error) {
    panic("implement me")
}

func (s *serverAPI) Register(
    ctx context.Context,
    in *ssov1.RegisterRequest,
) (*ssov1.RegisterResponse, error) {
    panic("implement me")
}
func (s *serverAPI) IsAdmin(
    ctx context.Context,
    in *ssov1.IsAdminRequest,
) (*ssov1.IsAdminResponse, error) {
    panic("implement me")
}