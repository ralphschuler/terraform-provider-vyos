package vyos

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceServiceDnsDynamic() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceDnsDynamicCreate,
		Read:   resourceServiceDnsDynamicRead,
		Update: resourceServiceDnsDynamicUpdate,
		Delete: resourceServiceDnsDynamicDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceServiceDnsDynamicCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceDnsDynamicRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceDnsDynamicUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceDnsDynamicDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
