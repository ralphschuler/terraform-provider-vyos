package vyos

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceBridge() *schema.Resource {
	return &schema.Resource{
		Create: resourceInterfaceBridgeCreate,
		Read:   resourceInterfaceBridgeRead,
		Update: resourceInterfaceBridgeUpdate,
		Delete: resourceInterfaceBridgeDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceInterfaceBridgeCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceInterfaceBridgeRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceInterfaceBridgeUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceInterfaceBridgeDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
