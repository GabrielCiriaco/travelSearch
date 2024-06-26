// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: contact.proto

package contact

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

type TravelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Origem            string `protobuf:"bytes,1,opt,name=origem,proto3" json:"origem,omitempty"`
	Destino           string `protobuf:"bytes,2,opt,name=destino,proto3" json:"destino,omitempty"`
	Cidade            string `protobuf:"bytes,3,opt,name=cidade,proto3" json:"cidade,omitempty"`
	DataIda           string `protobuf:"bytes,4,opt,name=data_ida,json=dataIda,proto3" json:"data_ida,omitempty"`
	DataVolta         string `protobuf:"bytes,5,opt,name=data_volta,json=dataVolta,proto3" json:"data_volta,omitempty"`
	QuantidadeAdultos int32  `protobuf:"varint,6,opt,name=quantidade_adultos,json=quantidadeAdultos,proto3" json:"quantidade_adultos,omitempty"`
}

func (x *TravelRequest) Reset() {
	*x = TravelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contact_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TravelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TravelRequest) ProtoMessage() {}

func (x *TravelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contact_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TravelRequest.ProtoReflect.Descriptor instead.
func (*TravelRequest) Descriptor() ([]byte, []int) {
	return file_contact_proto_rawDescGZIP(), []int{0}
}

func (x *TravelRequest) GetOrigem() string {
	if x != nil {
		return x.Origem
	}
	return ""
}

func (x *TravelRequest) GetDestino() string {
	if x != nil {
		return x.Destino
	}
	return ""
}

func (x *TravelRequest) GetCidade() string {
	if x != nil {
		return x.Cidade
	}
	return ""
}

func (x *TravelRequest) GetDataIda() string {
	if x != nil {
		return x.DataIda
	}
	return ""
}

func (x *TravelRequest) GetDataVolta() string {
	if x != nil {
		return x.DataVolta
	}
	return ""
}

func (x *TravelRequest) GetQuantidadeAdultos() int32 {
	if x != nil {
		return x.QuantidadeAdultos
	}
	return 0
}

type TravelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hoteis []*Hotel `protobuf:"bytes,1,rep,name=hoteis,proto3" json:"hoteis,omitempty"`
	Voos   []*Voo   `protobuf:"bytes,2,rep,name=voos,proto3" json:"voos,omitempty"`
}

func (x *TravelResponse) Reset() {
	*x = TravelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contact_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TravelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TravelResponse) ProtoMessage() {}

func (x *TravelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contact_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TravelResponse.ProtoReflect.Descriptor instead.
func (*TravelResponse) Descriptor() ([]byte, []int) {
	return file_contact_proto_rawDescGZIP(), []int{1}
}

func (x *TravelResponse) GetHoteis() []*Hotel {
	if x != nil {
		return x.Hoteis
	}
	return nil
}

func (x *TravelResponse) GetVoos() []*Voo {
	if x != nil {
		return x.Voos
	}
	return nil
}

type Hotel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NomeHotel         string  `protobuf:"bytes,1,opt,name=nome_hotel,json=nomeHotel,proto3" json:"nome_hotel,omitempty"`
	Cidade            string  `protobuf:"bytes,2,opt,name=cidade,proto3" json:"cidade,omitempty"`
	Endereco          string  `protobuf:"bytes,3,opt,name=endereco,proto3" json:"endereco,omitempty"`
	PrecoTotal        float32 `protobuf:"fixed32,4,opt,name=preco_total,json=precoTotal,proto3" json:"preco_total,omitempty"`
	DescricaoQuarto   string  `protobuf:"bytes,5,opt,name=descricao_quarto,json=descricaoQuarto,proto3" json:"descricao_quarto,omitempty"`
	CheckinDate       string  `protobuf:"bytes,6,opt,name=checkin_date,json=checkinDate,proto3" json:"checkin_date,omitempty"`
	CheckoutDate      string  `protobuf:"bytes,7,opt,name=checkout_date,json=checkoutDate,proto3" json:"checkout_date,omitempty"`
	QuantidadeAdultos int32   `protobuf:"varint,8,opt,name=quantidade_adultos,json=quantidadeAdultos,proto3" json:"quantidade_adultos,omitempty"`
	CategoriaQuarto   string  `protobuf:"bytes,9,opt,name=categoria_quarto,json=categoriaQuarto,proto3" json:"categoria_quarto,omitempty"`
}

