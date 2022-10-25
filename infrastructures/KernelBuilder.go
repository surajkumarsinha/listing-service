package infrastructures

import "sync"

type IKernelBuilder interface {
	Build(config map[string]any) IKernel
}

type kernelBuilder struct{}

var (
	k     *kernel
	kOnce sync.Once
)

func (kb *kernelBuilder) Build(config map[string]any) IKernel {
	if k == nil {
		kOnce.Do(func() {
			k = &kernel{}
			k.loadEnvVars()
		})
	}

	return k
}

func KernelBuilder() IKernelBuilder {
	return &kernelBuilder{}
}
