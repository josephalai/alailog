package alailog

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"
	"time"
)

// Debugger is a structure that provides functions for debugging code.
type Debugger struct{}

// GetFunctionName returns the name of the function that is 'steps' frames up the call stack.
func (d *Debugger) GetFunctionName(steps int) string {
	pc, _, _, _ := runtime.Caller(steps)
	return runtime.FuncForPC(pc).Name()
}

func (l *Logger) PrintFunctionName(steps int) {
	fnName := l.GetFunctionName(steps + 1) // adding 1 to steps to account for this function call
	Printf("Function name %s\n", fnName)
}

// GetCallingFunctionName gets the name of the function that calls it.
func (d *Debugger) GetCallingFunctionName() string {
	return d.GetFunctionName(2) // 2 steps up the call stack to get the function that calls the function that calls this function
}

// GetFileAndLineNumber returns the file name and line number of the callers up the stack.
func (d *Debugger) GetFileAndLineNumber(steps int) (string, int) {
	_, file, line, _ := runtime.Caller(steps)
	return file, line
}

// PrintFileAndLineNumber prints the file name and line number of the callers up the stack.
func (l *Logger) PrintFileAndLineNumber(steps int) {
	file, line := l.GetFileAndLineNumber(steps + 1) // account for this function call
	Printf("File: %s, Line: %d\n", file, line)
}

// LogVar logs a variable name and the corresponding value.
func (l *Logger) LogVar(name string, value interface{}) {
	Printf("Variable: %s, Value: %v\n", name, value)
}

// ElapsedExecutionTime tracks the execution time of a function or segment of code.
func (d *Debugger) elapsedExecutionTime(start time.Time) time.Duration {
	return time.Since(start)
}

// ElapsedExecutionTime tracks the execution time of a function or segment of code.
func (d *Logger) ElapsedExecutionTime(start time.Time, name string) {
	elapsed := time.Since(start)
	Printf("%s took %v\n", name, elapsed)
}

// DebugMessage prints a detailed message including function name,
// file name and line number of the caller.
func (d *Debugger) DebugMessage(goSkip ...int) {
	skip := 2
	if goSkip != nil {
		skip += goSkip[0]
	}
	functionName := d.GetFunctionName(skip)
	file, line := d.GetFileAndLineNumber(skip) // 2 steps back in call stack to account for this function call
	Printf("Debug Message - Function name: %s, File: %s, Line: %d\n", functionName, file, line)
}

// GoLogger represents a logger instance that provides logging functionalities.
//
// The GoLogger struct has two fields:
//   - Instance: An instance of the Logger struct which handles the actual logging operations.
//   - doOnce: A sync.Once struct used to ensure that the Logger instance is created only once.
//
// The GoLogger type should be used in the following way:
//   - Declare a variable of type GoLogger and initialize it using the loggerInstance variable.
//     Example: logger := loggerInstance
//   - The GoLogger instance can then be used to perform various logging operations by calling the corresponding methods.
//     Example: logger.Debug("Debug message")
//
// Note: The Logger struct and its methods are not listed here, but they are required for the proper functioning of the GoLogger instance.
type GoLogger struct {
	Instance *Logger
	doOnce   sync.Once
}

// Parameter represents the configuration options for logging.
type Parameter struct {
	Filename        string
	Level           Level
	Stdout          bool
	Stderror        bool
	IsColored       bool
	TextColor       Color
	BgColor         Color
	Timestamps      bool
	TimestampFormat string
}

// DefaultFile is a constant that represents the default file name used for logging. By default, it is set to "logs.txt".
const DefaultFile = "logs.txt"

