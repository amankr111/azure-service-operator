// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201201storage

import (
	"encoding/json"
	"github.com/Azure/azure-service-operator/v2/api/cache/v1alpha1api20210301storage"
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

func Test_Redis_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Redis via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedis, RedisGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedis runs a test to see if a specific instance of Redis round trips to JSON and back losslessly
func RunJSONSerializationTestForRedis(subject Redis) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Redis
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

// Generator of Redis instances for property testing - lazily instantiated by RedisGenerator()
var redisGenerator gopter.Gen

// RedisGenerator returns a generator of Redis instances for property testing.
func RedisGenerator() gopter.Gen {
	if redisGenerator != nil {
		return redisGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForRedis(generators)
	redisGenerator = gen.Struct(reflect.TypeOf(Redis{}), generators)

	return redisGenerator
}

// AddRelatedPropertyGeneratorsForRedis is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedis(gens map[string]gopter.Gen) {
	gens["Spec"] = RedisSpecGenerator()
	gens["Status"] = RedisResourceStatusGenerator()
}

func Test_RedisResource_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisResource_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisResourceStatus, RedisResourceStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisResourceStatus runs a test to see if a specific instance of RedisResource_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisResourceStatus(subject RedisResource_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisResource_Status
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

// Generator of RedisResource_Status instances for property testing - lazily instantiated by
//RedisResourceStatusGenerator()
var redisResourceStatusGenerator gopter.Gen

// RedisResourceStatusGenerator returns a generator of RedisResource_Status instances for property testing.
// We first initialize redisResourceStatusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisResourceStatusGenerator() gopter.Gen {
	if redisResourceStatusGenerator != nil {
		return redisResourceStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisResourceStatus(generators)
	redisResourceStatusGenerator = gen.Struct(reflect.TypeOf(RedisResource_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisResourceStatus(generators)
	AddRelatedPropertyGeneratorsForRedisResourceStatus(generators)
	redisResourceStatusGenerator = gen.Struct(reflect.TypeOf(RedisResource_Status{}), generators)

	return redisResourceStatusGenerator
}

// AddIndependentPropertyGeneratorsForRedisResourceStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisResourceStatus(gens map[string]gopter.Gen) {
	gens["EnableNonSslPort"] = gen.PtrOf(gen.Bool())
	gens["HostName"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["MinimumTlsVersion"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Port"] = gen.PtrOf(gen.Int())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.AlphaString())
	gens["RedisConfiguration"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["RedisVersion"] = gen.PtrOf(gen.AlphaString())
	gens["ReplicasPerMaster"] = gen.PtrOf(gen.Int())
	gens["ReplicasPerPrimary"] = gen.PtrOf(gen.Int())
	gens["ShardCount"] = gen.PtrOf(gen.Int())
	gens["SslPort"] = gen.PtrOf(gen.Int())
	gens["StaticIP"] = gen.PtrOf(gen.AlphaString())
	gens["SubnetId"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["TenantSettings"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Zones"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRedisResourceStatus is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisResourceStatus(gens map[string]gopter.Gen) {
	gens["Instances"] = gen.SliceOf(RedisInstanceDetailsStatusGenerator())
	gens["LinkedServers"] = gen.SliceOf(RedisLinkedServerStatusGenerator())
	gens["PrivateEndpointConnections"] = gen.SliceOf(PrivateEndpointConnectionStatusSubResourceEmbeddedGenerator())
	gens["Sku"] = gen.PtrOf(SkuStatusGenerator())
}

func Test_Redis_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Redis_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisSpec, RedisSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisSpec runs a test to see if a specific instance of Redis_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisSpec(subject Redis_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Redis_Spec
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

// Generator of Redis_Spec instances for property testing - lazily instantiated by RedisSpecGenerator()
var redisSpecGenerator gopter.Gen

// RedisSpecGenerator returns a generator of Redis_Spec instances for property testing.
// We first initialize redisSpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisSpecGenerator() gopter.Gen {
	if redisSpecGenerator != nil {
		return redisSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisSpec(generators)
	redisSpecGenerator = gen.Struct(reflect.TypeOf(Redis_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisSpec(generators)
	AddRelatedPropertyGeneratorsForRedisSpec(generators)
	redisSpecGenerator = gen.Struct(reflect.TypeOf(Redis_Spec{}), generators)

	return redisSpecGenerator
}

// AddIndependentPropertyGeneratorsForRedisSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisSpec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["EnableNonSslPort"] = gen.PtrOf(gen.Bool())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["MinimumTlsVersion"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.AlphaString())
	gens["RedisConfiguration"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["RedisVersion"] = gen.PtrOf(gen.AlphaString())
	gens["ReplicasPerMaster"] = gen.PtrOf(gen.Int())
	gens["ReplicasPerPrimary"] = gen.PtrOf(gen.Int())
	gens["ShardCount"] = gen.PtrOf(gen.Int())
	gens["StaticIP"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["TenantSettings"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Zones"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRedisSpec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisSpec(gens map[string]gopter.Gen) {
	gens["OperatorSpec"] = gen.PtrOf(RedisOperatorSpecGenerator())
	gens["Sku"] = gen.PtrOf(SkuGenerator())
}

func Test_PrivateEndpointConnection_Status_SubResourceEmbedded_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateEndpointConnection_Status_SubResourceEmbedded to PrivateEndpointConnection_Status_SubResourceEmbedded via AssignPropertiesToPrivateEndpointConnectionStatusSubResourceEmbedded & AssignPropertiesFromPrivateEndpointConnectionStatusSubResourceEmbedded returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateEndpointConnectionStatusSubResourceEmbedded, PrivateEndpointConnectionStatusSubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateEndpointConnectionStatusSubResourceEmbedded tests if a specific instance of PrivateEndpointConnection_Status_SubResourceEmbedded can be assigned to v1alpha1api20210301storage and back losslessly
func RunPropertyAssignmentTestForPrivateEndpointConnectionStatusSubResourceEmbedded(subject PrivateEndpointConnection_Status_SubResourceEmbedded) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20210301storage.PrivateEndpointConnection_Status_SubResourceEmbedded
	err := copied.AssignPropertiesToPrivateEndpointConnectionStatusSubResourceEmbedded(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateEndpointConnection_Status_SubResourceEmbedded
	err = actual.AssignPropertiesFromPrivateEndpointConnectionStatusSubResourceEmbedded(&other)
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

func Test_PrivateEndpointConnection_Status_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointConnection_Status_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointConnectionStatusSubResourceEmbedded, PrivateEndpointConnectionStatusSubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointConnectionStatusSubResourceEmbedded runs a test to see if a specific instance of PrivateEndpointConnection_Status_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointConnectionStatusSubResourceEmbedded(subject PrivateEndpointConnection_Status_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointConnection_Status_SubResourceEmbedded
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

// Generator of PrivateEndpointConnection_Status_SubResourceEmbedded instances for property testing - lazily
//instantiated by PrivateEndpointConnectionStatusSubResourceEmbeddedGenerator()
var privateEndpointConnectionStatusSubResourceEmbeddedGenerator gopter.Gen

// PrivateEndpointConnectionStatusSubResourceEmbeddedGenerator returns a generator of PrivateEndpointConnection_Status_SubResourceEmbedded instances for property testing.
func PrivateEndpointConnectionStatusSubResourceEmbeddedGenerator() gopter.Gen {
	if privateEndpointConnectionStatusSubResourceEmbeddedGenerator != nil {
		return privateEndpointConnectionStatusSubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointConnectionStatusSubResourceEmbedded(generators)
	privateEndpointConnectionStatusSubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointConnection_Status_SubResourceEmbedded{}), generators)

	return privateEndpointConnectionStatusSubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpointConnectionStatusSubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpointConnectionStatusSubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_RedisInstanceDetails_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisInstanceDetails_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisInstanceDetailsStatus, RedisInstanceDetailsStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisInstanceDetailsStatus runs a test to see if a specific instance of RedisInstanceDetails_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisInstanceDetailsStatus(subject RedisInstanceDetails_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisInstanceDetails_Status
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

// Generator of RedisInstanceDetails_Status instances for property testing - lazily instantiated by
//RedisInstanceDetailsStatusGenerator()
var redisInstanceDetailsStatusGenerator gopter.Gen

// RedisInstanceDetailsStatusGenerator returns a generator of RedisInstanceDetails_Status instances for property testing.
func RedisInstanceDetailsStatusGenerator() gopter.Gen {
	if redisInstanceDetailsStatusGenerator != nil {
		return redisInstanceDetailsStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisInstanceDetailsStatus(generators)
	redisInstanceDetailsStatusGenerator = gen.Struct(reflect.TypeOf(RedisInstanceDetails_Status{}), generators)

	return redisInstanceDetailsStatusGenerator
}

// AddIndependentPropertyGeneratorsForRedisInstanceDetailsStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisInstanceDetailsStatus(gens map[string]gopter.Gen) {
	gens["IsMaster"] = gen.PtrOf(gen.Bool())
	gens["IsPrimary"] = gen.PtrOf(gen.Bool())
	gens["NonSslPort"] = gen.PtrOf(gen.Int())
	gens["ShardId"] = gen.PtrOf(gen.Int())
	gens["SslPort"] = gen.PtrOf(gen.Int())
	gens["Zone"] = gen.PtrOf(gen.AlphaString())
}

func Test_RedisLinkedServer_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisLinkedServer_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisLinkedServerStatus, RedisLinkedServerStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisLinkedServerStatus runs a test to see if a specific instance of RedisLinkedServer_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisLinkedServerStatus(subject RedisLinkedServer_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisLinkedServer_Status
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

// Generator of RedisLinkedServer_Status instances for property testing - lazily instantiated by
//RedisLinkedServerStatusGenerator()
var redisLinkedServerStatusGenerator gopter.Gen

// RedisLinkedServerStatusGenerator returns a generator of RedisLinkedServer_Status instances for property testing.
func RedisLinkedServerStatusGenerator() gopter.Gen {
	if redisLinkedServerStatusGenerator != nil {
		return redisLinkedServerStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisLinkedServerStatus(generators)
	redisLinkedServerStatusGenerator = gen.Struct(reflect.TypeOf(RedisLinkedServer_Status{}), generators)

	return redisLinkedServerStatusGenerator
}

// AddIndependentPropertyGeneratorsForRedisLinkedServerStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisLinkedServerStatus(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_RedisOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisOperatorSpec, RedisOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisOperatorSpec runs a test to see if a specific instance of RedisOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisOperatorSpec(subject RedisOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisOperatorSpec
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

// Generator of RedisOperatorSpec instances for property testing - lazily instantiated by RedisOperatorSpecGenerator()
var redisOperatorSpecGenerator gopter.Gen

// RedisOperatorSpecGenerator returns a generator of RedisOperatorSpec instances for property testing.
func RedisOperatorSpecGenerator() gopter.Gen {
	if redisOperatorSpecGenerator != nil {
		return redisOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForRedisOperatorSpec(generators)
	redisOperatorSpecGenerator = gen.Struct(reflect.TypeOf(RedisOperatorSpec{}), generators)

	return redisOperatorSpecGenerator
}

// AddRelatedPropertyGeneratorsForRedisOperatorSpec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisOperatorSpec(gens map[string]gopter.Gen) {
	gens["Secrets"] = gen.PtrOf(RedisOperatorSecretsGenerator())
}

func Test_Sku_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Sku to Sku via AssignPropertiesToSku & AssignPropertiesFromSku returns original",
		prop.ForAll(RunPropertyAssignmentTestForSku, SkuGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSku tests if a specific instance of Sku can be assigned to v1alpha1api20210301storage and back losslessly
func RunPropertyAssignmentTestForSku(subject Sku) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20210301storage.Sku
	err := copied.AssignPropertiesToSku(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Sku
	err = actual.AssignPropertiesFromSku(&other)
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

func Test_Sku_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSku, SkuGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSku runs a test to see if a specific instance of Sku round trips to JSON and back losslessly
func RunJSONSerializationTestForSku(subject Sku) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku
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

// Generator of Sku instances for property testing - lazily instantiated by SkuGenerator()
var skuGenerator gopter.Gen

// SkuGenerator returns a generator of Sku instances for property testing.
func SkuGenerator() gopter.Gen {
	if skuGenerator != nil {
		return skuGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSku(generators)
	skuGenerator = gen.Struct(reflect.TypeOf(Sku{}), generators)

	return skuGenerator
}

// AddIndependentPropertyGeneratorsForSku is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSku(gens map[string]gopter.Gen) {
	gens["Capacity"] = gen.PtrOf(gen.Int())
	gens["Family"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

func Test_Sku_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Sku_Status to Sku_Status via AssignPropertiesToSkuStatus & AssignPropertiesFromSkuStatus returns original",
		prop.ForAll(RunPropertyAssignmentTestForSkuStatus, SkuStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSkuStatus tests if a specific instance of Sku_Status can be assigned to v1alpha1api20210301storage and back losslessly
func RunPropertyAssignmentTestForSkuStatus(subject Sku_Status) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20210301storage.Sku_Status
	err := copied.AssignPropertiesToSkuStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Sku_Status
	err = actual.AssignPropertiesFromSkuStatus(&other)
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

func Test_Sku_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSkuStatus, SkuStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSkuStatus runs a test to see if a specific instance of Sku_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForSkuStatus(subject Sku_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku_Status
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

// Generator of Sku_Status instances for property testing - lazily instantiated by SkuStatusGenerator()
var skuStatusGenerator gopter.Gen

// SkuStatusGenerator returns a generator of Sku_Status instances for property testing.
func SkuStatusGenerator() gopter.Gen {
	if skuStatusGenerator != nil {
		return skuStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSkuStatus(generators)
	skuStatusGenerator = gen.Struct(reflect.TypeOf(Sku_Status{}), generators)

	return skuStatusGenerator
}

// AddIndependentPropertyGeneratorsForSkuStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSkuStatus(gens map[string]gopter.Gen) {
	gens["Capacity"] = gen.PtrOf(gen.Int())
	gens["Family"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

func Test_RedisOperatorSecrets_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisOperatorSecrets via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisOperatorSecrets, RedisOperatorSecretsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisOperatorSecrets runs a test to see if a specific instance of RedisOperatorSecrets round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisOperatorSecrets(subject RedisOperatorSecrets) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisOperatorSecrets
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

// Generator of RedisOperatorSecrets instances for property testing - lazily instantiated by
//RedisOperatorSecretsGenerator()
var redisOperatorSecretsGenerator gopter.Gen

// RedisOperatorSecretsGenerator returns a generator of RedisOperatorSecrets instances for property testing.
func RedisOperatorSecretsGenerator() gopter.Gen {
	if redisOperatorSecretsGenerator != nil {
		return redisOperatorSecretsGenerator
	}

	generators := make(map[string]gopter.Gen)
	redisOperatorSecretsGenerator = gen.Struct(reflect.TypeOf(RedisOperatorSecrets{}), generators)

	return redisOperatorSecretsGenerator
}
