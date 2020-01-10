package commands

import (
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mmctl/client"
)

func getCommandFromCommandArg(c client.Client, commandArg string) *model.Command {
	// This may allow lookup via trigger in the future.
	cmd, resp := c.GetCommandById(commandArg)
	if resp != nil {
		return nil
	}
	return cmd
}
