package main

import "github.com/gregoryv/stamp"

func init() {
    s := &stamp.Stamp{
	    Package: "main",
	    Revision: "1801ee6",
	    ChangelogVersion: "Unreleased",
    }
    stamp.Use(s)
}

