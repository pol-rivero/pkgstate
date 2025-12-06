#!/bin/sh

REPO_NAME="pol-rivero/pkgstate"

OS=$(uname -s)

if [ "$OS" = "Linux" ]; then
    echo "Detected OS: Linux"
    base_name="pkgstate-linux"
else
    echo "Your OS is not supported. Consider compiling from source."
    exit 1
fi

ARCH=$(uname -m)

if [ "$ARCH" = "x86_64" ] || [ "$ARCH" = "amd64" ] || [ "$ARCH" = "x64" ]; then
    echo "Detected architecture: x86_64"
    base_name="$base_name-x86_64"
elif [ "$ARCH" = "aarch64" ] || [ "$ARCH" = "arm64" ]; then
    echo "Detected architecture: arm64 (aarch64)"
    base_name="$base_name-arm64"
else
    echo "Your CPU architecture is not supported. Consider compiling from source."
    exit 1
fi

echo "Downloading the latest release..."

download_url=$(curl -s https://api.github.com/repos/$REPO_NAME/releases/latest \
    | grep "browser_download_url.*$base_name" \
    | cut -d '"' -f 4)

if [ $? -ne 0 ]; then
    echo "Failed to fetch the latest release information. Please check your internet connection or the GitHub API status."
    exit 1
fi

if [ -z "$download_url" ]; then
    echo "No suitable binary found for your OS and architecture. Please check the GitHub releases page."
    exit 1
fi

curl -L -o pkgstate "$download_url"
if [ $? -ne 0 ]; then
    echo "Failed to download the binary. Please check your internet connection or the URL: $download_url"
    exit 1
fi

chmod +x pkgstate
if [ $? -ne 0 ]; then
    echo "Failed to make ./pkgstate executable. Please check your permissions."
    exit 1
fi

echo "The binary will now be moved to /usr/local/bin/pkgstate, you may be prompted for your password."
sudo mv pkgstate /usr/local/bin/pkgstate
if [ $? -ne 0 ]; then
    echo "Failed to move the binary to /usr/local/bin. Please check your permissions."
    exit 1
fi

echo -e "\033[0;32mInstallation complete! You can now run pkgstate from anywhere in your terminal.\033[0m"
