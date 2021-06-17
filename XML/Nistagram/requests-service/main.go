package main

import (
	"fmt"
	_ "github.com/antchfx/xpath"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/handler"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/repository"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/service"
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	_ "os"
	_ "strconv"
)

func initDB() *gorm.DB {
	dsn := initDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.InappropriateContentRequest{}, &model.PostICR{}, &model.StoryICR{},
		&model.CommentICR{}, &model.VerificationRequest{}, &model.AgentRegistrationRequest{}, &model.FollowRequest{})
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

func initInappropriateContentRequestRepo(database *gorm.DB) *repository.InappropriateContentRequestRepository {
	return &repository.InappropriateContentRequestRepository{Database: database}
}

func initPostICRRepo(database *gorm.DB) *repository.PostICRRepository {
	return &repository.PostICRRepository{Database: database}
}

func initStoryICRRepo(database *gorm.DB) *repository.StoryICRRepository {
	return &repository.StoryICRRepository{Database: database}
}

func initCommentICRRepo(database *gorm.DB) *repository.CommentICRRepository {
	return &repository.CommentICRRepository{Database: database}
}

func initVerificationRequestRepo(database *gorm.DB) *repository.VerificationRequestRepository {
	return &repository.VerificationRequestRepository{Database: database}
}

func initAgentRegistrationRequestRepo(database *gorm.DB) *repository.AgentRegistrationRequestRepository {
	return &repository.AgentRegistrationRequestRepository{Database: database}
}

func initFollowRequestRepo(database *gorm.DB) *repository.FollowRequestRepository {
	return &repository.FollowRequestRepository{Database: database}
}

func initInappropriateContentRequestServices(repo *repository.InappropriateContentRequestRepository) *service.InappropriateContentRequestService {
	return &service.InappropriateContentRequestService{Repo: repo}
}

func initPostICRServices(repo *repository.PostICRRepository) *service.PostICRService {
	return &service.PostICRService{Repo: repo}
}

func initStoryICRServices(repo *repository.StoryICRRepository) *service.StoryICRService {
	return &service.StoryICRService{Repo: repo}
}

func initCommentICRServices(repo *repository.CommentICRRepository) *service.CommentICRService {
	return &service.CommentICRService{Repo: repo}
}

func initVerificationRequestServices(repo *repository.VerificationRequestRepository) *service.VerificationRequestService {
	return &service.VerificationRequestService{Repo: repo}
}

func initAgentRegistrationRequestServices(repo *repository.AgentRegistrationRequestRepository) *service.AgentRegistrationRequestService {
	return &service.AgentRegistrationRequestService{Repo: repo}
}

func initFollowRequestServices(repo *repository.FollowRequestRepository) *service.FollowRequestService {
	return &service.FollowRequestService{Repo: repo}
}

func initInappropriateContentRequestHandler(service *service.InappropriateContentRequestService, LogInfo *logrus.Logger, LogError *logrus.Logger, validator *validator.Validate) *handler.InappropriateContentRequestHandler {
	return &handler.InappropriateContentRequestHandler{
		Service:   service,
		LogInfo:   LogInfo,
		LogError:  LogError,
		Validator: validator,
	}
}

func initPostICRHandler(service *service.PostICRService, LogInfo *logrus.Logger, LogError *logrus.Logger, validator *validator.Validate) *handler.PostICRHandler {
	return &handler.PostICRHandler{
		Service:   service,
		LogInfo:   LogInfo,
		LogError:  LogError,
		Validator: validator,
	}
}

func initStoryICRHandler(service *service.StoryICRService, LogInfo *logrus.Logger, LogError *logrus.Logger, validator *validator.Validate) *handler.StoryICRHandler {
	return &handler.StoryICRHandler{
		Service:   service,
		LogInfo:   LogInfo,
		LogError:  LogError,
		Validator: validator,
	}
}

func initCommentICRHandler(service *service.CommentICRService, LogInfo *logrus.Logger, LogError *logrus.Logger, validator *validator.Validate) *handler.CommentICRHandler {
	return &handler.CommentICRHandler{
		Service:   service,
		LogInfo:   LogInfo,
		LogError:  LogError,
		Validator: validator,
	}
}

func initVerificationRequestHandler(service *service.VerificationRequestService, LogInfo *logrus.Logger, LogError *logrus.Logger, validator *validator.Validate) *handler.VerificationRequestHandler {
	return &handler.VerificationRequestHandler{
		Service:   service,
		LogInfo:   LogInfo,
		LogError:  LogError,
		Validator: validator,
	}
}

func initAgentRegistrationRequestHandler(service *service.AgentRegistrationRequestService, LogInfo *logrus.Logger, LogError *logrus.Logger, validator *validator.Validate) *handler.AgentRegistrationRequestHandler {
	return &handler.AgentRegistrationRequestHandler{
		Service:   service,
		LogInfo:   LogInfo,
		LogError:  LogError,
		Validator: validator,
	}
}

func initFollowRequestHandler(service *service.FollowRequestService, LogInfo *logrus.Logger, LogError *logrus.Logger, validator *validator.Validate) *handler.FollowRequestHandler {
	return &handler.FollowRequestHandler{
		Service:   service,
		LogInfo:   LogInfo,
		LogError:  LogError,
		Validator: validator,
	}
}

