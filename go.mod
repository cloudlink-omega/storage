module github.com/cloudlink-omega/storage

go 1.24.1

replace github.com/cloudlink-omega/accounts => ../accounts

require (
	github.com/cloudlink-omega/accounts v0.0.0-00010101000000-000000000000
	github.com/gofiber/fiber/v2 v2.52.8
	github.com/oklog/ulid/v2 v2.1.1
	github.com/patrickmn/go-cache v2.1.0+incompatible
	gorm.io/gorm v1.26.1
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	golang.org/x/crypto v0.38.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.25.0 // indirect
)
