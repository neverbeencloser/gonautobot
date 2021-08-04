package dcim

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

// NewDcimVirtualChassisListParams creates a new DcimVirtualChassisListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimVirtualChassisListParams() *DcimVirtualChassisListParams {
	return &DcimVirtualChassisListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimVirtualChassisListParamsWithTimeout creates a new DcimVirtualChassisListParams object
// with the ability to set a timeout on a request.
func NewDcimVirtualChassisListParamsWithTimeout(timeout time.Duration) *DcimVirtualChassisListParams {
	return &DcimVirtualChassisListParams{
		timeout: timeout,
	}
}

// NewDcimVirtualChassisListParamsWithContext creates a new DcimVirtualChassisListParams object
// with the ability to set a context for a request.
func NewDcimVirtualChassisListParamsWithContext(ctx context.Context) *DcimVirtualChassisListParams {
	return &DcimVirtualChassisListParams{
		Context: ctx,
	}
}

// NewDcimVirtualChassisListParamsWithHTTPClient creates a new DcimVirtualChassisListParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimVirtualChassisListParamsWithHTTPClient(client *http.Client) *DcimVirtualChassisListParams {
	return &DcimVirtualChassisListParams{
		HTTPClient: client,
	}
}

/* DcimVirtualChassisListParams contains all the parameters to send to the API endpoint
   for the dcim virtual chassis list operation.

   Typically these are written to a http.Request.
*/
type DcimVirtualChassisListParams struct {

	// Domain.
	Domain *string

	// DomainIc.
	DomainIc *string

	// DomainIe.
	DomainIe *string

	// DomainIew.
	DomainIew *string

	// DomainIsw.
	DomainIsw *string

	// Domainn.
	Domainn *string

	// DomainNic.
	DomainNic *string

	// DomainNie.
	DomainNie *string

	// DomainNiew.
	DomainNiew *string

	// DomainNisw.
	DomainNisw *string

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

	/* Limit.

	   Number of results to return per page.
	*/
	Limit *int64

	// Master.
	Master *string

	// Mastern.
	Mastern *string

	// MasterID.
	MasterID *string

	// MasterIDn.
	MasterIDn *string

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

	// Tag.
	Tag *string

	// Tagn.
	Tagn *string

	// Tenant.
	Tenant *string

	// Tenantn.
	Tenantn *string

	// TenantID.
	TenantID *string

	// TenantIDn.
	TenantIDn *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim virtual chassis list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimVirtualChassisListParams) WithDefaults() *DcimVirtualChassisListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim virtual chassis list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimVirtualChassisListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithTimeout(timeout time.Duration) *DcimVirtualChassisListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithContext(ctx context.Context) *DcimVirtualChassisListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithHTTPClient(client *http.Client) *DcimVirtualChassisListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomain adds the domain to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithDomain(domain *string) *DcimVirtualChassisListParams {
	o.SetDomain(domain)
	return o
}

// SetDomain adds the domain to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetDomain(domain *string) {
	o.Domain = domain
}

// WithDomainIc adds the domainIc to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithDomainIc(domainIc *string) *DcimVirtualChassisListParams {
	o.SetDomainIc(domainIc)
	return o
}

// SetDomainIc adds the domainIc to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetDomainIc(domainIc *string) {
	o.DomainIc = domainIc
}

// WithDomainIe adds the domainIe to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithDomainIe(domainIe *string) *DcimVirtualChassisListParams {
	o.SetDomainIe(domainIe)
	return o
}

// SetDomainIe adds the domainIe to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetDomainIe(domainIe *string) {
	o.DomainIe = domainIe
}

// WithDomainIew adds the domainIew to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithDomainIew(domainIew *string) *DcimVirtualChassisListParams {
	o.SetDomainIew(domainIew)
	return o
}

// SetDomainIew adds the domainIew to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetDomainIew(domainIew *string) {
	o.DomainIew = domainIew
}

// WithDomainIsw adds the domainIsw to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithDomainIsw(domainIsw *string) *DcimVirtualChassisListParams {
	o.SetDomainIsw(domainIsw)
	return o
}

// SetDomainIsw adds the domainIsw to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetDomainIsw(domainIsw *string) {
	o.DomainIsw = domainIsw
}

// WithDomainn adds the domainn to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithDomainn(domainn *string) *DcimVirtualChassisListParams {
	o.SetDomainn(domainn)
	return o
}

// SetDomainn adds the domainN to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetDomainn(domainn *string) {
	o.Domainn = domainn
}

