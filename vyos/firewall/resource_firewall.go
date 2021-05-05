package vyos

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFirewall() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallCreate,
		Read:   resourceFirewallRead,
		Update: resourceFirewallUpdate,
		Delete: resourceFirewallDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceFirewallCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceFirewallRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceFirewallUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceFirewallDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
