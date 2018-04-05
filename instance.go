//go:generate $GOPATH/src/github.com/fortifi/go-api/generate.sh

package api

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/lucidcube/omni-flux/fortifiapi/client"
	"github.com/lucidcube/omni-flux/fortifiapi/client/operations"
	"github.com/lucidcube/omni-flux/fortifiapi/models"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

const (
	expiryBuffer             = 14400
	devHost                  = "api.fortifi.biz:9090"
	fortifiOrgHeader         = "x-fortifi-org"
	fortifiStrictHeader      = "x-fortifi-strict"
	fortifiAuthHeader        = "authorization"
	fortifiBearerTokenSchema = "Bearer %s"
)

var (
	debugFortifiRequests = false
	localSchemes         = []string{"http"}

	// map with relationship fidenttoken.ISS > ForitifInstance
	fortifiAPIMap = sync.Map{}
	productCache  = sync.Map{}
	groupCache    = sync.Map{}
)

// FortifiInstance is a single org API instance
type FortifiInstance struct {
	expiry              int64
	organisationFID     string
	authtoken           string
	user                string
	key                 string
	apiInstance         *client.Fortifi
	authenticator       *Authenticator
	hostingProductGroup string
	domainProductGroup  string
}

// Authenticator auths fortifi requests
type Authenticator struct {
	organisationFID string
	authtoken       string
}

// SetInstanceForIssuer is setter for instance relationship store
func SetInstanceForIssuer(iss string, instance FortifiInstance) {
	fortifiAPIMap.Store(iss, instance)
}

// GetInstanceForIssuer is getter for instance relationship store
func GetInstanceForIssuer(iss string) (*FortifiInstance, error) {
	i, ok := fortifiAPIMap.Load(iss)
	if !ok {
		return nil, errors.New("No Fortifi instance map for issuer")
	}

	in := i.(FortifiInstance)
	return &in, nil
}

// NewInstance creates a new FortifiInstance with embeded tokenhelper
func NewInstance(orgFid, user, key string, debug bool) (FortifiInstance, error) {
	inst := FortifiInstance{}
	inst.SetOrganisationFID(orgFid)
	inst.SetUser(user)
	inst.SetKey(key)
	inst.SetLocalDebug(debug)
	return inst, nil
}

// AuthenticateRequest handles request authentication
func (a *Authenticator) AuthenticateRequest(req runtime.ClientRequest, reg strfmt.Registry) error {
	req.SetHeaderParam(fortifiOrgHeader, a.organisationFID)
	req.SetHeaderParam(fortifiStrictHeader, "true")
	req.SetHeaderParam(fortifiAuthHeader, fmt.Sprintf(fortifiBearerTokenSchema, a.authtoken))
	return nil
}

// InitAPI called by main function to start the fortifi API and webhook listener
func (f *FortifiInstance) InitAPI() error {
	_, err := f.GetAPIInstance()
	return err
}

// GetAPIInstance returns current instance of fortifi API
func (f *FortifiInstance) GetAPIInstance() (*client.Fortifi, error) {
	if f.apiInstance != nil && time.Now().Unix() < (f.expiry-expiryBuffer) {
		return f.apiInstance, nil
	}

	if f.organisationFID == "" {
		return nil, errors.New("Organization FID is required but not set")
	}

	if f.user == "" {
		return nil, errors.New("User is required but not set")
	}

	if f.key == "" {
		return nil, errors.New("Key is required but not set")
	}

	transport := httptransport.New(f.GetAPIHost(), client.DefaultBasePath, f.GetSchemes())
	transport.DefaultAuthentication = f.GetAuthenticator()
	transport.SetDebug(debugFortifiRequests)
	err := f.getNewToken(transport)
	if err != nil {
		return nil, errors.New("Failed to authenticate with Fortifi")
	}

	f.apiInstance = client.New(transport, strfmt.Default)
	return f.apiInstance, nil
}

func (f *FortifiInstance) getNewToken(transport *httptransport.Runtime) error {
	c := client.New(transport, strfmt.Default)
	p := operations.NewGetServiceAuthTokenParams()
	p.Payload = &models.ServiceAccountCredentialsPayload{
		ID:  &f.user,
		Key: &f.key,
	}

	re, err := c.Operations.GetServiceAuthToken(p)
	if err != nil {
		return err
	}

	f.expiry = re.Payload.Expiry
	f.authtoken = re.Payload.Token
	return nil
}

// GetAuthenticator returns fortifi authenticator instance
func (f *FortifiInstance) GetAuthenticator() *Authenticator {
	return &Authenticator{organisationFID: f.organisationFID, authtoken: f.authtoken}
}

// GetSchemes returns the API protocol
func (f *FortifiInstance) GetSchemes() []string {
	if debugFortifiRequests {
		return localSchemes
	}
	return client.DefaultSchemes
}

// SetLocalDebug determines if local debug vars and logging are used for API
func (f *FortifiInstance) SetLocalDebug(s bool) {
	debugFortifiRequests = s
}

// GetAPIHost returns the API host URL
func (f *FortifiInstance) GetAPIHost() string {
	host := client.DefaultHost
	res := ""
	if debugFortifiRequests {
		host = devHost
	}
	res = host
	return res
}

// SetOrganisationFID sets the organisation fid used by the fortifi API
// this is case sensitive
func (f *FortifiInstance) SetOrganisationFID(of string) {
	f.organisationFID = of
}

// SetUser sets the user ID (This should be a service account) used by the fortifi API
func (f *FortifiInstance) SetUser(u string) {
	f.user = u
}

// SetKey sets the key for given User ID used by the fortifi API
func (f *FortifiInstance) SetKey(k string) {
	f.key = k
}
