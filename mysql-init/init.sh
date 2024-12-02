#!/bin/bash
# Replace the placeholders with environment variables
sed -i "s|\${BLUEPRINT_DB_DATABASE}|$BLUEPRINT_DB_DATABASE|g" /docker-entrypoint-initdb.d/init.template.sql
sed -i "s|\${BLUEPRINT_DB_USERNAME}|$BLUEPRINT_DB_USERNAME|g" /docker-entrypoint-initdb.d/init.template.sql
sed -i "s|\${BLUEPRINT_DB_PASSWORD}|$BLUEPRINT_DB_PASSWORD|g" /docker-entrypoint-initdb.d/init.template.sql

# Run the SQL script
mysql < /docker-entrypoint-initdb.d/init.template.sql
