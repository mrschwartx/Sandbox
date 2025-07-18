// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/api/course/v1/booking.proto

package v1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/anypb"
	_ "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Status int32

const (
	Status_BOOKING_UNSPECIFIED Status = 0
	Status_CREATED             Status = 1
	Status_RESERVED            Status = 2
	Status_COMPLETED           Status = 3
	Status_FAILED              Status = 4
	Status_EXPIRED             Status = 5
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "BOOKING_UNSPECIFIED",
		1: "CREATED",
		2: "RESERVED",
		3: "COMPLETED",
		4: "FAILED",
		5: "EXPIRED",
	}
	Status_value = map[string]int32{
		"BOOKING_UNSPECIFIED": 0,
		"CREATED":             1,
		"RESERVED":            2,
		"COMPLETED":           3,
		"FAILED":              4,
		"EXPIRED":             5,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_api_course_v1_booking_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_pkg_api_course_v1_booking_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_pkg_api_course_v1_booking_proto_rawDescGZIP(), []int{0}
}

type Booking struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Number        string                 `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Course        string                 `protobuf:"bytes,2,opt,name=course,proto3" json:"course,omitempty"`
	Batch         string                 `protobuf:"bytes,3,opt,name=batch,proto3" json:"batch,omitempty"`
	Price         float64                `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	Currency      string                 `protobuf:"bytes,5,opt,name=currency,proto3" json:"currency,omitempty"`
	Status        Status                 `protobuf:"varint,6,opt,name=status,proto3,enum=example.com.course.v1.Status" json:"status,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	ReservedAt    *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=reserved_at,json=reservedAt,proto3" json:"reserved_at,omitempty"`
	PaidAt        *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=paid_at,json=paidAt,proto3" json:"paid_at,omitempty"`
	Customer      *Customer              `protobuf:"bytes,10,opt,name=customer,proto3" json:"customer,omitempty"`
	Payment       *Payment               `protobuf:"bytes,11,opt,name=payment,proto3" json:"payment,omitempty"`
	ExpiredAt     *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at,omitempty"`
	FailedAt      *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=failed_at,json=failedAt,proto3" json:"failed_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Booking) Reset() {
	*x = Booking{}
	mi := &file_pkg_api_course_v1_booking_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Booking) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Booking) ProtoMessage() {}

func (x *Booking) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_course_v1_booking_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Booking.ProtoReflect.Descriptor instead.
func (*Booking) Descriptor() ([]byte, []int) {
	return file_pkg_api_course_v1_booking_proto_rawDescGZIP(), []int{0}
}

func (x *Booking) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *Booking) GetCourse() string {
	if x != nil {
		return x.Course
	}
	return ""
}

func (x *Booking) GetBatch() string {
	if x != nil {
		return x.Batch
	}
	return ""
}

func (x *Booking) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Booking) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Booking) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_BOOKING_UNSPECIFIED
}

func (x *Booking) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Booking) GetReservedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ReservedAt
	}
	return nil
}

func (x *Booking) GetPaidAt() *timestamppb.Timestamp {
	if x != nil {
		return x.PaidAt
	}
	return nil
}

func (x *Booking) GetCustomer() *Customer {
	if x != nil {
		return x.Customer
	}
	return nil
}

func (x *Booking) GetPayment() *Payment {
	if x != nil {
		return x.Payment
	}
	return nil
}

func (x *Booking) GetExpiredAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiredAt
	}
	return nil
}

func (x *Booking) GetFailedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.FailedAt
	}
	return nil
}

type Customer struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Name            string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email           string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber     string                 `protobuf:"bytes,3,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	ShippingAddress *Address               `protobuf:"bytes,4,opt,name=shipping_address,json=shippingAddress,proto3" json:"shipping_address,omitempty"`
	BillingAddress  *Address               `protobuf:"bytes,5,opt,name=billing_address,json=billingAddress,proto3" json:"billing_address,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *Customer) Reset() {
	*x = Customer{}
	mi := &file_pkg_api_course_v1_booking_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Customer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Customer) ProtoMessage() {}

func (x *Customer) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_course_v1_booking_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Customer.ProtoReflect.Descriptor instead.
func (*Customer) Descriptor() ([]byte, []int) {
	return file_pkg_api_course_v1_booking_proto_rawDescGZIP(), []int{1}
}

