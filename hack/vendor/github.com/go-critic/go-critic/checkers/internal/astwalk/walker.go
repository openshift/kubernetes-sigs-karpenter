/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package astwalk

import (
	"go/types"

	"github.com/go-critic/go-critic/linter"
)

// WalkerForFuncDecl returns file walker implementation for FuncDeclVisitor.
func WalkerForFuncDecl(v FuncDeclVisitor) linter.FileWalker {
	return &funcDeclWalker{visitor: v}
}

// WalkerForExpr returns file walker implementation for ExprVisitor.
func WalkerForExpr(v ExprVisitor) linter.FileWalker {
	return &exprWalker{visitor: v}
}

// WalkerForLocalExpr returns file walker implementation for LocalExprVisitor.
func WalkerForLocalExpr(v LocalExprVisitor) linter.FileWalker {
	return &localExprWalker{visitor: v}
}

// WalkerForStmtList returns file walker implementation for StmtListVisitor.
func WalkerForStmtList(v StmtListVisitor) linter.FileWalker {
	return &stmtListWalker{visitor: v}
}

// WalkerForStmt returns file walker implementation for StmtVisitor.
func WalkerForStmt(v StmtVisitor) linter.FileWalker {
	return &stmtWalker{visitor: v}
}

// WalkerForTypeExpr returns file walker implementation for TypeExprVisitor.
func WalkerForTypeExpr(v TypeExprVisitor, info *types.Info) linter.FileWalker {
	return &typeExprWalker{visitor: v, info: info}
}

// WalkerForLocalComment returns file walker implementation for LocalCommentVisitor.
func WalkerForLocalComment(v LocalCommentVisitor) linter.FileWalker {
	return &localCommentWalker{visitor: v}
}

// WalkerForComment returns file walker implementation for CommentVisitor.
func WalkerForComment(v CommentVisitor) linter.FileWalker {
	return &commentWalker{visitor: v}
}

// WalkerForDocComment returns file walker implementation for DocCommentVisitor.
func WalkerForDocComment(v DocCommentVisitor) linter.FileWalker {
	return &docCommentWalker{visitor: v}
}

// WalkerForLocalDef returns file walker implementation for LocalDefVisitor.
func WalkerForLocalDef(v LocalDefVisitor, info *types.Info) linter.FileWalker {
	return &localDefWalker{visitor: v, info: info}
}
