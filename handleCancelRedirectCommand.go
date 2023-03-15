package main

import (
	"fmt"

	"github.com/slack-go/slack"
)

// handleCancelRedirectCommand will take care of /cancel_redirect submissions and turns off the trough switching
func handleCancelRedirectCommand(command slack.SlashCommand, client *slack.Client) error {
	// The Input is found in the text field so
	// Create the attachment and assigned based on the message
	attachment := slack.Attachment{}

	// Greet the user
	attachment.Text = fmt.Sprintf("hi, this command doesnt exist yet. My creator is working on it")
	attachment.Color = "#4af030"

	// Send the message to the channel
	// The Channel is available in the command.ChannelID
	_, _, err := client.PostMessage(command.ChannelID, slack.MsgOptionAttachments(attachment))
	if err != nil {
		return fmt.Errorf("failed to post message: %w", err)
	}
	return nil
}
