# Lotso Airdrop Server

[![Go](https://github.com/BTI-US/Lotso_Airdrop_Server/actions/workflows/go.yml/badge.svg)](https://github.com/BTI-US/Lotso_Airdrop_Server/actions/workflows/go.yml)
[![Docker CI](https://github.com/BTI-US/Lotso_Airdrop_Server/actions/workflows/docker-image.yml/badge.svg)](https://github.com/BTI-US/Lotso_Airdrop_Server/actions/workflows/docker-image.yml)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

- Last Modified: 2024-04-26
- Author: Fu Ming

## Table of Contents

- [Lotso Airdrop Server](#lotso-airdrop-server)
  - [Table of Contents](#table-of-contents)
  - [Introduction](#introduction)
  - [Features](#features)
  - [Requirements](#requirements)
  - [Diagram](#diagram)
    - [Setting the docker network](#setting-the-docker-network)
  - [Usage](#usage)
  - [REST API Endpoints](#rest-api-endpoints)
  - [Shell Scripts](#shell-scripts)
    - [Manually Trigger Airdrop](#manually-trigger-airdrop)
    - [Add Whitelist for Address](#add-whitelist-for-address)
  - [License](#license)

## Introduction

This is a server for the Lotso Airdrop project. It is a simple server that listens for incoming requests and responds with the appropriate data. The server is written in Go, runs in Docker containers, and uses a MySQL database to store data, designed to be scalable and fault-tolerant, with the ability to handle a large number of requests simultaneously.

The server runs a scheduled task, which will obtain data from the MySQL database and distribute airdrops at odd hours.

## Features

- [] Need to be updated

## Requirements

- [Go 1.21](https://golang.org/)
- [Docker](https://www.docker.com/)
- API Key from [Pocket Network](https://www.pokt.network/) for the production environment
- API Key from [Alchemy](https://www.alchemy.com/) for the test-net environment

## Diagram

- [] Need to be updated

### Setting the docker network

We need to create a docker network to allow the containers to communicate with each other. Run the following command:
```bash
docker network create lotso-network
```

## Usage

(Base-Mainnet) Start the docker container for the production environment:

```bash
IMG_NAME=bti/lotso-airdrop-server:0.0.1 \
SERVER_PORT=8080 \
MYSQL_HOST=host.docker.internal \
MYSQL_PORT=3306 \
MYSQL_USER=root \
MYSQL_PASSWD=your_mysql_database_password \
MYSQL_DB=lotso_base_main \
API_URL=your_api_url \
PRIVATE_KEY=your_private_key \
CHAIN_ID=8453 CUTOFF_BLOCK=0xbede7c \
CONTRACT_ADDRESS=your_airdrop_contract_address \
DEBUG=False \
TZ=UTC \
API_MODE=1 \
DECIMALS=18 \
PAIR_ADDRESS=your_pair_address \
BUYER_REWARD_LIMIT=10000000 \
NOT_BUYER_REWARD_LIMIT=2000000 \
docker-compose -f docker-compose.yaml up -d
```

(Base-Mainnet) Stop the docker container for the production environment:

```bash
IMG_NAME=bti/lotso-airdrop-server:0.0.1 \
SERVER_PORT=8080 \
MYSQL_HOST=host.docker.internal \
MYSQL_PORT=3306 \
MYSQL_USER=root \
MYSQL_PASSWD=your_mysql_database_password \
MYSQL_DB=lotso_base_main \
API_URL=your_api_url \
PRIVATE_KEY=your_private_key \
CHAIN_ID=8453 CUTOFF_BLOCK=0xbede7c \
CONTRACT_ADDRESS=your_airdrop_contract_address \
DEBUG=False \
TZ=UTC \
API_MODE=1 \
DECIMALS=18 \
PAIR_ADDRESS=your_pair_address \
BUYER_REWARD_LIMIT=10000000 \
NOT_BUYER_REWARD_LIMIT=2000000 \
docker-compose -f docker-compose.yaml down
```

(ETH-Sepolia) Start the docker container for the test-net environment:

```bash
IMG_NAME=bti/lotso-airdrop-server:0.0.1 \
SERVER_PORT=8081 \
MYSQL_HOST=host.docker.internal \
MYSQL_PORT=3306 \
MYSQL_USER=root \
MYSQL_PASSWD=your_mysql_database_password \
MYSQL_DB=lotso_eth_sepolia \
API_URL=your_api_url \
PRIVATE_KEY=your_private_key \
CHAIN_ID=11155111 \
CUTOFF_BLOCK=0x556df8 \
CONTRACT_ADDRESS=your_airdrop_contract_address \
DEBUG=True \
TZ=UTC \
API_MODE=1 \
DECIMALS=18 \
PAIR_ADDRESS=your_pair_address \
BUYER_REWARD_LIMIT=10000000 \
NOT_BUYER_REWARD_LIMIT=2000000 \
docker-compose -f docker-compose.yaml up -d
```

(ETH-Sepolia) Stop the docker container for the test-net environment:

```bash
IMG_NAME=bti/lotso-airdrop-server:0.0.1 \
SERVER_PORT=8081 \
MYSQL_HOST=host.docker.internal \
MYSQL_PORT=3306 \
MYSQL_USER=root \
MYSQL_PASSWD=your_mysql_database_password \
MYSQL_DB=lotso_eth_sepolia \
API_URL=your_api_url \
PRIVATE_KEY=your_private_key \
CHAIN_ID=11155111 \
CUTOFF_BLOCK=0x556df8 \
CONTRACT_ADDRESS=your_airdrop_contract_address \
DEBUG=True \
TZ=UTC \
API_MODE=1 \
DECIMALS=18 \
PAIR_ADDRESS=your_pair_address \
BUYER_REWARD_LIMIT=10000000 \
NOT_BUYER_REWARD_LIMIT=2000000 \
docker-compose -f docker-compose.yaml down
```

## REST API Endpoints

1. **GET /v1/info/apply_airdrop**	*Only enabled when ApiMode=0*

   - **Description:** Retrieves the transaction count for a given address. If the transaction count meets the specified conditions, the airdrop will be issued to this address as planned.
   - **Parameters:**
     
     - `address` (String): The blockchain address to query.
   - **Response:**
     
     - `address`: Wallet address of the current user, should be the same as the input address.
     - `transaction_count` (Integer): The number of transactions associated with the address.
     - `airdrop_count` (Integer): The airdrop count for the address.
     - `has_airdropped_count` (Boolean): The number of airdrops that have been issued to this address.
     - `scheduled_delivery` (String): The next available time for airdropping.
   - **Example:**
     
     Request:
     
     ```bash
     curl --location 'http://127.0.0.1:1423/v1/info/apply_airdrop?address=0x0000000000000000000000000000000000000001'
     ```
     
     Response:
     
     ```json
     {
       "code": 0,
       "message": "Success",
       "error": "",
       "data": {
          "address": "0x0000000000000000000000000000000000000001",
          "transaction_count": 11,
          "airdrop_count": 100000,
          "has_airdropped_count": 0,
          "scheduled_delivery": "2024-04-01T08:00:00Z"
       }
     }
     ```

2. **POST /v1/info/set_airdrop**	*Only enabled when ApiMode=1*

   - **Description:** Issue a specified number of airdrops to a specified address. Each address can only be specified once.
   - **Parameters:**
     
     - `address` (String): The blockchain address that airdropped to.
     - `amount` (Integer): The amount of airdrops to be distributed, should <= 100000000.
   - **Response:**
     
     - `address`: Wallet address of the current user, should be the same as the input address.
     - `transaction_count` (Integer): The number of transactions associated with the address.
     - `airdrop_count` (Integer): The airdrop count for the address.
     - `has_airdropped_count` (Boolean): The number of airdrops that have been issued to this address.
     - `scheduled_delivery` (String): The next available time for airdropping.
   - **Example:**
     
     Request:
     
     ```bash
     curl --location 'http://127.0.0.1:1423/v1/info/set_airdrop' \
     --header 'Content-Type: application/json' \
     --data '{
         "address":"0x0000000000000000000000000000000000000001",
         "amount": 10000
     }'
     ```
     
     Response:
     
     ```json
     {
       "code": 0,
       "message": "Success",
       "error": "",
       "data": {
          "address": "0x0000000000000000000000000000000000000001",
          "transaction_count": 0,
          "airdrop_count": 10000,
          "has_airdropped_count": 0,
          "scheduled_delivery": "2024-04-01T08:00:00Z"
       }
     }
     ```

3. **POST /v1/info/append_airdrop**	*Only enabled when ApiMode=1*

   - **Description:** Additional airdrops of the specified number will be issued to the specified address. The address to be appended to the airdrop must first be `set_airdrop` called. Each address can be appended multiple times.
   - **Parameters:**
   
     - `address` (String): The blockchain address that will be appended airdrop to.
     - `amount` (Integer): The amount of airdrops to be appended, should <= 100000000.
   - **Response:**

     - `address`: Wallet address of the current user, should be the same as the input address.
     - `transaction_count` (Integer): The number of transactions associated with the address.
     - `airdrop_count` (Integer): The airdrop count for the address.
     - `has_airdropped_count` (Boolean): The number of airdrops that have been issued to this address.
     - `scheduled_delivery` (String): The next available time for airdropping.
   - **Example:**

     Request:

     ```bash
     curl --location 'http://127.0.0.1:1423/v1/info/append_airdrop' \
     --header 'Content-Type: application/json' \
     --data '{
         "address":"0x0000000000000000000000000000000000000001",
         "amount": 100
     }'
     ```

     Response:

     ```json
     {
       "code": 0,
       "message": "Success",
       "error": "",
       "data": {
          "address": "0x0000000000000000000000000000000000000001",
          "transaction_count": 0,
          "airdrop_count": 10100,
          "has_airdropped_count": 10000,
          "scheduled_delivery": "2024-04-01T08:00:00Z"
       }
     }
     ```

4. **GET /v1/info/check_eligibility**	*Only enabled when ApiMode=1*

   - **Description:** Retrieves whether the specified address purchased the first Lotso Token before the attack.

   - **Parameters:**

     - `address` (String): The blockchain address that will be checked.

   - **Response:**
     
     - `data` (Boolean): Whether the specified address purchased the first Lotso Token before the attack.
     
   - **Example:**
     
     Request:
     
     ```bash
     curl --location 'http://127.0.0.1:1423/v1/info/check_eligibility?address=0x0000000000000000000000000000000000000001'
     ```
     
     Response:
     ```json
     {
       "code": 0,
       "message": "Success",
       "error": "",
       "data": true
     }
     ```

5. **GET /v1/info/recipients_count**	*Only enabled when Debug=true*

   - **Description:** Retrieves the recipient count for the airdrop.

   - **Response:**
     - `data` (Integer): The number of recipients eligible for the airdrop.

   - **Example:**
     
     Request:
     
     ```bash
     curl --location 'http://127.0.0.1:1423/v1/info/recipients_count'
     ```
     
     Response:
     ```json
     {
       "code": 0,
       "message": "Success",
       "error": "",
       "data": 1
     }
     ```


5. **GET /v1/info/addresses_should_airdrop**	*Only enabled when Debug=true*

   - **Description:** Retrieve the address and quantity of the next airdrop. The amount of the next airdrop is `airdrop_count` minus `has_airdropped_count`.

   - **Response:**
     - `data` (List): The number of recipients who have manually claimed the airdrop.

   - **Example:**
     
     Request:
     
     ```bash
     curl --location 'http://127.0.0.1:1423/v1/info/recipients_count'
     ```
     
     Response:
     ```json
     {
       "code": 0,
       "message": "Success",
       "error": "",
       "data": 1
     }
     ```

6. **GET /v1/info/distribute_airdrops**	*Only enabled when Debug=true*

   - **Description:** Distribute airdrops.

   - **Response:**
     - `data` (List): The transaction hashes of distributing airdrops.

   - **Example:**
     
     Request:
     
     ```bash
     curl --location --request POST 'http://127.0.0.1:1423/v1/info/distribute_airdrops'
     ```
     
     Response:
     ```json
     {
       "code": 0,
       "message": "Success",
       "error": "",
       "data": [
         "0x58fad0f6535a3beda4dd8c1c39d1fe350c7d9518e55999e5ab11f40c4129fdb1"
       ]
     }
     ```

7. **GET /v1/info/distribute_airdrops_to**	*Only enabled when Debug=true

   - **Description:** Distribute a specified number of airdrops to a specified address.
   - **Parameters:**
        - `address` (String): The blockchain address which will be distributed airdrop to
        - `amount` (Integer): The amount of airdrops to be distributed.
   - **Response:**
     - `data` (String): The transaction hash of distributing airdrops.
   - **Example:**
     
        Request:
        
        ```bash
        curl --location 'http://127.0.0.1:1423/v1/info/distribute_airdrops_to' \
        --header 'Content-Type: application/json' \
        --data '{
            "address":"0x0000000000000000000000000000000000000001",
            "amount": 1000000
        }'
        ```
        
        Response:
        ```json
        {
          "code": 0,
          "message": "Success",
          "error": "",
          "data": "0x58fad0f6535a3beda4dd8c1c39d1fe350c7d9518e55999e5ab11f40c4129fdb1"
        }
        ```

8. **GET /v1/info/claim_airdrop**	*Only enabled when Debug=true

   - **Description:** Receive airdrops as a designated account
   - **Parameters:**
       - `private_key` (String): The blockchain address which will be distributed airdrop to.

   - **Response:**
    - `data` (String): The transaction hash of distributing airdrops.

   - **Example:**
    
       Request:
       
       ```bash
       curl --location 'http://127.0.0.1:1423/v1/info/claim_airdrop' \
       --header 'Content-Type: application/json' \
       --data '{
           "private_key":"0x0000000000000000000000000000000000000000000000000000000000000001"
       }'
       ```
       
       Response:
       ```json
       {
         "code": 0,
         "message": "Success",
         "error": "",
         "data": "0x58fad0f6535a3beda4dd8c1c39d1fe350c7d9518e55999e5ab11f40c4129fdb1"
       }
       ```

## Shell Scripts

For testing purposes, we provide two scripts: `manually_trigger_airdrop.sh` and `add_whitelist_for_address.sh`. The first script updates specific records in a MySQL database based on user input, while the second script allows for the insertion or deletion of records in a specific table. Both scripts are designed to be run directly from the terminal.

### Manually Trigger Airdrop

This script is used to interact with a MySQL database to update records in a specific table. The script prompts the user for an address and a time interval. It then updates the record with the specified address, setting the `transaction_count`, `has_airdropped_count`, and `scheduled_delivery` fields.

Run the script in a terminal:

```bash
./manually_trigger_airdrop.sh
```

You will be prompted to enter the MySQL root password if it's not set as an environment variable. Then, you will be asked to enter the address to update and the time interval for scheduled delivery.

### Add Whitelist for Address

This script is used to interact with a MySQL database to either insert or delete records in a specific table. The script prompts the user for an address and an action (either "update" or "delete"). If the user enters "update", the script inserts a new record into the table with the specified address. If the user enters "delete", the script deletes the record with the specified address from the table.

Run the script in a terminal:

```bash
./add_whitelist_for_address.sh
```

You will be prompted to enter the MySQL root password if it's not set as an environment variable. Then, you will be asked to enter the address to update and the action to perform (either "update" or "delete").

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details.