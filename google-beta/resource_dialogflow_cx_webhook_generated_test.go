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
)

func TestAccDialogflowCXWebhook_dialogflowcxWebhookFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDialogflowCXWebhookDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDialogflowCXWebhook_dialogflowcxWebhookFullExample(context),
			},
			{
				ResourceName:            "google_dialogflow_cx_webhook.basic_webhook",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parent"},
			},
		},
	})
}

func testAccDialogflowCXWebhook_dialogflowcxWebhookFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_dialogflow_cx_agent" "agent" {
  display_name = "tf-test-dialogflowcx-agent%{random_suffix}"
  location = "global"
  default_language_code = "en"
  supported_language_codes = ["it","de","es"]
  time_zone = "America/New_York"
  description = "Example description."
  avatar_uri = "https://cloud.google.com/_static/images/cloud/icons/favicons/onecloud/super_cloud.png"
  enable_stackdriver_logging = true
  enable_spell_correction    = true
	speech_to_text_settings {
		enable_speech_adaptation = true
	}
}


resource "google_dialogflow_cx_webhook" "basic_webhook" {
  parent       = google_dialogflow_cx_agent.agent.id
  display_name = "MyFlow"
  generic_web_service {
		uri = "https://example.com"
	}
}
`, context)
}

func testAccCheckDialogflowCXWebhookDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_dialogflow_cx_webhook" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{DialogflowCXBasePath}}{{parent}}/webhooks/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("DialogflowCXWebhook still exists at %s", url)
			}
		}

		return nil
	}
}