// loggerInstance represents a single instance of the GoLogger struct.
// Use this variable to access the logger instance.
// Example usage:
// createInstance(p) initializes the loggerInstance variable with a new Logger instance.
// GetInstance(p) returns the Logger instance stored in the loggerInstance variable.
// GoLogger declaration:
//
//	type GoLogger struct {
//	    Instance *Logger
//	    doOnce   sync.Once
//	}
//
// Logger declaration:
//
//	type Logger struct {
//	    file      *os.File
//	    level     Level
//	    stdout    bool
//	    stderr    bool
//	    color     bool
//	    textColor Color
//	    bgColor   Color
//	}
//
// SetTextColor(color Color) sets the text color of the logger.
// SetBgColor(color Color) sets the background color of the logger.
// Log(level Level, message interface{}) logs a message with the specified log level.
// LogColor(level Level, color Color, message interface{}) logs a colored message with the specified log level and color.
// Debug(message interface{}) logs a debug level message.
// Info(message interface{}) logs an info level message.
// Warn(message interface{}) logs a warning level message.
// Error(message interface{}) logs an error level message.
// Fatal(message interface{}) logs a fatal level message.
// DebugColor(color Color, message interface{}) logs a debug level message with the specified color.
// InfoColor(color Color, message interface{}) logs an info level message with the specified color.
// WarnColor(color Color, message interface{}) logs a warning level message with the specified color.
// ErrorColor(color Color, message interface{}) logs an error level message with the specified color.
// FatalColor(color Color, message interface{}) logs a fatal level message with the specified color.
// DebugBlack(message interface{}) logs a debug level message with black color.
// SetLevel(level Level) sets the logging level of the logger.
// SetStdout(stdout bool) sets whether to log to stdout.
// SetStderr(stderr bool) sets whether to log to stderr.
var loggerInstance GoLogger

