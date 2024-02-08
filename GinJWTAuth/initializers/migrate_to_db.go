package initializers

import "github.com/Edu58/Learn-Go/GinJWTAuth/models"

func MigrateModelsToDB() {
	DB.AutoMigrate(models.User{})
}
