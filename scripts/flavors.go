import (
    "github.com/gophercloud/gophercloud"
    "github.com/gophercloud/gophercloud/openstack"
    "time"
)

func collectFlavorData() {

    // Initialize OpenStack Compute (Nova) client using Gophercloud
    computeClient := openstack.NewComputeClient(provider)

    flavors := computeClient.ListFlavors()

    timestamp := currentUnixTime()

    for each flavor in flavors {

        // Collect flavor descriptive data
        saveFlavorDescription(
            timestamp,
            flavor.ID,
            flavor.Name
        )

        // Collect flavor hardware specifications defining
        // the virtual resource configuration
        saveFlavorSpecification(
            timestamp,
            flavor.ID,
            flavor.VCPUs,
            flavor.RAM,
            flavor.Disk
        )
    }
}