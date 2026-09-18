---
weight: 300
title: "preprocess summary"
icon: "article"
date: "2025-06-01T19:52:40+02:00"
lastmod: "2026-09-18T11:30:00+02:00"
description: "Generate comprehensive summary statistics for your dataset"
draft: false
toc: true
--- 

## Overview

The `preprocess summary` command is your **data exploration powerhouse**. It generates comprehensive summary statistics that help you understand your dataset's structure, quality, and characteristics before and after preprocessing.

For Data Scientists and Analysts, this command provides the insights you need to:
- Assess data quality
- Identify missing values
- Understand distributions
- Detect outliers
- Validate preprocessing results

## Basic Usage

### Quick Start

Generate a summary for your dataset:

```bash
preprocess summary --data ./my_dataset.csv
```

This creates a `Summaryfile.toml` file containing detailed statistics for each column.

### Command Reference

```bash
preprocess summary [flags]
```

| Flag | Shorthand | Default | Description |
|------|-----------|---------|-------------|
| `--data` | `-d` | | Path to the dataset file |
| `--prepfile` | `-f` | `Prepfile.toml` | Use a Prepfile for data specifications |
| `--sep` | `-s` | `,` | CSV separator |
| `--dsep` | `-m` | `.` | Decimal separator |
| `--encoding` | `-e` | `utf-8` | File encoding |
| `--output` | `-o` | `Summaryfile.toml` | Output filename |
| `--html` | `-t` | | Generate HTML report instead of TOML |
| `--no-browser` | `-b` | | Don't automatically open browser for HTML |
| `--exclude` | | | Exclude specific columns from summary |

## Output Formats

### TOML Format (Default)

The default output is a TOML file containing structured summary statistics:

```bash
preprocess summary --data ./data.csv --output stats.toml
```

This format is ideal for:
- Programmatic access to statistics
- Version control (can track changes in data characteristics)
- Integration with other tools
- Automated reporting

### HTML Format

Generate a beautiful, interactive HTML report:

```bash
preprocess summary --data ./data.csv --html --output report.html
```

The HTML report provides:
- Color-coded visualizations
- Sortable tables
- Interactive exploration
- Professional presentation
- Easy sharing with non-technical stakeholders

## What Statistics Are Generated?

### For All Columns

| Statistic | Description | Use Case |
|-----------|-------------|----------|
| `rows_count` | Number of non-null values | Identify missing data extent |
| `type` | Column data type | Understand data structure |
| `missing` | Count of missing values | Assess data completeness |

### For Numeric Columns

| Statistic | Description | Use Case |
|-----------|-------------|----------|
| `min` | Minimum value | Identify range, detect outliers |
| `max` | Maximum value | Identify range, detect outliers |
| `mean` | Arithmetic mean | Understand central tendency |
| `median` | Median value | Understand central tendency (robust to outliers) |
| `std` | Standard deviation | Understand variability |
| `q1` | First quartile (25th percentile) | Understand distribution |
| `q3` | Third quartile (75th percentile) | Understand distribution |
| `mode` | Most frequent value | Identify common values |

### For Text Columns

| Statistic | Description | Use Case |
|-----------|-------------|----------|
| `modality_count` | Number of unique values | Assess cardinality |
| `top_value` | Most frequent value | Identify common categories |
| `top_frequency` | Frequency of most common value | Understand distribution |
| `min_length` | Minimum string length | Identify formatting issues |
| `max_length` | Maximum string length | Identify formatting issues |
| `avg_length` | Average string length | Understand data characteristics |

## Practical Examples

### Example 1: Basic Summary Generation

```bash
preprocess summary --data ./sales_data.csv
```

This creates `Summaryfile.toml` with statistics for all columns.

### Example 2: Generate HTML Report

```bash
preprocess summary --data ./customer_data.csv --html
```

This creates `customer_data_report.html` and opens it in your default browser.

### Example 3: Save to Specific File

```bash
preprocess summary --data ./data.csv --output analysis_stats.toml
```

### Example 4: Exclude Columns

Exclude sensitive or irrelevant columns from the summary:

```bash
preprocess summary --data ./employees.csv \
  --exclude salary \
  --exclude ssn \
  --exclude password
```

### Example 5: Use Prepfile Specifications

Use the data specifications from your Prepfile:

```bash
preprocess summary --prepfile ./config/my_prep.toml
```

This is useful when you want to use the same data reading configuration as your preprocessing pipeline.

