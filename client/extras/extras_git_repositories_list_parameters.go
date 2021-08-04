package extras

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewExtrasGitRepositoriesListParams creates a new ExtrasGitRepositoriesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasGitRepositoriesListParams() *ExtrasGitRepositoriesListParams {
	return &ExtrasGitRepositoriesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasGitRepositoriesListParamsWithTimeout creates a new ExtrasGitRepositoriesListParams object
// with the ability to set a timeout on a request.
func NewExtrasGitRepositoriesListParamsWithTimeout(timeout time.Duration) *ExtrasGitRepositoriesListParams {
	return &ExtrasGitRepositoriesListParams{
		timeout: timeout,
	}
}

// NewExtrasGitRepositoriesListParamsWithContext creates a new ExtrasGitRepositoriesListParams object
// with the ability to set a context for a request.
func NewExtrasGitRepositoriesListParamsWithContext(ctx context.Context) *ExtrasGitRepositoriesListParams {
	return &ExtrasGitRepositoriesListParams{
		Context: ctx,
	}
}

// NewExtrasGitRepositoriesListParamsWithHTTPClient creates a new ExtrasGitRepositoriesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasGitRepositoriesListParamsWithHTTPClient(client *http.Client) *ExtrasGitRepositoriesListParams {
	return &ExtrasGitRepositoriesListParams{
		HTTPClient: client,
	}
}

/* ExtrasGitRepositoriesListParams contains all the parameters to send to the API endpoint
   for the extras git repositories list operation.

   Typically these are written to a http.Request.
*/
type ExtrasGitRepositoriesListParams struct {

	// Branch.
	Branch *string

	// BranchIc.
	BranchIc *string

	// BranchIe.
	BranchIe *string

	// BranchIew.
	BranchIew *string

	// BranchIsw.
	BranchIsw *string

	// Branchn.
	Branchn *string

	// BranchNic.
	BranchNic *string

	// BranchNie.
	BranchNie *string

	// BranchNiew.
	BranchNiew *string

	// BranchNisw.
	BranchNisw *string

	// Created.
	Created *string

	// CreatedGte.
	CreatedGte *string

	// CreatedLte.
	CreatedLte *string

	// ID.
	ID *string

	// IDIc.
	IDIc *string

	// IDIe.
	IDIe *string

	// IDIew.
	IDIew *string

	// IDIsw.
	IDIsw *string

	// IDn.
	IDn *string

	// IDNic.
	IDNic *string

	// IDNie.
	IDNie *string

	// IDNiew.
	IDNiew *string

	// IDNisw.
	IDNisw *string

	// LastUpdated.
	LastUpdated *string

	// LastUpdatedGte.
	LastUpdatedGte *string

	// LastUpdatedLte.
	LastUpdatedLte *string

	/* Limit.

	   Number of results to return per page.
	*/
	Limit *int64

	// Name.
	Name *string

	// NameIc.
	NameIc *string

	// NameIe.
	NameIe *string

	// NameIew.
	NameIew *string

	// NameIsw.
	NameIsw *string

	// Namen.
	Namen *string

	// NameNic.
	NameNic *string

	// NameNie.
	NameNie *string

	// NameNiew.
	NameNiew *string

	// NameNisw.
	NameNisw *string

	/* Offset.

	   The initial index from which to return the results.
	*/
	Offset *int64

	// Q.
	Q *string

	// RemoteURL.
	RemoteURL *string

	// RemoteURLIc.
	RemoteURLIc *string

	// RemoteURLIe.
	RemoteURLIe *string

	// RemoteURLIew.
	RemoteURLIew *string

	// RemoteURLIsw.
	RemoteURLIsw *string

	// RemoteURLn.
	RemoteURLn *string

	// RemoteURLNic.
	RemoteURLNic *string

	// RemoteURLNie.
	RemoteURLNie *string

	// RemoteURLNiew.
	RemoteURLNiew *string

	// RemoteURLNisw.
	RemoteURLNisw *string

	// Slug.
	Slug *string

	// SlugIc.
	SlugIc *string

	// SlugIe.
	SlugIe *string

	// SlugIew.
	SlugIew *string

	// SlugIsw.
	SlugIsw *string

	// Slugn.
	Slugn *string

	// SlugNic.
	SlugNic *string

	// SlugNie.
	SlugNie *string

	// SlugNiew.
	SlugNiew *string

	// SlugNisw.
	SlugNisw *string

	// Tag.
	Tag *string

	// Tagn.
	Tagn *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras git repositories list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasGitRepositoriesListParams) WithDefaults() *ExtrasGitRepositoriesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras git repositories list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasGitRepositoriesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithTimeout(timeout time.Duration) *ExtrasGitRepositoriesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithContext(ctx context.Context) *ExtrasGitRepositoriesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithHTTPClient(client *http.Client) *ExtrasGitRepositoriesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBranch adds the branch to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithBranch(branch *string) *ExtrasGitRepositoriesListParams {
	o.SetBranch(branch)
	return o
}

