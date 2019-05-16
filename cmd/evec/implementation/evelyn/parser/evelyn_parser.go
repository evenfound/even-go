// Code generated from Evelyn.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Evelyn
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 11, 39, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 20, 10, 2, 12, 2, 14, 2, 23, 11, 2, 3,
	3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	7, 3, 7, 3, 7, 2, 2, 8, 2, 4, 6, 8, 10, 12, 2, 3, 3, 3, 6, 6, 2, 33, 2,
	14, 3, 2, 2, 2, 4, 24, 3, 2, 2, 2, 6, 27, 3, 2, 2, 2, 8, 29, 3, 2, 2, 2,
	10, 31, 3, 2, 2, 2, 12, 36, 3, 2, 2, 2, 14, 15, 5, 4, 3, 2, 15, 21, 5,
	12, 7, 2, 16, 17, 5, 6, 4, 2, 17, 18, 5, 12, 7, 2, 18, 20, 3, 2, 2, 2,
	19, 16, 3, 2, 2, 2, 20, 23, 3, 2, 2, 2, 21, 19, 3, 2, 2, 2, 21, 22, 3,
	2, 2, 2, 22, 3, 3, 2, 2, 2, 23, 21, 3, 2, 2, 2, 24, 25, 7, 3, 2, 2, 25,
	26, 7, 7, 2, 2, 26, 5, 3, 2, 2, 2, 27, 28, 5, 8, 5, 2, 28, 7, 3, 2, 2,
	2, 29, 30, 5, 10, 6, 2, 30, 9, 3, 2, 2, 2, 31, 32, 7, 4, 2, 2, 32, 33,
	7, 7, 2, 2, 33, 34, 7, 5, 2, 2, 34, 35, 7, 7, 2, 2, 35, 11, 3, 2, 2, 2,
	36, 37, 9, 2, 2, 2, 37, 13, 3, 2, 2, 2, 3, 21,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'contract'", "'var'", "'='", "';'",
}
var symbolicNames = []string{
	"", "", "", "", "", "IDENTIFIER", "WS", "TERMINATOR", "COMMENT", "LINE_COMMENT",
}

var ruleNames = []string{
	"sourceFile", "contractClause", "topLevelDecl", "declaration", "varDecl",
	"eos",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type EvelynParser struct {
	*antlr.BaseParser
}

func NewEvelynParser(input antlr.TokenStream) *EvelynParser {
	this := new(EvelynParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Evelyn.g4"

	return this
}

// EvelynParser tokens.
const (
	EvelynParserEOF          = antlr.TokenEOF
	EvelynParserT__0         = 1
	EvelynParserT__1         = 2
	EvelynParserT__2         = 3
	EvelynParserT__3         = 4
	EvelynParserIDENTIFIER   = 5
	EvelynParserWS           = 6
	EvelynParserTERMINATOR   = 7
	EvelynParserCOMMENT      = 8
	EvelynParserLINE_COMMENT = 9
)

// EvelynParser rules.
const (
	EvelynParserRULE_sourceFile     = 0
	EvelynParserRULE_contractClause = 1
	EvelynParserRULE_topLevelDecl   = 2
	EvelynParserRULE_declaration    = 3
	EvelynParserRULE_varDecl        = 4
	EvelynParserRULE_eos            = 5
)

// ISourceFileContext is an interface to support dynamic dispatch.
type ISourceFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSourceFileContext differentiates from other interfaces.
	IsSourceFileContext()
}

type SourceFileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceFileContext() *SourceFileContext {
	var p = new(SourceFileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EvelynParserRULE_sourceFile
	return p
}

func (*SourceFileContext) IsSourceFileContext() {}

func NewSourceFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceFileContext {
	var p = new(SourceFileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvelynParserRULE_sourceFile

	return p
}

func (s *SourceFileContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceFileContext) ContractClause() IContractClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContractClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IContractClauseContext)
}

func (s *SourceFileContext) AllEos() []IEosContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEosContext)(nil)).Elem())
	var tst = make([]IEosContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEosContext)
		}
	}

	return tst
}

func (s *SourceFileContext) Eos(i int) IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *SourceFileContext) AllTopLevelDecl() []ITopLevelDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITopLevelDeclContext)(nil)).Elem())
	var tst = make([]ITopLevelDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITopLevelDeclContext)
		}
	}

	return tst
}

func (s *SourceFileContext) TopLevelDecl(i int) ITopLevelDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITopLevelDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITopLevelDeclContext)
}

func (s *SourceFileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceFileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceFileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvelynListener); ok {
		listenerT.EnterSourceFile(s)
	}
}

func (s *SourceFileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvelynListener); ok {
		listenerT.ExitSourceFile(s)
	}
}

