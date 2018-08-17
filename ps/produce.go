package ps

import (
	"errors"
	"fmt"

	"github.com/johan-lejdung/go-microservice-pubsub-template/protomsg"

	"cloud.google.com/go/pubsub"
	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

// Produce will place a message in the pubsub topic
func (ps *PubSub) Produce(msg *protomsg.Message) error {
	log.Debugf("About to place one message to pubsub on topic %s", ps.pubsubTopic.String())
	ctx := context.Background()

	byteArr, err := proto.Marshal(msg)
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not marshal message: %v", err), 500)
		return err
	}
	pubsubMsg := &pubsub.Message{
		Data: byteArr,
	}

	if ps.pubsubTopic == nil {
		log.Fatal(fmt.Sprintf("Failed to fetch topic when publishing message to pubsub: %v", err), 500)
		return errors.New("Failed to fetch topic")
	}

	if _, err := ps.pubsubTopic.Publish(ctx, pubsubMsg).Get(ctx); err != nil {
		log.Fatal(fmt.Sprintf("Could not publish message to pubsub: %v", err), 500)
		return err
	}
	log.Debugf("Placed one message to pubsub on topic %s", ps.pubsubTopic.String())
	return nil
}