func handleFunc(inappropriateContentRequestHandler *handler.InappropriateContentRequestHandler, postICRHandler *handler.PostICRHandler,
	storyICRHandler *handler.StoryICRHandler, commentICRHandler *handler.CommentICRHandler, verificationRequestHandler *handler.VerificationRequestHandler,
	agentRegistrationRequestHandler *handler.AgentRegistrationRequestHandler, followRequestHandler *handler.FollowRequestHandler) {

	router := mux.NewRouter().StrictSlash(true)

	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "Authorization", "Access-Control-Allow-Headers"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"https://localhost:8081"}),
		handlers.AllowCredentials(),
	)

	router.HandleFunc("/inappropriateContentRequest", inappropriateContentRequestHandler.CreateInappropriateContentRequest).Methods("POST")
	router.HandleFunc("/postICR", postICRHandler.CreatePostICR).Methods("POST")
	router.HandleFunc("/storyICR", storyICRHandler.CreateStoryICR).Methods("POST")
	router.HandleFunc("/commentICR", commentICRHandler.CreateCommentICR).Methods("POST")
	router.HandleFunc("/agentRegistrationRequestHandler", agentRegistrationRequestHandler.CreateAgentRegistrationRequest).Methods("POST")
	router.HandleFunc("/create_follow_request/", followRequestHandler.CreateFollowRequest).Methods("POST")
	router.HandleFunc("/find_all_pending_requests_for_user", followRequestHandler.FindAllPendingFollowerRequestsForUser).Methods("GET")
	router.HandleFunc("/find_request_by_id", followRequestHandler.FindFollowerRequestById).Methods("GET")
	router.HandleFunc("/find_all_requests_by_user_id/{userID}", followRequestHandler.FindAllFollowerRequestsForUser).Methods("GET")
	router.HandleFunc("/reject_follow_request", followRequestHandler.RejectFollowRequest).Methods("POST")
	router.HandleFunc("/find_request_by_classic_user_and_follower_user_ids/{classicUserID}/{followerUserID}", followRequestHandler.FindFollowRequestByIDsClassicUserAndHisFollower).Methods("GET")
	router.HandleFunc("/accept_follow_request/{requestID}", followRequestHandler.AcceptFollowRequest).Methods("POST")

	router.HandleFunc("/verificationRequest", verificationRequestHandler.CreateVerificationRequest).Methods("POST")
	router.HandleFunc("/accept_verification_request", verificationRequestHandler.AcceptVerificationRequest).Methods("POST")
	router.HandleFunc("/reject_verification_request/{requestID}", verificationRequestHandler.RejectVerificationRequest).Methods("POST")
	router.HandleFunc("/find_verification_request_by_id", verificationRequestHandler.FindVerificationRequestById).Methods("GET")
	router.HandleFunc("/find_all_pending_verification_requests", verificationRequestHandler.FindAllPendingVerificationRequests).Methods("POST")
	router.HandleFunc("/uploadOfficialDocument/", verificationRequestHandler.Upload).Methods("POST")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), cors(router)))
}

func main() {
	logInfo := logrus.New()
	logError := logrus.New()
	validator := validator.New()

	LogInfoFile, err := os.OpenFile(os.Getenv("LOG_URL")+"/logInfoREQUESTS.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("error opening file: %v", err)
	}

	LogErrorFile, err := os.OpenFile(os.Getenv("LOG_URL")+"/logErrorREQUESTS.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("error opening file: %v", err)
	}
	logInfo.Out = LogInfoFile
	logInfo.Formatter = &logrus.JSONFormatter{}
	logError.Out = LogErrorFile
	logError.Formatter = &logrus.JSONFormatter{}

	database := initDB()
	inappropriateContentRequestRepo := initInappropriateContentRequestRepo(database)
	postICRRepo := initPostICRRepo(database)
	storyICRRepo := initStoryICRRepo(database)
	commentICRRepo := initCommentICRRepo(database)
	verificationRequestRepo := initVerificationRequestRepo(database)
	agentRegistrationRequestRepo := initAgentRegistrationRequestRepo(database)
	followRequestRepo := initFollowRequestRepo(database)

	inappropriateContentRequestService := initInappropriateContentRequestServices(inappropriateContentRequestRepo)
	postICRService := initPostICRServices(postICRRepo)
	storyICRService := initStoryICRServices(storyICRRepo)
	commentICRService := initCommentICRServices(commentICRRepo)
	verificationRequestService := initVerificationRequestServices(verificationRequestRepo)
	agentRegistrationRequestService := initAgentRegistrationRequestServices(agentRegistrationRequestRepo)
	followRequestService := initFollowRequestServices(followRequestRepo)

	inappropriateContentRequestHandler := initInappropriateContentRequestHandler(inappropriateContentRequestService, logInfo, logError, validator)
	postICRHandler := initPostICRHandler(postICRService, logInfo, logError, validator)
	storyICRHandler := initStoryICRHandler(storyICRService, logInfo, logError, validator)
	commentICRHandler := initCommentICRHandler(commentICRService, logInfo, logError, validator)
	verificationRequestRHandler := initVerificationRequestHandler(verificationRequestService, logInfo, logError, validator)
	agentRegistrationRequestHandler := initAgentRegistrationRequestHandler(agentRegistrationRequestService, logInfo, logError, validator)
	followRequestHandler := initFollowRequestHandler(followRequestService, logInfo, logError, validator)

	handleFunc(inappropriateContentRequestHandler, postICRHandler, storyICRHandler, commentICRHandler,
		verificationRequestRHandler, agentRegistrationRequestHandler, followRequestHandler)
}
