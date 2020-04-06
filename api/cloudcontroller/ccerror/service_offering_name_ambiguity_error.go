package ccerror

import "fmt"

type ServiceOfferingNameAmbiguityError struct {
	Name               string
	ServiceBrokerNames []string
}

func (e ServiceOfferingNameAmbiguityError) Error() string {
	return fmt.Sprintf(`ambiguous name '%s' provided by service brokers: %v`, e.Name, e.ServiceBrokerNames)
}
