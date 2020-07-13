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

	migrateSchema()
	addConstraints()
}

func GetDB() *gorm.DB {
	return db
}

func migrateSchema() {
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

func addConstraints() {
	// base models with user
	// delete cascade
	db.Model(models.BeeFarm{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(models.PollenHarvest{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(models.Hive{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(models.Reminder{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(models.BeeFamily{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(models.Swarm{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(models.HoneyHarvest{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(models.HoneySale{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(models.ControlHarvest{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(models.FamilyDisease{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	// base models with custom
	// delete cascade too
	db.Model(models.BeeBreed{}).AddForeignKey("creator_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(models.SwarmStatus{}).AddForeignKey("creator_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(models.BeeFamilyStatus{}).AddForeignKey("creator_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(models.HiveFormat{}).AddForeignKey("creator_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(models.HiveFrameType{}).AddForeignKey("creator_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(models.HoneyType{}).AddForeignKey("creator_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(models.BeeFarmType{}).AddForeignKey("creator_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(models.BeeFarmSize{}).AddForeignKey("creator_id", "users(id)", "CASCADE", "CASCADE")
	// cascade on other models
	db.Model(models.Reminder{}).AddForeignKey("bee_farm_id", "bee_farms(id)", "CASCADE", "CASCADE")
	db.Model(models.Swarm{}).AddForeignKey("bee_family_id", "bee_families(id)", "CASCADE", "CASCADE")
	db.Model(models.FamilyDisease{}).AddForeignKey("bee_family_id", "bee_families(id)", "CASCADE", "CASCADE")
	// important restrictions on user model
	db.Model(models.User{}).AddForeignKey("subscription_type_id", "subscription_types(id)", "RESTRICT", "CASCADE")
	db.Model(models.User{}).AddForeignKey("subscription_status_id", "subscription_statuses(id)", "RESTRICT", "CASCADE")
	// and some other restrictions
	db.Model(models.WikiPage{}).AddForeignKey("wiki_section_id", "wiki_sections(id)", "RESTRICT", "CASCADE")
}