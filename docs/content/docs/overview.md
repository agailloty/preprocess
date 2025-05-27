---
weight: 100
title: "Overview"
icon: "article"
date: "2025-05-10T13:19:40+02:00"
lastmod: "2025-05-10T13:19:40+02:00"
description: "A brief presentation of preprocess and what it aims to solve"
draft: false
toc: true
---

preprocess is a fast command line tool to preprocess data. 
It aims to be simple to use, portable and fast. 

## Simple to use

Declare your operations in a Toml file and run these operations using `preprocess run`.

## Portable 

It ships as a standalone binary. You just need this binary on your system to be ready to apply preprocessing operations on your dataset. 

## Fast

Because it is written with Go, you expect faster execution time for these operations than if you were to write them from scratch in popular data analysis languages.