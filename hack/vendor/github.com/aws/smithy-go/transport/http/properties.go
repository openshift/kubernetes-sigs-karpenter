/*
Copyright The Kubernetes Authors.

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

package http

import smithy "github.com/aws/smithy-go"

type (
	sigV4SigningNameKey   struct{}
	sigV4SigningRegionKey struct{}

	sigV4ASigningNameKey    struct{}
	sigV4ASigningRegionsKey struct{}

	isUnsignedPayloadKey     struct{}
	disableDoubleEncodingKey struct{}
)

// GetSigV4SigningName gets the signing name from Properties.
func GetSigV4SigningName(p *smithy.Properties) (string, bool) {
	v, ok := p.Get(sigV4SigningNameKey{}).(string)
	return v, ok
}

// SetSigV4SigningName sets the signing name on Properties.
func SetSigV4SigningName(p *smithy.Properties, name string) {
	p.Set(sigV4SigningNameKey{}, name)
}

// GetSigV4SigningRegion gets the signing region from Properties.
func GetSigV4SigningRegion(p *smithy.Properties) (string, bool) {
	v, ok := p.Get(sigV4SigningRegionKey{}).(string)
	return v, ok
}

// SetSigV4SigningRegion sets the signing region on Properties.
func SetSigV4SigningRegion(p *smithy.Properties, region string) {
	p.Set(sigV4SigningRegionKey{}, region)
}

// GetSigV4ASigningName gets the v4a signing name from Properties.
func GetSigV4ASigningName(p *smithy.Properties) (string, bool) {
	v, ok := p.Get(sigV4ASigningNameKey{}).(string)
	return v, ok
}

// SetSigV4ASigningName sets the signing name on Properties.
func SetSigV4ASigningName(p *smithy.Properties, name string) {
	p.Set(sigV4ASigningNameKey{}, name)
}

// GetSigV4ASigningRegion gets the v4a signing region set from Properties.
func GetSigV4ASigningRegions(p *smithy.Properties) ([]string, bool) {
	v, ok := p.Get(sigV4ASigningRegionsKey{}).([]string)
	return v, ok
}

// SetSigV4ASigningRegions sets the v4a signing region set on Properties.
func SetSigV4ASigningRegions(p *smithy.Properties, regions []string) {
	p.Set(sigV4ASigningRegionsKey{}, regions)
}

// GetIsUnsignedPayload gets whether the payload is unsigned from Properties.
func GetIsUnsignedPayload(p *smithy.Properties) (bool, bool) {
	v, ok := p.Get(isUnsignedPayloadKey{}).(bool)
	return v, ok
}

// SetIsUnsignedPayload sets whether the payload is unsigned on Properties.
func SetIsUnsignedPayload(p *smithy.Properties, isUnsignedPayload bool) {
	p.Set(isUnsignedPayloadKey{}, isUnsignedPayload)
}

// GetDisableDoubleEncoding gets whether the payload is unsigned from Properties.
func GetDisableDoubleEncoding(p *smithy.Properties) (bool, bool) {
	v, ok := p.Get(disableDoubleEncodingKey{}).(bool)
	return v, ok
}

// SetDisableDoubleEncoding sets whether the payload is unsigned on Properties.
func SetDisableDoubleEncoding(p *smithy.Properties, disableDoubleEncoding bool) {
	p.Set(disableDoubleEncodingKey{}, disableDoubleEncoding)
}
