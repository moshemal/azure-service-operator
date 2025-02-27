// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200801preview

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_RoleAssignment_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RoleAssignment_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRoleAssignmentStatusARM, RoleAssignmentStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRoleAssignmentStatusARM runs a test to see if a specific instance of RoleAssignment_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRoleAssignmentStatusARM(subject RoleAssignment_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RoleAssignment_StatusARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of RoleAssignment_StatusARM instances for property testing - lazily instantiated by
//RoleAssignmentStatusARMGenerator()
var roleAssignmentStatusARMGenerator gopter.Gen

// RoleAssignmentStatusARMGenerator returns a generator of RoleAssignment_StatusARM instances for property testing.
// We first initialize roleAssignmentStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RoleAssignmentStatusARMGenerator() gopter.Gen {
	if roleAssignmentStatusARMGenerator != nil {
		return roleAssignmentStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRoleAssignmentStatusARM(generators)
	roleAssignmentStatusARMGenerator = gen.Struct(reflect.TypeOf(RoleAssignment_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRoleAssignmentStatusARM(generators)
	AddRelatedPropertyGeneratorsForRoleAssignmentStatusARM(generators)
	roleAssignmentStatusARMGenerator = gen.Struct(reflect.TypeOf(RoleAssignment_StatusARM{}), generators)

	return roleAssignmentStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForRoleAssignmentStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRoleAssignmentStatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRoleAssignmentStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRoleAssignmentStatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RoleAssignmentPropertiesStatusARMGenerator())
}

func Test_RoleAssignmentProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RoleAssignmentProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRoleAssignmentPropertiesStatusARM, RoleAssignmentPropertiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRoleAssignmentPropertiesStatusARM runs a test to see if a specific instance of RoleAssignmentProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRoleAssignmentPropertiesStatusARM(subject RoleAssignmentProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RoleAssignmentProperties_StatusARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of RoleAssignmentProperties_StatusARM instances for property testing - lazily instantiated by
//RoleAssignmentPropertiesStatusARMGenerator()
var roleAssignmentPropertiesStatusARMGenerator gopter.Gen

// RoleAssignmentPropertiesStatusARMGenerator returns a generator of RoleAssignmentProperties_StatusARM instances for property testing.
func RoleAssignmentPropertiesStatusARMGenerator() gopter.Gen {
	if roleAssignmentPropertiesStatusARMGenerator != nil {
		return roleAssignmentPropertiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRoleAssignmentPropertiesStatusARM(generators)
	roleAssignmentPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(RoleAssignmentProperties_StatusARM{}), generators)

	return roleAssignmentPropertiesStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForRoleAssignmentPropertiesStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRoleAssignmentPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["Condition"] = gen.PtrOf(gen.AlphaString())
	gens["ConditionVersion"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedOn"] = gen.PtrOf(gen.AlphaString())
	gens["DelegatedManagedIdentityResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalId"] = gen.AlphaString()
	gens["PrincipalType"] = gen.PtrOf(gen.OneConstOf(
		RoleAssignmentPropertiesStatusPrincipalTypeForeignGroup,
		RoleAssignmentPropertiesStatusPrincipalTypeGroup,
		RoleAssignmentPropertiesStatusPrincipalTypeServicePrincipal,
		RoleAssignmentPropertiesStatusPrincipalTypeUser))
	gens["RoleDefinitionId"] = gen.AlphaString()
	gens["Scope"] = gen.PtrOf(gen.AlphaString())
	gens["UpdatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["UpdatedOn"] = gen.PtrOf(gen.AlphaString())
}
