package v10

import(
	"github.com/ksang/goflow/openflow"
)
// Hello share the same interface and struct as echo
func NewHello(xid uint32) openflow.Echo {
	return &echo{
		Message: openflow.NewMessage(openflow.OF10_VERSION, OFPT_HELLO, xid),
	}
}