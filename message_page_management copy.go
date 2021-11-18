/**
用于生成要发送的页面控制指令

ver 1.0
*/

package message

import (
	"encoding/json"
)

/**
页面控制消息格式
*/
type Page_management_message struct {
	CMDType  int        `json:"type"`
	OpCode   int        `json:"op"`
	URL      string     `json:"url"`
	Pos      [2]float32 `json:"pos"`
	Message  string     `json:"msg"`
	Tiemout  float32    `json:"timeout"`
	DebugCMD string     `json:"cmd"`
}

/**
创建页面操作空指令消息 (op=0)
*/
func PageNoopMessage() ([]byte, error) {
	var message = &Page_management_message{
		CMDType: 3,
		OpCode:  0,
	}
	return json.Marshal(message)
}

/**
创建刷新页面指令(op=1)
*/
func PageRefreshMessage() ([]byte, error) {
	var message = &Page_management_message{
		CMDType: 3,
		OpCode:  1,
	}
	return json.Marshal(message)
}

/**
创建关闭页面指令(op=2)
*/
func PageCloseMessage() ([]byte, error) {
	var message = &Page_management_message{
		CMDType: 3,
		OpCode:  2,
	}
	return json.Marshal(message)
}

/**
创建页面跳转消息(op=3)
*/
func PageJumpToMessage(url string) ([]byte, error) {
	var message = &Page_management_message{
		CMDType: 3,
		OpCode:  3,
		URL:     url,
	}
	return json.Marshal(message)
}

/**
创建显示悬浮窗消息(op=4)
*/
func PageShowHoverWindowMessage(contentURL string, position [2]float32) ([]byte, error) {
	var message = &Page_management_message{
		CMDType: 3,
		URL:     contentURL,
		Pos:     position,
		OpCode:  4,
	}
	return json.Marshal(message)
}

/**
创建非阻塞式信息消息(op=5)
*/
func PageShowInfoMessage(content string, timeout float32) ([]byte, error) {
	var message = &Page_management_message{
		CMDType: 3,
		OpCode:  5,
		Message: content,
		Tiemout: timeout,
	}
	return json.Marshal(message)
}

/**
创建非阻塞式警告信息消息(op=6)
*/
func PageShowWarningMessage(content string, timeout float32) ([]byte, error) {
	var message = &Page_management_message{
		CMDType: 3,
		OpCode:  6,
		Message: content,
		Tiemout: timeout,
	}
	return json.Marshal(message)
}

/**
创建非阻塞式警告信息消息(op=165)
*/
func PageShowCriticalMessage(content string, timeout float32) ([]byte, error) {
	var message = &Page_management_message{
		CMDType: 3,
		OpCode:  165,
		Message: content,
		Tiemout: timeout,
	}
	return json.Marshal(message)
}

/**
创建调试指令（js代码）消息(op=255)
仅用于debug版本的前端类库
*/
func PageDebugMessage(js string) ([]byte, error) {
	var message = &Page_management_message{
		CMDType:  1,
		OpCode:   255,
		DebugCMD: js,
	}
	return json.Marshal(message)
}