func (x *Customer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Customer) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Customer) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *Customer) GetShippingAddress() *Address {
	if x != nil {
		return x.ShippingAddress
	}
	return nil
}

func (x *Customer) GetBillingAddress() *Address {
	if x != nil {
		return x.BillingAddress
	}
	return nil
}

type Address struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StreetAddress string                 `protobuf:"bytes,1,opt,name=street_address,json=streetAddress,proto3" json:"street_address,omitempty"`
	AptSuite      string                 `protobuf:"bytes,2,opt,name=apt_suite,json=aptSuite,proto3" json:"apt_suite,omitempty"`
	City          string                 `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Country       string                 `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	ZipCode       string                 `protobuf:"bytes,5,opt,name=zip_code,json=zipCode,proto3" json:"zip_code,omitempty"`
	State         string                 `protobuf:"bytes,6,opt,name=state,proto3" json:"state,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Address) Reset() {
	*x = Address{}
	mi := &file_pkg_api_course_v1_booking_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_course_v1_booking_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_pkg_api_course_v1_booking_proto_rawDescGZIP(), []int{2}
}

func (x *Address) GetStreetAddress() string {
	if x != nil {
		return x.StreetAddress
	}
	return ""
}

func (x *Address) GetAptSuite() string {
	if x != nil {
		return x.AptSuite
	}
	return ""
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Address) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Address) GetZipCode() string {
	if x != nil {
		return x.ZipCode
	}
	return ""
}

func (x *Address) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type Payment struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	InvoiceNumber string                 `protobuf:"bytes,1,opt,name=invoice_number,json=invoiceNumber,proto3" json:"invoice_number,omitempty"`
	Method        string                 `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Payment) Reset() {
	*x = Payment{}
	mi := &file_pkg_api_course_v1_booking_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Payment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payment) ProtoMessage() {}

func (x *Payment) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_course_v1_booking_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payment.ProtoReflect.Descriptor instead.
func (*Payment) Descriptor() ([]byte, []int) {
	return file_pkg_api_course_v1_booking_proto_rawDescGZIP(), []int{3}
}

func (x *Payment) GetInvoiceNumber() string {
	if x != nil {
		return x.InvoiceNumber
	}
	return ""
}

func (x *Payment) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

type ListBookingsParam struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Invoice       string                 `protobuf:"bytes,1,opt,name=invoice,proto3" json:"invoice,omitempty"`
	Status        Status                 `protobuf:"varint,2,opt,name=status,proto3,enum=example.com.course.v1.Status" json:"status,omitempty"`
	PageSize      uint64                 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken     string                 `protobuf:"bytes,4,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	OrderBy       string                 `protobuf:"bytes,5,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListBookingsParam) Reset() {
	*x = ListBookingsParam{}
	mi := &file_pkg_api_course_v1_booking_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListBookingsParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBookingsParam) ProtoMessage() {}

func (x *ListBookingsParam) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_course_v1_booking_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBookingsParam.ProtoReflect.Descriptor instead.
func (*ListBookingsParam) Descriptor() ([]byte, []int) {
	return file_pkg_api_course_v1_booking_proto_rawDescGZIP(), []int{4}
}

func (x *ListBookingsParam) GetInvoice() string {
	if x != nil {
		return x.Invoice
	}
	return ""
}

func (x *ListBookingsParam) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_BOOKING_UNSPECIFIED
}

func (x *ListBookingsParam) GetPageSize() uint64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListBookingsParam) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListBookingsParam) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

type ListBookingsResult struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Bookings      []*Booking             `protobuf:"bytes,1,rep,name=bookings,proto3" json:"bookings,omitempty"`
	NextPageToken string                 `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListBookingsResult) Reset() {
	*x = ListBookingsResult{}
	mi := &file_pkg_api_course_v1_booking_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListBookingsResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBookingsResult) ProtoMessage() {}

func (x *ListBookingsResult) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_course_v1_booking_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBookingsResult.ProtoReflect.Descriptor instead.
func (*ListBookingsResult) Descriptor() ([]byte, []int) {
	return file_pkg_api_course_v1_booking_proto_rawDescGZIP(), []int{5}
}

func (x *ListBookingsResult) GetBookings() []*Booking {
	if x != nil {
		return x.Bookings
	}
	return nil
}

func (x *ListBookingsResult) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type CreateBookingParam struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Booking       *Booking               `protobuf:"bytes,1,opt,name=booking,proto3" json:"booking,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateBookingParam) Reset() {
	*x = CreateBookingParam{}
	mi := &file_pkg_api_course_v1_booking_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBookingParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookingParam) ProtoMessage() {}

