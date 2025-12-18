# Reproducing Analysis

This repository contains all R Markdown scripts required to reproduce the figures from the paper "Dataset on Resource Allocation and Usage for a Private Cloud".

## Folder Structure

- **`r-markdown/`**: Contains individual `.Rmd` files for each figure.
  - Files are named matching the paper (e.g., `figure-9.Rmd` for Figure 9).
  - Each file generates one specific figure with complete code and analysis.

## Setup Instructions

1. **Run setup first**: Execute `r-markdown/setup-data.Rmd` before any figure files.
   - Installs all required libraries
   - Reads raw data files
   - Prepares datasets for plotting

2. **Generate figures**: Open and knit individual `figure-X.Rmd` files in order.
   - Each file is self-contained for its respective figure
   - Outputs match exactly those in the paper

## Requirements

- R with RStudio recommended
- All dependencies handled by `setup-data.Rmd`
- 8GB RAM

Run `setup-data.Rmd` -->  knit desired `figure-X.Rmd` files --> reproduce paper figures exactly.
