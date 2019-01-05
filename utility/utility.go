package utility

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// Thank you BomBonio!

// SetupInlineKeyboard is an utility function that makes it easier to build an appropriate set of buttons for reports
func SetupInlineKeyboard(subreddit string, permalink string) (keyboard tgbotapi.InlineKeyboardMarkup) {

	//The first button to show is the link to the reported message
	if subreddit != "" {
		row := tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonURL("/r/"+subreddit, "reddit.com"+permalink))
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row)
	}

	//We finally append the lower row to the keyboard
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard)
	return
}
