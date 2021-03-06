package prompts

import (
	"github.com/fatih/color"
	prompt "github.com/segmentio/go-prompt"
	"gopkg.in/salsita/go-pivotaltracker.v1/v5/pivotal"
)

// PivotalTracker prompts the user for pivotal tracker credentials
func PivotalTracker() (token string) {

	color.Green("Sign into Pivotal Tracker and visit https://www.pivotaltracker.com/profile")
	color.Green("Then copy and paste your API token below.")
	color.Green("")

	token = prompt.String("API token")
	return token
}

// PivotalTrackerBanner displays a banner for the project init
func PivotalTrackerBanner() {
	color.Cyan("Step #2: Pivotal Tracker Setup")
	color.Cyan("------------------------------")
}

func pivotalTrackerProjectNames(projects []*pivotal.Project) (names []string) {
	names = make([]string, len(projects))

	for i, v := range projects {
		names[i] = v.Name
	}

	return
}

// PivotalTrackerChooseProject prompts the user to choose a project
func PivotalTrackerChooseProject(projects []*pivotal.Project) (proj *pivotal.Project) {
	projectMenu := pivotalTrackerProjectNames(projects)
	selection := prompt.Choose("Choose a project", projectMenu)

	proj = projects[selection]
	return
}
