package systeminfo

type GPUDevice struct {
    VendorId float64 `json:"vendorId"`// PCI ID of the GPU vendor, if available; 0 otherwise.
    DeviceId float64 `json:"deviceId"`// PCI ID of the GPU device, if available; 0 otherwise.
    VendorString string `json:"vendorString"`// String description of the GPU vendor, if the PCI ID is not available.
    DeviceString string `json:"deviceString"`// String description of the GPU device, if the PCI ID is not available.
}

type GPUInfo struct {
    Devices []GPUDevice `json:"devices"`// The graphics devices on the system. Element 0 is the primary GPU.
    AuxAttributes *map[string]string `json:"auxAttributes,omitempty"`// An optional dictionary of additional GPU related attributes.
    FeatureStatus *map[string]string `json:"featureStatus,omitempty"`// An optional dictionary of graphics features and their status.
    DriverBugWorkarounds []string `json:"driverBugWorkarounds"`// An optional array of GPU driver bug workarounds.
}
