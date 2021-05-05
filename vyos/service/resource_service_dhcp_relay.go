package vyos

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceServiceDhcpRelay() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceDhcpRelayCreate,
		Read:   resourceServiceDhcpRelayRead,
		Update: resourceServiceDhcpRelayUpdate,
		Delete: resourceServiceDhcpRelayDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceServiceDhcpRelayCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceDhcpRelayRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceDhcpRelayUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceDhcpRelayDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
