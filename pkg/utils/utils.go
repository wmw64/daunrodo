package utils

import "net/url"

func IsURLValid(uri string) bool {

	u, err := url.Parse(uri)

	return err == nil && u.Scheme != "" && u.Host != "" && (u.Scheme == "http" || u.Scheme == "https")
}

// FixURL prepends https scheme to URL.
// Example: instagram.com => https://instagram.com
func FixURL(uri string) string {

	u, err := url.Parse(uri)
	if err == nil && (u.Scheme == "" || (u.Scheme != "http" && u.Scheme != "https")) {
		u.Scheme = "https"

		return u.String()
	}

	return uri
}
