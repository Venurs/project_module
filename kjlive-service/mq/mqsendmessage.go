package mq
import (
    "encoding/json"
    "kjlive-service/logs"
)

func SendMessageByManager(exchange string, routingKey string, task string, args []interface{}) {
    var mq1 *BaseMq
    mq1 = GetConnection("manager")
    channelContext := ChannelContext{Exchange: exchange,
        ExchangeType: "topic",
        RoutingKey:   routingKey,
        Reliable:     true,
        Durable:      true}
    body := []interface{}{args, map[string]string{}, nil}
    byteBody, err := json.Marshal(body)
    if err != nil {
        logs.Error(err)
    } else {
        mq1.Publish(&channelContext, byteBody, task)
    }
}