### Example 6: European Format Data

```bash
preprocess summary --data ./european_data.csv --sep ";" --dsep ","
```

## Real-World Scenarios

### Scenario 1: Initial Data Exploration

You've just received a new dataset and want to understand it:

```bash
# First, skim the data
preprocess skim --data ./new_dataset.csv

# Then generate detailed statistics
preprocess summary --data ./new_dataset.csv --html --output exploration.html
```

Open the HTML report to:
- See the complete structure of your data
- Identify columns with missing values
- Understand distributions
- Spot potential data quality issues

### Scenario 2: Pre- and Post-Processing Comparison

Compare your data before and after preprocessing:

```bash
# Generate summary of original data
preprocess summary --data ./original_data.csv --output before_summary.toml

# Run preprocessing
preprocess run --file my_prep.toml

# Generate summary of cleaned data
preprocess summary --data ./cleaned_data.csv --output after_summary.toml
```

Now you can compare the two TOML files to see how your preprocessing affected the data.

### Scenario 3: Automated Data Quality Reporting

Create a script to generate regular data quality reports:

```bash
#!/bin/bash
# generate_quality_report.sh

DATE=$(date +%Y%m%d)

preprocess summary \
  --data ./daily_data.csv \
  --html \
  --output ./reports/quality_report_$DATE.html \
  --no-browser

echo "Data quality report generated: ./reports/quality_report_$DATE.html"
```

### Scenario 4: Batch Processing Multiple Files

Generate summaries for multiple datasets:

```bash
#!/bin/bash
# batch_summary.sh

DATASETS=("customers.csv" "products.csv" "orders.csv")

for dataset in "${DATASETS[@]}"; do
  echo "Processing $dataset..."
  preprocess summary --data ./$dataset --output summaries/${dataset%.csv}_summary.toml
  echo "Summary created for $dataset"
done
```

## Understanding the Output

### TOML Output Structure

The TOML output has this structure:

```toml
[data]
filename = './my_dataset.csv'
csv_separator = ','
decimal_separator = '.'
encoding = 'utf-8'

[data_summary]
rows_count = 10000
columns_count = 15
numeric_columns = 10
string_columns = 5

[[columns]]
name = 'age'
type = 'numeric'
rows_count = 9500
missing = 500
min = 18.0
max = 85.0
mean = 42.5
median = 41.0
std = 12.3
q1 = 30.0
q3 = 55.0

[[columns]]
name = 'category'
type = 'string'
rows_count = 10000
missing = 0
modality_count = 25
min_length = 3
max_length = 45
avg_length = 12.5
top_value = "Electronics"
top_frequency = 2500
```

### HTML Output Features

The HTML report includes:

1. **Overview Section**:
   - Total rows and columns
   - Data types distribution
   - Missing values summary

2. **Numeric Columns Table**:
   - Sortable by any statistic
   - Color-coded missing values
   - Distribution indicators

3. **Text Columns Table**:
   - Cardinality indicators
   - Top values displayed
   - Length statistics

4. **Visualizations**:
   - Missing values heatmap
   - Distribution histograms for numeric columns
   - Bar charts for top values in text columns

## Best Practices

### 1. Always Explore Before Processing

Before writing any preprocessing logic:

```bash
# Step 1: Skim the data
preprocess skim --data ./data.csv

# Step 2: Generate summary statistics
preprocess summary --data ./data.csv --html

# Step 3: Review the report
# Then design your preprocessing pipeline
```

### 2. Use Exclusions Wisely

Exclude columns that don't provide valuable insights:

```bash
preprocess summary --data ./data.csv \
  --exclude id \
  --exclude created_at \
  --exclude internal_notes
```

This makes your summary more focused and easier to read.

### 3. Save Both Formats

For comprehensive analysis, save both TOML and HTML:

```bash
preprocess summary --data ./data.csv --output stats.toml
preprocess summary --data ./data.csv --html --output stats.html
```

- Use TOML for programmatic analysis
- Use HTML for presentations and sharing

### 4. Document Your Findings

Add notes to your summary files or include them in your project documentation:

```markdown
## Data Quality Assessment

Based on summary statistics from `preprocess summary`:

- **Missing Values**: 5% of age values are missing
- **Outliers**: Salary values range from 20K to 5M (investigate upper range)
- **Cardinality**: Category column has 150 unique values (high cardinality)
- **Data Types**: All numeric columns detected correctly

**Actions Taken**:
- Impute missing ages with median
- Investigate salary outliers
- Consider dimensionality reduction for category column
```

