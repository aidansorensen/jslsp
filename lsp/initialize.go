package lsp

type InitializeRequest struct {
    Request
    Params InitializeRequestParams `json:"params"`
    
}


type InitializeRequestParams struct {
    ClientInfo *ClientInfo `json:"clientinfo"`
}


type ClientInfo struct {
    Name string `json:"name"`
    Version string `json:"version"`
}

type InitializedResponse struct {
    Response 
    Result InitializeResult `json:"result"`
}

type InitializeResult struct {
    Capabilities ServerCapabilities `json:"capabilities"`
    ServerInfo ServerInfo `json:"serverInfo"`
}

type ServerCapabilities struct{
    TextDocumentSync int `json:"textDocumentSync"`

    HoverProvider bool `json:"hoverProvider"`
}

type ServerInfo struct {
    Name string `json:"name"`
    Version string `json:"version"`
}


func NewInitializeResponse (id int) InitializedResponse {
    return InitializedResponse{
        Response: Response{
            RPC: "2.0",
            ID: &id,
        },
        Result: InitializeResult{
            Capabilities: ServerCapabilities{
            TextDocumentSync: 1,//meaning full, documents are synced by sending full contents of file...
            HoverProvider: true,
            },
            ServerInfo: ServerInfo{
                Name: "jslsp",
                Version: "0.1.beta",
            },
        },
    }
}
