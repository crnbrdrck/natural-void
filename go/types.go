package naturalvoid

// DB Models
// Model representing a User in the DB
// Users come from LDAP so this just stores some information for mapping stories to their DMs
type User struct {
	ID int
}

// A story is a collection of episodes
// ASTTW Chapter 1 is a story, as well as Delve Into Delirium
type Story struct {
	Description []string
	DM          User
	ID          int
	Name        string
	ShortName   string
	Slug        string
}

// An Episode relates to a recorded file.
// Episodes are mapped to a Story for obvious reasons.
type Episode struct {
	Description []string
	Name        string
	Number      int
	Path        string
	Story       Story
}
