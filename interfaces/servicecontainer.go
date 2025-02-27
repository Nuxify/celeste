/*
|--------------------------------------------------------------------------
| Service Container
|--------------------------------------------------------------------------
|
| This file performs the compiled dependency injection for your middlewares,
| controllers, services, providers, repositories, etc..
|
*/
package interfaces

import (
	"log"
	"os"
	"sync"

	"celeste/infrastructures/database/mysql"
	"celeste/infrastructures/database/mysql/types"
	userRepository "celeste/module/user/infrastructure/repository"
	userService "celeste/module/user/infrastructure/service"
	userREST "celeste/module/user/interfaces/http/rest"
)

// ServiceContainerInterface contains the dependency injected instances
type ServiceContainerInterface interface {
	// // gRPC
	// RegisterUserGRPCCommandController() userGRPC.UserCommandController
	// RegisterUserGRPCQueryController() userGRPC.UserQueryController

	// REST
	RegisterUserRESTCommandController() userREST.UserCommandController
	RegisterUserRESTQueryController() userREST.UserQueryController
}

type kernel struct{}

var (
	m              sync.Mutex
	k              *kernel
	containerOnce  sync.Once
	mysqlDBHandler *mysql.MySQLDBHandler
)

// ================================= gRPC ===================================
// // RegisterUserGRPCCommandController performs dependency injection to the RegisterUserGRPCCommandController
// func (k *kernel) RegisterUserGRPCCommandController() userGRPC.UserCommandController {
// 	service := k.userCommandServiceContainer()

// 	controller := userGRPC.UserCommandController{
// 		UserCommandServiceInterface: service,
// 	}

// 	return controller
// }

// // RegisterUserGRPCQueryController performs dependency injection to the RegisterUserGRPCQueryController
// func (k *kernel) RegisterUserGRPCQueryController() userGRPC.UserQueryController {
// 	service := k.userQueryServiceContainer()

// 	controller := userGRPC.UserQueryController{
// 		UserQueryServiceInterface: service,
// 	}

// 	return controller
// }

// ==========================================================================
// ================================= REST ===================================
// RegisterUserRESTCommandController performs dependency injection to the RegisterUserRESTCommandController
func (k *kernel) RegisterUserRESTCommandController() userREST.UserCommandController {
	service := k.userCommandServiceContainer()

	controller := userREST.UserCommandController{
		UserCommandServiceInterface: service,
	}

	return controller
}

// RegisterUserRESTQueryController performs dependency injection to the RegisterUserRESTQueryController
func (k *kernel) RegisterUserRESTQueryController() userREST.UserQueryController {
	service := k.userQueryServiceContainer()

	controller := userREST.UserQueryController{
		UserQueryServiceInterface: service,
	}

	return controller
}

// ==========================================================================
func (k *kernel) userCommandServiceContainer() *userService.UserCommandService {
	repository := &userRepository.UserCommandRepository{
		MySQLDBHandlerInterface: mysqlDBHandler,
	}

	service := &userService.UserCommandService{
		UserCommandRepositoryInterface: &userRepository.UserCommandRepositoryCircuitBreaker{
			UserCommandRepositoryInterface: repository,
		},
	}

	return service
}

func (k *kernel) userQueryServiceContainer() *userService.UserQueryService {
	repository := &userRepository.UserQueryRepository{
		MySQLDBHandlerInterface: mysqlDBHandler,
	}

	service := &userService.UserQueryService{
		UserQueryRepositoryInterface: &userRepository.UserQueryRepositoryCircuitBreaker{
			UserQueryRepositoryInterface: repository,
		},
	}

	return service
}

func registerHandlers() {
	var err error

	// connect to database
	mysqlDBHandler = &mysql.MySQLDBHandler{}
	err = mysqlDBHandler.Connect(types.ConnectionParams{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBDatabase: os.Getenv("DB_DATABASE"),
		DBUsername: os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("[SERVER] mysql database is not responding: %v", err)
	}
}

// ServiceContainer export instantiated service container once
func ServiceContainer() ServiceContainerInterface {
	m.Lock()
	defer m.Unlock()

	if k == nil {
		containerOnce.Do(func() {
			// register container handlers
			registerHandlers()

			k = &kernel{}
		})
	}

	return k
}