func (x *CreateBookingParam) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_course_v1_booking_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookingParam.ProtoReflect.Descriptor instead.
func (*CreateBookingParam) Descriptor() ([]byte, []int) {
	return file_pkg_api_course_v1_booking_proto_rawDescGZIP(), []int{6}
}

func (x *CreateBookingParam) GetBooking() *Booking {
	if x != nil {
		return x.Booking
	}
	return nil
}

type GetBookingParam struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Booking       string                 `protobuf:"bytes,1,opt,name=booking,proto3" json:"booking,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBookingParam) Reset() {
	*x = GetBookingParam{}
	mi := &file_pkg_api_course_v1_booking_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBookingParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookingParam) ProtoMessage() {}

func (x *GetBookingParam) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_course_v1_booking_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookingParam.ProtoReflect.Descriptor instead.
func (*GetBookingParam) Descriptor() ([]byte, []int) {
	return file_pkg_api_course_v1_booking_proto_rawDescGZIP(), []int{7}
}

func (x *GetBookingParam) GetBooking() string {
	if x != nil {
		return x.Booking
	}
	return ""
}

type ReserveBookingParam struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Booking       string                 `protobuf:"bytes,1,opt,name=booking,proto3" json:"booking,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReserveBookingParam) Reset() {
	*x = ReserveBookingParam{}
	mi := &file_pkg_api_course_v1_booking_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReserveBookingParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReserveBookingParam) ProtoMessage() {}

func (x *ReserveBookingParam) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_course_v1_booking_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReserveBookingParam.ProtoReflect.Descriptor instead.
func (*ReserveBookingParam) Descriptor() ([]byte, []int) {
	return file_pkg_api_course_v1_booking_proto_rawDescGZIP(), []int{8}
}

func (x *ReserveBookingParam) GetBooking() string {
	if x != nil {
		return x.Booking
	}
	return ""
}

type ReserveBookingResult struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReserveBookingResult) Reset() {
	*x = ReserveBookingResult{}
	mi := &file_pkg_api_course_v1_booking_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReserveBookingResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReserveBookingResult) ProtoMessage() {}

func (x *ReserveBookingResult) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_course_v1_booking_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReserveBookingResult.ProtoReflect.Descriptor instead.
func (*ReserveBookingResult) Descriptor() ([]byte, []int) {
	return file_pkg_api_course_v1_booking_proto_rawDescGZIP(), []int{9}
}

type SetPaymentDetailParam struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Booking       string                 `protobuf:"bytes,1,opt,name=booking,proto3" json:"booking,omitempty"`
	Payment       *Payment               `protobuf:"bytes,2,opt,name=payment,proto3" json:"payment,omitempty"`
	Customer      *Customer              `protobuf:"bytes,3,opt,name=customer,proto3" json:"customer,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetPaymentDetailParam) Reset() {
	*x = SetPaymentDetailParam{}
	mi := &file_pkg_api_course_v1_booking_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetPaymentDetailParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPaymentDetailParam) ProtoMessage() {}

func (x *SetPaymentDetailParam) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_course_v1_booking_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPaymentDetailParam.ProtoReflect.Descriptor instead.
func (*SetPaymentDetailParam) Descriptor() ([]byte, []int) {
	return file_pkg_api_course_v1_booking_proto_rawDescGZIP(), []int{10}
}

func (x *SetPaymentDetailParam) GetBooking() string {
	if x != nil {
		return x.Booking
	}
	return ""
}

func (x *SetPaymentDetailParam) GetPayment() *Payment {
	if x != nil {
		return x.Payment
	}
	return nil
}

func (x *SetPaymentDetailParam) GetCustomer() *Customer {
	if x != nil {
		return x.Customer
	}
	return nil
}

type SetPaymentDetailResult struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetPaymentDetailResult) Reset() {
	*x = SetPaymentDetailResult{}
	mi := &file_pkg_api_course_v1_booking_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetPaymentDetailResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPaymentDetailResult) ProtoMessage() {}

func (x *SetPaymentDetailResult) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_course_v1_booking_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPaymentDetailResult.ProtoReflect.Descriptor instead.
func (*SetPaymentDetailResult) Descriptor() ([]byte, []int) {
	return file_pkg_api_course_v1_booking_proto_rawDescGZIP(), []int{11}
}

type ExpireBookingParam struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Booking       string                 `protobuf:"bytes,1,opt,name=booking,proto3" json:"booking,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExpireBookingParam) Reset() {
	*x = ExpireBookingParam{}
	mi := &file_pkg_api_course_v1_booking_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExpireBookingParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpireBookingParam) ProtoMessage() {}

