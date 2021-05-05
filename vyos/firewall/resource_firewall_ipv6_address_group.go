package vyos

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFirewallIpv6AddressGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallIpv6AddressGroupCreate,
		Read:   resourceFirewallIpv6AddressGroupRead,
		Update: resourceFirewallIpv6AddressGroupUpdate,
		Delete: resourceFirewallIpv6AddressGroupDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceFirewallIpv6AddressGroupCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceFirewallIpv6AddressGroupRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceFirewallIpv6AddressGroupUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceFirewallIpv6AddressGroupDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
