---
weight: 200
title: "preprocess run"
icon: "article"
date: "2025-05-13T19:56:40+02:00"
lastmod: "2026-09-18T11:30:00+02:00"
description: "Execute your preprocessing pipeline"
draft: false
toc: true
--- 

## Overview

The `preprocess run` command is the **workhorse** of the preprocess CLI. It executes all the preprocessing operations defined in your Prepfile and transforms your raw data into cleaned, ready-to-use datasets.

Whether you're preparing data for machine learning, analysis, or reporting, `preprocess run` applies all your specified transformations efficiently.

## Basic Usage

### Quick Start

Run preprocessing using your Prepfile:

```bash
preprocess run --file Prepfile.toml
```

If you have a Prepfile in your current directory, you can simply run:

```bash
preprocess run
```

The command will automatically look for `Prepfile.toml` in the current directory.

### Command Reference

```bash
preprocess run [flags]
```

| Flag | Shorthand | Default | Description |
|------|-----------|---------|-------------|
| `--file` | `-f` | `Prepfile.toml` | Path to the configuration file |
| `--data` | `-d` | | Path to the dataset (overrides Prepfile) |
| `--sep` | `-s` | `,` | CSV separator (overrides Prepfile) |
| `--dsep` | `-m` | `.` | Decimal separator (overrides Prepfile) |
| `--column` | | | Target specific column(s) for preprocessing |
| `--op` | | | Preprocessing operation(s) to apply |
| `--numerics` | | | Apply operations only on numeric columns |
| `--show-diff` | | | Show a summary of differences in HTML page |

## How It Works

The `run` command follows this workflow:

1. **Read Configuration**: Loads your Prepfile to understand what operations to perform
2. **Load Data**: Reads your dataset using the specifications from the `[data]` section
3. **Apply Operations**: Executes all preprocessing operations in the order they're defined
4. **Postprocess**: Applies any postprocessing steps (dropping columns, sorting, etc.)
5. **Export Results**: Saves the cleaned data to the specified output file

## Practical Examples

### Example 1: Basic Execution

You have a Prepfile in your current directory:

```bash
preprocess run
```

This will:
- Find `Prepfile.toml` automatically
- Read your dataset
- Apply all preprocessing operations
- Save the cleaned data to the file specified in the `[postprocess]` section

### Example 2: Specify Prepfile Path

Your Prepfile is located elsewhere:

```bash
preprocess run --file ./config/my_preprocessing.toml
```

### Example 3: Quick Processing Without Prepfile

For simple, one-off operations, you can specify everything via command line:

```bash
preprocess run --data ./sales.csv --column revenue --op fillna:method=mean
```

This fills missing values in the `revenue` column with the mean.

### Example 4: Apply Multiple Operations

```bash
preprocess run --data ./data.csv \
  --column age --op fillna:method=median \
  --column salary --op fillna:method=mean \
  --column salary --op scale:method=zscore
```

**Note**: Operations are applied in the order they appear on the command line.

### Example 5: Process All Numeric Columns

Apply operations to all numeric columns without specifying each one:

```bash
preprocess run --data ./data.csv --numerics --op fillna:method=mean
```

This fills missing values in all numeric columns with their respective means.

### Example 6: Show Differences

Want to see what changed? Use the `--show-diff` flag:

```bash
preprocess run --file my_prep.toml --show-diff
```

This opens a HTML page in your browser showing:
- Original vs. cleaned data comparison
- Statistics on changes made
- Visual indicators of modifications

## Real-World Scenarios

### Scenario 1: Machine Learning Pipeline

You're preparing data for a machine learning model:

```bash
# First, create your Prepfile
preprocess init --data ./training_data.csv --output ml_prep.toml

# Edit ml_prep.toml to include:
# - Missing value imputation
# - Feature scaling
# - Categorical encoding

# Then run the preprocessing
preprocess run --file ml_prep.toml
```

Your Prepfile might look like:

```toml
[data]
filename = './training_data.csv'
csv_separator = ','
decimal_separator = '.'
encoding = 'utf-8'

[preprocess]
# Handle missing values in all numeric columns
[preprocess.numerics]
operations = [
    {op = "fillna", method = "median"}
]

# Scale all numeric features
[preprocess.numerics]
operations = [
    {op = "scale", method = "zscore"}
]

# Encode categorical variables
[[preprocess.columns]]
name = 'category'
type = 'string'
operations = [
    {op = "dummy"}
]

[postprocess]
format = 'csv'
filename = 'training_data_cleaned.csv'
```

