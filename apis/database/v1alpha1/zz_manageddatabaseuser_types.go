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

type ManagedDatabaseUserInitParameters struct {

	// (String) MySQL only, authentication type.
	// MySQL only, authentication type.
	Authentication *string `json:"authentication,omitempty" tf:"authentication,omitempty"`

	// (Block List, Max: 1) OpenSearch access control object. (see below for nested schema)
	// OpenSearch access control object.
	OpensearchAccessControl []OpensearchAccessControlInitParameters `json:"opensearchAccessControl,omitempty" tf:"opensearch_access_control,omitempty"`

	// (String, Sensitive) Password for the database user. Defaults to a random value
	// Password for the database user. Defaults to a random value
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// (Block List, Max: 1) PostgreSQL access control object. (see below for nested schema)
	// PostgreSQL access control object.
	PgAccessControl []PgAccessControlInitParameters `json:"pgAccessControl,omitempty" tf:"pg_access_control,omitempty"`

	// (Block List, Max: 1) Redis access control object. (see below for nested schema)
	// Redis access control object.
	RedisAccessControl []RedisAccessControlInitParameters `json:"redisAccessControl,omitempty" tf:"redis_access_control,omitempty"`
}

type ManagedDatabaseUserObservation struct {

	// (String) MySQL only, authentication type.
	// MySQL only, authentication type.
	Authentication *string `json:"authentication,omitempty" tf:"authentication,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Block List, Max: 1) OpenSearch access control object. (see below for nested schema)
	// OpenSearch access control object.
	OpensearchAccessControl []OpensearchAccessControlObservation `json:"opensearchAccessControl,omitempty" tf:"opensearch_access_control,omitempty"`

	// (Block List, Max: 1) PostgreSQL access control object. (see below for nested schema)
	// PostgreSQL access control object.
	PgAccessControl []PgAccessControlObservation `json:"pgAccessControl,omitempty" tf:"pg_access_control,omitempty"`

	// (Block List, Max: 1) Redis access control object. (see below for nested schema)
	// Redis access control object.
	RedisAccessControl []RedisAccessControlObservation `json:"redisAccessControl,omitempty" tf:"redis_access_control,omitempty"`

	// (String) Service's UUID for which this user belongs to
	// The service to which the logical database belongs. Please note that reference fields (`serviceRef` and `serviceSelector`) only work for PostgreSQL databases. For other databases you need to leverage compositions and patches to pass database service ID to database user `service` field. See https://docs.crossplane.io/latest/concepts/patch-and-transform/#patching-between-resources for more info.
	Service *string `json:"service,omitempty" tf:"service,omitempty"`

	// (String) Type of the user. Only normal type users can be created
	// Type of the user. Only normal type users can be created
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ManagedDatabaseUserParameters struct {

	// (String) MySQL only, authentication type.
	// MySQL only, authentication type.
	// +kubebuilder:validation:Optional
	Authentication *string `json:"authentication,omitempty" tf:"authentication,omitempty"`

	// (Block List, Max: 1) OpenSearch access control object. (see below for nested schema)
	// OpenSearch access control object.
	// +kubebuilder:validation:Optional
	OpensearchAccessControl []OpensearchAccessControlParameters `json:"opensearchAccessControl,omitempty" tf:"opensearch_access_control,omitempty"`

	// (String, Sensitive) Password for the database user. Defaults to a random value
	// Password for the database user. Defaults to a random value
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// (Block List, Max: 1) PostgreSQL access control object. (see below for nested schema)
	// PostgreSQL access control object.
	// +kubebuilder:validation:Optional
	PgAccessControl []PgAccessControlParameters `json:"pgAccessControl,omitempty" tf:"pg_access_control,omitempty"`

	// (Block List, Max: 1) Redis access control object. (see below for nested schema)
	// Redis access control object.
	// +kubebuilder:validation:Optional
	RedisAccessControl []RedisAccessControlParameters `json:"redisAccessControl,omitempty" tf:"redis_access_control,omitempty"`

	// (String) Service's UUID for which this user belongs to
	// The service to which the logical database belongs. Please note that reference fields (`serviceRef` and `serviceSelector`) only work for PostgreSQL databases. For other databases you need to leverage compositions and patches to pass database service ID to database user `service` field. See https://docs.crossplane.io/latest/concepts/patch-and-transform/#patching-between-resources for more info.
	// +crossplane:generate:reference:type=github.com/UpCloudLtd/crossplane-provider-upcloud/apis/database/v1alpha1.ManagedDatabasePostgresql
	// +kubebuilder:validation:Optional
	Service *string `json:"service" tf:"service,omitempty"`

	// Reference to a ManagedDatabasePostgresql in database to populate service.
	// +kubebuilder:validation:Optional
	ServiceRef *v1.Reference `json:"serviceRef,omitempty" tf:"-"`

	// Selector for a ManagedDatabasePostgresql in database to populate service.
	// +kubebuilder:validation:Optional
	ServiceSelector *v1.Selector `json:"serviceSelector,omitempty" tf:"-"`
}

type OpensearchAccessControlInitParameters struct {

	// (Block List, Min: 1) Set user access control rules. (see below for nested schema)
	// Set user access control rules.
	Rules []RulesInitParameters `json:"rules,omitempty" tf:"rules,omitempty"`
}

type OpensearchAccessControlObservation struct {

	// (Block List, Min: 1) Set user access control rules. (see below for nested schema)
	// Set user access control rules.
	Rules []RulesObservation `json:"rules,omitempty" tf:"rules,omitempty"`
}

type OpensearchAccessControlParameters struct {

	// (Block List, Min: 1) Set user access control rules. (see below for nested schema)
	// Set user access control rules.
	// +kubebuilder:validation:Optional
	Rules []RulesParameters `json:"rules" tf:"rules,omitempty"`
}

type PgAccessControlInitParameters struct {

	// (Boolean) Grant replication privilege
	// Grant replication privilege
	AllowReplication *bool `json:"allowReplication,omitempty" tf:"allow_replication,omitempty"`
}

type PgAccessControlObservation struct {

	// (Boolean) Grant replication privilege
	// Grant replication privilege
	AllowReplication *bool `json:"allowReplication,omitempty" tf:"allow_replication,omitempty"`
}

type PgAccessControlParameters struct {

	// (Boolean) Grant replication privilege
	// Grant replication privilege
	// +kubebuilder:validation:Optional
	AllowReplication *bool `json:"allowReplication,omitempty" tf:"allow_replication,omitempty"`
}

type RedisAccessControlInitParameters struct {

	// (List of String) Set access control to all commands in specified categories.
	// Set access control to all commands in specified categories.
	Categories []*string `json:"categories,omitempty" tf:"categories,omitempty"`

	// (List of String) Set access control to Pub/Sub channels.
	// Set access control to Pub/Sub channels.
	Channels []*string `json:"channels,omitempty" tf:"channels,omitempty"`

	// (List of String) Set access control to commands.
	// Set access control to commands.
	Commands []*string `json:"commands,omitempty" tf:"commands,omitempty"`

	// (List of String) Set access control to keys.
	// Set access control to keys.
	Keys []*string `json:"keys,omitempty" tf:"keys,omitempty"`
}

type RedisAccessControlObservation struct {

	// (List of String) Set access control to all commands in specified categories.
	// Set access control to all commands in specified categories.
	Categories []*string `json:"categories,omitempty" tf:"categories,omitempty"`

	// (List of String) Set access control to Pub/Sub channels.
	// Set access control to Pub/Sub channels.
	Channels []*string `json:"channels,omitempty" tf:"channels,omitempty"`

	// (List of String) Set access control to commands.
	// Set access control to commands.
	Commands []*string `json:"commands,omitempty" tf:"commands,omitempty"`

	// (List of String) Set access control to keys.
	// Set access control to keys.
	Keys []*string `json:"keys,omitempty" tf:"keys,omitempty"`
}

type RedisAccessControlParameters struct {

	// (List of String) Set access control to all commands in specified categories.
	// Set access control to all commands in specified categories.
	// +kubebuilder:validation:Optional
	Categories []*string `json:"categories,omitempty" tf:"categories,omitempty"`

	// (List of String) Set access control to Pub/Sub channels.
	// Set access control to Pub/Sub channels.
	// +kubebuilder:validation:Optional
	Channels []*string `json:"channels,omitempty" tf:"channels,omitempty"`

	// (List of String) Set access control to commands.
	// Set access control to commands.
	// +kubebuilder:validation:Optional
	Commands []*string `json:"commands,omitempty" tf:"commands,omitempty"`

	// (List of String) Set access control to keys.
	// Set access control to keys.
	// +kubebuilder:validation:Optional
	Keys []*string `json:"keys,omitempty" tf:"keys,omitempty"`
}

type RulesInitParameters struct {

	// (String) Set index name, pattern or top level API.
	// Set index name, pattern or top level API.
	Index *string `json:"index,omitempty" tf:"index,omitempty"`

	// (String) Set permission access.
	// Set permission access.
	Permission *string `json:"permission,omitempty" tf:"permission,omitempty"`
}

type RulesObservation struct {

	// (String) Set index name, pattern or top level API.
	// Set index name, pattern or top level API.
	Index *string `json:"index,omitempty" tf:"index,omitempty"`

	// (String) Set permission access.
	// Set permission access.
	Permission *string `json:"permission,omitempty" tf:"permission,omitempty"`
}

type RulesParameters struct {

	// (String) Set index name, pattern or top level API.
	// Set index name, pattern or top level API.
	// +kubebuilder:validation:Optional
	Index *string `json:"index" tf:"index,omitempty"`

	// (String) Set permission access.
	// Set permission access.
	// +kubebuilder:validation:Optional
	Permission *string `json:"permission" tf:"permission,omitempty"`
}

// ManagedDatabaseUserSpec defines the desired state of ManagedDatabaseUser
type ManagedDatabaseUserSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagedDatabaseUserParameters `json:"forProvider"`
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
	InitProvider ManagedDatabaseUserInitParameters `json:"initProvider,omitempty"`
}

// ManagedDatabaseUserStatus defines the observed state of ManagedDatabaseUser.
type ManagedDatabaseUserStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagedDatabaseUserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ManagedDatabaseUser is the Schema for the ManagedDatabaseUsers API. This resource represents a user in managed database
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,upcloud}
type ManagedDatabaseUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagedDatabaseUserSpec   `json:"spec"`
	Status            ManagedDatabaseUserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedDatabaseUserList contains a list of ManagedDatabaseUsers
type ManagedDatabaseUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedDatabaseUser `json:"items"`
}

// Repository type metadata.
var (
	ManagedDatabaseUser_Kind             = "ManagedDatabaseUser"
	ManagedDatabaseUser_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ManagedDatabaseUser_Kind}.String()
	ManagedDatabaseUser_KindAPIVersion   = ManagedDatabaseUser_Kind + "." + CRDGroupVersion.String()
	ManagedDatabaseUser_GroupVersionKind = CRDGroupVersion.WithKind(ManagedDatabaseUser_Kind)
)

func init() {
	SchemeBuilder.Register(&ManagedDatabaseUser{}, &ManagedDatabaseUserList{})
}
