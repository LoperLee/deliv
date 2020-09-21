package spider

import "time"

// company code
type Company int

// status code
type Status int

// Company list
const (
	Cj = iota
	Kpost
	Logen
	Hanjin
	Lotte
)

// Parcel status code
const (
	In = iota
	Pickup
	Upper
	Down
	Ready
	Complete
	NotYet
)

type Parcel struct {
	Type    Company
	Traking string
}

type Tracker struct {
	Status      Status
	TrackingLog []struct {
		Date     time.Time
		Location string
		Status   Status
	}
}

// Create new parcel
func New(company Company, tracking string) Parcel {
	return Parcel{
		Type:    company,
		Traking: tracking,
	}
}

// Get Tracking history in website
func (p Parcel) GetTrack() {
	switch p.Type {
	case Cj:
		break
	case Kpost:
		break
	case Logen:
		break
	case Hanjin:
		break
	case Lotte:
		break
	}
}
