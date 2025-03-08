package lsp

type ClientCapabilities struct {
    Workspace *Workspace `json:"workspace,omitempty"`
    TextDocument *TextDocumentClientCapabilities `json:"textDocument,omitempty"`
    NotebookDocument *NotebookDocumentClientCapabilities `json:"notebookDocument,omitempty"`
    Window *WindowCapabilities `json:"window,omitempty"`
}

type WorkspaceEditClientCapabilities struct {   // TODO
}

type DidChangeConfigurationClientCapabilities struct {   // TODO
}

type DidChangeWatchedFilesClientCapabilities struct {   // TODO
}

type WorkspaceSymbolClientCapabilities struct {   // TODO
}

type ExecuteCommandClientCapabilities struct {   // TODO
}

type SemanticTokensWorkspaceClientCapabilities struct {   // TODO
}

type CodeLensWorkspaceClientCapabilities struct {   // TODO
}

type ServerCapabilities struct{
    TextDocumentSync int `json:"textDocumentSync"`
    HoverProvider bool `json:"hoverProvider"`
}

type InlineValueWorkspaceClientCapabilities struct {    // TODO
}

type InlayHintWorkspaceClientCapabilities struct {  // TODO
}

type DiagnosticWorkspaceClientCapabilities struct { // TODO
}
