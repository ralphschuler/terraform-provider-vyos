package vyos

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNat66DestinationRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceNat66DestinationRuleCreate,
		Read:   resourceNat66DestinationRuleRead,
		Update: resourceNat66DestinationRuleUpdate,
		Delete: resourceNat66DestinationRuleDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceNat66DestinationRuleCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceNat66DestinationRuleRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceNat66DestinationRuleUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceNat66DestinationRuleDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
