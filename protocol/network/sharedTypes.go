package network


import (
    "../security"
    "../runtime"
)
const (
    ConnectionTypeNone ConnectionType = "none"
    ConnectionTypeCellular2g ConnectionType = "cellular2g"
    ConnectionTypeCellular3g ConnectionType = "cellular3g"
    ConnectionTypeCellular4g ConnectionType = "cellular4g"
    ConnectionTypeBluetooth ConnectionType = "bluetooth"
    ConnectionTypeEthernet ConnectionType = "ethernet"
    ConnectionTypeWifi ConnectionType = "wifi"
    ConnectionTypeWimax ConnectionType = "wimax"
    ConnectionTypeOther ConnectionType = "other"
)

const (
    CookieSameSiteStrict CookieSameSite = "Strict"
    CookieSameSiteLax CookieSameSite = "Lax"
)

const (
    ResourcePriorityVeryLow ResourcePriority = "VeryLow"
    ResourcePriorityLow ResourcePriority = "Low"
    ResourcePriorityMedium ResourcePriority = "Medium"
    ResourcePriorityHigh ResourcePriority = "High"
    ResourcePriorityVeryHigh ResourcePriority = "VeryHigh"
)

type RequestMixedContentTypeEnum string
const (
    RequestMixedContentTypeBlockable RequestMixedContentTypeEnum = "blockable"
    RequestMixedContentTypeOptionallyDashblockable RequestMixedContentTypeEnum = "optionally-blockable"
    RequestMixedContentTypeNone RequestMixedContentTypeEnum = "none"
)

type RequestReferrerPolicyEnum string
const (
    RequestReferrerPolicyUnsafeDashurl RequestReferrerPolicyEnum = "unsafe-url"
    RequestReferrerPolicyNoDashreferrerwhendowngrade RequestReferrerPolicyEnum = "no-referrer-when-downgrade"
    RequestReferrerPolicyNoDashreferrer RequestReferrerPolicyEnum = "no-referrer"
    RequestReferrerPolicyOrigin RequestReferrerPolicyEnum = "origin"
    RequestReferrerPolicyOriginDashwhencrossorigin RequestReferrerPolicyEnum = "origin-when-cross-origin"
    RequestReferrerPolicyNoDashreferrerwhendowngradeoriginwhencrossorigin RequestReferrerPolicyEnum = "no-referrer-when-downgrade-origin-when-cross-origin"
)

const (
    BlockedReasonCsp BlockedReason = "csp"
    BlockedReasonMixedDashcontent BlockedReason = "mixed-content"
    BlockedReasonOrigin BlockedReason = "origin"
    BlockedReasonInspector BlockedReason = "inspector"
    BlockedReasonSubresourceDashfilter BlockedReason = "subresource-filter"
    BlockedReasonOther BlockedReason = "other"
)

type InitiatorTypeEnum string
const (
    InitiatorTypeParser InitiatorTypeEnum = "parser"
    InitiatorTypeScript InitiatorTypeEnum = "script"
    InitiatorTypePreload InitiatorTypeEnum = "preload"
    InitiatorTypeOther InitiatorTypeEnum = "other"
)
type LoaderId string

type RequestId string

type Timestamp float64

type Headers map[string]string

type ConnectionType string

type CookieSameSite string

type ResourceTiming struct {
    RequestTime float64 `json:"requestTime"`// Timing's requestTime is a baseline in seconds, while the other numbers are ticks in milliseconds relatively to this requestTime.
    ProxyStart float64 `json:"proxyStart"`// Started resolving proxy.
    ProxyEnd float64 `json:"proxyEnd"`// Finished resolving proxy.
    DnsStart float64 `json:"dnsStart"`// Started DNS address resolve.
    DnsEnd float64 `json:"dnsEnd"`// Finished DNS address resolve.
    ConnectStart float64 `json:"connectStart"`// Started connecting to the remote host.
    ConnectEnd float64 `json:"connectEnd"`// Connected to the remote host.
    SslStart float64 `json:"sslStart"`// Started SSL handshake.
    SslEnd float64 `json:"sslEnd"`// Finished SSL handshake.
    WorkerStart float64 `json:"workerStart"`// [Experimental] Started running ServiceWorker.
    WorkerReady float64 `json:"workerReady"`// [Experimental] Finished Starting ServiceWorker.
    SendStart float64 `json:"sendStart"`// Started sending request.
    SendEnd float64 `json:"sendEnd"`// Finished sending request.
    PushStart float64 `json:"pushStart"`// [Experimental] Time the server started pushing request.
    PushEnd float64 `json:"pushEnd"`// [Experimental] Time the server finished pushing request.
    ReceiveHeadersEnd float64 `json:"receiveHeadersEnd"`// Finished receiving response headers.
}

type ResourcePriority string

type Request struct {
    Url string `json:"url"`// Request URL.
    Method string `json:"method"`// HTTP request method.
    Headers Headers `json:"headers"`// HTTP request headers.
    PostData *string `json:"postData,omitempty"`// HTTP POST request data.
    MixedContentType *RequestMixedContentTypeEnum `json:"mixedContentType,omitempty"`// The mixed content status of the request, as defined in http://www.w3.org/TR/mixed-content/
    InitialPriority ResourcePriority `json:"initialPriority"`// Priority of the resource request at the time request is sent.
    ReferrerPolicy RequestReferrerPolicyEnum `json:"referrerPolicy"`// The referrer policy of the request, as defined in https://www.w3.org/TR/referrer-policy/
}

