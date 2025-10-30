#!/bin/bash

PKGBUILD="pkgstate-PKGBUILD"
PKGBUILD_BIN="pkgstate-bin-PKGBUILD"
PKGBUILD_GIT="pkgstate-git-PKGBUILD"

cp .github/pkgstate.PKGBUILD.template $PKGBUILD
cp .github/pkgstate-bin.PKGBUILD.template $PKGBUILD_BIN
cp .github/pkgstate-git.PKGBUILD.template $PKGBUILD_GIT

VERSION=$1
TARBALL_CHECKSUM=$2
LINUX_X86_CHECKSUM=$(sha256sum dist/pkgstate-linux-x86_64 | cut -d ' ' -f 1)
LINUX_ARM64_CHECKSUM=$(sha256sum dist/pkgstate-linux-arm64 | cut -d ' ' -f 1)

replace_var() {
    local var_name="$1"
    local var_value="${!var_name}"
    if [ -z "$var_value" ]; then
        echo "Error: Variable $var_name is not set."
        exit 1
    fi
    sed -i "s/{{$var_name}}/$var_value/g" "$2"
}
available_vars=(
    "VERSION"
    "TARBALL_CHECKSUM"
    "LINUX_X86_CHECKSUM"
    "LINUX_ARM64_CHECKSUM"
)

for var in "${available_vars[@]}"; do
    replace_var "$var" $PKGBUILD
    replace_var "$var" $PKGBUILD_BIN
    replace_var "$var" $PKGBUILD_GIT
done
