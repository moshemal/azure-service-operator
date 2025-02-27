/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package genruntime_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_ResourceHierarchy_ResourceGroupOnly(t *testing.T) {
	g := NewGomegaWithT(t)

	resourceGroupName := "myrg"

	a := createResourceGroup(resourceGroupName)
	hierarchy := genruntime.ResourceHierarchy{a}

	// This is expected to fail
	_, err := hierarchy.ResourceGroup()
	g.Expect(err).To(HaveOccurred())

	expectedARMID := fmt.Sprintf("/subscriptions/1234/resourceGroups/%s", resourceGroupName)

	location, err := hierarchy.Location()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(a.Spec.Location).To(Equal(location))
	g.Expect(hierarchy.AzureName()).To(Equal(resourceGroupName))
	g.Expect(hierarchy.FullyQualifiedARMID("1234")).To(Equal(expectedARMID))
}

func Test_ResourceHierarchy_ResourceGroup_TopLevelResource(t *testing.T) {
	g := NewGomegaWithT(t)

	resourceGroupName := "myrg"
	name := "myresource"

	a, b := createResourceGroupRootedResource(resourceGroupName, name)
	hierarchy := genruntime.ResourceHierarchy{a, b}

	expectedARMID := fmt.Sprintf("/subscriptions/1234/resourceGroups/%s/providers/Microsoft.Batch/batchAccounts/%s", resourceGroupName, name)

	rg, err := hierarchy.ResourceGroup()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(rg).To(Equal(resourceGroupName))
	g.Expect(hierarchy.AzureName()).To(Equal(name))
	g.Expect(hierarchy.FullyQualifiedARMID("1234")).To(Equal(expectedARMID))
}

func Test_ResourceHierarchy_ResourceGroup_NestedResource(t *testing.T) {
	g := NewGomegaWithT(t)

	resourceGroupName := "myrg"
	resourceName := "myresource"
	childResourceName := "mychildresource"

	hierarchy := createDeeplyNestedResource(resourceGroupName, resourceName, childResourceName)

	expectedARMID := fmt.Sprintf(
		"/subscriptions/1234/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s/blobServices/%s",
		resourceGroupName,
		resourceName,
		hierarchy[2].AzureName())

	rg, err := hierarchy.ResourceGroup()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(rg).To(Equal(resourceGroupName))
	g.Expect(hierarchy.AzureName()).To(Equal(hierarchy[2].AzureName()))
	g.Expect(hierarchy.FullyQualifiedARMID("1234")).To(Equal(expectedARMID))
}

func Test_ResourceHierarchy_ExtensionOnResourceGroup(t *testing.T) {
	g := NewGomegaWithT(t)

	resourceGroupName := "myrg"
	extensionName := "myextension"

	a, b := createExtensionResourceOnResourceGroup(resourceGroupName, extensionName)
	hierarchy := genruntime.ResourceHierarchy{a, b}

	expectedARMID := fmt.Sprintf(
		"/subscriptions/1234/resourceGroups/%s/providers/Microsoft.SimpleExtension/simpleExtensions/%s",
		resourceGroupName,
		extensionName)

	g.Expect(hierarchy.ResourceGroup()).To(Equal(resourceGroupName))
	g.Expect(hierarchy.FullyQualifiedARMID("1234")).To(Equal(expectedARMID))
	g.Expect(hierarchy.AzureName()).To(Equal(extensionName))
}

func Test_ResourceHierarchy_ExtensionOnResourceInResourceGroup(t *testing.T) {
	g := NewGomegaWithT(t)

	resourceGroupName := "myrg"
	resourceName := "myresource"
	extensionName := "myextension"

	hierarchy := createExtensionResourceOnResourceInResourceGroup(resourceGroupName, resourceName, extensionName)

	expectedARMID := fmt.Sprintf(
		"/subscriptions/1234/resourceGroups/%s/providers/Microsoft.Batch/batchAccounts/%s/providers/Microsoft.SimpleExtension/simpleExtensions/%s",
		resourceGroupName,
		resourceName,
		extensionName)

	g.Expect(hierarchy.ResourceGroup()).To(Equal(resourceGroupName))
	g.Expect(hierarchy.FullyQualifiedARMID("1234")).To(Equal(expectedARMID))
	g.Expect(hierarchy.AzureName()).To(Equal(extensionName))
}

func Test_ResourceHierarchy_ExtensionOnDeepHierarchy(t *testing.T) {
	g := NewGomegaWithT(t)

	resourceGroupName := "myrg"
	resourceName := "myresource"
	childResourceName := "mychildresource"
	extensionName := "myextension"

	hierarchy := createExtensionResourceOnDeepHierarchyInResourceGroup(resourceGroupName, resourceName, childResourceName, extensionName)

	expectedARMID := fmt.Sprintf(
		"/subscriptions/1234/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s/blobServices/%s/providers/Microsoft.SimpleExtension/simpleExtensions/%s",
		resourceGroupName,
		resourceName,
		hierarchy[2].AzureName(),
		extensionName)

	g.Expect(hierarchy.ResourceGroup()).To(Equal(resourceGroupName))
	g.Expect(hierarchy.FullyQualifiedARMID("1234")).To(Equal(expectedARMID))
	g.Expect(hierarchy.AzureName()).To(Equal(extensionName))
}
