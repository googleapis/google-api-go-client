// Copyright 2021 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build integration

// To run this test locally, you will need to do the following:
// • Navigate to your Google Cloud Project
// • Get a copy of a Service Account Key File for testing (should be in .json format)
// • If you are unable to obtain an existing key file, create one:
//    • > IAM and Admin > Service Accounts
//    • Under the needed Service Account > Actions > Manage Keys
//    • Add Key > Create New Key
//    • Select JSON, and the click Create
// • Look for an available VM Instance, or create one- > Compute > Compute Engine > VM Instances
// • On the VM Instance, click the SSH Button.  Then upload:
//    • Your Service Account Key File
//    • This script, along with setup.sh
//    • A copy of env.conf, containing the required environment variables (see existing skeleton)/
// • Set your environment variables (Usually this will be `source env.conf`)
// • Ensure that your VM is properly set up to run the integration test e.g.
//    • wget -c https://golang.org/dl/go1.15.2.linux-amd64.tar.gz
//       • Check https://golang.org/dl/for the latest version of Go
//    • sudo tar -C /usr/local -xvzf go1.15.2.linux-amd64.tar.gz
//    • go mod init google.golang.org/api/google-api-go-client
//    • go mod tidy
// • Run setup.sh
// • go test -tags integration`

package byoid

import (
	"context"
	//"crypto/hmac"
	//"crypto/sha256"
	//"encoding/hex"
	"encoding/json"
	"fmt"
	//"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	//neturl "net/url"
	"os"
	//"path"
	//"sort"
	//"strings"
	"testing"
	//"time"

	"google.golang.org/api/dns/v1"
	"google.golang.org/api/idtoken"
	"google.golang.org/api/option"
)

const (
	envCredentials  = "GCLOUD_TESTS_GOLANG_KEY"
	envAudienceOIDC = "GCLOUD_TESTS_GOLANG_AUDIENCE_OIDC"
	envProject      = "GCLOUD_TESTS_GOLANG_PROJECT_ID"
)

//// The following are specifically for AWS integration tests.
//const (
//	// AWS Signature Version 4 signing algorithm identifier.
//	awsAlgorithm = "AWS4-HMAC-SHA256"
//
//	// The termination string for the AWS credential scope value as defined in
//	// https://docs.aws.amazon.com/general/latest/gr/sigv4-create-string-to-sign.html
//	awsRequestType = "aws4_request"
//
//	// The AWS authorization header name for the security session token if available.
//	awsSecurityTokenHeader = "x-amz-security-token"
//
//	// The AWS authorization header name for the auto-generated date.
//	awsDateHeader = "x-amz-date"
//
//	awsTimeFormatLong  = "20060102T150405Z"
//	awsTimeFormatShort = "20060102"
//)

var (
	oidcAudience string
	oidcToken    string
	awsToken     string
	clientID     string
	projectID    string
)

// TestMain contains all of the setup code that needs to be run once before any of the tests are run
func TestMain(m *testing.M) {
	keyFileName := os.Getenv(envCredentials)
	if keyFileName == "" {
		log.Fatalf("Please set %s to your keyfile", envCredentials)
	}

	projectID = os.Getenv(envProject)
	if projectID == "" {
		log.Fatalf("Please set %s to the ID of the project", envProject)
	}

	oidcAudience = os.Getenv(envAudienceOIDC)
	if oidcAudience == "" {
		log.Fatalf("Please set %s to the OIDC Audience", envAudienceOIDC)
	}

	var err error

	clientID, err = getClientID(keyFileName)
	if err != nil {
		log.Fatalf("Error getting Client ID: %v", err)
	}

	oidcToken, err = generateGoogleToken(keyFileName)
	if err != nil {
		log.Fatalf("Error generating Google token: %v", err)
	}

	// This line runs all of our individual tests
	os.Exit(m.Run())
}

// keyFile is a struct to extract the relevant json fields for our ServiceAccount KeyFile
type keyFile struct {
	ClientEmail string `json:"client_email"`
	ClientID    string `json:"client_id"`
}

