package types

type ScreenOrientationTypeEnum string
const (
    ScreenOrientationTypePortraitPrimary ScreenOrientationTypeEnum = "portraitPrimary"
    ScreenOrientationTypePortraitSecondary ScreenOrientationTypeEnum = "portraitSecondary"
    ScreenOrientationTypeLandscapePrimary ScreenOrientationTypeEnum = "landscapePrimary"
    ScreenOrientationTypeLandscapeSecondary ScreenOrientationTypeEnum = "landscapeSecondary"
)

const (
    VirtualTimePolicyAdvance VirtualTimePolicy = "advance"
    VirtualTimePolicyPause VirtualTimePolicy = "pause"
    VirtualTimePolicyPauseIfNetworkFetchesPending VirtualTimePolicy = "pauseIfNetworkFetchesPending"
)

type SetTouchEmulationEnabledConfigurationEnum string
const (
    SetTouchEmulationEnabledConfigurationMobile SetTouchEmulationEnabledConfigurationEnum = "mobile"
    SetTouchEmulationEnabledConfigurationDesktop SetTouchEmulationEnabledConfigurationEnum = "desktop"
)
type ScreenOrientation struct {
    Type ScreenOrientationTypeEnum `json:"type"`// Orientation type.
    Angle int64 `json:"angle"`// Orientation angle.
}

type VirtualTimePolicy string
