// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package redis_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccRedisCluster_createClusterWithNodeType(t *testing.T) {

	t.Parallel()

	name := fmt.Sprintf("tf-test-%d", acctest.RandInt(t))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckRedisClusterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				// create cluster with replica count 1
				Config: createOrUpdateRedisCluster(&ClusterParams{name: name, replicaCount: 1, shardCount: 3, deletionProtectionEnabled: true, nodeType: "REDIS_STANDARD_SMALL", zoneDistributionMode: "MULTI_ZONE", maintenanceDay: "TUESDAY", maintenanceHours: 2, maintenanceMinutes: 0, maintenanceSeconds: 0, maintenanceNanos: 0}),
			},
			{
				ResourceName:            "google_redis_cluster.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"psc_configs"},
			},
			{
				// clean up the resource
				Config: createOrUpdateRedisCluster(&ClusterParams{name: name, replicaCount: 1, shardCount: 3, deletionProtectionEnabled: false, nodeType: "REDIS_STANDARD_SMALL", zoneDistributionMode: "MULTI_ZONE", maintenanceDay: "TUESDAY", maintenanceHours: 2, maintenanceMinutes: 0, maintenanceSeconds: 0, maintenanceNanos: 0}),
			},
		},
	})
}

// Validate zone distribution for the cluster.
func TestAccRedisCluster_createClusterWithZoneDistribution(t *testing.T) {
	t.Parallel()

	name := fmt.Sprintf("tf-test-%d", acctest.RandInt(t))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckRedisClusterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				// create cluster with replica count 1
				Config: createOrUpdateRedisCluster(&ClusterParams{name: name, replicaCount: 0, shardCount: 3, deletionProtectionEnabled: false, zoneDistributionMode: "SINGLE_ZONE", zone: "us-central1-b"}),
			},
			{
				ResourceName:            "google_redis_cluster.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"psc_configs"},
			},
			{
				// clean up the resource
				Config: createOrUpdateRedisCluster(&ClusterParams{name: name, replicaCount: 0, shardCount: 3, deletionProtectionEnabled: false, zoneDistributionMode: "SINGLE_ZONE", zone: "us-central1-b"}),
			},
		},
	})
}

// Validate that replica count is updated for the cluster
func TestAccRedisCluster_updateReplicaCount(t *testing.T) {
	t.Parallel()

	name := fmt.Sprintf("tf-test-%d", acctest.RandInt(t))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckRedisClusterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				// create cluster with replica count 1
				Config: createOrUpdateRedisCluster(&ClusterParams{name: name, replicaCount: 1, shardCount: 3, deletionProtectionEnabled: true, zoneDistributionMode: "MULTI_ZONE", maintenanceDay: "MONDAY", maintenanceHours: 1, maintenanceMinutes: 0, maintenanceSeconds: 0, maintenanceNanos: 0}),
			},
			{
				ResourceName:            "google_redis_cluster.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"psc_configs"},
			},
			{
				// update replica count to 2
				Config: createOrUpdateRedisCluster(&ClusterParams{name: name, replicaCount: 2, shardCount: 3, deletionProtectionEnabled: true, zoneDistributionMode: "MULTI_ZONE", maintenanceDay: "MONDAY", maintenanceHours: 1, maintenanceMinutes: 0, maintenanceSeconds: 0, maintenanceNanos: 0}),
			},
			{
				ResourceName:            "google_redis_cluster.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"psc_configs"},
			},
			{
				// update replica count to 0
				Config: createOrUpdateRedisCluster(&ClusterParams{name: name, replicaCount: 0, shardCount: 3, deletionProtectionEnabled: true, zoneDistributionMode: "MULTI_ZONE", maintenanceDay: "MONDAY", maintenanceHours: 1, maintenanceMinutes: 0, maintenanceSeconds: 0, maintenanceNanos: 0}),
			},
			{
				ResourceName:            "google_redis_cluster.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"psc_configs"},
			},
			{
				// clean up the resource
				Config: createOrUpdateRedisCluster(&ClusterParams{name: name, replicaCount: 0, shardCount: 3, deletionProtectionEnabled: false, zoneDistributionMode: "MULTI_ZONE", maintenanceDay: "MONDAY", maintenanceHours: 1, maintenanceMinutes: 0, maintenanceSeconds: 0, maintenanceNanos: 0}),
			},
		},
	})
}

