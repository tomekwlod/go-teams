package goteams

import (
	goteamsnotify "github.com/atc0005/go-teams-notify/v2"
)

type message struct {
	webhookURL  string
	title       string
	description string
	themeColor  string
}

func New(webhookURL, title, description string) message {

	m := message{webhookURL, title, description, "#56DF3D"} // default themeColor

	return m
}

func (m message) Send() error {

	// init the client
	mstClient := goteamsnotify.NewClient()

	// setup webhook url
	webhookURL := m.webhookURL

	// Disable webhook URL validation
	mstClient.SkipWebhookURLValidationOnSend(true)

	// setup message card
	msgCard := goteamsnotify.NewMessageCard()

	msgCard.Title = m.title
	msgCard.Text = m.description

	// msgCard.Text = "Here are some examples of formatted stuff like " +
	// 	"<br> * this list itself  <br> * **bold** <br> * *italic* <br> * ***bolditalic***"
	// msgCard.Text = "This is a **body** of this message (I will add something useful later)"
	msgCard.ThemeColor = m.themeColor

	if m.themeColor == "" {

		msgCard.ThemeColor = "#56DF3D" // green  - info
		// msgCard.ThemeColor = "#DF813D" // orange - warning
	}

	// send
	return mstClient.Send(webhookURL, msgCard)
}

func (m *message) SetTheme(hash string) {

	//
	// validation here!!
	//
	// also not codes but types: error,success,info,warning,itp
	//
	//

	m.themeColor = hash
}
