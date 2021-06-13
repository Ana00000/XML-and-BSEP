module user-service

replace github.com/xml/XML-and-BSEP/XML/Nistagram/user-service => ./

go 1.16

require (
	github.com/antchfx/xpath v1.1.11
	github.com/form3tech-oss/jwt-go v3.2.3+incompatible
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/google/uuid v1.2.0
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/lib/pq v1.10.2
	github.com/mikespook/gorbac v2.1.0+incompatible
	github.com/sirupsen/logrus v1.4.2
	github.com/xml/XML-and-BSEP/XML/Nistagram/user-service v0.0.0-00010101000000-000000000000
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0
	gopkg.in/mail.v2 v2.3.1
	gorm.io/driver/postgres v1.1.0
	gorm.io/gorm v1.21.10
)
