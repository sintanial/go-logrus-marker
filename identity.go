package logrus_msgid

import "github.com/sirupsen/logrus"

type MsgIdentityHook struct {
	Prefix string
}

func NewMsgIdentityHook(prefix string) *MsgIdentityHook {
	return &MsgIdentityHook{Prefix: prefix}
}

func (self *MsgIdentityHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (self *MsgIdentityHook) Fire(e *logrus.Entry) error {
	e.Message = self.Prefix + " " + e.Message
	return nil
}

func NewLog(prefix string) *logrus.Logger {
	l := logrus.New()
	l.AddHook(NewMsgIdentityHook(prefix))
	return l
}

