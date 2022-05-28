/*----------------------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Licensed under the MIT License. See LICENSE in the project root for license information.
 *---------------------------------------------------------------------------------------*/

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/microsoft/vscode-remote-try-go/hello"
)

func handle(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, hello.Hello())
}

func main() {
	portNumber := os.Getenv("PORT")
	if portNumber == "" {
		portNumber = "8000"
	}
	http.HandleFunc("/", handle)
	fmt.Println("Server listening on port hogehsssssaoge ", portNumber)
	http.ListenAndServe(":"+portNumber, nil)
}
