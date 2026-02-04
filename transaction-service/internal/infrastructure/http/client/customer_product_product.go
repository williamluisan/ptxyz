package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	entity "ptxyz/transaction-service/internal/domain/entity/product"
	repositoryProduct "ptxyz/transaction-service/internal/domain/repository/product"

	"github.com/spf13/viper"
)

type productImpl struct {
	baseURL string
	client *http.Client
}

func NewProduct(baseURL string) repositoryProduct.ProductRepository {
	return &productImpl{
		baseURL: baseURL,
		client: &http.Client{},
	}
}

func (c *productImpl) GetByPublicId(ctx context.Context, publicId string) (*entity.Product, error) {
	url := fmt.Sprintf(
		"%s%s/%s",
		c.baseURL,
		viper.GetString("CUSTOMER_PRODUCT_PRODUCT_EP"),
		publicId,
	)

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		url,
		nil, // no body
	)
	if err != nil {
		return nil, err
	}
	if token, ok := ctx.Value("authorization").(string); ok && token != "" {
		req.Header.Set("Authorization", token)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("gagal mendapatkan detail produk")
	}

	var res getProductResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	product := &entity.Product{
		PublicId: res.Data.PublicId,
		Name: res.Data.Name,
		HargaOtr: res.Data.HargaOtr,
		Category: res.Data.Category,
	}

	return product, err
}