func getClientID(keyFileName string) (string, error) {
	kf, err := os.Open(keyFileName)
	if err != nil {
		return "", err
	}
	defer kf.Close()

	decoder := json.NewDecoder(kf)
	var keyFileSettings keyFile
	err = decoder.Decode(&keyFileSettings)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("projects/-/serviceAccounts/%s", keyFileSettings.ClientEmail), nil
}

func generateGoogleToken(keyFileName string) (string, error) {
	ts, err := idtoken.NewTokenSource(context.Background(), oidcAudience, option.WithCredentialsFile(keyFileName))
	if err != nil {
		return "", nil
	}

	token, err := ts.Token()
	if err != nil {
		return "", nil
	}

	return token.AccessToken, nil
}

// testBYOID makes sure that the default credentials works for
// whatever preconditions have been set beforehand
// by using those credentials to run our client libraries.
//
// In each test we will set up whatever preconditions we need,
// and then use this function.
func testBYOID(t *testing.T, c config) {
	t.Helper()

	// Set up config file.
	configFile, err := ioutil.TempFile("", "config.json")
	if err != nil {
		t.Fatalf("Error creating config file: %v", err)
	}
	defer os.Remove(configFile.Name())

	err = json.NewEncoder(configFile).Encode(c)
	if err != nil {
		t.Errorf("Error writing to config file: %v", err)
	}
	configFile.Close()

	// Once the default credentials are obtained,
	// we should be able to access Google Cloud resources.
	dnsService, err := dns.NewService(context.Background(), option.WithCredentialsFile(configFile.Name()))
	if err != nil {
		t.Fatalf("Could not establish DNS Service: %v", err)
	}

	_, err = dnsService.Projects.Get(projectID).Do()
	if err != nil {
		t.Fatalf("DNS Service failed: %v", err)
	}
}

// These structs makes writing our config as json to a file much easier.
type config struct {
	Type                           string           `json:"type"`
	Audience                       string           `json:"audience"`
	SubjectTokenType               string           `json:"subject_token_type"`
	TokenURL                       string           `json:"token_url"`
	ServiceAccountImpersonationURL string           `json:"service_account_impersonation_url"`
	CredentialSource               credentialSource `json:"credential_source"`
}

type credentialSource struct {
	File string `json:"file,omitempty"`

	URL     string            `json:"url,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`

	EnvironmentID               string `json:"environment_id,omitempty"`
	RegionURL                   string `json:"region_url"`
	RegionalCredVerificationURL string `json:"regional_cred_verification_url,omitempty"`
	Format                      string `json:"format,omitempty"`
}

// Tests to make sure File based external credentials continues to work.
func TestFileBasedCredentials(t *testing.T) {
	// Set up Token as a file
	tokenFile, err := ioutil.TempFile("", "token.txt")
	if err != nil {
		t.Fatalf("Error creating token file:")
	}
	defer os.Remove(tokenFile.Name())

	tokenFile.WriteString(oidcToken)
	tokenFile.Close()

	// Run our test!
	testBYOID(t, config{
		Type:                           "external_account",
		Audience:                       oidcAudience,
		SubjectTokenType:               "urn:ietf:params:oauth:token-type:jwt",
		TokenURL:                       "https://sts.googleapis.com/v1beta/token",
		ServiceAccountImpersonationURL: fmt.Sprintf("https://iamcredentials.googleapis.com/v1/%s:generateAccessToken", clientID),
		CredentialSource: credentialSource{
			File: tokenFile.Name(),
		},
	})
}

// Tests to make sure URL based external credentials work properly.
func TestURLBasedCredentials(t *testing.T) {
	//Set up a server to return a token
	 ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		 if r.Method != "GET" {
			 t.Errorf("Unexpected request method, %v is found", r.Method)
		 }
		 w.Write([]byte(oidcToken))
	 }))

	 testBYOID(t, config{
	 	Type:			"external_account",
	 	Audience:		oidcAudience,
	 	SubjectTokenType:		"urn:ietf:params:oauth:token-type:jwt",
	 	TokenURL:		"https://sts.googleapis.com/v1beta/token",
	 	ServiceAccountImpersonationURL: fmt.Sprintf("https://iamcredentials.googleapis.com/v1/%s:generateAccessToken", clientID),
	 	CredentialSource: credentialSource{
	 		URL: ts.URL,
		},
	 })
}

