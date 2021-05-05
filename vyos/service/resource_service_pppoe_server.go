package vyos

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceServicePppoeServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceServicePppoeServerCreate,
		Read:   resourceServicePppoeServerRead,
		Update: resourceServicePppoeServerUpdate,
		Delete: resourceServicePppoeServerDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceServicePppoeServerCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServicePppoeServerRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServicePppoeServerUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServicePppoeServerDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
