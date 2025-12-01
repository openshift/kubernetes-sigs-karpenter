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

package sprig

import (
	"fmt"
	"net/url"
	"reflect"
)

func dictGetOrEmpty(dict map[string]interface{}, key string) string {
	value, ok := dict[key]
	if !ok {
		return ""
	}
	tp := reflect.TypeOf(value).Kind()
	if tp != reflect.String {
		panic(fmt.Sprintf("unable to parse %s key, must be of type string, but %s found", key, tp.String()))
	}
	return reflect.ValueOf(value).String()
}

// parses given URL to return dict object
func urlParse(v string) map[string]interface{} {
	dict := map[string]interface{}{}
	parsedURL, err := url.Parse(v)
	if err != nil {
		panic(fmt.Sprintf("unable to parse url: %s", err))
	}
	dict["scheme"] = parsedURL.Scheme
	dict["host"] = parsedURL.Host
	dict["hostname"] = parsedURL.Hostname()
	dict["path"] = parsedURL.Path
	dict["query"] = parsedURL.RawQuery
	dict["opaque"] = parsedURL.Opaque
	dict["fragment"] = parsedURL.Fragment
	if parsedURL.User != nil {
		dict["userinfo"] = parsedURL.User.String()
	} else {
		dict["userinfo"] = ""
	}

	return dict
}

// join given dict to URL string
func urlJoin(d map[string]interface{}) string {
	resURL := url.URL{
		Scheme:   dictGetOrEmpty(d, "scheme"),
		Host:     dictGetOrEmpty(d, "host"),
		Path:     dictGetOrEmpty(d, "path"),
		RawQuery: dictGetOrEmpty(d, "query"),
		Opaque:   dictGetOrEmpty(d, "opaque"),
		Fragment: dictGetOrEmpty(d, "fragment"),
	}
	userinfo := dictGetOrEmpty(d, "userinfo")
	var user *url.Userinfo
	if userinfo != "" {
		tempURL, err := url.Parse(fmt.Sprintf("proto://%s@host", userinfo))
		if err != nil {
			panic(fmt.Sprintf("unable to parse userinfo in dict: %s", err))
		}
		user = tempURL.User
	}

	resURL.User = user
	return resURL.String()
}
