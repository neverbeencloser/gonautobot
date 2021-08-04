package ipam

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

// NewIpamVlanGroupsListParams creates a new IpamVlanGroupsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamVlanGroupsListParams() *IpamVlanGroupsListParams {
	return &IpamVlanGroupsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamVlanGroupsListParamsWithTimeout creates a new IpamVlanGroupsListParams object
// with the ability to set a timeout on a request.
func NewIpamVlanGroupsListParamsWithTimeout(timeout time.Duration) *IpamVlanGroupsListParams {
	return &IpamVlanGroupsListParams{
		timeout: timeout,
	}
}

// NewIpamVlanGroupsListParamsWithContext creates a new IpamVlanGroupsListParams object
// with the ability to set a context for a request.
func NewIpamVlanGroupsListParamsWithContext(ctx context.Context) *IpamVlanGroupsListParams {
	return &IpamVlanGroupsListParams{
		Context: ctx,
	}
}

// NewIpamVlanGroupsListParamsWithHTTPClient creates a new IpamVlanGroupsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamVlanGroupsListParamsWithHTTPClient(client *http.Client) *IpamVlanGroupsListParams {
	return &IpamVlanGroupsListParams{
		HTTPClient: client,
	}
}

/* IpamVlanGroupsListParams contains all the parameters to send to the API endpoint
   for the ipam vlan groups list operation.

   Typically these are written to a http.Request.
*/
type IpamVlanGroupsListParams struct {

	// Created.
	Created *string

	// CreatedGte.
	CreatedGte *string

	// CreatedLte.
	CreatedLte *string

	// Description.
	Description *string

	// DescriptionIc.
	DescriptionIc *string

	// DescriptionIe.
	DescriptionIe *string

	// DescriptionIew.
	DescriptionIew *string

	// DescriptionIsw.
	DescriptionIsw *string

	// Descriptionn.
	Descriptionn *string

	// DescriptionNic.
	DescriptionNic *string

	// DescriptionNie.
	DescriptionNie *string

	// DescriptionNiew.
	DescriptionNiew *string

	// DescriptionNisw.
	DescriptionNisw *string

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

	// Region.
	Region *string

	// Regionn.
	Regionn *string

	// RegionID.
	RegionID *string

	// RegionIDn.
	RegionIDn *string

	// Site.
	Site *string

	// Siten.
	Siten *string

	// SiteID.
	SiteID *string

	// SiteIDn.
	SiteIDn *string

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam vlan groups list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVlanGroupsListParams) WithDefaults() *IpamVlanGroupsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam vlan groups list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVlanGroupsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithTimeout(timeout time.Duration) *IpamVlanGroupsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithContext(ctx context.Context) *IpamVlanGroupsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithHTTPClient(client *http.Client) *IpamVlanGroupsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreated adds the created to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithCreated(created *string) *IpamVlanGroupsListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGte adds the createdGte to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithCreatedGte(createdGte *string) *IpamVlanGroupsListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLte adds the createdLte to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithCreatedLte(createdLte *string) *IpamVlanGroupsListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithDescription adds the description to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithDescription(description *string) *IpamVlanGroupsListParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetDescription(description *string) {
	o.Description = description
}

// WithDescriptionIc adds the descriptionIc to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithDescriptionIc(descriptionIc *string) *IpamVlanGroupsListParams {
	o.SetDescriptionIc(descriptionIc)
	return o
}

// SetDescriptionIc adds the descriptionIc to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetDescriptionIc(descriptionIc *string) {
	o.DescriptionIc = descriptionIc
}

// WithDescriptionIe adds the descriptionIe to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithDescriptionIe(descriptionIe *string) *IpamVlanGroupsListParams {
	o.SetDescriptionIe(descriptionIe)
	return o
}

// SetDescriptionIe adds the descriptionIe to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetDescriptionIe(descriptionIe *string) {
	o.DescriptionIe = descriptionIe
}

// WithDescriptionIew adds the descriptionIew to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithDescriptionIew(descriptionIew *string) *IpamVlanGroupsListParams {
	o.SetDescriptionIew(descriptionIew)
	return o
}

