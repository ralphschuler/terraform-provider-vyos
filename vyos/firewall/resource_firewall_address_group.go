package vyos

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFirewallAddressGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallAddressGroupCreate,
		Read:   resourceFirewallAddressGroupRead,
		Update: resourceFirewallAddressGroupUpdate,
		Delete: resourceFirewallAddressGroupDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceFirewallAddressGroupCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceFirewallAddressGroupRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceFirewallAddressGroupUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceFirewallAddressGroupDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
