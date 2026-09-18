---
weight: 100
title: "preprocess init"
icon: "article"
date: "2025-05-13T19:56:40+02:00"
lastmod: "2026-09-18T11:30:00+02:00"
description: "Generate a Prepfile to declare your preprocessing pipeline"
draft: false
toc: true
--- 

## Overview

The `preprocess init` command is your starting point for using **preprocess CLI**. It generates a **Prepfile.toml** - a configuration file where you declare all the preprocessing operations you want to apply to your dataset.

Think of the Prepfile as your **data preprocessing recipe**: it contains everything preprocess needs to know about your data and what transformations to apply.

## Why Use a Prepfile?

For Data Scientists and Data Analysts:

- **Reproducibility**: Your preprocessing pipeline is documented and can be reused
- **Version Control**: Track changes to your data processing logic alongside your code
- **Collaboration**: Share your preprocessing configuration with team members
- **Automation**: Run the same preprocessing steps on new data without rewriting code

## Basic Usage

### Quick Start

Generate a Prepfile for your dataset with a single command:

```bash
preprocess init --data ./my_dataset.csv
```

This creates a `Prepfile.toml` in your current directory that automatically:
- Detects all columns in your dataset
- Identifies column types (numeric, text)
- Prepares the structure for adding preprocessing operations

### Command Reference

```bash
preprocess init [flags]
```

| Flag | Shorthand | Default | Description |
|------|-----------|---------|-------------|
| `--data` | `-d` | | **Required**: Path to your dataset file |
| `--sep` | `-s` | `,` | CSV separator character |
| `--dsep` | `-m` | `.` | Decimal separator (use `,` for European formats) |
| `--encoding` | `-e` | `utf-8` | File encoding (supports `utf-8`, `latin-1`) |
| `--output` | `-o` | `Prepfile.toml` | Output filename for the configuration |
| `--template` | `-t` | | Generate an example template with common operations |

## Practical Examples

### Example 1: Basic Prepfile Generation

You have a CSV file named `sales_data.csv` with standard formatting:

```bash
preprocess init --data ./sales_data.csv
```

This creates a Prepfile that looks like:

```toml
[data]
filename = './sales_data.csv'
csv_separator = ','
decimal_separator = '.'
encoding = 'utf-8'
missing_identifier = ''

[preprocess]
[[preprocess.columns]]
name = 'customer_id'
type = 'int'

[[preprocess.columns]]
name = 'purchase_amount'
type = 'float'

[[preprocess.columns]]
name = 'product_category'
type = 'string'

[postprocess]
format = 'csv'
filename = 'sales_data_cleaned.csv'
```

### Example 2: European Format Dataset

Your data uses semicolons as separators and commas for decimals:

```bash
preprocess init --data ./european_data.csv --sep ";" --dsep ","
```

Generated Prepfile:

```toml
[data]
filename = './european_data.csv'
csv_separator = ';'
decimal_separator = ','
encoding = 'utf-8'
missing_identifier = ''
```

### Example 3: Generate a Template with Common Operations

Want to see examples of all available operations? Use the `--template` flag:

```bash
preprocess init --data ./my_data.csv --template
```

This generates a comprehensive Prepfile with commented examples of:
- Missing value handling
- Scaling operations
- Text cleaning
- Column renaming
- And more...

### Example 4: Save to a Specific Location

```bash
preprocess init --data ./data/transactions.csv --output ./config/transactions_prep.toml
```

## Understanding the Generated Prepfile

### Data Section `[data]`

This section tells preprocess **how to read your dataset**:

```toml
[data]
filename = './my_dataset.csv'
csv_separator = ','
decimal_separator = '.'
encoding = 'utf-8'
missing_identifier = ''
```

| Parameter | Purpose | Common Values |
|-----------|---------|---------------|
| `filename` | Path to your dataset | Any valid file path |
| `csv_separator` | Column delimiter | `,`, `;`, `\t`, `|` |
| `decimal_separator` | Decimal point character | `.` or `,` |
| `encoding` | Text encoding | `utf-8`, `latin-1`, `iso-8859-1` |
| `missing_identifier` | How missing values are represented | `''`, `'NA'`, `'N/A'`, `'null'` |

**Pro Tip**: The `missing_identifier` is crucial! If your dataset uses `NA` or `N/A` for missing values, specify it here so preprocess can detect them correctly.

### Preprocess Section `[preprocess]`

This is where you **define your preprocessing operations**. The generated Prepfile lists all your columns:

```toml
[preprocess]
[[preprocess.columns]]
name = 'customer_id'
type = 'int'

[[preprocess.columns]]
name = 'age'
type = 'int'
operations = [
    {op = "fillna", method = "mean"}
]

[[preprocess.columns]]
name = 'category'
type = 'string'
operations = [
    {op = "clean", method = "trimws"},
    {op = "clean", method = "lower"}
]
```

You have **three ways** to apply operations:

1. **On specific columns** (as shown above)
2. **On all numeric columns** using `[preprocess.numerics]`
3. **On all text columns** using `[preprocess.texts]`

### Postprocess Section `[postprocess]`

Define what happens **after preprocessing**:

```toml
[postprocess]
format = 'csv'
filename = 'my_dataset_cleaned.csv'
dropcolumns = ["id", "temp_column"]
sortdataset = {descending = false}
```

