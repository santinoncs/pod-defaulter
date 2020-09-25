/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// NodePolicyProfileSpec defines the desired state of NodePolicyProfile
type NodePolicyProfileSpec struct {
	Tolerations  []corev1.Toleration `json:"tolerations,omitempty"`
	NodeAffinity corev1.NodeAffinity `json:"nodeAffinity,omitempty"`
	NodeSelector map[string]string   `json:"nodeSelector,omitempty"`
}

// NodePolicyProfileStatus defines the observed state of NodePolicyProfile
type NodePolicyProfileStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster

// NodePolicyProfile is the Schema for the nodepolicyprofiles API
type NodePolicyProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NodePolicyProfileSpec   `json:"spec,omitempty"`
	Status NodePolicyProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NodePolicyProfileList contains a list of NodePolicyProfile
type NodePolicyProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodePolicyProfile `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NodePolicyProfile{}, &NodePolicyProfileList{})
}
