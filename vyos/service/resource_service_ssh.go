package vyos

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceServiceSsh() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceSshCreate,
		Read:   resourceServiceSshRead,
		Update: resourceServiceSshUpdate,
		Delete: resourceServiceSshDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceServiceSshCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceSshRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceSshUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceSshDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
