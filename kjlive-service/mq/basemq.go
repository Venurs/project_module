package mq

import (
    "crypto/md5"
    "encoding/hex"
    "errors"
    "strconv"
    "sync"
    "time"
    "os"
    
    "github.com/streadway/amqp"
    "github.com/snluu/uuid"

    "kjlive-service/logs"
)

type MqConnection struct {
    Lock       sync.RWMutex
    Connection *amqp.Connection
    MqUri      string
}

type ChannelContext struct {
    Exchange     string
    ExchangeType string
    RoutingKey   string
    Reliable     bool
    Durable      bool
    ChannelId    string
    Channel      *amqp.Channel
}

type BaseMq struct {
    MqConnection *MqConnection
    
    //channel cache
    ChannelContexts map[string]*ChannelContext
}

func (bmq *BaseMq) Init() {
    bmq.ChannelContexts = make(map[string]*ChannelContext)
}

// One would typically keep a channel of publishings, a sequence number, and a
// set of unacknowledged sequence numbers and loop until the publishing channel
// is closed.
func (bmq *BaseMq) confirmOne(confirms <-chan amqp.Confirmation) {
    logs.Info("waiting for confirmation of one publishing")
    
    if confirmed := <-confirms; confirmed.Ack {
        logs.Infof("confirmed delivery with delivery tag: %d", confirmed.DeliveryTag)
    } else {
        logs.Errorf("failed delivery of delivery tag: %d", confirmed.DeliveryTag)
    }
}

/*
func (bmq *BaseMq) getMqUri() string {
	return "amqp://" + bmq.MqConnection.User + ":" + bmq.MqConnection.PassWord + "@" + bmq.MqConnection.Host + ":" + bmq.MqConnection.Port + "/"
}
*/
/*
get md5 from channel context
*/
func (bmq *BaseMq) generateChannelId(channelContext *ChannelContext) string {
    stringTag := channelContext.Exchange + ":" + channelContext.ExchangeType + ":" + channelContext.RoutingKey + ":" +
        strconv.FormatBool(channelContext.Durable) + ":" + strconv.FormatBool(channelContext.Reliable)
    hasher := md5.New()
    hasher.Write([]byte(stringTag))
    return hex.EncodeToString(hasher.Sum(nil))
}

/*
1. use old connection to generate channel
2. update connection then channel
*/
func (bmq *BaseMq) refreshConnectionAndChannel(channelContext *ChannelContext) error {
    bmq.MqConnection.Lock.Lock()
    defer bmq.MqConnection.Lock.Unlock()
    var err error
    
    if bmq.MqConnection.Connection != nil {
        channelContext.Channel, err = bmq.MqConnection.Connection.Channel()
    } else {
        err = errors.New("connection nil")
    }
    
    // reconnect connection
    if err != nil {
        for {
            bmq.MqConnection.Connection, err = amqp.Dial(bmq.MqConnection.MqUri)
            if err != nil {
                time.Sleep(10 * time.Second)
            } else {
                channelContext.Channel, _ = bmq.MqConnection.Connection.Channel()
                break
                
            }
        }
    }
    
    if err = channelContext.Channel.ExchangeDeclare(
        channelContext.Exchange,     // name
        channelContext.ExchangeType, // type
        channelContext.Durable,      // durable
        false,                       // auto-deleted
        false,                       // internal
        false,                       // noWait
        nil,                         // arguments
    ); err != nil {
        return err
    }
    
    //add channel to channel cache
    bmq.ChannelContexts[channelContext.ChannelId] = channelContext
    return nil
}

/*
publish message
*/
func (bmq *BaseMq) Publish(channelContext *ChannelContext, body []byte, task string) error {
    
    channelContext.ChannelId = bmq.generateChannelId(channelContext)
    if bmq.ChannelContexts[channelContext.ChannelId] == nil {
        bmq.refreshConnectionAndChannel(channelContext)
    } else {
        channelContext = bmq.ChannelContexts[channelContext.ChannelId]
    }
    var sguuid = uuid.Rand()
    host, _ := os.Hostname()
    if err := channelContext.Channel.Publish(
        channelContext.Exchange,   // publish to an exchange
        channelContext.RoutingKey, // routing to 0 or more queues
        false,                     // mandatory
        false,                     // immediate
        amqp.Publishing{
            Headers:         amqp.Table{
                "id":sguuid.Hex(),
                "lang":"go",
                "task":task,
                "argsrepr":string(body[:]),
                "kwargsrepr":"{}",
                "origin": strconv.Itoa(os.Getpid()) + "@" + host,
                
            },
            ContentType:     "application/json",
            ContentEncoding: "utf-8",
            Body:            body,
            DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
            Priority:        0,              // 0-9
            // a bunch of application/implementation-specific fields
        },
    ); err != nil {
        time.Sleep(10 * time.Second)
        bmq.refreshConnectionAndChannel(channelContext)
    }
    return nil
}
