package dynconfig

type ConfigProvider interface {
	LoadConfig() *Config
}

type Config struct {
	Metadata *Metadata      `yaml:"metadata"`
	Static   *StaticConfig  `yaml:"static"`
	Dynamic  *DynamicConfig `yaml:"dynamic"`
	Metrics  *Metrics       `yaml:"metrics"`
}

type StaticConfig struct {
	Client *Client `yaml:"client"`
}

type DynamicConfig struct {
	Client      *Client      `yaml:"client"`
	Read        *Read        `yaml:"read"`
	Write       *Write       `yaml:"write"`
	Query       *Query       `yaml:"query"`
	Scan        *Scan        `yaml:"scan"`
	BatchRead   *BatchRead   `yaml:"batch_read"`
	BatchWrite  *BatchWrite  `yaml:"batch_write"`
	BatchUdf    *BatchUdf    `yaml:"batch_udf"`
	BatchDelete *BatchDelete `yaml:"batch_delete"`
	TxnRoll     *TxnRoll     `yaml:"txn_roll"`
	TxnVerify   *TxnVerify   `yaml:"txn_verify"`
	Metrics     *Metrics     `yaml:"metrics"`
}

type Metadata struct {
	Name       *string `yaml:"name"`
	Version    *string `yaml:"version"`
	Generation *int    `yaml:"generation"`
}

type Client struct {
	// static config
	ConfigInterval             *int `yaml:"config_interval"`
	MaxConnectionsPerNode      *int `yaml:"max_connections_per_node"`
	MinConnectionsPerNode      *int `yaml:"min_connections_per_node"`
	AsyncMaxConnectionsPerNode *int `yaml:"async_max_connections_per_node"`
	AsyncMinConnectionsPerNode *int `yaml:"async_min_connections_per_node"`

	// dynamic config
	Timeout               *int   `yaml:"timeout"`
	ErrorRateWindow       *int   `yaml:"error_rate_window"`
	MaxErrorRate          *int   `yaml:"max_error_rate"`
	FailIfNotConnected    *bool  `yaml:"fail_if_not_connected"`
	LoginTimeout          *int   `yaml:"login_timeout"`
	MaxSocketIdle         *int   `yaml:"max_socket_idle"`
	RackAware             *bool  `yaml:"rack_aware"`
	RackIds               *[]int `yaml:"rack_ids"`
	TendInterval          *int   `yaml:"tend_interval"`
	UseServiceAlternative *bool  `yaml:"use_service_alternative"`
}

type Read struct {
	ReadModeAp          *ReadModeAp `yaml:"read_mode_ap"`
	ReadModeSc          *ReadModeSc `yaml:"read_mode_sc"`
	ConnectTimeout      *int        `yaml:"connect_timeout"`
	FailOnFilteredOut   *bool       `yaml:"fail_on_filtered_out"`
	Replica             *Replica    `yaml:"replica"`
	SleepBetweenRetries *int        `yaml:"sleep_between_retries"`
	SocketTimeout       *int        `yaml:"socket_timeout"`
	TimeoutDelay        *int        `yaml:"timeout_delay"`
	TotalTimeout        *int        `yaml:"total_timeout"`
	MaxRetries          *int        `yaml:"max_retries"`
}

type Write struct {
	ConnectTimeout      *int     `yaml:"connect_timeout"`
	FailOnFilteredOut   *bool    `yaml:"fail_on_filtered_out"`
	Replica             *Replica `yaml:"replica"`
	SendKey             *bool    `yaml:"send_key"`
	SleepBetweenRetries *int     `yaml:"sleep_between_retries"`
	SocketTimeout       *int     `yaml:"socket_timeout"`
	MaxRetries          *int     `yaml:"max_retries"`
	DurableDelete       *bool    `yaml:"durable_delete"`
}

type Query struct {
	ReadModeAp          *ReadModeAp `yaml:"read_mode_ap"`
	ReadModeSc          *ReadModeSc `yaml:"read_mode_sc"`
	ConnectTimeout      *int        `yaml:"connect_timeout"`
	Replica             *Replica    `yaml:"replica"`
	SleepBetweenRetries *int        `yaml:"sleep_between_retries"`
	SocketTimeout       *int        `yaml:"socket_timeout"`
	TimeoutDelay        *int        `yaml:"timeout_delay"`
	TotalTimeout        *int        `yaml:"total_timeout"`
	MaxRetries          *int        `yaml:"max_retries"`
	IncludeBinData      *bool       `yaml:"include_bin_data"`
	InfoTimeout         *int        `yaml:"info_timeout"`
	RecordQueueSize     *int        `yaml:"record_queue_size"`
	ExpectedDuration    *Duration   `yaml:"expected_duration"`
}

type Scan struct {
	ReadModeAp          *ReadModeAp `yaml:"read_mode_ap"`
	ReadModeSc          *ReadModeSc `yaml:"read_mode_sc"`
	ConnectTimeout      *int        `yaml:"connect_timeout"`
	Replica             *Replica    `yaml:"replica"`
	SleepBetweenRetries *int        `yaml:"sleep_between_retries"`
	SocketTimeout       *int        `yaml:"socket_timeout"`
	TimeoutDelay        *int        `yaml:"timeout_delay"`
	TotalTimeout        *int        `yaml:"total_timeout"`
	MaxRetries          *int        `yaml:"max_retries"`
	ConcurrentNodes     *int        `yaml:"concurrent_nodes"`
	MaxConcurrentNodes  *int        `yaml:"max_concurrent_nodes"`
}

