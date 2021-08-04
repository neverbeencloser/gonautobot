package extras

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new extras API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for extras API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ExtrasComputedFieldsBulkDelete(params *ExtrasComputedFieldsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasComputedFieldsBulkDeleteNoContent, error)

	ExtrasComputedFieldsBulkPartialUpdate(params *ExtrasComputedFieldsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasComputedFieldsBulkPartialUpdateOK, error)

	ExtrasComputedFieldsBulkUpdate(params *ExtrasComputedFieldsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasComputedFieldsBulkUpdateOK, error)

	ExtrasComputedFieldsCreate(params *ExtrasComputedFieldsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasComputedFieldsCreateCreated, error)

	ExtrasComputedFieldsDelete(params *ExtrasComputedFieldsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasComputedFieldsDeleteNoContent, error)

	ExtrasComputedFieldsList(params *ExtrasComputedFieldsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasComputedFieldsListOK, error)

	ExtrasComputedFieldsPartialUpdate(params *ExtrasComputedFieldsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasComputedFieldsPartialUpdateOK, error)

	ExtrasComputedFieldsRead(params *ExtrasComputedFieldsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasComputedFieldsReadOK, error)

	ExtrasComputedFieldsUpdate(params *ExtrasComputedFieldsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasComputedFieldsUpdateOK, error)

	ExtrasConfigContextSchemasBulkDelete(params *ExtrasConfigContextSchemasBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasConfigContextSchemasBulkDeleteNoContent, error)

	ExtrasConfigContextSchemasBulkPartialUpdate(params *ExtrasConfigContextSchemasBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasConfigContextSchemasBulkPartialUpdateOK, error)

	ExtrasConfigContextSchemasBulkUpdate(params *ExtrasConfigContextSchemasBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasConfigContextSchemasBulkUpdateOK, error)

	ExtrasConfigContextSchemasCreate(params *ExtrasConfigContextSchemasCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasConfigContextSchemasCreateCreated, error)

	ExtrasConfigContextSchemasDelete(params *ExtrasConfigContextSchemasDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasConfigContextSchemasDeleteNoContent, error)

	ExtrasConfigContextSchemasList(params *ExtrasConfigContextSchemasListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasConfigContextSchemasListOK, error)

	ExtrasConfigContextSchemasPartialUpdate(params *ExtrasConfigContextSchemasPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasConfigContextSchemasPartialUpdateOK, error)

	ExtrasConfigContextSchemasRead(params *ExtrasConfigContextSchemasReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasConfigContextSchemasReadOK, error)

	ExtrasConfigContextSchemasUpdate(params *ExtrasConfigContextSchemasUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasConfigContextSchemasUpdateOK, error)

	ExtrasConfigContextsBulkDelete(params *ExtrasConfigContextsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasConfigContextsBulkDeleteNoContent, error)

	ExtrasConfigContextsBulkPartialUpdate(params *ExtrasConfigContextsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasConfigContextsBulkPartialUpdateOK, error)

	ExtrasConfigContextsBulkUpdate(params *ExtrasConfigContextsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasConfigContextsBulkUpdateOK, error)

	ExtrasConfigContextsCreate(params *ExtrasConfigContextsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasConfigContextsCreateCreated, error)

	ExtrasConfigContextsDelete(params *ExtrasConfigContextsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasConfigContextsDeleteNoContent, error)

	ExtrasConfigContextsList(params *ExtrasConfigContextsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasConfigContextsListOK, error)

	ExtrasConfigContextsPartialUpdate(params *ExtrasConfigContextsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasConfigContextsPartialUpdateOK, error)

	ExtrasConfigContextsRead(params *ExtrasConfigContextsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasConfigContextsReadOK, error)

	ExtrasConfigContextsUpdate(params *ExtrasConfigContextsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasConfigContextsUpdateOK, error)

	ExtrasContentTypesList(params *ExtrasContentTypesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasContentTypesListOK, error)

	ExtrasContentTypesRead(params *ExtrasContentTypesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasContentTypesReadOK, error)

	ExtrasCustomFieldChoicesBulkDelete(params *ExtrasCustomFieldChoicesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomFieldChoicesBulkDeleteNoContent, error)

	ExtrasCustomFieldChoicesBulkPartialUpdate(params *ExtrasCustomFieldChoicesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomFieldChoicesBulkPartialUpdateOK, error)

	ExtrasCustomFieldChoicesBulkUpdate(params *ExtrasCustomFieldChoicesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomFieldChoicesBulkUpdateOK, error)

	ExtrasCustomFieldChoicesCreate(params *ExtrasCustomFieldChoicesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomFieldChoicesCreateCreated, error)

	ExtrasCustomFieldChoicesDelete(params *ExtrasCustomFieldChoicesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomFieldChoicesDeleteNoContent, error)

	ExtrasCustomFieldChoicesList(params *ExtrasCustomFieldChoicesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomFieldChoicesListOK, error)

	ExtrasCustomFieldChoicesPartialUpdate(params *ExtrasCustomFieldChoicesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomFieldChoicesPartialUpdateOK, error)

	ExtrasCustomFieldChoicesRead(params *ExtrasCustomFieldChoicesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomFieldChoicesReadOK, error)

	ExtrasCustomFieldChoicesUpdate(params *ExtrasCustomFieldChoicesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomFieldChoicesUpdateOK, error)

	ExtrasCustomFieldsBulkDelete(params *ExtrasCustomFieldsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomFieldsBulkDeleteNoContent, error)

	ExtrasCustomFieldsBulkPartialUpdate(params *ExtrasCustomFieldsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomFieldsBulkPartialUpdateOK, error)

	ExtrasCustomFieldsBulkUpdate(params *ExtrasCustomFieldsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomFieldsBulkUpdateOK, error)

	ExtrasCustomFieldsCreate(params *ExtrasCustomFieldsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomFieldsCreateCreated, error)

	ExtrasCustomFieldsDelete(params *ExtrasCustomFieldsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomFieldsDeleteNoContent, error)

	ExtrasCustomFieldsList(params *ExtrasCustomFieldsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomFieldsListOK, error)

	ExtrasCustomFieldsPartialUpdate(params *ExtrasCustomFieldsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomFieldsPartialUpdateOK, error)

	ExtrasCustomFieldsRead(params *ExtrasCustomFieldsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomFieldsReadOK, error)

	ExtrasCustomFieldsUpdate(params *ExtrasCustomFieldsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomFieldsUpdateOK, error)

	ExtrasCustomLinksBulkDelete(params *ExtrasCustomLinksBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomLinksBulkDeleteNoContent, error)

	ExtrasCustomLinksBulkPartialUpdate(params *ExtrasCustomLinksBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomLinksBulkPartialUpdateOK, error)

	ExtrasCustomLinksBulkUpdate(params *ExtrasCustomLinksBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomLinksBulkUpdateOK, error)

	ExtrasCustomLinksCreate(params *ExtrasCustomLinksCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomLinksCreateCreated, error)

	ExtrasCustomLinksDelete(params *ExtrasCustomLinksDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomLinksDeleteNoContent, error)

	ExtrasCustomLinksList(params *ExtrasCustomLinksListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomLinksListOK, error)

	ExtrasCustomLinksPartialUpdate(params *ExtrasCustomLinksPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomLinksPartialUpdateOK, error)

	ExtrasCustomLinksRead(params *ExtrasCustomLinksReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomLinksReadOK, error)

	ExtrasCustomLinksUpdate(params *ExtrasCustomLinksUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomLinksUpdateOK, error)

	ExtrasExportTemplatesBulkDelete(params *ExtrasExportTemplatesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasExportTemplatesBulkDeleteNoContent, error)

	ExtrasExportTemplatesBulkPartialUpdate(params *ExtrasExportTemplatesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasExportTemplatesBulkPartialUpdateOK, error)

	ExtrasExportTemplatesBulkUpdate(params *ExtrasExportTemplatesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasExportTemplatesBulkUpdateOK, error)

	ExtrasExportTemplatesCreate(params *ExtrasExportTemplatesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasExportTemplatesCreateCreated, error)

	ExtrasExportTemplatesDelete(params *ExtrasExportTemplatesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasExportTemplatesDeleteNoContent, error)

	ExtrasExportTemplatesList(params *ExtrasExportTemplatesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasExportTemplatesListOK, error)

	ExtrasExportTemplatesPartialUpdate(params *ExtrasExportTemplatesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasExportTemplatesPartialUpdateOK, error)

	ExtrasExportTemplatesRead(params *ExtrasExportTemplatesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasExportTemplatesReadOK, error)

	ExtrasExportTemplatesUpdate(params *ExtrasExportTemplatesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasExportTemplatesUpdateOK, error)

	ExtrasGitRepositoriesBulkDelete(params *ExtrasGitRepositoriesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasGitRepositoriesBulkDeleteNoContent, error)

	ExtrasGitRepositoriesBulkPartialUpdate(params *ExtrasGitRepositoriesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasGitRepositoriesBulkPartialUpdateOK, error)

	ExtrasGitRepositoriesBulkUpdate(params *ExtrasGitRepositoriesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasGitRepositoriesBulkUpdateOK, error)

	ExtrasGitRepositoriesCreate(params *ExtrasGitRepositoriesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasGitRepositoriesCreateCreated, error)

	ExtrasGitRepositoriesDelete(params *ExtrasGitRepositoriesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasGitRepositoriesDeleteNoContent, error)

	ExtrasGitRepositoriesList(params *ExtrasGitRepositoriesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasGitRepositoriesListOK, error)

	ExtrasGitRepositoriesPartialUpdate(params *ExtrasGitRepositoriesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasGitRepositoriesPartialUpdateOK, error)

	ExtrasGitRepositoriesRead(params *ExtrasGitRepositoriesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasGitRepositoriesReadOK, error)

	ExtrasGitRepositoriesSync(params *ExtrasGitRepositoriesSyncParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasGitRepositoriesSyncCreated, error)

	ExtrasGitRepositoriesUpdate(params *ExtrasGitRepositoriesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasGitRepositoriesUpdateOK, error)

	ExtrasGraphqlQueriesBulkDelete(params *ExtrasGraphqlQueriesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasGraphqlQueriesBulkDeleteNoContent, error)

	ExtrasGraphqlQueriesBulkPartialUpdate(params *ExtrasGraphqlQueriesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasGraphqlQueriesBulkPartialUpdateOK, error)

	ExtrasGraphqlQueriesBulkUpdate(params *ExtrasGraphqlQueriesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasGraphqlQueriesBulkUpdateOK, error)

	ExtrasGraphqlQueriesCreate(params *ExtrasGraphqlQueriesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasGraphqlQueriesCreateCreated, error)

	ExtrasGraphqlQueriesDelete(params *ExtrasGraphqlQueriesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasGraphqlQueriesDeleteNoContent, error)

	ExtrasGraphqlQueriesList(params *ExtrasGraphqlQueriesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasGraphqlQueriesListOK, error)

	ExtrasGraphqlQueriesPartialUpdate(params *ExtrasGraphqlQueriesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasGraphqlQueriesPartialUpdateOK, error)

	ExtrasGraphqlQueriesRead(params *ExtrasGraphqlQueriesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasGraphqlQueriesReadOK, error)

	ExtrasGraphqlQueriesRun(params *ExtrasGraphqlQueriesRunParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasGraphqlQueriesRunCreated, error)

	ExtrasGraphqlQueriesUpdate(params *ExtrasGraphqlQueriesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasGraphqlQueriesUpdateOK, error)

	ExtrasImageAttachmentsBulkDelete(params *ExtrasImageAttachmentsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasImageAttachmentsBulkDeleteNoContent, error)

	ExtrasImageAttachmentsBulkPartialUpdate(params *ExtrasImageAttachmentsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasImageAttachmentsBulkPartialUpdateOK, error)

	ExtrasImageAttachmentsBulkUpdate(params *ExtrasImageAttachmentsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasImageAttachmentsBulkUpdateOK, error)

	ExtrasImageAttachmentsCreate(params *ExtrasImageAttachmentsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasImageAttachmentsCreateCreated, error)

	ExtrasImageAttachmentsDelete(params *ExtrasImageAttachmentsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasImageAttachmentsDeleteNoContent, error)

	ExtrasImageAttachmentsList(params *ExtrasImageAttachmentsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasImageAttachmentsListOK, error)

	ExtrasImageAttachmentsPartialUpdate(params *ExtrasImageAttachmentsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasImageAttachmentsPartialUpdateOK, error)

	ExtrasImageAttachmentsRead(params *ExtrasImageAttachmentsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasImageAttachmentsReadOK, error)

	ExtrasImageAttachmentsUpdate(params *ExtrasImageAttachmentsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasImageAttachmentsUpdateOK, error)

	ExtrasJobResultsBulkDelete(params *ExtrasJobResultsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasJobResultsBulkDeleteNoContent, error)

	ExtrasJobResultsBulkPartialUpdate(params *ExtrasJobResultsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasJobResultsBulkPartialUpdateOK, error)

	ExtrasJobResultsBulkUpdate(params *ExtrasJobResultsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasJobResultsBulkUpdateOK, error)

	ExtrasJobResultsCreate(params *ExtrasJobResultsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasJobResultsCreateCreated, error)

	ExtrasJobResultsDelete(params *ExtrasJobResultsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasJobResultsDeleteNoContent, error)

	ExtrasJobResultsList(params *ExtrasJobResultsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasJobResultsListOK, error)

	ExtrasJobResultsPartialUpdate(params *ExtrasJobResultsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasJobResultsPartialUpdateOK, error)

	ExtrasJobResultsRead(params *ExtrasJobResultsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasJobResultsReadOK, error)

	ExtrasJobResultsUpdate(params *ExtrasJobResultsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasJobResultsUpdateOK, error)

	ExtrasJobsList(params *ExtrasJobsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasJobsListOK, error)

	ExtrasJobsRead(params *ExtrasJobsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasJobsReadOK, error)

	ExtrasJobsRun(params *ExtrasJobsRunParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasJobsRunCreated, error)

	ExtrasObjectChangesList(params *ExtrasObjectChangesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasObjectChangesListOK, error)

	ExtrasObjectChangesRead(params *ExtrasObjectChangesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasObjectChangesReadOK, error)

	ExtrasRelationshipAssociationsBulkDelete(params *ExtrasRelationshipAssociationsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasRelationshipAssociationsBulkDeleteNoContent, error)

	ExtrasRelationshipAssociationsBulkPartialUpdate(params *ExtrasRelationshipAssociationsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasRelationshipAssociationsBulkPartialUpdateOK, error)

	ExtrasRelationshipAssociationsBulkUpdate(params *ExtrasRelationshipAssociationsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasRelationshipAssociationsBulkUpdateOK, error)

	ExtrasRelationshipAssociationsCreate(params *ExtrasRelationshipAssociationsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasRelationshipAssociationsCreateCreated, error)

	ExtrasRelationshipAssociationsDelete(params *ExtrasRelationshipAssociationsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasRelationshipAssociationsDeleteNoContent, error)

	ExtrasRelationshipAssociationsList(params *ExtrasRelationshipAssociationsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasRelationshipAssociationsListOK, error)

	ExtrasRelationshipAssociationsPartialUpdate(params *ExtrasRelationshipAssociationsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasRelationshipAssociationsPartialUpdateOK, error)

	ExtrasRelationshipAssociationsRead(params *ExtrasRelationshipAssociationsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasRelationshipAssociationsReadOK, error)

	ExtrasRelationshipAssociationsUpdate(params *ExtrasRelationshipAssociationsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasRelationshipAssociationsUpdateOK, error)

	ExtrasRelationshipsBulkDelete(params *ExtrasRelationshipsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasRelationshipsBulkDeleteNoContent, error)

	ExtrasRelationshipsBulkPartialUpdate(params *ExtrasRelationshipsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasRelationshipsBulkPartialUpdateOK, error)

	ExtrasRelationshipsBulkUpdate(params *ExtrasRelationshipsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasRelationshipsBulkUpdateOK, error)

	ExtrasRelationshipsCreate(params *ExtrasRelationshipsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasRelationshipsCreateCreated, error)

	ExtrasRelationshipsDelete(params *ExtrasRelationshipsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasRelationshipsDeleteNoContent, error)

	ExtrasRelationshipsList(params *ExtrasRelationshipsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasRelationshipsListOK, error)

	ExtrasRelationshipsPartialUpdate(params *ExtrasRelationshipsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasRelationshipsPartialUpdateOK, error)

	ExtrasRelationshipsRead(params *ExtrasRelationshipsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasRelationshipsReadOK, error)

	ExtrasRelationshipsUpdate(params *ExtrasRelationshipsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasRelationshipsUpdateOK, error)

	ExtrasStatusesBulkDelete(params *ExtrasStatusesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasStatusesBulkDeleteNoContent, error)

	ExtrasStatusesBulkPartialUpdate(params *ExtrasStatusesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasStatusesBulkPartialUpdateOK, error)

	ExtrasStatusesBulkUpdate(params *ExtrasStatusesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasStatusesBulkUpdateOK, error)

	ExtrasStatusesCreate(params *ExtrasStatusesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasStatusesCreateCreated, error)

	ExtrasStatusesDelete(params *ExtrasStatusesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasStatusesDeleteNoContent, error)

	ExtrasStatusesList(params *ExtrasStatusesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasStatusesListOK, error)

	ExtrasStatusesPartialUpdate(params *ExtrasStatusesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasStatusesPartialUpdateOK, error)

	ExtrasStatusesRead(params *ExtrasStatusesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasStatusesReadOK, error)

	ExtrasStatusesUpdate(params *ExtrasStatusesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasStatusesUpdateOK, error)

	ExtrasTagsBulkDelete(params *ExtrasTagsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasTagsBulkDeleteNoContent, error)

	ExtrasTagsBulkPartialUpdate(params *ExtrasTagsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasTagsBulkPartialUpdateOK, error)

	ExtrasTagsBulkUpdate(params *ExtrasTagsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasTagsBulkUpdateOK, error)

	ExtrasTagsCreate(params *ExtrasTagsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasTagsCreateCreated, error)

	ExtrasTagsDelete(params *ExtrasTagsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasTagsDeleteNoContent, error)

	ExtrasTagsList(params *ExtrasTagsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasTagsListOK, error)

	ExtrasTagsPartialUpdate(params *ExtrasTagsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasTagsPartialUpdateOK, error)

	ExtrasTagsRead(params *ExtrasTagsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasTagsReadOK, error)

	ExtrasTagsUpdate(params *ExtrasTagsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasTagsUpdateOK, error)

	ExtrasWebhooksBulkDelete(params *ExtrasWebhooksBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasWebhooksBulkDeleteNoContent, error)

	ExtrasWebhooksBulkPartialUpdate(params *ExtrasWebhooksBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasWebhooksBulkPartialUpdateOK, error)

	ExtrasWebhooksBulkUpdate(params *ExtrasWebhooksBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasWebhooksBulkUpdateOK, error)

	ExtrasWebhooksCreate(params *ExtrasWebhooksCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasWebhooksCreateCreated, error)

	ExtrasWebhooksDelete(params *ExtrasWebhooksDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasWebhooksDeleteNoContent, error)

	ExtrasWebhooksList(params *ExtrasWebhooksListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasWebhooksListOK, error)

	ExtrasWebhooksPartialUpdate(params *ExtrasWebhooksPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasWebhooksPartialUpdateOK, error)

	ExtrasWebhooksRead(params *ExtrasWebhooksReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasWebhooksReadOK, error)

	ExtrasWebhooksUpdate(params *ExtrasWebhooksUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasWebhooksUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ExtrasComputedFieldsBulkDelete Manage Computed Fields through DELETE, GET, POST, PUT, and PATCH requests.
*/
func (a *Client) ExtrasComputedFieldsBulkDelete(params *ExtrasComputedFieldsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasComputedFieldsBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasComputedFieldsBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_computed-fields_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/extras/computed-fields/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasComputedFieldsBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*ExtrasComputedFieldsBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_computed-fields_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasComputedFieldsBulkPartialUpdate Manage Computed Fields through DELETE, GET, POST, PUT, and PATCH requests.
*/
func (a *Client) ExtrasComputedFieldsBulkPartialUpdate(params *ExtrasComputedFieldsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasComputedFieldsBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasComputedFieldsBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_computed-fields_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/extras/computed-fields/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasComputedFieldsBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasComputedFieldsBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_computed-fields_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasComputedFieldsBulkUpdate Manage Computed Fields through DELETE, GET, POST, PUT, and PATCH requests.
*/
func (a *Client) ExtrasComputedFieldsBulkUpdate(params *ExtrasComputedFieldsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasComputedFieldsBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasComputedFieldsBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_computed-fields_bulk_update",
		Method:             "PUT",
		PathPattern:        "/extras/computed-fields/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasComputedFieldsBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasComputedFieldsBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_computed-fields_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasComputedFieldsCreate Manage Computed Fields through DELETE, GET, POST, PUT, and PATCH requests.
*/
func (a *Client) ExtrasComputedFieldsCreate(params *ExtrasComputedFieldsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasComputedFieldsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasComputedFieldsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_computed-fields_create",
		Method:             "POST",
		PathPattern:        "/extras/computed-fields/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasComputedFieldsCreateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasComputedFieldsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_computed-fields_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasComputedFieldsDelete Manage Computed Fields through DELETE, GET, POST, PUT, and PATCH requests.
*/
func (a *Client) ExtrasComputedFieldsDelete(params *ExtrasComputedFieldsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasComputedFieldsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasComputedFieldsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_computed-fields_delete",
		Method:             "DELETE",
		PathPattern:        "/extras/computed-fields/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasComputedFieldsDeleteReader{formats: a.formats},
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
	success, ok := result.(*ExtrasComputedFieldsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_computed-fields_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasComputedFieldsList Manage Computed Fields through DELETE, GET, POST, PUT, and PATCH requests.
*/
func (a *Client) ExtrasComputedFieldsList(params *ExtrasComputedFieldsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasComputedFieldsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasComputedFieldsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_computed-fields_list",
		Method:             "GET",
		PathPattern:        "/extras/computed-fields/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasComputedFieldsListReader{formats: a.formats},
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
	success, ok := result.(*ExtrasComputedFieldsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_computed-fields_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasComputedFieldsPartialUpdate Manage Computed Fields through DELETE, GET, POST, PUT, and PATCH requests.
*/
func (a *Client) ExtrasComputedFieldsPartialUpdate(params *ExtrasComputedFieldsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasComputedFieldsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasComputedFieldsPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_computed-fields_partial_update",
		Method:             "PATCH",
		PathPattern:        "/extras/computed-fields/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasComputedFieldsPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasComputedFieldsPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_computed-fields_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasComputedFieldsRead Manage Computed Fields through DELETE, GET, POST, PUT, and PATCH requests.
*/
func (a *Client) ExtrasComputedFieldsRead(params *ExtrasComputedFieldsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasComputedFieldsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasComputedFieldsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_computed-fields_read",
		Method:             "GET",
		PathPattern:        "/extras/computed-fields/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasComputedFieldsReadReader{formats: a.formats},
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
	success, ok := result.(*ExtrasComputedFieldsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_computed-fields_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasComputedFieldsUpdate Manage Computed Fields through DELETE, GET, POST, PUT, and PATCH requests.
*/
func (a *Client) ExtrasComputedFieldsUpdate(params *ExtrasComputedFieldsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasComputedFieldsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasComputedFieldsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_computed-fields_update",
		Method:             "PUT",
		PathPattern:        "/extras/computed-fields/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasComputedFieldsUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasComputedFieldsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_computed-fields_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasConfigContextSchemasBulkDelete extras config context schemas bulk delete API
*/
func (a *Client) ExtrasConfigContextSchemasBulkDelete(params *ExtrasConfigContextSchemasBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasConfigContextSchemasBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasConfigContextSchemasBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_config-context-schemas_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/extras/config-context-schemas/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasConfigContextSchemasBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*ExtrasConfigContextSchemasBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_config-context-schemas_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasConfigContextSchemasBulkPartialUpdate extras config context schemas bulk partial update API
*/
func (a *Client) ExtrasConfigContextSchemasBulkPartialUpdate(params *ExtrasConfigContextSchemasBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasConfigContextSchemasBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasConfigContextSchemasBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_config-context-schemas_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/extras/config-context-schemas/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasConfigContextSchemasBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasConfigContextSchemasBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_config-context-schemas_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasConfigContextSchemasBulkUpdate extras config context schemas bulk update API
*/
func (a *Client) ExtrasConfigContextSchemasBulkUpdate(params *ExtrasConfigContextSchemasBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasConfigContextSchemasBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasConfigContextSchemasBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_config-context-schemas_bulk_update",
		Method:             "PUT",
		PathPattern:        "/extras/config-context-schemas/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasConfigContextSchemasBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasConfigContextSchemasBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_config-context-schemas_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasConfigContextSchemasCreate extras config context schemas create API
*/
func (a *Client) ExtrasConfigContextSchemasCreate(params *ExtrasConfigContextSchemasCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasConfigContextSchemasCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasConfigContextSchemasCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_config-context-schemas_create",
		Method:             "POST",
		PathPattern:        "/extras/config-context-schemas/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasConfigContextSchemasCreateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasConfigContextSchemasCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_config-context-schemas_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasConfigContextSchemasDelete extras config context schemas delete API
*/
func (a *Client) ExtrasConfigContextSchemasDelete(params *ExtrasConfigContextSchemasDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasConfigContextSchemasDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasConfigContextSchemasDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_config-context-schemas_delete",
		Method:             "DELETE",
		PathPattern:        "/extras/config-context-schemas/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasConfigContextSchemasDeleteReader{formats: a.formats},
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
	success, ok := result.(*ExtrasConfigContextSchemasDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_config-context-schemas_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasConfigContextSchemasList extras config context schemas list API
*/
func (a *Client) ExtrasConfigContextSchemasList(params *ExtrasConfigContextSchemasListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasConfigContextSchemasListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasConfigContextSchemasListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_config-context-schemas_list",
		Method:             "GET",
		PathPattern:        "/extras/config-context-schemas/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasConfigContextSchemasListReader{formats: a.formats},
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
	success, ok := result.(*ExtrasConfigContextSchemasListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_config-context-schemas_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasConfigContextSchemasPartialUpdate extras config context schemas partial update API
*/
func (a *Client) ExtrasConfigContextSchemasPartialUpdate(params *ExtrasConfigContextSchemasPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasConfigContextSchemasPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasConfigContextSchemasPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_config-context-schemas_partial_update",
		Method:             "PATCH",
		PathPattern:        "/extras/config-context-schemas/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasConfigContextSchemasPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasConfigContextSchemasPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_config-context-schemas_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasConfigContextSchemasRead extras config context schemas read API
*/
func (a *Client) ExtrasConfigContextSchemasRead(params *ExtrasConfigContextSchemasReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasConfigContextSchemasReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasConfigContextSchemasReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_config-context-schemas_read",
		Method:             "GET",
		PathPattern:        "/extras/config-context-schemas/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasConfigContextSchemasReadReader{formats: a.formats},
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
	success, ok := result.(*ExtrasConfigContextSchemasReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_config-context-schemas_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasConfigContextSchemasUpdate extras config context schemas update API
*/
func (a *Client) ExtrasConfigContextSchemasUpdate(params *ExtrasConfigContextSchemasUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasConfigContextSchemasUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasConfigContextSchemasUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_config-context-schemas_update",
		Method:             "PUT",
		PathPattern:        "/extras/config-context-schemas/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasConfigContextSchemasUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasConfigContextSchemasUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_config-context-schemas_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasConfigContextsBulkDelete extras config contexts bulk delete API
*/
func (a *Client) ExtrasConfigContextsBulkDelete(params *ExtrasConfigContextsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasConfigContextsBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasConfigContextsBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_config-contexts_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/extras/config-contexts/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasConfigContextsBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*ExtrasConfigContextsBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_config-contexts_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasConfigContextsBulkPartialUpdate extras config contexts bulk partial update API
*/
func (a *Client) ExtrasConfigContextsBulkPartialUpdate(params *ExtrasConfigContextsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasConfigContextsBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasConfigContextsBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_config-contexts_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/extras/config-contexts/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasConfigContextsBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasConfigContextsBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_config-contexts_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasConfigContextsBulkUpdate extras config contexts bulk update API
*/
func (a *Client) ExtrasConfigContextsBulkUpdate(params *ExtrasConfigContextsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasConfigContextsBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasConfigContextsBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_config-contexts_bulk_update",
		Method:             "PUT",
		PathPattern:        "/extras/config-contexts/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasConfigContextsBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasConfigContextsBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_config-contexts_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasConfigContextsCreate extras config contexts create API
*/
func (a *Client) ExtrasConfigContextsCreate(params *ExtrasConfigContextsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasConfigContextsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasConfigContextsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_config-contexts_create",
		Method:             "POST",
		PathPattern:        "/extras/config-contexts/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasConfigContextsCreateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasConfigContextsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_config-contexts_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasConfigContextsDelete extras config contexts delete API
*/
func (a *Client) ExtrasConfigContextsDelete(params *ExtrasConfigContextsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasConfigContextsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasConfigContextsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_config-contexts_delete",
		Method:             "DELETE",
		PathPattern:        "/extras/config-contexts/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasConfigContextsDeleteReader{formats: a.formats},
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
	success, ok := result.(*ExtrasConfigContextsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_config-contexts_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasConfigContextsList extras config contexts list API
*/
func (a *Client) ExtrasConfigContextsList(params *ExtrasConfigContextsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasConfigContextsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasConfigContextsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_config-contexts_list",
		Method:             "GET",
		PathPattern:        "/extras/config-contexts/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasConfigContextsListReader{formats: a.formats},
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
	success, ok := result.(*ExtrasConfigContextsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_config-contexts_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasConfigContextsPartialUpdate extras config contexts partial update API
*/
func (a *Client) ExtrasConfigContextsPartialUpdate(params *ExtrasConfigContextsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasConfigContextsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasConfigContextsPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_config-contexts_partial_update",
		Method:             "PATCH",
		PathPattern:        "/extras/config-contexts/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasConfigContextsPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasConfigContextsPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_config-contexts_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasConfigContextsRead extras config contexts read API
*/
func (a *Client) ExtrasConfigContextsRead(params *ExtrasConfigContextsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasConfigContextsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasConfigContextsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_config-contexts_read",
		Method:             "GET",
		PathPattern:        "/extras/config-contexts/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasConfigContextsReadReader{formats: a.formats},
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
	success, ok := result.(*ExtrasConfigContextsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_config-contexts_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasConfigContextsUpdate extras config contexts update API
*/
func (a *Client) ExtrasConfigContextsUpdate(params *ExtrasConfigContextsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasConfigContextsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasConfigContextsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_config-contexts_update",
		Method:             "PUT",
		PathPattern:        "/extras/config-contexts/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasConfigContextsUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasConfigContextsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_config-contexts_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasContentTypesList Read-only list of ContentTypes. Limit results to ContentTypes pertinent to Nautobot objects.
*/
func (a *Client) ExtrasContentTypesList(params *ExtrasContentTypesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasContentTypesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasContentTypesListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_content-types_list",
		Method:             "GET",
		PathPattern:        "/extras/content-types/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasContentTypesListReader{formats: a.formats},
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
	success, ok := result.(*ExtrasContentTypesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_content-types_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasContentTypesRead Read-only list of ContentTypes. Limit results to ContentTypes pertinent to Nautobot objects.
*/
func (a *Client) ExtrasContentTypesRead(params *ExtrasContentTypesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasContentTypesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasContentTypesReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_content-types_read",
		Method:             "GET",
		PathPattern:        "/extras/content-types/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasContentTypesReadReader{formats: a.formats},
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
	success, ok := result.(*ExtrasContentTypesReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_content-types_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasCustomFieldChoicesBulkDelete extras custom field choices bulk delete API
*/
func (a *Client) ExtrasCustomFieldChoicesBulkDelete(params *ExtrasCustomFieldChoicesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomFieldChoicesBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasCustomFieldChoicesBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_custom-field-choices_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/extras/custom-field-choices/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasCustomFieldChoicesBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*ExtrasCustomFieldChoicesBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_custom-field-choices_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasCustomFieldChoicesBulkPartialUpdate extras custom field choices bulk partial update API
*/
func (a *Client) ExtrasCustomFieldChoicesBulkPartialUpdate(params *ExtrasCustomFieldChoicesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomFieldChoicesBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasCustomFieldChoicesBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_custom-field-choices_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/extras/custom-field-choices/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasCustomFieldChoicesBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasCustomFieldChoicesBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_custom-field-choices_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasCustomFieldChoicesBulkUpdate extras custom field choices bulk update API
*/
func (a *Client) ExtrasCustomFieldChoicesBulkUpdate(params *ExtrasCustomFieldChoicesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomFieldChoicesBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasCustomFieldChoicesBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_custom-field-choices_bulk_update",
		Method:             "PUT",
		PathPattern:        "/extras/custom-field-choices/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasCustomFieldChoicesBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasCustomFieldChoicesBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_custom-field-choices_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasCustomFieldChoicesCreate extras custom field choices create API
*/
func (a *Client) ExtrasCustomFieldChoicesCreate(params *ExtrasCustomFieldChoicesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomFieldChoicesCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasCustomFieldChoicesCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_custom-field-choices_create",
		Method:             "POST",
		PathPattern:        "/extras/custom-field-choices/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasCustomFieldChoicesCreateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasCustomFieldChoicesCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_custom-field-choices_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasCustomFieldChoicesDelete extras custom field choices delete API
*/
func (a *Client) ExtrasCustomFieldChoicesDelete(params *ExtrasCustomFieldChoicesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomFieldChoicesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasCustomFieldChoicesDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_custom-field-choices_delete",
		Method:             "DELETE",
		PathPattern:        "/extras/custom-field-choices/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasCustomFieldChoicesDeleteReader{formats: a.formats},
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
	success, ok := result.(*ExtrasCustomFieldChoicesDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_custom-field-choices_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasCustomFieldChoicesList extras custom field choices list API
*/
func (a *Client) ExtrasCustomFieldChoicesList(params *ExtrasCustomFieldChoicesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomFieldChoicesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasCustomFieldChoicesListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_custom-field-choices_list",
		Method:             "GET",
		PathPattern:        "/extras/custom-field-choices/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasCustomFieldChoicesListReader{formats: a.formats},
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
	success, ok := result.(*ExtrasCustomFieldChoicesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_custom-field-choices_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasCustomFieldChoicesPartialUpdate extras custom field choices partial update API
*/
func (a *Client) ExtrasCustomFieldChoicesPartialUpdate(params *ExtrasCustomFieldChoicesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomFieldChoicesPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasCustomFieldChoicesPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_custom-field-choices_partial_update",
		Method:             "PATCH",
		PathPattern:        "/extras/custom-field-choices/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasCustomFieldChoicesPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasCustomFieldChoicesPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_custom-field-choices_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasCustomFieldChoicesRead extras custom field choices read API
*/
func (a *Client) ExtrasCustomFieldChoicesRead(params *ExtrasCustomFieldChoicesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomFieldChoicesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasCustomFieldChoicesReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_custom-field-choices_read",
		Method:             "GET",
		PathPattern:        "/extras/custom-field-choices/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasCustomFieldChoicesReadReader{formats: a.formats},
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
	success, ok := result.(*ExtrasCustomFieldChoicesReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_custom-field-choices_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasCustomFieldChoicesUpdate extras custom field choices update API
*/
func (a *Client) ExtrasCustomFieldChoicesUpdate(params *ExtrasCustomFieldChoicesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomFieldChoicesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasCustomFieldChoicesUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_custom-field-choices_update",
		Method:             "PUT",
		PathPattern:        "/extras/custom-field-choices/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasCustomFieldChoicesUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasCustomFieldChoicesUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_custom-field-choices_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasCustomFieldsBulkDelete extras custom fields bulk delete API
*/
func (a *Client) ExtrasCustomFieldsBulkDelete(params *ExtrasCustomFieldsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomFieldsBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasCustomFieldsBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_custom-fields_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/extras/custom-fields/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasCustomFieldsBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*ExtrasCustomFieldsBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_custom-fields_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasCustomFieldsBulkPartialUpdate extras custom fields bulk partial update API
*/
func (a *Client) ExtrasCustomFieldsBulkPartialUpdate(params *ExtrasCustomFieldsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomFieldsBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasCustomFieldsBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_custom-fields_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/extras/custom-fields/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasCustomFieldsBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasCustomFieldsBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_custom-fields_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasCustomFieldsBulkUpdate extras custom fields bulk update API
*/
func (a *Client) ExtrasCustomFieldsBulkUpdate(params *ExtrasCustomFieldsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomFieldsBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasCustomFieldsBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_custom-fields_bulk_update",
		Method:             "PUT",
		PathPattern:        "/extras/custom-fields/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasCustomFieldsBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasCustomFieldsBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_custom-fields_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasCustomFieldsCreate extras custom fields create API
*/
func (a *Client) ExtrasCustomFieldsCreate(params *ExtrasCustomFieldsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomFieldsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasCustomFieldsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_custom-fields_create",
		Method:             "POST",
		PathPattern:        "/extras/custom-fields/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasCustomFieldsCreateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasCustomFieldsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_custom-fields_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasCustomFieldsDelete extras custom fields delete API
*/
func (a *Client) ExtrasCustomFieldsDelete(params *ExtrasCustomFieldsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomFieldsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasCustomFieldsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_custom-fields_delete",
		Method:             "DELETE",
		PathPattern:        "/extras/custom-fields/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasCustomFieldsDeleteReader{formats: a.formats},
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
	success, ok := result.(*ExtrasCustomFieldsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_custom-fields_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasCustomFieldsList extras custom fields list API
*/
func (a *Client) ExtrasCustomFieldsList(params *ExtrasCustomFieldsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomFieldsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasCustomFieldsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_custom-fields_list",
		Method:             "GET",
		PathPattern:        "/extras/custom-fields/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasCustomFieldsListReader{formats: a.formats},
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
	success, ok := result.(*ExtrasCustomFieldsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_custom-fields_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasCustomFieldsPartialUpdate extras custom fields partial update API
*/
func (a *Client) ExtrasCustomFieldsPartialUpdate(params *ExtrasCustomFieldsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomFieldsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasCustomFieldsPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_custom-fields_partial_update",
		Method:             "PATCH",
		PathPattern:        "/extras/custom-fields/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasCustomFieldsPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasCustomFieldsPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_custom-fields_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasCustomFieldsRead extras custom fields read API
*/
func (a *Client) ExtrasCustomFieldsRead(params *ExtrasCustomFieldsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomFieldsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasCustomFieldsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_custom-fields_read",
		Method:             "GET",
		PathPattern:        "/extras/custom-fields/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasCustomFieldsReadReader{formats: a.formats},
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
	success, ok := result.(*ExtrasCustomFieldsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_custom-fields_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasCustomFieldsUpdate extras custom fields update API
*/
func (a *Client) ExtrasCustomFieldsUpdate(params *ExtrasCustomFieldsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomFieldsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasCustomFieldsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_custom-fields_update",
		Method:             "PUT",
		PathPattern:        "/extras/custom-fields/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasCustomFieldsUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasCustomFieldsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_custom-fields_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasCustomLinksBulkDelete Manage Custom Links through DELETE, GET, POST, PUT, and PATCH requests.
*/
func (a *Client) ExtrasCustomLinksBulkDelete(params *ExtrasCustomLinksBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomLinksBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasCustomLinksBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_custom-links_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/extras/custom-links/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasCustomLinksBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*ExtrasCustomLinksBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_custom-links_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasCustomLinksBulkPartialUpdate Manage Custom Links through DELETE, GET, POST, PUT, and PATCH requests.
*/
func (a *Client) ExtrasCustomLinksBulkPartialUpdate(params *ExtrasCustomLinksBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomLinksBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasCustomLinksBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_custom-links_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/extras/custom-links/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasCustomLinksBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasCustomLinksBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_custom-links_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasCustomLinksBulkUpdate Manage Custom Links through DELETE, GET, POST, PUT, and PATCH requests.
*/
func (a *Client) ExtrasCustomLinksBulkUpdate(params *ExtrasCustomLinksBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomLinksBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasCustomLinksBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_custom-links_bulk_update",
		Method:             "PUT",
		PathPattern:        "/extras/custom-links/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasCustomLinksBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasCustomLinksBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_custom-links_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasCustomLinksCreate Manage Custom Links through DELETE, GET, POST, PUT, and PATCH requests.
*/
func (a *Client) ExtrasCustomLinksCreate(params *ExtrasCustomLinksCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomLinksCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasCustomLinksCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_custom-links_create",
		Method:             "POST",
		PathPattern:        "/extras/custom-links/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasCustomLinksCreateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasCustomLinksCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_custom-links_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasCustomLinksDelete Manage Custom Links through DELETE, GET, POST, PUT, and PATCH requests.
*/
func (a *Client) ExtrasCustomLinksDelete(params *ExtrasCustomLinksDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomLinksDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasCustomLinksDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_custom-links_delete",
		Method:             "DELETE",
		PathPattern:        "/extras/custom-links/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasCustomLinksDeleteReader{formats: a.formats},
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
	success, ok := result.(*ExtrasCustomLinksDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_custom-links_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasCustomLinksList Manage Custom Links through DELETE, GET, POST, PUT, and PATCH requests.
*/
func (a *Client) ExtrasCustomLinksList(params *ExtrasCustomLinksListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomLinksListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasCustomLinksListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_custom-links_list",
		Method:             "GET",
		PathPattern:        "/extras/custom-links/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasCustomLinksListReader{formats: a.formats},
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
	success, ok := result.(*ExtrasCustomLinksListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_custom-links_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasCustomLinksPartialUpdate Manage Custom Links through DELETE, GET, POST, PUT, and PATCH requests.
*/
func (a *Client) ExtrasCustomLinksPartialUpdate(params *ExtrasCustomLinksPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomLinksPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasCustomLinksPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_custom-links_partial_update",
		Method:             "PATCH",
		PathPattern:        "/extras/custom-links/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasCustomLinksPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasCustomLinksPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_custom-links_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasCustomLinksRead Manage Custom Links through DELETE, GET, POST, PUT, and PATCH requests.
*/
func (a *Client) ExtrasCustomLinksRead(params *ExtrasCustomLinksReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomLinksReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasCustomLinksReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_custom-links_read",
		Method:             "GET",
		PathPattern:        "/extras/custom-links/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasCustomLinksReadReader{formats: a.formats},
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
	success, ok := result.(*ExtrasCustomLinksReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_custom-links_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasCustomLinksUpdate Manage Custom Links through DELETE, GET, POST, PUT, and PATCH requests.
*/
func (a *Client) ExtrasCustomLinksUpdate(params *ExtrasCustomLinksUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasCustomLinksUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasCustomLinksUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_custom-links_update",
		Method:             "PUT",
		PathPattern:        "/extras/custom-links/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasCustomLinksUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasCustomLinksUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_custom-links_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasExportTemplatesBulkDelete extras export templates bulk delete API
*/
func (a *Client) ExtrasExportTemplatesBulkDelete(params *ExtrasExportTemplatesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasExportTemplatesBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasExportTemplatesBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_export-templates_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/extras/export-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasExportTemplatesBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*ExtrasExportTemplatesBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_export-templates_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasExportTemplatesBulkPartialUpdate extras export templates bulk partial update API
*/
func (a *Client) ExtrasExportTemplatesBulkPartialUpdate(params *ExtrasExportTemplatesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasExportTemplatesBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasExportTemplatesBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_export-templates_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/extras/export-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasExportTemplatesBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasExportTemplatesBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_export-templates_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasExportTemplatesBulkUpdate extras export templates bulk update API
*/
func (a *Client) ExtrasExportTemplatesBulkUpdate(params *ExtrasExportTemplatesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasExportTemplatesBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasExportTemplatesBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_export-templates_bulk_update",
		Method:             "PUT",
		PathPattern:        "/extras/export-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasExportTemplatesBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasExportTemplatesBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_export-templates_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasExportTemplatesCreate extras export templates create API
*/
func (a *Client) ExtrasExportTemplatesCreate(params *ExtrasExportTemplatesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasExportTemplatesCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasExportTemplatesCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_export-templates_create",
		Method:             "POST",
		PathPattern:        "/extras/export-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasExportTemplatesCreateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasExportTemplatesCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_export-templates_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasExportTemplatesDelete extras export templates delete API
*/
func (a *Client) ExtrasExportTemplatesDelete(params *ExtrasExportTemplatesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasExportTemplatesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasExportTemplatesDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_export-templates_delete",
		Method:             "DELETE",
		PathPattern:        "/extras/export-templates/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasExportTemplatesDeleteReader{formats: a.formats},
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
	success, ok := result.(*ExtrasExportTemplatesDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_export-templates_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasExportTemplatesList extras export templates list API
*/
func (a *Client) ExtrasExportTemplatesList(params *ExtrasExportTemplatesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasExportTemplatesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasExportTemplatesListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_export-templates_list",
		Method:             "GET",
		PathPattern:        "/extras/export-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasExportTemplatesListReader{formats: a.formats},
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
	success, ok := result.(*ExtrasExportTemplatesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_export-templates_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasExportTemplatesPartialUpdate extras export templates partial update API
*/
func (a *Client) ExtrasExportTemplatesPartialUpdate(params *ExtrasExportTemplatesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasExportTemplatesPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasExportTemplatesPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_export-templates_partial_update",
		Method:             "PATCH",
		PathPattern:        "/extras/export-templates/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasExportTemplatesPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasExportTemplatesPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_export-templates_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasExportTemplatesRead extras export templates read API
*/
func (a *Client) ExtrasExportTemplatesRead(params *ExtrasExportTemplatesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasExportTemplatesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasExportTemplatesReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_export-templates_read",
		Method:             "GET",
		PathPattern:        "/extras/export-templates/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasExportTemplatesReadReader{formats: a.formats},
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
	success, ok := result.(*ExtrasExportTemplatesReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_export-templates_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasExportTemplatesUpdate extras export templates update API
*/
func (a *Client) ExtrasExportTemplatesUpdate(params *ExtrasExportTemplatesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasExportTemplatesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasExportTemplatesUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_export-templates_update",
		Method:             "PUT",
		PathPattern:        "/extras/export-templates/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasExportTemplatesUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasExportTemplatesUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_export-templates_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasGitRepositoriesBulkDelete Manage the use of Git repositories as external data sources.
*/
func (a *Client) ExtrasGitRepositoriesBulkDelete(params *ExtrasGitRepositoriesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasGitRepositoriesBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasGitRepositoriesBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_git-repositories_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/extras/git-repositories/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasGitRepositoriesBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*ExtrasGitRepositoriesBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_git-repositories_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasGitRepositoriesBulkPartialUpdate Manage the use of Git repositories as external data sources.
*/
func (a *Client) ExtrasGitRepositoriesBulkPartialUpdate(params *ExtrasGitRepositoriesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasGitRepositoriesBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasGitRepositoriesBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_git-repositories_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/extras/git-repositories/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasGitRepositoriesBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasGitRepositoriesBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_git-repositories_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasGitRepositoriesBulkUpdate Manage the use of Git repositories as external data sources.
*/
func (a *Client) ExtrasGitRepositoriesBulkUpdate(params *ExtrasGitRepositoriesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasGitRepositoriesBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasGitRepositoriesBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_git-repositories_bulk_update",
		Method:             "PUT",
		PathPattern:        "/extras/git-repositories/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasGitRepositoriesBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasGitRepositoriesBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_git-repositories_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasGitRepositoriesCreate Manage the use of Git repositories as external data sources.
*/
func (a *Client) ExtrasGitRepositoriesCreate(params *ExtrasGitRepositoriesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasGitRepositoriesCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasGitRepositoriesCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_git-repositories_create",
		Method:             "POST",
		PathPattern:        "/extras/git-repositories/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasGitRepositoriesCreateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasGitRepositoriesCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_git-repositories_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasGitRepositoriesDelete Manage the use of Git repositories as external data sources.
*/
func (a *Client) ExtrasGitRepositoriesDelete(params *ExtrasGitRepositoriesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasGitRepositoriesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasGitRepositoriesDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_git-repositories_delete",
		Method:             "DELETE",
		PathPattern:        "/extras/git-repositories/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasGitRepositoriesDeleteReader{formats: a.formats},
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
	success, ok := result.(*ExtrasGitRepositoriesDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_git-repositories_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasGitRepositoriesList Manage the use of Git repositories as external data sources.
*/
func (a *Client) ExtrasGitRepositoriesList(params *ExtrasGitRepositoriesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasGitRepositoriesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasGitRepositoriesListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_git-repositories_list",
		Method:             "GET",
		PathPattern:        "/extras/git-repositories/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasGitRepositoriesListReader{formats: a.formats},
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
	success, ok := result.(*ExtrasGitRepositoriesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_git-repositories_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasGitRepositoriesPartialUpdate Manage the use of Git repositories as external data sources.
*/
func (a *Client) ExtrasGitRepositoriesPartialUpdate(params *ExtrasGitRepositoriesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasGitRepositoriesPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasGitRepositoriesPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_git-repositories_partial_update",
		Method:             "PATCH",
		PathPattern:        "/extras/git-repositories/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasGitRepositoriesPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasGitRepositoriesPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_git-repositories_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasGitRepositoriesRead Manage the use of Git repositories as external data sources.
*/
func (a *Client) ExtrasGitRepositoriesRead(params *ExtrasGitRepositoriesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasGitRepositoriesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasGitRepositoriesReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_git-repositories_read",
		Method:             "GET",
		PathPattern:        "/extras/git-repositories/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasGitRepositoriesReadReader{formats: a.formats},
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
	success, ok := result.(*ExtrasGitRepositoriesReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_git-repositories_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasGitRepositoriesSync Enqueue pull git repository and refresh data.
*/
func (a *Client) ExtrasGitRepositoriesSync(params *ExtrasGitRepositoriesSyncParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasGitRepositoriesSyncCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasGitRepositoriesSyncParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_git-repositories_sync",
		Method:             "POST",
		PathPattern:        "/extras/git-repositories/{id}/sync/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasGitRepositoriesSyncReader{formats: a.formats},
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
	success, ok := result.(*ExtrasGitRepositoriesSyncCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_git-repositories_sync: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasGitRepositoriesUpdate Manage the use of Git repositories as external data sources.
*/
func (a *Client) ExtrasGitRepositoriesUpdate(params *ExtrasGitRepositoriesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasGitRepositoriesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasGitRepositoriesUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_git-repositories_update",
		Method:             "PUT",
		PathPattern:        "/extras/git-repositories/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasGitRepositoriesUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasGitRepositoriesUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_git-repositories_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasGraphqlQueriesBulkDelete extras graphql queries bulk delete API
*/
func (a *Client) ExtrasGraphqlQueriesBulkDelete(params *ExtrasGraphqlQueriesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasGraphqlQueriesBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasGraphqlQueriesBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_graphql-queries_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/extras/graphql-queries/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasGraphqlQueriesBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*ExtrasGraphqlQueriesBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_graphql-queries_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasGraphqlQueriesBulkPartialUpdate extras graphql queries bulk partial update API
*/
func (a *Client) ExtrasGraphqlQueriesBulkPartialUpdate(params *ExtrasGraphqlQueriesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasGraphqlQueriesBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasGraphqlQueriesBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_graphql-queries_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/extras/graphql-queries/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasGraphqlQueriesBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasGraphqlQueriesBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_graphql-queries_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasGraphqlQueriesBulkUpdate extras graphql queries bulk update API
*/
func (a *Client) ExtrasGraphqlQueriesBulkUpdate(params *ExtrasGraphqlQueriesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasGraphqlQueriesBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasGraphqlQueriesBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_graphql-queries_bulk_update",
		Method:             "PUT",
		PathPattern:        "/extras/graphql-queries/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasGraphqlQueriesBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasGraphqlQueriesBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_graphql-queries_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasGraphqlQueriesCreate extras graphql queries create API
*/
func (a *Client) ExtrasGraphqlQueriesCreate(params *ExtrasGraphqlQueriesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasGraphqlQueriesCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasGraphqlQueriesCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_graphql-queries_create",
		Method:             "POST",
		PathPattern:        "/extras/graphql-queries/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasGraphqlQueriesCreateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasGraphqlQueriesCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_graphql-queries_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasGraphqlQueriesDelete extras graphql queries delete API
*/
func (a *Client) ExtrasGraphqlQueriesDelete(params *ExtrasGraphqlQueriesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasGraphqlQueriesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasGraphqlQueriesDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_graphql-queries_delete",
		Method:             "DELETE",
		PathPattern:        "/extras/graphql-queries/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasGraphqlQueriesDeleteReader{formats: a.formats},
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
	success, ok := result.(*ExtrasGraphqlQueriesDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_graphql-queries_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasGraphqlQueriesList extras graphql queries list API
*/
func (a *Client) ExtrasGraphqlQueriesList(params *ExtrasGraphqlQueriesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasGraphqlQueriesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasGraphqlQueriesListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_graphql-queries_list",
		Method:             "GET",
		PathPattern:        "/extras/graphql-queries/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasGraphqlQueriesListReader{formats: a.formats},
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
	success, ok := result.(*ExtrasGraphqlQueriesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_graphql-queries_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasGraphqlQueriesPartialUpdate extras graphql queries partial update API
*/
func (a *Client) ExtrasGraphqlQueriesPartialUpdate(params *ExtrasGraphqlQueriesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasGraphqlQueriesPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasGraphqlQueriesPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_graphql-queries_partial_update",
		Method:             "PATCH",
		PathPattern:        "/extras/graphql-queries/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasGraphqlQueriesPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasGraphqlQueriesPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_graphql-queries_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasGraphqlQueriesRead extras graphql queries read API
*/
func (a *Client) ExtrasGraphqlQueriesRead(params *ExtrasGraphqlQueriesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasGraphqlQueriesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasGraphqlQueriesReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_graphql-queries_read",
		Method:             "GET",
		PathPattern:        "/extras/graphql-queries/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasGraphqlQueriesReadReader{formats: a.formats},
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
	success, ok := result.(*ExtrasGraphqlQueriesReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_graphql-queries_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasGraphqlQueriesRun extras graphql queries run API
*/
func (a *Client) ExtrasGraphqlQueriesRun(params *ExtrasGraphqlQueriesRunParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasGraphqlQueriesRunCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasGraphqlQueriesRunParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_graphql-queries_run",
		Method:             "POST",
		PathPattern:        "/extras/graphql-queries/{id}/run/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasGraphqlQueriesRunReader{formats: a.formats},
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
	success, ok := result.(*ExtrasGraphqlQueriesRunCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_graphql-queries_run: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasGraphqlQueriesUpdate extras graphql queries update API
*/
func (a *Client) ExtrasGraphqlQueriesUpdate(params *ExtrasGraphqlQueriesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasGraphqlQueriesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasGraphqlQueriesUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_graphql-queries_update",
		Method:             "PUT",
		PathPattern:        "/extras/graphql-queries/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasGraphqlQueriesUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasGraphqlQueriesUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_graphql-queries_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasImageAttachmentsBulkDelete extras image attachments bulk delete API
*/
func (a *Client) ExtrasImageAttachmentsBulkDelete(params *ExtrasImageAttachmentsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasImageAttachmentsBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasImageAttachmentsBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_image-attachments_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/extras/image-attachments/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasImageAttachmentsBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*ExtrasImageAttachmentsBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_image-attachments_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasImageAttachmentsBulkPartialUpdate extras image attachments bulk partial update API
*/
func (a *Client) ExtrasImageAttachmentsBulkPartialUpdate(params *ExtrasImageAttachmentsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasImageAttachmentsBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasImageAttachmentsBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_image-attachments_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/extras/image-attachments/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasImageAttachmentsBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasImageAttachmentsBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_image-attachments_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasImageAttachmentsBulkUpdate extras image attachments bulk update API
*/
func (a *Client) ExtrasImageAttachmentsBulkUpdate(params *ExtrasImageAttachmentsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasImageAttachmentsBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasImageAttachmentsBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_image-attachments_bulk_update",
		Method:             "PUT",
		PathPattern:        "/extras/image-attachments/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasImageAttachmentsBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasImageAttachmentsBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_image-attachments_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasImageAttachmentsCreate extras image attachments create API
*/
func (a *Client) ExtrasImageAttachmentsCreate(params *ExtrasImageAttachmentsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasImageAttachmentsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasImageAttachmentsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_image-attachments_create",
		Method:             "POST",
		PathPattern:        "/extras/image-attachments/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasImageAttachmentsCreateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasImageAttachmentsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_image-attachments_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasImageAttachmentsDelete extras image attachments delete API
*/
func (a *Client) ExtrasImageAttachmentsDelete(params *ExtrasImageAttachmentsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasImageAttachmentsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasImageAttachmentsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_image-attachments_delete",
		Method:             "DELETE",
		PathPattern:        "/extras/image-attachments/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasImageAttachmentsDeleteReader{formats: a.formats},
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
	success, ok := result.(*ExtrasImageAttachmentsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_image-attachments_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasImageAttachmentsList extras image attachments list API
*/
func (a *Client) ExtrasImageAttachmentsList(params *ExtrasImageAttachmentsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasImageAttachmentsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasImageAttachmentsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_image-attachments_list",
		Method:             "GET",
		PathPattern:        "/extras/image-attachments/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasImageAttachmentsListReader{formats: a.formats},
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
	success, ok := result.(*ExtrasImageAttachmentsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_image-attachments_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasImageAttachmentsPartialUpdate extras image attachments partial update API
*/
func (a *Client) ExtrasImageAttachmentsPartialUpdate(params *ExtrasImageAttachmentsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasImageAttachmentsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasImageAttachmentsPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_image-attachments_partial_update",
		Method:             "PATCH",
		PathPattern:        "/extras/image-attachments/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasImageAttachmentsPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasImageAttachmentsPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_image-attachments_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasImageAttachmentsRead extras image attachments read API
*/
func (a *Client) ExtrasImageAttachmentsRead(params *ExtrasImageAttachmentsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasImageAttachmentsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasImageAttachmentsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_image-attachments_read",
		Method:             "GET",
		PathPattern:        "/extras/image-attachments/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasImageAttachmentsReadReader{formats: a.formats},
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
	success, ok := result.(*ExtrasImageAttachmentsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_image-attachments_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasImageAttachmentsUpdate extras image attachments update API
*/
func (a *Client) ExtrasImageAttachmentsUpdate(params *ExtrasImageAttachmentsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasImageAttachmentsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasImageAttachmentsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_image-attachments_update",
		Method:             "PUT",
		PathPattern:        "/extras/image-attachments/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasImageAttachmentsUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasImageAttachmentsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_image-attachments_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasJobResultsBulkDelete Retrieve a list of job results
*/
func (a *Client) ExtrasJobResultsBulkDelete(params *ExtrasJobResultsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasJobResultsBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasJobResultsBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_job-results_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/extras/job-results/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasJobResultsBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*ExtrasJobResultsBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_job-results_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasJobResultsBulkPartialUpdate Retrieve a list of job results
*/
func (a *Client) ExtrasJobResultsBulkPartialUpdate(params *ExtrasJobResultsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasJobResultsBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasJobResultsBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_job-results_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/extras/job-results/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasJobResultsBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasJobResultsBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_job-results_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasJobResultsBulkUpdate Retrieve a list of job results
*/
func (a *Client) ExtrasJobResultsBulkUpdate(params *ExtrasJobResultsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasJobResultsBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasJobResultsBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_job-results_bulk_update",
		Method:             "PUT",
		PathPattern:        "/extras/job-results/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasJobResultsBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasJobResultsBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_job-results_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasJobResultsCreate Retrieve a list of job results
*/
func (a *Client) ExtrasJobResultsCreate(params *ExtrasJobResultsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasJobResultsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasJobResultsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_job-results_create",
		Method:             "POST",
		PathPattern:        "/extras/job-results/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasJobResultsCreateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasJobResultsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_job-results_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasJobResultsDelete Retrieve a list of job results
*/
func (a *Client) ExtrasJobResultsDelete(params *ExtrasJobResultsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasJobResultsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasJobResultsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_job-results_delete",
		Method:             "DELETE",
		PathPattern:        "/extras/job-results/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasJobResultsDeleteReader{formats: a.formats},
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
	success, ok := result.(*ExtrasJobResultsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_job-results_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasJobResultsList Retrieve a list of job results
*/
func (a *Client) ExtrasJobResultsList(params *ExtrasJobResultsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasJobResultsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasJobResultsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_job-results_list",
		Method:             "GET",
		PathPattern:        "/extras/job-results/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasJobResultsListReader{formats: a.formats},
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
	success, ok := result.(*ExtrasJobResultsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_job-results_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasJobResultsPartialUpdate Retrieve a list of job results
*/
func (a *Client) ExtrasJobResultsPartialUpdate(params *ExtrasJobResultsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasJobResultsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasJobResultsPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_job-results_partial_update",
		Method:             "PATCH",
		PathPattern:        "/extras/job-results/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasJobResultsPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasJobResultsPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_job-results_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasJobResultsRead Retrieve a list of job results
*/
func (a *Client) ExtrasJobResultsRead(params *ExtrasJobResultsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasJobResultsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasJobResultsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_job-results_read",
		Method:             "GET",
		PathPattern:        "/extras/job-results/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasJobResultsReadReader{formats: a.formats},
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
	success, ok := result.(*ExtrasJobResultsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_job-results_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasJobResultsUpdate Retrieve a list of job results
*/
func (a *Client) ExtrasJobResultsUpdate(params *ExtrasJobResultsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasJobResultsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasJobResultsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_job-results_update",
		Method:             "PUT",
		PathPattern:        "/extras/job-results/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasJobResultsUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasJobResultsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_job-results_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasJobsList extras jobs list API
*/
func (a *Client) ExtrasJobsList(params *ExtrasJobsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasJobsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasJobsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_jobs_list",
		Method:             "GET",
		PathPattern:        "/extras/jobs/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasJobsListReader{formats: a.formats},
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
	success, ok := result.(*ExtrasJobsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_jobs_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasJobsRead extras jobs read API
*/
func (a *Client) ExtrasJobsRead(params *ExtrasJobsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasJobsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasJobsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_jobs_read",
		Method:             "GET",
		PathPattern:        "/extras/jobs/{class_path}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasJobsReadReader{formats: a.formats},
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
	success, ok := result.(*ExtrasJobsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_jobs_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasJobsRun extras jobs run API
*/
func (a *Client) ExtrasJobsRun(params *ExtrasJobsRunParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasJobsRunCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasJobsRunParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_jobs_run",
		Method:             "POST",
		PathPattern:        "/extras/jobs/{class_path}/run/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasJobsRunReader{formats: a.formats},
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
	success, ok := result.(*ExtrasJobsRunCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_jobs_run: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasObjectChangesList Retrieve a list of recent changes.
*/
func (a *Client) ExtrasObjectChangesList(params *ExtrasObjectChangesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasObjectChangesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasObjectChangesListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_object-changes_list",
		Method:             "GET",
		PathPattern:        "/extras/object-changes/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasObjectChangesListReader{formats: a.formats},
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
	success, ok := result.(*ExtrasObjectChangesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_object-changes_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasObjectChangesRead Retrieve a list of recent changes.
*/
func (a *Client) ExtrasObjectChangesRead(params *ExtrasObjectChangesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasObjectChangesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasObjectChangesReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_object-changes_read",
		Method:             "GET",
		PathPattern:        "/extras/object-changes/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasObjectChangesReadReader{formats: a.formats},
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
	success, ok := result.(*ExtrasObjectChangesReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_object-changes_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasRelationshipAssociationsBulkDelete extras relationship associations bulk delete API
*/
func (a *Client) ExtrasRelationshipAssociationsBulkDelete(params *ExtrasRelationshipAssociationsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasRelationshipAssociationsBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasRelationshipAssociationsBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_relationship-associations_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/extras/relationship-associations/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasRelationshipAssociationsBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*ExtrasRelationshipAssociationsBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_relationship-associations_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasRelationshipAssociationsBulkPartialUpdate extras relationship associations bulk partial update API
*/
func (a *Client) ExtrasRelationshipAssociationsBulkPartialUpdate(params *ExtrasRelationshipAssociationsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasRelationshipAssociationsBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasRelationshipAssociationsBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_relationship-associations_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/extras/relationship-associations/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasRelationshipAssociationsBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasRelationshipAssociationsBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_relationship-associations_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasRelationshipAssociationsBulkUpdate extras relationship associations bulk update API
*/
func (a *Client) ExtrasRelationshipAssociationsBulkUpdate(params *ExtrasRelationshipAssociationsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasRelationshipAssociationsBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasRelationshipAssociationsBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_relationship-associations_bulk_update",
		Method:             "PUT",
		PathPattern:        "/extras/relationship-associations/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasRelationshipAssociationsBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasRelationshipAssociationsBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_relationship-associations_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasRelationshipAssociationsCreate extras relationship associations create API
*/
func (a *Client) ExtrasRelationshipAssociationsCreate(params *ExtrasRelationshipAssociationsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasRelationshipAssociationsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasRelationshipAssociationsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_relationship-associations_create",
		Method:             "POST",
		PathPattern:        "/extras/relationship-associations/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasRelationshipAssociationsCreateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasRelationshipAssociationsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_relationship-associations_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasRelationshipAssociationsDelete extras relationship associations delete API
*/
func (a *Client) ExtrasRelationshipAssociationsDelete(params *ExtrasRelationshipAssociationsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasRelationshipAssociationsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasRelationshipAssociationsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_relationship-associations_delete",
		Method:             "DELETE",
		PathPattern:        "/extras/relationship-associations/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasRelationshipAssociationsDeleteReader{formats: a.formats},
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
	success, ok := result.(*ExtrasRelationshipAssociationsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_relationship-associations_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasRelationshipAssociationsList extras relationship associations list API
*/
func (a *Client) ExtrasRelationshipAssociationsList(params *ExtrasRelationshipAssociationsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasRelationshipAssociationsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasRelationshipAssociationsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_relationship-associations_list",
		Method:             "GET",
		PathPattern:        "/extras/relationship-associations/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasRelationshipAssociationsListReader{formats: a.formats},
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
	success, ok := result.(*ExtrasRelationshipAssociationsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_relationship-associations_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasRelationshipAssociationsPartialUpdate extras relationship associations partial update API
*/
func (a *Client) ExtrasRelationshipAssociationsPartialUpdate(params *ExtrasRelationshipAssociationsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasRelationshipAssociationsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasRelationshipAssociationsPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_relationship-associations_partial_update",
		Method:             "PATCH",
		PathPattern:        "/extras/relationship-associations/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasRelationshipAssociationsPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasRelationshipAssociationsPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_relationship-associations_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasRelationshipAssociationsRead extras relationship associations read API
*/
func (a *Client) ExtrasRelationshipAssociationsRead(params *ExtrasRelationshipAssociationsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasRelationshipAssociationsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasRelationshipAssociationsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_relationship-associations_read",
		Method:             "GET",
		PathPattern:        "/extras/relationship-associations/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasRelationshipAssociationsReadReader{formats: a.formats},
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
	success, ok := result.(*ExtrasRelationshipAssociationsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_relationship-associations_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasRelationshipAssociationsUpdate extras relationship associations update API
*/
func (a *Client) ExtrasRelationshipAssociationsUpdate(params *ExtrasRelationshipAssociationsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasRelationshipAssociationsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasRelationshipAssociationsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_relationship-associations_update",
		Method:             "PUT",
		PathPattern:        "/extras/relationship-associations/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasRelationshipAssociationsUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasRelationshipAssociationsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_relationship-associations_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasRelationshipsBulkDelete extras relationships bulk delete API
*/
func (a *Client) ExtrasRelationshipsBulkDelete(params *ExtrasRelationshipsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasRelationshipsBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasRelationshipsBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_relationships_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/extras/relationships/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasRelationshipsBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*ExtrasRelationshipsBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_relationships_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasRelationshipsBulkPartialUpdate extras relationships bulk partial update API
*/
func (a *Client) ExtrasRelationshipsBulkPartialUpdate(params *ExtrasRelationshipsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasRelationshipsBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasRelationshipsBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_relationships_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/extras/relationships/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasRelationshipsBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasRelationshipsBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_relationships_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasRelationshipsBulkUpdate extras relationships bulk update API
*/
func (a *Client) ExtrasRelationshipsBulkUpdate(params *ExtrasRelationshipsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasRelationshipsBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasRelationshipsBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_relationships_bulk_update",
		Method:             "PUT",
		PathPattern:        "/extras/relationships/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasRelationshipsBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasRelationshipsBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_relationships_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasRelationshipsCreate extras relationships create API
*/
func (a *Client) ExtrasRelationshipsCreate(params *ExtrasRelationshipsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasRelationshipsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasRelationshipsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_relationships_create",
		Method:             "POST",
		PathPattern:        "/extras/relationships/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasRelationshipsCreateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasRelationshipsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_relationships_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasRelationshipsDelete extras relationships delete API
*/
func (a *Client) ExtrasRelationshipsDelete(params *ExtrasRelationshipsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasRelationshipsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasRelationshipsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_relationships_delete",
		Method:             "DELETE",
		PathPattern:        "/extras/relationships/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasRelationshipsDeleteReader{formats: a.formats},
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
	success, ok := result.(*ExtrasRelationshipsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_relationships_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasRelationshipsList extras relationships list API
*/
func (a *Client) ExtrasRelationshipsList(params *ExtrasRelationshipsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasRelationshipsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasRelationshipsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_relationships_list",
		Method:             "GET",
		PathPattern:        "/extras/relationships/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasRelationshipsListReader{formats: a.formats},
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
	success, ok := result.(*ExtrasRelationshipsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_relationships_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasRelationshipsPartialUpdate extras relationships partial update API
*/
func (a *Client) ExtrasRelationshipsPartialUpdate(params *ExtrasRelationshipsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasRelationshipsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasRelationshipsPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_relationships_partial_update",
		Method:             "PATCH",
		PathPattern:        "/extras/relationships/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasRelationshipsPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasRelationshipsPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_relationships_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasRelationshipsRead extras relationships read API
*/
func (a *Client) ExtrasRelationshipsRead(params *ExtrasRelationshipsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasRelationshipsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasRelationshipsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_relationships_read",
		Method:             "GET",
		PathPattern:        "/extras/relationships/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasRelationshipsReadReader{formats: a.formats},
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
	success, ok := result.(*ExtrasRelationshipsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_relationships_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasRelationshipsUpdate extras relationships update API
*/
func (a *Client) ExtrasRelationshipsUpdate(params *ExtrasRelationshipsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasRelationshipsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasRelationshipsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_relationships_update",
		Method:             "PUT",
		PathPattern:        "/extras/relationships/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasRelationshipsUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasRelationshipsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_relationships_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasStatusesBulkDelete View and manage custom status choices for objects with a `status` field.
*/
func (a *Client) ExtrasStatusesBulkDelete(params *ExtrasStatusesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasStatusesBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasStatusesBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_statuses_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/extras/statuses/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasStatusesBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*ExtrasStatusesBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_statuses_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasStatusesBulkPartialUpdate View and manage custom status choices for objects with a `status` field.
*/
func (a *Client) ExtrasStatusesBulkPartialUpdate(params *ExtrasStatusesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasStatusesBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasStatusesBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_statuses_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/extras/statuses/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasStatusesBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasStatusesBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_statuses_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasStatusesBulkUpdate View and manage custom status choices for objects with a `status` field.
*/
func (a *Client) ExtrasStatusesBulkUpdate(params *ExtrasStatusesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasStatusesBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasStatusesBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_statuses_bulk_update",
		Method:             "PUT",
		PathPattern:        "/extras/statuses/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasStatusesBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasStatusesBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_statuses_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasStatusesCreate View and manage custom status choices for objects with a `status` field.
*/
func (a *Client) ExtrasStatusesCreate(params *ExtrasStatusesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasStatusesCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasStatusesCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_statuses_create",
		Method:             "POST",
		PathPattern:        "/extras/statuses/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasStatusesCreateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasStatusesCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_statuses_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasStatusesDelete View and manage custom status choices for objects with a `status` field.
*/
func (a *Client) ExtrasStatusesDelete(params *ExtrasStatusesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasStatusesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasStatusesDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_statuses_delete",
		Method:             "DELETE",
		PathPattern:        "/extras/statuses/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasStatusesDeleteReader{formats: a.formats},
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
	success, ok := result.(*ExtrasStatusesDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_statuses_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasStatusesList View and manage custom status choices for objects with a `status` field.
*/
func (a *Client) ExtrasStatusesList(params *ExtrasStatusesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasStatusesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasStatusesListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_statuses_list",
		Method:             "GET",
		PathPattern:        "/extras/statuses/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasStatusesListReader{formats: a.formats},
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
	success, ok := result.(*ExtrasStatusesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_statuses_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasStatusesPartialUpdate View and manage custom status choices for objects with a `status` field.
*/
func (a *Client) ExtrasStatusesPartialUpdate(params *ExtrasStatusesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasStatusesPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasStatusesPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_statuses_partial_update",
		Method:             "PATCH",
		PathPattern:        "/extras/statuses/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasStatusesPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasStatusesPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_statuses_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasStatusesRead View and manage custom status choices for objects with a `status` field.
*/
func (a *Client) ExtrasStatusesRead(params *ExtrasStatusesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasStatusesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasStatusesReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_statuses_read",
		Method:             "GET",
		PathPattern:        "/extras/statuses/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasStatusesReadReader{formats: a.formats},
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
	success, ok := result.(*ExtrasStatusesReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_statuses_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasStatusesUpdate View and manage custom status choices for objects with a `status` field.
*/
func (a *Client) ExtrasStatusesUpdate(params *ExtrasStatusesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasStatusesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasStatusesUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_statuses_update",
		Method:             "PUT",
		PathPattern:        "/extras/statuses/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasStatusesUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasStatusesUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_statuses_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasTagsBulkDelete extras tags bulk delete API
*/
func (a *Client) ExtrasTagsBulkDelete(params *ExtrasTagsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasTagsBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasTagsBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_tags_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/extras/tags/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasTagsBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*ExtrasTagsBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_tags_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasTagsBulkPartialUpdate extras tags bulk partial update API
*/
func (a *Client) ExtrasTagsBulkPartialUpdate(params *ExtrasTagsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasTagsBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasTagsBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_tags_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/extras/tags/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasTagsBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasTagsBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_tags_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasTagsBulkUpdate extras tags bulk update API
*/
func (a *Client) ExtrasTagsBulkUpdate(params *ExtrasTagsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasTagsBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasTagsBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_tags_bulk_update",
		Method:             "PUT",
		PathPattern:        "/extras/tags/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasTagsBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasTagsBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_tags_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasTagsCreate extras tags create API
*/
func (a *Client) ExtrasTagsCreate(params *ExtrasTagsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasTagsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasTagsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_tags_create",
		Method:             "POST",
		PathPattern:        "/extras/tags/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasTagsCreateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasTagsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_tags_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasTagsDelete extras tags delete API
*/
func (a *Client) ExtrasTagsDelete(params *ExtrasTagsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasTagsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasTagsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_tags_delete",
		Method:             "DELETE",
		PathPattern:        "/extras/tags/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasTagsDeleteReader{formats: a.formats},
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
	success, ok := result.(*ExtrasTagsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_tags_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasTagsList extras tags list API
*/
func (a *Client) ExtrasTagsList(params *ExtrasTagsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasTagsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasTagsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_tags_list",
		Method:             "GET",
		PathPattern:        "/extras/tags/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasTagsListReader{formats: a.formats},
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
	success, ok := result.(*ExtrasTagsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_tags_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasTagsPartialUpdate extras tags partial update API
*/
func (a *Client) ExtrasTagsPartialUpdate(params *ExtrasTagsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasTagsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasTagsPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_tags_partial_update",
		Method:             "PATCH",
		PathPattern:        "/extras/tags/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasTagsPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasTagsPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_tags_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasTagsRead extras tags read API
*/
func (a *Client) ExtrasTagsRead(params *ExtrasTagsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasTagsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasTagsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_tags_read",
		Method:             "GET",
		PathPattern:        "/extras/tags/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasTagsReadReader{formats: a.formats},
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
	success, ok := result.(*ExtrasTagsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_tags_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasTagsUpdate extras tags update API
*/
func (a *Client) ExtrasTagsUpdate(params *ExtrasTagsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasTagsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasTagsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_tags_update",
		Method:             "PUT",
		PathPattern:        "/extras/tags/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasTagsUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasTagsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_tags_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasWebhooksBulkDelete Manage Webhooks through DELETE, GET, POST, PUT, and PATCH requests.
*/
func (a *Client) ExtrasWebhooksBulkDelete(params *ExtrasWebhooksBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasWebhooksBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasWebhooksBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_webhooks_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/extras/webhooks/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasWebhooksBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*ExtrasWebhooksBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_webhooks_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasWebhooksBulkPartialUpdate Manage Webhooks through DELETE, GET, POST, PUT, and PATCH requests.
*/
func (a *Client) ExtrasWebhooksBulkPartialUpdate(params *ExtrasWebhooksBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasWebhooksBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasWebhooksBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_webhooks_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/extras/webhooks/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasWebhooksBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasWebhooksBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_webhooks_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasWebhooksBulkUpdate Manage Webhooks through DELETE, GET, POST, PUT, and PATCH requests.
*/
func (a *Client) ExtrasWebhooksBulkUpdate(params *ExtrasWebhooksBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasWebhooksBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasWebhooksBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_webhooks_bulk_update",
		Method:             "PUT",
		PathPattern:        "/extras/webhooks/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasWebhooksBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasWebhooksBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_webhooks_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasWebhooksCreate Manage Webhooks through DELETE, GET, POST, PUT, and PATCH requests.
*/
func (a *Client) ExtrasWebhooksCreate(params *ExtrasWebhooksCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasWebhooksCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasWebhooksCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_webhooks_create",
		Method:             "POST",
		PathPattern:        "/extras/webhooks/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasWebhooksCreateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasWebhooksCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_webhooks_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasWebhooksDelete Manage Webhooks through DELETE, GET, POST, PUT, and PATCH requests.
*/
func (a *Client) ExtrasWebhooksDelete(params *ExtrasWebhooksDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasWebhooksDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasWebhooksDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_webhooks_delete",
		Method:             "DELETE",
		PathPattern:        "/extras/webhooks/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasWebhooksDeleteReader{formats: a.formats},
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
	success, ok := result.(*ExtrasWebhooksDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_webhooks_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasWebhooksList Manage Webhooks through DELETE, GET, POST, PUT, and PATCH requests.
*/
func (a *Client) ExtrasWebhooksList(params *ExtrasWebhooksListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasWebhooksListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasWebhooksListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_webhooks_list",
		Method:             "GET",
		PathPattern:        "/extras/webhooks/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasWebhooksListReader{formats: a.formats},
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
	success, ok := result.(*ExtrasWebhooksListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_webhooks_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasWebhooksPartialUpdate Manage Webhooks through DELETE, GET, POST, PUT, and PATCH requests.
*/
func (a *Client) ExtrasWebhooksPartialUpdate(params *ExtrasWebhooksPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasWebhooksPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasWebhooksPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_webhooks_partial_update",
		Method:             "PATCH",
		PathPattern:        "/extras/webhooks/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasWebhooksPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasWebhooksPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_webhooks_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasWebhooksRead Manage Webhooks through DELETE, GET, POST, PUT, and PATCH requests.
*/
func (a *Client) ExtrasWebhooksRead(params *ExtrasWebhooksReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasWebhooksReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasWebhooksReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_webhooks_read",
		Method:             "GET",
		PathPattern:        "/extras/webhooks/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasWebhooksReadReader{formats: a.formats},
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
	success, ok := result.(*ExtrasWebhooksReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_webhooks_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExtrasWebhooksUpdate Manage Webhooks through DELETE, GET, POST, PUT, and PATCH requests.
*/
func (a *Client) ExtrasWebhooksUpdate(params *ExtrasWebhooksUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExtrasWebhooksUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtrasWebhooksUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "extras_webhooks_update",
		Method:             "PUT",
		PathPattern:        "/extras/webhooks/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExtrasWebhooksUpdateReader{formats: a.formats},
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
	success, ok := result.(*ExtrasWebhooksUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for extras_webhooks_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