// Open or create the log file with write-only permissions, append mode, and permission 0666
func createInstance(p *Parameter) {
	file, err := os.OpenFile(p.Filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	loggerInstance.doOnce.Do(func() {
		loggerInstance.Instance = NewLogger(
			file,
			p.Level,
			p.Stdout,
			p.Stderror,
			p.IsColored,
			p.BgColor,
			p.TextColor,
			p.Timestamps,
			p.TimestampFormat,
		)
	})
}

// initInstance initializes the logger by calling the createInstance function and passing the given parameter.
func initInstance(p *Parameter) {
	createInstance(p)
}

// GetInstance returns an instance of the Logger. It initializes the Logger by calling initInstance with the provided Parameter
// Parameters:
//   - p: an optional pointer to a Parameter object that is used to initialize the Logger
//
// Returns:
//   - a pointer to the Logger instance
func GetInstance(p ...*Parameter) *Logger {
	if len(p) > 0 {
		initInstance(p[0])
	} else {
		initInstance(getDefaultParameter())
	}
	return loggerInstance.Instance
}

// getDefaultParameter generates a pointer to a new Parameter instance with default values.
// The default Parameter instance includes:
// Filename:  DefaultFile, which is the default file to write logs to.
// Level:     Info, setting the default logging level to informational messages.
// Stdout:    true, enabling logging to the standard output.
// Stderror:  false, disabling logging to the standard error.
// IsColored: false, disabling colored logging.
// TextColor: White, setting the default text color of logs to White.
// BgColor:   BgBlack, setting the default background color of logs to Black.
func getDefaultParameter() *Parameter {
	return &Parameter{
		Filename:        DefaultFile,
		Level:           InfoLvl,
		Stdout:          true,
		Stderror:        false,
		IsColored:       false,
		TextColor:       White,
		BgColor:         BgBlack,
		Timestamps:      true,
		TimestampFormat: "2006/01/02 15:04:05",
	}
}

// Path: gpt-split/util/log/logger.go
/**
 * This package allows for custom logging. Features include:
 *  - Logging to a file
 *  - Logging to stdout
 *  - Logging to stderr
 *  - Logging certain colors
 *  - Logging certain levels
 */
type Logger struct {
	// The file to log to
	file *os.File
	// The level to log at
	level Level
	// Whether to log to stdout
	stdout bool
	// Whether to log to stderr
	stderr bool
	// Whether to log in color
	color bool

	textColor Color
	bgColor   Color

	Timestamps      bool
	TimestampFormat string

	DebugMode bool

	Debugger
}

// EnableDebugMode turns on the debug logs output.
func (l *Logger) EnableDebugMode() {
	l.DebugMode = true
}

// DisableDebugMode turns off the debug logs output.
func (l *Logger) DisableDebugMode() {
	l.DebugMode = false
}

// SetTextColor sets the text color of the logger to the specified color.
// The color must be of type Color, which is a string representing a color code.
// The color codes are defined in the Color declaration.
// Example usage: logger.SetTextColor(Red)
func (l *Logger) SetTextColor(color Color) {
	l.textColor = color
}

// SetBgColor sets the background color for the logger.
func (l *Logger) SetBgColor(color Color) {
	l.bgColor = color
}

// Level represents the logging level. It is an integer value.
//
// Levels can have the following values:
//   - 0: Debug level for fine-grained logging
//   - 1: Info level for general purpose logging
//   - 2: Warning level for potential issues or unexpected behavior
//   - 3: Error level for critical issues or failures
//
// The default level is Info.
type Level int

const (
	// Log everything
	AllLvl Level = iota
	// Log debug and above
	DebugLvl
	// Log info and above
	InfoLvl
	// Log warn and above
	WarnLvl
	// Log error and above
	ErrorLvl
	// Log fatal and above
	FatalLvl
	// Log nothing
	OffLvl
)

// NewLogger creates a new instance of Logger with the given parameters
//
//	file: the log file to write to
//	level: the logging level to use
//	stdout: whether to log to stdout
//	stderr: whether to log to stderr
//	color: whether to use color for log output
//	bgColor: background color
//	textColor: text color
//	Returns a pointer to the newly created Logger instance
func NewLogger(file *os.File, level Level, stdout bool, stderr bool, color bool, bgColor Color, textColor Color, timestamps bool, timestampFormat string) *Logger {
	return &Logger{
		file:            file,
		level:           level,
		stdout:          stdout,
		stderr:          stderr,
		color:           color,
		bgColor:         bgColor,
		textColor:       textColor,
		Timestamps:      timestamps,
		TimestampFormat: timestampFormat,
		DebugMode:       true,
	}
}

// EnableTimestamps enables the timestamps in the logging.
func (l *Logger) EnableTimestamps() {
	l.Timestamps = true
}

// DisableTimestamps disables the timestamps in the logging.
func (l *Logger) DisableTimestamps() {
	l.Timestamps = false
}

// SetTimestampFormat sets the format of the timestamps in the logging.
func (l *Logger) SetTimestampFormat(format string) {
	l.TimestampFormat = format
}

// Log logs a message at the specified level.
// If the logger is configured to log in color, it calls the LogColor method with the specified level, text color, and message.
// Otherwise, it appends a newline character to the message and checks if the level is greater than or equal to the logger's configured level.
// If it is, it writes the message to the configured file, stdout, and stderr.
// If the level is lower than the configured level, it returns without writing the message.
func (l *Logger) Log(level Level, messageStr interface{}) {
	var message = fmt.Sprintf("%s", messageStr)
	if l.Timestamps { // Check if timestamps are to be added
		timestampFormat := l.TimestampFormat
		if timestampFormat == "" {
			timestampFormat = "2006-01-02 15:04:05"
		}
		message = fmt.Sprintf("[%s] %s", time.Now().Format(timestampFormat), message)
	}

	if l.color {
		l.LogColor(level, l.textColor, message)
	} else {
		if level >= l.level {
			if l.file != nil {
				l.file.WriteString(message)
			}
			if l.stdout {
				os.Stdout.WriteString(message)
			}
			if l.stderr {
				os.Stderr.WriteString(message)
			}
		} else {
			return
		}
	}
}

// LogColor logs a message with a specified color. If the logger is configured to log in color, the message will be displayed with the specified color and background. Otherwise, the
func (l *Logger) LogColor(level Level, color Color, messageStr interface{}) {
	var message = fmt.Sprintf("%s", messageStr)
	if l.color {
		l.Log(level, l.bgColor.String()+color.String()+message+Reset.String())
	} else {
		l.Log(level, message)
	}
}

func (l *Logger) DebugLog(skip ...int) (is bool) {
	is = false
	if l.DebugMode {
		l.DebugMessage(skip...)
		is = true
	}
	return
}

// Debug sends a debug message to the logger.
// It calls the Log method of the logger with the Debug level and the provided message.
func (l *Logger) Debug(message interface{}) {
	if l.DebugLog(4) {
		l.Log(DebugLvl, message)
	}
}

// Info logs an info-level message to the logger. It
// calls the Log method passing the Info log level and
// the specified message.
func (l *Logger) Info(message interface{}) {
	l.Log(InfoLvl, message)
}

func (l *Logger) Logf(format string, args ...interface{}) {
	l.Log(InfoLvl, fmt.Sprintf(format, args...))
}

// PrintMap
func (l *Logger) PrintMap(m map[string]interface{}) {
	for k, v := range m {
		l.Log(InfoLvl, fmt.Sprintf("%s: %v", k, v))
	}
}

// Warn logs a warning message.
// It calls the Log method with the Warn level and the provided message.
func (l *Logger) Warn(message interface{}) {
	l.Log(WarnLvl, message)
}

// Error logs an error message.
//
// The message parameter is a string containing the error message to be logged.
// This method internally calls the Log method of the Logger instance, passing
// the Error log level and the provided message as arguments.
func (l *Logger) Error(message interface{}) {
	l.Log(ErrorLvl, message)
}

// Fatal logs a fatal message.
// It calls the Log method with the Fatal level and the provided message.
//
// Usage Example:
//
//	l := &Logger{}
//	l.Fatal("This is a fatal message")
//
// Parameters:
//
//	message (string): The message to be logged
//
// Returns:
//
//	None
func (l *Logger) Fatal(message interface{}) {
	l.Log(FatalLvl, message)
}

// Color represents a type for ANSI escape codes that define text and background colors.
type Color string

// String returns the string representation of the Color.
func (c Color) String() string {
	return string(c)
}

// BgBlack represents the background color black
const (
	BgBlack   Color = "\033[40m"
	BgRed     Color = "\033[41m"
	BgGreen   Color = "\033[42m"
	BgYellow  Color = "\033[43m"
	BgBlue    Color = "\033[44m"
	BgMagenta Color = "\033[45m"
	BgCyan    Color = "\033[46m"
	BgWhite   Color = "\033[47m"
)

// code now contains the ANSI escape code for the color red
const (
	Black   Color = "\033[1;30m"
	Red     Color = "\033[1;31m"
	Green   Color = "\033[1;32m"
	Yellow  Color = "\033[1;33m"
	Blue    Color = "\033[1;34m"
	Magenta Color = "\033[1;35m"
	Cyan    Color = "\033[1;36m"
	Purple  Color = "\033[1;35m"
	White   Color = "\033[1;37m"
	Reset   Color = "\033[0m"
)

// Code returns the ANSI escape code for the specified color.
func (c Color) Code() string {
	switch c {
	case Black:
		return "30"
	case Red:
		return "31"
	case Green:
		return "32"
	case Yellow:
		return "33"
	case Blue:
		return "34"
	case Magenta:
		return "35"
	case Cyan:
		return "36"
	case White:
		return "37"
	default:
		return ""
	}
}

// DebugColor is a method of the Logger struct that logs a message with a specific color at the Debug level.
// It calls the LogColor method of the Logger struct, passing the Debug level, the specified color, and the message.
//
// Example usage:
//
//	logger.DebugColor(Red, "This is a debug message with red color")
//
// Parameters:
// - color: The color to apply to the message.
// - message: The message to be logged.
//
// Note: The color parameter should be one of the Color constants defined in the Color struct.
func (l *Logger) DebugColor(color Color, message interface{}) {
	l.LogColor(DebugLvl, color, message)
}

// InfoColor logs a message with the specified color at the Info level.
// The "color" parameter specifies the color of the message.
// The "message" parameter is the message to be logged.
func (l *Logger) InfoColor(color Color, message interface{}) {
	l.LogColor(InfoLvl, color, message)
}

// WarnColor sets the warning level log message with the specified color and message.
// It calls the LogColor method with the Level parameter set to Warn, the color parameter, and the message parameter.
func (l *Logger) WarnColor(color Color, message interface{}) {
	l.LogColor(WarnLvl, color, message)
}

// ErrorColor sets the text color for error messages and logs the message with the specified color.
// It delegates the logging to the LogColor method, passing the Error level, specified color, and the message.
// Example usage:
//
//	logger.ErrorColor(Red, "This is an error message")
func (l *Logger) ErrorColor(color Color, message interface{}) {
	l.LogColor(ErrorLvl, color, message)
}

// FatalColor sets the text color of the logger to the specified color and logs a fatal message with the specified color.
// The color must be of type Color, which is a string representing a color code.
// The color codes are defined in the Color declaration.
// Example usage: logger.FatalColor(Red, "Fatal error occurred")
// This method is used in the TestLogger_FatalColor function for testing purposes.
func (l *Logger) FatalColor(color Color, message interface{}) {
	l.LogColor(FatalLvl, color, message)
}

// DebugBlack logs a debug message with black text color.
//
// It calls the DebugColor method with the Black color and the given message.
// The debug message will only be logged if the logger's color is enabled.
//
// Example:
//
//	l := &Logger{}
//	l.DebugBlack("Debug message with black text color")
func (l *Logger) DebugBlack(message interface{}) {
	l.DebugColor(Black, message)
}

// SetLevel sets the log level of the Logger
func (l *Logger) SetLevel(level Level) {
	l.level = level
}

// SetStdout sets whether to log to stdout or not.
func (l *Logger) SetStdout(stdout bool) {
	l.stdout = stdout
}

// SetStderr enables or disables logging to standard error output
// based on the value of stderr parameter, which must be of type bool.
// If stderr is set to true, the logger will write log messages to the standard
// error output. If stderr is set to false, log messages will not be written to the
// standard error output.
// Example usage: logger.SetStderr(true)
func (l *Logger) SetStderr(stderr bool) {
	l.stderr = stderr
}

// Infof formats and logs a message at the info level with the specified format and arguments.
// The format argument is a string that specifies the format of the message.
// The args argument is a variadic parameter that represents the values to be formatted in the message.
// The formatted message is logged using the Info method of the logger.
// Example usage: logger.Infof("Received %d bytes", size)
func (l *Logger) Infof(format string, args ...interface{}) {
	l.Info(fmt.Sprintf(format, args...))
}

// Debugf formats a message according to the given format specifier and
// calls the Debug method to log it. It accepts a format string and any number
// of arguments to be formatted in the message.
// Example usage: logger.Debugf("This is a debug message: %v", variable)
func (l *Logger) Debugf(format string, args ...interface{}) {
	fmt.Println(fmt.Sprintf(format, args...))
}

// Printf formats the string according to the specified format string
// and writes the formatted string to the log with the Info level.
// The format string can contain verbs that are formatting instructions.
// The verbs are derived from the formats supported by the fmt package.
// The args parameter is a variadic argument of type interface{},
// which allows passing any number of arguments of any type.
// Example usage: logger.Printf("Value of x is %d and y is %f", x, y)
func (l *Logger) Printf(format string, args ...interface{}) {
	l.Infof(format, args)
}

// Errorf formats and logs an error message with the specified format and arguments.
// It uses the fmt.Sprintf function to format the message and then calls the Error method to log it.
// The format parameter is a string that specifies the format of the message.
// The args parameter is a variadic argument that represents the values to be formatted according to the format string.
// Example usage: logger.Errorf("Something went wrong
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.Error(fmt.Sprintf(format, args...))
}