func (p *EvelynParser) SourceFile() (localctx ISourceFileContext) {
	localctx = NewSourceFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, EvelynParserRULE_sourceFile)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(12)
		p.ContractClause()
	}
	{
		p.SetState(13)
		p.Eos()
	}
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == EvelynParserT__1 {
		{
			p.SetState(14)
			p.TopLevelDecl()
		}
		{
			p.SetState(15)
			p.Eos()
		}

		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IContractClauseContext is an interface to support dynamic dispatch.
type IContractClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContractClauseContext differentiates from other interfaces.
	IsContractClauseContext()
}

type ContractClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContractClauseContext() *ContractClauseContext {
	var p = new(ContractClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EvelynParserRULE_contractClause
	return p
}

func (*ContractClauseContext) IsContractClauseContext() {}

func NewContractClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContractClauseContext {
	var p = new(ContractClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvelynParserRULE_contractClause

	return p
}

func (s *ContractClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *ContractClauseContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(EvelynParserIDENTIFIER, 0)
}

func (s *ContractClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContractClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContractClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvelynListener); ok {
		listenerT.EnterContractClause(s)
	}
}

func (s *ContractClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvelynListener); ok {
		listenerT.ExitContractClause(s)
	}
}

func (p *EvelynParser) ContractClause() (localctx IContractClauseContext) {
	localctx = NewContractClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, EvelynParserRULE_contractClause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(22)
		p.Match(EvelynParserT__0)
	}
	{
		p.SetState(23)
		p.Match(EvelynParserIDENTIFIER)
	}

	return localctx
}

// ITopLevelDeclContext is an interface to support dynamic dispatch.
type ITopLevelDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTopLevelDeclContext differentiates from other interfaces.
	IsTopLevelDeclContext()
}

type TopLevelDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopLevelDeclContext() *TopLevelDeclContext {
	var p = new(TopLevelDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EvelynParserRULE_topLevelDecl
	return p
}

func (*TopLevelDeclContext) IsTopLevelDeclContext() {}

func NewTopLevelDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopLevelDeclContext {
	var p = new(TopLevelDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvelynParserRULE_topLevelDecl

	return p
}

func (s *TopLevelDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *TopLevelDeclContext) Declaration() IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *TopLevelDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopLevelDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TopLevelDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvelynListener); ok {
		listenerT.EnterTopLevelDecl(s)
	}
}

func (s *TopLevelDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvelynListener); ok {
		listenerT.ExitTopLevelDecl(s)
	}
}

func (p *EvelynParser) TopLevelDecl() (localctx ITopLevelDeclContext) {
	localctx = NewTopLevelDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, EvelynParserRULE_topLevelDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(25)
		p.Declaration()
	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EvelynParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvelynParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) VarDecl() IVarDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarDeclContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvelynListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvelynListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (p *EvelynParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, EvelynParserRULE_declaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(27)
		p.VarDecl()
	}

	return localctx
}

// IVarDeclContext is an interface to support dynamic dispatch.
type IVarDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarDeclContext differentiates from other interfaces.
	IsVarDeclContext()
}

type VarDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDeclContext() *VarDeclContext {
	var p = new(VarDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EvelynParserRULE_varDecl
	return p
}

func (*VarDeclContext) IsVarDeclContext() {}

func NewVarDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclContext {
	var p = new(VarDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvelynParserRULE_varDecl

	return p
}

func (s *VarDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(EvelynParserIDENTIFIER)
}

func (s *VarDeclContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(EvelynParserIDENTIFIER, i)
}

func (s *VarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvelynListener); ok {
		listenerT.EnterVarDecl(s)
	}
}

func (s *VarDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvelynListener); ok {
		listenerT.ExitVarDecl(s)
	}
}

func (p *EvelynParser) VarDecl() (localctx IVarDeclContext) {
	localctx = NewVarDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, EvelynParserRULE_varDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(29)
		p.Match(EvelynParserT__1)
	}
	{
		p.SetState(30)
		p.Match(EvelynParserIDENTIFIER)
	}
	{
		p.SetState(31)
		p.Match(EvelynParserT__2)
	}
	{
		p.SetState(32)
		p.Match(EvelynParserIDENTIFIER)
	}

	return localctx
}

// IEosContext is an interface to support dynamic dispatch.
type IEosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEosContext differentiates from other interfaces.
	IsEosContext()
}

type EosContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEosContext() *EosContext {
	var p = new(EosContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EvelynParserRULE_eos
	return p
}

func (*EosContext) IsEosContext() {}

func NewEosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EosContext {
	var p = new(EosContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvelynParserRULE_eos

	return p
}

func (s *EosContext) GetParser() antlr.Parser { return s.parser }

func (s *EosContext) EOF() antlr.TerminalNode {
	return s.GetToken(EvelynParserEOF, 0)
}

func (s *EosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvelynListener); ok {
		listenerT.EnterEos(s)
	}
}

func (s *EosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvelynListener); ok {
		listenerT.ExitEos(s)
	}
}

func (p *EvelynParser) Eos() (localctx IEosContext) {
	localctx = NewEosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, EvelynParserRULE_eos)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(34)
		_la = p.GetTokenStream().LA(1)

		if !(_la == EvelynParserEOF || _la == EvelynParserT__3) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
