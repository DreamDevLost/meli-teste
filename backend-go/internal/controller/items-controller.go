package controller

import (
	"backend-api/internal/dto"
	"backend-api/internal/service"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type ItemsController interface {
	GetItemByID(w http.ResponseWriter, r *http.Request)
	SearchItems(w http.ResponseWriter, r *http.Request)
}

type itemsController struct {
	meliAPI service.MeliAPI
}

func NewItemsController(meliAPI service.MeliAPI) ItemsController {
	return &itemsController{meliAPI: meliAPI}
}

func (i *itemsController) GetItemByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")
	log.Printf("[INFO][ITEMS-CONTROLLER][GET-ITEM-BY-ID] getting item by id=%s", id)
	item, err := i.meliAPI.GetItemByID(ctx, id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		log.Printf("[ERROR][ITEMS-CONTROLLER][GET-ITEM-BY-ID] error getting item by id: %v", err)
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if item == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	desc, err := i.meliAPI.GetItemDescription(ctx, id)
	if err != nil {
		log.Printf("[ERROR][ITEMS-CONTROLLER][GET-ITEM-BY-ID] error getting item description: %v", err)
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if desc == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	response := dto.GetItemResponseDto{
		Author: dto.Author{
			Name:     "John",
			LastName: "Doe",
		},
		Item: dto.SingleItem{
			ID:    item.ID,
			Title: item.Title,
			Price: dto.Price{
				Currency: item.CurrencyID,
				Amount:   item.Price,
				Decimals: 0,
			},
			Picture:      item.Thumbnail,
			Condition:    item.Condition,
			FreeShipping: item.Shipping.FreeShipping,
			SoldQuantity: item.InitialQuantity,
			Description:  desc.PlainText,
		},
	}

	resJSON, err := json.Marshal(response)
	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("[ERROR][ITEMS-CONTROLLER][GET-ITEM-BY-ID] error marshalling item: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resJSON)
}

func (i *itemsController) SearchItems(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	query := r.URL.Query().Get("search")
	log.Printf("[INFO][ITEMS-CONTROLLER][SEARCH-ITEMS] searching items with query=%s", query)
	items, err := i.meliAPI.SearchItems(ctx, query)
	if err != nil {
		log.Printf("[ERROR][ITEMS-CONTROLLER][SEARCH-ITEMS] error searching items: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(err.Error()))
		return
	}

	responseItems := make([]*dto.Item, len(items.Results))
	for i, item := range items.Results {
		responseItems[i] = &dto.Item{
			ID:    item.ID,
			Title: item.Title,
			Price: dto.Price{
				Currency: item.CurrencyID,
				Amount:   item.Price,
				Decimals: 0,
			},
			Picture:      item.Thumbnail,
			Condition:    item.Condition,
			FreeShipping: item.Shipping.FreeShipping,
		}
	}
	// map items to GetItemsResponseDto
	response := dto.GetItemsResponseDto{
		Author: dto.Author{
			Name:     "John",
			LastName: "Doe",
		},
		Categories: []string{"category1", "category2"},
		Items:      responseItems,
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("[ERROR][ITEMS-CONTROLLER][SEARCH-ITEMS] error marshalling response: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
