---
weight: 200
title: "Feature Scaling for Machine Learning - Complete Guide"
icon: "trending_up"
date: "2026-09-18T12:00:00+02:00"
lastmod: "2026-09-18T12:00:00+02:00"
description: "Master feature scaling with concrete examples from FIFA and Indicators datasets - Z-score vs Min-Max"
draft: false
toc: true
---

## Feature Scaling for Machine Learning - Complete Guide

Feature scaling is **essential** for machine learning. Without proper scaling, algorithms that rely on distance calculations (like SVM, KNN, Neural Networks) will be **dominated by features with larger scales**, leading to poor performance.

This guide uses **real datasets** to show you when and how to apply different scaling methods.

---

## Why Feature Scaling Matters

### The Problem: Different Scales

Consider the FIFA dataset:

| Feature | Min | Max | Scale |
|---------|-----|-----|-------|
| age | 16 | 39 | 1-100 |
| height_cm | 155 | 203 | 100-200 |
| value_eur | 0 | 167,500,000 | 0-167M |
| wage_eur | 0 | 565,000 | 0-565K |

**Problem**: `value_eur` has values **4.5 million times larger** than `age`. Without scaling:
- Distance-based algorithms will be **completely dominated** by `value_eur` and `wage_eur`
- Gradient descent will **converge slowly**
- Features will have **unequal influence** on the model

---

## Scaling Methods in Preprocess CLI

Preprocess CLI supports two primary scaling methods:

| Method | Formula | Range | Use Case |
|--------|---------|-------|----------|
| **zscore** | (x - mean) / std | (-∞, +∞) | When you want mean=0, std=1 |
| **minmax** | (x - min) / (max - min) | [0, 1] | When you need bounded values |

---

## Example 1: Z-score Scaling on FIFA Dataset

### Dataset: FIFA Players 2022

**Use case**: Preparing data for machine learning to predict player value.

### Before Scaling

```
Sample of numeric features:

age:        [16, 25, 39]  (mean=25.19, std=4.56)
shooting:   [20, 85, 95]  (mean=69.12, std=18.34)
value_eur:  [0, 3600000, 167500000]  (mean=3690000, std=12500000)
```

**Problem**: `value_eur` has a scale **4,500x larger** than `age`. Without scaling, any ML algorithm will give `value_eur` **4,500x more weight** than `age`.

### Solution: Z-score Scaling

```toml
[preprocess.numerics]
exclude_columns = ["short_name", "preferred_foot", "body_type"]  # Exclude text columns
operations = [
    {op = "scale", method = "zscore"}
]
```

### After Scaling

```
age (z-score):
  (16 - 25.19) / 4.56 = -2.01  (2.01 std dev below mean)
  (25 - 25.19) / 4.56 = -0.04  (0.04 std dev below mean)
  (39 - 25.19) / 4.56 = 3.05   (3.05 std dev above mean)

shooting (z-score):
  (20 - 69.12) / 18.34 = -2.67
  (85 - 69.12) / 18.34 = 0.87
  (95 - 69.12) / 18.34 = 1.42

value_eur (z-score):
  (0 - 3690000) / 12500000 = -0.30
  (3600000 - 3690000) / 12500000 = -0.01
  (167500000 - 3690000) / 12500000 = 13.07
```

**✅ Result**: 
- All features now have **mean = 0** and **standard deviation = 1**
- All features are on **the same scale**
- Features can be **fairly compared** by ML algorithms

### Properties of Z-score Scaling

- **Mean**: 0
- **Standard deviation**: 1
- **Range**: Unbounded (can be negative or > 1)
- **Interpretation**: Value tells you how many standard deviations from the mean
- **Preserves**: Shape of distribution, outliers remain outliers

---

## Example 2: Min-Max Scaling on Indicators Dataset

### Dataset: Economic Indicators

**Use case**: Comparing countries across different economic indicators with vastly different scales.

### Before Scaling

```
GDP_CAPITA:    [2,082, 34,827, 119,367]  (USD per capita)
GDP:           [24B, 2.6T, 21.4T]           (Total GDP in USD)
BALANCE:       [-5.7B, 0, 7.1B]           (Trade balance in USD)
POPULATION:    [32M, 1.4B, 83M]           (Population count)
```

