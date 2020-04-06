package ccerror

import "fmt"

type ServiceOfferingNotFoundError struct {
	Name, ServiceBrokerName string
}

func (e ServiceOfferingNotFoundError) Error() string {
	return fmt.Sprintf(`service offering '%s' for broker '%s' not found`, e.Name, e.ServiceBrokerName)
}
