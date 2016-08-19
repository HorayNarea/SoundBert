#!/bin/sh

# force static build
export CGO_ENABLED=0

export SRCDIR=$(pwd)
mkdir -p $SRCDIR/bin

release() {
	local TMPDIR=$(mktemp -d)
	mkdir $TMPDIR/SoundBert
	compile $TMPDIR $1 $2 $3
	pack $TMPDIR $1 $2
	rm -rf $TMPDIR
	echo -n "\n"
}

pack() {
	echo "Packing for $2 on $3..."
	cd $1
	cp $SRCDIR/config.default.toml $1/SoundBert/config.default.toml
	if [ "$2" = "windows" ]; then
		zip -rq SoundBert-$2-$3.zip SoundBert
	else
		tar cfz SoundBert-$2-$3.tar.gz SoundBert
	fi
	cd $SRCDIR
	cp -f $1/SoundBert-* $SRCDIR/bin/.
}

compile() {
	echo "Compiling for $2 on $3..."

	if [ "$2" = "windows" ]; then
		local EXE=".exe"
	fi

	if [ "$2-$3" = "linux-amd64" ]; then
		local CGO_ENABLED=1
		local CC=musl-clang
		local LDFLAGS="-ldflags '-linkmode=external -extldflags=-static'"
	fi

	GOOS=$2 GOARCH=$3 GOARM=$4 go build $LDFLAGS -o $1/SoundBert/SoundBert-$2-$3$EXE .
}

# Mac OS X doesn't need 32-bit
release darwin amd64

for OS in windows freebsd linux; do
	for ARCH in amd64 386; do
		release $OS $ARCH
	done
done

# build for Raspberry-Pi
release linux arm 6
