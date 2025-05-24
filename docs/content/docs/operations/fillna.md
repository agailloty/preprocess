---
weight: 100
title: "Filling missing values"
icon: "article"
date: "2025-05-13T20:41:40+02:00"
lastmod: "2025-05-24T14:41:40+02:00"
description: "Filling missing values"
draft: false
toc: true
---

## Text column

### Fill with value 

You can easily fill missing values of a particular of the dataset using value. In the following example I am replacing all missing values with the value "Unknown". 


```toml
[[preprocess.columns]]
name = 'short_name'
type = 'string'
operations = [
    {op = "fillna", value = "Unknown"}
]
```

## Numeric column 

For numeric columns, there are few missing value replacement. 

### Fill with value 

```toml
[[preprocess.columns]]
name = 'weight_kg'
type = 'int'
operations = [
    {op = "fillna", value = 0}
]
```

### method = `mean`
Replaces missing values in a numeric column with the mean of the non-missing values. This maintains dataset size while minimizing bias introduced by missing data. Suitable when data is missing at random and the distribution is approximately symmetric.

```toml
[[preprocess.columns]]
name = 'weight_kg'
type = 'int'
operations = [
    {op = "fillna", method = "mean"}
]
```


### method = `median`

Replaces missing values in a numeric column with the median of the non-missing values. Useful for skewed distributions or when outliers are present, as the median is more robust to extreme values than the mean.

```toml
[[preprocess.columns]]
name = 'weight_kg'
type = 'int'
operations = [
    {op = "fillna", method = "median"}
]
```

