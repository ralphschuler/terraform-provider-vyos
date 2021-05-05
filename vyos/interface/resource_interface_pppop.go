package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfacePppoe() *schema.Resource {
	return &schema.Resource{
		Create: resourceInterfacePppoeCreate,
		Read:   resourceInterfacePppoeRead,
		Update: resourceInterfacePppoeUpdate,
		Delete: resourceInterfacePppoeDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceInterfacePppoeCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceInterfacePppoeRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceInterfacePppoeUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceInterfacePppoeDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