### Scenario 2: Data Cleaning for Analysis

You need to clean survey data before analysis:

```bash
preprocess init --data ./survey.csv --output survey_clean.toml
```

Edit to add:
- Text cleaning (trim whitespace, standardize case)
- Missing value handling
- Column renaming

Then run:

```bash
preprocess run --file survey_clean.toml
```

### Scenario 3: Automated Data Ingestion

Set up a script to process new data files as they arrive:

```bash
#!/bin/bash
# process_new_data.sh

# Generate timestamp
TIMESTAMP=$(date +%Y%m%d_%H%M%S)

# Run preprocessing
preprocess run --file ./config/standard_prep.toml \
  --data ./new_data/incoming_$TIMESTAMP.csv \
  --output ./processed/cleaned_$TIMESTAMP.csv

# Log the operation
echo "Processed incoming_$TIMESTAMP.csv at $(date)" >> ./logs/processing.log
```

Make it executable:
```bash
chmod +x process_new_data.sh
```

## Advanced Usage

### Chaining Multiple Prepfiles

For complex pipelines, you can chain multiple Prepfiles:

```bash
# Step 1: Initial cleaning
preprocess run --file step1_cleaning.toml

# Step 2: Feature engineering
preprocess run --file step2_features.toml --data ./output_from_step1.csv

# Step 3: Final preparation
preprocess run --file step3_final.toml --data ./output_from_step2.csv
```

### Using with Different Data Formats

While preprocess works primarily with CSV files, you can process various formats:

```bash
# Tab-separated values
preprocess run --data ./data.tsv --sep $'\t'

# Semicolon-separated (common in European data)
preprocess run --data ./data.csv --sep ";" --dsep ","

# Pipe-separated
preprocess run --data ./data.txt --sep "|"
```

### Handling Large Datasets

For large datasets, preprocess is optimized for performance:

```bash
# Process a large file
preprocess run --file large_data_prep.toml
```

**Tips for large datasets:**
- Use `--show-diff` to verify changes without manually inspecting the entire dataset
- Process in batches if needed by creating multiple Prepfiles
- Ensure you have enough memory available

## Understanding Operation Syntax

### Single Operation Format

```bash
--op operation:method=value
```

Examples:
- `--op fillna:method=mean` - Fill missing values with mean
- `--op fillna:method=median` - Fill missing values with median
- `--op scale:method=zscore` - Standardize (z-score normalization)
- `--op scale:method=minmax` - Min-max normalization
- `--op clean:method=trimws` - Trim whitespace
- `--op clean:method=lower` - Convert to lowercase

### Multiple Operations on Same Column

```bash
--column revenue \
  --op fillna:method=mean \
  --op scale:method=zscore
```

This applies operations sequentially: first fill missing values, then scale.

### Operation Methods Reference

#### Fill Missing Values (`fillna`)

| Method | Applies To | Description |
|--------|------------|-------------|
| `mean` | Numeric | Replace with column mean |
| `median` | Numeric | Replace with column median |
| `value:<number>` | Numeric | Replace with specified value |
| `value:<string>` | Text | Replace with specified string |

#### Scaling (`scale`)

| Method | Description | Formula |
|--------|-------------|---------|
| `zscore` | Standardization | (x - mean) / std |
| `minmax` | Normalization | (x - min) / (max - min) |

#### Text Cleaning (`clean`)

| Method | Description |
|--------|-------------|
| `trimws` | Remove leading and trailing whitespace |
| `upper` | Convert to uppercase |
| `lower` | Convert to lowercase |
| `title` | Convert to title case |

#### Other Operations

| Operation | Method | Description |
|-----------|--------|-------------|
| `discretize` | `binning` | Create bins for continuous variables |
| `dummy` | | Create dummy variables from categorical |
| `group` | | Group operations |
| `rename` | | Rename columns |

## Best Practices

### 1. Start with a Prepfile

While you can use command-line flags for simple operations, for anything more complex:

```bash
# Do this:
preprocess init --data ./data.csv --output my_config.toml
# Edit my_config.toml
preprocess run --file my_config.toml

# Instead of this (for complex operations):
preprocess run --data ./data.csv --column col1 --op op1 --column col2 --op op2 ...
```

### 2. Use `--show-diff` for Verification

Always verify your preprocessing with `--show-diff`:

```bash
preprocess run --file my_prep.toml --show-diff
```

This helps you:
- Confirm operations are working as expected
- Catch configuration errors early
- Understand the impact of your preprocessing

