package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceMacvlan() *schema.Resource {
	return &schema.Resource{
		Create: resourceInterfaceMacvlanCreate,
		Read:   resourceInterfaceMacvlanRead,
		Update: resourceInterfaceMacvlanUpdate,
		Delete: resourceInterfaceMacvlanDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceInterfaceMacvlanCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceInterfaceMacvlanRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceInterfaceMacvlanUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceInterfaceMacvlanDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
