// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

type CourseResult interface {
	IsCourseResult()
}

type EnrollmentResult interface {
	IsEnrollmentResult()
}

type ProjectPageResult interface {
	IsProjectPageResult()
}

type RobboGroupResult interface {
	IsRobboGroupResult()
}

type RobboUnitResult interface {
	IsRobboUnitResult()
}

type AbsoluteMediaHTTP struct {
	ID          string `json:"id"`
	URI         string `json:"uri"`
	URIAbsolute string `json:"uri_absolute"`
}

type CourseAPIMediaCollectionHTTP struct {
	ID          string             `json:"id"`
	BannerImage *AbsoluteMediaHTTP `json:"banner_image"`
	CourseImage *MediaHTTP         `json:"course_image"`
	CourseVideo *MediaHTTP         `json:"course_video"`
	Image       *ImageHTTP         `json:"image"`
}

type CourseHTTP struct {
	ID               string                        `json:"id"`
	BlocksURL        string                        `json:"blocks_url"`
	Effort           string                        `json:"effort"`
	EnrollmentStart  string                        `json:"enrollment_start"`
	EnrollmentEnd    string                        `json:"enrollment_end"`
	End              string                        `json:"end"`
	Name             string                        `json:"name"`
	Number           string                        `json:"number"`
	Org              string                        `json:"org"`
	ShortDescription string                        `json:"short_description"`
	Start            string                        `json:"start"`
	StartDisplay     string                        `json:"start_display"`
	StartType        string                        `json:"start_type"`
	Pacing           string                        `json:"pacing"`
	MobileAvailable  bool                          `json:"mobile_available"`
	Hidden           bool                          `json:"hidden"`
	InvitationOnly   bool                          `json:"invitation_only"`
	Overview         *string                       `json:"overview"`
	CourseID         string                        `json:"course_id"`
	Media            *CourseAPIMediaCollectionHTTP `json:"media"`
}

func (CourseHTTP) IsCourseResult() {}

type CoursesListHTTP struct {
	Results    []*CourseHTTP `json:"results"`
	Pagination *Pagination   `json:"pagination"`
}

func (CoursesListHTTP) IsCourseResult() {}

type DeletedProjectPage struct {
	ProjectPageID string `json:"projectPageId"`
}

type DeletedRobboGroup struct {
	RobboGroupID string `json:"robboGroupId"`
}

type DeletedRobboUnit struct {
	RobboUnitID string `json:"robboUnitId"`
}

type EnrollmentHTTP struct {
	Created  string `json:"created"`
	Mode     string `json:"mode"`
	IsActive bool   `json:"isActive"`
	User     string `json:"user"`
	CourseID string `json:"course_id"`
}

type EnrollmentsListHTTP struct {
	Next     string            `json:"next"`
	Previous string            `json:"previous"`
	Results  []*EnrollmentHTTP `json:"results"`
}

func (EnrollmentsListHTTP) IsEnrollmentResult() {}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (Error) IsCourseResult() {}

func (Error) IsEnrollmentResult() {}

func (Error) IsProjectPageResult() {}

func (Error) IsRobboGroupResult() {}

func (Error) IsRobboUnitResult() {}

type ImageHTTP struct {
	ID    string `json:"id"`
	Raw   string `json:"raw"`
	Small string `json:"small"`
	Large string `json:"large"`
}

type MediaHTTP struct {
	ID  string `json:"id"`
	URI string `json:"uri"`
}

type NewParent struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	Nickname   string `json:"nickname"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Middlename string `json:"middlename"`
}

type NewRobboGroup struct {
	Name        string `json:"name"`
	RobboUnitID string `json:"robboUnitId"`
}

type NewRobboUnit struct {
	Name string `json:"name"`
	City string `json:"city"`
}

type NewStudent struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	Nickname   string `json:"nickname"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Middlename string `json:"middlename"`
	ParentID   string `json:"parentId"`
}

type NewTeacher struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	Nickname   string `json:"nickname"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Middlename string `json:"middlename"`
}

type NewUnitAdmin struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	Nickname   string `json:"nickname"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Middlename string `json:"middlename"`
}

type Pagination struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Count    int    `json:"count"`
	NumPages int    `json:"num_pages"`
}

