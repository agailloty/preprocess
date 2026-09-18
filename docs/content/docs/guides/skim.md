---
weight: 450
title: "preprocess skim"
icon: "article"
date: "2025-06-01T19:56:40+02:00"
lastmod: "2026-09-18T11:30:00+02:00"
description: "Quickly preview your dataset structure and content"
draft: false
toc: true
--- 

## Overview

The `preprocess skim` command provides a **quick preview** of your dataset, allowing you to understand its structure and content at a glance. It's the fastest way to get familiar with a new dataset or verify the results of your preprocessing.

For Data Scientists and Analysts, `skim` is your go-to command for:
- Quick data inspection
- Initial data quality checks
- Verifying data structure
- Comparing before/after preprocessing

## Basic Usage

### Quick Start

Preview your dataset:

```bash
preprocess skim --data ./my_dataset.csv
```

This displays a formatted table showing the first few rows and columns of your data.

### Command Reference

```bash
preprocess skim [flags]
```

| Flag | Shorthand | Default | Description |
|------|-----------|---------|-------------|
| `--data` | `-d` | | **Required**: Path to the dataset file |
| `--sep` | `-s` | `,` | CSV separator |
| `--dsep` | `-m` | `.` | Decimal separator |
| `--encoding` | `-e` | `utf-8` | File encoding |

## What You'll See

The `skim` command outputs:

1. **Dataset Overview**: Total number of columns and rows
2. **Column Headers**: Names of all columns
3. **Data Preview**: First few rows of data
4. **Formatting**: Clean, aligned output

### Example Output

```
Columns: 8 | Total rows: 1000

customer_id  age  salary     department   hire_date  is_active  score      notes
-----------  ---  -------    -----------   ---------  ---------  ------     -----
1            35   75000.50   Engineering    2020-01-15  true       85.5       Good
2            28   65000.00   Marketing      2019-05-22  true       78.0       Excellent
3            42   90000.00   Sales          2018-11-03  true       92.3       
4            31   68000.00   Engineering    2021-02-28  false      76.5       Average
5            29   72000.00   HR             2020-07-10  true       88.0       
```

## Practical Examples

### Example 1: Basic Skim

```bash
preprocess skim --data ./employees.csv
```

### Example 2: European Format Data

```bash
preprocess skim --data ./european_data.csv --sep ";" --dsep ","
```

### Example 3: Tab-Separated Data

```bash
preprocess skim --data ./data.tsv --sep $'\t'
```

### Example 4: Verify Preprocessing Results

After running preprocessing, verify the output:

```bash
# Run preprocessing
preprocess run --file my_prep.toml

# Skim the cleaned data
preprocess skim --data ./cleaned_data.csv
```

## Real-World Scenarios

### Scenario 1: Initial Data Inspection

You've just received a new dataset and want to quickly understand its structure:

```bash
# First command to run on any new dataset
preprocess skim --data ./new_dataset.csv
```

This immediately tells you:
- How many columns and rows
- Column names
- Data types (visible from the content)
- Sample values

### Scenario 2: Data Quality Spot Check

Quickly verify data quality after receiving new data:

```bash
# Check if the new batch looks correct
preprocess skim --data ./new_batch.csv
```

Look for:
- Expected number of columns
- Reasonable values in each column
- No obvious formatting issues

### Scenario 3: Pre- and Post-Processing Comparison

Compare data before and after preprocessing:

```bash
# Before preprocessing
echo "=== ORIGINAL DATA ==="
preprocess skim --data ./original.csv

# After preprocessing
echo "=== CLEANED DATA ==="
preprocess skim --data ./cleaned.csv
```

### Scenario 4: Verifying Column Operations

Check if your preprocessing operations worked correctly:

```bash
# Original data
preprocess skim --data ./data.csv

# After fillna operation
preprocess run --data ./data.csv --column age --op fillna:method=mean
preprocess skim --data ./data_cleaned.csv
```

Verify that:
- Missing values in the `age` column are now filled
- Other columns are unchanged

### Scenario 5: Batch Verification

Create a script to verify multiple processed files:

```bash
#!/bin/bash
# verify_processed.sh

FILES=("processed_1.csv" "processed_2.csv" "processed_3.csv")

for file in "${FILES[@]}"; do
  echo "=== Verifying $file ==="
  preprocess skim --data ./$file
  echo ""
done
```

## Understanding the Output

### Format Interpretation

The `skim` output uses alignment and truncation to display data cleanly:

```
Columns: 5 | Total rows: 100

id    name          value  category    date
----  ------------  -----  ----------  -----------
1     Product A     100.50  Electronics  2024-01-01
2     Product B     200.75  Electronics  2024-01-02
3     Product C     150.00  Furniture    2024-01-03
```

- **First line**: Shows total columns and rows
- **Second line**: Empty for readability
- **Third line**: Column headers
- **Fourth line**: Separator line (dashes)
- **Subsequent lines**: Data rows

### Truncation Rules

For large datasets:
- Only first 30 rows are displayed
- Only first 10 columns are displayed (use `...` to indicate more columns exist)
- Long strings are truncated with `...`

## Best Practices

### 1. Always Skim First

Before any analysis or preprocessing:

