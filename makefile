# Automatically set parallel jobs based on available CPU cores
MAKEFLAGS += -j$(shell nproc)

OUTPUT_DIR := dist

.DEFAULT_GOAL := build

ARCH_MAP_x86_64 := amd64
ARCH_MAP_arm64 := arm64

build: pkgstate-linux-x86_64 pkgstate-linux-arm64

check:
	staticcheck ./...

pkgstate-%:
	@GOOS=$(word 1,$(subst -, ,$*)) \
	GOARCH=$(ARCH_MAP_$(word 2,$(subst -, ,$*))) \
	go build -o $(OUTPUT_DIR)/pkgstate-$*

.PHONY: build