// Validate that shard count is updated for the cluster
func TestAccRedisCluster_updateShardCount(t *testing.T) {
	t.Parallel()

	name := fmt.Sprintf("tf-test-%d", acctest.RandInt(t))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckRedisClusterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				// create cluster with shard count 3
				Config: createOrUpdateRedisCluster(&ClusterParams{name: name, replicaCount: 1, shardCount: 3, deletionProtectionEnabled: true, zoneDistributionMode: "MULTI_ZONE", maintenanceDay: "MONDAY", maintenanceHours: 1, maintenanceMinutes: 0, maintenanceSeconds: 0, maintenanceNanos: 0}),
			},
			{
				ResourceName:            "google_redis_cluster.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"psc_configs"},
			},
			{
				// update shard count to 5
				Config: createOrUpdateRedisCluster(&ClusterParams{name: name, replicaCount: 1, shardCount: 5, deletionProtectionEnabled: true, zoneDistributionMode: "MULTI_ZONE", maintenanceDay: "MONDAY", maintenanceHours: 1, maintenanceMinutes: 0, maintenanceSeconds: 0, maintenanceNanos: 0}),
			},
			{
				ResourceName:            "google_redis_cluster.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"psc_configs"},
			},
			{
				// clean up the resource
				Config: createOrUpdateRedisCluster(&ClusterParams{name: name, replicaCount: 1, shardCount: 5, deletionProtectionEnabled: false, zoneDistributionMode: "MULTI_ZONE", maintenanceDay: "MONDAY", maintenanceHours: 1, maintenanceMinutes: 0, maintenanceSeconds: 0, maintenanceNanos: 0}),
			},
		},
	})
}

// Validate that redisConfigs is updated for the cluster
func TestAccRedisCluster_updateRedisConfigs(t *testing.T) {
	t.Parallel()

	name := fmt.Sprintf("tf-test-%d", acctest.RandInt(t))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckRedisClusterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				// create cluster
				Config: createOrUpdateRedisCluster(&ClusterParams{
					name:                 name,
					shardCount:           3,
					zoneDistributionMode: "MULTI_ZONE",
					maintenanceDay:       "MONDAY",
					maintenanceHours:     1,
					maintenanceMinutes:   0,
					maintenanceSeconds:   0,
					maintenanceNanos:     0,
					redisConfigs: map[string]string{
						"maxmemory-policy": "volatile-ttl",
					}}),
			},
			{
				ResourceName:            "google_redis_cluster.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"psc_configs"},
			},
			{
				// add a new redis config key-value pair and update existing redis config
				Config: createOrUpdateRedisCluster(&ClusterParams{
					name:                 name,
					shardCount:           3,
					zoneDistributionMode: "MULTI_ZONE",
					maintenanceDay:       "MONDAY",
					maintenanceHours:     1,
					maintenanceMinutes:   0,
					maintenanceSeconds:   0,
					maintenanceNanos:     0,
					redisConfigs: map[string]string{
						"maxmemory-policy":  "allkeys-lru",
						"maxmemory-clients": "90%",
					}}),
			},
			{
				ResourceName:            "google_redis_cluster.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"psc_configs"},
			},
			{
				// remove all redis configs
				Config: createOrUpdateRedisCluster(&ClusterParams{name: name, shardCount: 3, zoneDistributionMode: "MULTI_ZONE", maintenanceDay: "MONDAY", maintenanceHours: 1, maintenanceMinutes: 0, maintenanceSeconds: 0, maintenanceNanos: 0}),
			},
		},
	})
}

// Validate that deletion protection enabled/disabled cluster is created updated
func TestAccRedisCluster_createUpdateDeletionProtection(t *testing.T) {
	t.Parallel()

	name := fmt.Sprintf("tf-test-%d", acctest.RandInt(t))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckRedisClusterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				// create cluster with deletion protection set to false
				Config: createOrUpdateRedisCluster(&ClusterParams{name: name, replicaCount: 0, shardCount: 3, deletionProtectionEnabled: false, zoneDistributionMode: "MULTI_ZONE"}),
			},
			{
				ResourceName:            "google_redis_cluster.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"psc_configs"},
			},
			{
				// update deletion protection to true
				Config: createOrUpdateRedisCluster(&ClusterParams{name: name, replicaCount: 0, shardCount: 3, deletionProtectionEnabled: true, zoneDistributionMode: "MULTI_ZONE"}),
			},
			{
				ResourceName:            "google_redis_cluster.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"psc_configs"},
			},
			{
				// update deletion protection to false and delete the cluster
				Config: createOrUpdateRedisCluster(&ClusterParams{name: name, replicaCount: 0, shardCount: 3, deletionProtectionEnabled: false, zoneDistributionMode: "MULTI_ZONE"}),
			},
		},
	})
}

