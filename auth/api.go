package auth

import(
	"github.com/labstack/echo/v4"
	"net/http"
	"github.com/borgeslucaz/golang-api-structure/utils"
	"github.com/borgeslucaz/golang-api-structure/user"
	"errors"
	"time"
	"github.com/dgrijalva/jwt-go"

)

// Routes for api
func Routes(e *echo.Echo) {
	g := e.Group("/auth")
	g.POST("/login", login)
	g.POST("/register", user.CreateUser)
}

func login(c echo.Context) error {
	loginData := new(user.UserData)
	if err := c.Bind(loginData); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewError(err))
	}
	// Throws unauthorized error
	user, err := user.UserService.Login(loginData.Email, loginData.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewError(errors.New("Email or password not match")))
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}