// Copyright (c) 2019-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

package security

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/log"
	"github.com/beevik/etree"
)

// GetSecurityContext retrieves the security context fields from the given
// JSON-RPC request. Will return nil security context for an anonymous request.
func GetSecurityContext(jsonRequestBody map[string]interface{}) (map[string]interface{}, error) {
	paramValue, ok := jsonRequestBody[lib.JSONRPC_PARAMS]
	if !ok {
		log.Error("Field 'params' missing in the request")
		return nil, errors.New("field 'params' missing in the request")
	}
	paramMap, ok := paramValue.(map[string]interface{})
	if !ok {
		log.Errorf("Field 'params' must be a JSON object")
		return nil, errors.New("field 'params' must be a JSON object")
	}
	execContext, ok := paramMap[lib.EXECUTION_CONTEXT]
	if !ok {
		log.Error("Execution context missing in the request")
		return nil, errors.New("execution context missing in the request")
	}
	execContextMap, ok := execContext.(map[string]interface{})
	if !ok {
		log.Error("Execution context must be a JSON object")
		return nil, errors.New("execution context must be a JSON object")
	}
	securityContext, ok := execContextMap[lib.SECURITY_CONTEXT]
	if !ok {
		// anonymous API requests contain no security context
		return nil, nil
	}
	securityContextMap, ok := securityContext.(map[string]interface{})
	if !ok {
		log.Error("Security context must be a JSON object")
		return nil, errors.New("security context must be a JSON object")
	}
	return securityContextMap, nil
}

// Sets given security context to request body.
func SetSecurityContext(jsonRequestBody map[string]interface{}, securityContext map[string]interface{}) error {
	if paramValue, ok := jsonRequestBody[lib.JSONRPC_PARAMS]; ok {
		if _, ok := paramValue.(map[string]interface{}); !ok {
			log.Error("Value of json rpc parameter extracted from json request body failed assertion of type map[string]interface")
			return errors.New("Error setting security context")
		}
		paramMap, isObject := paramValue.(map[string]interface{})
		if !isObject {
			return errors.New("Error setting security context")
		}
		if execContext, ok := paramMap[lib.EXECUTION_CONTEXT]; ok {
			if _, ok := execContext.(map[string]interface{}); !ok {
				log.Error("Value of Execution Ctx extracted from json paramter map failed assertion of type map[string]interface")
				return errors.New("Error setting security context")
			}
			execContext := execContext.(map[string]interface{})
			execContext[lib.SECURITY_CONTEXT] = securityContext
		} else {
			return errors.New("Error setting security context")
		}
	}
	return nil
}

// Append proper prefix and suffix text to a private key. There is no
// standard way for storing certificates. OpenSSL expects the demarcation
// text. This method makes sure that the text the markers are present.
func preparePrivateKey(privateKey string) string {
	if !strings.HasPrefix(privateKey, "-----BEGIN RSA PRIVATE KEY-----") {
		return fmt.Sprintf("-----BEGIN RSA PRIVATE KEY-----\n%s\n-----END RSA PRIVATE KEY-----", privateKey)
	}
	return privateKey
}

// Extract Certificate of the principal from Holder of Key SAML token
func ExtractCertificate(samlToken string) (*x509.Certificate, error) {
	doc := etree.NewDocument()
	doc.ReadFromString(samlToken)
	subject := doc.FindElement("//saml2:SubjectConfirmationData")
	cert := subject.FindElement("./KeyInfo/X509Data/X509Certificate")
	certString := GetElementText(cert)
	block, _ := pem.Decode([]byte(PrepareCertificate(certString)))
	if block == nil {
		return nil, errors.New("failed to parse certificate PEM")
	}
	x509, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, errors.New("failed to parse certificate: " + err.Error())
	}
	return x509, nil
}

