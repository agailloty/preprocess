---
weight: 1
title: "Preprocess CLI Documentation"
icon: "home"
date: "2025-05-13T13:15:40+02:00"
lastmod: "2026-09-18T11:30:00+02:00"
draft: false
description: "Complete documentation for preprocess CLI - A fast data preprocessing tool for Data Scientists and Data Analysts"
publishdate: "2025-05-13T13:15:40+02:00"
--- 

## Welcome to Preprocess CLI

**preprocess** is a fast, cross-platform command-line tool designed specifically for **Data Scientists and Data Analysts** to preprocess datasets efficiently. Built with Go for performance, preprocess helps you clean, transform, and prepare your data for analysis and machine learning.

## What is Preprocess CLI?

Preprocess CLI is a **declarative data preprocessing tool** that allows you to:

- **Define preprocessing pipelines** using simple configuration files (Prepfiles)
- **Apply common data transformations** with a single command
- **Automate repetitive data cleaning tasks**
- **Ensure reproducibility** across your data projects
- **Collaborate with team members** by sharing preprocessing configurations

## Key Features

{{% table "table-hover" %}}
| Feature | Description | Benefit |
|---------|-------------|---------|
| **Declarative Configuration** | Define preprocessing in TOML files | Easy to read, write, and version control |
| **Multiple Operations** | Fill missing values, scale, clean text, encode categories, and more | Comprehensive data preparation |
| **Flexible Application** | Apply operations to specific columns, all numeric, or all text columns | Precisely control your preprocessing |
| **Batch Processing** | Process entire datasets with one command | Fast and efficient |
| **Data Exploration** | Built-in commands for exploring your data | Understand your data before preprocessing |
| **Visual Reports** | Generate HTML reports for data comparison and summary | Easy to share and interpret |
| **Cross-Platform** | Works on Windows, macOS, and Linux | Use your preferred operating system |
{{% /table %}}

## Getting Started in 5 Minutes

### Step 1: Install Preprocess CLI

Choose your operating system:

{{< tabs tabTotal="3">}}
{{% tab tabName="Linux/Mac" %}}

Use curl to install:

```bash
curl -LsSf https://preprocess-cli.netlify.app/install.sh | sh
```

Or with wget:

```bash
wget -qO- https://preprocess-cli.netlify.app/install.sh | sh
```

{{% /tab %}}
{{% tab tabName="Windows" %}}

Use PowerShell:

```powershell
powershell -ExecutionPolicy ByPass -c "irm https://preprocess-cli.netlify.app/install.ps1 | iex"
```

{{% /tab %}}
{{% tab tabName="Manual" %}}

