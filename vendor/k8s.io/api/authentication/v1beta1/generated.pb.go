/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: k8s.io/kubernetes/vendor/k8s.io/api/authentication/v1beta1/generated.proto

package v1beta1

import (
	fmt "fmt"

	io "io"

	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func (m *ExtraValue) Reset()      { *m = ExtraValue{} }
func (*ExtraValue) ProtoMessage() {}
func (*ExtraValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_77c9b20d3ad27844, []int{0}
}
func (m *ExtraValue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExtraValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ExtraValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtraValue.Merge(m, src)
}
func (m *ExtraValue) XXX_Size() int {
	return m.Size()
}
func (m *ExtraValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtraValue.DiscardUnknown(m)
}

var xxx_messageInfo_ExtraValue proto.InternalMessageInfo

func (m *TokenReview) Reset()      { *m = TokenReview{} }
func (*TokenReview) ProtoMessage() {}
func (*TokenReview) Descriptor() ([]byte, []int) {
	return fileDescriptor_77c9b20d3ad27844, []int{1}
}
func (m *TokenReview) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenReview) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *TokenReview) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenReview.Merge(m, src)
}
func (m *TokenReview) XXX_Size() int {
	return m.Size()
}
func (m *TokenReview) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenReview.DiscardUnknown(m)
}

var xxx_messageInfo_TokenReview proto.InternalMessageInfo

func (m *TokenReviewSpec) Reset()      { *m = TokenReviewSpec{} }
func (*TokenReviewSpec) ProtoMessage() {}
func (*TokenReviewSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_77c9b20d3ad27844, []int{2}
}
func (m *TokenReviewSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenReviewSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *TokenReviewSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenReviewSpec.Merge(m, src)
}
func (m *TokenReviewSpec) XXX_Size() int {
	return m.Size()
}
func (m *TokenReviewSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenReviewSpec.DiscardUnknown(m)
}

var xxx_messageInfo_TokenReviewSpec proto.InternalMessageInfo

func (m *TokenReviewStatus) Reset()      { *m = TokenReviewStatus{} }
func (*TokenReviewStatus) ProtoMessage() {}
func (*TokenReviewStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_77c9b20d3ad27844, []int{3}
}
func (m *TokenReviewStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenReviewStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *TokenReviewStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenReviewStatus.Merge(m, src)
}
func (m *TokenReviewStatus) XXX_Size() int {
	return m.Size()
}
func (m *TokenReviewStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenReviewStatus.DiscardUnknown(m)
}

var xxx_messageInfo_TokenReviewStatus proto.InternalMessageInfo

func (m *UserInfo) Reset()      { *m = UserInfo{} }
func (*UserInfo) ProtoMessage() {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_77c9b20d3ad27844, []int{4}
}
func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return m.Size()
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ExtraValue)(nil), "k8s.io.api.authentication.v1beta1.ExtraValue")
	proto.RegisterType((*TokenReview)(nil), "k8s.io.api.authentication.v1beta1.TokenReview")
	proto.RegisterType((*TokenReviewSpec)(nil), "k8s.io.api.authentication.v1beta1.TokenReviewSpec")
	proto.RegisterType((*TokenReviewStatus)(nil), "k8s.io.api.authentication.v1beta1.TokenReviewStatus")
	proto.RegisterType((*UserInfo)(nil), "k8s.io.api.authentication.v1beta1.UserInfo")
	proto.RegisterMapType((map[string]ExtraValue)(nil), "k8s.io.api.authentication.v1beta1.UserInfo.ExtraEntry")
}

func init() {
	proto.RegisterFile("k8s.io/kubernetes/vendor/k8s.io/api/authentication/v1beta1/generated.proto", fileDescriptor_77c9b20d3ad27844)
}

