package vyos

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFirewallIpv6Name() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallIpv6NameCreate,
		Read:   resourceFirewallIpv6NameRead,
		Update: resourceFirewallIpv6NameUpdate,
		Delete: resourceFirewallIpv6NameDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceFirewallIpv6NameCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceFirewallIpv6NameRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceFirewallIpv6NameUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceFirewallIpv6NameDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
