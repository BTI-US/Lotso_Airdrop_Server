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

# Define database and table variables
DATABASE="lotso2_eth_sepolia"
TABLE="airdrop_items"
USE_HOST=true

# Execute commands inside MySQL container or host based on USE_HOST variable
if [ "$USE_HOST" = true ]; then
    mysql -u root -p$MYSQL_ROOT_PASSWORD <<EOF
    USE $DATABASE;
    show tables;
    DELETE FROM \`$DATABASE\`.\`$TABLE\` WHERE address='$ADDRESS';
    SELECT ROW_COUNT();
    COMMIT;
EOF
else
    docker exec -i mysql_sepolia mysql -u root -p$MYSQL_ROOT_PASSWORD <<EOF
    USE $DATABASE;
    show tables;
    DELETE FROM \`$DATABASE\`.\`$TABLE\` WHERE address='$ADDRESS';
    SELECT ROW_COUNT();
    COMMIT;
EOF
fi

echo "Commands executed successfully."
