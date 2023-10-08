package gfilter

import (
	"log"
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func dbInit() {
	var err error
	db, err = gorm.Open(sqlite.Open("db::memory:?cache=shared"), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				ParameterizedQueries:      false,       // Don't include params in the SQL log
				Colorful:                  true,        // Disable color
			}),
	})
	if err != nil {
		slog.Error("", err)
		os.Exit(1)
	}
	db.AutoMigrate(&Mymodel{})
}
func Init() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {

		New(c, db.Model(&Mymodel{}))
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8080")
}

func TestXxx(t *testing.T) {
	dbInit()
	db.AutoMigrate(&Mymodel{})
	Init()

}
func TestXxx1(t *testing.T) {
	dbInit()
	New(nil, db.Model(&Mymodel{}))

}

type Mymodel struct {
	C int
}
