package vyos

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemConsole() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemConsoleCreate,
		Read:   resourceSystemConsoleRead,
		Update: resourceSystemConsoleUpdate,
		Delete: resourceSystemConsoleDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceSystemConsoleCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceSystemConsoleRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceSystemConsoleUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceSystemConsoleDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
