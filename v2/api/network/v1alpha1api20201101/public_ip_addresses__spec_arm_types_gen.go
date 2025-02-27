// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type PublicIPAddresses_SpecARM struct {
	//ExtendedLocation: The extended location of the public ip address.
	ExtendedLocation *ExtendedLocationARM `json:"extendedLocation,omitempty"`

	//Location: Location to deploy resource to
	Location string `json:"location,omitempty"`

	//Name: Name of the resource
	Name string `json:"name"`

	//Properties: Public IP address properties.
	Properties PublicIPAddressPropertiesFormatARM `json:"properties"`

	//Sku: The public IP address SKU.
	Sku *PublicIPAddressSkuARM `json:"sku,omitempty"`

	//Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`

	//Zones: A list of availability zones denoting the IP allocated for the resource
	//needs to come from.
	Zones []string `json:"zones,omitempty"`
}

var _ genruntime.ARMResourceSpec = &PublicIPAddresses_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (addresses PublicIPAddresses_SpecARM) GetAPIVersion() string {
	return "2020-11-01"
}

// GetName returns the Name of the resource
func (addresses PublicIPAddresses_SpecARM) GetName() string {
	return addresses.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/publicIPAddresses"
func (addresses PublicIPAddresses_SpecARM) GetType() string {
	return "Microsoft.Network/publicIPAddresses"
}

//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/PublicIPAddressPropertiesFormat
type PublicIPAddressPropertiesFormatARM struct {
	//DdosSettings: The DDoS protection custom policy associated with the public IP
	//address.
	DdosSettings *DdosSettingsARM `json:"ddosSettings,omitempty"`

	//DnsSettings: The FQDN of the DNS record associated with the public IP address.
	DnsSettings *PublicIPAddressDnsSettingsARM `json:"dnsSettings,omitempty"`

	//IdleTimeoutInMinutes: The idle timeout of the public IP address.
	IdleTimeoutInMinutes *int `json:"idleTimeoutInMinutes,omitempty"`

	//IpAddress: The IP address associated with the public IP address resource.
	IpAddress *string `json:"ipAddress,omitempty"`

	//IpTags: The list of tags associated with the public IP address.
	IpTags []IpTagARM `json:"ipTags,omitempty"`

	//PublicIPAddressVersion: The public IP address version.
	PublicIPAddressVersion *PublicIPAddressPropertiesFormatPublicIPAddressVersion `json:"publicIPAddressVersion,omitempty"`

	//PublicIPAllocationMethod: The public IP address allocation method.
	PublicIPAllocationMethod PublicIPAddressPropertiesFormatPublicIPAllocationMethod `json:"publicIPAllocationMethod"`

	//PublicIPPrefix: The Public IP Prefix this Public IP Address should be allocated
	//from.
	PublicIPPrefix *SubResourceARM `json:"publicIPPrefix,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/PublicIPAddressSku
type PublicIPAddressSkuARM struct {
	//Name: Name of a public IP address SKU.
	Name *PublicIPAddressSkuName `json:"name,omitempty"`

	//Tier: Tier of a public IP address SKU.
	Tier *PublicIPAddressSkuTier `json:"tier,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/DdosSettings
type DdosSettingsARM struct {
	//DdosCustomPolicy: The DDoS custom policy associated with the public IP.
	DdosCustomPolicy *SubResourceARM `json:"ddosCustomPolicy,omitempty"`

	//ProtectedIP: Enables DDoS protection on the public IP.
	ProtectedIP *bool `json:"protectedIP,omitempty"`

	//ProtectionCoverage: The DDoS protection policy customizability of the public IP.
	//Only standard coverage will have the ability to be customized.
	ProtectionCoverage *DdosSettingsProtectionCoverage `json:"protectionCoverage,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/IpTag
type IpTagARM struct {
	//IpTagType: The IP tag type. Example: FirstPartyUsage.
	IpTagType *string `json:"ipTagType,omitempty"`

	//Tag: The value of the IP tag associated with the public IP. Example: SQL.
	Tag *string `json:"tag,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/PublicIPAddressDnsSettings
type PublicIPAddressDnsSettingsARM struct {
	//DomainNameLabel: The domain name label. The concatenation of the domain name
	//label and the regionalized DNS zone make up the fully qualified domain name
	//associated with the public IP address. If a domain name label is specified, an A
	//DNS record is created for the public IP in the Microsoft Azure DNS system.
	DomainNameLabel string `json:"domainNameLabel"`

	//Fqdn: The Fully Qualified Domain Name of the A DNS record associated with the
	//public IP. This is the concatenation of the domainNameLabel and the regionalized
	//DNS zone.
	Fqdn *string `json:"fqdn,omitempty"`

	//ReverseFqdn: The reverse FQDN. A user-visible, fully qualified domain name that
	//resolves to this public IP address. If the reverseFqdn is specified, then a PTR
	//DNS record is created pointing from the IP address in the in-addr.arpa domain to
	//the reverse FQDN.
	ReverseFqdn *string `json:"reverseFqdn,omitempty"`
}

// +kubebuilder:validation:Enum={"Basic","Standard"}
type PublicIPAddressSkuName string

const (
	PublicIPAddressSkuNameBasic    = PublicIPAddressSkuName("Basic")
	PublicIPAddressSkuNameStandard = PublicIPAddressSkuName("Standard")
)

// +kubebuilder:validation:Enum={"Global","Regional"}
type PublicIPAddressSkuTier string

const (
	PublicIPAddressSkuTierGlobal   = PublicIPAddressSkuTier("Global")
	PublicIPAddressSkuTierRegional = PublicIPAddressSkuTier("Regional")
)
