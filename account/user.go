package account

import "context"


type User struct{
	ID string `json:"id,omitempty"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type Repository interface{
	createUser(ctx *context.Context,user User)(err error)
	getUser(ctx *context.Context,userID string)(string, error)
}
