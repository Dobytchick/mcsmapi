package mcsmapi

// PanelData represents the root structure of the MCSManager panel data
type PanelData struct {
	Status int   `json:"status"` // HTTP status code
	Data   Data  `json:"data"`   // Main data payload
	Time   int64 `json:"time"`   // Timestamp of the response
}

// Data contains all core panel information and metrics
type Data struct {
	Version                string      `json:"version"`                // Panel software version
	SpecifiedDaemonVersion string      `json:"specifiedDaemonVersion"` // Required daemon version
	Process                Process     `json:"process"`                // Panel process information
	Record                 Record      `json:"record"`                 // Access statistics
	System                 System      `json:"system"`                 // Host system information
	Chart                  Chart       `json:"chart"`                  // Performance charts data
	RemoteCount            RemoteCount `json:"remoteCount"`            // Daemon connection stats
	Remote                 []Remote    `json:"remote"`                 // List of connected daemons
}

// Process contains information about the panel process
type Process struct {
	CPU    int    `json:"cpu"`    // CPU usage percentage
	Memory int    `json:"memory"` // Memory usage in bytes (Panel Memory Usage)
	Cwd    string `json:"cwd"`    // Current working directory
}

// Record tracks authentication and security events
type Record struct {
	Logined       int `json:"logined"`       // Successful logins count
	IllegalAccess int `json:"illegalAccess"` // Unauthorized access attempts
	Banips        int `json:"banips"`        // Banned IP addresses count
	LoginFailed   int `json:"loginFailed"`   // Failed login attempts
}

// System contains detailed host system information
type System struct {
	User     User    `json:"user"`     // System user information
	Time     int64   `json:"time"`     // Timestamp (Memory usage on the panel)
	Totalmem int64   `json:"totalmem"` // Total system memory in bytes
	Freemem  int64   `json:"freemem"`  // Free memory in bytes
	Type     string  `json:"type"`     // OS type (Windows_NT/Linux)
	Version  string  `json:"version"`  // OS version
	Node     string  `json:"node"`     // Node.js version
	Hostname string  `json:"hostname"` // System hostname
	Loadavg  []int   `json:"loadavg"`  // System load averages (Linux only)
	Platform string  `json:"platform"` // Platform identifier
	Release  string  `json:"release"`  // Kernel release version
	Uptime   float64 `json:"uptime"`   // System uptime in seconds
	CPU      float64 `json:"cpu"`      // Overall CPU usage percentage
}

// User contains system user account information
type User struct {
	UID      int    `json:"uid"`      // User ID
	GID      int    `json:"gid"`      // Group ID
	Username string `json:"username"` // Username
	Homedir  string `json:"homedir"`  // Home directory path
	Shell    any    `json:"shell"`    // Shell path (nullable)
}

// Chart contains performance metrics for visualization
type Chart struct {
	System  []SystemChart  `json:"system"`  // System resource charts (Memory&CPU usage on the panel)
	Request []RequestChart `json:"request"` // Request statistics charts
}

// SystemChart represents a single data point for system metrics
type SystemChart struct {
	CPU float64 `json:"cpu"` // CPU usage percentage
	Mem float64 `json:"mem"` // Memory usage percentage
}

// RequestChart represents application request metrics
type RequestChart struct {
	Value           int `json:"value"`           // Current active requests
	TotalInstance   int `json:"totalInstance"`   // Total managed instances
	RunningInstance int `json:"runningInstance"` // Currently running instances
}

// RemoteCount tracks connected daemon statistics
type RemoteCount struct {
	Available int `json:"available"` // Number of available daemons
	Total     int `json:"total"`     // Total registered daemons
}

// Remote represents a connected daemon instance
type Remote struct {
	Version     string        `json:"version"`     // Daemon software version
	Process     DaemonProcess `json:"process"`     // Daemon process info
	Instance    Instance      `json:"instance"`    // Managed instances info
	System      DaemonSystem  `json:"system"`      // Daemon host system info (CPU and memory usage on the Daemon)
	CPUMemChart []CPUMemChart `json:"cpuMemChart"` // Performance metrics (CPU and memory usage on the Daemon Chart)
	UUID        string        `json:"uuid"`        // Unique daemon identifier
	IP          string        `json:"ip"`          // Connection IP/hostname
	Port        int           `json:"port"`        // Connection port
	Prefix      string        `json:"prefix"`      // URL prefix
	Available   bool          `json:"available"`   // Connection status
	Remarks     string        `json:"remarks"`     // Descriptive label
}

// DaemonProcess contains daemon-specific process information
type DaemonProcess struct {
	CPU    int    `json:"cpu"`    // CPU usage in nanoseconds
	Memory int    `json:"memory"` // Memory usage in bytes
	Cwd    string `json:"cwd"`    // Current working directory
}

// Instance contains Minecraft instance statistics
type Instance struct {
	Running int `json:"running"` // Number of running instances
	Total   int `json:"total"`   // Total configured instances
}

// DaemonSystem contains detailed daemon host system info
type DaemonSystem struct {
	Type       string    `json:"type"`       // OS type
	Hostname   string    `json:"hostname"`   // System hostname
	Platform   string    `json:"platform"`   // Platform identifier
	Release    string    `json:"release"`    // Kernel release version
	Uptime     float64   `json:"uptime"`     // System uptime in seconds
	Cwd        string    `json:"cwd"`        // Current working directory
	Loadavg    []float64 `json:"loadavg"`    // System load averages
	Freemem    int64     `json:"freemem"`    // Free memory in bytes
	CPUUsage   float64   `json:"cpuUsage"`   // Current CPU usage
	MemUsage   float64   `json:"memUsage"`   // Current memory usage
	Totalmem   int64     `json:"totalmem"`   // Total memory in bytes
	ProcessCPU float64   `json:"processCpu"` // Process CPU usage
	ProcessMem float64   `json:"processMem"` // Process memory usage
}

// CPUMemChart represents a single performance data point
type CPUMemChart struct {
	CPU float64 `json:"cpu"` // CPU usage percentage
	Mem float64 `json:"mem"` // Memory usage percentage
}
