---
weight: 100
title: "preprocess run"
icon: "article"
date: "2025-05-13T19:56"
lastmod: "2025-05-13T19:56"
description: "Generate Prefile"
draft: false
toc: true
---

Use preprocess run to run the operations on the Prepfile. If no Prepfile is specified, `preprocess run` will will use the Prepfile that is in the current folder.

```
Run operations using Prepfile

Usage:
  preprocess run [flags]

Flags:
      --column stringArray   Target column(s) for preprocessing
  -d, --data string          Path to the dataset file
  -f, --file string          Path to the configuration file (default "Prepfile.toml")
  -h, --help                 help for run
      --numerics             Apply operations only on numeric columns
      --op stringArray       Preprocessing operation(s) (e.g., fillna:method=mean)
  -s, --sep string           Csv separator (default ",")
```