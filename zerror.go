// Package zerror is an enhancement error library
package zerror

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
	"time"
)

// Enum type Severity is the message severity level: Fatal, Warn, and Info
type Severity int

const (
	SeverityFatal Severity = 0 + iota
	SeverityWarn
	SeverityInfo
)

var severities = [...]string{
	"Fatal",
	"Warn",
	"Info",
}

func (s Severity) String() string {
	return severities[s]
}

var defaultSeverity = SeverityInfo

func SetDefaultSeverity(s Severity) {
	defaultSeverity = s
}

func GetDefaultSeverity() Severity {
	return defaultSeverity
}

// time format in the error message
var timeFormat = time.RFC3339

func SetTimeFormat(format string) {
	timeFormat = format
}

func GetTimeFormat() string {
	return timeFormat
}

// Constant Notations are for formatting the error printable string
const (
	NotationTime     = "{time}"
	NotationSeverity = "{severity}"
	NotationMessage  = "{message}"
	NotationFunc     = "{func}"
	NotationLine     = "{line}"
	NotationFile     = "{file}"
)

// global format should compose of any of the notation options
var messageFormat = fmt.Sprintf("%s %s: %s(%s:%s)",
	NotationTime, NotationSeverity, NotationMessage, NotationFunc, NotationLine)

func SetMessageFormat(format string) {
	messageFormat = format
}

func GetMessageFormat() string {
	return messageFormat
}

type ZError struct {
	severity      Severity
	originalError error
	message       string
	callerFrame   runtime.Frame
	created       time.Time
}

func New(format string, a ...interface{}) *ZError {
	return newZError(defaultSeverity, RuntimeFrameIndirectCallerIndex, format, a...)
}

func NewFatal(format string, a ...interface{}) *ZError {
	return newZError(SeverityFatal, RuntimeFrameIndirectCallerIndex, format, a...)
}

func NewWarn(format string, a ...interface{}) *ZError {
	return newZError(SeverityWarn, RuntimeFrameIndirectCallerIndex, format, a...)
}

func NewInfo(format string, a ...interface{}) *ZError {
	return newZError(SeverityInfo, RuntimeFrameIndirectCallerIndex, format, a...)
}

func (z *ZError) SetSeverity(s Severity) {
	z.severity = s
}

func (z *ZError) GetSeverity() Severity {
	return z.severity
}

func (z *ZError) SetMessage(m string) {
	z.message = m
}

func (z *ZError) GetMessage() string {
	return z.message
}

func (z *ZError) GetFrame() runtime.Frame {
	return z.callerFrame
}

func (z *ZError) GetCreated() time.Time {
	return z.created
}

func (z *ZError) SetError(err error) {
	z.originalError = err
}

func (z *ZError) GetError() error {
	return z.originalError
}

func (z *ZError) Error() string {
	return z.message
}

func (z ZError) String() string {
	msg := strings.ReplaceAll(messageFormat, NotationTime, z.created.Format(timeFormat))
	msg = strings.ReplaceAll(msg, NotationSeverity, fmt.Sprintf("%-5v", z.severity.String()))
	msg = strings.ReplaceAll(msg, NotationMessage, z.message)
	msg = strings.ReplaceAll(msg, NotationFile, z.callerFrame.File)
	msg = strings.ReplaceAll(msg, NotationFunc, z.callerFrame.Function)
	msg = strings.ReplaceAll(msg, NotationLine, fmt.Sprintf("%d", z.callerFrame.Line))
	return msg
}

func newZError(s Severity, index int, format string, a ...interface{}) *ZError {
	msg := format
	if len(a) > 0 {
		msg = fmt.Sprintf(format, a...)
	}
	return &ZError{
		severity:      s,
		originalError: errors.New(msg),
		message:       msg,
		callerFrame:   getFrame(index),
		created:       time.Now(),
	}
}

const RuntimeFrameIndirectCallerIndex = 2

/*
  Getting from https://stackoverflow.com/questions/35212985/is-it-possible-get-information-about-caller-function-in-golang
*/
func getFrame(skipFrames int) runtime.Frame {
	// We need the frame at index skipFrames+2, since we never want runtime.Callers and getFrame
	targetFrameIndex := skipFrames + 2

	// Set size to targetFrameIndex+2 to ensure we have room for one more caller than we need
	programCounters := make([]uintptr, targetFrameIndex+2)
	n := runtime.Callers(0, programCounters)

	frame := runtime.Frame{Function: "unknown"}
	if n > 0 {
		frames := runtime.CallersFrames(programCounters[:n])
		//fmt.Printf("frames: %+v\n", frames)
		for more, frameIndex := true, 0; more && frameIndex <= targetFrameIndex; frameIndex++ {
			var frameCandidate runtime.Frame
			frameCandidate, more = frames.Next()
			//fmt.Printf("candidate: %+v\n", frameCandidate)
			if frameIndex == targetFrameIndex {
				frame = frameCandidate
			}
		}
	}

	return frame
}
