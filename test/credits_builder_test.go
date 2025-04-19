package test

import (
	xmltv "github.com/chris102994/go-xmltv/pkg/xmltv/models"
	"testing"
)

func TestCreditsBuilder(t *testing.T) {
	director := "Director 1"
	actor := xmltv.Actor{Text: "Actor 1"}
	writer := "Writer 1"
	adapter := "Adapter 1"
	producer := "Producer 1"
	composer := "Composer 1"
	editor := "Editor 1"
	presenter := "Presenter 1"
	commentator := "Commentator 1"
	guest := "Guest 1"

	credits := xmltv.NewCreditsBuilder().
		AddDirector(director).
		AddActor(&actor).
		AddWriter(writer).
		AddAdapter(adapter).
		AddProducer(producer).
		AddComposer(composer).
		AddEditor(editor).
		AddPresenter(presenter).
		AddCommentator(commentator).
		AddGuest(guest).
		Build()

	if len(credits.Directors) != 1 {
		t.Errorf("Expected 1 director, got %d", len(credits.Directors))
	} else {
		if credits.Directors[0] != director {
			t.Errorf("Expected director %s, got %s", director, credits.Directors[0])
		}
	}

	if len(credits.Actors) != 1 {
		t.Errorf("Expected 1 actor, got %d", len(credits.Actors))
	} else {
		if credits.Actors[0].Text != actor.Text {
			t.Errorf("Expected actor %s, got %s", actor.Text, credits.Actors[0].Text)
		}
	}

	if len(credits.Writers) != 1 {
		t.Errorf("Expected 1 writer, got %d", len(credits.Writers))
	} else {
		if credits.Writers[0] != writer {
			t.Errorf("Expected writer %s, got %s", writer, credits.Writers[0])
		}
	}

	if len(credits.Adapters) != 1 {
		t.Errorf("Expected 1 adapter, got %d", len(credits.Adapters))
	} else {
		if credits.Adapters[0] != adapter {
			t.Errorf("Expected adapter %s, got %s", adapter, credits.Adapters[0])
		}
	}

	if len(credits.Producers) != 1 {
		t.Errorf("Expected 1 producer, got %d", len(credits.Producers))
	} else {
		if credits.Producers[0] != producer {
			t.Errorf("Expected producer %s, got %s", producer, credits.Producers[0])
		}
	}

	if len(credits.Composers) != 1 {
		t.Errorf("Expected 1 composer, got %d", len(credits.Composers))
	} else {
		if credits.Composers[0] != composer {
			t.Errorf("Expected composer %s, got %s", composer, credits.Composers[0])
		}
	}

	if len(credits.Editors) != 1 {
		t.Errorf("Expected 1 editor, got %d", len(credits.Editors))
	} else {
		if credits.Editors[0] != editor {
			t.Errorf("Expected editor %s, got %s", editor, credits.Editors[0])
		}
	}

	if len(credits.Presenters) != 1 {
		t.Errorf("Expected 1 presenter, got %d", len(credits.Presenters))
	} else {
		if credits.Presenters[0] != presenter {
			t.Errorf("Expected presenter %s, got %s", presenter, credits.Presenters[0])
		}
	}

	if len(credits.Commentators) != 1 {
		t.Errorf("Expected 1 commentator, got %d", len(credits.Commentators))
	} else {
		if credits.Commentators[0] != commentator {
			t.Errorf("Expected commentator %s, got %s", commentator, credits.Commentators[0])
		}
	}

	if len(credits.Guests) != 1 {
		t.Errorf("Expected 1 guest, got %d", len(credits.Guests))
	} else {
		if credits.Guests[0] != guest {
			t.Errorf("Expected guest %s, got %s", guest, credits.Guests[0])
		}
	}
}
