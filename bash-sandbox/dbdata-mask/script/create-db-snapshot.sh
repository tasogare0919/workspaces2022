# /bin/bash
export RDS_CLUSTER_NAME="hoge"
export SNAPSHOT_NAME="hoge-final-snapshot"


aws rds create-db-cluster-snapshot \
    --db-cluster-identifier ${RDS_CLUSTER_NAME} \
    --db-cluster-snapshot-identifier ${SNAPSHOT_NAME}
 
aws rds wait db-cluster-snapshot-available --db-cluster-snapshot-identifier ${SNAPSHOT_NAME}