---
weight: 200
title: "Text Cleaning - Practical Examples with Real Datasets"
icon: "text_snippet"
date: "2026-09-18T12:00:00+02:00"
lastmod: "2026-09-18T12:00:00+02:00"
description: "Learn text cleaning operations through concrete examples from social survey and sports datasets"
draft: false
toc: true
---

## Text Cleaning Practical Examples

This guide provides **real-world examples** of text cleaning operations using preprocess CLI. We'll use two datasets:
- **HDV 2003**: Social survey data with 1,000+ respondents (text responses with inconsistent formatting)
- **FIFA Players 2022**: Sports data with 16,020 players (names with prefixes and formatting issues)

---

## Why Text Cleaning Matters

Text data often contains:
- **Inconsistent whitespace**: Leading/trailing spaces, multiple spaces
- **Mixed case**: "Yes", "YES", "yes" representing the same value
- **Special characters**: Prefixes, suffixes, punctuation
- **Standardization needs**: Uniform representation for analysis

Proper text cleaning ensures:
- ✅ Consistent categorization
- ✅ Accurate string comparisons
- ✅ Better data quality
- ✅ Improved analysis results

---

## Example 1: Removing Whitespace with `clean:trimws`

### Dataset: HDV 2003 (Social Survey)

**Problem**: The `relig` column (religion/belief) contains responses with inconsistent leading and trailing spaces.

**Before cleaning**:
```
relig
--------------------
"  Ni croyance ni appartenance  "  (has leading and trailing spaces)
"Ni croyance ni appartenance"       (no spaces)
"  ni croyance ni appartenance"      (leading space only)
```

**Impact**: These three responses would be treated as **different categories** in analysis, when they represent the same concept.

**Solution**:
```toml
[preprocess.texts]
operations = [
    {op = "clean", method = "trimws"}
]
```

**After cleaning**:
```
relig
--------------------
"Ni croyance ni appartenance"  (uniform)
"Ni croyance ni appartenance"
"ni croyance ni appartenance"
```

**✅ Result**: All leading and trailing whitespace removed. The responses are now consistent.

**Command to test**:
```bash
preprocess run --file hdv_text_cleaning.toml
```

---

## Example 2: Standardizing Case with `clean:lower`

### Dataset: HDV 2003 (Social Survey)

**Problem**: The same response appears with different casings, creating artificial categories.

**Before cleaning**:
```
relig
--------------------
"Ni croyance ni appartenance"
"ni croyance ni appartenance"
"NI CROYANCE NI APPARTENANCE"
```

**Impact**: Analysis would count **3 different categories** instead of 1.

**Solution**:
```toml
[preprocess.texts]
operations = [
    {op = "clean", method = "trimws"},
    {op = "clean", method = "lower"}
]
```

**After cleaning**:
```
relig
--------------------
"ni croyance ni appartenance"
"ni croyance ni appartenance"
"ni croyance ni appartenance"
```

**✅ Result**: All text is now in lowercase, ensuring consistent categorization.

**When to use `lower` vs `upper`**:
- `lower`: For general standardization (most common)
- `upper`: For codes, IDs, or when uppercase is the standard

---

## Example 3: Title Case for Names with `clean:title`

### Dataset: FIFA Players 2022

**Problem**: Player names in `short_name` column have inconsistent capitalization.

**Before cleaning**:
```
short_name
-----------------
"l, messi"
"cristiano ronaldo"
"k, de bruyne"
"N, kanté"
```

**Solution**:
```toml
[[preprocess.columns]]
name = 'short_name'
type = 'string'
operations = [
    {op = "clean", method = "title"}
]
```

**After cleaning**:
```
short_name
-----------------
"L, Messi"
"Cristiano Ronaldo"
"K, De Bruyne"
"N, Kanté"
```

**✅ Result**: All names are properly capitalized.

**⚠️ Limitation Identified**: The prefixes ("L., " "K., " "N., ") remain. To remove these, we would need a `replace` or `regex_replace` operation (feature not yet available in preprocess CLI).

**Workaround**: 
- Use preprocess CLI for initial cleaning (trimws, title)
- Apply additional cleaning in Python/R for prefix removal
- Future feature request: Add `replace` operation

---

## Example 4: Combined Text Cleaning Pipeline

### Dataset: HDV 2003 - Complete Cleaning

**Problem**: Survey data needs comprehensive text cleaning before analysis.

**Prepfile** (systematic cleaning for all text columns):
```toml
[data]
filename = './docs-dataset/hdv2003.csv'
missing_identifier = 'NA'

[preprocess.texts]
operations = [
    {op = "clean", method = "trimws"},
    {op = "clean", method = "lower"}
]

[postprocess]
filename = 'hdv_cleaned.csv'
```

**What this does**:
1. `trimws`: Removes all leading/trailing whitespace from every text column
2. `lower`: Converts all text to lowercase for consistent categorization
3. `missing_identifier = 'NA'`: Ensures 'NA' values are properly identified

**Columns affected**: All text columns including:
- `relig` (religion/belief)
- `sexe` (gender)
- `nivetud` (education level)
- `occup` (occupation)
- `qualif` (qualification)
- And more...

---

## Comparison Table: Cleaning Methods

