package vyos

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"vyos_interface_bond":     resourceInterfaceBond(),
			"vyos_interface_bridge":   resourceInterfaceBridge(),
			"vyos_interface_dummy":    resourceInterfaceDummy(),
			"vyos_interface_ethernet": resourceInterfaceEthernet(),
			"vyos_interface_geneve":   resourceInterfaceGeneve(),
			"vyos_interface_l2tpv3":   resourceInterfaceL2tpv3(),
			"vyos_interface_loopback": resourceInterfaceLoopback(),
			"vyos_interface_macsec":   resourceInterfaceMacsec(),
			"vyos_interface_macvlan":  resourceInterfaceMacvlan(),
			"vyos_interface_pppoe":    resourceInterfacePppoe(),
			"vyos_interface_tunnel":   resourceInterfaceTunnel(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
	}
}
