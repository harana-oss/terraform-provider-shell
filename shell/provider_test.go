package shell

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProvider.ConfigureFunc = testProviderConfigure

	testAccProviders = map[string]terraform.ResourceProvider{
		"shell": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testProviderConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{}
	config.Environment = map[string]interface{}{
		"TEST_ENV1": "Env1_Val01",
		"TEST_ENV2": "Env2_Val02",
	}

	return config.Client()
}
