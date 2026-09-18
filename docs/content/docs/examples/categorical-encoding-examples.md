---
weight: 200
title: "Categorical Encoding - Best Practices with Real Datasets"
icon: "label"
date: "2026-09-18T12:00:00+02:00"
lastmod: "2026-09-18T12:00:00+02:00"
description: "Learn categorical encoding with concrete examples from FIFA and HDV datasets - one-hot encoding done right"
draft: false
toc: true
---

## Categorical Encoding - Best Practices

Categorical encoding transforms **text categories** into **numeric representations** that machine learning algorithms can process. This guide shows you **when and how** to use categorical encoding with real-world examples.

---

## Why Categorical Encoding is Necessary

Machine learning algorithms work with **numeric data**. When you have categorical data like:
- Player foot preference: "Left", "Right"
- Body type: "Lean", "Normal", "Stocky"
- Gender: "Homme", "Femme"
- Education level: "Primary", "Secondary", "Higher"

You need to **encode** these categories as numbers.

**But there's a catch**: Simply mapping "Left" → 0, "Right" → 1 implies an **ordinal relationship** (0 < 1), which doesn't exist. The solution: **One-Hot Encoding**.

---

## One-Hot Encoding Explained

### The Concept

One-hot encoding creates **binary columns** for each category:
- Each category becomes a **separate column**
- Each column has **0 or 1** (False or True)
- Only **one column** has 1 for each observation (hence "one-hot")

### Example: Binary Categorical Variable

**Column**: `preferred_foot` (FIFA dataset)
- Categories: "Left", "Right"

**Before encoding**:
```
preferred_foot
--------------
Left
Right
Left
Right
```

**One-Hot Encoded**:
```
preferred_foot_Left | preferred_foot_Right
--------------------|----------------------
1                   | 0
0                   | 1
1                   | 0
0                   | 1
```

### Example: Multi-category Variable

**Column**: `body_type` (FIFA dataset)
- Categories: "Lean", "Normal", "Stocky", "Unique"

**Before encoding**:
```
body_type
----------
Lean
Normal
Stocky
Lean
Unique
```

**One-Hot Encoded**:
```
body_type_Lean | body_type_Normal | body_type_Stocky | body_type_Unique
---------------|------------------|----------------|----------------
1              | 0                | 0              | 0
0              | 1                | 0              | 0
0              | 0                | 1              | 0
1              | 0                | 0              | 0
0              | 0                | 0              | 1
```

---

## One-Hot Encoding in Preprocess CLI

Preprocess CLI uses the `dummy` operation for one-hot encoding.

### Basic Syntax

```toml
[[preprocess.columns]]
name = 'column_name'
type = 'string'
operations = [
    {op = "dummy"}
]
```

---

## Example 1: FIFA Dataset - Player Foot Preference

### Dataset: FIFA Players 2022

**Column**: `preferred_foot`
- **Categories**: "Left", "Right"
- **Rows**: 16,020
- **Missing**: 0

### Simple One-Hot Encoding

```toml
[[preprocess.columns]]
name = 'preferred_foot'
type = 'string'
operations = [
    {op = "dummy"}
]
```

**Result**:
- Creates 2 columns: `preferred_foot_Left`, `preferred_foot_Right`
- Each row has one 1 and one 0

**Problem**: This creates **multicollinearity** - the two columns are perfectly negatively correlated. If you know one, you know the other.

### Solution: Drop Last Category

```toml
[[preprocess.columns]]
name = 'preferred_foot'
type = 'string'
operations = [
    {op = "dummy", dummy_droplast = true}
]
```

**Result**:
- Creates **1 column**: `preferred_foot_Left`
- "Left" → 1, "Right" → 0 (reference category)

**✅ Benefits**:
- Eliminates multicollinearity
- Reduces dimensionality
- Preserves all information (Right = not Left)

---

## Example 2: FIFA Dataset - Body Type (Multiple Categories)

### Dataset: FIFA Players 2022

**Column**: `body_type`
- **Categories**: Lean, Normal, Stocky, Unique, and possibly others
- **Rows**: 16,020

### With Prefix and Drop Last

```toml
[[preprocess.columns]]
name = 'body_type'
type = 'string'
operations = [
    {op = "dummy", dummy_prefix = true, dummy_droplast = true}
]
```

**Result** (assuming 5 categories):
- Creates **4 columns**: `body_type_Lean`, `body_type_Normal`, `body_type_Stocky`, `body_type_Unique`
- Each row has one 1 and three 0s
- The 5th category (e.g., "Default") is the reference (all 0s)

**✅ Benefits of `dummy_prefix`**:
- Makes column names more descriptive
- Easier to identify which columns came from which categorical variable

---

## Example 3: HDV Dataset - Gender Encoding

### Dataset: HDV 2003 (Social Survey)

