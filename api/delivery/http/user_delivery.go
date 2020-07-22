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

	e.GET("/login", handler.Login)
}

type recvLogin struct {
	ID       string `json:"id" form:"id" query:"id"`
	Password string `json:"password" form:"password" query:"password"`
}

func (h *UserHandler) Login(c echo.Context) error {
	r := new(recvLogin)

	err := c.Bind(r)
	if err != nil {
		return err
	}

	ctx := c.Request().Context()

	user, err := h.uu.Login(ctx, r.ID, r.Password)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}
