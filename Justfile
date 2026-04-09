#!/usr/bin/env just --justfile

build:
  CC='clang' CXX='clang++' CGO_CXXFLAGS='-std=c++17' go build -ldflags "-s -w" -o dist/game .
