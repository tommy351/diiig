# diiig

[![Build Status](https://travis-ci.org/tommy351/diiig.svg?branch=master)](https://travis-ci.org/tommy351/diiig)

[Demo](https://diiig.herokuapp.com/)

## Installation

You have to [install dep](https://golang.github.io/dep/docs/installation.html) before getting started.

``` sh
dep ensure
```

## Usage

Start the server. The server listens on [localhost:4000](http://localhost:4000/) by default.

``` sh
go run cmd/diiig-server/main.go
```

Run tests.

``` sh
go test ./...
```
