package main

import (
	"log"

	delivery "github.com/rssh-jp/test-mng/api/delivery/http"
	repo "github.com/rssh-jp/test-mng/api/repository/mysql"
	usecase "github.com/rssh-jp/test-mng/api/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	ur := repo.NewUserMysqlRepository(nil, repo.OptionIsMock())
	uu := usecase.NewUserUsecase(ur)
	delivery.NewUserHandler(e, uu)

	log.Fatal(e.Start(":1323"))
}
