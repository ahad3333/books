package storage_testing

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"

	"add/models"
)

func TestBookInsert(t *testing.T) {

	tests := []struct {
		Name    string
		Input   *models.CreateBook
		WantErr bool
	}{
		{
			Name: "case 1",
			Input: &models.CreateBook{
				Name:        "Time",
				Price:       22000,
				Description: "OK",
			},
			WantErr: false,
		},
		{
			Name: "case 2",
			Input: &models.CreateBook{
				Price:       22000,
				Description: "OK",
			},
			WantErr: false,
		},
		{
			Name: "case 3",
			Input: &models.CreateBook{
				Price:       22000,
				Description: "OK",
			},
			WantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {

			got, err := bookRepo.Insert(context.Background(), tc.Input)
			if err != nil {
				t.Errorf("%s: expected: %v, got: %v", tc.Name, tc.WantErr, err)
				return
			}

			if got == "" {
				t.Errorf("%s: got: %v", tc.Name, got)
				return
			}
		})
	}
}

func TestBookGetById(t *testing.T) {

	tests := []struct {
		Name    string
		Input   *models.BookPrimeryKey
		Output  *models.Book
		WantErr bool
	}{
		{
			Name: "case 1",
			Input: &models.BookPrimeryKey{
				Id: "59930f89-8849-485c-ad0b-f05704fdffd4",
			},
			Output: &models.Book{
				Name:        "Time",
				Price:       23000,
				Description: "OK",
			},
			WantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {

			got, err := bookRepo.GetByID(context.Background(), tc.Input)
			if err != nil {
				t.Errorf("%s: expected: %v, got: %v", tc.Name, tc.WantErr, err)
				return
			}

			comparer := cmp.Comparer(func(x, y models.Book) bool {
				return x.Name == y.Name &&
					x.Price == y.Price &&
					x.Description == y.Description
			})

			if diff := cmp.Diff(tc.Output, got, comparer); diff != "" {
				t.Error(diff)
				return
			}
		})
	}
}
func TestBookUpdate(t *testing.T)  {
	tests := []struct{
		Name    string
		Input   *models.UpdateBook
		Output  *models.Book
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.UpdateBook{
				Id: "59930f89-8849-485c-ad0b-f05704fdffd4",
				Name:        "Time",
				Price:       23000,
				Description: "OK",
			},
			Output: &models.Book{
				Name:        "Time",
				Price:       23000,
				Description: "OK",
			},
			WantErr: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {

			 err := bookRepo.Update(context.Background(), tc.Input)
			if err != nil {
				t.Errorf("%s: expected: %v, got: %v", tc.Name, tc.WantErr, err)
				return
			}

			comparer := cmp.Comparer(func(x, y models.Book) bool {
				return x.Name == y.Name &&
					x.Price == y.Price &&
					x.Description == y.Description
			})

			if diff := cmp.Diff(tc.Output,  comparer); diff != "" {
				t.Error(diff)
				return
			}
		})
	}
}

func TestBookDelete(t *testing.T)  {
	
	tests := []struct {
		Name    string
		Input   *models.BookPrimeryKey
		Output  *models.Book
		WantErr bool
	}{
		{
			Name: "case 1",
			Input: &models.BookPrimeryKey{
				Id: "59930f89-8849-485c-ad0b-f05704fdffd4",
			},
			Output: &models.Book{
				Name:        "Time",
				Price:       23000,
				Description: "OK",
			},
			WantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {

			 err := bookRepo.Delete(context.Background(), tc.Input)
			if err != nil {
				t.Errorf("%s: expected: %v, got: %v", tc.Name, tc.WantErr, err)
				return
			}

			comparer := cmp.Comparer(func(x, y models.Book) bool {
				return x.Name == y.Name &&
					x.Price == y.Price &&
					x.Description == y.Description
			})

			if diff := cmp.Diff(tc.Output,  comparer); diff != "" {
				t.Error(diff)
				return
			}
		})
	}
}