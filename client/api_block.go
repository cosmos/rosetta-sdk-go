// Copyright 2022 Coinbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Generated by: OpenAPI Generator (https://openapi-generator.tech)

package client

import (
	_context "context"
	"fmt"
	_ioutil "io/ioutil"
	_nethttp "net/http"

	"github.com/cosmos/rosetta-sdk-go/types"
)

// Linger please
var (
	_ _context.Context
)

// BlockAPIService BlockAPI service
type BlockAPIService service

// Block Get a block by its Block Identifier. If transactions are returned in the same call to the
// node as fetching the block, the response should include these transactions in the Block object.
// If not, an array of Transaction Identifiers should be returned so /block/transaction fetches can
// be done to get all transaction information. When requesting a block by the hash component of the
// BlockIdentifier, this request MUST be idempotent: repeated invocations for the same
// hash-identified block must return the exact same block contents. No such restriction is imposed
// when requesting a block by height, given that a chain reorg event might cause the specific block
// at height n to be set to a different one.
func (a *BlockAPIService) Block(
	ctx _context.Context,
	blockRequest *types.BlockRequest,
) (*types.BlockResponse, *types.Error, error) {
	var (
		localVarPostBody interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/block"
	localVarHeaderParams := make(map[string]string)

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = blockRequest

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarPostBody, localVarHeaderParams)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to prepare request: %w", err)
	}

	localVarHTTPResponse, err := a.client.callAPI(ctx, r)
	if err != nil || localVarHTTPResponse == nil {
		return nil, nil, fmt.Errorf("failed to call API: %w", err)
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	defer localVarHTTPResponse.Body.Close()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read response: %w", err)
	}

	switch localVarHTTPResponse.StatusCode {
	case _nethttp.StatusOK:
		var v types.BlockResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, nil, fmt.Errorf(
				"failed to decode when hit status code 200, response body %s: %w",
				string(localVarBody),
				err,
			)
		}

		return &v, nil, nil
	case _nethttp.StatusInternalServerError:
		var v types.Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, nil, fmt.Errorf(
				"failed to decode when hit status code 500, response body %s: %w",
				string(localVarBody),
				err,
			)
		}

		return nil, &v, fmt.Errorf("error %+v", v)
	case _nethttp.StatusBadGateway,
		_nethttp.StatusServiceUnavailable,
		_nethttp.StatusGatewayTimeout,
		_nethttp.StatusRequestTimeout:
		return nil, nil, fmt.Errorf(
			"status code %d, response body %s: %w",
			localVarHTTPResponse.StatusCode,
			string(localVarBody),
			ErrRetriable,
		)
	default:
		return nil, nil, fmt.Errorf(
			"invalid status code %d, response body %s",
			localVarHTTPResponse.StatusCode,
			string(localVarBody),
		)
	}
}

// BlockTransaction Get a transaction in a block by its Transaction Identifier. This endpoint should
// only be used when querying a node for a block does not return all transactions contained within
// it. All transactions returned by this endpoint must be appended to any transactions returned by
// the /block method by consumers of this data. Fetching a transaction by hash is considered an
// Explorer Method (which is classified under the Future Work section). This method can be used to
// let consumers to paginate results when the  block trasactions count is too big to be returned in
// a single BlockResponse. Calling this endpoint requires reference to a BlockIdentifier because
// transaction parsing can change depending on which block contains the transaction. For example, in
// Bitcoin it is necessary to know which block contains a transaction to determine the destination
// of fee payments. Without specifying a block identifier, the node would have to infer which block
// to use (which could change during a re-org). Implementations that require fetching previous
// transactions to populate the response (ex: Previous UTXOs in Bitcoin) may find it useful to run a
// cache within the Rosetta server in the /data directory (on a path that does not conflict with the
// node).
func (a *BlockAPIService) BlockTransaction(
	ctx _context.Context,
	blockTransactionRequest *types.BlockTransactionRequest,
) (*types.BlockTransactionResponse, *types.Error, error) {
	var (
		localVarPostBody interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/block/transaction"
	localVarHeaderParams := make(map[string]string)

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = blockTransactionRequest

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarPostBody, localVarHeaderParams)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to prepare request: %w", err)
	}

	localVarHTTPResponse, err := a.client.callAPI(ctx, r)
	if err != nil || localVarHTTPResponse == nil {
		return nil, nil, fmt.Errorf("failed to call API: %w", err)
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	defer localVarHTTPResponse.Body.Close()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read response: %w", err)
	}

	switch localVarHTTPResponse.StatusCode {
	case _nethttp.StatusOK:
		var v types.BlockTransactionResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, nil, fmt.Errorf(
				"failed to decode when hit status code 200, response body %s: %w",
				string(localVarBody),
				err,
			)
		}

		return &v, nil, nil
	case _nethttp.StatusInternalServerError:
		var v types.Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, nil, fmt.Errorf(
				"failed to decode when hit status code 500, response body %s: %w",
				string(localVarBody),
				err,
			)
		}

		return nil, &v, fmt.Errorf("error %+v", v)
	case _nethttp.StatusBadGateway,
		_nethttp.StatusServiceUnavailable,
		_nethttp.StatusGatewayTimeout,
		_nethttp.StatusRequestTimeout:
		return nil, nil, fmt.Errorf(
			"status code %d, response body %s: %w",
			localVarHTTPResponse.StatusCode,
			string(localVarBody),
			ErrRetriable,
		)
	default:
		return nil, nil, fmt.Errorf(
			"invalid status code %d, response body %s",
			localVarHTTPResponse.StatusCode,
			string(localVarBody),
		)
	}
}
