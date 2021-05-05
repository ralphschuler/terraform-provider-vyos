package vyos

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceServiceTftpServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceTftpServerCreate,
		Read:   resourceServiceTftpServerRead,
		Update: resourceServiceTftpServerUpdate,
		Delete: resourceServiceTftpServerDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceServiceTftpServerCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceTftpServerRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceTftpServerUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceTftpServerDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
