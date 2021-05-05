package vyos

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNat66SourceRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceNat66SourceRuleCreate,
		Read:   resourceNat66SourceRuleRead,
		Update: resourceNat66SourceRuleUpdate,
		Delete: resourceNat66SourceRuleDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceNat66SourceRuleCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceNat66SourceRuleRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceNat66SourceRuleUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceNat66SourceRuleDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
