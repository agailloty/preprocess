---
weight: 300
title: "Command Reference"
description: "Complete guide to all preprocess CLI commands"
icon: "folder"
date: "2025-05-13T19:55:40+02:00"
lastmod: "2026-09-18T11:30:00+02:00"
draft: false
--- 

## Command Reference

This section contains detailed documentation for all **preprocess CLI commands**. Each command is designed to help Data Scientists and Data Analysts efficiently preprocess and understand their datasets.

## Available Commands

{{% table "table-hover" %}}
| Command | Purpose | When to Use |
|---------|---------|-------------|
| [`preprocess init`]({{% relref "init" %}}) | Generate a Prepfile configuration | Starting a new preprocessing project |
| [`preprocess run`]({{% relref "run" %}}) | Execute preprocessing operations | Running your preprocessing pipeline |
| [`preprocess summary`]({{% relref "summary" %}}) | Generate comprehensive statistics | Exploring data characteristics |
| [`preprocess skim`]({{% relref "skim" %}}) | Quickly preview dataset | Initial data inspection |
| [`preprocess diff`]({{% relref "diff" %}}) | Compare two datasets | Verifying preprocessing results |
| [`preprocess version`]({{% relref "versions" %}}) | Display CLI version | Checking installation |
{{% /table %}}

## Command Categories

### Configuration Commands

**Commands for setting up and managing preprocessing configurations:**

- **[`preprocess init`]({{% relref "init" %}})**: Create Prepfile configurations for your datasets
- **[`preprocess version`]({{% relref "versions" %}})**: Check your CLI version

### Execution Commands

**Commands for running preprocessing:**

- **[`preprocess run`]({{% relref "run" %}})**: Execute your preprocessing pipeline

### Exploration Commands

**Commands for understanding your data:**

- **[`preprocess skim`]({{% relref "skim" %}})**: Quick preview of your dataset
- **[`preprocess summary`]({{% relref "summary" %}})**: Generate detailed statistics

### Verification Commands

**Commands for verifying results:**

- **[`preprocess diff`]({{% relref "diff" %}})**: Compare datasets to see changes

## Quick Reference Guide

### Common Workflows

#### Workflow 1: Basic Preprocessing

```bash
# 1. Explore data
preprocess skim --data ./data.csv
preprocess summary --data ./data.csv --html

# 2. Create Prepfile
preprocess init --data ./data.csv --output my_pipeline.toml

# 3. Edit Prepfile (add operations)

# 4. Run preprocessing
preprocess run --file my_pipeline.toml

# 5. Verify results
preprocess diff --source ./data.csv --target ./data_cleaned.csv
```

#### Workflow 2: Data Quality Assessment

```bash
# Quick check
preprocess skim --data ./new_data.csv

# Detailed statistics
preprocess summary --data ./new_data.csv --html --output quality_report.html

# Identify issues and plan preprocessing
```

#### Workflow 3: Automated Processing

```bash
#!/bin/bash
# process_data.sh

# Process all CSV files in a directory
for file in *.csv; do
  preprocess init --data $file --output configs/${file%.csv}_prep.toml
  # Edit the Prepfile as needed
  preprocess run --file configs/${file%.csv}_prep.toml
  preprocess summary --data ${file%.csv}_cleaned.csv --html
done
```

### Command Comparison

{{% table "table-hover" %}}
| Task | Command | Output | Best For |
|------|---------|--------|----------|
| Quick look | `preprocess skim` | Console output | Initial inspection |
| Detailed stats | `preprocess summary` | TOML/HTML file | Data exploration |
| Create config | `preprocess init` | Prepfile.toml | Starting project |
| Run pipeline | `preprocess run` | Processed data | Preprocessing |
| Compare data | `preprocess diff` | HTML report | Verification |
| Check version | `preprocess version` | Version number | Troubleshooting |
{{% /table %}}

## Learning Path

### For Beginners

1. Start with **[`preprocess init`]({{% relref "init" %}})** to understand how to create configurations
2. Learn **[`preprocess skim`]({{% relref "skim" %}})** and **[`preprocess summary`]({{% relref "summary" %}})** for data exploration
3. Master **[`preprocess run`]({{% relref "run" %}})** to execute preprocessing

### For Intermediate Users

1. Use **[`preprocess diff`]({{% relref "diff" %}})** to verify preprocessing results
2. Combine commands for efficient workflows
3. Automate common tasks with scripts

### For Advanced Users

1. Chain multiple Prepfiles for complex pipelines
2. Integrate with CI/CD workflows
3. Create custom scripts combining multiple commands

## Tips for All Users

### Always Start with Exploration

Before any preprocessing, explore your data:

```bash
preprocess skim --data ./your_data.csv
preprocess summary --data ./your_data.csv --html
```

### Use Prepfiles for Reproducibility

Always create and use Prepfiles:

```bash
preprocess init --data ./data.csv --output my_pipeline.toml
# Edit the Prepfile
preprocess run --file my_pipeline.toml
```

### Verify Your Results

Always check your preprocessing worked as expected:

```bash
preprocess diff --source ./original.csv --target ./cleaned.csv
```

### Automate Repetitive Tasks

Create scripts for common workflows:

```bash
#!/bin/bash
# Clean and process all CSV files
for file in *.csv; do
  preprocess run --file config.toml --data $file
  mv output.csv ${file%.csv}_cleaned.csv
done
```

## Next Steps

- **[Prepfile Reference]({{% relref "prepreference" %}})**: Learn about configuration options
- **[Examples]({{% relref "examples" %}})**: See practical use cases
- **[Templates]({{% relref "templates" %}})**: Use pre-built configurations
