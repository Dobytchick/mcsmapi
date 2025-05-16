package mcsmapi

// UserQueryParams — parameters for searching/filtering users for /auth/search
type UserQueryParams struct {
	UserName *string `url:"userName,omitempty"` // Optional: Filter by username (partial match)
	Page     int     `url:"page"`               // Required: Page number (1-based index)
	PageSize int     `url:"page_size"`          // Required: Number of items per page
	Role     *string `url:"role,omitempty"`     // Optional: Filter by permission role ("1", "10", or "-1")
}

func (u *UserQueryParams) BuildQueryString() string {
	return BuildQueryString(u)
}
