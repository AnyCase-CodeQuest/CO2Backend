package message

import (
	"context"
	"github.com/AnyCase-Company-LTD/CO2Backend/src/values"
	"github.com/kubemq-io/kubemq-go"
	"log"
	"os"
	"strconv"
)

func SendMessageToQueue(message []byte) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	port, _ := strconv.Atoi(os.Getenv(values.EnvQueuePort))
	client, err := kubemq.NewClient(ctx,
		kubemq.WithAddress(os.Getenv(values.EnvQueueHost), port),
		kubemq.WithClientId(values.ClientId))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	sendResult, err := client.NewQueueMessage().
		SetChannel(os.Getenv(values.EnvQueueName)).
		SetBody(message).
		// message will expire within 20 seconds if will not consumed
		SetPolicyExpirationSeconds(values.QueueTTL).
		Send(ctx)
	if err != nil || sendResult.IsError {
		log.Println(sendResult.Error)
		log.Fatal(err)
	}
}
