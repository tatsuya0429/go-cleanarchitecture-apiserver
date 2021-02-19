package main

import (
	"fmt"
	"go-rest/infrastructure/persistence"
	"go-rest/interface/controller"
	"go-rest/usecase"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

func main() {
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		"user",
		"password",
		"db",
		"3306",
		"sampledb",
	)
	conn, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error)
	}

	err = conn.DB().Ping()
	if err != nil {
		panic(err)
	}

	conn.LogMode(true)
	conn.Set("gorm:table_options", "ENGINE=InnoDB")

	userRepository := persistence.NewUserRepository(conn)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)

	e := echo.New()
	controller.InitRouting(e, userController)
	e.Logger.Fatal(e.Start(":8080"))
}
