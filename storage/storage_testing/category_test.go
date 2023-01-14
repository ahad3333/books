package storage_testing

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"

	"add/models"
)

func TestCategoryInsert(t *testing.T) {

	tests := []struct {
		Name    string
		Input   *models.CreateCategory
		WantErr bool
	}{
		{
			Name: "case 1",
			Input: &models.CreateCategory{
				Name:        "Time",
			},
			WantErr: false,
		},{
			Name: "case 3",
			Input: &models.CreateCategory{
				Name: "Tame",
			},
			WantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {

			got, err := categoryRepo.Insert(context.Background(), tc.Input)
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

func TestCategoryGetById(t *testing.T) {

	tests := []struct {
		Name    string
		Input   *models.CategoryPrimeryKey
		Output  *models.Category
		WantErr bool
	}{
		{
			Name: "case 1",
			Input: &models.CategoryPrimeryKey{
				Id: "59930f89-8849-485c-ad0b-f05704fdffd4",
			},
			Output: &models.Category{
				Name:        "Time",
			},
			WantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {

			got, err := categoryRepo.GetByID(context.Background(), tc.Input)
			if err != nil {
				t.Errorf("%s: expected: %v, got: %v", tc.Name, tc.WantErr, err)
				return
			}

			comparer := cmp.Comparer(func(x, y models.Category) bool {
				return x.Name == y.Name
			})

			if diff := cmp.Diff(tc.Output, got, comparer); diff != "" {
				t.Error(diff)
				return
			}
		})
	}
}
func TestCategoryUpdate(t *testing.T)  {
	tests := []struct{
		Name    string
		Input   *models.UpdateCategory
		Output  *models.Category
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.UpdateCategory{
				Id: "59930f89-8849-485c-ad0b-f05704fdffd4",
				Name:        "Time",
			},
			Output: &models.Category{
				Name:        "Time",
			},
			WantErr: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {

			 err := categoryRepo.Update(context.Background(), tc.Input)
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

func TestCategoryDelete(t *testing.T)  {
	
	tests := []struct {
		Name    string
		Input   *models.CategoryPrimeryKey
		WantErr bool
	}{
		{
			Name: "case 1",
			Input: &models.CategoryPrimeryKey{
				Id: "59930f89-8849-485c-ad0b-f05704fdffd4",
			},
			WantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {

			 err := categoryRepo.Delete(context.Background(), tc.Input)
			if err != nil {
				t.Errorf("%s: expected: %v, got: %v", tc.Name, tc.WantErr, err)
				return
			}
		})
	}
}