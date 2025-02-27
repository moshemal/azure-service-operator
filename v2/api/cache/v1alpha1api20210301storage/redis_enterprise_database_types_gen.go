// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210301storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=cache.azure.com,resources=redisenterprisedatabases,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cache.azure.com,resources={redisenterprisedatabases/status,redisenterprisedatabases/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20210301.RedisEnterpriseDatabase
//Generated from: https://schema.management.azure.com/schemas/2021-03-01/Microsoft.Cache.Enterprise.json#/resourceDefinitions/redisEnterprise_databases
type RedisEnterpriseDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisEnterpriseDatabases_Spec `json:"spec,omitempty"`
	Status            Database_Status               `json:"status,omitempty"`
}

var _ conditions.Conditioner = &RedisEnterpriseDatabase{}

// GetConditions returns the conditions of the resource
func (database *RedisEnterpriseDatabase) GetConditions() conditions.Conditions {
	return database.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (database *RedisEnterpriseDatabase) SetConditions(conditions conditions.Conditions) {
	database.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &RedisEnterpriseDatabase{}

// AzureName returns the Azure name of the resource
func (database *RedisEnterpriseDatabase) AzureName() string {
	return database.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-03-01"
func (database RedisEnterpriseDatabase) GetAPIVersion() string {
	return "2021-03-01"
}

// GetResourceKind returns the kind of the resource
func (database *RedisEnterpriseDatabase) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (database *RedisEnterpriseDatabase) GetSpec() genruntime.ConvertibleSpec {
	return &database.Spec
}

// GetStatus returns the status of this resource
func (database *RedisEnterpriseDatabase) GetStatus() genruntime.ConvertibleStatus {
	return &database.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redisEnterprise/databases"
func (database *RedisEnterpriseDatabase) GetType() string {
	return "Microsoft.Cache/redisEnterprise/databases"
}

// NewEmptyStatus returns a new empty (blank) status
func (database *RedisEnterpriseDatabase) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Database_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (database *RedisEnterpriseDatabase) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(database.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  database.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (database *RedisEnterpriseDatabase) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Database_Status); ok {
		database.Status = *st
		return nil
	}

	// Convert status to required version
	var st Database_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	database.Status = st
	return nil
}

// Hub marks that this RedisEnterpriseDatabase is the hub type for conversion
func (database *RedisEnterpriseDatabase) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (database *RedisEnterpriseDatabase) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: database.Spec.OriginalVersion,
		Kind:    "RedisEnterpriseDatabase",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20210301.RedisEnterpriseDatabase
//Generated from: https://schema.management.azure.com/schemas/2021-03-01/Microsoft.Cache.Enterprise.json#/resourceDefinitions/redisEnterprise_databases
type RedisEnterpriseDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisEnterpriseDatabase `json:"items"`
}

//Storage version of v1alpha1api20210301.Database_Status
type Database_Status struct {
	ClientProtocol    *string                `json:"clientProtocol,omitempty"`
	ClusteringPolicy  *string                `json:"clusteringPolicy,omitempty"`
	Conditions        []conditions.Condition `json:"conditions,omitempty"`
	EvictionPolicy    *string                `json:"evictionPolicy,omitempty"`
	Id                *string                `json:"id,omitempty"`
	Modules           []Module_Status        `json:"modules,omitempty"`
	Name              *string                `json:"name,omitempty"`
	Persistence       *Persistence_Status    `json:"persistence,omitempty"`
	Port              *int                   `json:"port,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState *string                `json:"provisioningState,omitempty"`
	ResourceState     *string                `json:"resourceState,omitempty"`
	Type              *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Database_Status{}

// ConvertStatusFrom populates our Database_Status from the provided source
func (database *Database_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == database {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(database)
}

// ConvertStatusTo populates the provided destination from our Database_Status
func (database *Database_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == database {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(database)
}

//Storage version of v1alpha1api20210301.RedisEnterpriseDatabases_Spec
type RedisEnterpriseDatabases_Spec struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName        string   `json:"azureName"`
	ClientProtocol   *string  `json:"clientProtocol,omitempty"`
	ClusteringPolicy *string  `json:"clusteringPolicy,omitempty"`
	EvictionPolicy   *string  `json:"evictionPolicy,omitempty"`
	Location         *string  `json:"location,omitempty"`
	Modules          []Module `json:"modules,omitempty"`
	OriginalVersion  string   `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner       genruntime.KnownResourceReference `group:"cache.azure.com" json:"owner" kind:"RedisEnterprise"`
	Persistence *Persistence                      `json:"persistence,omitempty"`
	Port        *int                              `json:"port,omitempty"`
	PropertyBag genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	Tags        map[string]string                 `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &RedisEnterpriseDatabases_Spec{}

// ConvertSpecFrom populates our RedisEnterpriseDatabases_Spec from the provided source
func (databases *RedisEnterpriseDatabases_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == databases {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(databases)
}

// ConvertSpecTo populates the provided destination from our RedisEnterpriseDatabases_Spec
func (databases *RedisEnterpriseDatabases_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == databases {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(databases)
}

//Storage version of v1alpha1api20210301.Module
//Generated from: https://schema.management.azure.com/schemas/2021-03-01/Microsoft.Cache.Enterprise.json#/definitions/Module
type Module struct {
	Args        *string                `json:"args,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210301.Module_Status
type Module_Status struct {
	Args        *string                `json:"args,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Version     *string                `json:"version,omitempty"`
}

//Storage version of v1alpha1api20210301.Persistence
//Generated from: https://schema.management.azure.com/schemas/2021-03-01/Microsoft.Cache.Enterprise.json#/definitions/Persistence
type Persistence struct {
	AofEnabled   *bool                  `json:"aofEnabled,omitempty"`
	AofFrequency *string                `json:"aofFrequency,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RdbEnabled   *bool                  `json:"rdbEnabled,omitempty"`
	RdbFrequency *string                `json:"rdbFrequency,omitempty"`
}

//Storage version of v1alpha1api20210301.Persistence_Status
type Persistence_Status struct {
	AofEnabled   *bool                  `json:"aofEnabled,omitempty"`
	AofFrequency *string                `json:"aofFrequency,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RdbEnabled   *bool                  `json:"rdbEnabled,omitempty"`
	RdbFrequency *string                `json:"rdbFrequency,omitempty"`
}

func init() {
	SchemeBuilder.Register(&RedisEnterpriseDatabase{}, &RedisEnterpriseDatabaseList{})
}
