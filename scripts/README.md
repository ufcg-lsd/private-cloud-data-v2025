# Data Collection Scripts

This directory provides pseudocode samples of the data collection scripts used to build the dataset [Dataset on Resource Allocation and Usage for a Private Cloud
](http://doi.org/10.17632/trvb5k4x5m.1).  

The purpose of these scripts is to document the data collection methodology and to clarify how different data domains were obtained from the cloud infrastructure. In particular, the collection process captures projects, understood as organizational units that group users and define resource allocation quotas; users, who act as cloud clients and may be associated with one or more projects; flavors, which represent hardware configurations for virtual machines; and servers, understood here as virtual machine instances, together with their corresponding resource usage metrics over time.

The scripts presented here are not intended to be executed directly. Instead, they abstract the logic implemented in the original data collection tool, highlighting the operations and the external services involved at each stage of the data collection process.

---

## Scope and Design Principles

The data collection process was designed with the following principles: 

- Modularity: each data domain (servers, projects, users, flavors) is collected by an independent routine.
- Periodic execution: collection routines are designed to be executed repeatedly over time.
- Service-oriented acquisition: data are retrieved via OpenStack APIs and Prometheus endpoints.
- Timestamped records: all collected data are associated with a UNIX epoch timestamp.
- Separation of concerns: data collection, persistence and scheduling.

---

## Script Overview

The following scripts are provided in this directory (`data-collection/scripts/`):

- ```users.go```: collection of user metadata and user–project associations.
- ```projects.go```: collection of project metadata, quota limits, and allocated resources.
- ```servers.go```: collection of server metadata, ownership relations, and resource usage metrics.
- ```flavors.go```: collection of flavor metadata and hardware specifications.
- ```main.go```: orchestration of the data collection workflow.

---

## Environment Configuration

The collection process relies on environment-based configuration, including OpenStack authentication credentials, domain and region identifiers, database connection parameters, and logging settings.

The data collection service was executed as a background process on a server within the private cloud infrastructure, with the following characteristics:

- **Operating system:** Linux (Ubuntu 22.04)
- **Execution environment:** Server deployed in the private cloud
- **Server resources:** 1 vCPU, 2 GB RAM, 20 GB disk
- **Programming language:** Go (version 1.22.2)
- **Cloud platform:** OpenStack
- **OpenStack client library:** Gophercloud (Go SDK, version 1.11.0)
- **Metrics source:** Prometheus (libvirt-based exporters)
- **Database:** PostgreSQL (containerized, version 16.2)
- **Sampling interval:** Fixed five-minute intervals
- **Execution model:** Periodic execution via a scheduler (e.g., cron)

All authentication credentials, service endpoints, and other sensitive configuration parameters were supplied through environment variables in the production deployment and are intentionally omitted from this documentation.

## Reproducibility Note

These scripts are provided for methodological transparency and reproducibility and are intended as reference material.