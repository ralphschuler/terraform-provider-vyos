package vyos

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceServiceDhcpServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceDhcpServerCreate,
		Read:   resourceServiceDhcpServerRead,
		Update: resourceServiceDhcpServerUpdate,
		Delete: resourceServiceDhcpServerDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceServiceDhcpServerCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceDhcpServerRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceDhcpServerUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceDhcpServerDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
