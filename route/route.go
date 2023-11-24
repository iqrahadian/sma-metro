package route

type Route string

const (
	GreenLine Route = "Green"
	RedLine         = "Red"
)

type TravelRoute struct {
	From     Route
	To       Route
	TripTime string
}
