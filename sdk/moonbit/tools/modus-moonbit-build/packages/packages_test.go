package packages

import (
	"go/ast"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPackage(t *testing.T) {
	t.Parallel()
	mode := NeedName | NeedImports | NeedDeps | NeedTypes | NeedSyntax | NeedTypesInfo
	dir := "testdata"
	cfg := &Config{Mode: mode, Dir: dir}
	got, err := Load(cfg, ".")
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(got, wantPackages); diff != "" {
		t.Errorf("Load() mismatch (-want +got):\n%v", diff)
	}
}

var wantPackages = []*Package{
	{
		MoonBitFiles: []string{"testdata/simple-example.mbt"},
		Name:         "main",
		Syntax: []*ast.File{
			{
				Name: &ast.Ident{Name: "testdata/simple-example.mbt"},
				Decls: []ast.Decl{
					&ast.FuncDecl{
						Name: &ast.Ident{Name: "log_message"},
						Type: &ast.FuncType{
							Params: &ast.FieldList{
								List: []*ast.Field{
									{Names: []*ast.Ident{{Name: "message"}}, Type: &ast.Ident{Name: "String"}},
								},
							},
						},
					},
					&ast.FuncDecl{
						Name: &ast.Ident{Name: "add"},
						Type: &ast.FuncType{
							Params: &ast.FieldList{
								List: []*ast.Field{
									{Names: []*ast.Ident{{Name: "x"}}, Type: &ast.Ident{Name: "Int"}},
									{Names: []*ast.Ident{{Name: "y"}}, Type: &ast.Ident{Name: "Int"}},
								},
							},
							Results: &ast.FieldList{
								List: []*ast.Field{
									{Type: &ast.Ident{Name: "Int"}},
								},
							},
						},
					},
				},
				Imports: []*ast.ImportSpec{
					{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
					{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/interface/imports/wasi/clocks/wallClock"`}},
					{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
					{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
				},
			},
		},
	},
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
 ParseFile: (func(*token.FileSet, string, []uint8) (*ast.File, error)) <nil>,
 Tests: (bool) false,
 Overlay: (map[string][]uint8) <nil>,
 modFile: (string) "",
 modFlag: (string) ""
})

GML: getFunctionsNeedingWrappers: pkg.Syntax[0]=
(*ast.File)(0x14007fd88c0)({
*/

// var syntax0 = &ast.File{
//  Package: (token.Pos) 4231927,
//  Name: (*ast.Ident)(0x14007fa3a00)(main),
//  Decls: ([]ast.Decl) (len=20 cap=32) {
//   (*ast.FuncDecl)(0x14007fdda10)({
//    Name: (*ast.Ident)(0x14007fa3b20)(LogMessage),
//    Type: (*ast.FuncType)(0x14007fa3c00)({
//     Func: (token.Pos) 4232069,
//     Params: (*ast.FieldList)(0x14007fdd9b0)({
//      Opening: (token.Pos) 4232084,
//      List: ([]*ast.Field) (len=1 cap=1) {
//       (*ast.Field)(0x14007fe7680)({
//        Names: ([]*ast.Ident) (len=1 cap=1) {
//         (*ast.Ident)(0x14007fa3b40)(message)
//        },
//        Type: (*ast.Ident)(0x14007fa3b60)(string),
//       })
//      },
//      Closing: (token.Pos) 4232099
//     }),
//    }),
//   }),
//   (*ast.FuncDecl)(0x14007fddb30)({
//    Name: (*ast.Ident)(0x14007fa3c40)(Add),
//    Type: (*ast.FuncType)(0x14007fa3d40)({
//     Func: (token.Pos) 4232182,
//     Params: (*ast.FieldList)(0x14007fdda70)({
//      Opening: (token.Pos) 4232190,
//      List: ([]*ast.Field) (len=1 cap=1) {
//       (*ast.Field)(0x14007fe7700)({
//        Names: ([]*ast.Ident) (len=2 cap=2) {
//         (*ast.Ident)(0x14007fa3c60)(x),
//         (*ast.Ident)(0x14007fa3c80)(y)
//        },
//        Type: (*ast.Ident)(0x14007fa3ca0)(int),
//       })
//      },
//      Closing: (token.Pos) 4232199
//     }),
//     Results: (*ast.FieldList)(0x14007fddaa0)({
//      Opening: (token.Pos) 0,
//      List: ([]*ast.Field) (len=1 cap=1) {
//       (*ast.Field)(0x14007fe7740)({
//        Names: ([]*ast.Ident) <nil>,
//        Type: (*ast.Ident)(0x14007fa3cc0)(int),
//       })
//      },
//      Closing: (token.Pos) 0
//     })
//    }),
//   }),
//   (*ast.FuncDecl)(0x14007fddd10)({
//    Name: (*ast.Ident)(0x14007fa3d60)(Add3),
//    Type: (*ast.FuncType)(0x14007ffe000)({
//     Func: (token.Pos) 4232314,
//     Params: (*ast.FieldList)(0x14007fddb90)({
//      Opening: (token.Pos) 4232323,
//      List: ([]*ast.Field) (len=2 cap=2) {
//       (*ast.Field)(0x14007fe77c0)({
//        Names: ([]*ast.Ident) (len=2 cap=2) {
//         (*ast.Ident)(0x14007fa3d80)(a),
//         (*ast.Ident)(0x14007fa3da0)(b)
//        },
//        Type: (*ast.Ident)(0x14007fa3dc0)(int),
//       }),
//       (*ast.Field)(0x14007fe7800)({
//        Names: ([]*ast.Ident) (len=1 cap=1) {
//         (*ast.Ident)(0x14007fa3de0)(c)
//        },
//        Type: (*ast.StarExpr)(0x14007f9b500)({
//         Star: (token.Pos) 4232336,
//         X: (*ast.Ident)(0x14007fa3e00)(int)
//        }),
//       })
//      },
//      Closing: (token.Pos) 4232340
//     }),
//     Results: (*ast.FieldList)(0x14007fddbc0)({
//      Opening: (token.Pos) 0,
//      List: ([]*ast.Field) (len=1 cap=1) {
//       (*ast.Field)(0x14007fe7840)({
//        Names: ([]*ast.Ident) <nil>,
//        Type: (*ast.Ident)(0x14007fa3e20)(int),
//       })
//      },
//      Closing: (token.Pos) 0
//     })
//    }),
//   }),
//   (*ast.FuncDecl)(0x14007fdde00)({
//    Name: (*ast.Ident)(0x14007ffe020)(AddN),
//    Type: (*ast.FuncType)(0x14007ffe220)({
//     Func: (token.Pos) 4232467,
//     Params: (*ast.FieldList)(0x14007fddd40)({
//      Opening: (token.Pos) 4232476,
//      List: ([]*ast.Field) (len=1 cap=1) {
//       (*ast.Field)(0x14007fe7900)({
//        Names: ([]*ast.Ident) (len=1 cap=1) {
//         (*ast.Ident)(0x14007ffe040)(args)
//        },
//        Type: (*ast.Ellipsis)(0x14007f9b560)({
//         Ellipsis: (token.Pos) 4232482,
//         Elt: (*ast.Ident)(0x14007ffe060)(int)
//        }),
//       })
//      },
//      Closing: (token.Pos) 4232488
//     }),
//     Results: (*ast.FieldList)(0x14007fddd70)({
//      Opening: (token.Pos) 0,
//      List: ([]*ast.Field) (len=1 cap=1) {
//       (*ast.Field)(0x14007fe7940)({
//        Names: ([]*ast.Ident) <nil>,
//        Type: (*ast.Ident)(0x14007ffe080)(int),
//       })
//      },
//      Closing: (token.Pos) 0
//     })
//    }),
//   }),
//   (*ast.FuncDecl)(0x14007fddec0)({
//    Name: (*ast.Ident)(0x14007ffe2a0)(GetCurrentTime),
//    Type: (*ast.FuncType)(0x14007ffe340)({
//     Func: (token.Pos) 4232674,
//     Params: (*ast.FieldList)(0x14007fdde30)({
//      Opening: (token.Pos) 4232693,
//      List: ([]*ast.Field) <nil>,
//      Closing: (token.Pos) 4232694
//     }),
//     Results: (*ast.FieldList)(0x14007fdde60)({
//      Opening: (token.Pos) 0,
//      List: ([]*ast.Field) (len=1 cap=1) {
//       (*ast.Field)(0x14007fe7ac0)({
//        Names: ([]*ast.Ident) <nil>,
//        Type: (*ast.SelectorExpr)(0x14007f9b608)({
//         X: (*ast.Ident)(0x14007ffe2c0)(time),
//         Sel: (*ast.Ident)(0x14007ffe2e0)(Time)
//        }),
//       })
//      },
//      Closing: (token.Pos) 0
//     })
//    }),
//   }),
//   (*ast.FuncDecl)(0x14008478000)({
//    Name: (*ast.Ident)(0x14007ffe360)(GetCurrentTimeFormatted),
//    Type: (*ast.FuncType)(0x14007ffe440)({
//     Func: (token.Pos) 4232780,
//     Params: (*ast.FieldList)(0x14007fddef0)({
//      Opening: (token.Pos) 4232808,
//      List: ([]*ast.Field) <nil>,
//      Closing: (token.Pos) 4232809
//     }),
//     Results: (*ast.FieldList)(0x14007fddf20)({
//      Opening: (token.Pos) 0,
//      List: ([]*ast.Field) (len=1 cap=1) {
//       (*ast.Field)(0x14007fe7b40)({
//        Names: ([]*ast.Ident) <nil>,
//        Type: (*ast.Ident)(0x14007ffe380)(string),
//       })
//      },
//      Closing: (token.Pos) 0
//     })
//    }),
//   }),
//   (*ast.FuncDecl)(0x14008478150)({
//    Name: (*ast.Ident)(0x14007ffe460)(GetFullName),
//    Type: (*ast.FuncType)(0x14007ffe580)({
//     Func: (token.Pos) 4232939,
//     Params: (*ast.FieldList)(0x14008478060)({
//      Opening: (token.Pos) 4232955,
//      List: ([]*ast.Field) (len=1 cap=1) {
//       (*ast.Field)(0x14007fe7c00)({
//        Names: ([]*ast.Ident) (len=2 cap=2) {
//         (*ast.Ident)(0x14007ffe480)(firstName),
//         (*ast.Ident)(0x14007ffe4a0)(lastName)
//        },
//        Type: (*ast.Ident)(0x14007ffe4c0)(string),
//       })
//      },
//      Closing: (token.Pos) 4232982
//     }),
//     Results: (*ast.FieldList)(0x14008478090)({
//      Opening: (token.Pos) 0,
//      List: ([]*ast.Field) (len=1 cap=1) {
//       (*ast.Field)(0x14007fe7c40)({
//        Names: ([]*ast.Ident) <nil>,
//        Type: (*ast.Ident)(0x14007ffe4e0)(string),
//       })
//      },
//      Closing: (token.Pos) 0
//     })
//    }),
//   }),
//   (*ast.FuncDecl)(0x14008478300)({
//    Name: (*ast.Ident)(0x14007ffe5a0)(SayHello),
//    Type: (*ast.FuncType)(0x14007ffe720)({
//     Func: (token.Pos) 4233132,
//     Params: (*ast.FieldList)(0x14008478180)({
//      Opening: (token.Pos) 4233145,
//      List: ([]*ast.Field) (len=1 cap=1) {
//       (*ast.Field)(0x14007fe7c80)({
//        Names: ([]*ast.Ident) (len=1 cap=1) {
//         (*ast.Ident)(0x14007ffe5c0)(name)
//        },
//        Type: (*ast.StarExpr)(0x14007f9b710)({
//         Star: (token.Pos) 4233151,
//         X: (*ast.Ident)(0x14007ffe5e0)(string)
//        }),
//       })
//      },
//      Closing: (token.Pos) 4233158
//     }),
//     Results: (*ast.FieldList)(0x140084781b0)({
//      Opening: (token.Pos) 0,
//      List: ([]*ast.Field) (len=1 cap=1) {
//       (*ast.Field)(0x14007fe7cc0)({
//        Names: ([]*ast.Ident) <nil>,
//        Type: (*ast.Ident)(0x14007ffe600)(string),
//       })
//      },
//      Closing: (token.Pos) 0
//     })
//    }),
//   }),
//   (*ast.FuncDecl)(0x14008478480)({
//    Name: (*ast.Ident)(0x14007ffe8a0)(GetPerson),
//    Type: (*ast.FuncType)(0x14007ffea00)({
//     Func: (token.Pos) 4233521,
//     Params: (*ast.FieldList)(0x14008478360)({
//      Opening: (token.Pos) 4233535,
//      List: ([]*ast.Field) <nil>,
//      Closing: (token.Pos) 4233536
//     }),
//     Results: (*ast.FieldList)(0x14008478390)({
//      Opening: (token.Pos) 0,
//      List: ([]*ast.Field) (len=1 cap=1) {
//       (*ast.Field)(0x14007fe7e80)({
//        Names: ([]*ast.Ident) <nil>,
//        Type: (*ast.Ident)(0x14007ffe8c0)(Person),
//       })
//      },
//      Closing: (token.Pos) 0
//     })
//    }),
//   }),
//   (*ast.FuncDecl)(0x14008478570)({
//    Name: (*ast.Ident)(0x14007ffea20)(GetRandomPerson),
//    Type: (*ast.FuncType)(0x14007ffebc0)({
//     Func: (token.Pos) 4233681,
//     Params: (*ast.FieldList)(0x140084784b0)({
//      Opening: (token.Pos) 4233701,
//      List: ([]*ast.Field) <nil>,
//      Closing: (token.Pos) 4233702
//     }),
//     Results: (*ast.FieldList)(0x140084784e0)({
//      Opening: (token.Pos) 0,
//      List: ([]*ast.Field) (len=1 cap=1) {
//       (*ast.Field)(0x14007fe7f40)({
//        Names: ([]*ast.Ident) <nil>,
//        Type: (*ast.Ident)(0x14007ffea40)(Person),
//       })
//      },
//      Closing: (token.Pos) 0
//     })
//    }),
//   }),
//   (*ast.FuncDecl)(0x14008478840)({
//    Name: (*ast.Ident)(0x14007ffebe0)(GetPeople),
//    Type: (*ast.FuncType)(0x14007ffef20)({
//     Func: (token.Pos) 4233812,
//     Params: (*ast.FieldList)(0x140084785a0)({
//      Opening: (token.Pos) 4233826,
//      List: ([]*ast.Field) <nil>,
//      Closing: (token.Pos) 4233827
//     }),
//     Results: (*ast.FieldList)(0x14008478600)({
//      Opening: (token.Pos) 0,
//      List: ([]*ast.Field) (len=1 cap=1) {
//       (*ast.Field)(0x1400847a180)({
//        Names: ([]*ast.Ident) <nil>,
//        Type: (*ast.ArrayType)(0x140084785d0)({
//         Lbrack: (token.Pos) 4233829,
//         Len: (ast.Expr) <nil>,
//         Elt: (*ast.Ident)(0x14007ffec00)(Person)
//        }),
//       })
//      },
//      Closing: (token.Pos) 0
//     })
//    }),
//   }),
//   (*ast.FuncDecl)(0x14008478930)({
//    Name: (*ast.Ident)(0x14007ffef40)(GetNameAndAge),
//    Type: (*ast.FuncType)(0x14007fff180)({
//     Func: (token.Pos) 4234120,
//     Params: (*ast.FieldList)(0x14008478870)({
//      Opening: (token.Pos) 4234138,
//      List: ([]*ast.Field) <nil>,
//      Closing: (token.Pos) 4234139
//     }),
//     Results: (*ast.FieldList)(0x140084788d0)({
//      Opening: (token.Pos) 4234141,
//      List: ([]*ast.Field) (len=2 cap=2) {
//       (*ast.Field)(0x1400847a3c0)({
//        Names: ([]*ast.Ident) (len=1 cap=1) {
//         (*ast.Ident)(0x14007ffef60)(name)
//        },
//        Type: (*ast.Ident)(0x14007ffef80)(string),
//       }),
//       (*ast.Field)(0x1400847a400)({
//        Names: ([]*ast.Ident) (len=1 cap=1) {
//         (*ast.Ident)(0x14007ffefa0)(age)
//        },
//        Type: (*ast.Ident)(0x14007ffefc0)(int),
//       })
//      },
//      Closing: (token.Pos) 4234162
//     })
//    }),
//   }),
//   (*ast.FuncDecl)(0x14008478ab0)({
//    Name: (*ast.Ident)(0x14007fff1a0)(TestNormalError),
//    Type: (*ast.FuncType)(0x14007fff460)({
//     Func: (token.Pos) 4234268,
//     Params: (*ast.FieldList)(0x14008478960)({
//      Opening: (token.Pos) 4234288,
//      List: ([]*ast.Field) (len=1 cap=1) {
//       (*ast.Field)(0x1400847a500)({
//        Names: ([]*ast.Ident) (len=1 cap=1) {
//         (*ast.Ident)(0x14007fff1c0)(input)
//        },
//        Type: (*ast.Ident)(0x14007fff1e0)(string),
//       })
//      },
//      Closing: (token.Pos) 4234301
//     }),
//     Results: (*ast.FieldList)(0x140084789c0)({
//      Opening: (token.Pos) 4234303,
//      List: ([]*ast.Field) (len=2 cap=2) {
//       (*ast.Field)(0x1400847a540)({
//        Names: ([]*ast.Ident) <nil>,
//        Type: (*ast.Ident)(0x14007fff200)(string),
//       }),
//       (*ast.Field)(0x1400847a580)({
//        Names: ([]*ast.Ident) <nil>,
//        Type: (*ast.Ident)(0x14007fff220)(error),
//       })
//      },
//      Closing: (token.Pos) 4234317
//     })
//    }),
//   }),
//   (*ast.FuncDecl)(0x14008478c00)({
//    Name: (*ast.Ident)(0x14007fff480)(TestAlternativeError),
//    Type: (*ast.FuncType)(0x14007fff6c0)({
//     Func: (token.Pos) 4234821,
//     Params: (*ast.FieldList)(0x14008478ae0)({
//      Opening: (token.Pos) 4234846,
//      List: ([]*ast.Field) (len=1 cap=1) {
//       (*ast.Field)(0x1400847a700)({
//        Names: ([]*ast.Ident) (len=1 cap=1) {
//         (*ast.Ident)(0x14007fff4a0)(input)
//        },
//        Type: (*ast.Ident)(0x14007fff4c0)(string),
//       })
//      },
//      Closing: (token.Pos) 4234859
//     }),
//     Results: (*ast.FieldList)(0x14008478b10)({
//      Opening: (token.Pos) 0,
//      List: ([]*ast.Field) (len=1 cap=1) {
//       (*ast.Field)(0x1400847a740)({
//        Names: ([]*ast.Ident) <nil>,
//        Type: (*ast.Ident)(0x14007fff4e0)(string),
//       })
//      },
//      Closing: (token.Pos) 0
//     })
//    }),
//   }),
//   (*ast.FuncDecl)(0x14008478c90)({
//    Name: (*ast.Ident)(0x14007fff6e0)(TestPanic),
//    Type: (*ast.FuncType)(0x14007fff740)({
//     Func: (token.Pos) 4235146,
//     Params: (*ast.FieldList)(0x14008478c30)({
//      Opening: (token.Pos) 4235160,
//      List: ([]*ast.Field) <nil>,
//      Closing: (token.Pos) 4235161
//     }),
//    }),
//   }),
//   (*ast.FuncDecl)(0x14008478d20)({
//    Name: (*ast.Ident)(0x14007fff760)(TestExit),
//    Type: (*ast.FuncType)(0x14007fff8c0)({
//     Func: (token.Pos) 4235405,
//     Params: (*ast.FieldList)(0x14008478cc0)({
//      Opening: (token.Pos) 4235418,
//      List: ([]*ast.Field) <nil>,
//      Closing: (token.Pos) 4235419
//     }),
//    }),
//   }),
//   (*ast.FuncDecl)(0x14008478db0)({
//    Name: (*ast.Ident)(0x14007fff8e0)(TestLogging),
//    Type: (*ast.FuncType)(0x14007fffe40)({
//     Func: (token.Pos) 4235857,
//     Params: (*ast.FieldList)(0x14008478d50)({
//      Opening: (token.Pos) 4235873,
//      List: ([]*ast.Field) <nil>,
//      Closing: (token.Pos) 4235874
//     }),
//    }),
//   })
//  },
//  FileStart: (token.Pos) 4231691,
//  FileEnd: (token.Pos) 4237452,
//  Scope: (*ast.Scope)(0x1400847e070)(scope 0x1400847e070 {
// 	func TestPanic
// 	func Add3
// 	func GetCurrentTime
// 	func GetCurrentTimeFormatted
// 	func GetPerson
// 	func TestNormalError
// 	func TestAlternativeError
// 	func SayHello
// 	func GetPeople
// 	func TestExit
// 	func LogMessage
// 	func Add
// 	func AddN
// 	var nowFunc
// 	func GetFullName
// 	type Person
// 	func GetRandomPerson
// 	func GetNameAndAge
// 	func TestLogging
// }
// ),
//  Imports: ([]*ast.ImportSpec) (len=6 cap=8) {
//   (*ast.ImportSpec)(0x14007fdd890)({
//    Path: (*ast.BasicLit)(0x14007fa3a20)({
//     ValuePos: (token.Pos) 4231951,
//     Kind: (token.Token) STRING,
//     Value: (string) (len=8) "\"errors\""
//    }),
//    EndPos: (token.Pos) 0
//   }),
//   (*ast.ImportSpec)(0x14007fdd8c0)({
//    Path: (*ast.BasicLit)(0x14007fa3a40)({
//     ValuePos: (token.Pos) 4231961,
//     Kind: (token.Token) STRING,
//     Value: (string) (len=5) "\"fmt\""
//    }),
//    EndPos: (token.Pos) 0
//   }),
//   (*ast.ImportSpec)(0x14007fdd8f0)({
//    Path: (*ast.BasicLit)(0x14007fa3a80)({
//     ValuePos: (token.Pos) 4231968,
//     Kind: (token.Token) STRING,
//     Value: (string) (len=11) "\"math/rand\""
//    }),
//    EndPos: (token.Pos) 0
//   }),
//   (*ast.ImportSpec)(0x14007fdd920)({
//    Path: (*ast.BasicLit)(0x14007fa3ac0)({
//     ValuePos: (token.Pos) 4231981,
//     Kind: (token.Token) STRING,
//     Value: (string) (len=4) "\"os\""
//    }),
//    EndPos: (token.Pos) 0
//   }),
//   (*ast.ImportSpec)(0x14007fdd950)({
//    Path: (*ast.BasicLit)(0x14007fa3ae0)({
//     ValuePos: (token.Pos) 4231987,
//     Kind: (token.Token) STRING,
//     Value: (string) (len=6) "\"time\""
//    }),
//    EndPos: (token.Pos) 0
//   }),
//   (*ast.ImportSpec)(0x14007fdd980)({
//    Path: (*ast.BasicLit)(0x14007fa3b00)({
//     ValuePos: (token.Pos) 4231996,
//     Kind: (token.Token) STRING,
//     Value: (string) (len=50) "\"github.com/hypermodeinc/modus/sdk/go/pkg/console\""
//    }),
//    EndPos: (token.Pos) 0
//   })
//  },
//  Unresolved: ([]*ast.Ident) (len=51 cap=64) {
//   (*ast.Ident)(0x14007fa3b60)(string),
//   (*ast.Ident)(0x14007fa3b80)(console),
//   (*ast.Ident)(0x14007fa3ca0)(int),
//   (*ast.Ident)(0x14007fa3cc0)(int),
//   (*ast.Ident)(0x14007fa3dc0)(int),
//   (*ast.Ident)(0x14007fa3e00)(int),
//   (*ast.Ident)(0x14007fa3e20)(int),
//   (*ast.Ident)(0x14007fa3e60)(nil),
//   (*ast.Ident)(0x14007ffe060)(int),
//   (*ast.Ident)(0x14007ffe080)(int),
//   (*ast.Ident)(0x14007ffe260)(time),
//   (*ast.Ident)(0x14007ffe2c0)(time),
//   (*ast.Ident)(0x14007ffe380)(string),
//   (*ast.Ident)(0x14007ffe3e0)(time),
//   (*ast.Ident)(0x14007ffe4c0)(string),
//   (*ast.Ident)(0x14007ffe4e0)(string),
//   (*ast.Ident)(0x14007ffe5e0)(string),
//   (*ast.Ident)(0x14007ffe600)(string),
//   (*ast.Ident)(0x14007ffe640)(nil),
//   (*ast.Ident)(0x14007ffe780)(string),
//   (*ast.Ident)(0x14007ffe7e0)(string),
//   (*ast.Ident)(0x14007ffe840)(int),
//   (*ast.Ident)(0x14007ffeac0)(rand),
//   (*ast.Ident)(0x14007ffeb00)(len),
//   (*ast.Ident)(0x14007ffef80)(string),
//   (*ast.Ident)(0x14007ffefc0)(int),
//   (*ast.Ident)(0x14007fff1e0)(string),
//   (*ast.Ident)(0x14007fff200)(string),
//   (*ast.Ident)(0x14007fff220)(error),
//   (*ast.Ident)(0x14007fff2c0)(errors),
//   (*ast.Ident)(0x14007fff400)(nil),
//   (*ast.Ident)(0x14007fff4c0)(string),
//   (*ast.Ident)(0x14007fff4e0)(string),
//   (*ast.Ident)(0x14007fff540)(console),
//   (*ast.Ident)(0x14007fff700)(panic),
//   (*ast.Ident)(0x14007fff7a0)(console),
//   (*ast.Ident)(0x14007fff800)(os),
//   (*ast.Ident)(0x14007fff880)(println),
//   (*ast.Ident)(0x14007fff900)(console),
//   (*ast.Ident)(0x14007fff960)(console),
//   (*ast.Ident)(0x14007fff9e0)(console),
//   (*ast.Ident)(0x14007fffa40)(console),
//   (*ast.Ident)(0x14007fffaa0)(console),
//   (*ast.Ident)(0x14007fffb00)(console),
//   (*ast.Ident)(0x14007fffb60)(println),
//   (*ast.Ident)(0x14007fffba0)(fmt),
//   (*ast.Ident)(0x14007fffc00)(fmt),
//   (*ast.Ident)(0x14007fffca0)(fmt),
//   (*ast.Ident)(0x14007fffce0)(os),
//   (*ast.Ident)(0x14007fffd60)(fmt),
//   (*ast.Ident)(0x14007fffda0)(os)
//  },
//  Comments: ([]*ast.CommentGroup) (len=33 cap=64) {
//   (*ast.CommentGroup)(0x14007f9b3e0)({
//    List: ([]*ast.Comment) (len=1 cap=1) {
//     (*ast.Comment)(0x14007f9b3c8)({
//      Slash: (token.Pos) 4231691,
//      Text: (string) (len=234) "/*\n * This example is part of the Modus project, licensed under the Apache License 2.0.\n * You may modify and use this example in accordance with the license.\n * See the LICENSE file that accompanied this code for further details.\n * /"  //GML
//     })
//    }
//   }),
//   (*ast.CommentGroup)(0x14007f9b410)({
//    List: ([]*ast.Comment) (len=1 cap=1) {
//     (*ast.Comment)(0x14007f9b3f8)({
//      Slash: (token.Pos) 4232050,
//      Text: (string) (len=18) "// Logs a message."
//     })
//    }
//   }),
//   (*ast.CommentGroup)(0x14007f9b470)({
//    List: ([]*ast.Comment) (len=1 cap=1) {
//     (*ast.Comment)(0x14007f9b458)({
//      Slash: (token.Pos) 4232128,
//      Text: (string) (len=53) "// Adds two integers together and returns the result."
//     })
//    }
//   }),
//   (*ast.CommentGroup)(0x14007f9b4d0)({
//    List: ([]*ast.Comment) (len=2 cap=2) {
//     (*ast.Comment)(0x14007f9b4a0)({
//      Slash: (token.Pos) 4232224,
//      Text: (string) (len=55) "// Adds three integers together and returns the result."
//     }),
//     (*ast.Comment)(0x14007f9b4b8)({
//      Slash: (token.Pos) 4232280,
//      Text: (string) (len=33) "// The third integer is optional."
//     })
//    }
//   }),
//   (*ast.CommentGroup)(0x14007f9b548)({
//    List: ([]*ast.Comment) (len=1 cap=1) {
//     (*ast.Comment)(0x14007f9b530)({
//      Slash: (token.Pos) 4232403,
//      Text: (string) (len=63) "// Adds any number of integers together and returns the result."
//     })
//    }
//   }),
//   (*ast.CommentGroup)(0x14007f9b5a8)({
//    List: ([]*ast.Comment) (len=1 cap=1) {
//     (*ast.Comment)(0x14007f9b590)({
//      Slash: (token.Pos) 4232565,
//      Text: (string) (len=55) "// this indirection is so we can mock time.Now in tests"
//     })
//    }
//   }),
//   (*ast.CommentGroup)(0x14007f9b5f0)({
//    List: ([]*ast.Comment) (len=1 cap=1) {
//     (*ast.Comment)(0x14007f9b5d8)({
//      Slash: (token.Pos) 4232645,
//      Text: (string) (len=28) "// Returns the current time."
//     })
//    }
//   }),
//   (*ast.CommentGroup)(0x14007f9b638)({
//    List: ([]*ast.Comment) (len=1 cap=1) {
//     (*ast.Comment)(0x14007f9b620)({
//      Slash: (token.Pos) 4232729,
//      Text: (string) (len=50) "// Returns the current time formatted as a string."
//     })
//    }
//   }),
//   (*ast.CommentGroup)(0x14007f9b698)({
//    List: ([]*ast.Comment) (len=1 cap=1) {
//     (*ast.Comment)(0x14007f9b680)({
//      Slash: (token.Pos) 4232863,
//      Text: (string) (len=75) "// Combines the first and last name of a person, and returns the full name."
//     })
//    }
//   }),
//   (*ast.CommentGroup)(0x14007f9b6f8)({
//    List: ([]*ast.Comment) (len=2 cap=2) {
//     (*ast.Comment)(0x14007f9b6c8)({
//      Slash: (token.Pos) 4233031,
//      Text: (string) (len=34) "// Says hello to a person by name."
//     }),
//     (*ast.Comment)(0x14007f9b6e0)({
//      Slash: (token.Pos) 4233066,
//      Text: (string) (len=65) "// If the name is not provided, it will say hello without a name."
//     })
//    }
//   }),
//   (*ast.CommentGroup)(0x14007f9b770)({
//    List: ([]*ast.Comment) (len=1 cap=1) {
//     (*ast.Comment)(0x14007f9b758)({
//      Slash: (token.Pos) 4233254,
//      Text: (string) (len=41) "// A simple object representing a person."
//     })
//    }
//   }),
//   (*ast.CommentGroup)(0x14007f9b7a0)({
//    List: ([]*ast.Comment) (len=1 cap=1) {
//     (*ast.Comment)(0x14007f9b788)({
//      Slash: (token.Pos) 4233319,
//      Text: (string) (len=27) "// The person's first name."
//     })
//    }
//   }),
//   (*ast.CommentGroup)(0x14007f9b7d0)({
//    List: ([]*ast.Comment) (len=1 cap=1) {
//     (*ast.Comment)(0x14007f9b7b8)({
//      Slash: (token.Pos) 4233386,
//      Text: (string) (len=26) "// The person's last name."
//     })
//    }
//   }),
//   (*ast.CommentGroup)(0x14007f9b800)({
//    List: ([]*ast.Comment) (len=1 cap=1) {
//     (*ast.Comment)(0x14007f9b7e8)({
//      Slash: (token.Pos) 4233450,
//      Text: (string) (len=20) "// The person's age."
//     })
//    }
//   }),
//   (*ast.CommentGroup)(0x14007f9b848)({
//    List: ([]*ast.Comment) (len=1 cap=1) {
//     (*ast.Comment)(0x14007f9b830)({
//      Slash: (token.Pos) 4233496,
//      Text: (string) (len=24) "// Gets a person object."
//     })
//    }
//   }),
//   (*ast.CommentGroup)(0x14007f9b878)({
//    List: ([]*ast.Comment) (len=1 cap=1) {
//     (*ast.Comment)(0x14007f9b860)({
//      Slash: (token.Pos) 4233627,
//      Text: (string) (len=53) "// Gets a random person object from a list of people."
//     })
//    }
//   }),
//   (*ast.CommentGroup)(0x14007f9b8c0)({
//    List: ([]*ast.Comment) (len=1 cap=1) {
//     (*ast.Comment)(0x14007f9b8a8)({
//      Slash: (token.Pos) 4233786,
//      Text: (string) (len=25) "// Gets a list of people."
//     })
//    }
//   }),
//   (*ast.CommentGroup)(0x14007f9b8f0)({
//    List: ([]*ast.Comment) (len=1 cap=1) {
//     (*ast.Comment)(0x14007f9b8d8)({
//      Slash: (token.Pos) 4234082,
//      Text: (string) (len=37) "// Gets the name and age of a person."
//     })
//    }
//   }),
//   (*ast.CommentGroup)(0x14007f9b980)({
//    List: ([]*ast.Comment) (len=1 cap=1) {
//     (*ast.Comment)(0x14007f9b968)({
//      Slash: (token.Pos) 4234239,
//      Text: (string) (len=28) "// Tests returning an error."
//     })
//    }
//   }),
//   (*ast.CommentGroup)(0x14007f9ba40)({
//    List: ([]*ast.Comment) (len=5 cap=8) {
//     (*ast.Comment)(0x14007f9b9c8)({
//      Slash: (token.Pos) 4234323,
//      Text: (string) (len=59) "// This is the preferred way to handle errors in functions."
//     }),
//     (*ast.Comment)(0x14007f9b9e0)({
//      Slash: (token.Pos) 4234384,
//      Text: (string) (len=62) "// Simply declare an error interface as the last return value."
//     }),
//     (*ast.Comment)(0x14007f9b9f8)({
//      Slash: (token.Pos) 4234448,
//      Text: (string) (len=65) "// You can use any object that implements the Go error interface."
//     }),
//     (*ast.Comment)(0x14007f9ba10)({
//      Slash: (token.Pos) 4234515,
//      Text: (string) (len=70) "// For example, you can create a new error with errors.New(\"message\"),"
//     }),
//     (*ast.Comment)(0x14007f9ba28)({
//      Slash: (token.Pos) 4234587,
//      Text: (string) (len=55) "// or with fmt.Errorf(\"message with %s\", \"parameters\")."
//     })
//    }
//   }),
//   (*ast.CommentGroup)(0x14007f9ba88)({
//    List: ([]*ast.Comment) (len=1 cap=1) {
//     (*ast.Comment)(0x14007f9ba70)({
//      Slash: (token.Pos) 4234762,
//      Text: (string) (len=58) "// Tests an alternative way to handle errors in functions."
//     })
//    }
//   }),
//   (*ast.CommentGroup)(0x14007f9bae8)({
//    List: ([]*ast.Comment) (len=2 cap=2) {
//     (*ast.Comment)(0x14007f9bab8)({
//      Slash: (token.Pos) 4234872,
//      Text: (string) (len=60) "// This is an alternative way to handle errors in functions."
//     }),
//     (*ast.Comment)(0x14007f9bad0)({
//      Slash: (token.Pos) 4234934,
//      Text: (string) (len=75) "// It is identical in behavior to TestNormalError, but is not Go idiomatic."
//     })
//    }
//   }),
//   (*ast.CommentGroup)(0x14007f9bb30)({
//    List: ([]*ast.Comment) (len=1 cap=1) {
//     (*ast.Comment)(0x14007f9bb18)({
//      Slash: (token.Pos) 4235128,
//      Text: (string) (len=17) "// Tests a panic."
//     })
//    }
//   }),
//   (*ast.CommentGroup)(0x14007f9bb78)({
//    List: ([]*ast.Comment) (len=2 cap=2) {
//     (*ast.Comment)(0x14007f9bb48)({
//      Slash: (token.Pos) 4235167,
//      Text: (string) (len=71) "// This panics, will log the message as \"fatal\" and exits the function."
//     }),
//     (*ast.Comment)(0x14007f9bb60)({
//      Slash: (token.Pos) 4235240,
//      Text: (string) (len=35) "// Generally, you should not panic."
//     })
//    }
//   }),
//   (*ast.CommentGroup)(0x14007f9bba8)({
//    List: ([]*ast.Comment) (len=1 cap=1) {
//     (*ast.Comment)(0x14007f9bb90)({
//      Slash: (token.Pos) 4235361,
//      Text: (string) (len=43) "// Tests an exit with a non-zero exit code."
//     })
//    }
//   }),
//   (*ast.CommentGroup)(0x14007f9bc20)({
//    List: ([]*ast.Comment) (len=4 cap=4) {
//     (*ast.Comment)(0x14007f9bbc0)({
//      Slash: (token.Pos) 4235425,
//      Text: (string) (len=74) "// If you need to exit prematurely without panicking, you can use os.Exit."
//     }),
//     (*ast.Comment)(0x14007f9bbd8)({
//      Slash: (token.Pos) 4235501,
//      Text: (string) (len=74) "// However, you cannot return any values from the function, so if you want"
//     }),
//     (*ast.Comment)(0x14007f9bbf0)({
//      Slash: (token.Pos) 4235577,
//      Text: (string) (len=68) "// to log an error message, you should do so before calling os.Exit."
//     }),
//     (*ast.Comment)(0x14007f9bc08)({
//      Slash: (token.Pos) 4235647,
//      Text: (string) (len=67) "// The exit code should be 0 for success, and non-zero for failure."
//     })
//    }
//   }),
//   (*ast.CommentGroup)(0x14007f9bc80)({
//    List: ([]*ast.Comment) (len=1 cap=1) {
//     (*ast.Comment)(0x14007f9bc68)({
//      Slash: (token.Pos) 4235819,
//      Text: (string) (len=37) "// Tests logging at different levels."
//     })
//    }
//   }),
//   (*ast.CommentGroup)(0x14007f9bcb0)({
//    List: ([]*ast.Comment) (len=1 cap=1) {
//     (*ast.Comment)(0x14007f9bc98)({
//      Slash: (token.Pos) 4235879,
//      Text: (string) (len=49) "// This is a simple log message. It has no level."
//     })
//    }
//   }),
//   (*ast.CommentGroup)(0x14007f9bcf8)({
//    List: ([]*ast.Comment) (len=1 cap=1) {
//     (*ast.Comment)(0x14007f9bce0)({
//      Slash: (token.Pos) 4235977,
//      Text: (string) (len=49) "// These messages are logged at different levels."
//     })
//    }
//   }),
//   (*ast.CommentGroup)(0x14007f9bd70)({
//    List: ([]*ast.Comment) (len=1 cap=1) {
//     (*ast.Comment)(0x14007f9bd58)({
//      Slash: (token.Pos) 4236158,
//      Text: (string) (len=67) "// This logs an error message, but allows the function to continue."
//     })
//    }
//   }),
//   (*ast.CommentGroup)(0x14007f9bdd0)({
//    List: ([]*ast.Comment) (len=1 cap=1) {
//     (*ast.Comment)(0x14007f9bdb8)({
//      Slash: (token.Pos) 4236438,
//      Text: (string) (len=52) "// You can also use Go's built-in printing commands."
//     })
//    }
//   }),
//   (*ast.CommentGroup)(0x14007f9be30)({
//    List: ([]*ast.Comment) (len=1 cap=1) {
//     (*ast.Comment)(0x14007f9be18)({
//      Slash: (token.Pos) 4236651,
//      Text: (string) (len=96) "// You can even just use stdout/stderr and the log level will be \"info\" or \"error\" respectively."
//     })
//    }
//   }),
//   (*ast.CommentGroup)(0x14007f9bf68)({
//    List: ([]*ast.Comment) (len=8 cap=8) {
//     (*ast.Comment)(0x14007f9bea8)({
//      Slash: (token.Pos) 4236893,
//      Text: (string) (len=121) "// NOTE: The main difference between using console functions and Go's built-in functions, is in how newlines are handled."
//     }),
//     (*ast.Comment)(0x14007f9bec0)({
//      Slash: (token.Pos) 4237016,
//      Text: (string) (len=106) "// - Using console functions, newlines are preserved and the entire message is logged as a single message."
//     }),
//     (*ast.Comment)(0x14007f9bed8)({
//      Slash: (token.Pos) 4237124,
//      Text: (string) (len=78) "// - Using Go's built-in functions, each line is logged as a separate message."
//     }),
//     (*ast.Comment)(0x14007f9bef0)({
//      Slash: (token.Pos) 4237204,
//      Text: (string) (len=2) "//"
//     }),
//     (*ast.Comment)(0x14007f9bf08)({
//      Slash: (token.Pos) 4237208,
//      Text: (string) (len=95) "// Thus, if you are logging data for debugging, we highly recommend using the console functions"
//     }),
//     (*ast.Comment)(0x14007f9bf20)({
//      Slash: (token.Pos) 4237305,
//      Text: (string) (len=53) "// to keep the data together in a single log message."
//     }),
//     (*ast.Comment)(0x14007f9bf38)({
//      Slash: (token.Pos) 4237360,
//      Text: (string) (len=2) "//"
//     }),
//     (*ast.Comment)(0x14007f9bf50)({
//      Slash: (token.Pos) 4237364,
//      Text: (string) (len=85) "// The console functions also allow you to better control the reported logging level."
//     })
//    }
//   })
//  },
//  GoVersion: (string) ""
// }

/*

GML: getFunctionsNeedingWrappers: f.Imports[0]=
(*ast.ImportSpec)(0x14007fdd890)({
 Path: (*ast.BasicLit)(0x14007fa3a20)({
  ValuePos: (token.Pos) 4231951,
  Kind: (token.Token) STRING,
  Value: (string) (len=8) "\"errors\""
 }),
 EndPos: (token.Pos) 0
})

GML: getFunctionsNeedingWrappers: imports["errors"] = "errors"

GML: getFunctionsNeedingWrappers: f.Imports[1]=
(*ast.ImportSpec)(0x14007fdd8c0)({
 Path: (*ast.BasicLit)(0x14007fa3a40)({
  ValuePos: (token.Pos) 4231961,
  Kind: (token.Token) STRING,
  Value: (string) (len=5) "\"fmt\""
 }),
 EndPos: (token.Pos) 0
})

GML: getFunctionsNeedingWrappers: imports["fmt"] = "fmt"

GML: getFunctionsNeedingWrappers: f.Imports[2]=
(*ast.ImportSpec)(0x14007fdd8f0)({
 Path: (*ast.BasicLit)(0x14007fa3a80)({
  ValuePos: (token.Pos) 4231968,
  Kind: (token.Token) STRING,
  Value: (string) (len=11) "\"math/rand\""
 }),
 EndPos: (token.Pos) 0
})

GML: getFunctionsNeedingWrappers: imports["rand"] = "math/rand"

GML: getFunctionsNeedingWrappers: f.Imports[3]=
(*ast.ImportSpec)(0x14007fdd920)({
 Path: (*ast.BasicLit)(0x14007fa3ac0)({
  ValuePos: (token.Pos) 4231981,
  Kind: (token.Token) STRING,
  Value: (string) (len=4) "\"os\""
 }),
 EndPos: (token.Pos) 0
})

GML: getFunctionsNeedingWrappers: imports["os"] = "os"

GML: getFunctionsNeedingWrappers: f.Imports[4]=
(*ast.ImportSpec)(0x14007fdd950)({
 Path: (*ast.BasicLit)(0x14007fa3ae0)({
  ValuePos: (token.Pos) 4231987,
  Kind: (token.Token) STRING,
  Value: (string) (len=6) "\"time\""
 }),
 EndPos: (token.Pos) 0
})

GML: getFunctionsNeedingWrappers: imports["time"] = "time"

GML: getFunctionsNeedingWrappers: f.Imports[5]=
(*ast.ImportSpec)(0x14007fdd980)({
 Path: (*ast.BasicLit)(0x14007fa3b00)({
  ValuePos: (token.Pos) 4231996,
  Kind: (token.Token) STRING,
  Value: (string) (len=50) "\"github.com/hypermodeinc/modus/sdk/go/pkg/console\""
 }),
 EndPos: (token.Pos) 0
})

GML: getFunctionsNeedingWrappers: imports["console"] = "github.com/hypermodeinc/modus/sdk/go/pkg/console"

GML: getFunctionsNeedingWrappers: f.Decls[0]=

GML: getFunctionsNeedingWrappers: f.Decls[1]=
(*ast.FuncDecl)(0x14007fdda10)({
 Name: (*ast.Ident)(0x14007fa3b20)(LogMessage),
 Type: (*ast.FuncType)(0x14007fa3c00)({
  Func: (token.Pos) 4232069,
  Params: (*ast.FieldList)(0x14007fdd9b0)({
   Opening: (token.Pos) 4232084,
   List: ([]*ast.Field) (len=1 cap=1) {
    (*ast.Field)(0x14007fe7680)({
     Names: ([]*ast.Ident) (len=1 cap=1) {
      (*ast.Ident)(0x14007fa3b40)(message)
     },
     Type: (*ast.Ident)(0x14007fa3b60)(string),
    })
   },
   Closing: (token.Pos) 4232099
  }),
 }),
})

GML: getFunctionsNeedingWrappers: results.append: info=
(*codegen.funcInfo)(0x14008480618)({
 function: (*ast.FuncDecl)(0x14007fdda10)({
  Name: (*ast.Ident)(0x14007fa3b20)(LogMessage),
  Type: (*ast.FuncType)(0x14007fa3c00)({
   Func: (token.Pos) 4232069,
   Params: (*ast.FieldList)(0x14007fdd9b0)({
    Opening: (token.Pos) 4232084,
    List: ([]*ast.Field) (len=1 cap=1) {
     (*ast.Field)(0x14007fe7680)({
      Names: ([]*ast.Ident) (len=1 cap=1) {
       (*ast.Ident)(0x14007fa3b40)(message)
      },
      Type: (*ast.Ident)(0x14007fa3b60)(string),
     })
    },
    Closing: (token.Pos) 4232099
   }),
  }),
 }),
 imports: (map[string]string) {
 },
 aliases: (map[string]string) <nil>
})

GML: getFunctionsNeedingWrappers: f.Decls[2]=
(*ast.FuncDecl)(0x14007fddb30)({
 Name: (*ast.Ident)(0x14007fa3c40)(Add),
 Type: (*ast.FuncType)(0x14007fa3d40)({
  Func: (token.Pos) 4232182,
  Params: (*ast.FieldList)(0x14007fdda70)({
   Opening: (token.Pos) 4232190,
   List: ([]*ast.Field) (len=1 cap=1) {
    (*ast.Field)(0x14007fe7700)({
     Names: ([]*ast.Ident) (len=2 cap=2) {
      (*ast.Ident)(0x14007fa3c60)(x),
      (*ast.Ident)(0x14007fa3c80)(y)
     },
     Type: (*ast.Ident)(0x14007fa3ca0)(int),
    })
   },
   Closing: (token.Pos) 4232199
  }),
  Results: (*ast.FieldList)(0x14007fddaa0)({
   Opening: (token.Pos) 0,
   List: ([]*ast.Field) (len=1 cap=1) {
    (*ast.Field)(0x14007fe7740)({
     Names: ([]*ast.Ident) <nil>,
     Type: (*ast.Ident)(0x14007fa3cc0)(int),
    })
   },
   Closing: (token.Pos) 0
  })
 }),
})

GML: getFunctionsNeedingWrappers: results.append: info=
(*codegen.funcInfo)(0x14008480690)({
 function: (*ast.FuncDecl)(0x14007fddb30)({
  Name: (*ast.Ident)(0x14007fa3c40)(Add),
  Type: (*ast.FuncType)(0x14007fa3d40)({
   Func: (token.Pos) 4232182,
   Params: (*ast.FieldList)(0x14007fdda70)({
    Opening: (token.Pos) 4232190,
    List: ([]*ast.Field) (len=1 cap=1) {
     (*ast.Field)(0x14007fe7700)({
      Names: ([]*ast.Ident) (len=2 cap=2) {
       (*ast.Ident)(0x14007fa3c60)(x),
       (*ast.Ident)(0x14007fa3c80)(y)
      },
      Type: (*ast.Ident)(0x14007fa3ca0)(int),
     })
    },
    Closing: (token.Pos) 4232199
   }),
   Results: (*ast.FieldList)(0x14007fddaa0)({
    Opening: (token.Pos) 0,
    List: ([]*ast.Field) (len=1 cap=1) {
     (*ast.Field)(0x14007fe7740)({
      Names: ([]*ast.Ident) <nil>,
      Type: (*ast.Ident)(0x14007fa3cc0)(int),
     })
    },
    Closing: (token.Pos) 0
   })
  }),
 }),
 imports: (map[string]string) {
 },
 aliases: (map[string]string) <nil>
})

GML: getFunctionsNeedingWrappers: f.Decls[3]=
(*ast.FuncDecl)(0x14007fddd10)({
 Name: (*ast.Ident)(0x14007fa3d60)(Add3),
 Type: (*ast.FuncType)(0x14007ffe000)({
  Func: (token.Pos) 4232314,
  Params: (*ast.FieldList)(0x14007fddb90)({
   Opening: (token.Pos) 4232323,
   List: ([]*ast.Field) (len=2 cap=2) {
    (*ast.Field)(0x14007fe77c0)({
     Names: ([]*ast.Ident) (len=2 cap=2) {
      (*ast.Ident)(0x14007fa3d80)(a),
      (*ast.Ident)(0x14007fa3da0)(b)
     },
     Type: (*ast.Ident)(0x14007fa3dc0)(int),
    }),
    (*ast.Field)(0x14007fe7800)({
     Names: ([]*ast.Ident) (len=1 cap=1) {
      (*ast.Ident)(0x14007fa3de0)(c)
     },
     Type: (*ast.StarExpr)(0x14007f9b500)({
      Star: (token.Pos) 4232336,
      X: (*ast.Ident)(0x14007fa3e00)(int)
     }),
    })
   },
   Closing: (token.Pos) 4232340
  }),
  Results: (*ast.FieldList)(0x14007fddbc0)({
   Opening: (token.Pos) 0,
   List: ([]*ast.Field) (len=1 cap=1) {
    (*ast.Field)(0x14007fe7840)({
     Names: ([]*ast.Ident) <nil>,
     Type: (*ast.Ident)(0x14007fa3e20)(int),
    })
   },
   Closing: (token.Pos) 0
  })
 }),
})

GML: getFunctionsNeedingWrappers: results.append: info=
(*codegen.funcInfo)(0x14008480708)({
 function: (*ast.FuncDecl)(0x14007fddd10)({
  Name: (*ast.Ident)(0x14007fa3d60)(Add3),
  Type: (*ast.FuncType)(0x14007ffe000)({
   Func: (token.Pos) 4232314,
   Params: (*ast.FieldList)(0x14007fddb90)({
    Opening: (token.Pos) 4232323,
    List: ([]*ast.Field) (len=2 cap=2) {
     (*ast.Field)(0x14007fe77c0)({
      Names: ([]*ast.Ident) (len=2 cap=2) {
       (*ast.Ident)(0x14007fa3d80)(a),
       (*ast.Ident)(0x14007fa3da0)(b)
      },
      Type: (*ast.Ident)(0x14007fa3dc0)(int),
     }),
     (*ast.Field)(0x14007fe7800)({
      Names: ([]*ast.Ident) (len=1 cap=1) {
       (*ast.Ident)(0x14007fa3de0)(c)
      },
      Type: (*ast.StarExpr)(0x14007f9b500)({
       Star: (token.Pos) 4232336,
       X: (*ast.Ident)(0x14007fa3e00)(int)
      }),
     })
    },
    Closing: (token.Pos) 4232340
   }),
   Results: (*ast.FieldList)(0x14007fddbc0)({
    Opening: (token.Pos) 0,
    List: ([]*ast.Field) (len=1 cap=1) {
     (*ast.Field)(0x14007fe7840)({
      Names: ([]*ast.Ident) <nil>,
      Type: (*ast.Ident)(0x14007fa3e20)(int),
     })
    },
    Closing: (token.Pos) 0
   })
  }),
 }),
 imports: (map[string]string) {
 },
 aliases: (map[string]string) <nil>
})

GML: getFunctionsNeedingWrappers: f.Decls[4]=
(*ast.FuncDecl)(0x14007fdde00)({
 Name: (*ast.Ident)(0x14007ffe020)(AddN),
 Type: (*ast.FuncType)(0x14007ffe220)({
  Func: (token.Pos) 4232467,
  Params: (*ast.FieldList)(0x14007fddd40)({
   Opening: (token.Pos) 4232476,
   List: ([]*ast.Field) (len=1 cap=1) {
    (*ast.Field)(0x14007fe7900)({
     Names: ([]*ast.Ident) (len=1 cap=1) {
      (*ast.Ident)(0x14007ffe040)(args)
     },
     Type: (*ast.Ellipsis)(0x14007f9b560)({
      Ellipsis: (token.Pos) 4232482,
      Elt: (*ast.Ident)(0x14007ffe060)(int)
     }),
    })
   },
   Closing: (token.Pos) 4232488
  }),
  Results: (*ast.FieldList)(0x14007fddd70)({
   Opening: (token.Pos) 0,
   List: ([]*ast.Field) (len=1 cap=1) {
    (*ast.Field)(0x14007fe7940)({
     Names: ([]*ast.Ident) <nil>,
     Type: (*ast.Ident)(0x14007ffe080)(int),
    })
   },
   Closing: (token.Pos) 0
  })
 }),
})

GML: getFunctionsNeedingWrappers: results.append: info=
(*codegen.funcInfo)(0x14008480798)({
 function: (*ast.FuncDecl)(0x14007fdde00)({
  Name: (*ast.Ident)(0x14007ffe020)(AddN),
  Type: (*ast.FuncType)(0x14007ffe220)({
   Func: (token.Pos) 4232467,
   Params: (*ast.FieldList)(0x14007fddd40)({
    Opening: (token.Pos) 4232476,
    List: ([]*ast.Field) (len=1 cap=1) {
     (*ast.Field)(0x14007fe7900)({
      Names: ([]*ast.Ident) (len=1 cap=1) {
       (*ast.Ident)(0x14007ffe040)(args)
      },
      Type: (*ast.Ellipsis)(0x14007f9b560)({
       Ellipsis: (token.Pos) 4232482,
       Elt: (*ast.Ident)(0x14007ffe060)(int)
      }),
     })
    },
    Closing: (token.Pos) 4232488
   }),
   Results: (*ast.FieldList)(0x14007fddd70)({
    Opening: (token.Pos) 0,
    List: ([]*ast.Field) (len=1 cap=1) {
     (*ast.Field)(0x14007fe7940)({
      Names: ([]*ast.Ident) <nil>,
      Type: (*ast.Ident)(0x14007ffe080)(int),
     })
    },
    Closing: (token.Pos) 0
   })
  }),
 }),
 imports: (map[string]string) {
 },
 aliases: (map[string]string) <nil>
})

GML: getFunctionsNeedingWrappers: f.Decls[5]=

GML: getFunctionsNeedingWrappers: f.Decls[6]=
(*ast.FuncDecl)(0x14007fddec0)({
 Name: (*ast.Ident)(0x14007ffe2a0)(GetCurrentTime),
 Type: (*ast.FuncType)(0x14007ffe340)({
  Func: (token.Pos) 4232674,
  Params: (*ast.FieldList)(0x14007fdde30)({
   Opening: (token.Pos) 4232693,
   List: ([]*ast.Field) <nil>,
   Closing: (token.Pos) 4232694
  }),
  Results: (*ast.FieldList)(0x14007fdde60)({
   Opening: (token.Pos) 0,
   List: ([]*ast.Field) (len=1 cap=1) {
    (*ast.Field)(0x14007fe7ac0)({
     Names: ([]*ast.Ident) <nil>,
     Type: (*ast.SelectorExpr)(0x14007f9b608)({
      X: (*ast.Ident)(0x14007ffe2c0)(time),
      Sel: (*ast.Ident)(0x14007ffe2e0)(Time)
     }),
    })
   },
   Closing: (token.Pos) 0
  })
 }),
})

GML: getFunctionsNeedingWrappers: results.append: info=
(*codegen.funcInfo)(0x14008480840)({
 function: (*ast.FuncDecl)(0x14007fddec0)({
  Name: (*ast.Ident)(0x14007ffe2a0)(GetCurrentTime),
  Type: (*ast.FuncType)(0x14007ffe340)({
   Func: (token.Pos) 4232674,
   Params: (*ast.FieldList)(0x14007fdde30)({
    Opening: (token.Pos) 4232693,
    List: ([]*ast.Field) <nil>,
    Closing: (token.Pos) 4232694
   }),
   Results: (*ast.FieldList)(0x14007fdde60)({
    Opening: (token.Pos) 0,
    List: ([]*ast.Field) (len=1 cap=1) {
     (*ast.Field)(0x14007fe7ac0)({
      Names: ([]*ast.Ident) <nil>,
      Type: (*ast.SelectorExpr)(0x14007f9b608)({
       X: (*ast.Ident)(0x14007ffe2c0)(time),
       Sel: (*ast.Ident)(0x14007ffe2e0)(Time)
      }),
     })
    },
    Closing: (token.Pos) 0
   })
  }),
 }),
 imports: (map[string]string) (len=1) {
  (string) (len=4) "time": (string) (len=4) "time"
 },
 aliases: (map[string]string) <nil>
})

GML: getFunctionsNeedingWrappers: f.Decls[7]=
(*ast.FuncDecl)(0x14008478000)({
 Name: (*ast.Ident)(0x14007ffe360)(GetCurrentTimeFormatted),
 Type: (*ast.FuncType)(0x14007ffe440)({
  Func: (token.Pos) 4232780,
  Params: (*ast.FieldList)(0x14007fddef0)({
   Opening: (token.Pos) 4232808,
   List: ([]*ast.Field) <nil>,
   Closing: (token.Pos) 4232809
  }),
  Results: (*ast.FieldList)(0x14007fddf20)({
   Opening: (token.Pos) 0,
   List: ([]*ast.Field) (len=1 cap=1) {
    (*ast.Field)(0x14007fe7b40)({
     Names: ([]*ast.Ident) <nil>,
     Type: (*ast.Ident)(0x14007ffe380)(string),
    })
   },
   Closing: (token.Pos) 0
  })
 }),
})

GML: getFunctionsNeedingWrappers: results.append: info=
(*codegen.funcInfo)(0x140084808d0)({
 function: (*ast.FuncDecl)(0x14008478000)({
  Name: (*ast.Ident)(0x14007ffe360)(GetCurrentTimeFormatted),
  Type: (*ast.FuncType)(0x14007ffe440)({
   Func: (token.Pos) 4232780,
   Params: (*ast.FieldList)(0x14007fddef0)({
    Opening: (token.Pos) 4232808,
    List: ([]*ast.Field) <nil>,
    Closing: (token.Pos) 4232809
   }),
   Results: (*ast.FieldList)(0x14007fddf20)({
    Opening: (token.Pos) 0,
    List: ([]*ast.Field) (len=1 cap=1) {
     (*ast.Field)(0x14007fe7b40)({
      Names: ([]*ast.Ident) <nil>,
      Type: (*ast.Ident)(0x14007ffe380)(string),
     })
    },
    Closing: (token.Pos) 0
   })
  }),
 }),
 imports: (map[string]string) {
 },
 aliases: (map[string]string) <nil>
})

GML: getFunctionsNeedingWrappers: f.Decls[8]=
(*ast.FuncDecl)(0x14008478150)({
 Name: (*ast.Ident)(0x14007ffe460)(GetFullName),
 Type: (*ast.FuncType)(0x14007ffe580)({
  Func: (token.Pos) 4232939,
  Params: (*ast.FieldList)(0x14008478060)({
   Opening: (token.Pos) 4232955,
   List: ([]*ast.Field) (len=1 cap=1) {
    (*ast.Field)(0x14007fe7c00)({
     Names: ([]*ast.Ident) (len=2 cap=2) {
      (*ast.Ident)(0x14007ffe480)(firstName),
      (*ast.Ident)(0x14007ffe4a0)(lastName)
     },
     Type: (*ast.Ident)(0x14007ffe4c0)(string),
    })
   },
   Closing: (token.Pos) 4232982
  }),
  Results: (*ast.FieldList)(0x14008478090)({
   Opening: (token.Pos) 0,
   List: ([]*ast.Field) (len=1 cap=1) {
    (*ast.Field)(0x14007fe7c40)({
     Names: ([]*ast.Ident) <nil>,
     Type: (*ast.Ident)(0x14007ffe4e0)(string),
    })
   },
   Closing: (token.Pos) 0
  })
 }),
})

GML: getFunctionsNeedingWrappers: results.append: info=
(*codegen.funcInfo)(0x14008480948)({
 function: (*ast.FuncDecl)(0x14008478150)({
  Name: (*ast.Ident)(0x14007ffe460)(GetFullName),
  Type: (*ast.FuncType)(0x14007ffe580)({
   Func: (token.Pos) 4232939,
   Params: (*ast.FieldList)(0x14008478060)({
    Opening: (token.Pos) 4232955,
    List: ([]*ast.Field) (len=1 cap=1) {
     (*ast.Field)(0x14007fe7c00)({
      Names: ([]*ast.Ident) (len=2 cap=2) {
       (*ast.Ident)(0x14007ffe480)(firstName),
       (*ast.Ident)(0x14007ffe4a0)(lastName)
      },
      Type: (*ast.Ident)(0x14007ffe4c0)(string),
     })
    },
    Closing: (token.Pos) 4232982
   }),
   Results: (*ast.FieldList)(0x14008478090)({
    Opening: (token.Pos) 0,
    List: ([]*ast.Field) (len=1 cap=1) {
     (*ast.Field)(0x14007fe7c40)({
      Names: ([]*ast.Ident) <nil>,
      Type: (*ast.Ident)(0x14007ffe4e0)(string),
     })
    },
    Closing: (token.Pos) 0
   })
  }),
 }),
 imports: (map[string]string) {
 },
 aliases: (map[string]string) <nil>
})

GML: getFunctionsNeedingWrappers: f.Decls[9]=
(*ast.FuncDecl)(0x14008478300)({
 Name: (*ast.Ident)(0x14007ffe5a0)(SayHello),
 Type: (*ast.FuncType)(0x14007ffe720)({
  Func: (token.Pos) 4233132,
  Params: (*ast.FieldList)(0x14008478180)({
   Opening: (token.Pos) 4233145,
   List: ([]*ast.Field) (len=1 cap=1) {
    (*ast.Field)(0x14007fe7c80)({
     Names: ([]*ast.Ident) (len=1 cap=1) {
      (*ast.Ident)(0x14007ffe5c0)(name)
     },
     Type: (*ast.StarExpr)(0x14007f9b710)({
      Star: (token.Pos) 4233151,
      X: (*ast.Ident)(0x14007ffe5e0)(string)
     }),
    })
   },
   Closing: (token.Pos) 4233158
  }),
  Results: (*ast.FieldList)(0x140084781b0)({
   Opening: (token.Pos) 0,
   List: ([]*ast.Field) (len=1 cap=1) {
    (*ast.Field)(0x14007fe7cc0)({
     Names: ([]*ast.Ident) <nil>,
     Type: (*ast.Ident)(0x14007ffe600)(string),
    })
   },
   Closing: (token.Pos) 0
  })
 }),
})

GML: getFunctionsNeedingWrappers: results.append: info=
(*codegen.funcInfo)(0x140084809d8)({
 function: (*ast.FuncDecl)(0x14008478300)({
  Name: (*ast.Ident)(0x14007ffe5a0)(SayHello),
  Type: (*ast.FuncType)(0x14007ffe720)({
   Func: (token.Pos) 4233132,
   Params: (*ast.FieldList)(0x14008478180)({
    Opening: (token.Pos) 4233145,
    List: ([]*ast.Field) (len=1 cap=1) {
     (*ast.Field)(0x14007fe7c80)({
      Names: ([]*ast.Ident) (len=1 cap=1) {
       (*ast.Ident)(0x14007ffe5c0)(name)
      },
      Type: (*ast.StarExpr)(0x14007f9b710)({
       Star: (token.Pos) 4233151,
       X: (*ast.Ident)(0x14007ffe5e0)(string)
      }),
     })
    },
    Closing: (token.Pos) 4233158
   }),
   Results: (*ast.FieldList)(0x140084781b0)({
    Opening: (token.Pos) 0,
    List: ([]*ast.Field) (len=1 cap=1) {
     (*ast.Field)(0x14007fe7cc0)({
      Names: ([]*ast.Ident) <nil>,
      Type: (*ast.Ident)(0x14007ffe600)(string),
     })
    },
    Closing: (token.Pos) 0
   })
  }),
 }),
 imports: (map[string]string) {
 },
 aliases: (map[string]string) <nil>
})

GML: getFunctionsNeedingWrappers: f.Decls[10]=

GML: getFunctionsNeedingWrappers: f.Decls[11]=
(*ast.FuncDecl)(0x14008478480)({
 Name: (*ast.Ident)(0x14007ffe8a0)(GetPerson),
 Type: (*ast.FuncType)(0x14007ffea00)({
  Func: (token.Pos) 4233521,
  Params: (*ast.FieldList)(0x14008478360)({
   Opening: (token.Pos) 4233535,
   List: ([]*ast.Field) <nil>,
   Closing: (token.Pos) 4233536
  }),
  Results: (*ast.FieldList)(0x14008478390)({
   Opening: (token.Pos) 0,
   List: ([]*ast.Field) (len=1 cap=1) {
    (*ast.Field)(0x14007fe7e80)({
     Names: ([]*ast.Ident) <nil>,
     Type: (*ast.Ident)(0x14007ffe8c0)(Person),
    })
   },
   Closing: (token.Pos) 0
  })
 }),
})

GML: getFunctionsNeedingWrappers: results.append: info=
(*codegen.funcInfo)(0x14008480a98)({
 function: (*ast.FuncDecl)(0x14008478480)({
  Name: (*ast.Ident)(0x14007ffe8a0)(GetPerson),
  Type: (*ast.FuncType)(0x14007ffea00)({
   Func: (token.Pos) 4233521,
   Params: (*ast.FieldList)(0x14008478360)({
    Opening: (token.Pos) 4233535,
    List: ([]*ast.Field) <nil>,
    Closing: (token.Pos) 4233536
   }),
   Results: (*ast.FieldList)(0x14008478390)({
    Opening: (token.Pos) 0,
    List: ([]*ast.Field) (len=1 cap=1) {
     (*ast.Field)(0x14007fe7e80)({
      Names: ([]*ast.Ident) <nil>,
      Type: (*ast.Ident)(0x14007ffe8c0)(Person),
     })
    },
    Closing: (token.Pos) 0
   })
  }),
 }),
 imports: (map[string]string) {
 },
 aliases: (map[string]string) <nil>
})

GML: getFunctionsNeedingWrappers: f.Decls[12]=
(*ast.FuncDecl)(0x14008478570)({
 Name: (*ast.Ident)(0x14007ffea20)(GetRandomPerson),
 Type: (*ast.FuncType)(0x14007ffebc0)({
  Func: (token.Pos) 4233681,
  Params: (*ast.FieldList)(0x140084784b0)({
   Opening: (token.Pos) 4233701,
   List: ([]*ast.Field) <nil>,
   Closing: (token.Pos) 4233702
  }),
  Results: (*ast.FieldList)(0x140084784e0)({
   Opening: (token.Pos) 0,
   List: ([]*ast.Field) (len=1 cap=1) {
    (*ast.Field)(0x14007fe7f40)({
     Names: ([]*ast.Ident) <nil>,
     Type: (*ast.Ident)(0x14007ffea40)(Person),
    })
   },
   Closing: (token.Pos) 0
  })
 }),
})

GML: getFunctionsNeedingWrappers: results.append: info=
(*codegen.funcInfo)(0x14008480b10)({
 function: (*ast.FuncDecl)(0x14008478570)({
  Name: (*ast.Ident)(0x14007ffea20)(GetRandomPerson),
  Type: (*ast.FuncType)(0x14007ffebc0)({
   Func: (token.Pos) 4233681,
   Params: (*ast.FieldList)(0x140084784b0)({
    Opening: (token.Pos) 4233701,
    List: ([]*ast.Field) <nil>,
    Closing: (token.Pos) 4233702
   }),
   Results: (*ast.FieldList)(0x140084784e0)({
    Opening: (token.Pos) 0,
    List: ([]*ast.Field) (len=1 cap=1) {
     (*ast.Field)(0x14007fe7f40)({
      Names: ([]*ast.Ident) <nil>,
      Type: (*ast.Ident)(0x14007ffea40)(Person),
     })
    },
    Closing: (token.Pos) 0
   })
  }),
 }),
 imports: (map[string]string) {
 },
 aliases: (map[string]string) <nil>
})

GML: getFunctionsNeedingWrappers: f.Decls[13]=
(*ast.FuncDecl)(0x14008478840)({
 Name: (*ast.Ident)(0x14007ffebe0)(GetPeople),
 Type: (*ast.FuncType)(0x14007ffef20)({
  Func: (token.Pos) 4233812,
  Params: (*ast.FieldList)(0x140084785a0)({
   Opening: (token.Pos) 4233826,
   List: ([]*ast.Field) <nil>,
   Closing: (token.Pos) 4233827
  }),
  Results: (*ast.FieldList)(0x14008478600)({
   Opening: (token.Pos) 0,
   List: ([]*ast.Field) (len=1 cap=1) {
    (*ast.Field)(0x1400847a180)({
     Names: ([]*ast.Ident) <nil>,
     Type: (*ast.ArrayType)(0x140084785d0)({
      Lbrack: (token.Pos) 4233829,
      Len: (ast.Expr) <nil>,
      Elt: (*ast.Ident)(0x14007ffec00)(Person)
     }),
    })
   },
   Closing: (token.Pos) 0
  })
 }),
})

GML: getFunctionsNeedingWrappers: results.append: info=
(*codegen.funcInfo)(0x14008480ba0)({
 function: (*ast.FuncDecl)(0x14008478840)({
  Name: (*ast.Ident)(0x14007ffebe0)(GetPeople),
  Type: (*ast.FuncType)(0x14007ffef20)({
   Func: (token.Pos) 4233812,
   Params: (*ast.FieldList)(0x140084785a0)({
    Opening: (token.Pos) 4233826,
    List: ([]*ast.Field) <nil>,
    Closing: (token.Pos) 4233827
   }),
   Results: (*ast.FieldList)(0x14008478600)({
    Opening: (token.Pos) 0,
    List: ([]*ast.Field) (len=1 cap=1) {
     (*ast.Field)(0x1400847a180)({
      Names: ([]*ast.Ident) <nil>,
      Type: (*ast.ArrayType)(0x140084785d0)({
       Lbrack: (token.Pos) 4233829,
       Len: (ast.Expr) <nil>,
       Elt: (*ast.Ident)(0x14007ffec00)(Person)
      }),
     })
    },
    Closing: (token.Pos) 0
   })
  }),
 }),
 imports: (map[string]string) {
 },
 aliases: (map[string]string) <nil>
})

GML: getFunctionsNeedingWrappers: f.Decls[14]=
(*ast.FuncDecl)(0x14008478930)({
 Name: (*ast.Ident)(0x14007ffef40)(GetNameAndAge),
 Type: (*ast.FuncType)(0x14007fff180)({
  Func: (token.Pos) 4234120,
  Params: (*ast.FieldList)(0x14008478870)({
   Opening: (token.Pos) 4234138,
   List: ([]*ast.Field) <nil>,
   Closing: (token.Pos) 4234139
  }),
  Results: (*ast.FieldList)(0x140084788d0)({
   Opening: (token.Pos) 4234141,
   List: ([]*ast.Field) (len=2 cap=2) {
    (*ast.Field)(0x1400847a3c0)({
     Names: ([]*ast.Ident) (len=1 cap=1) {
      (*ast.Ident)(0x14007ffef60)(name)
     },
     Type: (*ast.Ident)(0x14007ffef80)(string),
    }),
    (*ast.Field)(0x1400847a400)({
     Names: ([]*ast.Ident) (len=1 cap=1) {
      (*ast.Ident)(0x14007ffefa0)(age)
     },
     Type: (*ast.Ident)(0x14007ffefc0)(int),
    })
   },
   Closing: (token.Pos) 4234162
  })
 }),
})

GML: getFunctionsNeedingWrappers: results.append: info=
(*codegen.funcInfo)(0x14008480c30)({
 function: (*ast.FuncDecl)(0x14008478930)({
  Name: (*ast.Ident)(0x14007ffef40)(GetNameAndAge),
  Type: (*ast.FuncType)(0x14007fff180)({
   Func: (token.Pos) 4234120,
   Params: (*ast.FieldList)(0x14008478870)({
    Opening: (token.Pos) 4234138,
    List: ([]*ast.Field) <nil>,
    Closing: (token.Pos) 4234139
   }),
   Results: (*ast.FieldList)(0x140084788d0)({
    Opening: (token.Pos) 4234141,
    List: ([]*ast.Field) (len=2 cap=2) {
     (*ast.Field)(0x1400847a3c0)({
      Names: ([]*ast.Ident) (len=1 cap=1) {
       (*ast.Ident)(0x14007ffef60)(name)
      },
      Type: (*ast.Ident)(0x14007ffef80)(string),
     }),
     (*ast.Field)(0x1400847a400)({
      Names: ([]*ast.Ident) (len=1 cap=1) {
       (*ast.Ident)(0x14007ffefa0)(age)
      },
      Type: (*ast.Ident)(0x14007ffefc0)(int),
     })
    },
    Closing: (token.Pos) 4234162
   })
  }),
 }),
 imports: (map[string]string) {
 },
 aliases: (map[string]string) <nil>
})

GML: getFunctionsNeedingWrappers: f.Decls[15]=
(*ast.FuncDecl)(0x14008478ab0)({
 Name: (*ast.Ident)(0x14007fff1a0)(TestNormalError),
 Type: (*ast.FuncType)(0x14007fff460)({
  Func: (token.Pos) 4234268,
  Params: (*ast.FieldList)(0x14008478960)({
   Opening: (token.Pos) 4234288,
   List: ([]*ast.Field) (len=1 cap=1) {
    (*ast.Field)(0x1400847a500)({
     Names: ([]*ast.Ident) (len=1 cap=1) {
      (*ast.Ident)(0x14007fff1c0)(input)
     },
     Type: (*ast.Ident)(0x14007fff1e0)(string),
    })
   },
   Closing: (token.Pos) 4234301
  }),
  Results: (*ast.FieldList)(0x140084789c0)({
   Opening: (token.Pos) 4234303,
   List: ([]*ast.Field) (len=2 cap=2) {
    (*ast.Field)(0x1400847a540)({
     Names: ([]*ast.Ident) <nil>,
     Type: (*ast.Ident)(0x14007fff200)(string),
    }),
    (*ast.Field)(0x1400847a580)({
     Names: ([]*ast.Ident) <nil>,
     Type: (*ast.Ident)(0x14007fff220)(error),
    })
   },
   Closing: (token.Pos) 4234317
  })
 }),
})

GML: getFunctionsNeedingWrappers: results.append: info=
(*codegen.funcInfo)(0x14008480cc0)({
 function: (*ast.FuncDecl)(0x14008478ab0)({
  Name: (*ast.Ident)(0x14007fff1a0)(TestNormalError),
  Type: (*ast.FuncType)(0x14007fff460)({
   Func: (token.Pos) 4234268,
   Params: (*ast.FieldList)(0x14008478960)({
    Opening: (token.Pos) 4234288,
    List: ([]*ast.Field) (len=1 cap=1) {
     (*ast.Field)(0x1400847a500)({
      Names: ([]*ast.Ident) (len=1 cap=1) {
       (*ast.Ident)(0x14007fff1c0)(input)
      },
      Type: (*ast.Ident)(0x14007fff1e0)(string),
     })
    },
    Closing: (token.Pos) 4234301
   }),
   Results: (*ast.FieldList)(0x140084789c0)({
    Opening: (token.Pos) 4234303,
    List: ([]*ast.Field) (len=2 cap=2) {
     (*ast.Field)(0x1400847a540)({
      Names: ([]*ast.Ident) <nil>,
      Type: (*ast.Ident)(0x14007fff200)(string),
     }),
     (*ast.Field)(0x1400847a580)({
      Names: ([]*ast.Ident) <nil>,
      Type: (*ast.Ident)(0x14007fff220)(error),
     })
    },
    Closing: (token.Pos) 4234317
   })
  }),
 }),
 imports: (map[string]string) {
 },
 aliases: (map[string]string) <nil>
})

GML: getFunctionsNeedingWrappers: f.Decls[16]=
(*ast.FuncDecl)(0x14008478c00)({
 Name: (*ast.Ident)(0x14007fff480)(TestAlternativeError),
 Type: (*ast.FuncType)(0x14007fff6c0)({
  Func: (token.Pos) 4234821,
  Params: (*ast.FieldList)(0x14008478ae0)({
   Opening: (token.Pos) 4234846,
   List: ([]*ast.Field) (len=1 cap=1) {
    (*ast.Field)(0x1400847a700)({
     Names: ([]*ast.Ident) (len=1 cap=1) {
      (*ast.Ident)(0x14007fff4a0)(input)
     },
     Type: (*ast.Ident)(0x14007fff4c0)(string),
    })
   },
   Closing: (token.Pos) 4234859
  }),
  Results: (*ast.FieldList)(0x14008478b10)({
   Opening: (token.Pos) 0,
   List: ([]*ast.Field) (len=1 cap=1) {
    (*ast.Field)(0x1400847a740)({
     Names: ([]*ast.Ident) <nil>,
     Type: (*ast.Ident)(0x14007fff4e0)(string),
    })
   },
   Closing: (token.Pos) 0
  })
 }),
})

GML: getFunctionsNeedingWrappers: results.append: info=
(*codegen.funcInfo)(0x14008480d50)({
 function: (*ast.FuncDecl)(0x14008478c00)({
  Name: (*ast.Ident)(0x14007fff480)(TestAlternativeError),
  Type: (*ast.FuncType)(0x14007fff6c0)({
   Func: (token.Pos) 4234821,
   Params: (*ast.FieldList)(0x14008478ae0)({
    Opening: (token.Pos) 4234846,
    List: ([]*ast.Field) (len=1 cap=1) {
     (*ast.Field)(0x1400847a700)({
      Names: ([]*ast.Ident) (len=1 cap=1) {
       (*ast.Ident)(0x14007fff4a0)(input)
      },
      Type: (*ast.Ident)(0x14007fff4c0)(string),
     })
    },
    Closing: (token.Pos) 4234859
   }),
   Results: (*ast.FieldList)(0x14008478b10)({
    Opening: (token.Pos) 0,
    List: ([]*ast.Field) (len=1 cap=1) {
     (*ast.Field)(0x1400847a740)({
      Names: ([]*ast.Ident) <nil>,
      Type: (*ast.Ident)(0x14007fff4e0)(string),
     })
    },
    Closing: (token.Pos) 0
   })
  }),
 }),
 imports: (map[string]string) {
 },
 aliases: (map[string]string) <nil>
})

GML: getFunctionsNeedingWrappers: f.Decls[17]=
(*ast.FuncDecl)(0x14008478c90)({
 Name: (*ast.Ident)(0x14007fff6e0)(TestPanic),
 Type: (*ast.FuncType)(0x14007fff740)({
  Func: (token.Pos) 4235146,
  Params: (*ast.FieldList)(0x14008478c30)({
   Opening: (token.Pos) 4235160,
   List: ([]*ast.Field) <nil>,
   Closing: (token.Pos) 4235161
  }),
 }),
})

GML: getFunctionsNeedingWrappers: results.append: info=
(*codegen.funcInfo)(0x14008480de0)({
 function: (*ast.FuncDecl)(0x14008478c90)({
  Name: (*ast.Ident)(0x14007fff6e0)(TestPanic),
  Type: (*ast.FuncType)(0x14007fff740)({
   Func: (token.Pos) 4235146,
   Params: (*ast.FieldList)(0x14008478c30)({
    Opening: (token.Pos) 4235160,
    List: ([]*ast.Field) <nil>,
    Closing: (token.Pos) 4235161
   }),
  }),
 }),
 imports: (map[string]string) {
 },
 aliases: (map[string]string) <nil>
})

GML: getFunctionsNeedingWrappers: f.Decls[18]=
(*ast.FuncDecl)(0x14008478d20)({
 Name: (*ast.Ident)(0x14007fff760)(TestExit),
 Type: (*ast.FuncType)(0x14007fff8c0)({
  Func: (token.Pos) 4235405,
  Params: (*ast.FieldList)(0x14008478cc0)({
   Opening: (token.Pos) 4235418,
   List: ([]*ast.Field) <nil>,
   Closing: (token.Pos) 4235419
  }),
 }),
})

GML: getFunctionsNeedingWrappers: results.append: info=
(*codegen.funcInfo)(0x14008480e58)({
 function: (*ast.FuncDecl)(0x14008478d20)({
  Name: (*ast.Ident)(0x14007fff760)(TestExit),
  Type: (*ast.FuncType)(0x14007fff8c0)({
   Func: (token.Pos) 4235405,
   Params: (*ast.FieldList)(0x14008478cc0)({
    Opening: (token.Pos) 4235418,
    List: ([]*ast.Field) <nil>,
    Closing: (token.Pos) 4235419
   }),
  }),
 }),
 imports: (map[string]string) {
 },
 aliases: (map[string]string) <nil>
})

GML: getFunctionsNeedingWrappers: f.Decls[19]=
(*ast.FuncDecl)(0x14008478db0)({
 Name: (*ast.Ident)(0x14007fff8e0)(TestLogging),
 Type: (*ast.FuncType)(0x14007fffe40)({
  Func: (token.Pos) 4235857,
  Params: (*ast.FieldList)(0x14008478d50)({
   Opening: (token.Pos) 4235873,
   List: ([]*ast.Field) <nil>,
   Closing: (token.Pos) 4235874
  }),
 }),
})

GML: getFunctionsNeedingWrappers: results.append: info=
(*codegen.funcInfo)(0x14008480ed0)({
 function: (*ast.FuncDecl)(0x14008478db0)({
  Name: (*ast.Ident)(0x14007fff8e0)(TestLogging),
  Type: (*ast.FuncType)(0x14007fffe40)({
   Func: (token.Pos) 4235857,
   Params: (*ast.FieldList)(0x14008478d50)({
    Opening: (token.Pos) 4235873,
    List: ([]*ast.Field) <nil>,
    Closing: (token.Pos) 4235874
   }),
  }),
 }),
 imports: (map[string]string) {
 },
 aliases: (map[string]string) <nil>
})

GML: PreProcess: functions=
([]*codegen.funcInfo) (len=17 cap=32) {
 (*codegen.funcInfo)(0x14008480618)({
  function: (*ast.FuncDecl)(0x14007fdda10)({
   Name: (*ast.Ident)(0x14007fa3b20)(LogMessage),
   Type: (*ast.FuncType)(0x14007fa3c00)({
    Func: (token.Pos) 4232069,
    Params: (*ast.FieldList)(0x14007fdd9b0)({
     Opening: (token.Pos) 4232084,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x14007fe7680)({
       Names: ([]*ast.Ident) (len=1 cap=1) {
        (*ast.Ident)(0x14007fa3b40)(message)
       },
       Type: (*ast.Ident)(0x14007fa3b60)(string),
      })
     },
     Closing: (token.Pos) 4232099
    }),
   }),
  }),
  imports: (map[string]string) {
  },
  aliases: (map[string]string) <nil>
 }),
 (*codegen.funcInfo)(0x14008480690)({
  function: (*ast.FuncDecl)(0x14007fddb30)({
   Name: (*ast.Ident)(0x14007fa3c40)(Add),
   Type: (*ast.FuncType)(0x14007fa3d40)({
    Func: (token.Pos) 4232182,
    Params: (*ast.FieldList)(0x14007fdda70)({
     Opening: (token.Pos) 4232190,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x14007fe7700)({
       Names: ([]*ast.Ident) (len=2 cap=2) {
        (*ast.Ident)(0x14007fa3c60)(x),
        (*ast.Ident)(0x14007fa3c80)(y)
       },
       Type: (*ast.Ident)(0x14007fa3ca0)(int),
      })
     },
     Closing: (token.Pos) 4232199
    }),
    Results: (*ast.FieldList)(0x14007fddaa0)({
     Opening: (token.Pos) 0,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x14007fe7740)({
       Names: ([]*ast.Ident) <nil>,
       Type: (*ast.Ident)(0x14007fa3cc0)(int),
      })
     },
     Closing: (token.Pos) 0
    })
   }),
  }),
  imports: (map[string]string) {
  },
  aliases: (map[string]string) <nil>
 }),
 (*codegen.funcInfo)(0x14008480708)({
  function: (*ast.FuncDecl)(0x14007fddd10)({
   Name: (*ast.Ident)(0x14007fa3d60)(Add3),
   Type: (*ast.FuncType)(0x14007ffe000)({
    Func: (token.Pos) 4232314,
    Params: (*ast.FieldList)(0x14007fddb90)({
     Opening: (token.Pos) 4232323,
     List: ([]*ast.Field) (len=2 cap=2) {
      (*ast.Field)(0x14007fe77c0)({
       Names: ([]*ast.Ident) (len=2 cap=2) {
        (*ast.Ident)(0x14007fa3d80)(a),
        (*ast.Ident)(0x14007fa3da0)(b)
       },
       Type: (*ast.Ident)(0x14007fa3dc0)(int),
      }),
      (*ast.Field)(0x14007fe7800)({
       Names: ([]*ast.Ident) (len=1 cap=1) {
        (*ast.Ident)(0x14007fa3de0)(c)
       },
       Type: (*ast.StarExpr)(0x14007f9b500)({
        Star: (token.Pos) 4232336,
        X: (*ast.Ident)(0x14007fa3e00)(int)
       }),
      })
     },
     Closing: (token.Pos) 4232340
    }),
    Results: (*ast.FieldList)(0x14007fddbc0)({
     Opening: (token.Pos) 0,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x14007fe7840)({
       Names: ([]*ast.Ident) <nil>,
       Type: (*ast.Ident)(0x14007fa3e20)(int),
      })
     },
     Closing: (token.Pos) 0
    })
   }),
  }),
  imports: (map[string]string) {
  },
  aliases: (map[string]string) <nil>
 }),
 (*codegen.funcInfo)(0x14008480798)({
  function: (*ast.FuncDecl)(0x14007fdde00)({
   Name: (*ast.Ident)(0x14007ffe020)(AddN),
   Type: (*ast.FuncType)(0x14007ffe220)({
    Func: (token.Pos) 4232467,
    Params: (*ast.FieldList)(0x14007fddd40)({
     Opening: (token.Pos) 4232476,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x14007fe7900)({
       Names: ([]*ast.Ident) (len=1 cap=1) {
        (*ast.Ident)(0x14007ffe040)(args)
       },
       Type: (*ast.Ellipsis)(0x14007f9b560)({
        Ellipsis: (token.Pos) 4232482,
        Elt: (*ast.Ident)(0x14007ffe060)(int)
       }),
      })
     },
     Closing: (token.Pos) 4232488
    }),
    Results: (*ast.FieldList)(0x14007fddd70)({
     Opening: (token.Pos) 0,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x14007fe7940)({
       Names: ([]*ast.Ident) <nil>,
       Type: (*ast.Ident)(0x14007ffe080)(int),
      })
     },
     Closing: (token.Pos) 0
    })
   }),
  }),
  imports: (map[string]string) {
  },
  aliases: (map[string]string) <nil>
 }),
 (*codegen.funcInfo)(0x14008480840)({
  function: (*ast.FuncDecl)(0x14007fddec0)({
   Name: (*ast.Ident)(0x14007ffe2a0)(GetCurrentTime),
   Type: (*ast.FuncType)(0x14007ffe340)({
    Func: (token.Pos) 4232674,
    Params: (*ast.FieldList)(0x14007fdde30)({
     Opening: (token.Pos) 4232693,
     List: ([]*ast.Field) <nil>,
     Closing: (token.Pos) 4232694
    }),
    Results: (*ast.FieldList)(0x14007fdde60)({
     Opening: (token.Pos) 0,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x14007fe7ac0)({
       Names: ([]*ast.Ident) <nil>,
       Type: (*ast.SelectorExpr)(0x14007f9b608)({
        X: (*ast.Ident)(0x14007ffe2c0)(time),
        Sel: (*ast.Ident)(0x14007ffe2e0)(Time)
       }),
      })
     },
     Closing: (token.Pos) 0
    })
   }),
  }),
  imports: (map[string]string) (len=1) {
   (string) (len=4) "time": (string) (len=4) "time"
  },
  aliases: (map[string]string) <nil>
 }),
 (*codegen.funcInfo)(0x140084808d0)({
  function: (*ast.FuncDecl)(0x14008478000)({
   Name: (*ast.Ident)(0x14007ffe360)(GetCurrentTimeFormatted),
   Type: (*ast.FuncType)(0x14007ffe440)({
    Func: (token.Pos) 4232780,
    Params: (*ast.FieldList)(0x14007fddef0)({
     Opening: (token.Pos) 4232808,
     List: ([]*ast.Field) <nil>,
     Closing: (token.Pos) 4232809
    }),
    Results: (*ast.FieldList)(0x14007fddf20)({
     Opening: (token.Pos) 0,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x14007fe7b40)({
       Names: ([]*ast.Ident) <nil>,
       Type: (*ast.Ident)(0x14007ffe380)(string),
      })
     },
     Closing: (token.Pos) 0
    })
   }),
  }),
  imports: (map[string]string) {
  },
  aliases: (map[string]string) <nil>
 }),
 (*codegen.funcInfo)(0x14008480948)({
  function: (*ast.FuncDecl)(0x14008478150)({
   Name: (*ast.Ident)(0x14007ffe460)(GetFullName),
   Type: (*ast.FuncType)(0x14007ffe580)({
    Func: (token.Pos) 4232939,
    Params: (*ast.FieldList)(0x14008478060)({
     Opening: (token.Pos) 4232955,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x14007fe7c00)({
       Names: ([]*ast.Ident) (len=2 cap=2) {
        (*ast.Ident)(0x14007ffe480)(firstName),
        (*ast.Ident)(0x14007ffe4a0)(lastName)
       },
       Type: (*ast.Ident)(0x14007ffe4c0)(string),
      })
     },
     Closing: (token.Pos) 4232982
    }),
    Results: (*ast.FieldList)(0x14008478090)({
     Opening: (token.Pos) 0,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x14007fe7c40)({
       Names: ([]*ast.Ident) <nil>,
       Type: (*ast.Ident)(0x14007ffe4e0)(string),
      })
     },
     Closing: (token.Pos) 0
    })
   }),
  }),
  imports: (map[string]string) {
  },
  aliases: (map[string]string) <nil>
 }),
 (*codegen.funcInfo)(0x140084809d8)({
  function: (*ast.FuncDecl)(0x14008478300)({
   Name: (*ast.Ident)(0x14007ffe5a0)(SayHello),
   Type: (*ast.FuncType)(0x14007ffe720)({
    Func: (token.Pos) 4233132,
    Params: (*ast.FieldList)(0x14008478180)({
     Opening: (token.Pos) 4233145,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x14007fe7c80)({
       Names: ([]*ast.Ident) (len=1 cap=1) {
        (*ast.Ident)(0x14007ffe5c0)(name)
       },
       Type: (*ast.StarExpr)(0x14007f9b710)({
        Star: (token.Pos) 4233151,
        X: (*ast.Ident)(0x14007ffe5e0)(string)
       }),
      })
     },
     Closing: (token.Pos) 4233158
    }),
    Results: (*ast.FieldList)(0x140084781b0)({
     Opening: (token.Pos) 0,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x14007fe7cc0)({
       Names: ([]*ast.Ident) <nil>,
       Type: (*ast.Ident)(0x14007ffe600)(string),
      })
     },
     Closing: (token.Pos) 0
    })
   }),
  }),
  imports: (map[string]string) {
  },
  aliases: (map[string]string) <nil>
 }),
 (*codegen.funcInfo)(0x14008480a98)({
  function: (*ast.FuncDecl)(0x14008478480)({
   Name: (*ast.Ident)(0x14007ffe8a0)(GetPerson),
   Type: (*ast.FuncType)(0x14007ffea00)({
    Func: (token.Pos) 4233521,
    Params: (*ast.FieldList)(0x14008478360)({
     Opening: (token.Pos) 4233535,
     List: ([]*ast.Field) <nil>,
     Closing: (token.Pos) 4233536
    }),
    Results: (*ast.FieldList)(0x14008478390)({
     Opening: (token.Pos) 0,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x14007fe7e80)({
       Names: ([]*ast.Ident) <nil>,
       Type: (*ast.Ident)(0x14007ffe8c0)(Person),
      })
     },
     Closing: (token.Pos) 0
    })
   }),
  }),
  imports: (map[string]string) {
  },
  aliases: (map[string]string) <nil>
 }),
 (*codegen.funcInfo)(0x14008480b10)({
  function: (*ast.FuncDecl)(0x14008478570)({
   Name: (*ast.Ident)(0x14007ffea20)(GetRandomPerson),
   Type: (*ast.FuncType)(0x14007ffebc0)({
    Func: (token.Pos) 4233681,
    Params: (*ast.FieldList)(0x140084784b0)({
     Opening: (token.Pos) 4233701,
     List: ([]*ast.Field) <nil>,
     Closing: (token.Pos) 4233702
    }),
    Results: (*ast.FieldList)(0x140084784e0)({
     Opening: (token.Pos) 0,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x14007fe7f40)({
       Names: ([]*ast.Ident) <nil>,
       Type: (*ast.Ident)(0x14007ffea40)(Person),
      })
     },
     Closing: (token.Pos) 0
    })
   }),
  }),
  imports: (map[string]string) {
  },
  aliases: (map[string]string) <nil>
 }),
 (*codegen.funcInfo)(0x14008480ba0)({
  function: (*ast.FuncDecl)(0x14008478840)({
   Name: (*ast.Ident)(0x14007ffebe0)(GetPeople),
   Type: (*ast.FuncType)(0x14007ffef20)({
    Func: (token.Pos) 4233812,
    Params: (*ast.FieldList)(0x140084785a0)({
     Opening: (token.Pos) 4233826,
     List: ([]*ast.Field) <nil>,
     Closing: (token.Pos) 4233827
    }),
    Results: (*ast.FieldList)(0x14008478600)({
     Opening: (token.Pos) 0,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x1400847a180)({
       Names: ([]*ast.Ident) <nil>,
       Type: (*ast.ArrayType)(0x140084785d0)({
        Lbrack: (token.Pos) 4233829,
        Len: (ast.Expr) <nil>,
        Elt: (*ast.Ident)(0x14007ffec00)(Person)
       }),
      })
     },
     Closing: (token.Pos) 0
    })
   }),
  }),
  imports: (map[string]string) {
  },
  aliases: (map[string]string) <nil>
 }),
 (*codegen.funcInfo)(0x14008480c30)({
  function: (*ast.FuncDecl)(0x14008478930)({
   Name: (*ast.Ident)(0x14007ffef40)(GetNameAndAge),
   Type: (*ast.FuncType)(0x14007fff180)({
    Func: (token.Pos) 4234120,
    Params: (*ast.FieldList)(0x14008478870)({
     Opening: (token.Pos) 4234138,
     List: ([]*ast.Field) <nil>,
     Closing: (token.Pos) 4234139
    }),
    Results: (*ast.FieldList)(0x140084788d0)({
     Opening: (token.Pos) 4234141,
     List: ([]*ast.Field) (len=2 cap=2) {
      (*ast.Field)(0x1400847a3c0)({
       Names: ([]*ast.Ident) (len=1 cap=1) {
        (*ast.Ident)(0x14007ffef60)(name)
       },
       Type: (*ast.Ident)(0x14007ffef80)(string),
      }),
      (*ast.Field)(0x1400847a400)({
       Names: ([]*ast.Ident) (len=1 cap=1) {
        (*ast.Ident)(0x14007ffefa0)(age)
       },
       Type: (*ast.Ident)(0x14007ffefc0)(int),
      })
     },
     Closing: (token.Pos) 4234162
    })
   }),
  }),
  imports: (map[string]string) {
  },
  aliases: (map[string]string) <nil>
 }),
 (*codegen.funcInfo)(0x14008480cc0)({
  function: (*ast.FuncDecl)(0x14008478ab0)({
   Name: (*ast.Ident)(0x14007fff1a0)(TestNormalError),
   Type: (*ast.FuncType)(0x14007fff460)({
    Func: (token.Pos) 4234268,
    Params: (*ast.FieldList)(0x14008478960)({
     Opening: (token.Pos) 4234288,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x1400847a500)({
       Names: ([]*ast.Ident) (len=1 cap=1) {
        (*ast.Ident)(0x14007fff1c0)(input)
       },
       Type: (*ast.Ident)(0x14007fff1e0)(string),
      })
     },
     Closing: (token.Pos) 4234301
    }),
    Results: (*ast.FieldList)(0x140084789c0)({
     Opening: (token.Pos) 4234303,
     List: ([]*ast.Field) (len=2 cap=2) {
      (*ast.Field)(0x1400847a540)({
       Names: ([]*ast.Ident) <nil>,
       Type: (*ast.Ident)(0x14007fff200)(string),
      }),
      (*ast.Field)(0x1400847a580)({
       Names: ([]*ast.Ident) <nil>,
       Type: (*ast.Ident)(0x14007fff220)(error),
      })
     },
     Closing: (token.Pos) 4234317
    })
   }),
  }),
  imports: (map[string]string) {
  },
  aliases: (map[string]string) <nil>
 }),
 (*codegen.funcInfo)(0x14008480d50)({
  function: (*ast.FuncDecl)(0x14008478c00)({
   Name: (*ast.Ident)(0x14007fff480)(TestAlternativeError),
   Type: (*ast.FuncType)(0x14007fff6c0)({
    Func: (token.Pos) 4234821,
    Params: (*ast.FieldList)(0x14008478ae0)({
     Opening: (token.Pos) 4234846,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x1400847a700)({
       Names: ([]*ast.Ident) (len=1 cap=1) {
        (*ast.Ident)(0x14007fff4a0)(input)
       },
       Type: (*ast.Ident)(0x14007fff4c0)(string),
      })
     },
     Closing: (token.Pos) 4234859
    }),
    Results: (*ast.FieldList)(0x14008478b10)({
     Opening: (token.Pos) 0,
     List: ([]*ast.Field) (len=1 cap=1) {
      (*ast.Field)(0x1400847a740)({
       Names: ([]*ast.Ident) <nil>,
       Type: (*ast.Ident)(0x14007fff4e0)(string),
      })
     },
     Closing: (token.Pos) 0
    })
   }),
  }),
  imports: (map[string]string) {
  },
  aliases: (map[string]string) <nil>
 }),
 (*codegen.funcInfo)(0x14008480de0)({
  function: (*ast.FuncDecl)(0x14008478c90)({
   Name: (*ast.Ident)(0x14007fff6e0)(TestPanic),
   Type: (*ast.FuncType)(0x14007fff740)({
    Func: (token.Pos) 4235146,
    Params: (*ast.FieldList)(0x14008478c30)({
     Opening: (token.Pos) 4235160,
     List: ([]*ast.Field) <nil>,
     Closing: (token.Pos) 4235161
    }),
   }),
  }),
  imports: (map[string]string) {
  },
  aliases: (map[string]string) <nil>
 }),
 (*codegen.funcInfo)(0x14008480e58)({
  function: (*ast.FuncDecl)(0x14008478d20)({
   Name: (*ast.Ident)(0x14007fff760)(TestExit),
   Type: (*ast.FuncType)(0x14007fff8c0)({
    Func: (token.Pos) 4235405,
    Params: (*ast.FieldList)(0x14008478cc0)({
     Opening: (token.Pos) 4235418,
     List: ([]*ast.Field) <nil>,
     Closing: (token.Pos) 4235419
    }),
   }),
  }),
  imports: (map[string]string) {
  },
  aliases: (map[string]string) <nil>
 }),
 (*codegen.funcInfo)(0x14008480ed0)({
  function: (*ast.FuncDecl)(0x14008478db0)({
   Name: (*ast.Ident)(0x14007fff8e0)(TestLogging),
   Type: (*ast.FuncType)(0x14007fffe40)({
    Func: (token.Pos) 4235857,
    Params: (*ast.FieldList)(0x14008478d50)({
     Opening: (token.Pos) 4235873,
     List: ([]*ast.Field) <nil>,
     Closing: (token.Pos) 4235874
    }),
   }),
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