func (x *ExpireBookingParam) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_course_v1_booking_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpireBookingParam.ProtoReflect.Descriptor instead.
func (*ExpireBookingParam) Descriptor() ([]byte, []int) {
	return file_pkg_api_course_v1_booking_proto_rawDescGZIP(), []int{12}
}

func (x *ExpireBookingParam) GetBooking() string {
	if x != nil {
		return x.Booking
	}
	return ""
}

type ExpireBookingResult struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExpireBookingResult) Reset() {
	*x = ExpireBookingResult{}
	mi := &file_pkg_api_course_v1_booking_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExpireBookingResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpireBookingResult) ProtoMessage() {}

func (x *ExpireBookingResult) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_course_v1_booking_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpireBookingResult.ProtoReflect.Descriptor instead.
func (*ExpireBookingResult) Descriptor() ([]byte, []int) {
	return file_pkg_api_course_v1_booking_proto_rawDescGZIP(), []int{13}
}

var File_pkg_api_course_v1_booking_proto protoreflect.FileDescriptor

const file_pkg_api_course_v1_booking_proto_rawDesc = "" +
	"\n" +
	"\x1fpkg/api/course/v1/booking.proto\x12\x15example.com.course.v1\x1a\x1cgoogle/api/annotations.proto\x1a\x19google/api/resource.proto\x1a\x1fgoogle/api/field_behavior.proto\x1a\x17google/api/client.proto\x1a\x19google/protobuf/any.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a google/protobuf/field_mask.proto\x1a\x1fpkg/api/course/v1/catalog.proto\"\x8a\x06\n" +
	"\aBooking\x12\x1b\n" +
	"\x06number\x18\x01 \x01(\tB\x03\xe0A\x03R\x06number\x126\n" +
	"\x06course\x18\x02 \x01(\tB\x1e\xfaA\x1b\n" +
	"\x19course.example.com/CourseR\x06course\x129\n" +
	"\x05batch\x18\x03 \x01(\tB#\xfaA \n" +
	"\x1ecourse.example.com/CourseBatchR\x05batch\x12\x19\n" +
	"\x05price\x18\x04 \x01(\x01B\x03\xe0A\x03R\x05price\x12\x1f\n" +
	"\bcurrency\x18\x05 \x01(\tB\x03\xe0A\x03R\bcurrency\x12:\n" +
	"\x06status\x18\x06 \x01(\x0e2\x1d.example.com.course.v1.StatusB\x03\xe0A\x03R\x06status\x12>\n" +
	"\n" +
	"created_at\x18\a \x01(\v2\x1a.google.protobuf.TimestampB\x03\xe0A\x03R\tcreatedAt\x12@\n" +
	"\vreserved_at\x18\b \x01(\v2\x1a.google.protobuf.TimestampB\x03\xe0A\x03R\n" +
	"reservedAt\x128\n" +
	"\apaid_at\x18\t \x01(\v2\x1a.google.protobuf.TimestampB\x03\xe0A\x03R\x06paidAt\x12;\n" +
	"\bcustomer\x18\n" +
	" \x01(\v2\x1f.example.com.course.v1.CustomerR\bcustomer\x128\n" +
	"\apayment\x18\v \x01(\v2\x1e.example.com.course.v1.PaymentR\apayment\x12>\n" +
	"\n" +
	"expired_at\x18\f \x01(\v2\x1a.google.protobuf.TimestampB\x03\xe0A\x03R\texpiredAt\x12<\n" +
	"\tfailed_at\x18\r \x01(\v2\x1a.google.protobuf.TimestampB\x03\xe0A\x03R\bfailedAt:F\xeaAC\n" +
	"\x1acourse.example.com/Booking\x12\x12bookings/{booking}*\bbookings2\abooking\"\xeb\x01\n" +
	"\bCustomer\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x14\n" +
	"\x05email\x18\x02 \x01(\tR\x05email\x12!\n" +
	"\fphone_number\x18\x03 \x01(\tR\vphoneNumber\x12I\n" +
	"\x10shipping_address\x18\x04 \x01(\v2\x1e.example.com.course.v1.AddressR\x0fshippingAddress\x12G\n" +
	"\x0fbilling_address\x18\x05 \x01(\v2\x1e.example.com.course.v1.AddressR\x0ebillingAddress\"\xac\x01\n" +
	"\aAddress\x12%\n" +
	"\x0estreet_address\x18\x01 \x01(\tR\rstreetAddress\x12\x1b\n" +
	"\tapt_suite\x18\x02 \x01(\tR\baptSuite\x12\x12\n" +
	"\x04city\x18\x03 \x01(\tR\x04city\x12\x18\n" +
	"\acountry\x18\x04 \x01(\tR\acountry\x12\x19\n" +
	"\bzip_code\x18\x05 \x01(\tR\azipCode\x12\x14\n" +
	"\x05state\x18\x06 \x01(\tR\x05state\"M\n" +
	"\aPayment\x12*\n" +
	"\x0einvoice_number\x18\x01 \x01(\tB\x03\xe0A\x03R\rinvoiceNumber\x12\x16\n" +
	"\x06method\x18\x02 \x01(\tR\x06method\"\xe0\x01\n" +
	"\x11ListBookingsParam\x12=\n" +
	"\ainvoice\x18\x01 \x01(\tB#\xe0A\x01\xfaA\x1d\n" +
	"\x1bpayment.example.com/InvoiceR\ainvoice\x125\n" +
	"\x06status\x18\x02 \x01(\x0e2\x1d.example.com.course.v1.StatusR\x06status\x12\x1b\n" +
	"\tpage_size\x18\x03 \x01(\x04R\bpageSize\x12\x1d\n" +
	"\n" +
	"page_token\x18\x04 \x01(\tR\tpageToken\x12\x19\n" +
	"\border_by\x18\x05 \x01(\tR\aorderBy\"x\n" +
	"\x12ListBookingsResult\x12:\n" +
	"\bbookings\x18\x01 \x03(\v2\x1e.example.com.course.v1.BookingR\bbookings\x12&\n" +
	"\x0fnext_page_token\x18\x02 \x01(\tR\rnextPageToken\"S\n" +
	"\x12CreateBookingParam\x12=\n" +
	"\abooking\x18\x01 \x01(\v2\x1e.example.com.course.v1.BookingB\x03\xe0A\x02R\abooking\"O\n" +
	"\x0fGetBookingParam\x12<\n" +
	"\abooking\x18\x01 \x01(\tB\"\xe0A\x02\xfaA\x1c\n" +
	"\x1acourse.example.com/BookingR\abooking\"S\n" +
	"\x13ReserveBookingParam\x12<\n" +
	"\abooking\x18\x01 \x01(\tB\"\xe0A\x02\xfaA\x1c\n" +
	"\x1acourse.example.com/BookingR\abooking\"\x16\n" +
	"\x14ReserveBookingResult\"\xd6\x01\n" +
	"\x15SetPaymentDetailParam\x12<\n" +
	"\abooking\x18\x01 \x01(\tB\"\xe0A\x02\xfaA\x1c\n" +
	"\x1acourse.example.com/BookingR\abooking\x12=\n" +
	"\apayment\x18\x02 \x01(\v2\x1e.example.com.course.v1.PaymentB\x03\xe0A\x02R\apayment\x12@\n" +
	"\bcustomer\x18\x03 \x01(\v2\x1f.example.com.course.v1.CustomerB\x03\xe0A\x02R\bcustomer\"\x18\n" +
	"\x16SetPaymentDetailResult\"R\n" +
	"\x12ExpireBookingParam\x12<\n" +
	"\abooking\x18\x01 \x01(\tB\"\xe0A\x02\xfaA\x1c\n" +
	"\x1acourse.example.com/BookingR\abooking\"\x15\n" +
	"\x13ExpireBookingResult*d\n" +
	"\x06Status\x12\x17\n" +
	"\x13BOOKING_UNSPECIFIED\x10\x00\x12\v\n" +
	"\aCREATED\x10\x01\x12\f\n" +
	"\bRESERVED\x10\x02\x12\r\n" +
	"\tCOMPLETED\x10\x03\x12\n" +
	"\n" +
	"\x06FAILED\x10\x04\x12\v\n" +
	"\aEXPIRED\x10\x052\xbf\x06\n" +
	"\x0eBookingService\x12\x95\x01\n" +
	"\fListBookings\x12(.example.com.course.v1.ListBookingsParam\x1a).example.com.course.v1.ListBookingsResult\"0\x92A\x0e\x12\fList Booking\x82\xd3\xe4\x93\x02\x19\x12\x17/api/course/v1/bookings\x12\x9b\x01\n" +
	"\rCreateBooking\x12).example.com.course.v1.CreateBookingParam\x1a\x1e.example.com.course.v1.Booking\"?\x92A\x14\x12\x12Create New Booking\x82\xd3\xe4\x93\x02\":\abooking\"\x17/api/course/v1/bookings\x12\x8f\x01\n" +
	"\n" +
	"GetBooking\x12&.example.com.course.v1.GetBookingParam\x1a\x1e.example.com.course.v1.Booking\"9\x92A\r\x12\vGet Booking\x82\xd3\xe4\x93\x02#\x12!/api/course/v1/bookings/{booking}\x12\xb3\x01\n" +
	"\x0eReserveBooking\x12*.example.com.course.v1.ReserveBookingParam\x1a+.example.com.course.v1.ReserveBookingResult\"H\x92A\x11\x12\x0fReserve Booking\x82\xd3\xe4\x93\x02.:\x01*\")/api/course/v1/bookings/{booking}:reserve\x12\xae\x01\n" +
	"\rExpireBooking\x12).example.com.course.v1.ExpireBookingParam\x1a*.example.com.course.v1.ExpireBookingResult\"F\x92A\x10\x12\x0eExpire Booking\x82\xd3\xe4\x93\x02-:\x01*\"(/api/course/v1/bookings/{booking}:expireB\x1fZ\x1dexample.com/pkg/api/course/v1b\x06proto3"

