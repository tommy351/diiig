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

## Design

This project use a linked list to store the list of topics. The list is sorted on insertion. There're some reasons why I chose linked list for storage:

- It takes less effort to make the list ordered. We just need to insert the element to the right position.
- Only the top N topics need to be displayed. So, we don't have to iterate over the whole list.

What can be improved?

- [ ] Try [skip list](https://en.wikipedia.org/wiki/Skip_list)? (like zset in Redis)
- [ ] Update elements in the linked list directly rather than remove and insert.
