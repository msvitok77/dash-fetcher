package urls

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	fileURL       = "file"
	fileURLPrefix = fileURL + "://"

	https = "https"
)

// URLResourceReader gets the 'io.ReadCloser' for given URL or an error.
// It's caller's responsibility to close the returned 'io.ReadClose'.
func URLResourceReader(sourceURL string) (io.ReadCloser, error) {
	parsedURL, err := url.Parse(sourceURL)
	if err != nil {
		return nil, fmt.Errorf("parsing source URL: %w", err)
	}

	// web resource
	if parsedURL.Scheme != "" && parsedURL.Scheme != fileURL {
		client := http.DefaultClient

		// handle https
		if strings.HasPrefix(parsedURL.Scheme, https) {
			tr := &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
			client = &http.Client{Transport: tr}
		}

		resp, err := client.Get(sourceURL)
		if err != nil {
			return nil, fmt.Errorf("reading data from web %q %w", sourceURL, err)
		}

		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			return nil, fmt.Errorf("reading data from web %q status: %d", sourceURL, resp.StatusCode)
		}

		return resp.Body, nil
	}

	// file resource
	return os.Open(sourceURL[len(fileURLPrefix):])
}
