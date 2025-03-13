package lsp

type DefinitionRequest struct {
	Request
	Params DefinitionParams `json:"params"`
}

type DefinitionParams struct {
	TextDocumentPositionParams
}

type DefinitionResult struct {
	Contents string `json:"contents"`
}

type DefinitionResponse struct {
	Response
	Result Location `json:"result"`
}
