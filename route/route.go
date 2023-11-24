package route

type Route string

const (
	GreenLine Route = "green"
	RedLine         = "red"
)

type TravelRoute struct {
	From     string
	To       string
	TripTime string
}
