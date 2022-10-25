package infrastructures

type IKernel interface {
	loadEnvVars()
}

type kernel struct{}

var fname = ".env"

func (k *kernel) loadEnvVars() {
	if err := InitEnvLoader().LoadFromFile(fname); err != nil {
		println(err)
	}
}
