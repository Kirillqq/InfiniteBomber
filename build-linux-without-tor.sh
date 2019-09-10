export CGO_ENABLED=1
export CFLAGS='-w -s -O3'
export CGO_CFLAGS="$CFLAGS"
export CGO_CPPFLAGS="$CFLAGS"
export CGO_CXXFLAGS="$CFLAGS"
export CGO_FFLAGS="$CFLAGS"
export CGO_LDFLAGS="$CFLAGS"

echo 'Компиляция linux-x86 бинарника...'
mkdir -p ./builds/linux/Infinite-Bomber-x86-without-tor
GOOS=linux GOARCH=386 CGO_ENABLED=1 go build -tags withoutTor -o ./builds/linux/Infinite-Bomber-x86-without-tor/infinite-bomber -gcflags="all=-trimpath=$GOPATH" -asmflags="all=-trimpath=$GOPATH" -ldflags="-s -w"

echo 'Компиляция linux-x64 бинарника...'
mkdir -p ./builds/linux/Infinite-Bomber-x64-without-tor
GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -tags withoutTor -o ./builds/linux/Infinite-Bomber-x64-without-tor/infinite-bomber -gcflags="all=-trimpath=$GOPATH" -asmflags="all=-trimpath=$GOPATH" -ldflags="-s -w"

echo 'Готово!'
