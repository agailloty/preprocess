---
weight: 501
date: "2025-05-26T21:40:40+02:00"
draft: false
author: "Axel-Cleris Gailloty"
title: "Prepfile"
icon: "developer_guide"
toc: true
description: "A reference of all Prepfile tags"
publishdate: "2025-05-26T21:40:40+02:00"
---

## Introducting Prepfile

If you have worked with tools such as Docker you'll notice that they give the user the ability to declare operations they wish to perform using simple files such as `Dockerfile` or `compose.yml`. 

Preprocess CLI implements a similar idea with the Prepfile.toml

The Prepfile.toml contains three sections : 
- Data section [data]
- Preprocess section [preprocess]
- Postprocess section [postprocess]

In the data section you provide all necessary information to read the dataset on which you want to apply the operations. 
The preprocess section contains the operations that will be applied on the dataset. 
The postprocess section contains what to do after the preprocessing steps. 


## Prepfile tags

## [data]


{{< table "table-hover" >}}
| Parameter | Default Value | Description |
|---------|--------|------|
| `filename` | N/A | The path to the dataset. This is a mandatory parameter. |
| `csv_separator` | "," | Specify the separator used for reading the CSV file. Accepts any single character : `";"`, `"\t"`, `" "` |
| `decimal_separator` | `"."` | Specify the decimal separator when reading the dataset. Possible values may include `","`, `"."` |
| `encoding` | `"utf-8"` | Specify the character set encoding. Possible values are `"utf-8"`, `"latin-1"` |
| `missing_identifier` | `""` | Specify how missing values are represented in the dataset. For most data it's empty `""` but some dataset may explicitly fill missing values with  `"NA"`, `"N/A"` ...  |
{{< /table >}}

## [preprocess]

### [preprocess.numerics]

{{< table "table-hover" >}}
| Parameter | Default Value | Description |
|---------|--------|------|
| `op` | N/A | The name of the operations  |
| `csv_separator` | "," | Specify the separator used for reading the CSV file. Accepts any single character : `";"`, `"\t"`, `" "` |
| `decimal_separator` | `"."` | Specify the decimal separator when reading the dataset. Possible values may include `","`, `"."` |
| `encoding` | `"utf-8"` | Specify the character set encoding. Possible values are `"utf-8"`, `"latin-1"` |
| `missing_identifier` | `""` | Specify how missing values are represented in the dataset. For most data it's empty `""` but some dataset may explicitly fill missing values with  `"NA"`, `"N/A"` ...  |
{{< /table >}}

### [preprocess.texts]

The following syntax 

```toml
[preprocess.texts]
operations = [
    {op = "fillna", method = "mean"}
]
```

is equivalent to : 

```toml
[[preprocess.texts.operations]]
op = "fillna"
method = "mean"
```

But you cannot mix them. Choose one style and keep it throughout the Prepfile.

### [[preprocess.columns]]

## [postprocess]

```toml
[postprocess]
dropcolumns = ["BALANCE","SURFACE_AREA","EXPORTS_GOOD_SERVICES"]
format = 'csv'
sortdataset = {descending = false}
filename = 'indicators_numerics_cleaned.csv'
```