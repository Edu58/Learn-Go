package initializers

import (
	"github.com/Edu58/Learn-Go/GinJWTAuth/utils"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadEnvVariables() {
	err := godotenv.Load()
	utils.FatalErrorHandler(err)
}
