// Code generated from Evelyn.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Evelyn
import "github.com/antlr/antlr4/runtime/Go/antlr"

// EvelynListener is a complete listener for a parse tree produced by EvelynParser.
type EvelynListener interface {
	antlr.ParseTreeListener

	// EnterSourceFile is called when entering the sourceFile production.
	EnterSourceFile(c *SourceFileContext)

	// EnterContractClause is called when entering the contractClause production.
	EnterContractClause(c *ContractClauseContext)

	// EnterTopLevelDecl is called when entering the topLevelDecl production.
	EnterTopLevelDecl(c *TopLevelDeclContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterVarDecl is called when entering the varDecl production.
	EnterVarDecl(c *VarDeclContext)

	// EnterEos is called when entering the eos production.
	EnterEos(c *EosContext)

	// ExitSourceFile is called when exiting the sourceFile production.
	ExitSourceFile(c *SourceFileContext)

	// ExitContractClause is called when exiting the contractClause production.
	ExitContractClause(c *ContractClauseContext)

	// ExitTopLevelDecl is called when exiting the topLevelDecl production.
	ExitTopLevelDecl(c *TopLevelDeclContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitVarDecl is called when exiting the varDecl production.
	ExitVarDecl(c *VarDeclContext)

	// ExitEos is called when exiting the eos production.
	ExitEos(c *EosContext)
}