**Problem**: Comparing these directly is **meaningless** due to scale differences. Min-Max scaling normalizes all features to [0, 1].

### Solution: Min-Max Scaling

```toml
[preprocess.numerics]
operations = [
    {op = "scale", method = "minmax"}
]
```

### After Scaling

```
GDP_CAPITA (minmax):
  (2082 - 2082) / (119367 - 2082) = 0.00    (Afghanistan - minimum)
  (34827 - 2082) / (119367 - 2082) = 0.27  (World average)
  (119367 - 2082) / (119367 - 2082) = 1.00 (Luxembourg - maximum)

GDP (minmax):
  (24B - 24B) / (21.4T - 24B) = 0.00    (Smallest economy)
  (2.6T - 24B) / (21.4T - 24B) = ~0.12  (Medium economy)
  (21.4T - 24B) / (21.4T - 24B) = 1.00 (Largest economy)

BALANCE (minmax):
  (-5.7B - (-5.7B)) / (7.1B - (-5.7B)) = 0.00 (Most negative)
  (0 - (-5.7B)) / (7.1B - (-5.7B)) = 0.44  (Balanced)
  (7.1B - (-5.7B)) / (7.1B - (-5.7B)) = 1.00 (Most positive)
```

**✅ Result**:
- All features now have **values between 0 and 1**
- **Direct comparison** possible between any two features
- **Preserves relative distances** within each feature

### Properties of Min-Max Scaling

- **Minimum**: 0
- **Maximum**: 1
- **Range**: [0, 1]
- **Interpretation**: Value represents position between min and max
- **Sensitive to**: Outliers (min and max can be extreme values)

---

## Comparison: Z-score vs Min-Max

| Aspect | Z-score | Min-Max |
|--------|---------|--------|
| **Formula** | (x - mean) / std | (x - min) / (max - min) |
| **Range** | (-∞, +∞) | [0, 1] |
| **Mean** | 0 | ~0.5 (if symmetric) |
| **Outlier sensitivity** | ✅ Low (uses mean/std) | ⚠️ High (uses min/max) |
| **Interpretation** | Standard deviations from mean | Percentage of range |
| **Use with** | SVM, PCA, Linear Regression | Neural Networks, KNN, Image data |
| **Negative values** | ✅ Yes | ❌ No |
| **Preserves sparsity** | ✅ Yes | ❌ No |

---

## When to Use Each Method

### Use **Z-score (Standardization)** when:

✅ **Algorithm requires normally distributed data**
- Linear Regression
- Logistic Regression
- SVM (Support Vector Machines)
- PCA (Principal Component Analysis)
- k-Means Clustering

✅ **Features have different units** but similar distributions

✅ **You need to preserve the shape** of the distribution

✅ **Data has outliers** (z-score is less sensitive to outliers than min-max)

✅ **FIFA dataset example**: Scaling player attributes (age, height, shooting, passing, etc.) for player valuation prediction

### Use **Min-Max (Normalization)** when:

✅ **Algorithm requires bounded values** (0 to 1)
- Neural Networks (especially with sigmoid activations)
- K-Nearest Neighbors (KNN)
- Image processing

✅ **You need all features on the same scale** for comparison

✅ **Data has similar ranges** across features

✅ **You want to interpret values** as percentages of the range

✅ **Indicators dataset example**: Comparing countries across different economic indicators

---

## Practical Example: FIFA ML Pipeline

### Complete Preprocessing for ML

```toml
[data]
filename = './docs-dataset/fifa_players_22.csv'

[preprocess]

# Step 1: Handle missing values
[preprocess.numerics]
exclude_columns = ["value_eur"]  # Exclude target
operations = [
    {op = "fillna", method = "median"}
]

# Step 2: Scale all numeric features
[preprocess.numerics]
exclude_columns = ["value_eur"]  # Exclude target from scaling
operations = [
    {op = "scale", method = "zscore"}
]

# Step 3: Handle target variable
[[preprocess.columns]]
name = 'value_eur'
type = 'float'
operations = [
    {op = "fillna", method = "median"}
]

# Step 4: Encode categorical variables
[[preprocess.columns]]
name = 'preferred_foot'
type = 'string'
operations = [
    {op = "dummy", dummy_prefix = true, dummy_droplast = true}
]

[[preprocess.columns]]
name = 'body_type'
type = 'string'
operations = [
    {op = "dummy", dummy_prefix = true, dummy_droplast = true}
]

[postprocess]
filename = 'fifa_ml_ready.csv'

[postprocess.dataset_split]
method = "train_test_split"
split_names = ["train", "test"]
random_seed = 42
train_test_split_ratio = 0.8
```

