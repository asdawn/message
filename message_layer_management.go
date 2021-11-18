/**
用于生成要发送的图层控制指令

ver 1.0
*/

package message

import (
	"encoding/json"
)

/**
新建/更新对象
*/
type Layer_management_message struct {
	CMDType int    `json:"type"`
	OpCode  int    `json:"op"`
	Layer   string `json:"class"`
}

/**
创建对象操作空指令消息 (op=0)
*/
func LayerNoopMessage() ([]byte, error) {
	var message = &Layer_management_message{
		CMDType: 1,
		OpCode:  0,
	}
	return json.Marshal(message)
}

/**
创建刷新指定动态图层指令(op=1)
*/
func LayerRefreshDynamicMessage(layerName string) ([]byte, error) {
	var message = &Layer_management_message{
		CMDType: 1,
		Layer:   layerName,
		OpCode:  1,
	}
	return json.Marshal(message)
}

/**
刷新全部动态图层消息(op=2)
*/
func LayerRefreshAllDynamicMessage() ([]byte, error) {
	var message = &Layer_management_message{
		CMDType: 1,
		OpCode:  2,
	}
	return json.Marshal(message)
}

/**
刷新指定图层消息(op=3)
*/
func LayerRefreshMessage(layerName string) ([]byte, error) {
	var message = &Layer_management_message{
		CMDType: 1,
		Layer:   layerName,
		OpCode:  3,
	}
	return json.Marshal(message)
}

/**
刷新全部图层消息(op=4)
*/
func LayerRefreshAllMessage() ([]byte, error) {
	var message = &Layer_management_message{
		CMDType: 1,
		OpCode:  4,
	}
	return json.Marshal(message)
}

/**
显示指定图层消息(op=5)
*/
func LayerShowMessage(layerName string) ([]byte, error) {
	var message = &Layer_management_message{
		CMDType: 1,
		Layer:   layerName,
		OpCode:  5,
	}
	return json.Marshal(message)
}

/**
隐藏指定图层消息(op=6)
*/
func LayerHideMessage(layerName string) ([]byte, error) {
	var message = &Layer_management_message{
		CMDType: 1,
		Layer:   layerName,
		OpCode:  6,
	}
	return json.Marshal(message)
}

/**
清空指定图层消息(op=165)
*/
func LayerClearFeaturesMessage(layerName string) ([]byte, error) {
	var message = &Layer_management_message{
		CMDType: 1,
		Layer:   layerName,
		OpCode:  165,
	}
	return json.Marshal(message)
}
