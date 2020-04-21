package app

import (
	"math/rand"
	"reflect"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Url is domain struct for app layer
type Url struct {
	Base  string
	Short string
}

type UrlServicer interface {
	Get(key string) (string, error)
	Create(url string) string
}

//go:generate mockery -name UrlRepository -case underscore -inpkg
// UrlRepository interface for working with interfaces
type UrlRepository interface {
	Save(key, url string) error
	Get(key string) (string, error)
}

//go:generate mockery -name Generator -case underscore -inpkg
// Generator interface for generating a random strings
type Generator interface {
	Generate() string
}

// NewUrlService creates instance that allows work with app cases
func NewUrlService(repo UrlRepository, generator Generator, domain string) UrlService {
	if repo == nil || reflect.ValueOf(repo).IsNil() {
		panic("repo param is nil")
	}
	if generator == nil || reflect.ValueOf(generator).IsNil() {
		panic("generator param is nil")
	}
	if domain == "" {
		panic("domain param is empty")
	}

	return UrlService{
		repo:      repo,
		generator: generator,
		domain:    domain,
	}
}

// UrlService implements UrlRepository interface
type UrlService struct {
	repo      UrlRepository
	generator Generator
	domain    string
}

// Read reads base originUrl by short originUrl
func (us UrlService) Get(shortUrl string) (string, error) {
	return us.repo.Get(shortUrl)
}

// Create creates short originUrl by base originUrl
func (us UrlService) Create(originUrl string) string {
	for {
		shortUrl := us.generator.Generate()

		err := us.repo.Save(shortUrl, originUrl)
		if err == nil {
			return us.domain + shortUrl
		}

		url, _ := us.repo.Get(shortUrl)
		if url == originUrl {
			return us.domain + shortUrl
		}
	}
}

// UrlGenerator implements Generator interface
type UrlGenerator struct {
	pool       *sync.Pool
	characters string
}

// NewUrlGenerator creates new UrlGenerator instance
func NewUrlGenerator(characters string, bufSize int) UrlGenerator {
	return UrlGenerator{
		characters: characters,
		pool: func() *sync.Pool {
			return &sync.Pool{
				New: func() interface{} {
					return make([]byte, bufSize)
				},
			}
		}(),
	}
}

// Generate generates new random string
func (ug UrlGenerator) Generate() string {
	buf := ug.pool.Get().([]byte)

	for i := range buf {
		buf[i] = ug.characters[rand.Intn(len(ug.characters))]
	}

	r := string(buf)

	ug.pool.Put(buf)

	return r
}