// SetBranch adds the branch to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetBranch(branch *string) {
	o.Branch = branch
}

// WithBranchIc adds the branchIc to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithBranchIc(branchIc *string) *ExtrasGitRepositoriesListParams {
	o.SetBranchIc(branchIc)
	return o
}

// SetBranchIc adds the branchIc to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetBranchIc(branchIc *string) {
	o.BranchIc = branchIc
}

// WithBranchIe adds the branchIe to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithBranchIe(branchIe *string) *ExtrasGitRepositoriesListParams {
	o.SetBranchIe(branchIe)
	return o
}

// SetBranchIe adds the branchIe to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetBranchIe(branchIe *string) {
	o.BranchIe = branchIe
}

// WithBranchIew adds the branchIew to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithBranchIew(branchIew *string) *ExtrasGitRepositoriesListParams {
	o.SetBranchIew(branchIew)
	return o
}

// SetBranchIew adds the branchIew to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetBranchIew(branchIew *string) {
	o.BranchIew = branchIew
}

// WithBranchIsw adds the branchIsw to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithBranchIsw(branchIsw *string) *ExtrasGitRepositoriesListParams {
	o.SetBranchIsw(branchIsw)
	return o
}

// SetBranchIsw adds the branchIsw to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetBranchIsw(branchIsw *string) {
	o.BranchIsw = branchIsw
}

// WithBranchn adds the branchn to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithBranchn(branchn *string) *ExtrasGitRepositoriesListParams {
	o.SetBranchn(branchn)
	return o
}

// SetBranchn adds the branchN to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetBranchn(branchn *string) {
	o.Branchn = branchn
}

// WithBranchNic adds the branchNic to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithBranchNic(branchNic *string) *ExtrasGitRepositoriesListParams {
	o.SetBranchNic(branchNic)
	return o
}

// SetBranchNic adds the branchNic to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetBranchNic(branchNic *string) {
	o.BranchNic = branchNic
}

// WithBranchNie adds the branchNie to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithBranchNie(branchNie *string) *ExtrasGitRepositoriesListParams {
	o.SetBranchNie(branchNie)
	return o
}

// SetBranchNie adds the branchNie to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetBranchNie(branchNie *string) {
	o.BranchNie = branchNie
}

// WithBranchNiew adds the branchNiew to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithBranchNiew(branchNiew *string) *ExtrasGitRepositoriesListParams {
	o.SetBranchNiew(branchNiew)
	return o
}

// SetBranchNiew adds the branchNiew to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetBranchNiew(branchNiew *string) {
	o.BranchNiew = branchNiew
}

// WithBranchNisw adds the branchNisw to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithBranchNisw(branchNisw *string) *ExtrasGitRepositoriesListParams {
	o.SetBranchNisw(branchNisw)
	return o
}

// SetBranchNisw adds the branchNisw to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetBranchNisw(branchNisw *string) {
	o.BranchNisw = branchNisw
}

// WithCreated adds the created to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithCreated(created *string) *ExtrasGitRepositoriesListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGte adds the createdGte to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithCreatedGte(createdGte *string) *ExtrasGitRepositoriesListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLte adds the createdLte to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithCreatedLte(createdLte *string) *ExtrasGitRepositoriesListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithID adds the id to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithID(id *string) *ExtrasGitRepositoriesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithIDIc(iDIc *string) *ExtrasGitRepositoriesListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithIDIe(iDIe *string) *ExtrasGitRepositoriesListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithIDIew(iDIew *string) *ExtrasGitRepositoriesListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithIDIsw(iDIsw *string) *ExtrasGitRepositoriesListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithIDn(iDn *string) *ExtrasGitRepositoriesListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithIDNic(iDNic *string) *ExtrasGitRepositoriesListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithIDNie(iDNie *string) *ExtrasGitRepositoriesListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithIDNiew(iDNiew *string) *ExtrasGitRepositoriesListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithIDNisw(iDNisw *string) *ExtrasGitRepositoriesListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithLastUpdated adds the lastUpdated to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithLastUpdated(lastUpdated *string) *ExtrasGitRepositoriesListParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetLastUpdated(lastUpdated *string) {
	o.LastUpdated = lastUpdated
}

// WithLastUpdatedGte adds the lastUpdatedGte to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithLastUpdatedGte(lastUpdatedGte *string) *ExtrasGitRepositoriesListParams {
	o.SetLastUpdatedGte(lastUpdatedGte)
	return o
}

// SetLastUpdatedGte adds the lastUpdatedGte to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetLastUpdatedGte(lastUpdatedGte *string) {
	o.LastUpdatedGte = lastUpdatedGte
}

// WithLastUpdatedLte adds the lastUpdatedLte to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithLastUpdatedLte(lastUpdatedLte *string) *ExtrasGitRepositoriesListParams {
	o.SetLastUpdatedLte(lastUpdatedLte)
	return o
}

