package mysql

type config struct {
	isMock bool
}

type Option func(*config)

func OptionIsMock() Option {
	return func(conf *config) {
		conf.isMock = true
	}
}
