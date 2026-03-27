#!/usr/bin/env bash

set -e

REPO_OWNER="agailloty" 
REPO_NAME="preprocess"
VERSION="0.1.1"
BINARY_NAME="preprocess"
INSTALL_DIR="$HOME/.local/bin"

mkdir -p "$INSTALL_DIR"

# Detect OS
OS="$(uname -s)"
ARCH="$(uname -m)"

case "$OS" in
  Linux) OS="Linux" ;;
  Darwin) OS="Darwin" ;;
  *) echo "❌ Unsupported OS: $OS" && exit 1 ;;
esac

case "$ARCH" in
  x86_64) ARCH="x86_64" ;;
  arm64|aarch64) ARCH="arm64" ;;
  i386|i686) ARCH="i386" ;;
  *) echo "❌ Unsupported architecture: $ARCH" && exit 1 ;;
esac

FILENAME="${REPO_NAME}_${OS}_${ARCH}.tar.gz"
URL="https://github.com/${REPO_OWNER}/${REPO_NAME}/releases/download/${VERSION}/${FILENAME}"

TMP_DIR=$(mktemp -d)
cd "$TMP_DIR"

echo "➡️ Downloading $URL"
curl -sL "$URL" -o "$FILENAME"

echo "📦 Extracting archive..."
tar -xzf "$FILENAME"

echo "📁 Installing to $INSTALL_DIR"
mv "$BINARY_NAME" "$INSTALL_DIR/$BINARY_NAME"
chmod +x "$INSTALL_DIR/$BINARY_NAME"

cd -
rm -rf "$TMP_DIR"

# Check if $INSTALL_DIR is in PATH
if ! echo "$PATH" | grep -q "$INSTALL_DIR"; then
  echo "⚠️ $INSTALL_DIR is not in your PATH."
  echo "You may want to add this line to your shell profile:"
  echo "export PATH=\"\$HOME/.local/bin:\$PATH\""
fi

echo "✅ Installation complete. Run: $BINARY_NAME"