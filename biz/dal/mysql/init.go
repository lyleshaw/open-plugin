package mysql

import (
	"crypto/tls"
	mysql_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

var DB *gorm.DB

func Init() {
	_ = mysql_.RegisterTLSConfig("tidb", &tls.Config{
		MinVersion: tls.VersionTLS12,
		ServerName: "gateway01.us-east-1.prod.aws.tidbcloud.com",
	})
	var err error
	//err = godotenv.Load(".env")
	//if err != nil {
	//	fmt.Printf("Some error occurred. Err: %s", err)
	//}
	dsn := os.Getenv("DSN")
	if os.Getenv("ENV") == "dev" {
		dsn = os.Getenv("DSN_DEV")
	}

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
}
