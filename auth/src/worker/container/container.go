package container

import (
	"auth/src/shared/infrastructure"
	"auth/src/worker/application/services"
	"auth/src/worker/application/usecases"
	"auth/src/worker/infrastructure/databases"
	"auth/src/worker/infrastructure/servers"
)

func ContainerReadiness() *usecases.ReadinessUsecase {
	settings := infrastructure.NewSettings()
	mysqlDb := infrastructure.NewMySQLDatabase(settings.DriverName, settings.DataSourceName)
	mailServer := infrastructure.NewMailServer(
		settings.SmtpHost,
		settings.SmtpUser,
		settings.SmtpPass,
		settings.SmtpPort,
	)
	dbWorkerRepository := databases.NewMySQLWorkerRepository(mysqlDb)
	mailWorkerRepository := servers.NewMailWorkerRepository(mailServer)
	dbWorkerService := services.NewDBWorkerService(dbWorkerRepository)
	mailWorkerService := services.NewMailWorkerService(mailWorkerRepository)
	readinessUsecase := usecases.NewReadinessUsecase(dbWorkerService, mailWorkerService)
	return readinessUsecase
}

func ContainerSignUp() *usecases.SignUpUsecase {
	settings := infrastructure.NewSettings()
	mysqlDb := infrastructure.NewMySQLDatabase(settings.DriverName, settings.DataSourceName)
	mailServer := infrastructure.NewMailServer(
		settings.SmtpHost,
		settings.SmtpUser,
		settings.SmtpPass,
		settings.SmtpPort,
	)
	dbWorkerRepository := databases.NewMySQLWorkerRepository(mysqlDb)
	mailWorkerRepository := servers.NewMailWorkerRepository(mailServer)
	dbWorkerService := services.NewDBWorkerService(dbWorkerRepository)
	mailWorkerService := services.NewMailWorkerService(mailWorkerRepository)
	signupUsecase := usecases.NewSignUpUsecase(dbWorkerService, mailWorkerService, settings)
	return signupUsecase
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
