package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type CloudManagedAlertSpec struct {
	// Alert name
	// +kubebuilder:validation:Pattern=^.*$
	Name string `json:"name"`
	// Cluster name
	// +kubebuilder:validation:Pattern=^.*$
	Cluster string `json:"cluster"`
	// AlertEndpoint
	// +kubebuilder:validation:Pattern=^.*$
	AlertEndpoint string `json:"endpoint"`
}

// CloudManagedAlert defines the observed state of CloudManagedAlert
type CloudManagedAlertStatus struct {
	Status string `json:"status"`
}

// +kubebuilder:object:root=true
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.status",description="The cluster status"
// +kubebuilder:printcolumn:name="Name",type=string,JSONPath=`.spec.name`,description="Alert name"
// +kubebuilder:printcolumn:name="Cluster",type=string,JSONPath=`.spec.cluster`,description="Cluster name"
// +kubebuilder:printcolumn:name="AlertEndpoint",type=string,JSONPath=`.spec.endpoint`,description="Alert endpoint"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:shortName=cla
type CloudManagedAlert struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudManagedAlertSpec   `json:"spec,omitempty"`
	Status CloudManagedAlertStatus `json:"status,omitempty"`
}

// CloudManagedAlertList contains a list of CloudManagedAlert
// +kubebuilder:object:root=true
type CloudManagedAlertList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudManagedAlert `json:"items"`
}

func (cm *CloudManagedAlert) IsEqual(newStatus string) bool {
	return cm.Status.Status == newStatus
}

func (cm *CloudManagedAlert) SetStatus(newStatus string) {
	cm.Status.Status = newStatus
}

func (cm *CloudManagedAlert) GetStatus() string {
	return cm.Status.Status
}

func init() {
	SchemeBuilder.Register(&CloudManagedAlert{}, &CloudManagedAlertList{})
}
