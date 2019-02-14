package v6

import (
	"fmt"

	"code.cloudfoundry.org/cli/actor/sharedaction"
	"code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/command"
	"code.cloudfoundry.org/cli/command/flag"
	"code.cloudfoundry.org/cli/shim"
)

type UpdateServiceActor interface {
	GetServiceInstanceSummaryByNameAndSpace(name string, spaceGUID string) (v2action.ServiceInstanceSummary, v2action.Warnings, error)
	UpdateDone(name string)
}

type UpdateServiceCommand struct {
	RequiredArgs     flag.ServiceInstance `positional-args:"yes"`
	ParametersAsJSON flag.Path            `short:"c" description:"Valid JSON object containing service-specific configuration parameters, provided either in-line or in a file. For a list of supported configuration parameters, see documentation for the particular service offering."`
	Plan             string               `short:"p" description:"Change service plan for a service instance"`
	Tags             string               `short:"t" description:"User provided tags"`
	usage            interface{}          `usage:"CF_NAME update-service SERVICE_INSTANCE [-p NEW_PLAN] [-c PARAMETERS_AS_JSON] [-t TAGS]\n\n   Optionally provide service-specific configuration parameters in a valid JSON object in-line.\n   CF_NAME update-service -c '{\"name\":\"value\",\"name\":\"value\"}'\n\n   Optionally provide a file containing service-specific configuration parameters in a valid JSON object. \n   The path to the parameters file can be an absolute or relative path to a file.\n   CF_NAME update-service -c PATH_TO_FILE\n\n   Example of valid JSON object:\n   {\n      \"cluster_nodes\": {\n         \"count\": 5,\n         \"memory_mb\": 1024\n      }\n   }\n\n   Optionally provide a list of comma-delimited tags that will be written to the VCAP_SERVICES environment variable for any bound applications.\n\nEXAMPLES:\n   CF_NAME update-service mydb -p gold\n   CF_NAME update-service mydb -c '{\"ram_gb\":4}'\n   CF_NAME update-service mydb -c ~/workspace/tmp/instance_config.json\n   CF_NAME update-service mydb -t \"list, of, tags\""`
	relatedCommands  interface{}          `related_commands:"rename-service, services, update-user-provided-service"`

	UI          command.UI
	Config      command.Config
	SharedActor command.SharedActor
	Actor       UpdateServiceActor
}

func (cmd *UpdateServiceCommand) Setup(config command.Config, ui command.UI) error {
	cmd.UI = ui
	cmd.Config = config
	cmd.SharedActor = sharedaction.NewActor(config)
	cmd.Actor = shim.Actor{}

	return nil
}

func (cmd *UpdateServiceCommand) Execute(args []string) error {
	name := cmd.RequiredArgs.ServiceInstance
	serviceInstanceSummary, warnings, err := cmd.Actor.GetServiceInstanceSummaryByNameAndSpace(cmd.RequiredArgs.ServiceInstance, cmd.Config.TargetedSpace().GUID)
	cmd.UI.DisplayWarnings(warnings)
	if err != nil {
		return err
	}

	if !serviceInstanceSummary.UpdateAvailable {
		return fmt.Errorf("Service %s doesn't need to be updated", name)
	}

	cmd.UI.DisplayNewline()
	cmd.UI.DisplayText(fmt.Sprintf("You are about to update %s.", name))

	cmd.UI.DisplayNewline()
	cmd.UI.DisplayText("The following changes will be applied:")
	cmd.UI.DisplayText(serviceInstanceSummary.UpdateDetails)

	cmd.UI.DisplayNewline()
	cmd.UI.DisplayText("Warning: This operation may be long running and will block further operations on the service until complete")

	cmd.UI.DisplayNewline()
	response, promptErr := cmd.UI.DisplayBoolPrompt(false, "Really update service `{{.Name}}`", map[string]interface{}{"Name": name})
	if promptErr != nil {
		return promptErr
	}
	if !response {
		return nil
	}

	cmd.UI.DisplayNewline()
	cmd.UI.DisplayText("Updating service instance {{.Name}} as {{.User}}...", map[string]interface{}{
		"Name": name,
		"User": "developer",
	})
	cmd.UI.DisplayText("OK")

	cmd.Actor.UpdateDone(name)

	return nil
}
