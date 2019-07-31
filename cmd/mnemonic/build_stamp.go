package main

import "github.com/gregoryv/stamp"

func init() {
    s := &stamp.Stamp{
	    Package: "main",
	    Revision: "6c20c0e",
	    ChangelogVersion: "Unreleased",
    }
    stamp.Use(s)
}

