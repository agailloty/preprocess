---
weight: 500
title: "preprocess version"
icon: "article"
date: "2025-06-01T19:56:40+02:00"
lastmod: "2026-09-18T11:30:00+02:00"
description: "Display the current version of preprocess CLI"
draft: false
toc: true
--- 

## Overview

The `preprocess version` command displays version information about your preprocess CLI installation. This is useful for:
- Verifying your installation
- Troubleshooting issues
- Ensuring compatibility
- Reporting bugs

## Basic Usage

### Display Version

```bash
preprocess version
```

This outputs the current version of preprocess CLI.

### Example Output

```
preprocess version 0.1.1
```

## Why Version Information Matters

### For Troubleshooting

When reporting issues or seeking help:

```bash
# Always include your version when reporting bugs
preprocess version
```

This helps maintainers:
- Identify if you're using an outdated version
- Reproduce issues in the same version
- Determine if the issue has already been fixed

### For Compatibility

Check version before using specific features:

```bash
# Check if your version supports a specific feature
preprocess version
```

### For Documentation

When following tutorials, ensure you're using a compatible version:

```bash
# Verify your version matches the tutorial requirements
preprocess version
```

## Checking for Updates

To check if you have the latest version:

1. Note your current version:
   ```bash
   preprocess version
   ```

2. Check the latest release on GitHub:
   ```bash
   # Using curl
   curl -s https://api.github.com/repos/agailloty/preprocess/releases/latest | grep tag_name
   ```

3. Compare with your version

## Best Practices

### Include in Bug Reports

Always include version information when reporting issues:

```bash
# When reporting a bug, run:
echo "preprocess version:"
preprocess version

echo "Operating System:"
uname -a  # Linux/Mac
echo %OS%  # Windows
```

### Log Version in Scripts

In your data processing scripts, log the version for reproducibility:

```bash
#!/bin/bash
# process_data.sh

echo "=== Processing Script Started ==="
echo "preprocess version: $(preprocess version)"
echo "Date: $(date)"
echo ""

# Rest of your script...
```

### Verify in CI/CD Pipelines

In your continuous integration workflows:

```yaml
# .github/workflows/data-processing.yml
- name: Check preprocess version
  run: |
    echo "preprocess version: $(preprocess version)"
    # Verify minimum version if needed
```

## Next Steps

After checking your version:

- **[Update]({{% relref "quickstart" %}})**: If outdated, follow the installation instructions to update
- **[Explore commands]({{% relref "guides" %}})**: Learn about all available commands
- **[Get started]({{% relref "quickstart" %}})**: Begin using preprocess CLI

## Quick Reference

| Task | Command |
|------|---------|
| Display version | `preprocess version` |
| Include in scripts | `$(preprocess version)` |
| Check for updates | Visit GitHub releases page |
