package main

import (
	"fmt"
	_ "fmt"
	_ "github.com/antchfx/xpath"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/location-service/handler"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/location-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/location-service/repository"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/location-service/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	_ "os"
	_ "strconv"
)

func initDB() *gorm.DB{
	dsn := initDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.Location{})
	return db
}

func initDSN() string {
	host := "localhost"
	user := "postgres"
	password := "root"
	dbname := "nistagram-db"
	dbport := "5432"
	if os.Getenv("DBHOST") != "" && os.Getenv("USER") != "" && os.Getenv("PASSWORD") != "" &&
		os.Getenv("DBNAME") != "" && os.Getenv("DBPORT") != "" {
		host = os.Getenv("DBHOST")
		user = os.Getenv("USER")
		password = os.Getenv("PASSWORD")
		dbname = os.Getenv("DBNAME")
		dbport = os.Getenv("DBPORT")
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, user, password, dbname, dbport)

	return dsn
}

func initRepo(database *gorm.DB) *repository.LocationRepository{
	return &repository.LocationRepository { Database: database }
}

func initServices(repo *repository.LocationRepository) *service.LocationService{
	return &service.LocationService { Repo: repo }
}

func initHandler(service *service.LocationService) *handler.LocationHandler{
	return &handler.LocationHandler { Service: service }
}

func handleFunc(handler *handler.LocationHandler){
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.CreateLocation)
	mux.HandleFunc("/find_location_by_id", handler.FindByID)
	handlerVar := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(":8083", handlerVar))
}

func main() {
	database := initDB()
	repo := initRepo(database)
	service := initServices(repo)
	handler := initHandler(service)
	handleFunc(handler)
}