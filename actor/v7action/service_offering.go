package v7action

import (
	"code.cloudfoundry.org/cli/actor/actionerror"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccerror"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3"
)

type ServiceOffering ccv3.ServiceOffering

func (actor Actor) GetServiceOfferingByNameAndBroker(serviceOfferingName, serviceBrokerName string) (ServiceOffering, Warnings, error) {
	serviceOffering, warnings, err := actor.CloudControllerClient.GetServiceOfferingByNameAndBroker(serviceOfferingName, serviceBrokerName)
	return ServiceOffering(serviceOffering), Warnings(warnings), convertAPIErrors(err)
}

func (actor Actor) PurgeServiceOfferingByNameAndBroker(serviceOfferingName, serviceBrokerName string) (Warnings, error) {
	_, warnings, err := actor.CloudControllerClient.GetServiceOfferingByNameAndBroker(serviceOfferingName, serviceBrokerName)
	return Warnings(warnings), convertAPIErrors(err)
}

func convertAPIErrors(e error) error {
	switch e.(type) {
	case ccerror.ServiceOfferingNotFoundError:
		return actionerror.ServiceNotFoundError{
			Name: e.(ccerror.ServiceOfferingNotFoundError).Name,
			Broker: e.
		}
	case ccerror.ServiceOfferingNameAmbiguityError:
		return actionerror.DuplicateServiceError{
			Name:           e.(ccerror.ServiceOfferingNameAmbiguityError).Name,
			ServiceBrokers: e.(ccerror.ServiceOfferingNameAmbiguityError).ServiceBrokerNames,
		}
	default:
		return e
	}
}
