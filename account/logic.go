package account

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/gofrs/uuid"
)


type Service struct{
repository Repository
logger log.Logger	
}


func NewService(rep Repository,logger log.Logger) *Service{
return &Service{
	repository:rep,
	logger:logger,
}
}


func(s *Service) CreateUser(ctx *context.Context,email string, password string)(error string,err error){
	logger := log.With(s.logger,"method","CreateUser")

	uuid,_:= uuid.NewV4()
	id := uuid.String()
	user:= User{
		ID:id,
		Email:email,
		Password:password,
	}


	if err:= s.repository.CreateUser(ctx,user); err!= nil{
		level.Error(logger).Log("err",err)
		return "",err
	}
	logger.log("create user",id)
	return "success",nil
}

func(s *Service) GetUser(ctx *context.Context,id string)(error string,err error){
	logger := log.With(s.logger,"method","CreateUser")

	uuid:= uuid.NewV4()
	user:= User{
		ID:id,
		Email:email,
		Password:password,
	}


	if err:= s.repository.CreateUser(ctx,user); err!= nil{
		level.Error(logger).Log("err",err)
		return "",err
	}
	logger.Log("create user",id)
	return "success",nil
}