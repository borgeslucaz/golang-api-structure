package user

import(
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
	"github.com/borgeslucaz/golang-api-structure/utils"
	"github.com/borgeslucaz/golang-api-structure/models"
"errors"
)

// Routes for api
func Routes(e *echo.Echo) {
	g := e.Group("/v0/users")
	g.Use(middleware.JWT([]byte("secret")))

	g.GET("/:id", findUserByID)
	g.GET("/", getAllUsers)
	g.POST("/register", CreateUser)
}

// UserData for login and sign in
type UserData struct {
	Password  string `json:"password"`
	Email     string `json:"email"`
}

// CreateUser New User create handler
func CreateUser(c echo.Context) error {
	newUserData := new(UserData)

	if err := c.Bind(newUserData); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewError(err))
	}

	user, err := UserService.FindByEmail(newUserData.Email)
	if err == nil || user != nil {
		return c.JSON(http.StatusBadRequest, utils.NewError(errors.New("User already exists")))
	}

	newUser, err := models.NewUser(newUserData.Email, newUserData.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewError(err))
	}
	
	newUser, err = UserService.Create(newUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewError(err))
	}

	return c.JSON(http.StatusOK,newUser)
}

func findUserByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewError(err))
	}
	user, err := UserService.Find(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.ResourceNotFound("user"))
	}
	return c.JSON(http.StatusOK, user)
}

func getAllUsers(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}