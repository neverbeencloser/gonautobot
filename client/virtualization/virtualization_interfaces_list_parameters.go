package virtualization

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

// NewVirtualizationInterfacesListParams creates a new VirtualizationInterfacesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVirtualizationInterfacesListParams() *VirtualizationInterfacesListParams {
	return &VirtualizationInterfacesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationInterfacesListParamsWithTimeout creates a new VirtualizationInterfacesListParams object
// with the ability to set a timeout on a request.
func NewVirtualizationInterfacesListParamsWithTimeout(timeout time.Duration) *VirtualizationInterfacesListParams {
	return &VirtualizationInterfacesListParams{
		timeout: timeout,
	}
}

// NewVirtualizationInterfacesListParamsWithContext creates a new VirtualizationInterfacesListParams object
// with the ability to set a context for a request.
func NewVirtualizationInterfacesListParamsWithContext(ctx context.Context) *VirtualizationInterfacesListParams {
	return &VirtualizationInterfacesListParams{
		Context: ctx,
	}
}

// NewVirtualizationInterfacesListParamsWithHTTPClient creates a new VirtualizationInterfacesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewVirtualizationInterfacesListParamsWithHTTPClient(client *http.Client) *VirtualizationInterfacesListParams {
	return &VirtualizationInterfacesListParams{
		HTTPClient: client,
	}
}

/* VirtualizationInterfacesListParams contains all the parameters to send to the API endpoint
   for the virtualization interfaces list operation.

   Typically these are written to a http.Request.
*/
type VirtualizationInterfacesListParams struct {

	// Cluster.
	Cluster *string

	// Clustern.
	Clustern *string

	// ClusterID.
	ClusterID *string

	// ClusterIDn.
	ClusterIDn *string

	// Enabled.
	Enabled *string

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

	// MacAddress.
	MacAddress *string

	// MacAddressIc.
	MacAddressIc *string

	// MacAddressIe.
	MacAddressIe *string

	// MacAddressIew.
	MacAddressIew *string

	// MacAddressIsw.
	MacAddressIsw *string

	// MacAddressn.
	MacAddressn *string

	// MacAddressNic.
	MacAddressNic *string

	// MacAddressNie.
	MacAddressNie *string

	// MacAddressNiew.
	MacAddressNiew *string

	// MacAddressNisw.
	MacAddressNisw *string

	// Mtu.
	Mtu *string

	// MtuGt.
	MtuGt *string

	// MtuGte.
	MtuGte *string

	// MtuLt.
	MtuLt *string

	// MtuLte.
	MtuLte *string

	// Mtun.
	Mtun *string

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

	// Tag.
	Tag *string

	// Tagn.
	Tagn *string

	// VirtualMachine.
	VirtualMachine *string

	// VirtualMachinen.
	VirtualMachinen *string

	// VirtualMachineID.
	VirtualMachineID *string

	// VirtualMachineIDn.
	VirtualMachineIDn *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the virtualization interfaces list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationInterfacesListParams) WithDefaults() *VirtualizationInterfacesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the virtualization interfaces list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationInterfacesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithTimeout(timeout time.Duration) *VirtualizationInterfacesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithContext(ctx context.Context) *VirtualizationInterfacesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithHTTPClient(client *http.Client) *VirtualizationInterfacesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCluster adds the cluster to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithCluster(cluster *string) *VirtualizationInterfacesListParams {
	o.SetCluster(cluster)
	return o
}

// SetCluster adds the cluster to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetCluster(cluster *string) {
	o.Cluster = cluster
}

// WithClustern adds the clustern to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithClustern(clustern *string) *VirtualizationInterfacesListParams {
	o.SetClustern(clustern)
	return o
}

// SetClustern adds the clusterN to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetClustern(clustern *string) {
	o.Clustern = clustern
}

// WithClusterID adds the clusterID to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithClusterID(clusterID *string) *VirtualizationInterfacesListParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetClusterID(clusterID *string) {
	o.ClusterID = clusterID
}

