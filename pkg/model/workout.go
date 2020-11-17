package model

import "time"

//Workout is the model for how an entire swim workout will look
type Workout struct {
	WarmUp   Set   `json:"warmUp" bson:"warmpUp"`
	PreSet   []Set `json:"preSet" bson:"preSet"`
	MainSet  []Set `json:"mainSet" bson:"mainSet"`
	PostSet  []Set `json:"postSet" bson:"postSet"`
	CoolDown Set   `jons:"coolDown" bson:"coolDown"`
}

// Set is the basic setup of a swim set
type Set struct {
	Amount       int64      `json:"amount" bson:"amount"`
	Distance     int64      `json:"distance" bson:"distance"`
	Stroke       string     `json:"stroke" bson:"stroke"`
	Description  string     `json:"description" bson:"description"`
	Interval     *time.Time `json:"interval" bson:"interval"`
	EnergySystem string     `json:"type" bson:"type"`
	Class        string     `json:"class" bson:"class"`
	Intensity    int64      `json:"intensity" bson:"intensity"`
	Dry          bool       `json:"dry" bson:"dry"`
}
