package logrus_marker

import "github.com/sirupsen/logrus"

type MessageMarkerHook struct {
	Prefix string
}

func NewMessageMarkerHook(prefix string) *MessageMarkerHook {
	return &MessageMarkerHook{Prefix: prefix}
}

func (self *MessageMarkerHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (self *MessageMarkerHook) Fire(e *logrus.Entry) error {
	e.Message = self.Prefix + " " + e.Message
	return nil
}

func NewLog(prefix string) *logrus.Logger {
	l := logrus.New()
	l.AddHook(NewMessageMarkerHook(prefix))
	return l
}

func NewWithLevel(prefix string, level logrus.Level) *logrus.Logger {
	l := NewLog(prefix)
	l.SetLevel(level)
	return l
}