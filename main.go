package main
 
import (
      "github.com/hashicorp/terraform-plugin-sdk/plugin"
      "github.com/hashicorp/terraform-plugin-sdk/terraform"
      "terraform-provider-vyos/vyos"
)
 
func main() {
      plugin.Serve(&plugin.ServeOpts{
        ProviderFunc: func() terraform.ResourceProvider {
                return vyos.Provider()
        },
      })
}