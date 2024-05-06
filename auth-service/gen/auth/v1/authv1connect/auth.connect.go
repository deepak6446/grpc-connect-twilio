// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: auth/v1/auth.proto

package authv1connect

import (
	v1 "auth-service/gen/auth/v1"
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// AuthServiceName is the fully-qualified name of the AuthService service.
	AuthServiceName = "auth.v1.AuthService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AuthServiceSignupWithPhoneNumberProcedure is the fully-qualified name of the AuthService's
	// SignupWithPhoneNumber RPC.
	AuthServiceSignupWithPhoneNumberProcedure = "/auth.v1.AuthService/SignupWithPhoneNumber"
	// AuthServiceGetProfileProcedure is the fully-qualified name of the AuthService's GetProfile RPC.
	AuthServiceGetProfileProcedure = "/auth.v1.AuthService/GetProfile"
	// AuthServiceValidatePhoneNumberLoginProcedure is the fully-qualified name of the AuthService's
	// ValidatePhoneNumberLogin RPC.
	AuthServiceValidatePhoneNumberLoginProcedure = "/auth.v1.AuthService/ValidatePhoneNumberLogin"
	// AuthServiceLoginWithPhoneNumberProcedure is the fully-qualified name of the AuthService's
	// LoginWithPhoneNumber RPC.
	AuthServiceLoginWithPhoneNumberProcedure = "/auth.v1.AuthService/LoginWithPhoneNumber"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	authServiceServiceDescriptor                        = v1.File_auth_v1_auth_proto.Services().ByName("AuthService")
	authServiceSignupWithPhoneNumberMethodDescriptor    = authServiceServiceDescriptor.Methods().ByName("SignupWithPhoneNumber")
	authServiceGetProfileMethodDescriptor               = authServiceServiceDescriptor.Methods().ByName("GetProfile")
	authServiceValidatePhoneNumberLoginMethodDescriptor = authServiceServiceDescriptor.Methods().ByName("ValidatePhoneNumberLogin")
	authServiceLoginWithPhoneNumberMethodDescriptor     = authServiceServiceDescriptor.Methods().ByName("LoginWithPhoneNumber")
)

// AuthServiceClient is a client for the auth.v1.AuthService service.
type AuthServiceClient interface {
	// Signs up a user with a phone number.
	SignupWithPhoneNumber(context.Context, *connect.Request[v1.SignupRequest]) (*connect.Response[v1.SignupResponse], error)
	GetProfile(context.Context, *connect.Request[v1.GetProfileRequest]) (*connect.Response[v1.GetProfileResponse], error)
	ValidatePhoneNumberLogin(context.Context, *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error)
	LoginWithPhoneNumber(context.Context, *connect.Request[v1.VerifyPhoneNumberRequest]) (*connect.Response[v1.VerifyPhoneNumberResponse], error)
}

