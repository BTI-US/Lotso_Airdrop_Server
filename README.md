# Lotso Airdrop Server

[![Go](https://github.com/BTI-US/Lotso_Airdrop_Server/actions/workflows/go.yml/badge.svg)](https://github.com/BTI-US/Lotso_Airdrop_Server/actions/workflows/go.yml)
[![Docker CI](https://github.com/BTI-US/Lotso_Airdrop_Server/actions/workflows/docker-image.yml/badge.svg)](https://github.com/BTI-US/Lotso_Airdrop_Server/actions/workflows/docker-image.yml)
[![golangci-lint](https://github.com/BTI-US/Lotso_Airdrop_Server/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/BTI-US/Lotso_Airdrop_Server/actions/workflows/golangci-lint.yml)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

- Last Modified: 2024-05-06
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
    - [Required Environment Variables](#required-environment-variables)
    - [Bash command for docker-compose](#bash-command-for-docker-compose)
  - [REST API Endpoints](#rest-api-endpoints)
    - [Return Code](#return-code)
  - [Shell Scripts](#shell-scripts)
    - [Delete Airdrop Record](#delete-airdrop-record)
    - [Add Whitelist for Address](#add-whitelist-for-address)
    - [Change Distribution Time](#change-distribution-time)
  - [License](#license)

## Introduction

This is a server for the Lotso Airdrop project. It is a simple server that listens for incoming requests and responds with the appropriate data. The server is written in Go, runs in Docker containers, and uses a MySQL database to store data, designed to be scalable and fault-tolerant, with the ability to handle a large number of requests simultaneously.

The server runs a scheduled task, which will obtain data from the MySQL database and distribute airdrops at odd hours.

## Features

- [] Need to be updated

## Requirements

- [Go 1.23](https://golang.org/)
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

### Required Environment Variables

| Env Name               | Type     | Default     | Must be declared | Description                                                  |
| ---------------------- | -------- | ----------- | ---------------- | ------------------------------------------------------------ |
| SERVER_PORT            | String   | "1423"      |                  | Server port                                                  |
| SSL_CERT               | String   | ""          |                  | SSL cert path                                                |
| SSL_KEY                | String   | ""          |                  | SSL key path                                                 |
| DEBUG                  | Bool     | false       |                  | Enable debug API                                             |
| API_URL                | String   | ""          | Yes              | Api url which is used to connect to the RPC interface of the chain |
| PRIVATE_KEY            | String   | ""          | Yes              | Account private key used to issue airdrops regularly, hexadecimal string type |
| CHAIN_ID               | Int64    | 0           | Yes              | The chain ID which server will connect to                    |
| CUTOFF_BLOCK           | String   | "latest"    |                  | When counting the number of transactions, the number of transactions before this block will be counted |
| TOKEN_ADDRESS          | String   | ""          | Yes              | The airdrop contract address, hexadecimal string type        |
| MYSQL_HOST             | String   | "127.0.0.1" |                  | Mysql host                                                   |
| MYSQL_PORT             | Int      | 3306        |                  | Mysql port                                                   |
| MYSQL_USER             | String   | "root"      |                  | Mysql user                                                   |
| MYSQL_PASSWD           | String   | "123456"    |                  | Mysql password                                               |
| MYSQL_DB               | String   | "lotso"     |                  | Mysql database                                               |
| PAIR_ADDRESS           | String   | ""          | Yes              | The uniswapPair address used to count buyers                 |
| DECIMALS               | Uint64   | 18          |                  | Decimals of token                                            |
| BUYER_REWARD_LIMIT     | Uint64   | 10000000    |                  | Buyer reward limit                                           |
| NOT_BUYER_REWARD_LIMIT | Uint64   | 2000000     |                  | NotBuyer reward limit                                        |
| TRUSTED_PROXIES        | []String | []          |                  | Trusted proxies                                              |

### Bash command for docker-compose

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

5. **GET /v1/info/recipient_info**

   - **Description:** Retrieves the number of times the airdrop contract has been called to claim airdrops and the total amount of airdrops claimed.

   - **Response:**
     
     - `recipient_count` (Integer): The number of times the airdrop contract has been called to receive airdrops.
     - `airdrop_amount` (Integer): The total amount of airdrops has been claimed.
     
   - **Example:**
     
     Request:
     
     ```bash
     curl --location 'http://127.0.0.1:1423/v1/info/recipient_info'
     ```
     
     Response:
     ```json
     {
       "code": 0,
       "message": "Success",
       "error": "",
       "data": {
          "recipient_count": 0,
          "airdrop_amount": 0
        }
     }
     ```


6. **GET /v1/info/addresses_should_airdrop**	*Only enabled when Debug=true*

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

7. **POST /v1/info/distribute_airdrops**	*Only enabled when Debug=true*

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

8. **POST /v1/info/distribute_airdrops_to**	*Only enabled when Debug=true

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

9. **POST /v1/info/claim_airdrop**	*Only enabled when Debug=true

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

### Return Code

Return code includes two parts: code and message, which are used to indicate the information returned by the HTTP API.

All possible return codes are as follows:

| Name                      | Code  | Message                      |
| ------------------------- | ----- | ---------------------------- |
| SUCCESS                   | 0     | Success                      |
| UnknownError              | 10000 | Unknown error                |
| WrongParams               | 10001 | Wrong params                 |
| DialApiServerFailed       | 10002 | Dial API server failed       |
| InvalidAddress            | 10003 | Invalid address              |
| EthApiCallFailed          | 10004 | ETH API call failed          |
| ConvertFailed             | 10005 | Convert failed               |
| AirdropItemAlreadyExists  | 10006 | You have applied the airdrop |
| AirdropItemNotExists      | 10007 | Airdrop item not exists      |
| GetAirdropItemFailed      | 10008 | Get airdrop item failed      |
| SaveAirdropItemFailed     | 10009 | Save airdrop item failed     |
| DistributeAirdropsFailed  | 10010 | Distribute airdrops failed   |
| ClaimAirdropsFailed       | 10011 | Claim airdrops failed        |
| RecipientsCountFailed     | 10012 | Get recipients count failed  |
| GetTransactionTopicFailed | 10013 | Get transaction topic failed |
| RewardLimitReached        | 10014 | Reward limit reached         |

## Shell Scripts

For testing purposes, we provide two scripts: `delete_airdrop_record.sh` and `add_whitelist_for_address.sh`. The first script updates specific records in a MySQL database based on user input, while the second script allows for the insertion or deletion of records in a specific table. Both scripts are designed to be run directly from the terminal.

### Delete Airdrop Record

This script is used to interact with a MySQL database to update records in a specific table. The script prompts the user for an address. It then deletes the record with the specified address from the table.

Run the script in a terminal:

```bash
./delete_airdrop_record.sh
```

You will be prompted to enter the MySQL root password if it's not set as an environment variable. Then, you will be asked to enter the address to update and the time interval for scheduled delivery.

### Add Whitelist for Address

This script is used to interact with a MySQL database to either insert or delete records in a specific table. The script prompts the user for an address and an action (either "update" or "delete"). If the user enters "update", the script inserts a new record into the table with the specified address. If the user enters "delete", the script deletes the record with the specified address from the table.

Run the script in a terminal:

```bash
./add_whitelist_for_address.sh
```

You will be prompted to enter the MySQL root password if it's not set as an environment variable. Then, you will be asked to enter the address to update and the action to perform (either "update" or "delete").

### Change Distribution Time

This script is used to interact with a MySQL database to update records in a specific table. The script prompts the user for an address and a new distribution time interval based on the current UTC. It then updates the record with the specified address in the table with the new distribution time.

Run the script in a terminal:

```bash
./change_distribute_time.sh
```

You will be prompted to enter the MySQL root password if it's not set as an environment variable. Then, you will be asked to enter the address to update and the new distribution time interval in minutes for scheduled delivery (positive for the future, negative for the past).

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details.