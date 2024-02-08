package initializers

import (
	"fmt"
	"os"

	"github.com/Edu58/Learn-Go/GinJWTAuth/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:5432/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	Db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	utils.FatalErrorHandler(err)
	DB = Db
}
