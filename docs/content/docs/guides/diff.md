---
weight: 400
title: "preprocess diff"
icon: "article"
date: "2025-06-01T19:56:40+02:00"
lastmod: "2026-09-18T11:30:00+02:00"
description: "Compare two datasets to identify differences"
draft: false
toc: true
--- 

## Overview

The `preprocess diff` command compares two datasets and generates a **visual HTML report** showing the differences between them. This is essential for:
- Verifying preprocessing results
- Comparing data versions
- Tracking data changes over time
- Quality assurance
- Auditing data transformations

For Data Scientists and Analysts, `diff` provides a powerful way to understand exactly how your data has changed through preprocessing or between different versions.

## Basic Usage

### Quick Start

Compare two datasets:

```bash
preprocess diff --source ./original.csv --target ./processed.csv
```

This generates `htmldiff.html` and opens it in your default browser.

### Command Reference

```bash
preprocess diff [flags]
```

| Flag | Shorthand | Default | Description |
|------|-----------|---------|-------------|
| `--source` | | | **Required**: Path to the source (original) dataset |
| `--target` | | | **Required**: Path to the target (modified) dataset |
| `--sep` | `-s` | `,` | CSV separator for both files |
| `--dsep` | `-m` | `.` | Decimal separator for both files |
| `--html` | `-t` | | Generate HTML report (default: true) |
| `--no-browser` | `-b` | | Don't automatically open browser |

## What the Diff Report Shows

The generated HTML report provides a comprehensive comparison including:

1. **Overview Statistics**:
   - Rows added, removed, and unchanged
   - Columns added, removed, and modified
   - Summary of changes

2. **Detailed Changes**:
   - Cell-by-cell comparison
   - Value changes highlighted
   - Added/removed rows identified

3. **Visual Indicators**:
   - Color-coding for different types of changes
   - Side-by-side comparison where applicable
   - Summary statistics

4. **Column-Level Analysis**:
   - Changes in each column
   - Type changes
   - Missing value changes

## Practical Examples

### Example 1: Basic Comparison

```bash
preprocess diff --source ./original.csv --target ./cleaned.csv
```

This creates `htmldiff.html` showing all differences between the two files.

### Example 2: No Browser

Generate the report without automatically opening it:

```bash
preprocess diff --source ./v1.csv --target ./v2.csv --no-browser
```

### Example 3: Custom Separator

Compare files with non-standard separators:

```bash
preprocess diff --source ./data1.tsv --target ./data2.tsv --sep $'\t'
```

### Example 4: European Format

Compare files with European formatting:

```bash
preprocess diff --source ./eu_data1.csv --target ./eu_data2.csv --sep ";" --dsep ","
```

## Real-World Scenarios

### Scenario 1: Verify Preprocessing Results

The most common use case - verify that your preprocessing worked as expected:

```bash
# Generate Prepfile
preprocess init --data ./original.csv --output cleaning.toml

# Edit Prepfile with your preprocessing operations

# Run preprocessing
preprocess run --file cleaning.toml

# Compare original vs. cleaned
preprocess diff --source ./original.csv --target ./original_cleaned.csv
```

**What to look for in the report:**
- Missing values filled in
- Data type changes
- Value transformations (scaling, etc.)
- Columns added or removed

### Scenario 2: Compare Data Versions

Track how your data has changed over time:

```bash
# Compare weekly data snapshots
preprocess diff \
  --source ./data/week_42.csv \
  --target ./data/week_43.csv \
  --output week_42_to_43_diff.html
```

### Scenario 3: Quality Assurance

Verify that a data migration or transformation preserved data integrity:

```bash
# Before migration
preprocess summary --data ./before_migration.csv --output before_summary.toml

# After migration
preprocess summary --data ./after_migration.csv --output after_summary.toml

# Compare to ensure nothing unexpected changed
preprocess diff --source ./before_migration.csv --target ./after_migration.csv
```

### Scenario 4: Automated Comparison in Scripts

Integrate diff into your data processing scripts:

```bash
#!/bin/bash
# compare_datasets.sh

SOURCE=$1
TARGET=$2
OUTPUT=$3

echo "Comparing $SOURCE and $TARGET..."
preprocess diff --source $SOURCE --target $TARGET --no-browser --output $OUTPUT

echo "Comparison report saved to $OUTPUT"
```

Run it:
```bash
./compare_datasets.sh ./original.csv ./processed.csv ./reports/comparison.html
```

### Scenario 5: Batch Comparison

Compare multiple dataset pairs:

```bash
#!/bin/bash
# batch_compare.sh

PAIRS=(
  "raw:cleaned"
  "v1:v2"
  "before:after"
)

for pair in "${PAIRS[@]}"; do
  IFS=':' read -r source target <<< "$pair"
  echo "Comparing ${source}.csv and ${target}.csv..."
  preprocess diff \
    --source ./${source}.csv \
    --target ./${target}.csv \
    --no-browser \
    --output ./reports/${source}_to_${target}.html
  echo "Report saved."
done
```

## Understanding the Diff Report

### Report Structure

The HTML report is organized into several sections:

#### 1. Overview Dashboard

- **Total rows**: Count in each dataset
- **Total columns**: Count in each dataset
- **Changed rows**: Number of rows with differences
- **Added rows**: Rows present in target but not in source
- **Removed rows**: Rows present in source but not in target
- **Changed columns**: Columns with differences

#### 2. Column Comparison

For each column:
- **Status**: Added, Removed, Modified, or Unchanged
- **Type changes**: If data type changed
- **Missing value changes**: Changes in null/NA counts
- **Value changes**: Summary of value modifications

#### 3. Row-Level Differences

Detailed view showing:
- Side-by-side comparison for changed rows
- Values that were modified
- Context around changes

