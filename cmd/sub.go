package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/cobra"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

// subCmd represents the sub command
var subCmd = &cobra.Command{
	Use: "sub",
	Run: func(cmd *cobra.Command, args []string) {
		mqtt.ERROR = log.New(os.Stdout, "", 0)
		opts := mqtt.NewClientOptions().AddBroker("tcp://broker.emqx.io:1883").SetClientID(args[0])

		c := mqtt.NewClient(opts)
		if token := c.Connect(); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}

		// Subscribe to a topic
		if token := c.Subscribe("zahid/hassan", 0, f); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
			os.Exit(1)
		}

		stop := make(chan os.Signal)
		signal.Notify(stop, os.Interrupt)
		<-stop

		// Unscribe
		if token := c.Unsubscribe("zahid/hassan"); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
			os.Exit(1)
		}

		// Disconnect
		c.Disconnect(250)
		time.Sleep(1 * time.Second)
	},
}