**Column**: `sexe` (gender)
- **Categories**: "Homme", "Femme", "NA"
- **Rows**: 1,000+

### Complete Example with Missing Value Handling

```toml
[data]
filename = './docs-dataset/hdv2003.csv'
missing_identifier = 'NA'

[preprocess]

# Step 1: Fill missing values in categorical column
[preprocess.texts]
operations = [
    {op = "fillna", value = "non spécifié"}
]

# Step 2: Clean text (remove whitespace, lowercase)
[preprocess.texts]
operations = [
    {op = "clean", method = "trimws"},
    {op = "clean", method = "lower"}
]

# Step 3: Encode categorical variables
[[preprocess.columns]]
name = 'sexe'
type = 'string'
operations = [
    {op = "dummy", dummy_prefix = false, dummy_droplast = true}
]

[postprocess]
filename = 'hdv_encoded.csv'
```

**What this does**:
1. Fills missing values with "non spécifié"
2. Cleans text (removes spaces, lowercases)
3. Encodes `sexe` with one-hot encoding, dropping the last category

**Result**:
- Categories: "homme", "femme", "non spécifié"
- Creates 2 columns (dropping last): `sexe_homme`, `sexe_femme`
- "non spécifié" becomes the reference category

---

## Dummy Encoding Options in Preprocess CLI

| Option | Description | Use Case | Example |
|--------|-------------|----------|---------|
| `dummy` | Basic one-hot encoding | Simple encoding | `{op = "dummy"}` |
| `dummy_prefix` | Add column name prefix | Multiple categorical vars | `{op = "dummy", dummy_prefix = true}` |
| `dummy_droplast` | Drop last category | Avoid multicollinearity | `{op = "dummy", dummy_droplast = true}` |

### Combining Options

```toml
# Best practice for most cases
{op = "dummy", dummy_prefix = true, dummy_droplast = true}

# When you need all categories (rare)
{op = "dummy", dummy_prefix = true, dummy_droplast = false}

# When you want minimal column names
{op = "dummy", dummy_prefix = false, dummy_droplast = true}
```

---

## When to Use Dummy Encoding vs Other Methods

### Use One-Hot (Dummy) Encoding when:

✅ **Categories have no ordinal relationship** (most common case)
- Foot preference: Left, Right
- Body type: Lean, Normal, Stocky
- Gender: Male, Female
- Country: France, Germany, Italy

✅ **Using algorithms that don't handle categorical data natively**
- Linear Regression
- Logistic Regression
- Neural Networks
- SVM
- k-Means Clustering

### Use Other Encodings when:

⚠️ **Categories have ordinal relationship** → Use ordinal encoding
- Education: Primary (1), Secondary (2), Higher (3)
- Satisfaction: Low (1), Medium (2), High (3)

⚠️ **High cardinality** (many categories) → Use target encoding or embeddings
- ZIP codes
- User IDs
- Product IDs

⚠️ **Tree-based algorithms** → Often don't need encoding
- Random Forest
- XGBoost
- Decision Trees

---

## Practical Example: Complete FIFA ML Pipeline

### Scenario: Predict Player Value

We want to predict `value_eur` based on:
- Numeric features: age, height, weight, skills (shooting, passing, etc.)
- Categorical features: preferred_foot, body_type

### Complete Prepfile

```toml
[data]
filename = './docs-dataset/fifa_players_22.csv'

[preprocess]

# Step 1: Handle missing values in numeric columns
[preprocess.numerics]
exclude_columns = ["value_eur"]  # Exclude target
operations = [
    {op = "fillna", method = "median"}
]

# Step 2: Scale numeric features
[preprocess.numerics]
exclude_columns = ["value_eur", "preferred_foot", "body_type"]
operations = [
    {op = "scale", method = "zscore"}
]

# Step 3: Encode categorical variables

# Foot preference (binary: Left/Right)
[[preprocess.columns]]
name = 'preferred_foot'
type = 'string'
operations = [
    {op = "dummy", dummy_prefix = true, dummy_droplast = true}
]

# Body type (multiple categories)
[[preprocess.columns]]
name = 'body_type'
type = 'string'
operations = [
    {op = "dummy", dummy_prefix = true, dummy_droplast = true}
]

# Step 4: Handle target variable
[[preprocess.columns]]
name = 'value_eur'
type = 'float'
operations = [
    {op = "fillna", method = "median"}
]

[postprocess]
filename = 'fifa_ml_ready.csv'

[postprocess.dataset_split]
method = "train_test_split"
split_names = ["train", "test"]
random_seed = 42
train_test_split_ratio = 0.8
```

**Result**:
- Numeric features: Scaled (mean=0, std=1)
- `preferred_foot`: 1 column (`preferred_foot_Left`), 0=Right, 1=Left
- `body_type`: N-1 columns (where N = number of body types)
- `value_eur`: Target variable, filled but not scaled
- Dataset split: 80% train, 20% test

