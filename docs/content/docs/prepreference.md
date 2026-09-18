---
weight: 501
date: "2025-05-26T21:40:40+02:00"
draft: false
author: "Axel-Cleris Gailloty"
title: "Prepfile Reference Guide"
icon: "developer_guide"
toc: true
description: "Complete reference of all Prepfile configuration options"
publishdate: "2025-05-26T21:40:40+02:00"
lastmod: "2026-09-18T11:30:00+02:00"
--- 

## Introduction to Prepfile

The **Prepfile.toml** is the heart of preprocess CLI. It's a configuration file where you declare your entire data preprocessing pipeline in a simple, readable format. If you've used Docker's `Dockerfile` or docker-compose's `compose.yml`, the concept will be familiar.

Think of the Prepfile as your **data processing recipe**: it contains all the instructions preprocess needs to transform your raw data into cleaned, analysis-ready datasets.

## Why Use Prepfile?

For Data Scientists and Data Analysts:

- **Declarative**: Describe *what* you want, not *how* to do it
- **Reproducible**: Same configuration produces same results every time
- **Version Control Friendly**: Track changes to your preprocessing logic
- **Collaborative**: Share preprocessing pipelines with team members
- **Documented**: The Prepfile itself documents your data processing steps

## Prepfile Structure

A complete Prepfile has three main sections:

```toml
# Section 1: Data Configuration
[data]
filename = './my_dataset.csv'
csv_separator = ','
decimal_separator = '.'
encoding = 'utf-8'
missing_identifier = ''

# Section 2: Preprocessing Operations
[preprocess]
[preprocess.numerics]
operations = [
    {op = "fillna", method = "median"}
]

[preprocess.texts]
operations = [
    {op = "clean", method = "trimws"}
]

[[preprocess.columns]]
name = 'age'
type = 'int'
operations = [
    {op = "fillna", method = "mean"}
]

# Section 3: Postprocessing Configuration
[postprocess]
format = 'csv'
filename = 'my_dataset_cleaned.csv'
dropcolumns = ["temp_column"]
sortdataset = {descending = false}
```

Let's explore each section in detail.

---

## [data] Section

The `[data]` section tells preprocess **how to read your dataset**. This is the only mandatory section - every Prepfile must have it.

### Parameters

{{% table "table-hover" %}}
| Parameter | Default | Required | Description | Common Values |
|-----------|---------|----------|-------------|---------------|
| `filename` | N/A | **Yes** | Path to your dataset file | `./data.csv`, `/path/to/file.csv` |
| `csv_separator` | `,` | No | Character used to separate columns | `,`, `;`, `\t`, `|` |
| `decimal_separator` | `.` | No | Character used for decimal point | `.`, `,` |
| `encoding` | `utf-8` | No | Text encoding of the file | `utf-8`, `latin-1`, `iso-8859-1` |
| `missing_identifier` | `""` | No | How missing values are represented | `""`, `"NA"`, `"N/A"`, `"null"` |
{{% /table %}}

### Examples

#### Basic Configuration

```toml
[data]
filename = './sales_data.csv'
```

This uses all defaults: comma separator, dot decimal, UTF-8 encoding, empty string for missing values.

#### European Format Data

```toml
[data]
filename = './european_sales.csv'
csv_separator = ';'
decimal_separator = ','
encoding = 'utf-8'
```

#### Tab-Separated Values

```toml
[data]
filename = './data.tsv'
csv_separator = '\t'
```

**Note**: Use `\t` for tab character in TOML.

#### Legacy System Data

```toml
[data]
filename = './legacy_export.txt'
csv_separator = '|'
decimal_separator = '.'
encoding = 'latin-1'
missing_identifier = 'NULL'
```

### Understanding missing_identifier

The `missing_identifier` parameter is **critical** for accurate preprocessing. It tells preprocess how to recognize missing data in your file.

| Value | Meaning |
|-------|---------|
| `""` (empty string) | Empty cells are missing |
| `"NA"` | Cells containing exactly "NA" are missing |
| `"N/A"` | Cells containing exactly "N/A" are missing |
| `"null"` | Cells containing exactly "null" are missing |
| `""` | Any empty string is missing |

