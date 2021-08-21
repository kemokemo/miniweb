package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

const (
	exitCodeOK int = iota
	exitCodeFailed
	exitCodeArgumentsInvalid
)

const usage = `
Usage:
miniweb [Options] {Directory to serve} 

Examples:
# run with a default port '8080' and default directory 'public'
miniweb

# run with a default port '8080'
miniweb ./my/content/directory

# run with specified port by -p flag
miniweb -p 9000 ./my/content/directory

Options:
`

var (
	versionOut bool
	servePath  string
	port       string
)

func init() {
	flag.Usage = printUsageFunc

	flag.BoolVar(&versionOut, "v", false, "print the version number")
	flag.BoolVar(&versionOut, "version", false, "print the version number")
	flag.StringVar(&port, "p", "8080", "port to serve files")
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		servePath = args[0]
	} else {
		servePath = "public"
	}
}

func main() {
	os.Exit(run())
}

func printUsageFunc() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n%s", os.Args[0], usage)
	flag.PrintDefaults()
}

func run() int {
	if versionOut {
		fmt.Printf("%s.%s\n", Version, Revision)
		return exitCodeOK
	}

	if _, err := os.Stat(servePath); os.IsNotExist(err) {
		fmt.Printf("serving path does not exist: %v\n\n", servePath)
		printUsageFunc()
		return exitCodeArgumentsInvalid
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.Recover())
	e.Static("/", servePath)

	go func() {
		err := e.Start(fmt.Sprintf(":%s", port))
		if err != nil {
			e.Logger.Infof("shutting down this server:%v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Error(err)
		return exitCodeFailed
	}

	return exitCodeOK
}
