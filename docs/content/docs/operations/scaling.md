---
weight: 300
date: "2025-05-30T14:30:40+02:00"
draft: false
author: "Axel-Cleris Gailloty"
title: "Scaling dataset"
icon: "developer_guide"
toc: true
description: "A reference of all implemented operations"
publishdate: "2025-05-30T14:30:40+02:00"
---

## Scaling dataset columns

**Scaling** refers to transforming features so they have similar ranges or distributions. Common **scaling methods** include:

1. **Min-Max Scaling (Normalization)**
   Rescales features to a fixed range, usually \[0, 1].

2. **Standardization (Z-score Normalization)**
   Transforms features to have a mean of 0 and standard deviation of 1.

---

## Why Scaling is Useful

Scaling is crucial for many machine learning algorithms, especially those sensitive to the magnitude of features:

* **Principal Component Analysis (PCA)**:
  PCA projects data onto directions of maximum variance. If features are on different scales, PCA will be biased toward features with larger magnitudes.

* **K-Nearest Neighbors (KNN)** and **K-Means Clustering**:
  These rely on distance metrics. Unscaled features can dominate the distance computation.

* **Gradient Descent Optimization** (used in Logistic Regression, Neural Networks, etc.):
  Feature scaling ensures faster convergence because the loss function contours are more symmetrical.

* **Support Vector Machines (SVM)**:
  Scaling affects the computation of the margin and support vectors.

---

## Using preprocess

Preprocess offers you two ways to apply scaling operation on your dataset : 

- Selected numeric columns 
- All numeric columns 

### On selected numeric columns

If you wish to apply scaling operation, then apply `op = "scaling"` operation using the selected method : `minmax` or `zscore`. 

Here is an example where I wish to apply minmax scaling to the numeric columns `age` and `wage_eur` of the fifa dataset. 

```toml
[[preprocess.columns]]
name = "age"
operations = [
    {op = "scale", method = "minmax"}
]

[[preprocess.columns]]
name = "wage_eur"
operations = [
    {op = "scale", method = "minmax"}
]
```

You need to apply the operation on each of the column. This gives you full flexibility as to which operations to apply on which column. You may wish for example to apply the `fillna` operation on *wage_eur* in addition to the scale operation. 


```toml
[[preprocess.columns]]
name = "wage_eur"
operations = [
    {op = "fillna", method = "mean"},
    {op = "scale", method = "minmax"}
]
```

### On all numeric columns

In some case, you prefer to apply this same operation on all the numeric columns. 
To achieve this, you just need to apply the operation on the [preprocess.numeric] selector. 

Here's an example : 

```toml
[preprocess.numerics]
operations = [
    {op = "scale", method = "minmax"}
]
```

You just don't need to set the column name. 

An alternative syntax that TOML allows us to express this operation is the following 
```toml
[[preprocess.numerics.operations]]
op = "scale", 
method = "minmax"
```

Since [[preprocess.numerics.operations]] is a list of operations, you can use this syntax to define individually the operation. If you choose to use that syntax then you will need to add as many operations you like. 
