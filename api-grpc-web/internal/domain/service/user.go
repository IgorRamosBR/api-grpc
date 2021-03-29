package service

import (
	"api-grpc-web/internal/domain/model"
	"api-grpc-web/pkg/proto"
	"context"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
)

type UsersService interface {
	FindUser(ctx context.Context, username string) (user model.User, err error)
}

type UsersServiceImpl struct {

}

func NewUserService() UsersService {
	return UsersServiceImpl{}
}

func (UsersServiceImpl) FindUser(ctx context.Context, username string) (user model.User, err error) {
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		return model.User{}, err
	}
	defer conn.Close()

	client := proto.NewUserServiceClient(conn)
	githubUser, err := client.FindGithubUser(ctx, &proto.UserRequest{Username: username})
	if err != nil {
		log.Error(err)
	}

	return model.ToUser(githubUser), nil
}