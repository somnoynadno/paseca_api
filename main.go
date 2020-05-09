package main

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"paseca/controllers/CRUD"
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

	router.HandleFunc("/api/bee_breed", CRUD.BeeBreedQuery).Methods("GET")
	router.HandleFunc("/api/bee_breed", CRUD.BeeBreedCreate).Methods("POST")
	router.HandleFunc("/api/bee_breed/{id}", CRUD.BeeBreedRetrieve).Methods("GET")
	router.HandleFunc("/api/bee_breed/{id}", CRUD.BeeBreedUpdate).Methods("PUT")
	router.HandleFunc("/api/bee_breed/{id}", CRUD.BeeBreedDelete).Methods("DELETE")

	router.HandleFunc("/api/bee_disease", CRUD.BeeDiseaseQuery).Methods("GET")
	router.HandleFunc("/api/bee_disease", CRUD.BeeDiseaseCreate).Methods("POST")
	router.HandleFunc("/api/bee_disease/{id}", CRUD.BeeDiseaseRetrieve).Methods("GET")
	router.HandleFunc("/api/bee_disease/{id}", CRUD.BeeDiseaseUpdate).Methods("PUT")
	router.HandleFunc("/api/bee_disease/{id}", CRUD.BeeDiseaseDelete).Methods("DELETE")

	router.HandleFunc("/api/bee_family", CRUD.BeeFamilyQuery).Methods("GET")
	router.HandleFunc("/api/bee_family", CRUD.BeeFamilyCreate).Methods("POST")
	router.HandleFunc("/api/bee_family/{id}", CRUD.BeeFamilyRetrieve).Methods("GET")
	router.HandleFunc("/api/bee_family/{id}", CRUD.BeeFamilyUpdate).Methods("PUT")
	router.HandleFunc("/api/bee_family/{id}", CRUD.BeeFamilyDelete).Methods("DELETE")

	router.HandleFunc("/api/bee_family_status", CRUD.BeeFamilyStatusQuery).Methods("GET")
	router.HandleFunc("/api/bee_family_status", CRUD.BeeFamilyStatusCreate).Methods("POST")
	router.HandleFunc("/api/bee_family_status/{id}", CRUD.BeeFamilyStatusRetrieve).Methods("GET")
	router.HandleFunc("/api/bee_family_status/{id}", CRUD.BeeFamilyStatusUpdate).Methods("PUT")
	router.HandleFunc("/api/bee_family_status/{id}", CRUD.BeeFamilyStatusDelete).Methods("DELETE")

	router.HandleFunc("/api/bee_farm", CRUD.BeeFarmQuery).Methods("GET")
	router.HandleFunc("/api/bee_farm", CRUD.BeeFarmCreate).Methods("POST")
	router.HandleFunc("/api/bee_farm/{id}", CRUD.BeeFarmRetrieve).Methods("GET")
	router.HandleFunc("/api/bee_farm/{id}", CRUD.BeeFarmUpdate).Methods("PUT")
	router.HandleFunc("/api/bee_farm/{id}", CRUD.BeeFarmDelete).Methods("DELETE")

	router.HandleFunc("/api/bee_farm_size", CRUD.BeeFarmSizeQuery).Methods("GET")
	router.HandleFunc("/api/bee_farm_size", CRUD.BeeFarmSizeCreate).Methods("POST")
	router.HandleFunc("/api/bee_farm_size/{id}", CRUD.BeeFarmSizeRetrieve).Methods("GET")
	router.HandleFunc("/api/bee_farm_size/{id}", CRUD.BeeFarmSizeUpdate).Methods("PUT")
	router.HandleFunc("/api/bee_farm_size/{id}", CRUD.BeeFarmSizeDelete).Methods("DELETE")

	router.HandleFunc("/api/bee_farm_type", CRUD.BeeFarmTypeQuery).Methods("GET")
	router.HandleFunc("/api/bee_farm_type", CRUD.BeeFarmTypeCreate).Methods("POST")
	router.HandleFunc("/api/bee_farm_type/{id}", CRUD.BeeFarmTypeRetrieve).Methods("GET")
	router.HandleFunc("/api/bee_farm_type/{id}", CRUD.BeeFarmTypeUpdate).Methods("PUT")
	router.HandleFunc("/api/bee_farm_type/{id}", CRUD.BeeFarmTypeDelete).Methods("DELETE")

	router.HandleFunc("/api/family_disease", CRUD.FamilyDiseaseQuery).Methods("GET")
	router.HandleFunc("/api/family_disease", CRUD.FamilyDiseaseCreate).Methods("POST")
	router.HandleFunc("/api/family_disease/{id}", CRUD.FamilyDiseaseRetrieve).Methods("GET")
	router.HandleFunc("/api/family_disease/{id}", CRUD.FamilyDiseaseUpdate).Methods("PUT")
	router.HandleFunc("/api/family_disease/{id}", CRUD.FamilyDiseaseDelete).Methods("DELETE")

	router.HandleFunc("/api/hive", CRUD.HiveQuery).Methods("GET")
	router.HandleFunc("/api/hive", CRUD.HiveCreate).Methods("POST")
	router.HandleFunc("/api/hive/{id}", CRUD.HiveRetrieve).Methods("GET")
	router.HandleFunc("/api/hive/{id}", CRUD.HiveUpdate).Methods("PUT")
	router.HandleFunc("/api/hive/{id}", CRUD.HiveDelete).Methods("DELETE")

	router.HandleFunc("/api/hive_format", CRUD.HiveFormatQuery).Methods("GET")
	router.HandleFunc("/api/hive_format", CRUD.HiveFormatCreate).Methods("POST")
	router.HandleFunc("/api/hive_format/{id}", CRUD.HiveFormatRetrieve).Methods("GET")
	router.HandleFunc("/api/hive_format/{id}", CRUD.HiveFormatUpdate).Methods("PUT")
	router.HandleFunc("/api/hive_format/{id}", CRUD.HiveFormatDelete).Methods("DELETE")

	router.HandleFunc("/api/hive_frame_type", CRUD.HiveFrameTypeQuery).Methods("GET")
	router.HandleFunc("/api/hive_frame_type", CRUD.HiveFrameTypeCreate).Methods("POST")
	router.HandleFunc("/api/hive_frame_type/{id}", CRUD.HiveFrameTypeRetrieve).Methods("GET")
	router.HandleFunc("/api/hive_frame_type/{id}", CRUD.HiveFrameTypeUpdate).Methods("PUT")
	router.HandleFunc("/api/hive_frame_type/{id}", CRUD.HiveFrameTypeDelete).Methods("DELETE")

	router.HandleFunc("/api/honey_harvest", CRUD.HoneyHarvestQuery).Methods("GET")
	router.HandleFunc("/api/honey_harvest", CRUD.HoneyHarvestCreate).Methods("POST")
	router.HandleFunc("/api/honey_harvest/{id}", CRUD.HoneyHarvestRetrieve).Methods("GET")
	router.HandleFunc("/api/honey_harvest/{id}", CRUD.HoneyHarvestUpdate).Methods("PUT")
	router.HandleFunc("/api/honey_harvest/{id}", CRUD.HoneyHarvestDelete).Methods("DELETE")

	router.HandleFunc("/api/honey_sale", CRUD.HoneySaleQuery).Methods("GET")
	router.HandleFunc("/api/honey_sale", CRUD.HoneySaleCreate).Methods("POST")
	router.HandleFunc("/api/honey_sale/{id}", CRUD.HoneySaleRetrieve).Methods("GET")
	router.HandleFunc("/api/honey_sale/{id}", CRUD.HoneySaleUpdate).Methods("PUT")
	router.HandleFunc("/api/honey_sale/{id}", CRUD.HoneySaleDelete).Methods("DELETE")

	router.HandleFunc("/api/honey_type", CRUD.HoneyTypeQuery).Methods("GET")
	router.HandleFunc("/api/honey_type", CRUD.HoneyTypeCreate).Methods("POST")
	router.HandleFunc("/api/honey_type/{id}", CRUD.HoneyTypeRetrieve).Methods("GET")
	router.HandleFunc("/api/honey_type/{id}", CRUD.HoneyTypeUpdate).Methods("PUT")
	router.HandleFunc("/api/honey_type/{id}", CRUD.HoneyTypeDelete).Methods("DELETE")

	router.HandleFunc("/api/reminder", CRUD.ReminderQuery).Methods("GET")
	router.HandleFunc("/api/reminder", CRUD.ReminderCreate).Methods("POST")
	router.HandleFunc("/api/reminder/{id}", CRUD.ReminderRetrieve).Methods("GET")
	router.HandleFunc("/api/reminder/{id}", CRUD.ReminderUpdate).Methods("PUT")
	router.HandleFunc("/api/reminder/{id}", CRUD.ReminderDelete).Methods("DELETE")

	router.HandleFunc("/api/subscription_status", CRUD.SubscriptionStatusQuery).Methods("GET")
	router.HandleFunc("/api/subscription_status", CRUD.SubscriptionStatusCreate).Methods("POST")
	router.HandleFunc("/api/subscription_status/{id}", CRUD.SubscriptionStatusRetrieve).Methods("GET")
	router.HandleFunc("/api/subscription_status/{id}", CRUD.SubscriptionStatusUpdate).Methods("PUT")
	router.HandleFunc("/api/subscription_status/{id}", CRUD.SubscriptionStatusDelete).Methods("DELETE")

	router.HandleFunc("/api/subscription_type", CRUD.SubscriptionTypeQuery).Methods("GET")
	router.HandleFunc("/api/subscription_type", CRUD.SubscriptionTypeCreate).Methods("POST")
	router.HandleFunc("/api/subscription_type/{id}", CRUD.SubscriptionTypeRetrieve).Methods("GET")
	router.HandleFunc("/api/subscription_type/{id}", CRUD.SubscriptionTypeUpdate).Methods("PUT")
	router.HandleFunc("/api/subscription_type/{id}", CRUD.SubscriptionTypeDelete).Methods("DELETE")

	router.HandleFunc("/api/user", CRUD.UserQuery).Methods("GET")
	router.HandleFunc("/api/user", CRUD.UserCreate).Methods("POST")
	router.HandleFunc("/api/user/{id}", CRUD.UserRetrieve).Methods("GET")
	router.HandleFunc("/api/user/{id}", CRUD.UserUpdate).Methods("PUT")
	router.HandleFunc("/api/user/{id}", CRUD.UserDelete).Methods("DELETE")

	// middleware usage
	// do NOT modify the order
	router.Use(middleware.LogBody) // log HTTP body, method and URI
	router.Use(middleware.CORS)   // enable CORS headers

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
