/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ManagedObjectStorageUserPolicyInitParameters struct {

	// (String) Policy name.
	// Policy name.
	// +crossplane:generate:reference:type=github.com/UpCloudLtd/crossplane-provider-upcloud/apis/objectstorage/v1alpha1.ManagedObjectStoragePolicy
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Reference to a ManagedObjectStoragePolicy in objectstorage to populate name.
	// +kubebuilder:validation:Optional
	NameRef *v1.Reference `json:"nameRef,omitempty" tf:"-"`

	// Selector for a ManagedObjectStoragePolicy in objectstorage to populate name.
	// +kubebuilder:validation:Optional
	NameSelector *v1.Selector `json:"nameSelector,omitempty" tf:"-"`

	// (String) Managed Object Storage service UUID.
	// Managed Object Storage service UUID.
	// +crossplane:generate:reference:type=github.com/UpCloudLtd/crossplane-provider-upcloud/apis/objectstorage/v1alpha1.ManagedObjectStorage
	ServiceUUID *string `json:"serviceUuid,omitempty" tf:"service_uuid,omitempty"`

	// Reference to a ManagedObjectStorage in objectstorage to populate serviceUuid.
	// +kubebuilder:validation:Optional
	ServiceUUIDRef *v1.Reference `json:"serviceUuidRef,omitempty" tf:"-"`

	// Selector for a ManagedObjectStorage in objectstorage to populate serviceUuid.
	// +kubebuilder:validation:Optional
	ServiceUUIDSelector *v1.Selector `json:"serviceUuidSelector,omitempty" tf:"-"`

	// (String) Username.
	// Username.
	// +crossplane:generate:reference:type=github.com/UpCloudLtd/crossplane-provider-upcloud/apis/objectstorage/v1alpha1.ManagedObjectStorageUser
	Username *string `json:"username,omitempty" tf:"username,omitempty"`

	// Reference to a ManagedObjectStorageUser in objectstorage to populate username.
	// +kubebuilder:validation:Optional
	UsernameRef *v1.Reference `json:"usernameRef,omitempty" tf:"-"`

	// Selector for a ManagedObjectStorageUser in objectstorage to populate username.
	// +kubebuilder:validation:Optional
	UsernameSelector *v1.Selector `json:"usernameSelector,omitempty" tf:"-"`
}

type ManagedObjectStorageUserPolicyObservation struct {

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Policy name.
	// Policy name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Managed Object Storage service UUID.
	// Managed Object Storage service UUID.
	ServiceUUID *string `json:"serviceUuid,omitempty" tf:"service_uuid,omitempty"`

	// (String) Username.
	// Username.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type ManagedObjectStorageUserPolicyParameters struct {

	// (String) Policy name.
	// Policy name.
	// +crossplane:generate:reference:type=github.com/UpCloudLtd/crossplane-provider-upcloud/apis/objectstorage/v1alpha1.ManagedObjectStoragePolicy
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Reference to a ManagedObjectStoragePolicy in objectstorage to populate name.
	// +kubebuilder:validation:Optional
	NameRef *v1.Reference `json:"nameRef,omitempty" tf:"-"`

	// Selector for a ManagedObjectStoragePolicy in objectstorage to populate name.
	// +kubebuilder:validation:Optional
	NameSelector *v1.Selector `json:"nameSelector,omitempty" tf:"-"`

	// (String) Managed Object Storage service UUID.
	// Managed Object Storage service UUID.
	// +crossplane:generate:reference:type=github.com/UpCloudLtd/crossplane-provider-upcloud/apis/objectstorage/v1alpha1.ManagedObjectStorage
	// +kubebuilder:validation:Optional
	ServiceUUID *string `json:"serviceUuid" tf:"service_uuid,omitempty"`

	// Reference to a ManagedObjectStorage in objectstorage to populate serviceUuid.
	// +kubebuilder:validation:Optional
	ServiceUUIDRef *v1.Reference `json:"serviceUuidRef,omitempty" tf:"-"`

	// Selector for a ManagedObjectStorage in objectstorage to populate serviceUuid.
	// +kubebuilder:validation:Optional
	ServiceUUIDSelector *v1.Selector `json:"serviceUuidSelector,omitempty" tf:"-"`

	// (String) Username.
	// Username.
	// +crossplane:generate:reference:type=github.com/UpCloudLtd/crossplane-provider-upcloud/apis/objectstorage/v1alpha1.ManagedObjectStorageUser
	// +kubebuilder:validation:Optional
	Username *string `json:"username" tf:"username,omitempty"`

	// Reference to a ManagedObjectStorageUser in objectstorage to populate username.
	// +kubebuilder:validation:Optional
	UsernameRef *v1.Reference `json:"usernameRef,omitempty" tf:"-"`

	// Selector for a ManagedObjectStorageUser in objectstorage to populate username.
	// +kubebuilder:validation:Optional
	UsernameSelector *v1.Selector `json:"usernameSelector,omitempty" tf:"-"`
}

// ManagedObjectStorageUserPolicySpec defines the desired state of ManagedObjectStorageUserPolicy
type ManagedObjectStorageUserPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagedObjectStorageUserPolicyParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ManagedObjectStorageUserPolicyInitParameters `json:"initProvider,omitempty"`
}

// ManagedObjectStorageUserPolicyStatus defines the observed state of ManagedObjectStorageUserPolicy.
type ManagedObjectStorageUserPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagedObjectStorageUserPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ManagedObjectStorageUserPolicy is the Schema for the ManagedObjectStorageUserPolicys API. This resource represents an UpCloud Managed Object Storage user policy attachment.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,upcloud}
type ManagedObjectStorageUserPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagedObjectStorageUserPolicySpec   `json:"spec"`
	Status            ManagedObjectStorageUserPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedObjectStorageUserPolicyList contains a list of ManagedObjectStorageUserPolicys
type ManagedObjectStorageUserPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedObjectStorageUserPolicy `json:"items"`
}

// Repository type metadata.
var (
	ManagedObjectStorageUserPolicy_Kind             = "ManagedObjectStorageUserPolicy"
	ManagedObjectStorageUserPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ManagedObjectStorageUserPolicy_Kind}.String()
	ManagedObjectStorageUserPolicy_KindAPIVersion   = ManagedObjectStorageUserPolicy_Kind + "." + CRDGroupVersion.String()
	ManagedObjectStorageUserPolicy_GroupVersionKind = CRDGroupVersion.WithKind(ManagedObjectStorageUserPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&ManagedObjectStorageUserPolicy{}, &ManagedObjectStorageUserPolicyList{})
}
