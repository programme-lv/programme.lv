package controller

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/KrisjanisP/deikstra/service/scheduler"
	"gorm.io/gorm"

	"github.com/alexedwards/scs/v2"
	"github.com/gorilla/mux"
)

type Controller struct {
	scheduler    *scheduler.Scheduler
	database     *gorm.DB
	router       *mux.Router
	validate     *validator.Validate
	infoLogger   *log.Logger
	errorLogger  *log.Logger
	sessions     *scs.SessionManager
	passwordSalt string
}

func CreateAPIController(scheduler *scheduler.Scheduler, database *gorm.DB, passwSalt string) *Controller {
	router := mux.NewRouter().StrictSlash(true)
	sessions := scs.New()
	sessions.Lifetime = 24 * time.Hour
	sessions.Cookie.Persist = true
	sessions.Cookie.SameSite = http.SameSiteNoneMode
	sessions.Cookie.Secure = true
	sessions.Cookie.Name = "proglv-session"
	router.Use(sessions.LoadAndSave)
	controller := Controller{
		scheduler:    scheduler,
		router:       router,
		database:     database,
		validate:     validator.New(),
		infoLogger:   log.New(os.Stdout, "API INFO ", log.Ldate|log.Ltime),
		errorLogger:  log.New(os.Stderr, "API ERROR ", log.Ldate|log.Ltime|log.Lshortfile),
		sessions:     sessions,
		passwordSalt: passwSalt,
	}
	controller.registerAPIRoutes()
	return &controller
}

func (c *Controller) StartAPIServer(APIPort int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", APIPort))
	if err != nil {
		c.errorLogger.Fatalf("failed to listen: %v", err)
	}
	c.infoLogger.Println("HTTP server listening at", lis.Addr())
	if err := http.Serve(lis, c.router); err != nil {
		c.errorLogger.Fatalf("failed to serve: %v", err)
	}
}