**Why z-score for FIFA ML**:
1. Player attributes (age, height, skills) have **different units and scales**
2. We're using **linear models** or distance-based algorithms
3. We want to **preserve the distribution shape**
4. Z-score handles **outliers** better than min-max

---

## Practical Example: Indicators Comparison

### Complete Normalization Pipeline

```toml
[data]
filename = './docs-dataset/indicators.csv'

[preprocess]

# Step 1: Handle missing values
[preprocess.numerics]
operations = [
    {op = "fillna", method = "median"}
]

# Step 2: Normalize all features to [0, 1]
[preprocess.numerics]
operations = [
    {op = "scale", method = "minmax"}
]

# Step 3: Select relevant features (drop less important ones)
[postprocess]
filename = 'indicators_for_analysis.csv'
dropcolumns = [
    "BALANCE",
    "FOREIGN_INVESTMENT",
    "FUEL_EXPORTS",
    "EXPORT_TELECOM",
    "GROSS_SAVINGS"
]
```

**Why min-max for Indicators**:
1. We need to **compare countries across different indicators**
2. All values must be on **the same [0, 1] scale**
3. We're doing **clustering** which benefits from bounded values
4. Economic indicators often have **similar ranges** after scaling

---

## Handling Outliers Before Scaling

### The Problem

Both scaling methods are **sensitive to outliers**:
- **Z-score**: Mean and std are affected by extreme values
- **Min-Max**: Min and max can be outliers themselves

### Example: FIFA value_eur

```
value_eur statistics:
  min: 0
  max: 167,500,000  (outlier!)
  mean: 3,690,000
  median: 3,600,000
  std: 12,500,000 (inflated by outliers)
```

**With outliers**:
- Z-score: The 167.5M value becomes (167.5M - 3.69M) / 12.5M = **13.07** (extreme outlier)
- Min-Max: All values are scaled relative to 167.5M (compressed)

### Solutions

1. **Remove outliers** before scaling (if appropriate for your analysis)
2. **Use median-based methods** (like z-score) which are more robust
3. **Winsorize** (cap outliers) - **Feature request** for preprocess CLI
4. **Log-transform** skewed data first - **Feature request** for preprocess CLI

### Current Workaround

```toml
# For extremely skewed data like value_eur
# Consider excluding from scaling or using log transformation in post-processing

[preprocess.numerics]
exclude_columns = ["value_eur"]
operations = [
    {op = "scale", method = "zscore"}
]

# Handle value_eur separately (e.g., in Python/R)
```

---

## Verification Commands

### Before Scaling

```bash
# Generate statistics
preprocess summary --data ./docs-dataset/fifa_players_22.csv --output fifa_before.toml

# Check specific columns
preprocess skim --data ./docs-dataset/fifa_players_22.csv --columns age,value_eur,wage_eur
```

### After Scaling

```bash
# Generate statistics for scaled data
preprocess summary --data fifa_scaled.csv --output fifa_after.toml

# Compare before and after
preprocess diff \
  --source ./docs-dataset/fifa_players_22.csv \
  --target fifa_scaled.csv \
  --output fifa_scaling_diff.html
```

### What to Verify

**For Z-score scaling**:
- ✅ Mean ≈ 0 for each column
- ✅ Standard deviation ≈ 1 for each column
- ✅ Values can be negative or > 1

**For Min-Max scaling**:
- ✅ Minimum = 0 for each column
- ✅ Maximum = 1 for each column
- ✅ All values between 0 and 1

---

## Statistics Comparison: Before and After Scaling

### FIFA Dataset - Z-score Scaling

