// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package http

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-http/sdk/go/http/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetHttp(ctx *pulumi.Context, args *GetHttpArgs, opts ...pulumi.InvokeOption) (*GetHttpResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetHttpResult
	err := ctx.Invoke("http:index/getHttp:getHttp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getHttp.
type GetHttpArgs struct {
	// Certificate data of the Certificate Authority (CA) in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
	CaCertPem *string `pulumi:"caCertPem"`
	// Disables verification of the server's certificate chain and hostname. Defaults to `false`
	Insecure *bool `pulumi:"insecure"`
	// The HTTP Method for the request. Allowed methods are a subset of methods defined in [RFC7231](https://datatracker.ietf.org/doc/html/rfc7231#section-4.3) namely, `GET`, `HEAD`, and `POST`. `POST` support is only intended for read-only URLs, such as submitting a search.
	Method *string `pulumi:"method"`
	// The request body as a string.
	RequestBody *string `pulumi:"requestBody"`
	// A map of request header field names and values.
	RequestHeaders map[string]string `pulumi:"requestHeaders"`
	// The request timeout in milliseconds.
	RequestTimeoutMs *int          `pulumi:"requestTimeoutMs"`
	Retry            *GetHttpRetry `pulumi:"retry"`
	// The URL for the request. Supported schemes are `getHttp` and `https`.
	Url string `pulumi:"url"`
}

// A collection of values returned by getHttp.
type GetHttpResult struct {
	// The response body returned as a string. **NOTE**: This is deprecated, use `responseBody` instead.
	//
	// Deprecated: Use responseBody instead
	Body string `pulumi:"body"`
	// Certificate data of the Certificate Authority (CA) in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
	CaCertPem *string `pulumi:"caCertPem"`
	// The URL used for the request.
	Id string `pulumi:"id"`
	// Disables verification of the server's certificate chain and hostname. Defaults to `false`
	Insecure *bool `pulumi:"insecure"`
	// The HTTP Method for the request. Allowed methods are a subset of methods defined in [RFC7231](https://datatracker.ietf.org/doc/html/rfc7231#section-4.3) namely, `GET`, `HEAD`, and `POST`. `POST` support is only intended for read-only URLs, such as submitting a search.
	Method *string `pulumi:"method"`
	// The request body as a string.
	RequestBody *string `pulumi:"requestBody"`
	// A map of request header field names and values.
	RequestHeaders map[string]string `pulumi:"requestHeaders"`
	// The request timeout in milliseconds.
	RequestTimeoutMs *int `pulumi:"requestTimeoutMs"`
	// The response body returned as a string.
	ResponseBody string `pulumi:"responseBody"`
	// The response body encoded as base64 (standard) as defined in [RFC 4648](https://datatracker.ietf.org/doc/html/rfc4648#section-4).
	ResponseBodyBase64 string `pulumi:"responseBodyBase64"`
	// A map of response header field names and values. Duplicate headers are concatenated according to [RFC2616](https://www.w3.org/Protocols/rfc2616/rfc2616-sec4.html#sec4.2).
	ResponseHeaders map[string]string `pulumi:"responseHeaders"`
	Retry           *GetHttpRetry     `pulumi:"retry"`
	// The HTTP response status code.
	StatusCode int `pulumi:"statusCode"`
	// The URL for the request. Supported schemes are `getHttp` and `https`.
	Url string `pulumi:"url"`
}

func GetHttpOutput(ctx *pulumi.Context, args GetHttpOutputArgs, opts ...pulumi.InvokeOption) GetHttpResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetHttpResult, error) {
			args := v.(GetHttpArgs)
			r, err := GetHttp(ctx, &args, opts...)
			var s GetHttpResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetHttpResultOutput)
}

// A collection of arguments for invoking getHttp.
type GetHttpOutputArgs struct {
	// Certificate data of the Certificate Authority (CA) in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
	CaCertPem pulumi.StringPtrInput `pulumi:"caCertPem"`
	// Disables verification of the server's certificate chain and hostname. Defaults to `false`
	Insecure pulumi.BoolPtrInput `pulumi:"insecure"`
	// The HTTP Method for the request. Allowed methods are a subset of methods defined in [RFC7231](https://datatracker.ietf.org/doc/html/rfc7231#section-4.3) namely, `GET`, `HEAD`, and `POST`. `POST` support is only intended for read-only URLs, such as submitting a search.
	Method pulumi.StringPtrInput `pulumi:"method"`
	// The request body as a string.
	RequestBody pulumi.StringPtrInput `pulumi:"requestBody"`
	// A map of request header field names and values.
	RequestHeaders pulumi.StringMapInput `pulumi:"requestHeaders"`
	// The request timeout in milliseconds.
	RequestTimeoutMs pulumi.IntPtrInput   `pulumi:"requestTimeoutMs"`
	Retry            GetHttpRetryPtrInput `pulumi:"retry"`
	// The URL for the request. Supported schemes are `getHttp` and `https`.
	Url pulumi.StringInput `pulumi:"url"`
}

