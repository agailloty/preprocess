---
weight: 101
date: "2025-05-13T13:15:40+02:00"
draft: false
author: "Axel-Cleris Gailloty"
title: "Quickstart"
icon: "rocket_launch"
toc: true
description: "A quickstart guide to install preprocess"
publishdate: "2025-05-13T13:15:40+02:00"
tags: ["Beginners"]
---


## Installation

Install the [preprocess CLI](https://github.com/gohugoio/hugo/releases/latest), using the specific instructions for your operating system below:

{{< tabs tabTotal="4">}}
{{% tab tabName="Linux and MacOs" %}}

Use `curl` to download the script and execute it with sh:

```shell
curl -LsSf https://astral.sh/uv/install.sh | sh
```

If your system doesn't have `curl`, you can use wget:

```shell
wget -qO- https://astral.sh/uv/install.sh | sh
```

{{% /tab %}}
{{% tab tabName="Windows" %}}


```shell
powershell -ExecutionPolicy ByPass -c "irm https://astral.sh/uv/install.ps1 | iex"
```

{{% /tab %}}
{{< /tabs >}}

### Manual Installation

The preprocess GitHub repository contains pre-built versions of the preprocess command-line tool for various operating systems, which can be found on the [Releases page](https://github.com/gohugoio/hugo/releases/latest)

At the release section, download the release for your operating system. You may need to set the environment variable to the location where you save the binary so you can call `preprocess` from any terminal on your computer. 

