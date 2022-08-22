package service

import (
	"context"
	"fmt"
	"golang/api-go-routine/models"
	"golang/api-go-routine/repository"
	"sync"
	"time"
)

var ctx context.Context = context.Background()

func CreateProducts(product models.Products) {
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	for _, p := range product.Data {
		wg.Add(1)
		go func(p models.Product) {
			result, err := repository.Insert(ctx, p)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(fmt.Sprintf("Product %s created with ID %d", p.Description, result))
			wg.Done()
		}(p)
	}

	wg.Wait()
}

func GetProducts(page int, limit int) models.Products {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	result := repository.Select(ctx, page, limit)

	return result
}
