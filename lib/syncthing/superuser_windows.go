// Copyright (C) 2017 The Syncthing Authors.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at https://mozilla.org/MPL/2.0/.

package syncthing

import "syscall"

// https://msdn.microsoft.com/en-us/library/windows/desktop/aa379649(v=vs.85).aspx
const securityLocalSystemRID = "S-1-5-18"

func isSuperUser() bool {
	tok, err := syscall.OpenCurrentProcessToken()
	if err != nil {
		l.Debugln("OpenCurrentProcessToken:", err)
		return false
	}
	defer tok.Close()

	user, err := tok.GetTokenUser()
	if err != nil {
		l.Debugln("GetTokenUser:", err)
		return false
	}

	if user.User.Sid == nil {
		l.Debugln("sid is nil")
		return false
	}

	sid, err := user.User.Sid.String()
	if err != nil {
		l.Debugln("Sid.String():", err)
		return false
	}

	l.Debugf("SID: %q", sid)
	return sid == securityLocalSystemRID
}
