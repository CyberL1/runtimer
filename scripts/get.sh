#!/bin/sh

set -e

if ! command -v unzip > /dev/null; then
  echo "Error: Unzip is required to install Runtimer"
  exit 1
fi

build=false

runtimer_dir="$HOME/.runtimer/bin"
runtimer_zip="$runtimer_dir/runtimer.zip"
runtimer_exe="$runtimer_dir/runtimer"

if [ "$OS" = "Windows_NT" ]; then
  target="windows"
else
  case $(uname -sm) in
  "Darwin x86_64") target="darwin-amd64" ;;
  "Darwin arm64") target="dawin-arm64" ;;
  "Linux aarch64") target="linux-arm64" ;;
  *) target="linux-amd64"
  esac
fi

download_url="https://github.com/CyberL1/runtimer/releases/latest/download/runtimer-${target}.zip"

if [ ! -d $runtimer_dir ]; then
  mkdir -p $runtimer_dir
fi

if $build; then
  go build -o $runtimer_exe
else
  curl --fail --location --progress-bar --output $runtimer_zip $download_url
  unzip -d $runtimer_dir -o $runtimer_zip
  chmod +x $runtimer_exe
  rm $runtimer_zip
fi

echo "Runtimer was installed to $runtimer_exe"
if command -v runtimer > /dev/null; then
  echo "Run 'runtimer --help' to get started"
else
  case $SHELL in
  /bin/zsh) shell_profile=".zshrc" ;;
  *) shell_profile=".bashrc" ;;
  esac
  echo "export PATH=\"$runtimer_dir:\$PATH\"" >> $HOME/$shell_profile
fi