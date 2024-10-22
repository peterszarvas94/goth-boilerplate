# Project goth

This repo is archived in favor of [goat](http://github.com/peterszarvas94/goat)

This is a minimalist but somewhat opinionated starter template for building full stack applications

Stack:

- go
- templ
- htmx
- tailwind
- turso db

## Features

- sample pages
- auth
- go's built-in router
- middlewares
- json logger
- env variable parsing
- dark mode (free)

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## MakeFile

run all make commands with clean tests
```bash
make all build
```

build the application
```bash
make build
```

run the application
```bash
make run
```

run templ codegen
```bash
make templ
```

run tailwind codegen
```bash
make tw

```

live reload the application
```bash
make watch
```

run the test suite
```bash
make test
```

clean up binary from the last build
```bash
make clean
```
