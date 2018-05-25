package main

import (
	"os"
	"go-okta"
	"fmt"
)

var orgName = os.Getenv("OKTA_API_TEST_ORG")
var apiToken = os.Getenv("OKTA_API_TEST_TOKEN")
var isProductionOKTAORG = false

func main() {
	fmt.Println(okta.GetUserIdOfUser("asaad@splunk.com"))
	fmt.Println(okta.GetAppMetaDataForUser("asaad@splunk.com"))

	}