// Append proper prefix and suffix text to a certificate. There is no
// standard way for storing certificates. OpenSSL expects the demarcation
// text. This method makes sure that the text the markers are present.
func PrepareCertificate(certString string) string {
	if !strings.HasPrefix(certString, "-----BEGIN CERTIFICATE-----") {
		return fmt.Sprintf("-----BEGIN CERTIFICATE-----\n%s\n-----END CERTIFICATE-----", certString)
	}
	return certString
}

// Generate a timestamp for the request. This will be embedded in the security
// context of the request to protect it against replay attacks
func GenerateRequestTimeStamp() map[string]string {
	createdDate := time.Now().UTC()
	expires := createdDate.Add(time.Minute * REQUEST_VALIDITY)
	return map[string]string{TS_EXPIRES_KEY: expires.Format(bindings.VAPI_DATETIME_LAYOUT),
		TS_CREATED_KEY: createdDate.Format(bindings.VAPI_DATETIME_LAYOUT)}
}

// Verify signature of sig by generating signature using public key with toVerify
func VerifySignature(pubKey *rsa.PublicKey, algorithm crypto.Hash, toVerify []byte, sig []byte) error {
	h := algorithm.New()
	h.Write(toVerify)
	hashed := h.Sum(nil)
	return rsa.VerifyPKCS1v15(pubKey, algorithm, hashed, sig)
}

// GetElementText of a public certificate in samltoken
func GetElementText(element *etree.Element) string {
	var sb strings.Builder
	for _, cData := range element.Child {
		charData := cData.(*etree.CharData).Data
		sb.WriteString(charData)
	}
	result := sb.String()
	return strings.Replace(result, "\\n", "\n", -1)
}

// Generates signature for 'toSign' bytes using provided 'algorithm' and 'privateKey'
func Sign(toSign []byte, algorithm crypto.Hash, privateKey *rsa.PrivateKey) ([]byte, error) {
	hash := algorithm.New()
	hash.Write(toSign)
	hashedData := hash.Sum(nil)
	return rsa.SignPKCS1v15(rand.Reader, privateKey, algorithm, hashedData)
}

// ParsePrivateKey parses private key from given input string.
func ParsePrivateKey(pemData string) (crypto.PrivateKey, error) {
	block, _ := pem.Decode([]byte(pemData))
	if block == nil {
		return nil, errors.New("error decoding private key")
	}

	key1, err0 := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err0 == nil {
		return key1, nil
	}

	key2, err1 := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err1 == nil {
		return key2, nil
	}

	return nil, fmt.Errorf("private key parse error, got '%s' and '%s'", err0, err1)
}

// Parses TOKEN to check intrinsic SubjectConfirmation property to detect if token is of type hok or bearer
func GetSamlTokenType(samlToken string) (string, error) {

	doc := etree.NewDocument()
	doc.ReadFromString(samlToken)

	e := doc.FindElement("//saml2:Assertion/saml2:Subject/saml2:SubjectConfirmation")

	var samlType = e.Attr[0].Value

	if samlType == SAML_HOK_TYPE {
		return SAML_HOK_TOKEN, nil
	} else if samlType == SAML_BEARER_TYPE {
		return SAML_BEARER_TOKEN, nil
	}
	log.Error("Invalid subject confirmation data for saml token")
	return "", errors.New("Invalid subject confirmation data for saml token")
}

// check if token is of type Bearer
func isSamlBearerToken(samlToken string) bool {

	tokenType, er := GetSamlTokenType(samlToken)
	if er != nil {
		return false
	}
	if tokenType == SAML_BEARER_TOKEN {
		return true
	}
	return false
}

func assertInterfaceSliceIsStringSlice(objs []interface{}) ([]string, error) {
	strs := make([]string, len(objs))
	for i := range objs {
		str, ok := objs[i].(string)
		if !ok {
			return nil, fmt.Errorf("Assertion failed - '%v' is not a string", objs[i])
		}
		strs[i] = str
	}
	return strs, nil
}

func contains(arr []string, s string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}
	return false
}
