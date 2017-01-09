// Copyright (C) 2017  Sylvain Laurent

// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation; either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program; if not, write to the Free Software Foundation,
// Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301  USA

package main

import (
	"flag"
	"log"
	"net/http"
	"github.com/ethereum/go-ethereum/logger"
	"github.com/ethereum/go-ethereum/logger/glog"
	"github.com/ethereum/ethash"
)

func GenerateDAG(block_number uint64, dir string) {
	err := ethash.MakeDAG(block_number, dir)
	if err != nil {
		log.Printf("Error while generating DAG: %v", err)
	}
}

func main() {
	host := flag.String("host", "0.0.0.0", "Address to serve on")
	port := flag.String("port", "8091", "Port to server on")
	directory := flag.String("dir", "/tmp", "the directory of tmp dag file to host")
	flag.Parse()

	glog.SetV(logger.Info)
	glog.SetToStderr(true)


	GenerateDAG(0, *directory + "/dag/")
	http.Handle("/dag/", http.FileServer(http.Dir(*directory)))

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Fatal(http.ListenAndServe(*host + ":" + *port, nil))
}
