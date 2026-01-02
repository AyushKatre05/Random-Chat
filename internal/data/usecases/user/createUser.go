package user_data_usecase


type CreateUserUseCase struct {
	createUserRepository repositories_data.CreateUserRepository
}

func NewCreateUserUseCase(createUserRepository repositories_data.CreateUserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{createUserRepository: createUserRepository}
}

func (createUserUseCase *CreateUserUseCase) Execute(user *user_domain.User) error {
	if validateErr := user.Validate(); validateErr != nil {
		return validateErr
	}

	mappedUser, err := mappings.MapUserToCreateUserInputDTO(user)
	if err != nil {
		return err
	}

	return createUserUseCase.createUserRepository.Create(mappedUser)
}