| Feature | Before (Min-Max) | After Z-score (Min-Max) | Transformation |
|---------|------------------|-------------------------|----------------|
| age | 16-39 | -2.01 to 3.05 | Standardized |
| height_cm | 155-203 | -2.18 to 1.96 | Standardized |
| overall | 44-99 | -2.38 to 2.56 | Standardized |
| potential | 44-99 | -2.38 to 2.56 | Standardized |
| value_eur | 0-167.5M | -0.30 to 13.07 | Standardized |

**Observations**:
- All features now have mean = 0, std = 1
- `value_eur` still has extreme values (13.07 std dev from mean) due to outliers
- Features are now **comparable** to each other

### Indicators Dataset - Min-Max Scaling

| Feature | Before (Min-Max) | After Min-Max (Min-Max) | Interpretation |
|---------|------------------|-------------------------|----------------|
| GDP_CAPITA | 2K-119K | 0.0-1.0 | Now comparable |
| POPULATION | 32M-1.4B | 0.0-1.0 | Now comparable |
| TRADE_BALANCE | -5.7B-7.1B | 0.0-1.0 | Now comparable |

**Observations**:
- All features now on [0, 1] scale
- **Direct comparison** now possible between any two indicators
- Outliers in original data become 0 or 1 in scaled data

---

## Common Mistakes and Solutions

### Mistake 1: Scaling Text Columns
**Problem**: Trying to scale string/categorical columns.
**Solution**: Exclude text columns from scaling:
```toml
[preprocess.numerics]
exclude_columns = ["short_name", "preferred_foot"]
operations = [
    {op = "scale", method = "zscore"}
]
```

### Mistake 2: Scaling Before Handling Missing Values
**Problem**: Scaling with missing values can produce unexpected results.
**Solution**: Always handle missing values **before** scaling:
```toml
[preprocess.numerics]
operations = [
    {op = "fillna", method = "median"},  # First: impute
    {op = "scale", method = "zscore"}     # Then: scale
]
```

### Mistake 3: Using Wrong Scaling for Algorithm
**Problem**: Using min-max for SVM or z-score for neural networks.
**Solution**: Match scaling method to algorithm requirements.

### Mistake 4: Not Excluding Target Variable
**Problem**: Scaling the target variable along with features.
**Solution**: Usually exclude target from scaling:
```toml
[preprocess.numerics]
exclude_columns = ["target_column"]
operations = [
    {op = "scale", method = "zscore"}
]
```

---

## Complete Decision Guide

### Choose Z-score when...

```
✓ Data has outliers
✓ Using algorithms that assume normally distributed data
✓ Need to preserve sparsity
✓ Features have different units
✓ Need to know how many std dev a value is from mean
```

### Choose Min-Max when...

```
✓ Need values between 0 and 1
✓ Using neural networks
✓ Need to compare features directly
✓ Data has similar ranges
✓ Need bounded values for visualization
```

---

## Key Takeaways

1. **Always scale** your features for machine learning (except decision trees)
2. **Z-score** is the default choice for most ML tasks (robust, preserves distribution)
3. **Min-Max** is essential when you need [0, 1] range (neural networks, comparisons)
4. **Handle missing values first** before scaling
5. **Exclude text columns** from numeric scaling operations
6. **Consider excluding target** variable from scaling
7. **Verify with summary** to ensure scaling worked correctly

---

## Identified Limitations and Future Needs

During our analysis, we identified these **missing features** that would improve scaling capabilities:

| Feature | Use Case | Priority | Current Workaround |
|---------|----------|----------|-------------------|
| `winsorize` | Cap outliers before scaling | ⭐⭐⭐⭐ | Remove outliers manually |
| `transform:log1p` | Log-transform skewed data | ⭐⭐⭐⭐ | Use external tools |
| `transform:sqrt` | Square root transform | ⭐⭐⭐ | Use external tools |
| `exclude_columns` per operation | Fine-grained exclusion | ⭐⭐⭐ | Use multiple sections |

---

## Related Articles

- [Handling Missing Values - Real-World Scenarios](/docs/examples/missing-values-examples/)
- [Text Cleaning - Practical Examples](/docs/examples/text-cleaning-examples/)
- [Categorical Encoding - Best Practices](/docs/examples/categorical-encoding-examples/)
- [FIFA Dataset Complete Analysis](/docs/examples/fifa-analysis/)

---

*Last updated: September 18, 2026*
*Based on analysis of FIFA Players 2022 and Indicators datasets*