**Important**: The value must match exactly. If your data uses `" NA "` (with spaces), it won't be recognized as missing with `missing_identifier = "NA"`.

**Best Practice**: 
1. First, skim your data: `preprocess skim --data ./your_file.csv`
2. Identify how missing values appear
3. Set `missing_identifier` accordingly in your Prepfile

---

## [preprocess] Section

The `[preprocess]` section is where you **define your preprocessing operations**. This is the most powerful and flexible part of the Prepfile.

You have **three ways** to apply operations:

1. **On all numeric columns** using `[preprocess.numerics]`
2. **On all text columns** using `[preprocess.texts]`
3. **On specific columns** using `[[preprocess.columns]]`

You can use any combination of these approaches in a single Prepfile.

### 1. [preprocess.numerics] - Apply to All Numeric Columns

Apply the same operations to every numeric column in your dataset.

#### Syntax

```toml
[preprocess.numerics]
operations = [
    {op = "operation_name", method = "method_name"},
    {op = "another_operation", method = "another_method"}
]
exclude_columns = ["column1", "column2"]
```

#### Parameters

{{% table "table-hover" %}}
| Parameter | Required | Description |
|-----------|----------|-------------|
| `operations` | Yes | List of operations to apply |
| `exclude_columns` | No | Columns to exclude from these operations |
{{% /table %}}

#### Examples

##### Fill missing values in all numeric columns

```toml
[preprocess.numerics]
operations = [
    {op = "fillna", method = "median"}
]
```

##### Scale all numeric columns

```toml
[preprocess.numerics]
operations = [
    {op = "scale", method = "zscore"}
]
```

##### Multiple operations with exclusions

```toml
[preprocess.numerics]
operations = [
    {op = "fillna", method = "mean"},
    {op = "scale", method = "minmax"}
]
exclude_columns = ["id", "timestamp"]
```

**Note**: Operations are applied in the order they appear in the list.

### 2. [preprocess.texts] - Apply to All Text Columns

Apply the same operations to every text (string) column in your dataset.

#### Syntax

```toml
[preprocess.texts]
operations = [
    {op = "clean", method = "trimws"},
    {op = "clean", method = "lower"}
]
exclude_columns = ["id"]
```

#### Examples

##### Clean all text columns

```toml
[preprocess.texts]
operations = [
    {op = "clean", method = "trimws"},
    {op = "clean", method = "lower"}
]
```

This removes whitespace and converts all text to lowercase.

##### Standardize text

```toml
[preprocess.texts]
operations = [
    {op = "clean", method = "trimws"},
    {op = "clean", method = "title"}
]
exclude_columns = ["product_id", "category_code"]
```

### 3. [[preprocess.columns]] - Apply to Specific Columns

Apply operations to specific, individual columns. This gives you the most control.

#### Syntax

```toml
[[preprocess.columns]]
name = 'column_name'
type = 'column_type'  # 'int', 'float', or 'string'
new_name = 'new_column_name'  # Optional: rename the column
operations = [
    {op = "operation", method = "method"}
]

[[preprocess.columns]]
name = 'another_column'
type = 'float'
operations = [
    {op = "fillna", method = "mean"},
    {op = "scale", method = "zscore"}
]
```

#### Parameters

{{% table "table-hover" %}}
| Parameter | Required | Description |
|-----------|----------|-------------|
| `name` | **Yes** | Name of the column to process |
| `type` | **Yes** | Data type: `'int'`, `'float'`, or `'string'` |
| `new_name` | No | New name for the column (for renaming) |
| `operations` | No | List of operations to apply to this column |
{{% /table %}}

#### Examples

##### Single operation on specific column

```toml
[[preprocess.columns]]
name = 'age'
type = 'int'
operations = [
    {op = "fillna", method = "median"}
]
```

##### Multiple operations on one column

```toml
[[preprocess.columns]]
name = 'salary'
type = 'float'
operations = [
    {op = "fillna", method = "mean"},
    {op = "scale", method = "zscore"}
]
```

