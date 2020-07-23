package main

import (
	"log"

	delivery "github.com/rssh-jp/test-mng/api/delivery/http"
	"github.com/rssh-jp/test-mng/api/domain"
	mysqlRepo "github.com/rssh-jp/test-mng/api/repository/mysql"
	redisRepo "github.com/rssh-jp/test-mng/api/repository/redis"
	usecase "github.com/rssh-jp/test-mng/api/usecase"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	e := newRouter(viper.GetBool("is_mock"))

	log.Fatal(e.Start(viper.GetString("address")))
}

func newRouter(isMock bool) *echo.Echo {
	e := echo.New()

	var ur domain.UserRepository
	var tr domain.TokenRepository

	if isMock {
		log.Println("Mock mode")
		ur = mysqlRepo.NewUserMysqlRepository(nil, mysqlRepo.OptionIsMock())
		tr = redisRepo.NewTokenRedisRepository(redisRepo.OptionIsMock())
	} else {
		ur = mysqlRepo.NewUserMysqlRepository(nil)
		tr = redisRepo.NewTokenRedisRepository()
	}
	uu := usecase.NewUserUsecase(ur, tr)
	delivery.NewUserHandler(e, uu)

	return e
}