// NewAuthServiceClient constructs a client for the auth.v1.AuthService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAuthServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) AuthServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &authServiceClient{
		signupWithPhoneNumber: connect.NewClient[v1.SignupRequest, v1.SignupResponse](
			httpClient,
			baseURL+AuthServiceSignupWithPhoneNumberProcedure,
			connect.WithSchema(authServiceSignupWithPhoneNumberMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getProfile: connect.NewClient[v1.GetProfileRequest, v1.GetProfileResponse](
			httpClient,
			baseURL+AuthServiceGetProfileProcedure,
			connect.WithSchema(authServiceGetProfileMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		validatePhoneNumberLogin: connect.NewClient[v1.LoginRequest, v1.LoginResponse](
			httpClient,
			baseURL+AuthServiceValidatePhoneNumberLoginProcedure,
			connect.WithSchema(authServiceValidatePhoneNumberLoginMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		loginWithPhoneNumber: connect.NewClient[v1.VerifyPhoneNumberRequest, v1.VerifyPhoneNumberResponse](
			httpClient,
			baseURL+AuthServiceLoginWithPhoneNumberProcedure,
			connect.WithSchema(authServiceLoginWithPhoneNumberMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// authServiceClient implements AuthServiceClient.
type authServiceClient struct {
	signupWithPhoneNumber    *connect.Client[v1.SignupRequest, v1.SignupResponse]
	getProfile               *connect.Client[v1.GetProfileRequest, v1.GetProfileResponse]
	validatePhoneNumberLogin *connect.Client[v1.LoginRequest, v1.LoginResponse]
	loginWithPhoneNumber     *connect.Client[v1.VerifyPhoneNumberRequest, v1.VerifyPhoneNumberResponse]
}

// SignupWithPhoneNumber calls auth.v1.AuthService.SignupWithPhoneNumber.
func (c *authServiceClient) SignupWithPhoneNumber(ctx context.Context, req *connect.Request[v1.SignupRequest]) (*connect.Response[v1.SignupResponse], error) {
	return c.signupWithPhoneNumber.CallUnary(ctx, req)
}

// GetProfile calls auth.v1.AuthService.GetProfile.
func (c *authServiceClient) GetProfile(ctx context.Context, req *connect.Request[v1.GetProfileRequest]) (*connect.Response[v1.GetProfileResponse], error) {
	return c.getProfile.CallUnary(ctx, req)
}

// ValidatePhoneNumberLogin calls auth.v1.AuthService.ValidatePhoneNumberLogin.
func (c *authServiceClient) ValidatePhoneNumberLogin(ctx context.Context, req *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error) {
	return c.validatePhoneNumberLogin.CallUnary(ctx, req)
}

// LoginWithPhoneNumber calls auth.v1.AuthService.LoginWithPhoneNumber.
func (c *authServiceClient) LoginWithPhoneNumber(ctx context.Context, req *connect.Request[v1.VerifyPhoneNumberRequest]) (*connect.Response[v1.VerifyPhoneNumberResponse], error) {
	return c.loginWithPhoneNumber.CallUnary(ctx, req)
}

// AuthServiceHandler is an implementation of the auth.v1.AuthService service.
type AuthServiceHandler interface {
	// Signs up a user with a phone number.
	SignupWithPhoneNumber(context.Context, *connect.Request[v1.SignupRequest]) (*connect.Response[v1.SignupResponse], error)
	GetProfile(context.Context, *connect.Request[v1.GetProfileRequest]) (*connect.Response[v1.GetProfileResponse], error)
	ValidatePhoneNumberLogin(context.Context, *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error)
	LoginWithPhoneNumber(context.Context, *connect.Request[v1.VerifyPhoneNumberRequest]) (*connect.Response[v1.VerifyPhoneNumberResponse], error)
}

// NewAuthServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAuthServiceHandler(svc AuthServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	authServiceSignupWithPhoneNumberHandler := connect.NewUnaryHandler(
		AuthServiceSignupWithPhoneNumberProcedure,
		svc.SignupWithPhoneNumber,
		connect.WithSchema(authServiceSignupWithPhoneNumberMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	authServiceGetProfileHandler := connect.NewUnaryHandler(
		AuthServiceGetProfileProcedure,
		svc.GetProfile,
		connect.WithSchema(authServiceGetProfileMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	authServiceValidatePhoneNumberLoginHandler := connect.NewUnaryHandler(
		AuthServiceValidatePhoneNumberLoginProcedure,
		svc.ValidatePhoneNumberLogin,
		connect.WithSchema(authServiceValidatePhoneNumberLoginMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	authServiceLoginWithPhoneNumberHandler := connect.NewUnaryHandler(
		AuthServiceLoginWithPhoneNumberProcedure,
		svc.LoginWithPhoneNumber,
		connect.WithSchema(authServiceLoginWithPhoneNumberMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/auth.v1.AuthService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AuthServiceSignupWithPhoneNumberProcedure:
			authServiceSignupWithPhoneNumberHandler.ServeHTTP(w, r)
		case AuthServiceGetProfileProcedure:
			authServiceGetProfileHandler.ServeHTTP(w, r)
		case AuthServiceValidatePhoneNumberLoginProcedure:
			authServiceValidatePhoneNumberLoginHandler.ServeHTTP(w, r)
		case AuthServiceLoginWithPhoneNumberProcedure:
			authServiceLoginWithPhoneNumberHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAuthServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAuthServiceHandler struct{}

func (UnimplementedAuthServiceHandler) SignupWithPhoneNumber(context.Context, *connect.Request[v1.SignupRequest]) (*connect.Response[v1.SignupResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("auth.v1.AuthService.SignupWithPhoneNumber is not implemented"))
}

func (UnimplementedAuthServiceHandler) GetProfile(context.Context, *connect.Request[v1.GetProfileRequest]) (*connect.Response[v1.GetProfileResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("auth.v1.AuthService.GetProfile is not implemented"))
}

func (UnimplementedAuthServiceHandler) ValidatePhoneNumberLogin(context.Context, *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("auth.v1.AuthService.ValidatePhoneNumberLogin is not implemented"))
}

func (UnimplementedAuthServiceHandler) LoginWithPhoneNumber(context.Context, *connect.Request[v1.VerifyPhoneNumberRequest]) (*connect.Response[v1.VerifyPhoneNumberResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("auth.v1.AuthService.LoginWithPhoneNumber is not implemented"))
}
