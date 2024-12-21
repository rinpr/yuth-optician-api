package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Bill struct {
	Id         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Date       string             `json:"date,omitempty" bson:"date,omitempty"`
	EyeData    EyeData            `json:"eye_data,omitempty" bson:"eye_data,omitempty"`
	LensDesign LensDesign         `json:"lens_design,omitempty" bson:"lens_design,omitempty"`
}

type LensDesign struct {
	Description     string `json:"description,omitempty" bson:"description,omitempty"`
	LensStyle       string `json:"lens_style,omitempty" bson:"lens_style,omitempty"`
	LensMaterial    string `json:"lens_material,omitempty" bson:"lens_material,omitempty"`
	LensOption      string `json:"lens_option,omitempty" bson:"lens_option,omitempty"`
	TreatmentOption string `json:"treatment_option,omitempty" bson:"treatment_option,omitempty"`
	Tint            string `json:"tint,omitempty" bson:"tint,omitempty"`
	FrameStyle      string `json:"frame_style,omitempty" bson:"frame_style,omitempty"`
}

type EyeData struct {
	OculusDexter   SpectaclePrescription `json:"right_eye,omitempty" bson:"right_eye,omitempty"`
	OculusSinister SpectaclePrescription `json:"left_eye,omitempty" bson:"left_eye,omitempty"`
	OculusUterque  string                `json:"both_eye,omitempty" bson:"both_eye,omitempty"`
	VertexDistance string                `json:"vertex_distance,omitempty" bson:"vertex_distance,omitempty"`
}

type SpectaclePrescription struct {
	Sphere   float32 `json:"sphere,omitempty" bson:"sphere,omitempty"`
	Cylinder float32 `json:"cylinder,omitempty" bson:"cylinder,omitempty"`
	Axis     int     `json:"axis,omitempty" bson:"axis,omitempty"`
	Prism    float32 `json:"prism,omitempty" bson:"prism,omitempty"`
	Add      float32 `json:"add,omitempty" bson:"add,omitempty"`
	MonoVa   string  `json:"mono_va,omitempty" bson:"mono_va,omitempty"`
	DPD      int     `json:"dpd,omitempty" bson:"dpd,omitempty"`
	NPD      int     `json:"npd,omitempty" bson:"npd,omitempty"`
	FH       int     `json:"fh,omitempty" bson:"fh,omitempty"`
}