func (l *Logger) Warningf(format string, args ...interface{}) {
	l.Warn(fmt.Sprintf(format, args...))
}

// Fatalf formats the given string according to the specified format and arguments,
// then passes it to the Fatal method of the logger.
// The Fatal method logs the message as a fatal error, causing the program to exit.
// This method is useful for logging critical errors that require immediate action.
// The format argument is a string specifying the format of the message, while
// the args argument is a variadic slice of interface{} representing the arguments
// to be formatted in the message.
//
// Example usage:
// logger.Fatalf("Critical error: %s", err)
// The above code will log the formatted error message as a fatal error,
// causing the program to exit immediately.
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.Fatal(fmt.Sprintf(format, args...))
}

func (l *Logger) Infoln(args ...interface{}) {
	l.Info(fmt.Sprintln(args...))
}

func (l *Logger) Println(args ...interface{}) {
	l.Infoln(fmt.Sprintln(args...))
}

func (l *Logger) Errorln(args ...interface{}) {
	l.Error(fmt.Sprintln(args...))
}

func (l *Logger) Warningln(args ...interface{}) {
	l.Warn(fmt.Sprintln(args...))
}

// Fatalln logs a fatal message followed by a newline.
// The message is formatted using fmt.Sprintln(args...),
// which concatenates the arguments using spaces and adds a newline character at the end.
// The formatted message is then passed to the Logger's Fatal method to log the message as a fatal error.
//
// Example usage:
//
//	logger.Fatalln("Something went wrong")
//
// This will log the message "Something went wrong" as a fatal error.
func (l *Logger) Fatalln(args ...interface{}) {
	l.Fatal(fmt.Sprintln(args...))
}