##### Rename a column

```toml
[[preprocess.columns]]
name = 'customer_id'
type = 'string'
new_name = 'client_id'
```

##### Rename with operations

```toml
[[preprocess.columns]]
name = 'purchase_date'
type = 'string'
new_name = 'transaction_date'
operations = [
    {op = "clean", method = "trimws"}
]
```

### Operation Syntax

Operations use a consistent syntax across all sections:

```toml
{op = "operation_name", method = "method_name"}
```

Or for operations that don't need a method:

```toml
{op = "operation_name"}
```

Or for operations that take parameters:

```toml
{op = "fillna", value = 0}
{op = "fillna", method = "mean"}
```

### TOML Syntax Notes

TOML offers two equivalent syntaxes for arrays of tables. Both work, but **don't mix them**:

**Style 1: Inline array**
```toml
[preprocess.texts]
operations = [
    {op = "clean", method = "trimws"},
    {op = "clean", method = "lower"}
]
```

**Style 2: Array of tables**
```toml
[[preprocess.texts.operations]]
op = "clean"
method = "trimws"

[[preprocess.texts.operations]]
op = "clean"
method = "lower"
```

Choose one style and use it consistently throughout your Prepfile.

---

## Supported Operations

### Numeric Operations

#### fillna - Fill Missing Values

Replace missing values in numeric columns.

**Methods:**

{{% table "table-hover" %}}
| Method | Description | Best For |
|--------|-------------|----------|
| `mean` | Fill with column mean | Symmetric distributions without outliers |
| `median` | Fill with column median | Skewed distributions or with outliers |
| `value:<number>` | Fill with specified value | Custom imputation |
{{% /table %}}

**Examples:**

```toml
# Fill with mean
{op = "fillna", method = "mean"}

# Fill with median
{op = "fillna", method = "median"}

# Fill with zero
{op = "fillna", value = 0}

# Fill with specific value
{op = "fillna", value = -999}
```

#### scale - Feature Scaling

Transform numeric values to a common scale.

**Methods:**

{{% table "table-hover" %}}
| Method | Description | Formula | Use Case |
|--------|-------------|---------|----------|
| `zscore` | Standardization | (x - mean) / std | Machine learning, PCA, SVM |
| `minmax` | Normalization | (x - min) / (max - min) | Neural networks, distance-based algorithms |
{{% /table %}}

**Examples:**

```toml
# Standardize
{op = "scale", method = "zscore"}

# Normalize to [0, 1] range
{op = "scale", method = "minmax"}
```

#### discretize - Create Bins

Convert continuous numeric values into discrete bins.

**Methods:**

{{% table "table-hover" %}}
| Method | Description |
|--------|-------------|
| `binning` | Create custom bins with labels |
{{% /table %}}

**Example:**

```toml
[[preprocess.columns]]
name = 'age'
type = 'int'
operations = [
    {op = "discretize", method = "binning", bins = [
        {lower = 0, upper = 18, label = "Child"},
        {lower = 18, upper = 35, label = "Young Adult"},
        {lower = 35, upper = 65, label = "Adult"},
        {lower = 65, upper = 120, label = "Senior"}
    ]}
]
```

### Text Operations

#### fillna - Fill Missing Values

Replace missing values in text columns.

**Parameters:**

{{% table "table-hover" %}}
| Parameter | Description | Example |
|-----------|-------------|---------|
| `value` | Text to use for filling | `value = "Unknown"` |
{{% /table %}}

**Example:**

```toml
{op = "fillna", value = "Unknown"}
{op = "fillna", value = "N/A"}
```

#### clean - Text Cleaning

Clean and standardize text data.

**Methods:**

{{% table "table-hover" %}}
| Method | Description | Example |
|--------|-------------|---------|
| `trimws` | Remove leading and trailing whitespace | `" hello "` → `"hello"` |
| `upper` | Convert to uppercase | `"Hello"` → `"HELLO"` |
| `lower` | Convert to lowercase | `"Hello"` → `"hello"` |
| `title` | Convert to title case | `"hello world"` → `"Hello World"` |
{{% /table %}}

