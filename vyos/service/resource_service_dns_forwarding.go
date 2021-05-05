package vyos

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceServiceDnsForwarding() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceDnsForwardingCreate,
		Read:   resourceServiceDnsForwardingRead,
		Update: resourceServiceDnsForwardingUpdate,
		Delete: resourceServiceDnsForwardingDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceServiceDnsForwardingCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceDnsForwardingRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceDnsForwardingUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceDnsForwardingDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
