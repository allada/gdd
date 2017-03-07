package security


type SecurityStateChangedEvent struct {
    SecurityState SecurityState `json:"securityState"`// Security state.
    SchemeIsCryptographic bool `json:"schemeIsCryptographic"`// True if the page was loaded over cryptographic transport such as HTTPS.
    Explanations []SecurityStateExplanation `json:"explanations"`// List of explanations for the security state. If the overall security state is `insecure` or `warning`, at least one corresponding explanation should be included.
    InsecureContentStatus InsecureContentStatus `json:"insecureContentStatus"`// Information about insecure content on the page.
    Summary *string `json:"summary,omitempty"`// Overrides user-visible description of the state.
}