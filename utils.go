/**
 * Copyright 2024-present Coinbase Global, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package intx

import (
	"fmt"
)

func appendQueryParam(queryParams, key, value string) string {
	if queryParams == "" {
		return "?" + key + "=" + value
	} else {
		return queryParams + "&" + key + "=" + value
	}
}

func appendPaginationParams(v string, p *PaginationParams) string {
	if p == nil {
		return v
	}

	if p.RefDatetime != "" {
		v = appendQueryParam(v, "ref_datetime", p.RefDatetime)
	}

	if p.ResultLimit > 0 {
		v = appendQueryParam(v, "result_limit", fmt.Sprint(p.ResultLimit))
	}

	if p.ResultOffset > 0 {
		v = appendQueryParam(v, "result_offset", fmt.Sprint(p.ResultOffset))
	}

	return v
}
