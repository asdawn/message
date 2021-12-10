/**
用于生成要发送的地图对象控制指令

ver 1.2 追加keep参数
ver 1.1 添加用于格式转换的Version_convert方法和
ver 1.0
*/

package message

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/asdawn/device"
)

/**
新建/更新对象
*/
type Object_management_message struct {
	CMDType      int              `json:"type"`
	ObjectClass  string           `json:"class"`
	ValuesSet    []*device.Device `json:"set"`
	ValuesDelete []string         `json:"del"`
	ValuesKeep   []string         `json:"keep"`
	ValuesClear  bool             `json:"clear"`
}

/**
新建/更新对象（字符串时间戳版本）
*/
type Object_management_message1 struct {
	CMDType      int               `json:"type"`
	ObjectClass  string            `json:"class"`
	ValuesSet    []*device.Device1 `json:"set"`
	ValuesDelete []string          `json:"del"`
	ValuesKeep   []string          `json:"keep"`
	ValuesClear  bool              `json:"clear"`
}

/**
创建新建/更新对象消息
objectClass: 对象类名
values: 对象状态值数组，根据ID进行更新/创建
*/
func ObjectUpsertMessage(objectClass string, values []*device.Device) ([]byte, error) {
	var message = &Object_management_message{
		CMDType:     2,
		ObjectClass: objectClass,
		ValuesSet:   values,
	}
	return json.Marshal(message)
}

/**
创建新建/更新+删除对象消息
objectClass: 对象类名
toSet: 对象状态值数组，根据ID进行更新/创建
toDelete: id数组，用于删除
*/
func ObjectMessage(objectClass string, toSet []*device.Device, toDelete []string) ([]byte, error) {
	var message = &Object_management_message{
		CMDType:      2,
		ObjectClass:  objectClass,
		ValuesSet:    toSet,
		ValuesDelete: toDelete,
	}
	return json.Marshal(message)
}

/**
创建对象操作空指令消息
*/
func ObjectNoopMessage(objectClass string) ([]byte, error) {
	var message = &Object_management_message{
		CMDType:     2,
		ObjectClass: objectClass,
	}
	return json.Marshal(message)
}

/**
创建清空对象消息
*/
func ObjectClearMessage(objectClass string) ([]byte, error) {
	var message = &Object_management_message{
		CMDType:     2,
		ObjectClass: objectClass,
		ValuesClear: true,
	}
	return json.Marshal(message)
}

/**
创建删除对象消息
objectClass: 对象类名
ids: 要删除的对象ID数组
*/
func ObjectDeleteMessage(objectClass string, ids []string) ([]byte, error) {
	var message = &Object_management_message{
		CMDType:      2,
		ObjectClass:  objectClass,
		ValuesDelete: ids,
	}
	return json.Marshal(message)
}

/**
创建（仅）保留指定对象消息
objectClass: 对象类名
ids: 要保留的对象ID数组
*/
func ObjectKeepMessage(objectClass string, ids []string) ([]byte, error) {
	var message = &Object_management_message{
		CMDType:     2,
		ObjectClass: objectClass,
		ValuesKeep:  ids,
	}
	return json.Marshal(message)
}

/**
创建一类对象的全状态消息
deviceSet: 对象状体集
返回（消息，错误），如果deviceSet中无对象则消息返回nil而非空指令消息
*/
func ObjectFullStatusMessage(deviceSet *device.DeviceSet) ([]byte, error) {
	/*
		if deviceSet == nil {
			return nil, errors.New("deviceset should not be null")
		}
	*/
	objectClass := deviceSet.DeviceClass
	if len(deviceSet.Devices) == 0 {
		return ObjectNoopMessage(objectClass)
	} else {
		devices := deviceSet.GetDevices()
		return ObjectUpsertMessage(objectClass, devices)
	}
}

/*
消息格式转换，时间戳由字符串转为unix时间戳
data1: 字符串时间戳的消息（Object_management_message1）
返回: (Unix时间戳版消息，错误)
*/
func Version_convert(data1 []byte) ([]byte, error) {
	message1 := &Object_management_message1{}
	err := json.Unmarshal(data1, message1)
	if err != nil {
		log.Println("Invalid json:" + string(data1))
		return nil, err
	} else {
		message := &Object_management_message{
			CMDType:      message1.CMDType,
			ObjectClass:  message1.ObjectClass,
			ValuesDelete: message1.ValuesDelete,
			ValuesClear:  message1.ValuesClear,
			ValuesKeep:   message1.ValuesKeep,
		}
		devices1 := message1.ValuesSet
		devices := make([]*device.Device, 0)
		for _, device1 := range devices1 {
			tString := device1.T
			var t int64
			loc, _ := time.LoadLocation("Local")
			t1, _ := time.ParseInLocation("2006-01-02 15:04:05", tString, loc)
			t = t1.Unix()
			//坐标舍入到小数后6位
			x, err := keep6(device1.X)
			if err != nil {
				return nil, err
			}
			y, err := keep6(device1.Y)
			if err != nil {
				return nil, err
			}
			dvc := &device.Device{
				ID:     device1.ID,
				X:      x,
				Y:      y,
				R:      device1.R,
				Status: device1.Status,
				T:      t,
				Color:  device1.Color,
			}
			devices = append(devices, dvc)
		}
		(*message).ValuesSet = devices
		data, err := json.Marshal(message)
		if err != nil {
			return nil, err
		} else {
			return data, nil
		}
	}
}

/*
保留6位小数，相当于精确到分米
*/
func keep6(value float32) (float32, error) {
	s := fmt.Sprintf("%.6f", value)
	newValue, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return value, err
	} else {
		return float32(newValue), nil
	}
}

/*
func main() {
	var msg1 string = `{"set":[{"exc":"","r":0,"s":1,"t":"2021-11-28 21:59:00","ot":"1","jp":"0","x":124.4567890,"y":23.4567890,"id":"TT","exb":"","exa":""}],"type":2,"class":"truck"}`
	var msg, err = version_convert([]byte(msg1))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(msg))
	}
}
*/
