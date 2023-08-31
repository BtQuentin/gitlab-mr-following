package configuration

type Teams struct {
	API     string
	Headers TeamHeader
}

type TeamHeader struct {
	ContentType string
}

func InitTeam(f Flags) Teams {
	teams := Teams{
		API: f.Webhook,
		Headers: TeamHeader{
			ContentType: "application/json",
		},
	}

	return teams
}
