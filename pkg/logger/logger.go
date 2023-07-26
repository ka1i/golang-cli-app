package logger

func Init() {
	SystemLog()

	// middleware log setting
	GinLog()
}
