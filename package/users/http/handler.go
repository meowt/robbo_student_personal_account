package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/users"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	authDelegate  auth.Delegate
	usersDelegate users.Delegate
}

func NewUsersHandler(
	authDelegate auth.Delegate,
	usersDelegate users.Delegate,
) Handler {
	return Handler{
		authDelegate:  authDelegate,
		usersDelegate: usersDelegate,
	}
}

type getParentResponse struct {
	Parent models.ParentHTTP `json:"parent"`
}

type getTeacherResponse struct {
	Teacher models.TeacherHTTP `json:"teacher"`
}

type getUnitAdminResponse struct {
	UnitAdmin models.UnitAdminHTTP `json:"unitAdmin"`
}

type getStudentResponse struct {
	Student models.StudentHTTP `json:"student"`
}

type getSuperAdminResponse struct {
	SuperAdmin models.SuperAdminHTTP `json:"superAdmin"`
}

type getFreeListener struct {
	FreeListener models.FreeListenerHttp `json:"freeListener"`
}

func (h *Handler) InitUsersRoutes(router *gin.Engine) {
	users := router.Group("/users")
	{
		users.GET("/", h.GetUserByAccessToken)
		users.GET("/:userId", h.GetUserById)
		// TODO refactor
		users.POST("/student", h.CreateStudent)
		users.DELETE("/student/:studentId", h.DeleteStudent)
		users.GET("/student/:studentId", h.GetStudentById)
		users.GET("/student/search/:parentId/:studentEmail", h.SearchStudentByEmail)
		users.GET("/students/:parentId", h.GetStudentByParentId)
		users.PUT("/student", h.UpdateStudent)
		users.POST("/student/:studentId/robboGroup/:robboGroupId", h.SetRobboGroupIdForStudent)
		//users.DELETE("/student/:studentId/robboGroup/:robboGroupId", h.SetRobboGroupIdForStudent)

		users.POST("/teacher", h.CreateTeacher)
		users.GET("/teachers", h.GetAllTeachers)
		users.DELETE("/teacher/:teacherId", h.DeleteTeacher)
		users.PUT("/teacher", h.UpdateTeacher)
		users.GET("/teacher/:teacherId", h.GetTeacherById)

		users.POST("/parent", h.CreateParent)
		users.GET("/parent/:parentId", h.GetParentById)
		users.GET("/parents", h.GetAllParent)
		users.PUT("/parent", h.UpdateParent)
		users.DELETE("/parent/:parentId", h.DeleteParent)

		users.POST("/freeListener", h.CreateFreeListener)
		users.DELETE("/freeListener/:freeListenerId", h.DeleteFreeListener)
		users.PUT("/freeListener", h.UpdateFreeListener)
		users.GET("/freeListener/:freeListenerId", h.GetFreeListenerById)

		users.POST("/unitAdmin", h.CreateUnitAdmin)
		users.DELETE("/unitAdmin/:unitAdminId", h.DeleteUnitAdmin)
		users.PUT("/unitAdmin", h.UpdateUnitAdmin)
		users.GET("/unitAdmin/:unitAdminId", h.GetUnitAdminByID)
		users.GET("/unitAdmins/:robboUnitId", h.GetUnitAdminsByRobboUnitId)
		users.GET("/unitAdmins", h.GetAllUnitAdmins)
		users.GET("/unitAdmin/search/:unitAdminEmail", h.SearchUnitAdminByEmail)
		users.POST("/unitAdmin/setRelation", h.SetNewUnitAdminForRobboUnit)
		users.POST("/unitAdmin/deleteRelation", h.DeleteUnitAdminForRobboUnit)

		users.GET("/superAdmin/:superAdminId", h.GetSuperAdminById)
		users.PUT("/superAdmin", h.UpdateSuperAdmin)
		users.DELETE("/superAdmin/:superAdminId", h.DeleteSuperAdmin)

		// TODO rename
		users.POST("/relation", h.CreateRelation)
	}
}

