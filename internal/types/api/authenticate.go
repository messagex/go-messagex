package api

// AuthoriseRequest - AuthoriseRequest API struct
type AuthenticateRequest struct {
	APIKey    string `json:"apiKey"`
	APISecret string `json:"apiSecret"`
}

type AuthenticateResponse struct {
	Data *AuthenticateResponseData `json:"data"`
}

type AuthenticateResponseData struct {
	ID             int64 `json:"id"`
	ApiKeyID       int64 `json:"apiKeyId"`
	BearerToken    string `json:"bearerToken"`
	RefreshToken   string `json:"refreshToken"`
	ExpiresAt      string `json:"expiresAt"`
	CreatedAt      string `json:"createdAt"`
	UpdatedAt      string `json:"updatedAt"`
	CreatedAtEpoch int64  `json:"createdAtEpoch"`
	UpdatedAtEpoch int64  `json:"updatedAtEpoch"`
}
