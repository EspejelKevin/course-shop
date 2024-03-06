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
	phoneServer := infrastructure.NewPhoneServer()
	dbWorkerRepository := databases.NewMySQLWorkerRepository(mysqlDb)
	mailWorkerRepository := servers.NewMailWorkerRepository(mailServer)
	phoneWorkerRepository := servers.NewPhoneWorkerRepository(phoneServer)
	dbWorkerService := services.NewDBWorkerService(dbWorkerRepository)
	mailWorkerService := services.NewMailWorkerService(mailWorkerRepository)
	phoneWorkerService := services.NewPhoneWorkerService(phoneWorkerRepository)
	readinessUsecase := usecases.NewReadinessUsecase(dbWorkerService, mailWorkerService, phoneWorkerService)
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
	signUpUsecase := usecases.NewSignUpUsecase(dbWorkerService, mailWorkerService, settings)
	return signUpUsecase
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
	validateTokenUsecase := usecases.NewValidateTokenUsecase(dbWorkerService)
	return validateTokenUsecase
}

func ContainerValidateEmail() *usecases.ValidateEmailUsecase {
	settings := infrastructure.NewSettings()
	mysqlDb := infrastructure.NewMySQLDatabase(settings.DriverName, settings.DataSourceName)
	dbWorkerRepository := databases.NewMySQLWorkerRepository(mysqlDb)
	dbWorkerService := services.NewDBWorkerService(dbWorkerRepository)
	validateEmailUsecase := usecases.NewValidateEmailUsecase(dbWorkerService)
	return validateEmailUsecase
}

func ContainerConfirmPhone() *usecases.ConfirmPhoneUsecase {
	settings := infrastructure.NewSettings()
	mysqlDb := infrastructure.NewMySQLDatabase(settings.DriverName, settings.DataSourceName)
	phoneServer := infrastructure.NewPhoneServer()
	dbWorkerRepository := databases.NewMySQLWorkerRepository(mysqlDb)
	phoneWorkerRepository := servers.NewPhoneWorkerRepository(phoneServer)
	dbWorkerService := services.NewDBWorkerService(dbWorkerRepository)
	phoneWorkerService := services.NewPhoneWorkerService(phoneWorkerRepository)
	confirmPhoneUsecase := usecases.NewConfirmPhoneUsecase(dbWorkerService, phoneWorkerService, settings)
	return confirmPhoneUsecase
}

func ContainerConfirmEmail() *usecases.ConfirmEmailUsecase {
	settings := infrastructure.NewSettings()
	mailServer := infrastructure.NewMailServer(
		settings.SmtpHost,
		settings.SmtpUser,
		settings.SmtpPass,
		settings.SmtpPort,
	)
	mailWorkerRepository := servers.NewMailWorkerRepository(mailServer)
	mailWorkerService := services.NewMailWorkerService(mailWorkerRepository)
	confirmEmailUsecase := usecases.NewConfirmEmailUsecase(mailWorkerService, settings)
	return confirmEmailUsecase
}
