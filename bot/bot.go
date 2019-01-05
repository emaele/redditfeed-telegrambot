package bot

import (
	"fmt"
	"os"
	"strings"

	"github.com/emaele/redditfeed-telegrambot/utility"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/turnage/graw/reddit"

	conf "github.com/emaele/redditfeed-telegrambot/config"
	video "gitlab.com/emaele/telegram-videodl-bot/utility"
)

// PostingBot defines postbot structure
type PostingBot struct {
	RBot   reddit.Bot
	TBot   *tgbotapi.BotAPI
	Config conf.Config
}

// Post manages incoming reddit post
func (r *PostingBot) Post(p *reddit.Post) error {
	switch {
	case p.NSFW:
		// We hide NSFW content
		msg := tgbotapi.NewMessage(r.Config.ChatID, fmt.Sprintf("Uh oh, nsfw content! 🔞\n%s", p.URL))
		msg.DisableWebPagePreview = true
		msg.ReplyMarkup = utility.SetupInlineKeyboard(p.Subreddit, p.Permalink)
		r.TBot.Send(msg)
	case p.Media.RedditVideo.IsGIF:
		msg := tgbotapi.NewDocumentUpload(r.Config.ChatID, p.URL)
		msg.ReplyMarkup = utility.SetupInlineKeyboard(p.Subreddit, p.Permalink)
		r.TBot.Send(msg)
	case strings.Contains(p.URL, ".jpg") || strings.Contains(p.URL, ".png"):
		msg := tgbotapi.NewPhotoUpload(r.Config.ChatID, "")
		msg.FileID = p.URL
		msg.UseExisting = true
		msg.ReplyMarkup = utility.SetupInlineKeyboard(p.Subreddit, p.Permalink)
		r.TBot.Send(msg)
	default:
		if r.Config.VideoDownload {
			fileName, err := video.GetVideo(p.URL)
			if err != nil {
				fmt.Println(err)
			}
			videoPath := r.Config.DownloadPath + fileName

			msg := tgbotapi.NewVideoUpload(r.Config.ChatID, videoPath)
			msg.ReplyMarkup = utility.SetupInlineKeyboard(p.Subreddit, p.Permalink)

			r.TBot.Send(msg)
			os.Remove(videoPath)
		} else {
			msg := tgbotapi.NewMessage(r.Config.ChatID, p.URL)
			r.TBot.Send(msg)
		}
	}
	return nil
}
