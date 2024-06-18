// Copyright (c) 2022-2024 Focela Technologies, All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package multierr

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
// Note that do not use pointer as its receiver here.
func (err Error) MarshalJSON() ([]byte, error) {
	return []byte(`"` + err.Error() + `"`), nil
}
