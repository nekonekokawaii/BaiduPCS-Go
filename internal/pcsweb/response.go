package pcsweb

import (
	"encoding/json"
	"golang.org/x/net/websocket"
	"net/http"
	"github.com/qjfoidnh/BaiduPCS-Go/pcsverbose"
)

type pcsConfigJSON struct {
	Name string `json:"name"`
	EnName string `json:"en_name"`
	Value string `json:"value"`
	Desc string `json:"desc"`
}

type pcsOptsJSON struct {
	Name string `json:"name"`
	Value bool `json:"value"`
	Desc string `json:"desc"`
}

type Response struct {
	Code int         `json:"code"`
	Type int         `json:"type"`
	Status int       `json:"status"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (res *Response) JSON() (data []byte) {
	data, _ = json.Marshal(res)
	return
}


func sendResponse(conn *websocket.Conn, rtype int, rstatus int, msg string, data string) (err error){
	response := &Response{
		Code: 0,
		Type: rtype,
		Status: rstatus,
		Msg: msg,
		Data: data,
	}
	if err = websocket.Message.Send(conn, string(response.JSON())); err != nil {
		pcsverbose.Verbosef("websocket send error: %s\n", err.Error())
		return err
	}
	return nil
}

func sendErrorResponse(conn *websocket.Conn, rcode int, msg string) (err error){
	response := &Response{
		Code: rcode,
		Type: 0,
		Status: 0,
		Msg: msg,
		Data: "",
	}
	if err = websocket.Message.Send(conn, string(response.JSON())); err != nil {
		pcsverbose.Verbosef("websocket send error: %s\n", err.Error())
		return err
	}
	return nil
}

func sendHttpErrorResponse(w http.ResponseWriter, rcode int, msg string) {
	response := &Response{
		Code: rcode,
		Type: 0,
		Status: 0,
		Msg: msg,
		Data: "",
	}
	w.Write(response.JSON())
}

func sendHttpResponse(w http.ResponseWriter, msg string, data interface{}) {
	response := &Response{
		Code: 0,
		Type: 0,
		Status: 0,
		Msg: msg,
		Data: data,
	}
	w.Write(response.JSON())
}

var (
	NeedPass = -3
	NotLogin = -4
	LoginError = -5
)
