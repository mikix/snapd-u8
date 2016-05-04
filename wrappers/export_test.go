// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2016 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package wrappers

import (
	"time"
)

// some internal helper exposed for testing
var (
	// binaries
	GenerateSnapBinaryWrapper = generateSnapBinaryWrapper

	// services
	GenerateSnapServiceFile = generateSnapServiceFile
	GenerateSnapSocketFile  = generateSnapSocketFile
)

func MockKillWait(wait time.Duration) (restore func()) {
	oldKillWait := killWait
	killWait = wait
	return func() {
		killWait = oldKillWait
	}
}