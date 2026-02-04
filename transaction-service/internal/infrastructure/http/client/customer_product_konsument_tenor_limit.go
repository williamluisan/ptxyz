package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	entityKTL "ptxyz/transaction-service/internal/domain/entity/konsumen_tenor_limit"
	repositoryKTL "ptxyz/transaction-service/internal/domain/repository/konsumen_tenor_limit"

	"github.com/spf13/viper"
)

type konsumenTenorLimitImpl struct {
	baseURL	 string
	client 	*http.Client
}

func NewKonsumenTenorLimit(baseURL string) repositoryKTL.KonsumenTenorLimitRepository {
	return &konsumenTenorLimitImpl{
		baseURL: baseURL,
		client: &http.Client{},
	}
}

func (c *konsumenTenorLimitImpl) UpdateBalance(ctx context.Context, entityUpdate *entityKTL.KTLUpdateBalance) error {
	body, _ := json.Marshal(map[string]any{
		"public_id": entityUpdate.PublicId,
		"balance": entityUpdate.Balance,
		"used_amount": entityUpdate.UsedAmount,
	})

	req, _ := http.NewRequestWithContext(
		ctx,
		http.MethodPut,
		c.baseURL + viper.GetString("CUSTOMER_PRODUCT_KONSUMEN_TENOR_LIMIT_UPDATE_BALANCE_EP"),
		bytes.NewBuffer(body),
	)
	if token, ok := ctx.Value("authorization").(string); ok && token != "" {
		req.Header.Set("Authorization", token)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return errors.New("gagal update balance")
	}

	return nil
}