package cap

import (
	"database/sql/driver"
	"strings"

	"github.com/pkg/errors"
)

type Wing uint

const (
	ALWG Wing = 1 + iota
	AZWG
	ARWG
	CAWG
	COWG
	CTWG
	DEWG
	FLWG
	GAWG
	IDWG
	ILWG
	INWG
	IAWG
	KSWG
	KYWG
	LAWG
	MEWG
	MDWG
	MAWG
	MIWG
	MNWG
	MSWG
	MOWG
	MTWG
	DCWG
	NEWG
	NVWG
	NHWG
	NJWG
	NMWG
	NYWG
	NCWG
	NDWG
	OHWG
	OKWG
	ORWG
	PAWG
	RIWG
	SCWG
	SDWG
	TNWG
	TXWG
	UTWG
	VTWG
	VAWG
	WAWG
	WVWG
	WIWG
	WYWG
	AKWG
	HIWG
	PRWG
	NER Wing = 39 + iota
	MAR
	GLR
	SER
	NCR
	SWR
	RMR
	PCR
	NHQ
)

func ParseWing(s string) (Wing, error) {
	switch strings.ToUpper(s) {
	case "AL":
		return ALWG, nil
	case "AZ":
		return AZWG, nil
	case "AR":
		return ARWG, nil
	case "CA":
		return CAWG, nil
	case "CO":
		return COWG, nil
	case "CT":
		return CTWG, nil
	case "DE":
		return DEWG, nil
	case "FL":
		return FLWG, nil
	case "GA":
		return GAWG, nil
	case "ID":
		return IDWG, nil
	case "IL":
		return ILWG, nil
	case "IN":
		return INWG, nil
	case "IA":
		return IAWG, nil
	case "KS":
		return KSWG, nil
	case "KY":
		return KYWG, nil
	case "LA":
		return LAWG, nil
	case "ME":
		return MEWG, nil
	case "MD":
		return MDWG, nil
	case "MA":
		return MAWG, nil
	case "MI":
		return MIWG, nil
	case "MN":
		return MNWG, nil
	case "MS":
		return MSWG, nil
	case "MO":
		return MOWG, nil
	case "MT":
		return MTWG, nil
	case "DC":
		return DCWG, nil
	case "NE":
		return NEWG, nil
	case "NV":
		return NVWG, nil
	case "NH":
		return NHWG, nil
	case "NJ":
		return NJWG, nil
	case "NM":
		return NMWG, nil
	case "NY":
		return NYWG, nil
	case "NC":
		return NCWG, nil
	case "ND":
		return NDWG, nil
	case "OH":
		return OHWG, nil
	case "OK":
		return OKWG, nil
	case "OR":
		return ORWG, nil
	case "PA":
		return PAWG, nil
	case "RI":
		return RIWG, nil
	case "SC":
		return SCWG, nil
	case "SD":
		return SDWG, nil
	case "TN":
		return TNWG, nil
	case "TX":
		return TXWG, nil
	case "UT":
		return UTWG, nil
	case "VT":
		return VTWG, nil
	case "VA":
		return VAWG, nil
	case "WA":
		return WAWG, nil
	case "WV":
		return WVWG, nil
	case "WI":
		return WIWG, nil
	case "WY":
		return WYWG, nil
	case "AK":
		return AKWG, nil
	case "HI":
		return HIWG, nil
	case "PR":
		return PRWG, nil
	case "NER":
		return NER, nil
	case "MAR":
		return MAR, nil
	case "GLR":
		return GLR, nil
	case "SER":
		return SER, nil
	case "NCR":
		return NCR, nil
	case "SWR":
		return SWR, nil
	case "RMR":
		return RMR, nil
	case "PCR":
		return PCR, nil
	case "NHQ":
		return NHQ, nil
	default:
		return 0, errors.Errorf("invalid wing: %s", s)
	}
}

