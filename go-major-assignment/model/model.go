package model

// Issue ...
type Issue struct {
	Title       string `db:"title", json:"title,omitempty"`
	Description string `db:"description", json:"description,omitempty"`
	Type        string `db:"type", json:"type,omitempty"`
	Assignee    int    `db:"assignee", json:"assignee,omitempty"`
	Reporter    int    `db:"reporter", json:"reporter,omitempty"`
	Status      string `db:"status", json:"status,omitempty"`
	Project     int    `db:"project", json:"project,omitempty"`
}

// User ...
type User struct {
	Name     string `db:"name", json:"name,omitempty"`
	Password string `db:"password", json:"password,omitempty"`
	Email    string `db:"email", json:"email,omitempty"`
	Role     string `db:"role", json:"role,omitempty"`
}

// Project ...
type Project struct {
	Name      string `db:"name" json:"name,omitempty"`
	CreatedBy int    `db:"created_by" json:"created_by"`
}

// Report ...
type Report struct {
	Email       string `json:"email,omitempty"` // assignee's email
	ID          int32  `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Reporter    string `json:"reporter,omitempty"`
	CreatedOn   string `json:"created_on"`
	UpdatedOn   string `json:"updated_on"`
}

/*

structs for testing

*/

//TestUser ...
type TestUser struct {
	ID       int    `db:"id",json:"id,omitempty"`
	Name     string `db:"name", json:"name,omitempty"`
	Password string `db:"password", json:"password,omitempty"`
	Email    string `db:"email", json:"email,omitempty"`
	Role     string `db:"role", json:"role,omitempty"`
}

// TestIssue ...
type TestIssue struct {
	ID          int    `db:"id",json:"id,omitempty"`
	Title       string `db:"title", json:"title,omitempty"`
	Description string `db:"description", json:"description,omitempty"`
	Type        string `db:"type", json:"type,omitempty"`
	Assignee    int    `db:"assignee", json:"assignee,omitempty"`
	Reporter    int    `db:"reporter", json:"reporter,omitempty"`
	Status      string `db:"status", json:"status,omitempty"`
	Project     int    `db:"project", json:"project,omitempty"`
}
