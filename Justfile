#!/usr/bin/env just --justfile

qt_version := `qmake --version | grep 'Qt' | awk 'match($0, /[0-9]+\.[0-9]+\.[0-9]+/) {print substr($0, RSTART, RLENGTH)}'`
qtbase_lib := "/opt/homebrew/Cellar/qtbase" / qt_version / "lib"

export CC := "clang"
export CXX := "clang++"
export CGO_CXXFLAGS := f"-std=c++23 -F{{ qtbase_lib }}"
export CGO_LDFLAGS := f"-F{{ qtbase_lib }} -framework QtOpenGL -framework QtGui -framework QtWidgets -framework QtCore -framework QtOpenGLWidgets"

data_dir := x"~/Library/Application\\ Support/ua.tarch64.AfterTheEnd"
db_path := data_dir / 'application.db'

compile_qrc:
  go tool miqt-rcc -Qt6 -Package resources -Input ./app/resources/resources.qrc

build_envs:
  env | grep -E '^(CC|CXX|CGO_CXXFLAGS|CGO_LDFLAGS)=' > .env.build

build: build_envs compile_qrc
  go build -ldflags "-s -w" -o dist/game .

build_debug: build_envs compile_qrc
  go build -gcflags="all=-N -l" -o dist/game .

db_drop:
  rm {{ db_path }}
