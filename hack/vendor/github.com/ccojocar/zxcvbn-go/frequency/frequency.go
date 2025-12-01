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

package frequency

import (
	"encoding/json"
	"log"

	"github.com/ccojocar/zxcvbn-go/data"
)

// List holds a frequency list
type List struct {
	Name string
	List []string
}

// Lists holds all the frequency list in a map
var Lists = make(map[string]List)

func init() {
	maleFilePath := getAsset("data/MaleNames.json")
	femaleFilePath := getAsset("data/FemaleNames.json")
	surnameFilePath := getAsset("data/Surnames.json")
	englishFilePath := getAsset("data/English.json")
	passwordsFilePath := getAsset("data/Passwords.json")

	Lists["MaleNames"] = getStringListFromAsset(maleFilePath, "MaleNames")
	Lists["FemaleNames"] = getStringListFromAsset(femaleFilePath, "FemaleNames")
	Lists["Surname"] = getStringListFromAsset(surnameFilePath, "Surname")
	Lists["English"] = getStringListFromAsset(englishFilePath, "English")
	Lists["Passwords"] = getStringListFromAsset(passwordsFilePath, "Passwords")
}

func getAsset(name string) []byte {
	data, err := data.Asset(name)
	if err != nil {
		panic("Error getting asset " + name)
	}

	return data
}

func getStringListFromAsset(data []byte, name string) List {
	var tempList List
	err := json.Unmarshal(data, &tempList)
	if err != nil {
		log.Fatal(err)
	}
	tempList.Name = name
	return tempList
}
