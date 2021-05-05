package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceTunnel() *schema.Resource {
	return &schema.Resource{
		Create: resourceInterfaceTunnelCreate,
		Read:   resourceInterfaceTunnelRead,
		Update: resourceInterfaceTunnelUpdate,
		Delete: resourceInterfaceTunnelDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceInterfaceTunnelCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceInterfaceTunnelRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceInterfaceTunnelUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceInterfaceTunnelDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
