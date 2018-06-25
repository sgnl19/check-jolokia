package main

import (
	"io"
	"github.com/s8sg/go_jolokia"
	"github.com/sgnl04/check-jolokia/pkg/checks/queue/attributes"
	"github.com/spf13/cobra"
)

type (
	checkQueueAttributeCmd struct {
		out           io.Writer
		JolokiaClient go_jolokia.JolokiaClient
		Warning       string
		Critical      string
		Url           string
		Domain        string
		Queue         string
		Attribute     string
		Verbose       int
	}
)

func newCheckQueueAttributeCmd(out io.Writer) *cobra.Command {
	c := &checkQueueAttributeCmd{out: out}

	cmd := &cobra.Command{
		Use:          "queueAttribute",
		Short:        "check if a Jolokia queue attribute query result meets the thresholds",
		SilenceUsage: false,
		Args:         NameArgs(),
		PreRun: func(cmd *cobra.Command, args []string) {
			c.Url = args[0]
			client := go_jolokia.NewJolokiaClient(c.Url)
			c.JolokiaClient = *client
		},
		Run: func(cmd *cobra.Command, args []string) {
			c.run()
		},
	}

	cmd.Flags().StringVarP(&c.Critical, "critical", "c", "10:", "critical threshold for minimum amount of result")
	cmd.Flags().StringVarP(&c.Warning, "warning", "w", "5:", "warning threshold for minimum amount of result")
	cmd.Flags().StringVarP(&c.Domain, "domain", "d", "org.apache.activemq.artemis", "the domain of the queue to query")
	cmd.Flags().StringVarP(&c.Queue, "queue", "q", "*", "the queue to query")
	cmd.Flags().StringVarP(&c.Attribute, "attribute", "a", "*", "the attributes to query from the queue")
	cmd.Flags().CountVarP(&c.Verbose, "verbose", "v", "enable verbose output")

	return cmd
}

func (c *checkQueueAttributeCmd) run() {
	checkQueueAttribute := attributes.NewCheckQueueAttribute(c.JolokiaClient, c.Url)
	results := checkQueueAttribute.CheckQueueAttributeQuery(attributes.CheckQueueAttributeOptions{
		ThresholdWarning:  c.Warning,
		ThresholdCritical: c.Critical,
		Url:               c.Url,
		Domain:            c.Domain,
		Queue:             c.Queue,
		Attribute:         c.Attribute,
		Verbose:           c.Verbose,
	})
	results.Exit()
}
