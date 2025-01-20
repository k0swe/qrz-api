package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/k0swe/qrz-api"
	"os"
)

func main() {
	ctx := context.Background()
	user := flag.String("username", "", "QRZ.com login name")
	pw := flag.String("password", "", "QRZ.com password")
	call := flag.String("lookup", "", "The callsign to look up")
	flag.Parse()
	if *user == "" || *pw == "" || *call == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	lookupResp, err := qrz.Lookup(ctx, user, pw, call)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if *lookupResp.Callsign == (qrz.Callsign{}) {
		fmt.Println(*call, "was not found")
	} else {
		fmt.Println(*call, "is", *lookupResp.Callsign.Fname, *lookupResp.Callsign.Name)
	}
}
