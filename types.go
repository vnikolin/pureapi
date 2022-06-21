package pureapi

type Auth struct {
	Token string `json:"api_token,omitempty"`
}

// Declaring the Constant Values
const (
	StatusUnauthorized        = "Unauthorized"
	StatusInternalServerError = "Server Error"
	StatusBadRequest          = "Bad Request"
	StatusConflict            = "Conflict Already Exists"
	StatusNotFound            = "Not Found"
	StatusForbidden           = "Forbidden"
)

//ArrayInfo
type ArrayInfo struct {
	Revision  string `json:"revision"`
	Version   string `json:"version"`
	ArrayName string `json:"array_name"`
	ID        string `json:"id"`
}

//PhoneHomeInfo
type PhoneHomeInfo struct {
	Phonehome string `json:"phonehome"`
}

//NTPServerInfo
type NTPServerInfo struct {
	Ntpserver []string `json:"ntpserver"`
}

// Hardware
type HardwareInfo struct {
	Details     interface{} `json:"details"`
	Identify    string      `json:"identify"`
	Index       int64       `json:"index"`
	Name        string      `json:"name"`
	Slot        int64       `json:"slot"`
	Speed       int64       `json:"speed"`
	Status      string      `json:"status"`
	Temperature int64       `json:"temperature"`
}

//Network
type NetworkInfo struct {
	Address  string   `json:"address"`
	Enabled  bool     `json:"enabled"`
	Gateway  string   `json:"gateway"`
	Hwaddr   string   `json:"hwaddr"`
	Mtu      int64    `json:"mtu"`
	Name     string   `json:"name"`
	Netmask  string   `json:"netmask"`
	Services []string `json:"services"`
	Slaves   []string `json:"slaves"`
	Speed    int64    `json:"speed"`
	Subnet   string   `json:"subnet"`
}

//DNS
type DNSInfo struct {
	Domain      string   `json:"domain"`
	Nameservers []string `json:"nameservers"`
}

//SNMP
type SNMPInfo struct {
	AuthPassphrase    interface{} `json:"auth_passphrase"`
	AuthProtocol      interface{} `json:"auth_protocol"`
	Community         interface{} `json:"community"`
	Host              string      `json:"host"`
	Name              string      `json:"name"`
	PrivacyPassphrase interface{} `json:"privacy_passphrase"`
	PrivacyProtocol   interface{} `json:"privacy_protocol"`
	User              interface{} `json:"user"`
	Version           string      `json:"version"`
}
