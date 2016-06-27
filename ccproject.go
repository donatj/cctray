package ccproject

import "encoding/xml"

type CCProjects struct {
	XMLName  xml.Name     `xml:"Projects"`
	Projects []*CCProject `xml:"Project"`
}

type Activity string

const (
	ActSleeping Activity = "Sleeping"
	ActBuilding Activity = "Building"

	ActCheckingModifications Activity = "CheckingModifications"
)

type Status string

const (
	StatusSuccess   Status = "Success"
	StatusFailure   Status = "Failure"
	StatusException Status = "Exception"
	StatusUnknown   Status = "Unknown"
)

type CCProject struct {
	XMLName         xml.Name `xml:"Project"`
	Name            string   `xml:"name,attr"`
	Activity        Activity `xml:"activity,attr"`
	LastBuildStatus Status   `xml:"lastBuildStatus,attr"`
	LastBuildLabel  string   `xml:"lastBuildLabel,attr,omitempty"`
	LastBuildTime   string   `xml:"lastBuildTime,attr"`
	NextBuildTime   string   `xml:"nextBuildTime,attr,omitempty"`
	WebURL          string   `xml:"webUrl,attr"`
}
