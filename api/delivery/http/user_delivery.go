package http

import (
	"net/http"

	"github.com/rssh-jp/test-mng/api/domain"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	uu domain.UserUsecase
}

func NewUserHandler(e *echo.Echo, uu domain.UserUsecase) {
	handler := &UserHandler{
		uu: uu,
	}

	e.POST("/login", handler.Login)
	e.POST("/users/fetch", handler.UsersFetch)
	e.POST("/users/update", handler.UsersUpdate)
	e.POST("/users/getown", handler.UsersGetOwn)
}

type recvLogin struct {
	ID       string `json:"id" form:"id" query:"id"`
	Password string `json:"password" form:"password" query:"password"`
}

func (h *UserHandler) Login(c echo.Context) error {
	r := new(domain.RecvLogin)

	err := c.Bind(r)
	if err != nil {
		return err
	}

	ctx := c.Request().Context()

	token, err := h.uu.Login(ctx, r.ID, r.Password)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, domain.SendLogin{Message: "OK", Token: token})
}

type recvUsersFetch struct {
	Token string `json:"token" form:"token" query:"token"`
}

func (h *UserHandler) UsersFetch(c echo.Context) error {
	r := new(domain.RecvUsersFetch)

	err := c.Bind(r)
	if err != nil {
		return err
	}

	ctx := c.Request().Context()

	users, err := h.uu.Fetch(ctx, r.Token)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, domain.SendUsersFetch{Message: "OK", Users: users})
}

type recvUsersUpdate struct {
	Token string      `json:"token" form:"token" query:"token"`
	User  domain.User `json:"user" form:"user" query:"user"`
}

func (h *UserHandler) UsersUpdate(c echo.Context) error {
	r := new(domain.RecvUsersUpdate)

	err := c.Bind(r)
	if err != nil {
		return err
	}

	ctx := c.Request().Context()

	err = h.uu.Update(ctx, r.Token, &r.User)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, domain.SendUsersUpdate{Message: "OK"})
}

type recvUsersGetOwn struct {
	Token string `json:"token" form:"token" query:"token"`
}

func (h *UserHandler) UsersGetOwn(c echo.Context) error {
	r := new(domain.RecvUsersGetOwn)

	err := c.Bind(r)
	if err != nil {
		return err
	}

	ctx := c.Request().Context()

	user, err := h.uu.GetOwn(ctx, r.Token)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, domain.SendUsersGetOwn{Message: "OK", User: user})
}
