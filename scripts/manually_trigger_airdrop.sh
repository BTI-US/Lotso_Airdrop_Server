#!/bin/bash

# Prompt for MySQL password if not set as an environment variable
if [ -z "$MYSQL_ROOT_PASSWORD" ]; then
    read -sp "Enter MySQL root password: " MYSQL_ROOT_PASSWORD
    echo
fi

# Prompt for the address
read -p "Enter the address to update: " ADDRESS

# Trim whitespace and ensure address is in correct case
ADDRESS=$(echo "$ADDRESS" | xargs)
ADDRESS=${ADDRESS// /}
if [[ $ADDRESS == 0x* ]]; then
    ADDRESS=${ADDRESS:2}
fi

# Check if the address is entered
if [ -z "$ADDRESS" ]; then
    echo "No address entered. Exiting."
    exit 1
fi

# Get current time in UTC and prompt for time interval
CURRENT_TIME=$(date -u +"%Y-%m-%d %H:%M:%S.%3N")
echo "Current UTC time is: $CURRENT_TIME"

read -p "Enter time interval in minutes for scheduled delivery (positive for future, negative for past): " INTERVAL

# Check if the interval is entered and is a number
if ! [[ "$INTERVAL" =~ ^[0-9]+$ ]]; then
    echo "Invalid time interval. Exiting."
    exit 1
fi

# Calculate scheduled delivery time based on positive or negative interval
if [[ "$INTERVAL" -lt 0 ]]; then
    # Past time
    SCHEDULED_DELIVERY=$(date -u -d "$CURRENT_TIME $INTERVAL minutes" +"%Y-%m-%d %H:%M:%S.%3N")
else
    # Future time
    SCHEDULED_DELIVERY=$(date -u -d "$CURRENT_TIME + $INTERVAL minutes" +"%Y-%m-%d %H:%M:%S.%3N")
fi

# Define database and table variables
DATABASE="lotso2_eth_sepolia"
TABLE="airdrop_items"
TRANSACTION_COUNT=11
HAS_AIRDROPPED_COUNT=0
USE_HOST=true

# Execute commands inside MySQL container or host based on USE_HOST variable
if [ "$USE_HOST" = true ]; then
    mysql -u root -p$MYSQL_ROOT_PASSWORD <<EOF
    USE $DATABASE;
    show tables;
    UPDATE $TABLE SET transaction_count=$TRANSACTION_COUNT, has_airdropped_count=$HAS_AIRDROPPED_COUNT, scheduled_delivery='$SCHEDULED_DELIVERY' WHERE BINARY address='$ADDRESS';
    SELECT ROW_COUNT();
    COMMIT;
EOF
else
    docker exec -i mysql_sepolia mysql -u root -p$MYSQL_ROOT_PASSWORD <<EOF
    USE $DATABASE;
    show tables;
    UPDATE $TABLE SET transaction_count=$TRANSACTION_COUNT, has_airdropped_count=$HAS_AIRDROPPED_COUNT, scheduled_delivery='$SCHEDULED_DELIVERY' WHERE BINARY address='$ADDRESS';
    SELECT ROW_COUNT();
    COMMIT;
EOF
fi

echo "Commands executed successfully."
