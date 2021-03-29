package model

import "api-grpc-web/pkg/proto"

type User struct {
	Username   string
	Name       string
	Email      string
	Company    string
	AvatarUrl  string
	ProfileUrl string
	ReposUrl   string
	Followers  int64
	Following  int64
}

func ToUser(githubUser *proto.UserResponse) User {
	return User{
		Username:   githubUser.Username,
		Name:       githubUser.Name,
		Email:      githubUser.Email,
		Company:    githubUser.Company,
		AvatarUrl:  githubUser.AvatarUrl,
		ProfileUrl: githubUser.ProfileUrl,
		ReposUrl:   githubUser.ReposUrl,
		Followers:  githubUser.Followers,
		Following:  githubUser.Following,
	}
}