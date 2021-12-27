module poiyo-be

// +heroku goVersion go1.15
go 1.15

require (
	cloud.google.com/go/firestore v1.6.1 // indirect
	cloud.google.com/go/storage v1.18.2 // indirect
	firebase.google.com/go v3.13.0+incompatible
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/go-playground/validator/v10 v10.9.0
	github.com/joho/godotenv v1.4.0
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/gommon v0.3.1
	google.golang.org/api v0.59.0
	gorm.io/driver/postgres v1.2.3
	gorm.io/gorm v1.22.3
)
