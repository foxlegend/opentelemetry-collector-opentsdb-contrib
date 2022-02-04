package opentsdbreceiver

import (
	"github.com/reiver/go-telnet"
)

type OpenTSDBTelnetHandler struct{}

func (h OpenTSDBTelnetHandler) ServeTELNET(ctx telnet.Context, w telnet.Writer, r telnet.Reader) {
	var buffer [1]byte
	p := buffer[:]

	for {
		n, err := r.Read(p)

		if n > 0 {
			w.Write(p)
		}

		if nil != err {
			break
		}
	}
}
