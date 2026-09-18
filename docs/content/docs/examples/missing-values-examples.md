---
weight: 200
title: "Handling Missing Values - Real-World Scenarios"
icon: "query_stats"
date: "2026-09-18T12:00:00+02:00"
lastmod: "2026-09-18T12:00:00+02:00"
description: "Practical guide to handling missing values across different dataset types with concrete examples"
draft: false
toc: true
---

## Handling Missing Values - Real-World Scenarios

Missing data is a **universal challenge** in data preprocessing. This guide presents **concrete examples** from three different datasets, showing you how to handle missing values effectively using preprocess CLI.

---

## The Importance of Missing Value Handling

Missing values can:
- ❌ **Bias your analysis**: If not random, missing data can skew results
- ❌ **Reduce statistical power**: Fewer observations = less reliable conclusions
- ❌ **Break algorithms**: Many ML algorithms cannot handle NULL/NaN values

Proper imputation ensures:
- ✅ **Complete datasets** for analysis and modeling
- ✅ **Consistent results** across different runs
- ✅ **Better data quality** for downstream tasks

---

## Dataset Overview

We'll work with three real datasets, each presenting different missing value challenges:

| Dataset | Rows | Missing Values | Type | Challenge |
|---------|------|----------------|------|-----------|
| **FIFA Players 2022** | 16,020 | 170+ | Sports | High-value outliers (salaries) |
| **HDV 2003** | 1,000+ | 127+ | Social Survey | Categorical and numeric mixed |
| **Indicators** | 209 | Few | Economic | Multi-scale numeric data |

---

## Example 1: FIFA Dataset - Imputation with Median (Robust to Outliers)

### The Challenge

The FIFA dataset has **85 missing values** in two critical columns:
- `value_eur`: Player market value (€0 - €167,500,000)
- `wage_eur`: Weekly salary (€0 - €565,000)

**Problem**: These columns have **extreme outliers**. The maximum value (€167.5M) is 45x the median (€3.6M). Using the mean for imputation would be heavily influenced by these outliers.

### Statistics Before Imputation

```
value_eur:
  rows_count: 15,935  (85 missing)
  min: 0
  max: 167,500,000
  mean: 3,690,000
  median: 3,600,000
  std: 12,500,000 (very high due to outliers)

wage_eur:
  rows_count: 15,935  (85 missing)
  min: 0
  max: 565,000
  mean: 26,000
  median: 16,000
```

### Solution: Use Median Imputation

```toml
[preprocess.numerics]
operations = [
    {op = "fillna", method = "median"}
]
```

**Why median?**
- **Robust to outliers**: Unlike mean, median is not affected by extreme values
- **Preserves distribution**: Maintains the central tendency of the data
- **Common practice**: Standard approach for imputing missing values in skewed distributions

### After Imputation

```
value_eur:
  rows_count: 16,020  (0 missing)
  min: 0
  max: 167,500,000
  median: 3,600,000 (unchanged)
  
  The 85 missing values are now filled with 3,600,000
```

**✅ Result**: All missing values replaced with the median, preserving the dataset's statistical properties.

### Verification with `diff`

```bash
preprocess diff \
  --source ./docs-dataset/fifa_players_22.csv \
  --target fifa_imputed.csv \
  --output fifa_diff.html
```

**What you'll see**: The 85 rows with NULL in `value_eur` and `wage_eur` now have the value 3,600,000 and 16,000 respectively.

---

## Example 2: HDV Dataset - Type-Specific Imputation

### The Challenge

The HDV social survey dataset has **mixed data types** with missing values:
- **Numeric columns**: `heures.tv`, `age`, etc.
- **Text/categorical columns**: `relig`, `sexe`, `nivetud`, etc.

Different types require **different imputation strategies**.

### Numeric Columns: Median Imputation

**Column**: `heures.tv` (hours of TV per week)
- **Missing**: 2 values
- **Range**: 0-28 hours

**Solution**:
```toml
[preprocess.numerics]
operations = [
    {op = "fillna", method = "median"}
]
```

**Alternative for `heures.tv`**: Use `value = 0` since 0 hours of TV is a logical default:
```toml
[[preprocess.columns]]
name = 'heures.tv'
type = 'float'
operations = [
    {op = "fillna", value = 0}
]
```

**Recommendation**: Use `median` for most numeric columns, but consider `value = 0` when it makes semantic sense (like hours watched).

### Text/Categorical Columns: Custom Value Imputation

**Column**: `sexe` (gender)
- **Missing**: Several values
- **Categories**: "Homme", "Femme"

**Solution**:
```toml
[preprocess.texts]
operations = [
    {op = "fillna", value = "non spécifié"}
]
```

**Why "non spécifié"?**
- **Preserves information**: We know this value was missing, not just another category
- **Avoids misclassification**: Doesn't force the missing value into "Homme" or "Femme"
- **Analytical clarity**: Can filter out "non spécifié" in analysis if needed

