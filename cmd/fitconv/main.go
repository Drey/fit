// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/muktihari/fit/cmd/fitconv/fitcsv"
	"github.com/muktihari/fit/decoder"
)

var version = "dev"

const blockSize = 8 << 10

func main() {
	var flagVersion bool
	flag.BoolVar(&flagVersion, "v", false, "Show version")

	// Commands
	var flagFitToCsv bool
	flag.BoolVar(&flagFitToCsv, "f2c", false, "Convert FIT to CSV (default if not specified)")

	// FitToCsv Options
	var flagFitToCsvRawValue bool
	flag.BoolVar(&flagFitToCsvRawValue, "f2c-raw-value", false, "Use raw value instead of scaled value")

	var flagFitToCsvUseDisk bool
	flag.BoolVar(&flagFitToCsvUseDisk, "f2c-use-disk", false, "Use disk instead of load everything in memory")

	var flagFitToCsvUnknownNumber bool
	flag.BoolVar(&flagFitToCsvUnknownNumber, "f2c-unknown-number", false, "Print 'unknown(68)' instead of 'unknown'")

	flag.Parse()

	if flagVersion {
		fmt.Println(version)
		return
	}

	if !flagFitToCsv {
		flagFitToCsv = true // default
	}

	var fitToCsvOptions []fitcsv.Option
	if flagFitToCsvRawValue {
		fitToCsvOptions = append(fitToCsvOptions, fitcsv.WithRawValue())
	}
	if flagFitToCsvUnknownNumber {
		fitToCsvOptions = append(fitToCsvOptions, fitcsv.WithUnknownNumber())
	}
	if flagFitToCsvUseDisk {
		fitToCsvOptions = append(fitToCsvOptions, fitcsv.WithUseDisk(blockSize))
	}

	paths := flag.Args()

	if len(paths) == 0 {
		fatalf("missing file argument, e.g.: fitconv Activity.fit\n")
	}

	for _, path := range paths {
		if flagFitToCsv {
			if err := fitToCsv(path, fitToCsvOptions...); err != nil {
				fmt.Fprintf(os.Stderr, "could not convert %q to csv: %v\n", path, err)
			}
		}
	}
}

func fitToCsv(path string, opts ...fitcsv.Option) error {
	ff, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("could not open file: %s: %w", path, err)
	}
	defer ff.Close()

	base := filepath.Base(path)
	dir := filepath.Dir(path)
	ext := filepath.Ext(path)

	if ext != ".fit" {
		return fmt.Errorf("expected *.fit, got %s", ext)
	}

	name := base
	if len(ext) < len(base) {
		name = base[:len(base)-len(ext)]
	}

	namecsv := name + ".csv"
	pathcsv := filepath.Join(dir, namecsv)

	cf, err := os.OpenFile(pathcsv, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
	if err != nil {
		return fmt.Errorf("could not open file: %s: %w", pathcsv, err)
	}
	defer cf.Close()

	bw := bufio.NewWriterSize(cf, blockSize)
	conv := fitcsv.NewConverter(bw, opts...)

	dec := decoder.New(bufio.NewReaderSize(ff, blockSize),
		decoder.WithMesgDefListener(conv),
		decoder.WithMesgListener(conv),
		decoder.WithBroadcastOnly(),
	)

	var sequenceCounter int
	defer func() {
		if sequenceCounter == 0 {
			_ = os.Remove(pathcsv)
		}
	}()

	for dec.Next() {
		_, err = dec.Decode()
		if err != nil {
			return fmt.Errorf("decode failed: %w", err)
		}
		sequenceCounter++
	}

	conv.Wait()

	if err := conv.Err(); err != nil {
		return fmt.Errorf("convert done with error: %v", err)
	}

	if err := bw.Flush(); err != nil {
		return fmt.Errorf("could not flush buffered data: %w", err)
	}

	fmt.Printf("Converted! %s\n", filepath.Join(dir, namecsv))

	return nil
}

func fatalf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}
