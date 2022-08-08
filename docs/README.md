
**Table of Contents**

- [Dealership Service](#dealership-service)
  - [About the project](#about-the-project)
    - [API docs](#api-docs)
    - [Design](#design)
    - [Status](#status)
    - [See also](#see-also)
  - [Getting started](#getting-started)
    - [Layout](#layout)
  - [Notes](#notes)



# Dealership Service

## About the project

This is a service to provide information about vehicle manufacturer inventory.

The project's technical stack consists of the following technology:
- Go
- Google Remote Procedural Calls (GRPC)

### API docs

GenDocu GRPC documentation is available:
https://doc.gendocu.com/thanders/api/dealershipservice

### Design

The template follows project convention doc.

* [Repository Conventions](https://github.com/caicloud/engineering/blob/master/guidelines/repo_conventions.md)

### Status

The template project is in alpha status.

## Getting started

Below we describe the conventions or tools specific to golang project.

### Layout

```tree
├── .github
├── .gitignore
|── bin
|    └── client
|    └── server
|── build
|    └── snapcraft.yaml
├── CHANGELOG.md
├── docs
│   └── README.md
├── Makefile
├── pkg
|    └── client
|    └── proto
|    └── server
├── README.md
├── test
|   └── test_make.sh

```

A brief description of the layout:

* `.github` has two template files for creating PR and issue. Please see the files for more details.
* `.gitignore` varies per project, but all projects need to ignore `bin` directory.
* `.golangci.yml` is the golangci-lint config file.
* `Makefile` is used to build the project.
* `CHANGELOG.md` contains auto-generated changelog information.
* `README.md` is a detailed description of the project.
* `bin` is to hold binary build outputs.
* `build` contains scripts, yaml files, dockerfiles, etc, to build and package the project.
* `docs` for project documentations.
* `release` [chart](https://github.com/caicloud/charts) for production deployment.
* `test` holds all tests (except unit tests), e.g. integration, e2e tests.
* `third_party` for all third party libraries and tools, e.g. swagger ui, protocol buf, etc.

## Notes
* All build files **MUST** be put under `build` directory.


// builds the di/server
```make buildServer```
// builds the di/client
```make buildClient```

// generates the Proto Buffer files
```make generateServer```

go run instance.go

Environment variables for Go install

```export GOPATH=$HOME/go```

```export GOBIN=$GOPATH/bin```

```export PATH=$PATH:$GOROOT:$GOPATH:$GOBIN```

