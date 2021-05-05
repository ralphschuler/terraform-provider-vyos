package vyos

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceServiceIpoeServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceIpoeServerCreate,
		Read:   resourceServiceIpoeServerRead,
		Update: resourceServiceIpoeServerUpdate,
		Delete: resourceServiceIpoeServerDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceServiceIpoeServerCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceIpoeServerRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceIpoeServerUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceIpoeServerDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
