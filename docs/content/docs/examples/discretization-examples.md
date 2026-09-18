---
weight: 200
title: "Discretization - Transforming Continuous to Categorical"
icon: "segment"
date: "2026-09-18T12:00:00+02:00"
lastmod: "2026-09-18T12:00:00+02:00"
description: "Learn discretization (binning) with concrete examples from FIFA and HDV datasets"
draft: false
toc: true
---

## Discretization - Transforming Continuous to Categorical

Discretization (or **binning**) converts **continuous numeric variables** into **categorical bins**. This is useful for:
- Simplifying complex distributions
- Creating meaningful categories for analysis
- Reducing the impact of outliers
- Making continuous data more interpretable

This guide shows you **when and how** to use discretization with real-world examples.

---

## Why Discretize?

Continuous variables can sometimes be **too granular** for analysis. Discretization helps by:

✅ **Creating meaningful categories**: Age groups, income brackets, etc.
✅ **Simplifying analysis**: Fewer unique values = easier interpretation
✅ **Reducing noise**: Small variations in continuous data can be grouped
✅ **Improving readability**: "Young (18-25)" is more interpretable than "22.3"
✅ **Handling non-linear relationships**: Binned values can capture non-linear patterns

---

## Discretization Methods in Preprocess CLI

Preprocess CLI supports **custom binning** through the `discretize` operation:

```toml
{op = "discretize", method = "binning", bins = [
    {lower = 0, upper = 10, label = "Category 1"},
    {lower = 10, upper = 20, label = "Category 2"},
    {lower = 20, upper = 100, label = "Category 3"}
]}
```

---

## Example 1: FIFA Dataset - Player Age Binning

### Dataset: FIFA Players 2022

**Column**: `age`
- **Range**: 16-39 years
- **Distribution**: Relatively uniform with peak at 25
- **Use case**: Categorize players for analysis and scouting

### Age Categories for Soccer Players

In professional soccer, age matters significantly for player roles:
- **16-18**: Youth/Academy players
- **18-21**: Young talents (U21)
- **21-25**: Rising stars
- **25-30**: Prime age
- **30-35**: Experienced veterans
- **35+**: Veteran/End of career

### Prepfile for Age Discretization

```toml
[[preprocess.columns]]
name = 'age'
type = 'int'
operations = [
    {op = "discretize", method = "binning", bins = [
        {lower = 0, upper = 18, label = "Jeune (U21)"},
        {lower = 18, upper = 21, label = "Jeune Talent (21-25)"},
        {lower = 21, upper = 25, label = "Rising Star (21-25)"},
        {lower = 25, upper = 30, label = "Prime (25-30)"},
        {lower = 30, upper = 35, label = "Expérimenté (30-35)"},
        {lower = 35, upper = 100, label = "Vétéran (35+)"}
    ]}
]
```

### Before and After

**Before**:
```
age
---
16
17
20
25
30
35
39
```

**After**:
```
age
--------------
Jeune (U21)
Jeune (U21)
Rising Star (21-25)
Prime (25-30)
Expérimenté (30-35)
Vétéran (35+)
Vétéran (35+)
```

**✅ Result**: Continuous age values transformed into meaningful categories.

### Distribution After Discretization

| Category | Age Range | Typical Count | % of Players |
|----------|-----------|---------------|--------------|
| Jeune (U21) | 16-18 | ~500 | ~3% |
| Jeune Talent (21-25) | 18-21 | ~2,000 | ~12% |
| Rising Star (21-25) | 21-25 | ~4,000 | ~25% |
| Prime (25-30) | 25-30 | ~5,000 | ~31% |
| Expérimenté (30-35) | 30-35 | ~3,000 | ~19% |
| Vétéran (35+) | 35+ | ~500 | ~3% |

---

## Example 2: HDV Dataset - TV Watching Habits

### Dataset: HDV 2003 (Social Survey)

