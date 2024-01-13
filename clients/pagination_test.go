package clients

import (
	"testing"
)

func TestCreatePaginationResult(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name           string
		currentPage    int
		totalItems     int
		perPage        int
		expectedResult PaginationResult
		items          []interface{}
		hasError       bool
	}{
		{
			name:        "Test Case 1",
			currentPage: 1,
			totalItems:  10,
			perPage:     10,
			items:       make([]interface{}, 10),
			hasError:    false,
			expectedResult: PaginationResult{
				MetaData: PaginationMetaData{
					HasNext: true,
					HasPrev: false,
					Page:    1,
					PerPage: 10,
				},
			},
		},
		{
			name:        "Test Case 2",
			currentPage: 1,
			totalItems:  10,
			perPage:     10,
			items:       make([]interface{}, 8),
			hasError:    false,
			expectedResult: PaginationResult{
				MetaData: PaginationMetaData{
					HasNext: false,
					HasPrev: false,
					Page:    1,
					PerPage: 10,
				},
			},
		},
		{
			name:        "Test Case 3",
			currentPage: 2,
			totalItems:  10,
			perPage:     10,
			items:       make([]interface{}, 10),
			hasError:    false,
			expectedResult: PaginationResult{
				MetaData: PaginationMetaData{
					HasNext: true,
					HasPrev: true,
					Page:    2,
					PerPage: 10,
				},
			},
		},
		{
			name:        "Test Case 4",
			currentPage: 2,
			totalItems:  10,
			perPage:     10,
			items:       make([]interface{}, 10),
			hasError:    false,
			expectedResult: PaginationResult{
				MetaData: PaginationMetaData{
					HasNext: true,
					HasPrev: true,
					Page:    2,
					PerPage: 10,
				},
			},
		},
		{
			name:        "Test Error 1",
			currentPage: -1,
			totalItems:  10,
			perPage:     10,
			items:       make([]interface{}, 10),
			hasError:    true,
		},
		{
			name:        "Test Error 2",
			currentPage: 1,
			totalItems:  10,
			perPage:     -1,
			items:       make([]interface{}, 10),
			hasError:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err, result := CreatePaginationResult(tc.items, tc.currentPage, tc.perPage)
			if tc.hasError && err == nil {
				t.Errorf("For %s, expected error but got nil", tc.name)
			}

			if !tc.hasError {
				if result.MetaData != tc.expectedResult.MetaData {
					t.Errorf("For %s, expected %v but got %v", tc.name, tc.expectedResult.MetaData, result.MetaData)
				}
			}

		})
	}
}
