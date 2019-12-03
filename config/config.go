package config

type VerboseLevel uint8

const (
	// Debug Verbose Level
	VerboseLevelDebug = VerboseLevel(255)
	// Normal Verbose Level
	VerboseLevelVerbose = VerboseLevel(2)
	// Warnings Verbose Level
	VerboseLevelWarnings = VerboseLevel(1)
	// Error Verbose Level
	VerboseLevelError = VerboseLevel(0)
)

type ServerConfiguration struct {
	// Port of the server
	Port string
	// File name of log file for the server
	LogFile string
	// Verbose Level
	LogVerbose VerboseLevel
}

type AppInfo struct {
	// Version of App
	Version string
	// Version Array of App
	VersionArray [4]int
}

// Alert Value
// Type: String/Int
type AlertValue interface{}

type Alert struct {
	// Name of Monitored Resource to make alert for
	NameOfMonitoredResource string
	// Type of Alert
	Type string
	// Verbose Level of Alert
	Verbose VerboseLevel
	// Alert Value
	AlertValue interface{}
}

type AlertInfo struct {
	// Is Alerts Enabled
	Enabled bool
	// Alerts
	Alerts []Alert
	// File name of log file for the server
	LogFile string
}

type LogType string

const (
	// Standard Output
	LogTypeStandardOutput = LogType("STDOUT")
	// Standard Error
	LogTypeStandardError = LogType("STDERR")
	// Standard Output and Standard Error
	LogTypeStandardAny = LogType("STDANY")
	// Docker Log
	LogTypeDockerLog = LogType("DOCKER")
	// Command Output
	// With Return Code and STDANY
	LogTypeCommandOutput = LogType("COMMAND")
)

type LogSource struct {
	// Name of Monitored Resource to make alert for
	NameOfMonitoredResource string
	// Type of Log
	Type LogType
	// Verbose Level of Log
	Verbose VerboseLevel
}

type LogInfo struct {
	// Is Logs Enabled
	Enabled bool
	// Verbose Level of Logging
	Verbose VerboseLevel
	// Sources of Logs
	Sources []LogSource
	// File name of log file for the server
	LogFile string
}

type MonitorServerType string

const (
	// Remote Server (using SSH)
	MonitorRemoteServer = MonitorServerType("REMOTE")
	// Local Server
	MonitorLocalServer = MonitorServerType("LOCAL")
	// Docker Server
	MonitorDockerServer = MonitorServerType("DOCKER")
)

type MonitorServerUri string

type MonitorResourceType string

const (
	MonitorResourceLog    = MonitorResourceType("log")
	MonitorResourceString = MonitorResourceType("string")
	MonitorResourceNumber = MonitorResourceType("number") //float
	MonitorResourceBool   = MonitorResourceType("bool")
)

type MonitorResourceCommand string

type MonitorResourceRegexString string

type MonitorResource struct {
	// Name of Monitored Resource
	Name string
	// Type of Monitored Resource
	Type MonitorResourceType
	// Command
	Command MonitorResourceCommand
	// Parser Regex
	ParserRegex MonitorResourceRegexString
}

type MonitorInfo struct {
	// Type of Server that you want to monitor
	ServerType MonitorServerType
	// Uri of Server that you want to monitor
	ServerUri MonitorServerUri
	// Resources you want to monitor
	Resources []MonitorResource
	// File name of log file for the server
	LogFile string
}

type Configuration struct {
	// Log Server Configuration
	UiServer ServerConfiguration
	// Admin Server Configuration
	ApiServer ServerConfiguration
	// App Information
	App AppInfo
	// Alerts Information
	Alerts AlertInfo
	// Log Information
	Logs LogInfo
	// Monitor Information
	Monitor MonitorInfo
}

func FillConfigWithDefaults() Configuration {
	return Configuration{
		UiServer: ServerConfiguration{
			Port:       "7790",
			LogFile:    "ui_error.log",
			LogVerbose: VerboseLevelVerbose,
		},
		ApiServer: ServerConfiguration{
			Port:       "7791",
			LogFile:    "api_error.log",
			LogVerbose: VerboseLevelVerbose,
		},
		App: AppInfo{
			Version:      "0.1.0.0",
			VersionArray: [4]int{0, 1, 0, 0},
		},
		Monitor: MonitorInfo{
			ServerType: MonitorLocalServer,
			ServerUri:  "local",
			Resources:  nil,
			LogFile:    "monitor_error.log",
		},
		Logs: LogInfo{
			Enabled: true,
			Verbose: VerboseLevelVerbose,
			Sources: nil,
			LogFile: "messages.log",
		},
		Alerts: AlertInfo{
			Enabled: false,
			Alerts:  nil,
			LogFile: "alerts_error.log",
		},
	}
}

func FillConfigWithEnvironmentVars() Configuration {
	// Make Config from Default Values
	newConfig := FillConfigWithDefaults()
	// Get All Environment Variables
	//gelsLogServerPort, gelsLogServerPortExist := os.LookupEnv("GELS_LOG_SERVER_PORT")
	//gelsLogServerLogfile, gelsLogServerLogfileExist := os.LookupEnv("GELS_LOG_SERVER_LOGFILE")
	// Log Server Environment Variables Setup
	//if gelsLogServerPortExist {
	//	newConfig.LogServer.Port = strings.TrimSpace(gelsLogServerPort)
	//}
	//if gelsLogServerLogfileExist {
	//	newConfig.LogServer.LogFile = strings.TrimSpace(gelsLogServerLogfile)
	//}
	// Return Changed Config
	return newConfig
}

var config = FillConfigWithDefaults()

func SetConfig(newConfig Configuration) {
	config = newConfig
}

func GetConfig() Configuration {
	return config
}

func InitConfig() Configuration {
	SetConfig(FillConfigWithEnvironmentVars())
	return GetConfig()
}
