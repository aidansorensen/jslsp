package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"jslsp/analysis"
	"jslsp/lsp"
	"jslsp/rpc"
	"log"
	"os"
)

func main() {
    cwd, err := os.Getwd()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("cwd=", cwd)
    logger:=getLogger(cwd + "/log.txt")
    logger.Println("Hey, I started logging")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(rpc.Split)

    state := analysis.NewState()

    for scanner.Scan() {
        msg := scanner.Bytes()
        method, contents, err := rpc.DecodeMessage(msg)
        if err != nil {
            logger.Printf("Got an error: %s", err)
            continue
        }

        handleMessage(logger, state, method, contents)//pass logger and msg
    }
}

func handleMessage(logger *log.Logger,state analysis.State, method string, contents []byte) {
    logger.Printf("Received message with method %s", method)//literally just print the message to the log file

    switch method {
    case "initialize":
        var request lsp.InitializeRequest
        if err:= json.Unmarshal(contents, &request); err != nil {
            logger.Printf("hey, couldn't parse: %s", err)
        }
        logger.Printf("Connected to: %s %s", 
            request.Params.ClientInfo.Name,
            request.Params.ClientInfo.Version,)
    
        //reply now       
        msg := lsp.NewInitializeResponse(request.ID)
        reply := rpc.EncodeMessage(msg)

        writer := os.Stdout
        writer.Write([]byte(reply))

        logger.Print("Sent the reply")
    case "textDocument/didOpen":
        var request lsp.DidOpenTextDocumentNotification
        if err:= json.Unmarshal(contents, &request); err != nil {
            logger.Printf("didOpen: %s", err)
            return
        }

        logger.Printf("Opened: %s", request.Params.TextDocument.URI)
        state.OpenDocument(request.Params.TextDocument.URI, request.Params.TextDocument.Text) 
    case "textDocument/didChange":
        var request lsp.TextDocumentDidChangeNotification
        if err:= json.Unmarshal(contents, &request); err != nil {
            logger.Printf("didChange: %s", err)
            return
        }

        logger.Printf("Changed: %s", request.Params.TextDocument.URI, /*&request.Params.ContentChanges*/)
        for _, change := range request.Params.ContentChanges {//handle list of all changes
            state.UpdateDocument(request.Params.TextDocument.URI, change.Text)
        } 
        
    }
}

func getLogger(filename string) *log.Logger {
    logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
    if err != nil {
        panic("hey you didn't give me a good file")
    }
    return log.New(logfile, "[jslsp]", log.Ldate|log.LUTC|log.Lshortfile)
}
