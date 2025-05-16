package mcsmapi

type InstanceStatus int

const (
	InstanceStop     InstanceStatus = 0
	InstanceStarting InstanceStatus = 1
	InstanceRunning  InstanceStatus = 2
	// InstanceStopping Maybe exists status idk ¯\_(ツ)_/¯
	InstanceStopping InstanceStatus = 3
)

type InstanceConfig struct {
	Nickname          string        `json:"nickname"`
	StartCommand      string        `json:"startCommand"`
	StopCommand       string        `json:"stopCommand"`
	Cwd               string        `json:"cwd"`
	IE                string        `json:"ie"`
	OE                string        `json:"oe"`
	CreateDatetime    int64         `json:"createDatetime"`
	LastDatetime      int64         `json:"lastDatetime"`
	Type              string        `json:"type"`
	Tag               []string      `json:"tag"`
	EndTime           int64         `json:"endTime"`
	FileCode          string        `json:"fileCode"`
	ProcessType       string        `json:"processType"`
	UpdateCommand     string        `json:"updateCommand"`
	ActionCommandList []string      `json:"actionCommandList"`
	CrLf              int           `json:"crlf"`
	Docker            *DockerConfig `json:"docker"`

	// Steam RCON configuration
	EnableRcon   bool   `json:"enableRcon"`
	RconPassword string `json:"rconPassword"`
	RconPort     int    `json:"rconPort"`
	RconIp       string `json:"rconIp"`

	// Terminal options
	TerminalOption struct {
		HaveColor bool `json:"haveColor"`
		Pty       bool `json:"pty"`
	} `json:"terminalOption"`

	// Event task settings
	EventTask struct {
		AutoStart   bool `json:"autoStart"`
		AutoRestart bool `json:"autoRestart"`
		Ignore      bool `json:"ignore"`
	} `json:"eventTask"`

	// Ping config
	PingConfig struct {
		IP   string `json:"ip"`
		Port int    `json:"port"`
		Type int    `json:"type"`
	} `json:"pingConfig"`
}

// InstanceDetail represents detailed information about an instance.
type InstanceDetail struct {
	Config *InstanceConfig `json:"config"`
	Info   struct {
		CurrentPlayers int      `json:"currentPlayers"`
		FileLock       int      `json:"fileLock"`
		MaxPlayers     int      `json:"maxPlayers"`
		OpenFrpStatus  bool     `json:"openFrpStatus"`
		PlayersChart   []string `json:"playersChart"`
		Version        string   `json:"version"`
	} `json:"info"`
	InstanceUUID string `json:"instanceUuid"`
	ProcessInfo  struct {
		CPU       int   `json:"cpu"`
		Memory    int   `json:"memory"`
		PPID      int   `json:"ppid"`
		PID       int   `json:"pid"`
		CTime     int64 `json:"ctime"`
		Elapsed   int64 `json:"elapsed"`
		Timestamp int64 `json:"timestamp"`
	} `json:"processInfo"`
	Space   int            `json:"space"`
	Started int            `json:"started"`
	Status  InstanceStatus `json:"status"` // Instance status (0: stopped, 1: starting, 2: running, etc.)
}

// DockerConfig represents the configuration for Docker-based instances.
type DockerConfig struct {
	ContainerName  string   `json:"containerName"`
	Image          string   `json:"image"`
	Memory         int      `json:"memory"` // MB
	Ports          []string `json:"ports"`
	ExtraVolumes   []string `json:"extraVolumes"`
	MaxSpace       int      `json:"maxSpace"`
	Network        int      `json:"network"`
	IO             int      `json:"io"`
	NetworkMode    string   `json:"networkMode"`
	NetworkAliases []string `json:"networkAliases"`
	CpuSetCpus     string   `json:"cpusetCpus"`
	CPUUsage       int      `json:"cpuUsage"`
	WorkingDir     string   `json:"workingDir"`
	Env            []string `json:"env"`
}