// WithDomainNic adds the domainNic to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithDomainNic(domainNic *string) *DcimVirtualChassisListParams {
	o.SetDomainNic(domainNic)
	return o
}

// SetDomainNic adds the domainNic to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetDomainNic(domainNic *string) {
	o.DomainNic = domainNic
}

// WithDomainNie adds the domainNie to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithDomainNie(domainNie *string) *DcimVirtualChassisListParams {
	o.SetDomainNie(domainNie)
	return o
}

// SetDomainNie adds the domainNie to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetDomainNie(domainNie *string) {
	o.DomainNie = domainNie
}

// WithDomainNiew adds the domainNiew to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithDomainNiew(domainNiew *string) *DcimVirtualChassisListParams {
	o.SetDomainNiew(domainNiew)
	return o
}

// SetDomainNiew adds the domainNiew to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetDomainNiew(domainNiew *string) {
	o.DomainNiew = domainNiew
}

// WithDomainNisw adds the domainNisw to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithDomainNisw(domainNisw *string) *DcimVirtualChassisListParams {
	o.SetDomainNisw(domainNisw)
	return o
}

// SetDomainNisw adds the domainNisw to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetDomainNisw(domainNisw *string) {
	o.DomainNisw = domainNisw
}

// WithID adds the id to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithID(id *string) *DcimVirtualChassisListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithIDIc(iDIc *string) *DcimVirtualChassisListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithIDIe(iDIe *string) *DcimVirtualChassisListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithIDIew(iDIew *string) *DcimVirtualChassisListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithIDIsw(iDIsw *string) *DcimVirtualChassisListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithIDn(iDn *string) *DcimVirtualChassisListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithIDNic(iDNic *string) *DcimVirtualChassisListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithIDNie(iDNie *string) *DcimVirtualChassisListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithIDNiew(iDNiew *string) *DcimVirtualChassisListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithIDNisw(iDNisw *string) *DcimVirtualChassisListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithLimit adds the limit to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithLimit(limit *int64) *DcimVirtualChassisListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithMaster adds the master to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithMaster(master *string) *DcimVirtualChassisListParams {
	o.SetMaster(master)
	return o
}

// SetMaster adds the master to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetMaster(master *string) {
	o.Master = master
}

// WithMastern adds the mastern to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithMastern(mastern *string) *DcimVirtualChassisListParams {
	o.SetMastern(mastern)
	return o
}

// SetMastern adds the masterN to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetMastern(mastern *string) {
	o.Mastern = mastern
}

// WithMasterID adds the masterID to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithMasterID(masterID *string) *DcimVirtualChassisListParams {
	o.SetMasterID(masterID)
	return o
}

// SetMasterID adds the masterId to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetMasterID(masterID *string) {
	o.MasterID = masterID
}

// WithMasterIDn adds the masterIDn to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithMasterIDn(masterIDn *string) *DcimVirtualChassisListParams {
	o.SetMasterIDn(masterIDn)
	return o
}

// SetMasterIDn adds the masterIdN to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetMasterIDn(masterIDn *string) {
	o.MasterIDn = masterIDn
}

// WithName adds the name to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithName(name *string) *DcimVirtualChassisListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithNameIc(nameIc *string) *DcimVirtualChassisListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithNameIe(nameIe *string) *DcimVirtualChassisListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithNameIew(nameIew *string) *DcimVirtualChassisListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithNameIsw(nameIsw *string) *DcimVirtualChassisListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithNamen(namen *string) *DcimVirtualChassisListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithNameNic(nameNic *string) *DcimVirtualChassisListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithNameNie(nameNie *string) *DcimVirtualChassisListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithNameNiew(nameNiew *string) *DcimVirtualChassisListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithNameNisw(nameNisw *string) *DcimVirtualChassisListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithOffset(offset *int64) *DcimVirtualChassisListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithQ(q *string) *DcimVirtualChassisListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetQ(q *string) {
	o.Q = q
}

// WithRegion adds the region to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithRegion(region *string) *DcimVirtualChassisListParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetRegion(region *string) {
	o.Region = region
}

// WithRegionn adds the regionn to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithRegionn(regionn *string) *DcimVirtualChassisListParams {
	o.SetRegionn(regionn)
	return o
}

// SetRegionn adds the regionN to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetRegionn(regionn *string) {
	o.Regionn = regionn
}

