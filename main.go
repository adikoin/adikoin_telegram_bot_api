package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"telegram_bot_api/config"
	"telegram_bot_api/controller"
	"telegram_bot_api/handler"
	"telegram_bot_api/repository"
	"telegram_bot_api/routes"
	"telegram_bot_api/security"
	"telegram_bot_api/util"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// const (
// 	path_dir = "/home/dexter/go_projects/telegram_bot_api"
// )

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var userController *controller.UserController

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// t := &Template{
	// 	templates: template.Must(template.Must(template.ParseGlob("/home/dexter/go_projects/telegram_bot_api/public/views/*.html")).ParseGlob("/home/dexter/go_projects/telegram_bot_api/public/views/layouts/*.html")),
	// }

	// e.Static("/static", "/home/dexter/go_projects/telegram_bot_api/public/assets")

	// e.Renderer = t

	// err := godotenv.Load(filepath.Join(path_dir, ".env"))

	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// bot_token := os.Getenv("TELEGRAM_APITOKEN")

	e.HTTPErrorHandler = handler.ErrorHandler
	e.Validator = util.NewValidationUtil()
	config.CORSConfig(e)
	security.WebSecurityConfig(e)

	routes.GetUserApiRoutes(e, userController)

	// api := echotron.NewAPI(bot_token)

	// for u := range echotron.PollingUpdates(bot_token) {
	// 	if u.Message.Text == "/start" {
	// 		api.SendMessage("Hello world", u.ChatID(), nil)
	// 	}
	// }

	// e.GET("/", HomePage)
	// e.GET("/mining", MiningPage)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.ServerPort)))
}

// func HomePage(c echo.Context) error {
// 	return c.Render(http.StatusOK, "home", nil)
// }

// func MiningPage(c echo.Context) error {
// 	return c.Render(http.StatusOK, "mining", nil)
// }

func init() {
	mongoConnection, errorMongoConn := config.MongoConnection()

	if errorMongoConn != nil {
		log.Println("Error when connect mongo : ", errorMongoConn.Error())
	}

	userRepository := repository.NewUserRepository(mongoConnection)
	authValidator := security.NewAuthValidator(userRepository)
	userController = controller.NewUserController(userRepository, authValidator)

	// fileRepository := repository.NewFileRepository(mongoConnection)
	// fileController = controller.NewFileController(fileRepository)

	// postRepository := repository.NewPostRepository(mongoConnection)
	// postController = controller.NewPostController(postRepository)

	// ringRepository := repository.NewRingRepository(mongoConnection)
	// ringController = controller.NewRingController(ringRepository)

	// botRepository := repository.NewBotRepository(mongoConnection)
	// botController = controller.NewBotController(botRepository)
}
