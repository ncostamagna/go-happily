package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	stop := make(chan struct{})

	for i := 1; i <= 3; i++ {
		go ctxWorker(ctx, i)
		go chanWorker(stop, i)
	}

	time.Sleep(3 * time.Second)
	fmt.Println(runtime.NumGoroutine())
	fmt.Println("Stopping all workers")
	cancel() // Cancela todas las goroutines
	close(stop) // Cierra el canal, notificando a las goroutines

	time.Sleep(time.Second) // Esperar a que terminen los logs
}

// Flexible, soporta timeouts - Ligero overhead
func ctxWorker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker ctx %d stopped\n", id)
			return
		default:
			fmt.Printf("Worker ctx %d running\n", id)
			time.Sleep(time.Second)
		}
	}
}

// Simple y eficiente - Menos control que context
func chanWorker(stop chan struct{}, id int) {
	for {
		select {
		case <-stop:
			fmt.Printf("Worker chan %d stopped\n", id)
			return
		default:
			fmt.Printf("Worker chan %d running\n", id)
			time.Sleep(time.Second)
		}
	}
}