//// now aliases time.Now for testing
//var now = func() time.Time {
//	return time.Now().UTC()
//}

//// Tests to make sure AWS based external credentials work properly.
//func TestAWSBasedCredentials(t *testing.T) {
//	server := createDefaultAwsTestServer()
//	ts := httptest.NewServer(server)
//
//	testBYOID(t, config{
//		Type:			"external_account",
//		Audience:		oidcAudience,
//		SubjectTokenType:		"urn:ietf:params:oauth:token-type:jwt",
//		TokenURL:		"https://sts.googleapis.com/v1beta/token",
//		ServiceAccountImpersonationURL: fmt.Sprintf("https://iamcredentials.googleapis.com/v1/%s:generateAccessToken", clientID),
//		CredentialSource: credentialSource{
//			EnvironmentID:               "aws1",
//			URL:                         ts.URL + server.url,
//			RegionURL:                   ts.URL + server.regionURL,
//			RegionalCredVerificationURL: server.regionalCredVerificationURL,
//		},
//	})
//	//oldGetenv := getenv
//	//defer func() { getenv = oldGetenv }()
//	//getenv = setEnvironment(map[string]string{})
//	//
//	//base, err := tfc.parse(context.Background())
//	//if err != nil {
//	//	t.Fatalf("parse() failed %v", err)
//	//}
//	//
//	//out, err := base.subjectToken()
//}
//
//
//func setEnvironment(env map[string]string) func(string) string {
//	return func(key string) string {
//		value, _ := env[key]
//		return value
//	}
//}
//
//
//type awsSecurityCredentials struct {
//	AccessKeyID     string `json:"AccessKeyID"`
//	SecretAccessKey string `json:"SecretAccessKey"`
//	SecurityToken   string `json:"Token"`
//}
//
//// awsRequestSigner is a utility class to sign http requests using a AWS V4 signature.
//type awsRequestSigner struct {
//	RegionName             string
//	AwsSecurityCredentials awsSecurityCredentials
//}
//type awsRequestHeader struct {
//	Key   string `json:"key"`
//	Value string `json:"value"`
//}
//
//type awsRequest struct {
//	URL     string             `json:"url"`
//	Method  string             `json:"method"`
//	Headers []awsRequestHeader `json:"headers"`
//}
//
//var defaultRequestSigner = &awsRequestSigner{
//	RegionName: "us-east-1",
//	AwsSecurityCredentials: awsSecurityCredentials{
//		AccessKeyID:     "AKIDEXAMPLE",
//		SecretAccessKey: "wJalrXUtnFEMI/K7MDENG+bPxRfiCYEXAMPLEKEY",
//	},
//}
//
//const (
//	accessKeyID     = "ASIARD4OQDT6A77FR3CL"
//	secretAccessKey = "Y8AfSaucF37G4PpvfguKZ3/l7Id4uocLXxX0+VTx"
//	securityToken   = "IQoJb3JpZ2luX2VjEIz//////////wEaCXVzLWVhc3QtMiJGMEQCIH7MHX/Oy/OB8OlLQa9GrqU1B914+iMikqWQW7vPCKlgAiA/Lsv8Jcafn14owfxXn95FURZNKaaphj0ykpmS+Ki+CSq0AwhlEAAaDDA3NzA3MTM5MTk5NiIMx9sAeP1ovlMTMKLjKpEDwuJQg41/QUKx0laTZYjPlQvjwSqS3OB9P1KAXPWSLkliVMMqaHqelvMF/WO/glv3KwuTfQsavRNs3v5pcSEm4SPO3l7mCs7KrQUHwGP0neZhIKxEXy+Ls//1C/Bqt53NL+LSbaGv6RPHaX82laz2qElphg95aVLdYgIFY6JWV5fzyjgnhz0DQmy62/Vi8pNcM2/VnxeCQ8CC8dRDSt52ry2v+nc77vstuI9xV5k8mPtnaPoJDRANh0bjwY5Sdwkbp+mGRUJBAQRlNgHUJusefXQgVKBCiyJY4w3Csd8Bgj9IyDV+Azuy1jQqfFZWgP68LSz5bURyIjlWDQunO82stZ0BgplKKAa/KJHBPCp8Qi6i99uy7qh76FQAqgVTsnDuU6fGpHDcsDSGoCls2HgZjZFPeOj8mmRhFk1Xqvkbjuz8V1cJk54d3gIJvQt8gD2D6yJQZecnuGWd5K2e2HohvCc8Fc9kBl1300nUJPV+k4tr/A5R/0QfEKOZL1/k5lf1g9CREnrM8LVkGxCgdYMxLQow1uTL+QU67AHRRSp5PhhGX4Rek+01vdYSnJCMaPhSEgcLqDlQkhk6MPsyT91QMXcWmyO+cAZwUPwnRamFepuP4K8k2KVXs/LIJHLELwAZ0ekyaS7CptgOqS7uaSTFG3U+vzFZLEnGvWQ7y9IPNQZ+Dffgh4p3vF4J68y9049sI6Sr5d5wbKkcbm8hdCDHZcv4lnqohquPirLiFQ3q7B17V9krMPu3mz1cg4Ekgcrn/E09NTsxAqD8NcZ7C7ECom9r+X3zkDOxaajW6hu3Az8hGlyylDaMiFfRbBJpTIlxp7jfa7CxikNgNtEKLH9iCzvuSg2vhA=="
//)
//
//var requestSignerWithToken = &awsRequestSigner{
//	RegionName: "us-east-2",
//	AwsSecurityCredentials: awsSecurityCredentials{
//		AccessKeyID:     accessKeyID,
//		SecretAccessKey: secretAccessKey,
//		SecurityToken:   securityToken,
//	},
//}
//
//func setDefaultTime(req *http.Request) {
//	// Don't use time.Format for this
//	// Our output signature expects this to be a Monday, even though Sept 9, 2011 is a Friday
//	req.Header.Add("date", "Mon, 09 Sep 2011 23:36:00 GMT")
//}
//
//type testAwsServer struct {
//	url                         string
//	securityCredentialURL       string
//	regionURL                   string
//	regionalCredVerificationURL string
//
//	Credentials map[string]string
//
//	WriteRolename            func(http.ResponseWriter)
//	WriteSecurityCredentials func(http.ResponseWriter)
//	WriteRegion              func(http.ResponseWriter)
//}
//
//func createAwsTestServer(url, regionURL, regionalCredVerificationURL, rolename, region string, credentials map[string]string) *testAwsServer {
//	server := &testAwsServer{
//		url:                         url,
//		securityCredentialURL:       fmt.Sprintf("%s/%s", url, rolename),
//		regionURL:                   regionURL,
//		regionalCredVerificationURL: regionalCredVerificationURL,
//		Credentials:                 credentials,
//		WriteRolename: func(w http.ResponseWriter) {
//			w.Write([]byte(rolename))
//		},
//		WriteRegion: func(w http.ResponseWriter) {
//			w.Write([]byte(region))
//		},
//	}
//
//	server.WriteSecurityCredentials = func(w http.ResponseWriter) {
//		jsonCredentials, _ := json.Marshal(server.Credentials)
//		w.Write(jsonCredentials)
//	}
//
//	return server
//}
//
//func createDefaultAwsTestServer() *testAwsServer {
//	return createAwsTestServer(
//		"/latest/meta-data/iam/security-credentials",
//		"/latest/meta-data/placement/availability-zone",
//		"https://sts.{region}.amazonaws.com?Action=GetCallerIdentity&Version=2011-06-15",
//		"gcp-aws-role",
//		"us-east-2b",
//		map[string]string{
//			"SecretAccessKey": secretAccessKey,
//			"AccessKeyId":     accessKeyID,
//			"Token":           securityToken,
//		},
//	)
//}
//
//func (server *testAwsServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	switch p := r.URL.Path; p {
//	case server.url:
//		server.WriteRolename(w)
//	case server.securityCredentialURL:
//		server.WriteSecurityCredentials(w)
//	case server.regionURL:
//		server.WriteRegion(w)
//	}
//}
//
//func notFound(w http.ResponseWriter) {
//	w.WriteHeader(404)
//	w.Write([]byte("Not Found"))
//}
//
////func (server *testAwsServer) getCredentialSource(url string) CredentialSource {
////	return CredentialSource{
////		EnvironmentID:               "aws1",
////		URL:                         url + server.url,
////		RegionURL:                   url + server.regionURL,
////		RegionalCredVerificationURL: server.regionalCredVerificationURL,
////	}
////}
//
//func getExpectedSubjectToken(url, region, accessKeyID, secretAccessKey, securityToken string) string {
//	req, _ := http.NewRequest("POST", url, nil)
//	//TODO: Does this work?  Also 1 other swapped line.
//	//req.Header.Add("x-goog-cloud-target-resource", testFileConfig.Audience)
//	req.Header.Add("x-goog-cloud-target-resource", oidcAudience)
//	signer := &awsRequestSigner{
//		RegionName: region,
//		AwsSecurityCredentials: awsSecurityCredentials{
//			AccessKeyID:     accessKeyID,
//			SecretAccessKey: secretAccessKey,
//			SecurityToken:   securityToken,
//		},
//	}
//	signer.SignRequest(req)
//
//	result := awsRequest{
//		URL:    url,
//		Method: "POST",
//		Headers: []awsRequestHeader{
//			{
//				Key:   "Authorization",
//				Value: req.Header.Get("Authorization"),
//			}, {
//				Key:   "Host",
//				Value: req.Header.Get("Host"),
//			}, {
//				Key:   "X-Amz-Date",
//				Value: req.Header.Get("X-Amz-Date"),
//			},
//		},
//	}
//
//	if securityToken != "" {
//		result.Headers = append(result.Headers, awsRequestHeader{
//			Key:   "X-Amz-Security-Token",
//			Value: securityToken,
//		})
//	}
//
//	result.Headers = append(result.Headers, awsRequestHeader{
//		Key:   "X-Goog-Cloud-Target-Resource",
//		//Value: testFileConfig.Audience,
//		Value: oidcAudience,
//	})
//
//	str, _ := json.Marshal(result)
//	return neturl.QueryEscape(string(str))
//}
//
//func cloneRequest(r *http.Request) *http.Request {
//	r2 := new(http.Request)
//	*r2 = *r
//	if r.Header != nil {
//		r2.Header = make(http.Header, len(r.Header))
//
//		// Find total number of values.
//		headerCount := 0
//		for _, headerValues := range r.Header {
//			headerCount += len(headerValues)
//		}
//		copiedHeaders := make([]string, headerCount) // shared backing array for headers' values
//
//		for headerKey, headerValues := range r.Header {
//			headerCount = copy(copiedHeaders, headerValues)
//			r2.Header[headerKey] = copiedHeaders[:headerCount:headerCount]
//			copiedHeaders = copiedHeaders[headerCount:]
//		}
//	}
//	return r2
//}
//
//func requestDataHash(req *http.Request) (string, error) {
//	var requestData []byte
//	if req.Body != nil {
//		requestBody, err := req.GetBody()
//		if err != nil {
//			return "", err
//		}
//		defer requestBody.Close()
//
//		requestData, err = ioutil.ReadAll(io.LimitReader(requestBody, 1<<20))
//		if err != nil {
//			return "", err
//		}
//	}
//
//	return getSha256(requestData)
//}
//
//func requestHost(req *http.Request) string {
//	if req.Host != "" {
//		return req.Host
//	}
//	return req.URL.Host
//}
//
//// SignRequest adds the appropriate headers to an http.Request
//// or returns an error if something prevented this.
//func (rs *awsRequestSigner) SignRequest(req *http.Request) error {
//	signedRequest := cloneRequest(req)
//	timestamp := now()
//
//	signedRequest.Header.Add("host", requestHost(req))
//
//	if rs.AwsSecurityCredentials.SecurityToken != "" {
//		signedRequest.Header.Add(awsSecurityTokenHeader, rs.AwsSecurityCredentials.SecurityToken)
//	}
//
//	if signedRequest.Header.Get("date") == "" {
//		signedRequest.Header.Add(awsDateHeader, timestamp.Format(awsTimeFormatLong))
//	}
//
//	authorizationCode, err := rs.generateAuthentication(signedRequest, timestamp)
//	if err != nil {
//		return err
//	}
//	signedRequest.Header.Set("Authorization", authorizationCode)
//
//	req.Header = signedRequest.Header
//	return nil
//}
//
//func (rs *awsRequestSigner) generateAuthentication(req *http.Request, timestamp time.Time) (string, error) {
//	canonicalHeaderColumns, canonicalHeaderData := canonicalHeaders(req)
//
//	dateStamp := timestamp.Format(awsTimeFormatShort)
//	serviceName := ""
//	if splitHost := strings.Split(requestHost(req), "."); len(splitHost) > 0 {
//		serviceName = splitHost[0]
//	}
//
//	credentialScope := fmt.Sprintf("%s/%s/%s/%s", dateStamp, rs.RegionName, serviceName, awsRequestType)
//
//	requestString, err := canonicalRequest(req, canonicalHeaderColumns, canonicalHeaderData)
//	if err != nil {
//		return "", err
//	}
//	requestHash, err := getSha256([]byte(requestString))
//	if err != nil {
//		return "", err
//	}
//
//	stringToSign := fmt.Sprintf("%s\n%s\n%s\n%s", awsAlgorithm, timestamp.Format(awsTimeFormatLong), credentialScope, requestHash)
//
//	signingKey := []byte("AWS4" + rs.AwsSecurityCredentials.SecretAccessKey)
//	for _, signingInput := range []string{
//		dateStamp, rs.RegionName, serviceName, awsRequestType, stringToSign,
//	} {
//		signingKey, err = getHmacSha256(signingKey, []byte(signingInput))
//		if err != nil {
//			return "", err
//		}
//	}
//
//	return fmt.Sprintf("%s Credential=%s/%s, SignedHeaders=%s, Signature=%s", awsAlgorithm, rs.AwsSecurityCredentials.AccessKeyID, credentialScope, canonicalHeaderColumns, hex.EncodeToString(signingKey)), nil
//}
//
//func canonicalRequest(req *http.Request, canonicalHeaderColumns, canonicalHeaderData string) (string, error) {
//	dataHash, err := requestDataHash(req)
//	if err != nil {
//		return "", err
//	}
//
//	return fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s", req.Method, canonicalPath(req), canonicalQuery(req), canonicalHeaderData, canonicalHeaderColumns, dataHash), nil
//}
//
//func canonicalPath(req *http.Request) string {
//	result := req.URL.EscapedPath()
//	if result == "" {
//		return "/"
//	}
//	return path.Clean(result)
//}
//
//func canonicalQuery(req *http.Request) string {
//	queryValues := req.URL.Query()
//	for queryKey := range queryValues {
//		sort.Strings(queryValues[queryKey])
//	}
//	return queryValues.Encode()
//}
//
//func canonicalHeaders(req *http.Request) (string, string) {
//	// Header keys need to be sorted alphabetically.
//	var headers []string
//	lowerCaseHeaders := make(http.Header)
//	for k, v := range req.Header {
//		k := strings.ToLower(k)
//		if _, ok := lowerCaseHeaders[k]; ok {
//			// include additional values
//			lowerCaseHeaders[k] = append(lowerCaseHeaders[k], v...)
//		} else {
//			headers = append(headers, k)
//			lowerCaseHeaders[k] = v
//		}
//	}
//	sort.Strings(headers)
//
//	var fullHeaders strings.Builder
//	for _, header := range headers {
//		headerValue := strings.Join(lowerCaseHeaders[header], ",")
//		fullHeaders.WriteString(header)
//		fullHeaders.WriteRune(':')
//		fullHeaders.WriteString(headerValue)
//		fullHeaders.WriteRune('\n')
//	}
//
//	return strings.Join(headers, ";"), fullHeaders.String()
//}
//func getSha256(input []byte) (string, error) {
//	hash := sha256.New()
//	if _, err := hash.Write(input); err != nil {
//		return "", err
//	}
//	return hex.EncodeToString(hash.Sum(nil)), nil
//}
//
//func getHmacSha256(key, input []byte) ([]byte, error) {
//	hash := hmac.New(sha256.New, key)
//	if _, err := hash.Write(input); err != nil {
//		return nil, err
//	}
//	return hash.Sum(nil), nil
//}