// WithClusterIDn adds the clusterIDn to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithClusterIDn(clusterIDn *string) *VirtualizationInterfacesListParams {
	o.SetClusterIDn(clusterIDn)
	return o
}

// SetClusterIDn adds the clusterIdN to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetClusterIDn(clusterIDn *string) {
	o.ClusterIDn = clusterIDn
}

// WithEnabled adds the enabled to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithEnabled(enabled *string) *VirtualizationInterfacesListParams {
	o.SetEnabled(enabled)
	return o
}

// SetEnabled adds the enabled to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetEnabled(enabled *string) {
	o.Enabled = enabled
}

// WithID adds the id to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithID(id *string) *VirtualizationInterfacesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithIDIc(iDIc *string) *VirtualizationInterfacesListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithIDIe(iDIe *string) *VirtualizationInterfacesListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithIDIew(iDIew *string) *VirtualizationInterfacesListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithIDIsw(iDIsw *string) *VirtualizationInterfacesListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithIDn(iDn *string) *VirtualizationInterfacesListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithIDNic(iDNic *string) *VirtualizationInterfacesListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithIDNie(iDNie *string) *VirtualizationInterfacesListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithIDNiew(iDNiew *string) *VirtualizationInterfacesListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithIDNisw(iDNisw *string) *VirtualizationInterfacesListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithLimit adds the limit to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithLimit(limit *int64) *VirtualizationInterfacesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithMacAddress adds the macAddress to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithMacAddress(macAddress *string) *VirtualizationInterfacesListParams {
	o.SetMacAddress(macAddress)
	return o
}

// SetMacAddress adds the macAddress to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetMacAddress(macAddress *string) {
	o.MacAddress = macAddress
}

// WithMacAddressIc adds the macAddressIc to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithMacAddressIc(macAddressIc *string) *VirtualizationInterfacesListParams {
	o.SetMacAddressIc(macAddressIc)
	return o
}

// SetMacAddressIc adds the macAddressIc to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetMacAddressIc(macAddressIc *string) {
	o.MacAddressIc = macAddressIc
}

// WithMacAddressIe adds the macAddressIe to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithMacAddressIe(macAddressIe *string) *VirtualizationInterfacesListParams {
	o.SetMacAddressIe(macAddressIe)
	return o
}

// SetMacAddressIe adds the macAddressIe to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetMacAddressIe(macAddressIe *string) {
	o.MacAddressIe = macAddressIe
}

// WithMacAddressIew adds the macAddressIew to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithMacAddressIew(macAddressIew *string) *VirtualizationInterfacesListParams {
	o.SetMacAddressIew(macAddressIew)
	return o
}

// SetMacAddressIew adds the macAddressIew to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetMacAddressIew(macAddressIew *string) {
	o.MacAddressIew = macAddressIew
}

// WithMacAddressIsw adds the macAddressIsw to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithMacAddressIsw(macAddressIsw *string) *VirtualizationInterfacesListParams {
	o.SetMacAddressIsw(macAddressIsw)
	return o
}

// SetMacAddressIsw adds the macAddressIsw to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetMacAddressIsw(macAddressIsw *string) {
	o.MacAddressIsw = macAddressIsw
}

// WithMacAddressn adds the macAddressn to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithMacAddressn(macAddressn *string) *VirtualizationInterfacesListParams {
	o.SetMacAddressn(macAddressn)
	return o
}

// SetMacAddressn adds the macAddressN to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetMacAddressn(macAddressn *string) {
	o.MacAddressn = macAddressn
}

// WithMacAddressNic adds the macAddressNic to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithMacAddressNic(macAddressNic *string) *VirtualizationInterfacesListParams {
	o.SetMacAddressNic(macAddressNic)
	return o
}

// SetMacAddressNic adds the macAddressNic to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetMacAddressNic(macAddressNic *string) {
	o.MacAddressNic = macAddressNic
}

