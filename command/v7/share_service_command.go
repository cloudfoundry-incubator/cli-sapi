package v7

import (
	"fmt"

	"code.cloudfoundry.org/cli/command/flag"
)

type ShareServiceCommand struct {
	BaseCommand

	RequiredArgs    flag.ShareServiceArgs `positional-args:"yes"`
	SpaceName       string                `short:"s" hidden:"true" required:"false" description:"Space to share the service instance into"`
	OrgName         string                `short:"o" required:"false" description:"Org of the other space (Default: targeted org)"`
	relatedCommands interface{}           `related_commands:"bind-service, service, services, unshare-service"`
}

func (cmd ShareServiceCommand) Usage() string {
	return "CF_NAME share-service SERVICE_INSTANCE OTHER_SPACE [-o OTHER_ORG]"
}

func (cmd ShareServiceCommand) Execute(args []string) error {
	space, err := cmd.parseSpaceName()
	if err != nil {
		return err
	}
	cmd.UI.DisplayText(space)

	if err := cmd.SharedActor.CheckTarget(true, true); err != nil {
		return err
	}

	return nil
}

func (cmd ShareServiceCommand) parseSpaceName() (string, error) {
	switch {
	case cmd.RequiredArgs.SpaceName != "" && cmd.SpaceName != "":
		return "", fmt.Errorf("you can't specify the space twice")
	case cmd.RequiredArgs.SpaceName != "":
		return cmd.RequiredArgs.SpaceName, nil
	case cmd.SpaceName != "":
		return cmd.SpaceName, nil
	default:
		return "", fmt.Errorf("space missing")
	}
}
