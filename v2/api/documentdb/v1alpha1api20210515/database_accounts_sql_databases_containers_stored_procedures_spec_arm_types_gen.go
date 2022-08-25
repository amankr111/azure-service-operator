// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of DatabaseAccounts_SqlDatabases_Containers_StoredProcedures_Spec. Use v1beta20210515.DatabaseAccounts_SqlDatabases_Containers_StoredProcedures_Spec instead
type DatabaseAccounts_SqlDatabases_Containers_StoredProcedures_SpecARM struct {
	Location   *string                                      `json:"location,omitempty"`
	Name       string                                       `json:"name,omitempty"`
	Properties *SqlStoredProcedureCreateUpdatePropertiesARM `json:"properties,omitempty"`
	Tags       map[string]string                            `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &DatabaseAccounts_SqlDatabases_Containers_StoredProcedures_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (procedures DatabaseAccounts_SqlDatabases_Containers_StoredProcedures_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (procedures *DatabaseAccounts_SqlDatabases_Containers_StoredProcedures_SpecARM) GetName() string {
	return procedures.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/storedProcedures"
func (procedures *DatabaseAccounts_SqlDatabases_Containers_StoredProcedures_SpecARM) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/storedProcedures"
}

// Deprecated version of SqlStoredProcedureCreateUpdateProperties. Use v1beta20210515.SqlStoredProcedureCreateUpdateProperties instead
type SqlStoredProcedureCreateUpdatePropertiesARM struct {
	Options  *CreateUpdateOptionsARM        `json:"options,omitempty"`
	Resource *SqlStoredProcedureResourceARM `json:"resource,omitempty"`
}

// Deprecated version of SqlStoredProcedureResource. Use v1beta20210515.SqlStoredProcedureResource instead
type SqlStoredProcedureResourceARM struct {
	Body *string `json:"body,omitempty"`
	Id   *string `json:"id,omitempty"`
}