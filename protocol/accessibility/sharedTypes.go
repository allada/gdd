package accessibility


import (
    "../dom"
)
const (
    AXValueTypeBoolean AXValueType = "boolean"
    AXValueTypeTristate AXValueType = "tristate"
    AXValueTypeBooleanOrUndefined AXValueType = "booleanOrUndefined"
    AXValueTypeIdref AXValueType = "idref"
    AXValueTypeIdrefList AXValueType = "idrefList"
    AXValueTypeInteger AXValueType = "integer"
    AXValueTypeNode AXValueType = "node"
    AXValueTypeNodeList AXValueType = "nodeList"
    AXValueTypeNumber AXValueType = "number"
    AXValueTypeString AXValueType = "string"
    AXValueTypeComputedString AXValueType = "computedString"
    AXValueTypeToken AXValueType = "token"
    AXValueTypeTokenList AXValueType = "tokenList"
    AXValueTypeDomRelation AXValueType = "domRelation"
    AXValueTypeRole AXValueType = "role"
    AXValueTypeInternalRole AXValueType = "internalRole"
    AXValueTypeValueUndefined AXValueType = "valueUndefined"
)

const (
    AXValueSourceTypeAttribute AXValueSourceType = "attribute"
    AXValueSourceTypeImplicit AXValueSourceType = "implicit"
    AXValueSourceTypeStyle AXValueSourceType = "style"
    AXValueSourceTypeContents AXValueSourceType = "contents"
    AXValueSourceTypePlaceholder AXValueSourceType = "placeholder"
    AXValueSourceTypeRelatedElement AXValueSourceType = "relatedElement"
)

const (
    AXValueNativeSourceTypeFigcaption AXValueNativeSourceType = "figcaption"
    AXValueNativeSourceTypeLabel AXValueNativeSourceType = "label"
    AXValueNativeSourceTypeLabelfor AXValueNativeSourceType = "labelfor"
    AXValueNativeSourceTypeLabelwrapped AXValueNativeSourceType = "labelwrapped"
    AXValueNativeSourceTypeLegend AXValueNativeSourceType = "legend"
    AXValueNativeSourceTypeTablecaption AXValueNativeSourceType = "tablecaption"
    AXValueNativeSourceTypeTitle AXValueNativeSourceType = "title"
    AXValueNativeSourceTypeOther AXValueNativeSourceType = "other"
)

const (
    AXGlobalStatesDisabled AXGlobalStates = "disabled"
    AXGlobalStatesHidden AXGlobalStates = "hidden"
    AXGlobalStatesHiddenRoot AXGlobalStates = "hiddenRoot"
    AXGlobalStatesInvalid AXGlobalStates = "invalid"
    AXGlobalStatesKeyshortcuts AXGlobalStates = "keyshortcuts"
    AXGlobalStatesRoledescription AXGlobalStates = "roledescription"
)

const (
    AXLiveRegionAttributesLive AXLiveRegionAttributes = "live"
    AXLiveRegionAttributesAtomic AXLiveRegionAttributes = "atomic"
    AXLiveRegionAttributesRelevant AXLiveRegionAttributes = "relevant"
    AXLiveRegionAttributesBusy AXLiveRegionAttributes = "busy"
    AXLiveRegionAttributesRoot AXLiveRegionAttributes = "root"
)

const (
    AXWidgetAttributesAutocomplete AXWidgetAttributes = "autocomplete"
    AXWidgetAttributesHaspopup AXWidgetAttributes = "haspopup"
    AXWidgetAttributesLevel AXWidgetAttributes = "level"
    AXWidgetAttributesMultiselectable AXWidgetAttributes = "multiselectable"
    AXWidgetAttributesOrientation AXWidgetAttributes = "orientation"
    AXWidgetAttributesMultiline AXWidgetAttributes = "multiline"
    AXWidgetAttributesReadonly AXWidgetAttributes = "readonly"
    AXWidgetAttributesRequired AXWidgetAttributes = "required"
    AXWidgetAttributesValuemin AXWidgetAttributes = "valuemin"
    AXWidgetAttributesValuemax AXWidgetAttributes = "valuemax"
    AXWidgetAttributesValuetext AXWidgetAttributes = "valuetext"
)

