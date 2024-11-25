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

package assets

import (
	"context"

	"github.com/coinbase-samples/intx-sdk-go/client"
)

type AssetsService interface {
	GetAsset(ctx context.Context, request *GetAssetRequest) (*GetAssetResponse, error)
	GetSupportedNetworks(ctx context.Context, request *GetSupportedNetworksRequest) (*GetSupportedNetworksResponse, error)
	ListAssets(ctx context.Context, request *ListAssetsRequest) (*ListAssetsResponse, error)
}

func NewAssetsService(c client.RestClient) AssetsService {
	return &assetsServiceImpl{client: c}
}

type assetsServiceImpl struct {
	client client.RestClient
}