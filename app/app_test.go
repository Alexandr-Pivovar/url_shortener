package app

import (
	"errors"
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
			testName: "Should return originUrl if all is ok",
			wantUrl:  "http://google.com",
			url:      "asd123",
			repo: func() *MockUrlRepository {
				m := &MockUrlRepository{}
				m.On("Get", "asd123").
					Return("http://google.com", nil).Once()
				return m
			}(),
		},
		{
			testName: "Should return error if originUrl does not exist",
			wantUrl:  "",
			url:      "asd123",
			repo: func() *MockUrlRepository {
				m := &MockUrlRepository{}
				m.On("Get", "asd123").
					Return("", errors.New("key does not exist")).
					Once()
				return m
			}(),
			wantErr: errors.New("key does not exist"),
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

			c.repo.AssertExpectations(t)
		})
	}
}

func TestNewUrlService_Create(t *testing.T) {
	// arrange
	cases := []struct {
		testName  string
		repo      *MockUrlRepository
		gen       *MockGenerator
		wantUrl   string
		originUrl string
	}{
		{
			testName:  "Should return short url if repo does not return error",
			wantUrl:   "localhost/asd123",
			originUrl: "http://google.com",
			gen: func() *MockGenerator {
				m := &MockGenerator{}
				m.On("Generate").Return("asd123")
				return m
			}(),
			repo: func() *MockUrlRepository {
				m := &MockUrlRepository{}
				m.On("Save", "asd123", "http://google.com").
					Return(nil).Once()
				return m
			}(),
		},
		{
			testName:  "Should return short url if repo.Save returns error",
			wantUrl:   "localhost/asd123",
			originUrl: "http://google.com",
			repo: func() *MockUrlRepository {
				m := &MockUrlRepository{}
				m.On("Save", "asd123", "http://google.com").
					Return(errors.New("key does not exist")).
					Once()
				m.On("Get", "asd123").
					Return("http://google.com", nil).
					Once()
				return m
			}(),
			gen: func() *MockGenerator {
				m := &MockGenerator{}
				m.On("Generate").Return("asd123")
				return m
			}(),
		},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {

			// act
			s := NewUrlService(c.repo, c.gen, "localhost/")
			gotUrl := s.Create(c.originUrl)

			// assert
			assert.Equal(t, c.wantUrl, gotUrl)

			c.repo.AssertExpectations(t)
			c.gen.AssertExpectations(t)
		})
	}
}
