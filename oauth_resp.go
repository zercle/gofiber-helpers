package helpers

// OauthRequest oauth request by IETF
type OauthRequest struct {
	APIKey       string `json:"client_id" form:"client_id" query:"client_id"`
	APISecret    string `json:"client_secret" form:"client_secret" query:"client_secret"`
	ResponseType string `json:"response_type" form:"response_type" query:"response_type"`
	RedirectURI  string `json:"redirect_uri" form:"redirect_uri" query:"redirect_uri"`
	Scope        string `json:"scope" form:"scope" query:"scope"`
	State        string `json:"state" form:"state" query:"state"`
	Code         string `json:"code" form:"code" query:"code"`
	GrantType    string `json:"grant_type" form:"grant_type" query:"grant_type"`
	UserName     string `json:"username" form:"username" query:"username"`
	Password     string `json:"password" form:"password" query:"password"`
	RefreshToken string `json:"refresh_token" form:"refresh_token" query:"refresh_token"`
	Token        string `json:"token" form:"token" query:"token"`
}

// OauthResponse oauth request by IETF
type OauthResponse struct {
	Scope        string      `json:"scope,omitempty"`
	State        string      `json:"state,omitempty"`
	Code         string      `json:"code,omitempty"`
	Error        OauthErr    `json:"error,omitempty"`
	ErrorDesc    string      `json:"error_description,omitempty"`
	ErrorURI     string      `json:"error_uri,omitempty"`
	AccessToken  string      `json:"access_token,omitempty"`
	IDToken      string      `json:"id_token,omitempty"`
	TokenType    string      `json:"token_type,omitempty"`
	ExpiresIn    int         `json:"expires_in,omitempty"`
	RefreshToken string      `json:"refresh_token,omitempty"`
	Data         interface{} `json:"data,omitempty"`
}

// OauthErr Oauth error response
type OauthErr string

// Oauth Error Response
const (
	InvalidRequest     OauthErr = "invalid_request"
	UnauthorizedClient OauthErr = "unauthorized_client"
	AccessDenied       OauthErr = "access_denied"
	UnsupportType      OauthErr = "unsupported_response_type"
	InvalidScope       OauthErr = "invalid_scope"
	ServerError        OauthErr = "server_error"
	Unavailable        OauthErr = "temporarily_unavailable"
)

// GrantType Oauth grant type
type GrantType string

// GrantType constant
const (
	GrantTypeCode     GrantType = "authorization_code"
	GrantTypeClient   GrantType = "client_credentials"
	GrantTypeRefresh  GrantType = "refresh_token"
	GrantTypeaccess   GrantType = "access_token"
	GrantTypePassword GrantType = "password"
)

// ResponseType Oauth response type
type ResponseType string

// Responsetype constant
const (
	ResponseTypeCode  ResponseType = "code"
	ResponseTypeToken ResponseType = "token"
)
