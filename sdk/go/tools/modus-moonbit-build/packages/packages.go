// This file is based on: https://cs.opensource.google/go/x/tools/+/refs/tags/v0.28.0:go/packages/packages.go
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package packages

// See doc.go for package documentation and implementation notes.

import (
	"encoding/json"
	"go/ast"
	"go/token"
	"go/types"
	"log"
	"os"
	"sync"
	"time"
)

// TODO: Remove debugging
var gmlDebugEnv bool

func gmlPrintf(fmtStr string, args ...any) {
	// gmlDebugEnv = true
	sync.OnceFunc(func() {
		log.SetFlags(0)
		if os.Getenv("GML_DEBUG") == "true" {
			gmlDebugEnv = true
		}
	})
	if gmlDebugEnv {
		log.Printf(fmtStr, args...)
	}
}

// A LoadMode controls the amount of detail to return when loading.
// The bits below can be combined to specify which fields should be
// filled in the result packages.
//
// The zero value is a special case, equivalent to combining
// the NeedName, NeedFiles, and NeedCompiledMoonBitFiles bits.
//
// ID and Errors (if present) will always be filled.
// [Load] may return more information than requested.
//
// The Mode flag is a union of several bits named NeedName,
// NeedFiles, and so on, each of which determines whether
// a given field of Package (Name, Files, etc) should be
// populated.
//
// For convenience, we provide named constants for the most
// common combinations of Need flags:
//
//	[LoadFiles]     lists of files in each package
//	[LoadImports]   ... plus imports
//	[LoadTypes]     ... plus type information
//	[LoadSyntax]    ... plus type-annotated syntax
//	[LoadAllSyntax] ... for all dependencies
//
// Unfortunately there are a number of open bugs related to
// interactions among the LoadMode bits:
//   - https://github.com/golang/go/issues/56633
//   - https://github.com/golang/go/issues/56677
//   - https://github.com/golang/go/issues/58726
//   - https://github.com/golang/go/issues/63517
type LoadMode int

const (
	// NeedName adds Name and PkgPath.
	NeedName LoadMode = 1 << iota

	// NeedFiles adds Dir, MoonBitFiles, OtherFiles, and IgnoredFiles
	NeedFiles

	// NeedCompiledMoonBitFiles adds CompiledMoonBitFiles.
	NeedCompiledMoonBitFiles

	// NeedImports adds Imports. If NeedDeps is not set, the Imports field will contain
	// "placeholder" Packages with only the ID set.
	NeedImports

	// NeedDeps adds the fields requested by the LoadMode in the packages in Imports.
	NeedDeps

	// NeedExportFile adds ExportFile.
	NeedExportFile

	// NeedTypes adds Types, Fset, and IllTyped.
	NeedTypes

	// NeedSyntax adds Syntax and Fset.
	NeedSyntax

	// NeedTypesInfo adds TypesInfo and Fset.
	NeedTypesInfo

	// NeedTypesSizes adds TypesSizes.
	NeedTypesSizes

	// NeedForTest adds ForTest.
	//
	// Tests must also be set on the context for this field to be populated.
	NeedForTest

	// NeedModule adds Module.
	NeedModule

	// NeedEmbedFiles adds EmbedFiles.
	NeedEmbedFiles

	// NeedEmbedPatterns adds EmbedPatterns.
	NeedEmbedPatterns

	// Be sure to update loadmode_string.go when adding new items!
)

const (
	// LoadFiles loads the name and file names for the initial packages.
	LoadFiles = NeedName | NeedFiles | NeedCompiledMoonBitFiles

	// LoadImports loads the name, file names, and import mapping for the initial packages.
	LoadImports = LoadFiles | NeedImports

	// LoadTypes loads exported type information for the initial packages.
	LoadTypes = LoadImports | NeedTypes | NeedTypesSizes

	// LoadSyntax loads typed syntax for the initial packages.
	LoadSyntax = LoadTypes | NeedSyntax | NeedTypesInfo

	// LoadAllSyntax loads typed syntax for the initial packages and all dependencies.
	LoadAllSyntax = LoadSyntax | NeedDeps

	// Deprecated: NeedExportsFile is a historical misspelling of NeedExportFile.
	NeedExportsFile = NeedExportFile
)

