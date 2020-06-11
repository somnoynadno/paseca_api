package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
	"paseca/models"
)

var db *gorm.DB

func init() {
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	conn, err := gorm.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			dbHost, dbPort, username, dbName, password))
	if err != nil {
		panic(err)
	} else {
		db = conn
	}

	// Migrate the schema
	// Notice: many-to-many first
	db.AutoMigrate(
		models.FamilyDisease{},
		models.BeeBreed{},
		models.BeeDisease{},
		models.BeeFamily{},
		models.BeeFamilyStatus{},
		models.BeeFarm{},
		models.BeeFarmSize{},
		models.BeeFarmType{},
		models.ControlHarvest{},
		models.Hive{},
		models.HiveFormat{},
		models.HiveFrameType{},
		models.HoneyHarvest{},
		models.HoneySale{},
		models.HoneyType{},
		models.News{},
		models.PollenHarvest{},
		models.Reminder{},
		models.SubscriptionStatus{},
		models.SubscriptionType{},
		models.Swarm{},
		models.SwarmStatus{},
		models.User{},
		models.WikiPage{},
		models.WikiSection{},
		)
}

func GetDB() *gorm.DB {
	return db
}
