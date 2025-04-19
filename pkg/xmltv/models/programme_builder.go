package models

type ProgrammeBuilderInterface interface {
	NewProgrammeBuilder() *ProgrammeBuilder
	SetStart(start string) *ProgrammeBuilder
	SetStop(stop string) *ProgrammeBuilder
	SetPdcStart(pdcStart string) *ProgrammeBuilder
	SetVpsStart(vpsStart string) *ProgrammeBuilder
	SetShowview(showview string) *ProgrammeBuilder
	SetVideoplus(videoplus string) *ProgrammeBuilder
	SetChannel(channel string) *ProgrammeBuilder
	SetClumpidx(clumpidx string) *ProgrammeBuilder
	AddTitle(title *Title) *ProgrammeBuilder
	AddSubTitle(subTitle *SubTitle) *ProgrammeBuilder
	AddDesc(desc *Desc) *ProgrammeBuilder
	SetCredits(credits *Credits) *ProgrammeBuilder
	SetDate(date string) *ProgrammeBuilder
	AddCategory(category *Category) *ProgrammeBuilder
	AddKeyword(keyword *Keyword) *ProgrammeBuilder
	SetLanguage(language string) *ProgrammeBuilder
	SetOrigLanguage(origLanguage string) *ProgrammeBuilder
	SetLength(length *Length) *ProgrammeBuilder
	AddIcon(icon *Icon) *ProgrammeBuilder
	AddURL(url *URL) *ProgrammeBuilder
	AddCountry(country *Country) *ProgrammeBuilder
	AddEpisodeNum(episodeNum *EpisodeNum) *ProgrammeBuilder
	SetVideo(video *Video) *ProgrammeBuilder
	SetAudio(audio *Audio) *ProgrammeBuilder
	SetPreviouslyShown(previouslyShown *PreviouslyShown) *ProgrammeBuilder
	SetPremiere(premiere string) *ProgrammeBuilder
	AddLastChance(lastChance *LastChance) *ProgrammeBuilder
	SetNew(new string) *ProgrammeBuilder
	AddSubtitle(subtitle *Subtitle) *ProgrammeBuilder
	AddRating(rating *Rating) *ProgrammeBuilder
	AddStarRating(starRating *StarRating) *ProgrammeBuilder
	AddReview(review *Review) *ProgrammeBuilder
	AddImage(image *Image) *ProgrammeBuilder
	Build() *Programme
}

type ProgrammeBuilder struct {
	programme *Programme
}

func NewProgrammeBuilder() *ProgrammeBuilder {
	return &ProgrammeBuilder{
		programme: &Programme{
			Title:      make([]*Title, 0),
			SubTitle:   make([]*SubTitle, 0),
			Desc:       make([]*Desc, 0),
			Category:   make([]*Category, 0),
			Keyword:    make([]*Keyword, 0),
			Icons:      make([]*Icon, 0),
			URLs:       make([]*URL, 0),
			Country:    make([]*Country, 0),
			EpisodeNum: make([]*EpisodeNum, 0),
			LastChance: make([]*LastChance, 0),
			Subtitles:  make([]*Subtitle, 0),
			Rating:     make([]*Rating, 0),
			StarRating: make([]*StarRating, 0),
			Review:     make([]*Review, 0),
			Image:      make([]*Image, 0),
		},
	}
}

func (b *ProgrammeBuilder) SetStart(start string) *ProgrammeBuilder {
	b.programme.Start = start
	return b
}

func (b *ProgrammeBuilder) SetStop(stop string) *ProgrammeBuilder {
	b.programme.Stop = stop
	return b
}

func (b *ProgrammeBuilder) SetPdcStart(pdcStart string) *ProgrammeBuilder {
	b.programme.PdcStart = pdcStart
	return b
}

func (b *ProgrammeBuilder) SetVpsStart(vpsStart string) *ProgrammeBuilder {
	b.programme.VpsStart = vpsStart
	return b
}

func (b *ProgrammeBuilder) SetShowview(showview string) *ProgrammeBuilder {
	b.programme.Showview = showview
	return b
}

func (b *ProgrammeBuilder) SetVideoplus(videoplus string) *ProgrammeBuilder {
	b.programme.Videoplus = videoplus
	return b
}

func (b *ProgrammeBuilder) SetChannel(channel string) *ProgrammeBuilder {
	b.programme.Channel = channel
	return b
}

func (b *ProgrammeBuilder) SetClumpidx(clumpidx string) *ProgrammeBuilder {
	b.programme.Clumpidx = clumpidx
	return b
}

