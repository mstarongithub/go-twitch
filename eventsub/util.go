package eventsub

import "net/url"

// Guaranteed to have a scheme of ws/wss
// If no scheme is specified, the scheme is set to ws
func decodeURL(rawUrl string) (*url.URL, error) {
	u, err := url.Parse(rawUrl)
	if err != nil {
		return nil, err
	}
	if u.Scheme == "" {
		u.Scheme = "ws"
	}

	return u, nil
}
