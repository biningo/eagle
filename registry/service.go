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
}

type IP struct {
	PublicIP  string `json:"public_ip"`
	PrivateIP string `json:"private_ip"`
}
type Port struct {
	PublicPort  uint16 `json:"public_port"`
	PrivatePort uint16 `json:"private_port"`
}
type ServiceInstance struct {
	Service
	IP   IP
	Port Port
	ID   string `json:"id"`
}

func NewServiceInstance(name string, id string, privateIP string, privatePort uint16, opts ...ServiceOption) *ServiceInstance {
	svc := &ServiceInstance{}
	svc.ID = id
	svc.Service.Name = name
	svc.Service.Namespace = "default"
	svc.IP = IP{
		PublicIP:  privateIP,
		PrivateIP: privateIP,
	}
	svc.Port = Port{
		PublicPort:  privatePort,
		PrivatePort: privatePort,
	}

	for _, o := range opts {
		o(svc)
	}
	return svc
}

type ServiceOption func(*ServiceInstance)

func WithNamespace(namespace string) ServiceOption {
	return func(service *ServiceInstance) {
		service.Service.Namespace = namespace
	}
}

func WithPrivateIP(privateIP string) ServiceOption {
	return func(service *ServiceInstance) {
		service.IP.PrivateIP = privateIP
	}
}

func WithPrivatePort(privatePort uint16) ServiceOption {
	return func(service *ServiceInstance) {
		service.Port.PrivatePort = privatePort
	}
}
