package tendermintlogger

import tmlog "github.com/tendermint/tendermint/libs/log"

type DiscardLogger struct{}

func (l DiscardLogger) Debug(_ string, _ ...interface{})   {}
func (l DiscardLogger) Info(_ string, _ ...interface{})    {}
func (l DiscardLogger) Error(_ string, _ ...interface{})   {}
func (l DiscardLogger) With(_ ...interface{}) tmlog.Logger { return l }
