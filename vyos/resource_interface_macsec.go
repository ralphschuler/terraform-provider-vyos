package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceMacsec() *schema.Resource {
	return &schema.Resource{
		Create: resourceInterfaceMacsecCreate,
		Read:   resourceInterfaceMacsecRead,
		Update: resourceInterfaceMacsecUpdate,
		Delete: resourceInterfaceMacsecDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceInterfaceMacsecCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceInterfaceMacsecRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceInterfaceMacsecUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceInterfaceMacsecDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
