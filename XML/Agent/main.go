package main

import (
	"fmt"
	_ "fmt"
	_ "github.com/antchfx/xpath"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/xml/XML-and-BSEP/XML/Agent/handler"
	"github.com/xml/XML-and-BSEP/XML/Agent/model"
	"github.com/xml/XML-and-BSEP/XML/Agent/repository"
	"github.com/xml/XML-and-BSEP/XML/Agent/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	_ "os"
	_ "strconv"
)

func initDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=root dbname=nistagram-db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.AgentUser{}, &model.Product{})
	return db
}

func initAgentUserRepo(database *gorm.DB) *repository.AgentUserRepository {
	return &repository.AgentUserRepository{Database: database}
}

func initAgentUserServices(repo *repository.AgentUserRepository) *service.AgentUserService {
	return &service.AgentUserService{Repo: repo}
}

func initAgentUserHandler(service *service.AgentUserService) *handler.AgentUserHandler {
	return &handler.AgentUserHandler{Service: service}
}

func initProductRepo(database *gorm.DB) *repository.ProductRepository {
	return &repository.ProductRepository{Database: database}
}

func initProductServices(repo *repository.ProductRepository) *service.ProductService {
	return &service.ProductService{Repo: repo}
}

func initProductHandler(service *service.ProductService) *handler.ProductHandler {
	return &handler.ProductHandler{Service: service}
}

func handleFunc(handlerAgentUser *handler.AgentUserHandler, handlerProduct *handler.ProductHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/agent_user/", handlerAgentUser.CreateAgentUser).Methods("POST")
	router.HandleFunc("/product/", handlerProduct.CreateProduct).Methods("POST")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8087"), router))
}

func main() {
	database := initDB()

	repoAgentUser := initAgentUserRepo(database)
	serviceAgentUser := initAgentUserServices(repoAgentUser)
	handlerAgentUser := initAgentUserHandler(serviceAgentUser)

	repoProduct := initProductRepo(database)
	serviceProduct := initProductServices(repoProduct)
	handlerProduct := initProductHandler(serviceProduct)

	handleFunc(handlerAgentUser, handlerProduct)
}
