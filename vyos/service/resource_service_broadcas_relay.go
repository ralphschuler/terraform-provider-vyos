package vyos

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceServiceBroadcastRelay() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceBroadcastRelayCreate,
		Read:   resourceServiceBroadcastRelayRead,
		Update: resourceServiceBroadcastRelayUpdate,
		Delete: resourceServiceBroadcastRelayDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceServiceBroadcastRelayCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceBroadcastRelayRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceBroadcastRelayUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceBroadcastRelayDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
