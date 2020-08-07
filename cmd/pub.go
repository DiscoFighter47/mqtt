package cmd

import (
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// pubCmd represents the pub command
var pubCmd = &cobra.Command{
	Use: "pub",
	Run: func(cmd *cobra.Command, args []string) {
		mqtt.ERROR = log.New(os.Stdout, "", 0)
		opts := mqtt.NewClientOptions().AddBroker("tcp://broker.emqx.io:1883").SetClientID(args[0])

		c := mqtt.NewClient(opts)
		if token := c.Connect(); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}

		// Publish a message
		token := c.Publish("zahid/hassan", 0, false, "Hello World")
		token.Wait()

		// Disconnect
		c.Disconnect(250)
		time.Sleep(1 * time.Second)
	},
}
