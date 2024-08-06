#!/bin/bash
set -e

CASSANDRA_CONTAINER=${CASSANDRA_CONTAINER:-cassandra_node}
CASSANDRA_USER=${CASSANDRA_USER:-admin}
CASSANDRA_PASSWORD=${CASSANDRA_PASSWORD:-admin}

echo "Waiting for Cassandra to be ready..."
until docker exec $CASSANDRA_CONTAINER cqlsh -u "admin" -p "cassandra" -e "describe keyspaces" > /dev/null 2>&1; do
    echo "Cassandra is unavailable - sleeping"
    sleep 2
done

echo "Cassandra is up - initializing keyspace"

docker exec -i $CASSANDRA_CONTAINER cqlsh -u "admin" -p "cassandra" << EOF
CREATE KEYSPACE IF NOT EXISTS job_scheduler
WITH replication = {
    'class': 'SimpleStrategy',
    'replication_factor': 1
};

USE job_scheduler;

CREATE TABLE IF NOT EXISTS migrations (
    version text PRIMARY KEY,
    applied_at timestamp
);
EOF

echo "Keyspace job_scheduler initialized"