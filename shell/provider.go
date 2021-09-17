package shell

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/mutexkv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {

	// The actual provider
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"environment": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     schema.TypeString,
			},
			"sensitive_environment": {
				Type:      schema.TypeMap,
				Optional:  true,
				Sensitive: true,
				Elem:      schema.TypeString,
			},
			"shell_script": {
				Type:     schema.TypeString,
				Required: true,
			},
		},

		DataSourcesMap: map[string]*schema.Resource{},
		ResourcesMap:   map[string]*schema.Resource{},
		ConfigureFunc:  providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	var environment map[string]interface{}
	if v, ok := d.GetOk("environment"); ok {
		environment = v.(map[string]interface{})
	}
	var sensitiveEnvironment map[string]interface{}
	if v, ok := d.GetOk("sensitive_environment"); ok {
		sensitiveEnvironment = v.(map[string]interface{})
	}
	shellScript := d.Get("shell_script").(string)

	fmt.Println("Hello")

	config := Config{
		Environment:          environment,
		SensitiveEnvironment: sensitiveEnvironment,
		ShellScript:          shellScript,
	}
	return config.Client()
}

// This is a global MutexKV for use within this plugin.
var shellMutexKV = mutexkv.NewMutexKV()

const shellScriptMutexKey = "shellScriptMutexKey"
