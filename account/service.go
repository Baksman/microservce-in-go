package account

import "context"


type service interface{
	CreateUser(ctx *context.Context,email string, password string)(error string,err error)
	GetUser(ctx *context.Context,userID string)(error string,err error)
}