func (w Wing) Name() string {
	switch w {
	case ALWG:
		return "Alabama Wing"
	case AZWG:
		return "Arizona Wing"
	case ARWG:
		return "Arkansas Wing"
	case CAWG:
		return "California Wing"
	case COWG:
		return "Colorado Wing"
	case CTWG:
		return "Connecticut Wing"
	case DEWG:
		return "Delaware Wing"
	case FLWG:
		return "Florida Wing"
	case GAWG:
		return "Georgia Wing"
	case IDWG:
		return "Idaho Wing"
	case ILWG:
		return "Illinois Wing"
	case INWG:
		return "Indiana Wing"
	case IAWG:
		return "Iowa Wing"
	case KSWG:
		return "Kansas Wing"
	case KYWG:
		return "Kentucky Wing"
	case LAWG:
		return "Louisiana Wing"
	case MEWG:
		return "Maine Wing"
	case MDWG:
		return "Maryland Wing"
	case MAWG:
		return "Massachusetts Wing"
	case MIWG:
		return "Michigan Wing"
	case MNWG:
		return "Minnesota Wing"
	case MSWG:
		return "Mississippi Wing"
	case MOWG:
		return "Missouri Wing"
	case MTWG:
		return "Montana Wing"
	case DCWG:
		return "National Capital Wing"
	case NEWG:
		return "Nebraska Wing"
	case NVWG:
		return "Nevada Wing"
	case NHWG:
		return "New Hampshire Wing"
	case NJWG:
		return "New Jersey Wing"
	case NMWG:
		return "New Mexico Wing"
	case NYWG:
		return "New York Wing"
	case NCWG:
		return "North Carolina Wing"
	case NDWG:
		return "North Dakota Wing"
	case OHWG:
		return "Ohio Wing"
	case OKWG:
		return "Oklahoma Wing"
	case ORWG:
		return "Oregon Wing"
	case PAWG:
		return "Pennsylvania Wing"
	case RIWG:
		return "Rhode Island Wing"
	case SCWG:
		return "South Carolina Wing"
	case SDWG:
		return "South Dakota Wing"
	case TNWG:
		return "Tennessee Wing"
	case TXWG:
		return "Texas Wing"
	case UTWG:
		return "Utah Wing"
	case VTWG:
		return "Vermont Wing"
	case VAWG:
		return "Virginia Wing"
	case WAWG:
		return "Washington Wing"
	case WVWG:
		return "West Virginia Wing"
	case WIWG:
		return "Wisconsin Wing"
	case WYWG:
		return "Wyoming Wing"
	case AKWG:
		return "Alaska Wing"
	case HIWG:
		return "Hawaii Wing"
	case PRWG:
		return "Puerto Rico Wing"
	case NER:
		return "North East Region"
	case MAR:
		return "Mid-Atlantic Region"
	case GLR:
		return "Great Lakes Region"
	case SER:
		return "Southeast Region"
	case NCR:
		return "North Central Region"
	case SWR:
		return "Southwest Region"
	case RMR:
		return "Rocky Mountain Region"
	case PCR:
		return "Pacific Region"
	case NHQ:
		return "National Headquarters"
	default:
		return ""
	}
}

func (w Wing) String() string {
	switch w {
	case ALWG:
		return "AL"
	case AZWG:
		return "AZ"
	case ARWG:
		return "AR"
	case CAWG:
		return "CA"
	case COWG:
		return "CO"
	case CTWG:
		return "CT"
	case DEWG:
		return "DE"
	case FLWG:
		return "FL"
	case GAWG:
		return "GA"
	case IDWG:
		return "ID"
	case ILWG:
		return "IL"
	case INWG:
		return "IN"
	case IAWG:
		return "IA"
	case KSWG:
		return "KS"
	case KYWG:
		return "KY"
	case LAWG:
		return "LA"
	case MEWG:
		return "ME"
	case MDWG:
		return "MD"
	case MAWG:
		return "MA"
	case MIWG:
		return "MI"
	case MNWG:
		return "MN"
	case MSWG:
		return "MS"
	case MOWG:
		return "MO"
	case MTWG:
		return "MT"
	case DCWG:
		return "DC"
	case NEWG:
		return "NE"
	case NVWG:
		return "NV"
	case NHWG:
		return "NH"
	case NJWG:
		return "NJ"
	case NMWG:
		return "NM"
	case NYWG:
		return "NY"
	case NCWG:
		return "NC"
	case NDWG:
		return "ND"
	case OHWG:
		return "OH"
	case OKWG:
		return "OK"
	case ORWG:
		return "OR"
	case PAWG:
		return "PA"
	case RIWG:
		return "RI"
	case SCWG:
		return "SC"
	case SDWG:
		return "SD"
	case TNWG:
		return "TN"
	case TXWG:
		return "TX"
	case UTWG:
		return "UT"
	case VTWG:
		return "VT"
	case VAWG:
		return "VA"
	case WAWG:
		return "WA"
	case WVWG:
		return "WV"
	case WIWG:
		return "WI"
	case WYWG:
		return "WY"
	case AKWG:
		return "AK"
	case HIWG:
		return "HI"
	case PRWG:
		return "PR"
	case NER:
		return "NER"
	case MAR:
		return "MAR"
	case GLR:
		return "GLR"
	case SER:
		return "SER"
	case NCR:
		return "NCR"
	case SWR:
		return "SWR"
	case RMR:
		return "RMR"
	case PCR:
		return "PCR"
	case NHQ:
		return "NHQ"
	default:
		return ""
	}
}

func (w Wing) MarshalJSON() ([]byte, error) {
	return []byte(w.String()), nil
}

func (w *Wing) UnmarshalJSON(raw []byte) error {
	val, err := ParseWing(string(raw))
	if err != nil {
		return errors.WithStack(err)
	}

	*w = val
	return nil
}

func (w Wing) Value() (driver.Value, error) {
	return []byte(w.String()), nil
}

func (w *Wing) Scan(raw any) error {
	s, ok := raw.(string)
	if !ok {
		return errors.Errorf("scanned value is not a string: %v", raw)
	}

	val, err := ParseWing(s)
	if err != nil {
		return errors.WithStack(err)
	}

	*w = val
	return nil
}
