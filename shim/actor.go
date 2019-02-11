package shim

import (
	"code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv2"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv2/constant"
)

type Actor struct{}

func (a Actor) GetServiceInstancesSummaryBySpace(service, spaceGUID string) ([]v2action.ServiceInstanceSummary, v2action.Warnings, error) {
	switch service {
	case "mysql":
		return []v2action.ServiceInstanceSummary{
			{
				ServiceInstance: v2action.ServiceInstance{
					Name:          "sqltime",
					LastOperation: ccv2.LastOperation{Type: "create", State: constant.LastOperationSucceeded},
				},
				Service: v2action.Service{
					Label: "mysql",
				},
				ServicePlan: v2action.ServicePlan{
					Name: "massive",
				},
				BoundApplications: []v2action.BoundApplication{
					{AppName: "cache"},
				},
				Environment: "test",
				Org:         "test-1",
				Space:       "test",
			},
		}, nil, nil
	case "redis":
		return []v2action.ServiceInstanceSummary{
			{
				ServiceInstance: v2action.ServiceInstance{
					Name:          "will-db",
					LastOperation: ccv2.LastOperation{Type: "create", State: constant.LastOperationSucceeded},
				},
				Service: v2action.Service{
					Label: "redis",
				},
				ServicePlan: v2action.ServicePlan{
					Name: "large",
				},
				BoundApplications: []v2action.BoundApplication{
					{AppName: "test"},
				},
				Environment: "test",
				Org:         "test-1",
				Space:       "test",
			},
			{
				ServiceInstance: v2action.ServiceInstance{
					Name:          "mydb",
					LastOperation: ccv2.LastOperation{Type: "create", State: constant.LastOperationSucceeded},
				},
				Service: v2action.Service{
					Label: "redis",
				},
				ServicePlan: v2action.ServicePlan{
					Name: "medium",
				},
				BoundApplications: []v2action.BoundApplication{
					{AppName: "store-web"},
					{AppName: "store-worker"},
					{AppName: "cache"},
				},
				Environment: "test",
				Org:         "test-1",
				Space:       "test",
			},
		}, nil, nil
	case "":
		return []v2action.ServiceInstanceSummary{
			{
				ServiceInstance: v2action.ServiceInstance{
					Name:          "will-db",
					LastOperation: ccv2.LastOperation{Type: "create", State: constant.LastOperationSucceeded},
				},
				Service: v2action.Service{
					Label: "redis",
				},
				ServicePlan: v2action.ServicePlan{
					Name: "large",
				},
				BoundApplications: []v2action.BoundApplication{
					{AppName: "test"},
				},
				Environment: "test",
				Org:         "test-1",
				Space:       "test",
			},
			{
				ServiceInstance: v2action.ServiceInstance{
					Name:          "mydb",
					LastOperation: ccv2.LastOperation{Type: "create", State: constant.LastOperationSucceeded},
				},
				Service: v2action.Service{
					Label: "redis",
				},
				ServicePlan: v2action.ServicePlan{
					Name: "medium",
				},
				BoundApplications: []v2action.BoundApplication{
					{AppName: "store-web"},
					{AppName: "store-worker"},
					{AppName: "cache"},
				},
				Environment: "test",
				Org:         "test-1",
				Space:       "test",
			},
			{
				ServiceInstance: v2action.ServiceInstance{
					Name:          "sqltime",
					LastOperation: ccv2.LastOperation{Type: "create", State: constant.LastOperationSucceeded},
				},
				Service: v2action.Service{
					Label: "mysql",
				},
				ServicePlan: v2action.ServicePlan{
					Name: "massive",
				},
				BoundApplications: []v2action.BoundApplication{
					{AppName: "myapp"},
				},
				Environment: "test",
				Org:         "test-1",
				Space:       "test",
			},
		}, nil, nil
	default:
		return []v2action.ServiceInstanceSummary{}, nil, nil
	}
}

