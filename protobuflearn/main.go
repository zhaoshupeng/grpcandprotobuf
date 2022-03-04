package main

import (
	"fmt"
	"github.com/golang/protobuf/descriptor"
	"golearning/protobuflearn/option"
	"golearning/protobuflearn/score_server"
	// 导入protobuf依赖包
	"github.com/golang/protobuf/proto"
	// 导入我们刚才生成的go代码所在的包，注意你们自己的项目路径，可能跟本例子不一样
)

func main() {

	//getBaseMessage()
//-----------------------------------------------------
//	getFileOptions()

//-----------------------------------------------------

//	getMessageOptions()
//-----------------------------------------------------

	getFieldOptions()

}

func getBaseMessage() {
	// 初始化消息
	score_info := &score_server.BaseScoreInfoT{}
	score_info.WinCount = 10
	score_info.LoseCount = 1
	score_info.ExceptionCount = 2
	score_info.KillCount = 2
	score_info.DeathCount = 1
	score_info.AssistCount = 3
	score_info.Rating = 120

	// 同一时间只能设值一个值
	score_info.TestOneof = &score_server.BaseScoreInfoT_M1{
		M1:"this is m1",
	}

	fmt.Println("	m1:", score_info.GetM1())

	// 以字符串的形式打印消息
	fmt.Println("打印消息",score_info.String())

	// encode, 转换成二进制数据
	data, err := proto.Marshal(score_info)
	if err != nil {
		panic(err)
	}

	// decode, 将二进制数据转换成struct对象
	new_score_info := score_server.BaseScoreInfoT{}
	err = proto.Unmarshal(data, &new_score_info)
	if err != nil {
		panic(err)
	}

	// 以字符串的形式打印消息
	fmt.Println("以字符串的形式打印消息",new_score_info.String())
}

func getFileOptions() {
	fmt.Println("file options::")
	// 任何一个该包的结构体都可以
	//msg := &score_server.ExtObj{}
	msg := &score_server.BaseScoreInfoT{}
	md, _ := descriptor.MessageDescriptorProto(msg)
	stringOpt, _ := proto.GetExtension(md.Options, score_server.E_FileOptString)
	objOpt, _ := proto.GetExtension(md.Options, score_server.E_FileOptObj)
	fmt.Println("	obj.foo_string:", objOpt.(*score_server.ExtObjOri).FooString)
	fmt.Println("	obj.bar_int", objOpt.(*score_server.ExtObjOri).BarInt)
	fmt.Println("	string:", *stringOpt.(*string))
}

func getMessageOptions() {
	fmt.Println("message options:")
	msg := &option.MessageOption{}
	_, md := descriptor.MessageDescriptorProto(msg)
	objOpt, _ := proto.GetExtension(md.Options, option.E_MsgOptObj)
	stringOpt, _ := proto.GetExtension(md.Options, option.E_MsgOptString)
	fmt.Println("	obj.foo_string:", objOpt.(*option.ExtObj).FooString)
	fmt.Println("	obj.bar_int", objOpt.(*option.ExtObj).BarInt)
	fmt.Println("	string:", *stringOpt.(*string))
}

func getFieldOptions() {
	fmt.Println("field options:")
	msg := &option.FieldOption{}
	_, md := descriptor.MessageDescriptorProto(msg)
	stringOpt, _ := proto.GetExtension(md.Field[0].Options, option.E_FieldOptString)
	objOpt, _ := proto.GetExtension(md.Field[0].Options, option.E_FieldOptObj)
	fmt.Println("	obj.foo_string:", objOpt.(*option.ExtObj).FooString)
	fmt.Println("	obj.bar_int", objOpt.(*option.ExtObj).BarInt)
	fmt.Println("	string:", *stringOpt.(*string))
}

