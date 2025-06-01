---
weight: 110
title: "Numeric operations"
icon: "article"
date: "2025-06-01T20:06:40+02:00"
lastmod: "2025-06-01T20:06:40+02:00"
description: "Apply operations on numeric columns only"
draft: false
toc: true
---

This `Prepfile.toml` template lets you read the a dataset called `data.csv`and apply operations only on the numeric columns. 

Then I export the result into a new dataset called `data_cleaned.csv`

```toml
[data]
filename = 'data.csv'
csv_separator = ','
decimal_separator = '.'
encoding = 'utf-8'
missing_identifier = ''

[preprocess]
[[preprocess.numerics]]
operations = [
    {op = "fillna", method = "median"},
    {op = "scale", method = "minmax"}
]

[postprocess]
format = 'csv'
filename = 'data_cleaned.csv'
```