package config

type Logger struct {
	LowestLevel string
	RecordLineNumber bool
	LogFormatter string
	RuntimeLogFile string
	ErrorLogFile string
	MaxAge int
	Rotate struct{
		Enable bool
		Type string
		Size struct{
			MaxSize int
			MaxBackups int
			Compress bool
		}
		Date struct{
			Extend string
		}
	}
	Access struct{
		Enable bool
		LogFile string
	}
	Rpc struct{
		Enable bool
		RpcLogFile string
	}
}