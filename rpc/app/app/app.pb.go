/*
 * # 此文件版权属于ASO.DESIGN
 * # 文件：app.pb.go  项目：aso-zero
 * # 作用：
 * # 当前修改时间：2021年07月16日 00:54:39
 * # 上次修改时间：2021年07月15日 23:04:39
 * # 作者：thunur
 * # 此文件不可非法传播、倒卖、共享，否则我们将追究相应的法律责任。
 * # 您如果已获得ASO.DESIGN授权可在原有基础上进行修改使用
 * # 如果您还没获得授权请联系我们 thunur@qq.com
 * # Copyright (c) 2021 aso.design
 */

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app.proto

package app

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// 相册分类
type AppListTypeReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppListTypeReq) Reset()         { *m = AppListTypeReq{} }
func (m *AppListTypeReq) String() string { return proto.CompactTextString(m) }
func (*AppListTypeReq) ProtoMessage()    {}
func (*AppListTypeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f9056a14b86d47, []int{0}
}

func (m *AppListTypeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppListTypeReq.Unmarshal(m, b)
}
func (m *AppListTypeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppListTypeReq.Marshal(b, m, deterministic)
}
func (m *AppListTypeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppListTypeReq.Merge(m, src)
}
func (m *AppListTypeReq) XXX_Size() int {
	return xxx_messageInfo_AppListTypeReq.Size(m)
}
func (m *AppListTypeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AppListTypeReq.DiscardUnknown(m)
}

var xxx_messageInfo_AppListTypeReq proto.InternalMessageInfo

type AppListTypeData struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreateTime           string   `protobuf:"bytes,2,opt,name=createTime,proto3" json:"createTime,omitempty"`
	UpdateTime           string   `protobuf:"bytes,3,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppListTypeData) Reset()         { *m = AppListTypeData{} }
func (m *AppListTypeData) String() string { return proto.CompactTextString(m) }
func (*AppListTypeData) ProtoMessage()    {}
func (*AppListTypeData) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f9056a14b86d47, []int{1}
}

func (m *AppListTypeData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppListTypeData.Unmarshal(m, b)
}
func (m *AppListTypeData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppListTypeData.Marshal(b, m, deterministic)
}
func (m *AppListTypeData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppListTypeData.Merge(m, src)
}
func (m *AppListTypeData) XXX_Size() int {
	return xxx_messageInfo_AppListTypeData.Size(m)
}
func (m *AppListTypeData) XXX_DiscardUnknown() {
	xxx_messageInfo_AppListTypeData.DiscardUnknown(m)
}

var xxx_messageInfo_AppListTypeData proto.InternalMessageInfo

func (m *AppListTypeData) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AppListTypeData) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *AppListTypeData) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

func (m *AppListTypeData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type AppListTypeResp struct {
	List                 []*AppListTypeData `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AppListTypeResp) Reset()         { *m = AppListTypeResp{} }
func (m *AppListTypeResp) String() string { return proto.CompactTextString(m) }
func (*AppListTypeResp) ProtoMessage()    {}
func (*AppListTypeResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f9056a14b86d47, []int{2}
}

func (m *AppListTypeResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppListTypeResp.Unmarshal(m, b)
}
func (m *AppListTypeResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppListTypeResp.Marshal(b, m, deterministic)
}
func (m *AppListTypeResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppListTypeResp.Merge(m, src)
}
func (m *AppListTypeResp) XXX_Size() int {
	return xxx_messageInfo_AppListTypeResp.Size(m)
}
func (m *AppListTypeResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AppListTypeResp.DiscardUnknown(m)
}

var xxx_messageInfo_AppListTypeResp proto.InternalMessageInfo

func (m *AppListTypeResp) GetList() []*AppListTypeData {
	if m != nil {
		return m.List
	}
	return nil
}

type AppAddTypeReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppAddTypeReq) Reset()         { *m = AppAddTypeReq{} }
func (m *AppAddTypeReq) String() string { return proto.CompactTextString(m) }
func (*AppAddTypeReq) ProtoMessage()    {}
func (*AppAddTypeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f9056a14b86d47, []int{3}
}

func (m *AppAddTypeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppAddTypeReq.Unmarshal(m, b)
}
func (m *AppAddTypeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppAddTypeReq.Marshal(b, m, deterministic)
}
func (m *AppAddTypeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppAddTypeReq.Merge(m, src)
}
func (m *AppAddTypeReq) XXX_Size() int {
	return xxx_messageInfo_AppAddTypeReq.Size(m)
}
func (m *AppAddTypeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AppAddTypeReq.DiscardUnknown(m)
}

var xxx_messageInfo_AppAddTypeReq proto.InternalMessageInfo

func (m *AppAddTypeReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type AppAddTypeResp struct {
	Pong                 int64    `protobuf:"varint,1,opt,name=pong,proto3" json:"pong,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppAddTypeResp) Reset()         { *m = AppAddTypeResp{} }
func (m *AppAddTypeResp) String() string { return proto.CompactTextString(m) }
func (*AppAddTypeResp) ProtoMessage()    {}
func (*AppAddTypeResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f9056a14b86d47, []int{4}
}

func (m *AppAddTypeResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppAddTypeResp.Unmarshal(m, b)
}
func (m *AppAddTypeResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppAddTypeResp.Marshal(b, m, deterministic)
}
func (m *AppAddTypeResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppAddTypeResp.Merge(m, src)
}
func (m *AppAddTypeResp) XXX_Size() int {
	return xxx_messageInfo_AppAddTypeResp.Size(m)
}
func (m *AppAddTypeResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AppAddTypeResp.DiscardUnknown(m)
}

var xxx_messageInfo_AppAddTypeResp proto.InternalMessageInfo

func (m *AppAddTypeResp) GetPong() int64 {
	if m != nil {
		return m.Pong
	}
	return 0
}

type AppUpdateTypeReq struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppUpdateTypeReq) Reset()         { *m = AppUpdateTypeReq{} }
func (m *AppUpdateTypeReq) String() string { return proto.CompactTextString(m) }
func (*AppUpdateTypeReq) ProtoMessage()    {}
func (*AppUpdateTypeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f9056a14b86d47, []int{5}
}

func (m *AppUpdateTypeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppUpdateTypeReq.Unmarshal(m, b)
}
func (m *AppUpdateTypeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppUpdateTypeReq.Marshal(b, m, deterministic)
}
func (m *AppUpdateTypeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppUpdateTypeReq.Merge(m, src)
}
func (m *AppUpdateTypeReq) XXX_Size() int {
	return xxx_messageInfo_AppUpdateTypeReq.Size(m)
}
func (m *AppUpdateTypeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AppUpdateTypeReq.DiscardUnknown(m)
}

var xxx_messageInfo_AppUpdateTypeReq proto.InternalMessageInfo

func (m *AppUpdateTypeReq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AppUpdateTypeReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type AppUpdateTypeResp struct {
	Pong                 int64    `protobuf:"varint,1,opt,name=pong,proto3" json:"pong,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppUpdateTypeResp) Reset()         { *m = AppUpdateTypeResp{} }
func (m *AppUpdateTypeResp) String() string { return proto.CompactTextString(m) }
func (*AppUpdateTypeResp) ProtoMessage()    {}
func (*AppUpdateTypeResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f9056a14b86d47, []int{6}
}

func (m *AppUpdateTypeResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppUpdateTypeResp.Unmarshal(m, b)
}
func (m *AppUpdateTypeResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppUpdateTypeResp.Marshal(b, m, deterministic)
}
func (m *AppUpdateTypeResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppUpdateTypeResp.Merge(m, src)
}
func (m *AppUpdateTypeResp) XXX_Size() int {
	return xxx_messageInfo_AppUpdateTypeResp.Size(m)
}
func (m *AppUpdateTypeResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AppUpdateTypeResp.DiscardUnknown(m)
}

var xxx_messageInfo_AppUpdateTypeResp proto.InternalMessageInfo

func (m *AppUpdateTypeResp) GetPong() int64 {
	if m != nil {
		return m.Pong
	}
	return 0
}

type AppDeleteTypeReq struct {
	Ids                  []int64  `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppDeleteTypeReq) Reset()         { *m = AppDeleteTypeReq{} }
func (m *AppDeleteTypeReq) String() string { return proto.CompactTextString(m) }
func (*AppDeleteTypeReq) ProtoMessage()    {}
func (*AppDeleteTypeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f9056a14b86d47, []int{7}
}

func (m *AppDeleteTypeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppDeleteTypeReq.Unmarshal(m, b)
}
func (m *AppDeleteTypeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppDeleteTypeReq.Marshal(b, m, deterministic)
}
func (m *AppDeleteTypeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppDeleteTypeReq.Merge(m, src)
}
func (m *AppDeleteTypeReq) XXX_Size() int {
	return xxx_messageInfo_AppDeleteTypeReq.Size(m)
}
func (m *AppDeleteTypeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AppDeleteTypeReq.DiscardUnknown(m)
}

var xxx_messageInfo_AppDeleteTypeReq proto.InternalMessageInfo

func (m *AppDeleteTypeReq) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

type AppDeleteTypeResp struct {
	Pong                 int64    `protobuf:"varint,1,opt,name=pong,proto3" json:"pong,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppDeleteTypeResp) Reset()         { *m = AppDeleteTypeResp{} }
func (m *AppDeleteTypeResp) String() string { return proto.CompactTextString(m) }
func (*AppDeleteTypeResp) ProtoMessage()    {}
func (*AppDeleteTypeResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f9056a14b86d47, []int{8}
}

func (m *AppDeleteTypeResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppDeleteTypeResp.Unmarshal(m, b)
}
func (m *AppDeleteTypeResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppDeleteTypeResp.Marshal(b, m, deterministic)
}
func (m *AppDeleteTypeResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppDeleteTypeResp.Merge(m, src)
}
func (m *AppDeleteTypeResp) XXX_Size() int {
	return xxx_messageInfo_AppDeleteTypeResp.Size(m)
}
func (m *AppDeleteTypeResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AppDeleteTypeResp.DiscardUnknown(m)
}

var xxx_messageInfo_AppDeleteTypeResp proto.InternalMessageInfo

func (m *AppDeleteTypeResp) GetPong() int64 {
	if m != nil {
		return m.Pong
	}
	return 0
}

// end
// 相册列表
type AppPageItemReq struct {
	ClassifyId           int64    `protobuf:"varint,1,opt,name=classifyId,proto3" json:"classifyId,omitempty"`
	Page                 int64    `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Size                 int64    `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Total                int64    `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppPageItemReq) Reset()         { *m = AppPageItemReq{} }
func (m *AppPageItemReq) String() string { return proto.CompactTextString(m) }
func (*AppPageItemReq) ProtoMessage()    {}
func (*AppPageItemReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f9056a14b86d47, []int{9}
}

func (m *AppPageItemReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppPageItemReq.Unmarshal(m, b)
}
func (m *AppPageItemReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppPageItemReq.Marshal(b, m, deterministic)
}
func (m *AppPageItemReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppPageItemReq.Merge(m, src)
}
func (m *AppPageItemReq) XXX_Size() int {
	return xxx_messageInfo_AppPageItemReq.Size(m)
}
func (m *AppPageItemReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AppPageItemReq.DiscardUnknown(m)
}

var xxx_messageInfo_AppPageItemReq proto.InternalMessageInfo

func (m *AppPageItemReq) GetClassifyId() int64 {
	if m != nil {
		return m.ClassifyId
	}
	return 0
}

func (m *AppPageItemReq) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *AppPageItemReq) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *AppPageItemReq) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

type AppPageItemData struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreateTime           string   `protobuf:"bytes,2,opt,name=createTime,proto3" json:"createTime,omitempty"`
	UpdateTime           string   `protobuf:"bytes,3,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	Url                  string   `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	Type                 string   `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	ClassifyId           int64    `protobuf:"varint,6,opt,name=classifyId,proto3" json:"classifyId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppPageItemData) Reset()         { *m = AppPageItemData{} }
func (m *AppPageItemData) String() string { return proto.CompactTextString(m) }
func (*AppPageItemData) ProtoMessage()    {}
func (*AppPageItemData) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f9056a14b86d47, []int{10}
}

func (m *AppPageItemData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppPageItemData.Unmarshal(m, b)
}
func (m *AppPageItemData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppPageItemData.Marshal(b, m, deterministic)
}
func (m *AppPageItemData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppPageItemData.Merge(m, src)
}
func (m *AppPageItemData) XXX_Size() int {
	return xxx_messageInfo_AppPageItemData.Size(m)
}
func (m *AppPageItemData) XXX_DiscardUnknown() {
	xxx_messageInfo_AppPageItemData.DiscardUnknown(m)
}

var xxx_messageInfo_AppPageItemData proto.InternalMessageInfo

func (m *AppPageItemData) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AppPageItemData) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *AppPageItemData) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

