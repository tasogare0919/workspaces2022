#/bin/sh
export DB_ENDPOINT=`aws rds describe-db-clusters --db-cluster-identifier hoge --query DBClusters[].Endpoint --output text`
export PASSWORD=hogehoe

echo "Create user execution start"
mysql -h"${DB_ENDPOINT}" -u"root" -p"${PASSWORD}" << EOF
$( cat sql/create-db-user.sql )
EOF

echo "end"
