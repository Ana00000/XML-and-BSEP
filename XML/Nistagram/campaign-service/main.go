package main

import (
	"./handler"
	"./model"
	"./repository"
	"./service"
	"fmt"
	_ "fmt"
	_ "github.com/antchfx/xpath"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	_ "os"
	_ "strconv"
)

func initDB() *gorm.DB{
	dsn := "host=localhost user=postgres password=root dbname=nistagram-db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.Campaign{}, &model.Advertisement{}, &model.MultiUseCampaign{}, &model.DisposableCampaign{}, &model.CampaignChosenGroup{})
	return db
}

func initCampaignChosenGroupRepo(database *gorm.DB) *repository.CampaignChosenGroupRepository{
	return &repository.CampaignChosenGroupRepository { Database: database }
}

func initCampaignChosenGroupServices(repo *repository.CampaignChosenGroupRepository) *service.CampaignChosenGroupService{
	return &service.CampaignChosenGroupService { Repo: repo }
}

func initCampaignChosenGroupHandler(service *service.CampaignChosenGroupService) *handler.CampaignChosenGroupHandler{
	return &handler.CampaignChosenGroupHandler { Service: service }
}

func initCampaignRepo(database *gorm.DB) *repository.CampaignRepository{
	return &repository.CampaignRepository { Database: database }
}

func initDisposableCampaignRepo(database *gorm.DB) *repository.DisposableCampaignRepository{
	return &repository.DisposableCampaignRepository { Database: database }
}

func initMultiUseCampaignRepo(database *gorm.DB) *repository.MultiUseCampaignRepository{
	return &repository.MultiUseCampaignRepository { Database: database }
}

func initAdvertisementRepo(database *gorm.DB) *repository.AdvertisementRepository{
	return &repository.AdvertisementRepository { Database: database }
}

func initAdvertisementServices(repo *repository.AdvertisementRepository) *service.AdvertisementService{
	return &service.AdvertisementService { Repo: repo }
}

func initMultiUseCampaignServices(repo *repository.MultiUseCampaignRepository) *service.MultiUseCampaignService{
	return &service.MultiUseCampaignService { Repo: repo }
}

func initDisposableCampaignServices(repo *repository.DisposableCampaignRepository) *service.DisposableCampaignService{
	return &service.DisposableCampaignService { Repo: repo }
}

func initCampaignServices(repo *repository.CampaignRepository) *service.CampaignService{
	return &service.CampaignService { Repo: repo }
}

func initAdvertisementHandler(service *service.AdvertisementService) *handler.AdvertisementHandler{
	return &handler.AdvertisementHandler { Service: service }
}

func initCampaignHandler(service *service.CampaignService) *handler.CampaignHandler{
	return &handler.CampaignHandler { Service: service }
}

func initDisposableCampaignHandler(service *service.DisposableCampaignService) *handler.DisposableCampaignHandler{
	return &handler.DisposableCampaignHandler { Service: service }
}

func initMultiUseCampaignHandler(service *service.MultiUseCampaignService) *handler.MultiUseCampaignHandler{
	return &handler.MultiUseCampaignHandler { Service: service }
}



func handleFunc(handlerMultiUseCampaign *handler.MultiUseCampaignHandler,handlerDisposableCampaign *handler.DisposableCampaignHandler,handlerCampaign *handler.CampaignHandler,handlerAdvertisement *handler.AdvertisementHandler,
	handlerCampaignChosenGroup *handler.CampaignChosenGroupHandler){
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/multi_use_campaign/", handlerMultiUseCampaign.CreateMultiUseCampaign).Methods("POST")
	router.HandleFunc("/disposable_campaign/", handlerDisposableCampaign.CreateDisposableCampaign).Methods("POST")
	router.HandleFunc("/campaign/", handlerCampaign.CreateCampaign).Methods("POST")
	router.HandleFunc("/advertisement/", handlerAdvertisement.CreateAdvertisement).Methods("POST")
	router.HandleFunc("/campaign_chosen_group/", handlerCampaignChosenGroup.CreateCampaignChosenGroup).Methods("POST")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8087"), router))
}

func main() {
	database := initDB()

	repoMultiUseCampaign := initMultiUseCampaignRepo(database)
	repoDisposableCampaign := initDisposableCampaignRepo(database)
	repoCampaign := initCampaignRepo(database)
	repoAdvertisement := initAdvertisementRepo(database)

	serviceMultiUseCampaign := initMultiUseCampaignServices(repoMultiUseCampaign)
	serviceDisposableCampaign := initDisposableCampaignServices(repoDisposableCampaign)
	serviceCampaign := initCampaignServices(repoCampaign)
	serviceAdvertisement := initAdvertisementServices(repoAdvertisement)

	handlerMultiUseCampaign := initMultiUseCampaignHandler(serviceMultiUseCampaign)
	handlerDisposableCampaign := initDisposableCampaignHandler(serviceDisposableCampaign)
	handlerCampaign := initCampaignHandler(serviceCampaign)
	handlerAdvertisement := initAdvertisementHandler(serviceAdvertisement)

	repoCampaignChosenGroup := initCampaignChosenGroupRepo(database)
	serviceCampaignChosenGroup := initCampaignChosenGroupServices(repoCampaignChosenGroup)
	handlerCampaignChosenGroup := initCampaignChosenGroupHandler(serviceCampaignChosenGroup)
	handleFunc(handlerMultiUseCampaign,handlerDisposableCampaign,handlerCampaign,handlerAdvertisement,handlerCampaignChosenGroup)
}