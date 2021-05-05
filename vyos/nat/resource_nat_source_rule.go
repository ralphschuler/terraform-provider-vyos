package vyos

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNatSourceRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceNatSourceRuleCreate,
		Read:   resourceNatSourceRuleRead,
		Update: resourceNatSourceRuleUpdate,
		Delete: resourceNatSourceRuleDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceNatSourceRuleCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceNatSourceRuleRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceNatSourceRuleUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceNatSourceRuleDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