**Other categorical columns** (`relig`, `nivetud`, `occup`, `qualif`):
```toml
[preprocess.texts]
operations = [
    {op = "fillna", value = "non spécifié"}
]
```

### Complete HDV Imputation Prepfile

```toml
[data]
filename = './docs-dataset/hdv2003.csv'
missing_identifier = 'NA'

[preprocess]

# Numeric columns: use median
[preprocess.numerics]
operations = [
    {op = "fillna", method = "median"}
]

# Text columns: use custom value
[preprocess.texts]
operations = [
    {op = "fillna", value = "non spécifié"}
]

[postprocess]
filename = 'hdv_imputed.csv'
```

---

## Example 3: Indicators Dataset - Median for Economic Data

### The Challenge

Economic indicators often have:
- **Different scales**: GDP (2K-119K) vs BALANCE (-5.7B to 7.1B)
- **Few missing values**: But each one is critical
- **Skewed distributions**: Economic data often follows power-law distributions

**Dataset**: 209 countries, 26 indicators
- Most columns have **0-5 missing values**
- Some columns have **no missing values**

### Statistics Example

**GDP_CAPITA** (GDP per capita):
```
rows_count: 208 (1 missing)
min: 2,082
max: 119,367
mean: 34,827.56
median: 20,000
```

**Solution**: Use median for all numeric columns
```toml
[preprocess.numerics]
operations = [
    {op = "fillna", method = "median"}
]
```

**Why median works well for economic data**:
1. **Robust to outliers**: Economic indicators often have extreme values
2. **Preserves central tendency**: Mean can be misleading with skewed data
3. **Consistent across scales**: Works well regardless of the indicator's magnitude

---

## Comparison: Imputation Methods

| Method | Use Case | Pros | Cons | Best For |
|--------|----------|------|------|----------|
| `median` | Numeric, skewed data | Robust to outliers | May not preserve variance | FIFA, Indicators |
| `mean` | Numeric, symmetric data | Preserves mean | Sensitive to outliers | Normally distributed data |
| `value` | All types | Simple, semantic | Arbitrary choice | Categorical, logical defaults |

---

## Mean vs Median: When to Use Each

### Use `median` when:
- ✅ Data has **outliers** (salaries, GDP, house prices)
- ✅ Distribution is **skewed**
- ✅ You want **robust** imputation
- ✅ FIFA dataset, Indicators dataset

### Use `mean` when:
- ✅ Data is **normally distributed**
- ✅ No extreme outliers
- ✅ You want to preserve the **mean** of the dataset
- ✅ Example: Age in FIFA dataset (16-39, relatively symmetric)

### Use `value` when:
- ✅ You need a **semantic default** (0 hours of TV)
- ✅ For **categorical data** ("non spécifié")
- ✅ When business logic dictates a specific value

---

## Before/After Comparison Table

### FIFA Dataset

| Column | Missing Before | Imputation Method | Missing After | Value Used |
|--------|----------------|-------------------|---------------|------------|
| value_eur | 85 | median | 0 | 3,600,000 € |
| wage_eur | 85 | median | 0 | 16,000 € |

### HDV Dataset

| Column | Type | Missing Before | Imputation | Missing After | Value Used |
|--------|------|----------------|------------|---------------|------------|
| heures.tv | numeric | 2 | value | 0 | 0 |
| sexe | text | 5 | value | 0 | "non spécifié" |
| relig | text | 8 | value | 0 | "non spécifié" |
| age | numeric | 0 | - | 0 | - |

### Indicators Dataset

| Column | Missing | Imputation | Result |
|--------|---------|------------|--------|
| GDP_CAPITA | 1 | median | 20,000 |
| BALANCE | 0 | - | - |
| TRADE_BALANCE | 2 | median | Median value |

---

## Advanced: Excluding Specific Columns from Imputation

Sometimes you want to **exclude certain columns** from bulk imputation.

### Example: FIFA Dataset - Exclude Target Variable

When preparing data for machine learning, you might want to:
1. Impute all numeric features with median
2. But **exclude** the target variable (`value_eur`) from imputation
3. Handle the target separately

**Solution**:
```toml
[preprocess.numerics]
exclude_columns = ["value_eur"]
operations = [
    {op = "fillna", method = "median"}
]

# Handle target separately
[[preprocess.columns]]
name = 'value_eur'
type = 'float'
operations = [
    {op = "fillna", method = "median"}
]
```

Or impute with a different strategy:
```toml
[[preprocess.columns]]
name = 'value_eur'
type = 'float'
operations = [
    {op = "fillna", value = 0}  # Or use median
]
```

---

## Practical Recommendations by Dataset Type

### Sports/Data with Outliers (FIFA-like)

**Strategy**:
1. **Use median** for all numeric columns with outliers
2. **Verify distribution** with `summary` command first
3. **Exclude target** if doing ML, handle separately

