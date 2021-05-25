module post-service

replace github.com/xml/XML-and-BSEP/XML/Nistagram/post-service => ./

go 1.16

require (
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/google/uuid v1.2.0
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/xml/XML-and-BSEP/XML/Nistagram/post-service v0.0.0-00010101000000-000000000000
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gorm.io/driver/postgres v1.1.0
	gorm.io/gorm v1.21.10
)
