package response

import "User-API/domain/entity"

func NewUserResponse(user *entity.User) UserResponse {
	return UserResponse{
		Name: user.Name,
		Id:   user.Id,
		Mail: user.Mail,
	}
}

type UserResponse struct {
	Name string `json:"name"`
	Id   int64  `json:"id"`
	Mail string `json:"mail"`
}