var (
	file_pkg_api_course_v1_booking_proto_rawDescOnce sync.Once
	file_pkg_api_course_v1_booking_proto_rawDescData []byte
)

func file_pkg_api_course_v1_booking_proto_rawDescGZIP() []byte {
	file_pkg_api_course_v1_booking_proto_rawDescOnce.Do(func() {
		file_pkg_api_course_v1_booking_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_pkg_api_course_v1_booking_proto_rawDesc), len(file_pkg_api_course_v1_booking_proto_rawDesc)))
	})
	return file_pkg_api_course_v1_booking_proto_rawDescData
}

var file_pkg_api_course_v1_booking_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_api_course_v1_booking_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_pkg_api_course_v1_booking_proto_goTypes = []any{
	(Status)(0),                    // 0: example.com.course.v1.Status
	(*Booking)(nil),                // 1: example.com.course.v1.Booking
	(*Customer)(nil),               // 2: example.com.course.v1.Customer
	(*Address)(nil),                // 3: example.com.course.v1.Address
	(*Payment)(nil),                // 4: example.com.course.v1.Payment
	(*ListBookingsParam)(nil),      // 5: example.com.course.v1.ListBookingsParam
	(*ListBookingsResult)(nil),     // 6: example.com.course.v1.ListBookingsResult
	(*CreateBookingParam)(nil),     // 7: example.com.course.v1.CreateBookingParam
	(*GetBookingParam)(nil),        // 8: example.com.course.v1.GetBookingParam
	(*ReserveBookingParam)(nil),    // 9: example.com.course.v1.ReserveBookingParam
	(*ReserveBookingResult)(nil),   // 10: example.com.course.v1.ReserveBookingResult
	(*SetPaymentDetailParam)(nil),  // 11: example.com.course.v1.SetPaymentDetailParam
	(*SetPaymentDetailResult)(nil), // 12: example.com.course.v1.SetPaymentDetailResult
	(*ExpireBookingParam)(nil),     // 13: example.com.course.v1.ExpireBookingParam
	(*ExpireBookingResult)(nil),    // 14: example.com.course.v1.ExpireBookingResult
	(*timestamppb.Timestamp)(nil),  // 15: google.protobuf.Timestamp
}
var file_pkg_api_course_v1_booking_proto_depIdxs = []int32{
	0,  // 0: example.com.course.v1.Booking.status:type_name -> example.com.course.v1.Status
	15, // 1: example.com.course.v1.Booking.created_at:type_name -> google.protobuf.Timestamp
	15, // 2: example.com.course.v1.Booking.reserved_at:type_name -> google.protobuf.Timestamp
	15, // 3: example.com.course.v1.Booking.paid_at:type_name -> google.protobuf.Timestamp
	2,  // 4: example.com.course.v1.Booking.customer:type_name -> example.com.course.v1.Customer
	4,  // 5: example.com.course.v1.Booking.payment:type_name -> example.com.course.v1.Payment
	15, // 6: example.com.course.v1.Booking.expired_at:type_name -> google.protobuf.Timestamp
	15, // 7: example.com.course.v1.Booking.failed_at:type_name -> google.protobuf.Timestamp
	3,  // 8: example.com.course.v1.Customer.shipping_address:type_name -> example.com.course.v1.Address
	3,  // 9: example.com.course.v1.Customer.billing_address:type_name -> example.com.course.v1.Address
	0,  // 10: example.com.course.v1.ListBookingsParam.status:type_name -> example.com.course.v1.Status
	1,  // 11: example.com.course.v1.ListBookingsResult.bookings:type_name -> example.com.course.v1.Booking
	1,  // 12: example.com.course.v1.CreateBookingParam.booking:type_name -> example.com.course.v1.Booking
	4,  // 13: example.com.course.v1.SetPaymentDetailParam.payment:type_name -> example.com.course.v1.Payment
	2,  // 14: example.com.course.v1.SetPaymentDetailParam.customer:type_name -> example.com.course.v1.Customer
	5,  // 15: example.com.course.v1.BookingService.ListBookings:input_type -> example.com.course.v1.ListBookingsParam
	7,  // 16: example.com.course.v1.BookingService.CreateBooking:input_type -> example.com.course.v1.CreateBookingParam
	8,  // 17: example.com.course.v1.BookingService.GetBooking:input_type -> example.com.course.v1.GetBookingParam
	9,  // 18: example.com.course.v1.BookingService.ReserveBooking:input_type -> example.com.course.v1.ReserveBookingParam
	13, // 19: example.com.course.v1.BookingService.ExpireBooking:input_type -> example.com.course.v1.ExpireBookingParam
	6,  // 20: example.com.course.v1.BookingService.ListBookings:output_type -> example.com.course.v1.ListBookingsResult
	1,  // 21: example.com.course.v1.BookingService.CreateBooking:output_type -> example.com.course.v1.Booking
	1,  // 22: example.com.course.v1.BookingService.GetBooking:output_type -> example.com.course.v1.Booking
	10, // 23: example.com.course.v1.BookingService.ReserveBooking:output_type -> example.com.course.v1.ReserveBookingResult
	14, // 24: example.com.course.v1.BookingService.ExpireBooking:output_type -> example.com.course.v1.ExpireBookingResult
	20, // [20:25] is the sub-list for method output_type
	15, // [15:20] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_pkg_api_course_v1_booking_proto_init() }
func file_pkg_api_course_v1_booking_proto_init() {
	if File_pkg_api_course_v1_booking_proto != nil {
		return
	}
	file_pkg_api_course_v1_catalog_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_api_course_v1_booking_proto_rawDesc), len(file_pkg_api_course_v1_booking_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_api_course_v1_booking_proto_goTypes,
		DependencyIndexes: file_pkg_api_course_v1_booking_proto_depIdxs,
		EnumInfos:         file_pkg_api_course_v1_booking_proto_enumTypes,
		MessageInfos:      file_pkg_api_course_v1_booking_proto_msgTypes,
	}.Build()
	File_pkg_api_course_v1_booking_proto = out.File
	file_pkg_api_course_v1_booking_proto_goTypes = nil
	file_pkg_api_course_v1_booking_proto_depIdxs = nil
}