// SetDescriptionIew adds the descriptionIew to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetDescriptionIew(descriptionIew *string) {
	o.DescriptionIew = descriptionIew
}

// WithDescriptionIsw adds the descriptionIsw to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithDescriptionIsw(descriptionIsw *string) *IpamVlanGroupsListParams {
	o.SetDescriptionIsw(descriptionIsw)
	return o
}

// SetDescriptionIsw adds the descriptionIsw to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetDescriptionIsw(descriptionIsw *string) {
	o.DescriptionIsw = descriptionIsw
}

// WithDescriptionn adds the descriptionn to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithDescriptionn(descriptionn *string) *IpamVlanGroupsListParams {
	o.SetDescriptionn(descriptionn)
	return o
}

// SetDescriptionn adds the descriptionN to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetDescriptionn(descriptionn *string) {
	o.Descriptionn = descriptionn
}

// WithDescriptionNic adds the descriptionNic to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithDescriptionNic(descriptionNic *string) *IpamVlanGroupsListParams {
	o.SetDescriptionNic(descriptionNic)
	return o
}

// SetDescriptionNic adds the descriptionNic to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetDescriptionNic(descriptionNic *string) {
	o.DescriptionNic = descriptionNic
}

// WithDescriptionNie adds the descriptionNie to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithDescriptionNie(descriptionNie *string) *IpamVlanGroupsListParams {
	o.SetDescriptionNie(descriptionNie)
	return o
}

// SetDescriptionNie adds the descriptionNie to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetDescriptionNie(descriptionNie *string) {
	o.DescriptionNie = descriptionNie
}

// WithDescriptionNiew adds the descriptionNiew to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithDescriptionNiew(descriptionNiew *string) *IpamVlanGroupsListParams {
	o.SetDescriptionNiew(descriptionNiew)
	return o
}

// SetDescriptionNiew adds the descriptionNiew to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetDescriptionNiew(descriptionNiew *string) {
	o.DescriptionNiew = descriptionNiew
}

// WithDescriptionNisw adds the descriptionNisw to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithDescriptionNisw(descriptionNisw *string) *IpamVlanGroupsListParams {
	o.SetDescriptionNisw(descriptionNisw)
	return o
}

// SetDescriptionNisw adds the descriptionNisw to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetDescriptionNisw(descriptionNisw *string) {
	o.DescriptionNisw = descriptionNisw
}

// WithID adds the id to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithID(id *string) *IpamVlanGroupsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithIDIc(iDIc *string) *IpamVlanGroupsListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithIDIe(iDIe *string) *IpamVlanGroupsListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithIDIew(iDIew *string) *IpamVlanGroupsListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithIDIsw(iDIsw *string) *IpamVlanGroupsListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithIDn(iDn *string) *IpamVlanGroupsListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithIDNic(iDNic *string) *IpamVlanGroupsListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithIDNie(iDNie *string) *IpamVlanGroupsListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithIDNiew(iDNiew *string) *IpamVlanGroupsListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithIDNisw(iDNisw *string) *IpamVlanGroupsListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithLastUpdated adds the lastUpdated to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithLastUpdated(lastUpdated *string) *IpamVlanGroupsListParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetLastUpdated(lastUpdated *string) {
	o.LastUpdated = lastUpdated
}

// WithLastUpdatedGte adds the lastUpdatedGte to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithLastUpdatedGte(lastUpdatedGte *string) *IpamVlanGroupsListParams {
	o.SetLastUpdatedGte(lastUpdatedGte)
	return o
}

// SetLastUpdatedGte adds the lastUpdatedGte to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetLastUpdatedGte(lastUpdatedGte *string) {
	o.LastUpdatedGte = lastUpdatedGte
}

// WithLastUpdatedLte adds the lastUpdatedLte to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithLastUpdatedLte(lastUpdatedLte *string) *IpamVlanGroupsListParams {
	o.SetLastUpdatedLte(lastUpdatedLte)
	return o
}

