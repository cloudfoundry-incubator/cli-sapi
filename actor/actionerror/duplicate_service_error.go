package actionerror

import (
	"fmt"
	"strings"
)

type DuplicateServiceError struct {
	Name           string
	ServiceBrokers []string
}

func (e DuplicateServiceError) Error() string {
	var quotedBrokers []string
	for _, broker := range e.ServiceBrokers {
		quotedBrokers = append(quotedBrokers, fmt.Sprintf("'%s'", broker))
	}
	availableBrokers := strings.Join(quotedBrokers, ", ")
	return fmt.Sprintf(
		"Service '%s' is provided by multiple service brokers.\n"+
			"Specify a broker from available brokers %s by using the '-b' flag.", e.Name, availableBrokers)
}