func (h *Handler) GetUserByAccessToken(c *gin.Context) {
	log.Println("Get User by access token")
	userId, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}

	switch role {
	case models.Student:
		student, getStudentErr := h.usersDelegate.GetStudentById(userId)
		if getStudentErr != nil {
			ErrorHandling(getStudentErr, c)
			return
		}
		c.JSON(http.StatusOK, student)
	case models.Teacher:
		teacher, getTeacherErr := h.usersDelegate.GetTeacherById(userId)
		if getTeacherErr != nil {
			ErrorHandling(getTeacherErr, c)
			return
		}
		c.JSON(http.StatusOK, teacher)
		return
	case models.Parent:
		parent, getParentErr := h.usersDelegate.GetParentById(userId)
		if getParentErr != nil {
			ErrorHandling(getParentErr, c)
			return
		}
		c.JSON(http.StatusOK, parent)
		return
	case models.FreeListener:
		freeListener, getFreeListenerErr := h.usersDelegate.GetFreeListenerById(userId)
		if getFreeListenerErr != nil {
			ErrorHandling(getFreeListenerErr, c)
			return
		}
		c.JSON(http.StatusOK, freeListener)
	case models.UnitAdmin:
		unitAdmin, getUnitAdminErr := h.usersDelegate.GetUnitAdminById(userId)
		if getUnitAdminErr != nil {
			ErrorHandling(getUnitAdminErr, c)
			return
		}
		c.JSON(http.StatusOK, unitAdmin)
		return
	case models.SuperAdmin:
		superAdmin, getSuperAdminErr := h.usersDelegate.GetSuperAdminById(userId)
		if getSuperAdminErr != nil {
			ErrorHandling(getSuperAdminErr, c)
			return
		}
		c.JSON(http.StatusOK, superAdmin)
		return
	}
}

func (h *Handler) SearchStudentByEmail(c *gin.Context) {
	log.Println("Get Student By Email")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}
	studentEmail := c.Param("studentEmail")
	parentId := c.Param("parentId")
	students, err := h.usersDelegate.SearchStudentByEmail(studentEmail, parentId)
	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}
	c.JSON(http.StatusOK, students)
}

func (h *Handler) GetStudentById(c *gin.Context) {
	log.Println("Get Student By Id")
	_, _, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}

	studentId := c.Param("studentId")

	student, err := h.usersDelegate.GetStudentById(studentId)

	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	c.JSON(http.StatusOK, getStudentResponse{
		*student,
	})
}

func (h *Handler) GetStudentByParentId(c *gin.Context) {
	log.Println("Get Student By Parent Id")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}

	parentId := c.Param("parentId")
	students, err := h.usersDelegate.GetStudentByParentId(parentId)

	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	c.JSON(http.StatusOK, students)
}

type createStudentInput struct {
	Student  *models.StudentHTTP `json:"student"`
	ParentId string              `json:"parentId"`
}

func (h *Handler) CreateStudent(c *gin.Context) {
	log.Println("Create Student")

	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}
	input := new(createStudentInput)

	if err := c.BindJSON(input); err != nil {
		err = users.ErrBadRequestBody
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	fmt.Println(input.ParentId)

	studentId, err := h.usersDelegate.CreateStudent(input.Student, input.ParentId)
	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"studentId": studentId,
	})
}

func (h *Handler) UpdateStudent(c *gin.Context) {
	log.Println("Update Student")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}

	userHttp := models.UserHTTP{}

	if err := c.BindJSON(&userHttp); err != nil {
		err = users.ErrBadRequestBody
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	studentHttp := &models.StudentHTTP{
		UserHTTP: &userHttp,
	}

	err := h.usersDelegate.UpdateStudent(studentHttp)
	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) DeleteStudent(c *gin.Context) {
	log.Println("Delete Student")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}

	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}

	studentId := c.Param("studentId")
	id, atoiErr := strconv.Atoi(studentId)
	if atoiErr != nil {
		atoiErr = users.ErrBadRequest
		log.Println(atoiErr)
		ErrorHandling(atoiErr, c)
		return
	}
	err := h.usersDelegate.DeleteStudent(uint(id))

	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}
	c.Status(http.StatusOK)
}

type SetRobboGroupIdForStudentInput struct {
	RobboUnitId string `json:"robboUnitId"`
}

