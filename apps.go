package okta

type AppList []struct {
	ID    string `json:"id"` // 0oa10kemsjQ2IRTwG2p7
	Links struct {
		AppLinks []struct {
			Href string `json:"href"` // https://splunk-ciam.okta.com/home/oidc_client/0oa10kemsjQ2IRTwG2p7/aln177a159h7Zf52X0g8
			Name string `json:"name"` // oidc_client_link
			Type string `json:"type"` // text/html
		} `json:"appLinks"`
		Deactivate struct {
			Href string `json:"href"` // https://splunk-ciam.okta.com/api/v1/apps/0oa10kemsjQ2IRTwG2p7/lifecycle/deactivate
		} `json:"deactivate"`
		Groups struct {
			Href string `json:"href"` // https://splunk-ciam.okta.com/api/v1/apps/0oa10kemsjQ2IRTwG2p7/groups
		} `json:"groups"`
		Logo []struct {
			Href string `json:"href"` // https://ok6static.oktacdn.com/assets/img/logos/default.6770228fb0dab49a1695ef440a5279bb.png
			Name string `json:"name"` // medium
			Type string `json:"type"` // image/png
		} `json:"logo"`
		Users struct {
			Href string `json:"href"` // https://splunk-ciam.okta.com/api/v1/apps/0oa10kemsjQ2IRTwG2p7/users
		} `json:"users"`
	} `json:"_links"`
	Accessibility struct {
		ErrorRedirectURL interface{} `json:"errorRedirectUrl"` // <nil>
		LoginRedirectURL interface{} `json:"loginRedirectUrl"` // <nil>
		SelfService      bool        `json:"selfService"`      // false
	} `json:"accessibility"`
	Created     string `json:"created"` // 2018-04-18T17:15:47.000Z
	Credentials struct {
		OauthClient struct {
			AutoKeyRotation         bool   `json:"autoKeyRotation"`            // true
			ClientID                string `json:"client_id"`                  // 0oa10kemsjQ2IRTwG2p7
			TokenEndpointAuthMethod string `json:"token_endpoint_auth_method"` // client_secret_basic
		} `json:"oauthClient"`
		Signing struct {
			Kid string `json:"kid"` // sgkyXrUWgIZAfHbITkdcSztsDwVMqSCbRGls22Bp4XY
		} `json:"signing"`
		UserNameTemplate struct {
			Template string `json:"template"` // ${source.login}
			Type     string `json:"type"`     // BUILT_IN
		} `json:"userNameTemplate"`
	} `json:"credentials"`
	Features    []interface{} `json:"features"`    // <nil>
	Label       string        `json:"label"`       // SSC_Native
	LastUpdated string        `json:"lastUpdated"` // 2018-04-18T17:16:19.000Z
	Name        string        `json:"name"`        // oidc_client
	Settings    struct {
		App struct {
		} `json:"app"`
		Notifications struct {
			Vpn struct {
				HelpURL interface{} `json:"helpUrl"` // <nil>
				Message interface{} `json:"message"` // <nil>
				Network struct {
					Connection string `json:"connection"` // DISABLED
				} `json:"network"`
			} `json:"vpn"`
		} `json:"notifications"`
		OauthClient struct {
			ApplicationType  string      `json:"application_type"`   // native
			ClientURI        interface{} `json:"client_uri"`         // <nil>
			GrantTypes       []string    `json:"grant_types"`        // authorization_code
			InitiateLoginURI string      `json:"initiate_login_uri"` // http://localhost:9090
			LogoURI          interface{} `json:"logo_uri"`           // <nil>
			RedirectURIs     []string    `json:"redirect_uris"`      // http://localhost:9090
			ResponseTypes    []string    `json:"response_types"`     // code
		} `json:"oauthClient"`
	} `json:"settings"`
	SignOnMode string `json:"signOnMode"` // OPENID_CONNECT
	Status     string `json:"status"`     // ACTIVE
	Visibility struct {
		AppLinks struct {
			OidcClientLink bool `json:"oidc_client_link"` // true
		} `json:"appLinks"`
		AutoSubmitToolbar bool `json:"autoSubmitToolbar"` // false
		Hide              struct {
			IOS bool `json:"iOS"` // true
			Web bool `json:"web"` // true
		} `json:"hide"`
	} `json:"visibility"`
}