**Examples:**

```toml
# Remove whitespace
{op = "clean", method = "trimws"}

# Lowercase
{op = "clean", method = "lower"}

# Multiple cleaning operations
operations = [
    {op = "clean", method = "trimws"},
    {op = "clean", method = "lower"}
]
```

### Categorical Operations

#### dummy - Create Dummy Variables

Convert categorical variables into one-hot encoded dummy variables.

**Parameters:**

{{% table "table-hover" %}}
| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `dummy_prefix` | bool | false | Add column name prefix to dummy columns |
| `dummy_droplast` | bool | false | Drop last category to avoid multicollinearity |
| `continue_with_toomany` | bool | false | Continue if too many unique values |
{{% /table %}}

**Examples:**

##### Basic dummy encoding

```toml
[[preprocess.columns]]
name = 'category'
type = 'string'
operations = [
    {op = "dummy"}
]
```

This converts `category` with values `["A", "B", "C"]` into three columns: `A`, `B`, `C` with 0/1 values.

##### With prefix

```toml
[[preprocess.columns]]
name = 'department'
type = 'string'
operations = [
    {op = "dummy", dummy_prefix = true}
]
```

This creates columns like: `department_Sales`, `department_Marketing`, etc.

##### Drop last category

```toml
[[preprocess.columns]]
name = 'color'
type = 'string'
operations = [
    {op = "dummy", dummy_droplast = true}
]
```

If `color` has values `["Red", "Green", "Blue"]`, this creates only `Red` and `Green` columns (Blue is the reference category).

### Other Operations

#### rename - Rename Columns

Change column names.

**Usage**: Use the `new_name` parameter in the column definition:

```toml
[[preprocess.columns]]
name = 'old_name'
type = 'string'
new_name = 'new_name'
```

#### group - Group Operations

Perform operations on grouped data.

**Example:**

```toml
[[preprocess.columns]]
name = 'value'
type = 'float'
operations = [
    {op = "group", method = "mean", options = [
        {by = "category"}
    ]}
]
```

#### train_test_split - Split Dataset

Split your dataset into training and test sets. See [postprocess] section for details.

---

## [postprocess] Section

The `[postprocess]` section defines what happens **after all preprocessing operations are applied**. This is where you specify the final output format and any final transformations.

### Parameters

{{% table "table-hover" %}}
| Parameter | Default | Description |
|-----------|---------|-------------|
| `format` | `'csv'` | Output file format |
| `filename` | `'data_cleaned.csv'` | Name of the output file |
| `dropcolumns` | `[]` | List of columns to remove |
| `sortdataset` | `null` | Sorting configuration |
| `dataset_split` | `null` | Dataset splitting configuration |
{{% /table %}}

### Examples

#### Basic Configuration

```toml
[postprocess]
format = 'csv'
filename = 'my_data_cleaned.csv'
```

#### Drop Unwanted Columns

```toml
[postprocess]
format = 'csv'
filename = 'final_data.csv'
dropcolumns = ["id", "temp_column", "internal_notes"]
```

#### Sort Dataset

```toml
[postprocess]
format = 'csv'
filename = 'sorted_data.csv'
sortdataset = {descending = false}
```

Set `descending = true` for descending order.

#### Dataset Splitting

Split your dataset into training and test sets:

```toml
[postprocess]
format = 'csv'
filename = 'data_cleaned.csv'
[postprocess.dataset_split]
method = "train_test_split"
split_names = ["train", "test"]
random_seed = 42
train_test_split_ratio = 0.8
```

**Parameters:**

{{% table "table-hover" %}}
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `method` | string | Yes | Splitting method (currently only `train_test_split`) |
| `split_names` | []string | Yes | Names for each split (e.g., `["train", "test"]`) |
| `random_seed` | uint64 | Yes | Random seed for reproducibility |
| `train_test_split_ratio` | float64 | Yes | Proportion for first split (typically 0.7-0.8) |
{{% /table %}}