func (h *Handler) SetRobboGroupIdForStudent(c *gin.Context) {
	log.Println("Set RobboGroupId For Student")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}

	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}

	studentId := c.Param("studentId")
	robboGroupId := c.Param("robboGroupId")
	input := new(SetRobboGroupIdForStudentInput)

	if err := c.BindJSON(&input); err != nil {
		err = users.ErrBadRequestBody
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	// TODO rename method to set robboGroupId, robboUnitId
	err := h.usersDelegate.AddStudentToRobboGroup(studentId, robboGroupId, input.RobboUnitId)
	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}
	c.Status(http.StatusOK)
}

func (h *Handler) CreateTeacher(c *gin.Context) {
	log.Println("Create Teacher")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}

	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}

	userHttp := models.UserHTTP{
		Role: int(models.Teacher),
	}

	if err := c.BindJSON(&userHttp); err != nil {
		err = users.ErrBadRequestBody
		log.Println(err)
		ErrorHandling(err, c)
		return
	}
	teacherHttp := &models.TeacherHTTP{
		UserHTTP: &userHttp,
	}

	teacherId, err := h.usersDelegate.CreateTeacher(teacherHttp)
	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"teacherId": teacherId,
	})

}

func (h *Handler) DeleteTeacher(c *gin.Context) {
	log.Println("Delete Teacher")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}

	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}

	teacherId := c.Param("teacherId")
	id, atoiErr := strconv.Atoi(teacherId)
	if atoiErr != nil {
		atoiErr = users.ErrBadRequest
		log.Println(atoiErr)
		ErrorHandling(atoiErr, c)
		return
	}
	err := h.usersDelegate.DeleteTeacher(uint(id))

	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) GetTeacherById(c *gin.Context) {
	log.Println("Get Teacher By Id")
	_, _, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}
	teacherId := c.Param("teacherId")

	teacher, err := h.usersDelegate.GetTeacherById(teacherId)

	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	c.JSON(http.StatusOK, getTeacherResponse{
		*teacher,
	})
}

func (h *Handler) GetAllTeachers(c *gin.Context) {
	log.Println("Get All Teachers")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}

	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}
	teachers, err := h.usersDelegate.GetAllTeachers()
	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	c.JSON(http.StatusOK, teachers)
}

type updateTeacherInput struct {
	Teacher *models.TeacherHTTP `json:"teacher"`
}

func (h *Handler) UpdateTeacher(c *gin.Context) {
	log.Println("Update Teacher")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}

	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}
	userHttp := models.UserHTTP{}

	if err := c.BindJSON(&userHttp); err != nil {
		err = users.ErrBadRequestBody
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	teacherHttp := &models.TeacherHTTP{
		UserHTTP: &userHttp,
	}

	err := h.usersDelegate.UpdateTeacher(teacherHttp)
	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) GetParentById(c *gin.Context) {
	log.Println("Get Parent By Id")
	_, _, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}
	parentId := c.Param("parentId")

	parent, err := h.usersDelegate.GetParentById(parentId)

	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	c.JSON(http.StatusOK, getParentResponse{
		*parent,
	})
}

func (h *Handler) GetAllParent(c *gin.Context) {
	log.Println("Get All Parents")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}

	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}

	parents, err := h.usersDelegate.GetAllParent()
	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	c.JSON(http.StatusOK, parents)
}

func (h *Handler) CreateParent(c *gin.Context) {
	log.Println("Create Parent")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}

	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}
	userHttp := models.UserHTTP{
		Role: int(models.Parent),
	}

	if err := c.BindJSON(&userHttp); err != nil {
		err = users.ErrBadRequestBody
		log.Println(err)
		ErrorHandling(err, c)
		return
	}
	parentHttp := &models.ParentHTTP{
		UserHTTP: &userHttp,
	}

	parentId, err := h.usersDelegate.CreateParent(parentHttp)
	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"parentId": parentId,
	})

}

func (h *Handler) DeleteParent(c *gin.Context) {
	log.Println("Delete Parent")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}

	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}

	parentId := c.Param("parentId")
	id, atoiErr := strconv.Atoi(parentId)
	if atoiErr != nil {
		atoiErr = users.ErrBadRequest
		log.Println(atoiErr)
		ErrorHandling(atoiErr, c)
		return
	}
	err := h.usersDelegate.DeleteParent(uint(id))

	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	c.Status(http.StatusOK)
}