// SetLastUpdatedLte adds the lastUpdatedLte to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetLastUpdatedLte(lastUpdatedLte *string) {
	o.LastUpdatedLte = lastUpdatedLte
}

// WithLimit adds the limit to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithLimit(limit *int64) *ExtrasGitRepositoriesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithName(name *string) *ExtrasGitRepositoriesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithNameIc(nameIc *string) *ExtrasGitRepositoriesListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithNameIe(nameIe *string) *ExtrasGitRepositoriesListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithNameIew(nameIew *string) *ExtrasGitRepositoriesListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithNameIsw(nameIsw *string) *ExtrasGitRepositoriesListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithNamen(namen *string) *ExtrasGitRepositoriesListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithNameNic(nameNic *string) *ExtrasGitRepositoriesListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithNameNie(nameNie *string) *ExtrasGitRepositoriesListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithNameNiew(nameNiew *string) *ExtrasGitRepositoriesListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithNameNisw(nameNisw *string) *ExtrasGitRepositoriesListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithOffset(offset *int64) *ExtrasGitRepositoriesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithQ(q *string) *ExtrasGitRepositoriesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetQ(q *string) {
	o.Q = q
}

// WithRemoteURL adds the remoteURL to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithRemoteURL(remoteURL *string) *ExtrasGitRepositoriesListParams {
	o.SetRemoteURL(remoteURL)
	return o
}

// SetRemoteURL adds the remoteUrl to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetRemoteURL(remoteURL *string) {
	o.RemoteURL = remoteURL
}

// WithRemoteURLIc adds the remoteURLIc to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithRemoteURLIc(remoteURLIc *string) *ExtrasGitRepositoriesListParams {
	o.SetRemoteURLIc(remoteURLIc)
	return o
}

// SetRemoteURLIc adds the remoteUrlIc to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetRemoteURLIc(remoteURLIc *string) {
	o.RemoteURLIc = remoteURLIc
}

// WithRemoteURLIe adds the remoteURLIe to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithRemoteURLIe(remoteURLIe *string) *ExtrasGitRepositoriesListParams {
	o.SetRemoteURLIe(remoteURLIe)
	return o
}

// SetRemoteURLIe adds the remoteUrlIe to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetRemoteURLIe(remoteURLIe *string) {
	o.RemoteURLIe = remoteURLIe
}

// WithRemoteURLIew adds the remoteURLIew to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithRemoteURLIew(remoteURLIew *string) *ExtrasGitRepositoriesListParams {
	o.SetRemoteURLIew(remoteURLIew)
	return o
}

// SetRemoteURLIew adds the remoteUrlIew to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetRemoteURLIew(remoteURLIew *string) {
	o.RemoteURLIew = remoteURLIew
}

// WithRemoteURLIsw adds the remoteURLIsw to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithRemoteURLIsw(remoteURLIsw *string) *ExtrasGitRepositoriesListParams {
	o.SetRemoteURLIsw(remoteURLIsw)
	return o
}

// SetRemoteURLIsw adds the remoteUrlIsw to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetRemoteURLIsw(remoteURLIsw *string) {
	o.RemoteURLIsw = remoteURLIsw
}

// WithRemoteURLn adds the remoteURLn to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithRemoteURLn(remoteURLn *string) *ExtrasGitRepositoriesListParams {
	o.SetRemoteURLn(remoteURLn)
	return o
}

// SetRemoteURLn adds the remoteUrlN to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetRemoteURLn(remoteURLn *string) {
	o.RemoteURLn = remoteURLn
}

// WithRemoteURLNic adds the remoteURLNic to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithRemoteURLNic(remoteURLNic *string) *ExtrasGitRepositoriesListParams {
	o.SetRemoteURLNic(remoteURLNic)
	return o
}

// SetRemoteURLNic adds the remoteUrlNic to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetRemoteURLNic(remoteURLNic *string) {
	o.RemoteURLNic = remoteURLNic
}

// WithRemoteURLNie adds the remoteURLNie to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithRemoteURLNie(remoteURLNie *string) *ExtrasGitRepositoriesListParams {
	o.SetRemoteURLNie(remoteURLNie)
	return o
}

// SetRemoteURLNie adds the remoteUrlNie to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetRemoteURLNie(remoteURLNie *string) {
	o.RemoteURLNie = remoteURLNie
}

// WithRemoteURLNiew adds the remoteURLNiew to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithRemoteURLNiew(remoteURLNiew *string) *ExtrasGitRepositoriesListParams {
	o.SetRemoteURLNiew(remoteURLNiew)
	return o
}

// SetRemoteURLNiew adds the remoteUrlNiew to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetRemoteURLNiew(remoteURLNiew *string) {
	o.RemoteURLNiew = remoteURLNiew
}

