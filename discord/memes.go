package discord

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

const subreddit = "memes"

type redditResp struct {
	Data redditData
}

type redditData struct {
	Children []previewData
}

type previewData struct {
	Data meme
}

type meme struct {
	URL       string
	Title     string
	Permalink string
}

func (b *Bot) memes(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "meme") {
		if b.memerefreshtime == 0 { // first time since startup
			s.ChannelMessageSend(m.ChannelID, "Sorry, this is the first time someone has asked for a meme since the server started. It may take a moment for us to download the memes from reddit.")
			b.loadMemes(m)
		}
		if (time.Now().UnixNano() - b.memerefreshtime) > 3600 { // its been an hour
			go b.loadMemes(m)
		}

		// send message
		meme := b.memedat[rand.Intn(len(b.memedat))]
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			URL:   meme.Permalink,
			Type:  discordgo.EmbedTypeImage,
			Title: meme.Title,
			Image: &discordgo.MessageEmbedImage{
				URL: meme.URL,
			},
		})
	}
}

func (b *Bot) loadMemes(m *discordgo.MessageCreate) bool {
	b.memerefreshtime = time.Now().UnixNano()

	// Download
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://reddit.com/r/"+subreddit+"/hot.json", nil)
	if b.handle(err, m) {
		return false
	}
	req.Header.Set("User-Agent", "Nv7 Bot")
	res, err := client.Do(req)
	if b.handle(err, m) {
		return false
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if b.handle(err, m) {
		return false
	}

	// Process
	var dat redditResp
	err = json.Unmarshal(data, &dat)
	if b.handle(err, m) {
		return false
	}

	b.memedat = make([]meme, len(dat.Data.Children))
	for i, val := range dat.Data.Children {
		b.memedat[i] = val.Data
	}
	return true
}
