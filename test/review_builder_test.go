package test

import (
	xmltv "github.com/chris102994/go-xmltv/pkg/xmltv/models"
	"testing"
)

func TestReviewBuilder(t *testing.T) {
	var reviewType = "positive"
	var source = "example.com"
	var reviewer = "John Doe"
	var lang = "en"
	var text = "This is a great show!"

	review := xmltv.NewReviewBuilder().
		SetType(reviewType).
		SetSource(source).
		SetReviewer(reviewer).
		SetLang(lang).
		SetText(text).
		Build()

	if review.Type != reviewType {
		t.Errorf("expected Type to be '%s', got '%s'", reviewType, review.Type)
	}
	if review.Source != source {
		t.Errorf("expected Source to be '%s', got '%s'", source, review.Source)
	}
	if review.Reviewer != reviewer {
		t.Errorf("expected Reviewer to be '%s', got '%s'", reviewer, review.Reviewer)
	}
	if review.Lang != lang {
		t.Errorf("expected Lang to be '%s', got '%s'", lang, review.Lang)
	}
	if review.Text != text {
		t.Errorf("expected Text to be '%s', got '%s'", text, review.Text)
	}
}