// WithRegionID adds the regionID to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithRegionID(regionID *string) *DcimVirtualChassisListParams {
	o.SetRegionID(regionID)
	return o
}

// SetRegionID adds the regionId to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetRegionID(regionID *string) {
	o.RegionID = regionID
}

// WithRegionIDn adds the regionIDn to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithRegionIDn(regionIDn *string) *DcimVirtualChassisListParams {
	o.SetRegionIDn(regionIDn)
	return o
}

// SetRegionIDn adds the regionIdN to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetRegionIDn(regionIDn *string) {
	o.RegionIDn = regionIDn
}

// WithSite adds the site to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithSite(site *string) *DcimVirtualChassisListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetSite(site *string) {
	o.Site = site
}

// WithSiten adds the siten to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithSiten(siten *string) *DcimVirtualChassisListParams {
	o.SetSiten(siten)
	return o
}

// SetSiten adds the siteN to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetSiten(siten *string) {
	o.Siten = siten
}

// WithSiteID adds the siteID to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithSiteID(siteID *string) *DcimVirtualChassisListParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithSiteIDn adds the siteIDn to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithSiteIDn(siteIDn *string) *DcimVirtualChassisListParams {
	o.SetSiteIDn(siteIDn)
	return o
}

// SetSiteIDn adds the siteIdN to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetSiteIDn(siteIDn *string) {
	o.SiteIDn = siteIDn
}

// WithTag adds the tag to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithTag(tag *string) *DcimVirtualChassisListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTagn adds the tagn to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithTagn(tagn *string) *DcimVirtualChassisListParams {
	o.SetTagn(tagn)
	return o
}

// SetTagn adds the tagN to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetTagn(tagn *string) {
	o.Tagn = tagn
}

// WithTenant adds the tenant to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithTenant(tenant *string) *DcimVirtualChassisListParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetTenant(tenant *string) {
	o.Tenant = tenant
}

// WithTenantn adds the tenantn to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithTenantn(tenantn *string) *DcimVirtualChassisListParams {
	o.SetTenantn(tenantn)
	return o
}

// SetTenantn adds the tenantN to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetTenantn(tenantn *string) {
	o.Tenantn = tenantn
}

// WithTenantID adds the tenantID to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithTenantID(tenantID *string) *DcimVirtualChassisListParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WithTenantIDn adds the tenantIDn to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) WithTenantIDn(tenantIDn *string) *DcimVirtualChassisListParams {
	o.SetTenantIDn(tenantIDn)
	return o
}

// SetTenantIDn adds the tenantIdN to the dcim virtual chassis list params
func (o *DcimVirtualChassisListParams) SetTenantIDn(tenantIDn *string) {
	o.TenantIDn = tenantIDn
}