var fileDescriptor_77c9b20d3ad27844 = []byte{
	// 663 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xb6, 0xf3, 0x53, 0x92, 0x0d, 0x81, 0xb2, 0x12, 0x52, 0x14, 0x09, 0xa7, 0x84, 0x4b, 0xa5,
	0xd2, 0x35, 0xad, 0xaa, 0x52, 0x95, 0x53, 0x0d, 0x15, 0x2a, 0x52, 0x85, 0xb4, 0xb4, 0x1c, 0x80,
	0x03, 0x1b, 0x67, 0xea, 0x98, 0xe0, 0x1f, 0xad, 0xd7, 0x81, 0xde, 0xfa, 0x08, 0x1c, 0x39, 0x22,
	0xf1, 0x24, 0xdc, 0x7a, 0xec, 0xb1, 0x07, 0x14, 0x51, 0xf3, 0x04, 0xbc, 0x01, 0xda, 0xf5, 0xb6,
	0x4e, 0x1b, 0x41, 0xdb, 0x9b, 0xf7, 0x9b, 0xf9, 0xbe, 0x99, 0xf9, 0xc6, 0x83, 0x5e, 0x0c, 0xd7,
	0x12, 0xe2, 0x47, 0xf6, 0x30, 0xed, 0x01, 0x0f, 0x41, 0x40, 0x62, 0x8f, 0x20, 0xec, 0x47, 0xdc,
	0xd6, 0x01, 0x16, 0xfb, 0x36, 0x4b, 0xc5, 0x00, 0x42, 0xe1, 0xbb, 0x4c, 0xf8, 0x51, 0x68, 0x8f,
	0x96, 0x7a, 0x20, 0xd8, 0x92, 0xed, 0x41, 0x08, 0x9c, 0x09, 0xe8, 0x93, 0x98, 0x47, 0x22, 0xc2,
	0xf7, 0x73, 0x0a, 0x61, 0xb1, 0x4f, 0xce, 0x53, 0x88, 0xa6, 0xb4, 0x17, 0x3d, 0x5f, 0x0c, 0xd2,
	0x1e, 0x71, 0xa3, 0xc0, 0xf6, 0x22, 0x2f, 0xb2, 0x15, 0xb3, 0x97, 0xee, 0xa9, 0x97, 0x7a, 0xa8,
	0xaf, 0x5c, 0xb1, 0xbd, 0x52, 0x34, 0x11, 0x30, 0x77, 0xe0, 0x87, 0xc0, 0xf7, 0xed, 0x78, 0xe8,
	0x49, 0x20, 0xb1, 0x03, 0x10, 0xcc, 0x1e, 0x4d, 0xf5, 0xd1, 0xb6, 0xff, 0xc5, 0xe2, 0x69, 0x28,
	0xfc, 0x00, 0xa6, 0x08, 0xab, 0x97, 0x11, 0x12, 0x77, 0x00, 0x01, 0xbb, 0xc8, 0xeb, 0x3e, 0x46,
	0x68, 0xf3, 0xb3, 0xe0, 0xec, 0x35, 0xfb, 0x98, 0x02, 0xee, 0xa0, 0xaa, 0x2f, 0x20, 0x48, 0x5a,
	0xe6, 0x5c, 0x79, 0xbe, 0xee, 0xd4, 0xb3, 0x71, 0xa7, 0xba, 0x25, 0x01, 0x9a, 0xe3, 0xeb, 0xb5,
	0xaf, 0xdf, 0x3a, 0xc6, 0xc1, 0xcf, 0x39, 0xa3, 0xfb, 0xbd, 0x84, 0x1a, 0x3b, 0xd1, 0x10, 0x42,
	0x0a, 0x23, 0x1f, 0x3e, 0xe1, 0xf7, 0xa8, 0x26, 0x87, 0xe9, 0x33, 0xc1, 0x5a, 0xe6, 0x9c, 0x39,
	0xdf, 0x58, 0x7e, 0x44, 0x0a, 0x33, 0xcf, 0x7a, 0x22, 0xf1, 0xd0, 0x93, 0x40, 0x42, 0x64, 0x36,
	0x19, 0x2d, 0x91, 0x97, 0xbd, 0x0f, 0xe0, 0x8a, 0x6d, 0x10, 0xcc, 0xc1, 0x87, 0xe3, 0x8e, 0x91,
	0x8d, 0x3b, 0xa8, 0xc0, 0xe8, 0x99, 0x2a, 0xde, 0x41, 0x95, 0x24, 0x06, 0xb7, 0x55, 0x52, 0xea,
	0xcb, 0xe4, 0xd2, 0x55, 0x91, 0x89, 0xfe, 0x5e, 0xc5, 0xe0, 0x3a, 0x37, 0xb5, 0x7e, 0x45, 0xbe,
	0xa8, 0x52, 0xc3, 0xef, 0xd0, 0x4c, 0x22, 0x98, 0x48, 0x93, 0x56, 0x59, 0xe9, 0xae, 0x5c, 0x53,
	0x57, 0x71, 0x9d, 0x5b, 0x5a, 0x79, 0x26, 0x7f, 0x53, 0xad, 0xd9, 0x75, 0xd1, 0xed, 0x0b, 0x4d,
	0xe0, 0x07, 0xa8, 0x2a, 0x24, 0xa4, 0x5c, 0xaa, 0x3b, 0x4d, 0xcd, 0xac, 0xe6, 0x79, 0x79, 0x0c,
	0x2f, 0xa0, 0x3a, 0x4b, 0xfb, 0x3e, 0x84, 0x2e, 0x24, 0xad, 0x92, 0x5a, 0x46, 0x33, 0x1b, 0x77,
	0xea, 0x1b, 0xa7, 0x20, 0x2d, 0xe2, 0xdd, 0x3f, 0x26, 0xba, 0x33, 0xd5, 0x12, 0x7e, 0x82, 0x9a,
	0x13, 0xed, 0x43, 0x5f, 0xd5, 0xab, 0x39, 0x77, 0x75, 0xbd, 0xe6, 0xc6, 0x64, 0x90, 0x9e, 0xcf,
	0xc5, 0xdb, 0xa8, 0x92, 0x26, 0xc0, 0xb5, 0xd7, 0x0b, 0x57, 0xf0, 0x64, 0x37, 0x01, 0xbe, 0x15,
	0xee, 0x45, 0x85, 0xc9, 0x12, 0xa1, 0x4a, 0xe6, 0xfc, 0x38, 0x95, 0xff, 0x8f, 0x23, 0x0d, 0x02,
	0xce, 0x23, 0xae, 0x16, 0x32, 0x61, 0xd0, 0xa6, 0x04, 0x69, 0x1e, 0xeb, 0xfe, 0x28, 0xa1, 0xda,
	0x69, 0x49, 0xfc, 0x10, 0xd5, 0x64, 0x99, 0x90, 0x05, 0xa0, 0x5d, 0x9d, 0xd5, 0x24, 0x95, 0x23,
	0x71, 0x7a, 0x96, 0x81, 0xef, 0xa1, 0x72, 0xea, 0xf7, 0xd5, 0x68, 0x75, 0xa7, 0xa1, 0x13, 0xcb,
	0xbb, 0x5b, 0xcf, 0xa8, 0xc4, 0x71, 0x17, 0xcd, 0x78, 0x3c, 0x4a, 0x63, 0xf9, 0x43, 0xc8, 0x46,
	0x91, 0x5c, 0xeb, 0x73, 0x85, 0x50, 0x1d, 0xc1, 0x6f, 0x51, 0x15, 0xe4, 0xd5, 0xa8, 0x59, 0x1a,
	0xcb, 0xab, 0xd7, 0xf0, 0x87, 0xa8, 0x73, 0xdb, 0x0c, 0x05, 0xdf, 0x9f, 0x18, 0x4d, 0x62, 0x34,
	0xd7, 0x6c, 0x7b, 0xfa, 0x24, 0x55, 0x0e, 0x9e, 0x45, 0xe5, 0x21, 0xec, 0xe7, 0x63, 0x51, 0xf9,
	0x89, 0x9f, 0xa2, 0xea, 0x48, 0x5e, 0xab, 0x5e, 0xce, 0xe2, 0x15, 0x8a, 0x17, 0x27, 0x4e, 0x73,
	0xee, 0x7a, 0x69, 0xcd, 0x74, 0x16, 0x0f, 0x4f, 0x2c, 0xe3, 0xe8, 0xc4, 0x32, 0x8e, 0x4f, 0x2c,
	0xe3, 0x20, 0xb3, 0xcc, 0xc3, 0xcc, 0x32, 0x8f, 0x32, 0xcb, 0x3c, 0xce, 0x2c, 0xf3, 0x57, 0x66,
	0x99, 0x5f, 0x7e, 0x5b, 0xc6, 0x9b, 0x1b, 0x5a, 0xe4, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x66,
	0xbb, 0x89, 0x53, 0x68, 0x05, 0x00, 0x00,
}