func (m *AppPageItemData) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *AppPageItemData) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *AppPageItemData) GetClassifyId() int64 {
	if m != nil {
		return m.ClassifyId
	}
	return 0
}

type AppPageItemResp struct {
	List                 []*AppPageItemData `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Total                int64              `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AppPageItemResp) Reset()         { *m = AppPageItemResp{} }
func (m *AppPageItemResp) String() string { return proto.CompactTextString(m) }
func (*AppPageItemResp) ProtoMessage()    {}
func (*AppPageItemResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f9056a14b86d47, []int{11}
}

func (m *AppPageItemResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppPageItemResp.Unmarshal(m, b)
}
func (m *AppPageItemResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppPageItemResp.Marshal(b, m, deterministic)
}
func (m *AppPageItemResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppPageItemResp.Merge(m, src)
}
func (m *AppPageItemResp) XXX_Size() int {
	return xxx_messageInfo_AppPageItemResp.Size(m)
}
func (m *AppPageItemResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AppPageItemResp.DiscardUnknown(m)
}

var xxx_messageInfo_AppPageItemResp proto.InternalMessageInfo

func (m *AppPageItemResp) GetList() []*AppPageItemData {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *AppPageItemResp) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

type AppAddItemReq struct {
	ClassifyId           int64    `protobuf:"varint,1,opt,name=classifyId,proto3" json:"classifyId,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Url                  string   `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppAddItemReq) Reset()         { *m = AppAddItemReq{} }