---

## Practical Example: HDV Survey Analysis

### Scenario: Analyze Social Survey Data

We want to analyze relationships between:
- Demographics: sexe, nivetud (education level)
- Lifestyle: heures.tv (TV hours), occup (occupation)

### Complete Prepfile

```toml
[data]
filename = './docs-dataset/hdv2003.csv'
missing_identifier = 'NA'

[preprocess]

# Step 1: Clean all text columns
[preprocess.texts]
operations = [
    {op = "clean", method = "trimws"},
    {op = "clean", method = "lower"}
]

# Step 2: Handle missing values
[preprocess.numerics]
operations = [
    {op = "fillna", method = "median"}
]

[preprocess.texts]
operations = [
    {op = "fillna", value = "non spécifié"}
]

# Step 3: Encode categorical variables

[[preprocess.columns]]
name = 'sexe'
type = 'string'
operations = [
    {op = "dummy", dummy_prefix = false, dummy_droplast = true}
]

[[preprocess.columns]]
name = 'nivetud'
type = 'string'
operations = [
    {op = "dummy", dummy_prefix = false, dummy_droplast = true}
]

[[preprocess.columns]]
name = 'occup'
type = 'string'
operations = [
    {op = "dummy", dummy_prefix = false, dummy_droplast = true}
]

# Step 4: Discretize numeric variable
[[preprocess.columns]]
name = 'heures.tv'
type = 'float'
operations = [
    {op = "discretize", method = "binning", bins = [
        {lower = 0, upper = 2, label = "faible"},
        {lower = 2, upper = 7, label = "modéré"},
        {lower = 7, upper = 14, label = "élevé"},
        {lower = 14, upper = 100, label = "très élevé"}
    ]}
]

[postprocess]
filename = 'hdv_analysis_ready.csv'
```

**Result**:
- Text columns: Cleaned and missing values filled
- Numeric columns: Missing values imputed
- Categorical columns: One-hot encoded
- `heures.tv`: Discretized into 4 categories

---

## Before/After Comparison

### FIFA Dataset

| Aspect | Before Encoding | After Encoding | Change |
|--------|-----------------|----------------|--------|
| Total columns | 49 | ~60 | +11 |
| preferred_foot | 1 (text) | 1 (dummy) | Encoded |
| body_type | 1 (text) | 4-6 (dummy) | Encoded |
| Column types | 46 numeric, 3 text | ~59 numeric | All numeric |
| ML-ready | ❌ No | ✅ Yes | Ready |

### HDV Dataset

| Aspect | Before Encoding | After Encoding | Change |
|--------|-----------------|----------------|--------|
| Total columns | 18 | ~20-25 | +2-7 |
| sexe | 1 (text) | 1 (dummy) | Encoded |
| nivetud | 1 (text) | 3-4 (dummy) | Encoded |
| occup | 1 (text) | 5-10 (dummy) | Encoded |
| heures.tv | 1 (numeric) | 1 (categorical) | Discretized |

---

## Handling High Cardinality

### The Problem

Some categorical variables have **many categories**:
- Country names: 195+ possible values
- Player names: 16,020 unique values
- ZIP codes: Thousands of possible values

**Issues with one-hot encoding high cardinality**:
- ❌ Creates **thousands of columns** (curse of dimensionality)
- ❌ **Sparse data**: Most columns will be 0
- ❌ **Computationally expensive**
- ❌ **Overfitting risk**

### Solutions

1. **Group rare categories**: Combine infrequent categories into "Other"
2. **Use frequency encoding**: Replace category with its frequency
3. **Use target encoding**: Replace category with mean of target for that category
4. **Remove the column**: If not useful for analysis
5. **Use in tree-based models**: Random Forest, XGBoost handle categorical natively

### Current Workaround in Preprocess CLI

```toml
# For columns with too many categories, consider:
# 1. Not encoding (if using tree-based algorithms)
# 2. Removing the column
# 3. Grouping categories before encoding

[postprocess]
dropcolumns = ["player_name", "club", "nationality"]  # High cardinality columns
```

---

## Verification Commands

### Check Encoding Results

```bash
# Before encoding
preprocess skim --data ./docs-dataset/fifa_players_22.csv --limit 5

# After encoding
preprocess skim --data fifa_encoded.csv --limit 5
```

### Compare Column Counts

```bash
# Before
preprocess summary --data ./docs-dataset/fifa_players_22.csv --output fifa_before.toml

# After
preprocess summary --data fifa_encoded.csv --output fifa_after.toml
```

**What to verify**:
- ✅ Text columns are replaced with numeric dummy columns
- ✅ No text columns remain (or only those intentionally kept)
- ✅ Column count increased appropriately
- ✅ No missing values in encoded columns

