package werewolf

type APIResponse struct {
	Code    int         `json:"Code" xml:"Code"`
	SubCode string      `json:"SubCode" xml:"SubCode"`
	Message string      `json:"Message" xml:"Message"`
	Data    interface{} `json:"Data" xml:"Data"`
}
