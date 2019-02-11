package v6

import (
	"fmt"
	"sort"
	"strings"

	"code.cloudfoundry.org/cli/actor/sharedaction"
	"code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv2/constant"
	"code.cloudfoundry.org/cli/command"
	"code.cloudfoundry.org/cli/shim"
	"code.cloudfoundry.org/cli/util/sorting"
)

//go:generate counterfeiter . ServiceInstancesActor

type ServiceInstancesActor interface {
	GetServiceInstancesSummaryBySpace(service, spaceGUID string) ([]v2action.ServiceInstanceSummary, v2action.Warnings, error)
	GetServiceInstancesSummary(service string) ([]v2action.ServiceInstanceSummary, v2action.Warnings, error)
}

type ServicesCommand struct {
	ServiceOffering string      `long:"service" short:"s" description:"show only services of this service type"`
	AllEnvironments bool        `long:"all-environments" short:"a" description:"show services from across all environments"`
	usage           interface{} `usage:"CF_NAME services"`
	relatedCommands interface{} `related_commands:"create-service, marketplace"`

	UI          command.UI
	Config      command.Config
	SharedActor command.SharedActor
	Actor       ServiceInstancesActor
}

func (cmd *ServicesCommand) Setup(config command.Config, ui command.UI) error {
	cmd.Config = config
	cmd.UI = ui
	cmd.SharedActor = sharedaction.NewActor(config)
	cmd.Actor = shim.Actor{}

	return nil
}

func (cmd ServicesCommand) Execute(args []string) error {
	if cmd.AllEnvironments {
		cmd.UI.DisplayTextWithFlavor("Getting services in all orgs and spaces in all environments as {{.CurrentUser}}...",
			map[string]interface{}{
				"CurrentUser": "admin",
			})
	} else {
		cmd.UI.DisplayTextWithFlavor("Getting services in org {{.OrgName}} / space {{.SpaceName}} as {{.CurrentUser}}...",
			map[string]interface{}{
				"OrgName":     "test-1",
				"SpaceName":   "test",
				"CurrentUser": "admin",
			})
	}

	cmd.UI.DisplayNewline()

	var (
		instanceSummaries []v2action.ServiceInstanceSummary
		warnings          v2action.Warnings
		err               error
	)

	if cmd.AllEnvironments {
		instanceSummaries, warnings, err = cmd.Actor.GetServiceInstancesSummary(cmd.ServiceOffering)
	} else {
		instanceSummaries, warnings, err = cmd.Actor.GetServiceInstancesSummaryBySpace(cmd.ServiceOffering, cmd.Config.TargetedSpace().GUID)
	}

	cmd.UI.DisplayWarnings(warnings)
	if err != nil {
		return err
	}

	if len(instanceSummaries) == 0 {
		cmd.UI.DisplayText("No services found")
		return nil
	}

	sortServiceInstances(instanceSummaries)

	table := [][]string{{
		cmd.UI.TranslateText("name"),
		cmd.UI.TranslateText("service"),
		cmd.UI.TranslateText("plan"),
		cmd.UI.TranslateText("bound apps"),
		cmd.UI.TranslateText("last operation"),
	}}

	if cmd.AllEnvironments {
		table = [][]string{{
			cmd.UI.TranslateText("name"),
			cmd.UI.TranslateText("service"),
			cmd.UI.TranslateText("plan"),
			cmd.UI.TranslateText("bound apps"),
			cmd.UI.TranslateText("last operation"),
			cmd.UI.TranslateText("environment"),
			cmd.UI.TranslateText("org"),
			cmd.UI.TranslateText("space"),
		}}
	}

	var boundAppNames []string
	var lastEnvironment = instanceSummaries[0].Environment

	for _, summary := range instanceSummaries {
		if cmd.AllEnvironments && summary.Environment != lastEnvironment {
			table = append(table, []string{"", "", "", "", "", "", "", ""})
			lastEnvironment = summary.Environment
		}

		serviceLabel := summary.Service.Label
		if summary.ServiceInstance.Type == constant.ServiceInstanceTypeUserProvidedService {
			serviceLabel = "user-provided"
		}

		boundAppNames = []string{}
		for _, boundApplication := range summary.BoundApplications {
			boundAppNames = append(boundAppNames, boundApplication.AppName)
		}

		row := []string{
			summary.Name,
			serviceLabel,
			summary.ServicePlan.Name,
			strings.Join(boundAppNames, ", "),
			fmt.Sprintf("%s %s", summary.LastOperation.Type, summary.LastOperation.State),
		}

		if cmd.AllEnvironments {
			row = append(row, summary.Environment)
			row = append(row, summary.Org)
			row = append(row, summary.Space)
		}

		table = append(table, row)
	}
	cmd.UI.DisplayTableWithHeader("", table, 3)

	return nil
}

func sortServiceInstances(instanceSummaries []v2action.ServiceInstanceSummary) {
	sort.Slice(instanceSummaries, func(i, j int) bool {
		if instanceSummaries[i].Environment < instanceSummaries[j].Environment {
			return true
		}

		if instanceSummaries[i].Environment > instanceSummaries[j].Environment {
			return false
		}

		return sorting.LessIgnoreCase(instanceSummaries[i].Name, instanceSummaries[j].Name)
	})

	for _, instance := range instanceSummaries {
		sortBoundApps(instance)
	}
}

func sortBoundApps(serviceInstance v2action.ServiceInstanceSummary) {
	sort.Slice(
		serviceInstance.BoundApplications,
		func(i, j int) bool {
			return sorting.LessIgnoreCase(serviceInstance.BoundApplications[i].AppName, serviceInstance.BoundApplications[j].AppName)
		})
}
