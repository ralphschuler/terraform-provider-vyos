package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceGeneve() *schema.Resource {
	return &schema.Resource{
		Create: resourceInterfaceGeneveCreate,
		Read:   resourceInterfaceGeneveRead,
		Update: resourceInterfaceGeneveUpdate,
		Delete: resourceInterfaceGeneveDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceInterfaceGeneveCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceInterfaceGeneveRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceInterfaceGeneveUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceInterfaceGeneveDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
