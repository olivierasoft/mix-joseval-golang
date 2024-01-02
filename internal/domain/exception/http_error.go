package exception

type HttpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Show    bool   `json:"show"`
}
