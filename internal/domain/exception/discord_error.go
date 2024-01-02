package exception

type DiscordErrorRequest struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}
