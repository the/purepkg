package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/tools/go/packages"
)

const (
	successExitCode          = 0
	usageExitCode            = 1
	unallowedPackageExitCode = 2
	packageErrorExitCode     = 3
)

func main() {
	var (
		verboseFlag bool
		allowStdlib bool
		ignoreTests bool
	)
	flag.BoolVar(&verboseFlag, "v", false, "verbose output")
	flag.BoolVar(&allowStdlib, "stdlib", false, "include all standard library packages")
	flag.BoolVar(&ignoreTests, "notest", false, "ignore tests")

	var allowedPkgs allowList
	flag.Var(&allowedPkgs, "allow", "comma separated list of allowed `packages`")

	flag.Parse()
	flag.Usage = usage
	if flag.NArg() == 0 {
		flag.Usage()
	}

	cfg := &packages.Config{
		Mode:  packages.NeedName | packages.NeedImports,
		Tests: !ignoreTests,
	}
	pkgs, err := packages.Load(cfg, flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	// if all the checks go well
	exitCode := successExitCode

	for _, pkg := range pkgs {
		if len(pkg.Errors) > 0 {
			packages.PrintErrors(pkgs)
			os.Exit(packageErrorExitCode)
		}

		printPackage := true
		if verboseFlag {
			fmt.Printf("package %s\n", pkg.PkgPath)
			printPackage = false
		}

		for imported := range pkg.Imports {
			allowed := isStdlib(imported) && allowStdlib
			if !allowed {
				allowed = allowedPkgs.contains(imported)
			}

			if !allowed || verboseFlag {
				if printPackage {
					fmt.Printf("package %s\n", pkg.PkgPath)
					printPackage = false
				}
				fmt.Printf("  %s import %s\n", allowedStr(allowed), imported)
			}

			if !allowed {
				exitCode = unallowedPackageExitCode
			}
		}
	}

	os.Exit(exitCode)
}

func allowedStr(allowed bool) string {
	if allowed {
		return "✔"
	}
	return "✗"
}

func isStdlib(name string) bool {
	path := strings.Split(name, "/")
	return !strings.Contains(path[0], ".")
}

func usage() {
	fmt.Fprintf(os.Stderr, "purepkg verifies that only allowed Go packages are imported.\n\n")
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "  purepkg [-v] [-stdlib] [-notest] [-allow pkg1,pkg2,...] <package>\n\n")
	fmt.Fprintf(os.Stderr, "Options:\n")
	flag.PrintDefaults()
	os.Exit(usageExitCode)
}
