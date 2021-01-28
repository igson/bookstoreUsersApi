package users

import "encoding/json"

//PublicUser classe de usuário
type PublicUser struct {
	ID int64 `json:"id"`
	//	FirstName   string `json:"first_name"`
	//	LastName    string `json:"last_name"`
	//	Email       string `json:"email"`
	Status      string `json:"status"`
	DateCreated string `json:"date_created"`
	//Password    string `json:"password"`
}

//PrivatelUser classe de usuário
type PrivatelUser struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Status      string `json:"status"`
	DateCreated string `json:"date_created"`
	//	Password    string `json:"password"`
}

//Marshall serializar
func (u Users) Marshall(isPublic bool) []interface{} {

	usuarios := make([]interface{}, len(u))

	for index, usuario := range u {
		usuarios[index] = usuario.Marshall(isPublic)
	}

	return usuarios

}

//Marshall serializar
func (user *User) Marshall(isPublic bool) interface{} {

	if isPublic {
		return PublicUser{
			ID:          user.ID,
			DateCreated: user.DateCreated,
			Status:      user.Status,
		}
	}

	userJSON, _ := json.Marshal(user)

	var privateUser PrivatelUser

	json.Unmarshal(userJSON, &privateUser)

	return privateUser

}
