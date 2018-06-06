package okta

import (
	"log"
)

func GetUserIdOfUser(email string) string {
	clientInstance := GetOktaClient(orgName, apiToken)
	resp, _ := clientInstance.UserByUsername(email)
	if len(*resp) > 1 {
		log.Fatal(errorString{"Too many accounts returned with that username"})
	}
	return (*resp)[0].ID

}

func GetAppMetaDataForUser(email string) []AppMetaData {
	clientInstance := GetOktaClient(orgName, apiToken)
	uID := GetUserIdOfUser(email)
	resp, _ := clientInstance.GetAppsForUser(uID)
	app_list := []AppMetaData{}
	for _,app := range *resp{
		if app.Status == "ACTIVE" {
			app_list = append(app_list, AppMetaData{app.Name,app.Label, tempDB[app.ID],
			 app.Settings.OauthClient.InitiateLoginURI, app.ID,
			 app.Settings.OauthClient.ApplicationType, app.Links.Logo[0].Href})
		}
	}
	return app_list
}