### 3. Test Incrementally

Add operations one at a time and verify each step:

```bash
# Step 1: Add fillna for one column
# Test with --show-diff

# Step 2: Add scaling
# Test again

# Step 3: Add more operations
# Verify
```

### 4. Document Your Pipeline

Add comments to your Prepfile explaining each operation:

```toml
[preprocess]
# Data cleaning phase
[preprocess.numerics]
operations = [
    {op = "fillna", method = "median"}  # Handle missing values first
]

# Feature scaling for ML
[preprocess.numerics]
operations = [
    {op = "scale", method = "zscore"}  # Standardize for algorithm compatibility
]
```

## Troubleshooting

### Common Issues

**Issue: Prepfile not found**
```
Error: Failed to load config file 'Prepfile.toml'
```

**Solution**: Specify the path explicitly:
```bash
preprocess run --file ./path/to/your/Prepfile.toml
```

**Issue: Dataset not found**
```
Error: Failed to read dataset
```

**Solution**: Verify the path in your Prepfile or use `--data` flag:
```bash
preprocess run --file my_prep.toml --data ./correct/path/to/data.csv
```

**Issue: Column not found**
```
Error: Column 'column_name' not found in dataset
```

**Solution**: Check for typos in your column names. Remember:
- Column names are case-sensitive
- Verify column names with `preprocess skim --data ./your_data.csv`

**Issue: Type mismatch**
```
Error: Operation not applicable to column type
```

**Solution**: Ensure you're applying the right operations to the right column types:
- Numeric operations (`fillna:mean`, `scale`) only work on numeric columns
- Text operations (`clean`, `dummy`) only work on text columns

### Debugging Tips

1. **Verify your data first**:
   ```bash
   preprocess skim --data ./your_data.csv
   ```

2. **Check your Prepfile syntax**:
   - Ensure proper TOML formatting
   - Verify all required sections are present

3. **Test with minimal operations**:
   ```bash
   # Test with just one simple operation
   preprocess run --data ./data.csv --column col1 --op fillna:method=mean
   ```

4. **Use `--show-diff` to see what's happening**:
   ```bash
   preprocess run --file my_prep.toml --show-diff
   ```

## Performance Tips

### Optimizing Processing Time

1. **Exclude unnecessary columns**:
   ```toml
   [preprocess.numerics]
   exclude_columns = ["id", "timestamp", "unnecessary_col"]
   ```

2. **Process in the right order**:
   - Apply operations that reduce data size first (dropping columns)
   - Then apply transformations
   - Finally apply scaling

3. **Use appropriate data types**:
   - Ensure numeric columns are detected as `int` or `float`, not `string`

### Memory Usage

For very large datasets:
- Process in batches if possible
- Close other memory-intensive applications
- Consider using `--show-diff` instead of opening the full dataset

## Next Steps

After running preprocessing:

1. **[Verify with summary]({{% relref "summary" %}})**: Generate statistics to confirm your data is clean
2. **[Explore the output]({{% relref "skim" %}})**: Use `preprocess skim` on your cleaned data
3. **[Compare with original]({{% relref "diff" %}})**: Use `preprocess diff` to compare original and cleaned data
4. **[Iterate on your Prepfile]({{% relref "prepreference" %}})**: Refine your preprocessing pipeline

## Quick Reference Card

| Task | Command |
|------|---------|
| Run with default Prepfile | `preprocess run` |
| Run with specific Prepfile | `preprocess run --file my_prep.toml` |
| Run with data override | `preprocess run --data ./new_data.csv` |
| Apply operation to column | `preprocess run --data ./data.csv --column age --op fillna:method=mean` |
| Apply to all numeric columns | `preprocess run --data ./data.csv --numerics --op scale:method=zscore` |
| Show changes | `preprocess run --file my_prep.toml --show-diff` |
| Chain multiple operations | `preprocess run --column col1 --op op1 --op op2` |

## Supported Operations Summary

| Operation | Methods | Applies To | Use Case |
|-----------|---------|------------|----------|
| `fillna` | mean, median, value | Numeric | Handle missing values |
| `fillna` | value | Text | Replace missing text |
| `scale` | zscore, minmax | Numeric | Feature scaling |
| `clean` | trimws, upper, lower, title | Text | Text normalization |
| `discretize` | binning | Numeric | Create categories from continuous data |
| `dummy` | | Text | Create one-hot encoding |
| `group` | | Both | Group operations |
| `rename` | | Both | Rename columns |
