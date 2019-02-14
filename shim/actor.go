package shim

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv2"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv2/constant"
)

type Actor struct{}

var (
	s1 = v2action.ServiceInstanceSummary{
		ServiceInstance: v2action.ServiceInstance{
			Name: "s1",
			LastOperation: ccv2.LastOperation{
				Type:        "create",
				State:       constant.LastOperationSucceeded,
				Description: "The operation has finished",
				UpdatedAt:   "2019-02-11T11:24:37Z",
				CreatedAt:   "2019-02-11T11:24:37Z",
			},
			Type:         constant.ServiceInstanceTypeManagedService,
			DashboardURL: "http://example.org/dashboard",
		},
		Service: v2action.Service{
			Label:             "redis",
			Description:       "Some redis service.",
			DocumentationURL:  "http://example.org/documentation",
			ServiceBrokerName: "awesome-broker",
		},
		ServicePlan: v2action.ServicePlan{
			Name: "small",
		},
		BoundApplications: []v2action.BoundApplication{},
		UpdateAvailable:   true,
		UpdateDetails: `- redis version update from v0.9 to v1.1
- OS update`,
	}

	s2 = v2action.ServiceInstanceSummary{
		ServiceInstance: v2action.ServiceInstance{
			Name: "s2",
			LastOperation: ccv2.LastOperation{
				Type:        "create",
				State:       constant.LastOperationSucceeded,
				Description: "The operation has finished!",
				UpdatedAt:   "2019-02-11T11:24:37Z",
				CreatedAt:   "2019-02-11T11:24:37Z",
			},
			Type:         constant.ServiceInstanceTypeManagedService,
			DashboardURL: "http://example.org/dashboard",
		},
		Service: v2action.Service{
			Label:             "redis",
			Description:       "Some some redis service.",
			DocumentationURL:  "http://example.org/documentation",
			ServiceBrokerName: "awesome-broker",
		},
		ServicePlan: v2action.ServicePlan{
			Name: "small",
		},
		BoundApplications: []v2action.BoundApplication{},
		UpdateAvailable:   true,
		UpdateDetails:     `- redis version update from v1.0 to v1.1`,
	}

	s3 = v2action.ServiceInstanceSummary{
		ServiceInstance: v2action.ServiceInstance{
			Name: "s3",
			LastOperation: ccv2.LastOperation{
				Type:        "create",
				State:       constant.LastOperationSucceeded,
				Description: "The operation has finished!",
				UpdatedAt:   "2019-02-11T11:24:37Z",
				CreatedAt:   "2019-02-11T11:24:37Z",
			},
			Type:         constant.ServiceInstanceTypeManagedService,
			DashboardURL: "http://example.org/dashboard",
		},
		Service: v2action.Service{
			Label:             "mysql",
			Description:       "Some mysql service.",
			DocumentationURL:  "http://example.org/documentation",
			ServiceBrokerName: "mysql-broker",
		},
		ServicePlan: v2action.ServicePlan{
			Name: "large",
		},
		BoundApplications: []v2action.BoundApplication{},
		UpdateAvailable:   false,
		UpdateDetails:     ``,
	}

	services = map[string]*v2action.ServiceInstanceSummary{
		"s1": &s1,
		"s2": &s2,
		"s3": &s3,
	}
)

type ServiceState struct {
	UpdateStatus string
}

type PrototypeState struct {
	Services map[string]*ServiceState
}

var state PrototypeState

func ensureOk(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	file, err := os.Open("prototype-state.json")
	if err != nil {
		fmt.Println(err)
		file, err = os.OpenFile("prototype-state.json", os.O_RDWR|os.O_CREATE, 0644)
		ensureOk(err)

		state := PrototypeState{
			Services: map[string]*ServiceState{
				"s1": &ServiceState{UpdateStatus: "available"},
				"s2": &ServiceState{UpdateStatus: "available"},
				"s3": &ServiceState{UpdateStatus: "unavailable"},
			},
		}
		bytes, err := json.Marshal(state)
		ensureOk(err)

		_, err = file.Write(bytes)
		ensureOk(err)
		ensureOk(file.Close())

		file, err = os.Open("prototype-state.json")
		ensureOk(err)
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	ensureOk(err)
	ensureOk(json.Unmarshal(bytes, &state))

	s1.UpdateAvailable = state.Services["s1"].UpdateStatus == "available"
	if state.Services["s1"].UpdateStatus == "in-progress" {
		s1.ServiceInstance.LastOperation.Type = "update"
		s1.ServiceInstance.LastOperation.State = constant.LastOperationInProgress
		s1.ServiceInstance.LastOperation.Description = "The operation is in progress..."
	}

	s2.UpdateAvailable = state.Services["s2"].UpdateStatus == "available"
	if state.Services["s2"].UpdateStatus == "in-progress" {
		s2.ServiceInstance.LastOperation.Type = "update"
		s2.ServiceInstance.LastOperation.State = constant.LastOperationInProgress
		s2.ServiceInstance.LastOperation.Description = "The operation is in progress..."
	}

	s3.UpdateAvailable = state.Services["s3"].UpdateStatus == "available"
	if state.Services["s3"].UpdateStatus == "in-progress" {
		s3.ServiceInstance.LastOperation.Type = "update"
		s3.ServiceInstance.LastOperation.State = constant.LastOperationInProgress
		s3.ServiceInstance.LastOperation.Description = "The operation is in progress..."
	}
}

func (actor Actor) GetServiceInstancesSummaryBySpace(spaceGUID string) ([]v2action.ServiceInstanceSummary, v2action.Warnings, error) {
	return []v2action.ServiceInstanceSummary{s1, s2, s3}, nil, nil
}

func (actor Actor) GetServiceInstanceByNameAndSpace(name, spaceGUID string) (v2action.ServiceInstance, v2action.Warnings, error) {
	return services[name].ServiceInstance, nil, nil
}

func (actor Actor) GetServiceInstanceSummaryByNameAndSpace(name, spaceGUID string) (v2action.ServiceInstanceSummary, v2action.Warnings, error) {
	return *services[name], nil, nil
}

func (actor Actor) UpdateDone(name string) {
	serviceState := state.Services[name]
	serviceState.UpdateStatus = "in-progress"

	bytes, err := json.Marshal(state)
	ensureOk(err)
	ensureOk(ioutil.WriteFile("prototype-state.json", bytes, 0644))
}
