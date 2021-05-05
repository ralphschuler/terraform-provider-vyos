package vyos

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFirewallPortGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallPortGroupCreate,
		Read:   resourceFirewallPortGroupRead,
		Update: resourceFirewallPortGroupUpdate,
		Delete: resourceFirewallPortGroupDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceFirewallPortGroupCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceFirewallPortGroupRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceFirewallPortGroupUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceFirewallPortGroupDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
