package dpfm_api_input_reader

type Request struct {
	ConnectionKey     string           `json:"connection_key"`
	Result            bool             `json:"result"`
	RedisKey          string           `json:"redis_key"`
	Filepath          string           `json:"filepath"`
	APIStatusCode     int              `json:"api_status_code"`
	RuntimeSessionID  string           `json:"runtime_session_id"`
	BusinessPartnerID *int             `json:"business_partner"`
	ServiceLabel      string           `json:"service_label"`
	APIType           string           `json:"api_type"`
	DeliveryDocument  DeliveryDocument `json:"DeliveryDocument"`
	APISchema         string           `json:"api_schema"`
	Accepter          []string         `json:"accepter"`
	Deleted           bool             `json:"deleted"`
}

type DeliveryDocument struct {
	FilePath string `json:"FilePath"`
	API      string `json:"API"`
}
