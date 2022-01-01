package simplyhentai

import (
	"net/url"
	"testing"

	"github.com/gan-of-culture/get-sauce/test"
)

func TestParseURL(t *testing.T) {
	tests := []struct {
		Name string
		URL  string
		Want int
	}{
		{
			Name: "Single Gallery doujin.sexy",
			URL:  "https://doujin.sexy/fate-grand-order/fdo-fatedosukebe-order-vol80",
			Want: 1,
		}, {
			Name: "Overview doujin.sexy",
			URL:  "https://doujin.sexy/character/gudao",
			Want: 24,
		}, {
			Name: "Single Gallery simply-hentai.com",
			URL:  "https://www.simply-hentai.com/1-kimetsu-no-yaiba/saimin-onsen-kanroji-mitsuri",
			Want: 1,
		}, {
			Name: "Overview simply-hentai.com",
			URL:  "https://www.simply-hentai.com/tag/big-breasts",
			Want: 24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			u, err := url.Parse(tt.URL)
			if err != nil {
				t.Error(err)
			}

			site = "https://" + u.Host + "/"

			URLs := parseURL(tt.URL)
			if len(URLs) > tt.Want || len(URLs) == 0 {
				t.Errorf("Got: %v - Want: %v", len(URLs), tt.Want)
			}
		})
	}
}

func TestExtract(t *testing.T) {
	tests := []struct {
		Name string
		Args test.Args
	}{
		{
			Name: "Single Gallery simply-hentai.com",
			Args: test.Args{
				URL:     "https://www.simply-hentai.com/original-work/torotoro-ni-shite-ageru-ch1-3",
				Title:   "Torotoro ni Shite Ageru Ch.1-3",
				Quality: "",
				Size:    0,
			},
		},
		{
			Name: "Single Gallery doujin.sexy",
			Args: test.Args{
				URL:     "https://doujin.sexy/fate-grand-order/fdo-fatedosukebe-order-vol80",
				Title:   "FDO Fate/Dosukebe Order VOL.8.0",
				Quality: "",
				Size:    0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			data, err := New().Extract(tt.Args.URL)
			test.CheckError(t, err)
			test.Check(t, tt.Args, data[0])
		})
	}
}
