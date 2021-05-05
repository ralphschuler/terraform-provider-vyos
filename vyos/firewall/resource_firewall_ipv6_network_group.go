package vyos

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFirewallIpv6NetworkGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallIpv6NetworkGroupCreate,
		Read:   resourceFirewallIpv6NetworkGroupRead,
		Update: resourceFirewallIpv6NetworkGroupUpdate,
		Delete: resourceFirewallIpv6NetworkGroupDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceFirewallIpv6NetworkGroupCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceFirewallIpv6NetworkGroupRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceFirewallIpv6NetworkGroupUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceFirewallIpv6NetworkGroupDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
