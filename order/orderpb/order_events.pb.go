// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        (unknown)
// source: orderpb/order_events.proto

package orderpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OrderCreated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderID      string               `protobuf:"bytes,1,opt,name=OrderID,proto3" json:"OrderID,omitempty"`
	ConsumerID   string               `protobuf:"bytes,2,opt,name=ConsumerID,proto3" json:"ConsumerID,omitempty"`
	RestaurantID string               `protobuf:"bytes,3,opt,name=RestaurantID,proto3" json:"RestaurantID,omitempty"`
	OrderTotal   int64                `protobuf:"varint,4,opt,name=OrderTotal,proto3" json:"OrderTotal,omitempty"`
	Status       string               `protobuf:"bytes,5,opt,name=Status,proto3" json:"Status,omitempty"`
	Items        []*OrderCreated_Item `protobuf:"bytes,6,rep,name=items,proto3" json:"items,omitempty"`
	Address      *Address             `protobuf:"bytes,7,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *OrderCreated) Reset() {
	*x = OrderCreated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderpb_order_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderCreated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderCreated) ProtoMessage() {}

func (x *OrderCreated) ProtoReflect() protoreflect.Message {
	mi := &file_orderpb_order_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderCreated.ProtoReflect.Descriptor instead.
func (*OrderCreated) Descriptor() ([]byte, []int) {
	return file_orderpb_order_events_proto_rawDescGZIP(), []int{0}
}

func (x *OrderCreated) GetOrderID() string {
	if x != nil {
		return x.OrderID
	}
	return ""
}

func (x *OrderCreated) GetConsumerID() string {
	if x != nil {
		return x.ConsumerID
	}
	return ""
}

func (x *OrderCreated) GetRestaurantID() string {
	if x != nil {
		return x.RestaurantID
	}
	return ""
}

func (x *OrderCreated) GetOrderTotal() int64 {
	if x != nil {
		return x.OrderTotal
	}
	return 0
}

func (x *OrderCreated) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *OrderCreated) GetItems() []*OrderCreated_Item {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *OrderCreated) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

// Command
type ApproveOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderID  string `protobuf:"bytes,1,opt,name=OrderID,proto3" json:"OrderID,omitempty"`
	TicketID string `protobuf:"bytes,2,opt,name=TicketID,proto3" json:"TicketID,omitempty"`
}

func (x *ApproveOrder) Reset() {
	*x = ApproveOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderpb_order_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApproveOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApproveOrder) ProtoMessage() {}

func (x *ApproveOrder) ProtoReflect() protoreflect.Message {
	mi := &file_orderpb_order_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApproveOrder.ProtoReflect.Descriptor instead.
func (*ApproveOrder) Descriptor() ([]byte, []int) {
	return file_orderpb_order_events_proto_rawDescGZIP(), []int{1}
}

func (x *ApproveOrder) GetOrderID() string {
	if x != nil {
		return x.OrderID
	}
	return ""
}

func (x *ApproveOrder) GetTicketID() string {
	if x != nil {
		return x.TicketID
	}
	return ""
}

type RejectOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderID string `protobuf:"bytes,1,opt,name=OrderID,proto3" json:"OrderID,omitempty"`
}

func (x *RejectOrder) Reset() {
	*x = RejectOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderpb_order_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RejectOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RejectOrder) ProtoMessage() {}

func (x *RejectOrder) ProtoReflect() protoreflect.Message {
	mi := &file_orderpb_order_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RejectOrder.ProtoReflect.Descriptor instead.
func (*RejectOrder) Descriptor() ([]byte, []int) {
	return file_orderpb_order_events_proto_rawDescGZIP(), []int{2}
}

func (x *RejectOrder) GetOrderID() string {
	if x != nil {
		return x.OrderID
	}
	return ""
}

type OrderCreated_Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MenuItemId string `protobuf:"bytes,1,opt,name=menu_item_id,json=menuItemId,proto3" json:"menu_item_id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price      int64  `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
	Quantity   int64  `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *OrderCreated_Item) Reset() {
	*x = OrderCreated_Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderpb_order_events_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderCreated_Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderCreated_Item) ProtoMessage() {}

func (x *OrderCreated_Item) ProtoReflect() protoreflect.Message {
	mi := &file_orderpb_order_events_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderCreated_Item.ProtoReflect.Descriptor instead.
func (*OrderCreated_Item) Descriptor() ([]byte, []int) {
	return file_orderpb_order_events_proto_rawDescGZIP(), []int{0, 0}
}

func (x *OrderCreated_Item) GetMenuItemId() string {
	if x != nil {
		return x.MenuItemId
	}
	return ""
}

func (x *OrderCreated_Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OrderCreated_Item) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *OrderCreated_Item) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

var File_orderpb_order_events_proto protoreflect.FileDescriptor

var file_orderpb_order_events_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x70, 0x62, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x70, 0x62, 0x1a, 0x17, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x70, 0x62, 0x2f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf2,
	0x02, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x73,
	0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1e, 0x0a,
	0x0a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x16, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x30, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x2a, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x1a, 0x6e, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x20, 0x0a, 0x0c, 0x6d,
	0x65, 0x6e, 0x75, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6d, 0x65, 0x6e, 0x75, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x22, 0x44, 0x0a, 0x0c, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a,
	0x08, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x44, 0x22, 0x27, 0x0a, 0x0b, 0x52, 0x65, 0x6a,
	0x65, 0x63, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x44, 0x42, 0x93, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x70, 0x62, 0x42, 0x10, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x7a, 0x61, 0x41, 0x6d, 0x69, 0x72, 0x69, 0x31, 0x32, 0x33, 0x2f,
	0x66, 0x74, 0x67, 0x6f, 0x67, 0x6f, 0x56, 0x33, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x70, 0x62, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x70, 0x62, 0xa2, 0x02,
	0x03, 0x4f, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x70, 0x62, 0xca, 0x02,
	0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x70, 0x62, 0xe2, 0x02, 0x13, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x70, 0x62, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orderpb_order_events_proto_rawDescOnce sync.Once
	file_orderpb_order_events_proto_rawDescData = file_orderpb_order_events_proto_rawDesc
)

func file_orderpb_order_events_proto_rawDescGZIP() []byte {
	file_orderpb_order_events_proto_rawDescOnce.Do(func() {
		file_orderpb_order_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_orderpb_order_events_proto_rawDescData)
	})
	return file_orderpb_order_events_proto_rawDescData
}

var file_orderpb_order_events_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_orderpb_order_events_proto_goTypes = []interface{}{
	(*OrderCreated)(nil),      // 0: orderpb.OrderCreated
	(*ApproveOrder)(nil),      // 1: orderpb.ApproveOrder
	(*RejectOrder)(nil),       // 2: orderpb.RejectOrder
	(*OrderCreated_Item)(nil), // 3: orderpb.OrderCreated.Item
	(*Address)(nil),           // 4: orderpb.Address
}
var file_orderpb_order_events_proto_depIdxs = []int32{
	3, // 0: orderpb.OrderCreated.items:type_name -> orderpb.OrderCreated.Item
	4, // 1: orderpb.OrderCreated.address:type_name -> orderpb.Address
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_orderpb_order_events_proto_init() }
func file_orderpb_order_events_proto_init() {
	if File_orderpb_order_events_proto != nil {
		return
	}
	file_orderpb_order_api_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_orderpb_order_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderCreated); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_orderpb_order_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApproveOrder); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_orderpb_order_events_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RejectOrder); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_orderpb_order_events_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderCreated_Item); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_orderpb_order_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_orderpb_order_events_proto_goTypes,
		DependencyIndexes: file_orderpb_order_events_proto_depIdxs,
		MessageInfos:      file_orderpb_order_events_proto_msgTypes,
	}.Build()
	File_orderpb_order_events_proto = out.File
	file_orderpb_order_events_proto_rawDesc = nil
	file_orderpb_order_events_proto_goTypes = nil
	file_orderpb_order_events_proto_depIdxs = nil
}