type App struct {
	ID    string `json:"id"` // 0oa10kemsjQ2IRTwG2p7
	Links struct {
		AppLinks []struct {
			Href string `json:"href"` // https://splunk-ciam.okta.com/home/oidc_client/0oa10kemsjQ2IRTwG2p7/aln177a159h7Zf52X0g8
			Name string `json:"name"` // oidc_client_link
			Type string `json:"type"` // text/html
		} `json:"appLinks"`
		Deactivate struct {
			Href string `json:"href"` // https://splunk-ciam.okta.com/api/v1/apps/0oa10kemsjQ2IRTwG2p7/lifecycle/deactivate
		} `json:"deactivate"`
		Groups struct {
			Href string `json:"href"` // https://splunk-ciam.okta.com/api/v1/apps/0oa10kemsjQ2IRTwG2p7/groups
		} `json:"groups"`
		Logo []struct {
			Href string `json:"href"` // https://ok6static.oktacdn.com/assets/img/logos/default.6770228fb0dab49a1695ef440a5279bb.png
			Name string `json:"name"` // medium
			Type string `json:"type"` // image/png
		} `json:"logo"`
		Users struct {
			Href string `json:"href"` // https://splunk-ciam.okta.com/api/v1/apps/0oa10kemsjQ2IRTwG2p7/users
		} `json:"users"`
	} `json:"_links"`
	Accessibility struct {
		ErrorRedirectURL interface{} `json:"errorRedirectUrl"` // <nil>
		LoginRedirectURL interface{} `json:"loginRedirectUrl"` // <nil>
		SelfService      bool        `json:"selfService"`      // false
	} `json:"accessibility"`
	Created     string `json:"created"` // 2018-04-18T17:15:47.000Z
	Credentials struct {
		OauthClient struct {
			AutoKeyRotation         bool   `json:"autoKeyRotation"`            // true
			ClientID                string `json:"client_id"`                  // 0oa10kemsjQ2IRTwG2p7
			TokenEndpointAuthMethod string `json:"token_endpoint_auth_method"` // client_secret_basic
		} `json:"oauthClient"`
		Signing struct {
			Kid string `json:"kid"` // sgkyXrUWgIZAfHbITkdcSztsDwVMqSCbRGls22Bp4XY
		} `json:"signing"`
		UserNameTemplate struct {
			Template string `json:"template"` // ${source.login}
			Type     string `json:"type"`     // BUILT_IN
		} `json:"userNameTemplate"`
	} `json:"credentials"`
	Features    []interface{} `json:"features"`    // <nil>
	Label       string        `json:"label"`       // SSC_Native
	LastUpdated string        `json:"lastUpdated"` // 2018-04-18T17:16:19.000Z
	Name        string        `json:"name"`        // oidc_client
	Settings    struct {
		App struct {
		} `json:"app"`
		Notifications struct {
			Vpn struct {
				HelpURL interface{} `json:"helpUrl"` // <nil>
				Message interface{} `json:"message"` // <nil>
				Network struct {
					Connection string `json:"connection"` // DISABLED
				} `json:"network"`
			} `json:"vpn"`
		} `json:"notifications"`
		OauthClient struct {
			ApplicationType  string      `json:"application_type"`   // native
			ClientURI        interface{} `json:"client_uri"`         // <nil>
			GrantTypes       []string    `json:"grant_types"`        // authorization_code
			InitiateLoginURI string      `json:"initiate_login_uri"` // http://localhost:9090
			LogoURI          interface{} `json:"logo_uri"`           // <nil>
			RedirectURIs     []string    `json:"redirect_uris"`      // http://localhost:9090
			ResponseTypes    []string    `json:"response_types"`     // code
		} `json:"oauthClient"`
	} `json:"settings"`
	SignOnMode string `json:"signOnMode"` // OPENID_CONNECT
	Status     string `json:"status"`     // ACTIVE
	Visibility struct {
		AppLinks struct {
			OidcClientLink bool `json:"oidc_client_link"` // true
		} `json:"appLinks"`
		AutoSubmitToolbar bool `json:"autoSubmitToolbar"` // false
		Hide              struct {
			IOS bool `json:"iOS"` // true
			Web bool `json:"web"` // true
		} `json:"hide"`
	} `json:"visibility"`
}

type AppMetaData struct{
	Name string
	Label string
	Url string
	Id string
	ApplicationType string
	LogoUrl string

}