| Parameter | Purpose | Values |
|-----------|---------|--------|
| `format` | Output file format | `csv` (default) |
| `filename` | Name of the output file | Any valid filename |
| `dropcolumns` | Columns to remove after processing | Array of column names |
| `sortdataset` | Sort order | `{descending = true/false}` |

## Real-World Scenario: Data Cleaning Pipeline

Let's say you're working with customer survey data that needs cleaning:

### Step 1: Initialize

```bash
preprocess init --data ./survey_2024.csv --output survey_cleaning.toml
```

### Step 2: Edit the Prepfile

Modify the generated file to add your preprocessing logic:

```toml
[data]
filename = './survey_2024.csv'
csv_separator = ','
decimal_separator = '.'
encoding = 'utf-8'
missing_identifier = 'N/A'

[preprocess]
# Clean all text columns
[preprocess.texts]
operations = [
    {op = "clean", method = "trimws"},
    {op = "clean", method = "title"}
]
exclude_columns = ["customer_id"]

# Process numeric columns
[preprocess.numerics]
operations = [
    {op = "fillna", method = "median"}
]

# Specific handling for age column
[[preprocess.columns]]
name = 'age'
type = 'int'
operations = [
    {op = "fillna", method = "mean"}
]

[postprocess]
format = 'csv'
filename = 'survey_2024_cleaned.csv'
dropcolumns = ["raw_notes"]
sortdataset = {descending = false}
```

### Step 3: Run the Pipeline

```bash
preprocess run --file survey_cleaning.toml
```

## Common Use Cases

### Use Case 1: Machine Learning Data Preparation

```bash
# Create Prepfile for ML dataset
preprocess init --data ./training_data.csv --output ml_pipeline.toml
```

Edit to add:
- Missing value imputation
- Feature scaling
- Categorical encoding (dummy variables)

### Use Case 2: Data Migration Project

```bash
# Initialize for legacy data
preprocess init --data ./legacy_system_export.csv --sep "\t" --encoding latin-1
```

Edit to add:
- Text cleaning and standardization
- Date parsing (if stored as strings)
- Column renaming to match new system

### Use Case 3: Regular Data Ingestion

```bash
# Create reusable Prepfile for recurring data imports
preprocess init --data ./sample_monthly_data.csv --template
```

Save as `monthly_import.toml` and reuse for each month's data.

## Best Practices

### 1. Start Simple

Begin with a basic Prepfile and add operations incrementally:

```bash
# First, just generate the structure
preprocess init --data ./data.csv

# Test it works
preprocess run

# Then add one operation at a time
```

### 2. Use Meaningful Filenames

Instead of the default `Prepfile.toml`, use descriptive names:

```bash
preprocess init --data ./sales.csv --output sales_data_cleaning.toml
preprocess init --data ./customers.csv --output customer_preprocessing.toml
```

### 3. Keep Prepfile Under Version Control

Add your Prepfile to Git to track changes to your preprocessing logic:

```bash
git add sales_data_cleaning.toml
git commit -m "Update preprocessing: add outlier handling"
```

### 4. Document Your Prepfile

Add comments to explain your preprocessing decisions:

```toml
[preprocess]
# Fill missing ages with mean - approved by data team
[[preprocess.columns]]
name = 'age'
type = 'int'
operations = [
    {op = "fillna", method = "mean"}  # Using mean per business rules
]

# Standardize all text fields for consistency
[preprocess.texts]
operations = [
    {op = "clean", method = "trimws"},  # Remove leading/trailing spaces
    {op = "clean", method = "lower"}   # Lowercase for case-insensitive analysis
]
```

## Troubleshooting

### Common Issues

**Issue: File not found**
```
Error: Failed to read dataset
```

**Solution**: Verify the path is correct. Use absolute paths if needed:
```bash
preprocess init --data /full/path/to/your/dataset.csv
```

**Issue: Encoding problems**
```
Error: Invalid UTF-8 encoding
```

**Solution**: Try different encodings:
```bash
preprocess init --data ./data.csv --encoding latin-1
```

**Issue: Wrong column detection**
```
Columns detected as wrong types
```

**Solution**: Manually edit the Prepfile to correct column types, or check your CSV separator.

### Verifying the Prepfile

After generation, check your Prepfile:

```bash
# Skim your data first to verify it looks correct
preprocess skim --data ./your_dataset.csv

# Then generate the Prepfile
preprocess init --data ./your_dataset.csv

# Verify the generated file
cat Prepfile.toml  # or open in your editor
```

## Next Steps

Once you have your Prepfile:

1. **[Edit the Prepfile]({{% relref "prepreference" %}})**: Learn about all available configuration options
2. **[Run preprocessing]({{% relref "run" %}})**: Execute your preprocessing pipeline
3. **[Generate summaries]({{% relref "summary" %}})**: Create statistics reports for your cleaned data
4. **[Explore examples]({{% relref "examples" %}})**: See practical examples of preprocessing workflows

## Quick Reference Card

| Action | Command |
|--------|---------|
| Generate basic Prepfile | `preprocess init --data ./data.csv` |
| Generate with template | `preprocess init --data ./data.csv --template` |
| European format | `preprocess init --data ./data.csv --sep ";" --dsep ","` |
| Save to specific file | `preprocess init --data ./data.csv --output my_config.toml` |
| Check data first | `preprocess skim --data ./data.csv` |
