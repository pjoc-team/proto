// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pay_settlement.proto

package pay

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *SettlementRequest) Validate() error {
	return nil
}
func (this *SettlementPayOrder) Validate() error {
	if this.Order != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Order); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Order", err)
		}
	}
	return nil
}
func (this *SettlementResponse) Validate() error {
	if this.Result != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Result); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Result", err)
		}
	}
	return nil
}
