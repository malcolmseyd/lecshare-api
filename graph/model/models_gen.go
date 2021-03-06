// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Class struct {
	Title      string     `json:"title"`
	Code       string     `json:"code"`
	Instructor *User      `json:"instructor"`
	Lectures   []*Lecture `json:"lectures"`
}

type Lecture struct {
	Name          string         `json:"name"`
	Datetime      string         `json:"datetime"`
	Duration      int            `json:"duration"`
	Transcription *Transcription `json:"transcription"`
}

type Resource struct {
	ContentType *string `json:"contentType"`
	URL         *string `json:"url"`
}

type Transcription struct {
	Sections []*TranscriptionSection `json:"sections"`
}

type TranscriptionParagraph struct {
	Transcript string               `json:"transcript"`
	Confidence float64              `json:"confidence"`
	Words      []*TranscriptionWord `json:"words"`
}

type TranscriptionSection struct {
	Alternatives []*TranscriptionParagraph `json:"alternatives"`
	LanguageCode string                    `json:"languageCode"`
}

type TranscriptionWord struct {
	Starttime  *WordTime `json:"starttime"`
	Endtime    *WordTime `json:"endtime"`
	Word       string    `json:"word"`
	Confidence float64   `json:"confidence"`
}

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Suffix    string `json:"suffix"`
	Role      string `json:"role"`
}

type WordTime struct {
	Seconds *string `json:"seconds"`
	Nanos   *int    `json:"nanos"`
}