**Column**: `heures.tv`
- **Range**: 0-28 hours per week
- **Mean**: ~4 hours
- **Distribution**: Right-skewed (most people watch little TV)
- **Use case**: Categorize TV watching habits for social analysis

### TV Watching Categories

```toml
[[preprocess.columns]]
name = 'heures.tv'
type = 'float'
operations = [
    {op = "discretize", method = "binning", bins = [
        {lower = 0, upper = 2, label = "faible (0-2h)"},
        {lower = 2, upper = 7, label = "modéré (2-7h)"},
        {lower = 7, upper = 14, label = "élevé (7-14h)"},
        {lower = 14, upper = 100, label = "très élevé (14h+)"}
    ]}
]
```

### Before and After

**Before**:
```
heures.tv
---------
0
1
4
7
10
15
28
```

**After**:
```
heures.tv
--------------
faible (0-2h)
faible (0-2h)
modéré (2-7h)
élevé (7-14h)
élevé (7-14h)
très élevé (14h+)
très élevé (14h+)
```

**✅ Result**: TV watching habits categorized into 4 meaningful groups.

### Interpretation

- **faible (0-2h)**: Minimal TV watching
- **modéré (2-7h)**: Average TV consumption
- **élevé (7-14h)**: High TV consumption
- **très élevé (14h+)**: Very high consumption (potential concern)

---

## Example 3: Indicators Dataset - Economic Development Levels

### Dataset: Economic Indicators

**Column**: `GDP_CAPITA` (GDP per capita)
- **Range**: $2,082 - $119,367
- **Distribution**: Highly right-skewed
- **Use case**: Categorize countries by economic development level

### Economic Development Categories

```toml
[[preprocess.columns]]
name = 'GDP_CAPITA'
type = 'float'
operations = [
    {op = "discretize", method = "binning", bins = [
        {lower = 0, upper = 10000, label = "Low Income"},
        {lower = 10000, upper = 50000, label = "Lower Middle Income"},
        {lower = 50000, upper = 100000, label = "Upper Middle Income"},
        {lower = 100000, upper = 1000000, label = "High Income"}
    ]}
]
```

**⚠️ Note**: These thresholds are illustrative. Use standard economic classifications (World Bank, IMF) for real analysis.

---

## Choosing Bin Boundaries

The key to effective discretization is **choosing meaningful bin boundaries**. Here are several approaches:

### Approach 1: Domain Knowledge

Use **established categories** from the domain:
- Age: 18-25, 25-35, 35-65, 65+ (standard age groups)
- Income: Low, Middle, High (economic classifications)
- Blood pressure: Normal, Elevated, Hypertension (medical standards)

**FIFA age example**: Used soccer-specific categories based on player development stages.

### Approach 2: Equal Width Binning

Divide the range into **equal-sized intervals**:
```toml
# Age: 16-39, create 5 bins of width ~5
bins = [
    {lower = 16, upper = 21, label = "16-20"},
    {lower = 21, upper = 26, label = "21-25"},
    {lower = 26, upper = 31, label = "26-30"},
    {lower = 31, upper = 36, label = "31-35"},
    {lower = 36, upper = 40, label = "36-40"}
]
```

### Approach 3: Equal Frequency Binning (Quantiles)

Create bins with **approximately equal number of observations**:
- **Not directly supported** in preprocess CLI
- **Workaround**: Calculate quantiles in Python/R first, then use those as bin boundaries

### Approach 4: Data-Driven Boundaries

Use **natural breaks** in the data:
- Look at the distribution (histogram)
- Identify natural clusters
- Set boundaries at gaps in the distribution

---

## Number of Bins: Guidelines

| Number of Bins | When to Use | Risk |
|----------------|-------------|------|
| 2-3 | Very coarse categorization | Too simplistic |
| 4-6 | Most common | Good balance |
| 7-10 | Detailed categorization | Complexity, sparsity |
| 10+ | Very detailed | Too many categories, noise |