// Validate that persistence is updated for the cluster
func TestAccRedisCluster_persistenceUpdate(t *testing.T) {
	t.Parallel()

	name := fmt.Sprintf("tf-test-%d", acctest.RandInt(t))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckRedisClusterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				// create cluster with AOF enabled
				Config: createOrUpdateRedisCluster(&ClusterParams{name: name, replicaCount: 0, shardCount: 3, nodeType: "REDIS_STANDARD_SMALL", zoneDistributionMode: "MULTI_ZONE", persistenceBlock: "persistence_config {\nmode = \"AOF\"\naof_config{\nappend_fsync = \"EVERYSEC\"\n}\n}"}),
			},
			{
				ResourceName:            "google_redis_cluster.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"psc_configs"},
			},
			{
				// disable AOF
				Config: createOrUpdateRedisCluster(&ClusterParams{name: name, replicaCount: 0, shardCount: 3, nodeType: "REDIS_STANDARD_SMALL", zoneDistributionMode: "MULTI_ZONE", persistenceBlock: "persistence_config {\nmode = \"DISABLED\"\n}"}),
			},
			{
				ResourceName:            "google_redis_cluster.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"psc_configs"},
			}, {
				// update persistence to RDB
				Config: createOrUpdateRedisCluster(&ClusterParams{name: name, replicaCount: 0, shardCount: 3, nodeType: "REDIS_STANDARD_SMALL", zoneDistributionMode: "MULTI_ZONE", persistenceBlock: "persistence_config {\nmode = \"RDB\"\nrdb_config {\nrdb_snapshot_period = \"ONE_HOUR\"\nrdb_snapshot_start_time = \"2024-10-02T15:01:23Z\"\n}\n}"}),
			},
			{
				ResourceName:            "google_redis_cluster.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"psc_configs"},
			},
			{
				// clean up the resource
				Config: createOrUpdateRedisCluster(&ClusterParams{name: name, replicaCount: 0, shardCount: 3, nodeType: "REDIS_STANDARD_SMALL", zoneDistributionMode: "MULTI_ZONE", persistenceBlock: "persistence_config {\nmode = \"RDB\"\nrdb_config {\nrdb_snapshot_period = \"ONE_HOUR\"\nrdb_snapshot_start_time = \"2024-10-02T15:01:23Z\"\n}\n}"}),
			},
		},
	})
}

type ClusterParams struct {
	name                      string
	replicaCount              int
	shardCount                int
	deletionProtectionEnabled bool
	nodeType                  string
	redisConfigs              map[string]string
	zoneDistributionMode      string
	zone                      string
	maintenanceDay            string
	maintenanceHours          int
	maintenanceMinutes        int
	maintenanceSeconds        int
	maintenanceNanos          int
	persistenceBlock          string
}

func createOrUpdateRedisCluster(params *ClusterParams) string {
	var redsConfigsStrBuilder strings.Builder
	for key, value := range params.redisConfigs {
		redsConfigsStrBuilder.WriteString(fmt.Sprintf("%s =  \"%s\"\n", key, value))
	}

	zoneDistributionConfigBlock := ``
	if params.zoneDistributionMode != "" {
		zoneDistributionConfigBlock = fmt.Sprintf(`
		zone_distribution_config {
			mode = "%s"
			zone = "%s"
		}
		`, params.zoneDistributionMode, params.zone)
	}

	maintenancePolicyBlock := ``
	if params.maintenanceDay != "" {
		maintenancePolicyBlock = fmt.Sprintf(`
		maintenance_policy {
			weekly_maintenance_window {
				day = "%s"
				start_time {
					hours = %d
					minutes = %d
					seconds = %d
					nanos = %d
				}
			}
		}
		`, params.maintenanceDay, params.maintenanceHours, params.maintenanceMinutes, params.maintenanceSeconds, params.maintenanceNanos)
	}

	return fmt.Sprintf(`
resource "google_redis_cluster" "test" {
	provider = google-beta
	name           = "%s"
	replica_count = %d
	shard_count = %d
	node_type = "%s"
	deletion_protection_enabled = %v
	region         = "us-central1"
	psc_configs {
			network = google_compute_network.producer_net.id
	}
	redis_configs = {
		%s
	}
  %s
  %s
  %s
	depends_on = [
          google_network_connectivity_service_connection_policy.default
        ]
}

resource "google_network_connectivity_service_connection_policy" "default" {
	provider = google-beta
	name = "%s"
	location = "us-central1"
	service_class = "gcp-memorystore-redis"
	description   = "my basic service connection policy"
	network = google_compute_network.producer_net.id
	psc_config {
	subnetworks = [google_compute_subnetwork.producer_subnet.id]
	}
}

resource "google_compute_subnetwork" "producer_subnet" {
	provider      = google-beta
	name          = "%s"
	ip_cidr_range = "10.0.0.248/29"
	region        = "us-central1"
	network       = google_compute_network.producer_net.id
}

resource "google_compute_network" "producer_net" {
	provider                = google-beta
	name                    = "%s"
	auto_create_subnetworks = false
}
`, params.name, params.replicaCount, params.shardCount, params.nodeType, params.deletionProtectionEnabled, redsConfigsStrBuilder.String(), zoneDistributionConfigBlock, maintenancePolicyBlock, params.persistenceBlock, params.name, params.name, params.name)
}
