// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: auth.proto

package authpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *JWTHeader) Validate() error {
	return nil
}
func (this *Claims) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *JWT) Validate() error {
	if this.Header != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Header); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Header", err)
		}
	}
	if this.Claims != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Claims); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Claims", err)
		}
	}
	return nil
}
func (this *Credentials) Validate() error {
	return nil
}
func (this *JwtInfo) Validate() error {
	return nil
}
func (this *JwtEvent) Validate() error {
	if this.Info != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Info); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Info", err)
		}
	}
	return nil
}
func (this *FeedResponse) Validate() error {
	return nil
}
func (this *FindJWTRequest) Validate() error {
	return nil
}
func (this *FindJWTResponse) Validate() error {
	return nil
}
func (this *DeleteJWTRequest) Validate() error {
	return nil
}
func (this *DeleteJWTResponse) Validate() error {
	return nil
}
func (this *ListenRequest) Validate() error {
	return nil
}
func (this *CreateCredentialsRequest) Validate() error {
	if nil == this.Credentials {
		return github_com_mwitkow_go_proto_validators.FieldError("Credentials", fmt.Errorf("message must exist"))
	}
	if this.Credentials != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Credentials); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Credentials", err)
		}
	}
	return nil
}
func (this *CreateCredentialsResponse) Validate() error {
	return nil
}
func (this *AuthenticateRequest) Validate() error {
	if !(len(this.Subject) > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Subject", fmt.Errorf(`value '%v' must have a length greater than '0'`, this.Subject))
	}
	if !(len(this.Password) > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Password", fmt.Errorf(`value '%v' must have a length greater than '0'`, this.Password))
	}
	if !(len(this.Audience) > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Audience", fmt.Errorf(`value '%v' must have a length greater than '0'`, this.Audience))
	}
	for _, item := range this.Scope {
		if !(len(item) > 0) {
			return github_com_mwitkow_go_proto_validators.FieldError("Scope", fmt.Errorf(`value '%v' must have a length greater than '0'`, item))
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *GetTokenResponse) Validate() error {
	return nil
}
func (this *RevokeTokenRequest) Validate() error {
	if this.Jwt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Jwt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Jwt", err)
		}
	}
	return nil
}
func (this *RevokeTokenResponse) Validate() error {
	return nil
}
func (this *SetPasswordRequest) Validate() error {
	if !(len(this.User) > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("User", fmt.Errorf(`value '%v' must have a length greater than '0'`, this.User))
	}
	if !(len(this.Email) > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Email", fmt.Errorf(`value '%v' must have a length greater than '0'`, this.Email))
	}
	if !(len(this.OldPassword) > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("OldPassword", fmt.Errorf(`value '%v' must have a length greater than '0'`, this.OldPassword))
	}
	if !(len(this.NewPassword) > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("NewPassword", fmt.Errorf(`value '%v' must have a length greater than '0'`, this.NewPassword))
	}
	return nil
}
func (this *SetPasswordResponse) Validate() error {
	return nil
}
func (this *CreateTokenRequest) Validate() error {
	if !(len(this.Sub) > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Sub", fmt.Errorf(`value '%v' must have a length greater than '0'`, this.Sub))
	}
	if !(len(this.Audience) > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Audience", fmt.Errorf(`value '%v' must have a length greater than '0'`, this.Audience))
	}
	for _, item := range this.Scope {
		if !(len(item) > 0) {
			return github_com_mwitkow_go_proto_validators.FieldError("Scope", fmt.Errorf(`value '%v' must have a length greater than '0'`, item))
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *CreateTokenResponse) Validate() error {
	if this.Jwt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Jwt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Jwt", err)
		}
	}
	return nil
}