type updateParentInput struct {
	Parent *models.ParentHTTP `json:"parent"`
}

func (h *Handler) UpdateParent(c *gin.Context) {
	fmt.Println("Update Parent")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}

	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}
	userHttp := models.UserHTTP{}

	if err := c.BindJSON(&userHttp); err != nil {
		err = users.ErrBadRequestBody
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	parentHttp := &models.ParentHTTP{
		UserHTTP: &userHttp,
	}

	err := h.usersDelegate.UpdateParent(parentHttp)
	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) GetFreeListenerById(c *gin.Context) {
	log.Println("Get FreeListener By Id")
	_, _, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}
	freeListenerId := c.Param("freeListenerId")

	freeListener, err := h.usersDelegate.GetFreeListenerById(freeListenerId)

	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	c.JSON(http.StatusOK, getFreeListener{
		freeListener,
	})
}

func (h *Handler) CreateFreeListener(c *gin.Context) {
	log.Println("Create FreeListener")
	_, _, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}

	userHttp := models.UserHTTP{
		Role: int(models.FreeListener),
	}

	if err := c.BindJSON(&userHttp); err != nil {
		err = users.ErrBadRequestBody
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	freeListener := &models.FreeListenerHttp{
		UserHTTP: userHttp,
	}

	freeListenerId, err := h.usersDelegate.CreateFreeListener(freeListener)
	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"freeListenerId": freeListenerId,
	})

}

func (h *Handler) DeleteFreeListener(c *gin.Context) {
	log.Println("Delete Free Listener")
	_, _, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}

	freeListenerId := c.Param("freeListenerId")
	id, atoiErr := strconv.Atoi(freeListenerId)
	if atoiErr != nil {
		atoiErr = users.ErrBadRequest
		log.Println(atoiErr)
		ErrorHandling(atoiErr, c)
		return
	}
	err := h.usersDelegate.DeleteFreeListener(uint(id))

	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	c.Status(http.StatusOK)
}

type updateFreeListenerInput struct {
	FreeListener *models.FreeListenerHttp `json:"freeListener"`
}

func (h *Handler) UpdateFreeListener(c *gin.Context) {
	log.Println("Update Free Listener")
	_, _, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}
	userHttp := models.UserHTTP{}

	if err := c.BindJSON(&userHttp); err != nil {
		err = users.ErrBadRequestBody
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	freeListenerHttp := &models.FreeListenerHttp{
		UserHTTP: userHttp,
	}

	err := h.usersDelegate.UpdateFreeListener(freeListenerHttp)
	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) GetUnitAdminByID(c *gin.Context) {
	log.Println("Get Unit Admin By ID")
	_, _, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}
	unitAdminId := c.Param("unitAdminId")

	unitAdmin, err := h.usersDelegate.GetUnitAdminById(unitAdminId)

	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	c.JSON(http.StatusOK, getUnitAdminResponse{
		unitAdmin,
	})
}

func (h *Handler) GetUnitAdminsByRobboUnitId(c *gin.Context) {
	log.Println("Get UnitAdmins By RobboUnitId")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}

	allowedRoles := []models.Role{models.SuperAdmin}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}
	robboUnitId := c.Param("robboUnitId")

	unitAdmins, err := h.usersDelegate.GetUnitAdminByRobboUnitId(robboUnitId)

	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	c.JSON(http.StatusOK, unitAdmins)
}

func (h *Handler) GetAllUnitAdmins(c *gin.Context) {
	log.Println("Get All UnitAdmins")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}

	allowedRoles := []models.Role{models.SuperAdmin}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}
	unitAdmins, err := h.usersDelegate.GetAllUnitAdmins()
	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	c.JSON(http.StatusOK, unitAdmins)
}

type updateUnitAdminInput struct {
	UnitAdmin *models.UnitAdminHTTP `json:"unitAdmin"`
}