func (x *Hotel) Reset() {
	*x = Hotel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contact_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hotel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hotel) ProtoMessage() {}

func (x *Hotel) ProtoReflect() protoreflect.Message {
	mi := &file_contact_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hotel.ProtoReflect.Descriptor instead.
func (*Hotel) Descriptor() ([]byte, []int) {
	return file_contact_proto_rawDescGZIP(), []int{2}
}

func (x *Hotel) GetNomeHotel() string {
	if x != nil {
		return x.NomeHotel
	}
	return ""
}

func (x *Hotel) GetCidade() string {
	if x != nil {
		return x.Cidade
	}
	return ""
}

func (x *Hotel) GetEndereco() string {
	if x != nil {
		return x.Endereco
	}
	return ""
}

func (x *Hotel) GetPrecoTotal() float32 {
	if x != nil {
		return x.PrecoTotal
	}
	return 0
}

func (x *Hotel) GetDescricaoQuarto() string {
	if x != nil {
		return x.DescricaoQuarto
	}
	return ""
}

func (x *Hotel) GetCheckinDate() string {
	if x != nil {
		return x.CheckinDate
	}
	return ""
}

func (x *Hotel) GetCheckoutDate() string {
	if x != nil {
		return x.CheckoutDate
	}
	return ""
}

func (x *Hotel) GetQuantidadeAdultos() int32 {
	if x != nil {
		return x.QuantidadeAdultos
	}
	return 0
}

func (x *Hotel) GetCategoriaQuarto() string {
	if x != nil {
		return x.CategoriaQuarto
	}
	return ""
}

type Voo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HorarioPartida string `protobuf:"bytes,1,opt,name=horario_partida,json=horarioPartida,proto3" json:"horario_partida,omitempty"`
	HorarioChegada string `protobuf:"bytes,2,opt,name=horario_chegada,json=horarioChegada,proto3" json:"horario_chegada,omitempty"`
	Preco          string `protobuf:"bytes,3,opt,name=preco,proto3" json:"preco,omitempty"`
	Detalhes       string `protobuf:"bytes,4,opt,name=detalhes,proto3" json:"detalhes,omitempty"`
}

func (x *Voo) Reset() {
	*x = Voo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contact_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Voo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Voo) ProtoMessage() {}

