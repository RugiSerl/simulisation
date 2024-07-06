#UPDATED FOR VERSION r26d (linux version)

# The process is in 2 steps : compiling as shared library and then building apk

# You must edit those values
ANDROID_NDK_HOME=/opt/android-ndk-r26d # Path of the ndk
ANDROID_HOME=~/Android/Sdk/            # Path of the sdk
ANDROID_API=21                         # Make sure in the $BIN folder defined below that a compiler exists for that version
ARCH_64_BIT=false # doesn't work for now so keep to false


# setting directories variables
ANDROID_TOOLCHAIN=${ANDROID_NDK_HOME}/toolchains/llvm/prebuilt/linux-x86_64/bin
ANDROID_SYSROOT=${ANDROID_NDK_HOME}/toolchains/llvm/prebuilt/linux-x86_64/sysroot
ANDROID_TOOLCHAIN=${ANDROID_NDK_HOME}/toolchains/llvm/prebuilt/linux-x86_64
BIN=${ANDROID_NDK_HOME}/toolchains/llvm/prebuilt/linux-x86_64/bin

# COMPILING SHARED LIBRARY

if "$ARCH_64_BIT" 
then
        export ARCH="aarch64-linux-android"
        export GOARCH=arm64
else
        export ARCH="arm-linux-androideabi"
        export GOARCH=arm 
        
fi

export CC="${BIN}/armv7a-linux-androideabi${ANDROID_API}-clang"
export CGO_CFLAGS="-I${ANDROID_SYSROOT}/usr/include -I${ANDROID_SYSROOT}/usr/include/${ARCH} --sysroot=${ANDROID_SYSROOT}" 
export CGO_LDFLAGS="-L${ANDROID_SYSROOT}/usr/lib/${ARCH}/${ANDROID_API} -L${ANDROID_TOOLCHAIN}/lib --sysroot=${ANDROID_SYSROOT}" 
export CGO_ENABLED=1
export GOOS=android

go build -buildmode=c-shared -ldflags="-s -w -extldflags=-Wl,-soname,libexample.so" -o=android/libs/armeabi-v7a/libexample.so &&

# BUILDING APK
./gradlew assembleDebug

# Optionnal : install on your device (you need adb and debug mode on your device)
# You can delete those lines if you don't want them
adb install android/build/outputs/apk/debug/android-debug.apk
