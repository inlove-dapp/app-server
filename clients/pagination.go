package clients

import "fmt"

type PaginationMetaData struct {
	HasNext bool `json:"has_next"`
	HasPrev bool `json:"has_prev"`
	Page    int  `json:"page"`
	PerPage int  `json:"per_page"`
}

type PaginationResult struct {
	MetaData PaginationMetaData `json:"metadata"`
	Items    interface{}        `json:"items"`
}

// CreatePaginationResult creates a pagination result from the given items, page, and perPage.
func CreatePaginationResult(items []interface{}, page int, perPage int) (error, *PaginationResult) {
	if page < 1 {
		return fmt.Errorf("page must be greater than 0"), nil
	}

	if perPage < 1 {
		return fmt.Errorf("perPage must be greater than 0"), nil
	}

	if len(items) < perPage {
		return nil, &PaginationResult{
			MetaData: PaginationMetaData{
				HasNext: false,
				HasPrev: page > 1,
				Page:    page,
				PerPage: perPage,
			},
			Items: items,
		}
	}

	return nil, &PaginationResult{
		MetaData: PaginationMetaData{
			HasNext: true,
			HasPrev: page > 1,
			Page:    page,
			PerPage: perPage,
		},
		Items: items[:perPage],
	}
}
