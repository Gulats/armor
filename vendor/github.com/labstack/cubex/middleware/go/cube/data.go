package cube

type Data struct {
	Uptime           int64             `json:"uptime"`
	NewConnection    int64             `json:"new_connection"`
	ActiveConnection int64             `json:"active_connection"`
	IdleConnection   int64             `json:"idle_connection"`
	Request          int64             `json:"request"`
	BytesIn          int64             `json:"bytes_in"`
	BytesOut         int64             `json:"bytes_out"`
	AverageLatency   int64             `json:"average_latency"`
	Endpoint         map[string]int64  `json:"endpoint"`
	UserAgent        map[string]int64  `json:"user_agent"`
	RemoteIP         map[string]int64  `json:"remote_ip"`
	Status           map[int32]int64   `json:"status"`
	Tags             map[string]string `json:"tags"`
}