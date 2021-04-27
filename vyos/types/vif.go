package vyos

type vif struct {
	VlanID            int    `json:"vlanId"`
	Address           string `json:"address"`
	Description       string `json:"description"`
	Disable           bool   `json:"disable"`
	DisableLinkDetect bool   `json:"disableLinkDetect"`
	Mac               string `json:"mac"`
	Mtu               int    `json:"mtu"`
	Ip                ip     `json:"ip"`
	Ipv6              ipv6   `json:"ip"`
}
