package lsp

type TextDocumentItem struct {
    URI string `json:"uri"`
    LanguageID string `json:"languageId"`
    Version int`json:"version"`
    Text string`json:"text"`
}

type TextDocumentIdentifier struct {
	URI string `json:"uri"`
}
type VersionTextDocumentIdentifier struct {
	TextDocumentIdentifier
	Version int `json:"version"`
}
