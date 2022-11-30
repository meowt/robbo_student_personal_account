// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

type AbsoluteMediaHTTP struct {
	ID          string `json:"ID"`
	URI         string `json:"URI"`
	URIAbsolute string `json:"URI_Absolute"`
}

type CourseAPIMediaCollectionHTTP struct {
	ID          string             `json:"ID"`
	BannerImage *AbsoluteMediaHTTP `json:"Banner_Image"`
	CourseImage *MediaHTTP         `json:"Course_Image"`
	CourseVideo *MediaHTTP         `json:"Course_Video"`
	Image       *ImageHTTP         `json:"Image"`
}

type CourseHTTP struct {
	ID               string                        `json:"ID"`
	BlocksURL        string                        `json:"Blocks_URL"`
	Effort           string                        `json:"Effort"`
	EnrollmentStart  string                        `json:"Enrollment_Start"`
	EnrollmentEnd    string                        `json:"Enrollment_End"`
	End              string                        `json:"End"`
	Name             string                        `json:"Name"`
	Number           string                        `json:"Number"`
	Org              string                        `json:"Org"`
	ShortDescription string                        `json:"Short_Description"`
	Start            string                        `json:"Start"`
	StartDisplay     string                        `json:"Start_Display"`
	StartType        string                        `json:"Start_Type"`
	Pacing           string                        `json:"Pacing"`
	MobileAvailable  bool                          `json:"Mobile_Available"`
	Hidden           bool                          `json:"Hidden"`
	InvitationOnly   bool                          `json:"Invitation_Only"`
	Overview         *string                       `json:"Overview"`
	CourseID         string                        `json:"Course_ID"`
	Media            *CourseAPIMediaCollectionHTTP `json:"Media"`
}

type CoursesListHTTP struct {
	Results    []*CourseHTTP `json:"Results"`
	Pagination *Pagination   `json:"Pagination"`
}

type EnrollmentHTTP struct {
	Created  string `json:"Created"`
	Mode     string `json:"Mode"`
	IsActive bool   `json:"IsActive"`
	User     string `json:"User"`
	CourseID string `json:"Course_ID"`
}

type EnrollmentsListHTTP struct {
	Next     string            `json:"Next"`
	Previous string            `json:"Previous"`
	Results  []*EnrollmentHTTP `json:"Results"`
}

type ImageHTTP struct {
	ID    string `json:"ID"`
	Raw   string `json:"Raw"`
	Small string `json:"Small"`
	Large string `json:"Large"`
}

type MediaHTTP struct {
	ID  string `json:"ID"`
	URI string `json:"URI"`
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
	Next     string `json:"Next"`
	Previous string `json:"Previous"`
	Count    int    `json:"Count"`
	NumPages int    `json:"Num_Pages"`
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

type RobboGroupHTTP struct {
	ID           string         `json:"id"`
	LastModified string         `json:"lastModified"`
	Name         string         `json:"name"`
	RobboUnitID  string         `json:"robboUnitId"`
	Students     []*StudentHTTP `json:"students"`
}

type RobboUnitHTTP struct {
	ID           string `json:"id"`
	LastModified string `json:"lastModified"`
	Name         string `json:"name"`
	City         string `json:"city"`
}

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
	ProjectID   string `json:"ProjectID"`
	Instruction string `json:"Instruction"`
	Notes       string `json:"Notes"`
	Preview     string `json:"Preview"`
	LinkScratch string `json:"LinkScratch"`
	Title       string `json:"Title"`
	IsShared    bool   `json:"IsShared"`
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
