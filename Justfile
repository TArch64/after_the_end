#!/usr/bin/env just --justfile

qt_home := x"~/Qt/6.10.3/macos"
qt_bin := qt_home / 'libexec'
qt_lib := qt_home / 'lib'
qt_pkgconfig := qt_lib / 'pkgconfig'

export CC := "clang"
export CXX := "clang++"
export CGO_CXXFLAGS := "-std=c++23 -D__yield=__builtin_arm_yield"
export CGO_LDFLAGS := f"-Wl,-rpath,{{ qt_lib }}"
export PKG_CONFIG_PATH := qt_pkgconfig

data_dir := x"~/Library/Application\\ Support/ua.tarch64.AfterTheEnd"
db_path := data_dir / 'application.db'

compile_qrc:
  go tool miqt-rcc -Qt6 -RccBinary {{ qt_bin }}/rcc  -Package resources -Input ./app/resources/resources.qrc

build: compile_qrc
  go build -ldflags "-s -w" -o dist/game .

build_debug: compile_qrc
  go build -gcflags="all=-N -l" -o dist/game .

db_drop:
  rm {{ db_path }}
