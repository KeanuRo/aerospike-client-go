package dynconfig

type ConfigProvider interface {
	LoadConfig() (string, *Config)
}

type Config struct {
	Metadata Metadata
	Static   StaticConfig
	Dynamic  DynamicConfig
	Metrics  Metrics
}

type StaticConfig struct {
	Client Client
}

type DynamicConfig struct {
	Client      Client
	Read        Read
	Write       Write
	Query       Query
	Scan        Scan
	BatchRead   BatchRead
	BatchWrite  BatchWrite
	BatchUdf    BatchUdf
	BatchDelete BatchDelete
	TxnRoll     TxnRoll
	TxnVerify   TxnVerify
	Metrics     Metrics
}

type Metadata struct {
	Name       string
	Version    string
	Generation int
}

type Client struct {
	// static configuration
	ConfigInterval             int
	MaxConnectionsPerNode      int
	MinConnectionsPerNode      int
	AsyncMaxConnectionsPerNode int
	AsyncMinConnectionsPerNode int

	// dynamic configuration
	Timeout               int
	ErrorRateWindow       int
	MaxErrorRate          int
	FailIfNotConnected    bool
	LoginTimeout          int
	MaxSocketIdle         int
	RackAware             bool
	RackIds               []int
	TendInterval          int
	UseServiceAlternative bool
}

type Read struct {
	ReadModeAp          ReadModeAp
	ReadModeSc          ReadModeSc
	ConnectTimeout      int
	FailOnFilteredOut   bool
	Replica             Replica
	SleepBetweenRetries int
	SocketTimeout       int
	TimeoutDelay        int
	TotalTimeout        int
	MaxRetries          int
}

type Write struct {
	ConnectTimeout      int
	FailOnFilteredOut   bool
	Replica             Replica
	SendKey             bool
	SleepBetweenRetries int
	SocketTimeout       int
	MaxRetries          int
	DurableDelete       bool
}

type Query struct {
	ReadmModeAp         ReadModeAp
	ReadModeSc          ReadModeSc
	ConnectTimeout      int
	Replica             Replica
	SleepBetweenRetries int
	SocketTimeout       int
	TimeoutDelay        int
	TotalTimeout        int
	MaxRetries          int
	IncludeBinData      bool
	InfoTimeout         int
	RecordQueueSize     int
	ExpectedDuration    Duration
}

type Scan struct {
	ReadModeAp          ReadModeAp
	ReadModeSc          ReadModeSc
	ConnectTimeout      int
	Replica             Replica
	SleepBetweenRetries int
	SocketTimeout       int
	TimeoutDelay        int
	TotalTimeout        int
	MaxRetries          int
	ConcurrentNodes     int
	MaxConcurrentNodes  int
}

type BatchRead struct {
	ReadModeAp          ReadModeAp
	ReadModeSc          ReadModeSc
	ConnectTimeout      int
	FailOnFilteredOut   bool
	Replica             Replica
	SleepBetweenRetries int
	SocketTimeout       int
	TimeoutDelay        int
	TotalTimeout        int
	MaxRetries          int
	MaxConcurrentThread int
	AllowInline         bool
	AllowInlineSSD      bool
	RespondAllKeys      bool
}

type BatchWrite struct {
	ConnectTimeout      int
	FailOnFilteredOut   bool
	Replica             Replica
	SleepBetweenRetries int
	SocketTimeout       int
	TimeoutDelay        int
	TotalTimeout        int
	MaxRetries          int
	DurableDelete       bool
	SendKey             bool
	MaxConcurrentThread int
	AllowInline         bool
	AllowInlineSSD      bool
	RespondAllKeys      bool
}

type BatchUdf struct {
	DurableDelete bool
	SendKey       bool
}

type BatchDelete struct {
	DurableDelete bool
	SendKey       bool
}

type TxnRoll struct {
	ReadModeAp          ReadModeAp
	ReadModeSc          ReadModeSc
	ConnectTimeout      int
	Replica             Replica
	SleepBetweenRetries int
	SocketTimeout       int
	TimeoutDelay        int
	TotalTimeout        int
	MaxRetries          int
	MaxConcurrentThread int
	AllowInline         bool
	AllowInlineSSD      bool
	RespondAllKeys      bool
}

type TxnVerify struct {
	ReadModeAp          ReadModeAp
	ReadModeSc          ReadModeSc
	ConnectTimeout      int
	Replica             Replica
	SleepBetweenRetries int
	SocketTimeout       int
	TimeoutDelay        int
	TotalTimeout        int
	MaxRetries          int
	MaxConcurrentThread int
	AllowInline         bool
	AllowInlineSSD      bool
	RespondAllKeys      bool
}

type Metrics struct {
	Enable         bool
	LatencyShift   int
	LatencyColumns int
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
