/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"contract-shorturl-chainmaker/convert"
)

// 安装合约时会执行此方法，必须
//export init_contract
func initContract() {
	// 此处可写安装合约的初始化逻辑

}

// 升级合约时会执行此方法，必须
//export upgrade
func upgrade() {
	// 此处可写升级合约的逻辑
}

// shortUrl对象
type Short struct {
	shortUrl string
	longUrl string
	code string
	description string
	creator string
	version string
	createTime     int32 // second
	ec       *EasyCodec
}

// NewShort 新建Short对象
func NewShort(shortUrl string, longUrl string, code string,
	description string,creator string,version string,time int32) *Short {
	short := &Short{
		shortUrl: shortUrl,
		longUrl: longUrl,
		code:code,
		description:description,
		creator:creator,
		version:version,
		createTime:     time,
	}
	return short
}

// 获取序列化对象
func (f *Short) getEasyCodec() *EasyCodec {
	if f.ec == nil {
		f.ec = NewEasyCodec()
		f.ec.AddString("shortUrl", f.shortUrl)
		f.ec.AddString("longUrl", f.longUrl)
		f.ec.AddString("code", f.code)
		f.ec.AddString("description", f.description)
		f.ec.AddString("creator", f.creator)
		f.ec.AddString("version", f.version)
		f.ec.AddInt32("time", f.createTime)
	}
	return f.ec
}

// 序列化为json字符串
func (f *Short) toJson() string {
	return f.getEasyCodec().ToJson()
}

// 序列化为cmec编码
func (f *Short) marshal() []byte {
	return f.getEasyCodec().Marshal()
}

// 反序列化cmec为存证对象
func unmarshalToShort(data []byte) *Short {
	ec := NewEasyCodecWithBytes(data)
	shortUrl, _ := ec.GetString("shortUrl")
	longUrl, _ := ec.GetString("longUrl")
	code, _ := ec.GetString("code")
	description, _ := ec.GetString("description")
	creator, _ := ec.GetString("creator")
	version, _ := ec.GetString("version")
	time, _ := ec.GetInt32("time")

	short := &Short{
		shortUrl: shortUrl,
		longUrl: longUrl,
		code:code,
		description:description,
		creator:creator,
		version:version,
		createTime:     time,
	}
	return short
}

// 对外暴露 save 方法，供用户由 SDK 调用
//export save
func save() {
	// 获取上下文
	ctx := NewSimContext()

	// 获取参数
	shortUrl, err1 := ctx.ArgString("short_url")
	longUrl, err2 := ctx.ArgString("long_url")
	compressionCode, err3 := ctx.ArgString("code")
	description, err4 := ctx.ArgString("description")
	creator, err5 := ctx.ArgString("creator")
	version, err6 := ctx.ArgString("version")
	timeStr, err7 := ctx.ArgString("time")

	if err1 != SUCCESS || err2 != SUCCESS || err3 != SUCCESS ||
		err4 != SUCCESS || err5 != SUCCESS || err6 != SUCCESS || err7 != SUCCESS{
		ctx.Log("get arg fail.")
		ctx.ErrorResult("get arg fail.")
		return
	}
	time, err := convert.StringToInt32(timeStr)
	if err != nil {
		ctx.ErrorResult(err.Error())
		ctx.Log(err.Error())
		return
	}
	// 构建结构体
	short := NewShort(shortUrl,longUrl,compressionCode,description,creator,version, time)
	// 序列化：两种方式
	jsonStr := short.toJson()
	//发送事件
	ctx.EmitEvent("topic_short",short.code)
	// 存储数据
	ctx.PutState("short_json", short.code, jsonStr)
	// 记录日志
	ctx.Log("【save】 code=" + short.code)
	// 返回结果
	ctx.SuccessResult(short.code)
}

// 对外暴露 find_by_code 方法，供用户由 SDK 调用
//export find_by_code
func findByCode() {
	ctx := NewSimContext()
	// 获取参数
	code, _ := ctx.ArgString("code")
	// 查询Json
	if result, resultCode := ctx.GetStateByte("short_json", code); resultCode != SUCCESS {
		// 返回结果
		ctx.ErrorResult("failed to call get_state, only 64 letters and numbers are allowed. got key:" + "short" + ", field:" + code)
	} else {
		// 返回结果
		ctx.SuccessResultByte(result)
		// 记录日志
		ctx.Log("get val:" + string(result))
	}
}

func main() {

}
