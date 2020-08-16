module gitlab.com/pplayground/pet_tracking/user-service

go 1.14

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/buaazp/fasthttprouter v0.1.1
	github.com/go-sql-driver/mysql v1.5.0
	github.com/joho/godotenv v1.3.0
	github.com/valyala/fasthttp v1.15.1
	gitlab.com/pplayground/pet_tracking/main-framework v0.0.0-00010101000000-000000000000
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2
)

replace gitlab.com/pplayground/pet_tracking/main-framework => gitlab.com/pplayground/pet_tracking/main-framework.git v1.0.2
