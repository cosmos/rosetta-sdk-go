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

package server

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/cosmos/rosetta-sdk-go/asserter"
	"github.com/cosmos/rosetta-sdk-go/types"
)

// A ConstructionAPIController binds http requests to an api service and writes the service results
// to the http response
type ConstructionAPIController struct {
	service  ConstructionAPIServicer
	asserter *asserter.Asserter
}

// NewConstructionAPIController creates a default api controller
func NewConstructionAPIController(
	s ConstructionAPIServicer,
	asserter *asserter.Asserter,
) Router {
	return &ConstructionAPIController{
		service:  s,
		asserter: asserter,
	}
}

// Routes returns all of the api route for the ConstructionAPIController
func (c *ConstructionAPIController) Routes() Routes {
	return Routes{
		{
			"ConstructionCombine",
			strings.ToUpper("Post"),
			"/construction/combine",
			c.ConstructionCombine,
		},
		{
			"ConstructionDerive",
			strings.ToUpper("Post"),
			"/construction/derive",
			c.ConstructionDerive,
		},
		{
			"ConstructionHash",
			strings.ToUpper("Post"),
			"/construction/hash",
			c.ConstructionHash,
		},
		{
			"ConstructionMetadata",
			strings.ToUpper("Post"),
			"/construction/metadata",
			c.ConstructionMetadata,
		},
		{
			"ConstructionParse",
			strings.ToUpper("Post"),
			"/construction/parse",
			c.ConstructionParse,
		},
		{
			"ConstructionPayloads",
			strings.ToUpper("Post"),
			"/construction/payloads",
			c.ConstructionPayloads,
		},
		{
			"ConstructionPreprocess",
			strings.ToUpper("Post"),
			"/construction/preprocess",
			c.ConstructionPreprocess,
		},
		{
			"ConstructionSubmit",
			strings.ToUpper("Post"),
			"/construction/submit",
			c.ConstructionSubmit,
		},
	}
}

// ConstructionCombine - Create Network Transaction from Signatures
func (c *ConstructionAPIController) ConstructionCombine(w http.ResponseWriter, r *http.Request) {
	constructionCombineRequest := &types.ConstructionCombineRequest{}
	if err := json.NewDecoder(r.Body).Decode(&constructionCombineRequest); err != nil {
		EncodeJSONResponse(&types.Error{
			Message: err.Error(),
		}, http.StatusInternalServerError, w)

		return
	}

	// Assert that ConstructionCombineRequest is correct
	if err := c.asserter.ConstructionCombineRequest(constructionCombineRequest); err != nil {
		EncodeJSONResponse(&types.Error{
			Message: err.Error(),
		}, http.StatusInternalServerError, w)

		return
	}

	result, serviceErr := c.service.ConstructionCombine(r.Context(), constructionCombineRequest)
	if serviceErr != nil {
		EncodeJSONResponse(serviceErr, http.StatusInternalServerError, w)

		return
	}

	EncodeJSONResponse(result, http.StatusOK, w)
}

// ConstructionDerive - Derive an AccountIdentifier from a PublicKey
func (c *ConstructionAPIController) ConstructionDerive(w http.ResponseWriter, r *http.Request) {
	constructionDeriveRequest := &types.ConstructionDeriveRequest{}
	if err := json.NewDecoder(r.Body).Decode(&constructionDeriveRequest); err != nil {
		EncodeJSONResponse(&types.Error{
			Message: err.Error(),
		}, http.StatusInternalServerError, w)

		return
	}

	// Assert that ConstructionDeriveRequest is correct
	if err := c.asserter.ConstructionDeriveRequest(constructionDeriveRequest); err != nil {
		EncodeJSONResponse(&types.Error{
			Message: err.Error(),
		}, http.StatusInternalServerError, w)

		return
	}

	result, serviceErr := c.service.ConstructionDerive(r.Context(), constructionDeriveRequest)
	if serviceErr != nil {
		EncodeJSONResponse(serviceErr, http.StatusInternalServerError, w)

		return
	}

	EncodeJSONResponse(result, http.StatusOK, w)
}

// ConstructionHash - Get the Hash of a Signed Transaction
func (c *ConstructionAPIController) ConstructionHash(w http.ResponseWriter, r *http.Request) {
	constructionHashRequest := &types.ConstructionHashRequest{}
	if err := json.NewDecoder(r.Body).Decode(&constructionHashRequest); err != nil {
		EncodeJSONResponse(&types.Error{
			Message: err.Error(),
		}, http.StatusInternalServerError, w)

		return
	}

	// Assert that ConstructionHashRequest is correct
	if err := c.asserter.ConstructionHashRequest(constructionHashRequest); err != nil {
		EncodeJSONResponse(&types.Error{
			Message: err.Error(),
		}, http.StatusInternalServerError, w)

		return
	}

	result, serviceErr := c.service.ConstructionHash(r.Context(), constructionHashRequest)
	if serviceErr != nil {
		EncodeJSONResponse(serviceErr, http.StatusInternalServerError, w)

		return
	}

	EncodeJSONResponse(result, http.StatusOK, w)
}

