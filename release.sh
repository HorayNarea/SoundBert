#!/bin/sh

# force static build
export CGO_ENABLED=0

compile() {
	echo "Compiling for $1 on $2..."
	if [ "$1" = "windows" ]; then
		local EXE=".exe"
	fi
	GOOS=$1 GOARCH=$2 GOARM=$3 go build -o bin/SoundBert-$1-$2$EXE .
}

for OS in freebsd linux windows; do
	for ARCH in amd64 386; do
		compile $OS $ARCH
	done
done

# Mac OS X doesn't need 32-bit
compile darwin amd64

# build for Raspberry-Pi
compile linux arm 6
