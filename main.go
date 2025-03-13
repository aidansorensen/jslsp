package main

import (
	"bufio"
	"encoding/json"
	"io"

	//"fmt"
	"jslsp/analysis"
	"jslsp/lsp"
	"jslsp/rpc"
	"log"
	"os"
)

func main() {
	/*** commented out for now, need to figure out how to put this log in a reasonable place
	  cwd, err := os.Getwd()
	  if err != nil {
	      fmt.Println("Error:", err)
	      return
	  }
	  logger:=getLogger(cwd + "/jslsp_log.txt")
	*/
	logger := getLogger("/home/bmalt/repos/jslsp/log.txt")
	logger.Println("Hey, I started logging")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	state := analysis.NewState()
	writer := os.Stdout

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Got an error: %s", err)
			continue
		}

		handleMessage(logger, writer, state, method, contents) //pass logger and msg
	}
}

func handleMessage(logger *log.Logger, writer io.Writer, state analysis.State, method string, contents []byte) {
	logger.Printf("Received message with method %s", method) //literally just print the message to the log file

	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("hey, couldn't parse: %s", err)
		}
		logger.Printf("Connected to: %s %s",
			request.Params.ClientInfo.Name,
			request.Params.ClientInfo.Version)

		//reply now
		msg := lsp.NewInitializeResponse(request.ID)
		writeResponse(writer, msg)

		logger.Print("Sent the reply")
	case "textDocument/didOpen":
		var request lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("didOpen: %s", err)
			return
		}

		logger.Printf("Opened: %s", request.Params.TextDocument.URI)
		state.OpenDocument(request.Params.TextDocument.URI, request.Params.TextDocument.Text)
	case "textDocument/didChange":
		var request lsp.TextDocumentDidChangeNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("didChange: %s", err)
			return
		}

		logger.Printf("Changed: %s", request.Params.TextDocument.URI /*&request.Params.ContentChanges*/)
		for _, change := range request.Params.ContentChanges { //handle list of all changes
			state.UpdateDocument(request.Params.TextDocument.URI, change.Text)
		}
    case "textDocument/hover":
        var request lsp.HoverRequest
        if err := json.Unmarshal(contents, &request); err != nil {
            logger.Printf("textDocument/hover: %s", err)
            return
        }
        
        response := state.Hover(request.ID, request.Params.TextDocument.URI, request.Params.Position)

        writeResponse(writer, response)
    case "textDocument/definition":
        var request lsp.DefinitionRequest
        if err := json.Unmarshal(contents, &request); err != nil {
            logger.Printf("textDocument/definition: %s", err)
            return
        }
        
        response := state.Definition(request.ID, request.Params.TextDocument.URI, request.Params.Position)

        writeResponse(writer, response)
    }
}

func writeResponse(writer io.Writer, msg any) {
	reply := rpc.EncodeMessage(msg)
	writer.Write([]byte(reply))
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("hey you didn't give me a good file")
	}
	return log.New(logfile, "[jslsp] ", log.LstdFlags|log.Lmsgprefix|log.Lshortfile)
}