This will create multiple files with the specified names and the data split accordingly.

#### Complete Postprocess Example

```toml
[postprocess]
format = 'csv'
filename = 'final_output.csv'
dropcolumns = ["temp1", "temp2"]
sortdataset = {descending = true}

[postprocess.dataset_split]
method = "train_test_split"
split_names = ["train", "test", "validation"]
random_seed = 12345
train_test_split_ratio = 0.7
```

**Note**: When using dataset_split, the train_test_split_ratio applies to the first split (train), then the remaining data is split equally among the other names.

---

## Complete Prepfile Examples

### Example 1: Basic Data Cleaning

```toml
[data]
filename = './survey_data.csv'
csv_separator = ','
decimal_separator = '.'
encoding = 'utf-8'
missing_identifier = 'N/A'

[preprocess]
# Clean all text columns
[preprocess.texts]
operations = [
    {op = "clean", method = "trimws"}
]

# Fill missing values in numeric columns
[preprocess.numerics]
operations = [
    {op = "fillna", method = "median"}
]

# Specific handling for age
[[preprocess.columns]]
name = 'age'
type = 'int'
operations = [
    {op = "fillna", method = "mean"}
]

# Encode categorical variable
[[preprocess.columns]]
name = 'gender'
type = 'string'
operations = [
    {op = "dummy", dummy_prefix = true}
]

[postprocess]
format = 'csv'
filename = 'survey_data_cleaned.csv'
dropcolumns = ["internal_id", "notes"]
```

### Example 2: Machine Learning Pipeline

```toml
[data]
filename = './training_data.csv'
csv_separator = ','
decimal_separator = '.'
encoding = 'utf-8'

[preprocess]
# Handle missing values in all numeric features
[preprocess.numerics]
operations = [
    {op = "fillna", method = "median"}
]
exclude_columns = ["target"]

# Scale all numeric features
[preprocess.numerics]
operations = [
    {op = "scale", method = "zscore"}
]
exclude_columns = ["target"]

# Encode categorical variables
[[preprocess.columns]]
name = 'category'
type = 'string'
operations = [
    {op = "dummy", dummy_droplast = true}
]

[[preprocess.columns]]
name = 'region'
type = 'string'
operations = [
    {op = "dummy", dummy_prefix = true}
]

[postprocess]
format = 'csv'
filename = 'training_data_processed.csv'

[postprocess.dataset_split]
method = "train_test_split"
split_names = ["train", "test"]
random_seed = 42
train_test_split_ratio = 0.8
```

### Example 3: European Data with Specific Operations

```toml
[data]
filename = './european_sales.csv'
csv_separator = ';'
decimal_separator = ','
encoding = 'latin-1'
missing_identifier = ''

[preprocess]
# Clean text
[preprocess.texts]
operations = [
    {op = "clean", method = "trimws"},
    {op = "clean", method = "title"}
]

# Process numeric columns
[preprocess.numerics]
operations = [
    {op = "fillna", method = "mean"}
]

# Specific transformations
[[preprocess.columns]]
name = 'date'
type = 'string'
new_name = 'transaction_date'
operations = [
    {op = "clean", method = "trimws"}
]

[[preprocess.columns]]
name = 'amount'
type = 'float'
operations = [
    {op = "fillna", method = "median"},
    {op = "scale", method = "minmax"}
]

[postprocess]
format = 'csv'
filename = 'european_sales_cleaned.csv'
```

### Example 4: Data Migration Prepfile

```toml
[data]
filename = './legacy_export.txt'
csv_separator = '|'
decimal_separator = '.'
encoding = 'latin-1'
missing_identifier = 'NULL'

[preprocess]
# Standardize all text
[preprocess.texts]
operations = [
    {op = "clean", method = "trimws"},
    {op = "clean", method = "upper"}
]

# Rename columns to match new system
[[preprocess.columns]]
name = 'cust_id'
type = 'string'
new_name = 'customer_id'

[[preprocess.columns]]
name = 'ord_date'
type = 'string'
new_name = 'order_date'

[[preprocess.columns]]
name = 'ord_amt'
type = 'float'
new_name = 'order_amount'
operations = [
    {op = "fillna", value = 0}
]

[postprocess]
format = 'csv'
filename = 'migrated_data.csv'
dropcolumns = ["old_id", "legacy_code"]
```

