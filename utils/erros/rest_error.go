package erros

//StatusError retorno das mensages de erro conforme o padrão rest
type StatusError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Error      string `json:"error"`
}