Download the latest release for your OS from:
[GitHub Releases](https://github.com/agailloty/preprocess/releases/latest)

{{% /tab %}}
{{< /tabs >}}

### Step 2: Verify Installation


```bash
preprocess version
```

You should see the version number (e.g., `preprocess version 0.1.1`)

### Step 3: Explore Your Data

Download a sample dataset and explore it:

```bash
# Skim through the data
preprocess skim --data ./your_dataset.csv

# Generate detailed statistics
preprocess summary --data ./your_dataset.csv --html
```

### Step 4: Create Your First Prepfile

Generate a configuration file for your dataset:

```bash
preprocess init --data ./your_dataset.csv --output my_pipeline.toml
```

This creates a Prepfile that detects all columns in your dataset.

### Step 5: Add Preprocessing Operations

Edit the generated Prepfile (`my_pipeline.toml`) to add preprocessing operations. For example:

```toml
[preprocess]
# Fill missing values in all numeric columns
[preprocess.numerics]
operations = [
    {op = "fillna", method = "median"}
]

# Clean all text columns
[preprocess.texts]
operations = [
    {op = "clean", method = "trimws"}
]

# Scale numeric features
[preprocess.numerics]
operations = [
    {op = "scale", method = "zscore"}
]
```

### Step 6: Run Your Pipeline

Execute your preprocessing:

```bash
preprocess run --file my_pipeline.toml
```

This creates a cleaned dataset as specified in the `[postprocess]` section.

### Step 7: Verify Results

Compare your original and cleaned data:

```bash
preprocess diff --source ./your_dataset.csv --target ./your_dataset_cleaned.csv
```

This opens a visual HTML report showing all changes.

## Who Should Use Preprocess CLI?

### Data Scientists

Use preprocess CLI to:
- **Prepare data for machine learning** (fill missing values, scale features, encode categories)
- **Automate repetitive data cleaning** tasks
- **Ensure reproducible preprocessing** across projects
- **Document data transformations** for team members

### Data Analysts

Use preprocess CLI to:
- **Clean messy datasets** before analysis
- **Standardize data formats** across multiple sources
- **Generate summary statistics** for data exploration
- **Create reusable data preparation** workflows

### Data Engineers

Use preprocess CLI to:
- **Build data preprocessing pipelines**
- **Integrate with ETL workflows**
- **Standardize data across the organization**
- **Automate data quality checks**

## What Makes Preprocess CLI Special?

### 1. Performance

Built with **Go**, preprocess is **blazing fast**:
- Processes large datasets efficiently
- Optimized for performance
- Minimal memory overhead

### 2. Simplicity

**No programming required**: Define your preprocessing in simple TOML files

```toml
# Instead of writing code like this:
# df.fillna(df.mean())
# df['column'] = df['column'].str.strip()

# You write this:
[preprocess.numerics]
operations = [
    {op = "fillna", method = "mean"}
]

[preprocess.texts]
operations = [
    {op = "clean", method = "trimws"}
]
```

### 3. Reproducibility

Same Prepfile + same data = same results, every time.

```bash
# Share your preprocessing with team members
# They can run the exact same pipeline
preprocess run --file shared_pipeline.toml
```

### 4. Integration

Works seamlessly with your existing workflow:
- **Version control friendly** (Prepfiles are text files)
- **Scriptable** (easy to integrate into shell scripts)
- **Automatable** (perfect for CI/CD pipelines)

## Use Cases

### Machine Learning Data Preparation

```toml
[data]
filename = './training_data.csv'

[preprocess]
[preprocess.numerics]
operations = [
    {op = "fillna", method = "median"},
    {op = "scale", method = "zscore"}
]

[preprocess.texts]
operations = [
    {op = "clean", method = "trimws"}
]

[[preprocess.columns]]
name = 'category'
type = 'string'
operations = [
    {op = "dummy", dummy_droplast = true}
]

[postprocess]
filename = 'training_data_processed.csv'

[postprocess.dataset_split]
method = "train_test_split"
split_names = ["train", "test"]
random_seed = 42
train_test_split_ratio = 0.8
```

### Data Cleaning Pipeline

```toml
[data]
filename = './raw_survey_data.csv'
missing_identifier = 'N/A'

[preprocess]
[preprocess.numerics]
operations = [
    {op = "fillna", method = "mean"}
]

[preprocess.texts]
operations = [
    {op = "clean", method = "trimws"},
    {op = "clean", method = "title"}
]

[[preprocess.columns]]
name = 'age'
type = 'int'
operations = [
    {op = "fillna", value = 0}
]

[postprocess]
filename = 'survey_data_cleaned.csv'
dropcolumns = ["internal_id", "notes"]
```

### Data Migration

```toml
[data]
filename = './legacy_export.txt'
csv_separator = '|'
encoding = 'latin-1'

[preprocess]
[preprocess.texts]
operations = [
    {op = "clean", method = "trimws"},
    {op = "clean", method = "upper"}
]

[[preprocess.columns]]
name = 'cust_id'
new_name = 'customer_id'

[postprocess]
filename = 'migrated_data.csv'
```

## Command Overview

Preprocess CLI provides several commands for data preprocessing:

{{% table "table-hover" %}}
| Command | Description | Use Case |
|---------|-------------|----------|
| [`preprocess init`]({{% relref "init" %}}) | Generate a Prepfile | Start a new preprocessing project |
| [`preprocess run`]({{% relref "run" %}}) | Run preprocessing operations | Execute your preprocessing pipeline |
| [`preprocess summary`]({{% relref "summary" %}}) | Generate summary statistics | Explore and understand your data |
| [`preprocess skim`]({{% relref "skim" %}}) | Preview dataset | Quick data inspection |
| [`preprocess diff`]({{% relref "diff" %}}) | Compare datasets | Verify preprocessing results |
| [`preprocess version`]({{% relref "versions" %}}) | Display version | Check your installation |
{{% /table %}}

## Learning Path

### For Beginners

1. **[Install preprocess]({{% relref "quickstart" %}})**
2. **[Explore the commands]({{% relref "guides" %}})**
3. **[Learn about Prepfile]({{% relref "prepreference" %}})** - The heart of preprocess
4. **[Try the examples]({{% relref "examples" %}})**

### For Intermediate Users

1. **[Understand operations]({{% relref "operations" %}})** - All available preprocessing operations
2. **[Use templates]({{% relref "templates" %}})** - Pre-built configurations for common tasks
3. **[Explore examples]({{% relref "examples" %}})** - Real-world use cases

### For Advanced Users

1. **[Create complex pipelines]** - Combine multiple Prepfiles
2. **[Automate workflows]** - Integrate with scripts and CI/CD
3. **[Share with team]** - Collaborative preprocessing

## Community & Support

### Documentation

You're reading it! All documentation is available at:
- [Quickstart Guide]({{% relref "quickstart" %}})
- [Command Reference]({{% relref "guides" %}})
- [Prepfile Reference]({{% relref "prepreference" %}})
- [Examples]({{% relref "examples" %}})
- [Templates]({{% relref "templates" %}})

### GitHub

- **Repository**: [https://github.com/agailloty/preprocess](https://github.com/agailloty/preprocess)
- **Issues**: Report bugs and request features
- **Releases**: Check for the latest version

### Get Help

1. **Read the documentation** - Most questions are answered here
2. **Check the examples** - Real-world use cases are provided
3. **Review the Prepfile reference** - All configuration options are documented
4. **Report issues on GitHub** - For bugs and feature requests

## Best Practices

### 1. Start Simple

Begin with basic operations and add complexity gradually.

```toml
# Start with just the essentials
[data]
filename = './data.csv'

[preprocess]

[postprocess]

# Then add operations one at a time
```

### 2. Use Version Control

Track changes to your Prepfiles:

```bash
git add my_pipeline.toml
git commit -m "Add feature scaling"
```

### 3. Document Your Pipeline

Add comments to your Prepfile:

```toml
[preprocess]
# Data cleaning phase
[preprocess.numerics]
operations = [
    {op = "fillna", method = "median"}  # Using median for robustness
]
```

### 4. Verify Results

Always check your preprocessing worked as expected:

```bash
preprocess run --file my_pipeline.toml
preprocess diff --source ./original.csv --target ./cleaned.csv
```

### 5. Test Incrementally

Add operations one at a time and verify each step:

```bash
# Test with one operation
preprocess run --file pipeline_v1.toml
preprocess skim --data ./output.csv

# Add another operation
# Edit to pipeline_v2.toml
preprocess run --file pipeline_v2.toml
preprocess skim --data ./output.csv
```

### 6. Share with Team

Prepfiles are text files - easy to share:

```bash
# Share with colleague
cp my_pipeline.toml ../team_member/project/
```

## Common Questions

### How does preprocess compare to Python/R?

| Feature | Preprocess CLI | Python (Pandas) | R (dplyr) |
|---------|----------------|---------------|----------|
| **Learning Curve** | Low - declarative | Medium - requires coding | Medium - requires coding |
| **Performance** | Very Fast (Go) | Fast | Medium |
| **Reproducibility** | High - same Prepfile = same results | High - with script | High - with script |
| **Version Control** | Easy - text files | Requires script management | Requires script management |
| **Sharing** | Very Easy - share Prepfile | Moderate - share scripts | Moderate - share scripts |
| **Integration** | CLI - easy scripting | Library - flexible | Library - flexible |

**Use preprocess CLI when you want:**
- Simple, declarative preprocessing
- Fast performance
- Easy sharing and collaboration
- Reproducible pipelines

**Use Python/R when you need:**
- Complex custom transformations
- Integration with analysis code
- Interactive exploration

### Can I use preprocess with Python/R?

Absolutely! Preprocess CLI is designed to **complement** your existing workflows:

```python
# In Python
import subprocess

# Run preprocess from Python
subprocess.run(["preprocess", "run", "--file", "pipeline.toml"])

# Then continue with pandas
import pandas as pd
df = pd.read_csv("output_cleaned.csv")
```

```r
# In R
system("preprocess run --file pipeline.toml")

# Then continue with R
library(readr)
data <- read_csv("output_cleaned.csv")
```

### How do I handle large datasets?

Preprocess is optimized for performance:
- Built with Go for speed
- Efficient memory usage
- For very large datasets:
  - Process in batches
  - Use `--show-diff` to verify without opening full data
  - Ensure you have enough memory available

### Can I customize operations?

The current operations cover most common preprocessing needs. If you need custom operations:
1. **Check existing operations** - You might find what you need already exists
2. **Combine operations** - Many complex transformations can be achieved by combining existing operations
3. **Request new operations** - Open an issue on GitHub with your use case

### How do I debug my Prepfile?

1. **Start simple** - Begin with minimal configuration
2. **Add incrementally** - Add one section at a time
3. **Use `--show-diff`** - Visual verification of results
4. **Check error messages** - They often provide clear guidance
5. **Validate TOML syntax** - Use an online TOML validator

## Next Steps

Ready to dive in?

1. **[Install preprocess]({{% relref "quickstart" %}})** - Get started in minutes
2. **[Learn the commands]({{% relref "guides" %}})** - Explore all capabilities
3. **[Understand Prepfile]({{% relref "prepreference" %}})** - Master the configuration
4. **[Try examples]({{% relref "examples" %}})** - See real-world use cases

## Summary

Preprocess CLI is a **powerful, simple, and fast** tool for data preprocessing. It's designed specifically for **Data Scientists and Data Analysts** who want to:

- **Save time** on repetitive data cleaning tasks
- **Ensure reproducibility** in their data processing
- **Collaborate easily** with team members
- **Integrate seamlessly** with existing workflows

Whether you're preparing data for machine learning, analysis, or reporting, preprocess CLI provides the tools you need to clean and transform your data efficiently.

---

**Get started today and transform the way you preprocess data!**