// WithMacAddressNie adds the macAddressNie to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithMacAddressNie(macAddressNie *string) *VirtualizationInterfacesListParams {
	o.SetMacAddressNie(macAddressNie)
	return o
}

// SetMacAddressNie adds the macAddressNie to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetMacAddressNie(macAddressNie *string) {
	o.MacAddressNie = macAddressNie
}

// WithMacAddressNiew adds the macAddressNiew to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithMacAddressNiew(macAddressNiew *string) *VirtualizationInterfacesListParams {
	o.SetMacAddressNiew(macAddressNiew)
	return o
}

// SetMacAddressNiew adds the macAddressNiew to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetMacAddressNiew(macAddressNiew *string) {
	o.MacAddressNiew = macAddressNiew
}

// WithMacAddressNisw adds the macAddressNisw to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithMacAddressNisw(macAddressNisw *string) *VirtualizationInterfacesListParams {
	o.SetMacAddressNisw(macAddressNisw)
	return o
}

// SetMacAddressNisw adds the macAddressNisw to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetMacAddressNisw(macAddressNisw *string) {
	o.MacAddressNisw = macAddressNisw
}

// WithMtu adds the mtu to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithMtu(mtu *string) *VirtualizationInterfacesListParams {
	o.SetMtu(mtu)
	return o
}

// SetMtu adds the mtu to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetMtu(mtu *string) {
	o.Mtu = mtu
}

// WithMtuGt adds the mtuGt to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithMtuGt(mtuGt *string) *VirtualizationInterfacesListParams {
	o.SetMtuGt(mtuGt)
	return o
}

// SetMtuGt adds the mtuGt to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetMtuGt(mtuGt *string) {
	o.MtuGt = mtuGt
}

// WithMtuGte adds the mtuGte to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithMtuGte(mtuGte *string) *VirtualizationInterfacesListParams {
	o.SetMtuGte(mtuGte)
	return o
}

// SetMtuGte adds the mtuGte to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetMtuGte(mtuGte *string) {
	o.MtuGte = mtuGte
}

// WithMtuLt adds the mtuLt to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithMtuLt(mtuLt *string) *VirtualizationInterfacesListParams {
	o.SetMtuLt(mtuLt)
	return o
}

// SetMtuLt adds the mtuLt to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetMtuLt(mtuLt *string) {
	o.MtuLt = mtuLt
}

// WithMtuLte adds the mtuLte to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithMtuLte(mtuLte *string) *VirtualizationInterfacesListParams {
	o.SetMtuLte(mtuLte)
	return o
}

// SetMtuLte adds the mtuLte to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetMtuLte(mtuLte *string) {
	o.MtuLte = mtuLte
}

// WithMtun adds the mtun to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithMtun(mtun *string) *VirtualizationInterfacesListParams {
	o.SetMtun(mtun)
	return o
}

// SetMtun adds the mtuN to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetMtun(mtun *string) {
	o.Mtun = mtun
}

// WithName adds the name to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithName(name *string) *VirtualizationInterfacesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithNameIc(nameIc *string) *VirtualizationInterfacesListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithNameIe(nameIe *string) *VirtualizationInterfacesListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithNameIew(nameIew *string) *VirtualizationInterfacesListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithNameIsw(nameIsw *string) *VirtualizationInterfacesListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithNamen(namen *string) *VirtualizationInterfacesListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithNameNic(nameNic *string) *VirtualizationInterfacesListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithNameNie(nameNie *string) *VirtualizationInterfacesListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithNameNiew(nameNiew *string) *VirtualizationInterfacesListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithNameNisw(nameNisw *string) *VirtualizationInterfacesListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithOffset(offset *int64) *VirtualizationInterfacesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithQ(q *string) *VirtualizationInterfacesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetQ(q *string) {
	o.Q = q
}

