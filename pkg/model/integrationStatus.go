package main

//structure for integrationStatus//

type IntegrationStatus struct {
	IntegrationStatus_id     int    `json : "i_id"`
	App_url                  string `json : "app_url"`
	Comment                  string `json : "comment"`
	IntegrationStatus_app    string `json : "i_app"`
	IntegrationStatus_cms    int    `json :"i_cms"`
	IntegrationStatus_org    int    `json : "i_org"`
	IntegrationStatus_status int    `json : "i_status"`
}
