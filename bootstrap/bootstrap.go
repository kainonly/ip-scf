package bootstrap

import (
	"context"
	"github.com/caarlos0/env/v6"
	"github.com/go-resty/resty/v2"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"ip2region-mongo/common"
)

var Provides = wire.NewSet(
	UseMongoDB,
	UseDatabase,
	UseHttpClient,
)

// SetValues 初始化配置
func SetValues() (values *common.Values, err error) {
	values = new(common.Values)
	if err = env.Parse(values); err != nil {
		return
	}
	return
}

// UseMongoDB 设置 MongoDB
func UseMongoDB(values *common.Values) (*mongo.Client, error) {
	return mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(values.Database.Uri),
	)
}

// UseDatabase 指定数据库
func UseDatabase(client *mongo.Client, values *common.Values) (db *mongo.Database) {
	return client.Database(values.Database.DbName)
}

// UseHttpClient 创建 HttpClient
func UseHttpClient() *resty.Client {
	return resty.New().
		SetBaseURL(`https://raw.githubusercontent.com/lionsoul2014/ip2region/master`)
}