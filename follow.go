package twitterscraper

import (
	"io"
	"net/url"
	"strings"
)

func (s *Scraper) Follow(userID string) error {
	req, err := s.newRequest("POST", "https://twitter.com/i/api/1.1/friendships/create.json")
	if err != nil {
		return err
	}

	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	// Prepare URL-encoded form data
	data := url.Values{
		"include_profile_interstitial_type": {"1"},
		"include_blocking":                  {"1"},
		"include_blocked_by":                {"1"},
		"include_followed_by":               {"1"},
		"include_want_retweets":             {"1"},
		"include_mute_edge":                 {"1"},
		"include_can_dm":                    {"1"},
		"include_can_media_tag":             {"1"},
		"include_ext_is_blue_verified":      {"1"},
		"include_ext_verified_type":         {"1"},
		"include_ext_profile_image_shape":   {"1"},
		"skip_status":                       {"1"},
		"user_id":                           {userID},
	}

	req.Body = io.NopCloser(strings.NewReader(data.Encode()))
	req.ContentLength = int64(len(data.Encode()))

	var response interface{}
	return s.RequestAPI(req, &response)
}

func (s *Scraper) Unfollow(userID string) error {
	req, err := s.newRequest("POST", "https://twitter.com/i/api/1.1/friendships/destroy.json")
	if err != nil {
		return err
	}

	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	// Prepare URL-encoded form data
	data := url.Values{
		"include_profile_interstitial_type": {"1"},
		"include_blocking":                  {"1"},
		"include_blocked_by":                {"1"},
		"include_followed_by":               {"1"},
		"include_want_retweets":             {"1"},
		"include_mute_edge":                 {"1"},
		"include_can_dm":                    {"1"},
		"include_can_media_tag":             {"1"},
		"include_ext_is_blue_verified":      {"1"},
		"include_ext_verified_type":         {"1"},
		"include_ext_profile_image_shape":   {"1"},
		"skip_status":                       {"1"},
		"user_id":                           {userID},
	}

	req.Body = io.NopCloser(strings.NewReader(data.Encode()))
	req.ContentLength = int64(len(data.Encode()))

	var response interface{}
	return s.RequestAPI(req, &response)
}
