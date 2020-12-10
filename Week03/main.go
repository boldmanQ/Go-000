package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"context"
	"golang.org/x/sync/errgroup"
)

func serverAppWithTimeout(ctx context.Context, signalShotDown <-chan os.Signal) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Println(resp, "Start1: new courrency ServerApp —— Timeout 5seconds then quit")
	})
	s := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	go func() error {
		timeout, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
		defer cancel()
		select {
			case <- ctx.Done():
				fmt.Println("Timeout trick, Groutine Will Stop")
				s.Shutdown(timeout)
				return errors.New("App1 Stop, by Timeout")
			case <- signalShotDown:
				s.Shutdown(context.Background())
				return errors.New("App1 Stop, by signal")
		}
	}()
	fmt.Println("App1 starting")
	return s.ListenAndServe()
}

func serverAppWithSigalQuit(ctx context.Context, signalShotDown <-chan os.Signal) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Println(resp, "Start2 .new courrency ServerAppWithSingleQuit —— Timeout 5seconds then quit")
	})
	s := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	go func() error {
		timeout, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
		defer cancel()
		select {
			case <- ctx.Done():
				fmt.Println("Timeout trick, Groutine Will Stop")
				s.Shutdown(timeout)
				return errors.New("App2 Stop, by timeout")
			case <- signalShotDown:
				s.Shutdown(context.Background())
				return errors.New("App2 Stop, by signal")
		}
	}()
	fmt.Println("App2 Starting")
	return s.ListenAndServe()
}

func signalShotDownEvent(signalShotdown chan os.Signal) error {
	signal.Notify(signalShotdown,os.Interrupt,syscall.SIGINT,syscall.SIGHUP,syscall.SIGTERM,syscall.SIGQUIT,syscall.SIGKILL)
	select {
		case <- signalShotdown:
			return errors.New("Receive ShotDown singal, Stop Server")
	}
}

func main() {
	signalShotDown := make(chan os.Signal)
	group, ctx := errgroup.WithContext(context.Background())

	group.Go(func() error {
		return serverAppWithTimeout(ctx, signalShotDown)
	})

	group.Go(func() error {
		return serverAppWithSigalQuit(ctx, signalShotDown)
	})

	group.Go(func() error {
		return signalShotDownEvent(signalShotDown)
	})

	if err := group.Wait(); err != nil {
		fmt.Println("Event: %+v", err)
	}

}