package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

type MonitorServerShell string

type MonitorInfo struct {
	// Type of Server that you want to monitor
	ServerType MonitorServerType
	// Uri of Server that you want to monitor
	ServerUri MonitorServerUri
	// Resources you want to monitor
	Resources []MonitorResource
	// File name of log file for the server
	LogFile string
	// Server Command Shell
	ServerCommandShell MonitorServerShell
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
			ServerType:         MonitorLocalServer,
			ServerUri:          "local",
			Resources:          nil,
			LogFile:            "monitor_error.log",
			ServerCommandShell: nil,
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
	// UI Server
	gssUiServerPort, gssUiServerPortExist := os.LookupEnv("GSS_UI_SERVER_PORT")
	gssUiServerLogfile, gssUiServerLogfileExist := os.LookupEnv("GSS_UI_SERVER_LOGFILE")
	gssUiServerVerboselevel, gssUiServerVerboselevelExist := os.LookupEnv("GSS_UI_SERVER_VERBOSE")
	// API Server
	gssApiServerPort, gssApiServerPortExist := os.LookupEnv("GSS_API_SERVER_PORT")
	gssApiServerLogfile, gssApiServerLogfileExist := os.LookupEnv("GSS_API_SERVER_LOGFILE")
	gssApiServerVerboselevel, gssApiServerVerboselevelExist := os.LookupEnv("GSS_API_SERVER_VERBOSE")
	// Monitor Server
	gssMonitorServerType, gssMonitorServerTypeExist := os.LookupEnv("GSS_MONITOR_SERVER_TYPE")
	gssMonitorServerUri, gssMonitorServerUriExist := os.LookupEnv("GSS_MONITOR_SERVER_URI")
	gssMonitorServerLogfile, gssMonitorServerLogfileExist := os.LookupEnv("GSS_MONITOR_SERVER_LOGFILE")
	// Log Server
	gssLogServerEnabled, gssLogServerEnabledExist := os.LookupEnv("GSS_LOG_SERVER_ENABLED")
	gssLogServerVerboselevel, gssLogServerVerboselevelExist := os.LookupEnv("GSS_LOG_SERVER_VERBOSE")
	gssLogServerLogfile, gssLogServerLogfileExist := os.LookupEnv("GSS_LOG_SERVER_LOGFILE")
	// Alert Server
	gssAlertServerEnabled, gssAlertServerEnabledExist := os.LookupEnv("GSS_ALERT_SERVER_ENABLED")
	gssAlertServerLogfile, gssAlertServerLogfileExist := os.LookupEnv("GSS_ALERT_SERVER_LOGFILE")
	// UI Server Environment Variables Setup
	if gssUiServerPortExist {
		newConfig.UiServer.Port = strings.TrimSpace(gssUiServerPort)
	}
	if gssUiServerLogfileExist {
		newConfig.UiServer.LogFile = strings.TrimSpace(gssUiServerLogfile)
	}
	if gssUiServerVerboselevelExist {
		gssUiServerVerboselevel = strings.TrimSpace(gssUiServerVerboselevel)
		gssUiServerVerboselevelUint, err := strconv.ParseUint(gssUiServerVerboselevel, 10, 8)
		if err == nil {
			newConfig.UiServer.LogVerbose = VerboseLevel(gssUiServerVerboselevelUint)
		} else {
			gssError := fmt.Errorf("configuration: %s | Error: %s", "GSS_UI_SERVER_VERBOSE", err)
			fmt.Println(gssError.Error())
		}
	}
	// API Server Environment Variables Setup
	if gssApiServerPortExist {
		newConfig.ApiServer.Port = strings.TrimSpace(gssApiServerPort)
	}
	if gssApiServerLogfileExist {
		newConfig.ApiServer.LogFile = strings.TrimSpace(gssApiServerLogfile)
	}
	if gssApiServerVerboselevelExist {
		gssApiServerVerboselevel = strings.TrimSpace(gssApiServerVerboselevel)
		gssApiServerVerboselevelUint, err := strconv.ParseUint(gssApiServerVerboselevel, 10, 8)
		if err == nil {
			newConfig.ApiServer.LogVerbose = VerboseLevel(gssApiServerVerboselevelUint)
		} else {
			gssError := fmt.Errorf("configuration: %s | Error: %s", "GSS_API_SERVER_VERBOSE", err)
			fmt.Println(gssError.Error())
		}
	}
	// Monitor Server Environment Variables Setup
	if gssMonitorServerTypeExist {
		newConfig.Monitor.ServerType = MonitorServerType(strings.TrimSpace(gssMonitorServerType))
	}
	if gssMonitorServerUriExist {
		newConfig.Monitor.ServerUri = MonitorServerUri(strings.TrimSpace(gssMonitorServerUri))
	}
	if gssMonitorServerLogfileExist {
		newConfig.Monitor.LogFile = strings.TrimSpace(gssMonitorServerLogfile)
	}
	// Log Server Environment Variables Setup
	if gssLogServerEnabledExist {
		gssLogServerEnabledBool, err := strconv.ParseBool(strings.TrimSpace(gssLogServerEnabled))
		if err == nil {
			newConfig.Logs.Enabled = gssLogServerEnabledBool
		} else {
			gssError := fmt.Errorf("configuration: %s | Error: %s", "GSS_LOG_SERVER_ENABLED", err)
			fmt.Println(gssError.Error())
		}
	}
	if gssLogServerLogfileExist {
		newConfig.Logs.LogFile = strings.TrimSpace(gssLogServerLogfile)
	}
	if gssLogServerVerboselevelExist {
		gssLogServerVerboselevel = strings.TrimSpace(gssLogServerVerboselevel)
		gssLogServerVerboselevelUint, _ := strconv.ParseUint(gssLogServerVerboselevel, 10, 8)
		newConfig.Logs.Verbose = VerboseLevel(gssLogServerVerboselevelUint)
	}
	// Alert Server Environment Variables Setup
	if gssAlertServerEnabledExist {
		gssAlertServerEnabledBool, err := strconv.ParseBool(strings.TrimSpace(gssAlertServerEnabled))
		if err == nil {
			newConfig.Alerts.Enabled = gssAlertServerEnabledBool
		} else {
			gssError := fmt.Errorf("configuration: %s | Error: %s", "GSS_ALERT_SERVER_ENABLED", err)
			fmt.Println(gssError.Error())
		}
	}
	if gssAlertServerLogfileExist {
		newConfig.Alerts.LogFile = strings.TrimSpace(gssAlertServerLogfile)
	}
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
