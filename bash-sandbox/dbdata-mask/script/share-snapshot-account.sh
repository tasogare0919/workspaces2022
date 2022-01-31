# /bin/bash
export SNAPSHOT_NAME="hoge-final-snapshot"
export SHARE_ACCOUNT_ID="12345678910"


aws rds modify-db-cluster-snapshot-attribute \
    --db-cluster-snapshot-identifier ${SNAPSHOT_NAME} \
    --attribute-name restore \
    --values-to-add ${SHARE_ACCOUNT_ID}