// WithTag adds the tag to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithTag(tag *string) *VirtualizationInterfacesListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTagn adds the tagn to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithTagn(tagn *string) *VirtualizationInterfacesListParams {
	o.SetTagn(tagn)
	return o
}

// SetTagn adds the tagN to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetTagn(tagn *string) {
	o.Tagn = tagn
}

// WithVirtualMachine adds the virtualMachine to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithVirtualMachine(virtualMachine *string) *VirtualizationInterfacesListParams {
	o.SetVirtualMachine(virtualMachine)
	return o
}

// SetVirtualMachine adds the virtualMachine to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetVirtualMachine(virtualMachine *string) {
	o.VirtualMachine = virtualMachine
}

// WithVirtualMachinen adds the virtualMachinen to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithVirtualMachinen(virtualMachinen *string) *VirtualizationInterfacesListParams {
	o.SetVirtualMachinen(virtualMachinen)
	return o
}

// SetVirtualMachinen adds the virtualMachineN to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetVirtualMachinen(virtualMachinen *string) {
	o.VirtualMachinen = virtualMachinen
}

// WithVirtualMachineID adds the virtualMachineID to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithVirtualMachineID(virtualMachineID *string) *VirtualizationInterfacesListParams {
	o.SetVirtualMachineID(virtualMachineID)
	return o
}

// SetVirtualMachineID adds the virtualMachineId to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetVirtualMachineID(virtualMachineID *string) {
	o.VirtualMachineID = virtualMachineID
}

// WithVirtualMachineIDn adds the virtualMachineIDn to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) WithVirtualMachineIDn(virtualMachineIDn *string) *VirtualizationInterfacesListParams {
	o.SetVirtualMachineIDn(virtualMachineIDn)
	return o
}

