package keepaliveservice

import (
	k8sv1alpha1 "github.com/myback/k8svc/pkg/apis/k8s/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"reflect"
)

func keepAliveServiceEndpointSubsets(spec k8sv1alpha1.KeepAliveServiceSpec) []corev1.EndpointSubset {
	endpointAddrs := []corev1.EndpointAddress{}
	endpointPorts := []corev1.EndpointPort{}

	for _, host := range spec.Hosts {
		endpointAddrs = append(endpointAddrs, corev1.EndpointAddress{IP: host})
	}

	for _, port := range spec.Ports {
		ep := corev1.EndpointPort{
			Name:     port.Name,
			Port:     port.Port,
			Protocol: corev1.ProtocolTCP,
		}

		if port.Protocol != "" {
			ep.Protocol = port.Protocol
		}

		endpointPorts = append(endpointPorts, ep)
	}

	return []corev1.EndpointSubset{
		{
			Addresses: endpointAddrs,
			//NotReadyAddresses: endpointAddrs,
			Ports: endpointPorts,
		},
	}
}

func endpointsEqual(crd, k8s []corev1.EndpointSubset) bool {
	secondSubs := []corev1.EndpointSubset{}

	for _, subs := range k8s {
		endpointAddrs := []corev1.EndpointAddress{}
		endpointPorts := []corev1.EndpointPort{}

		for _, addr := range subs.Addresses {
			endpointAddrs = append(endpointAddrs, corev1.EndpointAddress{IP: addr.IP})
		}

		for _, port := range subs.Ports {
			endpointPorts = append(endpointPorts, corev1.EndpointPort{
				Name:     port.Name,
				Port:     port.Port,
				Protocol: port.Protocol,
			})
		}

		secondSubs = append(secondSubs, corev1.EndpointSubset{
			Addresses: endpointAddrs,
			//NotReadyAddresses: endpointAddrs,
			Ports: endpointPorts,
		})
	}

	return reflect.DeepEqual(crd, secondSubs)
}

func newEndpoint(name string, cr *k8sv1alpha1.KeepAliveService) *corev1.Endpoints {
	return &corev1.Endpoints{
		ObjectMeta: v1.ObjectMeta{
			Name:        name,
			Namespace:   cr.Namespace,
			Labels:      cr.Spec.Template.Labels,
			Annotations: cr.Spec.Template.Annotations,
		},
		Subsets: keepAliveServiceEndpointSubsets(cr.Spec),
	}
}

func keepAliveServiceServiceSpec(cr k8sv1alpha1.KeepAliveServiceSpec) corev1.ServiceSpec {
	portsSpec := []corev1.ServicePort{}
	emptyTargetPort := intstr.IntOrString{}

	for _, portSpec := range cr.Ports {
		spec := *portSpec.DeepCopyServicePort()
		if spec.TargetPort == emptyTargetPort {
			spec.TargetPort = intstr.IntOrString{IntVal: portSpec.Port}
		}

		portsSpec = append(portsSpec, spec)
	}

	svcType := corev1.ServiceTypeClusterIP
	if cr.Type != "" {
		svcType = cr.Type
	}

	return corev1.ServiceSpec{
		Ports:       portsSpec,
		Type:        svcType,
		Selector:    nil,
		ExternalIPs: nil,
	}
}

func newService(name string, cr *k8sv1alpha1.KeepAliveService) *corev1.Service {
	return &corev1.Service{
		ObjectMeta: v1.ObjectMeta{
			Name:        name,
			Namespace:   cr.Namespace,
			Labels:      cr.Spec.Template.Labels,
			Annotations: cr.Spec.Template.Annotations,
		},
		Spec: keepAliveServiceServiceSpec(cr.Spec),
	}
}

func serviceEqual(crd, k8s corev1.ServiceSpec) bool {
	svcSpec := corev1.ServiceSpec{
		Type: k8s.Type,
	}

	for _, port := range k8s.Ports {
		svcSpec.Ports = append(svcSpec.Ports, corev1.ServicePort{
			Name:       port.Name,
			NodePort:   port.NodePort,
			Protocol:   port.Protocol,
			Port:       port.Port,
			TargetPort: port.TargetPort,
		})
	}

	return reflect.DeepEqual(crd, svcSpec)
}
