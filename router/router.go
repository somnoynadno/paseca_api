package router

import (
	"github.com/gorilla/mux"
	"paseca/controllers/AdminController"
	"paseca/controllers/AuthController"
	"paseca/controllers/MainController/AnalyticControllers"
	"paseca/controllers/MainController/BeeFarmControllers"
	"paseca/controllers/MainController/ControllersLK"
	"paseca/controllers/MainController/CustomTypesControllers"
	"paseca/controllers/MainController/HarvestControllers"
	"paseca/controllers/MainController/OtherControllers"
	"paseca/controllers/MainController/WikiControllers"
	"paseca/middleware"
	u "paseca/utils"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()

	// Base API router
	api   := router.PathPrefix("/api").Subrouter()

	// Subrouters
	admin   := api.PathPrefix("/admin").Subrouter()
	auth    := api.PathPrefix("/auth").Subrouter()
	lk      := api.PathPrefix("/lk").Subrouter()
	landing := api.PathPrefix("/landing").Subrouter()

	// Handle options
	api.HandleFunc("/{any}", u.HandleOptions).Methods("OPTIONS")
	api.HandleFunc("/{any1}/{any2}", u.HandleOptions).Methods("OPTIONS")
	api.HandleFunc("/{any1}/{any2}/{any3}", u.HandleOptions).Methods("OPTIONS")
	api.HandleFunc("/{any1}/{any2}/{any3}/{any4}", u.HandleOptions).Methods("OPTIONS")

	// ADMIN AdminController
	admin.HandleFunc("/bee_breed", AdminController.BeeBreedQuery).Methods("GET")
	admin.HandleFunc("/bee_breed", AdminController.BeeBreedCreate).Methods("POST")
	admin.HandleFunc("/bee_breed/{id}", AdminController.BeeBreedRetrieve).Methods("GET")
	admin.HandleFunc("/bee_breed/{id}", AdminController.BeeBreedUpdate).Methods("PUT")
	admin.HandleFunc("/bee_breed/{id}", AdminController.BeeBreedDelete).Methods("DELETE")

	admin.HandleFunc("/bee_disease", AdminController.BeeDiseaseQuery).Methods("GET")
	admin.HandleFunc("/bee_disease", AdminController.BeeDiseaseCreate).Methods("POST")
	admin.HandleFunc("/bee_disease/{id}", AdminController.BeeDiseaseRetrieve).Methods("GET")
	admin.HandleFunc("/bee_disease/{id}", AdminController.BeeDiseaseUpdate).Methods("PUT")
	admin.HandleFunc("/bee_disease/{id}", AdminController.BeeDiseaseDelete).Methods("DELETE")

	admin.HandleFunc("/bee_family", AdminController.BeeFamilyQuery).Methods("GET")
	admin.HandleFunc("/bee_family", AdminController.BeeFamilyCreate).Methods("POST")
	admin.HandleFunc("/bee_family/{id}", AdminController.BeeFamilyRetrieve).Methods("GET")
	admin.HandleFunc("/bee_family/{id}", AdminController.BeeFamilyUpdate).Methods("PUT")
	admin.HandleFunc("/bee_family/{id}", AdminController.BeeFamilyDelete).Methods("DELETE")

	admin.HandleFunc("/bee_family_status", AdminController.BeeFamilyStatusQuery).Methods("GET")
	admin.HandleFunc("/bee_family_status", AdminController.BeeFamilyStatusCreate).Methods("POST")
	admin.HandleFunc("/bee_family_status/{id}", AdminController.BeeFamilyStatusRetrieve).Methods("GET")
	admin.HandleFunc("/bee_family_status/{id}", AdminController.BeeFamilyStatusUpdate).Methods("PUT")
	admin.HandleFunc("/bee_family_status/{id}", AdminController.BeeFamilyStatusDelete).Methods("DELETE")

	admin.HandleFunc("/bee_farm", AdminController.BeeFarmQuery).Methods("GET")
	admin.HandleFunc("/bee_farm", AdminController.BeeFarmCreate).Methods("POST")
	admin.HandleFunc("/bee_farm/{id}", AdminController.BeeFarmRetrieve).Methods("GET")
	admin.HandleFunc("/bee_farm/{id}", AdminController.BeeFarmUpdate).Methods("PUT")
	admin.HandleFunc("/bee_farm/{id}", AdminController.BeeFarmDelete).Methods("DELETE")

	admin.HandleFunc("/bee_farm_size", AdminController.BeeFarmSizeQuery).Methods("GET")
	admin.HandleFunc("/bee_farm_size", AdminController.BeeFarmSizeCreate).Methods("POST")
	admin.HandleFunc("/bee_farm_size/{id}", AdminController.BeeFarmSizeRetrieve).Methods("GET")
	admin.HandleFunc("/bee_farm_size/{id}", AdminController.BeeFarmSizeUpdate).Methods("PUT")
	admin.HandleFunc("/bee_farm_size/{id}", AdminController.BeeFarmSizeDelete).Methods("DELETE")

	admin.HandleFunc("/bee_farm_type", AdminController.BeeFarmTypeQuery).Methods("GET")
	admin.HandleFunc("/bee_farm_type", AdminController.BeeFarmTypeCreate).Methods("POST")
	admin.HandleFunc("/bee_farm_type/{id}", AdminController.BeeFarmTypeRetrieve).Methods("GET")
	admin.HandleFunc("/bee_farm_type/{id}", AdminController.BeeFarmTypeUpdate).Methods("PUT")
	admin.HandleFunc("/bee_farm_type/{id}", AdminController.BeeFarmTypeDelete).Methods("DELETE")

	admin.HandleFunc("/control_harvest", AdminController.ControlHarvestQuery).Methods("GET")
	admin.HandleFunc("/control_harvest", AdminController.ControlHarvestCreate).Methods("POST")
	admin.HandleFunc("/control_harvest/{id}", AdminController.ControlHarvestRetrieve).Methods("GET")
	admin.HandleFunc("/control_harvest/{id}", AdminController.ControlHarvestUpdate).Methods("PUT")
	admin.HandleFunc("/control_harvest/{id}", AdminController.ControlHarvestDelete).Methods("DELETE")

	admin.HandleFunc("/family_disease", AdminController.FamilyDiseaseQuery).Methods("GET")
	admin.HandleFunc("/family_disease", AdminController.FamilyDiseaseCreate).Methods("POST")
	admin.HandleFunc("/family_disease/{id}", AdminController.FamilyDiseaseRetrieve).Methods("GET")
	admin.HandleFunc("/family_disease/{id}", AdminController.FamilyDiseaseUpdate).Methods("PUT")
	admin.HandleFunc("/family_disease/{id}", AdminController.FamilyDiseaseDelete).Methods("DELETE")

	admin.HandleFunc("/hive", AdminController.HiveQuery).Methods("GET")
	admin.HandleFunc("/hive", AdminController.HiveCreate).Methods("POST")
	admin.HandleFunc("/hive/{id}", AdminController.HiveRetrieve).Methods("GET")
	admin.HandleFunc("/hive/{id}", AdminController.HiveUpdate).Methods("PUT")
	admin.HandleFunc("/hive/{id}", AdminController.HiveDelete).Methods("DELETE")

	admin.HandleFunc("/hive_format", AdminController.HiveFormatQuery).Methods("GET")
	admin.HandleFunc("/hive_format", AdminController.HiveFormatCreate).Methods("POST")
	admin.HandleFunc("/hive_format/{id}", AdminController.HiveFormatRetrieve).Methods("GET")
	admin.HandleFunc("/hive_format/{id}", AdminController.HiveFormatUpdate).Methods("PUT")
	admin.HandleFunc("/hive_format/{id}", AdminController.HiveFormatDelete).Methods("DELETE")

	admin.HandleFunc("/hive_frame_type", AdminController.HiveFrameTypeQuery).Methods("GET")
	admin.HandleFunc("/hive_frame_type", AdminController.HiveFrameTypeCreate).Methods("POST")
	admin.HandleFunc("/hive_frame_type/{id}", AdminController.HiveFrameTypeRetrieve).Methods("GET")
	admin.HandleFunc("/hive_frame_type/{id}", AdminController.HiveFrameTypeUpdate).Methods("PUT")
	admin.HandleFunc("/hive_frame_type/{id}", AdminController.HiveFrameTypeDelete).Methods("DELETE")

	admin.HandleFunc("/honey_harvest", AdminController.HoneyHarvestQuery).Methods("GET")
	admin.HandleFunc("/honey_harvest", AdminController.HoneyHarvestCreate).Methods("POST")
	admin.HandleFunc("/honey_harvest/{id}", AdminController.HoneyHarvestRetrieve).Methods("GET")
	admin.HandleFunc("/honey_harvest/{id}", AdminController.HoneyHarvestUpdate).Methods("PUT")
	admin.HandleFunc("/honey_harvest/{id}", AdminController.HoneyHarvestDelete).Methods("DELETE")

	admin.HandleFunc("/honey_sale", AdminController.HoneySaleQuery).Methods("GET")
	admin.HandleFunc("/honey_sale", AdminController.HoneySaleCreate).Methods("POST")
	admin.HandleFunc("/honey_sale/{id}", AdminController.HoneySaleRetrieve).Methods("GET")
	admin.HandleFunc("/honey_sale/{id}", AdminController.HoneySaleUpdate).Methods("PUT")
	admin.HandleFunc("/honey_sale/{id}", AdminController.HoneySaleDelete).Methods("DELETE")

	admin.HandleFunc("/honey_type", AdminController.HoneyTypeQuery).Methods("GET")
	admin.HandleFunc("/honey_type", AdminController.HoneyTypeCreate).Methods("POST")
	admin.HandleFunc("/honey_type/{id}", AdminController.HoneyTypeRetrieve).Methods("GET")
	admin.HandleFunc("/honey_type/{id}", AdminController.HoneyTypeUpdate).Methods("PUT")
	admin.HandleFunc("/honey_type/{id}", AdminController.HoneyTypeDelete).Methods("DELETE")

	admin.HandleFunc("/reminder", AdminController.ReminderQuery).Methods("GET")
	admin.HandleFunc("/reminder", AdminController.ReminderCreate).Methods("POST")
	admin.HandleFunc("/reminder/{id}", AdminController.ReminderRetrieve).Methods("GET")
	admin.HandleFunc("/reminder/{id}", AdminController.ReminderUpdate).Methods("PUT")
	admin.HandleFunc("/reminder/{id}", AdminController.ReminderDelete).Methods("DELETE")

	admin.HandleFunc("/news", AdminController.NewsQuery).Methods("GET")
	admin.HandleFunc("/news", AdminController.NewsCreate).Methods("POST")
	admin.HandleFunc("/news/{id}", AdminController.NewsRetrieve).Methods("GET")
	admin.HandleFunc("/news/{id}", AdminController.NewsUpdate).Methods("PUT")
	admin.HandleFunc("/news/{id}", AdminController.NewsDelete).Methods("DELETE")

	admin.HandleFunc("/pollen_harvest", AdminController.PollenHarvestQuery).Methods("GET")
	admin.HandleFunc("/pollen_harvest", AdminController.PollenHarvestCreate).Methods("POST")
	admin.HandleFunc("/pollen_harvest/{id}", AdminController.PollenHarvestRetrieve).Methods("GET")
	admin.HandleFunc("/pollen_harvest/{id}", AdminController.PollenHarvestUpdate).Methods("PUT")
	admin.HandleFunc("/pollen_harvest/{id}", AdminController.PollenHarvestDelete).Methods("DELETE")

	admin.HandleFunc("/subscription_status", AdminController.SubscriptionStatusQuery).Methods("GET")
	admin.HandleFunc("/subscription_status", AdminController.SubscriptionStatusCreate).Methods("POST")
	admin.HandleFunc("/subscription_status/{id}", AdminController.SubscriptionStatusRetrieve).Methods("GET")
	admin.HandleFunc("/subscription_status/{id}", AdminController.SubscriptionStatusUpdate).Methods("PUT")
	admin.HandleFunc("/subscription_status/{id}", AdminController.SubscriptionStatusDelete).Methods("DELETE")

	admin.HandleFunc("/subscription_type", AdminController.SubscriptionTypeQuery).Methods("GET")
	admin.HandleFunc("/subscription_type", AdminController.SubscriptionTypeCreate).Methods("POST")
	admin.HandleFunc("/subscription_type/{id}", AdminController.SubscriptionTypeRetrieve).Methods("GET")
	admin.HandleFunc("/subscription_type/{id}", AdminController.SubscriptionTypeUpdate).Methods("PUT")
	admin.HandleFunc("/subscription_type/{id}", AdminController.SubscriptionTypeDelete).Methods("DELETE")

	admin.HandleFunc("/swarm", AdminController.SwarmQuery).Methods("GET")
	admin.HandleFunc("/swarm", AdminController.SwarmCreate).Methods("POST")
	admin.HandleFunc("/swarm/{id}", AdminController.SwarmRetrieve).Methods("GET")
	admin.HandleFunc("/swarm/{id}", AdminController.SwarmUpdate).Methods("PUT")
	admin.HandleFunc("/swarm/{id}", AdminController.SwarmDelete).Methods("DELETE")

	admin.HandleFunc("/swarm_status", AdminController.SwarmStatusQuery).Methods("GET")
	admin.HandleFunc("/swarm_status", AdminController.SwarmStatusCreate).Methods("POST")
	admin.HandleFunc("/swarm_status/{id}", AdminController.SwarmStatusRetrieve).Methods("GET")
	admin.HandleFunc("/swarm_status/{id}", AdminController.SwarmStatusUpdate).Methods("PUT")
	admin.HandleFunc("/swarm_status/{id}", AdminController.SwarmStatusDelete).Methods("DELETE")

	admin.HandleFunc("/user", AdminController.UserQuery).Methods("GET")
	admin.HandleFunc("/user", AdminController.UserCreate).Methods("POST")
	admin.HandleFunc("/user/{id}", AdminController.UserRetrieve).Methods("GET")
	admin.HandleFunc("/user/{id}", AdminController.UserUpdate).Methods("PUT")
	admin.HandleFunc("/user/{id}", AdminController.UserDelete).Methods("DELETE")

	admin.HandleFunc("/wiki_page", AdminController.WikiPageQuery).Methods("GET")
	admin.HandleFunc("/wiki_page", AdminController.WikiPageCreate).Methods("POST")
	admin.HandleFunc("/wiki_page/{id}", AdminController.WikiPageRetrieve).Methods("GET")
	admin.HandleFunc("/wiki_page/{id}", AdminController.WikiPageUpdate).Methods("PUT")
	admin.HandleFunc("/wiki_page/{id}", AdminController.WikiPageDelete).Methods("DELETE")

	admin.HandleFunc("/wiki_section", AdminController.WikiSectionQuery).Methods("GET")
	admin.HandleFunc("/wiki_section", AdminController.WikiSectionCreate).Methods("POST")
	admin.HandleFunc("/wiki_section/{id}", AdminController.WikiSectionRetrieve).Methods("GET")
	admin.HandleFunc("/wiki_section/{id}", AdminController.WikiSectionUpdate).Methods("PUT")
	admin.HandleFunc("/wiki_section/{id}", AdminController.WikiSectionDelete).Methods("DELETE")

	// AUTH PATHS
	auth.HandleFunc("/login", AuthController.Authenticate).Methods("POST")
	auth.HandleFunc("/register", AuthController.Registration).Methods("POST")

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
	lk.HandleFunc("/bee_family/{id}", BeeFarmControllers.EditBeeFamily).Methods("PUT")

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
	lk.HandleFunc("/honey_harvest_stats", AnalyticControllers.GetHoneyHarvestsGroupByAmount).Methods("GET")

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

	lk.HandleFunc("/subscription_types", ControllersLK.GetSubscriptionTypes).Methods("GET")

	lk.HandleFunc("/swarm", BeeFarmControllers.CreateSwarm).Methods("POST")
	lk.HandleFunc("/swarm/{id}", BeeFarmControllers.DeleteSwarm).Methods("DELETE")
	lk.HandleFunc("/swarm/{id}", BeeFarmControllers.EditSwarm).Methods("PUT")
	lk.HandleFunc("/swarms_by_bee_farm_id/{id}", BeeFarmControllers.GetSwarmsByBeeFarmID).Methods("GET")

	lk.HandleFunc("/swarm_statuses", CustomTypesControllers.GetSwarmStatuses).Methods("GET")
	lk.HandleFunc("/swarm_status", CustomTypesControllers.CreateSwarmStatus).Methods("POST")
	lk.HandleFunc("/swarm_status/{id}", CustomTypesControllers.DeleteSwarmStatusByID).Methods("DELETE")

	lk.HandleFunc("/user", ControllersLK.GetUser).Methods("GET")

	lk.HandleFunc("/wiki_sections", WikiControllers.GetWikiSections).Methods("GET")
	lk.HandleFunc("/wiki_pages_by_section_id/{id}", WikiControllers.GetWikiPagesBySectionID).Methods("GET")
	lk.HandleFunc("/wiki_page/{id}", WikiControllers.GetWikiPageByID).Methods("GET")

	landing.HandleFunc("/get_stats", OtherControllers.GetStatsForLanding).Methods("GET")

	// middleware usage
	// do NOT modify the order
	api.Use(middleware.CORS)    // enable CORS headers
	api.Use(middleware.LogPath) // log HTTP request URI and method

	admin.Use(middleware.LogBody)                // log HTTP body
	admin.Use(middleware.JwtAuthentication)      // check JWT and create context
	admin.Use(middleware.CheckAdminPermissions)  // check permissions for API usage

	lk.Use(middleware.LogBody)                 // log HTTP body
	lk.Use(middleware.JwtAuthentication)       // check JWT and create context
	lk.Use(middleware.CheckPermissionsBySubscription) // check subscription status

	return router
}
