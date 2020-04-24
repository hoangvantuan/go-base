package container

import (
	"log"

	"github.com/hoangvantuan/go-base/domain/repo"
	"github.com/hoangvantuan/go-base/domain/service"
	"github.com/hoangvantuan/go-base/infra"
	"github.com/hoangvantuan/go-base/usercase"
	"github.com/sarulabs/di"
)

var dependences = []di.Def{
	{
		Name: "user-repository",
		Build: func(ctn di.Container)(interface{}, error) {
			return repo.NewUserRepository(infra.DB), nil
		},
	}, 
	{
		Name: "user-service",
		Build: func(ctn di.Container) (interface{}, error) {
			return service.NewUserService(ctn.Get("user-repository").(repo.UserRepository)), nil
		},
	},
	{
		Name: "user-usecase",
		Build: func(ctn di.Container) (interface{}, error) {
			return usercase.NewUserUsercase(
				ctn.Get("user-repository").(repo.UserRepository),
				 ctn.Get("user-service").(*service.UserService)), nil
		},
	},
}

func Boot() di.Container {
	builder, err := di.NewBuilder()
	if err != nil {
		log.Panic(err)
	}

	err = builder.Add(dependences...)
	if err != nil {
		log.Panic(err)
	}

	return builder.Build()
}