func (m *AppAddItemReq) String() string { return proto.CompactTextString(m) }
func (*AppAddItemReq) ProtoMessage()    {}
func (*AppAddItemReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f9056a14b86d47, []int{12}
}

func (m *AppAddItemReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppAddItemReq.Unmarshal(m, b)
}
func (m *AppAddItemReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppAddItemReq.Marshal(b, m, deterministic)
}
func (m *AppAddItemReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppAddItemReq.Merge(m, src)
}
func (m *AppAddItemReq) XXX_Size() int {
	return xxx_messageInfo_AppAddItemReq.Size(m)
}
func (m *AppAddItemReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AppAddItemReq.DiscardUnknown(m)
}

var xxx_messageInfo_AppAddItemReq proto.InternalMessageInfo

func (m *AppAddItemReq) GetClassifyId() int64 {
	if m != nil {
		return m.ClassifyId
	}
	return 0
}

func (m *AppAddItemReq) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *AppAddItemReq) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type AppAddItemResp struct {
	Pong                 int64    `protobuf:"varint,1,opt,name=pong,proto3" json:"pong,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppAddItemResp) Reset()         { *m = AppAddItemResp{} }
func (m *AppAddItemResp) String() string { return proto.CompactTextString(m) }
func (*AppAddItemResp) ProtoMessage()    {}
func (*AppAddItemResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f9056a14b86d47, []int{13}
}

func (m *AppAddItemResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppAddItemResp.Unmarshal(m, b)
}
func (m *AppAddItemResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppAddItemResp.Marshal(b, m, deterministic)
}
func (m *AppAddItemResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppAddItemResp.Merge(m, src)
}
func (m *AppAddItemResp) XXX_Size() int {
	return xxx_messageInfo_AppAddItemResp.Size(m)
}
func (m *AppAddItemResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AppAddItemResp.DiscardUnknown(m)
}

var xxx_messageInfo_AppAddItemResp proto.InternalMessageInfo

func (m *AppAddItemResp) GetPong() int64 {
	if m != nil {
		return m.Pong
	}
	return 0
}

func init() {
	proto.RegisterType((*AppListTypeReq)(nil), "app.AppListTypeReq")
	proto.RegisterType((*AppListTypeData)(nil), "app.AppListTypeData")
	proto.RegisterType((*AppListTypeResp)(nil), "app.AppListTypeResp")
	proto.RegisterType((*AppAddTypeReq)(nil), "app.AppAddTypeReq")
	proto.RegisterType((*AppAddTypeResp)(nil), "app.AppAddTypeResp")
	proto.RegisterType((*AppUpdateTypeReq)(nil), "app.AppUpdateTypeReq")
	proto.RegisterType((*AppUpdateTypeResp)(nil), "app.AppUpdateTypeResp")
	proto.RegisterType((*AppDeleteTypeReq)(nil), "app.AppDeleteTypeReq")
	proto.RegisterType((*AppDeleteTypeResp)(nil), "app.AppDeleteTypeResp")
	proto.RegisterType((*AppPageItemReq)(nil), "app.AppPageItemReq")
	proto.RegisterType((*AppPageItemData)(nil), "app.AppPageItemData")
	proto.RegisterType((*AppPageItemResp)(nil), "app.AppPageItemResp")
	proto.RegisterType((*AppAddItemReq)(nil), "app.AppAddItemReq")
	proto.RegisterType((*AppAddItemResp)(nil), "app.AppAddItemResp")
}

func init() { proto.RegisterFile("app.proto", fileDescriptor_e0f9056a14b86d47) }

var fileDescriptor_e0f9056a14b86d47 = []byte{
	// 478 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xcd, 0x6a, 0x1b, 0x31,
	0x10, 0x66, 0x57, 0x4e, 0x20, 0x13, 0x9a, 0xba, 0xb2, 0x5b, 0x16, 0x1f, 0x8a, 0xd9, 0x06, 0xea,
	0x53, 0x0e, 0x29, 0x2d, 0x85, 0x42, 0x61, 0x21, 0x97, 0x40, 0x0f, 0xad, 0x48, 0x1e, 0x40, 0xed,
	0xaa, 0x46, 0xb0, 0xb6, 0x27, 0x96, 0x7c, 0x70, 0x5f, 0xa8, 0x8f, 0xd6, 0xd7, 0x28, 0x1a, 0xad,
	0x56, 0xf2, 0x6e, 0x5c, 0xe8, 0x21, 0xb7, 0xf1, 0xa7, 0xf9, 0xf9, 0xbe, 0x9d, 0x6f, 0x0c, 0x67,
	0x12, 0xf1, 0x0a, 0xb7, 0x1b, 0xbb, 0xe1, 0x4c, 0x22, 0x96, 0x63, 0xb8, 0xa8, 0x10, 0xbf, 0x68,
	0x63, 0xef, 0xf6, 0xa8, 0x84, 0x7a, 0x28, 0x77, 0xf0, 0x3c, 0x41, 0x6e, 0xa4, 0x95, 0xfc, 0x02,
	0x72, 0x5d, 0x17, 0xd9, 0x3c, 0x5b, 0x30, 0x91, 0xeb, 0x9a, 0xbf, 0x06, 0xf8, 0xb1, 0x55, 0xd2,
	0xaa, 0x3b, 0xbd, 0x52, 0x45, 0x3e, 0xcf, 0x16, 0x67, 0x22, 0x41, 0xdc, 0xfb, 0x0e, 0xeb, 0xf0,
	0xce, 0xfc, 0x7b, 0x44, 0x38, 0x87, 0xd1, 0x5a, 0xae, 0x54, 0x31, 0xa2, 0x17, 0x8a, 0xcb, 0x4f,
	0x07, 0x63, 0x85, 0x32, 0xc8, 0x17, 0x30, 0x6a, 0xb4, 0xb1, 0x45, 0x36, 0x67, 0x8b, 0xf3, 0xeb,
	0xe9, 0x95, 0xa3, 0xde, 0xa3, 0x26, 0x28, 0xa3, 0x7c, 0x03, 0xcf, 0x2a, 0xc4, 0xaa, 0xae, 0x5b,
	0x11, 0xdd, 0x84, 0x2c, 0x99, 0x70, 0x49, 0x52, 0xbb, 0x24, 0x83, 0x2e, 0x0b, 0x37, 0xeb, 0x65,
	0xab, 0x8c, 0xe2, 0xf2, 0x03, 0x8c, 0x2b, 0xc4, 0x7b, 0x4f, 0xb6, 0xed, 0xd6, 0xd7, 0x1f, 0xba,
	0xe7, 0x49, 0xf7, 0xb7, 0xf0, 0xa2, 0x57, 0x77, 0x64, 0xc0, 0x25, 0x0d, 0xb8, 0x51, 0x8d, 0x8a,
	0x03, 0xc6, 0xc0, 0x74, 0x6d, 0x48, 0x28, 0x13, 0x2e, 0x6c, 0xdb, 0xa5, 0x59, 0x47, 0xda, 0xad,
	0x49, 0xd5, 0x57, 0xb9, 0x54, 0xb7, 0x56, 0xad, 0x5c, 0x33, 0xb7, 0x9d, 0x46, 0x1a, 0xa3, 0x7f,
	0xee, 0x6f, 0x03, 0xeb, 0x04, 0xa1, 0x2e, 0x72, 0xe9, 0xd9, 0xbb, 0x2e, 0x72, 0x49, 0x1b, 0x31,
	0xfa, 0x97, 0xdf, 0x15, 0x13, 0x14, 0xf3, 0x29, 0x9c, 0xd8, 0x8d, 0x95, 0x0d, 0xad, 0x89, 0x09,
	0xff, 0xa3, 0xfc, 0x9d, 0xd1, 0xa2, 0xc2, 0xc0, 0x27, 0xf1, 0xc7, 0x18, 0xd8, 0x6e, 0xdb, 0xb4,
	0xf6, 0x70, 0xa1, 0xe3, 0x67, 0xf7, 0xa8, 0x8a, 0x13, 0xff, 0xc5, 0x5d, 0xdc, 0xd3, 0x79, 0xda,
	0xd7, 0x59, 0x7e, 0x3b, 0x20, 0xfa, 0x2f, 0x47, 0xa5, 0x62, 0xbc, 0xa3, 0xa2, 0xf8, 0x3c, 0x15,
	0x7f, 0x1f, 0x7c, 0xf6, 0x1f, 0xdf, 0x9a, 0x78, 0xe7, 0x09, 0xef, 0x56, 0x1d, 0xeb, 0xd4, 0x45,
	0x67, 0x76, 0x44, 0x1f, 0xd9, 0xf4, 0xf5, 0x9f, 0x1c, 0x58, 0x85, 0xc8, 0x3f, 0xc2, 0x79, 0x72,
	0x05, 0x7c, 0xd2, 0xbf, 0x0b, 0xa1, 0x1e, 0x66, 0xd3, 0x21, 0x68, 0x90, 0xbf, 0x07, 0x88, 0x17,
	0xc0, 0x79, 0xc8, 0x89, 0x77, 0x33, 0x9b, 0x0c, 0x30, 0x83, 0xfc, 0x33, 0xa9, 0x8e, 0xd6, 0xe6,
	0x2f, 0x43, 0xd6, 0xc1, 0x99, 0xcc, 0x5e, 0x3d, 0x06, 0x77, 0xf5, 0xd1, 0xcb, 0xb1, 0xfe, 0xe0,
	0x0a, 0x62, 0x7d, 0xcf, 0xf6, 0x5e, 0x70, 0x58, 0x52, 0x14, 0x9c, 0x98, 0x7e, 0x36, 0x1d, 0x82,
	0xa9, 0x60, 0x2a, 0x4c, 0x05, 0x87, 0xba, 0xc9, 0x00, 0x33, 0xf8, 0xfd, 0x94, 0xfe, 0x20, 0xdf,
	0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x81, 0xce, 0xdd, 0x2d, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AppClient is the client API for App service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AppClient interface {
	// 相册类型
	AppListType(ctx context.Context, in *AppListTypeReq, opts ...grpc.CallOption) (*AppListTypeResp, error)
	// 相册类型添加
	AppAddType(ctx context.Context, in *AppAddTypeReq, opts ...grpc.CallOption) (*AppAddTypeResp, error)
	// 相册类型更新
	AppUpdateType(ctx context.Context, in *AppUpdateTypeReq, opts ...grpc.CallOption) (*AppUpdateTypeResp, error)
	// 相册类型删除
	AppDeleteType(ctx context.Context, in *AppDeleteTypeReq, opts ...grpc.CallOption) (*AppDeleteTypeResp, error)
	// 相册列表
	AppPageItem(ctx context.Context, in *AppPageItemReq, opts ...grpc.CallOption) (*AppPageItemResp, error)
	// 相册列表插入
	AppAddItem(ctx context.Context, in *AppAddItemReq, opts ...grpc.CallOption) (*AppAddItemResp, error)
}

type appClient struct {
	cc *grpc.ClientConn
}

func NewAppClient(cc *grpc.ClientConn) AppClient {
	return &appClient{cc}
}

func (c *appClient) AppListType(ctx context.Context, in *AppListTypeReq, opts ...grpc.CallOption) (*AppListTypeResp, error) {
	out := new(AppListTypeResp)
	err := c.cc.Invoke(ctx, "/app.App/AppListType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) AppAddType(ctx context.Context, in *AppAddTypeReq, opts ...grpc.CallOption) (*AppAddTypeResp, error) {
	out := new(AppAddTypeResp)
	err := c.cc.Invoke(ctx, "/app.App/AppAddType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) AppUpdateType(ctx context.Context, in *AppUpdateTypeReq, opts ...grpc.CallOption) (*AppUpdateTypeResp, error) {
	out := new(AppUpdateTypeResp)
	err := c.cc.Invoke(ctx, "/app.App/AppUpdateType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) AppDeleteType(ctx context.Context, in *AppDeleteTypeReq, opts ...grpc.CallOption) (*AppDeleteTypeResp, error) {
	out := new(AppDeleteTypeResp)
	err := c.cc.Invoke(ctx, "/app.App/AppDeleteType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) AppPageItem(ctx context.Context, in *AppPageItemReq, opts ...grpc.CallOption) (*AppPageItemResp, error) {
	out := new(AppPageItemResp)
	err := c.cc.Invoke(ctx, "/app.App/AppPageItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) AppAddItem(ctx context.Context, in *AppAddItemReq, opts ...grpc.CallOption) (*AppAddItemResp, error) {
	out := new(AppAddItemResp)
	err := c.cc.Invoke(ctx, "/app.App/AppAddItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppServer is the server API for App service.
type AppServer interface {
	// 相册类型
	AppListType(context.Context, *AppListTypeReq) (*AppListTypeResp, error)
	// 相册类型添加
	AppAddType(context.Context, *AppAddTypeReq) (*AppAddTypeResp, error)
	// 相册类型更新
	AppUpdateType(context.Context, *AppUpdateTypeReq) (*AppUpdateTypeResp, error)
	// 相册类型删除
	AppDeleteType(context.Context, *AppDeleteTypeReq) (*AppDeleteTypeResp, error)
	// 相册列表
	AppPageItem(context.Context, *AppPageItemReq) (*AppPageItemResp, error)
	// 相册列表插入
	AppAddItem(context.Context, *AppAddItemReq) (*AppAddItemResp, error)
}

// UnimplementedAppServer can be embedded to have forward compatible implementations.
type UnimplementedAppServer struct {
}

func (*UnimplementedAppServer) AppListType(ctx context.Context, req *AppListTypeReq) (*AppListTypeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppListType not implemented")
}
func (*UnimplementedAppServer) AppAddType(ctx context.Context, req *AppAddTypeReq) (*AppAddTypeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppAddType not implemented")
}
func (*UnimplementedAppServer) AppUpdateType(ctx context.Context, req *AppUpdateTypeReq) (*AppUpdateTypeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppUpdateType not implemented")
}
func (*UnimplementedAppServer) AppDeleteType(ctx context.Context, req *AppDeleteTypeReq) (*AppDeleteTypeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppDeleteType not implemented")
}
func (*UnimplementedAppServer) AppPageItem(ctx context.Context, req *AppPageItemReq) (*AppPageItemResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppPageItem not implemented")
}
func (*UnimplementedAppServer) AppAddItem(ctx context.Context, req *AppAddItemReq) (*AppAddItemResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppAddItem not implemented")
}

func RegisterAppServer(s *grpc.Server, srv AppServer) {
	s.RegisterService(&_App_serviceDesc, srv)
}

func _App_AppListType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppListTypeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).AppListType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.App/AppListType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).AppListType(ctx, req.(*AppListTypeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_AppAddType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppAddTypeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).AppAddType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.App/AppAddType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).AppAddType(ctx, req.(*AppAddTypeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_AppUpdateType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppUpdateTypeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).AppUpdateType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.App/AppUpdateType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).AppUpdateType(ctx, req.(*AppUpdateTypeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_AppDeleteType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppDeleteTypeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).AppDeleteType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.App/AppDeleteType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).AppDeleteType(ctx, req.(*AppDeleteTypeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_AppPageItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppPageItemReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).AppPageItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.App/AppPageItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).AppPageItem(ctx, req.(*AppPageItemReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_AppAddItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppAddItemReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).AppAddItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.App/AppAddItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).AppAddItem(ctx, req.(*AppAddItemReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _App_serviceDesc = grpc.ServiceDesc{
	ServiceName: "app.App",
	HandlerType: (*AppServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AppListType",
			Handler:    _App_AppListType_Handler,
		},
		{
			MethodName: "AppAddType",
			Handler:    _App_AppAddType_Handler,
		},
		{
			MethodName: "AppUpdateType",
			Handler:    _App_AppUpdateType_Handler,
		},
		{
			MethodName: "AppDeleteType",
			Handler:    _App_AppDeleteType_Handler,
		},
		{
			MethodName: "AppPageItem",
			Handler:    _App_AppPageItem_Handler,
		},
		{
			MethodName: "AppAddItem",
			Handler:    _App_AppAddItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app.proto",
}
