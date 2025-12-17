import (
    "github.com/gophercloud/gophercloud"
    "github.com/gophercloud/gophercloud/openstack"
    "time"
)

func collectServerData() {

    // Initialize OpenStack Compute (Nova) client using Gophercloud
    computeClient := openstack.NewComputeClient(provider)

    servers := computeClient.ListServers(allTenants = true)

    timestamp := currentUnixTime()

    for each server in servers {

        // Collect server descriptive data
        saveServerDescription(
            timestamp,
            server.ID,
            server.Name,
            server.FixedIPAddress
        )

        // Collect server specification data by associating each server
        // with its corresponding flavor identifier
        saveServerSpecification(
            timestamp,
            server.ID,
            server.FlavorID
        )

        // Collect ownership relationships linking servers to users
        // and projects through UUID-based identifiers
        saveServerOwnership(
            timestamp,
            server.ID,
            server.UserID,
            server.ProjectID
        )

        // Query Prometheus to retrieve the hypervisor domain associated
        // with the server UUID
        domain := queryPrometheus(
            "libvirt_domain_info_meta",
            server.ID
        )

        // Retrieve memory usage metrics for the corresponding domain
        ramUsage := queryPrometheus(
            "libvirt_domain_memory_stats_used_percent",
            domain
        )

        // Retrieve vCPU usage metrics using a five-minute rate window
        vcpuUsage := queryPrometheus(
            "libvirt_domain_info_cpu_time_seconds_total",
            domain,
            window = "5m"
        )

        // Persist timestamped resource usage metrics in the database
        saveServerUsage(
            timestamp,
            server.ID,
            ramUsage,
            vcpuUsage
        )
    }
}