package select_statement

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (string, error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
		case <- ping(a):
			return a, nil
		case <- ping(b):
			return b, nil
		case <- time.After(timeout):
			return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(s string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(s)
		close(ch)
	}()
	return ch
}
