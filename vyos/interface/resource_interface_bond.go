package vyos

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceBond() *schema.Resource {
	return &schema.Resource{
		Create: resourceInterfaceBondCreate,
		Read:   resourceInterfaceBondRead,
		Update: resourceInterfaceBondUpdate,
		Delete: resourceInterfaceBondDelete,

		Schema: map[string]*schema.Schema{
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"disable": &schema.Schema{
				Type:     schema.TypeBool,
				Required: false,
			},
			"disableFlowControl": &schema.Schema{
				Type:     schema.TypeBool,
				Required: false,
			},
			"disableLinkDetect": &schema.Schema{
				Type:     schema.TypeBool,
				Required: false,
			},
			"mac": &schema.Schema{
				Type:     schema.TypeBool,
				Required: false,
			},
			"mtu": &schema.Schema{
				Type:     schema.TypeInt,
				Required: false,
			},
			"ip": &schema.Resource{ //Continue here
				Type:     schema.TypeMap,
				Required: false,
			},
		},
	}
}

func resourceInterfaceBondCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceInterfaceBondRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceInterfaceBondUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceInterfaceBondDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
