// Package cctray implements the CCTray v1 Specification Standard.
//
// The standard is currently defined at https://cctray.org/v1/
package cctray

import (
	"encoding/xml"
	"time"
)

// Activity is the type for the Activity response attribute
type Activity string

const (
	// ActSleeping indicates activity state is "Sleeping"
	ActSleeping Activity = "Sleeping"
	// ActBuilding indicates activity state is "Building"
	ActBuilding Activity = "Building"
	// ActCheckingModifications indicates activity state is "CheckingModifications"
	ActCheckingModifications Activity = "CheckingModifications"
)

// Status is the type for the LastBuildStatus response attribute
type Status string

const (
	// StatusSuccess indicates last build state was successful
	StatusSuccess Status = "Success"
	// StatusFailure indicates last build state was an expected failure
	StatusFailure Status = "Failure"
	// StatusException indicates the last build was an unexpected failure
	StatusException Status = "Exception"
	// StatusUnknown indicates last build state is not known
	StatusUnknown Status = "Unknown"
)

// CCProjects represents the base XML <Projects> node which contains projects
type CCProjects struct {
	XMLName  xml.Name     `xml:"Projects"`
	Projects []*CCProject `xml:"Project"`
}

// CCProject represents a <Project> node
type CCProject struct {
	XMLName         xml.Name   `xml:"Project"`
	Name            string     `xml:"name,attr"`
	Activity        Activity   `xml:"activity,attr"`
	LastBuildStatus Status     `xml:"lastBuildStatus,attr"`
	LastBuildLabel  string     `xml:"lastBuildLabel,attr,omitempty"`
	LastBuildTime   time.Time  `xml:"lastBuildTime,attr"`
	NextBuildTime   *time.Time `xml:"nextBuildTime,attr,omitempty"`
	WebURL          string     `xml:"webUrl,attr"`
}
