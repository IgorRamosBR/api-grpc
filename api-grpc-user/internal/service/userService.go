package service

import (
	"api-grpc-user/pkg/proto"
	"context"
	"github.com/google/go-github/v33/github"
)

type UserServiceImpl struct {
	Client *github.Client
}

func NewUserService(client *github.Client) *UserServiceImpl {
	return &UserServiceImpl{Client: client}
}

func (s *UserServiceImpl) FindGithubUser(ctx context.Context, req *proto.UserRequest) (*proto.UserResponse, error) {
	user, _, err := s.Client.Users.Get(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	return createUserResponse(user), nil
}

func createUserResponse(user *github.User) *proto.UserResponse{
	return &proto.UserResponse{
		Username:   extractString(user.Login),
		AvatarUrl:  extractString(user.AvatarURL),
		ProfileUrl: extractString(user.HTMLURL),
		ReposUrl:   extractString(user.ReposURL),
		Name:       extractString(user.Name),
		Email:      extractString(user.Email),
		Company:    extractString(user.Company),
		Followers:  int64(extractInt(user.Followers)),
		Following:  int64(extractInt(user.Following)),
	}
}

func extractString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

func extractInt(i *int) int {
	if i != nil {
		return *i
	}
	return 0
}