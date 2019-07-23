package main

import "github.com/gregoryv/stamp"

func init() {
    s := &stamp.Stamp{
	    Package: "main",
	    Revision: "a2fe0a2",
	    ChangelogVersion: "0.1.0",
    }
    stamp.Use(s)
}