const (
    AXWidgetStatesChecked AXWidgetStates = "checked"
    AXWidgetStatesExpanded AXWidgetStates = "expanded"
    AXWidgetStatesModal AXWidgetStates = "modal"
    AXWidgetStatesPressed AXWidgetStates = "pressed"
    AXWidgetStatesSelected AXWidgetStates = "selected"
)

const (
    AXRelationshipAttributesActivedescendant AXRelationshipAttributes = "activedescendant"
    AXRelationshipAttributesControls AXRelationshipAttributes = "controls"
    AXRelationshipAttributesDescribedby AXRelationshipAttributes = "describedby"
    AXRelationshipAttributesDetails AXRelationshipAttributes = "details"
    AXRelationshipAttributesErrormessage AXRelationshipAttributes = "errormessage"
    AXRelationshipAttributesFlowto AXRelationshipAttributes = "flowto"
    AXRelationshipAttributesLabelledby AXRelationshipAttributes = "labelledby"
    AXRelationshipAttributesOwns AXRelationshipAttributes = "owns"
)
type AXNodeId string

type AXValueType string

type AXValueSourceType string

type AXValueNativeSourceType string

type AXValueSource struct {
    Type AXValueSourceType `json:"type"`// What type of source this is.
    Value *AXValue `json:"value,omitempty"`// The value of this property source.
    Attribute *string `json:"attribute,omitempty"`// The name of the relevant attribute, if any.
    AttributeValue *AXValue `json:"attributeValue,omitempty"`// The value of the relevant attribute, if any.
    Superseded *bool `json:"superseded,omitempty"`// Whether this source is superseded by a higher priority source.
    NativeSource *AXValueNativeSourceType `json:"nativeSource,omitempty"`// The native markup source for this value, e.g. a <label> element.
    NativeSourceValue *AXValue `json:"nativeSourceValue,omitempty"`// The value, such as a node or node list, of the native source.
    Invalid *bool `json:"invalid,omitempty"`// Whether the value for this property is invalid.
    InvalidReason *string `json:"invalidReason,omitempty"`// Reason for the value being invalid, if it is.
}

type AXRelatedNode struct {
    BackendDOMNodeId dom.BackendNodeId `json:"backendDOMNodeId"`// The BackendNodeId of the related DOM node.
    Idref *string `json:"idref,omitempty"`// The IDRef value provided, if any.
    Text *string `json:"text,omitempty"`// The text alternative of this node in the current context.
}

type AXProperty struct {
    Name string `json:"name"`// The name of this property.
    Value AXValue `json:"value"`// The value of this property.
}

type AXValue struct {
    Type AXValueType `json:"type"`// The type of this value.
    Value interface{} `json:"value"`// The computed value of this property.
    RelatedNodes *[]AXRelatedNode `json:"relatedNodes,omitempty"`// One or more related nodes, if applicable.
    Sources *[]AXValueSource `json:"sources,omitempty"`// The sources which contributed to the computation of this property.
}

type AXGlobalStates string

type AXLiveRegionAttributes string

type AXWidgetAttributes string

type AXWidgetStates string

type AXRelationshipAttributes string

type AXNode struct {
    NodeId AXNodeId `json:"nodeId"`// Unique identifier for this node.
    Ignored bool `json:"ignored"`// Whether this node is ignored for accessibility
    IgnoredReasons *[]AXProperty `json:"ignoredReasons,omitempty"`// Collection of reasons why this node is hidden.
    Role *AXValue `json:"role,omitempty"`// This <code>Node</code>'s role, whether explicit or implicit.
    Name *AXValue `json:"name,omitempty"`// The accessible name for this <code>Node</code>.
    Description *AXValue `json:"description,omitempty"`// The accessible description for this <code>Node</code>.
    Value *AXValue `json:"value,omitempty"`// The value for this <code>Node</code>.
    Properties *[]AXProperty `json:"properties,omitempty"`// All other properties
    ChildIds *[]AXNodeId `json:"childIds,omitempty"`// IDs for each of this node's child nodes.
    BackendDOMNodeId *dom.BackendNodeId `json:"backendDOMNodeId,omitempty"`// The backend ID for the associated DOM node, if any.
}
