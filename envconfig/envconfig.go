package envconfig

const (
	FILE_NAMES = ""
)

func init() {
	println("loading environment variables")
}

func LoadEnvFromFiles(filenames ...string) {
	if len(filenames) == 0 {
		for filename := range filenames {
			println(filename)
		}
	}
}
