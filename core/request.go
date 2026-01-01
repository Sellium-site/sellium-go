package core

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type RateLimit struct {
	Limit     int
	Remaining int
	ResetSec  int
}

type ResponseMeta struct {
	Headers   http.Header
	Status    int
	RateLimit *RateLimit
}

type envelope[T any] struct {
	Success bool          `json:"success"`
	Data    T             `json:"data"`
	Error   *APIErrorBody `json:"error,omitempty"`
}

func (c *Client) Do(ctx context.Context, method, path string, query url.Values, body any, out any) (*ResponseMeta, error) {
	u := c.BaseURL + path
	if query != nil && len(query) > 0 {
		u += "?" + query.Encode()
	}

	var rdr io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		rdr = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, u, rdr)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-API-Key", c.APIKey)
	req.Header.Set("X-Store-ID", c.StoreID)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	res, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	meta := &ResponseMeta{
		Headers:   res.Header.Clone(),
		Status:    res.StatusCode,
		RateLimit: parseRateLimit(res.Header),
	}

	raw, err := io.ReadAll(res.Body)
	if err != nil {
		return meta, err
	}
	if len(raw) == 0 {
		return meta, nil
	}

	var probe envelope[json.RawMessage]
	if err := json.Unmarshal(raw, &probe); err == nil {
		if probe.Success {
			if out == nil {
				return meta, nil
			}
			return meta, json.Unmarshal(raw, out)
		}

		if probe.Error != nil {
			return meta, &APIError{
				Status:  res.StatusCode,
				Code:    probe.Error.Code,
				Message: probe.Error.Message,
				Raw:     raw,
			}
		}
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return meta, &APIError{
			Status:  res.StatusCode,
			Code:    "HTTP_ERROR",
			Message: "request failed",
			Raw:     raw,
		}
	}

	if out == nil {
		return meta, nil
	}
	if err := json.Unmarshal(raw, out); err != nil {
		return meta, errors.New("failed to decode response: " + err.Error())
	}
	return meta, nil
}

func parseRateLimit(h http.Header) *RateLimit {
	limit, _ := strconv.Atoi(h.Get("X-RateLimit-Limit"))
	rem, _ := strconv.Atoi(h.Get("X-RateLimit-Remaining"))
	reset, _ := strconv.Atoi(h.Get("X-RateLimit-Reset"))
	if limit == 0 && rem == 0 && reset == 0 {
		return nil
	}
	return &RateLimit{Limit: limit, Remaining: rem, ResetSec: reset}
}
