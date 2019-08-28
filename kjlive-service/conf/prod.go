package conf

type configProd struct {
    GOModel               string `cfgDefault:"prod"`
    DatabaseSource        string `cfgDefault:"sg:sg123456@tcp(地址.mysql.rds.aliyuncs.com)/sellergrowth"` // read
    DatabaseSourceReplica string `cfgDefault:"sg:sg123456@tcp(地址.mysql.rds.aliyuncs.com)/sellergrowth"` // write
    RedisAddr             string `cfgDefault:"地址.redis.rds.aliyuncs.com:6379"`
    RedisPassword         string `cfgDefault:""`
    PolyvLiveAppId        string `cfgDefault:""`
    PolyvLiveAppSecret    string `cfgDefault:""`
    PolyvLiveUserId       string `cfgDefault:""`
    RABBITMQHOST          string `cfgDefault:"amqp://guest:guest@:5672/"`
    UnicornServiceRpcUrl  string `cfgDefault:"http:///rpc/"`
}
