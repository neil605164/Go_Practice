package basic

import "log"

type ILog interface {
	Info()
	Debug()
}

type logger struct {
	IBasicLog
}

func ProviderBasicLog(opt ...func(*logger)) ILog {
	a := &logger{}
	for _, o := range opt {
		o(a)
	}

	if a.IBasicLog == nil {
		a.IBasicLog = &basicLog{}
	}

	return a
}

func (b *logger) Info() {
	b.IBasicLog.Info()
}

func (b *logger) Debug() {
	b.IBasicLog.Debug()
}

func WithCustomerOption(l IBasicLog) func(*logger) {
	return func(log *logger) {
		log.IBasicLog = l
	}
}

//////////

type IBasicLog interface {
	Info()
	Debug()
}

type basicLog struct {
}

func (b *basicLog) Info() {
	log.Println("basic info log")
}

func (b *basicLog) Debug() {
	log.Println("basic debug log")
}

//////////////////

type A struct {
}

func (b *A) Info() {
	log.Println("a info log")
}

func (b *A) Debug() {
	log.Println("a debug log")
}

type B struct {
}

func (b *B) Info() {
	log.Println("b info log")
}

func (b *B) Debug() {
	log.Println("b debug log")
}
