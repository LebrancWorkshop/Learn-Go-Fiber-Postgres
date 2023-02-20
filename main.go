package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB;

type SqlLogger struct {
	logger.Interface
}

type Student struct {
	gorm.Model
	FirstName string
	LastName string
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc();
	fmt.Printf("%v\n=====================\n", sql);
}

func main() {
	// Open Database via GORM. 
	var err error;
	dsn := "host=localhost user=postgres password=P@ssw0rd dbname=postgres port=5400 sslmode=disable TimeZone=Asia/Bangkok"
	postgres := postgres.Open(dsn);
	db, err := gorm.Open(postgres, &gorm.Config{
		Logger: &SqlLogger{},
		DryRun: true,
	}); 
	if err != nil {
		panic(err);
	}
	// app := fiber.New();
	
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello API.");
	// });

	// app.Listen(":8090"); 
}