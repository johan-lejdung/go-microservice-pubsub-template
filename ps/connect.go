package ps

import (
	"os"
	"time"

	"github.com/johan-lejdung/go-microservice-pubsub-template/goservice"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"cloud.google.com/go/pubsub"
	"golang.org/x/net/context"
)

// IPubSub interface for Pubsub
type IPubSub interface {
	Init()
	Consume()
	Produce(msg *goservice.Message) error
}

// PubSub implements IPubSub
type PubSub struct {
	GoService          goservice.Services `inject:""`
	client             *pubsub.Client
	pubsubTopic        *pubsub.Topic
	pubsubSubscription *pubsub.Subscription
}

// Init connects to PubSub and populates the topics with names
func (ps *PubSub) Init() {
	log.Debugf("Initializing PubSub")

	ctx := context.Background()

	c, err := pubsub.NewClient(ctx, os.Getenv("GOOGLE_CLOUD_PROJECT"))
	if err != nil {
		log.Fatal(err)
	}
	// Set the client
	ps.client = c

	log.Debugf("PubSub client PRODUCER created")
	log.Debugf("Filling topics and subscriptions")

	ps.pubsubTopic = ps.createTopicIfNotExists(os.Getenv("PUBSUB_TOPIC"))
	ps.pubsubSubscription = ps.createSubscriptionIfNotExists(os.Getenv("PUBSUB_SUB"), ps.pubsubTopic)

	log.Debugf("PubSub ready for use")
}

func (ps *PubSub) createSubscriptionIfNotExists(subName string, topic *pubsub.Topic) *pubsub.Subscription {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()

	// Create a topic to subscribe to.
	s, err := ps.client.CreateSubscription(ctx, subName, pubsub.SubscriptionConfig{
		Topic:             topic,
		AckDeadline:       5 * time.Minute,
		RetentionDuration: time.Hour * 24 * 2,
	})
	if err != nil && grpc.Code(err) != codes.AlreadyExists {
		log.Warnf("Failed to create the subscription. %v", err)
		return nil
	} else if grpc.Code(err) == codes.AlreadyExists {
		return ps.client.Subscription(subName)
	}
	return s
}

func (ps *PubSub) createTopicIfNotExists(topicName string) *pubsub.Topic {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()

	log.Debugf("Creating pubsub topic %s", topicName)
	// Create a topic to subscribe to.
	t, err := ps.client.CreateTopic(ctx, topicName)
	if err != nil && grpc.Code(err) != codes.AlreadyExists {
		log.Fatalf("Failed to create the pubsub topic: %v", err)
		return nil
	} else if grpc.Code(err) == codes.AlreadyExists {
		return ps.client.Topic(topicName)
	}
	return t
}
