package vyos

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceServiceLldp() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceLldpCreate,
		Read:   resourceServiceLldpRead,
		Update: resourceServiceLldpUpdate,
		Delete: resourceServiceLldpDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceServiceLldpCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceLldpRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceLldpUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceLldpDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
