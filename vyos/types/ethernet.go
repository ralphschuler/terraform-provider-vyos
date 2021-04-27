package vyos

type Ethernet struct {
	Address            string `json:"address"` // <address | dhcp | dhcpv6>
	Description        string `json:"description"`
	Disable            bool   `json:"disable"`
	DisableFlowControl bool   `json:"disableFlowControl"`
	DisableLinkdetect  bool   `json:"disableLinkDeect"`
	Mac                string `json:"mac"`
	ArpCacheTimeout    int    `json:"arpCacheTimeout"`
	DisableArpFilter   bool   `json:"diableArpFilter"`
	DisableForwarding  bool   `json:"disableForwarding"`
	Enable
	Mtu           int           `json:"mtu"`
	Vrf           int           `json:"vrf"`
	Ip            ip            `json:"ip"`
	Ipv6          ipv6          `json:"ipv6"`
	DhcpOptions   dhcpOptions   `json:"dhcpOptions"`
	Dhcpv6Options dhcpv6Options `json:"dhcpv6Options"`
	Member        string        `json:"string"`
	Mode          string        `json:"mode"` // <802.3ad | active-backup | broadcast | round-robin | transmit-load-balance | adaptive-load-balance | xor-hash>
	MinLinks      int           `json:"minLinks"`
	HashPolicy    string        `json:"hashPolicy`
	Primary       string        `json:"primary"`
	ArpMonitor    arpMonitor    `json:"arpMonitor"`
	Xdp           string        `json:"xdp"`
	Vif           []vif         `json:"vif"`
	Mirror        mirror        `json:"mirror"`
}
