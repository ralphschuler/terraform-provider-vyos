package vyos

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceServiceWebproxy() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceWebproxyCreate,
		Read:   resourceServiceWebproxyRead,
		Update: resourceServiceWebproxyUpdate,
		Delete: resourceServiceWebproxyDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceServiceWebproxyCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceWebproxyRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceWebproxyUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceWebproxyDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
