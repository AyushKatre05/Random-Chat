package handlers


type userHandler struct{}

func NewUserHandler() *userHandler {
	return &userHandler{}
}

func (uh *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser webserver_dtos.CreateUserRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Error decodifying JSON", http.StatusBadRequest)
		return
	}

	createUserRepository := user_repository.NewCreateUserRepository(database.MongoInstance.Db)
	createUserUseCase := user_data_usecase.NewCreateUserUseCase(createUserRepository)

	mappedUser, err := webserver_mappings.MapCreateUserRequestDTOToUser(&newUser)
	if err != nil {
		http.Error(w, "Error mapping user", http.StatusBadRequest)
	}
	mappedUser.Id = uuid.New().String()

	err = createUserUseCase.Execute(mappedUser)
	if err != nil {
		http.Error(w, "Error creating new user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully."))
}
