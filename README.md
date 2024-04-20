# Lotso Airdrop Server

[![Go](https://github.com/BTI-US/Lotso_Airdrop_Server/actions/workflows/go.yml/badge.svg)](https://github.com/BTI-US/Lotso_Airdrop_Server/actions/workflows/go.yml)
[![Docker CI](https://github.com/BTI-US/Lotso_Airdrop_Server/actions/workflows/docker-image.yml/badge.svg)](https://github.com/BTI-US/Lotso_Airdrop_Server/actions/workflows/docker-image.yml)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

- Last Modified: 2024-04-20
- Author: Fu Ming

## Table of Contents

- [Lotso Airdrop Server](#lotso-airdrop-server)
  - [Table of Contents](#table-of-contents)
  - [Introduction](#introduction)
  - [Features](#features)
  - [Requirements](#requirements)
  - [Diagram](#diagram)
  - [Usage](#usage)
  - [License](#license)

## Introduction

This is a server for the Lotso Airdrop project. It is a simple server that listens for incoming requests and responds with the appropriate data. The server is written in Go, runs in Docker containers, and uses a MySQL database to store data, designed to be scalable and fault-tolerant, with the ability to handle a large number of requests simultaneously.

## Features

- [] Need to be updated

## Requirements

- [Go 1.21](https://golang.org/)
- [Docker](https://www.docker.com/)
- API Key from [Pocket Network](https://www.pokt.network/) for the production environment
- API Key from [Alchemy](https://www.alchemy.com/) for the test-net environment

## Diagram

- [] Need to be updated

## Usage

(Base) Start the docker container for the production environment:

```bash
TZ=UTC DEBUG=False IMG_NAME=bti/lotso-airdrop-server:0.0.1 MYSQL_PASSWD="your_mysql_database_password" DB_NAME=lotso_airdrop API_URL=https://base-pokt.nodies.app/ PRIVATE_KEY="your_private_key"  CHAIN_ID=8453 CUTOFF_BLOCK=0xbede7c CONTRACT_ADDRESS=0x23da3D470325660208c8beB29c3b80f70f08bcac docker-compose -f /home/fuming/docker/LotsoAirdropServer/docker-compose.yaml up -d
```

(Base) Stop the docker container for the production environment:

```bash
TZ=UTC DEBUG=False IMG_NAME=bti/lotso-airdrop-server:0.0.1 MYSQL_PASSWD="your_mysql_database_password" DB_NAME=lotso_airdrop API_URL=https://base-pokt.nodies.app/ PRIVATE_KEY="your_private_key"  CHAIN_ID=8453 CUTOFF_BLOCK=0xbede7c CONTRACT_ADDRESS=0x23da3D470325660208c8beB29c3b80f70f08bcac docker-compose -f /home/fuming/docker/LotsoAirdropServer/docker-compose.yaml down
```

(Sepolia) Start the docker container for the test-net environment:

```bash
TZ=UTC DEBUG=True IMG_NAME=bti/lotso-airdrop-server:0.0.1 MYSQL_PASSWD="your_mysql_database_password" DB_NAME=lotso_airdrop API_URL=https://eth-sepolia.g.alchemy.com/v2/crycCDOIbtpZREF8mwIO4AJdC3dK4ihU PRIVATE_KEY="your_private_key" CHAIN_ID=11155111 CUTOFF_BLOCK=0x556df8 CONTRACT_ADDRESS=0xA648a901DCd3dc15FBd0bee0FC0ee03279ce1d29 docker-compose -f /home/fuming/docker/LotsoAirdropServer_Sepolia/docker-compose.yaml up
```

(Sepolia) Stop the docker container for the test-net environment:

```bash
TZ=UTC DEBUG=True IMG_NAME=bti/lotso-airdrop-server:0.0.1 MYSQL_PASSWD="your_mysql_database_password" DB_NAME=lotso_airdrop API_URL=https://eth-sepolia.g.alchemy.com/v2/crycCDOIbtpZREF8mwIO4AJdC3dK4ihU PRIVATE_KEY="your_private_key" CHAIN_ID=11155111 CUTOFF_BLOCK=0x556df8 CONTRACT_ADDRESS=0xA648a901DCd3dc15FBd0bee0FC0ee03279ce1d29 docker-compose -f /home/fuming/docker/LotsoAirdropServer_Sepolia/docker-compose.yaml down
```

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details.