```bash
# Step 1: Always skim first
preprocess skim --data ./data.csv

# Step 2: Then proceed with other operations
```

### 2. Use for Quick Verification

After any preprocessing step:

```bash
preprocess run --file my_prep.toml
preprocess skim --data ./output.csv  # Verify it worked
```

### 3. Check Column Names

Use `skim` to verify column names before referencing them in Prepfile:

```bash
preprocess skim --data ./data.csv
```

Then use the exact column names you see in your Prepfile:

```toml
[[preprocess.columns]]
name = 'correct_column_name'  # Use the name you saw in skim output
type = 'int'
```

### 4. Combine with Other Commands

Use `skim` in combination with other commands for efficient workflows:

```bash
# Explore, then summarize, then process
preprocess skim --data ./data.csv
preprocess summary --data ./data.csv --html
preprocess init --data ./data.csv
```

## Troubleshooting

### Common Issues

**Issue: File not found**
```
Error: Failed to read dataset
```

**Solution**: Verify the path:
```bash
ls -la ./my_data.csv  # Check file exists
preprocess skim --data ./my_data.csv  # Use correct path
```

**Issue: Separator problems**
```
Columns not aligning correctly
```

**Solution**: Specify the correct separator:
```bash
preprocess skim --data ./data.csv --sep ";"
```

**Issue: Encoding problems**
```
Error: Invalid UTF-8 encoding
```

**Solution**: Try different encodings:
```bash
preprocess skim --data ./data.csv --encoding latin-1
```

**Issue: All data appears in one column**
```
All data is in the first column, other columns are empty
```

**Solution**: Your separator is wrong. Try common alternatives:
```bash
# Try semicolon
preprocess skim --data ./data.csv --sep ";"

# Try tab
preprocess skim --data ./data.csv --sep $'\t'

# Try pipe
preprocess skim --data ./data.csv --sep "|"
```

## Tips and Tricks

### Quick Column Count

Need to quickly count columns?

```bash
preprocess skim --data ./data.csv 2>&1 | head -1
```

This outputs just the first line which contains the column count.

### Check for Missing Values

While `skim` shows sample data, look for:
- Empty cells in the output
- Placeholder values like `NA`, `N/A`, or `null`

```bash
preprocess skim --data ./data.csv | grep -i "null\|na\|missing"
```

### Verify Data Types

The values displayed give you clues about data types:
- Numbers with decimals = `float`
- Numbers without decimals = `int`
- Text = `string`

### Compare File Sizes

Quickly verify if preprocessing changed the dataset size:

```bash
# Original
preprocess skim --data ./original.csv

# Processed
preprocess skim --data ./processed.csv
```

Compare the row counts in the first line of output.

## Next Steps

After skimming your data:

1. **[Generate summary statistics]({{% relref "summary" %}})**: Get detailed statistics for each column
2. **[Create Prepfile]({{% relref "init" %}})**: Start building your preprocessing configuration
3. **[Run preprocessing]({{% relref "run" %}})**: Execute your preprocessing pipeline
4. **[Compare differences]({{% relref "diff" %}})**: Use `preprocess diff` to compare datasets

## Quick Reference Card

| Task | Command |
|------|---------|
| Basic skim | `preprocess skim --data ./data.csv` |
| European format | `preprocess skim --data ./data.csv --sep ";" --dsep ","` |
| Tab-separated | `preprocess skim --data ./data.tsv --sep $'\t'` |
| Verify preprocessing | `preprocess skim --data ./cleaned.csv` |
| Check column names | `preprocess skim --data ./data.csv` |

## Output Interpretation Guide

### What to Look For

| Observation | What It Means | Action |
|-------------|---------------|--------|
| Row count is 0 | Empty dataset or wrong separator | Check file, try different separator |
| All columns in one | Wrong separator | Try different separator |
| Many empty cells | Missing data | Consider imputation strategy |
| Mixed data types in column | Inconsistent data | Investigate data quality |
| Unexpected values | Data entry errors | Investigate outliers |
| Inconsistent formatting | Formatting issues | Consider cleaning operations |

### Data Quality Quick Check

When you run `skim`, quickly check:

1. **Row count**: Does it match expectations?
2. **Column count**: Does it match the data dictionary?
3. **Column names**: Do they make sense?
4. **Sample values**: Do they look reasonable?
5. **Formatting**: Is the data aligned correctly?

If any of these look wrong, investigate before proceeding with preprocessing.

## Advanced: Integrating with Workflows

### In Data Processing Scripts

```bash
#!/bin/bash
# process_data.sh

echo "=== Data Processing Started ==="
echo "Input data preview:"
preprocess skim --data ./input.csv

echo ""
echo "=== Running preprocessing ==="
preprocess run --file config.toml

echo ""
echo "=== Output data preview ==="
preprocess skim --data ./output.csv

echo ""
echo "=== Processing Complete ==="
```

### In Jupyter Notebooks or Python Scripts

While preprocess is a CLI tool, you can call it from Python:

```python
import subprocess

# Skim data
result = subprocess.run(
    ["preprocess", "skim", "--data", "./data.csv"],
    capture_output=True,
    text=True
)

print(result.stdout)
```

This allows you to integrate preprocess insights into your Python workflows.
