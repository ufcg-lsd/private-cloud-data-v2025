import (
    "github.com/gophercloud/gophercloud"
    "github.com/gophercloud/gophercloud/openstack"
    "time"
)

func collectUserData() {

    // Initialize OpenStack Identity (Keystone) client using Gophercloud
    identityClient := openstack.NewIdentityClient(provider)

    users := identityClient.ListUsers(domainID)

    timestamp := currentUnixTime()

    for each user in users {

        // Collect user descriptive data
        saveUserDescription(
            timestamp,
            user.ID,
            user.Name
        )

        // Collect user–project relationships by querying all projects
        // associated with the current user identifier
        projects := identityClient.ListProjectsForUser(user.ID)

        for each project in projects {

            // Persist UUID-based user–project associations in the database
            saveUserProjectRelation(
                timestamp,
                user.ID,
                project.ID
            )
        }
    }
}