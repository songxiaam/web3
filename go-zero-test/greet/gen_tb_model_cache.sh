#根据数据库表创建model

# 表名
tables=$2
# 目录
modeldir=./internal/model
# 数据库配置
host=127.0.0.1
port=3307
dbname=$1
username=root
password=12345678

echo "开始创建 $dbname 的表 $tables"
goctl model mysql datasource -url="${username}:${password}@tcp(${host}:${port})/${dbname}" -table="${tables}" -dir="${modeldir}" --style=goZero --cache true