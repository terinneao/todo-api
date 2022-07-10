package users

import (
	db "myapp1/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {

	user, err := db.ReadAll()
	if err != nil {
		return c.String(http.StatusNotFound, "Not found")
	}

	return c.JSON(http.StatusOK, user)
}

//GET(/users/:id)
func GetUser(c echo.Context) error {
	id := c.Param("id")

	userDB, err := db.Read(id)
	if err != nil {
		return c.String(http.StatusNotFound, "Not found")
	}

	user := User{
		Id: userDB.Code,
		// Name:  userDB.Name,
		Email: userDB.Email,
		Todo1: userDB.Todo1,
		Todo2: userDB.Todo2,
		Todo3: userDB.Todo3,
	}
	return c.JSON(http.StatusOK, user)
}

//POST(/users)
func Save(c echo.Context) error {
	user := User{}
	if err := c.Bind(&user); err != nil {
		return err
	}

	userDB := db.UserDB{
		Code: user.Id,
		// Name:  user.Name,
		Email: user.Email,
		Todo1: user.Todo1,
		Todo2: user.Todo2,
		Todo3: user.Todo3,
	}

	if err := db.Create(userDB); err != nil {
		return c.String(http.StatusExpectationFailed, "create fail.")
	}

	return c.JSON(http.StatusCreated, user)
}

//PUT(/users/:id)
func Update(c echo.Context) error {
	id := c.Param("id")

	user := User{}
	if err := c.Bind(&user); err != nil {
		return err
	}

	userDB := db.UserDB{

		Code: user.Id,
		// Name:  user.Name,
		Email: user.Email,
		Todo1: user.Todo1,
		Todo2: user.Todo2,
		Todo3: user.Todo3,
	}

	if err := db.Update(id, userDB); err != nil {
		return c.String(http.StatusExpectationFailed, "create fail.")
	}

	return c.String(http.StatusOK, "OK")
}

//DELETE(/users/:id)
func Delete(c echo.Context) error {
	id := c.Param("id")
	userDB, err := db.Delete(id)
	if err != nil {
		return c.String(http.StatusNotFound, "Not found")
	}
	user := User{
		Id: userDB.Code,
		// Name:  userDB.Name,
		Email: userDB.Email,
	}
	return c.JSON(http.StatusOK, user)

}
