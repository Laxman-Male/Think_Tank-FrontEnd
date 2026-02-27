package allqueries

const (
	GetProfileDetails = `select email, role , mobile from users where id=?`
)
