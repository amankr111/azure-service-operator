// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210101preview

import (
	"encoding/json"
	"github.com/Azure/azure-service-operator/v2/api/servicebus/v1alpha1api20210101previewstorage"
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

func Test_NamespacesTopic_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesTopic to NamespacesTopic via AssignPropertiesToNamespacesTopic & AssignPropertiesFromNamespacesTopic returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesTopic, NamespacesTopicGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesTopic tests if a specific instance of NamespacesTopic can be assigned to v1alpha1api20210101previewstorage and back losslessly
func RunPropertyAssignmentTestForNamespacesTopic(subject NamespacesTopic) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20210101previewstorage.NamespacesTopic
	err := copied.AssignPropertiesToNamespacesTopic(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesTopic
	err = actual.AssignPropertiesFromNamespacesTopic(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NamespacesTopic_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesTopic via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesTopic, NamespacesTopicGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesTopic runs a test to see if a specific instance of NamespacesTopic round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesTopic(subject NamespacesTopic) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesTopic
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

// Generator of NamespacesTopic instances for property testing - lazily instantiated by NamespacesTopicGenerator()
var namespacesTopicGenerator gopter.Gen

// NamespacesTopicGenerator returns a generator of NamespacesTopic instances for property testing.
func NamespacesTopicGenerator() gopter.Gen {
	if namespacesTopicGenerator != nil {
		return namespacesTopicGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForNamespacesTopic(generators)
	namespacesTopicGenerator = gen.Struct(reflect.TypeOf(NamespacesTopic{}), generators)

	return namespacesTopicGenerator
}

// AddRelatedPropertyGeneratorsForNamespacesTopic is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesTopic(gens map[string]gopter.Gen) {
	gens["Spec"] = NamespacesTopicsSpecGenerator()
	gens["Status"] = SBTopicStatusGenerator()
}

func Test_NamespacesTopics_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesTopics_Spec to NamespacesTopics_Spec via AssignPropertiesToNamespacesTopicsSpec & AssignPropertiesFromNamespacesTopicsSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesTopicsSpec, NamespacesTopicsSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesTopicsSpec tests if a specific instance of NamespacesTopics_Spec can be assigned to v1alpha1api20210101previewstorage and back losslessly
func RunPropertyAssignmentTestForNamespacesTopicsSpec(subject NamespacesTopics_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20210101previewstorage.NamespacesTopics_Spec
	err := copied.AssignPropertiesToNamespacesTopicsSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesTopics_Spec
	err = actual.AssignPropertiesFromNamespacesTopicsSpec(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NamespacesTopics_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesTopics_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesTopicsSpec, NamespacesTopicsSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesTopicsSpec runs a test to see if a specific instance of NamespacesTopics_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesTopicsSpec(subject NamespacesTopics_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesTopics_Spec
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

// Generator of NamespacesTopics_Spec instances for property testing - lazily instantiated by
//NamespacesTopicsSpecGenerator()
var namespacesTopicsSpecGenerator gopter.Gen

// NamespacesTopicsSpecGenerator returns a generator of NamespacesTopics_Spec instances for property testing.
func NamespacesTopicsSpecGenerator() gopter.Gen {
	if namespacesTopicsSpecGenerator != nil {
		return namespacesTopicsSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesTopicsSpec(generators)
	namespacesTopicsSpecGenerator = gen.Struct(reflect.TypeOf(NamespacesTopics_Spec{}), generators)

	return namespacesTopicsSpecGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesTopicsSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesTopicsSpec(gens map[string]gopter.Gen) {
	gens["AutoDeleteOnIdle"] = gen.PtrOf(gen.AlphaString())
	gens["AzureName"] = gen.AlphaString()
	gens["DefaultMessageTimeToLive"] = gen.PtrOf(gen.AlphaString())
	gens["DuplicateDetectionHistoryTimeWindow"] = gen.PtrOf(gen.AlphaString())
	gens["EnableBatchedOperations"] = gen.PtrOf(gen.Bool())
	gens["EnableExpress"] = gen.PtrOf(gen.Bool())
	gens["EnablePartitioning"] = gen.PtrOf(gen.Bool())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["MaxSizeInMegabytes"] = gen.PtrOf(gen.Int())
	gens["RequiresDuplicateDetection"] = gen.PtrOf(gen.Bool())
	gens["SupportOrdering"] = gen.PtrOf(gen.Bool())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

func Test_SBTopic_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SBTopic_Status to SBTopic_Status via AssignPropertiesToSBTopicStatus & AssignPropertiesFromSBTopicStatus returns original",
		prop.ForAll(RunPropertyAssignmentTestForSBTopicStatus, SBTopicStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSBTopicStatus tests if a specific instance of SBTopic_Status can be assigned to v1alpha1api20210101previewstorage and back losslessly
func RunPropertyAssignmentTestForSBTopicStatus(subject SBTopic_Status) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20210101previewstorage.SBTopic_Status
	err := copied.AssignPropertiesToSBTopicStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SBTopic_Status
	err = actual.AssignPropertiesFromSBTopicStatus(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_SBTopic_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SBTopic_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSBTopicStatus, SBTopicStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSBTopicStatus runs a test to see if a specific instance of SBTopic_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForSBTopicStatus(subject SBTopic_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SBTopic_Status
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

// Generator of SBTopic_Status instances for property testing - lazily instantiated by SBTopicStatusGenerator()
var sbTopicStatusGenerator gopter.Gen

// SBTopicStatusGenerator returns a generator of SBTopic_Status instances for property testing.
// We first initialize sbTopicStatusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SBTopicStatusGenerator() gopter.Gen {
	if sbTopicStatusGenerator != nil {
		return sbTopicStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBTopicStatus(generators)
	sbTopicStatusGenerator = gen.Struct(reflect.TypeOf(SBTopic_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBTopicStatus(generators)
	AddRelatedPropertyGeneratorsForSBTopicStatus(generators)
	sbTopicStatusGenerator = gen.Struct(reflect.TypeOf(SBTopic_Status{}), generators)

	return sbTopicStatusGenerator
}

// AddIndependentPropertyGeneratorsForSBTopicStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSBTopicStatus(gens map[string]gopter.Gen) {
	gens["AccessedAt"] = gen.PtrOf(gen.AlphaString())
	gens["AutoDeleteOnIdle"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["DefaultMessageTimeToLive"] = gen.PtrOf(gen.AlphaString())
	gens["DuplicateDetectionHistoryTimeWindow"] = gen.PtrOf(gen.AlphaString())
	gens["EnableBatchedOperations"] = gen.PtrOf(gen.Bool())
	gens["EnableExpress"] = gen.PtrOf(gen.Bool())
	gens["EnablePartitioning"] = gen.PtrOf(gen.Bool())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["MaxSizeInMegabytes"] = gen.PtrOf(gen.Int())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["RequiresDuplicateDetection"] = gen.PtrOf(gen.Bool())
	gens["SizeInBytes"] = gen.PtrOf(gen.Int())
	gens["Status"] = gen.PtrOf(gen.OneConstOf(
		EntityStatus_StatusActive,
		EntityStatus_StatusCreating,
		EntityStatus_StatusDeleting,
		EntityStatus_StatusDisabled,
		EntityStatus_StatusReceiveDisabled,
		EntityStatus_StatusRenaming,
		EntityStatus_StatusRestoring,
		EntityStatus_StatusSendDisabled,
		EntityStatus_StatusUnknown))
	gens["SubscriptionCount"] = gen.PtrOf(gen.Int())
	gens["SupportOrdering"] = gen.PtrOf(gen.Bool())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["UpdatedAt"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSBTopicStatus is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSBTopicStatus(gens map[string]gopter.Gen) {
	gens["CountDetails"] = gen.PtrOf(MessageCountDetailsStatusGenerator())
	gens["SystemData"] = gen.PtrOf(SystemDataStatusGenerator())
}