package frbnc

import "github.com/dwdwow/cex/bnc"

type VIPPortmarAcctSimple struct {
	user   *bnc.User
	cfg    VIPPortmarAccountConfig
	chAcct chan VIPPortmarAcctWatcherMsg
}

func (v *VIPPortmarAcctSimple) start() {
	for {
		msg := <-v.chAcct
		if msg.Err != nil {

		}
	}
}

func (v *VIPPortmarAcctSimple) handle(acct *VIPPortmarAccount) {
}

func (v *VIPPortmarAcctSimple) handleLowMMR(acct *VIPPortmarAccount) {
	if acct.PortmarAccountInformation.UniMMR > v.cfg.MinUniMMR {
		return
	}
}

func (v *VIPPortmarAcctSimple) handleHighLtv(acct *VIPPortmarAccount) {}
