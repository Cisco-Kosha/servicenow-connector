package app

func (a *App) initializeRoutes() {
	var apiV1 = "/api/v1"

	// specification routes
	a.Router.HandleFunc(apiV1+"/specification/list", a.listConnectorSpecification).Methods("GET", "OPTIONS")
	// a.Router.HandleFunc(apiV1+"/specification/test", a.testConnectorSpecification).Methods("GET", "OPTIONS")

	a.Router.HandleFunc(apiV1+"/table/{tableName}", a.createTableRecords).Methods("POST", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/table/{tableName}/{sysId}", a.retrieveSingleTableRecord).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/table/{tableName}/{sysId}", a.modifyTableRecord).Methods("PUT", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/table/{tableName}/{sysId}", a.deleteTableRecord).Methods("DELETE", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/table/{tableName}", a.retrieveAllTableRecords).Methods("GET", "OPTIONS")
}
