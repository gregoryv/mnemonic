package main

import "github.com/gregoryv/stamp"

func init() {
    s := &stamp.Stamp{
	    Package: "main",
	    Revision: "cab9528",
	    ChangelogVersion: "0.1.1",
    }
    stamp.Use(s)
}

