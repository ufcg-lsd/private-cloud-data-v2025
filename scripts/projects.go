import (
    "github.com/gophercloud/gophercloud"
    "github.com/gophercloud/gophercloud/openstack"
    "time"
)

func collectProjectData() {

    // Initialize OpenStack Identity (Keystone) client using Gophercloud
    identityClient := openstack.NewIdentityClient(provider)

    // Initialize OpenStack Compute (Nova) client for quota operations
    computeClient := openstack.NewComputeClient(provider)

    projects := identityClient.ListProjects(domainID)

    timestamp := currentUnixTime()

    for each project in projects {

        // Collect project descriptive data
        saveProjectDescription(
            timestamp,
            project.ID,
            project.Name,
            project.Description
        )

        // Collect quota limits associated with the current project
        quotas := computeClient.GetProjectQuotas(project.ID)

        saveProjectQuota(
            timestamp,
            project.ID,
            quotas.RAM,
            quotas.VCPUs
        )

        // Collect project-level allocated resources aggregated over
        // a five-minute observation window
        allocated := computeClient.GetProjectAllocatedResources(
            project.ID,
            window = "5m"
        )

        saveProjectAllocatedResources(
            timestamp,
            project.ID,
            allocated.RAM,
            allocated.VCPUs
        )
    }
}