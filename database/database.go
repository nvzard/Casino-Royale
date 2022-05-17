package database

import (
	"os"

	"github.com/nvzard/casino-royale/model"
	"github.com/nvzard/casino-royale/utils"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

// DB instance
var DB *gorm.DB

// Connect creates the DB connection
func Connect() error {
	var err error
	dsn := getDatabaseCredentials()
	logger := zapgorm2.New(utils.Logger())
	logger.SetAsDefault()
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger})
	if err != nil {
		return err
	}
	zap.S().Info("Connected to database successfully")
	return nil
}

// Prepare sets up the database tables by dropping if configured and then migrating
func Prepare() error {
	if os.Getenv("DROP_TABLES") == "true" {
		err := drop()
		if err != nil {
			return err
		}
	}
	err := migrate()
	if err != nil {
		return err
	}
	zap.S().Info("Prepared database successfully")
	return nil
}

func drop() error {
	err := DB.Migrator().DropTable(&model.Deck{})
	if err != nil {
		return err
	}

	zap.S().Info("Dropped tables successfully")
	return nil
}

func migrate() error {
	err := DB.AutoMigrate(&model.Deck{})
	if err != nil {
		return err
	}

	zap.S().Info("Migrated tables successfully")
	return nil
}

func getDatabaseCredentials() string {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	db := os.Getenv("POSTGRES_DB")
	sslmode := os.Getenv("SSL_MODE")
	return "host=" + host + " user=" + user + " password=" + password + " dbname=" + db + " sslmode=" + sslmode + " TimeZone=Asia/Kolkata"
}
