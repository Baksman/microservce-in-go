package account

import "context"


type service interface{
	createUser(ctx *context.Context,email string, password string)(error string,err error)
	getUser(ctx *context.Context,userID string)(error string,err error)
}