### Visual Diff

```bash
preprocess diff \
  --source ./docs-dataset/fifa_players_22.csv \
  --target fifa_encoded.csv \
  --output fifa_encoding_diff.html
```

---

## Common Mistakes and Solutions

### Mistake 1: Not Dropping Last Category
**Problem**: Creating multicollinearity in your data.
**Solution**: Always use `dummy_droplast = true` unless you have a specific reason not to.

### Mistake 2: Encoding Before Cleaning
**Problem**: Encoding "Left", " left ", "LEFT" as different categories.
**Solution**: Clean text first, then encode:
```toml
# First: clean
[preprocess.texts]
operations = [
    {op = "clean", method = "trimws"},
    {op = "clean", method = "lower"}
]

# Then: encode
[[preprocess.columns]]
name = 'column_name'
type = 'string'
operations = [
    {op = "dummy", dummy_droplast = true}
]
```

### Mistake 3: Encoding All Text Columns
**Problem**: Encoding text columns that shouldn't be (names, IDs, descriptions).
**Solution**: Be selective about which columns to encode:
```toml
# Only encode categorical variables, not identifiers
[[preprocess.columns]]
name = 'preferred_foot'  # Categorical
operations = [{op = "dummy"}]

# Don't encode
# - short_name (player name - identifier)
# - player_id (unique ID)
# - description (free text)
```

### Mistake 4: Not Handling Missing Values First
**Problem**: Encoding treats missing as a category.
**Solution**: Fill missing values before encoding:
```toml
# First: fill NA
[preprocess.texts]
operations = [
    {op = "fillna", value = "non spécifié"}
]

# Then: encode
[[preprocess.columns]]
name = 'column_name'
operations = [{op = "dummy"}]
```

---

## Best Practices Summary

### Do:

✅ **Always use `dummy_droplast = true`** to avoid multicollinearity
✅ **Use `dummy_prefix = true`** for better column name clarity
✅ **Clean text before encoding** (trimws, lower/upper)
✅ **Handle missing values before encoding**
✅ **Be selective** about which columns to encode
✅ **Drop high-cardinality columns** or use alternative encodings

### Don't:

❌ **Encode identifier columns** (names, IDs)
❌ **Encode free-text columns** (descriptions, comments)
❌ **Encode before cleaning**
❌ **Encode without handling missing values**
❌ **Use one-hot for ordinal data**

---

## Complete Categorical Encoding Decision Tree

```
Is the column categorical?
  │
  ├── No → Don't encode
  │
  └── Yes →
      │
      ├── Does it have ordinal relationship?
      │   │
      │   ├── Yes → Use ordinal encoding (not in preprocess CLI yet)
      │   │
      │   └── No →
      │       │
      │       ├── Is cardinality > 20?
      │       │   │
      │       │   ├── Yes → Use target/frequency encoding or drop
      │       │   │
      │       │   └── No → Use one-hot encoding
      │       │
      │       ├── Are there missing values?
      │       │   │
      │       │   ├── Yes → Fill first, then encode
      │       │   │
      │       │   └── No → Encode directly
      │       │
      │       └── → Use {op = "dummy", dummy_prefix = true, dummy_droplast = true}
```

---

## Key Takeaways

1. **One-hot encoding** is the standard for categorical data without ordinal relationships
2. **Always drop last category** (`dummy_droplast = true`) to avoid multicollinearity
3. **Use prefix** (`dummy_prefix = true`) for better column organization
4. **Clean before encoding** - standardize text first
5. **Handle missing values first** - fill NA before encoding
6. **Be selective** - don't encode identifiers or free text
7. **Consider cardinality** - one-hot doesn't work for high-cardinality columns

---

## Identified Limitations and Future Needs

| Feature | Use Case | Priority | Current Workaround |
|---------|----------|----------|-------------------|
| Ordinal encoding | Encoded ordered categories | ⭐⭐⭐ | Manual mapping |
| Frequency encoding | Replace with frequency count | ⭐⭐⭐ | Manual calculation |
| Target encoding | Replace with mean of target | ⭐⭐⭐⭐ | Manual calculation |
| Group rare categories | Combine infrequent categories | ⭐⭐⭐ | Pre-process in Python/R |

---

## Related Articles

- [Handling Missing Values - Real-World Scenarios](/docs/examples/missing-values-examples/)
- [Text Cleaning - Practical Examples](/docs/examples/text-cleaning-examples/)
- [Feature Scaling for ML - Complete Guide](/docs/examples/feature-scaling-examples/)
- [FIFA Dataset Complete Analysis](/docs/examples/fifa-analysis/)

---

*Last updated: September 18, 2026*
*Based on analysis of FIFA Players 2022 and HDV 2003 datasets*