type SignedCertificateTimestamp struct {
    Status string `json:"status"`// Validation status.
    Origin string `json:"origin"`// Origin.
    LogDescription string `json:"logDescription"`// Log name / description.
    LogId string `json:"logId"`// Log ID.
    Timestamp Timestamp `json:"timestamp"`// Issuance date.
    HashAlgorithm string `json:"hashAlgorithm"`// Hash algorithm.
    SignatureAlgorithm string `json:"signatureAlgorithm"`// Signature algorithm.
    SignatureData string `json:"signatureData"`// Signature data.
}

type SecurityDetails struct {
    Protocol string `json:"protocol"`// Protocol name (e.g. "TLS 1.2" or "QUIC").
    KeyExchange string `json:"keyExchange"`// Key Exchange used by the connection, or the empty string if not applicable.
    KeyExchangeGroup *string `json:"keyExchangeGroup,omitempty"`// (EC)DH group used by the connection, if applicable.
    Cipher string `json:"cipher"`// Cipher name.
    Mac *string `json:"mac,omitempty"`// TLS MAC. Note that AEAD ciphers do not have separate MACs.
    CertificateId security.CertificateId `json:"certificateId"`// Certificate ID value.
    SubjectName string `json:"subjectName"`// Certificate subject name.
    SanList []string `json:"sanList"`// Subject Alternative Name (SAN) DNS names and IP addresses.
    Issuer string `json:"issuer"`// Name of the issuing CA.
    ValidFrom Timestamp `json:"validFrom"`// Certificate valid from date.
    ValidTo Timestamp `json:"validTo"`// Certificate valid to (expiration) date
    SignedCertificateTimestampList []SignedCertificateTimestamp `json:"signedCertificateTimestampList"`// List of signed certificate timestamps (SCTs).
}

type BlockedReason string

type Response struct {
    Url string `json:"url"`// Response URL. This URL can be different from CachedResource.url in case of redirect.
    Status float64 `json:"status"`// HTTP response status code.
    StatusText string `json:"statusText"`// HTTP response status text.
    Headers Headers `json:"headers"`// HTTP response headers.
    HeadersText *string `json:"headersText,omitempty"`// HTTP response headers text.
    MimeType string `json:"mimeType"`// Resource mimeType as determined by the browser.
    RequestHeaders *Headers `json:"requestHeaders,omitempty"`// Refined HTTP request headers that were actually transmitted over the network.
    RequestHeadersText *string `json:"requestHeadersText,omitempty"`// HTTP request headers text.
    ConnectionReused bool `json:"connectionReused"`// Specifies whether physical connection was actually reused for this request.
    ConnectionId float64 `json:"connectionId"`// Physical connection id that was actually used for this request.
    RemoteIPAddress *string `json:"remoteIPAddress,omitempty"`// [Experimental] Remote IP address.
    RemotePort *int64 `json:"remotePort,omitempty"`// [Experimental] Remote port.
    FromDiskCache *bool `json:"fromDiskCache,omitempty"`// Specifies that the request was served from the disk cache.
    FromServiceWorker *bool `json:"fromServiceWorker,omitempty"`// Specifies that the request was served from the ServiceWorker.
    EncodedDataLength float64 `json:"encodedDataLength"`// Total number of bytes received for this request so far.
    Timing *ResourceTiming `json:"timing,omitempty"`// Timing information for the given request.
    Protocol *string `json:"protocol,omitempty"`// Protocol used to fetch this request.
    SecurityState security.SecurityState `json:"securityState"`// Security state of the request resource.
    SecurityDetails *SecurityDetails `json:"securityDetails,omitempty"`// Security details for the request.
}

type WebSocketRequest struct {
    Headers Headers `json:"headers"`// HTTP request headers.
}

type WebSocketResponse struct {
    Status float64 `json:"status"`// HTTP response status code.
    StatusText string `json:"statusText"`// HTTP response status text.
    Headers Headers `json:"headers"`// HTTP response headers.
    HeadersText *string `json:"headersText,omitempty"`// HTTP response headers text.
    RequestHeaders *Headers `json:"requestHeaders,omitempty"`// HTTP request headers.
    RequestHeadersText *string `json:"requestHeadersText,omitempty"`// HTTP request headers text.
}

type WebSocketFrame struct {
    Opcode float64 `json:"opcode"`// WebSocket frame opcode.
    Mask bool `json:"mask"`// WebSocke frame mask.
    PayloadData string `json:"payloadData"`// WebSocke frame payload data.
}

type CachedResource struct {
    Url string `json:"url"`// Resource URL. This is the url of the original network request.
    Type runtime.ResourceType `json:"type"`// Type of this resource.
    Response *Response `json:"response,omitempty"`// Cached response data.
    BodySize float64 `json:"bodySize"`// Cached response body size.
}

type Initiator struct {
    Type InitiatorTypeEnum `json:"type"`// Type of this initiator.
    Stack *runtime.StackTrace `json:"stack,omitempty"`// Initiator JavaScript stack trace, set for Script only.
    Url *string `json:"url,omitempty"`// Initiator URL, set for Parser type only.
    LineNumber *float64 `json:"lineNumber,omitempty"`// Initiator line number, set for Parser type only (0-based).
}

type Cookie struct {
    Name string `json:"name"`// Cookie name.
    Value string `json:"value"`// Cookie value.
    Domain string `json:"domain"`// Cookie domain.
    Path string `json:"path"`// Cookie path.
    Expires float64 `json:"expires"`// Cookie expiration date as the number of seconds since the UNIX epoch.
    Size int64 `json:"size"`// Cookie size.
    HttpOnly bool `json:"httpOnly"`// True if cookie is http-only.
    Secure bool `json:"secure"`// True if cookie is secure.
    Session bool `json:"session"`// True in case of session cookie.
    SameSite *CookieSameSite `json:"sameSite,omitempty"`// Cookie SameSite type.
}
