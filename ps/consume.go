package ps

import (
	"cloud.google.com/go/pubsub"
	"github.com/golang/protobuf/proto"
	"github.com/johan-lejdung/go-microservice-pubsub-template/protomsg"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

//Consume will start listening for pubsub messages on the topic
func (ps *PubSub) Consume() {
	ctx := context.Background()

	log.Debugf("PubSub sub %s starting", ps.pubsubTopic)
	err := ps.pubsubSubscription.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		msg := protomsg.Message{}
		if err := proto.Unmarshal(m.Data, &msg); err != nil {
			m.Nack()
		} else if ps.GoService.TestFunction() != nil {
			m.Ack() // Acknowledge that we've consumed the message.
		}
	})
	if err != nil {
		log.Fatalf("PubSub %s FAILED to start: %v", ps.pubsubSubscription, err)
	}
}
