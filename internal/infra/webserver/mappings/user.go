package webserver_mappings

import (
	"errors"
)

const (
	ErrInvalidUser = "Invalid user"
)

func MapCreateUserRequestDTOToUser(createUserRequestDTO *webserver_dtos.CreateUserRequestDTO) (*user_domain.User, error) {
	if createUserRequestDTO == nil {
		return nil, errors.New(ErrInvalidUser)
	}

	return &user_domain.User{
		Nickname: createUserRequestDTO.Nickname,
		Password: createUserRequestDTO.Password,
	}, nil
}