func (h *Handler) UpdateUnitAdmin(c *gin.Context) {
	log.Println("Update Unit Admin")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}

	allowedRoles := []models.Role{models.SuperAdmin}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}
	userHttp := models.UserHTTP{}

	if err := c.BindJSON(&userHttp); err != nil {
		err = users.ErrBadRequestBody
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	unitAdminHttp := &models.UnitAdminHTTP{
		UserHTTP: &userHttp,
	}

	err := h.usersDelegate.UpdateUnitAdmin(unitAdminHttp)
	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) CreateUnitAdmin(c *gin.Context) {
	log.Println("Create Unit Admin")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}

	allowedRoles := []models.Role{models.SuperAdmin}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}

	userHttp := models.UserHTTP{
		Role: int(models.UnitAdmin),
	}
	if err := c.BindJSON(&userHttp); err != nil {
		err = users.ErrBadRequestBody
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	unitAdminHttp := &models.UnitAdminHTTP{
		UserHTTP: &userHttp,
	}

	unitAdminId, err := h.usersDelegate.CreateUnitAdmin(unitAdminHttp)
	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"unitAdminId": unitAdminId,
	})
}

func (h *Handler) DeleteUnitAdmin(c *gin.Context) {
	log.Println("Delete Unit Admin")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}

	allowedRoles := []models.Role{models.SuperAdmin}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}

	unitAdminId := c.Param("unitAdminId")
	id, atoiErr := strconv.Atoi(unitAdminId)
	if atoiErr != nil {
		atoiErr = users.ErrBadRequest
		log.Println(atoiErr)
		ErrorHandling(atoiErr, c)
		return
	}
	err := h.usersDelegate.DeleteUnitAdmin(uint(id))

	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}
	c.Status(http.StatusOK)
}

func (h *Handler) SearchUnitAdminByEmail(c *gin.Context) {
	log.Println("Search Unit Admin By Email")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}

	allowedRoles := []models.Role{models.SuperAdmin}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}
	unitAdminEmail := c.Param("unitAdminEmail")

	unitAdmins, err := h.usersDelegate.SearchUnitAdminByEmail(unitAdminEmail)
	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}
	c.JSON(http.StatusOK, unitAdmins)
}

func (h *Handler) GetSuperAdminById(c *gin.Context) {
	log.Println("Get Super Admin By Id")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}

	allowedRoles := []models.Role{models.SuperAdmin}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}
	superAdminId := c.Param("superAdminId")

	superAdmin, err := h.usersDelegate.GetSuperAdminById(superAdminId)

	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	c.JSON(http.StatusOK, getSuperAdminResponse{
		superAdmin,
	})
}

type updateSuperAdminInput struct {
	SuperAdmin *models.SuperAdminHTTP `json:"superAdmin"`
}

func (h *Handler) UpdateSuperAdmin(c *gin.Context) {
	log.Println("Update Super Admin")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}

	allowedRoles := []models.Role{models.SuperAdmin}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}
	userHttp := models.UserHTTP{}

	if err := c.BindJSON(&userHttp); err != nil {
		err = users.ErrBadRequestBody
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	superAdminHTTP := &models.SuperAdminHTTP{
		UserHTTP: &userHttp,
	}

	err := h.usersDelegate.UpdateSuperAdmin(superAdminHTTP)
	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) DeleteSuperAdmin(c *gin.Context) {
	log.Println("Delete Super Admin")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}

	allowedRoles := []models.Role{models.SuperAdmin}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}
	superAdminId := c.Param("AdminId")
	id, atoiErr := strconv.Atoi(superAdminId)
	if atoiErr != nil {
		atoiErr = users.ErrBadRequest
		log.Println(atoiErr)
		ErrorHandling(atoiErr, c)
		return
	}
	err := h.usersDelegate.DeleteSuperAdmin(uint(id))

	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}
	c.Status(http.StatusOK)
}

type createRelation struct {
	ParentId string `json:"parentId"`
	ChildId  string `json:"childId"`
}

func (h *Handler) CreateRelation(c *gin.Context) {
	log.Println("Create Relation")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}

	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}
	createRelationInput := new(createRelation)

	if err := c.BindJSON(createRelationInput); err != nil {
		err = users.ErrBadRequestBody
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	createRelationErr := h.usersDelegate.CreateRelation(createRelationInput.ParentId, createRelationInput.ChildId)

	if createRelationErr != nil {
		log.Println(createRelationErr)
		ErrorHandling(createRelationErr, c)
		return
	}

	c.Status(http.StatusOK)
}

