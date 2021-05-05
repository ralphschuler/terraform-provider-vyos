package vyos

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFirewallNetworkGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallNetworkGroupCreate,
		Read:   resourceFirewallNetworkGroupRead,
		Update: resourceFirewallNetworkGroupUpdate,
		Delete: resourceFirewallNetworkGroupDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceFirewallNetworkGroupCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceFirewallNetworkGroupRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceFirewallNetworkGroupUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceFirewallNetworkGroupDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
