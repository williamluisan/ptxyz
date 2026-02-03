package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"

	entity "ptxyz/main-service/internal/domain/entity/auth"
	repository "ptxyz/main-service/internal/domain/repository/auth"

	"github.com/spf13/viper"
)

type customerProductClientImpl struct {
	baseURL string
	client  *http.Client
}

func NewCustomerProductServiceAuth(baseURL string) repository.AuthRepository {
	return &customerProductClientImpl{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

func (c *customerProductClientImpl) VerifyCredential(ctx context.Context, nik string, password string) (*entity.AuthIdentity, error) {
	body, _ := json.Marshal(map[string]string{
		"nik":    nik,
		"password": password,
	})

	req, _ := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.baseURL + viper.GetString("CUSTOMER_PRODUCT_AUTH_LOGIN_EP"),
		bytes.NewBuffer(body),
	)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("username atau password salah")
	}

	var result entity.AuthIdentity
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
