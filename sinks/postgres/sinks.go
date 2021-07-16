// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

// Adapted for Orb project, modifications licensed under MPL v. 2.0:
/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package postgres

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/ns1labs/orb/pkg/db"
	"github.com/ns1labs/orb/pkg/errors"
	"github.com/ns1labs/orb/pkg/types"
	"github.com/ns1labs/orb/sinks"
	"go.uber.org/zap"
)

var _ sinks.Repository = (*sinksRepository)(nil)

type sinksRepository struct {
	db     *sqlx.DB
	logger *zap.Logger
}

func NewSinksRepository(db *sqlx.DB, log *zap.Logger) sinks.Repository {
	return &sinksRepository{db: db, logger: log}
}

func (cr sinksRepository) Save(ctx context.Context, sink sinks.Sink) (string, error) {
	q := `INSERT INTO sinks (name, mf_owner_id, metadata)         
			  VALUES (:name, :mf_owner_id, :metadata) RETURNING id`

	if !sink.Name.IsValid() || sink.MFOwnerID == "" {
		return "", errors.ErrMalformedEntity
	}

	dba, err := toDBSink(sink)
	if err != nil {
		return "", errors.Wrap(db.ErrSaveDB, err)
	}

	row, err := cr.db.NamedQueryContext(ctx, q, dba)
	if err != nil {
		pqErr, ok := err.(*pq.Error)
		if ok {
			switch pqErr.Code.Name() {
			case db.ErrInvalid, db.ErrTruncation:
				return "", errors.Wrap(errors.ErrMalformedEntity, err)
			case db.ErrDuplicate:
				return "", errors.Wrap(errors.ErrConflict, err)
			}
		}
		return "", errors.Wrap(db.ErrSaveDB, err)
	}

	defer row.Close()
	row.Next()
	var id string
	if err := row.Scan(&id); err != nil {
		return "", err
	}
	return id, nil

}

type dbSink struct {
	ID        string           `db:"id"`
	Name      types.Identifier `db:"name"`
	MFOwnerID string           `db:"mf_owner_id"`
	Metadata  db.Metadata      `db:"metadata"`
}

func toDBSink(sink sinks.Sink) (dbSink, error) {

	var uID uuid.UUID
	err := uID.Scan(sink.MFOwnerID)
	if err != nil {
		return dbSink{}, errors.Wrap(errors.ErrMalformedEntity, err)
	}

	return dbSink{
		ID:        sink.ID,
		Name:      sink.Name,
		MFOwnerID: uID.String(),
		Metadata:  db.Metadata(sink.Config),
	}, nil

}