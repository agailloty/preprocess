---
weight: 300
title: "preprocess summary"
icon: "article"
date: "2025-06-01T19:52:40+02:00"
lastmod: "2025-06-01T19:52:40+02:00"
description: "Produce summary statistics"
draft: false
toc: true
---

```powershell
preprocess summary --help
```

```powershell
Generate dataset summary statistics.
        By default it generates a TOML file that contains all summary statistics.
        You can choose to generate a HTML file with the --html flag.

Usage:
  preprocess summary [flags]

Flags:
  -d, --data string           Path to the dataset
  -m, --dsep string           Decimal separator (default ".")
  -e, --encoding string       Character encoding (default "utf-8")
      --exclude stringArray   Exclude columns from summary
  -h, --help                  help for summary
  -t, --html                  Generate HTML file
  -o, --output string         Output name for Summaryfile (default "Summaryfile.toml")
  -f, --prepfile string       Path to the configuration file (default "Prepfile.toml")
  -s, --sep string            Separator for csv file (default ",")
```