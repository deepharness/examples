package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"sync"

	calc "goa.design/examples/basic"
	calcsvc "goa.design/examples/basic/gen/calc"
)

func main() {
	// Define command line flags, add any other flag required to configure the
	// service.
	var (
		hostF     = flag.String("host", "development", "Server host (valid values: development, production)")
		domainF   = flag.String("domain", "", "Host domain name (overrides host domain specified in service design)")
		httpPortF = flag.String("http-port", "", "HTTP port (overrides host HTTP port specified in service design)")
		grpcPortF = flag.String("grpc-port", "", "gRPC port (overrides host gRPC port specified in service design)")
		versionF  = flag.String("version", "v1", "API version")
		secureF   = flag.Bool("secure", false, "Use secure scheme (https or grpcs)")
		dbgF      = flag.Bool("debug", false, "Log request and response bodies")
		daemonF   = flag.String("daemon", "127.0.0.1:2000", "X-Ray daemon address")
	)
	flag.Parse()

	// Setup logger. Replace logger with your own log package of choice.
	var (
		logger *log.Logger
	)
	{
		logger = log.New(os.Stderr, "[calc] ", log.Ltime)
	}

	// Initialize the services.
	var (
		calcSvc calcsvc.Service
	)
	{
		calcSvc = calc.NewCalc()
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		calcEndpoints *calcsvc.Endpoints
	)
	{
		calcEndpoints = calcsvc.NewEndpoints(calcSvc)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	// Start the servers and send errors (if any) to the error channel.
	switch *hostF {
	case "development":
		{
			addr := "http://localhost:8000/calc"
			u, err := url.Parse(addr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "invalid URL %#v: %s", addr, err)
				os.Exit(1)
			}
			if *secureF {
				u.Scheme = "https"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *httpPortF != "" {
				h := strings.Split(u.Host, ":")[0]
				u.Host = h + ":" + *httpPortF
			} else if u.Port() == "" {
				u.Host += ":80"
			}
			handleHTTPServer(ctx, u, calcEndpoints, &wg, errc, logger, *dbgF, *daemonF)
		}

		{
			addr := "grpc://localhost:8080"
			u, err := url.Parse(addr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "invalid URL %#v: %s", addr, err)
				os.Exit(1)
			}
			if *secureF {
				u.Scheme = "grpcs"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *grpcPortF != "" {
				h := strings.Split(u.Host, ":")[0]
				u.Host = h + ":" + *grpcPortF
			} else if u.Port() == "" {
				u.Host += ":8080"
			}
			handleGRPCServer(ctx, u, calcEndpoints, &wg, errc, logger, *dbgF, *daemonF)
		}

	case "production":
		{
			addr := "https://{version}.goa.design/calc"
			addr = strings.Replace(addr, "{version}", *versionF, -1)
			u, err := url.Parse(addr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "invalid URL %#v: %s", addr, err)
				os.Exit(1)
			}
			if *secureF {
				u.Scheme = "https"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *httpPortF != "" {
				h := strings.Split(u.Host, ":")[0]
				u.Host = h + ":" + *httpPortF
			} else if u.Port() == "" {
				u.Host += ":443"
			}
			handleHTTPServer(ctx, u, calcEndpoints, &wg, errc, logger, *dbgF, *daemonF)
		}

		{
			addr := "grpcs://{version}.goa.design"
			addr = strings.Replace(addr, "{version}", *versionF, -1)
			u, err := url.Parse(addr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "invalid URL %#v: %s", addr, err)
				os.Exit(1)
			}
			if *secureF {
				u.Scheme = "grpcs"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *grpcPortF != "" {
				h := strings.Split(u.Host, ":")[0]
				u.Host = h + ":" + *grpcPortF
			} else if u.Port() == "" {
				u.Host += ":8443"
			}
			handleGRPCServer(ctx, u, calcEndpoints, &wg, errc, logger, *dbgF, *daemonF)
		}

	default:
		fmt.Fprintf(os.Stderr, "invalid host argument: %q (valid hosts: development|production)", *hostF)
	}
	// Wait for signal.
	logger.Printf("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Println("exited")
}