// WriteToRequest writes these params to a swagger request
func (o *DcimVirtualChassisListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Domain != nil {

		// query param domain
		var qrDomain string

		if o.Domain != nil {
			qrDomain = *o.Domain
		}
		qDomain := qrDomain
		if qDomain != "" {

			if err := r.SetQueryParam("domain", qDomain); err != nil {
				return err
			}
		}
	}

	if o.DomainIc != nil {

		// query param domain__ic
		var qrDomainIc string

		if o.DomainIc != nil {
			qrDomainIc = *o.DomainIc
		}
		qDomainIc := qrDomainIc
		if qDomainIc != "" {

			if err := r.SetQueryParam("domain__ic", qDomainIc); err != nil {
				return err
			}
		}
	}

	if o.DomainIe != nil {

		// query param domain__ie
		var qrDomainIe string

		if o.DomainIe != nil {
			qrDomainIe = *o.DomainIe
		}
		qDomainIe := qrDomainIe
		if qDomainIe != "" {

			if err := r.SetQueryParam("domain__ie", qDomainIe); err != nil {
				return err
			}
		}
	}

	if o.DomainIew != nil {

		// query param domain__iew
		var qrDomainIew string

		if o.DomainIew != nil {
			qrDomainIew = *o.DomainIew
		}
		qDomainIew := qrDomainIew
		if qDomainIew != "" {

			if err := r.SetQueryParam("domain__iew", qDomainIew); err != nil {
				return err
			}
		}
	}

	if o.DomainIsw != nil {

		// query param domain__isw
		var qrDomainIsw string

		if o.DomainIsw != nil {
			qrDomainIsw = *o.DomainIsw
		}
		qDomainIsw := qrDomainIsw
		if qDomainIsw != "" {

			if err := r.SetQueryParam("domain__isw", qDomainIsw); err != nil {
				return err
			}
		}
	}

	if o.Domainn != nil {

		// query param domain__n
		var qrDomainn string

		if o.Domainn != nil {
			qrDomainn = *o.Domainn
		}
		qDomainn := qrDomainn
		if qDomainn != "" {

			if err := r.SetQueryParam("domain__n", qDomainn); err != nil {
				return err
			}
		}
	}

	if o.DomainNic != nil {

		// query param domain__nic
		var qrDomainNic string

		if o.DomainNic != nil {
			qrDomainNic = *o.DomainNic
		}
		qDomainNic := qrDomainNic
		if qDomainNic != "" {

			if err := r.SetQueryParam("domain__nic", qDomainNic); err != nil {
				return err
			}
		}
	}

	if o.DomainNie != nil {

		// query param domain__nie
		var qrDomainNie string

		if o.DomainNie != nil {
			qrDomainNie = *o.DomainNie
		}
		qDomainNie := qrDomainNie
		if qDomainNie != "" {

			if err := r.SetQueryParam("domain__nie", qDomainNie); err != nil {
				return err
			}
		}
	}

	if o.DomainNiew != nil {

		// query param domain__niew
		var qrDomainNiew string

		if o.DomainNiew != nil {
			qrDomainNiew = *o.DomainNiew
		}
		qDomainNiew := qrDomainNiew
		if qDomainNiew != "" {

			if err := r.SetQueryParam("domain__niew", qDomainNiew); err != nil {
				return err
			}
		}
	}

	if o.DomainNisw != nil {

		// query param domain__nisw
		var qrDomainNisw string

		if o.DomainNisw != nil {
			qrDomainNisw = *o.DomainNisw
		}
		qDomainNisw := qrDomainNisw
		if qDomainNisw != "" {

			if err := r.SetQueryParam("domain__nisw", qDomainNisw); err != nil {
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

	if o.Master != nil {

		// query param master
		var qrMaster string

		if o.Master != nil {
			qrMaster = *o.Master
		}
		qMaster := qrMaster
		if qMaster != "" {

			if err := r.SetQueryParam("master", qMaster); err != nil {
				return err
			}
		}
	}

	if o.Mastern != nil {

		// query param master__n
		var qrMastern string

		if o.Mastern != nil {
			qrMastern = *o.Mastern
		}
		qMastern := qrMastern
		if qMastern != "" {

			if err := r.SetQueryParam("master__n", qMastern); err != nil {
				return err
			}
		}
	}

	if o.MasterID != nil {

		// query param master_id
		var qrMasterID string

		if o.MasterID != nil {
			qrMasterID = *o.MasterID
		}
		qMasterID := qrMasterID
		if qMasterID != "" {

			if err := r.SetQueryParam("master_id", qMasterID); err != nil {
				return err
			}
		}
	}

	if o.MasterIDn != nil {

		// query param master_id__n
		var qrMasterIDn string

		if o.MasterIDn != nil {
			qrMasterIDn = *o.MasterIDn
		}
		qMasterIDn := qrMasterIDn
		if qMasterIDn != "" {

			if err := r.SetQueryParam("master_id__n", qMasterIDn); err != nil {
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

	if o.Tenant != nil {

		// query param tenant
		var qrTenant string

		if o.Tenant != nil {
			qrTenant = *o.Tenant
		}
		qTenant := qrTenant
		if qTenant != "" {

			if err := r.SetQueryParam("tenant", qTenant); err != nil {
				return err
			}
		}
	}

	if o.Tenantn != nil {

		// query param tenant__n
		var qrTenantn string

		if o.Tenantn != nil {
			qrTenantn = *o.Tenantn
		}
		qTenantn := qrTenantn
		if qTenantn != "" {

			if err := r.SetQueryParam("tenant__n", qTenantn); err != nil {
				return err
			}
		}
	}

	if o.TenantID != nil {

		// query param tenant_id
		var qrTenantID string

		if o.TenantID != nil {
			qrTenantID = *o.TenantID
		}
		qTenantID := qrTenantID
		if qTenantID != "" {

			if err := r.SetQueryParam("tenant_id", qTenantID); err != nil {
				return err
			}
		}
	}

	if o.TenantIDn != nil {

		// query param tenant_id__n
		var qrTenantIDn string

		if o.TenantIDn != nil {
			qrTenantIDn = *o.TenantIDn
		}
		qTenantIDn := qrTenantIDn
		if qTenantIDn != "" {

			if err := r.SetQueryParam("tenant_id__n", qTenantIDn); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
