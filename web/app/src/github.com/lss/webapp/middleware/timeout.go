package middleware

import (
	"context"
	"net/http"
	"time"
)

type TimeoutMiddleware struct {
	Next http.Handler
}

func (tm TimeoutMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//Checking if this is the last middleware on the chain
	//If yes, call the http.DefaultServerMux
	if tm.Next == nil {
		tm.Next = http.DefaultServeMux
	}

	//Accessing the current context
	ctx := r.Context()
	//Create a new context adding a timeout on its Done() channel adding 3 seconds 
	//after processing the request
	ctx, _ = context.WithTimeout(ctx, 3*time.Second)
	//replacing the request context with the timeout context(ctx)
	r.WithContext(ctx)
	//creating a new channel which receives a signal(empty struct) if the request process normally
	ch := make(chan struct{})
	//
	go func() {
		//passing http.ResponseWriter and the modified http.Request object with the new context
		tm.Next.ServeHTTP(w, r)
		//If it returns just send a message to the channel
		ch <- struct{}{}
	}()
	//The channel that gets the signal gets processed
	select {
	//If the request ends normally, just let it continue normally.
	case <-ch:
		return
	//If the request times out, send a http.StatusRequestTimeout
	case <-ctx.Done():
		w.WriteHeader(http.StatusRequestTimeout)
	}
}
