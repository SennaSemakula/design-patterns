package main


type Github struct {
	accessToken string
	teamId string
}

type Metric struct {
	accessToken string
	deployments int
	changeFailureRate int
	meanTimeToRecover int
	leadTimeChanges int
}

func (t *Github) WithAccessToken(token string) *Github {
	t.accessToken = token
	return t
}

func (t *Github) WithTeam(team string) *Github {
	t.teamId = team
	return t
}



