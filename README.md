# Dataset on Resource Allocation and Usage for a Private Cloud

## Introduction
The dataset consists of records collected from a private OpenStack-based cloud operated by the Distributed Systems Laboratory at the Federal University of Campina Grande, Brazil, over approximately one year of operation.

This repository contains sample data collection scripts and an [R Markdown notebook]() for generating plots based on the dataset.

<br>
<p align="center">
  <a href="https://doi.org/10.17632/trvb5k4x5m.1"><img alt="DOI" src="https://img.shields.io/badge/DOI-10.17632%2Ftrvb5k4x5m.1-blue"></a>
  <a href="#license"><img alt="Data License" src="https://img.shields.io/badge/Data%20License-CC%20BY%204.0-black"></a>
  <a href="#citation"><img alt="Cite" src="https://img.shields.io/badge/Cite-Please%20cite%20the%20dataset-green"></a>
</p>

## Using the data

### License

The data is made available and licensed under a [CC-BY Attribution 4.0 License](https://creativecommons.org/licenses/by/4.0/). By downloading it or using it, you agree to the terms of this license.

### Attribution

If you use the data for a publication or project, please cite the accompanying dataset:

```
Marques, Paola; Mendes, Mariana; Pereira, Thiago Emmanuel; Farias, Giovanni (2025),
“Dataset on Resource Allocation and Usage for a Private Cloud”, Mendeley Data,
V1, doi: 10.17632/trvb5k4x5m.1
```

### Unzipping

The dataset is available on Mendeley Data as a compressed archive.

After downloading the ZIP file, you can extract it using:

```
unzip "Dataset on Resource Allocation and Usage for a Private Cloud.zip"
```

### Schema

The dataset is organized into seven CSV files, each representing a distinct data table:

#### projects_quota.csv

| Field | Description |
|--|--|
| timestamp | Collection time in UNIX epoch seconds |
| project_id | Unique project identifier (UUID) |
| quota_ram | RAM quota available to the project (MB) |
| quota_vcpu | vCPU quota available to the project |

#### projects_quota_allocated.csv

| Field | Description |
|--|--|
| id | Unique row identifier |
| timestamp | Collection time in UNIX epoch seconds |
| project_id | Unique project identifier (UUID) |
| ram_allocated | RAM quota allocated per project (MB) |
| vcpu_allocated | vCPU quota allocated per project |

#### user_projects.csv

| Field | Description |
|--|--|
| id | Unique row identifier |
| timestamp | Collection time in UNIX epoch seconds |
| user_id | Unique user identifier (UUID) |
| project_id | Unique project identifier (UUID) |

#### servers_specs.csv

| Field | Description |
|--|--|
| id | Unique row identifier |
| timestamp | Collection time in UNIX epoch seconds |
| server_id | Unique server identifier (UUID) |
| flavor_id | Unique flavor identifier (UUID) |

#### servers_ownerships.csv

| Field | Description |
|--|--|
| id | Unique row identifier |
| timestamp | Collection time in UNIX epoch seconds |
| server_id | Unique server identifier (UUID) |
| user_id | Unique user identifier (UUID) |
| project_id | Unique project identifier (UUID) |

#### servers_usage.csv

| Field | Description |
|--|--|
| id | Unique row identifier |
| timestamp | Collection time in UNIX epoch seconds |
| server_id | Unique server identifier (UUID) |
| vcpu_usage | vCPU usage (%) |
| ram_usage | Memory (RAM) usage (%) |
| host_id | Compute host identifier |

#### flavors.csv

| Field | Description |
|--|--|
| flavor_id | Unique flavor identifier (UUID) |
| flavor_name | Name of the flavor |
| vcpu | Number of vCPUs defined in the flavor |
| ram | RAM capacity defined in the flavor (MB) |
| disk | Disk capacity defined in the flavor (GB) |

### Validation

The dataset corresponds to the data archived on Mendeley Data with this repository providing supporting scripts and documentation.
