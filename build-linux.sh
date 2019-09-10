export CGO_ENABLED=1
export CFLAGS='-w -s -O3'
export CGO_CFLAGS="$CFLAGS"
export CGO_CPPFLAGS="$CFLAGS"
export CGO_CXXFLAGS="$CFLAGS"
export CGO_FFLAGS="$CFLAGS"
export CGO_LDFLAGS="$CFLAGS"

# echo 'Обновление зависимостей go-libtor'
# go get -u github.com/ipsn/go-libtor

echo 'Компиляция linux-x86 бинарника...'
mkdir -p ./builds/linux/Infinite-Bomber-x86
GOOS=linux GOARCH=386 CGO_ENABLED=1 go build -o ./builds/linux/Infinite-Bomber-x86/infinite-bomber -gcflags="all=-trimpath=$GOPATH" -asmflags="all=-trimpath=$GOPATH" -ldflags="-s -w"

echo 'Копирование файлов тора...'
mkdir -p ./builds/linux/Infinite-Bomber-x86/tor
cp -r ./tor-files/Data ./builds/linux/Infinite-Bomber-x86/tor/
cp -r ./tor-files/torrc ./builds/linux/Infinite-Bomber-x86/tor/

echo 'Компиляция linux-x64 бинарника...'
mkdir -p ./builds/linux/Infinite-Bomber-x64
GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o ./builds/linux/Infinite-Bomber-x64/infinite-bomber -gcflags="all=-trimpath=$GOPATH" -asmflags="all=-trimpath=$GOPATH" -ldflags="-s -w"

echo 'Копирование файлов тора...'
mkdir -p ./builds/linux/Infinite-Bomber-x64/tor
cp -r ./tor-files/Data ./builds/linux/Infinite-Bomber-x64/tor/
cp -r ./tor-files/torrc ./builds/linux/Infinite-Bomber-x64/tor/

echo 'Готово!'