---

## Best Practices

### 1. Start Simple

Begin with basic operations and add complexity incrementally:

```toml
# Step 1: Basic structure
[data]
filename = './data.csv'

[preprocess]

[postprocess]

# Step 2: Add simple operations
[preprocess.numerics]
operations = [
    {op = "fillna", method = "median"}
]

# Step 3: Add more operations
```

### 2. Use Comments

Document your preprocessing decisions:

```toml
[preprocess]
# Data cleaning phase - handles missing values
[preprocess.numerics]
operations = [
    {op = "fillna", method = "median"}  # Using median as it's robust to outliers
]

# Feature scaling for machine learning
[preprocess.numerics]
operations = [
    {op = "scale", method = "zscore"}  # Standardization for SVM algorithm
]
```

### 3. Test Incrementally

Add one operation at a time and verify the results:

```bash
# Add one operation
preprocess run --file prepfile_v1.toml
preprocess skim --data ./output.csv

# Add another operation
# Edit prepfile to prepfile_v2.toml
preprocess run --file prepfile_v2.toml
preprocess skim --data ./output.csv
```

### 4. Use Version Control

Track changes to your Prepfile:

```bash
git add my_prepfile.toml
git commit -m "Add scaling to numeric features"
```

### 5. Validate with Diff

Always verify your preprocessing worked as expected:

```bash
preprocess run --file my_prep.toml
preprocess diff --source ./original.csv --target ./cleaned.csv
```

### 6. Organize for Complex Pipelines

For complex preprocessing, use multiple Prepfiles:

```
.
├── prepfiles/
│   ├── step1_cleaning.toml
│   ├── step2_feature_engineering.toml
│   └── step3_final_prep.toml
└── scripts/
    ├── run_pipeline.sh
    └── verify_results.sh
```

### 7. Share with Team

Prepfiles are text files - easy to share and discuss:

```bash
# Share with colleague
cp my_prepfile.toml ../colleague_project/

# Or via email/git/etc.
```

---

## Common Patterns

### Pattern 1: Data Cleaning Pipeline

```toml
[preprocess]
# Step 1: Handle missing values
[preprocess.numerics]
operations = [
    {op = "fillna", method = "median"}
]

[preprocess.texts]
operations = [
    {op = "fillna", value = "Unknown"}
]

# Step 2: Clean text
[preprocess.texts]
operations = [
    {op = "clean", method = "trimws"},
    {op = "clean", method = "lower"}
]

# Step 3: Scale numeric
[preprocess.numerics]
operations = [
    {op = "scale", method = "zscore"}
]
```

### Pattern 2: Feature Engineering

```toml
[preprocess]
# Encode categorical variables
[[preprocess.columns]]
name = 'category'
type = 'string'
operations = [
    {op = "dummy", dummy_prefix = true}
]

# Scale numeric features
[preprocess.numerics]
exclude_columns = ["target"]
operations = [
    {op = "scale", method = "zscore"}
]

# Create bins for continuous variables
[[preprocess.columns]]
name = 'age'
type = 'int'
operations = [
    {op = "discretize", method = "binning", bins = [
        {lower = 0, upper = 18, label = "0-18"},
        {lower = 18, upper = 35, label = "18-35"},
        {lower = 35, upper = 60, label = "35-60"},
        {lower = 60, upper = 100, label = "60+"}
    ]}
]
```

### Pattern 3: Data Standardization

```toml
[preprocess]
# Rename columns
[[preprocess.columns]]
name = 'old_name'
type = 'string'
new_name = 'standard_name'

# Standardize text
[preprocess.texts]
operations = [
    {op = "clean", method = "trimws"},
    {op = "clean", method = "title"}
]

# Standardize numeric
[preprocess.numerics]
operations = [
    {op = "fillna", method = "median"},
    {op = "scale", method = "minmax"}
]
```

