// Copyright (c) 2016 Intel Corporation. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package mntent

import (
	"github.com/edwardsp/go-lustre"
	"github.com/pkg/errors"
)

// GetMounted returns a slide of filesystem entries from
// the mounted fs table.
func GetMounted() (Entries, error) {
	return nil, errors.Wrap(lustre.ErrUnimplemented, "GetMounted")
}