func (a Actor) GetServiceInstancesSummary(service string) ([]v2action.ServiceInstanceSummary, v2action.Warnings, error) {
	mysqlServiceInstanceSummaries := []v2action.ServiceInstanceSummary{
		{
			ServiceInstance: v2action.ServiceInstance{
				Name:          "sqltime",
				LastOperation: ccv2.LastOperation{Type: "create", State: constant.LastOperationSucceeded},
			},
			Service: v2action.Service{
				Label: "mysql",
			},
			ServicePlan: v2action.ServicePlan{
				Name: "massive",
			},
			BoundApplications: []v2action.BoundApplication{
				{AppName: "timekeeper"},
			},
			Environment: "test",
			Org:         "test-1",
			Space:       "test",
		},
		{
			ServiceInstance: v2action.ServiceInstance{
				Name:          "sequeled",
				LastOperation: ccv2.LastOperation{Type: "create", State: constant.LastOperationSucceeded},
			},
			Service: v2action.Service{
				Label: "mysql",
			},
			ServicePlan: v2action.ServicePlan{
				Name: "massive",
			},
			BoundApplications: []v2action.BoundApplication{
				{AppName: "atlas"},
				{AppName: "compass"},
			},
			Environment: "prod",
			Org:         "org1",
			Space:       "space1",
		},
	}

	serviceInstanceSummaries := []v2action.ServiceInstanceSummary{}
	switch service {
	case "":
		serviceInstanceSummaries = append(serviceInstanceSummaries, mysqlServiceInstanceSummaries...)
		serviceInstanceSummaries = append(serviceInstanceSummaries, getLaurelServices()...)
		serviceInstanceSummaries = append(serviceInstanceSummaries, getWillServices()...)
		serviceInstanceSummaries = append(serviceInstanceSummaries, getAnnamarieServices()...)
	case "redis":
		serviceInstanceSummaries = append(serviceInstanceSummaries, getLaurelServices()...)
		serviceInstanceSummaries = append(serviceInstanceSummaries, getWillServices()...)
		serviceInstanceSummaries = append(serviceInstanceSummaries, getAnnamarieServices()...)
	case "mysql":
		serviceInstanceSummaries = append(serviceInstanceSummaries, mysqlServiceInstanceSummaries...)
	}

	return serviceInstanceSummaries, nil, nil
}

