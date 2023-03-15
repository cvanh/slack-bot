package main

import (
	"fmt"

	"github.com/slack-go/slack"
)

// handleSwitchPhoneCommand will take care of /redirect submissions and switches the phone
func handleSwitchPhoneCommand(command slack.SlashCommand, client *slack.Client) error {
	

	// The Input is found in the text field so
	// Create the attachment and assigned based on the message
	attachment := slack.Attachment{}
	// Add Some default context like user who mentioned the bot
	attachment.Fields = []slack.AttachmentField{
		{
			Title: "current phone",
			Value: "0623440513",
		}, {
			Title: "idk",
			Value: "idk",
		},
	}

	// Greet the user
	// attachment.Text = fmt.Sprintf("Hello %s", command.Text)
	attachment.Color = "#4af030"

	// Send the message to the channel
	// The Channel is available in the command.ChannelID
	_, _, err := client.PostMessage(command.ChannelID, slack.MsgOptionAttachments(attachment))
	if err != nil {
		return fmt.Errorf("failed to post message: %w", err)
	}
	return nil
}
