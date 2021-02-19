package controller

import (
	"go-rest/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// UserController Interface of an user controller
type UserController interface {
	Post() echo.HandlerFunc
	Get() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	Put() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type userController struct {
	userUsecase usecase.UserUsecase
}

// NewUserController Constructor of an user controller
func NewUserController(userUsecase usecase.UserUsecase) UserController {
	return &userController{userUsecase: userUsecase}
}

type requestUser struct {
	LastName  string `json:"lastname"`
	FirstName string `json:"firstname"`
}

type responseUser struct {
	ID        int    `json:"id"`
	LastName  string `json:"lastname"`
	FirstName string `json:"firstname"`
}

// Get Controller of getting user
func (userController *userController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi((c.Param("id")))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		foundUser, err := userController.userUsecase.ReadByID(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseUser{
			ID:        foundUser.ID,
			LastName:  foundUser.LastName,
			FirstName: foundUser.FirstName,
		}

		return c.JSON(http.StatusOK, res)
	}
}

// GetAll Controller of getting users
func (userController *userController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		foundUsers, err := userController.userUsecase.ReadAll()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := []responseUser{}
		for _, foundUser := range *foundUsers {
			res = append(res, responseUser{
				ID:        foundUser.ID,
				LastName:  foundUser.LastName,
				FirstName: foundUser.FirstName,
			})
		}

		return c.JSON(http.StatusOK, res)
	}
}

// Post Controller of posting user
func (userController *userController) Post() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req requestUser
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		createdUser, err := userController.userUsecase.Create(req.LastName, req.FirstName)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseUser{
			ID:        createdUser.ID,
			LastName:  createdUser.LastName,
			FirstName: createdUser.FirstName,
		}

		return c.JSON(http.StatusCreated, res)
	}
}

// Put Controller of updating user
func (userController *userController) Put() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		var req requestUser
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		updatedUser, err := userController.userUsecase.Update(id, req.LastName, req.FirstName)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseUser{
			ID:        updatedUser.ID,
			LastName:  updatedUser.LastName,
			FirstName: updatedUser.FirstName,
		}

		return c.JSON(http.StatusOK, res)
	}
}

// Delete Controller of deleting user
func (userController *userController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		err = userController.userUsecase.Delete(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.NoContent(http.StatusNoContent)
	}
}
