// goncurses - ncurses library for Go.
// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package goncurses is a new curses (ncurses) library for the Go programming
// language. It implements all the ncurses extension libraries: form, menu and
// panel.
//
// Minimal operation would consist of initializing the display:
//
// 	src, err := goncurses.Init()
// 	if err != nil {
// 		log.Fatal("init:", err)
// 	}
// 	defer goncurses.End()
//
// It is important to always call End() before your program exits. If you
// fail to do so, the terminal will not perform properly and will either
// need to be reset or restarted completely.
//
// CAUTION: Calls to ncurses functions are normally not atomic nor reentrant
// and therefore extreme care should be taken to ensure ncurses functions
// are not called concurrently. Specifically, never write data to the same
// window concurrently nor accept input and send output to the same window as
// both alter the underlying C data structures in a non safe manner.
//
// Ideally, you should structure your program to ensure all ncurses related
// calls happen in a single goroutine. This is probably most easily achieved
// via channels and Go's built-in select. Alternatively, or additionally, you
// can use a mutex to protect any calls in multiple goroutines from happening
// concurrently. Failure to do so will result in unpredictable and
// undefined behaviour in your program.
//
// The examples directory contains demontrations of many of the capabilities
// goncurses can provide.
package goncurses
