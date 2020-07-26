package v1alpha1

import corev1 "k8s.io/api/core/v1"

func (in *PortsSpec) DeepCopyServicePort() *corev1.ServicePort {
	if in == nil {
		return nil
	}
	out := new(corev1.ServicePort)
	in.ServicePort.DeepCopyInto(out)
	out.Name = in.Name
	return out
}
