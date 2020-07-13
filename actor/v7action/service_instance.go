package v7action

import (
	"code.cloudfoundry.org/cli/actor/actionerror"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccerror"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3"
	"code.cloudfoundry.org/cli/resources"
	"code.cloudfoundry.org/cli/types"
)

type ServiceInstanceWithRelationships struct {
	resources.ServiceInstance
	Parameters        types.OptionalObject
	ServiceOffering   resources.ServiceOffering
	ServicePlanName   string
	ServiceBrokerName string
}

func (actor Actor) GetServiceInstanceByNameAndSpace(serviceInstanceName string, spaceGUID string) (resources.ServiceInstance, Warnings, error) {
	serviceInstance, _, warnings, err := actor.CloudControllerClient.GetServiceInstanceByNameAndSpace(serviceInstanceName, spaceGUID)
	return serviceInstance, Warnings(warnings), err
}

func (actor Actor) GetServiceInstanceDetails(serviceInstanceName string, spaceGUID string) (ServiceInstanceWithRelationships, Warnings, error) {
	query := []ccv3.Query{
		{Key: ccv3.FieldsServicePlan, Values: []string{"name", "guid"}},
		{Key: ccv3.FieldsServicePlanServiceOffering, Values: []string{"name", "guid", "description", "documentation_url"}},
		{Key: ccv3.FieldsServicePlanServiceOfferingServiceBroker, Values: []string{"name", "guid"}},
	}

	serviceInstance, included, warnings, err := actor.CloudControllerClient.GetServiceInstanceByNameAndSpace(serviceInstanceName, spaceGUID, query...)
	switch err.(type) {
	case nil:
	case ccerror.ServiceInstanceNotFoundError:
		return ServiceInstanceWithRelationships{}, Warnings(warnings), actionerror.ServiceInstanceNotFoundError{Name: serviceInstanceName}
	default:
		return ServiceInstanceWithRelationships{}, Warnings(warnings), err
	}
	result := ServiceInstanceWithRelationships{ServiceInstance: serviceInstance}

	if shouldRequestServiceInstanceParameters(serviceInstance) {
		params, paramsWarnings, err := actor.CloudControllerClient.GetServiceInstanceParameters(serviceInstance.GUID)
		warnings = append(warnings, paramsWarnings...)
		// ignore errors because we currently we request parameters even when the broker does
		// not allow instance retrieval
		if err == nil {
			result.Parameters = types.NewOptionalObject(params)
		}
	}

	if len(included.ServicePlans) == 1 {
		result.ServicePlanName = included.ServicePlans[0].Name
	}

	if len(included.ServiceOfferings) == 1 {
		result.ServiceOffering = included.ServiceOfferings[0]
	}

	if len(included.ServiceBrokers) == 1 {
		result.ServiceBrokerName = included.ServiceBrokers[0].Name
	}

	return result, Warnings(warnings), nil
}

func (actor Actor) UnshareServiceInstanceByServiceInstanceAndSpace(serviceInstanceGUID string, sharedToSpaceGUID string) (Warnings, error) {
	warnings, err := actor.CloudControllerClient.DeleteServiceInstanceRelationshipsSharedSpace(serviceInstanceGUID, sharedToSpaceGUID)
	return Warnings(warnings), err
}

func (actor Actor) CreateUserProvidedServiceInstance(serviceInstance resources.ServiceInstance) (Warnings, error) {
	serviceInstance.Type = resources.UserProvidedServiceInstance
	_, warnings, err := actor.CloudControllerClient.CreateServiceInstance(serviceInstance)
	return Warnings(warnings), err
}

func (actor Actor) UpdateUserProvidedServiceInstance(serviceInstanceName, spaceGUID string, serviceInstanceUpdates resources.ServiceInstance) (Warnings, error) {
	original, _, warnings, err := actor.CloudControllerClient.GetServiceInstanceByNameAndSpace(serviceInstanceName, spaceGUID)
	if err != nil {
		return Warnings(warnings), err
	}

	if original.Type != resources.UserProvidedServiceInstance {
		return Warnings(warnings), actionerror.ServiceInstanceTypeError{
			Name:         serviceInstanceName,
			RequiredType: resources.UserProvidedServiceInstance,
		}
	}

	_, updateWarnings, err := actor.CloudControllerClient.UpdateServiceInstance(original.GUID, serviceInstanceUpdates)
	warnings = append(warnings, updateWarnings...)
	if err != nil {
		return Warnings(warnings), err
	}

	return Warnings(warnings), nil
}

func shouldRequestServiceInstanceParameters(s resources.ServiceInstance) bool {
	// TODO: should also check if the "instances_retrievable" flag is set in the service offering
	return s.Type == resources.ManagedServiceInstance &&
		s.LastOperation.State != resources.OperationInProgress
}
