package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceL2tpv3() *schema.Resource {
	return &schema.Resource{
		Create: resourceInterfaceL2tpv3Create,
		Read:   resourceInterfaceL2tpv3Read,
		Update: resourceInterfaceL2tpv3Update,
		Delete: resourceInterfaceL2tpv3Delete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceInterfaceL2tpv3Create(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceInterfaceL2tpv3Read(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceInterfaceL2tpv3Update(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceInterfaceL2tpv3Delete(d *schema.ResourceData, m interface{}) error {
	return nil
}