// WithRemoteURLNisw adds the remoteURLNisw to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithRemoteURLNisw(remoteURLNisw *string) *ExtrasGitRepositoriesListParams {
	o.SetRemoteURLNisw(remoteURLNisw)
	return o
}

// SetRemoteURLNisw adds the remoteUrlNisw to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetRemoteURLNisw(remoteURLNisw *string) {
	o.RemoteURLNisw = remoteURLNisw
}

// WithSlug adds the slug to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithSlug(slug *string) *ExtrasGitRepositoriesListParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetSlug(slug *string) {
	o.Slug = slug
}

// WithSlugIc adds the slugIc to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithSlugIc(slugIc *string) *ExtrasGitRepositoriesListParams {
	o.SetSlugIc(slugIc)
	return o
}

// SetSlugIc adds the slugIc to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetSlugIc(slugIc *string) {
	o.SlugIc = slugIc
}

// WithSlugIe adds the slugIe to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithSlugIe(slugIe *string) *ExtrasGitRepositoriesListParams {
	o.SetSlugIe(slugIe)
	return o
}

// SetSlugIe adds the slugIe to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetSlugIe(slugIe *string) {
	o.SlugIe = slugIe
}

// WithSlugIew adds the slugIew to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithSlugIew(slugIew *string) *ExtrasGitRepositoriesListParams {
	o.SetSlugIew(slugIew)
	return o
}

// SetSlugIew adds the slugIew to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetSlugIew(slugIew *string) {
	o.SlugIew = slugIew
}

// WithSlugIsw adds the slugIsw to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithSlugIsw(slugIsw *string) *ExtrasGitRepositoriesListParams {
	o.SetSlugIsw(slugIsw)
	return o
}

// SetSlugIsw adds the slugIsw to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetSlugIsw(slugIsw *string) {
	o.SlugIsw = slugIsw
}

// WithSlugn adds the slugn to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithSlugn(slugn *string) *ExtrasGitRepositoriesListParams {
	o.SetSlugn(slugn)
	return o
}

// SetSlugn adds the slugN to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetSlugn(slugn *string) {
	o.Slugn = slugn
}

// WithSlugNic adds the slugNic to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithSlugNic(slugNic *string) *ExtrasGitRepositoriesListParams {
	o.SetSlugNic(slugNic)
	return o
}

// SetSlugNic adds the slugNic to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetSlugNic(slugNic *string) {
	o.SlugNic = slugNic
}

// WithSlugNie adds the slugNie to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithSlugNie(slugNie *string) *ExtrasGitRepositoriesListParams {
	o.SetSlugNie(slugNie)
	return o
}

// SetSlugNie adds the slugNie to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetSlugNie(slugNie *string) {
	o.SlugNie = slugNie
}

// WithSlugNiew adds the slugNiew to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithSlugNiew(slugNiew *string) *ExtrasGitRepositoriesListParams {
	o.SetSlugNiew(slugNiew)
	return o
}

// SetSlugNiew adds the slugNiew to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetSlugNiew(slugNiew *string) {
	o.SlugNiew = slugNiew
}

// WithSlugNisw adds the slugNisw to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithSlugNisw(slugNisw *string) *ExtrasGitRepositoriesListParams {
	o.SetSlugNisw(slugNisw)
	return o
}

// SetSlugNisw adds the slugNisw to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetSlugNisw(slugNisw *string) {
	o.SlugNisw = slugNisw
}

// WithTag adds the tag to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithTag(tag *string) *ExtrasGitRepositoriesListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTagn adds the tagn to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) WithTagn(tagn *string) *ExtrasGitRepositoriesListParams {
	o.SetTagn(tagn)
	return o
}

