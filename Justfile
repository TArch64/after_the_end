#!/usr/bin/env just --justfile

compile_qrc:
  go tool miqt-rcc -Package resources -Input ./app/resources/resources.qrc

build: compile_qrc
  CC='clang' CXX='clang++' CGO_CXXFLAGS='-std=c++17' go build -ldflags "-s -w" -o dist/game .

build_debug: compile_qrc
  CC='clang' CXX='clang++' CGO_CXXFLAGS='-std=c++17' go build -gcflags="all=-N -l" -o dist/game .