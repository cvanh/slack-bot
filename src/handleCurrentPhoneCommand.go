package main

import (
	"fmt"

	"github.com/slack-go/slack"
)

// handleCurrentPhoneCommand will take care of /current_redirected_phone submissions and returns the phoen that is curently switched trough
func handleCurrentPhoneCommand(command slack.SlashCommand, client *slack.Client) error {
	attachment := slack.Attachment{}

	// Greet the user
	attachment.Text = fmt.Sprintf("switching: %s till 1700pm today", command.Text)
	attachment.Color = "#4af030"

	// Send the message to the channel
	_, _, err := client.PostMessage(command.ChannelID, slack.MsgOptionAttachments(attachment))
	if err != nil {
		return fmt.Errorf("failed to post message: %w", err)
	}
	return nil
}
