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

type ManagedObjectStorageUserInitParameters struct {
}

type ManagedObjectStorageUserObservation struct {

	// (String) User ARN.
	// User ARN.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// (String) Creation time.
	// Creation time.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Managed Object Storage service UUID.
	// Managed Object Storage service UUID.
	ServiceUUID *string `json:"serviceUuid,omitempty" tf:"service_uuid,omitempty"`
}

type ManagedObjectStorageUserParameters struct {

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
}

// ManagedObjectStorageUserSpec defines the desired state of ManagedObjectStorageUser
type ManagedObjectStorageUserSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagedObjectStorageUserParameters `json:"forProvider"`
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
	InitProvider ManagedObjectStorageUserInitParameters `json:"initProvider,omitempty"`
}

// ManagedObjectStorageUserStatus defines the observed state of ManagedObjectStorageUser.
type ManagedObjectStorageUserStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagedObjectStorageUserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ManagedObjectStorageUser is the Schema for the ManagedObjectStorageUsers API. This resource represents an UpCloud Managed Object Storage user. No relation to UpCloud API accounts.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,upcloud}
type ManagedObjectStorageUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagedObjectStorageUserSpec   `json:"spec"`
	Status            ManagedObjectStorageUserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedObjectStorageUserList contains a list of ManagedObjectStorageUsers
type ManagedObjectStorageUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedObjectStorageUser `json:"items"`
}

// Repository type metadata.
var (
	ManagedObjectStorageUser_Kind             = "ManagedObjectStorageUser"
	ManagedObjectStorageUser_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ManagedObjectStorageUser_Kind}.String()
	ManagedObjectStorageUser_KindAPIVersion   = ManagedObjectStorageUser_Kind + "." + CRDGroupVersion.String()
	ManagedObjectStorageUser_GroupVersionKind = CRDGroupVersion.WithKind(ManagedObjectStorageUser_Kind)
)

func init() {
	SchemeBuilder.Register(&ManagedObjectStorageUser{}, &ManagedObjectStorageUserList{})
}
