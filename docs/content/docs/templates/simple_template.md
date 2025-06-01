---
weight: 100
title: "Simple template"
icon: "article"
date: "2025-06-01T20:06:40+02:00"
lastmod: "2025-06-01T20:06:40+02:00"
description: "Use basic operations"
draft: false
toc: true
---

This simple template shows the content of a `Prepfile.toml` that lets you read the a dataset called `data.csv`. In this Prepfile I decide to rename `column1` to `rename_column1`.  
I also choose to fill the missing values of the column `age` with the mean of the column. 

Then I export the result into a new dataset called `data_cleaned.csv`

```toml
[data]
filename = 'data.csv'
csv_separator = ','
decimal_separator = '.'
encoding = 'utf-8'
missing_identifier = ''

[preprocess]
[[preprocess.columns]]
name = 'column1'
type = 'string'
new_name = 'rename_column1'

[[preprocess.columns]]
name = 'age'
type = 'int'
operations = [{op = "fillna", method = "mean"}]



[postprocess]
format = 'csv'
filename = 'data_cleaned.csv'
```