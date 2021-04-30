package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceLoopback() *schema.Resource {
	return &schema.Resource{
		Create: resourceInterfaceLoopbackCreate,
		Read:   resourceInterfaceLoopbackRead,
		Update: resourceInterfaceLoopbackUpdate,
		Delete: resourceInterfaceLoopbackDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceInterfaceLoopbackCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceInterfaceLoopbackRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceInterfaceLoopbackUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceInterfaceLoopbackDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
