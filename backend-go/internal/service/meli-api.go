package service

import (
	"backend-api/internal/service/responses"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type MeliAPI interface {
	GetItemByID(ctx context.Context, id string) (*responses.GetItemResponse, error)
	GetItemDescription(ctx context.Context, id string) (*responses.GetItemDescription, error)
	SearchItems(ctx context.Context, query string) (*responses.GetItemsResponse, error)
}

type meliAPI struct {
	c *http.Client
}

func NewMeliAPI(c *http.Client) MeliAPI {
	return &meliAPI{c: c}
}

var ErrNotFound = errors.New("item not found")

func (m *meliAPI) GetItemByID(ctx context.Context, id string) (*responses.GetItemResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.mercadolibre.com/items/"+id, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := m.c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error doing request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	response := &responses.GetItemResponse{}
	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return response, nil
}

func (m *meliAPI) GetItemDescription(ctx context.Context, id string) (*responses.GetItemDescription, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.mercadolibre.com/items/"+id+"/description", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := m.c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error doing request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	response := &responses.GetItemDescription{}
	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return response, nil
}

func (m *meliAPI) SearchItems(ctx context.Context, query string) (*responses.GetItemsResponse, error) {
	response := &responses.GetItemsResponse{}

	q := url.QueryEscape(query)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.mercadolibre.com/sites/MLA/search?q="+q, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := m.c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error doing request: %w", err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return response, nil
}
