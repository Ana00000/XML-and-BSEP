module message-service

replace github.com/xml/XML-and-BSEP/XML/Nistagram/message-service => ./

go 1.16

require (
	github.com/google/uuid v1.2.0
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/xml/XML-and-BSEP/XML/Nistagram/message-service v0.0.0-00010101000000-000000000000
	gorm.io/driver/postgres v1.1.0
	gorm.io/gorm v1.21.10
)
