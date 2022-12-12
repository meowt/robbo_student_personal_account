package users

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
)

type Delegate interface {
	GetIndividualStudentsByTeacherId(teacherId string) (students []*models.StudentHTTP, err error)
	SearchStudentByEmail(email string, parentId string) (students []*models.StudentHTTP, err error)
	CreateStudent(student *models.StudentHTTP, parentId string) (newStudent *models.StudentHTTP, err error)
	UpdateStudent(student *models.StudentHTTP) (studentUpdated *models.StudentHTTP, err error)
	DeleteStudent(studentId string) (err error)
	GetStudentById(studentId string) (student *models.StudentHTTP, err error)
	GetStudentsByRobboGroupId(robboGroupId string) (students []*models.StudentHTTP, err error)
	GetStudentsByRobboUnitId(robboUnitId string) (students []*models.StudentHTTP, err error)
	GetStudentsByTeacherId(teacherId string) (students []*models.StudentHTTP, err error)
	GetStudentByParentId(parentId string) (students []*models.StudentHTTP, err error)
	AddStudentToRobboGroup(studentId string, robboGroupId string, robboUnitId string) (err error)

	GetPairsStudentParentsByTeacherId(teacherId string) (pairsStudentParents []*models.StudentParentsHTTP, err error)

	CreateStudentTeacherRelation(studentId, teacherId string) (student *models.StudentHTTP, err error)
	DeleteStudentTeacherRelation(studentId, teacherId string) (student *models.StudentHTTP, err error)

	//GetTeacher(email, password string) (teacher models.TeacherHTTP, err error)
	GetTeacherById(teacherId string) (teacher *models.TeacherHTTP, err error)
	GetAllTeachers() (teachers []*models.TeacherHTTP, err error)
	CreateTeacher(teacher *models.TeacherHTTP) (newTeacher models.TeacherHTTP, err error)
	UpdateTeacher(teacherHTTP *models.TeacherHTTP) (teacherUpdated models.TeacherHTTP, err error)
	DeleteTeacher(teacherId string) (err error)
	GetTeacherByRobboGroupId(robboGroupId string) (teachers []*models.TeacherHTTP, err error)
	GetTeachersByStudentId(studentId string) (teachers []*models.TeacherHTTP, err error)

	CreateParent(parentHTTP *models.ParentHTTP) (newParent *models.ParentHTTP, err error)
	UpdateParent(parentHTTP *models.ParentHTTP) (parentUpdated *models.ParentHTTP, err error)
	DeleteParent(parentId string) (err error)
	GetParentById(parentId string) (parent *models.ParentHTTP, err error)
	GetAllParent() (parents []*models.ParentHTTP, err error)
	//GetParent(email, password string) (parent models.ParentHTTP, err error)

	CreateFreeListener(freeListenerHTTP *models.FreeListenerHttp) (newFreeListener *models.FreeListenerHttp, err error)
	UpdateFreeListener(freeListenerHTTP *models.FreeListenerHttp) (freeListenerUpdated *models.FreeListenerHttp, err error)
	DeleteFreeListener(freeListenerId string) (err error)
	GetFreeListenerById(freeListenerId string) (freeListener models.FreeListenerHttp, err error)
	//GetFreeListener(email, password string) (freeListener models.FreeListenerHttp, err error)

	CreateUnitAdmin(unitAdminHTTP *models.UnitAdminHTTP) (newUnitAdmin *models.UnitAdminHTTP, err error)
	UpdateUnitAdmin(unitAdminHTTP *models.UnitAdminHTTP) (unitAdminUpdated *models.UnitAdminHTTP, err error)
	DeleteUnitAdmin(unitAdminId string) (err error)
	GetUnitAdminById(unitAdminId string) (unitAdmin models.UnitAdminHTTP, err error)
	GetAllUnitAdmins() (unitAdmins []*models.UnitAdminHTTP, err error)
	SearchUnitAdminByEmail(email string, robboUnitId string) (unitAdmins []*models.UnitAdminHTTP, err error)
	GetUnitAdminByRobboUnitId(robboUnitId string) (unitAdmins []*models.UnitAdminHTTP, err error)

	//GetSuperAdmin(email, password string) (superAdmin models.SuperAdminHTTP, err error)
	GetSuperAdminById(superAdminId string) (superAdmin models.SuperAdminHTTP, err error)
	UpdateSuperAdmin(superAdminHTTP *models.SuperAdminHTTP) (superAdminUpdated *models.SuperAdminHTTP, err error)
	DeleteSuperAdmin(superAdminId string) (err error)

	CreateRelation(parentId, childrenId string) (err error)
	SetNewUnitAdminForRobboUnit(unitAdminId, robboUnitId string) (err error)
	DeleteUnitAdminForRobboUnit(unitAdminId, robboUnitId string) (err error)
}
