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
		}, {
			Slug:        "project-c",
			Name:        "Project C",
			Description: "This is the description of Project C",
		}, {
			Slug:        "project-d",
			Name:        "Project D",
			Description: "This is the description of Project D",
		},
	}
)
