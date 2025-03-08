package lsp

type InitializeRequest struct {
    Request
    Params InitializeRequestParams `json:"params"`
    
}

type InitializeRequestParams struct {
    ClientInfo *ClientInfo `json:"clientInfo,omitempty"`
    //Locale *Locale `json:"locale"`
    RootPath *string `json:"rootPath,omitempty"`
    RootUri *string `json:"rootUri,omitempty"`
    InitializationOptions *any `json:"initializationOptions,omitempty"`
    Capabilities ClientCapabilities `json:"capabilities"`
    WorkspaceFolders []WorkspaceFolder `json:"workspaceFolders"`
}

type ClientInfo struct {
    Name string `json:"name"`
    Version string `json:"version"`
}

type WorkspaceFolder struct {
    URI string `json:"uri"`         // we may later need to import something to handle parsing these
    Name string `json:"name"`
}

type Workspace struct {
    ApplyEdit *bool `json:"applyEdit,omitempty"`
    WorkspaceEdit *WorkspaceEditClientCapabilities `json:"workspaceEdit,omitempty"`
    DidChangeConfiguration *DidChangeConfigurationClientCapabilities `json:"didChangeConfiguration,omitempty"`
    DidChangeWatchedFiles *DidChangeWatchedFilesClientCapabilities `json:"didChangeWatchedFiles,omitempty"`
    Symbol *WorkspaceSymbolClientCapabilities `json:"symbol,omitempty"`
    ExecuteCommand *ExecuteCommandClientCapabilities `json:"executeCommand,omitempty"`
    WorkspaceFolders *bool `json:"workspaceFolders,omitempty"`
    Configuration *bool `json:"configuration,omitempty"`
    SemanticTokens *SemanticTokensWorkspaceClientCapabilities `json:"semanticTokens,omitempty"`
    CodeLens *CodeLensWorkspaceClientCapabilities `json:"codeLens,omitempty"`
    FileOperations *FileOperations `json:"fileOperations,omitempty"`
    InlineValue *InlineValueWorkspaceClientCapabilities `json:"inlineValue,omitempty"`
    InlayHint *InlayHintWorkspaceClientCapabilities `json:"inlayHint,omitempty"`
    Diagnostics *DiagnosticWorkspaceClientCapabilities `json:"diagnostics,omitempty"`
}

type FileOperations struct {
    DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
    DidCreate *bool `json:"didCreate,omitempty"`
    WillCreate *bool `json:"willCreate,omitempty"`
    DidRename *bool `json:"didRename,omitempty"`
    WillRename *bool `json:"willRename,omitempty"`
    DidDelete *bool `json:"didDelete,omitempty"`
    WillDelete *bool `json:"willDelete,omitempty"`
}

type InitializedResponse struct {
    Response 
    Result InitializeResult `json:"result"`
}

type InitializeResult struct {
    Capabilities ServerCapabilities `json:"capabilities"`
    ServerInfo ServerInfo `json:"serverInfo"`
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
