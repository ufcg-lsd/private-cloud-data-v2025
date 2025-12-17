import (
    "flag"
    "time"
)

func main() {

    initializeLogger(logFilePath)

    connectToDatabase()

    provider := initializeOpenStackProvider()

    flags := parseExecutionFlags()

    // Execute server-related data collection routines
    if flags.collectServers {
        collectServerData()
    }

    // Execute project-related data collection routines
    if flags.collectProjects {
        collectProjectData()
    }

    // Execute flavor-related data collection routines
    if flags.collectFlavors {
        collectFlavorData()
    }

    // Execute user-related data collection routines
    if flags.collectUsers {
        collectUserData()
    }

    // Optionally execute all collection routines in a single run
    if flags.collectAll {
        collectServerData()
        collectProjectData()
        collectFlavorData()
        collectUserData()
    }
}