// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPGetChecks) DeepCopyInto(out *HTTPGetChecks) {
	*out = *in
	if in.HTTPHeaders != nil {
		in, out := &in.HTTPHeaders, &out.HTTPHeaders
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PortsName != nil {
		in, out := &in.PortsName, &out.PortsName
		*out = make(PortsName, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPGetChecks.
func (in *HTTPGetChecks) DeepCopy() *HTTPGetChecks {
	if in == nil {
		return nil
	}
	out := new(HTTPGetChecks)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeepAliveService) DeepCopyInto(out *KeepAliveService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeepAliveService.
func (in *KeepAliveService) DeepCopy() *KeepAliveService {
	if in == nil {
		return nil
	}
	out := new(KeepAliveService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeepAliveService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeepAliveServiceList) DeepCopyInto(out *KeepAliveServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KeepAliveService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeepAliveServiceList.
func (in *KeepAliveServiceList) DeepCopy() *KeepAliveServiceList {
	if in == nil {
		return nil
	}
	out := new(KeepAliveServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeepAliveServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeepAliveServiceSpec) DeepCopyInto(out *KeepAliveServiceSpec) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]PortsSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Template.DeepCopyInto(&out.Template)
	in.ReadinessProbe.DeepCopyInto(&out.ReadinessProbe)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeepAliveServiceSpec.
func (in *KeepAliveServiceSpec) DeepCopy() *KeepAliveServiceSpec {
	if in == nil {
		return nil
	}
	out := new(KeepAliveServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeepAliveServiceStatus) DeepCopyInto(out *KeepAliveServiceStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeepAliveServiceStatus.
func (in *KeepAliveServiceStatus) DeepCopy() *KeepAliveServiceStatus {
	if in == nil {
		return nil
	}
	out := new(KeepAliveServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in PortsName) DeepCopyInto(out *PortsName) {
	{
		in := &in
		*out = make(PortsName, len(*in))
		copy(*out, *in)
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortsName.
func (in PortsName) DeepCopy() PortsName {
	if in == nil {
		return nil
	}
	out := new(PortsName)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortsSpec) DeepCopyInto(out *PortsSpec) {
	*out = *in
	in.ServicePort.DeepCopyInto(&out.ServicePort)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortsSpec.
func (in *PortsSpec) DeepCopy() *PortsSpec {
	if in == nil {
		return nil
	}
	out := new(PortsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReadinessProbe) DeepCopyInto(out *ReadinessProbe) {
	*out = *in
	in.HTTPGet.DeepCopyInto(&out.HTTPGet)
	in.TCPSocket.DeepCopyInto(&out.TCPSocket)
	if in.Script != nil {
		in, out := &in.Script, &out.Script
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReadinessProbe.
func (in *ReadinessProbe) DeepCopy() *ReadinessProbe {
	if in == nil {
		return nil
	}
	out := new(ReadinessProbe)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TCPSocketChecks) DeepCopyInto(out *TCPSocketChecks) {
	*out = *in
	if in.PortsName != nil {
		in, out := &in.PortsName, &out.PortsName
		*out = make(PortsName, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TCPSocketChecks.
func (in *TCPSocketChecks) DeepCopy() *TCPSocketChecks {
	if in == nil {
		return nil
	}
	out := new(TCPSocketChecks)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Template) DeepCopyInto(out *Template) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Template.
func (in *Template) DeepCopy() *Template {
	if in == nil {
		return nil
	}
	out := new(Template)
	in.DeepCopyInto(out)
	return out
}
