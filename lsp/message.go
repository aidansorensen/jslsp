package lsp

type Request struct{
    RPC string `json:"jsonrpc"`
    ID int `json:"id"`
    Method string `json:"method"`


    
}

type Response struct{
    RPC string `json:"jsonrpc"`
    ID *int `json:"id,omitempty"` 

    //result 
    //error


}

type Notification struct{
    RPC string `json:"jsonrpc"`
    Method string `json:"method"`
}