// A Config specifies details about how packages should be loaded.
// The zero value is a valid configuration.
//
// Calls to [Load] do not modify this struct.
type Config struct {
	// Mode controls the level of information returned for each package.
	Mode LoadMode

	// // Context specifies the context for the load operation.
	// // Cancelling the context may cause [Load] to abort and
	// // return an error.
	// Context context.Context

	// // Logf is the logger for the config.
	// // If the user provides a logger, debug logging is enabled.
	// // If the GOPACKAGESDEBUG environment variable is set to true,
	// // but the logger is nil, default to gmlPrintf.
	// Logf func(format string, args ...interface{})

	// RootAbsPath is the absolute path to the root directory of the initial user program.
	// This directory contains the ".mooncakes" subdirectory where all
	// the imports should be able to be found.
	RootAbsPath string

	// Dir is the directory in which to run the build system's query tool
	// that provides information about the packages.
	// If Dir is empty, the tool is run in the current directory.
	Dir string

	// PackageName is the name that this package is referred to, which is
	// found in the moon.mod.json file.
	PackageName string

	// PackageAlias is the alias that this package is referred to, which is
	// found in the moon.mod.json file.
	PackageAlias string

	// Env is the environment to use when invoking the build system's query tool.
	// If Env is nil, the current environment is used.
	// As in os/exec's Cmd, only the last value in the slice for
	// each environment key is used. To specify the setting of only
	// a few variables, append to the current environment, as in:
	//
	//	opt.Env = append(os.Environ(), "GOOS=plan9", "GOARCH=386")
	//
	Env []string

	// // BuildFlags is a list of command-line flags to be passed through to
	// // the build system's query tool.
	// BuildFlags []string

	// // Fset provides source position information for syntax trees and types.
	// // If Fset is nil, Load will use a new fileset, but preserve Fset's value.
	// Fset *token.FileSet

	// // ParseFile is called to read and parse each file
	// // when preparing a package's type-checked syntax tree.
	// // It must be safe to call ParseFile simultaneously from multiple goroutines.
	// // If ParseFile is nil, the loader will uses parser.ParseFile.
	// //
	// // ParseFile should parse the source from src and use filename only for
	// // recording position information.
	// //
	// // An application may supply a custom implementation of ParseFile
	// // to change the effective file contents or the behavior of the parser,
	// // or to modify the syntax tree. For example, selectively eliminating
	// // unwanted function bodies can significantly accelerate type checking.
	// ParseFile func(fset *token.FileSet, filename string, src []byte) (*ast.File, error)

	// // If Tests is set, the loader includes not just the packages
	// // matching a particular pattern but also any related test packages,
	// // including test-only variants of the package and the test executable.
	// //
	// // For example, when using the go command, loading "fmt" with Tests=true
	// // returns four packages, with IDs "fmt" (the standard package),
	// // "fmt [fmt.test]" (the package as compiled for the test),
	// // "fmt_test" (the test functions from source files in package fmt_test),
	// // and "fmt.test" (the test binary).
	// //
	// // In build systems with explicit names for tests,
	// // setting Tests may have no effect.
	// Tests bool

	// // Overlay is a mapping from absolute file paths to file contents.
	// //
	// // For each map entry, [Load] uses the alternative file
	// // contents provided by the overlay mapping instead of reading
	// // from the file system. This mechanism can be used to enable
	// // editor-integrated tools to correctly analyze the contents
	// // of modified but unsaved buffers, for example.
	// //
	// // The overlay mapping is passed to the build system's driver
	// // (see "The driver protocol") so that it too can report
	// // consistent package metadata about unsaved files. However,
	// // drivers may vary in their level of support for overlays.
	// Overlay map[string][]byte
}

