package vyos

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceServiceMdns() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceMdnsCreate,
		Read:   resourceServiceMdnsRead,
		Update: resourceServiceMdnsUpdate,
		Delete: resourceServiceMdnsDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceServiceMdnsCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceMdnsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceMdnsUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceMdnsDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
