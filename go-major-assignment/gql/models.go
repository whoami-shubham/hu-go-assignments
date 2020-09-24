package gql

// DbIssue ...
type DbIssue struct {
	ID          int    `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
	Assignee    int    `json:"assignee,omitempty"`
	Reporter    int    `json:"reporter,omitempty"`
	Status      string `json:"status,omitempty"`
	Project     int    `json:"project,omitempty"`
	CreatedOn   string `json:"created_on"`
	UpdatedOn   string `json:"updated_on"`
}

// DbUser ...
type DbUser struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Password  string `json:"password,omitempty"`
	Email     string `json:"email,omitempty"`
	Role      string `json:"role,omitempty"`
	CreatedOn string `json:"created_on"`
	UpdatedOn string `json:"updated_on"`
}

// DbProject ...
type DbProject struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	CreatedBy int    `json:"created_by"`
	CreatedOn string `json:"created_on"`
	UpdatedOn string `json:"updated_on"`
}

// DbComment ...
type DbComment struct {
	ID        int    `json:"id,omitempty"`
	Text      string `json:"text,omitempty"`
	Author    int    `json:"author"`
	IssueID   int    `json:"issue_id"`
	CreatedOn string `json:"created_on"`
	UpdatedOn string `json:"updated_on"`
}

// gqlReturnedIssue ...
type gqlReturnedIssue struct {
	ID          int          `json:"id,omitempty"`
	Title       string       `json:"title,omitempty"`
	Description string       `json:"description,omitempty"`
	Type        string       `json:"type,omitempty"`
	Assignee    DbUser       `json:"assignee,omitempty"`
	Reporter    DbUser       `json:"reporter,omitempty"`
	Status      string       `json:"status,omitempty"`
	Project     DbProject    `json:"project,omitempty"`
	Comments    []DbComment  `json:"comments"`
	Logs        []DbIssueLog `json:"logs"`
}

type gqlReturnedProjectWithIssue struct {
	ID        int                `json:"id,omitempty"`
	Name      string             `json:"name,omitempty"`
	CreatedBy int                `json:"created_by"`
	Issues    []gqlReturnedIssue `json:"issues"`
}

// DbIssueLog ...
type DbIssueLog struct {
	ID            int    `json:"id,omitempty"`
	UpdatedFeild  string `json:"updated_feild,omitempty"`
	PreviousValue string `json:"previous_value,omitempty"`
	NewValue      string `json:"new_value,omitempty"`
	IssueID       int    `json:"issue_id,omitempty"`
	UpdatedOn     string `json:"updated_on"`
}

//CurUser ...
type CurUser struct {
	ID    int
	Email string
	Role  string
}
