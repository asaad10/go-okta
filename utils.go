package okta

import (
	"log"
	"os"
)

var orgName = os.Getenv("OKTA_API_TEST_ORG")
var apiToken = os.Getenv("OKTA_API_TEST_TOKEN")



func GetUserIdOfUser(email string) string {
	client_instance := GetOktaClient(orgName,apiToken)
	resp, _ := client_instance.UserByUsername(email)
	if len(*resp) > 1 {
		log.Fatal(errorString{"Too many accounts returned with that username"})
	}
	return (*resp)[0].ID


}

//TEMPORARY, should hopefully be able to use userID directly.
func GetAppMetaDataForUser(email string) []AppMetaData {
	client_instance := GetOktaClient(orgName,apiToken)
	uID := GetUserIdOfUser(email)
	resp, _ := client_instance.GetAppsForUser(uID)
	app_list := []AppMetaData{}
	for _,app := range *resp{
		if app.Status == "ACTIVE" {
			app_list = append(app_list, AppMetaData{app.Name,app.Label, app.Settings.OauthClient.InitiateLoginURI,
				app.ID, app.Settings.OauthClient.ApplicationType, app.Links.Logo[0].Href})
		}
	}
	return app_list
}