**Recommendation**: Start with 4-6 bins, adjust based on your analysis needs.

---

## Discretization vs Other Techniques

| Technique | Pros | Cons | When to Use |
|-----------|------|------|-------------|
| **Discretization** | Simple, interpretable | Loss of information | Categorical analysis, grouping |
| **Scaling** | Preserves all information | Still continuous | ML algorithms needing numeric input |
| **Both** | Best of both worlds | More complex | ML with categorical features |

---

## Practical Example: Complete FIFA Analysis Pipeline

### Scenario: Player Profile Analysis

We want to:
1. Categorize players by age
2. Categorize players by overall rating
3. Analyze distribution of player types across categories

### Complete Prepfile

```toml
[data]
filename = './docs-dataset/fifa_players_22.csv'

[preprocess]

# Step 1: Discretize age
[[preprocess.columns]]
name = 'age'
type = 'int'
operations = [
    {op = "discretize", method = "binning", bins = [
        {lower = 0, upper = 18, label = "Jeune (U21)"},
        {lower = 18, upper = 25, label = "Jeune Talent (18-25)"},
        {lower = 25, upper = 30, label = "Prime (25-30)"},
        {lower = 30, upper = 35, label = "Expérimenté (30-35)"},
        {lower = 35, upper = 100, label = "Vétéran (35+)"}
    ]}
]

# Step 2: Discretize overall rating
[[preprocess.columns]]
name = 'overall'
type = 'int'
operations = [
    {op = "discretize", method = "binning", bins = [
        {lower = 0, upper = 60, label = "Low (0-60)"},
        {lower = 60, upper = 75, label = "Medium (60-75)"},
        {lower = 75, upper = 90, label = "High (75-90)"},
        {lower = 90, upper = 100, label = "Elite (90-100)"}
    ]}
]

# Step 3: Encode categorical variables
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
filename = 'fifa_categorized.csv'
```

**Result**: Players categorized by age and rating, with categorical variables encoded.

---

## Practical Example: HDV Complete Preprocessing

### Scenario: Social Habits Analysis

We want to analyze TV watching habits by demographic categories.

### Complete Prepfile

```toml
[data]
filename = './docs-dataset/hdv2003.csv'
missing_identifier = 'NA'

[preprocess]

# Step 1: Clean text
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

# Step 3: Discretize heures.tv
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

# Step 4: Encode categorical variables
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

[postprocess]
filename = 'hdv_categorized.csv'
```

**Result**: Survey data with TV habits categorized and demographics encoded.

---

## Verification Commands

### Check Discretization Results

```bash
# Before discretization
preprocess summary --data ./docs-dataset/fifa_players_22.csv --output fifa_before.toml

# After discretization
preprocess summary --data fifa_categorized.csv --output fifa_after.toml
```

**What to check**:
- ✅ Numeric columns are now text/categorical
- ✅ Unique value count matches number of bins
- ✅ No values outside bin ranges
- ✅ All values have been assigned to a bin

### Visual Diff

```bash
preprocess diff \
  --source ./docs-dataset/fifa_players_22.csv \
  --target fifa_categorized.csv \
  --output fifa_discretization_diff.html
```

---

## Common Mistakes and Solutions

### Mistake 1: Overlapping Bin Boundaries
**Problem**: Bins overlap, causing ambiguous categorization.
**Solution**: Ensure `upper` of one bin = `lower` of next bin:
```toml
# Correct: no overlap
{lower = 0, upper = 18, label = "Group 1"},
{lower = 18, upper = 25, label = "Group 2"},  # 18 is in Group 2, not Group 1

# Also correct: explicit non-overlap
{lower = 0, upper = 18, label = "Group 1"},
{lower = 18.01, upper = 25, label = "Group 2"},  # But floats might have precision issues
```

