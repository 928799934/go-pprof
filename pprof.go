package pprof

import (
	"context"
	"net"
	"net/http"
	"net/http/pprof"
	"strings"
	"time"
)

var (
	ss []*http.Server
)

// InitByString StartPprof start http pprof.
func InitByString(addrList []string) error {
	for _, addr := range addrList {
		l, err := getListener(addr)
		if err != nil {
			return err
		}
		s := &http.Server{
			Handler:        initServeMux(),
			ReadTimeout:    15 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		ss = append(ss, s)
		go func(s *http.Server, l net.Listener) {
			if err := s.Serve(l); err != nil {
				if err == http.ErrServerClosed {
					return
				}
				logf("s.ListenAndServe() error(%v)", err)
			}
		}(s, l)
	}
	return nil
}

// InitByListener ...
func InitByListener(socketList []net.Listener) error {
	for _, socket := range socketList {
		s := &http.Server{
			Handler:        initServeMux(),
			ReadTimeout:    15 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		ss = append(ss, s)
		go func(s *http.Server, l net.Listener) {
			if err := s.Serve(l); err != nil {
				if err == http.ErrServerClosed {
					return
				}
				logf("s.ListenAndServe() error(%v)", err)
			}
		}(s, socket)
	}
	return nil
}

// initServeMux ...
func initServeMux() http.Handler {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/debug/pprof/", pprof.Index)
	serveMux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	serveMux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	serveMux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	serveMux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	return serveMux
}

func getListener(addr string) (net.Listener, error) {
	var (
		l   net.Listener
		err error
	)
	if strings.Index(addr, "/") != -1 {
		l, err = net.Listen("unix", addr)
	} else {
		l, err = net.Listen("tcp", addr)
	}
	if err != nil {
		logf("net.Listen(%v) error(%v)", addr, err)
		return nil, err
	}
	return l, nil
}

// Close close the resource.
func Close() {
	for _, s := range ss {
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		if err := s.Shutdown(ctx); err != nil {
			logf("s.Shutdown(ctx) error(%v)", err)
		}
	}
	ss = []*http.Server{}
}
