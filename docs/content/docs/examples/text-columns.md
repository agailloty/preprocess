---
weight: 100
title: "Apply operations on text columns"
icon: "article"
date: "2025-05-20T19:40:40+02:00"
lastmod: "2025-05-20T19:40:40+02:00"
description: "Apply preprocessing operations on all the text columns of the dataset"
draft: false
toc: true
---

This allows you to apply the same operations on all the text columns that are in the dataset. You can specify as much operation needed. 
Operations are applied sequentially (in order) on each column. 

In the following example I apply the `triws` operation to remove all whitespace (left and reight) from all entries of the text columns. 

After that, I want all text column to be in upper case. 

```toml
[data.texts]
preprocess = [
    {op = "clean", method = "trimws"},
    {op = "clean", method = "upper"}
]
```