type BatchRead struct {
	ReadModeAp          *ReadModeAp `yaml:"read_mode_ap"`
	ReadModeSc          *ReadModeSc `yaml:"read_mode_sc"`
	ConnectTimeout      *int        `yaml:"connect_timeout"`
	FailOnFilteredOut   *bool       `yaml:"fail_on_filtered_out"`
	Replica             *Replica    `yaml:"replica"`
	SleepBetweenRetries *int        `yaml:"sleep_between_retries"`
	SocketTimeout       *int        `yaml:"socket_timeout"`
	TimeoutDelay        *int        `yaml:"timeout_delay"`
	TotalTimeout        *int        `yaml:"total_timeout"`
	MaxRetries          *int        `yaml:"max_retries"`
	MaxConcurrentThread *int        `yaml:"max_concurrent_thread"`
	AllowInline         *bool       `yaml:"allow_inline"`
	AllowInlineSSD      *bool       `yaml:"allow_inline_ssd"`
	RespondAllKeys      *bool       `yaml:"respond_all_keys"`
}

type BatchWrite struct {
	ConnectTimeout      *int     `yaml:"connect_timeout"`
	FailOnFilteredOut   *bool    `yaml:"fail_on_filtered_out"`
	Replica             *Replica `yaml:"replica"`
	SleepBetweenRetries *int     `yaml:"sleep_between_retries"`
	SocketTimeout       *int     `yaml:"socket_timeout"`
	TimeoutDelay        *int     `yaml:"timeout_delay"`
	TotalTimeout        *int     `yaml:"total_timeout"`
	MaxRetries          *int     `yaml:"max_retries"`
	DurableDelete       *bool    `yaml:"durable_delete"`
	SendKey             *bool    `yaml:"send_key"`
	MaxConcurrentThread *int     `yaml:"max_concurrent_thread"`
	AllowInline         *bool    `yaml:"allow_inline"`
	AllowInlineSSD      *bool    `yaml:"allow_inline_ssd"`
	RespondAllKeys      *bool    `yaml:"respond_all_keys"`
}

type BatchUdf struct {
	DurableDelete *bool `yaml:"durable_delete"`
	SendKey       *bool `yaml:"send_key"`
}

type BatchDelete struct {
	DurableDelete *bool `yaml:"durable_delete"`
	SendKey       *bool `yaml:"send_key"`
}

type TxnRoll struct {
	ReadModeAp          *ReadModeAp `yaml:"read_mode_ap"`
	ReadModeSc          *ReadModeSc `yaml:"read_mode_sc"`
	ConnectTimeout      *int        `yaml:"connect_timeout"`
	Replica             *Replica    `yaml:"replica"`
	SleepBetweenRetries *int        `yaml:"sleep_between_retries"`
	SocketTimeout       *int        `yaml:"socket_timeout"`
	TimeoutDelay        *int        `yaml:"timeout_delay"`
	TotalTimeout        *int        `yaml:"total_timeout"`
	MaxRetries          *int        `yaml:"max_retries"`
	MaxConcurrentThread *int        `yaml:"max_concurrent_thread"`
	AllowInline         *bool       `yaml:"allow_inline"`
	AllowInlineSSD      *bool       `yaml:"allow_inline_ssd"`
	RespondAllKeys      *bool       `yaml:"respond_all_keys"`
}

type TxnVerify struct {
	ReadModeAp          *ReadModeAp `yaml:"read_mode_ap"`
	ReadModeSc          *ReadModeSc `yaml:"read_mode_sc"`
	ConnectTimeout      *int        `yaml:"connect_timeout"`
	Replica             *Replica    `yaml:"replica"`
	SleepBetweenRetries *int        `yaml:"sleep_between_retries"`
	SocketTimeout       *int        `yaml:"socket_timeout"`
	TimeoutDelay        *int        `yaml:"timeout_delay"`
	TotalTimeout        *int        `yaml:"total_timeout"`
	MaxRetries          *int        `yaml:"max_retries"`
	MaxConcurrentThread *int        `yaml:"max_concurrent_thread"`
	AllowInline         *bool       `yaml:"allow_inline"`
	AllowInlineSSD      *bool       `yaml:"allow_inline_ssd"`
	RespondAllKeys      *bool       `yaml:"respond_all_keys"`
}

type Metrics struct {
	Enable         *bool `yaml:"enable"`
	LatencyShift   *int  `yaml:"latency_shift"`
	LatencyColumns *int  `yaml:"latency_columns"`
}

// Enum types
type ReadModeAp int

const (
	One ReadModeAp = iota
	All
)

var readModeAp = map[ReadModeAp]string{
	One: "one",
	All: "all",
}

type ReadModeSc int

const (
	Session ReadModeSc = iota
	Linearize
	AllowReplica
	AllowUnavailable
)

var readModeSc = map[ReadModeSc]string{
	Session:          "session",
	Linearize:        "linearize",
	AllowReplica:     "allow_replica",
	AllowUnavailable: "allow_unavailable",
}

type Replica int

const (
	Master Replica = iota
	MasterProles
	Sequence
	PreferRack
)

var replica = map[Replica]string{
	Master:       "master",
	MasterProles: "master_proles",
	Sequence:     "sequence",
	PreferRack:   "prefer_rack",
}

type Duration int

const (
	Long Duration = iota
	Short
	LongRelaxAp
)

var duration = map[Duration]string{
	Long:        "long",
	Short:       "short",
	LongRelaxAp: "long_relax_ap",
}
