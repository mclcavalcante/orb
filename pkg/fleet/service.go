// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

// Adapted for Orb project, modifications licensed under MPL v. 2.0:
/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package fleet

import (
	"context"
	"github.com/mainflux/mainflux"
	mfsdk "github.com/mainflux/mainflux/pkg/sdk/go"
	"github.com/ns1labs/orb/pkg/errors"
	"time"
)

var (
	// ErrNotFound indicates a non-existent entity request.
	ErrNotFound = errors.New("non-existent entity")

	// ErrConflict indicates that entity already exists.
	ErrConflict = errors.New("entity already exists")

	// ErrMalformedEntity indicates malformed entity specification.
	ErrMalformedEntity = errors.New("malformed entity specification")

	// ErrUnauthorizedAccess indicates missing or invalid credentials provided
	// when accessing a protected resource.
	ErrUnauthorizedAccess = errors.New("missing or invalid credentials provided")

	// ErrScanMetadata indicates problem with metadata in db.
	ErrScanMetadata = errors.New("failed to scan metadata")

	ErrCreateSelector = errors.New("failed to create selector")
)

// A flat kv pair object
type Tags map[string]interface{}

// Maybe a full object hierarchy
type Metadata map[string]interface{}

type Service interface {
	AgentService
	SelectorService
}

var _ Service = (*fleetService)(nil)

type fleetService struct {
	// for AuthN/AuthZ
	auth mainflux.AuthServiceClient
	// for Thing manipulation
	mfsdk mfsdk.SDK
	// Agents and Selectors
	agentRepo    AgentRepository
	selectorRepo SelectorRepository
}

func (svc fleetService) identify(token string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := svc.auth.Identify(ctx, &mainflux.Token{Value: token})
	if err != nil {
		return "", errors.Wrap(ErrUnauthorizedAccess, err)
	}

	return res.GetId(), nil
}

func (svc fleetService) CreateSelector(ctx context.Context, token string, s Selector) (Selector, error) {
	mfOwnerID, err := svc.identify(token)
	if err != nil {
		return Selector{}, err
	}

	s.MFOwnerID = mfOwnerID

	err = svc.selectorRepo.Save(ctx, s)
	if err != nil {
		return Selector{}, errors.Wrap(ErrCreateSelector, err)
	}

	return s, nil
}

func (svc fleetService) CreateAgent(ctx context.Context, token string, a Agent) (Agent, error) {
	panic("implement me")
}

func NewFleetService(auth mainflux.AuthServiceClient, agentRepo AgentRepository, selectorRepo SelectorRepository, mfsdk mfsdk.SDK) Service {
	return &fleetService{
		auth:         auth,
		agentRepo:    agentRepo,
		selectorRepo: selectorRepo,
		mfsdk:        mfsdk,
	}
}