func (GetHttpOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetHttpArgs)(nil)).Elem()
}

// A collection of values returned by getHttp.
type GetHttpResultOutput struct{ *pulumi.OutputState }

func (GetHttpResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetHttpResult)(nil)).Elem()
}

func (o GetHttpResultOutput) ToGetHttpResultOutput() GetHttpResultOutput {
	return o
}

func (o GetHttpResultOutput) ToGetHttpResultOutputWithContext(ctx context.Context) GetHttpResultOutput {
	return o
}

// The response body returned as a string. **NOTE**: This is deprecated, use `responseBody` instead.
//
// Deprecated: Use responseBody instead
func (o GetHttpResultOutput) Body() pulumi.StringOutput {
	return o.ApplyT(func(v GetHttpResult) string { return v.Body }).(pulumi.StringOutput)
}

// Certificate data of the Certificate Authority (CA) in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
func (o GetHttpResultOutput) CaCertPem() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetHttpResult) *string { return v.CaCertPem }).(pulumi.StringPtrOutput)
}

// The URL used for the request.
func (o GetHttpResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetHttpResult) string { return v.Id }).(pulumi.StringOutput)
}

// Disables verification of the server's certificate chain and hostname. Defaults to `false`
func (o GetHttpResultOutput) Insecure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetHttpResult) *bool { return v.Insecure }).(pulumi.BoolPtrOutput)
}

// The HTTP Method for the request. Allowed methods are a subset of methods defined in [RFC7231](https://datatracker.ietf.org/doc/html/rfc7231#section-4.3) namely, `GET`, `HEAD`, and `POST`. `POST` support is only intended for read-only URLs, such as submitting a search.
func (o GetHttpResultOutput) Method() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetHttpResult) *string { return v.Method }).(pulumi.StringPtrOutput)
}

// The request body as a string.
func (o GetHttpResultOutput) RequestBody() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetHttpResult) *string { return v.RequestBody }).(pulumi.StringPtrOutput)
}

// A map of request header field names and values.
func (o GetHttpResultOutput) RequestHeaders() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetHttpResult) map[string]string { return v.RequestHeaders }).(pulumi.StringMapOutput)
}

// The request timeout in milliseconds.
func (o GetHttpResultOutput) RequestTimeoutMs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetHttpResult) *int { return v.RequestTimeoutMs }).(pulumi.IntPtrOutput)
}

// The response body returned as a string.
func (o GetHttpResultOutput) ResponseBody() pulumi.StringOutput {
	return o.ApplyT(func(v GetHttpResult) string { return v.ResponseBody }).(pulumi.StringOutput)
}

// The response body encoded as base64 (standard) as defined in [RFC 4648](https://datatracker.ietf.org/doc/html/rfc4648#section-4).
func (o GetHttpResultOutput) ResponseBodyBase64() pulumi.StringOutput {
	return o.ApplyT(func(v GetHttpResult) string { return v.ResponseBodyBase64 }).(pulumi.StringOutput)
}

// A map of response header field names and values. Duplicate headers are concatenated according to [RFC2616](https://www.w3.org/Protocols/rfc2616/rfc2616-sec4.html#sec4.2).
func (o GetHttpResultOutput) ResponseHeaders() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetHttpResult) map[string]string { return v.ResponseHeaders }).(pulumi.StringMapOutput)
}

func (o GetHttpResultOutput) Retry() GetHttpRetryPtrOutput {
	return o.ApplyT(func(v GetHttpResult) *GetHttpRetry { return v.Retry }).(GetHttpRetryPtrOutput)
}

// The HTTP response status code.
func (o GetHttpResultOutput) StatusCode() pulumi.IntOutput {
	return o.ApplyT(func(v GetHttpResult) int { return v.StatusCode }).(pulumi.IntOutput)
}

// The URL for the request. Supported schemes are `getHttp` and `https`.
func (o GetHttpResultOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v GetHttpResult) string { return v.Url }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetHttpResultOutput{})
}
