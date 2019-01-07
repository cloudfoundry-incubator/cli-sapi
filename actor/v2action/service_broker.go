package v2action

import (
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv2"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv2/constant"
)

// ServiceBroker represents a CLI Service Broker.
type ServiceBroker ccv2.ServiceBroker

// CreateServiceBroker creates and registers a service broker with the specified properties.
func (actor Actor) CreateServiceBroker(serviceBrokerName, username, password, brokerURI, spaceGUID string) (ServiceBroker, Warnings, error) {
	serviceBroker, warnings, err := actor.CloudControllerClient.CreateServiceBroker(serviceBrokerName, username, password, brokerURI, spaceGUID)
	return ServiceBroker(serviceBroker), Warnings(warnings), err
}

// GetServiceBrokers returns all registered service brokers.
func (actor Actor) GetServiceBrokers() ([]ServiceBroker, Warnings, error) {
	brokers, warnings, err := actor.CloudControllerClient.GetServiceBrokers()
	if err != nil {
		return nil, Warnings(warnings), err
	}

	var brokersToReturn []ServiceBroker
	for _, b := range brokers {
		brokersToReturn = append(brokersToReturn, ServiceBroker(b))
	}

	return brokersToReturn, Warnings(warnings), nil
}

// GetServiceBrokerByName returns a service broker with the specified name (or nothing if no such broker exists).
func (actor Actor) GetServiceBrokerByName(brokerName string) (ServiceBroker, Warnings, error) {
	serviceBrokers, warnings, err := actor.CloudControllerClient.GetServiceBrokers(ccv2.Filter{
		Type:     constant.NameFilter,
		Operator: constant.EqualOperator,
		Values:   []string{brokerName},
	})
	if err != nil {
		return ServiceBroker{}, Warnings(warnings), err
	}

	return ServiceBroker(serviceBrokers[0]), Warnings(warnings), nil
}
