package lsp

type HoverRequest struct {
    Request
    Params HoverParams `json:"params"`
}

type HoverParams struct {
    TextDocumentPositionParams
}

type HoverResult struct {
    value string
}

type HoverResponse struct {
    Response
    Result HoverResult `json:"result"`
}
