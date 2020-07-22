<a href="#">
    <img src="docs/logo.svg" width="100%" height="140">
</a>

<p align="center">
    <a href="https://goreportcard.com/report/github.com/dsrhub/dsrhub" target="_blank">
        <img src="https://goreportcard.com/badge/github.com/dsrhub/dsrhub">
    </a>
    <a href="https://circleci.com/gh/dsrhub/dsrhub" target="_blank">
        <img alt="CircleCI" src="https://img.shields.io/circleci/build/github/dsrhub/dsrhub?style=flat-square">
    </a>
    <a href="https://github.com/dsrhub/dsrhub/releases" target="_blank">
        <img src="https://img.shields.io/github/v/release/dsrhub/dsrhub?color=green&sort=semver&style=flat-square">
    </a>
</p>

# Introduction
DSRHub is an open platform for DSR (Data Subject Requests). It helps to orchestrate the workflow of privacy related compliance processes (e.g. GDPR/CCPA) among your microservices. 

DSRHub was built to ease the pain of managing microservices' data privacy related processes. The best practice of microservices is that data should be owned by the service. This adds complexity to managing personal identifiable information (PII) at scale, because usually the APIs of hanlding the PII data is not standardized.

The goal of DSRHub is to extract the common logic of handling data subject requests within the microservices ecosystem into standard APIs and workflows. It ships a gRPC/swagger based API specification for your services and an YAML DSL workflow to orchestrate the data subject requests in a declarative manner.

# Features
Integration with the microservices.
- Standard gRPC/REST APIs with the support for identity resolution and DSR handling
- Standard YAML DSL based workflow definition
- Native support as a gRPC client to work with gRPC services
- Native support for OpenAPI/swagger client requests
- You own your data, 100%

Workflow orchestration (thanks to [uTask](https://github.com/ovh/utask)).
- Workflow orchestration
- Fault tolerance (retries, backoff)
- Encryption-at-rest for all the workflow's input/output
- Default 30 days (configurable) data retention TTL for max privacy protection 

# Quick start

## Run with docker-compose

```sh
wget https://raw.githubusercontent.com/dsrhub/dsrhub/master/docker-compose.example.yaml
docker-compose up
```

# Usage
WIP

# Config and Customization
WIP

# Metrics and Monitoring
WIP

# License
- [dsrhub/dsrhub](https://github.com/dsrhub/dsrhub): [Apache-2.0 License](https://github.com/dsrhub/dsrhub/blob/master/LICENSE)
- [ovh/utask](https://github.com/ovh/utask): [BSD 3-Clause License](https://github.com/ovh/utask/blob/master/LICENSE)