// ConstructionMetadata - Get Metadata for Transaction Construction
func (c *ConstructionAPIController) ConstructionMetadata(w http.ResponseWriter, r *http.Request) {
	constructionMetadataRequest := &types.ConstructionMetadataRequest{}
	if err := json.NewDecoder(r.Body).Decode(&constructionMetadataRequest); err != nil {
		EncodeJSONResponse(&types.Error{
			Message: err.Error(),
		}, http.StatusInternalServerError, w)

		return
	}

	// Assert that ConstructionMetadataRequest is correct
	if err := c.asserter.ConstructionMetadataRequest(constructionMetadataRequest); err != nil {
		EncodeJSONResponse(&types.Error{
			Message: err.Error(),
		}, http.StatusInternalServerError, w)

		return
	}

	result, serviceErr := c.service.ConstructionMetadata(r.Context(), constructionMetadataRequest)
	if serviceErr != nil {
		EncodeJSONResponse(serviceErr, http.StatusInternalServerError, w)

		return
	}

	EncodeJSONResponse(result, http.StatusOK, w)
}

// ConstructionParse - Parse a Transaction
func (c *ConstructionAPIController) ConstructionParse(w http.ResponseWriter, r *http.Request) {
	constructionParseRequest := &types.ConstructionParseRequest{}
	if err := json.NewDecoder(r.Body).Decode(&constructionParseRequest); err != nil {
		EncodeJSONResponse(&types.Error{
			Message: err.Error(),
		}, http.StatusInternalServerError, w)

		return
	}

	// Assert that ConstructionParseRequest is correct
	if err := c.asserter.ConstructionParseRequest(constructionParseRequest); err != nil {
		EncodeJSONResponse(&types.Error{
			Message: err.Error(),
		}, http.StatusInternalServerError, w)

		return
	}

	result, serviceErr := c.service.ConstructionParse(r.Context(), constructionParseRequest)
	if serviceErr != nil {
		EncodeJSONResponse(serviceErr, http.StatusInternalServerError, w)

		return
	}

	EncodeJSONResponse(result, http.StatusOK, w)
}

// ConstructionPayloads - Generate an Unsigned Transaction and Signing Payloads
func (c *ConstructionAPIController) ConstructionPayloads(w http.ResponseWriter, r *http.Request) {
	constructionPayloadsRequest := &types.ConstructionPayloadsRequest{}
	if err := json.NewDecoder(r.Body).Decode(&constructionPayloadsRequest); err != nil {
		EncodeJSONResponse(&types.Error{
			Message: err.Error(),
		}, http.StatusInternalServerError, w)

		return
	}

	// Assert that ConstructionPayloadsRequest is correct
	if err := c.asserter.ConstructionPayloadsRequest(constructionPayloadsRequest); err != nil {
		EncodeJSONResponse(&types.Error{
			Message: err.Error(),
		}, http.StatusInternalServerError, w)

		return
	}

	result, serviceErr := c.service.ConstructionPayloads(r.Context(), constructionPayloadsRequest)
	if serviceErr != nil {
		EncodeJSONResponse(serviceErr, http.StatusInternalServerError, w)

		return
	}

	EncodeJSONResponse(result, http.StatusOK, w)
}

// ConstructionPreprocess - Create a Request to Fetch Metadata
func (c *ConstructionAPIController) ConstructionPreprocess(w http.ResponseWriter, r *http.Request) {
	constructionPreprocessRequest := &types.ConstructionPreprocessRequest{}
	if err := json.NewDecoder(r.Body).Decode(&constructionPreprocessRequest); err != nil {
		EncodeJSONResponse(&types.Error{
			Message: err.Error(),
		}, http.StatusInternalServerError, w)

		return
	}

	// Assert that ConstructionPreprocessRequest is correct
	if err := c.asserter.ConstructionPreprocessRequest(constructionPreprocessRequest); err != nil {
		EncodeJSONResponse(&types.Error{
			Message: err.Error(),
		}, http.StatusInternalServerError, w)

		return
	}

	result, serviceErr := c.service.ConstructionPreprocess(
		r.Context(),
		constructionPreprocessRequest,
	)
	if serviceErr != nil {
		EncodeJSONResponse(serviceErr, http.StatusInternalServerError, w)

		return
	}

	EncodeJSONResponse(result, http.StatusOK, w)
}

// ConstructionSubmit - Submit a Signed Transaction
func (c *ConstructionAPIController) ConstructionSubmit(w http.ResponseWriter, r *http.Request) {
	constructionSubmitRequest := &types.ConstructionSubmitRequest{}
	if err := json.NewDecoder(r.Body).Decode(&constructionSubmitRequest); err != nil {
		EncodeJSONResponse(&types.Error{
			Message: err.Error(),
		}, http.StatusInternalServerError, w)

		return
	}

	// Assert that ConstructionSubmitRequest is correct
	if err := c.asserter.ConstructionSubmitRequest(constructionSubmitRequest); err != nil {
		EncodeJSONResponse(&types.Error{
			Message: err.Error(),
		}, http.StatusInternalServerError, w)

		return
	}

	result, serviceErr := c.service.ConstructionSubmit(r.Context(), constructionSubmitRequest)
	if serviceErr != nil {
		EncodeJSONResponse(serviceErr, http.StatusInternalServerError, w)

		return
	}

	EncodeJSONResponse(result, http.StatusOK, w)
}
