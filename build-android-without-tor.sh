export CGO_ENABLED=1
export CFLAGS='-w -s -O3'
export CGO_CFLAGS="$CFLAGS"
export CGO_CPPFLAGS="$CFLAGS"
export CGO_CXXFLAGS="$CFLAGS"
export CGO_FFLAGS="$CFLAGS"
export CGO_LDFLAGS="$CFLAGS"

export llvm_bin=$ANDROID_HOME/ndk-bundle/toolchains/llvm/prebuilt/linux-x86_64/bin/

echo 'Компиляция android-arm7 бинарника...'
export CC=$llvm_bin/armv7a-linux-androideabi16-clang
export CXX=$llvm_bin/armv7a-linux-androideabi16-clang++

mkdir -p ./builds/android/Infinite-Bomber-arm7-without-tor
GOOS=android GOARCH=arm GOARM=7 go build -tags withoutTor -o ./builds/android/Infinite-Bomber-arm7-without-tor/infinite-bomber -gcflags="all=-trimpath=$GOPATH" -asmflags="all=-trimpath=$GOPATH" -ldflags="-s -w"

echo 'Компиляция android-x86 бинарника...'
export CC=$llvm_bin/i686-linux-android16-clang
export CXX=$llvm_bin/i686-linux-android16-clang++

mkdir -p ./builds/android/Infinite-Bomber-x86-without-tor
GOOS=android GOARCH=386 go build -tags withoutTor -o ./builds/android/Infinite-Bomber-x86-without-tor/infinite-bomber -gcflags="all=-trimpath=$GOPATH" -asmflags="all=-trimpath=$GOPATH" -ldflags="-s -w"

echo 'Готово!'
