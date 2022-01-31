#!/bin/sh
export SNAPSHOT_NAME="hoge-final-snapshot"

aws rds delete-db-cluster-snapshot --db-cluster-snapshot-identifier ${SNAPSHOT_NAME}