// SetLastUpdatedLte adds the lastUpdatedLte to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetLastUpdatedLte(lastUpdatedLte *string) {
	o.LastUpdatedLte = lastUpdatedLte
}

// WithLimit adds the limit to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithLimit(limit *int64) *IpamVlanGroupsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithName(name *string) *IpamVlanGroupsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithNameIc(nameIc *string) *IpamVlanGroupsListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithNameIe(nameIe *string) *IpamVlanGroupsListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithNameIew(nameIew *string) *IpamVlanGroupsListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithNameIsw(nameIsw *string) *IpamVlanGroupsListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithNamen(namen *string) *IpamVlanGroupsListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithNameNic(nameNic *string) *IpamVlanGroupsListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithNameNie(nameNie *string) *IpamVlanGroupsListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithNameNiew(nameNiew *string) *IpamVlanGroupsListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithNameNisw(nameNisw *string) *IpamVlanGroupsListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithOffset(offset *int64) *IpamVlanGroupsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithQ(q *string) *IpamVlanGroupsListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetQ(q *string) {
	o.Q = q
}

// WithRegion adds the region to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithRegion(region *string) *IpamVlanGroupsListParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetRegion(region *string) {
	o.Region = region
}

// WithRegionn adds the regionn to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithRegionn(regionn *string) *IpamVlanGroupsListParams {
	o.SetRegionn(regionn)
	return o
}

// SetRegionn adds the regionN to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetRegionn(regionn *string) {
	o.Regionn = regionn
}

// WithRegionID adds the regionID to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithRegionID(regionID *string) *IpamVlanGroupsListParams {
	o.SetRegionID(regionID)
	return o
}

// SetRegionID adds the regionId to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetRegionID(regionID *string) {
	o.RegionID = regionID
}

// WithRegionIDn adds the regionIDn to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithRegionIDn(regionIDn *string) *IpamVlanGroupsListParams {
	o.SetRegionIDn(regionIDn)
	return o
}

// SetRegionIDn adds the regionIdN to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetRegionIDn(regionIDn *string) {
	o.RegionIDn = regionIDn
}

// WithSite adds the site to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithSite(site *string) *IpamVlanGroupsListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetSite(site *string) {
	o.Site = site
}

// WithSiten adds the siten to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithSiten(siten *string) *IpamVlanGroupsListParams {
	o.SetSiten(siten)
	return o
}

// SetSiten adds the siteN to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetSiten(siten *string) {
	o.Siten = siten
}

// WithSiteID adds the siteID to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithSiteID(siteID *string) *IpamVlanGroupsListParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithSiteIDn adds the siteIDn to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithSiteIDn(siteIDn *string) *IpamVlanGroupsListParams {
	o.SetSiteIDn(siteIDn)
	return o
}

// SetSiteIDn adds the siteIdN to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetSiteIDn(siteIDn *string) {
	o.SiteIDn = siteIDn
}

// WithSlug adds the slug to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithSlug(slug *string) *IpamVlanGroupsListParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetSlug(slug *string) {
	o.Slug = slug
}

// WithSlugIc adds the slugIc to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithSlugIc(slugIc *string) *IpamVlanGroupsListParams {
	o.SetSlugIc(slugIc)
	return o
}

// SetSlugIc adds the slugIc to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetSlugIc(slugIc *string) {
	o.SlugIc = slugIc
}

// WithSlugIe adds the slugIe to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithSlugIe(slugIe *string) *IpamVlanGroupsListParams {
	o.SetSlugIe(slugIe)
	return o
}

// SetSlugIe adds the slugIe to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetSlugIe(slugIe *string) {
	o.SlugIe = slugIe
}

// WithSlugIew adds the slugIew to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithSlugIew(slugIew *string) *IpamVlanGroupsListParams {
	o.SetSlugIew(slugIew)
	return o
}

// SetSlugIew adds the slugIew to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetSlugIew(slugIew *string) {
	o.SlugIew = slugIew
}