type ParentHTTP struct {
	UserHTTP *UserHTTP      `json:"userHttp"`
	Children []*StudentHTTP `json:"children"`
}

type ProjectPageHTTP struct {
	ProjectPageID string `json:"projectPageId"`
	LastModified  string `json:"lastModified"`
	ProjectID     string `json:"projectId"`
	Instruction   string `json:"instruction"`
	Notes         string `json:"notes"`
	Preview       string `json:"preview"`
	LinkScratch   string `json:"linkScratch"`
	Title         string `json:"title"`
	IsShared      bool   `json:"isShared"`
}

func (ProjectPageHTTP) IsProjectPageResult() {}

type ProjectPageHTTPList struct {
	ProjectPages []*ProjectPageHTTP `json:"projectPages"`
}

func (ProjectPageHTTPList) IsProjectPageResult() {}

type RobboGroupHTTP struct {
	ID           string         `json:"id"`
	LastModified string         `json:"lastModified"`
	Name         string         `json:"name"`
	RobboUnitID  string         `json:"robboUnitId"`
	Students     []*StudentHTTP `json:"students"`
}

func (RobboGroupHTTP) IsRobboGroupResult() {}

type RobboGroupHTTPList struct {
	RobboGroups []*RobboGroupHTTP `json:"robboGroups"`
}

func (RobboGroupHTTPList) IsRobboGroupResult() {}

type RobboUnitHTTP struct {
	ID           string `json:"id"`
	LastModified string `json:"lastModified"`
	Name         string `json:"name"`
	City         string `json:"city"`
}

func (RobboUnitHTTP) IsRobboUnitResult() {}

type RobboUnitHTTPList struct {
	RobboUnits []*RobboUnitHTTP `json:"robboUnits"`
}

func (RobboUnitHTTPList) IsRobboUnitResult() {}

type StudentHTTP struct {
	UserHTTP     *UserHTTP `json:"userHttp"`
	RobboGroupID string    `json:"robboGroupId"`
	RobboUnitID  string    `json:"robboUnitId"`
}

type SuperAdminHTTP struct {
	UserHTTP *UserHTTP `json:"userHttp"`
}

type TeacherHTTP struct {
	UserHTTP *UserHTTP `json:"userHttp"`
}

type UnitAdminHTTP struct {
	UserHTTP *UserHTTP `json:"userHttp"`
}

type UpdateParentHTTP struct {
	UserHTTP *UpdateUserHTTP `json:"userHttp"`
}

type UpdateParentInput struct {
	ParentHTTP *UpdateParentHTTP `json:"parentHttp"`
}

type UpdateProjectPage struct {
	ProjectID     string `json:"projectId"`
	ProjectPageID string `json:"projectPageId"`
	Instruction   string `json:"instruction"`
	Notes         string `json:"notes"`
	Title         string `json:"title"`
	IsShared      bool   `json:"isShared"`
}

type UpdateRobboGroup struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	RobboUnitID string `json:"robboUnitId"`
}

type UpdateRobboUnit struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

type UpdateStudentHTTP struct {
	UserHTTP *UpdateUserHTTP `json:"userHttp"`
}

type UpdateStudentInput struct {
	StudentHTTP *UpdateStudentHTTP `json:"studentHttp"`
}

type UpdateSuperAdminHTTP struct {
	UserHTTP *UpdateUserHTTP `json:"userHttp"`
}

type UpdateSuperAdminInput struct {
	SuperAdminHTTP *UpdateSuperAdminHTTP `json:"superAdminHttp"`
}

type UpdateTeacherHTTP struct {
	UserHTTP *UpdateUserHTTP `json:"userHttp"`
}

type UpdateTeacherInput struct {
	TeacherHTTP *UpdateTeacherHTTP `json:"teacherHttp"`
}

type UpdateUnitAdminHTTP struct {
	UserHTTP *UpdateUserHTTP `json:"userHttp"`
}

type UpdateUnitAdminInput struct {
	UnitAdminHTTP *UpdateUnitAdminHTTP `json:"unitAdminHttp"`
}

type UpdateUserHTTP struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	Nickname   string `json:"nickname"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Middlename string `json:"middlename"`
}

type UserHTTP struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Role       int    `json:"role"`
	Nickname   string `json:"nickname"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Middlename string `json:"middlename"`
	CreatedAt  string `json:"createdAt"`
}