func getLaurelServices() []v2action.ServiceInstanceSummary {
	return []v2action.ServiceInstanceSummary{
		{
			ServiceInstance: v2action.ServiceInstance{
				Name:          "mydb",
				LastOperation: ccv2.LastOperation{Type: "create", State: constant.LastOperationSucceeded},
			},
			Service: v2action.Service{
				Label: "redis",
			},
			ServicePlan: v2action.ServicePlan{
				Name: "large",
			},
			BoundApplications: []v2action.BoundApplication{
				{AppName: "store-web"},
				{AppName: "store-worker"},
				{AppName: "cache"},
			},
			Environment: "prod",
			Org:         "prod",
			Space:       "prod",
		},
		{
			ServiceInstance: v2action.ServiceInstance{
				Name:          "mydb",
				LastOperation: ccv2.LastOperation{Type: "create", State: constant.LastOperationSucceeded},
			},
			Service: v2action.Service{
				Label: "redis",
			},
			ServicePlan: v2action.ServicePlan{
				Name: "large",
			},
			BoundApplications: []v2action.BoundApplication{
				{AppName: "cache"},
			},
			Environment: "staging",
			Org:         "staging",
			Space:       "staging",
		},
		{
			ServiceInstance: v2action.ServiceInstance{
				Name:          "mydb",
				LastOperation: ccv2.LastOperation{Type: "create", State: constant.LastOperationSucceeded},
			},
			Service: v2action.Service{
				Label: "redis",
			},
			ServicePlan: v2action.ServicePlan{
				Name: "large",
			},
			BoundApplications: []v2action.BoundApplication{
				{AppName: "store-web"},
				{AppName: "store-worker"},
				{AppName: "cache"},
			},
			Environment: "test",
			Org:         "test-1",
			Space:       "test",
		},
		{
			ServiceInstance: v2action.ServiceInstance{
				Name:          "mydb",
				LastOperation: ccv2.LastOperation{Type: "create", State: constant.LastOperationSucceeded},
			},
			Service: v2action.Service{
				Label: "redis",
			},
			ServicePlan: v2action.ServicePlan{
				Name: "large",
			},
			BoundApplications: []v2action.BoundApplication{
				{AppName: "store-web"},
				{AppName: "store-worker"},
				{AppName: "cache"},
			},
			Environment: "never",
			Org:         "test",
			Space:       "test",
		},
		{
			ServiceInstance: v2action.ServiceInstance{
				Name:          "mydb",
				LastOperation: ccv2.LastOperation{Type: "create", State: constant.LastOperationSucceeded},
			},
			Service: v2action.Service{
				Label: "redis",
			},
			ServicePlan: v2action.ServicePlan{
				Name: "large",
			},
			BoundApplications: []v2action.BoundApplication{
				{AppName: "cache"},
			},
			Environment: "gonna",
			Org:         "test",
			Space:       "test",
		},
		{
			ServiceInstance: v2action.ServiceInstance{
				Name:          "mydb",
				LastOperation: ccv2.LastOperation{Type: "create", State: constant.LastOperationSucceeded},
			},
			Service: v2action.Service{
				Label: "redis",
			},
			ServicePlan: v2action.ServicePlan{
				Name: "large",
			},
			BoundApplications: []v2action.BoundApplication{
				{AppName: "store-web"},
				{AppName: "store-worker"},
				{AppName: "cache"},
			},
			Environment: "give",
			Org:         "test",
			Space:       "test",
		},
		{
			ServiceInstance: v2action.ServiceInstance{
				Name:          "mydb",
				LastOperation: ccv2.LastOperation{Type: "create", State: constant.LastOperationSucceeded},
			},
			Service: v2action.Service{
				Label: "redis",
			},
			ServicePlan: v2action.ServicePlan{
				Name: "large",
			},
			BoundApplications: []v2action.BoundApplication{
				{AppName: "store-web"},
				{AppName: "store-worker"},
			},
			Environment: "you",
			Org:         "test",
			Space:       "test",
		},
		{
			ServiceInstance: v2action.ServiceInstance{
				Name:          "mydb",
				LastOperation: ccv2.LastOperation{Type: "create", State: constant.LastOperationSucceeded},
			},
			Service: v2action.Service{
				Label: "redis",
			},
			ServicePlan: v2action.ServicePlan{
				Name: "large",
			},
			BoundApplications: []v2action.BoundApplication{
				{AppName: "store-web"},
				{AppName: "store-worker"},
			},
			Environment: "up",
			Org:         "test",
			Space:       "test",
		},
	}
}

