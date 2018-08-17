# go-microservice-pubsub-template


[![CircleCI](https://circleci.com/gh/johan-lejdung/go-microservice-pubsub-template.svg?style=svg)](https://circleci.com/gh/johan-lejdung/go-microservice-pubsub-template)
[![codecov](https://codecov.io/gh/johan-lejdung/wiseer/branch/master/graph/badge.svg?token=2b779e51-4fbb-442d-96e8-d1af21051285)](https://codecov.io/gh/johan-lejdung/go-microservice-pubsub-template)

A template for a go microservice with a PubSub connection.

For a service configured with an REST-API look at https://github.com/johan-lejdung/go-microservice-api-template

Features:
- Dependency injection
- Database connection
- Database migrations
- Fluentd logging
- PubSub setup with Protobuf
- Kubernetes compatible logging/liveness endpoints
- CricleCI + CodeCov configuration files

# Make it yours!

## Rename

Start by forking this repository or just download the code.

Replace all occourances of `johan-lejdung` with whatever your github username/organization is.

Replace all occourances of `go-microservice-pubsub-template` with whatever you are calling your project.

## Env Config

Open up `.env` and change the settings nessicary.

## CircleCI and CodeCov

Simple register an account at CircleCI and CodeCov, copy the CodeCov token to the environment variables on CircleCI like so:

`CODECOV_TOKEN` = `<TOKEN>`

# Installation
Add the following lines to ~/.bach_profile or ~/.zshrc (if using zsh)

    export GOPATH=/Users/username/go

    export PATH=$GOPATH/bin:$PATH

Where username is the username of your profile.

Then install dep:

```
brew install dep
```

Run this command in the correct folder:

```
dep ensure
```

# Testing

Install mockery

```
go get github.com/vektra/mockery/.../
```

Generate files
```
go generate ./....
```

Run tests
```
go test ./...
```

# Code Guide

## PubSub
The pubsub connection is found in the `ps folder` and has the service it uses injected to it's structure.

Take a look in the `consume` and `produce` folders to check out how to subscribe and connect to topics.

# Protobuf
A generated protobuf file is in the `protomsg` packages, together with the raw proto file.

To generate a proto file install `protoc` and execute `protoc --go_out=./go/ ./<PROTO_FILE>`

## Database
The database will automatically apply migrations if the variable in `.env` called `ENV` is either `dev` or `test`.

I am using https://github.com/mattes/migrate for database migrations.

Run a local database in docker with `docker-compose up`.

## Bootstrap
I am using https://github.com/facebookgo/inject for dependency injection.

By injecting implementations of interfaces in the `bootstrapApp` you can easily inject them in structs such as:

```
type TestStruct struct {
    Variable pkg.InterfaceType `inject:""`
}
```

## Vendor folder
I have the vendor folder checked into the repo, for reproducibility.
