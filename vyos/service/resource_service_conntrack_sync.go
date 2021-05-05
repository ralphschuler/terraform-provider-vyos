package vyos

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceServiceContrackSync() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceContrackSyncCreate,
		Read:   resourceServiceContrackSyncRead,
		Update: resourceServiceContrackSyncUpdate,
		Delete: resourceServiceContrackSyncDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceServiceContrackSyncCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceContrackSyncRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceContrackSyncUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceContrackSyncDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
