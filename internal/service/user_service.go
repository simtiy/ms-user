package service

import (
	"log"
	"ms-user/internal/dto"
	"ms-user/internal/model"
	"ms-user/internal/repository"
)

type UserService struct {
	userRepository repository.IUserRepository
	logger         *log.Logger
}

func NewUserService(userRepository repository.IUserRepository, logger *log.Logger) *UserService {
	return &UserService{
		userRepository: userRepository,
		logger:         logger,
	}
}

func (u *UserService) CreateUser(req dto.CreateUserRequest) (dto.CreateUserResponse, error) {
	u.logger.Printf("Create user body=%+v\n", req)

	//map to entity
	userEntity := model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	user, err := u.userRepository.Create(userEntity)
	if err != nil {
		return dto.CreateUserResponse{}, err // @todo create custom exception and handler
	}

	//map to response
	var res = dto.CreateUserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	u.logger.Printf("Create user response body=%+v\n", res)
	return res, nil
}

func (u *UserService) DeleteUser(id int64) error {
	u.logger.Printf("Delete user by ID=%d\\n", id)
	return u.userRepository.Delete(id)
}

func (u *UserService) GetAllUsers() ([]dto.CreateUserResponse, error) {
	u.logger.Println("Get all users")
	users, err := u.userRepository.GetAll()

	if err != nil {
		return nil, err // @todo create custom exception and handler
	}

	//declare response variable
	var res []dto.CreateUserResponse

	//created for loop for user response
	for _, user := range users {
		res = append(res, dto.CreateUserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		})
	}

	u.logger.Printf("Get all users response=%+v\n", res)
	return res, nil
}

func (u *UserService) GetUserById(id int64) (dto.CreateUserResponse, error) {
	u.logger.Printf("Get user by ID=%d\\n ", id)
	user, err := u.userRepository.GetByID(id)

	if err != nil {
		return dto.CreateUserResponse{}, err // @todo create custom exception and handler
	}

	// map to response
	var res = dto.CreateUserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
	u.logger.Printf("Get user by ID response=%+v\n", res)

	return res, nil
}
