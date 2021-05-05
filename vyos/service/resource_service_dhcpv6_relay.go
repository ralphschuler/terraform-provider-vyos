package vyos

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceServiceDhcpv6Relay() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceDhcpv6RelayCreate,
		Read:   resourceServiceDhcpv6RelayRead,
		Update: resourceServiceDhcpv6RelayUpdate,
		Delete: resourceServiceDhcpv6RelayDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceServiceDhcpv6RelayCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceDhcpv6RelayRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceDhcpv6RelayUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceDhcpv6RelayDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
