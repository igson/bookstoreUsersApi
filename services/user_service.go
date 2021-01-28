package services

import (
	"github.com/igson/bookstoreUsersApi/domain/users"
	"github.com/igson/bookstoreUsersApi/utils/crypto"
	"github.com/igson/bookstoreUsersApi/utils/dateutils"
	"github.com/igson/bookstoreUsersApi/utils/errors"
)

var (
	// UserService acesso a camada de serviço do usuário
	UserService userServiceInterface = &userService{}
)

type userService struct{}

type userServiceInterface interface {
	GetUser(int64) (*users.User, *errors.RestErroAPI)
	CreateUser(users.User) (*users.User, *errors.RestErroAPI)
	UpdateUser(bool, users.User) (*users.User, *errors.RestErroAPI)
	DeleteUser(userID int64) *errors.RestErroAPI
	Search(status string) (users.Users, *errors.RestErroAPI)
}

//GetUser buscar usuario pelo ID
func (s *userService) GetUser(userID int64) (*users.User, *errors.RestErroAPI) {

	user := &users.User{ID: userID}

	if erro := user.GetUser(); erro != nil {
		return nil, erro
	}

	return user, nil

}

//CreateUser camada de regra de négocio pra cadastro de usuários
func (s *userService) CreateUser(user users.User) (*users.User, *errors.RestErroAPI) {

	if erro := user.Validate(); erro != nil {
		return nil, erro
	}

	user.Status = users.UserStatus
	user.Password = crypto.GetMD5(user.Password)
	user.DateCreated = dateutils.GetNowDBFormat()

	if erro := user.Save(); erro != nil {
		return nil, erro
	}

	return &user, nil
}

//UpdateUser camada de regra de négocio pra atualizar cadastro de usuários
func (s *userService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErroAPI) {

	getUser := &users.User{ID: user.ID}

	if erro := getUser.GetUser(); erro != nil {
		return nil, erro
	}

	if isPartial {

		if user.FirstName != "" {
			getUser.FirstName = user.FirstName
		}
		if user.LastName != "" {
			getUser.LastName = user.LastName
		}
		if user.Email != "" {
			getUser.Email = user.Email
		}

	} else {
		getUser.Email = user.Email
		getUser.LastName = user.LastName
		getUser.FirstName = user.FirstName
	}

	if erro := getUser.Update(); erro != nil {
		return nil, erro
	}

	return getUser, nil

}

//DeleteUser deletar usuário
func (s *userService) DeleteUser(userID int64) *errors.RestErroAPI {
	user := &users.User{ID: userID}
	return user.Delete()
}

//Search listar usuários por status
func (s *userService) Search(status string) (users.Users, *errors.RestErroAPI) {
	dao := &users.User{}
	return dao.Search(status)
}
