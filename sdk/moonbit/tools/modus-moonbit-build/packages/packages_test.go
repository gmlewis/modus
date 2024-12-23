package packages

import (
	_ "embed"
	"testing"

	"golang.org/x/tools/go/packages"
)

//go:embed testdata/simple-example.mbt
var simpleExample string

func TestPackage(t *testing.T) {
	t.Parallel()
	mode := NeedName | NeedImports | NeedDeps | NeedTypes | NeedSyntax | NeedTypesInfo
	cfg := &packages.Config{Mode: mode, Dir: dir}
	pkgs, err := Load(cfg, ".")
	if err != nil {
		t.Fatal(err)
	}
}

/*

GML: Calling packages.Load(cfg,'.'): cfg=
(*packages.Config)(0x14000326500)({
 Mode: (packages.LoadMode) (NeedName|NeedImports|NeedDeps|NeedTypes|NeedSyntax|NeedTypesInfo),
 Context: (context.Context) <nil>,
 Logf: (func(string, ...interface {})) <nil>,
 Dir: (string) (len=69) "/Users/glenn/src/github.com/hypermodeinc/modus/sdk/go/examples/simple",
 Env: ([]string) <nil>,
 BuildFlags: ([]string) <nil>,
 Fset: (*token.FileSet)(<nil>),
 ParseFile: (func(*token.FileSet, string, []uint8) (*ast.File, error)) <nil>,
 Tests: (bool) false,
 Overlay: (map[string][]uint8) <nil>,
 modFile: (string) "",
 modFlag: (string) ""
})

GML: getFunctionsNeedingWrappers: pkg.Syntax[0]=
(*ast.File)(0x14007fd88c0)({
 Doc: (*ast.CommentGroup)(<nil>),
 Package: (token.Pos) 4231927,
 Name: (*ast.Ident)(0x14007fa3a00)(main),
 Decls: ([]ast.Decl) (len=20 cap=32) {
  (*ast.GenDecl)(0x14007fe7640)({
   Doc: (*ast.CommentGroup)(<nil>),
   TokPos: (token.Pos) 4231941,
   Tok: (token.Token) import,
   Lparen: (token.Pos) 4231948,
   Specs: ([]ast.Spec) (len=6 cap=8) {
    (*ast.ImportSpec)(0x14007fdd890)({
     Doc: (*ast.CommentGroup)(<nil>),
     Name: (*ast.Ident)(<nil>),
     Path: (*ast.BasicLit)(0x14007fa3a20)({
      ValuePos: (token.Pos) 4231951,
      Kind: (token.Token) STRING,
      Value: (string) (len=8) "\"errors\""
     }),
     Comment: (*ast.CommentGroup)(<nil>),
     EndPos: (token.Pos) 0
    }),
    (*ast.ImportSpec)(0x14007fdd8c0)({
     Doc: (*ast.CommentGroup)(<nil>),
     Name: (*ast.Ident)(<nil>),
     Path: (*ast.BasicLit)(0x14007fa3a40)({
      ValuePos: (token.Pos) 4231961,
      Kind: (token.Token) STRING,
      Value: (string) (len=5) "\"fmt\""
     }),
     Comment: (*ast.CommentGroup)(<nil>),
     EndPos: (token.Pos) 0
    }),
    (*ast.ImportSpec)(0x14007fdd8f0)({
     Doc: (*ast.CommentGroup)(<nil>),
     Name: (*ast.Ident)(<nil>),
     Path: (*ast.BasicLit)(0x14007fa3a80)({
      ValuePos: (token.Pos) 4231968,
      Kind: (token.Token) STRING,
      Value: (string) (len=11) "\"math/rand\""
     }),
     Comment: (*ast.CommentGroup)(<nil>),
     EndPos: (token.Pos) 0
    }),
    (*ast.ImportSpec)(0x14007fdd920)({
     Doc: (*ast.CommentGroup)(<nil>),
     Name: (*ast.Ident)(<nil>),
     Path: (*ast.BasicLit)(0x14007fa3ac0)({
      ValuePos: (token.Pos) 4231981,
      Kind: (token.Token) STRING,
      Value: (string) (len=4) "\"os\""
     }),
     Comment: (*ast.CommentGroup)(<nil>),
     EndPos: (token.Pos) 0
    }),
    (*ast.ImportSpec)(0x14007fdd950)({
     Doc: (*ast.CommentGroup)(<nil>),
     Name: (*ast.Ident)(<nil>),
     Path: (*ast.BasicLit)(0x14007fa3ae0)({
      ValuePos: (token.Pos) 4231987,
      Kind: (token.Token) STRING,
      Value: (string) (len=6) "\"time\""
     }),
     Comment: (*ast.CommentGroup)(<nil>),
     EndPos: (token.Pos) 0
    }),
    (*ast.ImportSpec)(0x14007fdd980)({
     Doc: (*ast.CommentGroup)(<nil>),
     Name: (*ast.Ident)(<nil>),
     Path: (*ast.BasicLit)(0x14007fa3b00)({
      ValuePos: (token.Pos) 4231996,
      Kind: (token.Token) STRING,
      Value: (string) (len=50) "\"github.com/hypermodeinc/modus/sdk/go/pkg/console\""
     }),
     Comment: (*ast.CommentGroup)(<nil>),
     EndPos: (token.Pos) 0
    })
   },
   Rparen: (token.Pos) 4232047
  }),
  (*ast.FuncDecl)(0x14007fdda10)({
   Doc: (*ast.CommentGroup)(0x14007f9b410)({
    List: ([]*ast.Comment) (len=1 cap=1) {
     (*ast.Comment)(0x14007f9b3f8)({
      Slash: (token.Pos) 4232050,
      Text: (string) (len=18) "// Logs a message."
     })
    }
   }),
   Recv: (*ast.FieldList)(<nil>),
   Name: (*ast.Ident)(0x14007fa3b20)(LogMessage),
   Type: (*ast.FuncType)(0x14007fa3c00)({
    Func: (token.Pos) 4232069,
    TypeParams: (*ast.FieldList)(<nil>),
    Params: (*ast.FieldList)(0x14007fdd9b0)({
     Opening: (token.Pos) 4232084,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x14007fe7680)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) (len=1 cap=1) {
        (*ast.Ident)(0x14007fa3b40)(message)
       },
       Type: (*ast.Ident)(0x14007fa3b60)(string),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      })
     },
     Closing: (token.Pos) 4232099
    }),
    Results: (*ast.FieldList)(<nil>)
   }),
   Body: (*ast.BlockStmt)(0x14007fdd9e0)({
    Lbrace: (token.Pos) 4232101,
    List: ([]ast.Stmt) (len=1 cap=1) {
     (*ast.ExprStmt)(0x14007fc7730)({
      X: (*ast.CallExpr)(0x14007fe76c0)({
       Fun: (*ast.SelectorExpr)(0x14007f9b440)({
        X: (*ast.Ident)(0x14007fa3b80)(console),
        Sel: (*ast.Ident)(0x14007fa3ba0)(Log)
       }),
       Lparen: (token.Pos) 4232115,
       Args: ([]ast.Expr) (len=1 cap=1) {
        (*ast.Ident)(0x14007fa3bc0)(message)
       },
       Ellipsis: (token.Pos) 0,
       Rparen: (token.Pos) 4232123
      })
     })
    },
    Rbrace: (token.Pos) 4232125
   })
  }),
  (*ast.FuncDecl)(0x14007fddb30)({
   Doc: (*ast.CommentGroup)(0x14007f9b470)({
    List: ([]*ast.Comment) (len=1 cap=1) {
     (*ast.Comment)(0x14007f9b458)({
      Slash: (token.Pos) 4232128,
      Text: (string) (len=53) "// Adds two integers together and returns the result."
     })
    }
   }),
   Recv: (*ast.FieldList)(<nil>),
   Name: (*ast.Ident)(0x14007fa3c40)(Add),
   Type: (*ast.FuncType)(0x14007fa3d40)({
    Func: (token.Pos) 4232182,
    TypeParams: (*ast.FieldList)(<nil>),
    Params: (*ast.FieldList)(0x14007fdda70)({
     Opening: (token.Pos) 4232190,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x14007fe7700)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) (len=2 cap=2) {
        (*ast.Ident)(0x14007fa3c60)(x),
        (*ast.Ident)(0x14007fa3c80)(y)
       },
       Type: (*ast.Ident)(0x14007fa3ca0)(int),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      })
     },
     Closing: (token.Pos) 4232199
    }),
    Results: (*ast.FieldList)(0x14007fddaa0)({
     Opening: (token.Pos) 0,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x14007fe7740)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) <nil>,
       Type: (*ast.Ident)(0x14007fa3cc0)(int),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      })
     },
     Closing: (token.Pos) 0
    })
   }),
   Body: (*ast.BlockStmt)(0x14007fddb00)({
    Lbrace: (token.Pos) 4232205,
    List: ([]ast.Stmt) (len=1 cap=1) {
     (*ast.ReturnStmt)(0x14007fa3d20)({
      Return: (token.Pos) 4232208,
      Results: ([]ast.Expr) (len=1 cap=1) {
       (*ast.BinaryExpr)(0x14007fddad0)({
        X: (*ast.Ident)(0x14007fa3ce0)(x),
        OpPos: (token.Pos) 4232217,
        Op: (token.Token) +,
        Y: (*ast.Ident)(0x14007fa3d00)(y)
       })
      }
     })
    },
    Rbrace: (token.Pos) 4232221
   })
  }),
  (*ast.FuncDecl)(0x14007fddd10)({
   Doc: (*ast.CommentGroup)(0x14007f9b4d0)({
    List: ([]*ast.Comment) (len=2 cap=2) {
     (*ast.Comment)(0x14007f9b4a0)({
      Slash: (token.Pos) 4232224,
      Text: (string) (len=55) "// Adds three integers together and returns the result."
     }),
     (*ast.Comment)(0x14007f9b4b8)({
      Slash: (token.Pos) 4232280,
      Text: (string) (len=33) "// The third integer is optional."
     })
    }
   }),
   Recv: (*ast.FieldList)(<nil>),
   Name: (*ast.Ident)(0x14007fa3d60)(Add3),
   Type: (*ast.FuncType)(0x14007ffe000)({
    Func: (token.Pos) 4232314,
    TypeParams: (*ast.FieldList)(<nil>),
    Params: (*ast.FieldList)(0x14007fddb90)({
     Opening: (token.Pos) 4232323,
     List: ([]*ast.Field) (len=2 cap=2) {
      (*ast.Field)(0x14007fe77c0)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) (len=2 cap=2) {
        (*ast.Ident)(0x14007fa3d80)(a),
        (*ast.Ident)(0x14007fa3da0)(b)
       },
       Type: (*ast.Ident)(0x14007fa3dc0)(int),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      }),
      (*ast.Field)(0x14007fe7800)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) (len=1 cap=1) {
        (*ast.Ident)(0x14007fa3de0)(c)
       },
       Type: (*ast.StarExpr)(0x14007f9b500)({
        Star: (token.Pos) 4232336,
        X: (*ast.Ident)(0x14007fa3e00)(int)
       }),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      })
     },
     Closing: (token.Pos) 4232340
    }),
    Results: (*ast.FieldList)(0x14007fddbc0)({
     Opening: (token.Pos) 0,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x14007fe7840)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) <nil>,
       Type: (*ast.Ident)(0x14007fa3e20)(int),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      })
     },
     Closing: (token.Pos) 0
    })
   }),
   Body: (*ast.BlockStmt)(0x14007fddce0)({
    Lbrace: (token.Pos) 4232346,
    List: ([]ast.Stmt) (len=2 cap=2) {
     (*ast.IfStmt)(0x14007fe7880)({
      If: (token.Pos) 4232349,
      Init: (ast.Stmt) <nil>,
      Cond: (*ast.BinaryExpr)(0x14007fddbf0)({
       X: (*ast.Ident)(0x14007fa3e40)(c),
       OpPos: (token.Pos) 4232354,
       Op: (token.Token) !=,
       Y: (*ast.Ident)(0x14007fa3e60)(nil)
      }),
      Body: (*ast.BlockStmt)(0x14007fddc80)({
       Lbrace: (token.Pos) 4232361,
       List: ([]ast.Stmt) (len=1 cap=1) {
        (*ast.ReturnStmt)(0x14007fa3ee0)({
         Return: (token.Pos) 4232365,
         Results: ([]ast.Expr) (len=1 cap=1) {
          (*ast.BinaryExpr)(0x14007fddc50)({
           X: (*ast.BinaryExpr)(0x14007fddc20)({
            X: (*ast.Ident)(0x14007fa3e80)(a),
            OpPos: (token.Pos) 4232374,
            Op: (token.Token) +,
            Y: (*ast.Ident)(0x14007fa3ea0)(b)
           }),
           OpPos: (token.Pos) 4232378,
           Op: (token.Token) +,
           Y: (*ast.StarExpr)(0x14007f9b518)({
            Star: (token.Pos) 4232380,
            X: (*ast.Ident)(0x14007fa3ec0)(c)
           })
          })
         }
        })
       },
       Rbrace: (token.Pos) 4232384
      }),
      Else: (ast.Stmt) <nil>
     }),
     (*ast.ReturnStmt)(0x14007fa3f40)({
      Return: (token.Pos) 4232387,
      Results: ([]ast.Expr) (len=1 cap=1) {
       (*ast.BinaryExpr)(0x14007fddcb0)({
        X: (*ast.Ident)(0x14007fa3f00)(a),
        OpPos: (token.Pos) 4232396,
        Op: (token.Token) +,
        Y: (*ast.Ident)(0x14007fa3f20)(b)
       })
      }
     })
    },
    Rbrace: (token.Pos) 4232400
   })
  }),
  (*ast.FuncDecl)(0x14007fdde00)({
   Doc: (*ast.CommentGroup)(0x14007f9b548)({
    List: ([]*ast.Comment) (len=1 cap=1) {
     (*ast.Comment)(0x14007f9b530)({
      Slash: (token.Pos) 4232403,
      Text: (string) (len=63) "// Adds any number of integers together and returns the result."
     })
    }
   }),
   Recv: (*ast.FieldList)(<nil>),
   Name: (*ast.Ident)(0x14007ffe020)(AddN),
   Type: (*ast.FuncType)(0x14007ffe220)({
    Func: (token.Pos) 4232467,
    TypeParams: (*ast.FieldList)(<nil>),
    Params: (*ast.FieldList)(0x14007fddd40)({
     Opening: (token.Pos) 4232476,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x14007fe7900)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) (len=1 cap=1) {
        (*ast.Ident)(0x14007ffe040)(args)
       },
       Type: (*ast.Ellipsis)(0x14007f9b560)({
        Ellipsis: (token.Pos) 4232482,
        Elt: (*ast.Ident)(0x14007ffe060)(int)
       }),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      })
     },
     Closing: (token.Pos) 4232488
    }),
    Results: (*ast.FieldList)(0x14007fddd70)({
     Opening: (token.Pos) 0,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x14007fe7940)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) <nil>,
       Type: (*ast.Ident)(0x14007ffe080)(int),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      })
     },
     Closing: (token.Pos) 0
    })
   }),
   Body: (*ast.BlockStmt)(0x14007fdddd0)({
    Lbrace: (token.Pos) 4232494,
    List: ([]ast.Stmt) (len=3 cap=4) {
     (*ast.AssignStmt)(0x14007fe7980)({
      Lhs: ([]ast.Expr) (len=1 cap=1) {
       (*ast.Ident)(0x14007ffe0a0)(sum)
      },
      TokPos: (token.Pos) 4232501,
      Tok: (token.Token) :=,
      Rhs: ([]ast.Expr) (len=1 cap=1) {
       (*ast.BasicLit)(0x14007ffe0c0)({
        ValuePos: (token.Pos) 4232504,
        Kind: (token.Token) INT,
        Value: (string) (len=1) "0"
       })
      }
     }),
     (*ast.RangeStmt)(0x14007fe5aa0)({
      For: (token.Pos) 4232507,
      Key: (*ast.Ident)(0x14007ffe0e0)(_),
      Value: (*ast.Ident)(0x14007ffe100)(arg),
      TokPos: (token.Pos) 4232518,
      Tok: (token.Token) :=,
      Range: (token.Pos) 4232521,
      X: (*ast.Ident)(0x14007ffe140)(args),
      Body: (*ast.BlockStmt)(0x14007fddda0)({
       Lbrace: (token.Pos) 4232532,
       List: ([]ast.Stmt) (len=1 cap=1) {
        (*ast.AssignStmt)(0x14007fe7a00)({
         Lhs: ([]ast.Expr) (len=1 cap=1) {
          (*ast.Ident)(0x14007ffe180)(sum)
         },
         TokPos: (token.Pos) 4232540,
         Tok: (token.Token) +=,
         Rhs: ([]ast.Expr) (len=1 cap=1) {
          (*ast.Ident)(0x14007ffe1a0)(arg)
         }
        })
       },
       Rbrace: (token.Pos) 4232548
      })
     }),
     (*ast.ReturnStmt)(0x14007ffe200)({
      Return: (token.Pos) 4232551,
      Results: ([]ast.Expr) (len=1 cap=1) {
       (*ast.Ident)(0x14007ffe1e0)(sum)
      }
     })
    },
    Rbrace: (token.Pos) 4232562
   })
  }),
  (*ast.GenDecl)(0x14007fe7a80)({
   Doc: (*ast.CommentGroup)(0x14007f9b5a8)({
    List: ([]*ast.Comment) (len=1 cap=1) {
     (*ast.Comment)(0x14007f9b590)({
      Slash: (token.Pos) 4232565,
      Text: (string) (len=55) "// this indirection is so we can mock time.Now in tests"
     })
    }
   }),
   TokPos: (token.Pos) 4232621,
   Tok: (token.Token) var,
   Lparen: (token.Pos) 0,
   Specs: ([]ast.Spec) (len=1 cap=1) {
    (*ast.ValueSpec)(0x14007fcb130)({
     Doc: (*ast.CommentGroup)(<nil>),
     Names: ([]*ast.Ident) (len=1 cap=1) {
      (*ast.Ident)(0x14007ffe240)(nowFunc)
     },
     Type: (ast.Expr) <nil>,
     Values: ([]ast.Expr) (len=1 cap=1) {
      (*ast.SelectorExpr)(0x14007f9b5c0)({
       X: (*ast.Ident)(0x14007ffe260)(time),
       Sel: (*ast.Ident)(0x14007ffe280)(Now)
      })
     },
     Comment: (*ast.CommentGroup)(<nil>)
    })
   },
   Rparen: (token.Pos) 0
  }),
  (*ast.FuncDecl)(0x14007fddec0)({
   Doc: (*ast.CommentGroup)(0x14007f9b5f0)({
    List: ([]*ast.Comment) (len=1 cap=1) {
     (*ast.Comment)(0x14007f9b5d8)({
      Slash: (token.Pos) 4232645,
      Text: (string) (len=28) "// Returns the current time."
     })
    }
   }),
   Recv: (*ast.FieldList)(<nil>),
   Name: (*ast.Ident)(0x14007ffe2a0)(GetCurrentTime),
   Type: (*ast.FuncType)(0x14007ffe340)({
    Func: (token.Pos) 4232674,
    TypeParams: (*ast.FieldList)(<nil>),
    Params: (*ast.FieldList)(0x14007fdde30)({
     Opening: (token.Pos) 4232693,
     List: ([]*ast.Field) <nil>,
     Closing: (token.Pos) 4232694
    }),
    Results: (*ast.FieldList)(0x14007fdde60)({
     Opening: (token.Pos) 0,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x14007fe7ac0)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) <nil>,
       Type: (*ast.SelectorExpr)(0x14007f9b608)({
        X: (*ast.Ident)(0x14007ffe2c0)(time),
        Sel: (*ast.Ident)(0x14007ffe2e0)(Time)
       }),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      })
     },
     Closing: (token.Pos) 0
    })
   }),
   Body: (*ast.BlockStmt)(0x14007fdde90)({
    Lbrace: (token.Pos) 4232706,
    List: ([]ast.Stmt) (len=1 cap=1) {
     (*ast.ReturnStmt)(0x14007ffe320)({
      Return: (token.Pos) 4232709,
      Results: ([]ast.Expr) (len=1 cap=1) {
       (*ast.CallExpr)(0x14007fe7b00)({
        Fun: (*ast.Ident)(0x14007ffe300)(nowFunc),
        Lparen: (token.Pos) 4232723,
        Args: ([]ast.Expr) <nil>,
        Ellipsis: (token.Pos) 0,
        Rparen: (token.Pos) 4232724
       })
      }
     })
    },
    Rbrace: (token.Pos) 4232726
   })
  }),
  (*ast.FuncDecl)(0x14008478000)({
   Doc: (*ast.CommentGroup)(0x14007f9b638)({
    List: ([]*ast.Comment) (len=1 cap=1) {
     (*ast.Comment)(0x14007f9b620)({
      Slash: (token.Pos) 4232729,
      Text: (string) (len=50) "// Returns the current time formatted as a string."
     })
    }
   }),
   Recv: (*ast.FieldList)(<nil>),
   Name: (*ast.Ident)(0x14007ffe360)(GetCurrentTimeFormatted),
   Type: (*ast.FuncType)(0x14007ffe440)({
    Func: (token.Pos) 4232780,
    TypeParams: (*ast.FieldList)(<nil>),
    Params: (*ast.FieldList)(0x14007fddef0)({
     Opening: (token.Pos) 4232808,
     List: ([]*ast.Field) <nil>,
     Closing: (token.Pos) 4232809
    }),
    Results: (*ast.FieldList)(0x14007fddf20)({
     Opening: (token.Pos) 0,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x14007fe7b40)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) <nil>,
       Type: (*ast.Ident)(0x14007ffe380)(string),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      })
     },
     Closing: (token.Pos) 0
    })
   }),
   Body: (*ast.BlockStmt)(0x14007fddf50)({
    Lbrace: (token.Pos) 4232818,
    List: ([]ast.Stmt) (len=1 cap=1) {
     (*ast.ReturnStmt)(0x14007ffe420)({
      Return: (token.Pos) 4232821,
      Results: ([]ast.Expr) (len=1 cap=1) {
       (*ast.CallExpr)(0x14007fe7bc0)({
        Fun: (*ast.SelectorExpr)(0x14007f9b650)({
         X: (*ast.CallExpr)(0x14007fe7b80)({
          Fun: (*ast.Ident)(0x14007ffe3a0)(nowFunc),
          Lparen: (token.Pos) 4232835,
          Args: ([]ast.Expr) <nil>,
          Ellipsis: (token.Pos) 0,
          Rparen: (token.Pos) 4232836
         }),
         Sel: (*ast.Ident)(0x14007ffe3c0)(Format)
        }),
        Lparen: (token.Pos) 4232844,
        Args: ([]ast.Expr) (len=1 cap=1) {
         (*ast.SelectorExpr)(0x14007f9b668)({
          X: (*ast.Ident)(0x14007ffe3e0)(time),
          Sel: (*ast.Ident)(0x14007ffe400)(DateTime)
         })
        },
        Ellipsis: (token.Pos) 0,
        Rparen: (token.Pos) 4232858
       })
      }
     })
    },
    Rbrace: (token.Pos) 4232860
   })
  }),
  (*ast.FuncDecl)(0x14008478150)({
   Doc: (*ast.CommentGroup)(0x14007f9b698)({
    List: ([]*ast.Comment) (len=1 cap=1) {
     (*ast.Comment)(0x14007f9b680)({
      Slash: (token.Pos) 4232863,
      Text: (string) (len=75) "// Combines the first and last name of a person, and returns the full name."
     })
    }
   }),
   Recv: (*ast.FieldList)(<nil>),
   Name: (*ast.Ident)(0x14007ffe460)(GetFullName),
   Type: (*ast.FuncType)(0x14007ffe580)({
    Func: (token.Pos) 4232939,
    TypeParams: (*ast.FieldList)(<nil>),
    Params: (*ast.FieldList)(0x14008478060)({
     Opening: (token.Pos) 4232955,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x14007fe7c00)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) (len=2 cap=2) {
        (*ast.Ident)(0x14007ffe480)(firstName),
        (*ast.Ident)(0x14007ffe4a0)(lastName)
       },
       Type: (*ast.Ident)(0x14007ffe4c0)(string),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      })
     },
     Closing: (token.Pos) 4232982
    }),
    Results: (*ast.FieldList)(0x14008478090)({
     Opening: (token.Pos) 0,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x14007fe7c40)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) <nil>,
       Type: (*ast.Ident)(0x14007ffe4e0)(string),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      })
     },
     Closing: (token.Pos) 0
    })
   }),
   Body: (*ast.BlockStmt)(0x14008478120)({
    Lbrace: (token.Pos) 4232991,
    List: ([]ast.Stmt) (len=1 cap=1) {
     (*ast.ReturnStmt)(0x14007ffe560)({
      Return: (token.Pos) 4232994,
      Results: ([]ast.Expr) (len=1 cap=1) {
       (*ast.BinaryExpr)(0x140084780f0)({
        X: (*ast.BinaryExpr)(0x140084780c0)({
         X: (*ast.Ident)(0x14007ffe500)(firstName),
         OpPos: (token.Pos) 4233011,
         Op: (token.Token) +,
         Y: (*ast.BasicLit)(0x14007ffe520)({
          ValuePos: (token.Pos) 4233013,
          Kind: (token.Token) STRING,
          Value: (string) (len=3) "\" \""
         })
        }),
        OpPos: (token.Pos) 4233017,
        Op: (token.Token) +,
        Y: (*ast.Ident)(0x14007ffe540)(lastName)
       })
      }
     })
    },
    Rbrace: (token.Pos) 4233028
   })
  }),
  (*ast.FuncDecl)(0x14008478300)({
   Doc: (*ast.CommentGroup)(0x14007f9b6f8)({
    List: ([]*ast.Comment) (len=2 cap=2) {
     (*ast.Comment)(0x14007f9b6c8)({
      Slash: (token.Pos) 4233031,
      Text: (string) (len=34) "// Says hello to a person by name."
     }),
     (*ast.Comment)(0x14007f9b6e0)({
      Slash: (token.Pos) 4233066,
      Text: (string) (len=65) "// If the name is not provided, it will say hello without a name."
     })
    }
   }),
   Recv: (*ast.FieldList)(<nil>),
   Name: (*ast.Ident)(0x14007ffe5a0)(SayHello),
   Type: (*ast.FuncType)(0x14007ffe720)({
    Func: (token.Pos) 4233132,
    TypeParams: (*ast.FieldList)(<nil>),
    Params: (*ast.FieldList)(0x14008478180)({
     Opening: (token.Pos) 4233145,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x14007fe7c80)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) (len=1 cap=1) {
        (*ast.Ident)(0x14007ffe5c0)(name)
       },
       Type: (*ast.StarExpr)(0x14007f9b710)({
        Star: (token.Pos) 4233151,
        X: (*ast.Ident)(0x14007ffe5e0)(string)
       }),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      })
     },
     Closing: (token.Pos) 4233158
    }),
    Results: (*ast.FieldList)(0x140084781b0)({
     Opening: (token.Pos) 0,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x14007fe7cc0)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) <nil>,
       Type: (*ast.Ident)(0x14007ffe600)(string),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      })
     },
     Closing: (token.Pos) 0
    })
   }),
   Body: (*ast.BlockStmt)(0x140084782d0)({
    Lbrace: (token.Pos) 4233167,
    List: ([]ast.Stmt) (len=1 cap=1) {
     (*ast.IfStmt)(0x14007fe7d00)({
      If: (token.Pos) 4233170,
      Init: (ast.Stmt) <nil>,
      Cond: (*ast.BinaryExpr)(0x140084781e0)({
       X: (*ast.Ident)(0x14007ffe620)(name),
       OpPos: (token.Pos) 4233178,
       Op: (token.Token) ==,
       Y: (*ast.Ident)(0x14007ffe640)(nil)
      }),
      Body: (*ast.BlockStmt)(0x14008478210)({
       Lbrace: (token.Pos) 4233185,
       List: ([]ast.Stmt) (len=1 cap=1) {
        (*ast.ReturnStmt)(0x14007ffe680)({
         Return: (token.Pos) 4233189,
         Results: ([]ast.Expr) (len=1 cap=1) {
          (*ast.BasicLit)(0x14007ffe660)({
           ValuePos: (token.Pos) 4233196,
           Kind: (token.Token) STRING,
           Value: (string) (len=8) "\"Hello!\""
          })
         }
        })
       },
       Rbrace: (token.Pos) 4233206
      }),
      Else: (*ast.BlockStmt)(0x140084782a0)({
       Lbrace: (token.Pos) 4233213,
       List: ([]ast.Stmt) (len=1 cap=1) {
        (*ast.ReturnStmt)(0x14007ffe700)({
         Return: (token.Pos) 4233217,
         Results: ([]ast.Expr) (len=1 cap=1) {
          (*ast.BinaryExpr)(0x14008478270)({
           X: (*ast.BinaryExpr)(0x14008478240)({
            X: (*ast.BasicLit)(0x14007ffe6a0)({
             ValuePos: (token.Pos) 4233224,
             Kind: (token.Token) STRING,
             Value: (string) (len=9) "\"Hello, \""
            }),
            OpPos: (token.Pos) 4233234,
            Op: (token.Token) +,
            Y: (*ast.StarExpr)(0x14007f9b740)({
             Star: (token.Pos) 4233236,
             X: (*ast.Ident)(0x14007ffe6c0)(name)
            })
           }),
           OpPos: (token.Pos) 4233242,
           Op: (token.Token) +,
           Y: (*ast.BasicLit)(0x14007ffe6e0)({
            ValuePos: (token.Pos) 4233244,
            Kind: (token.Token) STRING,
            Value: (string) (len=3) "\"!\""
           })
          })
         }
        })
       },
       Rbrace: (token.Pos) 4233249
      })
     })
    },
    Rbrace: (token.Pos) 4233251
   })
  }),
  (*ast.GenDecl)(0x14007fe7e40)({
   Doc: (*ast.CommentGroup)(0x14007f9b770)({
    List: ([]*ast.Comment) (len=1 cap=1) {
     (*ast.Comment)(0x14007f9b758)({
      Slash: (token.Pos) 4233254,
      Text: (string) (len=41) "// A simple object representing a person."
     })
    }
   }),
   TokPos: (token.Pos) 4233296,
   Tok: (token.Token) type,
   Lparen: (token.Pos) 0,
   Specs: ([]ast.Spec) (len=1 cap=1) {
    (*ast.TypeSpec)(0x14007fe7d40)({
     Doc: (*ast.CommentGroup)(<nil>),
     Name: (*ast.Ident)(0x14007ffe740)(Person),
     TypeParams: (*ast.FieldList)(<nil>),
     Assign: (token.Pos) 0,
     Type: (*ast.StructType)(0x14007f9b818)({
      Struct: (token.Pos) 4233308,
      Fields: (*ast.FieldList)(0x14008478330)({
       Opening: (token.Pos) 4233315,
       List: ([]*ast.Field) (len=3 cap=4) {
        (*ast.Field)(0x14007fe7d80)({
         Doc: (*ast.CommentGroup)(0x14007f9b7a0)({
          List: ([]*ast.Comment) (len=1 cap=1) {
           (*ast.Comment)(0x14007f9b788)({
            Slash: (token.Pos) 4233319,
            Text: (string) (len=27) "// The person's first name."
           })
          }
         }),
         Names: ([]*ast.Ident) (len=1 cap=1) {
          (*ast.Ident)(0x14007ffe760)(FirstName)
         },
         Type: (*ast.Ident)(0x14007ffe780)(string),
         Tag: (*ast.BasicLit)(0x14007ffe7a0)({
          ValuePos: (token.Pos) 4233365,
          Kind: (token.Token) STRING,
          Value: (string) (len=18) "`json:\"firstName\"`"
         }),
         Comment: (*ast.CommentGroup)(<nil>)
        }),
        (*ast.Field)(0x14007fe7dc0)({
         Doc: (*ast.CommentGroup)(0x14007f9b7d0)({
          List: ([]*ast.Comment) (len=1 cap=1) {
           (*ast.Comment)(0x14007f9b7b8)({
            Slash: (token.Pos) 4233386,
            Text: (string) (len=26) "// The person's last name."
           })
          }
         }),
         Names: ([]*ast.Ident) (len=1 cap=1) {
          (*ast.Ident)(0x14007ffe7c0)(LastName)
         },
         Type: (*ast.Ident)(0x14007ffe7e0)(string),
         Tag: (*ast.BasicLit)(0x14007ffe800)({
          ValuePos: (token.Pos) 4233430,
          Kind: (token.Token) STRING,
          Value: (string) (len=17) "`json:\"lastName\"`"
         }),
         Comment: (*ast.CommentGroup)(<nil>)
        }),
        (*ast.Field)(0x14007fe7e00)({
         Doc: (*ast.CommentGroup)(0x14007f9b800)({
          List: ([]*ast.Comment) (len=1 cap=1) {
           (*ast.Comment)(0x14007f9b7e8)({
            Slash: (token.Pos) 4233450,
            Text: (string) (len=20) "// The person's age."
           })
          }
         }),
         Names: ([]*ast.Ident) (len=1 cap=1) {
          (*ast.Ident)(0x14007ffe820)(Age)
         },
         Type: (*ast.Ident)(0x14007ffe840)(int),
         Tag: (*ast.BasicLit)(0x14007ffe860)({
          ValuePos: (token.Pos) 4233480,
          Kind: (token.Token) STRING,
          Value: (string) (len=12) "`json:\"age\"`"
         }),
         Comment: (*ast.CommentGroup)(<nil>)
        })
       },
       Closing: (token.Pos) 4233493
      }),
      Incomplete: (bool) false
     }),
     Comment: (*ast.CommentGroup)(<nil>)
    })
   },
   Rparen: (token.Pos) 0
  }),
  (*ast.FuncDecl)(0x14008478480)({
   Doc: (*ast.CommentGroup)(0x14007f9b848)({
    List: ([]*ast.Comment) (len=1 cap=1) {
     (*ast.Comment)(0x14007f9b830)({
      Slash: (token.Pos) 4233496,
      Text: (string) (len=24) "// Gets a person object."
     })
    }
   }),
   Recv: (*ast.FieldList)(<nil>),
   Name: (*ast.Ident)(0x14007ffe8a0)(GetPerson),
   Type: (*ast.FuncType)(0x14007ffea00)({
    Func: (token.Pos) 4233521,
    TypeParams: (*ast.FieldList)(<nil>),
    Params: (*ast.FieldList)(0x14008478360)({
     Opening: (token.Pos) 4233535,
     List: ([]*ast.Field) <nil>,
     Closing: (token.Pos) 4233536
    }),
    Results: (*ast.FieldList)(0x14008478390)({
     Opening: (token.Pos) 0,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x14007fe7e80)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) <nil>,
       Type: (*ast.Ident)(0x14007ffe8c0)(Person),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      })
     },
     Closing: (token.Pos) 0
    })
   }),
   Body: (*ast.BlockStmt)(0x14008478450)({
    Lbrace: (token.Pos) 4233545,
    List: ([]ast.Stmt) (len=1 cap=1) {
     (*ast.ReturnStmt)(0x14007ffe9e0)({
      Return: (token.Pos) 4233548,
      Results: ([]ast.Expr) (len=1 cap=1) {
       (*ast.CompositeLit)(0x14007fe7f00)({
        Type: (*ast.Ident)(0x14007ffe8e0)(Person),
        Lbrace: (token.Pos) 4233561,
        Elts: ([]ast.Expr) (len=3 cap=4) {
         (*ast.KeyValueExpr)(0x140084783c0)({
          Key: (*ast.Ident)(0x14007ffe900)(FirstName),
          Colon: (token.Pos) 4233574,
          Value: (*ast.BasicLit)(0x14007ffe920)({
           ValuePos: (token.Pos) 4233576,
           Kind: (token.Token) STRING,
           Value: (string) (len=6) "\"John\""
          })
         }),
         (*ast.KeyValueExpr)(0x140084783f0)({
          Key: (*ast.Ident)(0x14007ffe940)(LastName),
          Colon: (token.Pos) 4233594,
          Value: (*ast.BasicLit)(0x14007ffe960)({
           ValuePos: (token.Pos) 4233597,
           Kind: (token.Token) STRING,
           Value: (string) (len=5) "\"Doe\""
          })
         }),
         (*ast.KeyValueExpr)(0x14008478420)({
          Key: (*ast.Ident)(0x14007ffe9a0)(Age),
          Colon: (token.Pos) 4233609,
          Value: (*ast.BasicLit)(0x14007ffe9c0)({
           ValuePos: (token.Pos) 4233617,
           Kind: (token.Token) INT,
           Value: (string) (len=2) "42"
          })
         })
        },
        Rbrace: (token.Pos) 4233622,
        Incomplete: (bool) false
       })
      }
     })
    },
    Rbrace: (token.Pos) 4233624
   })
  }),
  (*ast.FuncDecl)(0x14008478570)({
   Doc: (*ast.CommentGroup)(0x14007f9b878)({
    List: ([]*ast.Comment) (len=1 cap=1) {
     (*ast.Comment)(0x14007f9b860)({
      Slash: (token.Pos) 4233627,
      Text: (string) (len=53) "// Gets a random person object from a list of people."
     })
    }
   }),
   Recv: (*ast.FieldList)(<nil>),
   Name: (*ast.Ident)(0x14007ffea20)(GetRandomPerson),
   Type: (*ast.FuncType)(0x14007ffebc0)({
    Func: (token.Pos) 4233681,
    TypeParams: (*ast.FieldList)(<nil>),
    Params: (*ast.FieldList)(0x140084784b0)({
     Opening: (token.Pos) 4233701,
     List: ([]*ast.Field) <nil>,
     Closing: (token.Pos) 4233702
    }),
    Results: (*ast.FieldList)(0x140084784e0)({
     Opening: (token.Pos) 0,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x14007fe7f40)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) <nil>,
       Type: (*ast.Ident)(0x14007ffea40)(Person),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      })
     },
     Closing: (token.Pos) 0
    })
   }),
   Body: (*ast.BlockStmt)(0x14008478540)({
    Lbrace: (token.Pos) 4233711,
    List: ([]ast.Stmt) (len=3 cap=4) {
     (*ast.AssignStmt)(0x1400847a040)({
      Lhs: ([]ast.Expr) (len=1 cap=1) {
       (*ast.Ident)(0x14007ffea60)(people)
      },
      TokPos: (token.Pos) 4233721,
      Tok: (token.Token) :=,
      Rhs: ([]ast.Expr) (len=1 cap=1) {
       (*ast.CallExpr)(0x1400847a000)({
        Fun: (*ast.Ident)(0x14007ffea80)(GetPeople),
        Lparen: (token.Pos) 4233733,
        Args: ([]ast.Expr) <nil>,
        Ellipsis: (token.Pos) 0,
        Rparen: (token.Pos) 4233734
       })
      }
     }),
     (*ast.AssignStmt)(0x1400847a100)({
      Lhs: ([]ast.Expr) (len=1 cap=1) {
       (*ast.Ident)(0x14007ffeaa0)(i)
      },
      TokPos: (token.Pos) 4233739,
      Tok: (token.Token) :=,
      Rhs: ([]ast.Expr) (len=1 cap=1) {
       (*ast.CallExpr)(0x1400847a0c0)({
        Fun: (*ast.SelectorExpr)(0x14007f9b890)({
         X: (*ast.Ident)(0x14007ffeac0)(rand),
         Sel: (*ast.Ident)(0x14007ffeae0)(Intn)
        }),
        Lparen: (token.Pos) 4233751,
        Args: ([]ast.Expr) (len=1 cap=1) {
         (*ast.CallExpr)(0x1400847a080)({
          Fun: (*ast.Ident)(0x14007ffeb00)(len),
          Lparen: (token.Pos) 4233755,
          Args: ([]ast.Expr) (len=1 cap=1) {
           (*ast.Ident)(0x14007ffeb20)(people)
          },
          Ellipsis: (token.Pos) 0,
          Rparen: (token.Pos) 4233762
         })
        },
        Ellipsis: (token.Pos) 0,
        Rparen: (token.Pos) 4233763
       })
      }
     }),
     (*ast.ReturnStmt)(0x14007ffeba0)({
      Return: (token.Pos) 4233766,
      Results: ([]ast.Expr) (len=1 cap=1) {
       (*ast.IndexExpr)(0x14008478510)({
        X: (*ast.Ident)(0x14007ffeb60)(people),
        Lbrack: (token.Pos) 4233779,
        Index: (*ast.Ident)(0x14007ffeb80)(i),
        Rbrack: (token.Pos) 4233781
       })
      }
     })
    },
    Rbrace: (token.Pos) 4233783
   })
  }),
  (*ast.FuncDecl)(0x14008478840)({
   Doc: (*ast.CommentGroup)(0x14007f9b8c0)({
    List: ([]*ast.Comment) (len=1 cap=1) {
     (*ast.Comment)(0x14007f9b8a8)({
      Slash: (token.Pos) 4233786,
      Text: (string) (len=25) "// Gets a list of people."
     })
    }
   }),
   Recv: (*ast.FieldList)(<nil>),
   Name: (*ast.Ident)(0x14007ffebe0)(GetPeople),
   Type: (*ast.FuncType)(0x14007ffef20)({
    Func: (token.Pos) 4233812,
    TypeParams: (*ast.FieldList)(<nil>),
    Params: (*ast.FieldList)(0x140084785a0)({
     Opening: (token.Pos) 4233826,
     List: ([]*ast.Field) <nil>,
     Closing: (token.Pos) 4233827
    }),
    Results: (*ast.FieldList)(0x14008478600)({
     Opening: (token.Pos) 0,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x1400847a180)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) <nil>,
       Type: (*ast.ArrayType)(0x140084785d0)({
        Lbrack: (token.Pos) 4233829,
        Len: (ast.Expr) <nil>,
        Elt: (*ast.Ident)(0x14007ffec00)(Person)
       }),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      })
     },
     Closing: (token.Pos) 0
    })
   }),
   Body: (*ast.BlockStmt)(0x14008478810)({
    Lbrace: (token.Pos) 4233838,
    List: ([]ast.Stmt) (len=1 cap=1) {
     (*ast.ReturnStmt)(0x14007ffef00)({
      Return: (token.Pos) 4233841,
      Results: ([]ast.Expr) (len=1 cap=1) {
       (*ast.CompositeLit)(0x1400847a380)({
        Type: (*ast.ArrayType)(0x14008478630)({
         Lbrack: (token.Pos) 4233848,
         Len: (ast.Expr) <nil>,
         Elt: (*ast.Ident)(0x14007ffec20)(Person)
        }),
        Lbrace: (token.Pos) 4233856,
        Elts: ([]ast.Expr) (len=3 cap=4) {
         (*ast.CompositeLit)(0x1400847a200)({
          Type: (ast.Expr) <nil>,
          Lbrace: (token.Pos) 4233860,
          Elts: ([]ast.Expr) (len=3 cap=4) {
           (*ast.KeyValueExpr)(0x14008478660)({
            Key: (*ast.Ident)(0x14007ffec40)(FirstName),
            Colon: (token.Pos) 4233874,
            Value: (*ast.BasicLit)(0x14007ffec60)({
             ValuePos: (token.Pos) 4233876,
             Kind: (token.Token) STRING,
             Value: (string) (len=5) "\"Bob\""
            })
           }),
           (*ast.KeyValueExpr)(0x14008478690)({
            Key: (*ast.Ident)(0x14007ffec80)(LastName),
            Colon: (token.Pos) 4233894,
            Value: (*ast.BasicLit)(0x14007ffeca0)({
             ValuePos: (token.Pos) 4233897,
             Kind: (token.Token) STRING,
             Value: (string) (len=7) "\"Smith\""
            })
           }),
           (*ast.KeyValueExpr)(0x140084786c0)({
            Key: (*ast.Ident)(0x14007ffece0)(Age),
            Colon: (token.Pos) 4233912,
            Value: (*ast.BasicLit)(0x14007ffed00)({
             ValuePos: (token.Pos) 4233920,
             Kind: (token.Token) INT,
             Value: (string) (len=2) "42"
            })
           })
          },
          Rbrace: (token.Pos) 4233926,
          Incomplete: (bool) false
         }),
         (*ast.CompositeLit)(0x1400847a280)({
          Type: (ast.Expr) <nil>,
          Lbrace: (token.Pos) 4233931,
          Elts: ([]ast.Expr) (len=3 cap=4) {
           (*ast.KeyValueExpr)(0x140084786f0)({
            Key: (*ast.Ident)(0x14007ffed20)(FirstName),
            Colon: (token.Pos) 4233945,
            Value: (*ast.BasicLit)(0x14007ffed40)({
             ValuePos: (token.Pos) 4233947,
             Kind: (token.Token) STRING,
             Value: (string) (len=7) "\"Alice\""
            })
           }),
           (*ast.KeyValueExpr)(0x14008478720)({
            Key: (*ast.Ident)(0x14007ffed60)(LastName),
            Colon: (token.Pos) 4233967,
            Value: (*ast.BasicLit)(0x14007ffed80)({
             ValuePos: (token.Pos) 4233970,
             Kind: (token.Token) STRING,
             Value: (string) (len=7) "\"Jones\""
            })
           }),
           (*ast.KeyValueExpr)(0x14008478750)({
            Key: (*ast.Ident)(0x14007ffedc0)(Age),
            Colon: (token.Pos) 4233985,
            Value: (*ast.BasicLit)(0x14007ffede0)({
             ValuePos: (token.Pos) 4233993,
             Kind: (token.Token) INT,
             Value: (string) (len=2) "35"
            })
           })
          },
          Rbrace: (token.Pos) 4233999,
          Incomplete: (bool) false
         }),
         (*ast.CompositeLit)(0x1400847a300)({
          Type: (ast.Expr) <nil>,
          Lbrace: (token.Pos) 4234004,
          Elts: ([]ast.Expr) (len=3 cap=4) {
           (*ast.KeyValueExpr)(0x14008478780)({
            Key: (*ast.Ident)(0x14007ffee20)(FirstName),
            Colon: (token.Pos) 4234018,
            Value: (*ast.BasicLit)(0x14007ffee40)({
             ValuePos: (token.Pos) 4234020,
             Kind: (token.Token) STRING,
             Value: (string) (len=9) "\"Charlie\""
            })
           }),
           (*ast.KeyValueExpr)(0x140084787b0)({
            Key: (*ast.Ident)(0x14007ffee60)(LastName),
            Colon: (token.Pos) 4234042,
            Value: (*ast.BasicLit)(0x14007ffee80)({
             ValuePos: (token.Pos) 4234045,
             Kind: (token.Token) STRING,
             Value: (string) (len=7) "\"Brown\""
            })
           }),
           (*ast.KeyValueExpr)(0x140084787e0)({
            Key: (*ast.Ident)(0x14007ffeec0)(Age),
            Colon: (token.Pos) 4234060,
            Value: (*ast.BasicLit)(0x14007ffeee0)({
             ValuePos: (token.Pos) 4234068,
             Kind: (token.Token) INT,
             Value: (string) (len=1) "8"
            })
           })
          },
          Rbrace: (token.Pos) 4234073,
          Incomplete: (bool) false
         })
        },
        Rbrace: (token.Pos) 4234077,
        Incomplete: (bool) false
       })
      }
     })
    },
    Rbrace: (token.Pos) 4234079
   })
  }),
  (*ast.FuncDecl)(0x14008478930)({
   Doc: (*ast.CommentGroup)(0x14007f9b8f0)({
    List: ([]*ast.Comment) (len=1 cap=1) {
     (*ast.Comment)(0x14007f9b8d8)({
      Slash: (token.Pos) 4234082,
      Text: (string) (len=37) "// Gets the name and age of a person."
     })
    }
   }),
   Recv: (*ast.FieldList)(<nil>),
   Name: (*ast.Ident)(0x14007ffef40)(GetNameAndAge),
   Type: (*ast.FuncType)(0x14007fff180)({
    Func: (token.Pos) 4234120,
    TypeParams: (*ast.FieldList)(<nil>),
    Params: (*ast.FieldList)(0x14008478870)({
     Opening: (token.Pos) 4234138,
     List: ([]*ast.Field) <nil>,
     Closing: (token.Pos) 4234139
    }),
    Results: (*ast.FieldList)(0x140084788d0)({
     Opening: (token.Pos) 4234141,
     List: ([]*ast.Field) (len=2 cap=2) {
      (*ast.Field)(0x1400847a3c0)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) (len=1 cap=1) {
        (*ast.Ident)(0x14007ffef60)(name)
       },
       Type: (*ast.Ident)(0x14007ffef80)(string),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      }),
      (*ast.Field)(0x1400847a400)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) (len=1 cap=1) {
        (*ast.Ident)(0x14007ffefa0)(age)
       },
       Type: (*ast.Ident)(0x14007ffefc0)(int),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      })
     },
     Closing: (token.Pos) 4234162
    })
   }),
   Body: (*ast.BlockStmt)(0x14008478900)({
    Lbrace: (token.Pos) 4234164,
    List: ([]ast.Stmt) (len=2 cap=2) {
     (*ast.AssignStmt)(0x1400847a480)({
      Lhs: ([]ast.Expr) (len=1 cap=1) {
       (*ast.Ident)(0x14007ffefe0)(p)
      },
      TokPos: (token.Pos) 4234169,
      Tok: (token.Token) :=,
      Rhs: ([]ast.Expr) (len=1 cap=1) {
       (*ast.CallExpr)(0x1400847a440)({
        Fun: (*ast.Ident)(0x14007fff000)(GetPerson),
        Lparen: (token.Pos) 4234181,
        Args: ([]ast.Expr) <nil>,
        Ellipsis: (token.Pos) 0,
        Rparen: (token.Pos) 4234182
       })
      }
     }),
     (*ast.ReturnStmt)(0x14007fff140)({
      Return: (token.Pos) 4234185,
      Results: ([]ast.Expr) (len=2 cap=2) {
       (*ast.CallExpr)(0x1400847a4c0)({
        Fun: (*ast.Ident)(0x14007fff020)(GetFullName),
        Lparen: (token.Pos) 4234203,
        Args: ([]ast.Expr) (len=2 cap=2) {
         (*ast.SelectorExpr)(0x14007f9b920)({
          X: (*ast.Ident)(0x14007fff040)(p),
          Sel: (*ast.Ident)(0x14007fff060)(FirstName)
         }),
         (*ast.SelectorExpr)(0x14007f9b938)({
          X: (*ast.Ident)(0x14007fff080)(p),
          Sel: (*ast.Ident)(0x14007fff0a0)(LastName)
         })
        },
        Ellipsis: (token.Pos) 0,
        Rparen: (token.Pos) 4234227
       }),
       (*ast.SelectorExpr)(0x14007f9b950)({
        X: (*ast.Ident)(0x14007fff0e0)(p),
        Sel: (*ast.Ident)(0x14007fff100)(Age)
       })
      }
     })
    },
    Rbrace: (token.Pos) 4234236
   })
  }),
  (*ast.FuncDecl)(0x14008478ab0)({
   Doc: (*ast.CommentGroup)(0x14007f9b980)({
    List: ([]*ast.Comment) (len=1 cap=1) {
     (*ast.Comment)(0x14007f9b968)({
      Slash: (token.Pos) 4234239,
      Text: (string) (len=28) "// Tests returning an error."
     })
    }
   }),
   Recv: (*ast.FieldList)(<nil>),
   Name: (*ast.Ident)(0x14007fff1a0)(TestNormalError),
   Type: (*ast.FuncType)(0x14007fff460)({
    Func: (token.Pos) 4234268,
    TypeParams: (*ast.FieldList)(<nil>),
    Params: (*ast.FieldList)(0x14008478960)({
     Opening: (token.Pos) 4234288,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x1400847a500)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) (len=1 cap=1) {
        (*ast.Ident)(0x14007fff1c0)(input)
       },
       Type: (*ast.Ident)(0x14007fff1e0)(string),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      })
     },
     Closing: (token.Pos) 4234301
    }),
    Results: (*ast.FieldList)(0x140084789c0)({
     Opening: (token.Pos) 4234303,
     List: ([]*ast.Field) (len=2 cap=2) {
      (*ast.Field)(0x1400847a540)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) <nil>,
       Type: (*ast.Ident)(0x14007fff200)(string),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      }),
      (*ast.Field)(0x1400847a580)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) <nil>,
       Type: (*ast.Ident)(0x14007fff220)(error),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      })
     },
     Closing: (token.Pos) 4234317
    })
   }),
   Body: (*ast.BlockStmt)(0x14008478a80)({
    Lbrace: (token.Pos) 4234319,
    List: ([]ast.Stmt) (len=3 cap=4) {
     (*ast.IfStmt)(0x1400847a640)({
      If: (token.Pos) 4234645,
      Init: (ast.Stmt) <nil>,
      Cond: (*ast.BinaryExpr)(0x140084789f0)({
       X: (*ast.Ident)(0x14007fff260)(input),
       OpPos: (token.Pos) 4234654,
       Op: (token.Token) ==,
       Y: (*ast.BasicLit)(0x14007fff280)({
        ValuePos: (token.Pos) 4234657,
        Kind: (token.Token) STRING,
        Value: (string) (len=2) "\"\""
       })
      }),
      Body: (*ast.BlockStmt)(0x14008478a20)({
       Lbrace: (token.Pos) 4234660,
       List: ([]ast.Stmt) (len=1 cap=1) {
        (*ast.ReturnStmt)(0x14007fff340)({
         Return: (token.Pos) 4234664,
         Results: ([]ast.Expr) (len=2 cap=2) {
          (*ast.BasicLit)(0x14007fff2a0)({
           ValuePos: (token.Pos) 4234671,
           Kind: (token.Token) STRING,
           Value: (string) (len=2) "\"\""
          }),
          (*ast.CallExpr)(0x1400847a600)({
           Fun: (*ast.SelectorExpr)(0x14007f9ba58)({
            X: (*ast.Ident)(0x14007fff2c0)(errors),
            Sel: (*ast.Ident)(0x14007fff2e0)(New)
           }),
           Lparen: (token.Pos) 4234685,
           Args: ([]ast.Expr) (len=1 cap=1) {
            (*ast.BasicLit)(0x14007fff300)({
             ValuePos: (token.Pos) 4234686,
             Kind: (token.Token) STRING,
             Value: (string) (len=16) "\"input is empty\""
            })
           },
           Ellipsis: (token.Pos) 0,
           Rparen: (token.Pos) 4234702
          })
         }
        })
       },
       Rbrace: (token.Pos) 4234705
      }),
      Else: (ast.Stmt) <nil>
     }),
     (*ast.AssignStmt)(0x1400847a680)({
      Lhs: ([]ast.Expr) (len=1 cap=1) {
       (*ast.Ident)(0x14007fff360)(output)
      },
      TokPos: (token.Pos) 4234715,
      Tok: (token.Token) :=,
      Rhs: ([]ast.Expr) (len=1 cap=1) {
       (*ast.BinaryExpr)(0x14008478a50)({
        X: (*ast.BasicLit)(0x14007fff380)({
         ValuePos: (token.Pos) 4234718,
         Kind: (token.Token) STRING,
         Value: (string) (len=12) "\"You said: \""
        }),
        OpPos: (token.Pos) 4234731,
        Op: (token.Token) +,
        Y: (*ast.Ident)(0x14007fff3a0)(input)
       })
      }
     }),
     (*ast.ReturnStmt)(0x14007fff440)({
      Return: (token.Pos) 4234740,
      Results: ([]ast.Expr) (len=2 cap=2) {
       (*ast.Ident)(0x14007fff3e0)(output),
       (*ast.Ident)(0x14007fff400)(nil)
      }
     })
    },
    Rbrace: (token.Pos) 4234759
   })
  }),
  (*ast.FuncDecl)(0x14008478c00)({
   Doc: (*ast.CommentGroup)(0x14007f9ba88)({
    List: ([]*ast.Comment) (len=1 cap=1) {
     (*ast.Comment)(0x14007f9ba70)({
      Slash: (token.Pos) 4234762,
      Text: (string) (len=58) "// Tests an alternative way to handle errors in functions."
     })
    }
   }),
   Recv: (*ast.FieldList)(<nil>),
   Name: (*ast.Ident)(0x14007fff480)(TestAlternativeError),
   Type: (*ast.FuncType)(0x14007fff6c0)({
    Func: (token.Pos) 4234821,
    TypeParams: (*ast.FieldList)(<nil>),
    Params: (*ast.FieldList)(0x14008478ae0)({
     Opening: (token.Pos) 4234846,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x1400847a700)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) (len=1 cap=1) {
        (*ast.Ident)(0x14007fff4a0)(input)
       },
       Type: (*ast.Ident)(0x14007fff4c0)(string),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      })
     },
     Closing: (token.Pos) 4234859
    }),
    Results: (*ast.FieldList)(0x14008478b10)({
     Opening: (token.Pos) 0,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x1400847a740)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) <nil>,
       Type: (*ast.Ident)(0x14007fff4e0)(string),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      })
     },
     Closing: (token.Pos) 0
    })
   }),
   Body: (*ast.BlockStmt)(0x14008478bd0)({
    Lbrace: (token.Pos) 4234868,
    List: ([]ast.Stmt) (len=3 cap=4) {
     (*ast.IfStmt)(0x1400847a7c0)({
      If: (token.Pos) 4235012,
      Init: (ast.Stmt) <nil>,
      Cond: (*ast.BinaryExpr)(0x14008478b40)({
       X: (*ast.Ident)(0x14007fff500)(input),
       OpPos: (token.Pos) 4235021,
       Op: (token.Token) ==,
       Y: (*ast.BasicLit)(0x14007fff520)({
        ValuePos: (token.Pos) 4235024,
        Kind: (token.Token) STRING,
        Value: (string) (len=2) "\"\""
       })
      }),
      Body: (*ast.BlockStmt)(0x14008478b70)({
       Lbrace: (token.Pos) 4235027,
       List: ([]ast.Stmt) (len=2 cap=2) {
        (*ast.ExprStmt)(0x14007fc7c50)({
         X: (*ast.CallExpr)(0x1400847a780)({
          Fun: (*ast.SelectorExpr)(0x14007f9bb00)({
           X: (*ast.Ident)(0x14007fff540)(console),
           Sel: (*ast.Ident)(0x14007fff560)(Error)
          }),
          Lparen: (token.Pos) 4235044,
          Args: ([]ast.Expr) (len=1 cap=1) {
           (*ast.BasicLit)(0x14007fff580)({
            ValuePos: (token.Pos) 4235045,
            Kind: (token.Token) STRING,
            Value: (string) (len=16) "\"input is empty\""
           })
          },
          Ellipsis: (token.Pos) 0,
          Rparen: (token.Pos) 4235061
         })
        }),
        (*ast.ReturnStmt)(0x14007fff5c0)({
         Return: (token.Pos) 4235065,
         Results: ([]ast.Expr) (len=1 cap=1) {
          (*ast.BasicLit)(0x14007fff5a0)({
           ValuePos: (token.Pos) 4235072,
           Kind: (token.Token) STRING,
           Value: (string) (len=2) "\"\""
          })
         }
        })
       },
       Rbrace: (token.Pos) 4235076
      }),
      Else: (ast.Stmt) <nil>
     }),
     (*ast.AssignStmt)(0x1400847a800)({
      Lhs: ([]ast.Expr) (len=1 cap=1) {
       (*ast.Ident)(0x14007fff600)(output)
      },
      TokPos: (token.Pos) 4235086,
      Tok: (token.Token) :=,
      Rhs: ([]ast.Expr) (len=1 cap=1) {
       (*ast.BinaryExpr)(0x14008478ba0)({
        X: (*ast.BasicLit)(0x14007fff620)({
         ValuePos: (token.Pos) 4235089,
         Kind: (token.Token) STRING,
         Value: (string) (len=12) "\"You said: \""
        }),
        OpPos: (token.Pos) 4235102,
        Op: (token.Token) +,
        Y: (*ast.Ident)(0x14007fff640)(input)
       })
      }
     }),
     (*ast.ReturnStmt)(0x14007fff6a0)({
      Return: (token.Pos) 4235111,
      Results: ([]ast.Expr) (len=1 cap=1) {
       (*ast.Ident)(0x14007fff680)(output)
      }
     })
    },
    Rbrace: (token.Pos) 4235125
   })
  }),
  (*ast.FuncDecl)(0x14008478c90)({
   Doc: (*ast.CommentGroup)(0x14007f9bb30)({
    List: ([]*ast.Comment) (len=1 cap=1) {
     (*ast.Comment)(0x14007f9bb18)({
      Slash: (token.Pos) 4235128,
      Text: (string) (len=17) "// Tests a panic."
     })
    }
   }),
   Recv: (*ast.FieldList)(<nil>),
   Name: (*ast.Ident)(0x14007fff6e0)(TestPanic),
   Type: (*ast.FuncType)(0x14007fff740)({
    Func: (token.Pos) 4235146,
    TypeParams: (*ast.FieldList)(<nil>),
    Params: (*ast.FieldList)(0x14008478c30)({
     Opening: (token.Pos) 4235160,
     List: ([]*ast.Field) <nil>,
     Closing: (token.Pos) 4235161
    }),
    Results: (*ast.FieldList)(<nil>)
   }),
   Body: (*ast.BlockStmt)(0x14008478c60)({
    Lbrace: (token.Pos) 4235163,
    List: ([]ast.Stmt) (len=1 cap=1) {
     (*ast.ExprStmt)(0x14007fc7cf0)({
      X: (*ast.CallExpr)(0x1400847a880)({
       Fun: (*ast.Ident)(0x14007fff700)(panic),
       Lparen: (token.Pos) 4235283,
       Args: ([]ast.Expr) (len=1 cap=1) {
        (*ast.BasicLit)(0x14007fff720)({
         ValuePos: (token.Pos) 4235284,
         Kind: (token.Token) STRING,
         Value: (string) (len=72) "\"This is a message from a panic.\\nThis is a second line from a panic.\\n\""
        })
       },
       Ellipsis: (token.Pos) 0,
       Rparen: (token.Pos) 4235356
      })
     })
    },
    Rbrace: (token.Pos) 4235358
   })
  }),
  (*ast.FuncDecl)(0x14008478d20)({
   Doc: (*ast.CommentGroup)(0x14007f9bba8)({
    List: ([]*ast.Comment) (len=1 cap=1) {
     (*ast.Comment)(0x14007f9bb90)({
      Slash: (token.Pos) 4235361,
      Text: (string) (len=43) "// Tests an exit with a non-zero exit code."
     })
    }
   }),
   Recv: (*ast.FieldList)(<nil>),
   Name: (*ast.Ident)(0x14007fff760)(TestExit),
   Type: (*ast.FuncType)(0x14007fff8c0)({
    Func: (token.Pos) 4235405,
    TypeParams: (*ast.FieldList)(<nil>),
    Params: (*ast.FieldList)(0x14008478cc0)({
     Opening: (token.Pos) 4235418,
     List: ([]*ast.Field) <nil>,
     Closing: (token.Pos) 4235419
    }),
    Results: (*ast.FieldList)(<nil>)
   }),
   Body: (*ast.BlockStmt)(0x14008478cf0)({
    Lbrace: (token.Pos) 4235421,
    List: ([]ast.Stmt) (len=3 cap=4) {
     (*ast.ExprStmt)(0x14007fc7d40)({
      X: (*ast.CallExpr)(0x1400847a8c0)({
       Fun: (*ast.SelectorExpr)(0x14007f9bc38)({
        X: (*ast.Ident)(0x14007fff7a0)(console),
        Sel: (*ast.Ident)(0x14007fff7c0)(Error)
       }),
       Lparen: (token.Pos) 4235730,
       Args: ([]ast.Expr) (len=1 cap=1) {
        (*ast.BasicLit)(0x14007fff7e0)({
         ValuePos: (token.Pos) 4235731,
         Kind: (token.Token) STRING,
         Value: (string) (len=27) "\"This is an error message.\""
        })
       },
       Ellipsis: (token.Pos) 0,
       Rparen: (token.Pos) 4235758
      })
     }),
     (*ast.ExprStmt)(0x14007fc7d80)({
      X: (*ast.CallExpr)(0x1400847a900)({
       Fun: (*ast.SelectorExpr)(0x14007f9bc50)({
        X: (*ast.Ident)(0x14007fff800)(os),
        Sel: (*ast.Ident)(0x14007fff820)(Exit)
       }),
       Lparen: (token.Pos) 4235768,
       Args: ([]ast.Expr) (len=1 cap=1) {
        (*ast.BasicLit)(0x14007fff840)({
         ValuePos: (token.Pos) 4235769,
         Kind: (token.Token) INT,
         Value: (string) (len=1) "1"
        })
       },
       Ellipsis: (token.Pos) 0,
       Rparen: (token.Pos) 4235770
      })
     }),
     (*ast.ExprStmt)(0x14007fc7db0)({
      X: (*ast.CallExpr)(0x1400847a940)({
       Fun: (*ast.Ident)(0x14007fff880)(println),
       Lparen: (token.Pos) 4235780,
       Args: ([]ast.Expr) (len=1 cap=1) {
        (*ast.BasicLit)(0x14007fff8a0)({
         ValuePos: (token.Pos) 4235781,
         Kind: (token.Token) STRING,
         Value: (string) (len=33) "\"This line will not be executed.\""
        })
       },
       Ellipsis: (token.Pos) 0,
       Rparen: (token.Pos) 4235814
      })
     })
    },
    Rbrace: (token.Pos) 4235816
   })
  }),
  (*ast.FuncDecl)(0x14008478db0)({
   Doc: (*ast.CommentGroup)(0x14007f9bc80)({
    List: ([]*ast.Comment) (len=1 cap=1) {
     (*ast.Comment)(0x14007f9bc68)({
      Slash: (token.Pos) 4235819,
      Text: (string) (len=37) "// Tests logging at different levels."
     })
    }
   }),
   Recv: (*ast.FieldList)(<nil>),
   Name: (*ast.Ident)(0x14007fff8e0)(TestLogging),
   Type: (*ast.FuncType)(0x14007fffe40)({
    Func: (token.Pos) 4235857,
    TypeParams: (*ast.FieldList)(<nil>),
    Params: (*ast.FieldList)(0x14008478d50)({
     Opening: (token.Pos) 4235873,
     List: ([]*ast.Field) <nil>,
     Closing: (token.Pos) 4235874
    }),
    Results: (*ast.FieldList)(<nil>)
   }),
   Body: (*ast.BlockStmt)(0x14008478d80)({
    Lbrace: (token.Pos) 4235876,
    List: ([]ast.Stmt) (len=11 cap=16) {
     (*ast.ExprStmt)(0x14007fc7de0)({
      X: (*ast.CallExpr)(0x1400847a9c0)({
       Fun: (*ast.SelectorExpr)(0x14007f9bcc8)({
        X: (*ast.Ident)(0x14007fff900)(console),
        Sel: (*ast.Ident)(0x14007fff920)(Log)
       }),
       Lparen: (token.Pos) 4235941,
       Args: ([]ast.Expr) (len=1 cap=1) {
        (*ast.BasicLit)(0x14007fff940)({
         ValuePos: (token.Pos) 4235942,
         Kind: (token.Token) STRING,
         Value: (string) (len=31) "\"This is a simple log message.\""
        })
       },
       Ellipsis: (token.Pos) 0,
       Rparen: (token.Pos) 4235973
      })
     }),
     (*ast.ExprStmt)(0x14007fc7e20)({
      X: (*ast.CallExpr)(0x1400847aa00)({
       Fun: (*ast.SelectorExpr)(0x14007f9bd10)({
        X: (*ast.Ident)(0x14007fff960)(console),
        Sel: (*ast.Ident)(0x14007fff980)(Debug)
       }),
       Lparen: (token.Pos) 4236041,
       Args: ([]ast.Expr) (len=1 cap=1) {
        (*ast.BasicLit)(0x14007fff9a0)({
         ValuePos: (token.Pos) 4236042,
         Kind: (token.Token) STRING,
         Value: (string) (len=26) "\"This is a debug message.\""
        })
       },
       Ellipsis: (token.Pos) 0,
       Rparen: (token.Pos) 4236068
      })
     }),
     (*ast.ExprStmt)(0x14007fc7e50)({
      X: (*ast.CallExpr)(0x1400847aa40)({
       Fun: (*ast.SelectorExpr)(0x14007f9bd28)({
        X: (*ast.Ident)(0x14007fff9e0)(console),
        Sel: (*ast.Ident)(0x14007fffa00)(Info)
       }),
       Lparen: (token.Pos) 4236083,
       Args: ([]ast.Expr) (len=1 cap=1) {
        (*ast.BasicLit)(0x14007fffa20)({
         ValuePos: (token.Pos) 4236084,
         Kind: (token.Token) STRING,
         Value: (string) (len=26) "\"This is an info message.\""
        })
       },
       Ellipsis: (token.Pos) 0,
       Rparen: (token.Pos) 4236110
      })
     }),
     (*ast.ExprStmt)(0x14007fc7e80)({
      X: (*ast.CallExpr)(0x1400847aac0)({
       Fun: (*ast.SelectorExpr)(0x14007f9bd40)({
        X: (*ast.Ident)(0x14007fffa40)(console),
        Sel: (*ast.Ident)(0x14007fffa60)(Warn)
       }),
       Lparen: (token.Pos) 4236125,
       Args: ([]ast.Expr) (len=1 cap=1) {
        (*ast.BasicLit)(0x14007fffa80)({
         ValuePos: (token.Pos) 4236126,
         Kind: (token.Token) STRING,
         Value: (string) (len=28) "\"This is a warning message.\""
        })
       },
       Ellipsis: (token.Pos) 0,
       Rparen: (token.Pos) 4236154
      })
     }),
     (*ast.ExprStmt)(0x14007fc7eb0)({
      X: (*ast.CallExpr)(0x1400847ab00)({
       Fun: (*ast.SelectorExpr)(0x14007f9bd88)({
        X: (*ast.Ident)(0x14007fffaa0)(console),
        Sel: (*ast.Ident)(0x14007fffac0)(Error)
       }),
       Lparen: (token.Pos) 4236240,
       Args: ([]ast.Expr) (len=1 cap=1) {
        (*ast.BasicLit)(0x14007fffae0)({
         ValuePos: (token.Pos) 4236241,
         Kind: (token.Token) STRING,
         Value: (string) (len=27) "\"This is an error message.\""
        })
       },
       Ellipsis: (token.Pos) 0,
       Rparen: (token.Pos) 4236268
      })
     }),
     (*ast.ExprStmt)(0x14007fc7ee0)({
      X: (*ast.CallExpr)(0x1400847ab40)({
       Fun: (*ast.SelectorExpr)(0x14007f9bda0)({
        X: (*ast.Ident)(0x14007fffb00)(console),
        Sel: (*ast.Ident)(0x14007fffb20)(Error)
       }),
       Lparen: (token.Pos) 4236284,
       Args: ([]ast.Expr) (len=1 cap=1) {
        (*ast.BasicLit)(0x14007fffb40)({
         ValuePos: (token.Pos) 4236288,
         Kind: (token.Token) STRING,
         Value: (string) (len=143) "`This is line 1 of a multi-line error message.\n  This is line 2 of a multi-line error message.\n  This is line 3 of a multi-line error message.`"
        })
       },
       Ellipsis: (token.Pos) 0,
       Rparen: (token.Pos) 4236434
      })
     }),
     (*ast.ExprStmt)(0x14007fc7f10)({
      X: (*ast.CallExpr)(0x1400847ab80)({
       Fun: (*ast.Ident)(0x14007fffb60)(println),
       Lparen: (token.Pos) 4236499,
       Args: ([]ast.Expr) (len=1 cap=1) {
        (*ast.BasicLit)(0x14007fffb80)({
         ValuePos: (token.Pos) 4236500,
         Kind: (token.Token) STRING,
         Value: (string) (len=28) "\"This is a println message.\""
        })
       },
       Ellipsis: (token.Pos) 0,
       Rparen: (token.Pos) 4236528
      })
     }),
     (*ast.ExprStmt)(0x14007fc7f40)({
      X: (*ast.CallExpr)(0x1400847abc0)({
       Fun: (*ast.SelectorExpr)(0x14007f9bde8)({
        X: (*ast.Ident)(0x14007fffba0)(fmt),
        Sel: (*ast.Ident)(0x14007fffbc0)(Println)
       }),
       Lparen: (token.Pos) 4236542,
       Args: ([]ast.Expr) (len=1 cap=1) {
        (*ast.BasicLit)(0x14007fffbe0)({
         ValuePos: (token.Pos) 4236543,
         Kind: (token.Token) STRING,
         Value: (string) (len=32) "\"This is a fmt.Println message.\""
        })
       },
       Ellipsis: (token.Pos) 0,
       Rparen: (token.Pos) 4236575
      })
     }),
     (*ast.ExprStmt)(0x14007fc7f70)({
      X: (*ast.CallExpr)(0x1400847ac00)({
       Fun: (*ast.SelectorExpr)(0x14007f9be00)({
        X: (*ast.Ident)(0x14007fffc00)(fmt),
        Sel: (*ast.Ident)(0x14007fffc20)(Printf)
       }),
       Lparen: (token.Pos) 4236588,
       Args: ([]ast.Expr) (len=2 cap=2) {
        (*ast.BasicLit)(0x14007fffc40)({
         ValuePos: (token.Pos) 4236589,
         Kind: (token.Token) STRING,
         Value: (string) (len=38) "\"This is a fmt.Printf message (%s).\\n\""
        }),
        (*ast.BasicLit)(0x14007fffc60)({
         ValuePos: (token.Pos) 4236629,
         Kind: (token.Token) STRING,
         Value: (string) (len=18) "\"with a parameter\""
        })
       },
       Ellipsis: (token.Pos) 0,
       Rparen: (token.Pos) 4236647
      })
     }),
     (*ast.ExprStmt)(0x1400847e020)({
      X: (*ast.CallExpr)(0x1400847ac40)({
       Fun: (*ast.SelectorExpr)(0x14007f9be48)({
        X: (*ast.Ident)(0x14007fffca0)(fmt),
        Sel: (*ast.Ident)(0x14007fffcc0)(Fprintln)
       }),
       Lparen: (token.Pos) 4236761,
       Args: ([]ast.Expr) (len=2 cap=2) {
        (*ast.SelectorExpr)(0x14007f9be60)({
         X: (*ast.Ident)(0x14007fffce0)(os),
         Sel: (*ast.Ident)(0x14007fffd00)(Stdout)
        }),
        (*ast.BasicLit)(0x14007fffd20)({
         ValuePos: (token.Pos) 4236773,
         Kind: (token.Token) STRING,
         Value: (string) (len=44) "\"This is an info message printed to stdout.\""
        })
       },
       Ellipsis: (token.Pos) 0,
       Rparen: (token.Pos) 4236817
      })
     }),
     (*ast.ExprStmt)(0x1400847e050)({
      X: (*ast.CallExpr)(0x1400847ac80)({
       Fun: (*ast.SelectorExpr)(0x14007f9be78)({
        X: (*ast.Ident)(0x14007fffd60)(fmt),
        Sel: (*ast.Ident)(0x14007fffd80)(Fprintln)
       }),
       Lparen: (token.Pos) 4236832,
       Args: ([]ast.Expr) (len=2 cap=2) {
        (*ast.SelectorExpr)(0x14007f9be90)({
         X: (*ast.Ident)(0x14007fffda0)(os),
         Sel: (*ast.Ident)(0x14007fffdc0)(Stderr)
        }),
        (*ast.BasicLit)(0x14007fffde0)({
         ValuePos: (token.Pos) 4236844,
         Kind: (token.Token) STRING,
         Value: (string) (len=45) "\"This is an error message printed to stderr.\""
        })
       },
       Ellipsis: (token.Pos) 0,
       Rparen: (token.Pos) 4236889
      })
     })
    },
    Rbrace: (token.Pos) 4237450
   })
  })
 },
 FileStart: (token.Pos) 4231691,
 FileEnd: (token.Pos) 4237452,
 Scope: (*ast.Scope)(0x1400847e070)(scope 0x1400847e070 {
	func TestPanic
	func Add3
	func GetCurrentTime
	func GetCurrentTimeFormatted
	func GetPerson
	func TestNormalError
	func TestAlternativeError
	func SayHello
	func GetPeople
	func TestExit
	func LogMessage
	func Add
	func AddN
	var nowFunc
	func GetFullName
	type Person
	func GetRandomPerson
	func GetNameAndAge
	func TestLogging
}
),
 Imports: ([]*ast.ImportSpec) (len=6 cap=8) {
  (*ast.ImportSpec)(0x14007fdd890)({
   Doc: (*ast.CommentGroup)(<nil>),
   Name: (*ast.Ident)(<nil>),
   Path: (*ast.BasicLit)(0x14007fa3a20)({
    ValuePos: (token.Pos) 4231951,
    Kind: (token.Token) STRING,
    Value: (string) (len=8) "\"errors\""
   }),
   Comment: (*ast.CommentGroup)(<nil>),
   EndPos: (token.Pos) 0
  }),
  (*ast.ImportSpec)(0x14007fdd8c0)({
   Doc: (*ast.CommentGroup)(<nil>),
   Name: (*ast.Ident)(<nil>),
   Path: (*ast.BasicLit)(0x14007fa3a40)({
    ValuePos: (token.Pos) 4231961,
    Kind: (token.Token) STRING,
    Value: (string) (len=5) "\"fmt\""
   }),
   Comment: (*ast.CommentGroup)(<nil>),
   EndPos: (token.Pos) 0
  }),
  (*ast.ImportSpec)(0x14007fdd8f0)({
   Doc: (*ast.CommentGroup)(<nil>),
   Name: (*ast.Ident)(<nil>),
   Path: (*ast.BasicLit)(0x14007fa3a80)({
    ValuePos: (token.Pos) 4231968,
    Kind: (token.Token) STRING,
    Value: (string) (len=11) "\"math/rand\""
   }),
   Comment: (*ast.CommentGroup)(<nil>),
   EndPos: (token.Pos) 0
  }),
  (*ast.ImportSpec)(0x14007fdd920)({
   Doc: (*ast.CommentGroup)(<nil>),
   Name: (*ast.Ident)(<nil>),
   Path: (*ast.BasicLit)(0x14007fa3ac0)({
    ValuePos: (token.Pos) 4231981,
    Kind: (token.Token) STRING,
    Value: (string) (len=4) "\"os\""
   }),
   Comment: (*ast.CommentGroup)(<nil>),
   EndPos: (token.Pos) 0
  }),
  (*ast.ImportSpec)(0x14007fdd950)({
   Doc: (*ast.CommentGroup)(<nil>),
   Name: (*ast.Ident)(<nil>),
   Path: (*ast.BasicLit)(0x14007fa3ae0)({
    ValuePos: (token.Pos) 4231987,
    Kind: (token.Token) STRING,
    Value: (string) (len=6) "\"time\""
   }),
   Comment: (*ast.CommentGroup)(<nil>),
   EndPos: (token.Pos) 0
  }),
  (*ast.ImportSpec)(0x14007fdd980)({
   Doc: (*ast.CommentGroup)(<nil>),
   Name: (*ast.Ident)(<nil>),
   Path: (*ast.BasicLit)(0x14007fa3b00)({
    ValuePos: (token.Pos) 4231996,
    Kind: (token.Token) STRING,
    Value: (string) (len=50) "\"github.com/hypermodeinc/modus/sdk/go/pkg/console\""
   }),
   Comment: (*ast.CommentGroup)(<nil>),
   EndPos: (token.Pos) 0
  })
 },
 Unresolved: ([]*ast.Ident) (len=51 cap=64) {
  (*ast.Ident)(0x14007fa3b60)(string),
  (*ast.Ident)(0x14007fa3b80)(console),
  (*ast.Ident)(0x14007fa3ca0)(int),
  (*ast.Ident)(0x14007fa3cc0)(int),
  (*ast.Ident)(0x14007fa3dc0)(int),
  (*ast.Ident)(0x14007fa3e00)(int),
  (*ast.Ident)(0x14007fa3e20)(int),
  (*ast.Ident)(0x14007fa3e60)(nil),
  (*ast.Ident)(0x14007ffe060)(int),
  (*ast.Ident)(0x14007ffe080)(int),
  (*ast.Ident)(0x14007ffe260)(time),
  (*ast.Ident)(0x14007ffe2c0)(time),
  (*ast.Ident)(0x14007ffe380)(string),
  (*ast.Ident)(0x14007ffe3e0)(time),
  (*ast.Ident)(0x14007ffe4c0)(string),
  (*ast.Ident)(0x14007ffe4e0)(string),
  (*ast.Ident)(0x14007ffe5e0)(string),
  (*ast.Ident)(0x14007ffe600)(string),
  (*ast.Ident)(0x14007ffe640)(nil),
  (*ast.Ident)(0x14007ffe780)(string),
  (*ast.Ident)(0x14007ffe7e0)(string),
  (*ast.Ident)(0x14007ffe840)(int),
  (*ast.Ident)(0x14007ffeac0)(rand),
  (*ast.Ident)(0x14007ffeb00)(len),
  (*ast.Ident)(0x14007ffef80)(string),
  (*ast.Ident)(0x14007ffefc0)(int),
  (*ast.Ident)(0x14007fff1e0)(string),
  (*ast.Ident)(0x14007fff200)(string),
  (*ast.Ident)(0x14007fff220)(error),
  (*ast.Ident)(0x14007fff2c0)(errors),
  (*ast.Ident)(0x14007fff400)(nil),
  (*ast.Ident)(0x14007fff4c0)(string),
  (*ast.Ident)(0x14007fff4e0)(string),
  (*ast.Ident)(0x14007fff540)(console),
  (*ast.Ident)(0x14007fff700)(panic),
  (*ast.Ident)(0x14007fff7a0)(console),
  (*ast.Ident)(0x14007fff800)(os),
  (*ast.Ident)(0x14007fff880)(println),
  (*ast.Ident)(0x14007fff900)(console),
  (*ast.Ident)(0x14007fff960)(console),
  (*ast.Ident)(0x14007fff9e0)(console),
  (*ast.Ident)(0x14007fffa40)(console),
  (*ast.Ident)(0x14007fffaa0)(console),
  (*ast.Ident)(0x14007fffb00)(console),
  (*ast.Ident)(0x14007fffb60)(println),
  (*ast.Ident)(0x14007fffba0)(fmt),
  (*ast.Ident)(0x14007fffc00)(fmt),
  (*ast.Ident)(0x14007fffca0)(fmt),
  (*ast.Ident)(0x14007fffce0)(os),
  (*ast.Ident)(0x14007fffd60)(fmt),
  (*ast.Ident)(0x14007fffda0)(os)
 },
 Comments: ([]*ast.CommentGroup) (len=33 cap=64) {
  (*ast.CommentGroup)(0x14007f9b3e0)({
   List: ([]*ast.Comment) (len=1 cap=1) {
    (*ast.Comment)(0x14007f9b3c8)({
     Slash: (token.Pos) 4231691,
     Text: (string) (len=234) "/*\n * This example is part of the Modus project, licensed under the Apache License 2.0.\n * You may modify and use this example in accordance with the license.\n * See the LICENSE file that accompanied this code for further details.\n * /"  //GML
    })
   }
  }),
  (*ast.CommentGroup)(0x14007f9b410)({
   List: ([]*ast.Comment) (len=1 cap=1) {
    (*ast.Comment)(0x14007f9b3f8)({
     Slash: (token.Pos) 4232050,
     Text: (string) (len=18) "// Logs a message."
    })
   }
  }),
  (*ast.CommentGroup)(0x14007f9b470)({
   List: ([]*ast.Comment) (len=1 cap=1) {
    (*ast.Comment)(0x14007f9b458)({
     Slash: (token.Pos) 4232128,
     Text: (string) (len=53) "// Adds two integers together and returns the result."
    })
   }
  }),
  (*ast.CommentGroup)(0x14007f9b4d0)({
   List: ([]*ast.Comment) (len=2 cap=2) {
    (*ast.Comment)(0x14007f9b4a0)({
     Slash: (token.Pos) 4232224,
     Text: (string) (len=55) "// Adds three integers together and returns the result."
    }),
    (*ast.Comment)(0x14007f9b4b8)({
     Slash: (token.Pos) 4232280,
     Text: (string) (len=33) "// The third integer is optional."
    })
   }
  }),
  (*ast.CommentGroup)(0x14007f9b548)({
   List: ([]*ast.Comment) (len=1 cap=1) {
    (*ast.Comment)(0x14007f9b530)({
     Slash: (token.Pos) 4232403,
     Text: (string) (len=63) "// Adds any number of integers together and returns the result."
    })
   }
  }),
  (*ast.CommentGroup)(0x14007f9b5a8)({
   List: ([]*ast.Comment) (len=1 cap=1) {
    (*ast.Comment)(0x14007f9b590)({
     Slash: (token.Pos) 4232565,
     Text: (string) (len=55) "// this indirection is so we can mock time.Now in tests"
    })
   }
  }),
  (*ast.CommentGroup)(0x14007f9b5f0)({
   List: ([]*ast.Comment) (len=1 cap=1) {
    (*ast.Comment)(0x14007f9b5d8)({
     Slash: (token.Pos) 4232645,
     Text: (string) (len=28) "// Returns the current time."
    })
   }
  }),
  (*ast.CommentGroup)(0x14007f9b638)({
   List: ([]*ast.Comment) (len=1 cap=1) {
    (*ast.Comment)(0x14007f9b620)({
     Slash: (token.Pos) 4232729,
     Text: (string) (len=50) "// Returns the current time formatted as a string."
    })
   }
  }),
  (*ast.CommentGroup)(0x14007f9b698)({
   List: ([]*ast.Comment) (len=1 cap=1) {
    (*ast.Comment)(0x14007f9b680)({
     Slash: (token.Pos) 4232863,
     Text: (string) (len=75) "// Combines the first and last name of a person, and returns the full name."
    })
   }
  }),
  (*ast.CommentGroup)(0x14007f9b6f8)({
   List: ([]*ast.Comment) (len=2 cap=2) {
    (*ast.Comment)(0x14007f9b6c8)({
     Slash: (token.Pos) 4233031,
     Text: (string) (len=34) "// Says hello to a person by name."
    }),
    (*ast.Comment)(0x14007f9b6e0)({
     Slash: (token.Pos) 4233066,
     Text: (string) (len=65) "// If the name is not provided, it will say hello without a name."
    })
   }
  }),
  (*ast.CommentGroup)(0x14007f9b770)({
   List: ([]*ast.Comment) (len=1 cap=1) {
    (*ast.Comment)(0x14007f9b758)({
     Slash: (token.Pos) 4233254,
     Text: (string) (len=41) "// A simple object representing a person."
    })
   }
  }),
  (*ast.CommentGroup)(0x14007f9b7a0)({
   List: ([]*ast.Comment) (len=1 cap=1) {
    (*ast.Comment)(0x14007f9b788)({
     Slash: (token.Pos) 4233319,
     Text: (string) (len=27) "// The person's first name."
    })
   }
  }),
  (*ast.CommentGroup)(0x14007f9b7d0)({
   List: ([]*ast.Comment) (len=1 cap=1) {
    (*ast.Comment)(0x14007f9b7b8)({
     Slash: (token.Pos) 4233386,
     Text: (string) (len=26) "// The person's last name."
    })
   }
  }),
  (*ast.CommentGroup)(0x14007f9b800)({
   List: ([]*ast.Comment) (len=1 cap=1) {
    (*ast.Comment)(0x14007f9b7e8)({
     Slash: (token.Pos) 4233450,
     Text: (string) (len=20) "// The person's age."
    })
   }
  }),
  (*ast.CommentGroup)(0x14007f9b848)({
   List: ([]*ast.Comment) (len=1 cap=1) {
    (*ast.Comment)(0x14007f9b830)({
     Slash: (token.Pos) 4233496,
     Text: (string) (len=24) "// Gets a person object."
    })
   }
  }),
  (*ast.CommentGroup)(0x14007f9b878)({
   List: ([]*ast.Comment) (len=1 cap=1) {
    (*ast.Comment)(0x14007f9b860)({
     Slash: (token.Pos) 4233627,
     Text: (string) (len=53) "// Gets a random person object from a list of people."
    })
   }
  }),
  (*ast.CommentGroup)(0x14007f9b8c0)({
   List: ([]*ast.Comment) (len=1 cap=1) {
    (*ast.Comment)(0x14007f9b8a8)({
     Slash: (token.Pos) 4233786,
     Text: (string) (len=25) "// Gets a list of people."
    })
   }
  }),
  (*ast.CommentGroup)(0x14007f9b8f0)({
   List: ([]*ast.Comment) (len=1 cap=1) {
    (*ast.Comment)(0x14007f9b8d8)({
     Slash: (token.Pos) 4234082,
     Text: (string) (len=37) "// Gets the name and age of a person."
    })
   }
  }),
  (*ast.CommentGroup)(0x14007f9b980)({
   List: ([]*ast.Comment) (len=1 cap=1) {
    (*ast.Comment)(0x14007f9b968)({
     Slash: (token.Pos) 4234239,
     Text: (string) (len=28) "// Tests returning an error."
    })
   }
  }),
  (*ast.CommentGroup)(0x14007f9ba40)({
   List: ([]*ast.Comment) (len=5 cap=8) {
    (*ast.Comment)(0x14007f9b9c8)({
     Slash: (token.Pos) 4234323,
     Text: (string) (len=59) "// This is the preferred way to handle errors in functions."
    }),
    (*ast.Comment)(0x14007f9b9e0)({
     Slash: (token.Pos) 4234384,
     Text: (string) (len=62) "// Simply declare an error interface as the last return value."
    }),
    (*ast.Comment)(0x14007f9b9f8)({
     Slash: (token.Pos) 4234448,
     Text: (string) (len=65) "// You can use any object that implements the Go error interface."
    }),
    (*ast.Comment)(0x14007f9ba10)({
     Slash: (token.Pos) 4234515,
     Text: (string) (len=70) "// For example, you can create a new error with errors.New(\"message\"),"
    }),
    (*ast.Comment)(0x14007f9ba28)({
     Slash: (token.Pos) 4234587,
     Text: (string) (len=55) "// or with fmt.Errorf(\"message with %s\", \"parameters\")."
    })
   }
  }),
  (*ast.CommentGroup)(0x14007f9ba88)({
   List: ([]*ast.Comment) (len=1 cap=1) {
    (*ast.Comment)(0x14007f9ba70)({
     Slash: (token.Pos) 4234762,
     Text: (string) (len=58) "// Tests an alternative way to handle errors in functions."
    })
   }
  }),
  (*ast.CommentGroup)(0x14007f9bae8)({
   List: ([]*ast.Comment) (len=2 cap=2) {
    (*ast.Comment)(0x14007f9bab8)({
     Slash: (token.Pos) 4234872,
     Text: (string) (len=60) "// This is an alternative way to handle errors in functions."
    }),
    (*ast.Comment)(0x14007f9bad0)({
     Slash: (token.Pos) 4234934,
     Text: (string) (len=75) "// It is identical in behavior to TestNormalError, but is not Go idiomatic."
    })
   }
  }),
  (*ast.CommentGroup)(0x14007f9bb30)({
   List: ([]*ast.Comment) (len=1 cap=1) {
    (*ast.Comment)(0x14007f9bb18)({
     Slash: (token.Pos) 4235128,
     Text: (string) (len=17) "// Tests a panic."
    })
   }
  }),
  (*ast.CommentGroup)(0x14007f9bb78)({
   List: ([]*ast.Comment) (len=2 cap=2) {
    (*ast.Comment)(0x14007f9bb48)({
     Slash: (token.Pos) 4235167,
     Text: (string) (len=71) "// This panics, will log the message as \"fatal\" and exits the function."
    }),
    (*ast.Comment)(0x14007f9bb60)({
     Slash: (token.Pos) 4235240,
     Text: (string) (len=35) "// Generally, you should not panic."
    })
   }
  }),
  (*ast.CommentGroup)(0x14007f9bba8)({
   List: ([]*ast.Comment) (len=1 cap=1) {
    (*ast.Comment)(0x14007f9bb90)({
     Slash: (token.Pos) 4235361,
     Text: (string) (len=43) "// Tests an exit with a non-zero exit code."
    })
   }
  }),
  (*ast.CommentGroup)(0x14007f9bc20)({
   List: ([]*ast.Comment) (len=4 cap=4) {
    (*ast.Comment)(0x14007f9bbc0)({
     Slash: (token.Pos) 4235425,
     Text: (string) (len=74) "// If you need to exit prematurely without panicking, you can use os.Exit."
    }),
    (*ast.Comment)(0x14007f9bbd8)({
     Slash: (token.Pos) 4235501,
     Text: (string) (len=74) "// However, you cannot return any values from the function, so if you want"
    }),
    (*ast.Comment)(0x14007f9bbf0)({
     Slash: (token.Pos) 4235577,
     Text: (string) (len=68) "// to log an error message, you should do so before calling os.Exit."
    }),
    (*ast.Comment)(0x14007f9bc08)({
     Slash: (token.Pos) 4235647,
     Text: (string) (len=67) "// The exit code should be 0 for success, and non-zero for failure."
    })
   }
  }),
  (*ast.CommentGroup)(0x14007f9bc80)({
   List: ([]*ast.Comment) (len=1 cap=1) {
    (*ast.Comment)(0x14007f9bc68)({
     Slash: (token.Pos) 4235819,
     Text: (string) (len=37) "// Tests logging at different levels."
    })
   }
  }),
  (*ast.CommentGroup)(0x14007f9bcb0)({
   List: ([]*ast.Comment) (len=1 cap=1) {
    (*ast.Comment)(0x14007f9bc98)({
     Slash: (token.Pos) 4235879,
     Text: (string) (len=49) "// This is a simple log message. It has no level."
    })
   }
  }),
  (*ast.CommentGroup)(0x14007f9bcf8)({
   List: ([]*ast.Comment) (len=1 cap=1) {
    (*ast.Comment)(0x14007f9bce0)({
     Slash: (token.Pos) 4235977,
     Text: (string) (len=49) "// These messages are logged at different levels."
    })
   }
  }),
  (*ast.CommentGroup)(0x14007f9bd70)({
   List: ([]*ast.Comment) (len=1 cap=1) {
    (*ast.Comment)(0x14007f9bd58)({
     Slash: (token.Pos) 4236158,
     Text: (string) (len=67) "// This logs an error message, but allows the function to continue."
    })
   }
  }),
  (*ast.CommentGroup)(0x14007f9bdd0)({
   List: ([]*ast.Comment) (len=1 cap=1) {
    (*ast.Comment)(0x14007f9bdb8)({
     Slash: (token.Pos) 4236438,
     Text: (string) (len=52) "// You can also use Go's built-in printing commands."
    })
   }
  }),
  (*ast.CommentGroup)(0x14007f9be30)({
   List: ([]*ast.Comment) (len=1 cap=1) {
    (*ast.Comment)(0x14007f9be18)({
     Slash: (token.Pos) 4236651,
     Text: (string) (len=96) "// You can even just use stdout/stderr and the log level will be \"info\" or \"error\" respectively."
    })
   }
  }),
  (*ast.CommentGroup)(0x14007f9bf68)({
   List: ([]*ast.Comment) (len=8 cap=8) {
    (*ast.Comment)(0x14007f9bea8)({
     Slash: (token.Pos) 4236893,
     Text: (string) (len=121) "// NOTE: The main difference between using console functions and Go's built-in functions, is in how newlines are handled."
    }),
    (*ast.Comment)(0x14007f9bec0)({
     Slash: (token.Pos) 4237016,
     Text: (string) (len=106) "// - Using console functions, newlines are preserved and the entire message is logged as a single message."
    }),
    (*ast.Comment)(0x14007f9bed8)({
     Slash: (token.Pos) 4237124,
     Text: (string) (len=78) "// - Using Go's built-in functions, each line is logged as a separate message."
    }),
    (*ast.Comment)(0x14007f9bef0)({
     Slash: (token.Pos) 4237204,
     Text: (string) (len=2) "//"
    }),
    (*ast.Comment)(0x14007f9bf08)({
     Slash: (token.Pos) 4237208,
     Text: (string) (len=95) "// Thus, if you are logging data for debugging, we highly recommend using the console functions"
    }),
    (*ast.Comment)(0x14007f9bf20)({
     Slash: (token.Pos) 4237305,
     Text: (string) (len=53) "// to keep the data together in a single log message."
    }),
    (*ast.Comment)(0x14007f9bf38)({
     Slash: (token.Pos) 4237360,
     Text: (string) (len=2) "//"
    }),
    (*ast.Comment)(0x14007f9bf50)({
     Slash: (token.Pos) 4237364,
     Text: (string) (len=85) "// The console functions also allow you to better control the reported logging level."
    })
   }
  })
 },
 GoVersion: (string) ""
})

GML: getFunctionsNeedingWrappers: f.Imports[0]=
(*ast.ImportSpec)(0x14007fdd890)({
 Doc: (*ast.CommentGroup)(<nil>),
 Name: (*ast.Ident)(<nil>),
 Path: (*ast.BasicLit)(0x14007fa3a20)({
  ValuePos: (token.Pos) 4231951,
  Kind: (token.Token) STRING,
  Value: (string) (len=8) "\"errors\""
 }),
 Comment: (*ast.CommentGroup)(<nil>),
 EndPos: (token.Pos) 0
})

GML: getFunctionsNeedingWrappers: imports["errors"] = "errors"

GML: getFunctionsNeedingWrappers: f.Imports[1]=
(*ast.ImportSpec)(0x14007fdd8c0)({
 Doc: (*ast.CommentGroup)(<nil>),
 Name: (*ast.Ident)(<nil>),
 Path: (*ast.BasicLit)(0x14007fa3a40)({
  ValuePos: (token.Pos) 4231961,
  Kind: (token.Token) STRING,
  Value: (string) (len=5) "\"fmt\""
 }),
 Comment: (*ast.CommentGroup)(<nil>),
 EndPos: (token.Pos) 0
})

GML: getFunctionsNeedingWrappers: imports["fmt"] = "fmt"

GML: getFunctionsNeedingWrappers: f.Imports[2]=
(*ast.ImportSpec)(0x14007fdd8f0)({
 Doc: (*ast.CommentGroup)(<nil>),
 Name: (*ast.Ident)(<nil>),
 Path: (*ast.BasicLit)(0x14007fa3a80)({
  ValuePos: (token.Pos) 4231968,
  Kind: (token.Token) STRING,
  Value: (string) (len=11) "\"math/rand\""
 }),
 Comment: (*ast.CommentGroup)(<nil>),
 EndPos: (token.Pos) 0
})

GML: getFunctionsNeedingWrappers: imports["rand"] = "math/rand"

GML: getFunctionsNeedingWrappers: f.Imports[3]=
(*ast.ImportSpec)(0x14007fdd920)({
 Doc: (*ast.CommentGroup)(<nil>),
 Name: (*ast.Ident)(<nil>),
 Path: (*ast.BasicLit)(0x14007fa3ac0)({
  ValuePos: (token.Pos) 4231981,
  Kind: (token.Token) STRING,
  Value: (string) (len=4) "\"os\""
 }),
 Comment: (*ast.CommentGroup)(<nil>),
 EndPos: (token.Pos) 0
})

GML: getFunctionsNeedingWrappers: imports["os"] = "os"

GML: getFunctionsNeedingWrappers: f.Imports[4]=
(*ast.ImportSpec)(0x14007fdd950)({
 Doc: (*ast.CommentGroup)(<nil>),
 Name: (*ast.Ident)(<nil>),
 Path: (*ast.BasicLit)(0x14007fa3ae0)({
  ValuePos: (token.Pos) 4231987,
  Kind: (token.Token) STRING,
  Value: (string) (len=6) "\"time\""
 }),
 Comment: (*ast.CommentGroup)(<nil>),
 EndPos: (token.Pos) 0
})

GML: getFunctionsNeedingWrappers: imports["time"] = "time"

GML: getFunctionsNeedingWrappers: f.Imports[5]=
(*ast.ImportSpec)(0x14007fdd980)({
 Doc: (*ast.CommentGroup)(<nil>),
 Name: (*ast.Ident)(<nil>),
 Path: (*ast.BasicLit)(0x14007fa3b00)({
  ValuePos: (token.Pos) 4231996,
  Kind: (token.Token) STRING,
  Value: (string) (len=50) "\"github.com/hypermodeinc/modus/sdk/go/pkg/console\""
 }),
 Comment: (*ast.CommentGroup)(<nil>),
 EndPos: (token.Pos) 0
})

GML: getFunctionsNeedingWrappers: imports["console"] = "github.com/hypermodeinc/modus/sdk/go/pkg/console"

GML: getFunctionsNeedingWrappers: f.Decls[0]=
(*ast.GenDecl)(0x14007fe7640)({
 Doc: (*ast.CommentGroup)(<nil>),
 TokPos: (token.Pos) 4231941,
 Tok: (token.Token) import,
 Lparen: (token.Pos) 4231948,
 Specs: ([]ast.Spec) (len=6 cap=8) {
  (*ast.ImportSpec)(0x14007fdd890)({
   Doc: (*ast.CommentGroup)(<nil>),
   Name: (*ast.Ident)(<nil>),
   Path: (*ast.BasicLit)(0x14007fa3a20)({
    ValuePos: (token.Pos) 4231951,
    Kind: (token.Token) STRING,
    Value: (string) (len=8) "\"errors\""
   }),
   Comment: (*ast.CommentGroup)(<nil>),
   EndPos: (token.Pos) 0
  }),
  (*ast.ImportSpec)(0x14007fdd8c0)({
   Doc: (*ast.CommentGroup)(<nil>),
   Name: (*ast.Ident)(<nil>),
   Path: (*ast.BasicLit)(0x14007fa3a40)({
    ValuePos: (token.Pos) 4231961,
    Kind: (token.Token) STRING,
    Value: (string) (len=5) "\"fmt\""
   }),
   Comment: (*ast.CommentGroup)(<nil>),
   EndPos: (token.Pos) 0
  }),
  (*ast.ImportSpec)(0x14007fdd8f0)({
   Doc: (*ast.CommentGroup)(<nil>),
   Name: (*ast.Ident)(<nil>),
   Path: (*ast.BasicLit)(0x14007fa3a80)({
    ValuePos: (token.Pos) 4231968,
    Kind: (token.Token) STRING,
    Value: (string) (len=11) "\"math/rand\""
   }),
   Comment: (*ast.CommentGroup)(<nil>),
   EndPos: (token.Pos) 0
  }),
  (*ast.ImportSpec)(0x14007fdd920)({
   Doc: (*ast.CommentGroup)(<nil>),
   Name: (*ast.Ident)(<nil>),
   Path: (*ast.BasicLit)(0x14007fa3ac0)({
    ValuePos: (token.Pos) 4231981,
    Kind: (token.Token) STRING,
    Value: (string) (len=4) "\"os\""
   }),
   Comment: (*ast.CommentGroup)(<nil>),
   EndPos: (token.Pos) 0
  }),
  (*ast.ImportSpec)(0x14007fdd950)({
   Doc: (*ast.CommentGroup)(<nil>),
   Name: (*ast.Ident)(<nil>),
   Path: (*ast.BasicLit)(0x14007fa3ae0)({
    ValuePos: (token.Pos) 4231987,
    Kind: (token.Token) STRING,
    Value: (string) (len=6) "\"time\""
   }),
   Comment: (*ast.CommentGroup)(<nil>),
   EndPos: (token.Pos) 0
  }),
  (*ast.ImportSpec)(0x14007fdd980)({
   Doc: (*ast.CommentGroup)(<nil>),
   Name: (*ast.Ident)(<nil>),
   Path: (*ast.BasicLit)(0x14007fa3b00)({
    ValuePos: (token.Pos) 4231996,
    Kind: (token.Token) STRING,
    Value: (string) (len=50) "\"github.com/hypermodeinc/modus/sdk/go/pkg/console\""
   }),
   Comment: (*ast.CommentGroup)(<nil>),
   EndPos: (token.Pos) 0
  })
 },
 Rparen: (token.Pos) 4232047
})

GML: getFunctionsNeedingWrappers: f.Decls[1]=
(*ast.FuncDecl)(0x14007fdda10)({
 Doc: (*ast.CommentGroup)(0x14007f9b410)({
  List: ([]*ast.Comment) (len=1 cap=1) {
   (*ast.Comment)(0x14007f9b3f8)({
    Slash: (token.Pos) 4232050,
    Text: (string) (len=18) "// Logs a message."
   })
  }
 }),
 Recv: (*ast.FieldList)(<nil>),
 Name: (*ast.Ident)(0x14007fa3b20)(LogMessage),
 Type: (*ast.FuncType)(0x14007fa3c00)({
  Func: (token.Pos) 4232069,
  TypeParams: (*ast.FieldList)(<nil>),
  Params: (*ast.FieldList)(0x14007fdd9b0)({
   Opening: (token.Pos) 4232084,
   List: ([]*ast.Field) (len=1 cap=1) {
    (*ast.Field)(0x14007fe7680)({
     Doc: (*ast.CommentGroup)(<nil>),
     Names: ([]*ast.Ident) (len=1 cap=1) {
      (*ast.Ident)(0x14007fa3b40)(message)
     },
     Type: (*ast.Ident)(0x14007fa3b60)(string),
     Tag: (*ast.BasicLit)(<nil>),
     Comment: (*ast.CommentGroup)(<nil>)
    })
   },
   Closing: (token.Pos) 4232099
  }),
  Results: (*ast.FieldList)(<nil>)
 }),
 Body: (*ast.BlockStmt)(0x14007fdd9e0)({
  Lbrace: (token.Pos) 4232101,
  List: ([]ast.Stmt) (len=1 cap=1) {
   (*ast.ExprStmt)(0x14007fc7730)({
    X: (*ast.CallExpr)(0x14007fe76c0)({
     Fun: (*ast.SelectorExpr)(0x14007f9b440)({
      X: (*ast.Ident)(0x14007fa3b80)(console),
      Sel: (*ast.Ident)(0x14007fa3ba0)(Log)
     }),
     Lparen: (token.Pos) 4232115,
     Args: ([]ast.Expr) (len=1 cap=1) {
      (*ast.Ident)(0x14007fa3bc0)(message)
     },
     Ellipsis: (token.Pos) 0,
     Rparen: (token.Pos) 4232123
    })
   })
  },
  Rbrace: (token.Pos) 4232125
 })
})

GML: getFunctionsNeedingWrappers: results.append: info=
(*codegen.funcInfo)(0x14008480618)({
 function: (*ast.FuncDecl)(0x14007fdda10)({
  Doc: (*ast.CommentGroup)(0x14007f9b410)({
   List: ([]*ast.Comment) (len=1 cap=1) {
    (*ast.Comment)(0x14007f9b3f8)({
     Slash: (token.Pos) 4232050,
     Text: (string) (len=18) "// Logs a message."
    })
   }
  }),
  Recv: (*ast.FieldList)(<nil>),
  Name: (*ast.Ident)(0x14007fa3b20)(LogMessage),
  Type: (*ast.FuncType)(0x14007fa3c00)({
   Func: (token.Pos) 4232069,
   TypeParams: (*ast.FieldList)(<nil>),
   Params: (*ast.FieldList)(0x14007fdd9b0)({
    Opening: (token.Pos) 4232084,
    List: ([]*ast.Field) (len=1 cap=1) {
     (*ast.Field)(0x14007fe7680)({
      Doc: (*ast.CommentGroup)(<nil>),
      Names: ([]*ast.Ident) (len=1 cap=1) {
       (*ast.Ident)(0x14007fa3b40)(message)
      },
      Type: (*ast.Ident)(0x14007fa3b60)(string),
      Tag: (*ast.BasicLit)(<nil>),
      Comment: (*ast.CommentGroup)(<nil>)
     })
    },
    Closing: (token.Pos) 4232099
   }),
   Results: (*ast.FieldList)(<nil>)
  }),
  Body: (*ast.BlockStmt)(0x14007fdd9e0)({
   Lbrace: (token.Pos) 4232101,
   List: ([]ast.Stmt) (len=1 cap=1) {
    (*ast.ExprStmt)(0x14007fc7730)({
     X: (*ast.CallExpr)(0x14007fe76c0)({
      Fun: (*ast.SelectorExpr)(0x14007f9b440)({
       X: (*ast.Ident)(0x14007fa3b80)(console),
       Sel: (*ast.Ident)(0x14007fa3ba0)(Log)
      }),
      Lparen: (token.Pos) 4232115,
      Args: ([]ast.Expr) (len=1 cap=1) {
       (*ast.Ident)(0x14007fa3bc0)(message)
      },
      Ellipsis: (token.Pos) 0,
      Rparen: (token.Pos) 4232123
     })
    })
   },
   Rbrace: (token.Pos) 4232125
  })
 }),
 imports: (map[string]string) {
 },
 aliases: (map[string]string) <nil>
})

GML: getFunctionsNeedingWrappers: f.Decls[2]=
(*ast.FuncDecl)(0x14007fddb30)({
 Doc: (*ast.CommentGroup)(0x14007f9b470)({
  List: ([]*ast.Comment) (len=1 cap=1) {
   (*ast.Comment)(0x14007f9b458)({
    Slash: (token.Pos) 4232128,
    Text: (string) (len=53) "// Adds two integers together and returns the result."
   })
  }
 }),
 Recv: (*ast.FieldList)(<nil>),
 Name: (*ast.Ident)(0x14007fa3c40)(Add),
 Type: (*ast.FuncType)(0x14007fa3d40)({
  Func: (token.Pos) 4232182,
  TypeParams: (*ast.FieldList)(<nil>),
  Params: (*ast.FieldList)(0x14007fdda70)({
   Opening: (token.Pos) 4232190,
   List: ([]*ast.Field) (len=1 cap=1) {
    (*ast.Field)(0x14007fe7700)({
     Doc: (*ast.CommentGroup)(<nil>),
     Names: ([]*ast.Ident) (len=2 cap=2) {
      (*ast.Ident)(0x14007fa3c60)(x),
      (*ast.Ident)(0x14007fa3c80)(y)
     },
     Type: (*ast.Ident)(0x14007fa3ca0)(int),
     Tag: (*ast.BasicLit)(<nil>),
     Comment: (*ast.CommentGroup)(<nil>)
    })
   },
   Closing: (token.Pos) 4232199
  }),
  Results: (*ast.FieldList)(0x14007fddaa0)({
   Opening: (token.Pos) 0,
   List: ([]*ast.Field) (len=1 cap=1) {
    (*ast.Field)(0x14007fe7740)({
     Doc: (*ast.CommentGroup)(<nil>),
     Names: ([]*ast.Ident) <nil>,
     Type: (*ast.Ident)(0x14007fa3cc0)(int),
     Tag: (*ast.BasicLit)(<nil>),
     Comment: (*ast.CommentGroup)(<nil>)
    })
   },
   Closing: (token.Pos) 0
  })
 }),
 Body: (*ast.BlockStmt)(0x14007fddb00)({
  Lbrace: (token.Pos) 4232205,
  List: ([]ast.Stmt) (len=1 cap=1) {
   (*ast.ReturnStmt)(0x14007fa3d20)({
    Return: (token.Pos) 4232208,
    Results: ([]ast.Expr) (len=1 cap=1) {
     (*ast.BinaryExpr)(0x14007fddad0)({
      X: (*ast.Ident)(0x14007fa3ce0)(x),
      OpPos: (token.Pos) 4232217,
      Op: (token.Token) +,
      Y: (*ast.Ident)(0x14007fa3d00)(y)
     })
    }
   })
  },
  Rbrace: (token.Pos) 4232221
 })
})

GML: getFunctionsNeedingWrappers: results.append: info=
(*codegen.funcInfo)(0x14008480690)({
 function: (*ast.FuncDecl)(0x14007fddb30)({
  Doc: (*ast.CommentGroup)(0x14007f9b470)({
   List: ([]*ast.Comment) (len=1 cap=1) {
    (*ast.Comment)(0x14007f9b458)({
     Slash: (token.Pos) 4232128,
     Text: (string) (len=53) "// Adds two integers together and returns the result."
    })
   }
  }),
  Recv: (*ast.FieldList)(<nil>),
  Name: (*ast.Ident)(0x14007fa3c40)(Add),
  Type: (*ast.FuncType)(0x14007fa3d40)({
   Func: (token.Pos) 4232182,
   TypeParams: (*ast.FieldList)(<nil>),
   Params: (*ast.FieldList)(0x14007fdda70)({
    Opening: (token.Pos) 4232190,
    List: ([]*ast.Field) (len=1 cap=1) {
     (*ast.Field)(0x14007fe7700)({
      Doc: (*ast.CommentGroup)(<nil>),
      Names: ([]*ast.Ident) (len=2 cap=2) {
       (*ast.Ident)(0x14007fa3c60)(x),
       (*ast.Ident)(0x14007fa3c80)(y)
      },
      Type: (*ast.Ident)(0x14007fa3ca0)(int),
      Tag: (*ast.BasicLit)(<nil>),
      Comment: (*ast.CommentGroup)(<nil>)
     })
    },
    Closing: (token.Pos) 4232199
   }),
   Results: (*ast.FieldList)(0x14007fddaa0)({
    Opening: (token.Pos) 0,
    List: ([]*ast.Field) (len=1 cap=1) {
     (*ast.Field)(0x14007fe7740)({
      Doc: (*ast.CommentGroup)(<nil>),
      Names: ([]*ast.Ident) <nil>,
      Type: (*ast.Ident)(0x14007fa3cc0)(int),
      Tag: (*ast.BasicLit)(<nil>),
      Comment: (*ast.CommentGroup)(<nil>)
     })
    },
    Closing: (token.Pos) 0
   })
  }),
  Body: (*ast.BlockStmt)(0x14007fddb00)({
   Lbrace: (token.Pos) 4232205,
   List: ([]ast.Stmt) (len=1 cap=1) {
    (*ast.ReturnStmt)(0x14007fa3d20)({
     Return: (token.Pos) 4232208,
     Results: ([]ast.Expr) (len=1 cap=1) {
      (*ast.BinaryExpr)(0x14007fddad0)({
       X: (*ast.Ident)(0x14007fa3ce0)(x),
       OpPos: (token.Pos) 4232217,
       Op: (token.Token) +,
       Y: (*ast.Ident)(0x14007fa3d00)(y)
      })
     }
    })
   },
   Rbrace: (token.Pos) 4232221
  })
 }),
 imports: (map[string]string) {
 },
 aliases: (map[string]string) <nil>
})

GML: getFunctionsNeedingWrappers: f.Decls[3]=
(*ast.FuncDecl)(0x14007fddd10)({
 Doc: (*ast.CommentGroup)(0x14007f9b4d0)({
  List: ([]*ast.Comment) (len=2 cap=2) {
   (*ast.Comment)(0x14007f9b4a0)({
    Slash: (token.Pos) 4232224,
    Text: (string) (len=55) "// Adds three integers together and returns the result."
   }),
   (*ast.Comment)(0x14007f9b4b8)({
    Slash: (token.Pos) 4232280,
    Text: (string) (len=33) "// The third integer is optional."
   })
  }
 }),
 Recv: (*ast.FieldList)(<nil>),
 Name: (*ast.Ident)(0x14007fa3d60)(Add3),
 Type: (*ast.FuncType)(0x14007ffe000)({
  Func: (token.Pos) 4232314,
  TypeParams: (*ast.FieldList)(<nil>),
  Params: (*ast.FieldList)(0x14007fddb90)({
   Opening: (token.Pos) 4232323,
   List: ([]*ast.Field) (len=2 cap=2) {
    (*ast.Field)(0x14007fe77c0)({
     Doc: (*ast.CommentGroup)(<nil>),
     Names: ([]*ast.Ident) (len=2 cap=2) {
      (*ast.Ident)(0x14007fa3d80)(a),
      (*ast.Ident)(0x14007fa3da0)(b)
     },
     Type: (*ast.Ident)(0x14007fa3dc0)(int),
     Tag: (*ast.BasicLit)(<nil>),
     Comment: (*ast.CommentGroup)(<nil>)
    }),
    (*ast.Field)(0x14007fe7800)({
     Doc: (*ast.CommentGroup)(<nil>),
     Names: ([]*ast.Ident) (len=1 cap=1) {
      (*ast.Ident)(0x14007fa3de0)(c)
     },
     Type: (*ast.StarExpr)(0x14007f9b500)({
      Star: (token.Pos) 4232336,
      X: (*ast.Ident)(0x14007fa3e00)(int)
     }),
     Tag: (*ast.BasicLit)(<nil>),
     Comment: (*ast.CommentGroup)(<nil>)
    })
   },
   Closing: (token.Pos) 4232340
  }),
  Results: (*ast.FieldList)(0x14007fddbc0)({
   Opening: (token.Pos) 0,
   List: ([]*ast.Field) (len=1 cap=1) {
    (*ast.Field)(0x14007fe7840)({
     Doc: (*ast.CommentGroup)(<nil>),
     Names: ([]*ast.Ident) <nil>,
     Type: (*ast.Ident)(0x14007fa3e20)(int),
     Tag: (*ast.BasicLit)(<nil>),
     Comment: (*ast.CommentGroup)(<nil>)
    })
   },
   Closing: (token.Pos) 0
  })
 }),
 Body: (*ast.BlockStmt)(0x14007fddce0)({
  Lbrace: (token.Pos) 4232346,
  List: ([]ast.Stmt) (len=2 cap=2) {
   (*ast.IfStmt)(0x14007fe7880)({
    If: (token.Pos) 4232349,
    Init: (ast.Stmt) <nil>,
    Cond: (*ast.BinaryExpr)(0x14007fddbf0)({
     X: (*ast.Ident)(0x14007fa3e40)(c),
     OpPos: (token.Pos) 4232354,
     Op: (token.Token) !=,
     Y: (*ast.Ident)(0x14007fa3e60)(nil)
    }),
    Body: (*ast.BlockStmt)(0x14007fddc80)({
     Lbrace: (token.Pos) 4232361,
     List: ([]ast.Stmt) (len=1 cap=1) {
      (*ast.ReturnStmt)(0x14007fa3ee0)({
       Return: (token.Pos) 4232365,
       Results: ([]ast.Expr) (len=1 cap=1) {
        (*ast.BinaryExpr)(0x14007fddc50)({
         X: (*ast.BinaryExpr)(0x14007fddc20)({
          X: (*ast.Ident)(0x14007fa3e80)(a),
          OpPos: (token.Pos) 4232374,
          Op: (token.Token) +,
          Y: (*ast.Ident)(0x14007fa3ea0)(b)
         }),
         OpPos: (token.Pos) 4232378,
         Op: (token.Token) +,
         Y: (*ast.StarExpr)(0x14007f9b518)({
          Star: (token.Pos) 4232380,
          X: (*ast.Ident)(0x14007fa3ec0)(c)
         })
        })
       }
      })
     },
     Rbrace: (token.Pos) 4232384
    }),
    Else: (ast.Stmt) <nil>
   }),
   (*ast.ReturnStmt)(0x14007fa3f40)({
    Return: (token.Pos) 4232387,
    Results: ([]ast.Expr) (len=1 cap=1) {
     (*ast.BinaryExpr)(0x14007fddcb0)({
      X: (*ast.Ident)(0x14007fa3f00)(a),
      OpPos: (token.Pos) 4232396,
      Op: (token.Token) +,
      Y: (*ast.Ident)(0x14007fa3f20)(b)
     })
    }
   })
  },
  Rbrace: (token.Pos) 4232400
 })
})

GML: getFunctionsNeedingWrappers: results.append: info=
(*codegen.funcInfo)(0x14008480708)({
 function: (*ast.FuncDecl)(0x14007fddd10)({
  Doc: (*ast.CommentGroup)(0x14007f9b4d0)({
   List: ([]*ast.Comment) (len=2 cap=2) {
    (*ast.Comment)(0x14007f9b4a0)({
     Slash: (token.Pos) 4232224,
     Text: (string) (len=55) "// Adds three integers together and returns the result."
    }),
    (*ast.Comment)(0x14007f9b4b8)({
     Slash: (token.Pos) 4232280,
     Text: (string) (len=33) "// The third integer is optional."
    })
   }
  }),
  Recv: (*ast.FieldList)(<nil>),
  Name: (*ast.Ident)(0x14007fa3d60)(Add3),
  Type: (*ast.FuncType)(0x14007ffe000)({
   Func: (token.Pos) 4232314,
   TypeParams: (*ast.FieldList)(<nil>),
   Params: (*ast.FieldList)(0x14007fddb90)({
    Opening: (token.Pos) 4232323,
    List: ([]*ast.Field) (len=2 cap=2) {
     (*ast.Field)(0x14007fe77c0)({
      Doc: (*ast.CommentGroup)(<nil>),
      Names: ([]*ast.Ident) (len=2 cap=2) {
       (*ast.Ident)(0x14007fa3d80)(a),
       (*ast.Ident)(0x14007fa3da0)(b)
      },
      Type: (*ast.Ident)(0x14007fa3dc0)(int),
      Tag: (*ast.BasicLit)(<nil>),
      Comment: (*ast.CommentGroup)(<nil>)
     }),
     (*ast.Field)(0x14007fe7800)({
      Doc: (*ast.CommentGroup)(<nil>),
      Names: ([]*ast.Ident) (len=1 cap=1) {
       (*ast.Ident)(0x14007fa3de0)(c)
      },
      Type: (*ast.StarExpr)(0x14007f9b500)({
       Star: (token.Pos) 4232336,
       X: (*ast.Ident)(0x14007fa3e00)(int)
      }),
      Tag: (*ast.BasicLit)(<nil>),
      Comment: (*ast.CommentGroup)(<nil>)
     })
    },
    Closing: (token.Pos) 4232340
   }),
   Results: (*ast.FieldList)(0x14007fddbc0)({
    Opening: (token.Pos) 0,
    List: ([]*ast.Field) (len=1 cap=1) {
     (*ast.Field)(0x14007fe7840)({
      Doc: (*ast.CommentGroup)(<nil>),
      Names: ([]*ast.Ident) <nil>,
      Type: (*ast.Ident)(0x14007fa3e20)(int),
      Tag: (*ast.BasicLit)(<nil>),
      Comment: (*ast.CommentGroup)(<nil>)
     })
    },
    Closing: (token.Pos) 0
   })
  }),
  Body: (*ast.BlockStmt)(0x14007fddce0)({
   Lbrace: (token.Pos) 4232346,
   List: ([]ast.Stmt) (len=2 cap=2) {
    (*ast.IfStmt)(0x14007fe7880)({
     If: (token.Pos) 4232349,
     Init: (ast.Stmt) <nil>,
     Cond: (*ast.BinaryExpr)(0x14007fddbf0)({
      X: (*ast.Ident)(0x14007fa3e40)(c),
      OpPos: (token.Pos) 4232354,
      Op: (token.Token) !=,
      Y: (*ast.Ident)(0x14007fa3e60)(nil)
     }),
     Body: (*ast.BlockStmt)(0x14007fddc80)({
      Lbrace: (token.Pos) 4232361,
      List: ([]ast.Stmt) (len=1 cap=1) {
       (*ast.ReturnStmt)(0x14007fa3ee0)({
        Return: (token.Pos) 4232365,
        Results: ([]ast.Expr) (len=1 cap=1) {
         (*ast.BinaryExpr)(0x14007fddc50)({
          X: (*ast.BinaryExpr)(0x14007fddc20)({
           X: (*ast.Ident)(0x14007fa3e80)(a),
           OpPos: (token.Pos) 4232374,
           Op: (token.Token) +,
           Y: (*ast.Ident)(0x14007fa3ea0)(b)
          }),
          OpPos: (token.Pos) 4232378,
          Op: (token.Token) +,
          Y: (*ast.StarExpr)(0x14007f9b518)({
           Star: (token.Pos) 4232380,
           X: (*ast.Ident)(0x14007fa3ec0)(c)
          })
         })
        }
       })
      },
      Rbrace: (token.Pos) 4232384
     }),
     Else: (ast.Stmt) <nil>
    }),
    (*ast.ReturnStmt)(0x14007fa3f40)({
     Return: (token.Pos) 4232387,
     Results: ([]ast.Expr) (len=1 cap=1) {
      (*ast.BinaryExpr)(0x14007fddcb0)({
       X: (*ast.Ident)(0x14007fa3f00)(a),
       OpPos: (token.Pos) 4232396,
       Op: (token.Token) +,
       Y: (*ast.Ident)(0x14007fa3f20)(b)
      })
     }
    })
   },
   Rbrace: (token.Pos) 4232400
  })
 }),
 imports: (map[string]string) {
 },
 aliases: (map[string]string) <nil>
})

GML: getFunctionsNeedingWrappers: f.Decls[4]=
(*ast.FuncDecl)(0x14007fdde00)({
 Doc: (*ast.CommentGroup)(0x14007f9b548)({
  List: ([]*ast.Comment) (len=1 cap=1) {
   (*ast.Comment)(0x14007f9b530)({
    Slash: (token.Pos) 4232403,
    Text: (string) (len=63) "// Adds any number of integers together and returns the result."
   })
  }
 }),
 Recv: (*ast.FieldList)(<nil>),
 Name: (*ast.Ident)(0x14007ffe020)(AddN),
 Type: (*ast.FuncType)(0x14007ffe220)({
  Func: (token.Pos) 4232467,
  TypeParams: (*ast.FieldList)(<nil>),
  Params: (*ast.FieldList)(0x14007fddd40)({
   Opening: (token.Pos) 4232476,
   List: ([]*ast.Field) (len=1 cap=1) {
    (*ast.Field)(0x14007fe7900)({
     Doc: (*ast.CommentGroup)(<nil>),
     Names: ([]*ast.Ident) (len=1 cap=1) {
      (*ast.Ident)(0x14007ffe040)(args)
     },
     Type: (*ast.Ellipsis)(0x14007f9b560)({
      Ellipsis: (token.Pos) 4232482,
      Elt: (*ast.Ident)(0x14007ffe060)(int)
     }),
     Tag: (*ast.BasicLit)(<nil>),
     Comment: (*ast.CommentGroup)(<nil>)
    })
   },
   Closing: (token.Pos) 4232488
  }),
  Results: (*ast.FieldList)(0x14007fddd70)({
   Opening: (token.Pos) 0,
   List: ([]*ast.Field) (len=1 cap=1) {
    (*ast.Field)(0x14007fe7940)({
     Doc: (*ast.CommentGroup)(<nil>),
     Names: ([]*ast.Ident) <nil>,
     Type: (*ast.Ident)(0x14007ffe080)(int),
     Tag: (*ast.BasicLit)(<nil>),
     Comment: (*ast.CommentGroup)(<nil>)
    })
   },
   Closing: (token.Pos) 0
  })
 }),
 Body: (*ast.BlockStmt)(0x14007fdddd0)({
  Lbrace: (token.Pos) 4232494,
  List: ([]ast.Stmt) (len=3 cap=4) {
   (*ast.AssignStmt)(0x14007fe7980)({
    Lhs: ([]ast.Expr) (len=1 cap=1) {
     (*ast.Ident)(0x14007ffe0a0)(sum)
    },
    TokPos: (token.Pos) 4232501,
    Tok: (token.Token) :=,
    Rhs: ([]ast.Expr) (len=1 cap=1) {
     (*ast.BasicLit)(0x14007ffe0c0)({
      ValuePos: (token.Pos) 4232504,
      Kind: (token.Token) INT,
      Value: (string) (len=1) "0"
     })
    }
   }),
   (*ast.RangeStmt)(0x14007fe5aa0)({
    For: (token.Pos) 4232507,
    Key: (*ast.Ident)(0x14007ffe0e0)(_),
    Value: (*ast.Ident)(0x14007ffe100)(arg),
    TokPos: (token.Pos) 4232518,
    Tok: (token.Token) :=,
    Range: (token.Pos) 4232521,
    X: (*ast.Ident)(0x14007ffe140)(args),
    Body: (*ast.BlockStmt)(0x14007fddda0)({
     Lbrace: (token.Pos) 4232532,
     List: ([]ast.Stmt) (len=1 cap=1) {
      (*ast.AssignStmt)(0x14007fe7a00)({
       Lhs: ([]ast.Expr) (len=1 cap=1) {
        (*ast.Ident)(0x14007ffe180)(sum)
       },
       TokPos: (token.Pos) 4232540,
       Tok: (token.Token) +=,
       Rhs: ([]ast.Expr) (len=1 cap=1) {
        (*ast.Ident)(0x14007ffe1a0)(arg)
       }
      })
     },
     Rbrace: (token.Pos) 4232548
    })
   }),
   (*ast.ReturnStmt)(0x14007ffe200)({
    Return: (token.Pos) 4232551,
    Results: ([]ast.Expr) (len=1 cap=1) {
     (*ast.Ident)(0x14007ffe1e0)(sum)
    }
   })
  },
  Rbrace: (token.Pos) 4232562
 })
})

GML: getFunctionsNeedingWrappers: results.append: info=
(*codegen.funcInfo)(0x14008480798)({
 function: (*ast.FuncDecl)(0x14007fdde00)({
  Doc: (*ast.CommentGroup)(0x14007f9b548)({
   List: ([]*ast.Comment) (len=1 cap=1) {
    (*ast.Comment)(0x14007f9b530)({
     Slash: (token.Pos) 4232403,
     Text: (string) (len=63) "// Adds any number of integers together and returns the result."
    })
   }
  }),
  Recv: (*ast.FieldList)(<nil>),
  Name: (*ast.Ident)(0x14007ffe020)(AddN),
  Type: (*ast.FuncType)(0x14007ffe220)({
   Func: (token.Pos) 4232467,
   TypeParams: (*ast.FieldList)(<nil>),
   Params: (*ast.FieldList)(0x14007fddd40)({
    Opening: (token.Pos) 4232476,
    List: ([]*ast.Field) (len=1 cap=1) {
     (*ast.Field)(0x14007fe7900)({
      Doc: (*ast.CommentGroup)(<nil>),
      Names: ([]*ast.Ident) (len=1 cap=1) {
       (*ast.Ident)(0x14007ffe040)(args)
      },
      Type: (*ast.Ellipsis)(0x14007f9b560)({
       Ellipsis: (token.Pos) 4232482,
       Elt: (*ast.Ident)(0x14007ffe060)(int)
      }),
      Tag: (*ast.BasicLit)(<nil>),
      Comment: (*ast.CommentGroup)(<nil>)
     })
    },
    Closing: (token.Pos) 4232488
   }),
   Results: (*ast.FieldList)(0x14007fddd70)({
    Opening: (token.Pos) 0,
    List: ([]*ast.Field) (len=1 cap=1) {
     (*ast.Field)(0x14007fe7940)({
      Doc: (*ast.CommentGroup)(<nil>),
      Names: ([]*ast.Ident) <nil>,
      Type: (*ast.Ident)(0x14007ffe080)(int),
      Tag: (*ast.BasicLit)(<nil>),
      Comment: (*ast.CommentGroup)(<nil>)
     })
    },
    Closing: (token.Pos) 0
   })
  }),
  Body: (*ast.BlockStmt)(0x14007fdddd0)({
   Lbrace: (token.Pos) 4232494,
   List: ([]ast.Stmt) (len=3 cap=4) {
    (*ast.AssignStmt)(0x14007fe7980)({
     Lhs: ([]ast.Expr) (len=1 cap=1) {
      (*ast.Ident)(0x14007ffe0a0)(sum)
     },
     TokPos: (token.Pos) 4232501,
     Tok: (token.Token) :=,
     Rhs: ([]ast.Expr) (len=1 cap=1) {
      (*ast.BasicLit)(0x14007ffe0c0)({
       ValuePos: (token.Pos) 4232504,
       Kind: (token.Token) INT,
       Value: (string) (len=1) "0"
      })
     }
    }),
    (*ast.RangeStmt)(0x14007fe5aa0)({
     For: (token.Pos) 4232507,
     Key: (*ast.Ident)(0x14007ffe0e0)(_),
     Value: (*ast.Ident)(0x14007ffe100)(arg),
     TokPos: (token.Pos) 4232518,
     Tok: (token.Token) :=,
     Range: (token.Pos) 4232521,
     X: (*ast.Ident)(0x14007ffe140)(args),
     Body: (*ast.BlockStmt)(0x14007fddda0)({
      Lbrace: (token.Pos) 4232532,
      List: ([]ast.Stmt) (len=1 cap=1) {
       (*ast.AssignStmt)(0x14007fe7a00)({
        Lhs: ([]ast.Expr) (len=1 cap=1) {
         (*ast.Ident)(0x14007ffe180)(sum)
        },
        TokPos: (token.Pos) 4232540,
        Tok: (token.Token) +=,
        Rhs: ([]ast.Expr) (len=1 cap=1) {
         (*ast.Ident)(0x14007ffe1a0)(arg)
        }
       })
      },
      Rbrace: (token.Pos) 4232548
     })
    }),
    (*ast.ReturnStmt)(0x14007ffe200)({
     Return: (token.Pos) 4232551,
     Results: ([]ast.Expr) (len=1 cap=1) {
      (*ast.Ident)(0x14007ffe1e0)(sum)
     }
    })
   },
   Rbrace: (token.Pos) 4232562
  })
 }),
 imports: (map[string]string) {
 },
 aliases: (map[string]string) <nil>
})

GML: getFunctionsNeedingWrappers: f.Decls[5]=
(*ast.GenDecl)(0x14007fe7a80)({
 Doc: (*ast.CommentGroup)(0x14007f9b5a8)({
  List: ([]*ast.Comment) (len=1 cap=1) {
   (*ast.Comment)(0x14007f9b590)({
    Slash: (token.Pos) 4232565,
    Text: (string) (len=55) "// this indirection is so we can mock time.Now in tests"
   })
  }
 }),
 TokPos: (token.Pos) 4232621,
 Tok: (token.Token) var,
 Lparen: (token.Pos) 0,
 Specs: ([]ast.Spec) (len=1 cap=1) {
  (*ast.ValueSpec)(0x14007fcb130)({
   Doc: (*ast.CommentGroup)(<nil>),
   Names: ([]*ast.Ident) (len=1 cap=1) {
    (*ast.Ident)(0x14007ffe240)(nowFunc)
   },
   Type: (ast.Expr) <nil>,
   Values: ([]ast.Expr) (len=1 cap=1) {
    (*ast.SelectorExpr)(0x14007f9b5c0)({
     X: (*ast.Ident)(0x14007ffe260)(time),
     Sel: (*ast.Ident)(0x14007ffe280)(Now)
    })
   },
   Comment: (*ast.CommentGroup)(<nil>)
  })
 },
 Rparen: (token.Pos) 0
})

GML: getFunctionsNeedingWrappers: f.Decls[6]=
(*ast.FuncDecl)(0x14007fddec0)({
 Doc: (*ast.CommentGroup)(0x14007f9b5f0)({
  List: ([]*ast.Comment) (len=1 cap=1) {
   (*ast.Comment)(0x14007f9b5d8)({
    Slash: (token.Pos) 4232645,
    Text: (string) (len=28) "// Returns the current time."
   })
  }
 }),
 Recv: (*ast.FieldList)(<nil>),
 Name: (*ast.Ident)(0x14007ffe2a0)(GetCurrentTime),
 Type: (*ast.FuncType)(0x14007ffe340)({
  Func: (token.Pos) 4232674,
  TypeParams: (*ast.FieldList)(<nil>),
  Params: (*ast.FieldList)(0x14007fdde30)({
   Opening: (token.Pos) 4232693,
   List: ([]*ast.Field) <nil>,
   Closing: (token.Pos) 4232694
  }),
  Results: (*ast.FieldList)(0x14007fdde60)({
   Opening: (token.Pos) 0,
   List: ([]*ast.Field) (len=1 cap=1) {
    (*ast.Field)(0x14007fe7ac0)({
     Doc: (*ast.CommentGroup)(<nil>),
     Names: ([]*ast.Ident) <nil>,
     Type: (*ast.SelectorExpr)(0x14007f9b608)({
      X: (*ast.Ident)(0x14007ffe2c0)(time),
      Sel: (*ast.Ident)(0x14007ffe2e0)(Time)
     }),
     Tag: (*ast.BasicLit)(<nil>),
     Comment: (*ast.CommentGroup)(<nil>)
    })
   },
   Closing: (token.Pos) 0
  })
 }),
 Body: (*ast.BlockStmt)(0x14007fdde90)({
  Lbrace: (token.Pos) 4232706,
  List: ([]ast.Stmt) (len=1 cap=1) {
   (*ast.ReturnStmt)(0x14007ffe320)({
    Return: (token.Pos) 4232709,
    Results: ([]ast.Expr) (len=1 cap=1) {
     (*ast.CallExpr)(0x14007fe7b00)({
      Fun: (*ast.Ident)(0x14007ffe300)(nowFunc),
      Lparen: (token.Pos) 4232723,
      Args: ([]ast.Expr) <nil>,
      Ellipsis: (token.Pos) 0,
      Rparen: (token.Pos) 4232724
     })
    }
   })
  },
  Rbrace: (token.Pos) 4232726
 })
})

GML: getFunctionsNeedingWrappers: results.append: info=
(*codegen.funcInfo)(0x14008480840)({
 function: (*ast.FuncDecl)(0x14007fddec0)({
  Doc: (*ast.CommentGroup)(0x14007f9b5f0)({
   List: ([]*ast.Comment) (len=1 cap=1) {
    (*ast.Comment)(0x14007f9b5d8)({
     Slash: (token.Pos) 4232645,
     Text: (string) (len=28) "// Returns the current time."
    })
   }
  }),
  Recv: (*ast.FieldList)(<nil>),
  Name: (*ast.Ident)(0x14007ffe2a0)(GetCurrentTime),
  Type: (*ast.FuncType)(0x14007ffe340)({
   Func: (token.Pos) 4232674,
   TypeParams: (*ast.FieldList)(<nil>),
   Params: (*ast.FieldList)(0x14007fdde30)({
    Opening: (token.Pos) 4232693,
    List: ([]*ast.Field) <nil>,
    Closing: (token.Pos) 4232694
   }),
   Results: (*ast.FieldList)(0x14007fdde60)({
    Opening: (token.Pos) 0,
    List: ([]*ast.Field) (len=1 cap=1) {
     (*ast.Field)(0x14007fe7ac0)({
      Doc: (*ast.CommentGroup)(<nil>),
      Names: ([]*ast.Ident) <nil>,
      Type: (*ast.SelectorExpr)(0x14007f9b608)({
       X: (*ast.Ident)(0x14007ffe2c0)(time),
       Sel: (*ast.Ident)(0x14007ffe2e0)(Time)
      }),
      Tag: (*ast.BasicLit)(<nil>),
      Comment: (*ast.CommentGroup)(<nil>)
     })
    },
    Closing: (token.Pos) 0
   })
  }),
  Body: (*ast.BlockStmt)(0x14007fdde90)({
   Lbrace: (token.Pos) 4232706,
   List: ([]ast.Stmt) (len=1 cap=1) {
    (*ast.ReturnStmt)(0x14007ffe320)({
     Return: (token.Pos) 4232709,
     Results: ([]ast.Expr) (len=1 cap=1) {
      (*ast.CallExpr)(0x14007fe7b00)({
       Fun: (*ast.Ident)(0x14007ffe300)(nowFunc),
       Lparen: (token.Pos) 4232723,
       Args: ([]ast.Expr) <nil>,
       Ellipsis: (token.Pos) 0,
       Rparen: (token.Pos) 4232724
      })
     }
    })
   },
   Rbrace: (token.Pos) 4232726
  })
 }),
 imports: (map[string]string) (len=1) {
  (string) (len=4) "time": (string) (len=4) "time"
 },
 aliases: (map[string]string) <nil>
})

GML: getFunctionsNeedingWrappers: f.Decls[7]=
(*ast.FuncDecl)(0x14008478000)({
 Doc: (*ast.CommentGroup)(0x14007f9b638)({
  List: ([]*ast.Comment) (len=1 cap=1) {
   (*ast.Comment)(0x14007f9b620)({
    Slash: (token.Pos) 4232729,
    Text: (string) (len=50) "// Returns the current time formatted as a string."
   })
  }
 }),
 Recv: (*ast.FieldList)(<nil>),
 Name: (*ast.Ident)(0x14007ffe360)(GetCurrentTimeFormatted),
 Type: (*ast.FuncType)(0x14007ffe440)({
  Func: (token.Pos) 4232780,
  TypeParams: (*ast.FieldList)(<nil>),
  Params: (*ast.FieldList)(0x14007fddef0)({
   Opening: (token.Pos) 4232808,
   List: ([]*ast.Field) <nil>,
   Closing: (token.Pos) 4232809
  }),
  Results: (*ast.FieldList)(0x14007fddf20)({
   Opening: (token.Pos) 0,
   List: ([]*ast.Field) (len=1 cap=1) {
    (*ast.Field)(0x14007fe7b40)({
     Doc: (*ast.CommentGroup)(<nil>),
     Names: ([]*ast.Ident) <nil>,
     Type: (*ast.Ident)(0x14007ffe380)(string),
     Tag: (*ast.BasicLit)(<nil>),
     Comment: (*ast.CommentGroup)(<nil>)
    })
   },
   Closing: (token.Pos) 0
  })
 }),
 Body: (*ast.BlockStmt)(0x14007fddf50)({
  Lbrace: (token.Pos) 4232818,
  List: ([]ast.Stmt) (len=1 cap=1) {
   (*ast.ReturnStmt)(0x14007ffe420)({
    Return: (token.Pos) 4232821,
    Results: ([]ast.Expr) (len=1 cap=1) {
     (*ast.CallExpr)(0x14007fe7bc0)({
      Fun: (*ast.SelectorExpr)(0x14007f9b650)({
       X: (*ast.CallExpr)(0x14007fe7b80)({
        Fun: (*ast.Ident)(0x14007ffe3a0)(nowFunc),
        Lparen: (token.Pos) 4232835,
        Args: ([]ast.Expr) <nil>,
        Ellipsis: (token.Pos) 0,
        Rparen: (token.Pos) 4232836
       }),
       Sel: (*ast.Ident)(0x14007ffe3c0)(Format)
      }),
      Lparen: (token.Pos) 4232844,
      Args: ([]ast.Expr) (len=1 cap=1) {
       (*ast.SelectorExpr)(0x14007f9b668)({
        X: (*ast.Ident)(0x14007ffe3e0)(time),
        Sel: (*ast.Ident)(0x14007ffe400)(DateTime)
       })
      },
      Ellipsis: (token.Pos) 0,
      Rparen: (token.Pos) 4232858
     })
    }
   })
  },
  Rbrace: (token.Pos) 4232860
 })
})

GML: getFunctionsNeedingWrappers: results.append: info=
(*codegen.funcInfo)(0x140084808d0)({
 function: (*ast.FuncDecl)(0x14008478000)({
  Doc: (*ast.CommentGroup)(0x14007f9b638)({
   List: ([]*ast.Comment) (len=1 cap=1) {
    (*ast.Comment)(0x14007f9b620)({
     Slash: (token.Pos) 4232729,
     Text: (string) (len=50) "// Returns the current time formatted as a string."
    })
   }
  }),
  Recv: (*ast.FieldList)(<nil>),
  Name: (*ast.Ident)(0x14007ffe360)(GetCurrentTimeFormatted),
  Type: (*ast.FuncType)(0x14007ffe440)({
   Func: (token.Pos) 4232780,
   TypeParams: (*ast.FieldList)(<nil>),
   Params: (*ast.FieldList)(0x14007fddef0)({
    Opening: (token.Pos) 4232808,
    List: ([]*ast.Field) <nil>,
    Closing: (token.Pos) 4232809
   }),
   Results: (*ast.FieldList)(0x14007fddf20)({
    Opening: (token.Pos) 0,
    List: ([]*ast.Field) (len=1 cap=1) {
     (*ast.Field)(0x14007fe7b40)({
      Doc: (*ast.CommentGroup)(<nil>),
      Names: ([]*ast.Ident) <nil>,
      Type: (*ast.Ident)(0x14007ffe380)(string),
      Tag: (*ast.BasicLit)(<nil>),
      Comment: (*ast.CommentGroup)(<nil>)
     })
    },
    Closing: (token.Pos) 0
   })
  }),
  Body: (*ast.BlockStmt)(0x14007fddf50)({
   Lbrace: (token.Pos) 4232818,
   List: ([]ast.Stmt) (len=1 cap=1) {
    (*ast.ReturnStmt)(0x14007ffe420)({
     Return: (token.Pos) 4232821,
     Results: ([]ast.Expr) (len=1 cap=1) {
      (*ast.CallExpr)(0x14007fe7bc0)({
       Fun: (*ast.SelectorExpr)(0x14007f9b650)({
        X: (*ast.CallExpr)(0x14007fe7b80)({
         Fun: (*ast.Ident)(0x14007ffe3a0)(nowFunc),
         Lparen: (token.Pos) 4232835,
         Args: ([]ast.Expr) <nil>,
         Ellipsis: (token.Pos) 0,
         Rparen: (token.Pos) 4232836
        }),
        Sel: (*ast.Ident)(0x14007ffe3c0)(Format)
       }),
       Lparen: (token.Pos) 4232844,
       Args: ([]ast.Expr) (len=1 cap=1) {
        (*ast.SelectorExpr)(0x14007f9b668)({
         X: (*ast.Ident)(0x14007ffe3e0)(time),
         Sel: (*ast.Ident)(0x14007ffe400)(DateTime)
        })
       },
       Ellipsis: (token.Pos) 0,
       Rparen: (token.Pos) 4232858
      })
     }
    })
   },
   Rbrace: (token.Pos) 4232860
  })
 }),
 imports: (map[string]string) {
 },
 aliases: (map[string]string) <nil>
})

GML: getFunctionsNeedingWrappers: f.Decls[8]=
(*ast.FuncDecl)(0x14008478150)({
 Doc: (*ast.CommentGroup)(0x14007f9b698)({
  List: ([]*ast.Comment) (len=1 cap=1) {
   (*ast.Comment)(0x14007f9b680)({
    Slash: (token.Pos) 4232863,
    Text: (string) (len=75) "// Combines the first and last name of a person, and returns the full name."
   })
  }
 }),
 Recv: (*ast.FieldList)(<nil>),
 Name: (*ast.Ident)(0x14007ffe460)(GetFullName),
 Type: (*ast.FuncType)(0x14007ffe580)({
  Func: (token.Pos) 4232939,
  TypeParams: (*ast.FieldList)(<nil>),
  Params: (*ast.FieldList)(0x14008478060)({
   Opening: (token.Pos) 4232955,
   List: ([]*ast.Field) (len=1 cap=1) {
    (*ast.Field)(0x14007fe7c00)({
     Doc: (*ast.CommentGroup)(<nil>),
     Names: ([]*ast.Ident) (len=2 cap=2) {
      (*ast.Ident)(0x14007ffe480)(firstName),
      (*ast.Ident)(0x14007ffe4a0)(lastName)
     },
     Type: (*ast.Ident)(0x14007ffe4c0)(string),
     Tag: (*ast.BasicLit)(<nil>),
     Comment: (*ast.CommentGroup)(<nil>)
    })
   },
   Closing: (token.Pos) 4232982
  }),
  Results: (*ast.FieldList)(0x14008478090)({
   Opening: (token.Pos) 0,
   List: ([]*ast.Field) (len=1 cap=1) {
    (*ast.Field)(0x14007fe7c40)({
     Doc: (*ast.CommentGroup)(<nil>),
     Names: ([]*ast.Ident) <nil>,
     Type: (*ast.Ident)(0x14007ffe4e0)(string),
     Tag: (*ast.BasicLit)(<nil>),
     Comment: (*ast.CommentGroup)(<nil>)
    })
   },
   Closing: (token.Pos) 0
  })
 }),
 Body: (*ast.BlockStmt)(0x14008478120)({
  Lbrace: (token.Pos) 4232991,
  List: ([]ast.Stmt) (len=1 cap=1) {
   (*ast.ReturnStmt)(0x14007ffe560)({
    Return: (token.Pos) 4232994,
    Results: ([]ast.Expr) (len=1 cap=1) {
     (*ast.BinaryExpr)(0x140084780f0)({
      X: (*ast.BinaryExpr)(0x140084780c0)({
       X: (*ast.Ident)(0x14007ffe500)(firstName),
       OpPos: (token.Pos) 4233011,
       Op: (token.Token) +,
       Y: (*ast.BasicLit)(0x14007ffe520)({
        ValuePos: (token.Pos) 4233013,
        Kind: (token.Token) STRING,
        Value: (string) (len=3) "\" \""
       })
      }),
      OpPos: (token.Pos) 4233017,
      Op: (token.Token) +,
      Y: (*ast.Ident)(0x14007ffe540)(lastName)
     })
    }
   })
  },
  Rbrace: (token.Pos) 4233028
 })
})

GML: getFunctionsNeedingWrappers: results.append: info=
(*codegen.funcInfo)(0x14008480948)({
 function: (*ast.FuncDecl)(0x14008478150)({
  Doc: (*ast.CommentGroup)(0x14007f9b698)({
   List: ([]*ast.Comment) (len=1 cap=1) {
    (*ast.Comment)(0x14007f9b680)({
     Slash: (token.Pos) 4232863,
     Text: (string) (len=75) "// Combines the first and last name of a person, and returns the full name."
    })
   }
  }),
  Recv: (*ast.FieldList)(<nil>),
  Name: (*ast.Ident)(0x14007ffe460)(GetFullName),
  Type: (*ast.FuncType)(0x14007ffe580)({
   Func: (token.Pos) 4232939,
   TypeParams: (*ast.FieldList)(<nil>),
   Params: (*ast.FieldList)(0x14008478060)({
    Opening: (token.Pos) 4232955,
    List: ([]*ast.Field) (len=1 cap=1) {
     (*ast.Field)(0x14007fe7c00)({
      Doc: (*ast.CommentGroup)(<nil>),
      Names: ([]*ast.Ident) (len=2 cap=2) {
       (*ast.Ident)(0x14007ffe480)(firstName),
       (*ast.Ident)(0x14007ffe4a0)(lastName)
      },
      Type: (*ast.Ident)(0x14007ffe4c0)(string),
      Tag: (*ast.BasicLit)(<nil>),
      Comment: (*ast.CommentGroup)(<nil>)
     })
    },
    Closing: (token.Pos) 4232982
   }),
   Results: (*ast.FieldList)(0x14008478090)({
    Opening: (token.Pos) 0,
    List: ([]*ast.Field) (len=1 cap=1) {
     (*ast.Field)(0x14007fe7c40)({
      Doc: (*ast.CommentGroup)(<nil>),
      Names: ([]*ast.Ident) <nil>,
      Type: (*ast.Ident)(0x14007ffe4e0)(string),
      Tag: (*ast.BasicLit)(<nil>),
      Comment: (*ast.CommentGroup)(<nil>)
     })
    },
    Closing: (token.Pos) 0
   })
  }),
  Body: (*ast.BlockStmt)(0x14008478120)({
   Lbrace: (token.Pos) 4232991,
   List: ([]ast.Stmt) (len=1 cap=1) {
    (*ast.ReturnStmt)(0x14007ffe560)({
     Return: (token.Pos) 4232994,
     Results: ([]ast.Expr) (len=1 cap=1) {
      (*ast.BinaryExpr)(0x140084780f0)({
       X: (*ast.BinaryExpr)(0x140084780c0)({
        X: (*ast.Ident)(0x14007ffe500)(firstName),
        OpPos: (token.Pos) 4233011,
        Op: (token.Token) +,
        Y: (*ast.BasicLit)(0x14007ffe520)({
         ValuePos: (token.Pos) 4233013,
         Kind: (token.Token) STRING,
         Value: (string) (len=3) "\" \""
        })
       }),
       OpPos: (token.Pos) 4233017,
       Op: (token.Token) +,
       Y: (*ast.Ident)(0x14007ffe540)(lastName)
      })
     }
    })
   },
   Rbrace: (token.Pos) 4233028
  })
 }),
 imports: (map[string]string) {
 },
 aliases: (map[string]string) <nil>
})

GML: getFunctionsNeedingWrappers: f.Decls[9]=
(*ast.FuncDecl)(0x14008478300)({
 Doc: (*ast.CommentGroup)(0x14007f9b6f8)({
  List: ([]*ast.Comment) (len=2 cap=2) {
   (*ast.Comment)(0x14007f9b6c8)({
    Slash: (token.Pos) 4233031,
    Text: (string) (len=34) "// Says hello to a person by name."
   }),
   (*ast.Comment)(0x14007f9b6e0)({
    Slash: (token.Pos) 4233066,
    Text: (string) (len=65) "// If the name is not provided, it will say hello without a name."
   })
  }
 }),
 Recv: (*ast.FieldList)(<nil>),
 Name: (*ast.Ident)(0x14007ffe5a0)(SayHello),
 Type: (*ast.FuncType)(0x14007ffe720)({
  Func: (token.Pos) 4233132,
  TypeParams: (*ast.FieldList)(<nil>),
  Params: (*ast.FieldList)(0x14008478180)({
   Opening: (token.Pos) 4233145,
   List: ([]*ast.Field) (len=1 cap=1) {
    (*ast.Field)(0x14007fe7c80)({
     Doc: (*ast.CommentGroup)(<nil>),
     Names: ([]*ast.Ident) (len=1 cap=1) {
      (*ast.Ident)(0x14007ffe5c0)(name)
     },
     Type: (*ast.StarExpr)(0x14007f9b710)({
      Star: (token.Pos) 4233151,
      X: (*ast.Ident)(0x14007ffe5e0)(string)
     }),
     Tag: (*ast.BasicLit)(<nil>),
     Comment: (*ast.CommentGroup)(<nil>)
    })
   },
   Closing: (token.Pos) 4233158
  }),
  Results: (*ast.FieldList)(0x140084781b0)({
   Opening: (token.Pos) 0,
   List: ([]*ast.Field) (len=1 cap=1) {
    (*ast.Field)(0x14007fe7cc0)({
     Doc: (*ast.CommentGroup)(<nil>),
     Names: ([]*ast.Ident) <nil>,
     Type: (*ast.Ident)(0x14007ffe600)(string),
     Tag: (*ast.BasicLit)(<nil>),
     Comment: (*ast.CommentGroup)(<nil>)
    })
   },
   Closing: (token.Pos) 0
  })
 }),
 Body: (*ast.BlockStmt)(0x140084782d0)({
  Lbrace: (token.Pos) 4233167,
  List: ([]ast.Stmt) (len=1 cap=1) {
   (*ast.IfStmt)(0x14007fe7d00)({
    If: (token.Pos) 4233170,
    Init: (ast.Stmt) <nil>,
    Cond: (*ast.BinaryExpr)(0x140084781e0)({
     X: (*ast.Ident)(0x14007ffe620)(name),
     OpPos: (token.Pos) 4233178,
     Op: (token.Token) ==,
     Y: (*ast.Ident)(0x14007ffe640)(nil)
    }),
    Body: (*ast.BlockStmt)(0x14008478210)({
     Lbrace: (token.Pos) 4233185,
     List: ([]ast.Stmt) (len=1 cap=1) {
      (*ast.ReturnStmt)(0x14007ffe680)({
       Return: (token.Pos) 4233189,
       Results: ([]ast.Expr) (len=1 cap=1) {
        (*ast.BasicLit)(0x14007ffe660)({
         ValuePos: (token.Pos) 4233196,
         Kind: (token.Token) STRING,
         Value: (string) (len=8) "\"Hello!\""
        })
       }
      })
     },
     Rbrace: (token.Pos) 4233206
    }),
    Else: (*ast.BlockStmt)(0x140084782a0)({
     Lbrace: (token.Pos) 4233213,
     List: ([]ast.Stmt) (len=1 cap=1) {
      (*ast.ReturnStmt)(0x14007ffe700)({
       Return: (token.Pos) 4233217,
       Results: ([]ast.Expr) (len=1 cap=1) {
        (*ast.BinaryExpr)(0x14008478270)({
         X: (*ast.BinaryExpr)(0x14008478240)({
          X: (*ast.BasicLit)(0x14007ffe6a0)({
           ValuePos: (token.Pos) 4233224,
           Kind: (token.Token) STRING,
           Value: (string) (len=9) "\"Hello, \""
          }),
          OpPos: (token.Pos) 4233234,
          Op: (token.Token) +,
          Y: (*ast.StarExpr)(0x14007f9b740)({
           Star: (token.Pos) 4233236,
           X: (*ast.Ident)(0x14007ffe6c0)(name)
          })
         }),
         OpPos: (token.Pos) 4233242,
         Op: (token.Token) +,
         Y: (*ast.BasicLit)(0x14007ffe6e0)({
          ValuePos: (token.Pos) 4233244,
          Kind: (token.Token) STRING,
          Value: (string) (len=3) "\"!\""
         })
        })
       }
      })
     },
     Rbrace: (token.Pos) 4233249
    })
   })
  },
  Rbrace: (token.Pos) 4233251
 })
})

GML: getFunctionsNeedingWrappers: results.append: info=
(*codegen.funcInfo)(0x140084809d8)({
 function: (*ast.FuncDecl)(0x14008478300)({
  Doc: (*ast.CommentGroup)(0x14007f9b6f8)({
   List: ([]*ast.Comment) (len=2 cap=2) {
    (*ast.Comment)(0x14007f9b6c8)({
     Slash: (token.Pos) 4233031,
     Text: (string) (len=34) "// Says hello to a person by name."
    }),
    (*ast.Comment)(0x14007f9b6e0)({
     Slash: (token.Pos) 4233066,
     Text: (string) (len=65) "// If the name is not provided, it will say hello without a name."
    })
   }
  }),
  Recv: (*ast.FieldList)(<nil>),
  Name: (*ast.Ident)(0x14007ffe5a0)(SayHello),
  Type: (*ast.FuncType)(0x14007ffe720)({
   Func: (token.Pos) 4233132,
   TypeParams: (*ast.FieldList)(<nil>),
   Params: (*ast.FieldList)(0x14008478180)({
    Opening: (token.Pos) 4233145,
    List: ([]*ast.Field) (len=1 cap=1) {
     (*ast.Field)(0x14007fe7c80)({
      Doc: (*ast.CommentGroup)(<nil>),
      Names: ([]*ast.Ident) (len=1 cap=1) {
       (*ast.Ident)(0x14007ffe5c0)(name)
      },
      Type: (*ast.StarExpr)(0x14007f9b710)({
       Star: (token.Pos) 4233151,
       X: (*ast.Ident)(0x14007ffe5e0)(string)
      }),
      Tag: (*ast.BasicLit)(<nil>),
      Comment: (*ast.CommentGroup)(<nil>)
     })
    },
    Closing: (token.Pos) 4233158
   }),
   Results: (*ast.FieldList)(0x140084781b0)({
    Opening: (token.Pos) 0,
    List: ([]*ast.Field) (len=1 cap=1) {
     (*ast.Field)(0x14007fe7cc0)({
      Doc: (*ast.CommentGroup)(<nil>),
      Names: ([]*ast.Ident) <nil>,
      Type: (*ast.Ident)(0x14007ffe600)(string),
      Tag: (*ast.BasicLit)(<nil>),
      Comment: (*ast.CommentGroup)(<nil>)
     })
    },
    Closing: (token.Pos) 0
   })
  }),
  Body: (*ast.BlockStmt)(0x140084782d0)({
   Lbrace: (token.Pos) 4233167,
   List: ([]ast.Stmt) (len=1 cap=1) {
    (*ast.IfStmt)(0x14007fe7d00)({
     If: (token.Pos) 4233170,
     Init: (ast.Stmt) <nil>,
     Cond: (*ast.BinaryExpr)(0x140084781e0)({
      X: (*ast.Ident)(0x14007ffe620)(name),
      OpPos: (token.Pos) 4233178,
      Op: (token.Token) ==,
      Y: (*ast.Ident)(0x14007ffe640)(nil)
     }),
     Body: (*ast.BlockStmt)(0x14008478210)({
      Lbrace: (token.Pos) 4233185,
      List: ([]ast.Stmt) (len=1 cap=1) {
       (*ast.ReturnStmt)(0x14007ffe680)({
        Return: (token.Pos) 4233189,
        Results: ([]ast.Expr) (len=1 cap=1) {
         (*ast.BasicLit)(0x14007ffe660)({
          ValuePos: (token.Pos) 4233196,
          Kind: (token.Token) STRING,
          Value: (string) (len=8) "\"Hello!\""
         })
        }
       })
      },
      Rbrace: (token.Pos) 4233206
     }),
     Else: (*ast.BlockStmt)(0x140084782a0)({
      Lbrace: (token.Pos) 4233213,
      List: ([]ast.Stmt) (len=1 cap=1) {
       (*ast.ReturnStmt)(0x14007ffe700)({
        Return: (token.Pos) 4233217,
        Results: ([]ast.Expr) (len=1 cap=1) {
         (*ast.BinaryExpr)(0x14008478270)({
          X: (*ast.BinaryExpr)(0x14008478240)({
           X: (*ast.BasicLit)(0x14007ffe6a0)({
            ValuePos: (token.Pos) 4233224,
            Kind: (token.Token) STRING,
            Value: (string) (len=9) "\"Hello, \""
           }),
           OpPos: (token.Pos) 4233234,
           Op: (token.Token) +,
           Y: (*ast.StarExpr)(0x14007f9b740)({
            Star: (token.Pos) 4233236,
            X: (*ast.Ident)(0x14007ffe6c0)(name)
           })
          }),
          OpPos: (token.Pos) 4233242,
          Op: (token.Token) +,
          Y: (*ast.BasicLit)(0x14007ffe6e0)({
           ValuePos: (token.Pos) 4233244,
           Kind: (token.Token) STRING,
           Value: (string) (len=3) "\"!\""
          })
         })
        }
       })
      },
      Rbrace: (token.Pos) 4233249
     })
    })
   },
   Rbrace: (token.Pos) 4233251
  })
 }),
 imports: (map[string]string) {
 },
 aliases: (map[string]string) <nil>
})

GML: getFunctionsNeedingWrappers: f.Decls[10]=
(*ast.GenDecl)(0x14007fe7e40)({
 Doc: (*ast.CommentGroup)(0x14007f9b770)({
  List: ([]*ast.Comment) (len=1 cap=1) {
   (*ast.Comment)(0x14007f9b758)({
    Slash: (token.Pos) 4233254,
    Text: (string) (len=41) "// A simple object representing a person."
   })
  }
 }),
 TokPos: (token.Pos) 4233296,
 Tok: (token.Token) type,
 Lparen: (token.Pos) 0,
 Specs: ([]ast.Spec) (len=1 cap=1) {
  (*ast.TypeSpec)(0x14007fe7d40)({
   Doc: (*ast.CommentGroup)(<nil>),
   Name: (*ast.Ident)(0x14007ffe740)(Person),
   TypeParams: (*ast.FieldList)(<nil>),
   Assign: (token.Pos) 0,
   Type: (*ast.StructType)(0x14007f9b818)({
    Struct: (token.Pos) 4233308,
    Fields: (*ast.FieldList)(0x14008478330)({
     Opening: (token.Pos) 4233315,
     List: ([]*ast.Field) (len=3 cap=4) {
      (*ast.Field)(0x14007fe7d80)({
       Doc: (*ast.CommentGroup)(0x14007f9b7a0)({
        List: ([]*ast.Comment) (len=1 cap=1) {
         (*ast.Comment)(0x14007f9b788)({
          Slash: (token.Pos) 4233319,
          Text: (string) (len=27) "// The person's first name."
         })
        }
       }),
       Names: ([]*ast.Ident) (len=1 cap=1) {
        (*ast.Ident)(0x14007ffe760)(FirstName)
       },
       Type: (*ast.Ident)(0x14007ffe780)(string),
       Tag: (*ast.BasicLit)(0x14007ffe7a0)({
        ValuePos: (token.Pos) 4233365,
        Kind: (token.Token) STRING,
        Value: (string) (len=18) "`json:\"firstName\"`"
       }),
       Comment: (*ast.CommentGroup)(<nil>)
      }),
      (*ast.Field)(0x14007fe7dc0)({
       Doc: (*ast.CommentGroup)(0x14007f9b7d0)({
        List: ([]*ast.Comment) (len=1 cap=1) {
         (*ast.Comment)(0x14007f9b7b8)({
          Slash: (token.Pos) 4233386,
          Text: (string) (len=26) "// The person's last name."
         })
        }
       }),
       Names: ([]*ast.Ident) (len=1 cap=1) {
        (*ast.Ident)(0x14007ffe7c0)(LastName)
       },
       Type: (*ast.Ident)(0x14007ffe7e0)(string),
       Tag: (*ast.BasicLit)(0x14007ffe800)({
        ValuePos: (token.Pos) 4233430,
        Kind: (token.Token) STRING,
        Value: (string) (len=17) "`json:\"lastName\"`"
       }),
       Comment: (*ast.CommentGroup)(<nil>)
      }),
      (*ast.Field)(0x14007fe7e00)({
       Doc: (*ast.CommentGroup)(0x14007f9b800)({
        List: ([]*ast.Comment) (len=1 cap=1) {
         (*ast.Comment)(0x14007f9b7e8)({
          Slash: (token.Pos) 4233450,
          Text: (string) (len=20) "// The person's age."
         })
        }
       }),
       Names: ([]*ast.Ident) (len=1 cap=1) {
        (*ast.Ident)(0x14007ffe820)(Age)
       },
       Type: (*ast.Ident)(0x14007ffe840)(int),
       Tag: (*ast.BasicLit)(0x14007ffe860)({
        ValuePos: (token.Pos) 4233480,
        Kind: (token.Token) STRING,
        Value: (string) (len=12) "`json:\"age\"`"
       }),
       Comment: (*ast.CommentGroup)(<nil>)
      })
     },
     Closing: (token.Pos) 4233493
    }),
    Incomplete: (bool) false
   }),
   Comment: (*ast.CommentGroup)(<nil>)
  })
 },
 Rparen: (token.Pos) 0
})

GML: getFunctionsNeedingWrappers: f.Decls[11]=
(*ast.FuncDecl)(0x14008478480)({
 Doc: (*ast.CommentGroup)(0x14007f9b848)({
  List: ([]*ast.Comment) (len=1 cap=1) {
   (*ast.Comment)(0x14007f9b830)({
    Slash: (token.Pos) 4233496,
    Text: (string) (len=24) "// Gets a person object."
   })
  }
 }),
 Recv: (*ast.FieldList)(<nil>),
 Name: (*ast.Ident)(0x14007ffe8a0)(GetPerson),
 Type: (*ast.FuncType)(0x14007ffea00)({
  Func: (token.Pos) 4233521,
  TypeParams: (*ast.FieldList)(<nil>),
  Params: (*ast.FieldList)(0x14008478360)({
   Opening: (token.Pos) 4233535,
   List: ([]*ast.Field) <nil>,
   Closing: (token.Pos) 4233536
  }),
  Results: (*ast.FieldList)(0x14008478390)({
   Opening: (token.Pos) 0,
   List: ([]*ast.Field) (len=1 cap=1) {
    (*ast.Field)(0x14007fe7e80)({
     Doc: (*ast.CommentGroup)(<nil>),
     Names: ([]*ast.Ident) <nil>,
     Type: (*ast.Ident)(0x14007ffe8c0)(Person),
     Tag: (*ast.BasicLit)(<nil>),
     Comment: (*ast.CommentGroup)(<nil>)
    })
   },
   Closing: (token.Pos) 0
  })
 }),
 Body: (*ast.BlockStmt)(0x14008478450)({
  Lbrace: (token.Pos) 4233545,
  List: ([]ast.Stmt) (len=1 cap=1) {
   (*ast.ReturnStmt)(0x14007ffe9e0)({
    Return: (token.Pos) 4233548,
    Results: ([]ast.Expr) (len=1 cap=1) {
     (*ast.CompositeLit)(0x14007fe7f00)({
      Type: (*ast.Ident)(0x14007ffe8e0)(Person),
      Lbrace: (token.Pos) 4233561,
      Elts: ([]ast.Expr) (len=3 cap=4) {
       (*ast.KeyValueExpr)(0x140084783c0)({
        Key: (*ast.Ident)(0x14007ffe900)(FirstName),
        Colon: (token.Pos) 4233574,
        Value: (*ast.BasicLit)(0x14007ffe920)({
         ValuePos: (token.Pos) 4233576,
         Kind: (token.Token) STRING,
         Value: (string) (len=6) "\"John\""
        })
       }),
       (*ast.KeyValueExpr)(0x140084783f0)({
        Key: (*ast.Ident)(0x14007ffe940)(LastName),
        Colon: (token.Pos) 4233594,
        Value: (*ast.BasicLit)(0x14007ffe960)({
         ValuePos: (token.Pos) 4233597,
         Kind: (token.Token) STRING,
         Value: (string) (len=5) "\"Doe\""
        })
       }),
       (*ast.KeyValueExpr)(0x14008478420)({
        Key: (*ast.Ident)(0x14007ffe9a0)(Age),
        Colon: (token.Pos) 4233609,
        Value: (*ast.BasicLit)(0x14007ffe9c0)({
         ValuePos: (token.Pos) 4233617,
         Kind: (token.Token) INT,
         Value: (string) (len=2) "42"
        })
       })
      },
      Rbrace: (token.Pos) 4233622,
      Incomplete: (bool) false
     })
    }
   })
  },
  Rbrace: (token.Pos) 4233624
 })
})

GML: getFunctionsNeedingWrappers: results.append: info=
(*codegen.funcInfo)(0x14008480a98)({
 function: (*ast.FuncDecl)(0x14008478480)({
  Doc: (*ast.CommentGroup)(0x14007f9b848)({
   List: ([]*ast.Comment) (len=1 cap=1) {
    (*ast.Comment)(0x14007f9b830)({
     Slash: (token.Pos) 4233496,
     Text: (string) (len=24) "// Gets a person object."
    })
   }
  }),
  Recv: (*ast.FieldList)(<nil>),
  Name: (*ast.Ident)(0x14007ffe8a0)(GetPerson),
  Type: (*ast.FuncType)(0x14007ffea00)({
   Func: (token.Pos) 4233521,
   TypeParams: (*ast.FieldList)(<nil>),
   Params: (*ast.FieldList)(0x14008478360)({
    Opening: (token.Pos) 4233535,
    List: ([]*ast.Field) <nil>,
    Closing: (token.Pos) 4233536
   }),
   Results: (*ast.FieldList)(0x14008478390)({
    Opening: (token.Pos) 0,
    List: ([]*ast.Field) (len=1 cap=1) {
     (*ast.Field)(0x14007fe7e80)({
      Doc: (*ast.CommentGroup)(<nil>),
      Names: ([]*ast.Ident) <nil>,
      Type: (*ast.Ident)(0x14007ffe8c0)(Person),
      Tag: (*ast.BasicLit)(<nil>),
      Comment: (*ast.CommentGroup)(<nil>)
     })
    },
    Closing: (token.Pos) 0
   })
  }),
  Body: (*ast.BlockStmt)(0x14008478450)({
   Lbrace: (token.Pos) 4233545,
   List: ([]ast.Stmt) (len=1 cap=1) {
    (*ast.ReturnStmt)(0x14007ffe9e0)({
     Return: (token.Pos) 4233548,
     Results: ([]ast.Expr) (len=1 cap=1) {
      (*ast.CompositeLit)(0x14007fe7f00)({
       Type: (*ast.Ident)(0x14007ffe8e0)(Person),
       Lbrace: (token.Pos) 4233561,
       Elts: ([]ast.Expr) (len=3 cap=4) {
        (*ast.KeyValueExpr)(0x140084783c0)({
         Key: (*ast.Ident)(0x14007ffe900)(FirstName),
         Colon: (token.Pos) 4233574,
         Value: (*ast.BasicLit)(0x14007ffe920)({
          ValuePos: (token.Pos) 4233576,
          Kind: (token.Token) STRING,
          Value: (string) (len=6) "\"John\""
         })
        }),
        (*ast.KeyValueExpr)(0x140084783f0)({
         Key: (*ast.Ident)(0x14007ffe940)(LastName),
         Colon: (token.Pos) 4233594,
         Value: (*ast.BasicLit)(0x14007ffe960)({
          ValuePos: (token.Pos) 4233597,
          Kind: (token.Token) STRING,
          Value: (string) (len=5) "\"Doe\""
         })
        }),
        (*ast.KeyValueExpr)(0x14008478420)({
         Key: (*ast.Ident)(0x14007ffe9a0)(Age),
         Colon: (token.Pos) 4233609,
         Value: (*ast.BasicLit)(0x14007ffe9c0)({
          ValuePos: (token.Pos) 4233617,
          Kind: (token.Token) INT,
          Value: (string) (len=2) "42"
         })
        })
       },
       Rbrace: (token.Pos) 4233622,
       Incomplete: (bool) false
      })
     }
    })
   },
   Rbrace: (token.Pos) 4233624
  })
 }),
 imports: (map[string]string) {
 },
 aliases: (map[string]string) <nil>
})

GML: getFunctionsNeedingWrappers: f.Decls[12]=
(*ast.FuncDecl)(0x14008478570)({
 Doc: (*ast.CommentGroup)(0x14007f9b878)({
  List: ([]*ast.Comment) (len=1 cap=1) {
   (*ast.Comment)(0x14007f9b860)({
    Slash: (token.Pos) 4233627,
    Text: (string) (len=53) "// Gets a random person object from a list of people."
   })
  }
 }),
 Recv: (*ast.FieldList)(<nil>),
 Name: (*ast.Ident)(0x14007ffea20)(GetRandomPerson),
 Type: (*ast.FuncType)(0x14007ffebc0)({
  Func: (token.Pos) 4233681,
  TypeParams: (*ast.FieldList)(<nil>),
  Params: (*ast.FieldList)(0x140084784b0)({
   Opening: (token.Pos) 4233701,
   List: ([]*ast.Field) <nil>,
   Closing: (token.Pos) 4233702
  }),
  Results: (*ast.FieldList)(0x140084784e0)({
   Opening: (token.Pos) 0,
   List: ([]*ast.Field) (len=1 cap=1) {
    (*ast.Field)(0x14007fe7f40)({
     Doc: (*ast.CommentGroup)(<nil>),
     Names: ([]*ast.Ident) <nil>,
     Type: (*ast.Ident)(0x14007ffea40)(Person),
     Tag: (*ast.BasicLit)(<nil>),
     Comment: (*ast.CommentGroup)(<nil>)
    })
   },
   Closing: (token.Pos) 0
  })
 }),
 Body: (*ast.BlockStmt)(0x14008478540)({
  Lbrace: (token.Pos) 4233711,
  List: ([]ast.Stmt) (len=3 cap=4) {
   (*ast.AssignStmt)(0x1400847a040)({
    Lhs: ([]ast.Expr) (len=1 cap=1) {
     (*ast.Ident)(0x14007ffea60)(people)
    },
    TokPos: (token.Pos) 4233721,
    Tok: (token.Token) :=,
    Rhs: ([]ast.Expr) (len=1 cap=1) {
     (*ast.CallExpr)(0x1400847a000)({
      Fun: (*ast.Ident)(0x14007ffea80)(GetPeople),
      Lparen: (token.Pos) 4233733,
      Args: ([]ast.Expr) <nil>,
      Ellipsis: (token.Pos) 0,
      Rparen: (token.Pos) 4233734
     })
    }
   }),
   (*ast.AssignStmt)(0x1400847a100)({
    Lhs: ([]ast.Expr) (len=1 cap=1) {
     (*ast.Ident)(0x14007ffeaa0)(i)
    },
    TokPos: (token.Pos) 4233739,
    Tok: (token.Token) :=,
    Rhs: ([]ast.Expr) (len=1 cap=1) {
     (*ast.CallExpr)(0x1400847a0c0)({
      Fun: (*ast.SelectorExpr)(0x14007f9b890)({
       X: (*ast.Ident)(0x14007ffeac0)(rand),
       Sel: (*ast.Ident)(0x14007ffeae0)(Intn)
      }),
      Lparen: (token.Pos) 4233751,
      Args: ([]ast.Expr) (len=1 cap=1) {
       (*ast.CallExpr)(0x1400847a080)({
        Fun: (*ast.Ident)(0x14007ffeb00)(len),
        Lparen: (token.Pos) 4233755,
        Args: ([]ast.Expr) (len=1 cap=1) {
         (*ast.Ident)(0x14007ffeb20)(people)
        },
        Ellipsis: (token.Pos) 0,
        Rparen: (token.Pos) 4233762
       })
      },
      Ellipsis: (token.Pos) 0,
      Rparen: (token.Pos) 4233763
     })
    }
   }),
   (*ast.ReturnStmt)(0x14007ffeba0)({
    Return: (token.Pos) 4233766,
    Results: ([]ast.Expr) (len=1 cap=1) {
     (*ast.IndexExpr)(0x14008478510)({
      X: (*ast.Ident)(0x14007ffeb60)(people),
      Lbrack: (token.Pos) 4233779,
      Index: (*ast.Ident)(0x14007ffeb80)(i),
      Rbrack: (token.Pos) 4233781
     })
    }
   })
  },
  Rbrace: (token.Pos) 4233783
 })
})

GML: getFunctionsNeedingWrappers: results.append: info=
(*codegen.funcInfo)(0x14008480b10)({
 function: (*ast.FuncDecl)(0x14008478570)({
  Doc: (*ast.CommentGroup)(0x14007f9b878)({
   List: ([]*ast.Comment) (len=1 cap=1) {
    (*ast.Comment)(0x14007f9b860)({
     Slash: (token.Pos) 4233627,
     Text: (string) (len=53) "// Gets a random person object from a list of people."
    })
   }
  }),
  Recv: (*ast.FieldList)(<nil>),
  Name: (*ast.Ident)(0x14007ffea20)(GetRandomPerson),
  Type: (*ast.FuncType)(0x14007ffebc0)({
   Func: (token.Pos) 4233681,
   TypeParams: (*ast.FieldList)(<nil>),
   Params: (*ast.FieldList)(0x140084784b0)({
    Opening: (token.Pos) 4233701,
    List: ([]*ast.Field) <nil>,
    Closing: (token.Pos) 4233702
   }),
   Results: (*ast.FieldList)(0x140084784e0)({
    Opening: (token.Pos) 0,
    List: ([]*ast.Field) (len=1 cap=1) {
     (*ast.Field)(0x14007fe7f40)({
      Doc: (*ast.CommentGroup)(<nil>),
      Names: ([]*ast.Ident) <nil>,
      Type: (*ast.Ident)(0x14007ffea40)(Person),
      Tag: (*ast.BasicLit)(<nil>),
      Comment: (*ast.CommentGroup)(<nil>)
     })
    },
    Closing: (token.Pos) 0
   })
  }),
  Body: (*ast.BlockStmt)(0x14008478540)({
   Lbrace: (token.Pos) 4233711,
   List: ([]ast.Stmt) (len=3 cap=4) {
    (*ast.AssignStmt)(0x1400847a040)({
     Lhs: ([]ast.Expr) (len=1 cap=1) {
      (*ast.Ident)(0x14007ffea60)(people)
     },
     TokPos: (token.Pos) 4233721,
     Tok: (token.Token) :=,
     Rhs: ([]ast.Expr) (len=1 cap=1) {
      (*ast.CallExpr)(0x1400847a000)({
       Fun: (*ast.Ident)(0x14007ffea80)(GetPeople),
       Lparen: (token.Pos) 4233733,
       Args: ([]ast.Expr) <nil>,
       Ellipsis: (token.Pos) 0,
       Rparen: (token.Pos) 4233734
      })
     }
    }),
    (*ast.AssignStmt)(0x1400847a100)({
     Lhs: ([]ast.Expr) (len=1 cap=1) {
      (*ast.Ident)(0x14007ffeaa0)(i)
     },
     TokPos: (token.Pos) 4233739,
     Tok: (token.Token) :=,
     Rhs: ([]ast.Expr) (len=1 cap=1) {
      (*ast.CallExpr)(0x1400847a0c0)({
       Fun: (*ast.SelectorExpr)(0x14007f9b890)({
        X: (*ast.Ident)(0x14007ffeac0)(rand),
        Sel: (*ast.Ident)(0x14007ffeae0)(Intn)
       }),
       Lparen: (token.Pos) 4233751,
       Args: ([]ast.Expr) (len=1 cap=1) {
        (*ast.CallExpr)(0x1400847a080)({
         Fun: (*ast.Ident)(0x14007ffeb00)(len),
         Lparen: (token.Pos) 4233755,
         Args: ([]ast.Expr) (len=1 cap=1) {
          (*ast.Ident)(0x14007ffeb20)(people)
         },
         Ellipsis: (token.Pos) 0,
         Rparen: (token.Pos) 4233762
        })
       },
       Ellipsis: (token.Pos) 0,
       Rparen: (token.Pos) 4233763
      })
     }
    }),
    (*ast.ReturnStmt)(0x14007ffeba0)({
     Return: (token.Pos) 4233766,
     Results: ([]ast.Expr) (len=1 cap=1) {
      (*ast.IndexExpr)(0x14008478510)({
       X: (*ast.Ident)(0x14007ffeb60)(people),
       Lbrack: (token.Pos) 4233779,
       Index: (*ast.Ident)(0x14007ffeb80)(i),
       Rbrack: (token.Pos) 4233781
      })
     }
    })
   },
   Rbrace: (token.Pos) 4233783
  })
 }),
 imports: (map[string]string) {
 },
 aliases: (map[string]string) <nil>
})

GML: getFunctionsNeedingWrappers: f.Decls[13]=
(*ast.FuncDecl)(0x14008478840)({
 Doc: (*ast.CommentGroup)(0x14007f9b8c0)({
  List: ([]*ast.Comment) (len=1 cap=1) {
   (*ast.Comment)(0x14007f9b8a8)({
    Slash: (token.Pos) 4233786,
    Text: (string) (len=25) "// Gets a list of people."
   })
  }
 }),
 Recv: (*ast.FieldList)(<nil>),
 Name: (*ast.Ident)(0x14007ffebe0)(GetPeople),
 Type: (*ast.FuncType)(0x14007ffef20)({
  Func: (token.Pos) 4233812,
  TypeParams: (*ast.FieldList)(<nil>),
  Params: (*ast.FieldList)(0x140084785a0)({
   Opening: (token.Pos) 4233826,
   List: ([]*ast.Field) <nil>,
   Closing: (token.Pos) 4233827
  }),
  Results: (*ast.FieldList)(0x14008478600)({
   Opening: (token.Pos) 0,
   List: ([]*ast.Field) (len=1 cap=1) {
    (*ast.Field)(0x1400847a180)({
     Doc: (*ast.CommentGroup)(<nil>),
     Names: ([]*ast.Ident) <nil>,
     Type: (*ast.ArrayType)(0x140084785d0)({
      Lbrack: (token.Pos) 4233829,
      Len: (ast.Expr) <nil>,
      Elt: (*ast.Ident)(0x14007ffec00)(Person)
     }),
     Tag: (*ast.BasicLit)(<nil>),
     Comment: (*ast.CommentGroup)(<nil>)
    })
   },
   Closing: (token.Pos) 0
  })
 }),
 Body: (*ast.BlockStmt)(0x14008478810)({
  Lbrace: (token.Pos) 4233838,
  List: ([]ast.Stmt) (len=1 cap=1) {
   (*ast.ReturnStmt)(0x14007ffef00)({
    Return: (token.Pos) 4233841,
    Results: ([]ast.Expr) (len=1 cap=1) {
     (*ast.CompositeLit)(0x1400847a380)({
      Type: (*ast.ArrayType)(0x14008478630)({
       Lbrack: (token.Pos) 4233848,
       Len: (ast.Expr) <nil>,
       Elt: (*ast.Ident)(0x14007ffec20)(Person)
      }),
      Lbrace: (token.Pos) 4233856,
      Elts: ([]ast.Expr) (len=3 cap=4) {
       (*ast.CompositeLit)(0x1400847a200)({
        Type: (ast.Expr) <nil>,
        Lbrace: (token.Pos) 4233860,
        Elts: ([]ast.Expr) (len=3 cap=4) {
         (*ast.KeyValueExpr)(0x14008478660)({
          Key: (*ast.Ident)(0x14007ffec40)(FirstName),
          Colon: (token.Pos) 4233874,
          Value: (*ast.BasicLit)(0x14007ffec60)({
           ValuePos: (token.Pos) 4233876,
           Kind: (token.Token) STRING,
           Value: (string) (len=5) "\"Bob\""
          })
         }),
         (*ast.KeyValueExpr)(0x14008478690)({
          Key: (*ast.Ident)(0x14007ffec80)(LastName),
          Colon: (token.Pos) 4233894,
          Value: (*ast.BasicLit)(0x14007ffeca0)({
           ValuePos: (token.Pos) 4233897,
           Kind: (token.Token) STRING,
           Value: (string) (len=7) "\"Smith\""
          })
         }),
         (*ast.KeyValueExpr)(0x140084786c0)({
          Key: (*ast.Ident)(0x14007ffece0)(Age),
          Colon: (token.Pos) 4233912,
          Value: (*ast.BasicLit)(0x14007ffed00)({
           ValuePos: (token.Pos) 4233920,
           Kind: (token.Token) INT,
           Value: (string) (len=2) "42"
          })
         })
        },
        Rbrace: (token.Pos) 4233926,
        Incomplete: (bool) false
       }),
       (*ast.CompositeLit)(0x1400847a280)({
        Type: (ast.Expr) <nil>,
        Lbrace: (token.Pos) 4233931,
        Elts: ([]ast.Expr) (len=3 cap=4) {
         (*ast.KeyValueExpr)(0x140084786f0)({
          Key: (*ast.Ident)(0x14007ffed20)(FirstName),
          Colon: (token.Pos) 4233945,
          Value: (*ast.BasicLit)(0x14007ffed40)({
           ValuePos: (token.Pos) 4233947,
           Kind: (token.Token) STRING,
           Value: (string) (len=7) "\"Alice\""
          })
         }),
         (*ast.KeyValueExpr)(0x14008478720)({
          Key: (*ast.Ident)(0x14007ffed60)(LastName),
          Colon: (token.Pos) 4233967,
          Value: (*ast.BasicLit)(0x14007ffed80)({
           ValuePos: (token.Pos) 4233970,
           Kind: (token.Token) STRING,
           Value: (string) (len=7) "\"Jones\""
          })
         }),
         (*ast.KeyValueExpr)(0x14008478750)({
          Key: (*ast.Ident)(0x14007ffedc0)(Age),
          Colon: (token.Pos) 4233985,
          Value: (*ast.BasicLit)(0x14007ffede0)({
           ValuePos: (token.Pos) 4233993,
           Kind: (token.Token) INT,
           Value: (string) (len=2) "35"
          })
         })
        },
        Rbrace: (token.Pos) 4233999,
        Incomplete: (bool) false
       }),
       (*ast.CompositeLit)(0x1400847a300)({
        Type: (ast.Expr) <nil>,
        Lbrace: (token.Pos) 4234004,
        Elts: ([]ast.Expr) (len=3 cap=4) {
         (*ast.KeyValueExpr)(0x14008478780)({
          Key: (*ast.Ident)(0x14007ffee20)(FirstName),
          Colon: (token.Pos) 4234018,
          Value: (*ast.BasicLit)(0x14007ffee40)({
           ValuePos: (token.Pos) 4234020,
           Kind: (token.Token) STRING,
           Value: (string) (len=9) "\"Charlie\""
          })
         }),
         (*ast.KeyValueExpr)(0x140084787b0)({
          Key: (*ast.Ident)(0x14007ffee60)(LastName),
          Colon: (token.Pos) 4234042,
          Value: (*ast.BasicLit)(0x14007ffee80)({
           ValuePos: (token.Pos) 4234045,
           Kind: (token.Token) STRING,
           Value: (string) (len=7) "\"Brown\""
          })
         }),
         (*ast.KeyValueExpr)(0x140084787e0)({
          Key: (*ast.Ident)(0x14007ffeec0)(Age),
          Colon: (token.Pos) 4234060,
          Value: (*ast.BasicLit)(0x14007ffeee0)({
           ValuePos: (token.Pos) 4234068,
           Kind: (token.Token) INT,
           Value: (string) (len=1) "8"
          })
         })
        },
        Rbrace: (token.Pos) 4234073,
        Incomplete: (bool) false
       })
      },
      Rbrace: (token.Pos) 4234077,
      Incomplete: (bool) false
     })
    }
   })
  },
  Rbrace: (token.Pos) 4234079
 })
})

GML: getFunctionsNeedingWrappers: results.append: info=
(*codegen.funcInfo)(0x14008480ba0)({
 function: (*ast.FuncDecl)(0x14008478840)({
  Doc: (*ast.CommentGroup)(0x14007f9b8c0)({
   List: ([]*ast.Comment) (len=1 cap=1) {
    (*ast.Comment)(0x14007f9b8a8)({
     Slash: (token.Pos) 4233786,
     Text: (string) (len=25) "// Gets a list of people."
    })
   }
  }),
  Recv: (*ast.FieldList)(<nil>),
  Name: (*ast.Ident)(0x14007ffebe0)(GetPeople),
  Type: (*ast.FuncType)(0x14007ffef20)({
   Func: (token.Pos) 4233812,
   TypeParams: (*ast.FieldList)(<nil>),
   Params: (*ast.FieldList)(0x140084785a0)({
    Opening: (token.Pos) 4233826,
    List: ([]*ast.Field) <nil>,
    Closing: (token.Pos) 4233827
   }),
   Results: (*ast.FieldList)(0x14008478600)({
    Opening: (token.Pos) 0,
    List: ([]*ast.Field) (len=1 cap=1) {
     (*ast.Field)(0x1400847a180)({
      Doc: (*ast.CommentGroup)(<nil>),
      Names: ([]*ast.Ident) <nil>,
      Type: (*ast.ArrayType)(0x140084785d0)({
       Lbrack: (token.Pos) 4233829,
       Len: (ast.Expr) <nil>,
       Elt: (*ast.Ident)(0x14007ffec00)(Person)
      }),
      Tag: (*ast.BasicLit)(<nil>),
      Comment: (*ast.CommentGroup)(<nil>)
     })
    },
    Closing: (token.Pos) 0
   })
  }),
  Body: (*ast.BlockStmt)(0x14008478810)({
   Lbrace: (token.Pos) 4233838,
   List: ([]ast.Stmt) (len=1 cap=1) {
    (*ast.ReturnStmt)(0x14007ffef00)({
     Return: (token.Pos) 4233841,
     Results: ([]ast.Expr) (len=1 cap=1) {
      (*ast.CompositeLit)(0x1400847a380)({
       Type: (*ast.ArrayType)(0x14008478630)({
        Lbrack: (token.Pos) 4233848,
        Len: (ast.Expr) <nil>,
        Elt: (*ast.Ident)(0x14007ffec20)(Person)
       }),
       Lbrace: (token.Pos) 4233856,
       Elts: ([]ast.Expr) (len=3 cap=4) {
        (*ast.CompositeLit)(0x1400847a200)({
         Type: (ast.Expr) <nil>,
         Lbrace: (token.Pos) 4233860,
         Elts: ([]ast.Expr) (len=3 cap=4) {
          (*ast.KeyValueExpr)(0x14008478660)({
           Key: (*ast.Ident)(0x14007ffec40)(FirstName),
           Colon: (token.Pos) 4233874,
           Value: (*ast.BasicLit)(0x14007ffec60)({
            ValuePos: (token.Pos) 4233876,
            Kind: (token.Token) STRING,
            Value: (string) (len=5) "\"Bob\""
           })
          }),
          (*ast.KeyValueExpr)(0x14008478690)({
           Key: (*ast.Ident)(0x14007ffec80)(LastName),
           Colon: (token.Pos) 4233894,
           Value: (*ast.BasicLit)(0x14007ffeca0)({
            ValuePos: (token.Pos) 4233897,
            Kind: (token.Token) STRING,
            Value: (string) (len=7) "\"Smith\""
           })
          }),
          (*ast.KeyValueExpr)(0x140084786c0)({
           Key: (*ast.Ident)(0x14007ffece0)(Age),
           Colon: (token.Pos) 4233912,
           Value: (*ast.BasicLit)(0x14007ffed00)({
            ValuePos: (token.Pos) 4233920,
            Kind: (token.Token) INT,
            Value: (string) (len=2) "42"
           })
          })
         },
         Rbrace: (token.Pos) 4233926,
         Incomplete: (bool) false
        }),
        (*ast.CompositeLit)(0x1400847a280)({
         Type: (ast.Expr) <nil>,
         Lbrace: (token.Pos) 4233931,
         Elts: ([]ast.Expr) (len=3 cap=4) {
          (*ast.KeyValueExpr)(0x140084786f0)({
           Key: (*ast.Ident)(0x14007ffed20)(FirstName),
           Colon: (token.Pos) 4233945,
           Value: (*ast.BasicLit)(0x14007ffed40)({
            ValuePos: (token.Pos) 4233947,
            Kind: (token.Token) STRING,
            Value: (string) (len=7) "\"Alice\""
           })
          }),
          (*ast.KeyValueExpr)(0x14008478720)({
           Key: (*ast.Ident)(0x14007ffed60)(LastName),
           Colon: (token.Pos) 4233967,
           Value: (*ast.BasicLit)(0x14007ffed80)({
            ValuePos: (token.Pos) 4233970,
            Kind: (token.Token) STRING,
            Value: (string) (len=7) "\"Jones\""
           })
          }),
          (*ast.KeyValueExpr)(0x14008478750)({
           Key: (*ast.Ident)(0x14007ffedc0)(Age),
           Colon: (token.Pos) 4233985,
           Value: (*ast.BasicLit)(0x14007ffede0)({
            ValuePos: (token.Pos) 4233993,
            Kind: (token.Token) INT,
            Value: (string) (len=2) "35"
           })
          })
         },
         Rbrace: (token.Pos) 4233999,
         Incomplete: (bool) false
        }),
        (*ast.CompositeLit)(0x1400847a300)({
         Type: (ast.Expr) <nil>,
         Lbrace: (token.Pos) 4234004,
         Elts: ([]ast.Expr) (len=3 cap=4) {
          (*ast.KeyValueExpr)(0x14008478780)({
           Key: (*ast.Ident)(0x14007ffee20)(FirstName),
           Colon: (token.Pos) 4234018,
           Value: (*ast.BasicLit)(0x14007ffee40)({
            ValuePos: (token.Pos) 4234020,
            Kind: (token.Token) STRING,
            Value: (string) (len=9) "\"Charlie\""
           })
          }),
          (*ast.KeyValueExpr)(0x140084787b0)({
           Key: (*ast.Ident)(0x14007ffee60)(LastName),
           Colon: (token.Pos) 4234042,
           Value: (*ast.BasicLit)(0x14007ffee80)({
            ValuePos: (token.Pos) 4234045,
            Kind: (token.Token) STRING,
            Value: (string) (len=7) "\"Brown\""
           })
          }),
          (*ast.KeyValueExpr)(0x140084787e0)({
           Key: (*ast.Ident)(0x14007ffeec0)(Age),
           Colon: (token.Pos) 4234060,
           Value: (*ast.BasicLit)(0x14007ffeee0)({
            ValuePos: (token.Pos) 4234068,
            Kind: (token.Token) INT,
            Value: (string) (len=1) "8"
           })
          })
         },
         Rbrace: (token.Pos) 4234073,
         Incomplete: (bool) false
        })
       },
       Rbrace: (token.Pos) 4234077,
       Incomplete: (bool) false
      })
     }
    })
   },
   Rbrace: (token.Pos) 4234079
  })
 }),
 imports: (map[string]string) {
 },
 aliases: (map[string]string) <nil>
})

GML: getFunctionsNeedingWrappers: f.Decls[14]=
(*ast.FuncDecl)(0x14008478930)({
 Doc: (*ast.CommentGroup)(0x14007f9b8f0)({
  List: ([]*ast.Comment) (len=1 cap=1) {
   (*ast.Comment)(0x14007f9b8d8)({
    Slash: (token.Pos) 4234082,
    Text: (string) (len=37) "// Gets the name and age of a person."
   })
  }
 }),
 Recv: (*ast.FieldList)(<nil>),
 Name: (*ast.Ident)(0x14007ffef40)(GetNameAndAge),
 Type: (*ast.FuncType)(0x14007fff180)({
  Func: (token.Pos) 4234120,
  TypeParams: (*ast.FieldList)(<nil>),
  Params: (*ast.FieldList)(0x14008478870)({
   Opening: (token.Pos) 4234138,
   List: ([]*ast.Field) <nil>,
   Closing: (token.Pos) 4234139
  }),
  Results: (*ast.FieldList)(0x140084788d0)({
   Opening: (token.Pos) 4234141,
   List: ([]*ast.Field) (len=2 cap=2) {
    (*ast.Field)(0x1400847a3c0)({
     Doc: (*ast.CommentGroup)(<nil>),
     Names: ([]*ast.Ident) (len=1 cap=1) {
      (*ast.Ident)(0x14007ffef60)(name)
     },
     Type: (*ast.Ident)(0x14007ffef80)(string),
     Tag: (*ast.BasicLit)(<nil>),
     Comment: (*ast.CommentGroup)(<nil>)
    }),
    (*ast.Field)(0x1400847a400)({
     Doc: (*ast.CommentGroup)(<nil>),
     Names: ([]*ast.Ident) (len=1 cap=1) {
      (*ast.Ident)(0x14007ffefa0)(age)
     },
     Type: (*ast.Ident)(0x14007ffefc0)(int),
     Tag: (*ast.BasicLit)(<nil>),
     Comment: (*ast.CommentGroup)(<nil>)
    })
   },
   Closing: (token.Pos) 4234162
  })
 }),
 Body: (*ast.BlockStmt)(0x14008478900)({
  Lbrace: (token.Pos) 4234164,
  List: ([]ast.Stmt) (len=2 cap=2) {
   (*ast.AssignStmt)(0x1400847a480)({
    Lhs: ([]ast.Expr) (len=1 cap=1) {
     (*ast.Ident)(0x14007ffefe0)(p)
    },
    TokPos: (token.Pos) 4234169,
    Tok: (token.Token) :=,
    Rhs: ([]ast.Expr) (len=1 cap=1) {
     (*ast.CallExpr)(0x1400847a440)({
      Fun: (*ast.Ident)(0x14007fff000)(GetPerson),
      Lparen: (token.Pos) 4234181,
      Args: ([]ast.Expr) <nil>,
      Ellipsis: (token.Pos) 0,
      Rparen: (token.Pos) 4234182
     })
    }
   }),
   (*ast.ReturnStmt)(0x14007fff140)({
    Return: (token.Pos) 4234185,
    Results: ([]ast.Expr) (len=2 cap=2) {
     (*ast.CallExpr)(0x1400847a4c0)({
      Fun: (*ast.Ident)(0x14007fff020)(GetFullName),
      Lparen: (token.Pos) 4234203,
      Args: ([]ast.Expr) (len=2 cap=2) {
       (*ast.SelectorExpr)(0x14007f9b920)({
        X: (*ast.Ident)(0x14007fff040)(p),
        Sel: (*ast.Ident)(0x14007fff060)(FirstName)
       }),
       (*ast.SelectorExpr)(0x14007f9b938)({
        X: (*ast.Ident)(0x14007fff080)(p),
        Sel: (*ast.Ident)(0x14007fff0a0)(LastName)
       })
      },
      Ellipsis: (token.Pos) 0,
      Rparen: (token.Pos) 4234227
     }),
     (*ast.SelectorExpr)(0x14007f9b950)({
      X: (*ast.Ident)(0x14007fff0e0)(p),
      Sel: (*ast.Ident)(0x14007fff100)(Age)
     })
    }
   })
  },
  Rbrace: (token.Pos) 4234236
 })
})

GML: getFunctionsNeedingWrappers: results.append: info=
(*codegen.funcInfo)(0x14008480c30)({
 function: (*ast.FuncDecl)(0x14008478930)({
  Doc: (*ast.CommentGroup)(0x14007f9b8f0)({
   List: ([]*ast.Comment) (len=1 cap=1) {
    (*ast.Comment)(0x14007f9b8d8)({
     Slash: (token.Pos) 4234082,
     Text: (string) (len=37) "// Gets the name and age of a person."
    })
   }
  }),
  Recv: (*ast.FieldList)(<nil>),
  Name: (*ast.Ident)(0x14007ffef40)(GetNameAndAge),
  Type: (*ast.FuncType)(0x14007fff180)({
   Func: (token.Pos) 4234120,
   TypeParams: (*ast.FieldList)(<nil>),
   Params: (*ast.FieldList)(0x14008478870)({
    Opening: (token.Pos) 4234138,
    List: ([]*ast.Field) <nil>,
    Closing: (token.Pos) 4234139
   }),
   Results: (*ast.FieldList)(0x140084788d0)({
    Opening: (token.Pos) 4234141,
    List: ([]*ast.Field) (len=2 cap=2) {
     (*ast.Field)(0x1400847a3c0)({
      Doc: (*ast.CommentGroup)(<nil>),
      Names: ([]*ast.Ident) (len=1 cap=1) {
       (*ast.Ident)(0x14007ffef60)(name)
      },
      Type: (*ast.Ident)(0x14007ffef80)(string),
      Tag: (*ast.BasicLit)(<nil>),
      Comment: (*ast.CommentGroup)(<nil>)
     }),
     (*ast.Field)(0x1400847a400)({
      Doc: (*ast.CommentGroup)(<nil>),
      Names: ([]*ast.Ident) (len=1 cap=1) {
       (*ast.Ident)(0x14007ffefa0)(age)
      },
      Type: (*ast.Ident)(0x14007ffefc0)(int),
      Tag: (*ast.BasicLit)(<nil>),
      Comment: (*ast.CommentGroup)(<nil>)
     })
    },
    Closing: (token.Pos) 4234162
   })
  }),
  Body: (*ast.BlockStmt)(0x14008478900)({
   Lbrace: (token.Pos) 4234164,
   List: ([]ast.Stmt) (len=2 cap=2) {
    (*ast.AssignStmt)(0x1400847a480)({
     Lhs: ([]ast.Expr) (len=1 cap=1) {
      (*ast.Ident)(0x14007ffefe0)(p)
     },
     TokPos: (token.Pos) 4234169,
     Tok: (token.Token) :=,
     Rhs: ([]ast.Expr) (len=1 cap=1) {
      (*ast.CallExpr)(0x1400847a440)({
       Fun: (*ast.Ident)(0x14007fff000)(GetPerson),
       Lparen: (token.Pos) 4234181,
       Args: ([]ast.Expr) <nil>,
       Ellipsis: (token.Pos) 0,
       Rparen: (token.Pos) 4234182
      })
     }
    }),
    (*ast.ReturnStmt)(0x14007fff140)({
     Return: (token.Pos) 4234185,
     Results: ([]ast.Expr) (len=2 cap=2) {
      (*ast.CallExpr)(0x1400847a4c0)({
       Fun: (*ast.Ident)(0x14007fff020)(GetFullName),
       Lparen: (token.Pos) 4234203,
       Args: ([]ast.Expr) (len=2 cap=2) {
        (*ast.SelectorExpr)(0x14007f9b920)({
         X: (*ast.Ident)(0x14007fff040)(p),
         Sel: (*ast.Ident)(0x14007fff060)(FirstName)
        }),
        (*ast.SelectorExpr)(0x14007f9b938)({
         X: (*ast.Ident)(0x14007fff080)(p),
         Sel: (*ast.Ident)(0x14007fff0a0)(LastName)
        })
       },
       Ellipsis: (token.Pos) 0,
       Rparen: (token.Pos) 4234227
      }),
      (*ast.SelectorExpr)(0x14007f9b950)({
       X: (*ast.Ident)(0x14007fff0e0)(p),
       Sel: (*ast.Ident)(0x14007fff100)(Age)
      })
     }
    })
   },
   Rbrace: (token.Pos) 4234236
  })
 }),
 imports: (map[string]string) {
 },
 aliases: (map[string]string) <nil>
})

GML: getFunctionsNeedingWrappers: f.Decls[15]=
(*ast.FuncDecl)(0x14008478ab0)({
 Doc: (*ast.CommentGroup)(0x14007f9b980)({
  List: ([]*ast.Comment) (len=1 cap=1) {
   (*ast.Comment)(0x14007f9b968)({
    Slash: (token.Pos) 4234239,
    Text: (string) (len=28) "// Tests returning an error."
   })
  }
 }),
 Recv: (*ast.FieldList)(<nil>),
 Name: (*ast.Ident)(0x14007fff1a0)(TestNormalError),
 Type: (*ast.FuncType)(0x14007fff460)({
  Func: (token.Pos) 4234268,
  TypeParams: (*ast.FieldList)(<nil>),
  Params: (*ast.FieldList)(0x14008478960)({
   Opening: (token.Pos) 4234288,
   List: ([]*ast.Field) (len=1 cap=1) {
    (*ast.Field)(0x1400847a500)({
     Doc: (*ast.CommentGroup)(<nil>),
     Names: ([]*ast.Ident) (len=1 cap=1) {
      (*ast.Ident)(0x14007fff1c0)(input)
     },
     Type: (*ast.Ident)(0x14007fff1e0)(string),
     Tag: (*ast.BasicLit)(<nil>),
     Comment: (*ast.CommentGroup)(<nil>)
    })
   },
   Closing: (token.Pos) 4234301
  }),
  Results: (*ast.FieldList)(0x140084789c0)({
   Opening: (token.Pos) 4234303,
   List: ([]*ast.Field) (len=2 cap=2) {
    (*ast.Field)(0x1400847a540)({
     Doc: (*ast.CommentGroup)(<nil>),
     Names: ([]*ast.Ident) <nil>,
     Type: (*ast.Ident)(0x14007fff200)(string),
     Tag: (*ast.BasicLit)(<nil>),
     Comment: (*ast.CommentGroup)(<nil>)
    }),
    (*ast.Field)(0x1400847a580)({
     Doc: (*ast.CommentGroup)(<nil>),
     Names: ([]*ast.Ident) <nil>,
     Type: (*ast.Ident)(0x14007fff220)(error),
     Tag: (*ast.BasicLit)(<nil>),
     Comment: (*ast.CommentGroup)(<nil>)
    })
   },
   Closing: (token.Pos) 4234317
  })
 }),
 Body: (*ast.BlockStmt)(0x14008478a80)({
  Lbrace: (token.Pos) 4234319,
  List: ([]ast.Stmt) (len=3 cap=4) {
   (*ast.IfStmt)(0x1400847a640)({
    If: (token.Pos) 4234645,
    Init: (ast.Stmt) <nil>,
    Cond: (*ast.BinaryExpr)(0x140084789f0)({
     X: (*ast.Ident)(0x14007fff260)(input),
     OpPos: (token.Pos) 4234654,
     Op: (token.Token) ==,
     Y: (*ast.BasicLit)(0x14007fff280)({
      ValuePos: (token.Pos) 4234657,
      Kind: (token.Token) STRING,
      Value: (string) (len=2) "\"\""
     })
    }),
    Body: (*ast.BlockStmt)(0x14008478a20)({
     Lbrace: (token.Pos) 4234660,
     List: ([]ast.Stmt) (len=1 cap=1) {
      (*ast.ReturnStmt)(0x14007fff340)({
       Return: (token.Pos) 4234664,
       Results: ([]ast.Expr) (len=2 cap=2) {
        (*ast.BasicLit)(0x14007fff2a0)({
         ValuePos: (token.Pos) 4234671,
         Kind: (token.Token) STRING,
         Value: (string) (len=2) "\"\""
        }),
        (*ast.CallExpr)(0x1400847a600)({
         Fun: (*ast.SelectorExpr)(0x14007f9ba58)({
          X: (*ast.Ident)(0x14007fff2c0)(errors),
          Sel: (*ast.Ident)(0x14007fff2e0)(New)
         }),
         Lparen: (token.Pos) 4234685,
         Args: ([]ast.Expr) (len=1 cap=1) {
          (*ast.BasicLit)(0x14007fff300)({
           ValuePos: (token.Pos) 4234686,
           Kind: (token.Token) STRING,
           Value: (string) (len=16) "\"input is empty\""
          })
         },
         Ellipsis: (token.Pos) 0,
         Rparen: (token.Pos) 4234702
        })
       }
      })
     },
     Rbrace: (token.Pos) 4234705
    }),
    Else: (ast.Stmt) <nil>
   }),
   (*ast.AssignStmt)(0x1400847a680)({
    Lhs: ([]ast.Expr) (len=1 cap=1) {
     (*ast.Ident)(0x14007fff360)(output)
    },
    TokPos: (token.Pos) 4234715,
    Tok: (token.Token) :=,
    Rhs: ([]ast.Expr) (len=1 cap=1) {
     (*ast.BinaryExpr)(0x14008478a50)({
      X: (*ast.BasicLit)(0x14007fff380)({
       ValuePos: (token.Pos) 4234718,
       Kind: (token.Token) STRING,
       Value: (string) (len=12) "\"You said: \""
      }),
      OpPos: (token.Pos) 4234731,
      Op: (token.Token) +,
      Y: (*ast.Ident)(0x14007fff3a0)(input)
     })
    }
   }),
   (*ast.ReturnStmt)(0x14007fff440)({
    Return: (token.Pos) 4234740,
    Results: ([]ast.Expr) (len=2 cap=2) {
     (*ast.Ident)(0x14007fff3e0)(output),
     (*ast.Ident)(0x14007fff400)(nil)
    }
   })
  },
  Rbrace: (token.Pos) 4234759
 })
})

GML: getFunctionsNeedingWrappers: results.append: info=
(*codegen.funcInfo)(0x14008480cc0)({
 function: (*ast.FuncDecl)(0x14008478ab0)({
  Doc: (*ast.CommentGroup)(0x14007f9b980)({
   List: ([]*ast.Comment) (len=1 cap=1) {
    (*ast.Comment)(0x14007f9b968)({
     Slash: (token.Pos) 4234239,
     Text: (string) (len=28) "// Tests returning an error."
    })
   }
  }),
  Recv: (*ast.FieldList)(<nil>),
  Name: (*ast.Ident)(0x14007fff1a0)(TestNormalError),
  Type: (*ast.FuncType)(0x14007fff460)({
   Func: (token.Pos) 4234268,
   TypeParams: (*ast.FieldList)(<nil>),
   Params: (*ast.FieldList)(0x14008478960)({
    Opening: (token.Pos) 4234288,
    List: ([]*ast.Field) (len=1 cap=1) {
     (*ast.Field)(0x1400847a500)({
      Doc: (*ast.CommentGroup)(<nil>),
      Names: ([]*ast.Ident) (len=1 cap=1) {
       (*ast.Ident)(0x14007fff1c0)(input)
      },
      Type: (*ast.Ident)(0x14007fff1e0)(string),
      Tag: (*ast.BasicLit)(<nil>),
      Comment: (*ast.CommentGroup)(<nil>)
     })
    },
    Closing: (token.Pos) 4234301
   }),
   Results: (*ast.FieldList)(0x140084789c0)({
    Opening: (token.Pos) 4234303,
    List: ([]*ast.Field) (len=2 cap=2) {
     (*ast.Field)(0x1400847a540)({
      Doc: (*ast.CommentGroup)(<nil>),
      Names: ([]*ast.Ident) <nil>,
      Type: (*ast.Ident)(0x14007fff200)(string),
      Tag: (*ast.BasicLit)(<nil>),
      Comment: (*ast.CommentGroup)(<nil>)
     }),
     (*ast.Field)(0x1400847a580)({
      Doc: (*ast.CommentGroup)(<nil>),
      Names: ([]*ast.Ident) <nil>,
      Type: (*ast.Ident)(0x14007fff220)(error),
      Tag: (*ast.BasicLit)(<nil>),
      Comment: (*ast.CommentGroup)(<nil>)
     })
    },
    Closing: (token.Pos) 4234317
   })
  }),
  Body: (*ast.BlockStmt)(0x14008478a80)({
   Lbrace: (token.Pos) 4234319,
   List: ([]ast.Stmt) (len=3 cap=4) {
    (*ast.IfStmt)(0x1400847a640)({
     If: (token.Pos) 4234645,
     Init: (ast.Stmt) <nil>,
     Cond: (*ast.BinaryExpr)(0x140084789f0)({
      X: (*ast.Ident)(0x14007fff260)(input),
      OpPos: (token.Pos) 4234654,
      Op: (token.Token) ==,
      Y: (*ast.BasicLit)(0x14007fff280)({
       ValuePos: (token.Pos) 4234657,
       Kind: (token.Token) STRING,
       Value: (string) (len=2) "\"\""
      })
     }),
     Body: (*ast.BlockStmt)(0x14008478a20)({
      Lbrace: (token.Pos) 4234660,
      List: ([]ast.Stmt) (len=1 cap=1) {
       (*ast.ReturnStmt)(0x14007fff340)({
        Return: (token.Pos) 4234664,
        Results: ([]ast.Expr) (len=2 cap=2) {
         (*ast.BasicLit)(0x14007fff2a0)({
          ValuePos: (token.Pos) 4234671,
          Kind: (token.Token) STRING,
          Value: (string) (len=2) "\"\""
         }),
         (*ast.CallExpr)(0x1400847a600)({
          Fun: (*ast.SelectorExpr)(0x14007f9ba58)({
           X: (*ast.Ident)(0x14007fff2c0)(errors),
           Sel: (*ast.Ident)(0x14007fff2e0)(New)
          }),
          Lparen: (token.Pos) 4234685,
          Args: ([]ast.Expr) (len=1 cap=1) {
           (*ast.BasicLit)(0x14007fff300)({
            ValuePos: (token.Pos) 4234686,
            Kind: (token.Token) STRING,
            Value: (string) (len=16) "\"input is empty\""
           })
          },
          Ellipsis: (token.Pos) 0,
          Rparen: (token.Pos) 4234702
         })
        }
       })
      },
      Rbrace: (token.Pos) 4234705
     }),
     Else: (ast.Stmt) <nil>
    }),
    (*ast.AssignStmt)(0x1400847a680)({
     Lhs: ([]ast.Expr) (len=1 cap=1) {
      (*ast.Ident)(0x14007fff360)(output)
     },
     TokPos: (token.Pos) 4234715,
     Tok: (token.Token) :=,
     Rhs: ([]ast.Expr) (len=1 cap=1) {
      (*ast.BinaryExpr)(0x14008478a50)({
       X: (*ast.BasicLit)(0x14007fff380)({
        ValuePos: (token.Pos) 4234718,
        Kind: (token.Token) STRING,
        Value: (string) (len=12) "\"You said: \""
       }),
       OpPos: (token.Pos) 4234731,
       Op: (token.Token) +,
       Y: (*ast.Ident)(0x14007fff3a0)(input)
      })
     }
    }),
    (*ast.ReturnStmt)(0x14007fff440)({
     Return: (token.Pos) 4234740,
     Results: ([]ast.Expr) (len=2 cap=2) {
      (*ast.Ident)(0x14007fff3e0)(output),
      (*ast.Ident)(0x14007fff400)(nil)
     }
    })
   },
   Rbrace: (token.Pos) 4234759
  })
 }),
 imports: (map[string]string) {
 },
 aliases: (map[string]string) <nil>
})

GML: getFunctionsNeedingWrappers: f.Decls[16]=
(*ast.FuncDecl)(0x14008478c00)({
 Doc: (*ast.CommentGroup)(0x14007f9ba88)({
  List: ([]*ast.Comment) (len=1 cap=1) {
   (*ast.Comment)(0x14007f9ba70)({
    Slash: (token.Pos) 4234762,
    Text: (string) (len=58) "// Tests an alternative way to handle errors in functions."
   })
  }
 }),
 Recv: (*ast.FieldList)(<nil>),
 Name: (*ast.Ident)(0x14007fff480)(TestAlternativeError),
 Type: (*ast.FuncType)(0x14007fff6c0)({
  Func: (token.Pos) 4234821,
  TypeParams: (*ast.FieldList)(<nil>),
  Params: (*ast.FieldList)(0x14008478ae0)({
   Opening: (token.Pos) 4234846,
   List: ([]*ast.Field) (len=1 cap=1) {
    (*ast.Field)(0x1400847a700)({
     Doc: (*ast.CommentGroup)(<nil>),
     Names: ([]*ast.Ident) (len=1 cap=1) {
      (*ast.Ident)(0x14007fff4a0)(input)
     },
     Type: (*ast.Ident)(0x14007fff4c0)(string),
     Tag: (*ast.BasicLit)(<nil>),
     Comment: (*ast.CommentGroup)(<nil>)
    })
   },
   Closing: (token.Pos) 4234859
  }),
  Results: (*ast.FieldList)(0x14008478b10)({
   Opening: (token.Pos) 0,
   List: ([]*ast.Field) (len=1 cap=1) {
    (*ast.Field)(0x1400847a740)({
     Doc: (*ast.CommentGroup)(<nil>),
     Names: ([]*ast.Ident) <nil>,
     Type: (*ast.Ident)(0x14007fff4e0)(string),
     Tag: (*ast.BasicLit)(<nil>),
     Comment: (*ast.CommentGroup)(<nil>)
    })
   },
   Closing: (token.Pos) 0
  })
 }),
 Body: (*ast.BlockStmt)(0x14008478bd0)({
  Lbrace: (token.Pos) 4234868,
  List: ([]ast.Stmt) (len=3 cap=4) {
   (*ast.IfStmt)(0x1400847a7c0)({
    If: (token.Pos) 4235012,
    Init: (ast.Stmt) <nil>,
    Cond: (*ast.BinaryExpr)(0x14008478b40)({
     X: (*ast.Ident)(0x14007fff500)(input),
     OpPos: (token.Pos) 4235021,
     Op: (token.Token) ==,
     Y: (*ast.BasicLit)(0x14007fff520)({
      ValuePos: (token.Pos) 4235024,
      Kind: (token.Token) STRING,
      Value: (string) (len=2) "\"\""
     })
    }),
    Body: (*ast.BlockStmt)(0x14008478b70)({
     Lbrace: (token.Pos) 4235027,
     List: ([]ast.Stmt) (len=2 cap=2) {
      (*ast.ExprStmt)(0x14007fc7c50)({
       X: (*ast.CallExpr)(0x1400847a780)({
        Fun: (*ast.SelectorExpr)(0x14007f9bb00)({
         X: (*ast.Ident)(0x14007fff540)(console),
         Sel: (*ast.Ident)(0x14007fff560)(Error)
        }),
        Lparen: (token.Pos) 4235044,
        Args: ([]ast.Expr) (len=1 cap=1) {
         (*ast.BasicLit)(0x14007fff580)({
          ValuePos: (token.Pos) 4235045,
          Kind: (token.Token) STRING,
          Value: (string) (len=16) "\"input is empty\""
         })
        },
        Ellipsis: (token.Pos) 0,
        Rparen: (token.Pos) 4235061
       })
      }),
      (*ast.ReturnStmt)(0x14007fff5c0)({
       Return: (token.Pos) 4235065,
       Results: ([]ast.Expr) (len=1 cap=1) {
        (*ast.BasicLit)(0x14007fff5a0)({
         ValuePos: (token.Pos) 4235072,
         Kind: (token.Token) STRING,
         Value: (string) (len=2) "\"\""
        })
       }
      })
     },
     Rbrace: (token.Pos) 4235076
    }),
    Else: (ast.Stmt) <nil>
   }),
   (*ast.AssignStmt)(0x1400847a800)({
    Lhs: ([]ast.Expr) (len=1 cap=1) {
     (*ast.Ident)(0x14007fff600)(output)
    },
    TokPos: (token.Pos) 4235086,
    Tok: (token.Token) :=,
    Rhs: ([]ast.Expr) (len=1 cap=1) {
     (*ast.BinaryExpr)(0x14008478ba0)({
      X: (*ast.BasicLit)(0x14007fff620)({
       ValuePos: (token.Pos) 4235089,
       Kind: (token.Token) STRING,
       Value: (string) (len=12) "\"You said: \""
      }),
      OpPos: (token.Pos) 4235102,
      Op: (token.Token) +,
      Y: (*ast.Ident)(0x14007fff640)(input)
     })
    }
   }),
   (*ast.ReturnStmt)(0x14007fff6a0)({
    Return: (token.Pos) 4235111,
    Results: ([]ast.Expr) (len=1 cap=1) {
     (*ast.Ident)(0x14007fff680)(output)
    }
   })
  },
  Rbrace: (token.Pos) 4235125
 })
})

GML: getFunctionsNeedingWrappers: results.append: info=
(*codegen.funcInfo)(0x14008480d50)({
 function: (*ast.FuncDecl)(0x14008478c00)({
  Doc: (*ast.CommentGroup)(0x14007f9ba88)({
   List: ([]*ast.Comment) (len=1 cap=1) {
    (*ast.Comment)(0x14007f9ba70)({
     Slash: (token.Pos) 4234762,
     Text: (string) (len=58) "// Tests an alternative way to handle errors in functions."
    })
   }
  }),
  Recv: (*ast.FieldList)(<nil>),
  Name: (*ast.Ident)(0x14007fff480)(TestAlternativeError),
  Type: (*ast.FuncType)(0x14007fff6c0)({
   Func: (token.Pos) 4234821,
   TypeParams: (*ast.FieldList)(<nil>),
   Params: (*ast.FieldList)(0x14008478ae0)({
    Opening: (token.Pos) 4234846,
    List: ([]*ast.Field) (len=1 cap=1) {
     (*ast.Field)(0x1400847a700)({
      Doc: (*ast.CommentGroup)(<nil>),
      Names: ([]*ast.Ident) (len=1 cap=1) {
       (*ast.Ident)(0x14007fff4a0)(input)
      },
      Type: (*ast.Ident)(0x14007fff4c0)(string),
      Tag: (*ast.BasicLit)(<nil>),
      Comment: (*ast.CommentGroup)(<nil>)
     })
    },
    Closing: (token.Pos) 4234859
   }),
   Results: (*ast.FieldList)(0x14008478b10)({
    Opening: (token.Pos) 0,
    List: ([]*ast.Field) (len=1 cap=1) {
     (*ast.Field)(0x1400847a740)({
      Doc: (*ast.CommentGroup)(<nil>),
      Names: ([]*ast.Ident) <nil>,
      Type: (*ast.Ident)(0x14007fff4e0)(string),
      Tag: (*ast.BasicLit)(<nil>),
      Comment: (*ast.CommentGroup)(<nil>)
     })
    },
    Closing: (token.Pos) 0
   })
  }),
  Body: (*ast.BlockStmt)(0x14008478bd0)({
   Lbrace: (token.Pos) 4234868,
   List: ([]ast.Stmt) (len=3 cap=4) {
    (*ast.IfStmt)(0x1400847a7c0)({
     If: (token.Pos) 4235012,
     Init: (ast.Stmt) <nil>,
     Cond: (*ast.BinaryExpr)(0x14008478b40)({
      X: (*ast.Ident)(0x14007fff500)(input),
      OpPos: (token.Pos) 4235021,
      Op: (token.Token) ==,
      Y: (*ast.BasicLit)(0x14007fff520)({
       ValuePos: (token.Pos) 4235024,
       Kind: (token.Token) STRING,
       Value: (string) (len=2) "\"\""
      })
     }),
     Body: (*ast.BlockStmt)(0x14008478b70)({
      Lbrace: (token.Pos) 4235027,
      List: ([]ast.Stmt) (len=2 cap=2) {
       (*ast.ExprStmt)(0x14007fc7c50)({
        X: (*ast.CallExpr)(0x1400847a780)({
         Fun: (*ast.SelectorExpr)(0x14007f9bb00)({
          X: (*ast.Ident)(0x14007fff540)(console),
          Sel: (*ast.Ident)(0x14007fff560)(Error)
         }),
         Lparen: (token.Pos) 4235044,
         Args: ([]ast.Expr) (len=1 cap=1) {
          (*ast.BasicLit)(0x14007fff580)({
           ValuePos: (token.Pos) 4235045,
           Kind: (token.Token) STRING,
           Value: (string) (len=16) "\"input is empty\""
          })
         },
         Ellipsis: (token.Pos) 0,
         Rparen: (token.Pos) 4235061
        })
       }),
       (*ast.ReturnStmt)(0x14007fff5c0)({
        Return: (token.Pos) 4235065,
        Results: ([]ast.Expr) (len=1 cap=1) {
         (*ast.BasicLit)(0x14007fff5a0)({
          ValuePos: (token.Pos) 4235072,
          Kind: (token.Token) STRING,
          Value: (string) (len=2) "\"\""
         })
        }
       })
      },
      Rbrace: (token.Pos) 4235076
     }),
     Else: (ast.Stmt) <nil>
    }),
    (*ast.AssignStmt)(0x1400847a800)({
     Lhs: ([]ast.Expr) (len=1 cap=1) {
      (*ast.Ident)(0x14007fff600)(output)
     },
     TokPos: (token.Pos) 4235086,
     Tok: (token.Token) :=,
     Rhs: ([]ast.Expr) (len=1 cap=1) {
      (*ast.BinaryExpr)(0x14008478ba0)({
       X: (*ast.BasicLit)(0x14007fff620)({
        ValuePos: (token.Pos) 4235089,
        Kind: (token.Token) STRING,
        Value: (string) (len=12) "\"You said: \""
       }),
       OpPos: (token.Pos) 4235102,
       Op: (token.Token) +,
       Y: (*ast.Ident)(0x14007fff640)(input)
      })
     }
    }),
    (*ast.ReturnStmt)(0x14007fff6a0)({
     Return: (token.Pos) 4235111,
     Results: ([]ast.Expr) (len=1 cap=1) {
      (*ast.Ident)(0x14007fff680)(output)
     }
    })
   },
   Rbrace: (token.Pos) 4235125
  })
 }),
 imports: (map[string]string) {
 },
 aliases: (map[string]string) <nil>
})

GML: getFunctionsNeedingWrappers: f.Decls[17]=
(*ast.FuncDecl)(0x14008478c90)({
 Doc: (*ast.CommentGroup)(0x14007f9bb30)({
  List: ([]*ast.Comment) (len=1 cap=1) {
   (*ast.Comment)(0x14007f9bb18)({
    Slash: (token.Pos) 4235128,
    Text: (string) (len=17) "// Tests a panic."
   })
  }
 }),
 Recv: (*ast.FieldList)(<nil>),
 Name: (*ast.Ident)(0x14007fff6e0)(TestPanic),
 Type: (*ast.FuncType)(0x14007fff740)({
  Func: (token.Pos) 4235146,
  TypeParams: (*ast.FieldList)(<nil>),
  Params: (*ast.FieldList)(0x14008478c30)({
   Opening: (token.Pos) 4235160,
   List: ([]*ast.Field) <nil>,
   Closing: (token.Pos) 4235161
  }),
  Results: (*ast.FieldList)(<nil>)
 }),
 Body: (*ast.BlockStmt)(0x14008478c60)({
  Lbrace: (token.Pos) 4235163,
  List: ([]ast.Stmt) (len=1 cap=1) {
   (*ast.ExprStmt)(0x14007fc7cf0)({
    X: (*ast.CallExpr)(0x1400847a880)({
     Fun: (*ast.Ident)(0x14007fff700)(panic),
     Lparen: (token.Pos) 4235283,
     Args: ([]ast.Expr) (len=1 cap=1) {
      (*ast.BasicLit)(0x14007fff720)({
       ValuePos: (token.Pos) 4235284,
       Kind: (token.Token) STRING,
       Value: (string) (len=72) "\"This is a message from a panic.\\nThis is a second line from a panic.\\n\""
      })
     },
     Ellipsis: (token.Pos) 0,
     Rparen: (token.Pos) 4235356
    })
   })
  },
  Rbrace: (token.Pos) 4235358
 })
})

GML: getFunctionsNeedingWrappers: results.append: info=
(*codegen.funcInfo)(0x14008480de0)({
 function: (*ast.FuncDecl)(0x14008478c90)({
  Doc: (*ast.CommentGroup)(0x14007f9bb30)({
   List: ([]*ast.Comment) (len=1 cap=1) {
    (*ast.Comment)(0x14007f9bb18)({
     Slash: (token.Pos) 4235128,
     Text: (string) (len=17) "// Tests a panic."
    })
   }
  }),
  Recv: (*ast.FieldList)(<nil>),
  Name: (*ast.Ident)(0x14007fff6e0)(TestPanic),
  Type: (*ast.FuncType)(0x14007fff740)({
   Func: (token.Pos) 4235146,
   TypeParams: (*ast.FieldList)(<nil>),
   Params: (*ast.FieldList)(0x14008478c30)({
    Opening: (token.Pos) 4235160,
    List: ([]*ast.Field) <nil>,
    Closing: (token.Pos) 4235161
   }),
   Results: (*ast.FieldList)(<nil>)
  }),
  Body: (*ast.BlockStmt)(0x14008478c60)({
   Lbrace: (token.Pos) 4235163,
   List: ([]ast.Stmt) (len=1 cap=1) {
    (*ast.ExprStmt)(0x14007fc7cf0)({
     X: (*ast.CallExpr)(0x1400847a880)({
      Fun: (*ast.Ident)(0x14007fff700)(panic),
      Lparen: (token.Pos) 4235283,
      Args: ([]ast.Expr) (len=1 cap=1) {
       (*ast.BasicLit)(0x14007fff720)({
        ValuePos: (token.Pos) 4235284,
        Kind: (token.Token) STRING,
        Value: (string) (len=72) "\"This is a message from a panic.\\nThis is a second line from a panic.\\n\""
       })
      },
      Ellipsis: (token.Pos) 0,
      Rparen: (token.Pos) 4235356
     })
    })
   },
   Rbrace: (token.Pos) 4235358
  })
 }),
 imports: (map[string]string) {
 },
 aliases: (map[string]string) <nil>
})

GML: getFunctionsNeedingWrappers: f.Decls[18]=
(*ast.FuncDecl)(0x14008478d20)({
 Doc: (*ast.CommentGroup)(0x14007f9bba8)({
  List: ([]*ast.Comment) (len=1 cap=1) {
   (*ast.Comment)(0x14007f9bb90)({
    Slash: (token.Pos) 4235361,
    Text: (string) (len=43) "// Tests an exit with a non-zero exit code."
   })
  }
 }),
 Recv: (*ast.FieldList)(<nil>),
 Name: (*ast.Ident)(0x14007fff760)(TestExit),
 Type: (*ast.FuncType)(0x14007fff8c0)({
  Func: (token.Pos) 4235405,
  TypeParams: (*ast.FieldList)(<nil>),
  Params: (*ast.FieldList)(0x14008478cc0)({
   Opening: (token.Pos) 4235418,
   List: ([]*ast.Field) <nil>,
   Closing: (token.Pos) 4235419
  }),
  Results: (*ast.FieldList)(<nil>)
 }),
 Body: (*ast.BlockStmt)(0x14008478cf0)({
  Lbrace: (token.Pos) 4235421,
  List: ([]ast.Stmt) (len=3 cap=4) {
   (*ast.ExprStmt)(0x14007fc7d40)({
    X: (*ast.CallExpr)(0x1400847a8c0)({
     Fun: (*ast.SelectorExpr)(0x14007f9bc38)({
      X: (*ast.Ident)(0x14007fff7a0)(console),
      Sel: (*ast.Ident)(0x14007fff7c0)(Error)
     }),
     Lparen: (token.Pos) 4235730,
     Args: ([]ast.Expr) (len=1 cap=1) {
      (*ast.BasicLit)(0x14007fff7e0)({
       ValuePos: (token.Pos) 4235731,
       Kind: (token.Token) STRING,
       Value: (string) (len=27) "\"This is an error message.\""
      })
     },
     Ellipsis: (token.Pos) 0,
     Rparen: (token.Pos) 4235758
    })
   }),
   (*ast.ExprStmt)(0x14007fc7d80)({
    X: (*ast.CallExpr)(0x1400847a900)({
     Fun: (*ast.SelectorExpr)(0x14007f9bc50)({
      X: (*ast.Ident)(0x14007fff800)(os),
      Sel: (*ast.Ident)(0x14007fff820)(Exit)
     }),
     Lparen: (token.Pos) 4235768,
     Args: ([]ast.Expr) (len=1 cap=1) {
      (*ast.BasicLit)(0x14007fff840)({
       ValuePos: (token.Pos) 4235769,
       Kind: (token.Token) INT,
       Value: (string) (len=1) "1"
      })
     },
     Ellipsis: (token.Pos) 0,
     Rparen: (token.Pos) 4235770
    })
   }),
   (*ast.ExprStmt)(0x14007fc7db0)({
    X: (*ast.CallExpr)(0x1400847a940)({
     Fun: (*ast.Ident)(0x14007fff880)(println),
     Lparen: (token.Pos) 4235780,
     Args: ([]ast.Expr) (len=1 cap=1) {
      (*ast.BasicLit)(0x14007fff8a0)({
       ValuePos: (token.Pos) 4235781,
       Kind: (token.Token) STRING,
       Value: (string) (len=33) "\"This line will not be executed.\""
      })
     },
     Ellipsis: (token.Pos) 0,
     Rparen: (token.Pos) 4235814
    })
   })
  },
  Rbrace: (token.Pos) 4235816
 })
})

GML: getFunctionsNeedingWrappers: results.append: info=
(*codegen.funcInfo)(0x14008480e58)({
 function: (*ast.FuncDecl)(0x14008478d20)({
  Doc: (*ast.CommentGroup)(0x14007f9bba8)({
   List: ([]*ast.Comment) (len=1 cap=1) {
    (*ast.Comment)(0x14007f9bb90)({
     Slash: (token.Pos) 4235361,
     Text: (string) (len=43) "// Tests an exit with a non-zero exit code."
    })
   }
  }),
  Recv: (*ast.FieldList)(<nil>),
  Name: (*ast.Ident)(0x14007fff760)(TestExit),
  Type: (*ast.FuncType)(0x14007fff8c0)({
   Func: (token.Pos) 4235405,
   TypeParams: (*ast.FieldList)(<nil>),
   Params: (*ast.FieldList)(0x14008478cc0)({
    Opening: (token.Pos) 4235418,
    List: ([]*ast.Field) <nil>,
    Closing: (token.Pos) 4235419
   }),
   Results: (*ast.FieldList)(<nil>)
  }),
  Body: (*ast.BlockStmt)(0x14008478cf0)({
   Lbrace: (token.Pos) 4235421,
   List: ([]ast.Stmt) (len=3 cap=4) {
    (*ast.ExprStmt)(0x14007fc7d40)({
     X: (*ast.CallExpr)(0x1400847a8c0)({
      Fun: (*ast.SelectorExpr)(0x14007f9bc38)({
       X: (*ast.Ident)(0x14007fff7a0)(console),
       Sel: (*ast.Ident)(0x14007fff7c0)(Error)
      }),
      Lparen: (token.Pos) 4235730,
      Args: ([]ast.Expr) (len=1 cap=1) {
       (*ast.BasicLit)(0x14007fff7e0)({
        ValuePos: (token.Pos) 4235731,
        Kind: (token.Token) STRING,
        Value: (string) (len=27) "\"This is an error message.\""
       })
      },
      Ellipsis: (token.Pos) 0,
      Rparen: (token.Pos) 4235758
     })
    }),
    (*ast.ExprStmt)(0x14007fc7d80)({
     X: (*ast.CallExpr)(0x1400847a900)({
      Fun: (*ast.SelectorExpr)(0x14007f9bc50)({
       X: (*ast.Ident)(0x14007fff800)(os),
       Sel: (*ast.Ident)(0x14007fff820)(Exit)
      }),
      Lparen: (token.Pos) 4235768,
      Args: ([]ast.Expr) (len=1 cap=1) {
       (*ast.BasicLit)(0x14007fff840)({
        ValuePos: (token.Pos) 4235769,
        Kind: (token.Token) INT,
        Value: (string) (len=1) "1"
       })
      },
      Ellipsis: (token.Pos) 0,
      Rparen: (token.Pos) 4235770
     })
    }),
    (*ast.ExprStmt)(0x14007fc7db0)({
     X: (*ast.CallExpr)(0x1400847a940)({
      Fun: (*ast.Ident)(0x14007fff880)(println),
      Lparen: (token.Pos) 4235780,
      Args: ([]ast.Expr) (len=1 cap=1) {
       (*ast.BasicLit)(0x14007fff8a0)({
        ValuePos: (token.Pos) 4235781,
        Kind: (token.Token) STRING,
        Value: (string) (len=33) "\"This line will not be executed.\""
       })
      },
      Ellipsis: (token.Pos) 0,
      Rparen: (token.Pos) 4235814
     })
    })
   },
   Rbrace: (token.Pos) 4235816
  })
 }),
 imports: (map[string]string) {
 },
 aliases: (map[string]string) <nil>
})

GML: getFunctionsNeedingWrappers: f.Decls[19]=
(*ast.FuncDecl)(0x14008478db0)({
 Doc: (*ast.CommentGroup)(0x14007f9bc80)({
  List: ([]*ast.Comment) (len=1 cap=1) {
   (*ast.Comment)(0x14007f9bc68)({
    Slash: (token.Pos) 4235819,
    Text: (string) (len=37) "// Tests logging at different levels."
   })
  }
 }),
 Recv: (*ast.FieldList)(<nil>),
 Name: (*ast.Ident)(0x14007fff8e0)(TestLogging),
 Type: (*ast.FuncType)(0x14007fffe40)({
  Func: (token.Pos) 4235857,
  TypeParams: (*ast.FieldList)(<nil>),
  Params: (*ast.FieldList)(0x14008478d50)({
   Opening: (token.Pos) 4235873,
   List: ([]*ast.Field) <nil>,
   Closing: (token.Pos) 4235874
  }),
  Results: (*ast.FieldList)(<nil>)
 }),
 Body: (*ast.BlockStmt)(0x14008478d80)({
  Lbrace: (token.Pos) 4235876,
  List: ([]ast.Stmt) (len=11 cap=16) {
   (*ast.ExprStmt)(0x14007fc7de0)({
    X: (*ast.CallExpr)(0x1400847a9c0)({
     Fun: (*ast.SelectorExpr)(0x14007f9bcc8)({
      X: (*ast.Ident)(0x14007fff900)(console),
      Sel: (*ast.Ident)(0x14007fff920)(Log)
     }),
     Lparen: (token.Pos) 4235941,
     Args: ([]ast.Expr) (len=1 cap=1) {
      (*ast.BasicLit)(0x14007fff940)({
       ValuePos: (token.Pos) 4235942,
       Kind: (token.Token) STRING,
       Value: (string) (len=31) "\"This is a simple log message.\""
      })
     },
     Ellipsis: (token.Pos) 0,
     Rparen: (token.Pos) 4235973
    })
   }),
   (*ast.ExprStmt)(0x14007fc7e20)({
    X: (*ast.CallExpr)(0x1400847aa00)({
     Fun: (*ast.SelectorExpr)(0x14007f9bd10)({
      X: (*ast.Ident)(0x14007fff960)(console),
      Sel: (*ast.Ident)(0x14007fff980)(Debug)
     }),
     Lparen: (token.Pos) 4236041,
     Args: ([]ast.Expr) (len=1 cap=1) {
      (*ast.BasicLit)(0x14007fff9a0)({
       ValuePos: (token.Pos) 4236042,
       Kind: (token.Token) STRING,
       Value: (string) (len=26) "\"This is a debug message.\""
      })
     },
     Ellipsis: (token.Pos) 0,
     Rparen: (token.Pos) 4236068
    })
   }),
   (*ast.ExprStmt)(0x14007fc7e50)({
    X: (*ast.CallExpr)(0x1400847aa40)({
     Fun: (*ast.SelectorExpr)(0x14007f9bd28)({
      X: (*ast.Ident)(0x14007fff9e0)(console),
      Sel: (*ast.Ident)(0x14007fffa00)(Info)
     }),
     Lparen: (token.Pos) 4236083,
     Args: ([]ast.Expr) (len=1 cap=1) {
      (*ast.BasicLit)(0x14007fffa20)({
       ValuePos: (token.Pos) 4236084,
       Kind: (token.Token) STRING,
       Value: (string) (len=26) "\"This is an info message.\""
      })
     },
     Ellipsis: (token.Pos) 0,
     Rparen: (token.Pos) 4236110
    })
   }),
   (*ast.ExprStmt)(0x14007fc7e80)({
    X: (*ast.CallExpr)(0x1400847aac0)({
     Fun: (*ast.SelectorExpr)(0x14007f9bd40)({
      X: (*ast.Ident)(0x14007fffa40)(console),
      Sel: (*ast.Ident)(0x14007fffa60)(Warn)
     }),
     Lparen: (token.Pos) 4236125,
     Args: ([]ast.Expr) (len=1 cap=1) {
      (*ast.BasicLit)(0x14007fffa80)({
       ValuePos: (token.Pos) 4236126,
       Kind: (token.Token) STRING,
       Value: (string) (len=28) "\"This is a warning message.\""
      })
     },
     Ellipsis: (token.Pos) 0,
     Rparen: (token.Pos) 4236154
    })
   }),
   (*ast.ExprStmt)(0x14007fc7eb0)({
    X: (*ast.CallExpr)(0x1400847ab00)({
     Fun: (*ast.SelectorExpr)(0x14007f9bd88)({
      X: (*ast.Ident)(0x14007fffaa0)(console),
      Sel: (*ast.Ident)(0x14007fffac0)(Error)
     }),
     Lparen: (token.Pos) 4236240,
     Args: ([]ast.Expr) (len=1 cap=1) {
      (*ast.BasicLit)(0x14007fffae0)({
       ValuePos: (token.Pos) 4236241,
       Kind: (token.Token) STRING,
       Value: (string) (len=27) "\"This is an error message.\""
      })
     },
     Ellipsis: (token.Pos) 0,
     Rparen: (token.Pos) 4236268
    })
   }),
   (*ast.ExprStmt)(0x14007fc7ee0)({
    X: (*ast.CallExpr)(0x1400847ab40)({
     Fun: (*ast.SelectorExpr)(0x14007f9bda0)({
      X: (*ast.Ident)(0x14007fffb00)(console),
      Sel: (*ast.Ident)(0x14007fffb20)(Error)
     }),
     Lparen: (token.Pos) 4236284,
     Args: ([]ast.Expr) (len=1 cap=1) {
      (*ast.BasicLit)(0x14007fffb40)({
       ValuePos: (token.Pos) 4236288,
       Kind: (token.Token) STRING,
       Value: (string) (len=143) "`This is line 1 of a multi-line error message.\n  This is line 2 of a multi-line error message.\n  This is line 3 of a multi-line error message.`"
      })
     },
     Ellipsis: (token.Pos) 0,
     Rparen: (token.Pos) 4236434
    })
   }),
   (*ast.ExprStmt)(0x14007fc7f10)({
    X: (*ast.CallExpr)(0x1400847ab80)({
     Fun: (*ast.Ident)(0x14007fffb60)(println),
     Lparen: (token.Pos) 4236499,
     Args: ([]ast.Expr) (len=1 cap=1) {
      (*ast.BasicLit)(0x14007fffb80)({
       ValuePos: (token.Pos) 4236500,
       Kind: (token.Token) STRING,
       Value: (string) (len=28) "\"This is a println message.\""
      })
     },
     Ellipsis: (token.Pos) 0,
     Rparen: (token.Pos) 4236528
    })
   }),
   (*ast.ExprStmt)(0x14007fc7f40)({
    X: (*ast.CallExpr)(0x1400847abc0)({
     Fun: (*ast.SelectorExpr)(0x14007f9bde8)({
      X: (*ast.Ident)(0x14007fffba0)(fmt),
      Sel: (*ast.Ident)(0x14007fffbc0)(Println)
     }),
     Lparen: (token.Pos) 4236542,
     Args: ([]ast.Expr) (len=1 cap=1) {
      (*ast.BasicLit)(0x14007fffbe0)({
       ValuePos: (token.Pos) 4236543,
       Kind: (token.Token) STRING,
       Value: (string) (len=32) "\"This is a fmt.Println message.\""
      })
     },
     Ellipsis: (token.Pos) 0,
     Rparen: (token.Pos) 4236575
    })
   }),
   (*ast.ExprStmt)(0x14007fc7f70)({
    X: (*ast.CallExpr)(0x1400847ac00)({
     Fun: (*ast.SelectorExpr)(0x14007f9be00)({
      X: (*ast.Ident)(0x14007fffc00)(fmt),
      Sel: (*ast.Ident)(0x14007fffc20)(Printf)
     }),
     Lparen: (token.Pos) 4236588,
     Args: ([]ast.Expr) (len=2 cap=2) {
      (*ast.BasicLit)(0x14007fffc40)({
       ValuePos: (token.Pos) 4236589,
       Kind: (token.Token) STRING,
       Value: (string) (len=38) "\"This is a fmt.Printf message (%s).\\n\""
      }),
      (*ast.BasicLit)(0x14007fffc60)({
       ValuePos: (token.Pos) 4236629,
       Kind: (token.Token) STRING,
       Value: (string) (len=18) "\"with a parameter\""
      })
     },
     Ellipsis: (token.Pos) 0,
     Rparen: (token.Pos) 4236647
    })
   }),
   (*ast.ExprStmt)(0x1400847e020)({
    X: (*ast.CallExpr)(0x1400847ac40)({
     Fun: (*ast.SelectorExpr)(0x14007f9be48)({
      X: (*ast.Ident)(0x14007fffca0)(fmt),
      Sel: (*ast.Ident)(0x14007fffcc0)(Fprintln)
     }),
     Lparen: (token.Pos) 4236761,
     Args: ([]ast.Expr) (len=2 cap=2) {
      (*ast.SelectorExpr)(0x14007f9be60)({
       X: (*ast.Ident)(0x14007fffce0)(os),
       Sel: (*ast.Ident)(0x14007fffd00)(Stdout)
      }),
      (*ast.BasicLit)(0x14007fffd20)({
       ValuePos: (token.Pos) 4236773,
       Kind: (token.Token) STRING,
       Value: (string) (len=44) "\"This is an info message printed to stdout.\""
      })
     },
     Ellipsis: (token.Pos) 0,
     Rparen: (token.Pos) 4236817
    })
   }),
   (*ast.ExprStmt)(0x1400847e050)({
    X: (*ast.CallExpr)(0x1400847ac80)({
     Fun: (*ast.SelectorExpr)(0x14007f9be78)({
      X: (*ast.Ident)(0x14007fffd60)(fmt),
      Sel: (*ast.Ident)(0x14007fffd80)(Fprintln)
     }),
     Lparen: (token.Pos) 4236832,
     Args: ([]ast.Expr) (len=2 cap=2) {
      (*ast.SelectorExpr)(0x14007f9be90)({
       X: (*ast.Ident)(0x14007fffda0)(os),
       Sel: (*ast.Ident)(0x14007fffdc0)(Stderr)
      }),
      (*ast.BasicLit)(0x14007fffde0)({
       ValuePos: (token.Pos) 4236844,
       Kind: (token.Token) STRING,
       Value: (string) (len=45) "\"This is an error message printed to stderr.\""
      })
     },
     Ellipsis: (token.Pos) 0,
     Rparen: (token.Pos) 4236889
    })
   })
  },
  Rbrace: (token.Pos) 4237450
 })
})

GML: getFunctionsNeedingWrappers: results.append: info=
(*codegen.funcInfo)(0x14008480ed0)({
 function: (*ast.FuncDecl)(0x14008478db0)({
  Doc: (*ast.CommentGroup)(0x14007f9bc80)({
   List: ([]*ast.Comment) (len=1 cap=1) {
    (*ast.Comment)(0x14007f9bc68)({
     Slash: (token.Pos) 4235819,
     Text: (string) (len=37) "// Tests logging at different levels."
    })
   }
  }),
  Recv: (*ast.FieldList)(<nil>),
  Name: (*ast.Ident)(0x14007fff8e0)(TestLogging),
  Type: (*ast.FuncType)(0x14007fffe40)({
   Func: (token.Pos) 4235857,
   TypeParams: (*ast.FieldList)(<nil>),
   Params: (*ast.FieldList)(0x14008478d50)({
    Opening: (token.Pos) 4235873,
    List: ([]*ast.Field) <nil>,
    Closing: (token.Pos) 4235874
   }),
   Results: (*ast.FieldList)(<nil>)
  }),
  Body: (*ast.BlockStmt)(0x14008478d80)({
   Lbrace: (token.Pos) 4235876,
   List: ([]ast.Stmt) (len=11 cap=16) {
    (*ast.ExprStmt)(0x14007fc7de0)({
     X: (*ast.CallExpr)(0x1400847a9c0)({
      Fun: (*ast.SelectorExpr)(0x14007f9bcc8)({
       X: (*ast.Ident)(0x14007fff900)(console),
       Sel: (*ast.Ident)(0x14007fff920)(Log)
      }),
      Lparen: (token.Pos) 4235941,
      Args: ([]ast.Expr) (len=1 cap=1) {
       (*ast.BasicLit)(0x14007fff940)({
        ValuePos: (token.Pos) 4235942,
        Kind: (token.Token) STRING,
        Value: (string) (len=31) "\"This is a simple log message.\""
       })
      },
      Ellipsis: (token.Pos) 0,
      Rparen: (token.Pos) 4235973
     })
    }),
    (*ast.ExprStmt)(0x14007fc7e20)({
     X: (*ast.CallExpr)(0x1400847aa00)({
      Fun: (*ast.SelectorExpr)(0x14007f9bd10)({
       X: (*ast.Ident)(0x14007fff960)(console),
       Sel: (*ast.Ident)(0x14007fff980)(Debug)
      }),
      Lparen: (token.Pos) 4236041,
      Args: ([]ast.Expr) (len=1 cap=1) {
       (*ast.BasicLit)(0x14007fff9a0)({
        ValuePos: (token.Pos) 4236042,
        Kind: (token.Token) STRING,
        Value: (string) (len=26) "\"This is a debug message.\""
       })
      },
      Ellipsis: (token.Pos) 0,
      Rparen: (token.Pos) 4236068
     })
    }),
    (*ast.ExprStmt)(0x14007fc7e50)({
     X: (*ast.CallExpr)(0x1400847aa40)({
      Fun: (*ast.SelectorExpr)(0x14007f9bd28)({
       X: (*ast.Ident)(0x14007fff9e0)(console),
       Sel: (*ast.Ident)(0x14007fffa00)(Info)
      }),
      Lparen: (token.Pos) 4236083,
      Args: ([]ast.Expr) (len=1 cap=1) {
       (*ast.BasicLit)(0x14007fffa20)({
        ValuePos: (token.Pos) 4236084,
        Kind: (token.Token) STRING,
        Value: (string) (len=26) "\"This is an info message.\""
       })
      },
      Ellipsis: (token.Pos) 0,
      Rparen: (token.Pos) 4236110
     })
    }),
    (*ast.ExprStmt)(0x14007fc7e80)({
     X: (*ast.CallExpr)(0x1400847aac0)({
      Fun: (*ast.SelectorExpr)(0x14007f9bd40)({
       X: (*ast.Ident)(0x14007fffa40)(console),
       Sel: (*ast.Ident)(0x14007fffa60)(Warn)
      }),
      Lparen: (token.Pos) 4236125,
      Args: ([]ast.Expr) (len=1 cap=1) {
       (*ast.BasicLit)(0x14007fffa80)({
        ValuePos: (token.Pos) 4236126,
        Kind: (token.Token) STRING,
        Value: (string) (len=28) "\"This is a warning message.\""
       })
      },
      Ellipsis: (token.Pos) 0,
      Rparen: (token.Pos) 4236154
     })
    }),
    (*ast.ExprStmt)(0x14007fc7eb0)({
     X: (*ast.CallExpr)(0x1400847ab00)({
      Fun: (*ast.SelectorExpr)(0x14007f9bd88)({
       X: (*ast.Ident)(0x14007fffaa0)(console),
       Sel: (*ast.Ident)(0x14007fffac0)(Error)
      }),
      Lparen: (token.Pos) 4236240,
      Args: ([]ast.Expr) (len=1 cap=1) {
       (*ast.BasicLit)(0x14007fffae0)({
        ValuePos: (token.Pos) 4236241,
        Kind: (token.Token) STRING,
        Value: (string) (len=27) "\"This is an error message.\""
       })
      },
      Ellipsis: (token.Pos) 0,
      Rparen: (token.Pos) 4236268
     })
    }),
    (*ast.ExprStmt)(0x14007fc7ee0)({
     X: (*ast.CallExpr)(0x1400847ab40)({
      Fun: (*ast.SelectorExpr)(0x14007f9bda0)({
       X: (*ast.Ident)(0x14007fffb00)(console),
       Sel: (*ast.Ident)(0x14007fffb20)(Error)
      }),
      Lparen: (token.Pos) 4236284,
      Args: ([]ast.Expr) (len=1 cap=1) {
       (*ast.BasicLit)(0x14007fffb40)({
        ValuePos: (token.Pos) 4236288,
        Kind: (token.Token) STRING,
        Value: (string) (len=143) "`This is line 1 of a multi-line error message.\n  This is line 2 of a multi-line error message.\n  This is line 3 of a multi-line error message.`"
       })
      },
      Ellipsis: (token.Pos) 0,
      Rparen: (token.Pos) 4236434
     })
    }),
    (*ast.ExprStmt)(0x14007fc7f10)({
     X: (*ast.CallExpr)(0x1400847ab80)({
      Fun: (*ast.Ident)(0x14007fffb60)(println),
      Lparen: (token.Pos) 4236499,
      Args: ([]ast.Expr) (len=1 cap=1) {
       (*ast.BasicLit)(0x14007fffb80)({
        ValuePos: (token.Pos) 4236500,
        Kind: (token.Token) STRING,
        Value: (string) (len=28) "\"This is a println message.\""
       })
      },
      Ellipsis: (token.Pos) 0,
      Rparen: (token.Pos) 4236528
     })
    }),
    (*ast.ExprStmt)(0x14007fc7f40)({
     X: (*ast.CallExpr)(0x1400847abc0)({
      Fun: (*ast.SelectorExpr)(0x14007f9bde8)({
       X: (*ast.Ident)(0x14007fffba0)(fmt),
       Sel: (*ast.Ident)(0x14007fffbc0)(Println)
      }),
      Lparen: (token.Pos) 4236542,
      Args: ([]ast.Expr) (len=1 cap=1) {
       (*ast.BasicLit)(0x14007fffbe0)({
        ValuePos: (token.Pos) 4236543,
        Kind: (token.Token) STRING,
        Value: (string) (len=32) "\"This is a fmt.Println message.\""
       })
      },
      Ellipsis: (token.Pos) 0,
      Rparen: (token.Pos) 4236575
     })
    }),
    (*ast.ExprStmt)(0x14007fc7f70)({
     X: (*ast.CallExpr)(0x1400847ac00)({
      Fun: (*ast.SelectorExpr)(0x14007f9be00)({
       X: (*ast.Ident)(0x14007fffc00)(fmt),
       Sel: (*ast.Ident)(0x14007fffc20)(Printf)
      }),
      Lparen: (token.Pos) 4236588,
      Args: ([]ast.Expr) (len=2 cap=2) {
       (*ast.BasicLit)(0x14007fffc40)({
        ValuePos: (token.Pos) 4236589,
        Kind: (token.Token) STRING,
        Value: (string) (len=38) "\"This is a fmt.Printf message (%s).\\n\""
       }),
       (*ast.BasicLit)(0x14007fffc60)({
        ValuePos: (token.Pos) 4236629,
        Kind: (token.Token) STRING,
        Value: (string) (len=18) "\"with a parameter\""
       })
      },
      Ellipsis: (token.Pos) 0,
      Rparen: (token.Pos) 4236647
     })
    }),
    (*ast.ExprStmt)(0x1400847e020)({
     X: (*ast.CallExpr)(0x1400847ac40)({
      Fun: (*ast.SelectorExpr)(0x14007f9be48)({
       X: (*ast.Ident)(0x14007fffca0)(fmt),
       Sel: (*ast.Ident)(0x14007fffcc0)(Fprintln)
      }),
      Lparen: (token.Pos) 4236761,
      Args: ([]ast.Expr) (len=2 cap=2) {
       (*ast.SelectorExpr)(0x14007f9be60)({
        X: (*ast.Ident)(0x14007fffce0)(os),
        Sel: (*ast.Ident)(0x14007fffd00)(Stdout)
       }),
       (*ast.BasicLit)(0x14007fffd20)({
        ValuePos: (token.Pos) 4236773,
        Kind: (token.Token) STRING,
        Value: (string) (len=44) "\"This is an info message printed to stdout.\""
       })
      },
      Ellipsis: (token.Pos) 0,
      Rparen: (token.Pos) 4236817
     })
    }),
    (*ast.ExprStmt)(0x1400847e050)({
     X: (*ast.CallExpr)(0x1400847ac80)({
      Fun: (*ast.SelectorExpr)(0x14007f9be78)({
       X: (*ast.Ident)(0x14007fffd60)(fmt),
       Sel: (*ast.Ident)(0x14007fffd80)(Fprintln)
      }),
      Lparen: (token.Pos) 4236832,
      Args: ([]ast.Expr) (len=2 cap=2) {
       (*ast.SelectorExpr)(0x14007f9be90)({
        X: (*ast.Ident)(0x14007fffda0)(os),
        Sel: (*ast.Ident)(0x14007fffdc0)(Stderr)
       }),
       (*ast.BasicLit)(0x14007fffde0)({
        ValuePos: (token.Pos) 4236844,
        Kind: (token.Token) STRING,
        Value: (string) (len=45) "\"This is an error message printed to stderr.\""
       })
      },
      Ellipsis: (token.Pos) 0,
      Rparen: (token.Pos) 4236889
     })
    })
   },
   Rbrace: (token.Pos) 4237450
  })
 }),
 imports: (map[string]string) {
 },
 aliases: (map[string]string) <nil>
})

GML: PreProcess: functions=
([]*codegen.funcInfo) (len=17 cap=32) {
 (*codegen.funcInfo)(0x14008480618)({
  function: (*ast.FuncDecl)(0x14007fdda10)({
   Doc: (*ast.CommentGroup)(0x14007f9b410)({
    List: ([]*ast.Comment) (len=1 cap=1) {
     (*ast.Comment)(0x14007f9b3f8)({
      Slash: (token.Pos) 4232050,
      Text: (string) (len=18) "// Logs a message."
     })
    }
   }),
   Recv: (*ast.FieldList)(<nil>),
   Name: (*ast.Ident)(0x14007fa3b20)(LogMessage),
   Type: (*ast.FuncType)(0x14007fa3c00)({
    Func: (token.Pos) 4232069,
    TypeParams: (*ast.FieldList)(<nil>),
    Params: (*ast.FieldList)(0x14007fdd9b0)({
     Opening: (token.Pos) 4232084,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x14007fe7680)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) (len=1 cap=1) {
        (*ast.Ident)(0x14007fa3b40)(message)
       },
       Type: (*ast.Ident)(0x14007fa3b60)(string),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      })
     },
     Closing: (token.Pos) 4232099
    }),
    Results: (*ast.FieldList)(<nil>)
   }),
   Body: (*ast.BlockStmt)(0x14007fdd9e0)({
    Lbrace: (token.Pos) 4232101,
    List: ([]ast.Stmt) (len=1 cap=1) {
     (*ast.ExprStmt)(0x14007fc7730)({
      X: (*ast.CallExpr)(0x14007fe76c0)({
       Fun: (*ast.SelectorExpr)(0x14007f9b440)({
        X: (*ast.Ident)(0x14007fa3b80)(console),
        Sel: (*ast.Ident)(0x14007fa3ba0)(Log)
       }),
       Lparen: (token.Pos) 4232115,
       Args: ([]ast.Expr) (len=1 cap=1) {
        (*ast.Ident)(0x14007fa3bc0)(message)
       },
       Ellipsis: (token.Pos) 0,
       Rparen: (token.Pos) 4232123
      })
     })
    },
    Rbrace: (token.Pos) 4232125
   })
  }),
  imports: (map[string]string) {
  },
  aliases: (map[string]string) <nil>
 }),
 (*codegen.funcInfo)(0x14008480690)({
  function: (*ast.FuncDecl)(0x14007fddb30)({
   Doc: (*ast.CommentGroup)(0x14007f9b470)({
    List: ([]*ast.Comment) (len=1 cap=1) {
     (*ast.Comment)(0x14007f9b458)({
      Slash: (token.Pos) 4232128,
      Text: (string) (len=53) "// Adds two integers together and returns the result."
     })
    }
   }),
   Recv: (*ast.FieldList)(<nil>),
   Name: (*ast.Ident)(0x14007fa3c40)(Add),
   Type: (*ast.FuncType)(0x14007fa3d40)({
    Func: (token.Pos) 4232182,
    TypeParams: (*ast.FieldList)(<nil>),
    Params: (*ast.FieldList)(0x14007fdda70)({
     Opening: (token.Pos) 4232190,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x14007fe7700)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) (len=2 cap=2) {
        (*ast.Ident)(0x14007fa3c60)(x),
        (*ast.Ident)(0x14007fa3c80)(y)
       },
       Type: (*ast.Ident)(0x14007fa3ca0)(int),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      })
     },
     Closing: (token.Pos) 4232199
    }),
    Results: (*ast.FieldList)(0x14007fddaa0)({
     Opening: (token.Pos) 0,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x14007fe7740)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) <nil>,
       Type: (*ast.Ident)(0x14007fa3cc0)(int),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      })
     },
     Closing: (token.Pos) 0
    })
   }),
   Body: (*ast.BlockStmt)(0x14007fddb00)({
    Lbrace: (token.Pos) 4232205,
    List: ([]ast.Stmt) (len=1 cap=1) {
     (*ast.ReturnStmt)(0x14007fa3d20)({
      Return: (token.Pos) 4232208,
      Results: ([]ast.Expr) (len=1 cap=1) {
       (*ast.BinaryExpr)(0x14007fddad0)({
        X: (*ast.Ident)(0x14007fa3ce0)(x),
        OpPos: (token.Pos) 4232217,
        Op: (token.Token) +,
        Y: (*ast.Ident)(0x14007fa3d00)(y)
       })
      }
     })
    },
    Rbrace: (token.Pos) 4232221
   })
  }),
  imports: (map[string]string) {
  },
  aliases: (map[string]string) <nil>
 }),
 (*codegen.funcInfo)(0x14008480708)({
  function: (*ast.FuncDecl)(0x14007fddd10)({
   Doc: (*ast.CommentGroup)(0x14007f9b4d0)({
    List: ([]*ast.Comment) (len=2 cap=2) {
     (*ast.Comment)(0x14007f9b4a0)({
      Slash: (token.Pos) 4232224,
      Text: (string) (len=55) "// Adds three integers together and returns the result."
     }),
     (*ast.Comment)(0x14007f9b4b8)({
      Slash: (token.Pos) 4232280,
      Text: (string) (len=33) "// The third integer is optional."
     })
    }
   }),
   Recv: (*ast.FieldList)(<nil>),
   Name: (*ast.Ident)(0x14007fa3d60)(Add3),
   Type: (*ast.FuncType)(0x14007ffe000)({
    Func: (token.Pos) 4232314,
    TypeParams: (*ast.FieldList)(<nil>),
    Params: (*ast.FieldList)(0x14007fddb90)({
     Opening: (token.Pos) 4232323,
     List: ([]*ast.Field) (len=2 cap=2) {
      (*ast.Field)(0x14007fe77c0)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) (len=2 cap=2) {
        (*ast.Ident)(0x14007fa3d80)(a),
        (*ast.Ident)(0x14007fa3da0)(b)
       },
       Type: (*ast.Ident)(0x14007fa3dc0)(int),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      }),
      (*ast.Field)(0x14007fe7800)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) (len=1 cap=1) {
        (*ast.Ident)(0x14007fa3de0)(c)
       },
       Type: (*ast.StarExpr)(0x14007f9b500)({
        Star: (token.Pos) 4232336,
        X: (*ast.Ident)(0x14007fa3e00)(int)
       }),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      })
     },
     Closing: (token.Pos) 4232340
    }),
    Results: (*ast.FieldList)(0x14007fddbc0)({
     Opening: (token.Pos) 0,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x14007fe7840)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) <nil>,
       Type: (*ast.Ident)(0x14007fa3e20)(int),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      })
     },
     Closing: (token.Pos) 0
    })
   }),
   Body: (*ast.BlockStmt)(0x14007fddce0)({
    Lbrace: (token.Pos) 4232346,
    List: ([]ast.Stmt) (len=2 cap=2) {
     (*ast.IfStmt)(0x14007fe7880)({
      If: (token.Pos) 4232349,
      Init: (ast.Stmt) <nil>,
      Cond: (*ast.BinaryExpr)(0x14007fddbf0)({
       X: (*ast.Ident)(0x14007fa3e40)(c),
       OpPos: (token.Pos) 4232354,
       Op: (token.Token) !=,
       Y: (*ast.Ident)(0x14007fa3e60)(nil)
      }),
      Body: (*ast.BlockStmt)(0x14007fddc80)({
       Lbrace: (token.Pos) 4232361,
       List: ([]ast.Stmt) (len=1 cap=1) {
        (*ast.ReturnStmt)(0x14007fa3ee0)({
         Return: (token.Pos) 4232365,
         Results: ([]ast.Expr) (len=1 cap=1) {
          (*ast.BinaryExpr)(0x14007fddc50)({
           X: (*ast.BinaryExpr)(0x14007fddc20)({
            X: (*ast.Ident)(0x14007fa3e80)(a),
            OpPos: (token.Pos) 4232374,
            Op: (token.Token) +,
            Y: (*ast.Ident)(0x14007fa3ea0)(b)
           }),
           OpPos: (token.Pos) 4232378,
           Op: (token.Token) +,
           Y: (*ast.StarExpr)(0x14007f9b518)({
            Star: (token.Pos) 4232380,
            X: (*ast.Ident)(0x14007fa3ec0)(c)
           })
          })
         }
        })
       },
       Rbrace: (token.Pos) 4232384
      }),
      Else: (ast.Stmt) <nil>
     }),
     (*ast.ReturnStmt)(0x14007fa3f40)({
      Return: (token.Pos) 4232387,
      Results: ([]ast.Expr) (len=1 cap=1) {
       (*ast.BinaryExpr)(0x14007fddcb0)({
        X: (*ast.Ident)(0x14007fa3f00)(a),
        OpPos: (token.Pos) 4232396,
        Op: (token.Token) +,
        Y: (*ast.Ident)(0x14007fa3f20)(b)
       })
      }
     })
    },
    Rbrace: (token.Pos) 4232400
   })
  }),
  imports: (map[string]string) {
  },
  aliases: (map[string]string) <nil>
 }),
 (*codegen.funcInfo)(0x14008480798)({
  function: (*ast.FuncDecl)(0x14007fdde00)({
   Doc: (*ast.CommentGroup)(0x14007f9b548)({
    List: ([]*ast.Comment) (len=1 cap=1) {
     (*ast.Comment)(0x14007f9b530)({
      Slash: (token.Pos) 4232403,
      Text: (string) (len=63) "// Adds any number of integers together and returns the result."
     })
    }
   }),
   Recv: (*ast.FieldList)(<nil>),
   Name: (*ast.Ident)(0x14007ffe020)(AddN),
   Type: (*ast.FuncType)(0x14007ffe220)({
    Func: (token.Pos) 4232467,
    TypeParams: (*ast.FieldList)(<nil>),
    Params: (*ast.FieldList)(0x14007fddd40)({
     Opening: (token.Pos) 4232476,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x14007fe7900)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) (len=1 cap=1) {
        (*ast.Ident)(0x14007ffe040)(args)
       },
       Type: (*ast.Ellipsis)(0x14007f9b560)({
        Ellipsis: (token.Pos) 4232482,
        Elt: (*ast.Ident)(0x14007ffe060)(int)
       }),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      })
     },
     Closing: (token.Pos) 4232488
    }),
    Results: (*ast.FieldList)(0x14007fddd70)({
     Opening: (token.Pos) 0,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x14007fe7940)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) <nil>,
       Type: (*ast.Ident)(0x14007ffe080)(int),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      })
     },
     Closing: (token.Pos) 0
    })
   }),
   Body: (*ast.BlockStmt)(0x14007fdddd0)({
    Lbrace: (token.Pos) 4232494,
    List: ([]ast.Stmt) (len=3 cap=4) {
     (*ast.AssignStmt)(0x14007fe7980)({
      Lhs: ([]ast.Expr) (len=1 cap=1) {
       (*ast.Ident)(0x14007ffe0a0)(sum)
      },
      TokPos: (token.Pos) 4232501,
      Tok: (token.Token) :=,
      Rhs: ([]ast.Expr) (len=1 cap=1) {
       (*ast.BasicLit)(0x14007ffe0c0)({
        ValuePos: (token.Pos) 4232504,
        Kind: (token.Token) INT,
        Value: (string) (len=1) "0"
       })
      }
     }),
     (*ast.RangeStmt)(0x14007fe5aa0)({
      For: (token.Pos) 4232507,
      Key: (*ast.Ident)(0x14007ffe0e0)(_),
      Value: (*ast.Ident)(0x14007ffe100)(arg),
      TokPos: (token.Pos) 4232518,
      Tok: (token.Token) :=,
      Range: (token.Pos) 4232521,
      X: (*ast.Ident)(0x14007ffe140)(args),
      Body: (*ast.BlockStmt)(0x14007fddda0)({
       Lbrace: (token.Pos) 4232532,
       List: ([]ast.Stmt) (len=1 cap=1) {
        (*ast.AssignStmt)(0x14007fe7a00)({
         Lhs: ([]ast.Expr) (len=1 cap=1) {
          (*ast.Ident)(0x14007ffe180)(sum)
         },
         TokPos: (token.Pos) 4232540,
         Tok: (token.Token) +=,
         Rhs: ([]ast.Expr) (len=1 cap=1) {
          (*ast.Ident)(0x14007ffe1a0)(arg)
         }
        })
       },
       Rbrace: (token.Pos) 4232548
      })
     }),
     (*ast.ReturnStmt)(0x14007ffe200)({
      Return: (token.Pos) 4232551,
      Results: ([]ast.Expr) (len=1 cap=1) {
       (*ast.Ident)(0x14007ffe1e0)(sum)
      }
     })
    },
    Rbrace: (token.Pos) 4232562
   })
  }),
  imports: (map[string]string) {
  },
  aliases: (map[string]string) <nil>
 }),
 (*codegen.funcInfo)(0x14008480840)({
  function: (*ast.FuncDecl)(0x14007fddec0)({
   Doc: (*ast.CommentGroup)(0x14007f9b5f0)({
    List: ([]*ast.Comment) (len=1 cap=1) {
     (*ast.Comment)(0x14007f9b5d8)({
      Slash: (token.Pos) 4232645,
      Text: (string) (len=28) "// Returns the current time."
     })
    }
   }),
   Recv: (*ast.FieldList)(<nil>),
   Name: (*ast.Ident)(0x14007ffe2a0)(GetCurrentTime),
   Type: (*ast.FuncType)(0x14007ffe340)({
    Func: (token.Pos) 4232674,
    TypeParams: (*ast.FieldList)(<nil>),
    Params: (*ast.FieldList)(0x14007fdde30)({
     Opening: (token.Pos) 4232693,
     List: ([]*ast.Field) <nil>,
     Closing: (token.Pos) 4232694
    }),
    Results: (*ast.FieldList)(0x14007fdde60)({
     Opening: (token.Pos) 0,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x14007fe7ac0)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) <nil>,
       Type: (*ast.SelectorExpr)(0x14007f9b608)({
        X: (*ast.Ident)(0x14007ffe2c0)(time),
        Sel: (*ast.Ident)(0x14007ffe2e0)(Time)
       }),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      })
     },
     Closing: (token.Pos) 0
    })
   }),
   Body: (*ast.BlockStmt)(0x14007fdde90)({
    Lbrace: (token.Pos) 4232706,
    List: ([]ast.Stmt) (len=1 cap=1) {
     (*ast.ReturnStmt)(0x14007ffe320)({
      Return: (token.Pos) 4232709,
      Results: ([]ast.Expr) (len=1 cap=1) {
       (*ast.CallExpr)(0x14007fe7b00)({
        Fun: (*ast.Ident)(0x14007ffe300)(nowFunc),
        Lparen: (token.Pos) 4232723,
        Args: ([]ast.Expr) <nil>,
        Ellipsis: (token.Pos) 0,
        Rparen: (token.Pos) 4232724
       })
      }
     })
    },
    Rbrace: (token.Pos) 4232726
   })
  }),
  imports: (map[string]string) (len=1) {
   (string) (len=4) "time": (string) (len=4) "time"
  },
  aliases: (map[string]string) <nil>
 }),
 (*codegen.funcInfo)(0x140084808d0)({
  function: (*ast.FuncDecl)(0x14008478000)({
   Doc: (*ast.CommentGroup)(0x14007f9b638)({
    List: ([]*ast.Comment) (len=1 cap=1) {
     (*ast.Comment)(0x14007f9b620)({
      Slash: (token.Pos) 4232729,
      Text: (string) (len=50) "// Returns the current time formatted as a string."
     })
    }
   }),
   Recv: (*ast.FieldList)(<nil>),
   Name: (*ast.Ident)(0x14007ffe360)(GetCurrentTimeFormatted),
   Type: (*ast.FuncType)(0x14007ffe440)({
    Func: (token.Pos) 4232780,
    TypeParams: (*ast.FieldList)(<nil>),
    Params: (*ast.FieldList)(0x14007fddef0)({
     Opening: (token.Pos) 4232808,
     List: ([]*ast.Field) <nil>,
     Closing: (token.Pos) 4232809
    }),
    Results: (*ast.FieldList)(0x14007fddf20)({
     Opening: (token.Pos) 0,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x14007fe7b40)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) <nil>,
       Type: (*ast.Ident)(0x14007ffe380)(string),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      })
     },
     Closing: (token.Pos) 0
    })
   }),
   Body: (*ast.BlockStmt)(0x14007fddf50)({
    Lbrace: (token.Pos) 4232818,
    List: ([]ast.Stmt) (len=1 cap=1) {
     (*ast.ReturnStmt)(0x14007ffe420)({
      Return: (token.Pos) 4232821,
      Results: ([]ast.Expr) (len=1 cap=1) {
       (*ast.CallExpr)(0x14007fe7bc0)({
        Fun: (*ast.SelectorExpr)(0x14007f9b650)({
         X: (*ast.CallExpr)(0x14007fe7b80)({
          Fun: (*ast.Ident)(0x14007ffe3a0)(nowFunc),
          Lparen: (token.Pos) 4232835,
          Args: ([]ast.Expr) <nil>,
          Ellipsis: (token.Pos) 0,
          Rparen: (token.Pos) 4232836
         }),
         Sel: (*ast.Ident)(0x14007ffe3c0)(Format)
        }),
        Lparen: (token.Pos) 4232844,
        Args: ([]ast.Expr) (len=1 cap=1) {
         (*ast.SelectorExpr)(0x14007f9b668)({
          X: (*ast.Ident)(0x14007ffe3e0)(time),
          Sel: (*ast.Ident)(0x14007ffe400)(DateTime)
         })
        },
        Ellipsis: (token.Pos) 0,
        Rparen: (token.Pos) 4232858
       })
      }
     })
    },
    Rbrace: (token.Pos) 4232860
   })
  }),
  imports: (map[string]string) {
  },
  aliases: (map[string]string) <nil>
 }),
 (*codegen.funcInfo)(0x14008480948)({
  function: (*ast.FuncDecl)(0x14008478150)({
   Doc: (*ast.CommentGroup)(0x14007f9b698)({
    List: ([]*ast.Comment) (len=1 cap=1) {
     (*ast.Comment)(0x14007f9b680)({
      Slash: (token.Pos) 4232863,
      Text: (string) (len=75) "// Combines the first and last name of a person, and returns the full name."
     })
    }
   }),
   Recv: (*ast.FieldList)(<nil>),
   Name: (*ast.Ident)(0x14007ffe460)(GetFullName),
   Type: (*ast.FuncType)(0x14007ffe580)({
    Func: (token.Pos) 4232939,
    TypeParams: (*ast.FieldList)(<nil>),
    Params: (*ast.FieldList)(0x14008478060)({
     Opening: (token.Pos) 4232955,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x14007fe7c00)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) (len=2 cap=2) {
        (*ast.Ident)(0x14007ffe480)(firstName),
        (*ast.Ident)(0x14007ffe4a0)(lastName)
       },
       Type: (*ast.Ident)(0x14007ffe4c0)(string),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      })
     },
     Closing: (token.Pos) 4232982
    }),
    Results: (*ast.FieldList)(0x14008478090)({
     Opening: (token.Pos) 0,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x14007fe7c40)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) <nil>,
       Type: (*ast.Ident)(0x14007ffe4e0)(string),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      })
     },
     Closing: (token.Pos) 0
    })
   }),
   Body: (*ast.BlockStmt)(0x14008478120)({
    Lbrace: (token.Pos) 4232991,
    List: ([]ast.Stmt) (len=1 cap=1) {
     (*ast.ReturnStmt)(0x14007ffe560)({
      Return: (token.Pos) 4232994,
      Results: ([]ast.Expr) (len=1 cap=1) {
       (*ast.BinaryExpr)(0x140084780f0)({
        X: (*ast.BinaryExpr)(0x140084780c0)({
         X: (*ast.Ident)(0x14007ffe500)(firstName),
         OpPos: (token.Pos) 4233011,
         Op: (token.Token) +,
         Y: (*ast.BasicLit)(0x14007ffe520)({
          ValuePos: (token.Pos) 4233013,
          Kind: (token.Token) STRING,
          Value: (string) (len=3) "\" \""
         })
        }),
        OpPos: (token.Pos) 4233017,
        Op: (token.Token) +,
        Y: (*ast.Ident)(0x14007ffe540)(lastName)
       })
      }
     })
    },
    Rbrace: (token.Pos) 4233028
   })
  }),
  imports: (map[string]string) {
  },
  aliases: (map[string]string) <nil>
 }),
 (*codegen.funcInfo)(0x140084809d8)({
  function: (*ast.FuncDecl)(0x14008478300)({
   Doc: (*ast.CommentGroup)(0x14007f9b6f8)({
    List: ([]*ast.Comment) (len=2 cap=2) {
     (*ast.Comment)(0x14007f9b6c8)({
      Slash: (token.Pos) 4233031,
      Text: (string) (len=34) "// Says hello to a person by name."
     }),
     (*ast.Comment)(0x14007f9b6e0)({
      Slash: (token.Pos) 4233066,
      Text: (string) (len=65) "// If the name is not provided, it will say hello without a name."
     })
    }
   }),
   Recv: (*ast.FieldList)(<nil>),
   Name: (*ast.Ident)(0x14007ffe5a0)(SayHello),
   Type: (*ast.FuncType)(0x14007ffe720)({
    Func: (token.Pos) 4233132,
    TypeParams: (*ast.FieldList)(<nil>),
    Params: (*ast.FieldList)(0x14008478180)({
     Opening: (token.Pos) 4233145,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x14007fe7c80)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) (len=1 cap=1) {
        (*ast.Ident)(0x14007ffe5c0)(name)
       },
       Type: (*ast.StarExpr)(0x14007f9b710)({
        Star: (token.Pos) 4233151,
        X: (*ast.Ident)(0x14007ffe5e0)(string)
       }),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      })
     },
     Closing: (token.Pos) 4233158
    }),
    Results: (*ast.FieldList)(0x140084781b0)({
     Opening: (token.Pos) 0,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x14007fe7cc0)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) <nil>,
       Type: (*ast.Ident)(0x14007ffe600)(string),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      })
     },
     Closing: (token.Pos) 0
    })
   }),
   Body: (*ast.BlockStmt)(0x140084782d0)({
    Lbrace: (token.Pos) 4233167,
    List: ([]ast.Stmt) (len=1 cap=1) {
     (*ast.IfStmt)(0x14007fe7d00)({
      If: (token.Pos) 4233170,
      Init: (ast.Stmt) <nil>,
      Cond: (*ast.BinaryExpr)(0x140084781e0)({
       X: (*ast.Ident)(0x14007ffe620)(name),
       OpPos: (token.Pos) 4233178,
       Op: (token.Token) ==,
       Y: (*ast.Ident)(0x14007ffe640)(nil)
      }),
      Body: (*ast.BlockStmt)(0x14008478210)({
       Lbrace: (token.Pos) 4233185,
       List: ([]ast.Stmt) (len=1 cap=1) {
        (*ast.ReturnStmt)(0x14007ffe680)({
         Return: (token.Pos) 4233189,
         Results: ([]ast.Expr) (len=1 cap=1) {
          (*ast.BasicLit)(0x14007ffe660)({
           ValuePos: (token.Pos) 4233196,
           Kind: (token.Token) STRING,
           Value: (string) (len=8) "\"Hello!\""
          })
         }
        })
       },
       Rbrace: (token.Pos) 4233206
      }),
      Else: (*ast.BlockStmt)(0x140084782a0)({
       Lbrace: (token.Pos) 4233213,
       List: ([]ast.Stmt) (len=1 cap=1) {
        (*ast.ReturnStmt)(0x14007ffe700)({
         Return: (token.Pos) 4233217,
         Results: ([]ast.Expr) (len=1 cap=1) {
          (*ast.BinaryExpr)(0x14008478270)({
           X: (*ast.BinaryExpr)(0x14008478240)({
            X: (*ast.BasicLit)(0x14007ffe6a0)({
             ValuePos: (token.Pos) 4233224,
             Kind: (token.Token) STRING,
             Value: (string) (len=9) "\"Hello, \""
            }),
            OpPos: (token.Pos) 4233234,
            Op: (token.Token) +,
            Y: (*ast.StarExpr)(0x14007f9b740)({
             Star: (token.Pos) 4233236,
             X: (*ast.Ident)(0x14007ffe6c0)(name)
            })
           }),
           OpPos: (token.Pos) 4233242,
           Op: (token.Token) +,
           Y: (*ast.BasicLit)(0x14007ffe6e0)({
            ValuePos: (token.Pos) 4233244,
            Kind: (token.Token) STRING,
            Value: (string) (len=3) "\"!\""
           })
          })
         }
        })
       },
       Rbrace: (token.Pos) 4233249
      })
     })
    },
    Rbrace: (token.Pos) 4233251
   })
  }),
  imports: (map[string]string) {
  },
  aliases: (map[string]string) <nil>
 }),
 (*codegen.funcInfo)(0x14008480a98)({
  function: (*ast.FuncDecl)(0x14008478480)({
   Doc: (*ast.CommentGroup)(0x14007f9b848)({
    List: ([]*ast.Comment) (len=1 cap=1) {
     (*ast.Comment)(0x14007f9b830)({
      Slash: (token.Pos) 4233496,
      Text: (string) (len=24) "// Gets a person object."
     })
    }
   }),
   Recv: (*ast.FieldList)(<nil>),
   Name: (*ast.Ident)(0x14007ffe8a0)(GetPerson),
   Type: (*ast.FuncType)(0x14007ffea00)({
    Func: (token.Pos) 4233521,
    TypeParams: (*ast.FieldList)(<nil>),
    Params: (*ast.FieldList)(0x14008478360)({
     Opening: (token.Pos) 4233535,
     List: ([]*ast.Field) <nil>,
     Closing: (token.Pos) 4233536
    }),
    Results: (*ast.FieldList)(0x14008478390)({
     Opening: (token.Pos) 0,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x14007fe7e80)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) <nil>,
       Type: (*ast.Ident)(0x14007ffe8c0)(Person),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      })
     },
     Closing: (token.Pos) 0
    })
   }),
   Body: (*ast.BlockStmt)(0x14008478450)({
    Lbrace: (token.Pos) 4233545,
    List: ([]ast.Stmt) (len=1 cap=1) {
     (*ast.ReturnStmt)(0x14007ffe9e0)({
      Return: (token.Pos) 4233548,
      Results: ([]ast.Expr) (len=1 cap=1) {
       (*ast.CompositeLit)(0x14007fe7f00)({
        Type: (*ast.Ident)(0x14007ffe8e0)(Person),
        Lbrace: (token.Pos) 4233561,
        Elts: ([]ast.Expr) (len=3 cap=4) {
         (*ast.KeyValueExpr)(0x140084783c0)({
          Key: (*ast.Ident)(0x14007ffe900)(FirstName),
          Colon: (token.Pos) 4233574,
          Value: (*ast.BasicLit)(0x14007ffe920)({
           ValuePos: (token.Pos) 4233576,
           Kind: (token.Token) STRING,
           Value: (string) (len=6) "\"John\""
          })
         }),
         (*ast.KeyValueExpr)(0x140084783f0)({
          Key: (*ast.Ident)(0x14007ffe940)(LastName),
          Colon: (token.Pos) 4233594,
          Value: (*ast.BasicLit)(0x14007ffe960)({
           ValuePos: (token.Pos) 4233597,
           Kind: (token.Token) STRING,
           Value: (string) (len=5) "\"Doe\""
          })
         }),
         (*ast.KeyValueExpr)(0x14008478420)({
          Key: (*ast.Ident)(0x14007ffe9a0)(Age),
          Colon: (token.Pos) 4233609,
          Value: (*ast.BasicLit)(0x14007ffe9c0)({
           ValuePos: (token.Pos) 4233617,
           Kind: (token.Token) INT,
           Value: (string) (len=2) "42"
          })
         })
        },
        Rbrace: (token.Pos) 4233622,
        Incomplete: (bool) false
       })
      }
     })
    },
    Rbrace: (token.Pos) 4233624
   })
  }),
  imports: (map[string]string) {
  },
  aliases: (map[string]string) <nil>
 }),
 (*codegen.funcInfo)(0x14008480b10)({
  function: (*ast.FuncDecl)(0x14008478570)({
   Doc: (*ast.CommentGroup)(0x14007f9b878)({
    List: ([]*ast.Comment) (len=1 cap=1) {
     (*ast.Comment)(0x14007f9b860)({
      Slash: (token.Pos) 4233627,
      Text: (string) (len=53) "// Gets a random person object from a list of people."
     })
    }
   }),
   Recv: (*ast.FieldList)(<nil>),
   Name: (*ast.Ident)(0x14007ffea20)(GetRandomPerson),
   Type: (*ast.FuncType)(0x14007ffebc0)({
    Func: (token.Pos) 4233681,
    TypeParams: (*ast.FieldList)(<nil>),
    Params: (*ast.FieldList)(0x140084784b0)({
     Opening: (token.Pos) 4233701,
     List: ([]*ast.Field) <nil>,
     Closing: (token.Pos) 4233702
    }),
    Results: (*ast.FieldList)(0x140084784e0)({
     Opening: (token.Pos) 0,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x14007fe7f40)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) <nil>,
       Type: (*ast.Ident)(0x14007ffea40)(Person),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      })
     },
     Closing: (token.Pos) 0
    })
   }),
   Body: (*ast.BlockStmt)(0x14008478540)({
    Lbrace: (token.Pos) 4233711,
    List: ([]ast.Stmt) (len=3 cap=4) {
     (*ast.AssignStmt)(0x1400847a040)({
      Lhs: ([]ast.Expr) (len=1 cap=1) {
       (*ast.Ident)(0x14007ffea60)(people)
      },
      TokPos: (token.Pos) 4233721,
      Tok: (token.Token) :=,
      Rhs: ([]ast.Expr) (len=1 cap=1) {
       (*ast.CallExpr)(0x1400847a000)({
        Fun: (*ast.Ident)(0x14007ffea80)(GetPeople),
        Lparen: (token.Pos) 4233733,
        Args: ([]ast.Expr) <nil>,
        Ellipsis: (token.Pos) 0,
        Rparen: (token.Pos) 4233734
       })
      }
     }),
     (*ast.AssignStmt)(0x1400847a100)({
      Lhs: ([]ast.Expr) (len=1 cap=1) {
       (*ast.Ident)(0x14007ffeaa0)(i)
      },
      TokPos: (token.Pos) 4233739,
      Tok: (token.Token) :=,
      Rhs: ([]ast.Expr) (len=1 cap=1) {
       (*ast.CallExpr)(0x1400847a0c0)({
        Fun: (*ast.SelectorExpr)(0x14007f9b890)({
         X: (*ast.Ident)(0x14007ffeac0)(rand),
         Sel: (*ast.Ident)(0x14007ffeae0)(Intn)
        }),
        Lparen: (token.Pos) 4233751,
        Args: ([]ast.Expr) (len=1 cap=1) {
         (*ast.CallExpr)(0x1400847a080)({
          Fun: (*ast.Ident)(0x14007ffeb00)(len),
          Lparen: (token.Pos) 4233755,
          Args: ([]ast.Expr) (len=1 cap=1) {
           (*ast.Ident)(0x14007ffeb20)(people)
          },
          Ellipsis: (token.Pos) 0,
          Rparen: (token.Pos) 4233762
         })
        },
        Ellipsis: (token.Pos) 0,
        Rparen: (token.Pos) 4233763
       })
      }
     }),
     (*ast.ReturnStmt)(0x14007ffeba0)({
      Return: (token.Pos) 4233766,
      Results: ([]ast.Expr) (len=1 cap=1) {
       (*ast.IndexExpr)(0x14008478510)({
        X: (*ast.Ident)(0x14007ffeb60)(people),
        Lbrack: (token.Pos) 4233779,
        Index: (*ast.Ident)(0x14007ffeb80)(i),
        Rbrack: (token.Pos) 4233781
       })
      }
     })
    },
    Rbrace: (token.Pos) 4233783
   })
  }),
  imports: (map[string]string) {
  },
  aliases: (map[string]string) <nil>
 }),
 (*codegen.funcInfo)(0x14008480ba0)({
  function: (*ast.FuncDecl)(0x14008478840)({
   Doc: (*ast.CommentGroup)(0x14007f9b8c0)({
    List: ([]*ast.Comment) (len=1 cap=1) {
     (*ast.Comment)(0x14007f9b8a8)({
      Slash: (token.Pos) 4233786,
      Text: (string) (len=25) "// Gets a list of people."
     })
    }
   }),
   Recv: (*ast.FieldList)(<nil>),
   Name: (*ast.Ident)(0x14007ffebe0)(GetPeople),
   Type: (*ast.FuncType)(0x14007ffef20)({
    Func: (token.Pos) 4233812,
    TypeParams: (*ast.FieldList)(<nil>),
    Params: (*ast.FieldList)(0x140084785a0)({
     Opening: (token.Pos) 4233826,
     List: ([]*ast.Field) <nil>,
     Closing: (token.Pos) 4233827
    }),
    Results: (*ast.FieldList)(0x14008478600)({
     Opening: (token.Pos) 0,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x1400847a180)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) <nil>,
       Type: (*ast.ArrayType)(0x140084785d0)({
        Lbrack: (token.Pos) 4233829,
        Len: (ast.Expr) <nil>,
        Elt: (*ast.Ident)(0x14007ffec00)(Person)
       }),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      })
     },
     Closing: (token.Pos) 0
    })
   }),
   Body: (*ast.BlockStmt)(0x14008478810)({
    Lbrace: (token.Pos) 4233838,
    List: ([]ast.Stmt) (len=1 cap=1) {
     (*ast.ReturnStmt)(0x14007ffef00)({
      Return: (token.Pos) 4233841,
      Results: ([]ast.Expr) (len=1 cap=1) {
       (*ast.CompositeLit)(0x1400847a380)({
        Type: (*ast.ArrayType)(0x14008478630)({
         Lbrack: (token.Pos) 4233848,
         Len: (ast.Expr) <nil>,
         Elt: (*ast.Ident)(0x14007ffec20)(Person)
        }),
        Lbrace: (token.Pos) 4233856,
        Elts: ([]ast.Expr) (len=3 cap=4) {
         (*ast.CompositeLit)(0x1400847a200)({
          Type: (ast.Expr) <nil>,
          Lbrace: (token.Pos) 4233860,
          Elts: ([]ast.Expr) (len=3 cap=4) {
           (*ast.KeyValueExpr)(0x14008478660)({
            Key: (*ast.Ident)(0x14007ffec40)(FirstName),
            Colon: (token.Pos) 4233874,
            Value: (*ast.BasicLit)(0x14007ffec60)({
             ValuePos: (token.Pos) 4233876,
             Kind: (token.Token) STRING,
             Value: (string) (len=5) "\"Bob\""
            })
           }),
           (*ast.KeyValueExpr)(0x14008478690)({
            Key: (*ast.Ident)(0x14007ffec80)(LastName),
            Colon: (token.Pos) 4233894,
            Value: (*ast.BasicLit)(0x14007ffeca0)({
             ValuePos: (token.Pos) 4233897,
             Kind: (token.Token) STRING,
             Value: (string) (len=7) "\"Smith\""
            })
           }),
           (*ast.KeyValueExpr)(0x140084786c0)({
            Key: (*ast.Ident)(0x14007ffece0)(Age),
            Colon: (token.Pos) 4233912,
            Value: (*ast.BasicLit)(0x14007ffed00)({
             ValuePos: (token.Pos) 4233920,
             Kind: (token.Token) INT,
             Value: (string) (len=2) "42"
            })
           })
          },
          Rbrace: (token.Pos) 4233926,
          Incomplete: (bool) false
         }),
         (*ast.CompositeLit)(0x1400847a280)({
          Type: (ast.Expr) <nil>,
          Lbrace: (token.Pos) 4233931,
          Elts: ([]ast.Expr) (len=3 cap=4) {
           (*ast.KeyValueExpr)(0x140084786f0)({
            Key: (*ast.Ident)(0x14007ffed20)(FirstName),
            Colon: (token.Pos) 4233945,
            Value: (*ast.BasicLit)(0x14007ffed40)({
             ValuePos: (token.Pos) 4233947,
             Kind: (token.Token) STRING,
             Value: (string) (len=7) "\"Alice\""
            })
           }),
           (*ast.KeyValueExpr)(0x14008478720)({
            Key: (*ast.Ident)(0x14007ffed60)(LastName),
            Colon: (token.Pos) 4233967,
            Value: (*ast.BasicLit)(0x14007ffed80)({
             ValuePos: (token.Pos) 4233970,
             Kind: (token.Token) STRING,
             Value: (string) (len=7) "\"Jones\""
            })
           }),
           (*ast.KeyValueExpr)(0x14008478750)({
            Key: (*ast.Ident)(0x14007ffedc0)(Age),
            Colon: (token.Pos) 4233985,
            Value: (*ast.BasicLit)(0x14007ffede0)({
             ValuePos: (token.Pos) 4233993,
             Kind: (token.Token) INT,
             Value: (string) (len=2) "35"
            })
           })
          },
          Rbrace: (token.Pos) 4233999,
          Incomplete: (bool) false
         }),
         (*ast.CompositeLit)(0x1400847a300)({
          Type: (ast.Expr) <nil>,
          Lbrace: (token.Pos) 4234004,
          Elts: ([]ast.Expr) (len=3 cap=4) {
           (*ast.KeyValueExpr)(0x14008478780)({
            Key: (*ast.Ident)(0x14007ffee20)(FirstName),
            Colon: (token.Pos) 4234018,
            Value: (*ast.BasicLit)(0x14007ffee40)({
             ValuePos: (token.Pos) 4234020,
             Kind: (token.Token) STRING,
             Value: (string) (len=9) "\"Charlie\""
            })
           }),
           (*ast.KeyValueExpr)(0x140084787b0)({
            Key: (*ast.Ident)(0x14007ffee60)(LastName),
            Colon: (token.Pos) 4234042,
            Value: (*ast.BasicLit)(0x14007ffee80)({
             ValuePos: (token.Pos) 4234045,
             Kind: (token.Token) STRING,
             Value: (string) (len=7) "\"Brown\""
            })
           }),
           (*ast.KeyValueExpr)(0x140084787e0)({
            Key: (*ast.Ident)(0x14007ffeec0)(Age),
            Colon: (token.Pos) 4234060,
            Value: (*ast.BasicLit)(0x14007ffeee0)({
             ValuePos: (token.Pos) 4234068,
             Kind: (token.Token) INT,
             Value: (string) (len=1) "8"
            })
           })
          },
          Rbrace: (token.Pos) 4234073,
          Incomplete: (bool) false
         })
        },
        Rbrace: (token.Pos) 4234077,
        Incomplete: (bool) false
       })
      }
     })
    },
    Rbrace: (token.Pos) 4234079
   })
  }),
  imports: (map[string]string) {
  },
  aliases: (map[string]string) <nil>
 }),
 (*codegen.funcInfo)(0x14008480c30)({
  function: (*ast.FuncDecl)(0x14008478930)({
   Doc: (*ast.CommentGroup)(0x14007f9b8f0)({
    List: ([]*ast.Comment) (len=1 cap=1) {
     (*ast.Comment)(0x14007f9b8d8)({
      Slash: (token.Pos) 4234082,
      Text: (string) (len=37) "// Gets the name and age of a person."
     })
    }
   }),
   Recv: (*ast.FieldList)(<nil>),
   Name: (*ast.Ident)(0x14007ffef40)(GetNameAndAge),
   Type: (*ast.FuncType)(0x14007fff180)({
    Func: (token.Pos) 4234120,
    TypeParams: (*ast.FieldList)(<nil>),
    Params: (*ast.FieldList)(0x14008478870)({
     Opening: (token.Pos) 4234138,
     List: ([]*ast.Field) <nil>,
     Closing: (token.Pos) 4234139
    }),
    Results: (*ast.FieldList)(0x140084788d0)({
     Opening: (token.Pos) 4234141,
     List: ([]*ast.Field) (len=2 cap=2) {
      (*ast.Field)(0x1400847a3c0)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) (len=1 cap=1) {
        (*ast.Ident)(0x14007ffef60)(name)
       },
       Type: (*ast.Ident)(0x14007ffef80)(string),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      }),
      (*ast.Field)(0x1400847a400)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) (len=1 cap=1) {
        (*ast.Ident)(0x14007ffefa0)(age)
       },
       Type: (*ast.Ident)(0x14007ffefc0)(int),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      })
     },
     Closing: (token.Pos) 4234162
    })
   }),
   Body: (*ast.BlockStmt)(0x14008478900)({
    Lbrace: (token.Pos) 4234164,
    List: ([]ast.Stmt) (len=2 cap=2) {
     (*ast.AssignStmt)(0x1400847a480)({
      Lhs: ([]ast.Expr) (len=1 cap=1) {
       (*ast.Ident)(0x14007ffefe0)(p)
      },
      TokPos: (token.Pos) 4234169,
      Tok: (token.Token) :=,
      Rhs: ([]ast.Expr) (len=1 cap=1) {
       (*ast.CallExpr)(0x1400847a440)({
        Fun: (*ast.Ident)(0x14007fff000)(GetPerson),
        Lparen: (token.Pos) 4234181,
        Args: ([]ast.Expr) <nil>,
        Ellipsis: (token.Pos) 0,
        Rparen: (token.Pos) 4234182
       })
      }
     }),
     (*ast.ReturnStmt)(0x14007fff140)({
      Return: (token.Pos) 4234185,
      Results: ([]ast.Expr) (len=2 cap=2) {
       (*ast.CallExpr)(0x1400847a4c0)({
        Fun: (*ast.Ident)(0x14007fff020)(GetFullName),
        Lparen: (token.Pos) 4234203,
        Args: ([]ast.Expr) (len=2 cap=2) {
         (*ast.SelectorExpr)(0x14007f9b920)({
          X: (*ast.Ident)(0x14007fff040)(p),
          Sel: (*ast.Ident)(0x14007fff060)(FirstName)
         }),
         (*ast.SelectorExpr)(0x14007f9b938)({
          X: (*ast.Ident)(0x14007fff080)(p),
          Sel: (*ast.Ident)(0x14007fff0a0)(LastName)
         })
        },
        Ellipsis: (token.Pos) 0,
        Rparen: (token.Pos) 4234227
       }),
       (*ast.SelectorExpr)(0x14007f9b950)({
        X: (*ast.Ident)(0x14007fff0e0)(p),
        Sel: (*ast.Ident)(0x14007fff100)(Age)
       })
      }
     })
    },
    Rbrace: (token.Pos) 4234236
   })
  }),
  imports: (map[string]string) {
  },
  aliases: (map[string]string) <nil>
 }),
 (*codegen.funcInfo)(0x14008480cc0)({
  function: (*ast.FuncDecl)(0x14008478ab0)({
   Doc: (*ast.CommentGroup)(0x14007f9b980)({
    List: ([]*ast.Comment) (len=1 cap=1) {
     (*ast.Comment)(0x14007f9b968)({
      Slash: (token.Pos) 4234239,
      Text: (string) (len=28) "// Tests returning an error."
     })
    }
   }),
   Recv: (*ast.FieldList)(<nil>),
   Name: (*ast.Ident)(0x14007fff1a0)(TestNormalError),
   Type: (*ast.FuncType)(0x14007fff460)({
    Func: (token.Pos) 4234268,
    TypeParams: (*ast.FieldList)(<nil>),
    Params: (*ast.FieldList)(0x14008478960)({
     Opening: (token.Pos) 4234288,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x1400847a500)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) (len=1 cap=1) {
        (*ast.Ident)(0x14007fff1c0)(input)
       },
       Type: (*ast.Ident)(0x14007fff1e0)(string),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      })
     },
     Closing: (token.Pos) 4234301
    }),
    Results: (*ast.FieldList)(0x140084789c0)({
     Opening: (token.Pos) 4234303,
     List: ([]*ast.Field) (len=2 cap=2) {
      (*ast.Field)(0x1400847a540)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) <nil>,
       Type: (*ast.Ident)(0x14007fff200)(string),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      }),
      (*ast.Field)(0x1400847a580)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) <nil>,
       Type: (*ast.Ident)(0x14007fff220)(error),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      })
     },
     Closing: (token.Pos) 4234317
    })
   }),
   Body: (*ast.BlockStmt)(0x14008478a80)({
    Lbrace: (token.Pos) 4234319,
    List: ([]ast.Stmt) (len=3 cap=4) {
     (*ast.IfStmt)(0x1400847a640)({
      If: (token.Pos) 4234645,
      Init: (ast.Stmt) <nil>,
      Cond: (*ast.BinaryExpr)(0x140084789f0)({
       X: (*ast.Ident)(0x14007fff260)(input),
       OpPos: (token.Pos) 4234654,
       Op: (token.Token) ==,
       Y: (*ast.BasicLit)(0x14007fff280)({
        ValuePos: (token.Pos) 4234657,
        Kind: (token.Token) STRING,
        Value: (string) (len=2) "\"\""
       })
      }),
      Body: (*ast.BlockStmt)(0x14008478a20)({
       Lbrace: (token.Pos) 4234660,
       List: ([]ast.Stmt) (len=1 cap=1) {
        (*ast.ReturnStmt)(0x14007fff340)({
         Return: (token.Pos) 4234664,
         Results: ([]ast.Expr) (len=2 cap=2) {
          (*ast.BasicLit)(0x14007fff2a0)({
           ValuePos: (token.Pos) 4234671,
           Kind: (token.Token) STRING,
           Value: (string) (len=2) "\"\""
          }),
          (*ast.CallExpr)(0x1400847a600)({
           Fun: (*ast.SelectorExpr)(0x14007f9ba58)({
            X: (*ast.Ident)(0x14007fff2c0)(errors),
            Sel: (*ast.Ident)(0x14007fff2e0)(New)
           }),
           Lparen: (token.Pos) 4234685,
           Args: ([]ast.Expr) (len=1 cap=1) {
            (*ast.BasicLit)(0x14007fff300)({
             ValuePos: (token.Pos) 4234686,
             Kind: (token.Token) STRING,
             Value: (string) (len=16) "\"input is empty\""
            })
           },
           Ellipsis: (token.Pos) 0,
           Rparen: (token.Pos) 4234702
          })
         }
        })
       },
       Rbrace: (token.Pos) 4234705
      }),
      Else: (ast.Stmt) <nil>
     }),
     (*ast.AssignStmt)(0x1400847a680)({
      Lhs: ([]ast.Expr) (len=1 cap=1) {
       (*ast.Ident)(0x14007fff360)(output)
      },
      TokPos: (token.Pos) 4234715,
      Tok: (token.Token) :=,
      Rhs: ([]ast.Expr) (len=1 cap=1) {
       (*ast.BinaryExpr)(0x14008478a50)({
        X: (*ast.BasicLit)(0x14007fff380)({
         ValuePos: (token.Pos) 4234718,
         Kind: (token.Token) STRING,
         Value: (string) (len=12) "\"You said: \""
        }),
        OpPos: (token.Pos) 4234731,
        Op: (token.Token) +,
        Y: (*ast.Ident)(0x14007fff3a0)(input)
       })
      }
     }),
     (*ast.ReturnStmt)(0x14007fff440)({
      Return: (token.Pos) 4234740,
      Results: ([]ast.Expr) (len=2 cap=2) {
       (*ast.Ident)(0x14007fff3e0)(output),
       (*ast.Ident)(0x14007fff400)(nil)
      }
     })
    },
    Rbrace: (token.Pos) 4234759
   })
  }),
  imports: (map[string]string) {
  },
  aliases: (map[string]string) <nil>
 }),
 (*codegen.funcInfo)(0x14008480d50)({
  function: (*ast.FuncDecl)(0x14008478c00)({
   Doc: (*ast.CommentGroup)(0x14007f9ba88)({
    List: ([]*ast.Comment) (len=1 cap=1) {
     (*ast.Comment)(0x14007f9ba70)({
      Slash: (token.Pos) 4234762,
      Text: (string) (len=58) "// Tests an alternative way to handle errors in functions."
     })
    }
   }),
   Recv: (*ast.FieldList)(<nil>),
   Name: (*ast.Ident)(0x14007fff480)(TestAlternativeError),
   Type: (*ast.FuncType)(0x14007fff6c0)({
    Func: (token.Pos) 4234821,
    TypeParams: (*ast.FieldList)(<nil>),
    Params: (*ast.FieldList)(0x14008478ae0)({
     Opening: (token.Pos) 4234846,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x1400847a700)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) (len=1 cap=1) {
        (*ast.Ident)(0x14007fff4a0)(input)
       },
       Type: (*ast.Ident)(0x14007fff4c0)(string),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      })
     },
     Closing: (token.Pos) 4234859
    }),
    Results: (*ast.FieldList)(0x14008478b10)({
     Opening: (token.Pos) 0,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x1400847a740)({
       Doc: (*ast.CommentGroup)(<nil>),
       Names: ([]*ast.Ident) <nil>,
       Type: (*ast.Ident)(0x14007fff4e0)(string),
       Tag: (*ast.BasicLit)(<nil>),
       Comment: (*ast.CommentGroup)(<nil>)
      })
     },
     Closing: (token.Pos) 0
    })
   }),
   Body: (*ast.BlockStmt)(0x14008478bd0)({
    Lbrace: (token.Pos) 4234868,
    List: ([]ast.Stmt) (len=3 cap=4) {
     (*ast.IfStmt)(0x1400847a7c0)({
      If: (token.Pos) 4235012,
      Init: (ast.Stmt) <nil>,
      Cond: (*ast.BinaryExpr)(0x14008478b40)({
       X: (*ast.Ident)(0x14007fff500)(input),
       OpPos: (token.Pos) 4235021,
       Op: (token.Token) ==,
       Y: (*ast.BasicLit)(0x14007fff520)({
        ValuePos: (token.Pos) 4235024,
        Kind: (token.Token) STRING,
        Value: (string) (len=2) "\"\""
       })
      }),
      Body: (*ast.BlockStmt)(0x14008478b70)({
       Lbrace: (token.Pos) 4235027,
       List: ([]ast.Stmt) (len=2 cap=2) {
        (*ast.ExprStmt)(0x14007fc7c50)({
         X: (*ast.CallExpr)(0x1400847a780)({
          Fun: (*ast.SelectorExpr)(0x14007f9bb00)({
           X: (*ast.Ident)(0x14007fff540)(console),
           Sel: (*ast.Ident)(0x14007fff560)(Error)
          }),
          Lparen: (token.Pos) 4235044,
          Args: ([]ast.Expr) (len=1 cap=1) {
           (*ast.BasicLit)(0x14007fff580)({
            ValuePos: (token.Pos) 4235045,
            Kind: (token.Token) STRING,
            Value: (string) (len=16) "\"input is empty\""
           })
          },
          Ellipsis: (token.Pos) 0,
          Rparen: (token.Pos) 4235061
         })
        }),
        (*ast.ReturnStmt)(0x14007fff5c0)({
         Return: (token.Pos) 4235065,
         Results: ([]ast.Expr) (len=1 cap=1) {
          (*ast.BasicLit)(0x14007fff5a0)({
           ValuePos: (token.Pos) 4235072,
           Kind: (token.Token) STRING,
           Value: (string) (len=2) "\"\""
          })
         }
        })
       },
       Rbrace: (token.Pos) 4235076
      }),
      Else: (ast.Stmt) <nil>
     }),
     (*ast.AssignStmt)(0x1400847a800)({
      Lhs: ([]ast.Expr) (len=1 cap=1) {
       (*ast.Ident)(0x14007fff600)(output)
      },
      TokPos: (token.Pos) 4235086,
      Tok: (token.Token) :=,
      Rhs: ([]ast.Expr) (len=1 cap=1) {
       (*ast.BinaryExpr)(0x14008478ba0)({
        X: (*ast.BasicLit)(0x14007fff620)({
         ValuePos: (token.Pos) 4235089,
         Kind: (token.Token) STRING,
         Value: (string) (len=12) "\"You said: \""
        }),
        OpPos: (token.Pos) 4235102,
        Op: (token.Token) +,
        Y: (*ast.Ident)(0x14007fff640)(input)
       })
      }
     }),
     (*ast.ReturnStmt)(0x14007fff6a0)({
      Return: (token.Pos) 4235111,
      Results: ([]ast.Expr) (len=1 cap=1) {
       (*ast.Ident)(0x14007fff680)(output)
      }
     })
    },
    Rbrace: (token.Pos) 4235125
   })
  }),
  imports: (map[string]string) {
  },
  aliases: (map[string]string) <nil>
 }),
 (*codegen.funcInfo)(0x14008480de0)({
  function: (*ast.FuncDecl)(0x14008478c90)({
   Doc: (*ast.CommentGroup)(0x14007f9bb30)({
    List: ([]*ast.Comment) (len=1 cap=1) {
     (*ast.Comment)(0x14007f9bb18)({
      Slash: (token.Pos) 4235128,
      Text: (string) (len=17) "// Tests a panic."
     })
    }
   }),
   Recv: (*ast.FieldList)(<nil>),
   Name: (*ast.Ident)(0x14007fff6e0)(TestPanic),
   Type: (*ast.FuncType)(0x14007fff740)({
    Func: (token.Pos) 4235146,
    TypeParams: (*ast.FieldList)(<nil>),
    Params: (*ast.FieldList)(0x14008478c30)({
     Opening: (token.Pos) 4235160,
     List: ([]*ast.Field) <nil>,
     Closing: (token.Pos) 4235161
    }),
    Results: (*ast.FieldList)(<nil>)
   }),
   Body: (*ast.BlockStmt)(0x14008478c60)({
    Lbrace: (token.Pos) 4235163,
    List: ([]ast.Stmt) (len=1 cap=1) {
     (*ast.ExprStmt)(0x14007fc7cf0)({
      X: (*ast.CallExpr)(0x1400847a880)({
       Fun: (*ast.Ident)(0x14007fff700)(panic),
       Lparen: (token.Pos) 4235283,
       Args: ([]ast.Expr) (len=1 cap=1) {
        (*ast.BasicLit)(0x14007fff720)({
         ValuePos: (token.Pos) 4235284,
         Kind: (token.Token) STRING,
         Value: (string) (len=72) "\"This is a message from a panic.\\nThis is a second line from a panic.\\n\""
        })
       },
       Ellipsis: (token.Pos) 0,
       Rparen: (token.Pos) 4235356
      })
     })
    },
    Rbrace: (token.Pos) 4235358
   })
  }),
  imports: (map[string]string) {
  },
  aliases: (map[string]string) <nil>
 }),
 (*codegen.funcInfo)(0x14008480e58)({
  function: (*ast.FuncDecl)(0x14008478d20)({
   Doc: (*ast.CommentGroup)(0x14007f9bba8)({
    List: ([]*ast.Comment) (len=1 cap=1) {
     (*ast.Comment)(0x14007f9bb90)({
      Slash: (token.Pos) 4235361,
      Text: (string) (len=43) "// Tests an exit with a non-zero exit code."
     })
    }
   }),
   Recv: (*ast.FieldList)(<nil>),
   Name: (*ast.Ident)(0x14007fff760)(TestExit),
   Type: (*ast.FuncType)(0x14007fff8c0)({
    Func: (token.Pos) 4235405,
    TypeParams: (*ast.FieldList)(<nil>),
    Params: (*ast.FieldList)(0x14008478cc0)({
     Opening: (token.Pos) 4235418,
     List: ([]*ast.Field) <nil>,
     Closing: (token.Pos) 4235419
    }),
    Results: (*ast.FieldList)(<nil>)
   }),
   Body: (*ast.BlockStmt)(0x14008478cf0)({
    Lbrace: (token.Pos) 4235421,
    List: ([]ast.Stmt) (len=3 cap=4) {
     (*ast.ExprStmt)(0x14007fc7d40)({
      X: (*ast.CallExpr)(0x1400847a8c0)({
       Fun: (*ast.SelectorExpr)(0x14007f9bc38)({
        X: (*ast.Ident)(0x14007fff7a0)(console),
        Sel: (*ast.Ident)(0x14007fff7c0)(Error)
       }),
       Lparen: (token.Pos) 4235730,
       Args: ([]ast.Expr) (len=1 cap=1) {
        (*ast.BasicLit)(0x14007fff7e0)({
         ValuePos: (token.Pos) 4235731,
         Kind: (token.Token) STRING,
         Value: (string) (len=27) "\"This is an error message.\""
        })
       },
       Ellipsis: (token.Pos) 0,
       Rparen: (token.Pos) 4235758
      })
     }),
     (*ast.ExprStmt)(0x14007fc7d80)({
      X: (*ast.CallExpr)(0x1400847a900)({
       Fun: (*ast.SelectorExpr)(0x14007f9bc50)({
        X: (*ast.Ident)(0x14007fff800)(os),
        Sel: (*ast.Ident)(0x14007fff820)(Exit)
       }),
       Lparen: (token.Pos) 4235768,
       Args: ([]ast.Expr) (len=1 cap=1) {
        (*ast.BasicLit)(0x14007fff840)({
         ValuePos: (token.Pos) 4235769,
         Kind: (token.Token) INT,
         Value: (string) (len=1) "1"
        })
       },
       Ellipsis: (token.Pos) 0,
       Rparen: (token.Pos) 4235770
      })
     }),
     (*ast.ExprStmt)(0x14007fc7db0)({
      X: (*ast.CallExpr)(0x1400847a940)({
       Fun: (*ast.Ident)(0x14007fff880)(println),
       Lparen: (token.Pos) 4235780,
       Args: ([]ast.Expr) (len=1 cap=1) {
        (*ast.BasicLit)(0x14007fff8a0)({
         ValuePos: (token.Pos) 4235781,
         Kind: (token.Token) STRING,
         Value: (string) (len=33) "\"This line will not be executed.\""
        })
       },
       Ellipsis: (token.Pos) 0,
       Rparen: (token.Pos) 4235814
      })
     })
    },
    Rbrace: (token.Pos) 4235816
   })
  }),
  imports: (map[string]string) {
  },
  aliases: (map[string]string) <nil>
 }),
 (*codegen.funcInfo)(0x14008480ed0)({
  function: (*ast.FuncDecl)(0x14008478db0)({
   Doc: (*ast.CommentGroup)(0x14007f9bc80)({
    List: ([]*ast.Comment) (len=1 cap=1) {
     (*ast.Comment)(0x14007f9bc68)({
      Slash: (token.Pos) 4235819,
      Text: (string) (len=37) "// Tests logging at different levels."
     })
    }
   }),
   Recv: (*ast.FieldList)(<nil>),
   Name: (*ast.Ident)(0x14007fff8e0)(TestLogging),
   Type: (*ast.FuncType)(0x14007fffe40)({
    Func: (token.Pos) 4235857,
    TypeParams: (*ast.FieldList)(<nil>),
    Params: (*ast.FieldList)(0x14008478d50)({
     Opening: (token.Pos) 4235873,
     List: ([]*ast.Field) <nil>,
     Closing: (token.Pos) 4235874
    }),
    Results: (*ast.FieldList)(<nil>)
   }),
   Body: (*ast.BlockStmt)(0x14008478d80)({
    Lbrace: (token.Pos) 4235876,
    List: ([]ast.Stmt) (len=11 cap=16) {
     (*ast.ExprStmt)(0x14007fc7de0)({
      X: (*ast.CallExpr)(0x1400847a9c0)({
       Fun: (*ast.SelectorExpr)(0x14007f9bcc8)({
        X: (*ast.Ident)(0x14007fff900)(console),
        Sel: (*ast.Ident)(0x14007fff920)(Log)
       }),
       Lparen: (token.Pos) 4235941,
       Args: ([]ast.Expr) (len=1 cap=1) {
        (*ast.BasicLit)(0x14007fff940)({
         ValuePos: (token.Pos) 4235942,
         Kind: (token.Token) STRING,
         Value: (string) (len=31) "\"This is a simple log message.\""
        })
       },
       Ellipsis: (token.Pos) 0,
       Rparen: (token.Pos) 4235973
      })
     }),
     (*ast.ExprStmt)(0x14007fc7e20)({
      X: (*ast.CallExpr)(0x1400847aa00)({
       Fun: (*ast.SelectorExpr)(0x14007f9bd10)({
        X: (*ast.Ident)(0x14007fff960)(console),
        Sel: (*ast.Ident)(0x14007fff980)(Debug)
       }),
       Lparen: (token.Pos) 4236041,
       Args: ([]ast.Expr) (len=1 cap=1) {
        (*ast.BasicLit)(0x14007fff9a0)({
         ValuePos: (token.Pos) 4236042,
         Kind: (token.Token) STRING,
         Value: (string) (len=26) "\"This is a debug message.\""
        })
       },
       Ellipsis: (token.Pos) 0,
       Rparen: (token.Pos) 4236068
      })
     }),
     (*ast.ExprStmt)(0x14007fc7e50)({
      X: (*ast.CallExpr)(0x1400847aa40)({
       Fun: (*ast.SelectorExpr)(0x14007f9bd28)({
        X: (*ast.Ident)(0x14007fff9e0)(console),
        Sel: (*ast.Ident)(0x14007fffa00)(Info)
       }),
       Lparen: (token.Pos) 4236083,
       Args: ([]ast.Expr) (len=1 cap=1) {
        (*ast.BasicLit)(0x14007fffa20)({
         ValuePos: (token.Pos) 4236084,
         Kind: (token.Token) STRING,
         Value: (string) (len=26) "\"This is an info message.\""
        })
       },
       Ellipsis: (token.Pos) 0,
       Rparen: (token.Pos) 4236110
      })
     }),
     (*ast.ExprStmt)(0x14007fc7e80)({
      X: (*ast.CallExpr)(0x1400847aac0)({
       Fun: (*ast.SelectorExpr)(0x14007f9bd40)({
        X: (*ast.Ident)(0x14007fffa40)(console),
        Sel: (*ast.Ident)(0x14007fffa60)(Warn)
       }),
       Lparen: (token.Pos) 4236125,
       Args: ([]ast.Expr) (len=1 cap=1) {
        (*ast.BasicLit)(0x14007fffa80)({
         ValuePos: (token.Pos) 4236126,
         Kind: (token.Token) STRING,
         Value: (string) (len=28) "\"This is a warning message.\""
        })
       },
       Ellipsis: (token.Pos) 0,
       Rparen: (token.Pos) 4236154
      })
     }),
     (*ast.ExprStmt)(0x14007fc7eb0)({
      X: (*ast.CallExpr)(0x1400847ab00)({
       Fun: (*ast.SelectorExpr)(0x14007f9bd88)({
        X: (*ast.Ident)(0x14007fffaa0)(console),
        Sel: (*ast.Ident)(0x14007fffac0)(Error)
       }),
       Lparen: (token.Pos) 4236240,
       Args: ([]ast.Expr) (len=1 cap=1) {
        (*ast.BasicLit)(0x14007fffae0)({
         ValuePos: (token.Pos) 4236241,
         Kind: (token.Token) STRING,
         Value: (string) (len=27) "\"This is an error message.\""
        })
       },
       Ellipsis: (token.Pos) 0,
       Rparen: (token.Pos) 4236268
      })
     }),
     (*ast.ExprStmt)(0x14007fc7ee0)({
      X: (*ast.CallExpr)(0x1400847ab40)({
       Fun: (*ast.SelectorExpr)(0x14007f9bda0)({
        X: (*ast.Ident)(0x14007fffb00)(console),
        Sel: (*ast.Ident)(0x14007fffb20)(Error)
       }),
       Lparen: (token.Pos) 4236284,
       Args: ([]ast.Expr) (len=1 cap=1) {
        (*ast.BasicLit)(0x14007fffb40)({
         ValuePos: (token.Pos) 4236288,
         Kind: (token.Token) STRING,
         Value: (string) (len=143) "`This is line 1 of a multi-line error message.\n  This is line 2 of a multi-line error message.\n  This is line 3 of a multi-line error message.`"
        })
       },
       Ellipsis: (token.Pos) 0,
       Rparen: (token.Pos) 4236434
      })
     }),
     (*ast.ExprStmt)(0x14007fc7f10)({
      X: (*ast.CallExpr)(0x1400847ab80)({
       Fun: (*ast.Ident)(0x14007fffb60)(println),
       Lparen: (token.Pos) 4236499,
       Args: ([]ast.Expr) (len=1 cap=1) {
        (*ast.BasicLit)(0x14007fffb80)({
         ValuePos: (token.Pos) 4236500,
         Kind: (token.Token) STRING,
         Value: (string) (len=28) "\"This is a println message.\""
        })
       },
       Ellipsis: (token.Pos) 0,
       Rparen: (token.Pos) 4236528
      })
     }),
     (*ast.ExprStmt)(0x14007fc7f40)({
      X: (*ast.CallExpr)(0x1400847abc0)({
       Fun: (*ast.SelectorExpr)(0x14007f9bde8)({
        X: (*ast.Ident)(0x14007fffba0)(fmt),
        Sel: (*ast.Ident)(0x14007fffbc0)(Println)
       }),
       Lparen: (token.Pos) 4236542,
       Args: ([]ast.Expr) (len=1 cap=1) {
        (*ast.BasicLit)(0x14007fffbe0)({
         ValuePos: (token.Pos) 4236543,
         Kind: (token.Token) STRING,
         Value: (string) (len=32) "\"This is a fmt.Println message.\""
        })
       },
       Ellipsis: (token.Pos) 0,
       Rparen: (token.Pos) 4236575
      })
     }),
     (*ast.ExprStmt)(0x14007fc7f70)({
      X: (*ast.CallExpr)(0x1400847ac00)({
       Fun: (*ast.SelectorExpr)(0x14007f9be00)({
        X: (*ast.Ident)(0x14007fffc00)(fmt),
        Sel: (*ast.Ident)(0x14007fffc20)(Printf)
       }),
       Lparen: (token.Pos) 4236588,
       Args: ([]ast.Expr) (len=2 cap=2) {
        (*ast.BasicLit)(0x14007fffc40)({
         ValuePos: (token.Pos) 4236589,
         Kind: (token.Token) STRING,
         Value: (string) (len=38) "\"This is a fmt.Printf message (%s).\\n\""
        }),
        (*ast.BasicLit)(0x14007fffc60)({
         ValuePos: (token.Pos) 4236629,
         Kind: (token.Token) STRING,
         Value: (string) (len=18) "\"with a parameter\""
        })
       },
       Ellipsis: (token.Pos) 0,
       Rparen: (token.Pos) 4236647
      })
     }),
     (*ast.ExprStmt)(0x1400847e020)({
      X: (*ast.CallExpr)(0x1400847ac40)({
       Fun: (*ast.SelectorExpr)(0x14007f9be48)({
        X: (*ast.Ident)(0x14007fffca0)(fmt),
        Sel: (*ast.Ident)(0x14007fffcc0)(Fprintln)
       }),
       Lparen: (token.Pos) 4236761,
       Args: ([]ast.Expr) (len=2 cap=2) {
        (*ast.SelectorExpr)(0x14007f9be60)({
         X: (*ast.Ident)(0x14007fffce0)(os),
         Sel: (*ast.Ident)(0x14007fffd00)(Stdout)
        }),
        (*ast.BasicLit)(0x14007fffd20)({
         ValuePos: (token.Pos) 4236773,
         Kind: (token.Token) STRING,
         Value: (string) (len=44) "\"This is an info message printed to stdout.\""
        })
       },
       Ellipsis: (token.Pos) 0,
       Rparen: (token.Pos) 4236817
      })
     }),
     (*ast.ExprStmt)(0x1400847e050)({
      X: (*ast.CallExpr)(0x1400847ac80)({
       Fun: (*ast.SelectorExpr)(0x14007f9be78)({
        X: (*ast.Ident)(0x14007fffd60)(fmt),
        Sel: (*ast.Ident)(0x14007fffd80)(Fprintln)
       }),
       Lparen: (token.Pos) 4236832,
       Args: ([]ast.Expr) (len=2 cap=2) {
        (*ast.SelectorExpr)(0x14007f9be90)({
         X: (*ast.Ident)(0x14007fffda0)(os),
         Sel: (*ast.Ident)(0x14007fffdc0)(Stderr)
        }),
        (*ast.BasicLit)(0x14007fffde0)({
         ValuePos: (token.Pos) 4236844,
         Kind: (token.Token) STRING,
         Value: (string) (len=45) "\"This is an error message printed to stderr.\""
        })
       },
       Ellipsis: (token.Pos) 0,
       Rparen: (token.Pos) 4236889
      })
     })
    },
    Rbrace: (token.Pos) 4237450
   })
  }),
  imports: (map[string]string) {
  },
  aliases: (map[string]string) <nil>
 })
}

GML: getRequiredImports: alias="time", pkgPath="time"

GML: getRequiredImports: imports["time"]="time", names["time"]="time"

GML: PreProcess: imports=
(map[string]string) (len=1) {
 (string) (len=4) "time": (string) (len=4) "time"
}

GML: writeMainFunc: add a main function
Metadata:
  Plugin Name:     simple-example
  Go Module:       simple-example
  Modus SDK:       modus-sdk-go@0.15.0
  Build ID:        ctkntvnrack5flhp25ig
  Build Timestamp: 2024-12-23T15:17:18.907Z
  Git Repository:  https://github.com/hypermodeinc/modus
  Git Commit:      9fb68eefae318a92f10ade39a569a22725606f75

Functions:
  add(x int, y int) int
  add3(a int, b int, c *int) int
  addN(args []int) int
  getCurrentTime() time.Time
  getCurrentTimeFormatted() string
  getFullName(firstName string, lastName string) string
  getNameAndAge() (name string, age int)
  getPeople() []Person
  getPerson() Person
  getRandomPerson() Person
  logMessage(message string)
  sayHello(name *string) string
  testAlternativeError(input string) string
  testExit()
  testLogging()
  testNormalError(input string) string
  testPanic()

Custom Types:
  Person { firstName string, lastName string, age int }

*/
