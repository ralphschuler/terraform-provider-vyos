package vyos

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceServiceSaltSnmp() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceSaltSnmpCreate,
		Read:   resourceServiceSaltSnmpRead,
		Update: resourceServiceSaltSnmpUpdate,
		Delete: resourceServiceSaltSnmpDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceServiceSaltSnmpCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceSaltSnmpRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceSaltSnmpUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceSaltSnmpDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
