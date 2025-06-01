---
weight: 100
title: "preprocess init"
icon: "article"
date: "2025-05-13T19:56:40+02:00"
lastmod: "2025-06-01T20:00:40+02:00"
description: "Generate Prefile"
draft: false
toc: true
---

Use preprocess init to generate a Prepfile. A Prepfile is a file that contains all the specifications to be applied to the data. 


```powershell
This command is used to generate a Prepfile in which you can specify the preprocessing computations either on specifics columns of the dataset or on whole numeric or text columns.

Usage:
  preprocess init [flags]

Flags:
  -d, --data string       Path to the dataset
  -m, --dsep string       Decimal separator (default ".")
  -e, --encoding string   Character encoding (default "utf-8")
  -h, --help              help for init
  -o, --output string     Output name for Prepfile (default "Prepfile.toml")
  -s, --sep string        Separator for csv file (default ",")
  -t, --template          Generate example Prepfile.toml template
```