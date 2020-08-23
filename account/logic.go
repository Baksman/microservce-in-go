package account

import (
	"github.com/go-kit/kit/log"
)


type service struct{
repository Repository
logger log.Logger	
}


func NewService(rep Repository,logger log.Logger) *Service{
return &Service{
	repository:rep,
	logger:logger
}
}