package registry

/**
*@Author lyer
*@Date 7/19/21 13:40
*@Describe
**/

type Service struct {
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	Labels    map[string]string `json:"labels"`
	Check     *ServiceCheck
}

type ServiceInstance struct {
	Service
	ID   string `json:"id"`
	IP   string `json:"ip"`
	Port int16  `json:"port"`
}

type ServiceOption func(*ServiceInstance)

func Namespace(namespace string) ServiceOption {
	return func(service *ServiceInstance) {
		service.Namespace = namespace
	}
}

func Name(name string) ServiceOption {
	return func(service *ServiceInstance) {
		service.Name = name
	}
}

func Check(check *ServiceCheck) ServiceOption {
	return func(service *ServiceInstance) {
		service.Check = check
	}
}
