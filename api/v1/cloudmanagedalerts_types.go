package v1

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type CloudManagedAlertSpec struct {
	// Alert name
	// +kubebuilder:validation:Pattern=^.*$
	AlertName string `json:"alertname"`
	// Value
	// +kubebuilder:validation:Pattern=^.*$
	AlertValue string `json:"alertvalue"`
	// Cluster name
	// +kubebuilder:validation:Pattern=^.*$
	Cluster string `json:"cluster"`
	// Pod
	// +kubebuilder:validation:Pattern=^.*$
	Pod string `json:"pod"`
}

// CloudManagedAlert defines the observed state of CloudManagedAlert
type CloudManagedAlertStatus struct {
	Status string `json:"status"`
}

// +kubebuilder:object:root=true
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.status",description="The cluster status"
// +kubebuilder:printcolumn:name="AlertName",type="string",JSONPath=".spec.alertname",description="Alert name"
// +kubebuilder:printcolumn:name="AlertValue",type="string",JSONPath=".spec.alertvalue",description="Alert value"
// +kubebuilder:printcolumn:name="Cluster",type="string",JSONPath=".spec.cluster",description="Cluster name"
// +kubebuilder:printcolumn:name="Pod",type="string",JSONPath=".spec.pod",description="Affected Pod Name"
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

func (cm *CloudManagedAlert) Process() func() string {
	switch cm.Status.Status {
	case "":
		return cm.newProcess
	case AlertCreatedStatus:
		return cm.createdStatusProcess
	case AlertAckedStatus:
		return cm.ackedStatusProcess
	default:
		return func() string {
			fmt.Printf("Unknown status!")
			return "unknown"
		}
	}
}

func (cm *CloudManagedAlert) newProcess() string {
	fmt.Printf("Doing new process!")
	return AlertCreatedStatus
}

func (cm *CloudManagedAlert) createdStatusProcess() string {
	fmt.Printf("Doing created process")
	return AlertAckedStatus
}

func (cm *CloudManagedAlert) ackedStatusProcess() string {
	fmt.Printf("Doing acked process")
	return AlertResolvedStatus
}

func init() {
	SchemeBuilder.Register(&CloudManagedAlert{}, &CloudManagedAlertList{})
}
