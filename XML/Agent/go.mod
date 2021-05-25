module agent-app

replace github.com/xml/XML-and-BSEP/XML/Agent => ./

go 1.16

require (
	github.com/antchfx/xpath v1.1.11
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/google/uuid v1.2.0
	github.com/gorilla/mux v1.8.0
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/lib/pq v1.10.2
	github.com/xml/XML-and-BSEP/XML/Agent v0.0.0-00010101000000-000000000000
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0
	gorm.io/driver/postgres v1.1.0
	gorm.io/gorm v1.21.10
)