func (b *ProgrammeBuilder) AddTitle(title *Title) *ProgrammeBuilder {
	b.programme.Title = append(b.programme.Title, title)
	return b
}

func (b *ProgrammeBuilder) AddSubTitle(subTitle *SubTitle) *ProgrammeBuilder {
	b.programme.SubTitle = append(b.programme.SubTitle, subTitle)
	return b
}

func (b *ProgrammeBuilder) AddDesc(desc *Desc) *ProgrammeBuilder {
	b.programme.Desc = append(b.programme.Desc, desc)
	return b
}

func (b *ProgrammeBuilder) SetCredits(credits *Credits) *ProgrammeBuilder {
	b.programme.Credits = credits
	return b
}

func (b *ProgrammeBuilder) SetDate(date string) *ProgrammeBuilder {
	b.programme.Date = date
	return b
}

func (b *ProgrammeBuilder) AddCategory(category *Category) *ProgrammeBuilder {
	b.programme.Category = append(b.programme.Category, category)
	return b
}

func (b *ProgrammeBuilder) AddKeyword(keyword *Keyword) *ProgrammeBuilder {
	b.programme.Keyword = append(b.programme.Keyword, keyword)
	return b
}

func (b *ProgrammeBuilder) SetLanguage(language string) *ProgrammeBuilder {
	b.programme.Language = language
	return b
}

func (b *ProgrammeBuilder) SetOrigLanguage(origLanguage string) *ProgrammeBuilder {
	b.programme.OrigLanguage = origLanguage
	return b
}

func (b *ProgrammeBuilder) SetLength(length *Length) *ProgrammeBuilder {
	b.programme.Length = length
	return b
}

func (b *ProgrammeBuilder) AddIcon(icon *Icon) *ProgrammeBuilder {
	b.programme.Icons = append(b.programme.Icons, icon)
	return b
}

func (b *ProgrammeBuilder) AddURL(url *URL) *ProgrammeBuilder {
	b.programme.URLs = append(b.programme.URLs, url)
	return b
}

func (b *ProgrammeBuilder) AddCountry(country *Country) *ProgrammeBuilder {
	b.programme.Country = append(b.programme.Country, country)
	return b
}

func (b *ProgrammeBuilder) AddEpisodeNum(episodeNum *EpisodeNum) *ProgrammeBuilder {
	b.programme.EpisodeNum = append(b.programme.EpisodeNum, episodeNum)
	return b
}

func (b *ProgrammeBuilder) SetVideo(video *Video) *ProgrammeBuilder {
	b.programme.Video = video
	return b
}

func (b *ProgrammeBuilder) SetAudio(audio *Audio) *ProgrammeBuilder {
	b.programme.Audio = audio
	return b
}

func (b *ProgrammeBuilder) SetPreviouslyShown(previouslyShown *PreviouslyShown) *ProgrammeBuilder {
	b.programme.PreviouslyShown = previouslyShown
	return b
}

func (b *ProgrammeBuilder) SetPremiere(premiere string) *ProgrammeBuilder {
	b.programme.Premiere = premiere
	return b
}

func (b *ProgrammeBuilder) AddLastChance(lastChance *LastChance) *ProgrammeBuilder {
	b.programme.LastChance = append(b.programme.LastChance, lastChance)
	return b
}

func (b *ProgrammeBuilder) SetNew(new string) *ProgrammeBuilder {
	b.programme.New = new
	return b
}

func (b *ProgrammeBuilder) AddSubtitle(subtitle *Subtitle) *ProgrammeBuilder {
	b.programme.Subtitles = append(b.programme.Subtitles, subtitle)
	return b
}

func (b *ProgrammeBuilder) AddRating(rating *Rating) *ProgrammeBuilder {
	b.programme.Rating = append(b.programme.Rating, rating)
	return b
}

func (b *ProgrammeBuilder) AddStarRating(starRating *StarRating) *ProgrammeBuilder {
	b.programme.StarRating = append(b.programme.StarRating, starRating)
	return b
}

func (b *ProgrammeBuilder) AddReview(review *Review) *ProgrammeBuilder {
	b.programme.Review = append(b.programme.Review, review)
	return b
}

func (b *ProgrammeBuilder) AddImage(image *Image) *ProgrammeBuilder {
	b.programme.Image = append(b.programme.Image, image)
	return b
}

func (b *ProgrammeBuilder) Build() *Programme {
	return b.programme
}
