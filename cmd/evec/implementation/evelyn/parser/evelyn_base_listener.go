// Code generated from Evelyn.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Evelyn
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseEvelynListener is a complete listener for a parse tree produced by EvelynParser.
type BaseEvelynListener struct{}

var _ EvelynListener = &BaseEvelynListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseEvelynListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseEvelynListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseEvelynListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseEvelynListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSourceFile is called when production sourceFile is entered.
func (s *BaseEvelynListener) EnterSourceFile(ctx *SourceFileContext) {}

// ExitSourceFile is called when production sourceFile is exited.
func (s *BaseEvelynListener) ExitSourceFile(ctx *SourceFileContext) {}

// EnterContractClause is called when production contractClause is entered.
func (s *BaseEvelynListener) EnterContractClause(ctx *ContractClauseContext) {}

// ExitContractClause is called when production contractClause is exited.
func (s *BaseEvelynListener) ExitContractClause(ctx *ContractClauseContext) {}

// EnterTopLevelDecl is called when production topLevelDecl is entered.
func (s *BaseEvelynListener) EnterTopLevelDecl(ctx *TopLevelDeclContext) {}

// ExitTopLevelDecl is called when production topLevelDecl is exited.
func (s *BaseEvelynListener) ExitTopLevelDecl(ctx *TopLevelDeclContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseEvelynListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseEvelynListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterVarDecl is called when production varDecl is entered.
func (s *BaseEvelynListener) EnterVarDecl(ctx *VarDeclContext) {}

// ExitVarDecl is called when production varDecl is exited.
func (s *BaseEvelynListener) ExitVarDecl(ctx *VarDeclContext) {}

// EnterEos is called when production eos is entered.
func (s *BaseEvelynListener) EnterEos(ctx *EosContext) {}

// ExitEos is called when production eos is exited.
func (s *BaseEvelynListener) ExitEos(ctx *EosContext) {}
