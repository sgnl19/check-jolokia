package attributes

import (
	"fmt"
	"github.com/benkeil/icinga-checks-library"
	"github.com/s8sg/go_jolokia"
	"github.com/sgnl04/check-jolokia/pkg/utils"
	"log"
)

type (
	// CheckObjectProperty interface to check a query string
	CheckObjectProperty interface {
		CheckQueueAttributeQuery(CheckQueueAttributeOptions) icinga.Result
	}

	checkQueueAttributeImpl struct {
		JolokiaClient go_jolokia.JolokiaClient
		Url           string
	}
)

// NewCheckQueueAttribute creates a new instance of CheckObjectProperty
func NewCheckQueueAttribute(client go_jolokia.JolokiaClient, url string) CheckObjectProperty {
	return &checkQueueAttributeImpl{JolokiaClient: client, Url: url}
}

// CheckAvailableAddressesOptions contains options needed to run CheckAvailableAddresses check
type CheckQueueAttributeOptions struct {
	ThresholdWarning  string
	ThresholdCritical string
	Url               string
	Domain            string
	Queue             string
	Attribute         string
	Verbose           int
}

// CheckAvailableAddresses checks if the deployment has a minimum of available replicas
func (c *checkQueueAttributeImpl) CheckQueueAttributeQuery(options CheckQueueAttributeOptions) icinga.Result {
	name := "Queue.Attributes"

	statusCheck, err := icinga.NewStatusCheck(options.ThresholdWarning, options.ThresholdCritical)
	if err != nil {
		return icinga.NewResult(name, icinga.ServiceStatusUnknown, fmt.Sprintf("can't check status: %v", err))
	}

	searchResult, err := c.JolokiaClient.GetAttr(options.Domain, []string{options.Queue}, options.Attribute)
	if err != nil {
		return icinga.NewResult(name, icinga.ServiceStatusUnknown, fmt.Sprintf("can't query Jolokia: %v", err))
	}

	result, err := utils.ToFloat(searchResult)
	if err != nil {
		if (options.Verbose > 0) {
			log.Printf("An error occured with result [%v]", searchResult)
		}
		return icinga.NewResult(name, icinga.ServiceStatusUnknown, fmt.Sprintf("query result is invalid: %v", err))
	}

	message := fmt.Sprintf("Search produced: %v", searchResult)
	status := statusCheck.Check(result)

	return icinga.NewResult(name, status, message)
}
