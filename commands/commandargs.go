package commands

import (
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mmctl/client"
)

func getCommandFromCommandArg(c client.Client, commandArg string) *model.Command {
	// This is here so we can more easily allow lookup via trigger in the future.
	cmd, resp := c.GetCommandById(commandArg)
	if resp.Error != nil {
		return nil
	}
	return cmd
}
