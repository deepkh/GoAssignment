#!/bin/bash

export DATABASE="RecommendationSystem"
# [Warning] Using a password on the command line interface can be insecure.
export MYSQL_BIN="mysql -h 127.0.0.1 -P 3306 -uroot -p1234 "
export SQLS_PATH=./

create_db() {
    echo "=== create_db"
    $MYSQL_BIN -e "CREATE DATABASE IF NOT EXISTS RecommendationSystem;"
}

create_tables() {
    echo "=== create_tables"
    $MYSQL_BIN -D $DATABASE < $SQLS_PATH/recommendation-system-tables.sql

    # add default ten thoundsands recommendation items
    for i in {1..1000}
    do
        $MYSQL_BIN -D $DATABASE < $SQLS_PATH/recommendation-system_add_default_recommendations.sql
    done
}

drop_db() {
    echo "=== drop_db"
    $MYSQL_BIN -e "drop database RecommendationSystem;"
}

all() {
    drop_db
    create_db
    create_tables
}

if [ "$#" -ne 1 ]; then
    all
else
    $@
fi

