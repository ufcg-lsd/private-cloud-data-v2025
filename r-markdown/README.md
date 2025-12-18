# Reproducing Analysis

This repository contains all R Markdown scripts required to reproduce exploratory analysis with ["Dataset on Resource Allocation and Usage for a Private Cloud"](https://data.mendeley.com/datasets/trvb5k4x5m/1)

## Setup Instructions

1. **Run setup first**: Execute `r-markdown/setup-data.Rmd` before any figure files.
   - Installs all required libraries
   - Reads raw data files
   - Prepares datasets for plotting

2. **Generate figures**: Open and knit individual `<plot-name>.Rmd` files.
   - Each file is self-contained for its respective figure

## Requirements

- R with RStudio recommended
- All dependencies handled by `setup-data.Rmd`
- 8GB RAM


### Please consider to cite our dataset in your paper! :) 

```
Marques, Paola; Mendes, Mariana; Pereira, Thiago Emmanuel; Farias, Giovanni (2025), “Dataset on Resource Allocation and Usage for a Private Cloud”, Mendeley Data, V1, doi: 10.17632/trvb5k4x5m.1
```
