---
weight: 350
title: "Prepfile Templates"
description: "Ready-to-use templates for common preprocessing scenarios"
icon: "pattern"
date: "2025-06-01T20:10:40+02:00"
lastmod: "2026-09-18T11:30:00+02:00"
draft: false
--- 

## Prepfile Templates

Jumpstart your data preprocessing with **ready-to-use templates**. These templates provide starting points for common data preprocessing scenarios that Data Scientists and Data Analysts encounter regularly.

## Available Templates

{{% table "table-hover" %}}
| Template | Description | Best For |
|----------|-------------|----------|
| [Simple Template]({{% relref "simple_template" %}}) | Basic preprocessing with column renaming | Beginners, simple cleaning |
| [Numeric Operations]({{% relref "numeric_operations" %}}) | Focused on numeric column operations | Feature scaling, missing value handling |
{{% /table %}}

## How to Use Templates

### Quick Start with Templates

1. **Find a template** that matches your use case
2. **Copy the configuration** to your Prepfile
3. **Customize** for your specific dataset
4. **Run** with `preprocess run --file your_prepfile.toml`

### Example Workflow

```bash
# Step 1: Copy a template
cp templates/numeric_operations.toml my_project_prep.toml

# Step 2: Edit for your dataset
nano my_project_prep.toml
# Update filename, adjust operations, etc.

# Step 3: Run preprocessing
preprocess run --file my_project_prep.toml

# Step 4: Verify results
preprocess diff --source ./original.csv --target ./output.csv
```

## Template Categories

### Beginner Templates

Start here if you're new to preprocess CLI:
- **[Simple Template]({{% relref "simple_template" %}})**: Basic operations and column renaming

### Intermediate Templates

For users familiar with the basics:
- **[Numeric Operations]({{% relref "numeric_operations" %}})**: Common numeric preprocessing tasks

### Advanced Templates

Combine multiple templates for complex workflows.

## Creating Custom Templates

### Best Practices for Template Creation

1. **Start with a working Prepfile**: Ensure it runs correctly
2. **Remove dataset-specific details**: Make it reusable
3. **Add clear comments**: Explain what each part does
4. **Document assumptions**: Note any requirements
5. **Test thoroughly**: Verify it works with different datasets

### Template Structure

A good template follows this structure:

```toml
# ============================================
# Template: Description of what this template does
# Use Case: When to use this template
# Requirements: Any prerequisites or assumptions
# ============================================

[data]
filename = './your_data.csv'  # REPLACE with your filename
csv_separator = ','
decimal_separator = '.'
encoding = 'utf-8'
missing_identifier = ''

[preprocess]
# Section 1: Description of these operations
[preprocess.numerics]
operations = [
    {op = "fillna", method = "median"}  # Comment explaining this operation
]

# Section 2: Description of these operations
[preprocess.texts]
operations = [
    {op = "clean", method = "trimws"}
]

[postprocess]
format = 'csv'
filename = 'your_data_cleaned.csv'  # REPLACE with desired output filename
```

### Template Documentation

When creating templates, include documentation:

```markdown
## Template: [Template Name]

### Description
Brief explanation of what this template accomplishes.

### Use Case
When to use this template:
- Scenario 1
- Scenario 2
- Scenario 3

### Requirements
- Dataset format: CSV/TSV/etc.
- Column types: Expected column types
- Encoding: Expected encoding

### Customization
How to adapt this template:
1. Update `filename` in [data] section
2. Adjust operations as needed
3. Modify [postprocess] for desired output

### Example
```toml
[Template configuration here]
```
```

## Template Gallery

### Machine Learning Templates

**Coming Soon**: Templates for common ML preprocessing tasks:
- Classification data preparation
- Regression data preparation
- Time series data preprocessing
- Image data preprocessing (if supported)

### Data Cleaning Templates

**Coming Soon**: Templates for data cleaning scenarios:
- Missing value handling
- Outlier detection and treatment
- Data standardization
- Data migration

### Industry-Specific Templates