// SetTagn adds the tagN to the extras git repositories list params
func (o *ExtrasGitRepositoriesListParams) SetTagn(tagn *string) {
	o.Tagn = tagn
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasGitRepositoriesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Branch != nil {

		// query param branch
		var qrBranch string

		if o.Branch != nil {
			qrBranch = *o.Branch
		}
		qBranch := qrBranch
		if qBranch != "" {

			if err := r.SetQueryParam("branch", qBranch); err != nil {
				return err
			}
		}
	}

	if o.BranchIc != nil {

		// query param branch__ic
		var qrBranchIc string

		if o.BranchIc != nil {
			qrBranchIc = *o.BranchIc
		}
		qBranchIc := qrBranchIc
		if qBranchIc != "" {

			if err := r.SetQueryParam("branch__ic", qBranchIc); err != nil {
				return err
			}
		}
	}

	if o.BranchIe != nil {

		// query param branch__ie
		var qrBranchIe string

		if o.BranchIe != nil {
			qrBranchIe = *o.BranchIe
		}
		qBranchIe := qrBranchIe
		if qBranchIe != "" {

			if err := r.SetQueryParam("branch__ie", qBranchIe); err != nil {
				return err
			}
		}
	}

	if o.BranchIew != nil {

		// query param branch__iew
		var qrBranchIew string

		if o.BranchIew != nil {
			qrBranchIew = *o.BranchIew
		}
		qBranchIew := qrBranchIew
		if qBranchIew != "" {

			if err := r.SetQueryParam("branch__iew", qBranchIew); err != nil {
				return err
			}
		}
	}

	if o.BranchIsw != nil {

		// query param branch__isw
		var qrBranchIsw string

		if o.BranchIsw != nil {
			qrBranchIsw = *o.BranchIsw
		}
		qBranchIsw := qrBranchIsw
		if qBranchIsw != "" {

			if err := r.SetQueryParam("branch__isw", qBranchIsw); err != nil {
				return err
			}
		}
	}

	if o.Branchn != nil {

		// query param branch__n
		var qrBranchn string

		if o.Branchn != nil {
			qrBranchn = *o.Branchn
		}
		qBranchn := qrBranchn
		if qBranchn != "" {

			if err := r.SetQueryParam("branch__n", qBranchn); err != nil {
				return err
			}
		}
	}

	if o.BranchNic != nil {

		// query param branch__nic
		var qrBranchNic string

		if o.BranchNic != nil {
			qrBranchNic = *o.BranchNic
		}
		qBranchNic := qrBranchNic
		if qBranchNic != "" {

			if err := r.SetQueryParam("branch__nic", qBranchNic); err != nil {
				return err
			}
		}
	}

	if o.BranchNie != nil {

		// query param branch__nie
		var qrBranchNie string

		if o.BranchNie != nil {
			qrBranchNie = *o.BranchNie
		}
		qBranchNie := qrBranchNie
		if qBranchNie != "" {

			if err := r.SetQueryParam("branch__nie", qBranchNie); err != nil {
				return err
			}
		}
	}

	if o.BranchNiew != nil {

		// query param branch__niew
		var qrBranchNiew string

		if o.BranchNiew != nil {
			qrBranchNiew = *o.BranchNiew
		}
		qBranchNiew := qrBranchNiew
		if qBranchNiew != "" {

			if err := r.SetQueryParam("branch__niew", qBranchNiew); err != nil {
				return err
			}
		}
	}

	if o.BranchNisw != nil {

		// query param branch__nisw
		var qrBranchNisw string

		if o.BranchNisw != nil {
			qrBranchNisw = *o.BranchNisw
		}
		qBranchNisw := qrBranchNisw
		if qBranchNisw != "" {

			if err := r.SetQueryParam("branch__nisw", qBranchNisw); err != nil {
				return err
			}
		}
	}

	if o.Created != nil {

		// query param created
		var qrCreated string

		if o.Created != nil {
			qrCreated = *o.Created
		}
		qCreated := qrCreated
		if qCreated != "" {

			if err := r.SetQueryParam("created", qCreated); err != nil {
				return err
			}
		}
	}

	if o.CreatedGte != nil {

		// query param created__gte
		var qrCreatedGte string

		if o.CreatedGte != nil {
			qrCreatedGte = *o.CreatedGte
		}
		qCreatedGte := qrCreatedGte
		if qCreatedGte != "" {

			if err := r.SetQueryParam("created__gte", qCreatedGte); err != nil {
				return err
			}
		}
	}

	if o.CreatedLte != nil {

		// query param created__lte
		var qrCreatedLte string

		if o.CreatedLte != nil {
			qrCreatedLte = *o.CreatedLte
		}
		qCreatedLte := qrCreatedLte
		if qCreatedLte != "" {

			if err := r.SetQueryParam("created__lte", qCreatedLte); err != nil {
				return err
			}
		}
	}

	if o.ID != nil {

		// query param id
		var qrID string

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
	}

	if o.IDIc != nil {

		// query param id__ic
		var qrIDIc string

		if o.IDIc != nil {
			qrIDIc = *o.IDIc
		}
		qIDIc := qrIDIc
		if qIDIc != "" {

			if err := r.SetQueryParam("id__ic", qIDIc); err != nil {
				return err
			}
		}
	}

	if o.IDIe != nil {

		// query param id__ie
		var qrIDIe string

		if o.IDIe != nil {
			qrIDIe = *o.IDIe
		}
		qIDIe := qrIDIe
		if qIDIe != "" {

			if err := r.SetQueryParam("id__ie", qIDIe); err != nil {
				return err
			}
		}
	}

	if o.IDIew != nil {

		// query param id__iew
		var qrIDIew string

		if o.IDIew != nil {
			qrIDIew = *o.IDIew
		}
		qIDIew := qrIDIew
		if qIDIew != "" {

			if err := r.SetQueryParam("id__iew", qIDIew); err != nil {
				return err
			}
		}
	}

	if o.IDIsw != nil {

		// query param id__isw
		var qrIDIsw string

		if o.IDIsw != nil {
			qrIDIsw = *o.IDIsw
		}
		qIDIsw := qrIDIsw
		if qIDIsw != "" {

			if err := r.SetQueryParam("id__isw", qIDIsw); err != nil {
				return err
			}
		}
	}

	if o.IDn != nil {

		// query param id__n
		var qrIDn string

		if o.IDn != nil {
			qrIDn = *o.IDn
		}
		qIDn := qrIDn
		if qIDn != "" {

			if err := r.SetQueryParam("id__n", qIDn); err != nil {
				return err
			}
		}
	}

	if o.IDNic != nil {

		// query param id__nic
		var qrIDNic string

		if o.IDNic != nil {
			qrIDNic = *o.IDNic
		}
		qIDNic := qrIDNic
		if qIDNic != "" {

			if err := r.SetQueryParam("id__nic", qIDNic); err != nil {
				return err
			}
		}
	}

	if o.IDNie != nil {

		// query param id__nie
		var qrIDNie string

		if o.IDNie != nil {
			qrIDNie = *o.IDNie
		}
		qIDNie := qrIDNie
		if qIDNie != "" {

			if err := r.SetQueryParam("id__nie", qIDNie); err != nil {
				return err
			}
		}
	}

	if o.IDNiew != nil {

		// query param id__niew
		var qrIDNiew string

		if o.IDNiew != nil {
			qrIDNiew = *o.IDNiew
		}
		qIDNiew := qrIDNiew
		if qIDNiew != "" {

			if err := r.SetQueryParam("id__niew", qIDNiew); err != nil {
				return err
			}
		}
	}

	if o.IDNisw != nil {

		// query param id__nisw
		var qrIDNisw string

		if o.IDNisw != nil {
			qrIDNisw = *o.IDNisw
		}
		qIDNisw := qrIDNisw
		if qIDNisw != "" {

			if err := r.SetQueryParam("id__nisw", qIDNisw); err != nil {
				return err
			}
		}
	}

	if o.LastUpdated != nil {

		// query param last_updated
		var qrLastUpdated string

		if o.LastUpdated != nil {
			qrLastUpdated = *o.LastUpdated
		}
		qLastUpdated := qrLastUpdated
		if qLastUpdated != "" {

			if err := r.SetQueryParam("last_updated", qLastUpdated); err != nil {
				return err
			}
		}
	}

	if o.LastUpdatedGte != nil {

		// query param last_updated__gte
		var qrLastUpdatedGte string

		if o.LastUpdatedGte != nil {
			qrLastUpdatedGte = *o.LastUpdatedGte
		}
		qLastUpdatedGte := qrLastUpdatedGte
		if qLastUpdatedGte != "" {

			if err := r.SetQueryParam("last_updated__gte", qLastUpdatedGte); err != nil {
				return err
			}
		}
	}

	if o.LastUpdatedLte != nil {

		// query param last_updated__lte
		var qrLastUpdatedLte string

		if o.LastUpdatedLte != nil {
			qrLastUpdatedLte = *o.LastUpdatedLte
		}
		qLastUpdatedLte := qrLastUpdatedLte
		if qLastUpdatedLte != "" {

			if err := r.SetQueryParam("last_updated__lte", qLastUpdatedLte); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.NameIc != nil {

		// query param name__ic
		var qrNameIc string

		if o.NameIc != nil {
			qrNameIc = *o.NameIc
		}
		qNameIc := qrNameIc
		if qNameIc != "" {

			if err := r.SetQueryParam("name__ic", qNameIc); err != nil {
				return err
			}
		}
	}

	if o.NameIe != nil {

		// query param name__ie
		var qrNameIe string

		if o.NameIe != nil {
			qrNameIe = *o.NameIe
		}
		qNameIe := qrNameIe
		if qNameIe != "" {

			if err := r.SetQueryParam("name__ie", qNameIe); err != nil {
				return err
			}
		}
	}

	if o.NameIew != nil {

		// query param name__iew
		var qrNameIew string

		if o.NameIew != nil {
			qrNameIew = *o.NameIew
		}
		qNameIew := qrNameIew
		if qNameIew != "" {

			if err := r.SetQueryParam("name__iew", qNameIew); err != nil {
				return err
			}
		}
	}

	if o.NameIsw != nil {

		// query param name__isw
		var qrNameIsw string

		if o.NameIsw != nil {
			qrNameIsw = *o.NameIsw
		}
		qNameIsw := qrNameIsw
		if qNameIsw != "" {

			if err := r.SetQueryParam("name__isw", qNameIsw); err != nil {
				return err
			}
		}
	}

	if o.Namen != nil {

		// query param name__n
		var qrNamen string

		if o.Namen != nil {
			qrNamen = *o.Namen
		}
		qNamen := qrNamen
		if qNamen != "" {

			if err := r.SetQueryParam("name__n", qNamen); err != nil {
				return err
			}
		}
	}

	if o.NameNic != nil {

		// query param name__nic
		var qrNameNic string

		if o.NameNic != nil {
			qrNameNic = *o.NameNic
		}
		qNameNic := qrNameNic
		if qNameNic != "" {

			if err := r.SetQueryParam("name__nic", qNameNic); err != nil {
				return err
			}
		}
	}

	if o.NameNie != nil {

		// query param name__nie
		var qrNameNie string

		if o.NameNie != nil {
			qrNameNie = *o.NameNie
		}
		qNameNie := qrNameNie
		if qNameNie != "" {

			if err := r.SetQueryParam("name__nie", qNameNie); err != nil {
				return err
			}
		}
	}

	if o.NameNiew != nil {

		// query param name__niew
		var qrNameNiew string

		if o.NameNiew != nil {
			qrNameNiew = *o.NameNiew
		}
		qNameNiew := qrNameNiew
		if qNameNiew != "" {

			if err := r.SetQueryParam("name__niew", qNameNiew); err != nil {
				return err
			}
		}
	}

	if o.NameNisw != nil {

		// query param name__nisw
		var qrNameNisw string

		if o.NameNisw != nil {
			qrNameNisw = *o.NameNisw
		}
		qNameNisw := qrNameNisw
		if qNameNisw != "" {

			if err := r.SetQueryParam("name__nisw", qNameNisw); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}
	}

	if o.RemoteURL != nil {

		// query param remote_url
		var qrRemoteURL string

		if o.RemoteURL != nil {
			qrRemoteURL = *o.RemoteURL
		}
		qRemoteURL := qrRemoteURL
		if qRemoteURL != "" {

			if err := r.SetQueryParam("remote_url", qRemoteURL); err != nil {
				return err
			}
		}
	}

	if o.RemoteURLIc != nil {

		// query param remote_url__ic
		var qrRemoteURLIc string

		if o.RemoteURLIc != nil {
			qrRemoteURLIc = *o.RemoteURLIc
		}
		qRemoteURLIc := qrRemoteURLIc
		if qRemoteURLIc != "" {

			if err := r.SetQueryParam("remote_url__ic", qRemoteURLIc); err != nil {
				return err
			}
		}
	}

	if o.RemoteURLIe != nil {

		// query param remote_url__ie
		var qrRemoteURLIe string

		if o.RemoteURLIe != nil {
			qrRemoteURLIe = *o.RemoteURLIe
		}
		qRemoteURLIe := qrRemoteURLIe
		if qRemoteURLIe != "" {

			if err := r.SetQueryParam("remote_url__ie", qRemoteURLIe); err != nil {
				return err
			}
		}
	}

	if o.RemoteURLIew != nil {

		// query param remote_url__iew
		var qrRemoteURLIew string

		if o.RemoteURLIew != nil {
			qrRemoteURLIew = *o.RemoteURLIew
		}
		qRemoteURLIew := qrRemoteURLIew
		if qRemoteURLIew != "" {

			if err := r.SetQueryParam("remote_url__iew", qRemoteURLIew); err != nil {
				return err
			}
		}
	}

	if o.RemoteURLIsw != nil {

		// query param remote_url__isw
		var qrRemoteURLIsw string

		if o.RemoteURLIsw != nil {
			qrRemoteURLIsw = *o.RemoteURLIsw
		}
		qRemoteURLIsw := qrRemoteURLIsw
		if qRemoteURLIsw != "" {

			if err := r.SetQueryParam("remote_url__isw", qRemoteURLIsw); err != nil {
				return err
			}
		}
	}

	if o.RemoteURLn != nil {

		// query param remote_url__n
		var qrRemoteURLn string

		if o.RemoteURLn != nil {
			qrRemoteURLn = *o.RemoteURLn
		}
		qRemoteURLn := qrRemoteURLn
		if qRemoteURLn != "" {

			if err := r.SetQueryParam("remote_url__n", qRemoteURLn); err != nil {
				return err
			}
		}
	}

	if o.RemoteURLNic != nil {

		// query param remote_url__nic
		var qrRemoteURLNic string

		if o.RemoteURLNic != nil {
			qrRemoteURLNic = *o.RemoteURLNic
		}
		qRemoteURLNic := qrRemoteURLNic
		if qRemoteURLNic != "" {

			if err := r.SetQueryParam("remote_url__nic", qRemoteURLNic); err != nil {
				return err
			}
		}
	}

	if o.RemoteURLNie != nil {

		// query param remote_url__nie
		var qrRemoteURLNie string

		if o.RemoteURLNie != nil {
			qrRemoteURLNie = *o.RemoteURLNie
		}
		qRemoteURLNie := qrRemoteURLNie
		if qRemoteURLNie != "" {

			if err := r.SetQueryParam("remote_url__nie", qRemoteURLNie); err != nil {
				return err
			}
		}
	}

	if o.RemoteURLNiew != nil {

		// query param remote_url__niew
		var qrRemoteURLNiew string

		if o.RemoteURLNiew != nil {
			qrRemoteURLNiew = *o.RemoteURLNiew
		}
		qRemoteURLNiew := qrRemoteURLNiew
		if qRemoteURLNiew != "" {

			if err := r.SetQueryParam("remote_url__niew", qRemoteURLNiew); err != nil {
				return err
			}
		}
	}

	if o.RemoteURLNisw != nil {

		// query param remote_url__nisw
		var qrRemoteURLNisw string

		if o.RemoteURLNisw != nil {
			qrRemoteURLNisw = *o.RemoteURLNisw
		}
		qRemoteURLNisw := qrRemoteURLNisw
		if qRemoteURLNisw != "" {

			if err := r.SetQueryParam("remote_url__nisw", qRemoteURLNisw); err != nil {
				return err
			}
		}
	}

	if o.Slug != nil {

		// query param slug
		var qrSlug string

		if o.Slug != nil {
			qrSlug = *o.Slug
		}
		qSlug := qrSlug
		if qSlug != "" {

			if err := r.SetQueryParam("slug", qSlug); err != nil {
				return err
			}
		}
	}

	if o.SlugIc != nil {

		// query param slug__ic
		var qrSlugIc string

		if o.SlugIc != nil {
			qrSlugIc = *o.SlugIc
		}
		qSlugIc := qrSlugIc
		if qSlugIc != "" {

			if err := r.SetQueryParam("slug__ic", qSlugIc); err != nil {
				return err
			}
		}
	}

	if o.SlugIe != nil {

		// query param slug__ie
		var qrSlugIe string

		if o.SlugIe != nil {
			qrSlugIe = *o.SlugIe
		}
		qSlugIe := qrSlugIe
		if qSlugIe != "" {

			if err := r.SetQueryParam("slug__ie", qSlugIe); err != nil {
				return err
			}
		}
	}

	if o.SlugIew != nil {

		// query param slug__iew
		var qrSlugIew string

		if o.SlugIew != nil {
			qrSlugIew = *o.SlugIew
		}
		qSlugIew := qrSlugIew
		if qSlugIew != "" {

			if err := r.SetQueryParam("slug__iew", qSlugIew); err != nil {
				return err
			}
		}
	}

	if o.SlugIsw != nil {

		// query param slug__isw
		var qrSlugIsw string

		if o.SlugIsw != nil {
			qrSlugIsw = *o.SlugIsw
		}
		qSlugIsw := qrSlugIsw
		if qSlugIsw != "" {

			if err := r.SetQueryParam("slug__isw", qSlugIsw); err != nil {
				return err
			}
		}
	}

	if o.Slugn != nil {

		// query param slug__n
		var qrSlugn string

		if o.Slugn != nil {
			qrSlugn = *o.Slugn
		}
		qSlugn := qrSlugn
		if qSlugn != "" {

			if err := r.SetQueryParam("slug__n", qSlugn); err != nil {
				return err
			}
		}
	}

	if o.SlugNic != nil {

		// query param slug__nic
		var qrSlugNic string

		if o.SlugNic != nil {
			qrSlugNic = *o.SlugNic
		}
		qSlugNic := qrSlugNic
		if qSlugNic != "" {

			if err := r.SetQueryParam("slug__nic", qSlugNic); err != nil {
				return err
			}
		}
	}

	if o.SlugNie != nil {

		// query param slug__nie
		var qrSlugNie string

		if o.SlugNie != nil {
			qrSlugNie = *o.SlugNie
		}
		qSlugNie := qrSlugNie
		if qSlugNie != "" {

			if err := r.SetQueryParam("slug__nie", qSlugNie); err != nil {
				return err
			}
		}
	}

	if o.SlugNiew != nil {

		// query param slug__niew
		var qrSlugNiew string

		if o.SlugNiew != nil {
			qrSlugNiew = *o.SlugNiew
		}
		qSlugNiew := qrSlugNiew
		if qSlugNiew != "" {

			if err := r.SetQueryParam("slug__niew", qSlugNiew); err != nil {
				return err
			}
		}
	}

	if o.SlugNisw != nil {

		// query param slug__nisw
		var qrSlugNisw string

		if o.SlugNisw != nil {
			qrSlugNisw = *o.SlugNisw
		}
		qSlugNisw := qrSlugNisw
		if qSlugNisw != "" {

			if err := r.SetQueryParam("slug__nisw", qSlugNisw); err != nil {
				return err
			}
		}
	}

	if o.Tag != nil {

		// query param tag
		var qrTag string

		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {

			if err := r.SetQueryParam("tag", qTag); err != nil {
				return err
			}
		}
	}

	if o.Tagn != nil {

		// query param tag__n
		var qrTagn string

		if o.Tagn != nil {
			qrTagn = *o.Tagn
		}
		qTagn := qrTagn
		if qTagn != "" {

			if err := r.SetQueryParam("tag__n", qTagn); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
