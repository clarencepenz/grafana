package definitions

// swagger:route GET /v1/notifications/time-intervals notifications stable RouteNotificationsGetTimeIntervals
//
// Get all the time intervals
//
//     Responses:
//       200: GetAllIntervalsResponse
//       403: ForbiddenError

// swagger:route GET /v1/notifications/time-intervals/{name} notifications stable RouteNotificationsGetTimeInterval
//
// Get a time interval by name.
//
//     Responses:
//       200: GetIntervalsByNameResponse
//       404: NotFound
//       403: ForbiddenError

// swagger:parameters stable RouteNotificationsGetTimeInterval
type RouteTimeIntervalNameParam struct {
	// Time interval name
	// in:path
	Name string `json:"name"`
}

// swagger:response GetAllIntervalsResponse
type GetAllIntervalsResponse struct {
	// in:body
	Body []GettableTimeIntervals
}

// swagger:response GetIntervalsByNameResponse
type GetIntervalsByNameResponse struct {
	// in:body
	Body GettableTimeIntervals
}

// swagger:model
type PostableTimeIntervals struct {
	Name          string             `json:"name" hcl:"name"`
	TimeIntervals []TimeIntervalItem `json:"time_intervals" hcl:"intervals,block"`
}

type TimeIntervalItem struct {
	Times       []TimeIntervalTimeRange `json:"times,omitempty" hcl:"times,block"`
	Weekdays    *[]string               `json:"weekdays,omitempty" hcl:"weekdays"`
	DaysOfMonth *[]string               `json:"days_of_month,omitempty" hcl:"days_of_month"`
	Months      *[]string               `json:"months,omitempty" hcl:"months"`
	Years       *[]string               `json:"years,omitempty" hcl:"years"`
	Location    *string                 `json:"location,omitempty" hcl:"location"`
}

type TimeIntervalTimeRange struct {
	StartMinute string `json:"start_time" hcl:"start"`
	EndMinute   string `json:"end_time" hcl:"end"`
}

// swagger:model
type GettableTimeIntervals struct {
	Name          string             `json:"name" hcl:"name"`
	TimeIntervals []TimeIntervalItem `json:"time_intervals" hcl:"intervals,block"`
	Provenance    Provenance         `json:"provenance,omitempty"`
}
