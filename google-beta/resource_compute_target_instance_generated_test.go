// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

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

func TestAccComputeTargetInstance_targetInstanceBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeTargetInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeTargetInstance_targetInstanceBasicExample(context),
			},
			{
				ResourceName:            "google_compute_target_instance.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"instance", "zone"},
			},
		},
	})
}

func testAccComputeTargetInstance_targetInstanceBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_target_instance" "default" {
  name     = "target%{random_suffix}"
  instance = google_compute_instance.target-vm.id
}

data "google_compute_image" "vmimage" {
  family  = "debian-11"
  project = "debian-cloud"
}

resource "google_compute_instance" "target-vm" {
  name         = "tf-test-target-vm%{random_suffix}"
  machine_type = "e2-medium"
  zone         = "us-central1-a"

  boot_disk {
    initialize_params {
      image = data.google_compute_image.vmimage.self_link
    }
  }

  network_interface {
    network = "default"
  }
}
`, context)
}

func TestAccComputeTargetInstance_targetInstanceCustomNetworkExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeTargetInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeTargetInstance_targetInstanceCustomNetworkExample(context),
			},
			{
				ResourceName:            "google_compute_target_instance.custom_network",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"instance", "zone"},
			},
		},
	})
}

func testAccComputeTargetInstance_targetInstanceCustomNetworkExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_target_instance" "custom_network" {
  provider = google-beta
  name     = "tf-test-custom-network%{random_suffix}"
  instance = google_compute_instance.target-vm.id
  network  = data.google_compute_network.target-vm.self_link
}

data "google_compute_network" "target-vm" {
  provider = google-beta
  name = "default"
}

data "google_compute_image" "vmimage" {
  provider = google-beta
  family  = "debian-10"
  project = "debian-cloud"
}

resource "google_compute_instance" "target-vm" {
  provider = google-beta
  name         = "tf-test-custom-network-target-vm%{random_suffix}"
  machine_type = "e2-medium"
  zone         = "us-central1-a"

  boot_disk {
    initialize_params {
      image = data.google_compute_image.vmimage.self_link
    }
  }

  network_interface {
    network = "default"
  }
}
`, context)
}

func testAccCheckComputeTargetInstanceDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_target_instance" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/targetInstances/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("ComputeTargetInstance still exists at %s", url)
			}
		}

		return nil
	}
}
