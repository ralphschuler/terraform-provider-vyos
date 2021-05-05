package vyos

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceDummy() *schema.Resource {
	return &schema.Resource{
		Create: resourceInterfaceDummyCreate,
		Read:   resourceInterfaceDummyRead,
		Update: resourceInterfaceDummyUpdate,
		Delete: resourceInterfaceDummyDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceInterfaceDummyCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceInterfaceDummyRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceInterfaceDummyUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceInterfaceDummyDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