---

## Troubleshooting Prepfiles

### Common Errors

**Error: Failed to load config file**
```
Error: Failed to load config file 'Prepfile.toml': <error details>
```

**Causes and Solutions:**
- **Syntax error**: Check TOML syntax (use a TOML validator)
- **File not found**: Verify path and filename
- **Invalid values**: Ensure all values are valid for their types

**Error: Column not found**
```
Error: Column 'column_name' not found in dataset
```

**Solutions:**
- Check column name spelling (case-sensitive!)
- Verify column exists in your dataset
- Use `preprocess skim --data ./your_data.csv` to see actual column names

**Error: Operation not applicable**
```
Error: Operation 'fillna:mean' not applicable to column type 'string'
```

**Solution**: Ensure you're applying the right operations to the right column types:
- Numeric operations (`fillna:mean`, `scale`) only work on numeric columns
- Text operations (`clean`, `dummy`) only work on text columns

**Error: Missing required parameter**
```
Error: Missing required parameter 'filename' in [data] section
```

**Solution**: Add the required parameter. All Prepfiles must have:
```toml
[data]
filename = './your_data.csv'  # This is required
```

### Validation Tips

1. **Use TOML linter**: Validate your TOML syntax with an online linter
2. **Start minimal**: Begin with just the `[data]` section and test
3. **Add incrementally**: Add one section at a time and test each
4. **Check types**: Ensure column types match actual data types

---

## Prepfile Templates

### Template 1: Quick Start

```toml
[data]
filename = './your_data.csv'
csv_separator = ','
decimal_separator = '.'
encoding = 'utf-8'
missing_identifier = ''

[preprocess]
[preprocess.numerics]
operations = [
    {op = "fillna", method = "median"}
]

[preprocess.texts]
operations = [
    {op = "clean", method = "trimws"}
]

[postprocess]
format = 'csv'
filename = 'your_data_cleaned.csv'
```

### Template 2: Machine Learning

```toml
[data]
filename = './training_data.csv'
csv_separator = ','
decimal_separator = '.'
encoding = 'utf-8'

[preprocess]
[preprocess.numerics]
exclude_columns = ["id", "target"]
operations = [
    {op = "fillna", method = "median"},
    {op = "scale", method = "zscore"}
]

[preprocess.texts]
operations = [
    {op = "clean", method = "trimws"},
    {op = "clean", method = "lower"}
]

[postprocess]
format = 'csv'
filename = 'training_data_processed.csv'

[postprocess.dataset_split]
method = "train_test_split"
split_names = ["train", "test"]
random_seed = 42
train_test_split_ratio = 0.8
```

### Template 3: Data Cleaning

```toml
[data]
filename = './raw_data.csv'
csv_separator = ','
decimal_separator = '.'
encoding = 'utf-8'
missing_identifier = 'NA'

[preprocess]
# Handle missing values
[preprocess.numerics]
operations = [
    {op = "fillna", method = "median"}
]

[preprocess.texts]
operations = [
    {op = "fillna", value = "Unknown"}
]

# Clean text
[preprocess.texts]
operations = [
    {op = "clean", method = "trimws"},
    {op = "clean", method = "title"}
]

# Column-specific operations
[[preprocess.columns]]
name = 'date'
type = 'string'
operations = [
    {op = "clean", method = "trimws"}
]

[postprocess]
format = 'csv'
filename = 'cleaned_data.csv'
dropcolumns = ["temp1", "temp2"]
```

---

## Next Steps

Now that you understand Prepfile:

1. **[Create your first Prepfile]({{% relref "init" %}})**: Use `preprocess init` to generate a starting point
2. **[Run preprocessing]({{% relref "run" %}})**: Execute your preprocessing pipeline
3. **[Verify results]({{% relref "diff" %}})**: Use `preprocess diff` to check your preprocessing worked
4. **[Explore examples]({{% relref "examples" %}})**: See practical examples of preprocessing workflows
5. **[Check templates]({{% relref "templates" %}})**: Use pre-built templates for common scenarios
