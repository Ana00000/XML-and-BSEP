module content-service

replace github.com/xml/XML-and-BSEP/XML/Nistagram/content-service => ./

go 1.16

require (
	github.com/antchfx/xpath v1.1.11
	github.com/google/uuid v1.2.0
	github.com/gorilla/mux v1.8.0
	github.com/lib/pq v1.10.2
	github.com/rs/cors v1.7.0
	github.com/xml/XML-and-BSEP/XML/Nistagram/content-service v0.0.0-00010101000000-000000000000
	gorm.io/driver/postgres v1.1.0
	gorm.io/gorm v1.21.10
)
