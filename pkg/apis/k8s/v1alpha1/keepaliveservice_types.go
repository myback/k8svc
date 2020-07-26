package v1alpha1

import (
	apiv1 "k8s.io/api/core/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// KeepAliveServiceSpec defines the desired state of KeepAliveService
type KeepAliveServiceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	Hosts          []string           `json:"hosts"`
	Ports          []PortsSpec        `json:"ports"`
	Type           corev1.ServiceType `json:"type,omitempty"`
	Template       Template           `json:"template,omitempty"`
	ReadinessProbe ReadinessProbe     `json:"readinessProbe,omitempty"`
}

type PortsSpec struct {
	apiv1.ServicePort `json:",inline"`
	Name              string `json:"name"`
}

type Template struct {
	Name        string            `json:"name,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
}

type ReadinessProbe struct {
	HTTPGet          HTTPGetChecks   `json:"httpGet,omitempty"`
	TCPSocket        TCPSocketChecks `json:"tcpSocket,omitempty"`
	Script           []string        `json:"script,omitempty"`
	Timeout          int32           `json:"timeout,omitempty"`
	PeriodSeconds    int32           `json:"periodSeconds,omitempty"`
	PeriodExtend     int32           `json:"periodExtend,omitempty"`
	FailureThreshold int32           `json:"failureThreshold,omitempty"`
	SuccessThreshold int32           `json:"successThreshold,omitempty"`
}

type PortsName []string

type HTTPGetChecks struct {
	Path        string            `json:"path,omitempty"`
	HTTPHeaders map[string]string `json:"httpHeaders,omitempty"`
	PortsName   PortsName         `json:"portsName"`
}

type TCPSocketChecks struct {
	PortsName PortsName `json:"portsName"`
}

// KeepAliveServiceStatus defines the observed state of KeepAliveService
type KeepAliveServiceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KeepAliveService is the Schema for the keepaliveservices API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=keepaliveservices,scope=Namespaced
type KeepAliveService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KeepAliveServiceSpec   `json:"spec,omitempty"`
	Status KeepAliveServiceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KeepAliveServiceList contains a list of KeepAliveService
type KeepAliveServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KeepAliveService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KeepAliveService{}, &KeepAliveServiceList{})
}
