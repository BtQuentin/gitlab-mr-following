package configuration

type Gitlab struct {
	API     string
	Headers Header
}

type Header struct {
	PrivateToken string
	Accept       string
}

// InitGitlab Define Gitlab details
//
// @param f Flags
// @param e Environment
//
// return Gitlab struct
func InitGitlab(f Flags, e Environment) Gitlab {
	gitlab := Gitlab{
		API: e.API + "api/v4/",
		Headers: Header{
			PrivateToken: f.Token,
			Accept:       "application/json",
		},
	}

	return gitlab
}
