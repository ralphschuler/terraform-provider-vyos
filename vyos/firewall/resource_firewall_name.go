package vyos

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFirewallName() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallNameCreate,
		Read:   resourceFirewallNameRead,
		Update: resourceFirewallNameUpdate,
		Delete: resourceFirewallNameDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceFirewallNameCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceFirewallNameRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceFirewallNameUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceFirewallNameDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
