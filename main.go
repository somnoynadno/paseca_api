package main

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"paseca/controllers/AuthController"
	"paseca/controllers/CRUD"
	"paseca/controllers/ControllerLK"
	"paseca/db"
	"paseca/middleware"
	u "paseca/utils"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	router := mux.NewRouter()

	// Handle options
	router.HandleFunc("/api/{any}", u.HandleOptions).Methods("OPTIONS")
	router.HandleFunc("/api/{any1}/{any2}", u.HandleOptions).Methods("OPTIONS")
	router.HandleFunc("/api/{any1}/{any2}/{any3}", u.HandleOptions).Methods("OPTIONS")

	// ADMIN CRUD
	router.HandleFunc("/api/admin/bee_breed", CRUD.BeeBreedQuery).Methods("GET")
	router.HandleFunc("/api/admin/bee_breed", CRUD.BeeBreedCreate).Methods("POST")
	router.HandleFunc("/api/admin/bee_breed/{id}", CRUD.BeeBreedRetrieve).Methods("GET")
	router.HandleFunc("/api/admin/bee_breed/{id}", CRUD.BeeBreedUpdate).Methods("PUT")
	router.HandleFunc("/api/admin/bee_breed/{id}", CRUD.BeeBreedDelete).Methods("DELETE")

	router.HandleFunc("/api/admin/bee_disease", CRUD.BeeDiseaseQuery).Methods("GET")
	router.HandleFunc("/api/admin/bee_disease", CRUD.BeeDiseaseCreate).Methods("POST")
	router.HandleFunc("/api/admin/bee_disease/{id}", CRUD.BeeDiseaseRetrieve).Methods("GET")
	router.HandleFunc("/api/admin/bee_disease/{id}", CRUD.BeeDiseaseUpdate).Methods("PUT")
	router.HandleFunc("/api/admin/bee_disease/{id}", CRUD.BeeDiseaseDelete).Methods("DELETE")

	router.HandleFunc("/api/admin/bee_family", CRUD.BeeFamilyQuery).Methods("GET")
	router.HandleFunc("/api/admin/bee_family", CRUD.BeeFamilyCreate).Methods("POST")
	router.HandleFunc("/api/admin/bee_family/{id}", CRUD.BeeFamilyRetrieve).Methods("GET")
	router.HandleFunc("/api/admin/bee_family/{id}", CRUD.BeeFamilyUpdate).Methods("PUT")
	router.HandleFunc("/api/admin/bee_family/{id}", CRUD.BeeFamilyDelete).Methods("DELETE")

	router.HandleFunc("/api/admin/bee_family_status", CRUD.BeeFamilyStatusQuery).Methods("GET")
	router.HandleFunc("/api/admin/bee_family_status", CRUD.BeeFamilyStatusCreate).Methods("POST")
	router.HandleFunc("/api/admin/bee_family_status/{id}", CRUD.BeeFamilyStatusRetrieve).Methods("GET")
	router.HandleFunc("/api/admin/bee_family_status/{id}", CRUD.BeeFamilyStatusUpdate).Methods("PUT")
	router.HandleFunc("/api/admin/bee_family_status/{id}", CRUD.BeeFamilyStatusDelete).Methods("DELETE")

	router.HandleFunc("/api/admin/bee_farm", CRUD.BeeFarmQuery).Methods("GET")
	router.HandleFunc("/api/admin/bee_farm", CRUD.BeeFarmCreate).Methods("POST")
	router.HandleFunc("/api/admin/bee_farm/{id}", CRUD.BeeFarmRetrieve).Methods("GET")
	router.HandleFunc("/api/admin/bee_farm/{id}", CRUD.BeeFarmUpdate).Methods("PUT")
	router.HandleFunc("/api/admin/bee_farm/{id}", CRUD.BeeFarmDelete).Methods("DELETE")

	router.HandleFunc("/api/admin/bee_farm_size", CRUD.BeeFarmSizeQuery).Methods("GET")
	router.HandleFunc("/api/admin/bee_farm_size", CRUD.BeeFarmSizeCreate).Methods("POST")
	router.HandleFunc("/api/admin/bee_farm_size/{id}", CRUD.BeeFarmSizeRetrieve).Methods("GET")
	router.HandleFunc("/api/admin/bee_farm_size/{id}", CRUD.BeeFarmSizeUpdate).Methods("PUT")
	router.HandleFunc("/api/admin/bee_farm_size/{id}", CRUD.BeeFarmSizeDelete).Methods("DELETE")

	router.HandleFunc("/api/admin/bee_farm_type", CRUD.BeeFarmTypeQuery).Methods("GET")
	router.HandleFunc("/api/admin/bee_farm_type", CRUD.BeeFarmTypeCreate).Methods("POST")
	router.HandleFunc("/api/admin/bee_farm_type/{id}", CRUD.BeeFarmTypeRetrieve).Methods("GET")
	router.HandleFunc("/api/admin/bee_farm_type/{id}", CRUD.BeeFarmTypeUpdate).Methods("PUT")
	router.HandleFunc("/api/admin/bee_farm_type/{id}", CRUD.BeeFarmTypeDelete).Methods("DELETE")

	router.HandleFunc("/api/admin/control_harvest", CRUD.ControlHarvestQuery).Methods("GET")
	router.HandleFunc("/api/admin/control_harvest", CRUD.ControlHarvestCreate).Methods("POST")
	router.HandleFunc("/api/admin/control_harvest/{id}", CRUD.ControlHarvestRetrieve).Methods("GET")
	router.HandleFunc("/api/admin/control_harvest/{id}", CRUD.ControlHarvestUpdate).Methods("PUT")
	router.HandleFunc("/api/admin/control_harvest/{id}", CRUD.ControlHarvestDelete).Methods("DELETE")

	router.HandleFunc("/api/admin/family_disease", CRUD.FamilyDiseaseQuery).Methods("GET")
	router.HandleFunc("/api/admin/family_disease", CRUD.FamilyDiseaseCreate).Methods("POST")
	router.HandleFunc("/api/admin/family_disease/{id}", CRUD.FamilyDiseaseRetrieve).Methods("GET")
	router.HandleFunc("/api/admin/family_disease/{id}", CRUD.FamilyDiseaseUpdate).Methods("PUT")
	router.HandleFunc("/api/admin/family_disease/{id}", CRUD.FamilyDiseaseDelete).Methods("DELETE")

	router.HandleFunc("/api/admin/hive", CRUD.HiveQuery).Methods("GET")
	router.HandleFunc("/api/admin/hive", CRUD.HiveCreate).Methods("POST")
	router.HandleFunc("/api/admin/hive/{id}", CRUD.HiveRetrieve).Methods("GET")
	router.HandleFunc("/api/admin/hive/{id}", CRUD.HiveUpdate).Methods("PUT")
	router.HandleFunc("/api/admin/hive/{id}", CRUD.HiveDelete).Methods("DELETE")

	router.HandleFunc("/api/admin/hive_format", CRUD.HiveFormatQuery).Methods("GET")
	router.HandleFunc("/api/admin/hive_format", CRUD.HiveFormatCreate).Methods("POST")
	router.HandleFunc("/api/admin/hive_format/{id}", CRUD.HiveFormatRetrieve).Methods("GET")
	router.HandleFunc("/api/admin/hive_format/{id}", CRUD.HiveFormatUpdate).Methods("PUT")
	router.HandleFunc("/api/admin/hive_format/{id}", CRUD.HiveFormatDelete).Methods("DELETE")

	router.HandleFunc("/api/admin/hive_frame_type", CRUD.HiveFrameTypeQuery).Methods("GET")
	router.HandleFunc("/api/admin/hive_frame_type", CRUD.HiveFrameTypeCreate).Methods("POST")
	router.HandleFunc("/api/admin/hive_frame_type/{id}", CRUD.HiveFrameTypeRetrieve).Methods("GET")
	router.HandleFunc("/api/admin/hive_frame_type/{id}", CRUD.HiveFrameTypeUpdate).Methods("PUT")
	router.HandleFunc("/api/admin/hive_frame_type/{id}", CRUD.HiveFrameTypeDelete).Methods("DELETE")

	router.HandleFunc("/api/admin/honey_harvest", CRUD.HoneyHarvestQuery).Methods("GET")
	router.HandleFunc("/api/admin/honey_harvest", CRUD.HoneyHarvestCreate).Methods("POST")
	router.HandleFunc("/api/admin/honey_harvest/{id}", CRUD.HoneyHarvestRetrieve).Methods("GET")
	router.HandleFunc("/api/admin/honey_harvest/{id}", CRUD.HoneyHarvestUpdate).Methods("PUT")
	router.HandleFunc("/api/admin/honey_harvest/{id}", CRUD.HoneyHarvestDelete).Methods("DELETE")

	router.HandleFunc("/api/admin/honey_sale", CRUD.HoneySaleQuery).Methods("GET")
	router.HandleFunc("/api/admin/honey_sale", CRUD.HoneySaleCreate).Methods("POST")
	router.HandleFunc("/api/admin/honey_sale/{id}", CRUD.HoneySaleRetrieve).Methods("GET")
	router.HandleFunc("/api/admin/honey_sale/{id}", CRUD.HoneySaleUpdate).Methods("PUT")
	router.HandleFunc("/api/admin/honey_sale/{id}", CRUD.HoneySaleDelete).Methods("DELETE")

	router.HandleFunc("/api/admin/honey_type", CRUD.HoneyTypeQuery).Methods("GET")
	router.HandleFunc("/api/admin/honey_type", CRUD.HoneyTypeCreate).Methods("POST")
	router.HandleFunc("/api/admin/honey_type/{id}", CRUD.HoneyTypeRetrieve).Methods("GET")
	router.HandleFunc("/api/admin/honey_type/{id}", CRUD.HoneyTypeUpdate).Methods("PUT")
	router.HandleFunc("/api/admin/honey_type/{id}", CRUD.HoneyTypeDelete).Methods("DELETE")

	router.HandleFunc("/api/admin/reminder", CRUD.ReminderQuery).Methods("GET")
	router.HandleFunc("/api/admin/reminder", CRUD.ReminderCreate).Methods("POST")
	router.HandleFunc("/api/admin/reminder/{id}", CRUD.ReminderRetrieve).Methods("GET")
	router.HandleFunc("/api/admin/reminder/{id}", CRUD.ReminderUpdate).Methods("PUT")
	router.HandleFunc("/api/admin/reminder/{id}", CRUD.ReminderDelete).Methods("DELETE")

	router.HandleFunc("/api/admin/news", CRUD.NewsQuery).Methods("GET")
	router.HandleFunc("/api/admin/news", CRUD.NewsCreate).Methods("POST")
	router.HandleFunc("/api/admin/news/{id}", CRUD.NewsRetrieve).Methods("GET")
	router.HandleFunc("/api/admin/news/{id}", CRUD.NewsUpdate).Methods("PUT")
	router.HandleFunc("/api/admin/news/{id}", CRUD.NewsDelete).Methods("DELETE")

	router.HandleFunc("/api/admin/subscription_status", CRUD.SubscriptionStatusQuery).Methods("GET")
	router.HandleFunc("/api/admin/subscription_status", CRUD.SubscriptionStatusCreate).Methods("POST")
	router.HandleFunc("/api/admin/subscription_status/{id}", CRUD.SubscriptionStatusRetrieve).Methods("GET")
	router.HandleFunc("/api/admin/subscription_status/{id}", CRUD.SubscriptionStatusUpdate).Methods("PUT")
	router.HandleFunc("/api/admin/subscription_status/{id}", CRUD.SubscriptionStatusDelete).Methods("DELETE")

	router.HandleFunc("/api/admin/subscription_type", CRUD.SubscriptionTypeQuery).Methods("GET")
	router.HandleFunc("/api/admin/subscription_type", CRUD.SubscriptionTypeCreate).Methods("POST")
	router.HandleFunc("/api/admin/subscription_type/{id}", CRUD.SubscriptionTypeRetrieve).Methods("GET")
	router.HandleFunc("/api/admin/subscription_type/{id}", CRUD.SubscriptionTypeUpdate).Methods("PUT")
	router.HandleFunc("/api/admin/subscription_type/{id}", CRUD.SubscriptionTypeDelete).Methods("DELETE")

	router.HandleFunc("/api/admin/user", CRUD.UserQuery).Methods("GET")
	router.HandleFunc("/api/admin/user", CRUD.UserCreate).Methods("POST")
	router.HandleFunc("/api/admin/user/{id}", CRUD.UserRetrieve).Methods("GET")
	router.HandleFunc("/api/admin/user/{id}", CRUD.UserUpdate).Methods("PUT")
	router.HandleFunc("/api/admin/user/{id}", CRUD.UserDelete).Methods("DELETE")

	// AUTH PATHS
	router.HandleFunc("/api/auth/login", AuthController.Authenticate).Methods("POST")

	// LK PATHS
	router.HandleFunc("/api/lk/news", ControllerLK.GetLastNews).Methods("GET")
	router.HandleFunc("/api/lk/bee_farm_types", ControllerLK.GetBeeFarmTypes).Methods("GET")
	router.HandleFunc("/api/lk/bee_breeds", ControllerLK.GetBeeBreeds).Methods("GET")
	router.HandleFunc("/api/lk/honey_harvests", ControllerLK.GetUsersHoneyHarvests).Methods("GET")
	router.HandleFunc("/api/lk/bee_farm_sizes", ControllerLK.GetBeeFarmSizes).Methods("GET")
	router.HandleFunc("/api/lk/bee_farms", ControllerLK.GetUserBeeFarms).Methods("GET")
	router.HandleFunc("/api/lk/bee_farm/{id}", ControllerLK.GetBeeFarmByID).Methods("GET")
	router.HandleFunc("/api/lk/user", ControllerLK.GetUser).Methods("GET")
	router.HandleFunc("/api/lk/bee_farm", ControllerLK.CreateBeeFarm).Methods("POST")

	// middleware usage
	// do NOT modify the order
	router.Use(middleware.LogBody) // log HTTP body, method and URI
	router.Use(middleware.CORS)   // enable CORS headers
	router.Use(middleware.JwtAuthentication)   // check JWT and create context
	router.Use(middleware.CheckPermissions)   // check permissions for API usage

	// check connection
	con := db.GetDB()
	errors := con.GetErrors()
	if errors != nil && len(errors) > 0 {
		panic(errors[0])
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "6060" // localhost
	}

	log.Info("Listening on: ", port)

	err := http.ListenAndServe(":" + port, router)
	if err != nil {
		log.Panic(err)
	}
}
