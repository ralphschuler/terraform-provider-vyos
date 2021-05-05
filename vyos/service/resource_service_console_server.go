package vyos

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceServiceConsoleServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceConsoleServerCreate,
		Read:   resourceServiceConsoleServerRead,
		Update: resourceServiceConsoleServerUpdate,
		Delete: resourceServiceConsoleServerDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceServiceConsoleServerCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceConsoleServerRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceConsoleServerUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceConsoleServerDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
