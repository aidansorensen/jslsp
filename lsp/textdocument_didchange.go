package lsp

type TextDocumentDidChangeNotification struct {
    Notification
    Params DidChangeTextDocumentParams `json:"params"`
}

type DidChangeTextDocumentParams struct {
    TextDocument VersionTextDocumentIdentifier `json:"textDocument"`
    ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges"`
}

type TextDocumentContentChangeEvent struct {
	/**
	 * The range of the document that changed.
	 */
//	range: Range;

	/**
	 * The optional length of the range that got replaced.
	 *
	 * @deprecated use range instead.
	 */
//	rangeLength?: uinteger;

	/**
	 * The new text for the provided range.
	 */
//	text: string;
//} | {
	/**
	 * The new text of the whole document.
	 */
//	text: string;
    Text string `json:"text"`
};
