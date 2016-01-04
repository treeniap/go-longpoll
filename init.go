// Copyright (c) 2015 Ventu.io, Oleg Sklyar, contributors
// The use of this source code is governed by a MIT style license found in the LICENSE file

package longpoll

import (
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("longpoll")

const (
	no int32 = iota
	yes
)

const Version = 1.1
