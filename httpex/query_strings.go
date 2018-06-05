package httpex

import "net/url"

// MergeQueryStrings merge query strings
func MergeQueryStrings(origin *string, qs map[string]string) error {
	// Parse origin url
	u, err := url.Parse(*origin)
	if err != nil {
		return err
	}
	// Add query strings
	params := u.Query()
	for k, v := range qs {
		params.Set(k, v)
	}
	// Set raw query
	u.RawQuery = params.Encode()
	// Update origin
	*origin = u.String()

	return nil
}
