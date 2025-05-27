---
weight: 100
title: "Apply operations on numeric columns"
icon: "article"
date: "2025-05-20T18:40:40+02:00"
lastmod: "2025-05-24T14:50:40+02:00"
description: "Apply preprocessing operations on all the numeric columns of the dataset"
draft: false
toc: true
---

## Apply operations

```toml
[data]
file = '.\fifa_players_22.csv'
separator = ','

[preprocess.numerics]
operations = [
    {op = "fillna", method = "mean"},
    {op = "normalize", method = "zscore"}
]

[postprocess]
dropcolumns = ["preferred_foot", "body_type"]
format = 'csv'
filename = 'fifa_players_22_cleaned.csv'
```

```toml
[data]
file = '.\indicators_numerics.csv'
separator = ','

[preprocess.numerics]
operations = [
    {op = "normalize", method = "zscore"}
]

[postprocess]
dropcolumns = ["BALANCE","SURFACE_AREA","EXPORTS_GOOD_SERVICES"]
format = 'csv'
sortdataset = {descending = false}
filename = 'indicators_numerics_cleaned.csv'
```