---
weight: 100
title: "preprocess init"
icon: "article"
date: "2025-05-13T19:56"
lastmod: "2025-05-13T19:56"
description: "Generate Prefile"
draft: false
toc: true
---

Use preprocess init to generate a Prepfile. A Prepfile is a file that contains all the specifications to be applied to the data. 


```
Generate a Prepfile

Usage:
  preprocess init [flags]

Flags:
  -f, --file string     Path to the dataset
  -h, --help            help for init
  -o, --output string   Output name for Prepfile (default "Prepfile.toml")
  -s, --sep string      Separator for csv file (default ",")
  -t, --template        Generate example Prepfile.toml template
```