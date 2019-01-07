package v2action

import "code.cloudfoundry.org/cli/api/cloudcontroller/ccv2"

// ServiceInstanceSharedFrom represents a CLI ServiceInstanceSharedFrom.
type ServiceInstanceSharedFrom ccv2.ServiceInstanceSharedFrom

// GetServiceInstanceSharedFromByServiceInstance returns details of the org and space the specified service instance is shared from.
func (actor Actor) GetServiceInstanceSharedFromByServiceInstance(serviceInstanceGUID string) (ServiceInstanceSharedFrom, Warnings, error) {
	serviceInstanceSharedFrom, warnings, err := actor.CloudControllerClient.GetServiceInstanceSharedFrom(serviceInstanceGUID)
	return ServiceInstanceSharedFrom(serviceInstanceSharedFrom), Warnings(warnings), err
}
