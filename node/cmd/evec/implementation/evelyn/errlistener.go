package evelyn

import (
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// simpleErrorListener stores syntax error(s) instead of printing.
type simpleErrorListener struct {
	messages []string
}

func newErrorListener() *simpleErrorListener {
	return &simpleErrorListener{
		messages: []string{},
	}
}

//func newErrorListener() antlr.ErrorListener {
//	return antlr.NewDiagnosticErrorListener(true)
//}

func (s *simpleErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	s.messages = append(s.messages, "line "+strconv.Itoa(line)+":"+strconv.Itoa(column)+" "+msg)
}

func (s *simpleErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
}

func (s *simpleErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
}

func (s *simpleErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
}

func (s simpleErrorListener) Empty() bool {
	return len(s.messages) == 0
}

func (s simpleErrorListener) FirstMessage() string {
	return s.messages[0]
}