// Info logs the provided arguments using the logger instance.
// The logger instance is obtained by calling GetInstance() function.
// The arguments are converted to a string using fmt.Sprint() function
// and then passed to the logger's Info method.
func Info(args ...interface{}) {
	logger := GetInstance()
	logger.Info(fmt.Sprint(args...))
}

// Example usage:
// Infof("User %s logged in", username)
// The format parameter specifies the format of the log message
// args specifies optional arguments to format into the log message.
func Infof(format string, args ...interface{}) {
	logger := GetInstance()
	logger.Infof(format, args...)
}

func Infoln(args ...interface{}) {
	logger := GetInstance()
	logger.Infoln(args...)
}

// Print writes the arguments to the logger at the Info level
//
// Usage:
//
//	Print(args ...interface{})
//
// Arguments:
//   - args: The values to be printed
func Print(args ...interface{}) {
	logger := GetInstance()
	logger.Info(fmt.Sprint(args...))
}

// Printf writes a formatted string to the log file using the logger's
func Printf(format string, args ...interface{}) {
	logger := GetInstance()
	logger.Infof(format, args...)
}

// Println prints the given arguments to the log file, stdout, and stderr using the Infoln method of the logger instance obtained from GetInstance.
// Parameters:
//   - args: The values to be printed
//
// Example:
//
//	Println("Hello", "World") will print "Hello World" to the log file, stdout, and stderr.
func Println(args ...interface{}) {
	logger := GetInstance()
	logger.Infoln(args...)
}

