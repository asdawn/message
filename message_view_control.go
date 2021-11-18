/**
用于生成要发送的地图视图控制指令

ver 1.0
*/

package message

import "encoding/json"

/**
地图操作0: 无操作
*/
type View_control_message struct {
	CMDType   int        `json:"type"`
	Operation int        `json:"op"`
	Zoom      int        `json:"zoom"`
	Rotation  int        `json:"rr"`
	Position  [2]float32 `json:"pos"`
	Extent    [4]float32 `json:"extent"`
}

/**
创建地图操作0：无操作消息
*/
func ViewNoopMessage() ([]byte, error) {
	var message = &View_control_message{
		CMDType:   4,
		Operation: 0,
	}
	return json.Marshal(message)
}

/**
创建地图操作1：刷新view消息
*/
func ViewRefreshMessage() ([]byte, error) {
	var message = &View_control_message{
		CMDType:   4,
		Operation: 1,
	}
	return json.Marshal(message)
}

/**
创建地图操作2：放大一级消息
*/
func ViewZoomInMessage() ([]byte, error) {
	var message = &View_control_message{
		CMDType:   4,
		Operation: 2,
	}
	return json.Marshal(message)
}

/**
创建地图操作3：缩小一级消息
*/
func ViewZoomOutMessage() ([]byte, error) {
	var message = &View_control_message{
		CMDType:   4,
		Operation: 3,
	}
	return json.Marshal(message)
}

/**
创建地图操作4: 缩放到指定级别消息
zoom: 目标缩放级别
*/
func ViewZoomToMessage(zoom int) ([]byte, error) {
	var message = &View_control_message{
		CMDType:   4,
		Operation: 4,
		Zoom:      zoom,
	}
	return json.Marshal(message)
}

/**
创建地图操作5: 设定地图旋转角度消息
rotation: 目标旋转角度
*/
func ViewRotationMessage(rotation int) ([]byte, error) {
	var message = &View_control_message{
		CMDType:   4,
		Operation: 5,
		Rotation:  rotation,
	}
	return json.Marshal(message)
}

/**
创建地图操作6: 设定地图中心点消息
x: 经度
y: 纬度
*/
func ViewMoveToMessage(x float32, y float32) ([]byte, error) {
	var message = &View_control_message{
		CMDType:   4,
		Operation: 6,
		Position:  [2]float32{x, y},
	}
	return json.Marshal(message)
}

/**
创建地图操作7: 设定地图显示范围消息
xmin: 最小经度
ymin: 最小纬度
xmax: 最大经度
ymax: 最大纬度
*/
func ViewSetExtentMessage(xmin float32, ymin float32, xmax float32, ymax float32) ([]byte, error) {
	var message = &View_control_message{
		CMDType:   4,
		Operation: 7,
		Extent:    [4]float32{xmin, ymin, xmax, ymax},
	}
	return json.Marshal(message)
}
