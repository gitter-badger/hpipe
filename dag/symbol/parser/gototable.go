
/*
*/
package parser

const numNTSymbols = 9
type(
	gotoTable [numStates]gotoRow
	gotoRow	[numNTSymbols] int
)

var gotoTab = gotoTable{
	gotoRow{ // S0
		
		-1, // S'
		2, // StmtList
		3, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S1
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S2
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S3
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S4
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S5
		
		-1, // S'
		-1, // StmtList
		7, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S6
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		9, // Expr
		10, // Term
		11, // Factor
		17, // Id
		15, // Date
		16, // Env
		

	},
	gotoRow{ // S7
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S8
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S9
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S10
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S11
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S12
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		27, // Expr
		28, // Term
		29, // Factor
		35, // Id
		33, // Date
		34, // Env
		

	},
	gotoRow{ // S13
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S14
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S15
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S16
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S17
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S18
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S19
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S20
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S21
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S22
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		41, // Term
		11, // Factor
		17, // Id
		15, // Date
		16, // Env
		

	},
	gotoRow{ // S23
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		42, // Term
		11, // Factor
		17, // Id
		15, // Date
		16, // Env
		

	},
	gotoRow{ // S24
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		43, // Factor
		17, // Id
		15, // Date
		16, // Env
		

	},
	gotoRow{ // S25
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		44, // Factor
		17, // Id
		15, // Date
		16, // Env
		

	},
	gotoRow{ // S26
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S27
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S28
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S29
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S30
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		51, // Expr
		28, // Term
		29, // Factor
		35, // Id
		33, // Date
		34, // Env
		

	},
	gotoRow{ // S31
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S32
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S33
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S34
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S35
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S36
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S37
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S38
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S39
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S40
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S41
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S42
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S43
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S44
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S45
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S46
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		56, // Term
		29, // Factor
		35, // Id
		33, // Date
		34, // Env
		

	},
	gotoRow{ // S47
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		57, // Term
		29, // Factor
		35, // Id
		33, // Date
		34, // Env
		

	},
	gotoRow{ // S48
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S49
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		58, // Factor
		35, // Id
		33, // Date
		34, // Env
		

	},
	gotoRow{ // S50
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		59, // Factor
		35, // Id
		33, // Date
		34, // Env
		

	},
	gotoRow{ // S51
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S52
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S53
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S54
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S55
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S56
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S57
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S58
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S59
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S60
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S61
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	gotoRow{ // S62
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Id
		-1, // Date
		-1, // Env
		

	},
	
}
