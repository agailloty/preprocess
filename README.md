preprocess is a fast cross-platform data analysis preprocessing tool.

## Installation

Install the [preprocess CLI](https://github.com/agailloty/preprocess/releases/latest), using the specific instructions for your operating system below:

### Linux and MacOs"

Use `curl` to download the script and execute it with sh:

```sh
curl -LsSf https://preprocess-cli.netlify.app/install.sh | sh
```

If your system doesn't have `curl`, you can use wget:

```sh
wget -qO- https://preprocess-cli.netlify.app/install.sh | sh
```

### Windows"


```sh
powershell -ExecutionPolicy ByPass -c "irm https://preprocess-cli.netlify.app/install.ps1 | iex"
```


### Manual Installation

The preprocess GitHub repository contains pre-built versions of the preprocess command-line tool for various operating systems, which can be found on the [Releases page](https://github.com/agailloty/preprocess/releases/latest)

At the release section, download the release for your operating system. You may need to set the environment variable to the location where you save the binary so you can call `preprocess` from any terminal on your computer. 