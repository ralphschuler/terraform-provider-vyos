package vyos

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceServiceDhcpv6Server() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceDhcpv6ServerCreate,
		Read:   resourceServiceDhcpv6ServerRead,
		Update: resourceServiceDhcpv6ServerUpdate,
		Delete: resourceServiceDhcpv6ServerDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceServiceDhcpv6ServerCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceDhcpv6ServerRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceDhcpv6ServerUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceDhcpv6ServerDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
