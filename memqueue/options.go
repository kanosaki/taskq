package memqueue

import "gopkg.in/queue.v1/processor"

type Storager interface {
	Exists(key string) bool
}

type Options struct {
	Name    string
	Storage Storager

	Processor  processor.Options
	AlwaysSync bool // if true messages are always processed synchronously
}

func (opt *Options) init() {}
