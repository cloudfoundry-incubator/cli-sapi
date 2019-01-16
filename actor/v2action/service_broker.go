package v2action

import (
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv2"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv2/constant"
	"fmt"
)

type ServiceBroker ccv2.ServiceBroker

func (actor Actor) CreateServiceBroker(serviceBrokerName, username, password, brokerURI, spaceGUID string) (ServiceBroker, Warnings, error) {
	serviceBroker, warnings, err := actor.CloudControllerClient.CreateServiceBroker(serviceBrokerName, username, password, brokerURI, spaceGUID)
	return ServiceBroker(serviceBroker), Warnings(warnings), err
}

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

func (actor Actor) GetServiceBrokerByName(brokerName string) (ServiceBroker, Warnings, error) {
	serviceBrokers, warnings, err := actor.CloudControllerClient.GetServiceBrokers(ccv2.Filter{
		Type:     constant.NameFilter,
		Operator: constant.EqualOperator,
		Values:   []string{brokerName},
	})
	if err != nil {
		return ServiceBroker{}, Warnings(warnings), err
	}

	if len(serviceBrokers) < 1 {
		return ServiceBroker{}, Warnings(warnings), fmt.Errorf("No service broker with name %s", brokerName)
	}

	return ServiceBroker(serviceBrokers[0]), Warnings(warnings), nil
}

func (actor Actor) MigrateServiceBrokerByName(brokerName string) error {
	broker, _, err := actor.GetServiceBrokerByName(brokerName)
	if err != nil {
		return err
	}

	return actor.CloudControllerClient.MigrateServiceBroker(broker.GUID)
}
