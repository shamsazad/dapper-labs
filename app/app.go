package app

import (
	"dapper-labs/dao"
	"dapper-labs/handlers"
	"dapper-labs/middleware"
	"dapper-labs/models"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

type App struct {
	Router *mux.Router
	DAO    dao.DaoInterface
}

var app *App

func NewApp() *App {
	db := createDbConnection()
	runMigration(db)
	return &App{
		Router: mux.NewRouter().StrictSlash(true),
		DAO: &dao.Dao{
			DB: db,
		},
	}
}

func runMigration(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{}, &models.UserCredential{})
	if err != nil {
		errorMsg := fmt.Sprintf("could not run the migration %v", err)
		panic(errorMsg)
	}
}

func createDbConnection() *gorm.DB {
	dsn := createDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		errorMsg := fmt.Sprintf("could not connect to the database %v", err)
		panic(errorMsg)
	}
	return db
}

func createDSN() string {

	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatalf("Error loading .env file %v", err)
	//}

	host := getEnv("DB_HOST", "localhost")
	user := getEnv("DB_USER", "shamsazad")
	password := getEnv("DB_PASSWORD", "postgres")
	dbname := getEnv("DB_NAME", "postgres")
	port := getEnv("DB_PORT", "5432")

	return fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", user, password, host, port, dbname)
	//return fmt.Sprintf("postgres://test_dapper_user:123@database:5432/dapper_lab?sslmode=disable")
}

func getEnv(key string, defaultValue string) string {
	configValue := os.Getenv(key)
	if configValue == "" {
		return defaultValue
	} else {
		return configValue
	}
}

func (a *App) HandleRequests() {
	if app == nil {
		app = NewApp()
	}
	log.Println("inside app")
	authenticatedRouter := app.Router.PathPrefix("/auth").Subrouter()
	authenticatedRouter.Use(middleware.AuthMiddleware())
	authenticatedRouter.HandleFunc("/dapper-lab/update/user", handlers.UpdateUser(app.DAO)).Methods(http.MethodPost)
	authenticatedRouter.HandleFunc("/dapper-lab/users", handlers.GetAllUsers(app.DAO)).Methods(http.MethodGet)
	app.Router.HandleFunc("/dapper-lab/user", handlers.SignUp(app.DAO)).Methods(http.MethodPost)
	app.Router.HandleFunc("/dapper-lab/login", handlers.Login(app.DAO)).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":10000", app.Router))
}
