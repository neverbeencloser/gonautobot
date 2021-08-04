package plugins

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new plugins API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for plugins API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	PluginsCircuitMaintenanceCircuitimpactBulkDelete(params *PluginsCircuitMaintenanceCircuitimpactBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceCircuitimpactBulkDeleteNoContent, error)

	PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdate(params *PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateOK, error)

	PluginsCircuitMaintenanceCircuitimpactBulkUpdate(params *PluginsCircuitMaintenanceCircuitimpactBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceCircuitimpactBulkUpdateOK, error)

	PluginsCircuitMaintenanceCircuitimpactCreate(params *PluginsCircuitMaintenanceCircuitimpactCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceCircuitimpactCreateCreated, error)

	PluginsCircuitMaintenanceCircuitimpactDelete(params *PluginsCircuitMaintenanceCircuitimpactDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceCircuitimpactDeleteNoContent, error)

	PluginsCircuitMaintenanceCircuitimpactList(params *PluginsCircuitMaintenanceCircuitimpactListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceCircuitimpactListOK, error)

	PluginsCircuitMaintenanceCircuitimpactPartialUpdate(params *PluginsCircuitMaintenanceCircuitimpactPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceCircuitimpactPartialUpdateOK, error)

	PluginsCircuitMaintenanceCircuitimpactRead(params *PluginsCircuitMaintenanceCircuitimpactReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceCircuitimpactReadOK, error)

	PluginsCircuitMaintenanceCircuitimpactUpdate(params *PluginsCircuitMaintenanceCircuitimpactUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceCircuitimpactUpdateOK, error)

	PluginsCircuitMaintenanceMaintenanceBulkDelete(params *PluginsCircuitMaintenanceMaintenanceBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceMaintenanceBulkDeleteNoContent, error)

	PluginsCircuitMaintenanceMaintenanceBulkPartialUpdate(params *PluginsCircuitMaintenanceMaintenanceBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceMaintenanceBulkPartialUpdateOK, error)

	PluginsCircuitMaintenanceMaintenanceBulkUpdate(params *PluginsCircuitMaintenanceMaintenanceBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceMaintenanceBulkUpdateOK, error)

	PluginsCircuitMaintenanceMaintenanceCreate(params *PluginsCircuitMaintenanceMaintenanceCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceMaintenanceCreateCreated, error)

	PluginsCircuitMaintenanceMaintenanceDelete(params *PluginsCircuitMaintenanceMaintenanceDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceMaintenanceDeleteNoContent, error)

	PluginsCircuitMaintenanceMaintenanceList(params *PluginsCircuitMaintenanceMaintenanceListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceMaintenanceListOK, error)

	PluginsCircuitMaintenanceMaintenancePartialUpdate(params *PluginsCircuitMaintenanceMaintenancePartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceMaintenancePartialUpdateOK, error)

	PluginsCircuitMaintenanceMaintenanceRead(params *PluginsCircuitMaintenanceMaintenanceReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceMaintenanceReadOK, error)

	PluginsCircuitMaintenanceMaintenanceUpdate(params *PluginsCircuitMaintenanceMaintenanceUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceMaintenanceUpdateOK, error)

	PluginsCircuitMaintenanceNoteBulkDelete(params *PluginsCircuitMaintenanceNoteBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceNoteBulkDeleteNoContent, error)

	PluginsCircuitMaintenanceNoteBulkPartialUpdate(params *PluginsCircuitMaintenanceNoteBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceNoteBulkPartialUpdateOK, error)

	PluginsCircuitMaintenanceNoteBulkUpdate(params *PluginsCircuitMaintenanceNoteBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceNoteBulkUpdateOK, error)

	PluginsCircuitMaintenanceNoteCreate(params *PluginsCircuitMaintenanceNoteCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceNoteCreateCreated, error)

	PluginsCircuitMaintenanceNoteDelete(params *PluginsCircuitMaintenanceNoteDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceNoteDeleteNoContent, error)

	PluginsCircuitMaintenanceNoteList(params *PluginsCircuitMaintenanceNoteListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceNoteListOK, error)

	PluginsCircuitMaintenanceNotePartialUpdate(params *PluginsCircuitMaintenanceNotePartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceNotePartialUpdateOK, error)

	PluginsCircuitMaintenanceNoteRead(params *PluginsCircuitMaintenanceNoteReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceNoteReadOK, error)

	PluginsCircuitMaintenanceNoteUpdate(params *PluginsCircuitMaintenanceNoteUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceNoteUpdateOK, error)

	PluginsDataValidationEngineRulesMinMaxBulkDelete(params *PluginsDataValidationEngineRulesMinMaxBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDataValidationEngineRulesMinMaxBulkDeleteNoContent, error)

	PluginsDataValidationEngineRulesMinMaxBulkPartialUpdate(params *PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateOK, error)

	PluginsDataValidationEngineRulesMinMaxBulkUpdate(params *PluginsDataValidationEngineRulesMinMaxBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDataValidationEngineRulesMinMaxBulkUpdateOK, error)

	PluginsDataValidationEngineRulesMinMaxCreate(params *PluginsDataValidationEngineRulesMinMaxCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDataValidationEngineRulesMinMaxCreateCreated, error)

	PluginsDataValidationEngineRulesMinMaxDelete(params *PluginsDataValidationEngineRulesMinMaxDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDataValidationEngineRulesMinMaxDeleteNoContent, error)

	PluginsDataValidationEngineRulesMinMaxList(params *PluginsDataValidationEngineRulesMinMaxListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDataValidationEngineRulesMinMaxListOK, error)

	PluginsDataValidationEngineRulesMinMaxPartialUpdate(params *PluginsDataValidationEngineRulesMinMaxPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDataValidationEngineRulesMinMaxPartialUpdateOK, error)

	PluginsDataValidationEngineRulesMinMaxRead(params *PluginsDataValidationEngineRulesMinMaxReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDataValidationEngineRulesMinMaxReadOK, error)

	PluginsDataValidationEngineRulesMinMaxUpdate(params *PluginsDataValidationEngineRulesMinMaxUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDataValidationEngineRulesMinMaxUpdateOK, error)

	PluginsDataValidationEngineRulesRegexBulkDelete(params *PluginsDataValidationEngineRulesRegexBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDataValidationEngineRulesRegexBulkDeleteNoContent, error)

	PluginsDataValidationEngineRulesRegexBulkPartialUpdate(params *PluginsDataValidationEngineRulesRegexBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDataValidationEngineRulesRegexBulkPartialUpdateOK, error)

	PluginsDataValidationEngineRulesRegexBulkUpdate(params *PluginsDataValidationEngineRulesRegexBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDataValidationEngineRulesRegexBulkUpdateOK, error)

	PluginsDataValidationEngineRulesRegexCreate(params *PluginsDataValidationEngineRulesRegexCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDataValidationEngineRulesRegexCreateCreated, error)

	PluginsDataValidationEngineRulesRegexDelete(params *PluginsDataValidationEngineRulesRegexDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDataValidationEngineRulesRegexDeleteNoContent, error)

	PluginsDataValidationEngineRulesRegexList(params *PluginsDataValidationEngineRulesRegexListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDataValidationEngineRulesRegexListOK, error)

	PluginsDataValidationEngineRulesRegexPartialUpdate(params *PluginsDataValidationEngineRulesRegexPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDataValidationEngineRulesRegexPartialUpdateOK, error)

	PluginsDataValidationEngineRulesRegexRead(params *PluginsDataValidationEngineRulesRegexReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDataValidationEngineRulesRegexReadOK, error)

	PluginsDataValidationEngineRulesRegexUpdate(params *PluginsDataValidationEngineRulesRegexUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDataValidationEngineRulesRegexUpdateOK, error)

	PluginsDeviceOnboardingOnboardingCreate(params *PluginsDeviceOnboardingOnboardingCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDeviceOnboardingOnboardingCreateCreated, error)

	PluginsDeviceOnboardingOnboardingDelete(params *PluginsDeviceOnboardingOnboardingDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDeviceOnboardingOnboardingDeleteNoContent, error)

	PluginsDeviceOnboardingOnboardingList(params *PluginsDeviceOnboardingOnboardingListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDeviceOnboardingOnboardingListOK, error)

	PluginsDeviceOnboardingOnboardingRead(params *PluginsDeviceOnboardingOnboardingReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDeviceOnboardingOnboardingReadOK, error)

	PluginsGoldenConfigComplianceFeatureBulkDelete(params *PluginsGoldenConfigComplianceFeatureBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigComplianceFeatureBulkDeleteNoContent, error)

	PluginsGoldenConfigComplianceFeatureBulkPartialUpdate(params *PluginsGoldenConfigComplianceFeatureBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigComplianceFeatureBulkPartialUpdateOK, error)

	PluginsGoldenConfigComplianceFeatureBulkUpdate(params *PluginsGoldenConfigComplianceFeatureBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigComplianceFeatureBulkUpdateOK, error)

	PluginsGoldenConfigComplianceFeatureCreate(params *PluginsGoldenConfigComplianceFeatureCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigComplianceFeatureCreateCreated, error)

	PluginsGoldenConfigComplianceFeatureDelete(params *PluginsGoldenConfigComplianceFeatureDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigComplianceFeatureDeleteNoContent, error)

	PluginsGoldenConfigComplianceFeatureList(params *PluginsGoldenConfigComplianceFeatureListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigComplianceFeatureListOK, error)

	PluginsGoldenConfigComplianceFeaturePartialUpdate(params *PluginsGoldenConfigComplianceFeaturePartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigComplianceFeaturePartialUpdateOK, error)

	PluginsGoldenConfigComplianceFeatureRead(params *PluginsGoldenConfigComplianceFeatureReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigComplianceFeatureReadOK, error)

	PluginsGoldenConfigComplianceFeatureUpdate(params *PluginsGoldenConfigComplianceFeatureUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigComplianceFeatureUpdateOK, error)

	PluginsGoldenConfigComplianceRuleBulkDelete(params *PluginsGoldenConfigComplianceRuleBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigComplianceRuleBulkDeleteNoContent, error)

	PluginsGoldenConfigComplianceRuleBulkPartialUpdate(params *PluginsGoldenConfigComplianceRuleBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigComplianceRuleBulkPartialUpdateOK, error)

	PluginsGoldenConfigComplianceRuleBulkUpdate(params *PluginsGoldenConfigComplianceRuleBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigComplianceRuleBulkUpdateOK, error)

	PluginsGoldenConfigComplianceRuleCreate(params *PluginsGoldenConfigComplianceRuleCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigComplianceRuleCreateCreated, error)

	PluginsGoldenConfigComplianceRuleDelete(params *PluginsGoldenConfigComplianceRuleDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigComplianceRuleDeleteNoContent, error)

	PluginsGoldenConfigComplianceRuleList(params *PluginsGoldenConfigComplianceRuleListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigComplianceRuleListOK, error)

	PluginsGoldenConfigComplianceRulePartialUpdate(params *PluginsGoldenConfigComplianceRulePartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigComplianceRulePartialUpdateOK, error)

	PluginsGoldenConfigComplianceRuleRead(params *PluginsGoldenConfigComplianceRuleReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigComplianceRuleReadOK, error)

	PluginsGoldenConfigComplianceRuleUpdate(params *PluginsGoldenConfigComplianceRuleUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigComplianceRuleUpdateOK, error)

	PluginsGoldenConfigConfigComplianceBulkDelete(params *PluginsGoldenConfigConfigComplianceBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigComplianceBulkDeleteNoContent, error)

	PluginsGoldenConfigConfigComplianceBulkPartialUpdate(params *PluginsGoldenConfigConfigComplianceBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigComplianceBulkPartialUpdateOK, error)

	PluginsGoldenConfigConfigComplianceBulkUpdate(params *PluginsGoldenConfigConfigComplianceBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigComplianceBulkUpdateOK, error)

	PluginsGoldenConfigConfigComplianceCreate(params *PluginsGoldenConfigConfigComplianceCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigComplianceCreateCreated, error)

	PluginsGoldenConfigConfigComplianceDelete(params *PluginsGoldenConfigConfigComplianceDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigComplianceDeleteNoContent, error)

	PluginsGoldenConfigConfigComplianceList(params *PluginsGoldenConfigConfigComplianceListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigComplianceListOK, error)

	PluginsGoldenConfigConfigCompliancePartialUpdate(params *PluginsGoldenConfigConfigCompliancePartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigCompliancePartialUpdateOK, error)

	PluginsGoldenConfigConfigComplianceRead(params *PluginsGoldenConfigConfigComplianceReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigComplianceReadOK, error)

	PluginsGoldenConfigConfigComplianceUpdate(params *PluginsGoldenConfigConfigComplianceUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigComplianceUpdateOK, error)

	PluginsGoldenConfigConfigRemoveBulkDelete(params *PluginsGoldenConfigConfigRemoveBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigRemoveBulkDeleteNoContent, error)

	PluginsGoldenConfigConfigRemoveBulkPartialUpdate(params *PluginsGoldenConfigConfigRemoveBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigRemoveBulkPartialUpdateOK, error)

	PluginsGoldenConfigConfigRemoveBulkUpdate(params *PluginsGoldenConfigConfigRemoveBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigRemoveBulkUpdateOK, error)

	PluginsGoldenConfigConfigRemoveCreate(params *PluginsGoldenConfigConfigRemoveCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigRemoveCreateCreated, error)

	PluginsGoldenConfigConfigRemoveDelete(params *PluginsGoldenConfigConfigRemoveDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigRemoveDeleteNoContent, error)

	PluginsGoldenConfigConfigRemoveList(params *PluginsGoldenConfigConfigRemoveListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigRemoveListOK, error)

	PluginsGoldenConfigConfigRemovePartialUpdate(params *PluginsGoldenConfigConfigRemovePartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigRemovePartialUpdateOK, error)

	PluginsGoldenConfigConfigRemoveRead(params *PluginsGoldenConfigConfigRemoveReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigRemoveReadOK, error)

	PluginsGoldenConfigConfigRemoveUpdate(params *PluginsGoldenConfigConfigRemoveUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigRemoveUpdateOK, error)

	PluginsGoldenConfigConfigReplaceBulkDelete(params *PluginsGoldenConfigConfigReplaceBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigReplaceBulkDeleteNoContent, error)

	PluginsGoldenConfigConfigReplaceBulkPartialUpdate(params *PluginsGoldenConfigConfigReplaceBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigReplaceBulkPartialUpdateOK, error)

	PluginsGoldenConfigConfigReplaceBulkUpdate(params *PluginsGoldenConfigConfigReplaceBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigReplaceBulkUpdateOK, error)

	PluginsGoldenConfigConfigReplaceCreate(params *PluginsGoldenConfigConfigReplaceCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigReplaceCreateCreated, error)

	PluginsGoldenConfigConfigReplaceDelete(params *PluginsGoldenConfigConfigReplaceDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigReplaceDeleteNoContent, error)

	PluginsGoldenConfigConfigReplaceList(params *PluginsGoldenConfigConfigReplaceListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigReplaceListOK, error)

	PluginsGoldenConfigConfigReplacePartialUpdate(params *PluginsGoldenConfigConfigReplacePartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigReplacePartialUpdateOK, error)

	PluginsGoldenConfigConfigReplaceRead(params *PluginsGoldenConfigConfigReplaceReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigReplaceReadOK, error)

	PluginsGoldenConfigConfigReplaceUpdate(params *PluginsGoldenConfigConfigReplaceUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigReplaceUpdateOK, error)

	PluginsGoldenConfigGoldenConfigSettingsBulkDelete(params *PluginsGoldenConfigGoldenConfigSettingsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigGoldenConfigSettingsBulkDeleteNoContent, error)

	PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdate(params *PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateOK, error)

	PluginsGoldenConfigGoldenConfigSettingsBulkUpdate(params *PluginsGoldenConfigGoldenConfigSettingsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigGoldenConfigSettingsBulkUpdateOK, error)

	PluginsGoldenConfigGoldenConfigSettingsCreate(params *PluginsGoldenConfigGoldenConfigSettingsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigGoldenConfigSettingsCreateCreated, error)

	PluginsGoldenConfigGoldenConfigSettingsDelete(params *PluginsGoldenConfigGoldenConfigSettingsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigGoldenConfigSettingsDeleteNoContent, error)

	PluginsGoldenConfigGoldenConfigSettingsList(params *PluginsGoldenConfigGoldenConfigSettingsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigGoldenConfigSettingsListOK, error)

	PluginsGoldenConfigGoldenConfigSettingsPartialUpdate(params *PluginsGoldenConfigGoldenConfigSettingsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigGoldenConfigSettingsPartialUpdateOK, error)

	PluginsGoldenConfigGoldenConfigSettingsRead(params *PluginsGoldenConfigGoldenConfigSettingsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigGoldenConfigSettingsReadOK, error)

	PluginsGoldenConfigGoldenConfigSettingsUpdate(params *PluginsGoldenConfigGoldenConfigSettingsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigGoldenConfigSettingsUpdateOK, error)

	PluginsGoldenConfigGoldenConfigBulkDelete(params *PluginsGoldenConfigGoldenConfigBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigGoldenConfigBulkDeleteNoContent, error)

	PluginsGoldenConfigGoldenConfigBulkPartialUpdate(params *PluginsGoldenConfigGoldenConfigBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigGoldenConfigBulkPartialUpdateOK, error)

	PluginsGoldenConfigGoldenConfigBulkUpdate(params *PluginsGoldenConfigGoldenConfigBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigGoldenConfigBulkUpdateOK, error)

	PluginsGoldenConfigGoldenConfigCreate(params *PluginsGoldenConfigGoldenConfigCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigGoldenConfigCreateCreated, error)

	PluginsGoldenConfigGoldenConfigDelete(params *PluginsGoldenConfigGoldenConfigDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigGoldenConfigDeleteNoContent, error)

	PluginsGoldenConfigGoldenConfigList(params *PluginsGoldenConfigGoldenConfigListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigGoldenConfigListOK, error)

	PluginsGoldenConfigGoldenConfigPartialUpdate(params *PluginsGoldenConfigGoldenConfigPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigGoldenConfigPartialUpdateOK, error)

	PluginsGoldenConfigGoldenConfigRead(params *PluginsGoldenConfigGoldenConfigReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigGoldenConfigReadOK, error)

	PluginsGoldenConfigGoldenConfigUpdate(params *PluginsGoldenConfigGoldenConfigUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigGoldenConfigUpdateOK, error)

	PluginsGoldenConfigSotaggRead(params *PluginsGoldenConfigSotaggReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigSotaggReadOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  PluginsCircuitMaintenanceCircuitimpactBulkDelete API view for Circuit Impact CRUD operations.
*/
func (a *Client) PluginsCircuitMaintenanceCircuitimpactBulkDelete(params *PluginsCircuitMaintenanceCircuitimpactBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceCircuitimpactBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsCircuitMaintenanceCircuitimpactBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_circuit-maintenance_circuitimpact_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/plugins/circuit-maintenance/circuitimpact/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsCircuitMaintenanceCircuitimpactBulkDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsCircuitMaintenanceCircuitimpactBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_circuit-maintenance_circuitimpact_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdate API view for Circuit Impact CRUD operations.
*/
func (a *Client) PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdate(params *PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_circuit-maintenance_circuitimpact_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/plugins/circuit-maintenance/circuitimpact/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsCircuitMaintenanceCircuitimpactBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_circuit-maintenance_circuitimpact_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsCircuitMaintenanceCircuitimpactBulkUpdate API view for Circuit Impact CRUD operations.
*/
func (a *Client) PluginsCircuitMaintenanceCircuitimpactBulkUpdate(params *PluginsCircuitMaintenanceCircuitimpactBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceCircuitimpactBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsCircuitMaintenanceCircuitimpactBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_circuit-maintenance_circuitimpact_bulk_update",
		Method:             "PUT",
		PathPattern:        "/plugins/circuit-maintenance/circuitimpact/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsCircuitMaintenanceCircuitimpactBulkUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsCircuitMaintenanceCircuitimpactBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_circuit-maintenance_circuitimpact_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsCircuitMaintenanceCircuitimpactCreate API view for Circuit Impact CRUD operations.
*/
func (a *Client) PluginsCircuitMaintenanceCircuitimpactCreate(params *PluginsCircuitMaintenanceCircuitimpactCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceCircuitimpactCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsCircuitMaintenanceCircuitimpactCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_circuit-maintenance_circuitimpact_create",
		Method:             "POST",
		PathPattern:        "/plugins/circuit-maintenance/circuitimpact/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsCircuitMaintenanceCircuitimpactCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsCircuitMaintenanceCircuitimpactCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_circuit-maintenance_circuitimpact_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsCircuitMaintenanceCircuitimpactDelete API view for Circuit Impact CRUD operations.
*/
func (a *Client) PluginsCircuitMaintenanceCircuitimpactDelete(params *PluginsCircuitMaintenanceCircuitimpactDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceCircuitimpactDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsCircuitMaintenanceCircuitimpactDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_circuit-maintenance_circuitimpact_delete",
		Method:             "DELETE",
		PathPattern:        "/plugins/circuit-maintenance/circuitimpact/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsCircuitMaintenanceCircuitimpactDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsCircuitMaintenanceCircuitimpactDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_circuit-maintenance_circuitimpact_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsCircuitMaintenanceCircuitimpactList API view for Circuit Impact CRUD operations.
*/
func (a *Client) PluginsCircuitMaintenanceCircuitimpactList(params *PluginsCircuitMaintenanceCircuitimpactListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceCircuitimpactListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsCircuitMaintenanceCircuitimpactListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_circuit-maintenance_circuitimpact_list",
		Method:             "GET",
		PathPattern:        "/plugins/circuit-maintenance/circuitimpact/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsCircuitMaintenanceCircuitimpactListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsCircuitMaintenanceCircuitimpactListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_circuit-maintenance_circuitimpact_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsCircuitMaintenanceCircuitimpactPartialUpdate API view for Circuit Impact CRUD operations.
*/
func (a *Client) PluginsCircuitMaintenanceCircuitimpactPartialUpdate(params *PluginsCircuitMaintenanceCircuitimpactPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceCircuitimpactPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsCircuitMaintenanceCircuitimpactPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_circuit-maintenance_circuitimpact_partial_update",
		Method:             "PATCH",
		PathPattern:        "/plugins/circuit-maintenance/circuitimpact/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsCircuitMaintenanceCircuitimpactPartialUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsCircuitMaintenanceCircuitimpactPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_circuit-maintenance_circuitimpact_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsCircuitMaintenanceCircuitimpactRead API view for Circuit Impact CRUD operations.
*/
func (a *Client) PluginsCircuitMaintenanceCircuitimpactRead(params *PluginsCircuitMaintenanceCircuitimpactReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceCircuitimpactReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsCircuitMaintenanceCircuitimpactReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_circuit-maintenance_circuitimpact_read",
		Method:             "GET",
		PathPattern:        "/plugins/circuit-maintenance/circuitimpact/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsCircuitMaintenanceCircuitimpactReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsCircuitMaintenanceCircuitimpactReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_circuit-maintenance_circuitimpact_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsCircuitMaintenanceCircuitimpactUpdate API view for Circuit Impact CRUD operations.
*/
func (a *Client) PluginsCircuitMaintenanceCircuitimpactUpdate(params *PluginsCircuitMaintenanceCircuitimpactUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceCircuitimpactUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsCircuitMaintenanceCircuitimpactUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_circuit-maintenance_circuitimpact_update",
		Method:             "PUT",
		PathPattern:        "/plugins/circuit-maintenance/circuitimpact/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsCircuitMaintenanceCircuitimpactUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsCircuitMaintenanceCircuitimpactUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_circuit-maintenance_circuitimpact_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsCircuitMaintenanceMaintenanceBulkDelete API view for Circuit Maintenance CRUD operations.
*/
func (a *Client) PluginsCircuitMaintenanceMaintenanceBulkDelete(params *PluginsCircuitMaintenanceMaintenanceBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceMaintenanceBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsCircuitMaintenanceMaintenanceBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_circuit-maintenance_maintenance_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/plugins/circuit-maintenance/maintenance/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsCircuitMaintenanceMaintenanceBulkDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsCircuitMaintenanceMaintenanceBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_circuit-maintenance_maintenance_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsCircuitMaintenanceMaintenanceBulkPartialUpdate API view for Circuit Maintenance CRUD operations.
*/
func (a *Client) PluginsCircuitMaintenanceMaintenanceBulkPartialUpdate(params *PluginsCircuitMaintenanceMaintenanceBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceMaintenanceBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsCircuitMaintenanceMaintenanceBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_circuit-maintenance_maintenance_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/plugins/circuit-maintenance/maintenance/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsCircuitMaintenanceMaintenanceBulkPartialUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsCircuitMaintenanceMaintenanceBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_circuit-maintenance_maintenance_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsCircuitMaintenanceMaintenanceBulkUpdate API view for Circuit Maintenance CRUD operations.
*/
func (a *Client) PluginsCircuitMaintenanceMaintenanceBulkUpdate(params *PluginsCircuitMaintenanceMaintenanceBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceMaintenanceBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsCircuitMaintenanceMaintenanceBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_circuit-maintenance_maintenance_bulk_update",
		Method:             "PUT",
		PathPattern:        "/plugins/circuit-maintenance/maintenance/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsCircuitMaintenanceMaintenanceBulkUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsCircuitMaintenanceMaintenanceBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_circuit-maintenance_maintenance_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsCircuitMaintenanceMaintenanceCreate API view for Circuit Maintenance CRUD operations.
*/
func (a *Client) PluginsCircuitMaintenanceMaintenanceCreate(params *PluginsCircuitMaintenanceMaintenanceCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceMaintenanceCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsCircuitMaintenanceMaintenanceCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_circuit-maintenance_maintenance_create",
		Method:             "POST",
		PathPattern:        "/plugins/circuit-maintenance/maintenance/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsCircuitMaintenanceMaintenanceCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsCircuitMaintenanceMaintenanceCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_circuit-maintenance_maintenance_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsCircuitMaintenanceMaintenanceDelete API view for Circuit Maintenance CRUD operations.
*/
func (a *Client) PluginsCircuitMaintenanceMaintenanceDelete(params *PluginsCircuitMaintenanceMaintenanceDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceMaintenanceDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsCircuitMaintenanceMaintenanceDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_circuit-maintenance_maintenance_delete",
		Method:             "DELETE",
		PathPattern:        "/plugins/circuit-maintenance/maintenance/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsCircuitMaintenanceMaintenanceDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsCircuitMaintenanceMaintenanceDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_circuit-maintenance_maintenance_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsCircuitMaintenanceMaintenanceList API view for Circuit Maintenance CRUD operations.
*/
func (a *Client) PluginsCircuitMaintenanceMaintenanceList(params *PluginsCircuitMaintenanceMaintenanceListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceMaintenanceListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsCircuitMaintenanceMaintenanceListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_circuit-maintenance_maintenance_list",
		Method:             "GET",
		PathPattern:        "/plugins/circuit-maintenance/maintenance/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsCircuitMaintenanceMaintenanceListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsCircuitMaintenanceMaintenanceListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_circuit-maintenance_maintenance_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsCircuitMaintenanceMaintenancePartialUpdate API view for Circuit Maintenance CRUD operations.
*/
func (a *Client) PluginsCircuitMaintenanceMaintenancePartialUpdate(params *PluginsCircuitMaintenanceMaintenancePartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceMaintenancePartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsCircuitMaintenanceMaintenancePartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_circuit-maintenance_maintenance_partial_update",
		Method:             "PATCH",
		PathPattern:        "/plugins/circuit-maintenance/maintenance/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsCircuitMaintenanceMaintenancePartialUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsCircuitMaintenanceMaintenancePartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_circuit-maintenance_maintenance_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsCircuitMaintenanceMaintenanceRead API view for Circuit Maintenance CRUD operations.
*/
func (a *Client) PluginsCircuitMaintenanceMaintenanceRead(params *PluginsCircuitMaintenanceMaintenanceReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceMaintenanceReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsCircuitMaintenanceMaintenanceReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_circuit-maintenance_maintenance_read",
		Method:             "GET",
		PathPattern:        "/plugins/circuit-maintenance/maintenance/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsCircuitMaintenanceMaintenanceReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsCircuitMaintenanceMaintenanceReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_circuit-maintenance_maintenance_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsCircuitMaintenanceMaintenanceUpdate API view for Circuit Maintenance CRUD operations.
*/
func (a *Client) PluginsCircuitMaintenanceMaintenanceUpdate(params *PluginsCircuitMaintenanceMaintenanceUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceMaintenanceUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsCircuitMaintenanceMaintenanceUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_circuit-maintenance_maintenance_update",
		Method:             "PUT",
		PathPattern:        "/plugins/circuit-maintenance/maintenance/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsCircuitMaintenanceMaintenanceUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsCircuitMaintenanceMaintenanceUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_circuit-maintenance_maintenance_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsCircuitMaintenanceNoteBulkDelete API view for Circuit Note CRUD operations.
*/
func (a *Client) PluginsCircuitMaintenanceNoteBulkDelete(params *PluginsCircuitMaintenanceNoteBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceNoteBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsCircuitMaintenanceNoteBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_circuit-maintenance_note_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/plugins/circuit-maintenance/note/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsCircuitMaintenanceNoteBulkDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsCircuitMaintenanceNoteBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_circuit-maintenance_note_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsCircuitMaintenanceNoteBulkPartialUpdate API view for Circuit Note CRUD operations.
*/
func (a *Client) PluginsCircuitMaintenanceNoteBulkPartialUpdate(params *PluginsCircuitMaintenanceNoteBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceNoteBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsCircuitMaintenanceNoteBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_circuit-maintenance_note_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/plugins/circuit-maintenance/note/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsCircuitMaintenanceNoteBulkPartialUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsCircuitMaintenanceNoteBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_circuit-maintenance_note_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsCircuitMaintenanceNoteBulkUpdate API view for Circuit Note CRUD operations.
*/
func (a *Client) PluginsCircuitMaintenanceNoteBulkUpdate(params *PluginsCircuitMaintenanceNoteBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceNoteBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsCircuitMaintenanceNoteBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_circuit-maintenance_note_bulk_update",
		Method:             "PUT",
		PathPattern:        "/plugins/circuit-maintenance/note/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsCircuitMaintenanceNoteBulkUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsCircuitMaintenanceNoteBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_circuit-maintenance_note_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsCircuitMaintenanceNoteCreate API view for Circuit Note CRUD operations.
*/
func (a *Client) PluginsCircuitMaintenanceNoteCreate(params *PluginsCircuitMaintenanceNoteCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceNoteCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsCircuitMaintenanceNoteCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_circuit-maintenance_note_create",
		Method:             "POST",
		PathPattern:        "/plugins/circuit-maintenance/note/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsCircuitMaintenanceNoteCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsCircuitMaintenanceNoteCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_circuit-maintenance_note_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsCircuitMaintenanceNoteDelete API view for Circuit Note CRUD operations.
*/
func (a *Client) PluginsCircuitMaintenanceNoteDelete(params *PluginsCircuitMaintenanceNoteDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceNoteDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsCircuitMaintenanceNoteDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_circuit-maintenance_note_delete",
		Method:             "DELETE",
		PathPattern:        "/plugins/circuit-maintenance/note/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsCircuitMaintenanceNoteDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsCircuitMaintenanceNoteDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_circuit-maintenance_note_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsCircuitMaintenanceNoteList API view for Circuit Note CRUD operations.
*/
func (a *Client) PluginsCircuitMaintenanceNoteList(params *PluginsCircuitMaintenanceNoteListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceNoteListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsCircuitMaintenanceNoteListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_circuit-maintenance_note_list",
		Method:             "GET",
		PathPattern:        "/plugins/circuit-maintenance/note/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsCircuitMaintenanceNoteListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsCircuitMaintenanceNoteListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_circuit-maintenance_note_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsCircuitMaintenanceNotePartialUpdate API view for Circuit Note CRUD operations.
*/
func (a *Client) PluginsCircuitMaintenanceNotePartialUpdate(params *PluginsCircuitMaintenanceNotePartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceNotePartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsCircuitMaintenanceNotePartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_circuit-maintenance_note_partial_update",
		Method:             "PATCH",
		PathPattern:        "/plugins/circuit-maintenance/note/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsCircuitMaintenanceNotePartialUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsCircuitMaintenanceNotePartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_circuit-maintenance_note_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsCircuitMaintenanceNoteRead API view for Circuit Note CRUD operations.
*/
func (a *Client) PluginsCircuitMaintenanceNoteRead(params *PluginsCircuitMaintenanceNoteReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceNoteReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsCircuitMaintenanceNoteReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_circuit-maintenance_note_read",
		Method:             "GET",
		PathPattern:        "/plugins/circuit-maintenance/note/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsCircuitMaintenanceNoteReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsCircuitMaintenanceNoteReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_circuit-maintenance_note_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsCircuitMaintenanceNoteUpdate API view for Circuit Note CRUD operations.
*/
func (a *Client) PluginsCircuitMaintenanceNoteUpdate(params *PluginsCircuitMaintenanceNoteUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsCircuitMaintenanceNoteUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsCircuitMaintenanceNoteUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_circuit-maintenance_note_update",
		Method:             "PUT",
		PathPattern:        "/plugins/circuit-maintenance/note/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsCircuitMaintenanceNoteUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsCircuitMaintenanceNoteUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_circuit-maintenance_note_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsDataValidationEngineRulesMinMaxBulkDelete View to manage min max expression validation rules
*/
func (a *Client) PluginsDataValidationEngineRulesMinMaxBulkDelete(params *PluginsDataValidationEngineRulesMinMaxBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDataValidationEngineRulesMinMaxBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsDataValidationEngineRulesMinMaxBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_data-validation-engine_rules_min-max_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/plugins/data-validation-engine/rules/min-max/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsDataValidationEngineRulesMinMaxBulkDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsDataValidationEngineRulesMinMaxBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_data-validation-engine_rules_min-max_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsDataValidationEngineRulesMinMaxBulkPartialUpdate View to manage min max expression validation rules
*/
func (a *Client) PluginsDataValidationEngineRulesMinMaxBulkPartialUpdate(params *PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsDataValidationEngineRulesMinMaxBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_data-validation-engine_rules_min-max_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/plugins/data-validation-engine/rules/min-max/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsDataValidationEngineRulesMinMaxBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_data-validation-engine_rules_min-max_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsDataValidationEngineRulesMinMaxBulkUpdate View to manage min max expression validation rules
*/
func (a *Client) PluginsDataValidationEngineRulesMinMaxBulkUpdate(params *PluginsDataValidationEngineRulesMinMaxBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDataValidationEngineRulesMinMaxBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsDataValidationEngineRulesMinMaxBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_data-validation-engine_rules_min-max_bulk_update",
		Method:             "PUT",
		PathPattern:        "/plugins/data-validation-engine/rules/min-max/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsDataValidationEngineRulesMinMaxBulkUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsDataValidationEngineRulesMinMaxBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_data-validation-engine_rules_min-max_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsDataValidationEngineRulesMinMaxCreate View to manage min max expression validation rules
*/
func (a *Client) PluginsDataValidationEngineRulesMinMaxCreate(params *PluginsDataValidationEngineRulesMinMaxCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDataValidationEngineRulesMinMaxCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsDataValidationEngineRulesMinMaxCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_data-validation-engine_rules_min-max_create",
		Method:             "POST",
		PathPattern:        "/plugins/data-validation-engine/rules/min-max/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsDataValidationEngineRulesMinMaxCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsDataValidationEngineRulesMinMaxCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_data-validation-engine_rules_min-max_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsDataValidationEngineRulesMinMaxDelete View to manage min max expression validation rules
*/
func (a *Client) PluginsDataValidationEngineRulesMinMaxDelete(params *PluginsDataValidationEngineRulesMinMaxDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDataValidationEngineRulesMinMaxDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsDataValidationEngineRulesMinMaxDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_data-validation-engine_rules_min-max_delete",
		Method:             "DELETE",
		PathPattern:        "/plugins/data-validation-engine/rules/min-max/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsDataValidationEngineRulesMinMaxDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsDataValidationEngineRulesMinMaxDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_data-validation-engine_rules_min-max_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsDataValidationEngineRulesMinMaxList View to manage min max expression validation rules
*/
func (a *Client) PluginsDataValidationEngineRulesMinMaxList(params *PluginsDataValidationEngineRulesMinMaxListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDataValidationEngineRulesMinMaxListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsDataValidationEngineRulesMinMaxListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_data-validation-engine_rules_min-max_list",
		Method:             "GET",
		PathPattern:        "/plugins/data-validation-engine/rules/min-max/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsDataValidationEngineRulesMinMaxListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsDataValidationEngineRulesMinMaxListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_data-validation-engine_rules_min-max_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsDataValidationEngineRulesMinMaxPartialUpdate View to manage min max expression validation rules
*/
func (a *Client) PluginsDataValidationEngineRulesMinMaxPartialUpdate(params *PluginsDataValidationEngineRulesMinMaxPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDataValidationEngineRulesMinMaxPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsDataValidationEngineRulesMinMaxPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_data-validation-engine_rules_min-max_partial_update",
		Method:             "PATCH",
		PathPattern:        "/plugins/data-validation-engine/rules/min-max/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsDataValidationEngineRulesMinMaxPartialUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsDataValidationEngineRulesMinMaxPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_data-validation-engine_rules_min-max_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsDataValidationEngineRulesMinMaxRead View to manage min max expression validation rules
*/
func (a *Client) PluginsDataValidationEngineRulesMinMaxRead(params *PluginsDataValidationEngineRulesMinMaxReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDataValidationEngineRulesMinMaxReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsDataValidationEngineRulesMinMaxReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_data-validation-engine_rules_min-max_read",
		Method:             "GET",
		PathPattern:        "/plugins/data-validation-engine/rules/min-max/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsDataValidationEngineRulesMinMaxReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsDataValidationEngineRulesMinMaxReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_data-validation-engine_rules_min-max_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsDataValidationEngineRulesMinMaxUpdate View to manage min max expression validation rules
*/
func (a *Client) PluginsDataValidationEngineRulesMinMaxUpdate(params *PluginsDataValidationEngineRulesMinMaxUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDataValidationEngineRulesMinMaxUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsDataValidationEngineRulesMinMaxUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_data-validation-engine_rules_min-max_update",
		Method:             "PUT",
		PathPattern:        "/plugins/data-validation-engine/rules/min-max/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsDataValidationEngineRulesMinMaxUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsDataValidationEngineRulesMinMaxUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_data-validation-engine_rules_min-max_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsDataValidationEngineRulesRegexBulkDelete View to manage regular expression validation rules
*/
func (a *Client) PluginsDataValidationEngineRulesRegexBulkDelete(params *PluginsDataValidationEngineRulesRegexBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDataValidationEngineRulesRegexBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsDataValidationEngineRulesRegexBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_data-validation-engine_rules_regex_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/plugins/data-validation-engine/rules/regex/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsDataValidationEngineRulesRegexBulkDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsDataValidationEngineRulesRegexBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_data-validation-engine_rules_regex_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsDataValidationEngineRulesRegexBulkPartialUpdate View to manage regular expression validation rules
*/
func (a *Client) PluginsDataValidationEngineRulesRegexBulkPartialUpdate(params *PluginsDataValidationEngineRulesRegexBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDataValidationEngineRulesRegexBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsDataValidationEngineRulesRegexBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_data-validation-engine_rules_regex_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/plugins/data-validation-engine/rules/regex/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsDataValidationEngineRulesRegexBulkPartialUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsDataValidationEngineRulesRegexBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_data-validation-engine_rules_regex_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsDataValidationEngineRulesRegexBulkUpdate View to manage regular expression validation rules
*/
func (a *Client) PluginsDataValidationEngineRulesRegexBulkUpdate(params *PluginsDataValidationEngineRulesRegexBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDataValidationEngineRulesRegexBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsDataValidationEngineRulesRegexBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_data-validation-engine_rules_regex_bulk_update",
		Method:             "PUT",
		PathPattern:        "/plugins/data-validation-engine/rules/regex/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsDataValidationEngineRulesRegexBulkUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsDataValidationEngineRulesRegexBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_data-validation-engine_rules_regex_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsDataValidationEngineRulesRegexCreate View to manage regular expression validation rules
*/
func (a *Client) PluginsDataValidationEngineRulesRegexCreate(params *PluginsDataValidationEngineRulesRegexCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDataValidationEngineRulesRegexCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsDataValidationEngineRulesRegexCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_data-validation-engine_rules_regex_create",
		Method:             "POST",
		PathPattern:        "/plugins/data-validation-engine/rules/regex/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsDataValidationEngineRulesRegexCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsDataValidationEngineRulesRegexCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_data-validation-engine_rules_regex_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsDataValidationEngineRulesRegexDelete View to manage regular expression validation rules
*/
func (a *Client) PluginsDataValidationEngineRulesRegexDelete(params *PluginsDataValidationEngineRulesRegexDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDataValidationEngineRulesRegexDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsDataValidationEngineRulesRegexDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_data-validation-engine_rules_regex_delete",
		Method:             "DELETE",
		PathPattern:        "/plugins/data-validation-engine/rules/regex/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsDataValidationEngineRulesRegexDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsDataValidationEngineRulesRegexDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_data-validation-engine_rules_regex_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsDataValidationEngineRulesRegexList View to manage regular expression validation rules
*/
func (a *Client) PluginsDataValidationEngineRulesRegexList(params *PluginsDataValidationEngineRulesRegexListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDataValidationEngineRulesRegexListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsDataValidationEngineRulesRegexListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_data-validation-engine_rules_regex_list",
		Method:             "GET",
		PathPattern:        "/plugins/data-validation-engine/rules/regex/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsDataValidationEngineRulesRegexListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsDataValidationEngineRulesRegexListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_data-validation-engine_rules_regex_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsDataValidationEngineRulesRegexPartialUpdate View to manage regular expression validation rules
*/
func (a *Client) PluginsDataValidationEngineRulesRegexPartialUpdate(params *PluginsDataValidationEngineRulesRegexPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDataValidationEngineRulesRegexPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsDataValidationEngineRulesRegexPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_data-validation-engine_rules_regex_partial_update",
		Method:             "PATCH",
		PathPattern:        "/plugins/data-validation-engine/rules/regex/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsDataValidationEngineRulesRegexPartialUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsDataValidationEngineRulesRegexPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_data-validation-engine_rules_regex_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsDataValidationEngineRulesRegexRead View to manage regular expression validation rules
*/
func (a *Client) PluginsDataValidationEngineRulesRegexRead(params *PluginsDataValidationEngineRulesRegexReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDataValidationEngineRulesRegexReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsDataValidationEngineRulesRegexReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_data-validation-engine_rules_regex_read",
		Method:             "GET",
		PathPattern:        "/plugins/data-validation-engine/rules/regex/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsDataValidationEngineRulesRegexReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsDataValidationEngineRulesRegexReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_data-validation-engine_rules_regex_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsDataValidationEngineRulesRegexUpdate View to manage regular expression validation rules
*/
func (a *Client) PluginsDataValidationEngineRulesRegexUpdate(params *PluginsDataValidationEngineRulesRegexUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDataValidationEngineRulesRegexUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsDataValidationEngineRulesRegexUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_data-validation-engine_rules_regex_update",
		Method:             "PUT",
		PathPattern:        "/plugins/data-validation-engine/rules/regex/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsDataValidationEngineRulesRegexUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsDataValidationEngineRulesRegexUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_data-validation-engine_rules_regex_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsDeviceOnboardingOnboardingCreate creates check status of and delete onboarding tasks

  In-place updates (PUT, PATCH) of tasks are not permitted.
*/
func (a *Client) PluginsDeviceOnboardingOnboardingCreate(params *PluginsDeviceOnboardingOnboardingCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDeviceOnboardingOnboardingCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsDeviceOnboardingOnboardingCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_device-onboarding_onboarding_create",
		Method:             "POST",
		PathPattern:        "/plugins/device-onboarding/onboarding/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsDeviceOnboardingOnboardingCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsDeviceOnboardingOnboardingCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_device-onboarding_onboarding_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsDeviceOnboardingOnboardingDelete creates check status of and delete onboarding tasks

  In-place updates (PUT, PATCH) of tasks are not permitted.
*/
func (a *Client) PluginsDeviceOnboardingOnboardingDelete(params *PluginsDeviceOnboardingOnboardingDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDeviceOnboardingOnboardingDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsDeviceOnboardingOnboardingDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_device-onboarding_onboarding_delete",
		Method:             "DELETE",
		PathPattern:        "/plugins/device-onboarding/onboarding/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsDeviceOnboardingOnboardingDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsDeviceOnboardingOnboardingDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_device-onboarding_onboarding_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsDeviceOnboardingOnboardingList creates check status of and delete onboarding tasks

  In-place updates (PUT, PATCH) of tasks are not permitted.
*/
func (a *Client) PluginsDeviceOnboardingOnboardingList(params *PluginsDeviceOnboardingOnboardingListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDeviceOnboardingOnboardingListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsDeviceOnboardingOnboardingListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_device-onboarding_onboarding_list",
		Method:             "GET",
		PathPattern:        "/plugins/device-onboarding/onboarding/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsDeviceOnboardingOnboardingListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsDeviceOnboardingOnboardingListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_device-onboarding_onboarding_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsDeviceOnboardingOnboardingRead creates check status of and delete onboarding tasks

  In-place updates (PUT, PATCH) of tasks are not permitted.
*/
func (a *Client) PluginsDeviceOnboardingOnboardingRead(params *PluginsDeviceOnboardingOnboardingReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsDeviceOnboardingOnboardingReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsDeviceOnboardingOnboardingReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_device-onboarding_onboarding_read",
		Method:             "GET",
		PathPattern:        "/plugins/device-onboarding/onboarding/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsDeviceOnboardingOnboardingReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsDeviceOnboardingOnboardingReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_device-onboarding_onboarding_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigComplianceFeatureBulkDelete API viewset for interacting with ComplianceFeature objects.
*/
func (a *Client) PluginsGoldenConfigComplianceFeatureBulkDelete(params *PluginsGoldenConfigComplianceFeatureBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigComplianceFeatureBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigComplianceFeatureBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_compliance-feature_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/plugins/golden-config/compliance-feature/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigComplianceFeatureBulkDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigComplianceFeatureBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_compliance-feature_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigComplianceFeatureBulkPartialUpdate API viewset for interacting with ComplianceFeature objects.
*/
func (a *Client) PluginsGoldenConfigComplianceFeatureBulkPartialUpdate(params *PluginsGoldenConfigComplianceFeatureBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigComplianceFeatureBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigComplianceFeatureBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_compliance-feature_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/plugins/golden-config/compliance-feature/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigComplianceFeatureBulkPartialUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigComplianceFeatureBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_compliance-feature_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigComplianceFeatureBulkUpdate API viewset for interacting with ComplianceFeature objects.
*/
func (a *Client) PluginsGoldenConfigComplianceFeatureBulkUpdate(params *PluginsGoldenConfigComplianceFeatureBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigComplianceFeatureBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigComplianceFeatureBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_compliance-feature_bulk_update",
		Method:             "PUT",
		PathPattern:        "/plugins/golden-config/compliance-feature/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigComplianceFeatureBulkUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigComplianceFeatureBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_compliance-feature_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigComplianceFeatureCreate API viewset for interacting with ComplianceFeature objects.
*/
func (a *Client) PluginsGoldenConfigComplianceFeatureCreate(params *PluginsGoldenConfigComplianceFeatureCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigComplianceFeatureCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigComplianceFeatureCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_compliance-feature_create",
		Method:             "POST",
		PathPattern:        "/plugins/golden-config/compliance-feature/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigComplianceFeatureCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigComplianceFeatureCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_compliance-feature_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigComplianceFeatureDelete API viewset for interacting with ComplianceFeature objects.
*/
func (a *Client) PluginsGoldenConfigComplianceFeatureDelete(params *PluginsGoldenConfigComplianceFeatureDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigComplianceFeatureDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigComplianceFeatureDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_compliance-feature_delete",
		Method:             "DELETE",
		PathPattern:        "/plugins/golden-config/compliance-feature/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigComplianceFeatureDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigComplianceFeatureDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_compliance-feature_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigComplianceFeatureList API viewset for interacting with ComplianceFeature objects.
*/
func (a *Client) PluginsGoldenConfigComplianceFeatureList(params *PluginsGoldenConfigComplianceFeatureListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigComplianceFeatureListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigComplianceFeatureListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_compliance-feature_list",
		Method:             "GET",
		PathPattern:        "/plugins/golden-config/compliance-feature/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigComplianceFeatureListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigComplianceFeatureListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_compliance-feature_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigComplianceFeaturePartialUpdate API viewset for interacting with ComplianceFeature objects.
*/
func (a *Client) PluginsGoldenConfigComplianceFeaturePartialUpdate(params *PluginsGoldenConfigComplianceFeaturePartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigComplianceFeaturePartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigComplianceFeaturePartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_compliance-feature_partial_update",
		Method:             "PATCH",
		PathPattern:        "/plugins/golden-config/compliance-feature/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigComplianceFeaturePartialUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigComplianceFeaturePartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_compliance-feature_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigComplianceFeatureRead API viewset for interacting with ComplianceFeature objects.
*/
func (a *Client) PluginsGoldenConfigComplianceFeatureRead(params *PluginsGoldenConfigComplianceFeatureReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigComplianceFeatureReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigComplianceFeatureReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_compliance-feature_read",
		Method:             "GET",
		PathPattern:        "/plugins/golden-config/compliance-feature/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigComplianceFeatureReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigComplianceFeatureReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_compliance-feature_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigComplianceFeatureUpdate API viewset for interacting with ComplianceFeature objects.
*/
func (a *Client) PluginsGoldenConfigComplianceFeatureUpdate(params *PluginsGoldenConfigComplianceFeatureUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigComplianceFeatureUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigComplianceFeatureUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_compliance-feature_update",
		Method:             "PUT",
		PathPattern:        "/plugins/golden-config/compliance-feature/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigComplianceFeatureUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigComplianceFeatureUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_compliance-feature_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigComplianceRuleBulkDelete API viewset for interacting with ComplianceRule objects.
*/
func (a *Client) PluginsGoldenConfigComplianceRuleBulkDelete(params *PluginsGoldenConfigComplianceRuleBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigComplianceRuleBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigComplianceRuleBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_compliance-rule_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/plugins/golden-config/compliance-rule/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigComplianceRuleBulkDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigComplianceRuleBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_compliance-rule_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigComplianceRuleBulkPartialUpdate API viewset for interacting with ComplianceRule objects.
*/
func (a *Client) PluginsGoldenConfigComplianceRuleBulkPartialUpdate(params *PluginsGoldenConfigComplianceRuleBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigComplianceRuleBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigComplianceRuleBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_compliance-rule_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/plugins/golden-config/compliance-rule/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigComplianceRuleBulkPartialUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigComplianceRuleBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_compliance-rule_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigComplianceRuleBulkUpdate API viewset for interacting with ComplianceRule objects.
*/
func (a *Client) PluginsGoldenConfigComplianceRuleBulkUpdate(params *PluginsGoldenConfigComplianceRuleBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigComplianceRuleBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigComplianceRuleBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_compliance-rule_bulk_update",
		Method:             "PUT",
		PathPattern:        "/plugins/golden-config/compliance-rule/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigComplianceRuleBulkUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigComplianceRuleBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_compliance-rule_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigComplianceRuleCreate API viewset for interacting with ComplianceRule objects.
*/
func (a *Client) PluginsGoldenConfigComplianceRuleCreate(params *PluginsGoldenConfigComplianceRuleCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigComplianceRuleCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigComplianceRuleCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_compliance-rule_create",
		Method:             "POST",
		PathPattern:        "/plugins/golden-config/compliance-rule/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigComplianceRuleCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigComplianceRuleCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_compliance-rule_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigComplianceRuleDelete API viewset for interacting with ComplianceRule objects.
*/
func (a *Client) PluginsGoldenConfigComplianceRuleDelete(params *PluginsGoldenConfigComplianceRuleDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigComplianceRuleDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigComplianceRuleDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_compliance-rule_delete",
		Method:             "DELETE",
		PathPattern:        "/plugins/golden-config/compliance-rule/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigComplianceRuleDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigComplianceRuleDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_compliance-rule_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigComplianceRuleList API viewset for interacting with ComplianceRule objects.
*/
func (a *Client) PluginsGoldenConfigComplianceRuleList(params *PluginsGoldenConfigComplianceRuleListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigComplianceRuleListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigComplianceRuleListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_compliance-rule_list",
		Method:             "GET",
		PathPattern:        "/plugins/golden-config/compliance-rule/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigComplianceRuleListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigComplianceRuleListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_compliance-rule_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigComplianceRulePartialUpdate API viewset for interacting with ComplianceRule objects.
*/
func (a *Client) PluginsGoldenConfigComplianceRulePartialUpdate(params *PluginsGoldenConfigComplianceRulePartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigComplianceRulePartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigComplianceRulePartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_compliance-rule_partial_update",
		Method:             "PATCH",
		PathPattern:        "/plugins/golden-config/compliance-rule/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigComplianceRulePartialUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigComplianceRulePartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_compliance-rule_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigComplianceRuleRead API viewset for interacting with ComplianceRule objects.
*/
func (a *Client) PluginsGoldenConfigComplianceRuleRead(params *PluginsGoldenConfigComplianceRuleReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigComplianceRuleReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigComplianceRuleReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_compliance-rule_read",
		Method:             "GET",
		PathPattern:        "/plugins/golden-config/compliance-rule/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigComplianceRuleReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigComplianceRuleReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_compliance-rule_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigComplianceRuleUpdate API viewset for interacting with ComplianceRule objects.
*/
func (a *Client) PluginsGoldenConfigComplianceRuleUpdate(params *PluginsGoldenConfigComplianceRuleUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigComplianceRuleUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigComplianceRuleUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_compliance-rule_update",
		Method:             "PUT",
		PathPattern:        "/plugins/golden-config/compliance-rule/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigComplianceRuleUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigComplianceRuleUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_compliance-rule_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigConfigComplianceBulkDelete API viewset for interacting with ConfigCompliance objects.
*/
func (a *Client) PluginsGoldenConfigConfigComplianceBulkDelete(params *PluginsGoldenConfigConfigComplianceBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigComplianceBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigConfigComplianceBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_config-compliance_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/plugins/golden-config/config-compliance/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigConfigComplianceBulkDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigConfigComplianceBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_config-compliance_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigConfigComplianceBulkPartialUpdate API viewset for interacting with ConfigCompliance objects.
*/
func (a *Client) PluginsGoldenConfigConfigComplianceBulkPartialUpdate(params *PluginsGoldenConfigConfigComplianceBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigComplianceBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigConfigComplianceBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_config-compliance_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/plugins/golden-config/config-compliance/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigConfigComplianceBulkPartialUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigConfigComplianceBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_config-compliance_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigConfigComplianceBulkUpdate API viewset for interacting with ConfigCompliance objects.
*/
func (a *Client) PluginsGoldenConfigConfigComplianceBulkUpdate(params *PluginsGoldenConfigConfigComplianceBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigComplianceBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigConfigComplianceBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_config-compliance_bulk_update",
		Method:             "PUT",
		PathPattern:        "/plugins/golden-config/config-compliance/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigConfigComplianceBulkUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigConfigComplianceBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_config-compliance_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigConfigComplianceCreate API viewset for interacting with ConfigCompliance objects.
*/
func (a *Client) PluginsGoldenConfigConfigComplianceCreate(params *PluginsGoldenConfigConfigComplianceCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigComplianceCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigConfigComplianceCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_config-compliance_create",
		Method:             "POST",
		PathPattern:        "/plugins/golden-config/config-compliance/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigConfigComplianceCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigConfigComplianceCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_config-compliance_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigConfigComplianceDelete API viewset for interacting with ConfigCompliance objects.
*/
func (a *Client) PluginsGoldenConfigConfigComplianceDelete(params *PluginsGoldenConfigConfigComplianceDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigComplianceDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigConfigComplianceDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_config-compliance_delete",
		Method:             "DELETE",
		PathPattern:        "/plugins/golden-config/config-compliance/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigConfigComplianceDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigConfigComplianceDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_config-compliance_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigConfigComplianceList API viewset for interacting with ConfigCompliance objects.
*/
func (a *Client) PluginsGoldenConfigConfigComplianceList(params *PluginsGoldenConfigConfigComplianceListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigComplianceListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigConfigComplianceListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_config-compliance_list",
		Method:             "GET",
		PathPattern:        "/plugins/golden-config/config-compliance/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigConfigComplianceListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigConfigComplianceListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_config-compliance_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigConfigCompliancePartialUpdate API viewset for interacting with ConfigCompliance objects.
*/
func (a *Client) PluginsGoldenConfigConfigCompliancePartialUpdate(params *PluginsGoldenConfigConfigCompliancePartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigCompliancePartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigConfigCompliancePartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_config-compliance_partial_update",
		Method:             "PATCH",
		PathPattern:        "/plugins/golden-config/config-compliance/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigConfigCompliancePartialUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigConfigCompliancePartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_config-compliance_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigConfigComplianceRead API viewset for interacting with ConfigCompliance objects.
*/
func (a *Client) PluginsGoldenConfigConfigComplianceRead(params *PluginsGoldenConfigConfigComplianceReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigComplianceReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigConfigComplianceReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_config-compliance_read",
		Method:             "GET",
		PathPattern:        "/plugins/golden-config/config-compliance/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigConfigComplianceReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigConfigComplianceReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_config-compliance_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigConfigComplianceUpdate API viewset for interacting with ConfigCompliance objects.
*/
func (a *Client) PluginsGoldenConfigConfigComplianceUpdate(params *PluginsGoldenConfigConfigComplianceUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigComplianceUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigConfigComplianceUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_config-compliance_update",
		Method:             "PUT",
		PathPattern:        "/plugins/golden-config/config-compliance/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigConfigComplianceUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigConfigComplianceUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_config-compliance_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigConfigRemoveBulkDelete API viewset for interacting with ConfigRemove objects.
*/
func (a *Client) PluginsGoldenConfigConfigRemoveBulkDelete(params *PluginsGoldenConfigConfigRemoveBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigRemoveBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigConfigRemoveBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_config-remove_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/plugins/golden-config/config-remove/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigConfigRemoveBulkDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigConfigRemoveBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_config-remove_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigConfigRemoveBulkPartialUpdate API viewset for interacting with ConfigRemove objects.
*/
func (a *Client) PluginsGoldenConfigConfigRemoveBulkPartialUpdate(params *PluginsGoldenConfigConfigRemoveBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigRemoveBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigConfigRemoveBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_config-remove_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/plugins/golden-config/config-remove/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigConfigRemoveBulkPartialUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigConfigRemoveBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_config-remove_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigConfigRemoveBulkUpdate API viewset for interacting with ConfigRemove objects.
*/
func (a *Client) PluginsGoldenConfigConfigRemoveBulkUpdate(params *PluginsGoldenConfigConfigRemoveBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigRemoveBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigConfigRemoveBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_config-remove_bulk_update",
		Method:             "PUT",
		PathPattern:        "/plugins/golden-config/config-remove/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigConfigRemoveBulkUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigConfigRemoveBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_config-remove_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigConfigRemoveCreate API viewset for interacting with ConfigRemove objects.
*/
func (a *Client) PluginsGoldenConfigConfigRemoveCreate(params *PluginsGoldenConfigConfigRemoveCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigRemoveCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigConfigRemoveCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_config-remove_create",
		Method:             "POST",
		PathPattern:        "/plugins/golden-config/config-remove/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigConfigRemoveCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigConfigRemoveCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_config-remove_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigConfigRemoveDelete API viewset for interacting with ConfigRemove objects.
*/
func (a *Client) PluginsGoldenConfigConfigRemoveDelete(params *PluginsGoldenConfigConfigRemoveDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigRemoveDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigConfigRemoveDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_config-remove_delete",
		Method:             "DELETE",
		PathPattern:        "/plugins/golden-config/config-remove/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigConfigRemoveDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigConfigRemoveDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_config-remove_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigConfigRemoveList API viewset for interacting with ConfigRemove objects.
*/
func (a *Client) PluginsGoldenConfigConfigRemoveList(params *PluginsGoldenConfigConfigRemoveListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigRemoveListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigConfigRemoveListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_config-remove_list",
		Method:             "GET",
		PathPattern:        "/plugins/golden-config/config-remove/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigConfigRemoveListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigConfigRemoveListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_config-remove_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigConfigRemovePartialUpdate API viewset for interacting with ConfigRemove objects.
*/
func (a *Client) PluginsGoldenConfigConfigRemovePartialUpdate(params *PluginsGoldenConfigConfigRemovePartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigRemovePartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigConfigRemovePartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_config-remove_partial_update",
		Method:             "PATCH",
		PathPattern:        "/plugins/golden-config/config-remove/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigConfigRemovePartialUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigConfigRemovePartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_config-remove_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigConfigRemoveRead API viewset for interacting with ConfigRemove objects.
*/
func (a *Client) PluginsGoldenConfigConfigRemoveRead(params *PluginsGoldenConfigConfigRemoveReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigRemoveReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigConfigRemoveReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_config-remove_read",
		Method:             "GET",
		PathPattern:        "/plugins/golden-config/config-remove/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigConfigRemoveReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigConfigRemoveReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_config-remove_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigConfigRemoveUpdate API viewset for interacting with ConfigRemove objects.
*/
func (a *Client) PluginsGoldenConfigConfigRemoveUpdate(params *PluginsGoldenConfigConfigRemoveUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigRemoveUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigConfigRemoveUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_config-remove_update",
		Method:             "PUT",
		PathPattern:        "/plugins/golden-config/config-remove/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigConfigRemoveUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigConfigRemoveUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_config-remove_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigConfigReplaceBulkDelete API viewset for interacting with ConfigReplace objects.
*/
func (a *Client) PluginsGoldenConfigConfigReplaceBulkDelete(params *PluginsGoldenConfigConfigReplaceBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigReplaceBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigConfigReplaceBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_config-replace_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/plugins/golden-config/config-replace/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigConfigReplaceBulkDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigConfigReplaceBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_config-replace_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigConfigReplaceBulkPartialUpdate API viewset for interacting with ConfigReplace objects.
*/
func (a *Client) PluginsGoldenConfigConfigReplaceBulkPartialUpdate(params *PluginsGoldenConfigConfigReplaceBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigReplaceBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigConfigReplaceBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_config-replace_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/plugins/golden-config/config-replace/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigConfigReplaceBulkPartialUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigConfigReplaceBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_config-replace_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigConfigReplaceBulkUpdate API viewset for interacting with ConfigReplace objects.
*/
func (a *Client) PluginsGoldenConfigConfigReplaceBulkUpdate(params *PluginsGoldenConfigConfigReplaceBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigReplaceBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigConfigReplaceBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_config-replace_bulk_update",
		Method:             "PUT",
		PathPattern:        "/plugins/golden-config/config-replace/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigConfigReplaceBulkUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigConfigReplaceBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_config-replace_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigConfigReplaceCreate API viewset for interacting with ConfigReplace objects.
*/
func (a *Client) PluginsGoldenConfigConfigReplaceCreate(params *PluginsGoldenConfigConfigReplaceCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigReplaceCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigConfigReplaceCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_config-replace_create",
		Method:             "POST",
		PathPattern:        "/plugins/golden-config/config-replace/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigConfigReplaceCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigConfigReplaceCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_config-replace_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigConfigReplaceDelete API viewset for interacting with ConfigReplace objects.
*/
func (a *Client) PluginsGoldenConfigConfigReplaceDelete(params *PluginsGoldenConfigConfigReplaceDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigReplaceDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigConfigReplaceDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_config-replace_delete",
		Method:             "DELETE",
		PathPattern:        "/plugins/golden-config/config-replace/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigConfigReplaceDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigConfigReplaceDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_config-replace_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigConfigReplaceList API viewset for interacting with ConfigReplace objects.
*/
func (a *Client) PluginsGoldenConfigConfigReplaceList(params *PluginsGoldenConfigConfigReplaceListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigReplaceListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigConfigReplaceListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_config-replace_list",
		Method:             "GET",
		PathPattern:        "/plugins/golden-config/config-replace/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigConfigReplaceListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigConfigReplaceListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_config-replace_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigConfigReplacePartialUpdate API viewset for interacting with ConfigReplace objects.
*/
func (a *Client) PluginsGoldenConfigConfigReplacePartialUpdate(params *PluginsGoldenConfigConfigReplacePartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigReplacePartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigConfigReplacePartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_config-replace_partial_update",
		Method:             "PATCH",
		PathPattern:        "/plugins/golden-config/config-replace/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigConfigReplacePartialUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigConfigReplacePartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_config-replace_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigConfigReplaceRead API viewset for interacting with ConfigReplace objects.
*/
func (a *Client) PluginsGoldenConfigConfigReplaceRead(params *PluginsGoldenConfigConfigReplaceReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigReplaceReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigConfigReplaceReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_config-replace_read",
		Method:             "GET",
		PathPattern:        "/plugins/golden-config/config-replace/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigConfigReplaceReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigConfigReplaceReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_config-replace_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigConfigReplaceUpdate API viewset for interacting with ConfigReplace objects.
*/
func (a *Client) PluginsGoldenConfigConfigReplaceUpdate(params *PluginsGoldenConfigConfigReplaceUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigConfigReplaceUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigConfigReplaceUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_config-replace_update",
		Method:             "PUT",
		PathPattern:        "/plugins/golden-config/config-replace/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigConfigReplaceUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigConfigReplaceUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_config-replace_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigGoldenConfigSettingsBulkDelete API viewset for interacting with GoldenConfigSetting objects.
*/
func (a *Client) PluginsGoldenConfigGoldenConfigSettingsBulkDelete(params *PluginsGoldenConfigGoldenConfigSettingsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigGoldenConfigSettingsBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigGoldenConfigSettingsBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_golden-config-settings_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/plugins/golden-config/golden-config-settings/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigGoldenConfigSettingsBulkDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigGoldenConfigSettingsBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_golden-config-settings_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdate API viewset for interacting with GoldenConfigSetting objects.
*/
func (a *Client) PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdate(params *PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_golden-config-settings_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/plugins/golden-config/golden-config-settings/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigGoldenConfigSettingsBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_golden-config-settings_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigGoldenConfigSettingsBulkUpdate API viewset for interacting with GoldenConfigSetting objects.
*/
func (a *Client) PluginsGoldenConfigGoldenConfigSettingsBulkUpdate(params *PluginsGoldenConfigGoldenConfigSettingsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigGoldenConfigSettingsBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigGoldenConfigSettingsBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_golden-config-settings_bulk_update",
		Method:             "PUT",
		PathPattern:        "/plugins/golden-config/golden-config-settings/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigGoldenConfigSettingsBulkUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigGoldenConfigSettingsBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_golden-config-settings_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigGoldenConfigSettingsCreate API viewset for interacting with GoldenConfigSetting objects.
*/
func (a *Client) PluginsGoldenConfigGoldenConfigSettingsCreate(params *PluginsGoldenConfigGoldenConfigSettingsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigGoldenConfigSettingsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigGoldenConfigSettingsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_golden-config-settings_create",
		Method:             "POST",
		PathPattern:        "/plugins/golden-config/golden-config-settings/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigGoldenConfigSettingsCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigGoldenConfigSettingsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_golden-config-settings_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigGoldenConfigSettingsDelete API viewset for interacting with GoldenConfigSetting objects.
*/
func (a *Client) PluginsGoldenConfigGoldenConfigSettingsDelete(params *PluginsGoldenConfigGoldenConfigSettingsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigGoldenConfigSettingsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigGoldenConfigSettingsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_golden-config-settings_delete",
		Method:             "DELETE",
		PathPattern:        "/plugins/golden-config/golden-config-settings/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigGoldenConfigSettingsDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigGoldenConfigSettingsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_golden-config-settings_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigGoldenConfigSettingsList API viewset for interacting with GoldenConfigSetting objects.
*/
func (a *Client) PluginsGoldenConfigGoldenConfigSettingsList(params *PluginsGoldenConfigGoldenConfigSettingsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigGoldenConfigSettingsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigGoldenConfigSettingsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_golden-config-settings_list",
		Method:             "GET",
		PathPattern:        "/plugins/golden-config/golden-config-settings/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigGoldenConfigSettingsListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigGoldenConfigSettingsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_golden-config-settings_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigGoldenConfigSettingsPartialUpdate API viewset for interacting with GoldenConfigSetting objects.
*/
func (a *Client) PluginsGoldenConfigGoldenConfigSettingsPartialUpdate(params *PluginsGoldenConfigGoldenConfigSettingsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigGoldenConfigSettingsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigGoldenConfigSettingsPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_golden-config-settings_partial_update",
		Method:             "PATCH",
		PathPattern:        "/plugins/golden-config/golden-config-settings/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigGoldenConfigSettingsPartialUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigGoldenConfigSettingsPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_golden-config-settings_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigGoldenConfigSettingsRead API viewset for interacting with GoldenConfigSetting objects.
*/
func (a *Client) PluginsGoldenConfigGoldenConfigSettingsRead(params *PluginsGoldenConfigGoldenConfigSettingsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigGoldenConfigSettingsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigGoldenConfigSettingsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_golden-config-settings_read",
		Method:             "GET",
		PathPattern:        "/plugins/golden-config/golden-config-settings/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigGoldenConfigSettingsReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigGoldenConfigSettingsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_golden-config-settings_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigGoldenConfigSettingsUpdate API viewset for interacting with GoldenConfigSetting objects.
*/
func (a *Client) PluginsGoldenConfigGoldenConfigSettingsUpdate(params *PluginsGoldenConfigGoldenConfigSettingsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigGoldenConfigSettingsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigGoldenConfigSettingsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_golden-config-settings_update",
		Method:             "PUT",
		PathPattern:        "/plugins/golden-config/golden-config-settings/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigGoldenConfigSettingsUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigGoldenConfigSettingsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_golden-config-settings_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigGoldenConfigBulkDelete API viewset for interacting with GoldenConfig objects.
*/
func (a *Client) PluginsGoldenConfigGoldenConfigBulkDelete(params *PluginsGoldenConfigGoldenConfigBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigGoldenConfigBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigGoldenConfigBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_golden-config_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/plugins/golden-config/golden-config/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigGoldenConfigBulkDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigGoldenConfigBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_golden-config_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigGoldenConfigBulkPartialUpdate API viewset for interacting with GoldenConfig objects.
*/
func (a *Client) PluginsGoldenConfigGoldenConfigBulkPartialUpdate(params *PluginsGoldenConfigGoldenConfigBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigGoldenConfigBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigGoldenConfigBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_golden-config_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/plugins/golden-config/golden-config/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigGoldenConfigBulkPartialUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigGoldenConfigBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_golden-config_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigGoldenConfigBulkUpdate API viewset for interacting with GoldenConfig objects.
*/
func (a *Client) PluginsGoldenConfigGoldenConfigBulkUpdate(params *PluginsGoldenConfigGoldenConfigBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigGoldenConfigBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigGoldenConfigBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_golden-config_bulk_update",
		Method:             "PUT",
		PathPattern:        "/plugins/golden-config/golden-config/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigGoldenConfigBulkUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigGoldenConfigBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_golden-config_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigGoldenConfigCreate API viewset for interacting with GoldenConfig objects.
*/
func (a *Client) PluginsGoldenConfigGoldenConfigCreate(params *PluginsGoldenConfigGoldenConfigCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigGoldenConfigCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigGoldenConfigCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_golden-config_create",
		Method:             "POST",
		PathPattern:        "/plugins/golden-config/golden-config/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigGoldenConfigCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigGoldenConfigCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_golden-config_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigGoldenConfigDelete API viewset for interacting with GoldenConfig objects.
*/
func (a *Client) PluginsGoldenConfigGoldenConfigDelete(params *PluginsGoldenConfigGoldenConfigDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigGoldenConfigDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigGoldenConfigDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_golden-config_delete",
		Method:             "DELETE",
		PathPattern:        "/plugins/golden-config/golden-config/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigGoldenConfigDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigGoldenConfigDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_golden-config_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigGoldenConfigList API viewset for interacting with GoldenConfig objects.
*/
func (a *Client) PluginsGoldenConfigGoldenConfigList(params *PluginsGoldenConfigGoldenConfigListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigGoldenConfigListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigGoldenConfigListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_golden-config_list",
		Method:             "GET",
		PathPattern:        "/plugins/golden-config/golden-config/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigGoldenConfigListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigGoldenConfigListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_golden-config_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigGoldenConfigPartialUpdate API viewset for interacting with GoldenConfig objects.
*/
func (a *Client) PluginsGoldenConfigGoldenConfigPartialUpdate(params *PluginsGoldenConfigGoldenConfigPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigGoldenConfigPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigGoldenConfigPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_golden-config_partial_update",
		Method:             "PATCH",
		PathPattern:        "/plugins/golden-config/golden-config/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigGoldenConfigPartialUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigGoldenConfigPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_golden-config_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigGoldenConfigRead API viewset for interacting with GoldenConfig objects.
*/
func (a *Client) PluginsGoldenConfigGoldenConfigRead(params *PluginsGoldenConfigGoldenConfigReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigGoldenConfigReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigGoldenConfigReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_golden-config_read",
		Method:             "GET",
		PathPattern:        "/plugins/golden-config/golden-config/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigGoldenConfigReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigGoldenConfigReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_golden-config_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigGoldenConfigUpdate API viewset for interacting with GoldenConfig objects.
*/
func (a *Client) PluginsGoldenConfigGoldenConfigUpdate(params *PluginsGoldenConfigGoldenConfigUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigGoldenConfigUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigGoldenConfigUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_golden-config_update",
		Method:             "PUT",
		PathPattern:        "/plugins/golden-config/golden-config/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigGoldenConfigUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigGoldenConfigUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_golden-config_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PluginsGoldenConfigSotaggRead Get method serialize for a dictionary to json response.
*/
func (a *Client) PluginsGoldenConfigSotaggRead(params *PluginsGoldenConfigSotaggReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsGoldenConfigSotaggReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsGoldenConfigSotaggReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_golden-config_sotagg_read",
		Method:             "GET",
		PathPattern:        "/plugins/golden-config/sotagg/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PluginsGoldenConfigSotaggReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PluginsGoldenConfigSotaggReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_golden-config_sotagg_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