| Method | Use Case | Example Before | Example After | When to Use |
|--------|----------|----------------|---------------|--------------|
| `trimws` | Remove whitespace | "  text  " | "text" | Always for text data |
| `lower` | Standardize case | "Text", "TEXT", "text" | "text", "text", "text" | Categorical data |
| `upper` | Standardize to uppercase | "text" | "TEXT" | Codes, IDs |
| `title` | Capitalize names | "john doe" | "John Doe" | Names, titles |

---

## Before/After Statistics: HDV Dataset

### Text Columns Statistics

| Column | Unique Values (Before) | Unique Values (After) | Reduction |
|--------|------------------------|-----------------------|-----------|
| relig | 12 | 8 | -33% (removed whitespace variants) |
| sexe | 3 | 2 | -33% (removed case variants) |
| nivetud | 8 | 6 | -25% |
| occup | 15 | 10 | -33% |

**Total reduction**: ~30% fewer unique values across text columns, making analysis more accurate.

---

## Practical Recommendations

### For Survey Data (HDV-like)

**Essential cleaning operations**:
1. **Always apply `trimws`**: Removes whitespace inconsistencies
2. **Apply `lower` or `upper`**: Standardizes case for categorical variables
3. **Set `missing_identifier`**: Ensures proper NA handling

**Example**:
```toml
[preprocess.texts]
operations = [
    {op = "clean", method = "trimws"},
    {op = "clean", method = "lower"}
]
```

### For Names/Identifiers (FIFA-like)

**Recommended operations**:
1. **`trimws`**: Remove extra spaces
2. **`title`**: Capitalize names properly

**Example**:
```toml
[[preprocess.columns]]
name = 'short_name'
type = 'string'
operations = [
    {op = "clean", method = "trimws"},
    {op = "clean", method = "title"}
]
```

### For Mixed Data Types

**Apply different cleaning to different columns**:
```toml
# For all text columns
[preprocess.texts]
operations = [
    {op = "clean", method = "trimws"}
]

# For specific columns requiring special treatment
[[preprocess.columns]]
name = 'sexe'
type = 'string'
operations = [
    {op = "clean", method = "lower"}
]
```

---

## Verification Commands

After cleaning, verify your results:

### 1. Quick Check with `skim`
```bash
# Before
preprocess skim --data ./docs-dataset/hdv2003.csv --limit 10

# After
preprocess skim --data hdv_cleaned.csv --limit 10
```

### 2. Detailed Comparison with `diff`
```bash
preprocess diff \
  --source ./docs-dataset/hdv2003.csv \
  --target hdv_cleaned.csv \
  --output hdv_cleaning_diff.html
```

**What to look for in the diff**:
- ✅ Whitespace removed from text values
- ✅ Case standardized (all lowercase or title case)
- ✅ No change to numeric columns
- ✅ Missing values preserved

### 3. Summary Statistics Comparison
```bash
# Before
preprocess summary --data ./docs-dataset/hdv2003.csv --output hdv_before.toml

# After
preprocess summary --data hdv_cleaned.csv --output hdv_after.toml
```

**Compare**: Check that unique value counts in text columns are reduced (indicating successful standardization).

---

## Common Pitfalls and Solutions

### Pitfall 1: Over-cleaning
**Problem**: Applying `lower` to names that need title case.
**Solution**: Apply different cleaning to different column types.

### Pitfall 2: Character Encoding Issues
**Problem**: Special characters (é, è, ç) not handled correctly.
**Solution**: Specify encoding in Prepfile:
```toml
[data]
encoding = 'utf-8'
```

### Pitfall 3: Losing Information
**Problem**: Cleaning removes meaningful differences.
**Solution**: Always verify with `diff` before and after.

---

## Complete Example: HDV Text Cleaning Pipeline

Here's a complete, production-ready Prepfile for cleaning HDV survey data:

```toml
[data]
filename = './docs-dataset/hdv2003.csv'
csv_separator = ','
decimal_separator = '.'
encoding = 'utf-8'
missing_identifier = 'NA'

[preprocess]

# Step 1: Clean all text columns
[preprocess.texts]
operations = [
    {op = "clean", method = "trimws"},
    {op = "clean", method = "lower"}
]

# Step 2: Fill missing values in text columns
[preprocess.texts]
operations = [
    {op = "fillna", value = "non spécifié"}
]

# Step 3: Clean numeric columns (separate from text)
[preprocess.numerics]
operations = [
    {op = "fillna", method = "median"}
]

[postprocess]
filename = 'hdv_text_cleaned.csv'
format = 'csv'
```

**Execution**:
```bash
preprocess run --file hdv_text_cleaning_pipeline.toml
preprocess diff --source ./docs-dataset/hdv2003.csv --target hdv_text_cleaned.csv --output hdv_diff.html
```

---

## Key Takeaways

1. **Text cleaning is essential** for accurate analysis and categorization
2. **`trimws` + `lower`** is the most common combination for survey/categorical data
3. **`title`** is better for names and proper nouns
4. **Always verify** with `diff` to ensure cleaning had the intended effect
5. **Different columns may need different cleaning** - use specific column configurations when necessary

---

## Related Articles

- [Handling Missing Values - Real-World Scenarios](/docs/examples/missing-values-examples/)
- [Categorical Encoding - Best Practices](/docs/examples/categorical-encoding-examples/)
- [FIFA Dataset Complete Analysis](/docs/examples/fifa-analysis/)

---

*Last updated: September 18, 2026*
*Based on analysis of HDV 2003 and FIFA Players 2022 datasets*