// SetVirtualMachineIDn adds the virtualMachineIdN to the virtualization interfaces list params
func (o *VirtualizationInterfacesListParams) SetVirtualMachineIDn(virtualMachineIDn *string) {
	o.VirtualMachineIDn = virtualMachineIDn
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationInterfacesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cluster != nil {

		// query param cluster
		var qrCluster string

		if o.Cluster != nil {
			qrCluster = *o.Cluster
		}
		qCluster := qrCluster
		if qCluster != "" {

			if err := r.SetQueryParam("cluster", qCluster); err != nil {
				return err
			}
		}
	}

	if o.Clustern != nil {

		// query param cluster__n
		var qrClustern string

		if o.Clustern != nil {
			qrClustern = *o.Clustern
		}
		qClustern := qrClustern
		if qClustern != "" {

			if err := r.SetQueryParam("cluster__n", qClustern); err != nil {
				return err
			}
		}
	}

	if o.ClusterID != nil {

		// query param cluster_id
		var qrClusterID string

		if o.ClusterID != nil {
			qrClusterID = *o.ClusterID
		}
		qClusterID := qrClusterID
		if qClusterID != "" {

			if err := r.SetQueryParam("cluster_id", qClusterID); err != nil {
				return err
			}
		}
	}

	if o.ClusterIDn != nil {

		// query param cluster_id__n
		var qrClusterIDn string

		if o.ClusterIDn != nil {
			qrClusterIDn = *o.ClusterIDn
		}
		qClusterIDn := qrClusterIDn
		if qClusterIDn != "" {

			if err := r.SetQueryParam("cluster_id__n", qClusterIDn); err != nil {
				return err
			}
		}
	}

	if o.Enabled != nil {

		// query param enabled
		var qrEnabled string

		if o.Enabled != nil {
			qrEnabled = *o.Enabled
		}
		qEnabled := qrEnabled
		if qEnabled != "" {

			if err := r.SetQueryParam("enabled", qEnabled); err != nil {
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

	if o.MacAddress != nil {

		// query param mac_address
		var qrMacAddress string

		if o.MacAddress != nil {
			qrMacAddress = *o.MacAddress
		}
		qMacAddress := qrMacAddress
		if qMacAddress != "" {

			if err := r.SetQueryParam("mac_address", qMacAddress); err != nil {
				return err
			}
		}
	}

	if o.MacAddressIc != nil {

		// query param mac_address__ic
		var qrMacAddressIc string

		if o.MacAddressIc != nil {
			qrMacAddressIc = *o.MacAddressIc
		}
		qMacAddressIc := qrMacAddressIc
		if qMacAddressIc != "" {

			if err := r.SetQueryParam("mac_address__ic", qMacAddressIc); err != nil {
				return err
			}
		}
	}

	if o.MacAddressIe != nil {

		// query param mac_address__ie
		var qrMacAddressIe string

		if o.MacAddressIe != nil {
			qrMacAddressIe = *o.MacAddressIe
		}
		qMacAddressIe := qrMacAddressIe
		if qMacAddressIe != "" {

			if err := r.SetQueryParam("mac_address__ie", qMacAddressIe); err != nil {
				return err
			}
		}
	}

	if o.MacAddressIew != nil {

		// query param mac_address__iew
		var qrMacAddressIew string

		if o.MacAddressIew != nil {
			qrMacAddressIew = *o.MacAddressIew
		}
		qMacAddressIew := qrMacAddressIew
		if qMacAddressIew != "" {

			if err := r.SetQueryParam("mac_address__iew", qMacAddressIew); err != nil {
				return err
			}
		}
	}

	if o.MacAddressIsw != nil {

		// query param mac_address__isw
		var qrMacAddressIsw string

		if o.MacAddressIsw != nil {
			qrMacAddressIsw = *o.MacAddressIsw
		}
		qMacAddressIsw := qrMacAddressIsw
		if qMacAddressIsw != "" {

			if err := r.SetQueryParam("mac_address__isw", qMacAddressIsw); err != nil {
				return err
			}
		}
	}

	if o.MacAddressn != nil {

		// query param mac_address__n
		var qrMacAddressn string

		if o.MacAddressn != nil {
			qrMacAddressn = *o.MacAddressn
		}
		qMacAddressn := qrMacAddressn
		if qMacAddressn != "" {

			if err := r.SetQueryParam("mac_address__n", qMacAddressn); err != nil {
				return err
			}
		}
	}

	if o.MacAddressNic != nil {

		// query param mac_address__nic
		var qrMacAddressNic string

		if o.MacAddressNic != nil {
			qrMacAddressNic = *o.MacAddressNic
		}
		qMacAddressNic := qrMacAddressNic
		if qMacAddressNic != "" {

			if err := r.SetQueryParam("mac_address__nic", qMacAddressNic); err != nil {
				return err
			}
		}
	}

	if o.MacAddressNie != nil {

		// query param mac_address__nie
		var qrMacAddressNie string

		if o.MacAddressNie != nil {
			qrMacAddressNie = *o.MacAddressNie
		}
		qMacAddressNie := qrMacAddressNie
		if qMacAddressNie != "" {

			if err := r.SetQueryParam("mac_address__nie", qMacAddressNie); err != nil {
				return err
			}
		}
	}

	if o.MacAddressNiew != nil {

		// query param mac_address__niew
		var qrMacAddressNiew string

		if o.MacAddressNiew != nil {
			qrMacAddressNiew = *o.MacAddressNiew
		}
		qMacAddressNiew := qrMacAddressNiew
		if qMacAddressNiew != "" {

			if err := r.SetQueryParam("mac_address__niew", qMacAddressNiew); err != nil {
				return err
			}
		}
	}

	if o.MacAddressNisw != nil {

		// query param mac_address__nisw
		var qrMacAddressNisw string

		if o.MacAddressNisw != nil {
			qrMacAddressNisw = *o.MacAddressNisw
		}
		qMacAddressNisw := qrMacAddressNisw
		if qMacAddressNisw != "" {

			if err := r.SetQueryParam("mac_address__nisw", qMacAddressNisw); err != nil {
				return err
			}
		}
	}

	if o.Mtu != nil {

		// query param mtu
		var qrMtu string

		if o.Mtu != nil {
			qrMtu = *o.Mtu
		}
		qMtu := qrMtu
		if qMtu != "" {

			if err := r.SetQueryParam("mtu", qMtu); err != nil {
				return err
			}
		}
	}

	if o.MtuGt != nil {

		// query param mtu__gt
		var qrMtuGt string

		if o.MtuGt != nil {
			qrMtuGt = *o.MtuGt
		}
		qMtuGt := qrMtuGt
		if qMtuGt != "" {

			if err := r.SetQueryParam("mtu__gt", qMtuGt); err != nil {
				return err
			}
		}
	}

	if o.MtuGte != nil {

		// query param mtu__gte
		var qrMtuGte string

		if o.MtuGte != nil {
			qrMtuGte = *o.MtuGte
		}
		qMtuGte := qrMtuGte
		if qMtuGte != "" {

			if err := r.SetQueryParam("mtu__gte", qMtuGte); err != nil {
				return err
			}
		}
	}

	if o.MtuLt != nil {

		// query param mtu__lt
		var qrMtuLt string

		if o.MtuLt != nil {
			qrMtuLt = *o.MtuLt
		}
		qMtuLt := qrMtuLt
		if qMtuLt != "" {

			if err := r.SetQueryParam("mtu__lt", qMtuLt); err != nil {
				return err
			}
		}
	}

	if o.MtuLte != nil {

		// query param mtu__lte
		var qrMtuLte string

		if o.MtuLte != nil {
			qrMtuLte = *o.MtuLte
		}
		qMtuLte := qrMtuLte
		if qMtuLte != "" {

			if err := r.SetQueryParam("mtu__lte", qMtuLte); err != nil {
				return err
			}
		}
	}

	if o.Mtun != nil {

		// query param mtu__n
		var qrMtun string

		if o.Mtun != nil {
			qrMtun = *o.Mtun
		}
		qMtun := qrMtun
		if qMtun != "" {

			if err := r.SetQueryParam("mtu__n", qMtun); err != nil {
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

	if o.VirtualMachine != nil {

		// query param virtual_machine
		var qrVirtualMachine string

		if o.VirtualMachine != nil {
			qrVirtualMachine = *o.VirtualMachine
		}
		qVirtualMachine := qrVirtualMachine
		if qVirtualMachine != "" {

			if err := r.SetQueryParam("virtual_machine", qVirtualMachine); err != nil {
				return err
			}
		}
	}

	if o.VirtualMachinen != nil {

		// query param virtual_machine__n
		var qrVirtualMachinen string

		if o.VirtualMachinen != nil {
			qrVirtualMachinen = *o.VirtualMachinen
		}
		qVirtualMachinen := qrVirtualMachinen
		if qVirtualMachinen != "" {

			if err := r.SetQueryParam("virtual_machine__n", qVirtualMachinen); err != nil {
				return err
			}
		}
	}

	if o.VirtualMachineID != nil {

		// query param virtual_machine_id
		var qrVirtualMachineID string

		if o.VirtualMachineID != nil {
			qrVirtualMachineID = *o.VirtualMachineID
		}
		qVirtualMachineID := qrVirtualMachineID
		if qVirtualMachineID != "" {

			if err := r.SetQueryParam("virtual_machine_id", qVirtualMachineID); err != nil {
				return err
			}
		}
	}

	if o.VirtualMachineIDn != nil {

		// query param virtual_machine_id__n
		var qrVirtualMachineIDn string

		if o.VirtualMachineIDn != nil {
			qrVirtualMachineIDn = *o.VirtualMachineIDn
		}
		qVirtualMachineIDn := qrVirtualMachineIDn
		if qVirtualMachineIDn != "" {

			if err := r.SetQueryParam("virtual_machine_id__n", qVirtualMachineIDn); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
