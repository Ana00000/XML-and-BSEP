package main

import (
	_ "github.com/antchfx/xpath"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/campaign-service/handler"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/campaign-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/campaign-service/repository"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/campaign-service/service"
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

	db.AutoMigrate(&model.Campaign{}, &model.Advertisement{}, &model.MultiUseCampaign{}, &model.DisposableCampaign{}, &model.CampaignChosenGroup{})
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

	mux := http.NewServeMux()

	mux.HandleFunc("/multi_use_campaign/", handlerMultiUseCampaign.CreateMultiUseCampaign)
	mux.HandleFunc("/disposable_campaign/", handlerDisposableCampaign.CreateDisposableCampaign)
	mux.HandleFunc("/campaign/", handlerCampaign.CreateCampaign)
	mux.HandleFunc("/advertisement/", handlerAdvertisement.CreateAdvertisement)
	mux.HandleFunc("/campaign_chosen_group/", handlerCampaignChosenGroup.CreateCampaignChosenGroup)

	handlerVar := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(":8090", handlerVar))
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