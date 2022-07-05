package data

type Project struct {
	Slug        string
	Name        string
	Description string
}

var (
	Projects = []Project{
		{
			Slug:        "project-a",
			Name:        "Project A",
			Description: "This is the description of Project A",
		}, {
			Slug:        "project-b",
			Name:        "Project B",
			Description: "This is the description of Project B",
		},
	}
)
