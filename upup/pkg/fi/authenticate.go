/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package fi

// Authenticator generates authentication credentials for requests.
type Authenticator interface {
	CreateToken(body []byte) (string, error)
}

// VerifyResult is the result of a successfully verified request.
type VerifyResult struct {
	// Nodename is the name that this node is authorized to use.
	NodeName string

	// InstanceGroupName is the name of the kops InstanceGroup this node is a member of.
	InstanceGroupName string

	// CertificateNames is the names the node is authorized to use for certificates.
	CertificateNames []string
}

// Verifier verifies authentication credentials for requests.
type Verifier interface {
	VerifyToken(token string, body []byte) (*VerifyResult, error)
}
