package types

const (
    SecurityStateUnknown SecurityState = "unknown"
    SecurityStateNeutral SecurityState = "neutral"
    SecurityStateInsecure SecurityState = "insecure"
    SecurityStateWarning SecurityState = "warning"
    SecurityStateSecure SecurityState = "secure"
    SecurityStateInfo SecurityState = "info"
)
type CertificateId int64

type SecurityState string

type SecurityStateExplanation struct {
    SecurityState SecurityState `json:"securityState"`// Security state representing the severity of the factor being explained.
    Summary string `json:"summary"`// Short phrase describing the type of factor.
    Description string `json:"description"`// Full text explanation of the factor.
    HasCertificate bool `json:"hasCertificate"`// True if the page has a certificate.
}

type InsecureContentStatus struct {
    RanMixedContent bool `json:"ranMixedContent"`// True if the page was loaded over HTTPS and ran mixed (HTTP) content such as scripts.
    DisplayedMixedContent bool `json:"displayedMixedContent"`// True if the page was loaded over HTTPS and displayed mixed (HTTP) content such as images.
    RanContentWithCertErrors bool `json:"ranContentWithCertErrors"`// True if the page was loaded over HTTPS without certificate errors, and ran content such as scripts that were loaded with certificate errors.
    DisplayedContentWithCertErrors bool `json:"displayedContentWithCertErrors"`// True if the page was loaded over HTTPS without certificate errors, and displayed content such as images that were loaded with certificate errors.
    RanInsecureContentStyle SecurityState `json:"ranInsecureContentStyle"`// Security state representing a page that ran insecure content.
    DisplayedInsecureContentStyle SecurityState `json:"displayedInsecureContentStyle"`// Security state representing a page that displayed insecure content.
}
