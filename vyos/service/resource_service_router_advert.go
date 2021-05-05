package vyos

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceServiceRouterAdvert() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceRouterAdvertCreate,
		Read:   resourceServiceRouterAdvertRead,
		Update: resourceServiceRouterAdvertUpdate,
		Delete: resourceServiceRouterAdvertDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceServiceRouterAdvertCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceRouterAdvertRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceRouterAdvertUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceRouterAdvertDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
