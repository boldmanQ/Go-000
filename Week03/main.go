package main

import (
	"fmt"
	"net/http"
	"time"

	//"log"
	//"net/http"
	//"os/signal"
	"context"
	"golang.org/x/sync/errgroup"
)

func serverAppWithTimeout(ctx context.Context) error {
	mux := http.NewServeMux()
	mux.Handle("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Println(resp, "Start1: new courrency ServerApp —— Timeout 5seconds then quit")
	})
	s := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	go func() {
		timeout, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
		defer cancel()
		select {
			case <- ctx.Done():
				fmt.Println("Timeout trick, Groutine Will Stop")
				s.Shutdown(timeout)
		}
	}()
	fmt.Println("App1 starting")
	return s.ListenAndServe()
}

func serverAppWithSigalQuit(ctx context.Context) error {
	mux := http.NewServeMux()
	mux.Handle("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Println(resp, "Start2 .new courrency ServerAppWithSingleQuit —— Timeout 5seconds then quit")
	})
	s := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	go func() {
		timeout, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
		defer cancel()
		select {
			case <- ctx.Done():
				fmt.Println("Timeout trick, Groutine Will Stop")
				s.Shutdown(timeout)
			case <- signalEvent:
				return errors.New("Process Quit with Signal")
		}
	}()
	fmt.Println("App2 Starting")
}


func main() {
	group, ctx := errgroup.WithContext(context.Background())

	group.Go(func() error {
		return serverApp(ctx)
	})

	group.Go(func() error {
		return serverDebug(ctx)
	})
	fmt.Println(group, ctx)

	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Println(w, "hello, go cocurrence")
	//})
	//go func() {
	//	if err:= http.ListenAndServe(":8080", nil); err != nil {
	//		log.Fatal(err)
	//	}
	//}()
	//
	//select {}
}