// WithSlugIsw adds the slugIsw to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithSlugIsw(slugIsw *string) *IpamVlanGroupsListParams {
	o.SetSlugIsw(slugIsw)
	return o
}

// SetSlugIsw adds the slugIsw to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetSlugIsw(slugIsw *string) {
	o.SlugIsw = slugIsw
}

// WithSlugn adds the slugn to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithSlugn(slugn *string) *IpamVlanGroupsListParams {
	o.SetSlugn(slugn)
	return o
}

// SetSlugn adds the slugN to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetSlugn(slugn *string) {
	o.Slugn = slugn
}

// WithSlugNic adds the slugNic to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithSlugNic(slugNic *string) *IpamVlanGroupsListParams {
	o.SetSlugNic(slugNic)
	return o
}

// SetSlugNic adds the slugNic to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetSlugNic(slugNic *string) {
	o.SlugNic = slugNic
}

// WithSlugNie adds the slugNie to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithSlugNie(slugNie *string) *IpamVlanGroupsListParams {
	o.SetSlugNie(slugNie)
	return o
}

// SetSlugNie adds the slugNie to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetSlugNie(slugNie *string) {
	o.SlugNie = slugNie
}

// WithSlugNiew adds the slugNiew to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithSlugNiew(slugNiew *string) *IpamVlanGroupsListParams {
	o.SetSlugNiew(slugNiew)
	return o
}

// SetSlugNiew adds the slugNiew to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetSlugNiew(slugNiew *string) {
	o.SlugNiew = slugNiew
}

// WithSlugNisw adds the slugNisw to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) WithSlugNisw(slugNisw *string) *IpamVlanGroupsListParams {
	o.SetSlugNisw(slugNisw)
	return o
}

// SetSlugNisw adds the slugNisw to the ipam vlan groups list params
func (o *IpamVlanGroupsListParams) SetSlugNisw(slugNisw *string) {
	o.SlugNisw = slugNisw
}

