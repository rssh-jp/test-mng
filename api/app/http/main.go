package main

import (
	"log"

	delivery "github.com/rssh-jp/test-mng/api/delivery/http"
	mysqlRepo "github.com/rssh-jp/test-mng/api/repository/mysql"
	redisRepo "github.com/rssh-jp/test-mng/api/repository/redis"
	usecase "github.com/rssh-jp/test-mng/api/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	ur := mysqlRepo.NewUserMysqlRepository(nil, mysqlRepo.OptionIsMock())
	tr := redisRepo.NewTokenRedisRepository(redisRepo.OptionIsMock())
	uu := usecase.NewUserUsecase(ur, tr)
	delivery.NewUserHandler(e, uu)

	log.Fatal(e.Start(":1323"))
}
