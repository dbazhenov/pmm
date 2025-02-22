// Copyright (C) 2017 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package models

import (
	"time"

	"gopkg.in/reform.v1"
)

//go:generate ../../bin/reform

// OnboardingSystemTip represents tip for user which can be completed once
// time per system.
//
//reform:onboarding_system_tips
type OnboardingSystemTip struct {
	ID          int64 `reform:"id,pk"`
	IsCompleted bool  `reform:"is_completed"`

	CreatedAt time.Time `reform:"created_at"`
	UpdatedAt time.Time `reform:"updated_at"`
}

// BeforeInsert implements reform.BeforeInserter interface.
func (t *OnboardingSystemTip) BeforeInsert() error {
	now := Now()
	t.CreatedAt = now
	t.UpdatedAt = now

	return nil
}

// BeforeUpdate implements reform.BeforeUpdater interface.
func (t *OnboardingSystemTip) BeforeUpdate() error {
	t.UpdatedAt = Now()

	return nil
}

// check interfaces.
var (
	_ reform.BeforeInserter = (*OnboardingSystemTip)(nil)
	_ reform.BeforeUpdater  = (*OnboardingSystemTip)(nil)
)