func getWillServices() []v2action.ServiceInstanceSummary {
	return []v2action.ServiceInstanceSummary{
		{
			ServiceInstance: v2action.ServiceInstance{
				Name:          "will-db",
				LastOperation: ccv2.LastOperation{Type: "create", State: constant.LastOperationSucceeded},
			},
			Service: v2action.Service{
				Label: "redis",
			},
			ServicePlan: v2action.ServicePlan{
				Name: "large",
			},
			BoundApplications: []v2action.BoundApplication{
				{AppName: "cache"},
				{AppName: "store-worker"},
			},
			Environment: "prod",
			Org:         "prod",
			Space:       "prod",
		},
		{
			ServiceInstance: v2action.ServiceInstance{
				Name:          "will-db",
				LastOperation: ccv2.LastOperation{Type: "create", State: constant.LastOperationSucceeded},
			},
			Service: v2action.Service{
				Label: "redis",
			},
			ServicePlan: v2action.ServicePlan{
				Name: "large",
			},
			BoundApplications: []v2action.BoundApplication{
				{AppName: "cache"},
				{AppName: "store-worker"},
			},
			Environment: "staging",
			Org:         "staging",
			Space:       "staging",
		},
		{
			ServiceInstance: v2action.ServiceInstance{
				Name:          "will-db",
				LastOperation: ccv2.LastOperation{Type: "create", State: constant.LastOperationSucceeded},
			},
			Service: v2action.Service{
				Label: "redis",
			},
			ServicePlan: v2action.ServicePlan{
				Name: "large",
			},
			BoundApplications: []v2action.BoundApplication{
				{AppName: "cache"},
			},
			Environment: "test",
			Org:         "test-1",
			Space:       "test",
		},
		{
			ServiceInstance: v2action.ServiceInstance{
				Name:          "will-db",
				LastOperation: ccv2.LastOperation{Type: "create", State: constant.LastOperationSucceeded},
			},
			Service: v2action.Service{
				Label: "redis",
			},
			ServicePlan: v2action.ServicePlan{
				Name: "large",
			},
			BoundApplications: []v2action.BoundApplication{
				{AppName: "cache"},
			},
			Environment: "never",
			Org:         "test",
			Space:       "test",
		},
		{
			ServiceInstance: v2action.ServiceInstance{
				Name:          "will-db",
				LastOperation: ccv2.LastOperation{Type: "create", State: constant.LastOperationSucceeded},
			},
			Service: v2action.Service{
				Label: "redis",
			},
			ServicePlan: v2action.ServicePlan{
				Name: "large",
			},
			BoundApplications: []v2action.BoundApplication{
				{AppName: "cache"},
			},
			Environment: "gonna",
			Org:         "test",
			Space:       "test",
		},
		{
			ServiceInstance: v2action.ServiceInstance{
				Name:          "will-db",
				LastOperation: ccv2.LastOperation{Type: "create", State: constant.LastOperationSucceeded},
			},
			Service: v2action.Service{
				Label: "redis",
			},
			ServicePlan: v2action.ServicePlan{
				Name: "large",
			},
			BoundApplications: []v2action.BoundApplication{
				{AppName: "cache"},
				{AppName: "store-worker"},
			},
			Environment: "give",
			Org:         "test",
			Space:       "test",
		},
		{
			ServiceInstance: v2action.ServiceInstance{
				Name:          "will-db",
				LastOperation: ccv2.LastOperation{Type: "create", State: constant.LastOperationSucceeded},
			},
			Service: v2action.Service{
				Label: "redis",
			},
			ServicePlan: v2action.ServicePlan{
				Name: "large",
			},
			BoundApplications: []v2action.BoundApplication{
				{AppName: "cache"},
				{AppName: "store-web"},
			},
			Environment: "you",
			Org:         "test",
			Space:       "test",
		},
		{
			ServiceInstance: v2action.ServiceInstance{
				Name:          "will-db",
				LastOperation: ccv2.LastOperation{Type: "create", State: constant.LastOperationSucceeded},
			},
			Service: v2action.Service{
				Label: "redis",
			},
			ServicePlan: v2action.ServicePlan{
				Name: "large",
			},
			BoundApplications: []v2action.BoundApplication{
				{AppName: "cache"},
			},
			Environment: "up",
			Org:         "test",
			Space:       "test",
		},
	}
}