#### 4. Statistics Summary

- Most changed columns
- Distribution of change types
- Impact assessment

### Color Coding

The report uses color coding to make differences easy to spot:

| Color | Meaning |
|-------|---------|
| Green | Added/Created |
| Red | Removed/Deleted |
| Yellow/Orange | Modified/Changed |
| Blue | Information/Neutral |
| Gray | Unchanged |

## Best Practices

### 1. Always Compare After Preprocessing

Make diff comparison part of your preprocessing workflow:

```bash
# Standard workflow
preprocess run --file my_prep.toml
preprocess diff --source ./original.csv --target ./cleaned.csv
```

### 2. Name Your Reports Descriptively

Use meaningful output filenames:

```bash
preprocess diff \
  --source ./raw_data.csv \
  --target ./cleaned_data.csv \
  --output ./reports/preprocessing_changes_$(date +%Y%m%d).html
```

### 3. Review Before Production

Before deploying preprocessing to production:

```bash
# Test on sample data
preprocess run --file new_prep.toml --data ./sample.csv

# Compare with expected output
preprocess diff \
  --source ./sample_expected.csv \
  --target ./sample_cleaned.csv

# Only proceed if differences are expected and correct
```

### 4. Archive Reports

Save diff reports for audit purposes:

```bash
#!/bin/bash
# archive_comparison.sh

DATE=$(date +%Y%m%d_%H%M%S)
mkdir -p ./archive/comparisons

preprocess diff \
  --source ./original.csv \
  --target ./processed.csv \
  --no-browser \
  --output ./archive/comparisons/comparison_$DATE.html

echo "Archived comparison report: comparison_$DATE.html"
```

## Troubleshooting

### Common Issues

**Issue: Files must have same structure**
```
Error: Datasets have different structures
```

**Solution**: The diff command works best with datasets that have the same or similar structure. If your files have completely different columns, consider:
- Adding missing columns with default values
- Removing extra columns before comparison
- Using `preprocess run` to standardize first

**Issue: Large datasets**
```
Processing takes a long time
```

**Solution**: For very large datasets:
- Compare a sample first
- Focus on specific columns of interest
- Ensure you have enough memory

**Issue: Browser doesn't open**
```
HTML generated but browser didn't open
```

**Solution**: Manually open the file or use `--no-browser` and open it yourself:
```bash
preprocess diff --source ./a.csv --target ./b.csv --no-browser
# Then open htmldiff.html in your browser
```

**Issue: No differences detected**
```
Report shows no differences but I expected some
```

**Solution**: Check:
- Are you comparing the right files?
- Did your preprocessing actually modify the data?
- Are the files truly different? (Use `diff` command to verify)

## Advanced Usage

### Comparing Specific Columns

While `diff` compares all columns, you can focus on specific ones by preprocessing first:

```bash
# Create a Prepfile that keeps only columns of interest
preprocess init --data ./source.csv --output diff_columns.toml

# Edit to drop columns you don't want to compare
# Then run preprocessing
preprocess run --file diff_columns.toml

# Now compare only those columns
preprocess diff --source ./source_selected.csv --target ./target_selected.csv
```

### Chaining with Other Commands

Combine diff with other preprocess commands:

```bash
# Full workflow: skim, summary, run, diff
preprocess skim --data ./original.csv
preprocess summary --data ./original.csv --html
preprocess run --file cleaning.toml
preprocess diff --source ./original.csv --target ./cleaned.csv
```

### Programmatic Access

Access diff results programmatically by parsing the HTML or using summary statistics:

```bash
# Generate TOML summary of both files
preprocess summary --data ./source.csv --output source_summary.toml
preprocess summary --data ./target.csv --output target_summary.toml

# Compare the TOML files with standard tools
# This gives you programmatic access to the differences
```

## Next Steps

After using diff:

1. **[Review the report]**: Understand the differences
2. **[Validate changes]**: Ensure preprocessing worked as expected
3. **[Iterate on Prepfile]**: If changes aren't what you expected, modify your Prepfile
4. **[Document findings]**: Save the diff report for future reference
5. **[Share with team]**: Use the HTML report to communicate changes

## Quick Reference Card

| Task | Command |
|------|---------|
| Basic comparison | `preprocess diff --source ./a.csv --target ./b.csv` |
| No auto-open | `preprocess diff --source ./a.csv --target ./b.csv --no-browser` |
| Custom output | `preprocess diff --source ./a.csv --target ./b.csv --output my_diff.html` |
| Tab-separated | `preprocess diff --source ./a.tsv --target ./b.tsv --sep $'\t'` |
| European format | `preprocess diff --source ./a.csv --target ./b.csv --sep ";" --dsep ","` |

## Interpretation Guide

### Understanding Change Types

| Change Type | What It Means | Severity |
|-------------|---------------|----------|
| Rows added | New data in target | Low-Medium |
| Rows removed | Data removed from target | Medium-High |
| Column added | New column in target | Low |
| Column removed | Column removed from target | Medium |
| Column type changed | Data type transformation | Medium |
| Values changed | Data modifications | Depends |
| Missing values changed | Null handling | Depends |

### When to Investigate

| Observation | Action |
|-------------|--------|
| Unexpected row count changes | Investigate data loss |
| Column type changes not intended | Check preprocessing logic |
| Value changes in ID columns | High priority investigation |
| Large number of changed values | Verify preprocessing is working correctly |
| Missing values increased | Check why values are being set to null |

### Quality Check Checklist

When reviewing a diff report, check:

- [ ] Row counts are as expected
- [ ] All important columns are present
- [ ] Data types are correct
- [ ] Missing values are handled appropriately
- [ ] Value transformations are working as intended
- [ ] No unintended changes occurred
- [ ] Changes align with your preprocessing goals
