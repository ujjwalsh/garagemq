package config

// Config represents server changeable se
type Config struct {
	Users      []User
	TCP        TCPConfig
	Queue      Queue
	Db         Db
	Vhost      Vhost
	Security   Security
	Connection Connection
}

// User for auth check
type User struct {
	Username string
	Password string
}

// TCPConfig represents properties for tune network connections
type TCPConfig struct {
	IP           string `yaml:"ip"`
	Port         string
	Nodelay      bool
	ReadBufSize  int `yaml:"readBufSize"`
	WriteBufSize int `yaml:"writeBufSize"`
}

// Queue settings
type Queue struct {
	ShardSize        int    `yaml:"shardSize"`
	MaxMessagesInRam uint64 `yaml:"maxMessagesInRam"`
}

// Db settings, such as path to load/save and engine
type Db struct {
	DefaultPath string `yaml:"defaultPath"`
	Engine      string `yaml:"engine"`
}

// Vhost settings
type Vhost struct {
	DefaultPath string `yaml:"defaultPath"`
}

// Security settings
type Security struct {
	PasswordCheck string `yaml:"passwordCheck"`
}

// Connection settings for AMQP-connection
type Connection struct {
	ChannelsMax  uint16 `yaml:"channelsMax"`
	FrameMaxSize uint32 `yaml:"frameMaxSize"`
}
