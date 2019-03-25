// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package prometheus

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "prometheus", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "eJy0krGum0AQRXu+4oo6jw+gSJP+KVLKKEJruMDKyw7ZGZ7kv48INsYEy03elMPsPYfZfcOZlxJjkoHWc9IMMG+BJfLvazPPgIZaJz+al1jiawYAP8yZopYQWBsbtEkG3E8VGaC9JKtqia3vSrQuKDMgMdApS3RunqGZj52W+JmrhvwL8t5szH9lQOsZGi3/4t4Q3cCd7Fx2GeesJNN47RzIPmYteY8et+5C0fnn1u4R5CloqWU77iSTwXpuFgNl+mAqNuOryYnmNv2989YwivnW126G68PEM9sXxnO9b0KXHRS7mSOlrdbviROrwNhZ/8/QzSxI7A4+vpCb69uUEqMtGCyYveJdpkkyjmw+weN9Gk5MkPbGuBrxg3G/tM3LralKLWRkrNrm+NoOpF7d2iozB6P1gesRSU9s1CS5jkXdT/GslUk1MqlX++9SAwdJFywgWO8MLnF+wLjQcMWygQkar+ci+xMAAP//EEQ9qA=="
}