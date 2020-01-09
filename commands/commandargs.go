package commands

import (
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mmctl/client"
)

func getCommandFromCommandArg(c client.Client, commandId string) *model.Command {
	cmd, resp := c.GetCommandById(commandId)
	if resp != nil {
		return nil
	}
	return cmd
}
