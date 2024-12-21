package models

type BillV1 struct {
	Id         int 			`gorm:"type:int;primary_key"`
	Date       string       `gorm:"type:varchar(255)"`
	EyeData    EyeDataV1      `gorm:"type:varchar(255)"`
	LensDesign LensDesignV1   `gorm:"type:varchar(255)"`
}

type LensDesignV1 struct {
	Description     string `json:"description`
	LensStyle       string `json:"lens_style"`
	LensMaterial    string `json:"lens_material"`
	LensOption      string `json:"lens_option"`
	TreatmentOption string `json:"treatment_option"`
	Tint            string `json:"tint"`
	FrameStyle      string `json:"frame_style"`
}

type EyeDataV1 struct {
	OculusDexter   SpectaclePrescriptionV1 `json:"right_eye"`
	OculusSinister SpectaclePrescriptionV1 `json:"left_eye"`
	OculusUterque  string                `json:"both_eye"`
	VertexDistance string                `json:"vertex_distance"`
}

type SpectaclePrescriptionV1 struct {
	Sphere   float32 `json:"sphere"`
	Cylinder float32 `json:"cylinder"`
	Axis     int     `json:"axis"`
	Prism    float32 `json:"prism"`
	Add      float32 `json:"add"`
	MonoVa   string  `json:"mono_va"`
	DPD      int     `json:"dpd"`
	NPD      int     `json:"npd"`
	FH       int     `json:"fh"`
}