// Error writes an error log message.
// It retrieves an instance of the logger using GetInstance() function.
// It then calls the Error() method of the logger, passing in the provided arguments as a string.
// The Error() method of the logger writes the log message to the log file,
// and optionally to stdout and stderr based on the logger configuration.
// If color logging is
func Error(args ...interface{}) {
	logger := GetInstance()
	logger.Error(fmt.Sprint(args...))
}

// Errorf formats the error message with the given format string and optional arguments using Logger's Error method. It delegates the logging operation to the GetInstance function to
func Errorf(format string, args ...interface{}) {
	logger := GetInstance()
	logger.Errorf(format, args...)
}

// Errorln calls the Errorln method of the logger instance to log a message with level Error.
// The args parameter holds the
func Errorln(args ...interface{}) {
	logger := GetInstance()
	logger.Errorln(args...)
}

// Warning logs a warning message.
// It retrieves the logger instance using GetInstance() function.
// The logger instance is used to call the Warn() method,
// which logs the message with the warning level.
// The message is formatted using fmt.Sprint(args...) and passed to the Warn() method
func Warning(args ...interface{}) {
	logger := GetInstance()
	logger.Warn(fmt.Sprint(args...))
}

// Warningf formats and logs a warning message with the given format and arguments
func Warningf(format string, args ...interface{}) {
	logger := GetInstance()
	logger.Warningf(format, args...)
}

// Warningln calls the Warningln method of the logger instance to log a warning message using the default logger configuration
func Warningln(args ...interface{}) {
	logger := GetInstance()
	logger.Warningln(args...)
}

// Fatal logs a message with the "Fatal" log level using the logger instance.
// It formats the message using fmt.Sprint(args...) and calls the logger's Fatal method with the formatted message.
// The message will be written to the log file, stdout, and stderr if the log level is equal to or higher than the logger's level.
// The log message can include timestamps if enabled, formatted according to the logger's TimestampFormat value.
// The logger instance is retrieved using the GetInstance function.
// The logger instance
func Fatal(args ...interface{}) {
	logger := GetInstance()
	logger.Fatal(fmt.Sprint(args...))
}

// Fatalf logs a formatted message at the Fatal level and then terminates the program.
// `format` is the format string for the log message.
//
// `args` are the arguments to format the log message.
//
// It retrieves the logger instance using the GetInstance() function and
// calls the logger's Fatalf() method to log the formatted message and exit the program.
func Fatalf(format string, args ...interface{}) {
	logger := GetInstance()
	logger.Fatalf(format, args...)
}

// Fatalln logs a message at the Fatal level and terminates the program.
// It uses the GetInstance function to retrieve the logger instance.
// The logger instance's Fatalln method is then called with the passed args.
func Fatalln(args ...interface{}) {
	logger := GetInstance()
	logger.Fatalln(args...)
}

func (l *Logger) Debugln(args ...interface{}) {
	if l.DebugLog(4) {
		l.Println(fmt.Sprintln(args...))
	}
}

func Debug(args ...interface{}) {
	logger := GetInstance()
	if logger.DebugLog(4) {
		logger.Debug(args)
	}
}

func Debugf(format string, args ...interface{}) {
	logger := GetInstance()
	if logger.DebugLog(4) {
		logger.Debugf(format, args...)
	}
}

// Debugln writes the debug message to the logger's output.
// The debug message is constructed by passing the args to fmt.Sprintln.
//
// Example usage:
// logger.Debugln("This is a debug message")
func Debugln(args ...interface{}) {
	logger := GetInstance()
	if logger.DebugLog(4) {
		logger.Debugln(args...)
	}
}

// EnableDebugMode turns on the debug logs output.
func EnableDebugMode() {
	logger := GetInstance()
	logger.DebugMode = true
}

// DisableDebugMode turns off the debug logs output.
func DisableDebugMode() {
	logger := GetInstance()
	logger.DebugMode = false
}
