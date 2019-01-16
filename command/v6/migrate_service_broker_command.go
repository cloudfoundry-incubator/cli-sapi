package v6

import (
	"code.cloudfoundry.org/cli/actor/sharedaction"
	"code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/command"
	"code.cloudfoundry.org/cli/command/flag"
	"code.cloudfoundry.org/cli/command/v6/shared"
)

//go:generate counterfeiter . CreateServiceBrokerActor

type MigrateServiceBrokerActor interface {
	MigrateServiceBrokerByName(serviceBrokerName string)error
}

type MigrateServiceBrokerCommand struct {
	RequiredArgs    flag.MigrateServiceBrokerArgs `positional-args:"yes"`
	usage           interface{}            `usage:"CF_NAME migrate-service-broker SERVICE_BROKER"`
	relatedCommands interface{}            `related_commands:"service-brokers, target"`

	UI          command.UI
	Config      command.Config
	SharedActor command.SharedActor
	Actor       MigrateServiceBrokerActor
}

func (cmd *MigrateServiceBrokerCommand) Setup(config command.Config, ui command.UI) error {
	cmd.UI = ui
	cmd.Config = config

	cmd.SharedActor = sharedaction.NewActor(config)

	ccClient, uaaClient, err := shared.NewClients(config, ui, true)
	if err != nil {
		return err
	}

	cmd.Actor = v2action.NewActor(ccClient, uaaClient, config)

	return nil
}

func (cmd MigrateServiceBrokerCommand) Execute(args []string) error {
  if err := cmd.SharedActor.CheckTarget(false, false); err != nil {
	  return err
	}

	user, err := cmd.Config.CurrentUser()
	if err != nil {
		return err
	}

		cmd.UI.DisplayTextWithFlavor("Migrating service broker {{.ServiceBrokerName}} as {{.User}}...",
			map[string]interface{}{
				"ServiceBrokerName": cmd.RequiredArgs.ServiceBroker,
				"User":              user.Name,
			})

	err = cmd.Actor.MigrateServiceBrokerByName(cmd.RequiredArgs.ServiceBroker)
	if err != nil {
		return err
	}

	cmd.UI.DisplayOK()
	return nil
}
