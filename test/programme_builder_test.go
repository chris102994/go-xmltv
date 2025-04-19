package test

import (
	xmltv "github.com/chris102994/go-xmltv/pkg/xmltv/models"
	"testing"
)

func TestProgrammeBuilder(t *testing.T) {
	var start = "20231001000000 +0000"
	var stop = "20231001010000 +0000"
	var pdcStart = "20231001000000 +0000"
	var vpsStart = "20231001000000 +0000"
	var showview = "1234567"
	var videoplus = "1234567"
	var channel = "channel1"
	var clumpidx = "1"
	var title = &xmltv.Title{Text: "Example Title"}
	var subTitle = &xmltv.SubTitle{Text: "Example Subtitle"}
	var desc = &xmltv.Desc{Text: "Example Description"}
	var credits = &xmltv.Credits{Directors: []string{"Director 1", "Director 2"}}
	var date = "20231001"
	var category = &xmltv.Category{Text: "Example Category"}
	var keyword = &xmltv.Keyword{Text: "Example Keyword"}
	var language = "en"
	var origLanguage = "en"
	var length = &xmltv.Length{Units: "minutes", Text: "60"}
	var icon = &xmltv.Icon{Src: "http://example.com/icon.png"}
	var url = &xmltv.URL{Text: "http://example.com"}
	var country = &xmltv.Country{Text: "US"}
	var episodeNum = &xmltv.EpisodeNum{Text: "S01E01"}
	var video = &xmltv.Video{}
	var audio = &xmltv.Audio{}
	var previouslyShown = &xmltv.PreviouslyShown{}
	var premiere = "yes"
	var lastChance = &xmltv.LastChance{Text: "yes"}
	var new = "yes"
	var subtitles = &xmltv.Subtitle{Type: "teletext"}
	var rating = &xmltv.Rating{System: "MPAA", Value: "PG-13"}
	var starRating = &xmltv.StarRating{Value: "5"}
	var review = &xmltv.Review{Text: "Example Review"}
	var image = &xmltv.Image{Text: "http://example.com/image.png"}

	programme := xmltv.NewProgrammeBuilder().
		SetStart(start).
		SetStop(stop).
		SetPdcStart(pdcStart).
		SetVpsStart(vpsStart).
		SetShowview(showview).
		SetVideoplus(videoplus).
		SetChannel(channel).
		SetClumpidx(clumpidx).
		AddTitle(title).
		AddSubTitle(subTitle).
		AddDesc(desc).
		SetCredits(credits).
		SetDate(date).
		AddCategory(category).
		AddKeyword(keyword).
		SetLanguage(language).
		SetOrigLanguage(origLanguage).
		SetLength(length).
		AddIcon(icon).
		AddURL(url).
		AddCountry(country).
		AddEpisodeNum(episodeNum).
		SetVideo(video).
		SetAudio(audio).
		SetPreviouslyShown(previouslyShown).
		SetPremiere(premiere).
		AddLastChance(lastChance).
		SetNew(new).
		AddSubtitle(subtitles).
		AddRating(rating).
		AddStarRating(starRating).
		AddReview(review).
		AddImage(image).
		Build()

	if programme.Start != start {
		t.Errorf("expected Start to be '%s', got '%s'", start, programme.Start)
	}
	if programme.Stop != stop {
		t.Errorf("expected Stop to be '%s', got '%s'", stop, programme.Stop)
	}
	if programme.PdcStart != pdcStart {
		t.Errorf("expected PdcStart to be '%s', got '%s'", pdcStart, programme.PdcStart)
	}
	if programme.VpsStart != vpsStart {
		t.Errorf("expected VpsStart to be '%s', got '%s'", vpsStart, programme.VpsStart)
	}
	if programme.Showview != showview {
		t.Errorf("expected Showview to be '%s', got '%s'", showview, programme.Showview)
	}
	if programme.Videoplus != videoplus {
		t.Errorf("expected Videoplus to be '%s', got '%s'", videoplus, programme.Videoplus)
	}
	if programme.Channel != channel {
		t.Errorf("expected Channel to be '%s', got '%s'", channel, programme.Channel)
	}
	if programme.Clumpidx != clumpidx {
		t.Errorf("expected Clumpidx to be '%s', got '%s'", clumpidx, programme.Clumpidx)
	}
	if len(programme.Title) != 1 || programme.Title[0] != title {
		t.Errorf("expected Title to contain one title: '%v', got '%v'", title, programme.Title[0])
	}
	if len(programme.SubTitle) != 1 || programme.SubTitle[0] != subTitle {
		t.Errorf("expected SubTitle to contain one subtitle: '%v', got '%v'", subTitle, programme.SubTitle[0])
	}
	if len(programme.Desc) != 1 || programme.Desc[0] != desc {
		t.Errorf("expected Desc to contain one description: '%v', got '%v'", desc, programme.Desc[0])
	}
	if len(programme.Credits.Directors) != 2 || programme.Credits.Directors[0] != "Director 1" {
		t.Errorf("expected Credits to contain two directors: '%v', got '%v'", credits.Directors, programme.Credits.Directors)
	}
	if programme.Date != date {
		t.Errorf("expected Date to be '%s', got '%s'", date, programme.Date)
	}
	if len(programme.Category) != 1 || programme.Category[0] != category {
		t.Errorf("expected Category to contain one category: '%v', got '%v'", category, programme.Category[0])
	}
	if len(programme.Keyword) != 1 || programme.Keyword[0] != keyword {
		t.Errorf("expected Keyword to contain one keyword: '%v', got '%v'", keyword, programme.Keyword[0])
	}
	if programme.Language != language {
		t.Errorf("expected Language to be '%s', got '%s'", language, programme.Language)
	}
	if programme.OrigLanguage != origLanguage {
		t.Errorf("expected OrigLanguage to be '%s', got '%s'", origLanguage, programme.OrigLanguage)
	}
	if programme.Length.Units != "minutes" || programme.Length.Text != "60" {
		t.Errorf("expected Length to be '%v', got '%v'", length, programme.Length)
	}
	if len(programme.Icons) != 1 || programme.Icons[0] != icon {
		t.Errorf("expected Icons to contain one icon: '%v', got '%v'", icon, programme.Icons[0])
	}
	if len(programme.URLs) != 1 || programme.URLs[0] != url {
		t.Errorf("expected URLs to contain one URL: '%v', got '%v'", url, programme.URLs[0])
	}
	if len(programme.Country) != 1 || programme.Country[0] != country {
		t.Errorf("expected Country to contain one country: '%v', got '%v'", country, programme.Country[0])
	}
	if len(programme.EpisodeNum) != 1 || programme.EpisodeNum[0] != episodeNum {
		t.Errorf("expected EpisodeNum to contain one episode number: '%v', got '%v'", episodeNum, programme.EpisodeNum[0])
	}
	if programme.Video != video {
		t.Errorf("expected Video to be '%v', got '%v'", video, programme.Video)
	}
	if programme.Audio != audio {
		t.Errorf("expected Audio to be '%v', got '%v'", audio, programme.Audio)
	}
	if programme.PreviouslyShown != previouslyShown {
		t.Errorf("expected PreviouslyShown to be '%v', got '%v'", previouslyShown, programme.PreviouslyShown)
	}
	if programme.Premiere != premiere {
		t.Errorf("expected Premiere to be '%s', got '%s'", premiere, programme.Premiere)
	}
	if len(programme.LastChance) != 1 || programme.LastChance[0] != lastChance {
		t.Errorf("expected LastChance to contain one last chance: '%v', got '%v'", lastChance, programme.LastChance[0])
	}
	if programme.New != new {
		t.Errorf("expected New to be '%s', got '%s'", new, programme.New)
	}
	if len(programme.Subtitles) != 1 || programme.Subtitles[0] != subtitles {
		t.Errorf("expected Subtitles to contain one subtitle: '%v', got '%v'", subtitles, programme.Subtitles[0])
	}
	if len(programme.Rating) != 1 || programme.Rating[0] != rating {
		t.Errorf("expected Rating to contain one rating: '%v', got '%v'", rating, programme.Rating[0])
	}
	if len(programme.StarRating) != 1 || programme.StarRating[0] != starRating {
		t.Errorf("expected StarRating to contain one star rating: '%v', got '%v'", starRating, programme.StarRating[0])
	}
	if len(programme.Review) != 1 || programme.Review[0] != review {
		t.Errorf("expected Review to contain one review: '%v', got '%v'", review, programme.Review[0])
	}
	if len(programme.Image) != 1 || programme.Image[0] != image {
		t.Errorf("expected Image to contain one image: '%v', got '%v'", image, programme.Image[0])
	}
}
