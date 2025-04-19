package main

import (
	"encoding/xml"
	xmltv "github.com/chris102994/go-xmltv/pkg/xmltv/models"
)

func main() {
	channel13 := xmltv.NewChannelBuilder().
		SetID("I10436.labs.zap2it.com").
		AddDisplayName(xmltv.NewDisplayNameBuilder().
			SetText("13 KERA").
			Build()).
		AddDisplayName(xmltv.NewDisplayNameBuilder().
			SetText("13 KERA TX42822:-").
			Build()).
		AddDisplayName(xmltv.NewDisplayNameBuilder().
			SetText("13").
			Build()).
		AddDisplayName(xmltv.NewDisplayNameBuilder().
			SetText("13 KERA fcc").
			Build()).
		AddDisplayName(xmltv.NewDisplayNameBuilder().
			SetText("KERA").
			Build()).
		AddDisplayName(xmltv.NewDisplayNameBuilder().
			SetText("KERA").
			Build()).
		AddDisplayName(xmltv.NewDisplayNameBuilder().
			SetText("PBS Affiliate").
			Build()).
		AddIcon(xmltv.NewIconBuilder().
			SetSrc("file://C:\\Perl\\site/share/xmltv/icons/KERA.gif").
			Build()).
		Build()

	channel11 := xmltv.NewChannelBuilder().
		SetID("I10759.labs.zap2it.com").
		AddDisplayName(xmltv.NewDisplayNameBuilder().
			SetText("11 KTVT").
			Build()).
		AddDisplayName(xmltv.NewDisplayNameBuilder().
			SetText("11 KTVT TX42822:-").
			Build()).
		AddDisplayName(xmltv.NewDisplayNameBuilder().
			SetText("11").
			Build()).
		AddDisplayName(xmltv.NewDisplayNameBuilder().
			SetText("11 KTVT fcc").
			Build()).
		AddDisplayName(xmltv.NewDisplayNameBuilder().
			SetText("KTVT").
			Build()).
		AddDisplayName(xmltv.NewDisplayNameBuilder().
			SetText("KTVT").
			Build()).
		AddDisplayName(xmltv.NewDisplayNameBuilder().
			SetText("CBS Affiliate").
			Build()).
		AddIcon(xmltv.NewIconBuilder().
			SetSrc("file://C:\\Perl\\site/share/xmltv/icons/KTVT.gif").
			Build()).
		Build()

	tv := xmltv.NewTVBuilder().
		SetSourceInfoURL("http://www.schedulesdirect.org").
		SetSourceInfoName("Schedules Direct").
		SetGeneratorInfoName("XMLTV/$Id: tv_grab_na_dd.in,v 1.70 2008/03/03 15:21:41 rmeden Exp $").
		SetGeneratorInfoURL("http://xmltv.org").
		AddChannel(channel13).
		AddChannel(channel11).
		AddProgramme(xmltv.NewProgrammeBuilder().
			SetStart("20080715003000 -0600").
			SetStop("20080715010000 -0600").
			SetChannel(channel13.ID).
			AddTitle(xmltv.NewTitleBuilder().
				SetLang("en").
				SetText("NOW on PBS").
				Build()).
			AddDesc(xmltv.NewDescBuilder().
				SetLang("en").
				SetText("Jordan's Queen Rania has made job creation a priority to help curb the staggering unemployment rates among youths in the Middle East.").
				Build()).
			SetDate("20080711").
			AddCategory(xmltv.NewCategoryBuilder().
				SetLang("en").
				SetText("Newsmagazine").
				Build()).
			AddCategory(xmltv.NewCategoryBuilder().
				SetLang("en").
				SetText("Interview").
				Build()).
			AddCategory(xmltv.NewCategoryBuilder().
				SetLang("en").
				SetText("Public affairs").
				Build()).
			AddCategory(xmltv.NewCategoryBuilder().
				SetLang("en").
				SetText("Series").
				Build()).
			AddEpisodeNum(xmltv.NewEpisodeNumBuilder().
				SetSystem("dd_progid").
				SetText("EP01006886.0028").
				Build()).
			AddEpisodeNum(xmltv.NewEpisodeNumBuilder().
				SetSystem("onscreen").
				SetText("427").
				Build()).
			SetAudio(xmltv.NewAudioBuilder().
				SetStereo("stereo").
				Build()).
			SetPreviouslyShown(xmltv.NewPreviouslyShownBuilder().
				SetStart("20080711000000").
				Build()).
			AddSubtitle(xmltv.NewSubtitleBuilder().
				SetType("teletext").
				Build()).
			Build()).
		AddProgramme(xmltv.NewProgrammeBuilder().
			SetStart("20080715010000 -0600").
			SetStop("20080715023000 -0600").
			SetChannel(channel13.ID).
			AddTitle(xmltv.NewTitleBuilder().
				SetLang("en").
				SetText("Mystery!").
				Build()).
			AddSubTitle(xmltv.NewSubTitleBuilder().
				SetLang("en").
				SetText("Foyle's War, Series IV: Bleak Midwinter").
				Build()).
			AddDesc(xmltv.NewDescBuilder().
				SetLang("en").
				SetText("Foyle investigates an explosion at a munitions factory, which he comes to believe may have been premeditated.").
				Build()).
			SetDate("20070701").
			AddCategory(xmltv.NewCategoryBuilder().
				SetLang("en").
				SetText("Anthology").
				Build()).
			AddCategory(xmltv.NewCategoryBuilder().
				SetLang("en").
				SetText("Mystery").
				Build()).
			AddCategory(xmltv.NewCategoryBuilder().
				SetLang("en").
				SetText("Series").
				Build()).
			AddEpisodeNum(xmltv.NewEpisodeNumBuilder().
				SetSystem("dd_progid").
				SetText("EP00003026.0665").
				Build()).
			AddEpisodeNum(xmltv.NewEpisodeNumBuilder().
				SetSystem("onscreen").
				SetText("2705").
				Build()).
			SetAudio(xmltv.NewAudioBuilder().
				SetStereo("stereo").
				Build()).
			SetPreviouslyShown(xmltv.NewPreviouslyShownBuilder().
				SetStart("20070701000000").
				Build()).
			AddSubtitle(xmltv.NewSubtitleBuilder().
				SetType("teletext").
				Build()).
			Build()).
		AddProgramme(xmltv.NewProgrammeBuilder().
			SetStart("20080715023000 -0600").
			SetStop("20080715040000 -0600").
			SetChannel(channel13.ID).
			AddTitle(xmltv.NewTitleBuilder().
				SetLang("en").
				SetText("Mystery!").
				Build()).
			AddSubTitle(xmltv.NewSubTitleBuilder().
				SetLang("en").
				SetText("Foyle's War, Series IV: Casualties of War").
				Build()).
			AddDesc(xmltv.NewDescBuilder().
				SetLang("en").
				SetText("The murder of a prominent scientist may have been due to a gambling debt.").
				Build()).
			SetDate("20070708").
			AddCategory(xmltv.NewCategoryBuilder().
				SetLang("en").
				SetText("Anthology").
				Build()).
			AddCategory(xmltv.NewCategoryBuilder().
				SetLang("en").
				SetText("Mystery").
				Build()).
			AddCategory(xmltv.NewCategoryBuilder().
				SetLang("en").
				SetText("Series").
				Build()).
			AddEpisodeNum(xmltv.NewEpisodeNumBuilder().
				SetSystem("dd_progid").
				SetText("EP00003026.0666").
				Build()).
			AddEpisodeNum(xmltv.NewEpisodeNumBuilder().
				SetSystem("onscreen").
				SetText("2706").
				Build()).
			SetAudio(xmltv.NewAudioBuilder().
				SetStereo("stereo").
				Build()).
			SetPreviouslyShown(xmltv.NewPreviouslyShownBuilder().
				SetStart("20070708000000").
				Build()).
			AddSubtitle(xmltv.NewSubtitleBuilder().
				SetType("teletext").
				Build()).
			Build()).
		AddProgramme(xmltv.NewProgrammeBuilder().
			SetStart("20080715040000 -0600").
			SetStop("20080715043000 -0600").
			SetChannel(channel13.ID).
			AddTitle(xmltv.NewTitleBuilder().
				SetLang("en").
				SetText("BBC World News").
				Build()).
			AddDesc(xmltv.NewDescBuilder().
				SetLang("en").
				SetText("International issues.").
				Build()).
			AddCategory(xmltv.NewCategoryBuilder().
				SetLang("en").
				SetText("News").
				Build()).
			AddCategory(xmltv.NewCategoryBuilder().
				SetLang("en").
				SetText("Series").
				Build()).
			AddEpisodeNum(xmltv.NewEpisodeNumBuilder().
				SetSystem("dd_progid").
				SetText("SH00315789.0000").
				Build()).
			AddSubtitle(xmltv.NewSubtitleBuilder().
				SetType("teletext").
				Build()).
			Build()).
		AddProgramme(xmltv.NewProgrammeBuilder().
			SetStart("20080715060000 -0600").
			SetStop("20080715080000 -0600").
			SetChannel(channel11.ID).
			AddTitle(xmltv.NewTitleBuilder().
				SetLang("en").
				SetText("The Early Show").
				Build()).
			AddDesc(xmltv.NewDescBuilder().
				SetLang("en").
				SetText("Republican candidate John McCain; premiere of the film \"The Dark Knight.\"").
				Build()).
			SetDate("20080715").
			AddCategory(xmltv.NewCategoryBuilder().
				SetLang("en").
				SetText("Talk").
				Build()).
			AddCategory(xmltv.NewCategoryBuilder().
				SetLang("en").
				SetText("News").
				Build()).
			AddCategory(xmltv.NewCategoryBuilder().
				SetLang("en").
				SetText("Series").
				Build()).
			AddEpisodeNum(xmltv.NewEpisodeNumBuilder().
				SetSystem("dd_progid").
				SetText("EP00337003.2361").
				Build()).
			SetAudio(xmltv.NewAudioBuilder().
				SetStereo("stereo").
				Build()).
			AddSubtitle(xmltv.NewSubtitleBuilder().
				SetType("teletext").
				Build()).
			Build()).
		AddProgramme(xmltv.NewProgrammeBuilder().
			SetStart("20080715080000 -0600").
			SetStop("20080715090000 -0600").
			SetChannel(channel11.ID).
			AddTitle(xmltv.NewTitleBuilder().
				SetLang("en").
				SetText("Rachael Ray").
				Build()).
			AddDesc(xmltv.NewDescBuilder().
				SetLang("en").
				SetText("Actresses Kim Raver, Brooke Shields and Lindsay Price (\"Lipstick Jungle\"); women in their 40s tell why they got breast implants; a 30-minute meal.").
				Build()).
			SetDate("20080306").
			AddCategory(xmltv.NewCategoryBuilder().
				SetLang("en").
				SetText("Talk").
				Build()).
			AddCategory(xmltv.NewCategoryBuilder().
				SetLang("en").
				SetText("Series").
				Build()).
			AddEpisodeNum(xmltv.NewEpisodeNumBuilder().
				SetSystem("dd_progid").
				SetText("EP00847333.0303").
				Build()).
			AddEpisodeNum(xmltv.NewEpisodeNumBuilder().
				SetSystem("onscreen").
				SetText("2119").
				Build()).
			SetAudio(xmltv.NewAudioBuilder().
				SetStereo("stereo").
				Build()).
			SetPreviouslyShown(xmltv.NewPreviouslyShownBuilder().
				SetStart("20080306000000").
				Build()).
			AddSubtitle(xmltv.NewSubtitleBuilder().
				SetType("teletext").
				Build()).
			Build()).
		AddProgramme(xmltv.NewProgrammeBuilder().
			SetStart("20080715090000 -0600").
			SetStop("20080715100000 -0600").
			SetChannel(channel11.ID).
			AddTitle(xmltv.NewTitleBuilder().
				SetLang("en").
				SetText("The Price Is Right").
				Build()).
			AddDesc(xmltv.NewDescBuilder().
				SetLang("en").
				SetText("Contestants bid for prizes then compete for fabulous showcases.").
				Build()).
			SetCredits(xmltv.NewCreditsBuilder().
				AddDirector("Bart Eskander").
				AddProducer("Roger Dobkowitz").
				AddPresenter("Drew Carey").
				Build()).
			AddCategory(xmltv.NewCategoryBuilder().
				SetLang("en").
				SetText("Game show").
				Build()).
			AddCategory(xmltv.NewCategoryBuilder().
				SetLang("en").
				SetText("Series").
				Build()).
			AddEpisodeNum(xmltv.NewEpisodeNumBuilder().
				SetSystem("dd_progid").
				SetText("SH00004372.0000").
				Build()).
			SetAudio(xmltv.NewAudioBuilder().
				SetStereo("stereo").
				Build()).
			AddSubtitle(xmltv.NewSubtitleBuilder().
				SetType("teletext").
				Build()).
			AddRating(xmltv.NewRatingBuilder().
				SetSystem("VCHIP").
				SetValue("TV-G").
				Build()).
			Build()).
		AddProgramme(xmltv.NewProgrammeBuilder().
			SetStart("20080715100000 -0600").
			SetStop("20080715103000 -0600").
			SetChannel(channel11.ID).
			AddTitle(xmltv.NewTitleBuilder().
				SetLang("en").
				SetText("Jeopardy!").
				Build()).
			SetCredits(xmltv.NewCreditsBuilder().
				AddPresenter("Alex Trebek").
				Build()).
			SetDate("20080715").
			AddCategory(xmltv.NewCategoryBuilder().
				SetLang("en").
				SetText("Game show").
				Build()).
			AddCategory(xmltv.NewCategoryBuilder().
				SetLang("en").
				SetText("Series").
				Build()).
			AddEpisodeNum(xmltv.NewEpisodeNumBuilder().
				SetSystem("dd_progid").
				SetText("EP00002348.1700").
				Build()).
			AddEpisodeNum(xmltv.NewEpisodeNumBuilder().
				SetSystem("onscreen").
				SetText("5507").
				Build()).
			AddSubtitle(xmltv.NewSubtitleBuilder().
				SetType("teletext").
				Build()).
			AddRating(xmltv.NewRatingBuilder().
				SetSystem("VCHIP").
				SetValue("TV-G").
				Build()).
			Build()).
		AddProgramme(xmltv.NewProgrammeBuilder().
			SetStart("20080715103000 -0600").
			SetStop("20080715113000 -0600").
			SetChannel(channel11.ID).
			AddTitle(xmltv.NewTitleBuilder().
				SetLang("en").
				SetText("The Young and the Restless").
				Build()).
			AddSubTitle(xmltv.NewSubTitleBuilder().
				SetLang("en").
				SetText("Sabrina Offers Victoria a Truce").
				Build()).
			AddDesc(xmltv.NewDescBuilder().
				SetLang("en").
				SetText("Jeff thinks Kyon stole the face cream; Nikki asks Nick to give David a chance; Amber begs Adrian to go to Australia.").
				Build()).
			SetCredits(xmltv.NewCreditsBuilder().
				AddActor(xmltv.NewActorBuilder().
					SetText("Peter Bergman").
					Build()).
				AddActor(xmltv.NewActorBuilder().
					SetText("Eric Braeden").
					Build()).
				AddActor(xmltv.NewActorBuilder().
					SetText("Jeanne Cooper").
					Build()).
				AddActor(xmltv.NewActorBuilder().
					SetText("Melody Thomas Scott").
					Build()).
				Build()).
			SetDate("20080715").
			AddCategory(xmltv.NewCategoryBuilder().
				SetLang("en").
				SetText("Soap").
				Build()).
			AddCategory(xmltv.NewCategoryBuilder().
				SetLang("en").
				SetText("Series").
				Build()).
			AddEpisodeNum(xmltv.NewEpisodeNumBuilder().
				SetSystem("dd_progid").
				SetText("EP00004422.1359").
				Build()).
			AddEpisodeNum(xmltv.NewEpisodeNumBuilder().
				SetSystem("onscreen").
				SetText("8937").
				Build()).
			SetAudio(xmltv.NewAudioBuilder().
				SetStereo("stereo").
				Build()).
			AddSubtitle(xmltv.NewSubtitleBuilder().
				SetType("teletext").
				Build()).
			AddRating(xmltv.NewRatingBuilder().
				SetSystem("VCHIP").
				SetValue("TV-14").
				Build()).
			Build()).
		Build()

	output, err := xml.MarshalIndent(tv, "", "    ")
	if err != nil {
		panic(err)
	}
	println(string(output))
}