**Prepfile template**:
```toml
[preprocess.numerics]
exclude_columns = ["target_column"]
operations = [
    {op = "fillna", method = "median"}
]

[[preprocess.columns]]
name = 'target_column'
type = 'float'
operations = [
    {op = "fillna", method = "median"}
]
```

### Survey Data (HDV-like)

**Strategy**:
1. **Numeric columns**: median imputation
2. **Text columns**: "non spécifié" or similar
3. **Set `missing_identifier`**: Ensure 'NA' is recognized

**Prepfile template**:
```toml
[data]
missing_identifier = 'NA'

[preprocess.numerics]
operations = [
    {op = "fillna", method = "median"}
]

[preprocess.texts]
operations = [
    {op = "fillna", value = "non spécifié"}
]
```

### Economic Data (Indicators-like)

**Strategy**:
1. **Use median** for all numeric columns (robust to scale differences)
2. **Consider feature selection** after imputation
3. **Be aware of**: Very different scales across indicators

**Prepfile template**:
```toml
[preprocess.numerics]
operations = [
    {op = "fillna", method = "median"}
]
```

---

## Verification Workflow

Always verify your imputation results:

### Step 1: Check Summary Before
```bash
preprocess summary --data ./docs-dataset/fifa_players_22.csv --output fifa_before.toml
```

### Step 2: Apply Imputation
```bash
preprocess run --file fifa_imputation.toml
```

### Step 3: Check Summary After
```bash
preprocess summary --data fifa_imputed.csv --output fifa_after.toml
```

### Step 4: Compare
```bash
# Check the TOML files
# Missing counts should be 0 for all columns you imputed
```

### Step 5: Visual Diff
```bash
preprocess diff \
  --source ./docs-dataset/fifa_players_22.csv \
  --target fifa_imputed.csv \
  --output fifa_imputation_diff.html
```

---

## Common Mistakes to Avoid

### Mistake 1: Using Mean for Skewed Data
**Problem**: Using mean for salary data (with outliers at 167M€) will skew your imputation.
**Solution**: Always check distribution first with `summary`, use median for skewed data.

### Mistake 2: Imputing Categorical with Mean/Median
**Problem**: Trying to use `method = "median"` on text columns.
**Solution**: Use `value` parameter for text columns: `{op = "fillna", value = "default"}`

### Mistake 3: Forgetting to Set Missing Identifier
**Problem**: Missing values are not recognized if `missing_identifier` is not set correctly.
**Solution**: Set in Prepfile: `missing_identifier = 'NA'` or `'NULL'` or whatever your dataset uses.

### Mistake 4: Imputing Without Verification
**Problem**: Applying imputation without checking the results.
**Solution**: Always use `diff` or `summary` to verify imputation worked as expected.

---

## Complete Examples

### FIFA Complete Imputation Pipeline

```toml
[data]
filename = './docs-dataset/fifa_players_22.csv'

[preprocess]

# Impute all numeric columns with median
[preprocess.numerics]
operations = [
    {op = "fillna", method = "median"}
]

# Clean text columns
[preprocess.texts]
operations = [
    {op = "clean", method = "trimws"}
]

[postprocess]
filename = 'fifa_imputed.csv'
```

### HDV Complete Imputation Pipeline

```toml
[data]
filename = './docs-dataset/hdv2003.csv'
missing_identifier = 'NA'

[preprocess]

# Numeric: median
[preprocess.numerics]
operations = [
    {op = "fillna", method = "median"}
]

# Text: custom value
[preprocess.texts]
operations = [
    {op = "fillna", value = "non spécifié"}
]

# Clean text
[preprocess.texts]
operations = [
    {op = "clean", method = "trimws"},
    {op = "clean", method = "lower"}
]

[postprocess]
filename = 'hdv_imputed.csv'
```

---

## Key Takeaways

1. **Median is your friend**: Use it for numeric data with outliers (most common case)
2. **Mean has its place**: Use for normally distributed numeric data
3. **Custom values for text**: Always use semantic defaults for categorical data
4. **Verify, verify, verify**: Use `summary` and `diff` to confirm imputation worked
5. **Set missing_identifier**: Crucial for proper NA detection
6. **Exclude when needed**: Use `exclude_columns` for special handling

---

## Identified Limitations and Future Needs

While preprocess CLI handles most missing value scenarios well, we identified these **feature requests** during analysis:

| Feature | Use Case | Priority |
|---------|----------|----------|
| `winsorize` | Cap outliers before imputation | ⭐⭐⭐ |
| `transform:log1p` | Log-transform for exponential distributions | ⭐⭐⭐⭐ |
| Conditional imputation | Different strategies based on conditions | ⭐⭐ |

---

## Related Articles

- [Text Cleaning - Practical Examples](/docs/examples/text-cleaning-examples/)
- [Feature Scaling for ML - Complete Guide](/docs/examples/feature-scaling-examples/)
- [FIFA Dataset Complete Analysis](/docs/examples/fifa-analysis/)

---

*Last updated: September 18, 2026*
*Based on analysis of FIFA Players 2022, HDV 2003, and Indicators datasets*
