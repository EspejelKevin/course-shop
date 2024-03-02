package container

import (
	"auth/src/shared/infrastructure"
	"auth/src/worker/application/services"
	"auth/src/worker/application/usecases"
	"auth/src/worker/infrastructure/databases"
)

func ContainerReadiness() *usecases.ReadinessUsecase {
	settings := infrastructure.NewSettings()
	mysqlDb := infrastructure.NewMySQLDatabase(settings.DriverName, settings.DataSourceName)
	dbWorkerRepository := databases.NewMySQLWorkerRepository(mysqlDb)
	dbWorkerService := services.NewDBWorkerService(dbWorkerRepository)
	readinessUsecase := usecases.NewReadinessUsecase(dbWorkerService)
	return readinessUsecase
}

func ContainerSignIn() *usecases.SignInUsecase {
	settings := infrastructure.NewSettings()
	mysqlDb := infrastructure.NewMySQLDatabase(settings.DriverName, settings.DataSourceName)
	dbWorkerRepository := databases.NewMySQLWorkerRepository(mysqlDb)
	dbWorkerService := services.NewDBWorkerService(dbWorkerRepository)
	signinUsecase := usecases.NewSignUpUsecase(dbWorkerService)
	return signinUsecase
}

func ContainerLogIn() *usecases.LogInUsecase {
	settings := infrastructure.NewSettings()
	mysqlDb := infrastructure.NewMySQLDatabase(settings.DriverName, settings.DataSourceName)
	dbWorkerRepository := databases.NewMySQLWorkerRepository(mysqlDb)
	dbWorkerService := services.NewDBWorkerService(dbWorkerRepository)
	loginUsecase := usecases.NewLogInUsecase(dbWorkerService)
	return loginUsecase
}

func ContainerValidateToken() *usecases.ValidateTokenUsecase {
	settings := infrastructure.NewSettings()
	mysqlDb := infrastructure.NewMySQLDatabase(settings.DriverName, settings.DataSourceName)
	dbWorkerRepository := databases.NewMySQLWorkerRepository(mysqlDb)
	dbWorkerService := services.NewDBWorkerService(dbWorkerRepository)
	validatetokenusecase := usecases.NewValidateTokenUsecase(dbWorkerService)
	return validatetokenusecase
}
