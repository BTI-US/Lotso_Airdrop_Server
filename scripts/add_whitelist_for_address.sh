#!/bin/bash

# Prompt for MySQL password if not set as an environment variable
if [ -z "$MYSQL_ROOT_PASSWORD" ]; then
    read -sp "Enter MySQL root password: " MYSQL_ROOT_PASSWORD
    echo
fi

# Prompt for the address
read -p "Enter the address to update: " ADDRESS

# Trim whitespace and ensure address is in correct case
TO_ADDRESS=$(echo "$ADDRESS" | xargs)
TO_ADDRESS=${TO_ADDRESS// /}
if [[ $TO_ADDRESS == 0x* ]]; then
    TO_ADDRESS=${TO_ADDRESS:2}
fi

# Check if the address is entered
if [ -z "$TO_ADDRESS" ]; then
    echo "No address entered. Exiting."
    exit 1
fi

# Prompt for the action
read -p "Enter the action (insert/delete): " ACTION

# Define database and table variables
DATABASE="lotso2_eth_sepolia"
TABLE="transfer_topics"
BLOCK_HASH="182e7a9eabd0a8754e101507972b8e0b3500bb409756950691c5cd8b4ada4b43"
BLOCK_NUMBER=12800534
TRANSACTION_HASH="23ad64db3f4cdf8d306d33a50ea7b0ab2b8254d070e256e381a12f4ee997aadd"
FROM_ADDRESS="c53d6ab5c23faca7f545fa20a523c87c16edf3e2"
DATA="0000000000000000000000000000000000000000005e5051263506fefe8f7602"
USE_HOST=true

# Execute commands inside MySQL container or host based on USE_HOST variable
if [ "$USE_HOST" = true ]; then
    if [ "$ACTION" = "insert" ]; then
        mysql -u root -p$MYSQL_ROOT_PASSWORD <<EOF
        USE $DATABASE;
        INSERT INTO $TABLE (\`block_hash\`, \`block_number\`, \`transaction_hash\`, \`from\`, \`to\`, \`data\`) VALUES ('$BLOCK_HASH', $BLOCK_NUMBER, '$TRANSACTION_HASH', '$FROM_ADDRESS', '$TO_ADDRESS', '$DATA');
        SELECT ROW_COUNT();
        COMMIT;
EOF
    elif [ "$ACTION" = "delete" ]; then
        mysql -u root -p$MYSQL_ROOT_PASSWORD <<EOF
        USE $DATABASE;
        DELETE FROM $TABLE WHERE \`to\`='$TO_ADDRESS';
        SELECT ROW_COUNT();
        COMMIT;
EOF
    else
        echo "Invalid action. Exiting."
        exit 1
    fi
else
    if [ "$ACTION" = "insert" ]; then
        docker exec -i mysql_sepolia mysql -u root -p$MYSQL_ROOT_PASSWORD <<EOF
        USE $DATABASE;
        INSERT INTO $TABLE (\`block_hash\`, \`block_number\`, \`transaction_hash\`, \`from\`, \`to\`, \`data\`) VALUES ('$BLOCK_HASH', $BLOCK_NUMBER, '$TRANSACTION_HASH', '$FROM_ADDRESS', '$TO_ADDRESS', '$DATA');
        SELECT ROW_COUNT();
        COMMIT;
EOF
    elif [ "$ACTION" = "delete" ]; then
        docker exec -i mysql_sepolia mysql -u root -p$MYSQL_ROOT_PASSWORD <<EOF
        USE $DATABASE;
        DELETE FROM $TABLE WHERE \`to\`='$TO_ADDRESS';
        SELECT ROW_COUNT();
        COMMIT;
EOF
    else
        echo "Invalid action. Exiting."
        exit 1
    fi
fi

echo "Commands executed successfully."
