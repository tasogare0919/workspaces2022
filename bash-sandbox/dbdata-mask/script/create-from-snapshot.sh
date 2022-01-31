# /bin/bash
export DATE=`date +"%Y-%m-%d"`
export RDS_CLUSTER_NAME="hoge"
export SNAPSHOT_NAME="hoge-snapshot-${DATE}"
export SUBNET_GROUP_NAME="hoge-dbsbunet"
export SECURITY_GROUP_ID="sg-xxx"
export KMS_KEY_ID="xxx"
export DB_INSTANCE_CLASS="db.t3.small"

aws rds restore-db-cluster-from-snapshot \
    --snapshot-identifier ${SNAPSHOT_NAME} \
    --db-cluster-identifier "${RDS_CLUSTER_NAME}" \
    --engine "aurora" \
    --engine-version "5.6.mysql_aurora.1.22.2" \
    --db-subnet-group-name "${SUBNET_GROUP_NAME}" \
    --vpc-security-group-ids "${SECURITY_GROUP_ID}" \
    --db-cluster-parameter-group-name "hoge-parameter" \
    --kms-key-id "${KMS_KEY_ID}" \
    --no-deletion-protection
aws rds create-db-instance  \
    --db-instance-identifier "${RDS_CLUSTER_NAME}" \
    --db-cluster-identifier "${RDS_CLUSTER_NAME}" \
    --db-instance-class "${DB_INSTANCE_CLASS}" \
    --engine "aurora" \
    --no-publicly-accessible
aws rds wait db-instance-available \
    --db-instance-identifier ${RDS_CLUSTER_NAME}