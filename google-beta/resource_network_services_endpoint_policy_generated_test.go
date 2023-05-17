// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccNetworkServicesEndpointPolicy_networkServicesEndpointPolicyBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckNetworkServicesEndpointPolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkServicesEndpointPolicy_networkServicesEndpointPolicyBasicExample(context),
			},
			{
				ResourceName:            "google_network_services_endpoint_policy.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name"},
			},
		},
	})
}

func testAccNetworkServicesEndpointPolicy_networkServicesEndpointPolicyBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_network_services_endpoint_policy" "default" {
  provider               = google-beta
  name                   = "tf-test-my-endpoint-policy%{random_suffix}"
  labels                 = {
    foo = "bar"
  }
  description            = "my description"
  type                   = "SIDECAR_PROXY"
  traffic_port_selector {
    ports = ["8081"]
  }
  endpoint_matcher {
    metadata_label_matcher {
      metadata_label_match_criteria = "MATCH_ANY"
      metadata_labels {
        label_name = "foo"
        label_value = "bar"
      }
    }
  }
}
  
`, context)
}

func TestAccNetworkServicesEndpointPolicy_networkServicesEndpointPolicyEmptyMatchExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckNetworkServicesEndpointPolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkServicesEndpointPolicy_networkServicesEndpointPolicyEmptyMatchExample(context),
			},
			{
				ResourceName:            "google_network_services_endpoint_policy.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name"},
			},
		},
	})
}

func testAccNetworkServicesEndpointPolicy_networkServicesEndpointPolicyEmptyMatchExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_network_services_endpoint_policy" "default" {
  provider               = google-beta
  name                   = "tf-test-my-endpoint-policy%{random_suffix}"
  labels                 = {
    foo = "bar"
  }
  description            = "my description"
  type                   = "SIDECAR_PROXY"
  traffic_port_selector {
    ports = ["8081"]
  }
  endpoint_matcher {
    metadata_label_matcher {
      metadata_label_match_criteria = "MATCH_ANY"
    }
  }
}
  
`, context)
}

func testAccCheckNetworkServicesEndpointPolicyDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_network_services_endpoint_policy" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{NetworkServicesBasePath}}projects/{{project}}/locations/global/endpointPolicies/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("NetworkServicesEndpointPolicy still exists at %s", url)
			}
		}

		return nil
	}
}