func getAnnamarieServices() []v2action.ServiceInstanceSummary {
	return []v2action.ServiceInstanceSummary{
		{
			ServiceInstance: v2action.ServiceInstance{
				Name:          "awesome-db",
				LastOperation: ccv2.LastOperation{Type: "create", State: constant.LastOperationSucceeded},
			},
			Service: v2action.Service{
				Label: "redis",
			},
			ServicePlan: v2action.ServicePlan{
				Name: "small",
			},
			BoundApplications: []v2action.BoundApplication{
				{AppName: "design-studio"},
			},
			Environment: "test",
			Org:         "test",
			Space:       "test",
		},
		{
			ServiceInstance: v2action.ServiceInstance{
				Name:          "awesome-db",
				LastOperation: ccv2.LastOperation{Type: "create", State: constant.LastOperationSucceeded},
			},
			Service: v2action.Service{
				Label: "redis",
			},
			ServicePlan: v2action.ServicePlan{
				Name: "small",
			},
			BoundApplications: []v2action.BoundApplication{
				{AppName: "design-studio"},
			},
			Environment: "test",
			Org:         "test-1",
			Space:       "test",
		},
		{
			ServiceInstance: v2action.ServiceInstance{
				Name:          "awesome-db",
				LastOperation: ccv2.LastOperation{Type: "create", State: constant.LastOperationSucceeded},
			},
			Service: v2action.Service{
				Label: "redis",
			},
			ServicePlan: v2action.ServicePlan{
				Name: "small",
			},
			BoundApplications: []v2action.BoundApplication{
				{AppName: "design-studio"},
			},
			Environment: "staging",
			Org:         "staging",
			Space:       "staging",
		},
		{
			ServiceInstance: v2action.ServiceInstance{
				Name:          "awesome-db",
				LastOperation: ccv2.LastOperation{Type: "create", State: constant.LastOperationSucceeded},
			},
			Service: v2action.Service{
				Label: "redis",
			},
			ServicePlan: v2action.ServicePlan{
				Name: "medium",
			},
			BoundApplications: []v2action.BoundApplication{
				{AppName: "design-studio"},
			},
			Environment: "staging",
			Org:         "staging-1",
			Space:       "staging",
		},
		{
			ServiceInstance: v2action.ServiceInstance{
				Name:          "awesome-db",
				LastOperation: ccv2.LastOperation{Type: "create", State: constant.LastOperationSucceeded},
			},
			Service: v2action.Service{
				Label: "redis",
			},
			ServicePlan: v2action.ServicePlan{
				Name: "medium",
			},
			BoundApplications: []v2action.BoundApplication{
				{AppName: "design-studio"},
			},
			Environment: "prod",
			Org:         "prod",
			Space:       "prod",
		},
		{
			ServiceInstance: v2action.ServiceInstance{
				Name:          "awesome-db",
				LastOperation: ccv2.LastOperation{Type: "create", State: constant.LastOperationSucceeded},
			},
			Service: v2action.Service{
				Label: "redis",
			},
			ServicePlan: v2action.ServicePlan{
				Name: "large",
			},
			BoundApplications: []v2action.BoundApplication{
				{AppName: "design-studio"},
			},
			Environment: "prod",
			Org:         "prod-HA",
			Space:       "prod-HA",
		},
		{
			ServiceInstance: v2action.ServiceInstance{
				Name:          "awesome-db",
				LastOperation: ccv2.LastOperation{Type: "create", State: constant.LastOperationSucceeded},
			},
			Service: v2action.Service{
				Label: "redis",
			},
			ServicePlan: v2action.ServicePlan{
				Name: "large",
			},
			BoundApplications: []v2action.BoundApplication{
				{AppName: "design-studio"},
			},
			Environment: "you",
			Org:         "test",
			Space:       "test",
		},
		{
			ServiceInstance: v2action.ServiceInstance{
				Name:          "awesome-db",
				LastOperation: ccv2.LastOperation{Type: "create", State: constant.LastOperationSucceeded},
			},
			Service: v2action.Service{
				Label: "redis",
			},
			ServicePlan: v2action.ServicePlan{
				Name: "large",
			},
			BoundApplications: []v2action.BoundApplication{
				{AppName: "design-studio"},
			},
			Environment: "give",
			Org:         "test",
			Space:       "test",
		},
	}
}
