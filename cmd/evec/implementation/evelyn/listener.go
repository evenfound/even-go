package evelyn

import (
	"io"

	"github.com/evenfound/even-go/node/cmd/evec/implementation/evelyn/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func newListener(sw io.StringWriter) antlr.ParseTreeListener {
	return &tengoEmitterListener{out: sw}
}

// tengoEmitterListener generates Tengo source code
// following syntactic events from the Evelyn parser.
type tengoEmitterListener struct {
	*parser.BaseEvelynListener

	out io.StringWriter
}

func (t *tengoEmitterListener) ExitContractClause(ctx *parser.ContractClauseContext) {
	if _, err := t.out.WriteString("XXX\n"); err != nil {
		panic(err)
	}
}