// WriteToRequest writes these params to a swagger request
func (o *IpamVlanGroupsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.Description != nil {

		// query param description
		var qrDescription string

		if o.Description != nil {
			qrDescription = *o.Description
		}
		qDescription := qrDescription
		if qDescription != "" {

			if err := r.SetQueryParam("description", qDescription); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIc != nil {

		// query param description__ic
		var qrDescriptionIc string

		if o.DescriptionIc != nil {
			qrDescriptionIc = *o.DescriptionIc
		}
		qDescriptionIc := qrDescriptionIc
		if qDescriptionIc != "" {

			if err := r.SetQueryParam("description__ic", qDescriptionIc); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIe != nil {

		// query param description__ie
		var qrDescriptionIe string

		if o.DescriptionIe != nil {
			qrDescriptionIe = *o.DescriptionIe
		}
		qDescriptionIe := qrDescriptionIe
		if qDescriptionIe != "" {

			if err := r.SetQueryParam("description__ie", qDescriptionIe); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIew != nil {

		// query param description__iew
		var qrDescriptionIew string

		if o.DescriptionIew != nil {
			qrDescriptionIew = *o.DescriptionIew
		}
		qDescriptionIew := qrDescriptionIew
		if qDescriptionIew != "" {

			if err := r.SetQueryParam("description__iew", qDescriptionIew); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIsw != nil {

		// query param description__isw
		var qrDescriptionIsw string

		if o.DescriptionIsw != nil {
			qrDescriptionIsw = *o.DescriptionIsw
		}
		qDescriptionIsw := qrDescriptionIsw
		if qDescriptionIsw != "" {

			if err := r.SetQueryParam("description__isw", qDescriptionIsw); err != nil {
				return err
			}
		}
	}

	if o.Descriptionn != nil {

		// query param description__n
		var qrDescriptionn string

		if o.Descriptionn != nil {
			qrDescriptionn = *o.Descriptionn
		}
		qDescriptionn := qrDescriptionn
		if qDescriptionn != "" {

			if err := r.SetQueryParam("description__n", qDescriptionn); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNic != nil {

		// query param description__nic
		var qrDescriptionNic string

		if o.DescriptionNic != nil {
			qrDescriptionNic = *o.DescriptionNic
		}
		qDescriptionNic := qrDescriptionNic
		if qDescriptionNic != "" {

			if err := r.SetQueryParam("description__nic", qDescriptionNic); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNie != nil {

		// query param description__nie
		var qrDescriptionNie string

		if o.DescriptionNie != nil {
			qrDescriptionNie = *o.DescriptionNie
		}
		qDescriptionNie := qrDescriptionNie
		if qDescriptionNie != "" {

			if err := r.SetQueryParam("description__nie", qDescriptionNie); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNiew != nil {

		// query param description__niew
		var qrDescriptionNiew string

		if o.DescriptionNiew != nil {
			qrDescriptionNiew = *o.DescriptionNiew
		}
		qDescriptionNiew := qrDescriptionNiew
		if qDescriptionNiew != "" {

			if err := r.SetQueryParam("description__niew", qDescriptionNiew); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNisw != nil {

		// query param description__nisw
		var qrDescriptionNisw string

		if o.DescriptionNisw != nil {
			qrDescriptionNisw = *o.DescriptionNisw
		}
		qDescriptionNisw := qrDescriptionNisw
		if qDescriptionNisw != "" {

			if err := r.SetQueryParam("description__nisw", qDescriptionNisw); err != nil {
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

	if o.Region != nil {

		// query param region
		var qrRegion string

		if o.Region != nil {
			qrRegion = *o.Region
		}
		qRegion := qrRegion
		if qRegion != "" {

			if err := r.SetQueryParam("region", qRegion); err != nil {
				return err
			}
		}
	}

	if o.Regionn != nil {

		// query param region__n
		var qrRegionn string

		if o.Regionn != nil {
			qrRegionn = *o.Regionn
		}
		qRegionn := qrRegionn
		if qRegionn != "" {

			if err := r.SetQueryParam("region__n", qRegionn); err != nil {
				return err
			}
		}
	}

	if o.RegionID != nil {

		// query param region_id
		var qrRegionID string

		if o.RegionID != nil {
			qrRegionID = *o.RegionID
		}
		qRegionID := qrRegionID
		if qRegionID != "" {

			if err := r.SetQueryParam("region_id", qRegionID); err != nil {
				return err
			}
		}
	}

	if o.RegionIDn != nil {

		// query param region_id__n
		var qrRegionIDn string

		if o.RegionIDn != nil {
			qrRegionIDn = *o.RegionIDn
		}
		qRegionIDn := qrRegionIDn
		if qRegionIDn != "" {

			if err := r.SetQueryParam("region_id__n", qRegionIDn); err != nil {
				return err
			}
		}
	}

	if o.Site != nil {

		// query param site
		var qrSite string

		if o.Site != nil {
			qrSite = *o.Site
		}
		qSite := qrSite
		if qSite != "" {

			if err := r.SetQueryParam("site", qSite); err != nil {
				return err
			}
		}
	}

	if o.Siten != nil {

		// query param site__n
		var qrSiten string

		if o.Siten != nil {
			qrSiten = *o.Siten
		}
		qSiten := qrSiten
		if qSiten != "" {

			if err := r.SetQueryParam("site__n", qSiten); err != nil {
				return err
			}
		}
	}

	if o.SiteID != nil {

		// query param site_id
		var qrSiteID string

		if o.SiteID != nil {
			qrSiteID = *o.SiteID
		}
		qSiteID := qrSiteID
		if qSiteID != "" {

			if err := r.SetQueryParam("site_id", qSiteID); err != nil {
				return err
			}
		}
	}

	if o.SiteIDn != nil {

		// query param site_id__n
		var qrSiteIDn string

		if o.SiteIDn != nil {
			qrSiteIDn = *o.SiteIDn
		}
		qSiteIDn := qrSiteIDn
		if qSiteIDn != "" {

			if err := r.SetQueryParam("site_id__n", qSiteIDn); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
