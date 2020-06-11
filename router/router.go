package router

import (
	"github.com/gorilla/mux"
	"paseca/controllers/AuthController"
	"paseca/controllers/CRUD"
	"paseca/controllers/ControllerLK/BeeFarmControllers"
	"paseca/controllers/ControllerLK/HarvestControllers"
	"paseca/controllers/ControllerLK/CustomTypesControllers"
	"paseca/controllers/ControllerLK/OtherControllers"
	"paseca/controllers/ControllerLK/WikiControllers"
	"paseca/middleware"
	u "paseca/utils"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()

	// Base API router
	api   := router.PathPrefix("/api").Subrouter()

	// Subrouters
	admin := api.PathPrefix("/admin").Subrouter()
	auth  := api.PathPrefix("/auth").Subrouter()
	lk    := api.PathPrefix("/lk").Subrouter()

	// Handle options
	api.HandleFunc("/{any}", u.HandleOptions).Methods("OPTIONS")
	api.HandleFunc("/{any1}/{any2}", u.HandleOptions).Methods("OPTIONS")
	api.HandleFunc("/{any1}/{any2}/{any3}", u.HandleOptions).Methods("OPTIONS")
	api.HandleFunc("/{any1}/{any2}/{any3}/{any4}", u.HandleOptions).Methods("OPTIONS")

	// ADMIN CRUD
	admin.HandleFunc("/bee_breed", CRUD.BeeBreedQuery).Methods("GET")
	admin.HandleFunc("/bee_breed", CRUD.BeeBreedCreate).Methods("POST")
	admin.HandleFunc("/bee_breed/{id}", CRUD.BeeBreedRetrieve).Methods("GET")
	admin.HandleFunc("/bee_breed/{id}", CRUD.BeeBreedUpdate).Methods("PUT")
	admin.HandleFunc("/bee_breed/{id}", CRUD.BeeBreedDelete).Methods("DELETE")

	admin.HandleFunc("/bee_disease", CRUD.BeeDiseaseQuery).Methods("GET")
	admin.HandleFunc("/bee_disease", CRUD.BeeDiseaseCreate).Methods("POST")
	admin.HandleFunc("/bee_disease/{id}", CRUD.BeeDiseaseRetrieve).Methods("GET")
	admin.HandleFunc("/bee_disease/{id}", CRUD.BeeDiseaseUpdate).Methods("PUT")
	admin.HandleFunc("/bee_disease/{id}", CRUD.BeeDiseaseDelete).Methods("DELETE")

	admin.HandleFunc("/bee_family", CRUD.BeeFamilyQuery).Methods("GET")
	admin.HandleFunc("/bee_family", CRUD.BeeFamilyCreate).Methods("POST")
	admin.HandleFunc("/bee_family/{id}", CRUD.BeeFamilyRetrieve).Methods("GET")
	admin.HandleFunc("/bee_family/{id}", CRUD.BeeFamilyUpdate).Methods("PUT")
	admin.HandleFunc("/bee_family/{id}", CRUD.BeeFamilyDelete).Methods("DELETE")

	admin.HandleFunc("/bee_family_status", CRUD.BeeFamilyStatusQuery).Methods("GET")
	admin.HandleFunc("/bee_family_status", CRUD.BeeFamilyStatusCreate).Methods("POST")
	admin.HandleFunc("/bee_family_status/{id}", CRUD.BeeFamilyStatusRetrieve).Methods("GET")
	admin.HandleFunc("/bee_family_status/{id}", CRUD.BeeFamilyStatusUpdate).Methods("PUT")
	admin.HandleFunc("/bee_family_status/{id}", CRUD.BeeFamilyStatusDelete).Methods("DELETE")

	admin.HandleFunc("/bee_farm", CRUD.BeeFarmQuery).Methods("GET")
	admin.HandleFunc("/bee_farm", CRUD.BeeFarmCreate).Methods("POST")
	admin.HandleFunc("/bee_farm/{id}", CRUD.BeeFarmRetrieve).Methods("GET")
	admin.HandleFunc("/bee_farm/{id}", CRUD.BeeFarmUpdate).Methods("PUT")
	admin.HandleFunc("/bee_farm/{id}", CRUD.BeeFarmDelete).Methods("DELETE")

	admin.HandleFunc("/bee_farm_size", CRUD.BeeFarmSizeQuery).Methods("GET")
	admin.HandleFunc("/bee_farm_size", CRUD.BeeFarmSizeCreate).Methods("POST")
	admin.HandleFunc("/bee_farm_size/{id}", CRUD.BeeFarmSizeRetrieve).Methods("GET")
	admin.HandleFunc("/bee_farm_size/{id}", CRUD.BeeFarmSizeUpdate).Methods("PUT")
	admin.HandleFunc("/bee_farm_size/{id}", CRUD.BeeFarmSizeDelete).Methods("DELETE")

	admin.HandleFunc("/bee_farm_type", CRUD.BeeFarmTypeQuery).Methods("GET")
	admin.HandleFunc("/bee_farm_type", CRUD.BeeFarmTypeCreate).Methods("POST")
	admin.HandleFunc("/bee_farm_type/{id}", CRUD.BeeFarmTypeRetrieve).Methods("GET")
	admin.HandleFunc("/bee_farm_type/{id}", CRUD.BeeFarmTypeUpdate).Methods("PUT")
	admin.HandleFunc("/bee_farm_type/{id}", CRUD.BeeFarmTypeDelete).Methods("DELETE")

	admin.HandleFunc("/control_harvest", CRUD.ControlHarvestQuery).Methods("GET")
	admin.HandleFunc("/control_harvest", CRUD.ControlHarvestCreate).Methods("POST")
	admin.HandleFunc("/control_harvest/{id}", CRUD.ControlHarvestRetrieve).Methods("GET")
	admin.HandleFunc("/control_harvest/{id}", CRUD.ControlHarvestUpdate).Methods("PUT")
	admin.HandleFunc("/control_harvest/{id}", CRUD.ControlHarvestDelete).Methods("DELETE")

	admin.HandleFunc("/family_disease", CRUD.FamilyDiseaseQuery).Methods("GET")
	admin.HandleFunc("/family_disease", CRUD.FamilyDiseaseCreate).Methods("POST")
	admin.HandleFunc("/family_disease/{id}", CRUD.FamilyDiseaseRetrieve).Methods("GET")
	admin.HandleFunc("/family_disease/{id}", CRUD.FamilyDiseaseUpdate).Methods("PUT")
	admin.HandleFunc("/family_disease/{id}", CRUD.FamilyDiseaseDelete).Methods("DELETE")

	admin.HandleFunc("/hive", CRUD.HiveQuery).Methods("GET")
	admin.HandleFunc("/hive", CRUD.HiveCreate).Methods("POST")
	admin.HandleFunc("/hive/{id}", CRUD.HiveRetrieve).Methods("GET")
	admin.HandleFunc("/hive/{id}", CRUD.HiveUpdate).Methods("PUT")
	admin.HandleFunc("/hive/{id}", CRUD.HiveDelete).Methods("DELETE")

	admin.HandleFunc("/hive_format", CRUD.HiveFormatQuery).Methods("GET")
	admin.HandleFunc("/hive_format", CRUD.HiveFormatCreate).Methods("POST")
	admin.HandleFunc("/hive_format/{id}", CRUD.HiveFormatRetrieve).Methods("GET")
	admin.HandleFunc("/hive_format/{id}", CRUD.HiveFormatUpdate).Methods("PUT")
	admin.HandleFunc("/hive_format/{id}", CRUD.HiveFormatDelete).Methods("DELETE")

	admin.HandleFunc("/hive_frame_type", CRUD.HiveFrameTypeQuery).Methods("GET")
	admin.HandleFunc("/hive_frame_type", CRUD.HiveFrameTypeCreate).Methods("POST")
	admin.HandleFunc("/hive_frame_type/{id}", CRUD.HiveFrameTypeRetrieve).Methods("GET")
	admin.HandleFunc("/hive_frame_type/{id}", CRUD.HiveFrameTypeUpdate).Methods("PUT")
	admin.HandleFunc("/hive_frame_type/{id}", CRUD.HiveFrameTypeDelete).Methods("DELETE")

	admin.HandleFunc("/honey_harvest", CRUD.HoneyHarvestQuery).Methods("GET")
	admin.HandleFunc("/honey_harvest", CRUD.HoneyHarvestCreate).Methods("POST")
	admin.HandleFunc("/honey_harvest/{id}", CRUD.HoneyHarvestRetrieve).Methods("GET")
	admin.HandleFunc("/honey_harvest/{id}", CRUD.HoneyHarvestUpdate).Methods("PUT")
	admin.HandleFunc("/honey_harvest/{id}", CRUD.HoneyHarvestDelete).Methods("DELETE")

	admin.HandleFunc("/honey_sale", CRUD.HoneySaleQuery).Methods("GET")
	admin.HandleFunc("/honey_sale", CRUD.HoneySaleCreate).Methods("POST")
	admin.HandleFunc("/honey_sale/{id}", CRUD.HoneySaleRetrieve).Methods("GET")
	admin.HandleFunc("/honey_sale/{id}", CRUD.HoneySaleUpdate).Methods("PUT")
	admin.HandleFunc("/honey_sale/{id}", CRUD.HoneySaleDelete).Methods("DELETE")

	admin.HandleFunc("/honey_type", CRUD.HoneyTypeQuery).Methods("GET")
	admin.HandleFunc("/honey_type", CRUD.HoneyTypeCreate).Methods("POST")
	admin.HandleFunc("/honey_type/{id}", CRUD.HoneyTypeRetrieve).Methods("GET")
	admin.HandleFunc("/honey_type/{id}", CRUD.HoneyTypeUpdate).Methods("PUT")
	admin.HandleFunc("/honey_type/{id}", CRUD.HoneyTypeDelete).Methods("DELETE")

	admin.HandleFunc("/reminder", CRUD.ReminderQuery).Methods("GET")
	admin.HandleFunc("/reminder", CRUD.ReminderCreate).Methods("POST")
	admin.HandleFunc("/reminder/{id}", CRUD.ReminderRetrieve).Methods("GET")
	admin.HandleFunc("/reminder/{id}", CRUD.ReminderUpdate).Methods("PUT")
	admin.HandleFunc("/reminder/{id}", CRUD.ReminderDelete).Methods("DELETE")

	admin.HandleFunc("/news", CRUD.NewsQuery).Methods("GET")
	admin.HandleFunc("/news", CRUD.NewsCreate).Methods("POST")
	admin.HandleFunc("/news/{id}", CRUD.NewsRetrieve).Methods("GET")
	admin.HandleFunc("/news/{id}", CRUD.NewsUpdate).Methods("PUT")
	admin.HandleFunc("/news/{id}", CRUD.NewsDelete).Methods("DELETE")

	admin.HandleFunc("/pollen_harvest", CRUD.PollenHarvestQuery).Methods("GET")
	admin.HandleFunc("/pollen_harvest", CRUD.PollenHarvestCreate).Methods("POST")
	admin.HandleFunc("/pollen_harvest/{id}", CRUD.PollenHarvestRetrieve).Methods("GET")
	admin.HandleFunc("/pollen_harvest/{id}", CRUD.PollenHarvestUpdate).Methods("PUT")
	admin.HandleFunc("/pollen_harvest/{id}", CRUD.PollenHarvestDelete).Methods("DELETE")

	admin.HandleFunc("/subscription_status", CRUD.SubscriptionStatusQuery).Methods("GET")
	admin.HandleFunc("/subscription_status", CRUD.SubscriptionStatusCreate).Methods("POST")
	admin.HandleFunc("/subscription_status/{id}", CRUD.SubscriptionStatusRetrieve).Methods("GET")
	admin.HandleFunc("/subscription_status/{id}", CRUD.SubscriptionStatusUpdate).Methods("PUT")
	admin.HandleFunc("/subscription_status/{id}", CRUD.SubscriptionStatusDelete).Methods("DELETE")

	admin.HandleFunc("/subscription_type", CRUD.SubscriptionTypeQuery).Methods("GET")
	admin.HandleFunc("/subscription_type", CRUD.SubscriptionTypeCreate).Methods("POST")
	admin.HandleFunc("/subscription_type/{id}", CRUD.SubscriptionTypeRetrieve).Methods("GET")
	admin.HandleFunc("/subscription_type/{id}", CRUD.SubscriptionTypeUpdate).Methods("PUT")
	admin.HandleFunc("/subscription_type/{id}", CRUD.SubscriptionTypeDelete).Methods("DELETE")

	admin.HandleFunc("/swarm", CRUD.SwarmQuery).Methods("GET")
	admin.HandleFunc("/swarm", CRUD.SwarmCreate).Methods("POST")
	admin.HandleFunc("/swarm/{id}", CRUD.SwarmRetrieve).Methods("GET")
	admin.HandleFunc("/swarm/{id}", CRUD.SwarmUpdate).Methods("PUT")
	admin.HandleFunc("/swarm/{id}", CRUD.SwarmDelete).Methods("DELETE")

	admin.HandleFunc("/swarm_status", CRUD.SwarmStatusQuery).Methods("GET")
	admin.HandleFunc("/swarm_status", CRUD.SwarmStatusCreate).Methods("POST")
	admin.HandleFunc("/swarm_status/{id}", CRUD.SwarmStatusRetrieve).Methods("GET")
	admin.HandleFunc("/swarm_status/{id}", CRUD.SwarmStatusUpdate).Methods("PUT")
	admin.HandleFunc("/swarm_status/{id}", CRUD.SwarmStatusDelete).Methods("DELETE")

	admin.HandleFunc("/user", CRUD.UserQuery).Methods("GET")
	admin.HandleFunc("/user", CRUD.UserCreate).Methods("POST")
	admin.HandleFunc("/user/{id}", CRUD.UserRetrieve).Methods("GET")
	admin.HandleFunc("/user/{id}", CRUD.UserUpdate).Methods("PUT")
	admin.HandleFunc("/user/{id}", CRUD.UserDelete).Methods("DELETE")

	admin.HandleFunc("/wiki_page", CRUD.WikiPageQuery).Methods("GET")
	admin.HandleFunc("/wiki_page", CRUD.WikiPageCreate).Methods("POST")
	admin.HandleFunc("/wiki_page/{id}", CRUD.WikiPageRetrieve).Methods("GET")
	admin.HandleFunc("/wiki_page/{id}", CRUD.WikiPageUpdate).Methods("PUT")
	admin.HandleFunc("/wiki_page/{id}", CRUD.WikiPageDelete).Methods("DELETE")

	admin.HandleFunc("/wiki_section", CRUD.WikiSectionQuery).Methods("GET")
	admin.HandleFunc("/wiki_section", CRUD.WikiSectionCreate).Methods("POST")
	admin.HandleFunc("/wiki_section/{id}", CRUD.WikiSectionRetrieve).Methods("GET")
	admin.HandleFunc("/wiki_section/{id}", CRUD.WikiSectionUpdate).Methods("PUT")
	admin.HandleFunc("/wiki_section/{id}", CRUD.WikiSectionDelete).Methods("DELETE")

	// AUTH PATHS
	auth.HandleFunc("/login", AuthController.Authenticate).Methods("POST")

	// LK PATHS
	lk.HandleFunc("/bee_breeds", CustomTypesControllers.GetBeeBreeds).Methods("GET")
	lk.HandleFunc("/bee_breed", CustomTypesControllers.CreateBeeBreed).Methods("POST")
	lk.HandleFunc("/bee_breed/{id}", CustomTypesControllers.DeleteBeeBreedByID).Methods("DELETE")

	lk.HandleFunc("/bee_diseases", CustomTypesControllers.GetBeeDiseases).Methods("GET")
	lk.HandleFunc("/bee_disease", CustomTypesControllers.CreateBeeDisease).Methods("POST")
	lk.HandleFunc("/bee_disease/{id}", CustomTypesControllers.DeleteBeeDiseaseByID).Methods("DELETE")

	lk.HandleFunc("/bee_families", BeeFarmControllers.GetUsersBeeFamilies).Methods("GET")
	lk.HandleFunc("/bee_families_without_hives", BeeFarmControllers.GetUsersBeeFamiliesWithoutHives).Methods("GET")
	lk.HandleFunc("/bee_family/{id}", BeeFarmControllers.GetBeeFamilyByID).Methods("GET")
	lk.HandleFunc("/bee_family", BeeFarmControllers.CreateBeeFamily).Methods("POST")
	lk.HandleFunc("/bee_family/{id}", BeeFarmControllers.DeleteBeeFamily).Methods("DELETE")
	lk.HandleFunc("/do_bee_family_inspection/{id}", BeeFarmControllers.DoBeeFamilyInspection).Methods("POST")
	lk.HandleFunc("/bee_families_by_bee_farm_id/{id}", BeeFarmControllers.GetBeeFamiliesByBeeFarmID).Methods("GET")

	lk.HandleFunc("/bee_family_statuses", CustomTypesControllers.GetBeeFamilyStatuses).Methods("GET")
	lk.HandleFunc("/bee_family_statuses", CustomTypesControllers.CreateBeeFamilyStatus).Methods("POST")
	lk.HandleFunc("/bee_family_status/{id}", CustomTypesControllers.DeleteBeeFamilyStatusByID).Methods("DELETE")

	lk.HandleFunc("/bee_farms", BeeFarmControllers.GetUsersBeeFarms).Methods("GET")
	lk.HandleFunc("/bee_farm/{id}", BeeFarmControllers.GetBeeFarmByID).Methods("GET")
	lk.HandleFunc("/bee_farm", BeeFarmControllers.CreateBeeFarm).Methods("POST")
	lk.HandleFunc("/bee_farm/{id}", BeeFarmControllers.DeleteBeeFarm).Methods("DELETE")
	lk.HandleFunc("/bee_farm/{id}", BeeFarmControllers.EditBeeFarm).Methods("PUT")

	lk.HandleFunc("/bee_farm_sizes", CustomTypesControllers.GetBeeFarmSizes).Methods("GET")
	lk.HandleFunc("/bee_farm_size", CustomTypesControllers.CreateBeeFarmSize).Methods("POST")
	lk.HandleFunc("/bee_farm_size/{id}", CustomTypesControllers.DeleteBeeFarmSizeByID).Methods("DELETE")

	lk.HandleFunc("/bee_farm_types", CustomTypesControllers.GetBeeFarmTypes).Methods("GET")
	lk.HandleFunc("/bee_farm_type", CustomTypesControllers.CreateBeeFarmType).Methods("POST")
	lk.HandleFunc("/bee_farm_type/{id}", CustomTypesControllers.DeleteBeeFarmTypeByID).Methods("DELETE")

	lk.HandleFunc("/control_harvests", HarvestControllers.GetUsersControlHarvests).Methods("GET")
	lk.HandleFunc("/control_harvest", HarvestControllers.CreateControlHarvest).Methods("POST")
	lk.HandleFunc("/control_harvest/{id}", HarvestControllers.DeleteControlHarvest).Methods("DELETE")

	lk.HandleFunc("/family_disease", BeeFarmControllers.CreateFamilyDisease).Methods("POST")
	lk.HandleFunc("/family_disease/{id}", BeeFarmControllers.DeleteFamilyDisease).Methods("DELETE")
	lk.HandleFunc("/family_diseases_by_bee_farm_id/{id}", BeeFarmControllers.GetFamilyDiseasesByBeeFarmID).Methods("GET")

	lk.HandleFunc("/hives", BeeFarmControllers.GetUsersFreeHives).Methods("GET")
	lk.HandleFunc("/hive", BeeFarmControllers.CreateHive).Methods("POST")
	lk.HandleFunc("/hive/{id}", BeeFarmControllers.DeleteHive).Methods("DELETE")
	lk.HandleFunc("/hive/set_coords", BeeFarmControllers.SetHiveCoords).Methods("POST")
	lk.HandleFunc("/hive/set_hive_bee_family", BeeFarmControllers.SetHiveBeeFamily).Methods("POST")
	lk.HandleFunc("/hives_by_bee_farm_id/{id}", BeeFarmControllers.GetHivesByBeeFarmID).Methods("GET")

	lk.HandleFunc("/hive_formats", CustomTypesControllers.GetHiveFormats).Methods("GET")
	lk.HandleFunc("/hive_format", CustomTypesControllers.CreateHiveFormat).Methods("POST")
	lk.HandleFunc("/hive_format/{id}", CustomTypesControllers.DeleteHiveFormatByID).Methods("DELETE")

	lk.HandleFunc("/hive_frame_types", CustomTypesControllers.GetHiveFrameTypes).Methods("GET")
	lk.HandleFunc("/hive_frame_type", CustomTypesControllers.CreateHiveFrameType).Methods("POST")
	lk.HandleFunc("/hive_frame_type/{id}", CustomTypesControllers.DeleteHiveFrameTypeByID).Methods("DELETE")

	lk.HandleFunc("/honey_harvests", HarvestControllers.GetUsersHoneyHarvests).Methods("GET")
	lk.HandleFunc("/honey_harvest", HarvestControllers.CreateHoneyHarvest).Methods("POST")
	lk.HandleFunc("/honey_harvest/{id}", HarvestControllers.DeleteHoneyHarvest).Methods("DELETE")

	lk.HandleFunc("/honey_sales", HarvestControllers.GetUsersHoneySales).Methods("GET")
	lk.HandleFunc("/honey_sale", HarvestControllers.CreateHoneySale).Methods("POST")
	lk.HandleFunc("/honey_sale/{id}", HarvestControllers.DeleteHoneySale).Methods("DELETE")

	lk.HandleFunc("/honey_types", CustomTypesControllers.GetHoneyTypes).Methods("GET")
	lk.HandleFunc("/honey_type", CustomTypesControllers.CreateHoneyType).Methods("POST")
	lk.HandleFunc("/honey_type/{id}", CustomTypesControllers.DeleteHoneyTypeByID).Methods("DELETE")

	lk.HandleFunc("/news", OtherControllers.GetLastNews).Methods("GET")

	lk.HandleFunc("/pollen_harvests", HarvestControllers.GetUsersPollenHarvests).Methods("GET")
	lk.HandleFunc("/pollen_harvest", HarvestControllers.CreatePollenHarvest).Methods("POST")
	lk.HandleFunc("/pollen_harvest/{id}", HarvestControllers.DeletePollenHarvest).Methods("DELETE")

	lk.HandleFunc("/reminder", BeeFarmControllers.CreateReminder).Methods("POST")
	lk.HandleFunc("/reminder/{id}", BeeFarmControllers.DeleteReminder).Methods("DELETE")
	lk.HandleFunc("/check_reminder/{id}", BeeFarmControllers.CheckReminder).Methods("POST")
	lk.HandleFunc("/reminders_by_bee_farm_id/{id}", BeeFarmControllers.GetRemindersByBeeFarmID).Methods("GET")

	lk.HandleFunc("/swarm", BeeFarmControllers.CreateSwarm).Methods("POST")
	lk.HandleFunc("/swarm/{id}", BeeFarmControllers.DeleteSwarm).Methods("DELETE")
	lk.HandleFunc("/swarms_by_bee_farm_id/{id}", BeeFarmControllers.GetSwarmsByBeeFarmID).Methods("GET")

	lk.HandleFunc("/swarm_statuses", CustomTypesControllers.GetSwarmStatuses).Methods("GET")
	lk.HandleFunc("/swarm_status", CustomTypesControllers.CreateSwarmStatus).Methods("POST")
	lk.HandleFunc("/swarm_status/{id}", CustomTypesControllers.DeleteSwarmStatusByID).Methods("DELETE")

	lk.HandleFunc("/user", OtherControllers.GetUser).Methods("GET")

	lk.HandleFunc("/wiki_sections", WikiControllers.GetWikiSections).Methods("GET")
	lk.HandleFunc("/wiki_pages_by_section_id/{id}", WikiControllers.GetWikiPagesBySectionID).Methods("GET")
	lk.HandleFunc("/wiki_page/{id}", WikiControllers.GetWikiPageByID).Methods("GET")

	// middleware usage
	// do NOT modify the order
	api.Use(middleware.LogBody) // log HTTP body, method and URI
	api.Use(middleware.CORS)   // enable CORS headers
	admin.Use(middleware.JwtAuthentication)      // check JWT and create context
	admin.Use(middleware.CheckAdminPermissions) // check permissions for API usage
	lk.Use(middleware.JwtAuthentication)       // check JWT and create context

	return router
}
