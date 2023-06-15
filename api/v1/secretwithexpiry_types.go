package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SecretWithExpiry is a specification for a SecretWithExpiry resource
type SecretWithExpiry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec SecretWithExpirySpec `json:"spec"`
}

// SecretWithExpirySpec is the spec for a SecretWithExpiry resource
type SecretWithExpirySpec struct {
	SecretName string     `json:"secretName"`
	ExpiryDate metav1.Time `json:"expiryDate"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SecretWithExpiryList is a list of SecretWithExpiry resources
type SecretWithExpiryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []SecretWithExpiry `json:"items"`
}