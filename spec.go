package main

type DeploymentSpec struct {
	Name string `yaml:"name"`

	VPC VPCSpec `yaml:"vpc"`

	InternetGateways []InternetGatewaySpec `yaml:"internet_gateways"`

	Subnets []SubnetSpec `yaml:"subnets"`

	DNS []DNSSpec `yaml:"dns"`

	SecurityGroups []SecurityGroupSpec `yaml:"security_groups"`

	LoadBalancers []LoadBalancerSpec `yaml:"load_balancers"`
}

type VPCSpec struct {
	CIDR            string `yaml:"cidr"`
	InternetGateway string `yaml:"internet_gateway"`
}

type InternetGatewaySpec struct {
	Name string `yaml:"name"`
}

type SubnetSpec struct {
	Name             string          `yaml:"name"`
	CIDR             string          `yaml:"cidr"`
	AvailabilityZone string          `yaml:"availability_zone"`
	RouteTable       *RouteTableSpec `yaml:"route_table,omitempty"`
	NAT              *SubnetNATSpec  `yaml:"nat,omitempty"`
}

type RouteTableSpec struct {
	InternetGateway *string `yaml:"internet_gateway,omitempty"`
	Instance        *string `yaml:"instance,omitempty"`
}

type DNSSpec string

type SubnetNATSpec struct {
	Name          string `yaml:"name"`
	InstanceType  string `yaml:"type"`
	IP            string `yaml:"ip"`
	KeyPairName   string `yaml:"key_pair_name"`
	SecurityGroup string `yaml:"security_group"`
}

type SecurityGroupSpec struct {
	Name    string                   `yaml:"name"`
	Ingress []SecurityGroupEntrySpec `yaml:"ingress,omitempty"`
	Egress  []SecurityGroupEntrySpec `yaml:"egress,omitempty"`
}

type SecurityGroupEntrySpec struct {
	CIDR     string `yaml:"cidr"`
	Protocol string `yaml:"protocol,omitempty"`
	Ports    string `yaml:"ports,omitempty"`
}

type LoadBalancerSpec struct {
	Name           string                      `yaml:"name"`
	Listeners      []LoadBalancerListenerSpec  `yaml:"listeners"`
	HealthCheck    LoadBalancerHealthCheckSpec `yaml:"health_check"`
	Subnets        []string                    `yaml:"subnets"`
	SecurityGroups []string                    `yaml:"security_groups"`
}

type LoadBalancerListenerSpec struct {
	Protocol            string  `yaml:"protocol"`
	Port                string  `yaml:"port"`
	DestinationProtocol *string `yaml:"destination_protocol,omitempty"`
	DestinationPort     *string `yaml:"destination_port,omitempty"`
}

type LoadBalancerHealthCheckSpec struct {
	Target             LoadBalancerHealthCheckTargetSpec `yaml:"target"`
	Timeout            string                            `yaml:"timeout"`
	Interval           string                            `yaml:"interval"`
	HealthyThreshold   string                            `yaml:"healthy_threshold"`
	UnhealthyThreshold string                            `yaml:"unhealthy_threshold"`
}

type LoadBalancerHealthCheckTargetSpec struct {
	Type string `yaml:"type"`
	Port string `yaml:"port"`
}
