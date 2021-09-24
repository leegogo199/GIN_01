package dto

import (
	"GINVUE/Model"
)

type UserDto struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

func ToUserDto(User Model.User) UserDto {
	return UserDto{
		Name:      User.Name,
		Telephone: User.Telephone,
	}

}
