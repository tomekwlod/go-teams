package goteams

import (
	"errors"
	"strings"

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

	if m.title == "" {

		return errors.New("To send a message you need to provide a title")
	}
	if m.description == "" {

		return errors.New("To send a message you need to provide a description")
	}

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
		// msgCard.ThemeColor = "#E42828" // red - error
	}

	// send
	return mstClient.Send(webhookURL, msgCard)
}

func (m *message) SetTheme(hash string) {

	color := ""

	switch strings.ToLower(hash) {

	case "error":
	case "fail":
	case "red":

		color = "#E42828"
		break

	case "warning":
	case "orange":

		color = "#DF813D"
		break

	case "info":
	case "success":
	case "ok":
	case "green":

		color = "#56DF3D"
		break

	default:

		// maybe we should actually check if the hash is truly a hash?
		color = hash
		break
	}

	m.themeColor = color
}
