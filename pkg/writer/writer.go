package writer

type (
	Config struct {
		reference string
		args      string
		// packages  []packages.Package
	}

	Writer struct {
		conf *Config
	}
)

func New(config *Config) *Writer {
	return &Writer{
		conf: config,
	}
}

func (w *Writer) Write(tag string, sources *[]source) error {

	return nil
}
