module post-service

replace github.com/xml/XML-and-BSEP/XML/Nistagram/post-service => ./

go 1.16

require (
	github.com/google/uuid v1.2.0
	github.com/gorilla/mux v1.8.0
	github.com/rs/cors v1.7.0
	github.com/xml/XML-and-BSEP/XML/Nistagram/post-service v0.0.0-00010101000000-000000000000
	gorm.io/driver/postgres v1.1.0
	gorm.io/gorm v1.21.10
)
