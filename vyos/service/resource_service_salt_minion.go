package vyos

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceServiceSaltMinion() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceSaltMinionCreate,
		Read:   resourceServiceSaltMinionRead,
		Update: resourceServiceSaltMinionUpdate,
		Delete: resourceServiceSaltMinionDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceServiceSaltMinionCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceSaltMinionRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceSaltMinionUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceSaltMinionDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
