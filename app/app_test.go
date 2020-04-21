package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUrlService(t *testing.T) {
	// arrange
	repo := &MockUrlRepository{}
	gen := &MockGenerator{}
	domain := "localhost"

	wantValue := UrlService{
		domain:    "localhost",
		repo:      repo,
		generator: gen,
	}

	// act
	got := NewUrlService(repo, gen, domain)

	// assert
	assert.NotNil(t, got)
	assert.Equal(t, wantValue, got)
}

func TestNewUrlServiceWithPacic(t *testing.T) {
	// arrange
	cases := []struct {
		testName  string
		domain    string
		generator *MockGenerator
		repo      *MockUrlRepository
		wantMsg   string
	}{
		{
			testName:  "Should call panic if repo is nil",
			domain:    "localhost",
			generator: &MockGenerator{},
			wantMsg:   "repo param is nil",
		},
		{
			testName: "Should call panic if generator is nil",
			domain:   "localhost",
			repo:     &MockUrlRepository{},
			wantMsg:  "generator param is nil",
		},
		{
			testName:  "Should call panic if domain is empty",
			repo:      &MockUrlRepository{},
			generator: &MockGenerator{},
			wantMsg:   "domain param is empty",
		},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {

			// assert
			assert.PanicsWithValue(t, c.wantMsg, func() {
				// act
				NewUrlService(c.repo, c.generator, c.domain)
			})
		})
	}
}

func TestUrlService_Get(t *testing.T) {
	// arrange
	cases := []struct {
		testName string
		repo     *MockUrlRepository
		wantErr  error
		wantUrl  string
		url      string
	}{
		{
			testName: "Should return url if all is ok",
			wantUrl:  "http://google.com",
			url:      "asd123",
			repo: func() *MockUrlRepository {
				m := &MockUrlRepository{}
				m.On("Get", "asd123").
					Return("http://google.com", nil)
				return m
			}(),
		},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			// act
			s := NewUrlService(c.repo, &MockGenerator{}, "localhost")
			gotUrl, gotErr := s.Get(c.url)

			// assert
			assert.Equal(t, c.wantErr, gotErr)
			assert.Equal(t, c.wantUrl, gotUrl)
		})
	}
}
