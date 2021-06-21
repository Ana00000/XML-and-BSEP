module requests-service

replace github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service => ./

go 1.16

require (
	github.com/antchfx/xpath v1.1.11
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/google/uuid v1.2.0
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/lib/pq v1.10.2
	github.com/sirupsen/logrus v1.4.2
	github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service v0.0.0-00010101000000-000000000000
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0
	gopkg.in/mail.v2 v2.3.1
	gorm.io/driver/postgres v1.1.0
	gorm.io/gorm v1.21.10
)