func (x *Voo) ProtoReflect() protoreflect.Message {
	mi := &file_contact_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Voo.ProtoReflect.Descriptor instead.
func (*Voo) Descriptor() ([]byte, []int) {
	return file_contact_proto_rawDescGZIP(), []int{3}
}

func (x *Voo) GetHorarioPartida() string {
	if x != nil {
		return x.HorarioPartida
	}
	return ""
}

func (x *Voo) GetHorarioChegada() string {
	if x != nil {
		return x.HorarioChegada
	}
	return ""
}

func (x *Voo) GetPreco() string {
	if x != nil {
		return x.Preco
	}
	return ""
}

func (x *Voo) GetDetalhes() string {
	if x != nil {
		return x.Detalhes
	}
	return ""
}

var File_contact_proto protoreflect.FileDescriptor

var file_contact_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc2, 0x01, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x65, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x69, 0x64, 0x61, 0x64, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x69, 0x64, 0x61, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x69, 0x64, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64,
	0x61, 0x74, 0x61, 0x49, 0x64, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x76,
	0x6f, 0x6c, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61,
	0x56, 0x6f, 0x6c, 0x74, 0x61, 0x12, 0x2d, 0x0a, 0x12, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x64,
	0x61, 0x64, 0x65, 0x5f, 0x61, 0x64, 0x75, 0x6c, 0x74, 0x6f, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x11, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x64, 0x61, 0x64, 0x65, 0x41, 0x64, 0x75,
	0x6c, 0x74, 0x6f, 0x73, 0x22, 0x4a, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x06, 0x68, 0x6f, 0x74, 0x65, 0x69, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x52, 0x06,
	0x68, 0x6f, 0x74, 0x65, 0x69, 0x73, 0x12, 0x18, 0x0a, 0x04, 0x76, 0x6f, 0x6f, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x04, 0x2e, 0x56, 0x6f, 0x6f, 0x52, 0x04, 0x76, 0x6f, 0x6f, 0x73,
	0x22, 0xc8, 0x02, 0x0a, 0x05, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x6f,
	0x6d, 0x65, 0x5f, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6e, 0x6f, 0x6d, 0x65, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x69, 0x64,
	0x61, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x69, 0x64, 0x61, 0x64,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x65, 0x63, 0x6f, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x65, 0x63, 0x6f, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x72, 0x65, 0x63, 0x6f, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0a, 0x70, 0x72, 0x65, 0x63, 0x6f, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x29,
	0x0a, 0x10, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x63, 0x61, 0x6f, 0x5f, 0x71, 0x75, 0x61, 0x72,
	0x74, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x63, 0x61, 0x6f, 0x51, 0x75, 0x61, 0x72, 0x74, 0x6f, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x69, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x2d, 0x0a, 0x12, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x64, 0x61, 0x64, 0x65, 0x5f,
	0x61, 0x64, 0x75, 0x6c, 0x74, 0x6f, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x64, 0x61, 0x64, 0x65, 0x41, 0x64, 0x75, 0x6c, 0x74, 0x6f, 0x73,
	0x12, 0x29, 0x0a, 0x10, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x61, 0x5f, 0x71, 0x75,
	0x61, 0x72, 0x74, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x69, 0x61, 0x51, 0x75, 0x61, 0x72, 0x74, 0x6f, 0x22, 0x89, 0x01, 0x0a, 0x03,
	0x56, 0x6f, 0x6f, 0x12, 0x27, 0x0a, 0x0f, 0x68, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x5f, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x64, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x68, 0x6f,
	0x72, 0x61, 0x72, 0x69, 0x6f, 0x50, 0x61, 0x72, 0x74, 0x69, 0x64, 0x61, 0x12, 0x27, 0x0a, 0x0f,
	0x68, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x5f, 0x63, 0x68, 0x65, 0x67, 0x61, 0x64, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x68, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x43, 0x68,
	0x65, 0x67, 0x61, 0x64, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x65, 0x63, 0x6f, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x65, 0x63, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x65, 0x74, 0x61, 0x6c, 0x68, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x65, 0x74, 0x61, 0x6c, 0x68, 0x65, 0x73, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x46, 0x65, 0x6c, 0x69, 0x70, 0x65, 0x46, 0x41, 0x6c, 0x76,
	0x65, 0x73, 0x2f, 0x76, 0x69, 0x61, 0x67, 0x65, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contact_proto_rawDescOnce sync.Once
	file_contact_proto_rawDescData = file_contact_proto_rawDesc
)

func file_contact_proto_rawDescGZIP() []byte {
	file_contact_proto_rawDescOnce.Do(func() {
		file_contact_proto_rawDescData = protoimpl.X.CompressGZIP(file_contact_proto_rawDescData)
	})
	return file_contact_proto_rawDescData
}

var file_contact_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_contact_proto_goTypes = []interface{}{
	(*TravelRequest)(nil),  // 0: TravelRequest
	(*TravelResponse)(nil), // 1: TravelResponse
	(*Hotel)(nil),          // 2: Hotel
	(*Voo)(nil),            // 3: Voo
}
var file_contact_proto_depIdxs = []int32{
	2, // 0: TravelResponse.hoteis:type_name -> Hotel
	3, // 1: TravelResponse.voos:type_name -> Voo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_contact_proto_init() }
func file_contact_proto_init() {
	if File_contact_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contact_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TravelRequest); i {
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
		file_contact_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TravelResponse); i {
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
		file_contact_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hotel); i {
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
		file_contact_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Voo); i {
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
			RawDescriptor: file_contact_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_contact_proto_goTypes,
		DependencyIndexes: file_contact_proto_depIdxs,
		MessageInfos:      file_contact_proto_msgTypes,
	}.Build()
	File_contact_proto = out.File
	file_contact_proto_rawDesc = nil
	file_contact_proto_goTypes = nil
	file_contact_proto_depIdxs = nil
}
