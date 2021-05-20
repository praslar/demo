package model

// Please use the body/param/schema field you need for the request
type DemoCreateRequest struct {
	Body   DemoBody
	//Param  DemoParam
	//Scheme DemoScheme
}
type DemoUpdateRequest struct {
	Body   DemoBody
	//Param  DemoParam
	Scheme DemoScheme
}
type DemoGetRequest struct {
	//Body   DemoBody
	Param  DemoParam
	//Scheme DemoScheme
}
type DemoDeleteRequest struct {
	//Body   DemoBody
	//Param  DemoParam
	Scheme DemoScheme
}

// Define your request body here
type DemoBody struct{}

// Define your request param here
type DemoParam struct{}

// Define your request scheme here
type DemoScheme struct{}
