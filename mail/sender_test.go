package mail

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSendEmailWithGmail(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	sender := NewGmailSender("Admin", "1640.sad@gmail.com", "hxjccxravqupuyos")

	subject := "A test email"
	content := `
	<h1>Hello world</h1>
	<p>This is a test message from <a href="http://techschool.guru">Tech School</a></p>
	`
	to := []string{"yoakearsene@gmail.com"}
	attachFiles := []string{"../README.md"}

	err := sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	require.NoError(t, err)
}
