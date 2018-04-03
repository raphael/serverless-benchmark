package design

import (
	. "goa.design/goa/http/design"
	. "goa.design/goa/http/dsl"
)

var _ = API("recorder", func() {
	Title("Recorder API")
	Description("Recorder records arbitrary data together with the recording timestamp")
})

var _ = Service("recorder", func() {
	Method("record-data", func() {
		Description("RecordData creates a new datapoint.")
		Payload(Datapoint)
		Result(Empty)
		HTTP(func() {
			POST("/data")
			Response(StatusOK)
		})
	})
	Method("list", func() {
		Description("List lists all recorded datapoints.")
		Payload(Series)
		Result(ArrayOf(Float64))
		HTTP(func() {
			GET("/data")
			Param("service")
			Param("name")
			Response(StatusOK)
		})
	})
})

var Series = Type("Series", func() {
	Description("Series represent a time series.")
	Attribute("service", String, "Service that created datapoint.", func() {
		Example("lambda")
	})
	Attribute("name", String, "Name is the name of the datapoint.", func() {
		Example("duration")
	})
	Required("service", "name")
})

var Datapoint = Type("Datapoint", func() {
	Description("Datapoint describes a single recording datapoint.")
	Attribute("service", String, "Service that created datapoint.", func() {
		Example("lambda")
	})
	Attribute("value", Float64, "Datapoint value.")
	Attribute("name", String, "Name is the name of the datapoint.", func() {
		Example("duration")
	})
	Required("service", "name", "value")
})
