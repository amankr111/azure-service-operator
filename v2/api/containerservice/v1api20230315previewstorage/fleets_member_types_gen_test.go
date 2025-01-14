// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230315previewstorage

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

func Test_FleetsMember_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FleetsMember via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFleetsMember, FleetsMemberGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFleetsMember runs a test to see if a specific instance of FleetsMember round trips to JSON and back losslessly
func RunJSONSerializationTestForFleetsMember(subject FleetsMember) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FleetsMember
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

// Generator of FleetsMember instances for property testing - lazily instantiated by FleetsMemberGenerator()
var fleetsMemberGenerator gopter.Gen

// FleetsMemberGenerator returns a generator of FleetsMember instances for property testing.
func FleetsMemberGenerator() gopter.Gen {
	if fleetsMemberGenerator != nil {
		return fleetsMemberGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForFleetsMember(generators)
	fleetsMemberGenerator = gen.Struct(reflect.TypeOf(FleetsMember{}), generators)

	return fleetsMemberGenerator
}

// AddRelatedPropertyGeneratorsForFleetsMember is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFleetsMember(gens map[string]gopter.Gen) {
	gens["Spec"] = Fleets_Member_SpecGenerator()
	gens["Status"] = Fleets_Member_STATUSGenerator()
}

func Test_Fleets_Member_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Fleets_Member_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFleets_Member_Spec, Fleets_Member_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFleets_Member_Spec runs a test to see if a specific instance of Fleets_Member_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForFleets_Member_Spec(subject Fleets_Member_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Fleets_Member_Spec
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

// Generator of Fleets_Member_Spec instances for property testing - lazily instantiated by Fleets_Member_SpecGenerator()
var fleets_Member_SpecGenerator gopter.Gen

// Fleets_Member_SpecGenerator returns a generator of Fleets_Member_Spec instances for property testing.
func Fleets_Member_SpecGenerator() gopter.Gen {
	if fleets_Member_SpecGenerator != nil {
		return fleets_Member_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFleets_Member_Spec(generators)
	fleets_Member_SpecGenerator = gen.Struct(reflect.TypeOf(Fleets_Member_Spec{}), generators)

	return fleets_Member_SpecGenerator
}

// AddIndependentPropertyGeneratorsForFleets_Member_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFleets_Member_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Group"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
}

func Test_Fleets_Member_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Fleets_Member_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFleets_Member_STATUS, Fleets_Member_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFleets_Member_STATUS runs a test to see if a specific instance of Fleets_Member_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForFleets_Member_STATUS(subject Fleets_Member_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Fleets_Member_STATUS
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

// Generator of Fleets_Member_STATUS instances for property testing - lazily instantiated by
// Fleets_Member_STATUSGenerator()
var fleets_Member_STATUSGenerator gopter.Gen

// Fleets_Member_STATUSGenerator returns a generator of Fleets_Member_STATUS instances for property testing.
// We first initialize fleets_Member_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Fleets_Member_STATUSGenerator() gopter.Gen {
	if fleets_Member_STATUSGenerator != nil {
		return fleets_Member_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFleets_Member_STATUS(generators)
	fleets_Member_STATUSGenerator = gen.Struct(reflect.TypeOf(Fleets_Member_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFleets_Member_STATUS(generators)
	AddRelatedPropertyGeneratorsForFleets_Member_STATUS(generators)
	fleets_Member_STATUSGenerator = gen.Struct(reflect.TypeOf(Fleets_Member_STATUS{}), generators)

	return fleets_Member_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForFleets_Member_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFleets_Member_STATUS(gens map[string]gopter.Gen) {
	gens["ClusterResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["ETag"] = gen.PtrOf(gen.AlphaString())
	gens["Group"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFleets_Member_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFleets_Member_STATUS(gens map[string]gopter.Gen) {
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}