// A Package describes a loaded MoonBit package.
//
// It also defines part of the JSON schema of [DriverResponse].
// See the package documentation for an overview.
type Package struct {
	// MoonPkgJSON is the parsed moon.pkg.json file.
	MoonPkgJSON MoonPkgJSON

	// ID is a unique identifier for a package,
	// in a syntax provided by the underlying build system.
	//
	// Because the syntax varies based on the build system,
	// clients should treat IDs as opaque and not attempt to
	// interpret them.
	ID string

	// Name is the package name as it appears in the package source code.
	Name string

	// PkgPath is the package path as used by the go/types package.
	// It is "" for the main (user) package, and otherwise e.g. "@neo4j".
	PkgPath string

	// // Dir is the directory associated with the package, if it exists.
	// //
	// // For packages listed by the go command, this is the directory containing
	// // the package files.
	// Dir string

	// // Errors contains any errors encountered querying the metadata
	// // of the package, or while parsing or type-checking its files.
	// Errors []Error

	// // TypeErrors contains the subset of errors produced during type checking.
	// TypeErrors []types.Error

	// MoonBitFiles lists the absolute file paths of the package's MoonBit source files.
	// It may include files that should not be compiled, for example because
	// they contain non-matching build tags, are documentary pseudo-files such as
	// unsafe/unsafe.go or builtin/builtin.go, or are subject to cgo preprocessing.
	MoonBitFiles []string

	// Since this package is a hack to parse MoonBit files without the assistance
	// of the MoonBit compiler, and since we are shoe-horning the MoonBit AST
	// into an AST designed for the Go programming language, this map provides
	// assistance in resolving publicly-exported structs so that they can be
	// uniquely-identified by package.
	StructLookup map[string]*ast.TypeSpec

	// This is an ugly workaround, but due to the hack manner of processing MoonBit
	// files in potentially random orders combined with forward references to
	// custom types that are not resolved until all parsing has been completed,
	// this map contains a list of underlying types that may need to be added to
	// the metadata so that the Modus Runtime can prepare handlers for them.
	// It is not known if they are needed until all the source code processing has
	// been completed.
	PossiblyMissingUnderlyingTypes map[string]struct{}

	// // CompiledMoonBitFiles lists the absolute file paths of the package's source
	// // files that are suitable for type checking.
	// // This may differ from MoonBitFiles if files are processed before compilation.
	// CompiledMoonBitFiles []string

	// // OtherFiles lists the absolute file paths of the package's non-MoonBit source files,
	// // including assembly, C, C++, Fortran, Objective-C, SWIG, and so on.
	// OtherFiles []string

	// // EmbedFiles lists the absolute file paths of the package's files
	// // embedded with go:embed.
	// EmbedFiles []string

	// // EmbedPatterns lists the absolute file patterns of the package's
	// // files embedded with go:embed.
	// EmbedPatterns []string

	// // IgnoredFiles lists source files that are not part of the package
	// // using the current build configuration but that might be part of
	// // the package using other build configurations.
	// IgnoredFiles []string

	// // ExportFile is the absolute path to a file containing type
	// // information for the package as provided by the build system.
	// ExportFile string

	// Imports maps import paths appearing in the package's MoonBit source files
	// to corresponding loaded Packages.
	Imports map[string]*Package

	// // Module is the module information for the package if it exists.
	// //
	// // Note: it may be missing for std and cmd; see MoonBit issue #65816.
	// Module *Module

	// // -- The following fields are not part of the driver JSON schema. --

	// // Types provides type information for the package.
	// // The NeedTypes LoadMode bit sets this field for packages matching the
	// // patterns; type information for dependencies may be missing or incomplete,
	// // unless NeedDeps and NeedImports are also set.
	// //
	// // Each call to [Load] returns a consistent set of type
	// // symbols, as defined by the comment at [types.Identical].
	// // Avoid mixing type information from two or more calls to [Load].
	// Types *types.Package `json:"-"`

	// Fset provides position information for Types, TypesInfo, and Syntax.
	// It is set only when Types is set.
	Fset *token.FileSet `json:"-"`

	// // IllTyped indicates whether the package or any dependency contains errors.
	// // It is set only when Types is set.
	// IllTyped bool `json:"-"`

	// Syntax is the package's syntax trees, for the files listed in CompiledMoonBitFiles.
	//
	// The NeedSyntax LoadMode bit populates this field for packages matching the patterns.
	// If NeedDeps and NeedImports are also set, this field will also be populated
	// for dependencies.
	//
	// Syntax is kept in the same order as CompiledMoonBitFiles, with the caveat that nils are
	// removed.  If parsing returned nil, Syntax may be shorter than CompiledMoonBitFiles.
	Syntax []*ast.File `json:"-"`

	// TypesInfo provides type information about the package's syntax trees.
	// It is set only when Syntax is set.
	TypesInfo *types.Info `json:"-"`

	// // TypesSizes provides the effective size function for types in TypesInfo.
	// TypesSizes types.Sizes `json:"-"`

	// // -- internal --

	// // ForTest is the package under test, if any.
	// ForTest string
}

func (p *Package) AddPossiblyMissingUnderlyingType(typ string) {
	if p.PossiblyMissingUnderlyingTypes == nil {
		p.PossiblyMissingUnderlyingTypes = map[string]struct{}{}
	}
	p.PossiblyMissingUnderlyingTypes[typ] = struct{}{}
}

// Module provides module information for a package.
//
// It also defines part of the JSON schema of [DriverResponse].
// See the package documentation for an overview.
type Module struct {
	Path           string       // module path
	Version        string       // module version
	Replace        *Module      // replaced by this module
	Time           *time.Time   // time version was created
	Main           bool         // is this the main module?
	Indirect       bool         // is this module only an indirect dependency of main module?
	Dir            string       // directory holding files for this module, if any
	MoonBitMod     string       // path to go.mod file used when loading this module, if any
	MoonBitVersion string       // go version used in module
	Error          *ModuleError // error loading module
}

