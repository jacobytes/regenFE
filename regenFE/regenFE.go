package regenfe

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	Version          = "v0.1.0"
	defaultBaseUrl   = "https://api.formula-e.pulselive.com/formula-e/v1/"
	defaultUserAgent = "regenFE" + "/" + Version
)

var errNonNilContext = errors.New("Context must be non-nill")

type Client struct {
	client    *http.Client
	BaseUrl   *url.URL
	UserAgent string

	common service

	Teams         *TeamService
	Drivers       *DriverService
	Championships *ChampionshipService
	Series        *SeriesService
	Races         *RaceService
}

type service struct {
	client *Client
}

func NewClient() Client {

	httpClient := &http.Client{}

	c := Client{client: httpClient}
	c.initialize()

	return c
}

func (c *Client) initialize() {

	if c.client == nil {
		c.client = &http.Client{}
	}

	if c.BaseUrl == nil {
		c.BaseUrl, _ = url.Parse(defaultBaseUrl)
	}

	if c.UserAgent == "" {
		c.UserAgent = defaultUserAgent
	}

	c.common.client = c
	c.Teams = (*TeamService)(&c.common)
	c.Drivers = (*DriverService)(&c.common)
	c.Championships = (*ChampionshipService)(&c.common)
	c.Series = (*SeriesService)(&c.common)
	c.Races = (*RaceService)(&c.common)
}

type RequestOptions func(req *http.Request)

func (c *Client) NewRequest(method, urlStr string, body any, opts ...RequestOptions) (*http.Request, error) {

	if !strings.HasSuffix(c.BaseUrl.Path, "/") {
		return nil, fmt.Errorf("baseURL must have a trailing slash, but %q does not", c.BaseUrl)
	}

	url, err := c.BaseUrl.Parse(urlStr)

	if err != nil {
		return nil, err
	}

	var buffer io.ReadWriter
	if body != nil {
		buffer = &bytes.Buffer{}
		encoder := json.NewEncoder(buffer)
		encoder.SetEscapeHTML(false)
		err := encoder.Encode(body)

		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url.String(), buffer)

	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	for _, opt := range opts {
		opt(req)
	}

	return req, nil
}

func (c *Client) bareDo(ctx context.Context, req *http.Request) (*Response, error) {

	if ctx == nil {
		return nil, errNonNilContext
	}

	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)

	var response *Response

	if resp != nil {
		response = newResponse(resp)
	}

	if err != nil {
		select {
		case <-ctx.Done():
			return response, ctx.Err()

		default:
		}
	}

	err = CheckResponse(resp)

	if err != nil {
		return nil, err
	}

	return response, err
}

func (c *Client) Do(ctx context.Context, req *http.Request, value any) (*Response, error) {
	resp, err := c.bareDo(ctx, req)

	if err != nil {
		return resp, err
	}

	defer resp.Body.Close()

	decodeErr := json.NewDecoder(resp.Body).Decode(value)

	if decodeErr == io.EOF {
		decodeErr = nil // ignore EOF errors caused by empty response body
	}
	if decodeErr != nil {
		err = decodeErr
	}

	return resp, err
}

type Response struct {
	*http.Response
}

func newResponse(resp *http.Response) *Response {
	response := &Response{Response: resp}
	return response
}

func CheckResponse(r *http.Response) error {

	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	bodyBytes, _ := io.ReadAll(r.Body)
	return fmt.Errorf("HTTP %d: %s", r.StatusCode, string(bodyBytes))
}

type ListOptions struct {
	Size int
	Page int
}

type PaginationResponse struct {
	PageInfo *Pagination `json:"pageInfo"`
}

type Pagination struct {
	Page       *int `json:"page"`
	NumPages   *int `json:"numPages"`
	PageSize   *int `json:"pageSize"`
	NumEntries *int `json:"numEntries"`
}

func Ptr[T any](v T) *T {
	return &v
}
