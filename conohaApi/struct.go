package ConohaApi

import "time"

var token string
var tenantId string

//vps server ID list
var ServerIDList []string

var ApiURLs = struct {
	IdentityToken string
	NetWorks      string
	Compute       string
}{
	IdentityToken: "https://identity.tyo1.conoha.io/v2.0/tokens",
	//NetWorks:"https://networking.tyo1.conoha.io/v2.0/networks",
	NetWorks: "http://ip.zhuikan.com/t.php",
	Compute:  "https://compute.tyo1.conoha.io/v2",
}

//'{"auth":{"passwordCredentials":{"username":"ConoHa","password":"paSSword123456#$%"},"tenantId":"487727e3921d44e3bfe7ebb337bf085e"}}'

var FileList = struct {
	Token string
}{
	Token: "token.json",
}

type DataLogin = struct {
	PasswordCredentials DataPasswordCredentials `json:"passwordCredentials"`
	TenantId            string                  `json:"tenantId"`
}

type DataPasswordCredentials = struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ServerInfo struct {
	Name        string            `json:"name"`
	Status      string              `json:"status"`
	IPV4Address map[string]string `json:"ipv_4_address"`
	IPV6Address map[string]string `json:"ipv_6_address"`
	Created     time.Time         `json:"created"`
	Updated     time.Time         `json:"updated"`
}

/**********************************************************************************************************************/
/************************************ conoha api json struct **********************************************************/
/**********************************************************************************************************************/

type RespError struct {
	BadRequest struct {
		Code    int64  `json:"code"`
		Message string `json:"message"`
	} `json:"badRequest"`
}

type RespServerList struct {
	Servers []struct {
		ID    string `json:"id"`
		Links []struct {
			Href string `json:"href"`
			Rel  string `json:"rel"`
		} `json:"links"`
		Name string `json:"name"`
	} `json:"servers"`
}

type RespServerInfoAddress struct {
	OS_EXT_IPS_MAC_macAddr string `json:"OS-EXT-IPS-MAC:mac_addr"`
	OS_EXT_IPS_type        string `json:"OS-EXT-IPS:type"`
	Addr                   string `json:"addr"`
	Version                int64  `json:"version"`
}
