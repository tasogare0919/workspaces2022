#!/bin/sh
export RDS_CLUSTER_NAME="hoge"

aws rds delete-db-instance \
    --db-instance-identifier "${RDS_CLUSTER_NAME}" \
    --skip-final-snapshot || true

aws rds wait db-instance-deleted \
    --db-instance-identifier "${RDS_CLUSTER_NAME}"

aws rds delete-db-cluster \
    --db-cluster-identifier "${RDS_CLUSTER_NAME}" \
    --skip-final-snapshot || true