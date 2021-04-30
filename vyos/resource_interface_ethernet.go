package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceEthernet() *schema.Resource {
	return &schema.Resource{
		Create: resourceInterfaceEthernetCreate,
		Read:   resourceInterfaceEthernetRead,
		Update: resourceInterfaceEthernetUpdate,
		Delete: resourceInterfaceEthernetDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceInterfaceEthernetCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceInterfaceEthernetRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceInterfaceEthernetUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceInterfaceEthernetDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
