package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SecretWithExpirySpec defines the desired state of SecretWithExpiry
type SecretWithExpirySpec struct {
	SecretName string     `json:"secretName"`
	ExpiryDate metav1.Time `json:"expiryDate"`
}

// SecretWithExpiryStatus defines the observed state of SecretWithExpiry
type SecretWithExpiryStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// SecretWithExpiry is the Schema for the secretwithexpiries API
type SecretWithExpiry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SecretWithExpirySpec   `json:"spec,omitempty"`
	Status SecretWithExpiryStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SecretWithExpiryList contains a list of SecretWithExpiry
type SecretWithExpiryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretWithExpiry `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SecretWithExpiry{}, &SecretWithExpiryList{})
}
