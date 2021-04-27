package vyos

type ipv6 struct {
	ArpCacheTimeout   bool   `json:"arpCacheTimeout"`
	DisableArpFilter  bool   `json:"disableArpFilter"`
	DisableForwarding bool   `json:"disableForwarding"`
	EnableArpAccept   bool   `json:"enableArpAccept"`
	EnableArpIgnore   bool   `json:"enableArpEnable"`
	EnableProxyArp    bool   `json:"enableProxyArp"`
	ProxyArpPVlan     bool   `json:"proxyArpPVlan"`
	SourceValidation  string `json:"sourceValidation"` // <strict | loose | disable>
}