## Troubleshooting

### Common Issues

**Issue: No data source specified**
```
No data source specified. Please provide a data file or a prepfile.
```

**Solution**: Provide either `--data` or `--prepfile`:
```bash
preprocess summary --data ./my_data.csv
# OR
preprocess summary --prepfile ./my_prep.toml
```

**Issue: Encoding problems**
```
Error: Invalid UTF-8 encoding
```

**Solution**: Try different encodings:
```bash
preprocess summary --data ./data.csv --encoding latin-1
```

**Issue: Separator problems**
```
Error: Failed to parse CSV
```

**Solution**: Specify the correct separator:
```bash
preprocess summary --data ./data.csv --sep ";"
```

**Issue: HTML generation fails**
```
Error: Failed to generate HTML
```

**Solution**: Ensure you have write permissions in the directory and try again.

## Performance Considerations

### Large Datasets

For large datasets, summary generation can be resource-intensive:

```bash
# For very large files, be patient
preprocess summary --data ./large_dataset.csv
```

**Tips:**
- Exclude columns you don't need with `--exclude`
- Consider sampling your data for initial exploration
- Use TOML format (faster than HTML)

### Memory Usage

Summary generation loads the entire dataset into memory. For extremely large datasets:
- Process in batches
- Use the `exclude` flag to reduce memory footprint
- Consider generating summaries for subsets of your data

## Next Steps

After generating summary statistics:

1. **[Review the output]**: Understand your data characteristics
2. **[Identify issues]**: Look for missing values, outliers, data quality problems
3. **[Design preprocessing]**: Use insights to inform your Prepfile configuration
4. **[Verify preprocessing]**: Generate new summaries after preprocessing to confirm improvements
5. **[Create Prepfile]({{% relref "init" %}})**: Start building your preprocessing pipeline

## Quick Reference Card

| Task | Command |
|------|---------|
| Basic summary (TOML) | `preprocess summary --data ./data.csv` |
| HTML report | `preprocess summary --data ./data.csv --html` |
| Save to file | `preprocess summary --data ./data.csv --output my_stats.toml` |
| Exclude columns | `preprocess summary --data ./data.csv --exclude col1 --exclude col2` |
| Use Prepfile | `preprocess summary --prepfile ./my_prep.toml` |
| No browser | `preprocess summary --data ./data.csv --html --no-browser` |
| Custom separator | `preprocess summary --data ./data.csv --sep ";"` |
| European decimals | `preprocess summary --data ./data.csv --dsep ","` |

## Summary Statistics Reference

### Numeric Column Statistics

| Statistic | Formula/Calculation | What It Tells You |
|-----------|---------------------|-------------------|
| `min` | Minimum observed value | Lower bound, potential outliers |
| `max` | Maximum observed value | Upper bound, potential outliers |
| `mean` | Sum of all values / count | Central tendency (affected by outliers) |
| `median` | Middle value | Central tendency (robust to outliers) |
| `std` | Square root of variance | Spread/variability of data |
| `q1` | 25th percentile | Lower quartile boundary |
| `q3` | 75th percentile | Upper quartile boundary |
| `mode` | Most frequent value | Most common observation |
| `missing` | Count of null/NA values | Data completeness |

### Text Column Statistics

| Statistic | Calculation | What It Tells You |
|-----------|-------------|-------------------|
| `modality_count` | Count of unique values | Cardinality, potential for encoding |
| `top_value` | Most frequent string | Dominant category |
| `top_frequency` | Count of top value | Distribution concentration |
| `min_length` | Shortest string | Formatting consistency |
| `max_length` | Longest string | Formatting consistency, potential errors |
| `avg_length` | Average string length | General string characteristics |
| `missing` | Count of null/NA values | Data completeness |

### Data Quality Indicators

| Indicator | How to Identify | Action |
|-----------|-----------------|--------|
| High missing rate | `missing` > 10% of `rows_count` | Investigate, consider imputation or removal |
| Outliers | `min` or `max` far from `q1`/`q3` | Investigate, consider winsorization or removal |
| High cardinality | `modality_count` very high | Consider grouping or encoding strategies |
| Zero variance | `min` = `max` | Feature provides no information, consider removal |
| Inconsistent lengths | Large range between `min_length` and `max_length` | Investigate formatting, potential data entry errors |