### Mistake 2: Gaps Between Bins
**Problem**: Values fall between bins and get no label.
**Solution**: Cover the entire range:
```toml
# Start from min
{lower = 0, upper = 18, label = "Group 1"},
{lower = 18, upper = 25, label = "Group 2"},
# ...
{lower = 35, upper = 100, label = "Group N"}  # End with high upper bound
```

### Mistake 3: Not Handling Missing Values First
**Problem**: Missing values cause errors in discretization.
**Solution**: Fill missing values before discretization:
```toml
[[preprocess.columns]]
name = 'age'
type = 'int'
operations = [
    {op = "fillna", method = "median"},  # First: handle missing
    {op = "discretize", method = "binning", bins = [...]}
]
```

### Mistake 4: Too Many or Too Few Bins
**Problem**: Bins are not meaningful for analysis.
**Solution**: Use domain knowledge or analyze distribution first:
```bash
# Check distribution first
preprocess summary --data ./docs-dataset/fifa_players_22.csv
preprocess skim --data ./docs-dataset/fifa_players_22.csv --columns age
```

---

## Best Practices

### Do:

✅ **Use domain knowledge** to define meaningful categories
✅ **Cover the entire range** with your bins
✅ **Ensure no overlaps** between bins
✅ **Handle missing values first**
✅ **Start with 4-6 bins** and adjust as needed
✅ **Use descriptive labels** for categories
✅ **Verify with summary** after discretization

### Don't:

❌ **Use arbitrary bin boundaries** without justification
❌ **Create overlapping bins**
❌ **Leave gaps between bins**
❌ **Use too many bins** (causes sparsity)
❌ **Use too few bins** (loses information)
❌ **Discretize without a purpose**

---

## Discretization Decision Guide

```
Should you discretize?
  │
  ├── No → Keep continuous
  │    │
  │    └── When:
  │        - Using algorithms that need numeric input (most ML)
  │        - Need precise values
  │        - Data has no natural categories
  │
  └── Yes →
       │
       ├── What's the purpose?
       │   │
       │   ├── Analysis/Visualization → Use meaningful categories
       │   ├── Simplification → Use equal-width or quantile bins
       │   └── Grouping → Use domain-specific categories
       │
       ├── How many observations?
       │   │
       │   ├── < 100 → Use 2-4 bins
       │   ├── 100-1000 → Use 4-6 bins
       │   └── > 1000 → Use 6-10 bins
       │
       └── → Define bins based on domain knowledge or distribution
```

---

## Key Takeaways

1. **Discretization converts continuous to categorical** for easier analysis
2. **Use domain knowledge** to define meaningful categories
3. **Always cover the entire range** with non-overlapping bins
4. **Handle missing values first** before discretization
5. **Start with 4-6 bins** for most analyses
6. **Verify results** with summary and diff commands
7. **Use descriptive labels** for better interpretability

---

## Identified Limitations and Future Needs

| Feature | Use Case | Priority | Current Workaround |
|---------|----------|----------|-------------------|
| Quantile binning | Equal frequency bins | ⭐⭐⭐⭐ | Calculate in Python/R first |
| Optimal binning (WOE) | Statistically optimal bins | ⭐⭐⭐ | Use external tools |
| K-means binning | Data-driven cluster bins | ⭐⭐⭐ | Pre-cluster, then bin |
| Automatic bin selection | Choose number of bins automatically | ⭐⭐ | Manual selection |

---

## Related Articles

- [Handling Missing Values - Real-World Scenarios](/docs/examples/missing-values-examples/)
- [Text Cleaning - Practical Examples](/docs/examples/text-cleaning-examples/)
- [Feature Scaling for ML - Complete Guide](/docs/examples/feature-scaling-examples/)
- [Categorical Encoding - Best Practices](/docs/examples/categorical-encoding-examples/)
- [FIFA Dataset Complete Analysis](/docs/examples/fifa-analysis/)

---

*Last updated: September 18, 2026*
*Based on analysis of FIFA Players 2022 and HDV 2003 datasets*
