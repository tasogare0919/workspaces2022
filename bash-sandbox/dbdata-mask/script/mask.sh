#/bin/sh
export DB_ENDPOINT=`aws rds describe-db-clusters --db-cluster-identifier hoge --query DBClusters[].Endpoint --output text`
export STG_PASSWORD=hoge

echo "Masking execution start"
mysql -h"${DB_ENDPOINT}" -u"stgdb-user" -p"${STG_PASSWORD}" << EOF
$( cat sql/all-mask.sql )
EOF

echo "end"