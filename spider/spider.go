package spider

type Company int

const (
	Cj = iota
	Kpost
	Logen
	Hanjin
	Lotte
)

type Parcel struct {
	Type    Company
	Traking string
}

type Tracker struct {
	Status      int
	TrackingLog struct {
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
