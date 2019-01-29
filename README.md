MPExporter [![GoDoc](https://godoc.org/github.com/jtsiros/mpexporter?status.svg)](https://godoc.org/github.com/jtsiros/mpexporter) ![Version](https://img.shields.io/badge/version-0.1.0-green.svg)
====

MPExporter exports your Macpass/Keepass compatible XML export file to CSV.  This is particularly useful for password managers such as 1Password or LastPass that can import your data from CSV. Only 1Password is supported at the moment.

## Installation
    go get github.com/jtsiros/mpexporter

## Build
    go build cmd/mpexporter/mpexporter.go

## Usage

`./mpexporter --export 1password --i macpass.xml --o sample.csv`

