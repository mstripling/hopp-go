#!/bin/bash

# Load environment variables
if [ -f .env ]; then
  export $(grep -v '^#' .env | xargs)
fi

envsubst < mysql-init/init.template.sql > mysql-init/init.sql