type setNewUnitAdminForRobboUnitRequest struct {
	UnitAdminId string `json:"unitAdminId"`
	RobboUnitId string `json:"robboUnitId"`
}

func (h *Handler) SetNewUnitAdminForRobboUnit(c *gin.Context) {
	log.Println("Set NewUnitAdmin For RobboUnit")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}

	allowedRoles := []models.Role{models.SuperAdmin}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}
	input := new(setNewUnitAdminForRobboUnitRequest)

	if err := c.BindJSON(input); err != nil {
		err = users.ErrBadRequestBody
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	createRelationErr := h.usersDelegate.SetNewUnitAdminForRobboUnit(input.UnitAdminId, input.RobboUnitId)

	if createRelationErr != nil {
		log.Println(createRelationErr)
		ErrorHandling(createRelationErr, c)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) DeleteUnitAdminForRobboUnit(c *gin.Context) {
	log.Println("Delete UnitAdmin For RobboUnit")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}

	allowedRoles := []models.Role{models.SuperAdmin}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}

	input := new(setNewUnitAdminForRobboUnitRequest)

	if err := c.BindJSON(input); err != nil {
		err = users.ErrBadRequestBody
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	deleteRelationErr := h.usersDelegate.DeleteUnitAdminForRobboUnit(input.UnitAdminId, input.RobboUnitId)

	if deleteRelationErr != nil {
		log.Println(deleteRelationErr)
		ErrorHandling(deleteRelationErr, c)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) GetUserById(c *gin.Context) {
	log.Println("Get User by id")
	_, role, errUserIdentityErr := h.authDelegate.UserIdentity(c)
	if errUserIdentityErr != nil {
		log.Println(errUserIdentityErr)
		ErrorHandling(errUserIdentityErr, c)
		return
	}
	allowedRoles := []models.Role{models.SuperAdmin, models.UnitAdmin, models.Teacher}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}
	userId := c.Param("userId")
	userRole, _ := strconv.Atoi(c.Query("role"))

	switch models.Role(userRole) {
	case models.Student:
		student, getStudentErr := h.usersDelegate.GetStudentById(userId)
		if getStudentErr != nil {
			ErrorHandling(getStudentErr, c)
			return
		}
		c.JSON(http.StatusOK, student)
	case models.Teacher:
		teacher, getTeacherErr := h.usersDelegate.GetTeacherById(userId)
		if getTeacherErr != nil {
			ErrorHandling(getTeacherErr, c)
			return
		}
		c.JSON(http.StatusOK, teacher)
		return
	case models.Parent:
		parent, getParentErr := h.usersDelegate.GetParentById(userId)
		if getParentErr != nil {
			ErrorHandling(getParentErr, c)
			return
		}
		c.JSON(http.StatusOK, parent)
		return
	case models.FreeListener:
		freeListener, getFreeListenerErr := h.usersDelegate.GetFreeListenerById(userId)
		if getFreeListenerErr != nil {
			ErrorHandling(getFreeListenerErr, c)
			return
		}
		c.JSON(http.StatusOK, freeListener)
	case models.UnitAdmin:
		unitAdmin, getUnitAdminErr := h.usersDelegate.GetUnitAdminById(userId)
		if getUnitAdminErr != nil {
			ErrorHandling(getUnitAdminErr, c)
			return
		}
		c.JSON(http.StatusOK, unitAdmin)
		return
	case models.SuperAdmin:
		superAdmin, getSuperAdminErr := h.usersDelegate.GetSuperAdminById(userId)
		if getSuperAdminErr != nil {
			ErrorHandling(getSuperAdminErr, c)
			return
		}
		c.JSON(http.StatusOK, superAdmin)
		return
	}
}

func ErrorHandling(err error, c *gin.Context) {
	switch err {
	case users.ErrBadRequest:
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	case users.ErrInternalServerLevel:
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	case users.ErrBadRequestBody:
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	case auth.ErrInvalidAccessToken:
		c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
	case auth.ErrTokenNotFound:
		c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
	case auth.ErrUserNotFound:
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	case auth.ErrNotAccess:
		c.AbortWithStatusJSON(http.StatusForbidden, err.Error())
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
}