**Coming Soon**: Templates tailored for specific industries:
- Healthcare data preprocessing
- Financial data preprocessing
- Retail data preprocessing
- Social media data preprocessing

## Contributing Templates

### Share Your Templates

We welcome contributions of useful templates:

1. **Create a template** that solves a common problem
2. **Test it thoroughly** with various datasets
3. **Document it clearly** with use cases and requirements
4. **Submit a pull request** to the documentation repository

### Template Contribution Guidelines

- **Name clearly**: Use descriptive names
- **Document thoroughly**: Explain use cases and requirements
- **Keep simple**: Focus on one specific use case per template
- **Test**: Ensure the template works correctly
- **Follow conventions**: Use consistent formatting and style

## Template Examples

### Example: Creating a Template from Scratch

Let's create a template for **customer data preprocessing**:

```toml
# ============================================
# Template: Customer Data Cleaning
# Use Case: Cleaning customer datasets for analysis
# Requirements: CSV file with customer data
# Assumes: Standard column names (can be customized)
# ============================================

[data]
filename = './customers.csv'
csv_separator = ','
decimal_separator = '.'
encoding = 'utf-8'
missing_identifier = 'N/A'

[preprocess]
# Clean all text fields
[preprocess.texts]
operations = [
    {op = "clean", method = "trimws"},  # Remove whitespace
    {op = "clean", method = "title"}   # Standardize case
]

# Handle missing numeric values
[preprocess.numerics]
operations = [
    {op = "fillna", method = "median"}  # Fill with median
]

# Specific handling for age column
[[preprocess.columns]]
name = 'age'
type = 'int'
operations = [
    {op = "fillna", value = 0}  # Fill missing ages with 0
]

# Encode gender column
[[preprocess.columns]]
name = 'gender'
type = 'string'
operations = [
    {op = "dummy", dummy_prefix = true}
]

[postprocess]
format = 'csv'
filename = 'customers_cleaned.csv'
dropcolumns = ["internal_id", "temp_column"]
```

### Example: Machine Learning Template

```toml
# ============================================
# Template: Machine Learning Data Preparation
# Use Case: Preparing data for ML algorithms
# Requirements: CSV file with features and target
# Note: Assumes 'target' column exists
# ============================================

[data]
filename = './training_data.csv'
csv_separator = ','
decimal_separator = '.'
encoding = 'utf-8'

[preprocess]
# Handle missing values in features (not target)
[preprocess.numerics]
exclude_columns = ["target", "id"]
operations = [
    {op = "fillna", method = "median"}
]

# Scale all numeric features
[preprocess.numerics]
exclude_columns = ["target", "id"]
operations = [
    {op = "scale", method = "zscore"}
]

# Encode categorical variables
[[preprocess.columns]]
name = 'category'
type = 'string'
operations = [
    {op = "dummy", dummy_droplast = true}
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

## Tips for Using Templates

### 1. Start with the Closest Match

Find the template that most closely matches your use case, then customize it.

### 2. Understand Before Customizing

Read through the template and understand what each part does before modifying it.

### 3. Test Incrementally

When customizing a template:
1. Make one change at a time
2. Test that it works as expected
3. Continue with next change

### 4. Save Your Customizations

Save your customized templates for future use:

```bash
# Save as a new template
cp my_custom_pipeline.toml templates/my_custom_template.toml
```

### 5. Document Your Changes

Add comments to explain your customizations:

```toml
# Customized from: simple_template.toml
# Changes made:
# - Added scaling for numeric features
# - Changed missing value handling to mean
[preprocess.numerics]
operations = [
    {op = "fillna", method = "mean"}  # Changed from median to mean
]
```

## Next Steps

After exploring templates:

1. **[Try a template]({{% relref "simple_template" %}})** with your own data
2. **[Create your own]({{% relref "prepreference" %}})** custom templates
3. **[Share with the community]** by contributing to the documentation
4. **[Explore examples]({{% relref "examples" %}})** for more practical use cases