// ModuleError holds errors loading a module.
type ModuleError struct {
	Err string // the error itself
}

// func init() {
// packagesinternal.GetDepsErrors = func(p interface{}) []*packagesinternal.PackageError {
// 	return p.(*Package).depsErrors
// }
// packagesinternal.SetModFile = func(config interface{}, value string) {
// 	config.(*Config).modFile = value
// }
// packagesinternal.SetModFlag = func(config interface{}, value string) {
// 	config.(*Config).modFlag = value
// }
// packagesinternal.TypecheckCgo = int(typecheckCgo)
// packagesinternal.DepsErrors = int(needInternalDepsErrors)
// }

// An Error describes a problem with a package's metadata, syntax, or types.
type Error struct {
	Pos  string // "file:line:col" or "file:line" or "" or "-"
	Msg  string
	Kind ErrorKind
}

// ErrorKind describes the source of the error, allowing the user to
// differentiate between errors generated by the driver, the parser, or the
// type-checker.
type ErrorKind int

const (
	UnknownError ErrorKind = iota
	ListError
	ParseError
	TypeError
)

func (err Error) Error() string {
	pos := err.Pos
	if pos == "" {
		pos = "-" // like token.Position{}.String()
	}
	return pos + ": " + err.Msg
}

// flatPackage is the JSON form of Package
// It drops all the type and syntax fields, and transforms the Imports
//
// TODO(adonovan): identify this struct with Package, effectively
// publishing the JSON protocol.
type flatPackage struct {
	ID                   string
	Name                 string            `json:",omitempty"`
	PkgPath              string            `json:",omitempty"`
	Errors               []Error           `json:",omitempty"`
	MoonBitFiles         []string          `json:",omitempty"`
	CompiledMoonBitFiles []string          `json:",omitempty"`
	OtherFiles           []string          `json:",omitempty"`
	EmbedFiles           []string          `json:",omitempty"`
	EmbedPatterns        []string          `json:",omitempty"`
	IgnoredFiles         []string          `json:",omitempty"`
	ExportFile           string            `json:",omitempty"`
	Imports              map[string]string `json:",omitempty"`
}

// MarshalJSON returns the Package in its JSON form.
// For the most part, the structure fields are written out unmodified, and
// the type and syntax fields are skipped.
// The imports are written out as just a map of path to package id.
// The errors are written using a custom type that tries to preserve the
// structure of error types we know about.
//
// This method exists to enable support for additional build systems.  It is
// not intended for use by clients of the API and we may change the format.
func (p *Package) MarshalJSON() ([]byte, error) {
	flat := &flatPackage{
		ID:   p.ID,
		Name: p.Name,
		// PkgPath:              p.PkgPath,
		// Errors:               p.Errors,
		MoonBitFiles: p.MoonBitFiles,
		// CompiledMoonBitFiles: p.CompiledMoonBitFiles,
		// OtherFiles:           p.OtherFiles,
		// EmbedFiles:           p.EmbedFiles,
		// EmbedPatterns:        p.EmbedPatterns,
		// IgnoredFiles:         p.IgnoredFiles,
		// ExportFile:           p.ExportFile,
	}
	if len(p.Imports) > 0 {
		flat.Imports = make(map[string]string, len(p.Imports))
		for path, ipkg := range p.Imports {
			flat.Imports[path] = ipkg.ID
		}
	}
	return json.Marshal(flat)
}

// UnmarshalJSON reads in a Package from its JSON format.
// See MarshalJSON for details about the format accepted.
func (p *Package) UnmarshalJSON(b []byte) error {
	flat := &flatPackage{}
	if err := json.Unmarshal(b, &flat); err != nil {
		return err
	}
	*p = Package{
		ID:   flat.ID,
		Name: flat.Name,
		// PkgPath:              flat.PkgPath,
		// Errors:               flat.Errors,
		MoonBitFiles: flat.MoonBitFiles,
		// CompiledMoonBitFiles: flat.CompiledMoonBitFiles,
		// OtherFiles:           flat.OtherFiles,
		// EmbedFiles:           flat.EmbedFiles,
		// EmbedPatterns:        flat.EmbedPatterns,
		// IgnoredFiles:         flat.IgnoredFiles,
		// ExportFile:           flat.ExportFile,
	}
	if len(flat.Imports) > 0 {
		p.Imports = make(map[string]*Package, len(flat.Imports))
		for path, id := range flat.Imports {
			p.Imports[path] = &Package{ID: id, StructLookup: map[string]*ast.TypeSpec{}}
		}
	}
	return nil
}

func (p *Package) String() string { return p.ID }
