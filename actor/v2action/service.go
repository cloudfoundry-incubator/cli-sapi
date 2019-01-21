package v2action

import (
	"code.cloudfoundry.org/cli/actor/actionerror"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv2"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv2/constant"
)

// Service represents a CLI Service.
type Service ccv2.Service

// GetService returns a service with the provided GUID.
func (actor Actor) GetService(serviceGUID string) (Service, Warnings, error) {
	service, warnings, err := actor.CloudControllerClient.GetService(serviceGUID)
	return Service(service), Warnings(warnings), err
}

// GetServiceByName returns a service based on the name provided.
// If there are no services, an ServiceNotFoundError will be returned.
// If there are multiple services, the first will be returned.
func (actor Actor) GetServiceByName(serviceName string) (Service, Warnings, error) {
	services, warnings, err := actor.CloudControllerClient.GetServices(ccv2.Filter{
		Type:     constant.LabelFilter,
		Operator: constant.EqualOperator,
		Values:   []string{serviceName},
	})
	if err != nil {
		return Service{}, Warnings(warnings), err
	}

	if len(services) == 0 {
		return Service{}, Warnings(warnings), actionerror.ServiceNotFoundError{Name: serviceName}
	}

	return Service(services[0]), Warnings(warnings), nil
}

func (actor Actor) getServiceByNameForSpace(serviceName, spaceGUID string) (Service, Warnings, error) {
	services, warnings, err := actor.CloudControllerClient.GetSpaceServices(spaceGUID, ccv2.Filter{
		Type:     constant.LabelFilter,
		Operator: constant.EqualOperator,
		Values:   []string{serviceName},
	})
	if err != nil {
		return Service{}, Warnings(warnings), err
	}

	if len(services) == 0 {
		return Service{}, Warnings(warnings), actionerror.ServiceNotFoundError{Name: serviceName}
	}

	return Service(services[0]), Warnings(warnings), nil
}

// ServicesWithPlans is a map of services and a slice of corresponding service plans.
type ServicesWithPlans map[Service][]ServicePlan

// GetServicesWithPlansForBroker returns all services and their corresponding services exposed by the specified broker.
func (actor Actor) GetServicesWithPlansForBroker(brokerGUID string) (ServicesWithPlans, Warnings, error) {
	var allWarnings Warnings
	services, warnings, err := actor.CloudControllerClient.GetServices(ccv2.Filter{
		Type:     constant.ServiceBrokerGUIDFilter,
		Operator: constant.EqualOperator,
		Values:   []string{brokerGUID},
	})
	allWarnings = append(allWarnings, Warnings(warnings)...)
	if err != nil {
		return nil, allWarnings, err
	}

	servicesWithPlans := ServicesWithPlans{}
	for _, service := range services {
		servicePlans, warnings, err := actor.CloudControllerClient.GetServicePlans(ccv2.Filter{
			Type:     constant.ServiceGUIDFilter,
			Operator: constant.EqualOperator,
			Values:   []string{service.GUID},
		})
		allWarnings = append(allWarnings, Warnings(warnings)...)
		if err != nil {
			return nil, allWarnings, err
		}

		plansToReturn := []ServicePlan{}
		for _, plan := range servicePlans {
			plansToReturn = append(plansToReturn, ServicePlan(plan))
		}

		servicesWithPlans[Service(service)] = plansToReturn
	}

	return servicesWithPlans, allWarnings, nil
}
