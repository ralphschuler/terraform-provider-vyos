package vyos

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNatDesctinationRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceNatDesctinationRuleCreate,
		Read:   resourceNatDesctinationRuleRead,
		Update: resourceNatDesctinationRuleUpdate,
		Delete: resourceNatDesctinationRuleDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceNatDesctinationRuleCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceNatDesctinationRuleRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceNatDesctinationRuleUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceNatDesctinationRuleDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