func (m ExtraValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m ExtraValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m ExtraValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m) > 0 {
		for iNdEx := len(m) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m[iNdEx])
			copy(dAtA[i:], m[iNdEx])
			i = encodeVarintGenerated(dAtA, i, uint64(len(m[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *TokenReview) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenReview) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenReview) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Status.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.Spec.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.ObjectMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *TokenReviewSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenReviewSpec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenReviewSpec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Audiences) > 0 {
		for iNdEx := len(m.Audiences) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Audiences[iNdEx])
			copy(dAtA[i:], m.Audiences[iNdEx])
			i = encodeVarintGenerated(dAtA, i, uint64(len(m.Audiences[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	i -= len(m.Token)
	copy(dAtA[i:], m.Token)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Token)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *TokenReviewStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenReviewStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenReviewStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Audiences) > 0 {
		for iNdEx := len(m.Audiences) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Audiences[iNdEx])
			copy(dAtA[i:], m.Audiences[iNdEx])
			i = encodeVarintGenerated(dAtA, i, uint64(len(m.Audiences[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	i -= len(m.Error)
	copy(dAtA[i:], m.Error)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Error)))
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.User.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	i--
	if m.Authenticated {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}

func (m *UserInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Extra) > 0 {
		keysForExtra := make([]string, 0, len(m.Extra))
		for k := range m.Extra {
			keysForExtra = append(keysForExtra, string(k))
		}
		github_com_gogo_protobuf_sortkeys.Strings(keysForExtra)
		for iNdEx := len(keysForExtra) - 1; iNdEx >= 0; iNdEx-- {
			v := m.Extra[string(keysForExtra[iNdEx])]
			baseI := i
			{
				size, err := (&v).MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenerated(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
			i -= len(keysForExtra[iNdEx])
			copy(dAtA[i:], keysForExtra[iNdEx])
			i = encodeVarintGenerated(dAtA, i, uint64(len(keysForExtra[iNdEx])))
			i--
			dAtA[i] = 0xa
			i = encodeVarintGenerated(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Groups) > 0 {
		for iNdEx := len(m.Groups) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Groups[iNdEx])
			copy(dAtA[i:], m.Groups[iNdEx])
			i = encodeVarintGenerated(dAtA, i, uint64(len(m.Groups[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	i -= len(m.UID)
	copy(dAtA[i:], m.UID)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.UID)))
	i--
	dAtA[i] = 0x12
	i -= len(m.Username)
	copy(dAtA[i:], m.Username)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Username)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenerated(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m ExtraValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m) > 0 {
		for _, s := range m {
			l = len(s)
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *TokenReview) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Spec.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Status.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *TokenReviewSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Token)
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Audiences) > 0 {
		for _, s := range m.Audiences {
			l = len(s)
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *TokenReviewStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 2
	l = m.User.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Error)
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Audiences) > 0 {
		for _, s := range m.Audiences {
			l = len(s)
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *UserInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Username)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.UID)
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Groups) > 0 {
		for _, s := range m.Groups {
			l = len(s)
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if len(m.Extra) > 0 {
		for k, v := range m.Extra {
			_ = k
			_ = v
			l = v.Size()
			mapEntrySize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + l + sovGenerated(uint64(l))
			n += mapEntrySize + 1 + sovGenerated(uint64(mapEntrySize))
		}
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *TokenReview) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TokenReview{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ObjectMeta), "ObjectMeta", "v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Spec:` + strings.Replace(strings.Replace(this.Spec.String(), "TokenReviewSpec", "TokenReviewSpec", 1), `&`, ``, 1) + `,`,
		`Status:` + strings.Replace(strings.Replace(this.Status.String(), "TokenReviewStatus", "TokenReviewStatus", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *TokenReviewSpec) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TokenReviewSpec{`,
		`Token:` + fmt.Sprintf("%v", this.Token) + `,`,
		`Audiences:` + fmt.Sprintf("%v", this.Audiences) + `,`,
		`}`,
	}, "")
	return s
}
func (this *TokenReviewStatus) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TokenReviewStatus{`,
		`Authenticated:` + fmt.Sprintf("%v", this.Authenticated) + `,`,
		`User:` + strings.Replace(strings.Replace(this.User.String(), "UserInfo", "UserInfo", 1), `&`, ``, 1) + `,`,
		`Error:` + fmt.Sprintf("%v", this.Error) + `,`,
		`Audiences:` + fmt.Sprintf("%v", this.Audiences) + `,`,
		`}`,
	}, "")
	return s
}
func (this *UserInfo) String() string {
	if this == nil {
		return "nil"
	}
	keysForExtra := make([]string, 0, len(this.Extra))
	for k := range this.Extra {
		keysForExtra = append(keysForExtra, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForExtra)
	mapStringForExtra := "map[string]ExtraValue{"
	for _, k := range keysForExtra {
		mapStringForExtra += fmt.Sprintf("%v: %v,", k, this.Extra[k])
	}
	mapStringForExtra += "}"
	s := strings.Join([]string{`&UserInfo{`,
		`Username:` + fmt.Sprintf("%v", this.Username) + `,`,
		`UID:` + fmt.Sprintf("%v", this.UID) + `,`,
		`Groups:` + fmt.Sprintf("%v", this.Groups) + `,`,
		`Extra:` + mapStringForExtra + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ExtraValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ExtraValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExtraValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			*m = append(*m, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TokenReview) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TokenReview: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenReview: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TokenReviewSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TokenReviewSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenReviewSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Audiences", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Audiences = append(m.Audiences, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TokenReviewStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TokenReviewStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenReviewStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authenticated", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Authenticated = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.User.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Audiences", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Audiences = append(m.Audiences, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UserInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UserInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Username = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Groups", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Groups = append(m.Groups, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Extra", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Extra == nil {
				m.Extra = make(map[string]ExtraValue)
			}
			var mapkey string
			mapvalue := &ExtraValue{}
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenerated
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthGenerated
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthGenerated
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenerated
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthGenerated
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthGenerated
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &ExtraValue{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipGenerated(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthGenerated
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Extra[mapkey] = *mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenerated
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenerated
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenerated        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenerated = fmt.Errorf("proto: unexpected end of group")
)
