package logs

import (
	"encoding/json"
	"sync"
)

type Logs struct {
	lines []string
	mu    sync.RWMutex
}

var (
	logInst *Logs
	logOnce sync.Once
)

func getLogInstance() *Logs {
	logOnce.Do(func() {
		logInst = newLogs()
	})
	return logInst
}

func newLogs() *Logs {
	return &Logs{lines: make([]string, 0)}
}

func (l *Logs) add(text string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.lines = append(l.lines, text)
}

func (l *Logs) remove(text string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	for i, v := range l.lines {
		if v == text {
			l.lines = append(l.lines[:i], l.lines[i+1:]...)
			break
		}
	}
}

func (l *Logs) contains(text string) bool {
	l.mu.RLock()
	defer l.mu.RUnlock()
	for _, v := range l.lines {
		if v == text {
			return true
		}
	}
	return false
}

func (l *Logs) length() int {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return len(l.lines)
}

func (l *Logs) toJson() string {
	jsonData, err := json.Marshal(l.lines)
	if err != nil {
		return ""
	}
	return string(jsonData)
}

// Public =================================================

func AddText(text string) {
	getLogInstance().add(text)
}

func AddTextAndJump(text string) {
	getLogInstance().add(text)
	getLogInstance().add("")
}

func ToJson() string {
	return getLogInstance().toJson()
}

func Length() int {
	return getLogInstance().length()
}

func Remove(text string) {
	getLogInstance().remove(text)
}

func Contains(text string) bool {
	return getLogInstance().contains(text)
}
