// Package classroom provides access to the Google Classroom API.
//
// Usage example:
//
//   import "google.golang.org/api/classroom/v1"
//   ...
//   classroomService, err := classroom.New(oauthHttpClient)
package classroom

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Background

const apiId = "classroom:v1"
const apiName = "classroom"
const apiVersion = "v1"
const basePath = "https://classroom.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// Manage your Google Classroom classes
	ClassroomCoursesScope = "https://www.googleapis.com/auth/classroom.courses"

	// View your Google Classroom classes
	ClassroomCoursesReadonlyScope = "https://www.googleapis.com/auth/classroom.courses.readonly"

	// View the email addresses of people in your classes
	ClassroomProfileEmailsScope = "https://www.googleapis.com/auth/classroom.profile.emails"

	// View the profile photos of people in your classes
	ClassroomProfilePhotosScope = "https://www.googleapis.com/auth/classroom.profile.photos"

	// Manage your Google Classroom class rosters
	ClassroomRostersScope = "https://www.googleapis.com/auth/classroom.rosters"

	// View your Google Classroom class rosters
	ClassroomRostersReadonlyScope = "https://www.googleapis.com/auth/classroom.rosters.readonly"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Courses = NewCoursesService(s)
	s.Invitations = NewInvitationsService(s)
	s.UserProfiles = NewUserProfilesService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Courses *CoursesService

	Invitations *InvitationsService

	UserProfiles *UserProfilesService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewCoursesService(s *Service) *CoursesService {
	rs := &CoursesService{s: s}
	rs.Aliases = NewCoursesAliasesService(s)
	rs.Students = NewCoursesStudentsService(s)
	rs.Teachers = NewCoursesTeachersService(s)
	return rs
}

type CoursesService struct {
	s *Service

	Aliases *CoursesAliasesService

	Students *CoursesStudentsService

	Teachers *CoursesTeachersService
}

func NewCoursesAliasesService(s *Service) *CoursesAliasesService {
	rs := &CoursesAliasesService{s: s}
	return rs
}

type CoursesAliasesService struct {
	s *Service
}

func NewCoursesStudentsService(s *Service) *CoursesStudentsService {
	rs := &CoursesStudentsService{s: s}
	return rs
}

type CoursesStudentsService struct {
	s *Service
}

func NewCoursesTeachersService(s *Service) *CoursesTeachersService {
	rs := &CoursesTeachersService{s: s}
	return rs
}

type CoursesTeachersService struct {
	s *Service
}

func NewInvitationsService(s *Service) *InvitationsService {
	rs := &InvitationsService{s: s}
	return rs
}

type InvitationsService struct {
	s *Service
}

func NewUserProfilesService(s *Service) *UserProfilesService {
	rs := &UserProfilesService{s: s}
	return rs
}

type UserProfilesService struct {
	s *Service
}

// Course: A Course in Classroom.
type Course struct {
	// AlternateLink: Absolute link to this course in the Classroom web UI.
	// Read-only.
	AlternateLink string `json:"alternateLink,omitempty"`

	// CourseState: State of the course. If unspecified, the default state
	// will be `PROVISIONED`.
	//
	// Possible values:
	//   "COURSE_STATE_UNSPECIFIED"
	//   "ACTIVE"
	//   "ARCHIVED"
	//   "PROVISIONED"
	//   "DECLINED"
	CourseState string `json:"courseState,omitempty"`

	// CreationTime: Creation time of the course. Specifying this field in a
	// course update mask will result in an error. Read-only.
	CreationTime string `json:"creationTime,omitempty"`

	// Description: Optional description. For example, "We'll be learning
	// about about the structure of living creatures from a combination of
	// textbooks, guest lectures, and lab work. Expect to be excited!" If
	// set, this field must be a valid UTF-8 string and no longer than
	// 30,000 characters.
	Description string `json:"description,omitempty"`

	// DescriptionHeading: Optional heading for the description. For
	// example, "Welcome to 10th Grade Biology" If set, this field must be a
	// valid UTF-8 string and no longer than 3600 characters.
	DescriptionHeading string `json:"descriptionHeading,omitempty"`

	// EnrollmentCode: Enrollment code to use when joining this course.
	// Specifying this field in a course update mask will result in an
	// error. Read-only.
	EnrollmentCode string `json:"enrollmentCode,omitempty"`

	// Id: Unique identifier for this course assigned by Classroom. You may
	// optionally set this to an [alias
	// string][google.classroom.v1.CourseAlias] as part of [creating a
	// course][google.classroom.v1.Courses.CreateCourse], creating a
	// corresponding alias. The `ID` cannot be updated after a course is
	// created. Specifying this field in a course update mask will result in
	// an error.
	Id string `json:"id,omitempty"`

	// Name: Name of the course. For example, "10th Grade Biology". This is
	// required and must be between 1 and 750 characters and a valid UTF-8
	// string.
	Name string `json:"name,omitempty"`

	// OwnerId: The identifier of the owner (and primary teacher) of a
	// course. When specified as a parameter of CreateCourseRequest, this
	// field is required. It may be the numeric identifier for the user, or
	// an alias that identifies the owner. The following aliases are
	// supported: * the e-mail address of the user * the string literal
	// "me", indicating that the requesting user This must be set in a
	// CreateRequest; specifying this field in a course update mask will
	// result in an error.
	OwnerId string `json:"ownerId,omitempty"`

	// Room: Optional room location. For example, "301" If set, this field
	// must be a valid UTF-8 string and no longer than 650 characters.
	Room string `json:"room,omitempty"`

	// Section: Section of the course For example, "Period 2". If set, this
	// field must be a valid UTF-8 string and no longer than 2800
	// characters.
	Section string `json:"section,omitempty"`

	// UpdateTime: Time of the most recent update to this course. Specifying
	// this field in a course update mask will result in an error.
	// Read-only.
	UpdateTime string `json:"updateTime,omitempty"`
}

// CourseAlias: Alternative identifier for a course. An alias uniquely
// identifies a course. It will be unique within one of the following
// scopes: * domain: A domain-scoped alias is visible to all users
// within the alias creator's domain and may only be created by a domain
// admin. A domain-scoped alias is often used when a course has an
// identifier external to Classroom. * project: A project-scoped alias
// is visible to any request from an application using the Developer
// Console Project ID that created the alias and may be created by any
// project. A project-scoped alias is often used when an application has
// alternative identifiers. A random value can also be used to avoid
// duplicate courses in the event of transmission failures, as retrying
// a request will return ALREADY_EXISTS if a previous one has succeeded.
type CourseAlias struct {
	// Alias: Alias string. The format of the string indicated the desired
	// alias scoping. * "d:" indicates a domain-scoped alias. Example:
	// d:math_101 * "p:" indicates a project-scoped alias. Example: p:abc123
	// This field has a maximum length of 256 characters.
	Alias string `json:"alias,omitempty"`
}

// Empty: A generic empty message that you can re-use to avoid defining
// duplicated empty messages in your APIs. A typical example is to use
// it as the request or the response type of an API method. For
// instance: service Foo { rpc Bar(google.protobuf.Empty) returns
// (google.protobuf.Empty); } The JSON representation for `Empty` is
// empty JSON object `{}`.
type Empty struct {
}

// GlobalPermission: Global user permission description.
type GlobalPermission struct {
	// Permission: Permission value.
	//
	// Possible values:
	//   "PERMISSION_UNSPECIFIED"
	//   "CREATE_COURSE"
	Permission string `json:"permission,omitempty"`
}

// Invitation: An invitation to join a course.
type Invitation struct {
	// CourseId: Identifier of the course to invite the user to.
	CourseId string `json:"courseId,omitempty"`

	// Id: Unique identifier assigned by Classroom. Read-only
	Id string `json:"id,omitempty"`

	// Role: Role to invite the user to have. Must not be
	// `COURSE_ROLE_UNSPECIFIED`.
	//
	// Possible values:
	//   "COURSE_ROLE_UNSPECIFIED"
	//   "STUDENT"
	//   "TEACHER"
	Role string `json:"role,omitempty"`

	// UserId: Identifier of the invited user. When specified as a parameter
	// of a request, this may be set to an alias that identifies the user to
	// invite. The supported aliases are: * the e-mail address of the user *
	// the string literal "me", indicating that the requesting user
	UserId string `json:"userId,omitempty"`
}

// ListCourseAliasesResponse: Response when listing course aliases.
type ListCourseAliasesResponse struct {
	// Aliases: The course aliases.
	Aliases []*CourseAlias `json:"aliases,omitempty"`

	// NextPageToken: Token identifying the next page of results to return.
	// If empty, no further results are available.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

// ListCoursesResponse: Response when listing courses.
type ListCoursesResponse struct {
	// Courses: Courses that match the request.
	Courses []*Course `json:"courses,omitempty"`

	// NextPageToken: Token identifying the next page of results to return.
	// If empty, no further results are available.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

// ListInvitationsResponse: Response when listing invitations.
type ListInvitationsResponse struct {
	// Invitations: Invitations that match the request.
	Invitations []*Invitation `json:"invitations,omitempty"`

	// NextPageToken: Token identifying the next page of results to return.
	// If empty, no further results are available.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

// ListStudentsResponse: Response when listing students.
type ListStudentsResponse struct {
	// NextPageToken: Token identifying the next page of results to return.
	// If empty, no further results are available.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Students: The students who match the list request.
	Students []*Student `json:"students,omitempty"`
}

// ListTeachersResponse: Response when listing teachers.
type ListTeachersResponse struct {
	// NextPageToken: Token identifying the next page of results to return.
	// If empty, no further results are available.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Teachers: The teachers who match the list request.
	Teachers []*Teacher `json:"teachers,omitempty"`
}

// Name: Details of the user's name.
type Name struct {
	// FamilyName: The user's last name. Read-only
	FamilyName string `json:"familyName,omitempty"`

	// FullName: The user's full name formed by concatenating the first and
	// last name values. Read-only
	FullName string `json:"fullName,omitempty"`

	// GivenName: The user's first name. Read-only
	GivenName string `json:"givenName,omitempty"`
}

// Student: Student in a course.
type Student struct {
	// CourseId: Unique identifier of the course. Read-only
	CourseId string `json:"courseId,omitempty"`

	// Profile: Global user information for the student. Read-only
	Profile *UserProfile `json:"profile,omitempty"`

	// UserId: The identifier of the user. When specified as a parameter of
	// request, this field may be set to an alias that identifies the
	// student. The following are supported: * the e-mail address of the
	// user * the string literal "me", indicating that the requesting user
	UserId string `json:"userId,omitempty"`
}

// Teacher: Teacher of a course.
type Teacher struct {
	// CourseId: Unique identifier of the course. Read-only
	CourseId string `json:"courseId,omitempty"`

	// Profile: Global user information for the teacher. Read-only
	Profile *UserProfile `json:"profile,omitempty"`

	// UserId: The identifier of the user. When specified as a parameter of
	// request, this field may be set to an alias that identifies the
	// teacher. The following are supported: * the e-mail address of the
	// user * the string literal "me", indicating the requesting user
	UserId string `json:"userId,omitempty"`
}

// UserProfile: Global information for a user.
type UserProfile struct {
	// EmailAddress: E-mail address of the user. Read-only
	EmailAddress string `json:"emailAddress,omitempty"`

	// Id: Unique identifier of the user. Read-only
	Id string `json:"id,omitempty"`

	// Name: Name of the user. Read-only
	Name *Name `json:"name,omitempty"`

	// Permissions: Global permissions of the user. Read-only
	Permissions []*GlobalPermission `json:"permissions,omitempty"`

	// PhotoUrl: Url of user's profile photo. Read-only
	PhotoUrl string `json:"photoUrl,omitempty"`
}

// method id "classroom.courses.create":

type CoursesCreateCall struct {
	s      *Service
	course *Course
	opt_   map[string]interface{}
}

// Create: Creates a course. The user specified as the primary teacher
// in `primary_teacher_id` is the owner of the created course and added
// as a teacher. This method returns the following error codes: *
// `PERMISSION_DENIED` if the requesting user is not permitted to create
// courses. * `NOT_FOUND` if the primary teacher is not a valid user. *
// `ALREADY_EXISTS` if an alias was specified and already exists.
func (r *CoursesService) Create(course *Course) *CoursesCreateCall {
	c := &CoursesCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.course = course
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CoursesCreateCall) Fields(s ...googleapi.Field) *CoursesCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *CoursesCreateCall) Do() (*Course, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.course)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/courses")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Course
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a course. The user specified as the primary teacher in `primary_teacher_id` is the owner of the created course and added as a teacher. This method returns the following error codes: * `PERMISSION_DENIED` if the requesting user is not permitted to create courses. * `NOT_FOUND` if the primary teacher is not a valid user. * `ALREADY_EXISTS` if an alias was specified and already exists.",
	//   "httpMethod": "POST",
	//   "id": "classroom.courses.create",
	//   "path": "v1/courses",
	//   "request": {
	//     "$ref": "Course"
	//   },
	//   "response": {
	//     "$ref": "Course"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/classroom.courses"
	//   ]
	// }

}

// method id "classroom.courses.delete":

type CoursesDeleteCall struct {
	s    *Service
	id   string
	opt_ map[string]interface{}
}

// Delete: Deletes a course. This method returns the following error
// codes: * `PERMISSION_DENIED` if the requesting user is not permitted
// to delete the requested course. * `NOT_FOUND` if no course exists
// with the requested ID.
func (r *CoursesService) Delete(id string) *CoursesDeleteCall {
	c := &CoursesDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CoursesDeleteCall) Fields(s ...googleapi.Field) *CoursesDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *CoursesDeleteCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/courses/{id}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Empty
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes a course. This method returns the following error codes: * `PERMISSION_DENIED` if the requesting user is not permitted to delete the requested course. * `NOT_FOUND` if no course exists with the requested ID.",
	//   "httpMethod": "DELETE",
	//   "id": "classroom.courses.delete",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Identifier of the course to delete. This may either be the Classroom-assigned identifier or an [alias][google.classroom.v1.CourseAlias].",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/courses/{id}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/classroom.courses"
	//   ]
	// }

}

// method id "classroom.courses.get":

type CoursesGetCall struct {
	s    *Service
	id   string
	opt_ map[string]interface{}
}

// Get: Returns a course. This method returns the following error codes:
// * `PERMISSION_DENIED` if the requesting user is not permitted to
// access the requested course. * `NOT_FOUND` if no course exists with
// the requested ID.
func (r *CoursesService) Get(id string) *CoursesGetCall {
	c := &CoursesGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CoursesGetCall) Fields(s ...googleapi.Field) *CoursesGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *CoursesGetCall) Do() (*Course, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/courses/{id}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Course
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns a course. This method returns the following error codes: * `PERMISSION_DENIED` if the requesting user is not permitted to access the requested course. * `NOT_FOUND` if no course exists with the requested ID.",
	//   "httpMethod": "GET",
	//   "id": "classroom.courses.get",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Identifier of the course to return. This may either be the Classroom-assigned identifier or an [alias][google.classroom.v1.CourseAlias].",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/courses/{id}",
	//   "response": {
	//     "$ref": "Course"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/classroom.courses",
	//     "https://www.googleapis.com/auth/classroom.courses.readonly"
	//   ]
	// }

}

// method id "classroom.courses.list":

type CoursesListCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// List: Returns a list of courses that the requesting user is permitted
// to view, restricted to those that match the request. This method
// returns the following error codes: * `INVALID_ARGUMENT` if the query
// argument is malformed. * `NOT_FOUND` if any users specified in the
// query arguments do not exist.
func (r *CoursesService) List() *CoursesListCall {
	c := &CoursesListCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// PageSize sets the optional parameter "pageSize": Maximum number of
// items to return. Zero or unspecified indicates that the server may
// assign a maximum. The server may return fewer than the specified
// number of results.
func (c *CoursesListCall) PageSize(pageSize int64) *CoursesListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken":
// [nextPageToken][google.classroom.v1.ListCoursesResponse.next_page_toke
// n] value returned from a previous
// [list][google.classroom.v1.Courses.ListCourses] call, indicating that
// the subsequent page of results should be returned. The
// [list][google.classroom.v1.Courses.ListCourses] request must be
// identical to the one which resulted in this token.
func (c *CoursesListCall) PageToken(pageToken string) *CoursesListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// StudentId sets the optional parameter "studentId": Restricts returned
// courses to those having a student with the specified identifier, or
// an alias that identifies a student. The following aliases are
// supported: * the e-mail address of the user * the string literal
// "me", indicating that the requesting user
func (c *CoursesListCall) StudentId(studentId string) *CoursesListCall {
	c.opt_["studentId"] = studentId
	return c
}

// TeacherId sets the optional parameter "teacherId": Restricts returned
// courses to those having a teacher with the specified identifier, or
// an alias that identifies a teacher. The following aliases are
// supported: * the e-mail address of the user * the string literal
// "me", indicating that the requesting user
func (c *CoursesListCall) TeacherId(teacherId string) *CoursesListCall {
	c.opt_["teacherId"] = teacherId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CoursesListCall) Fields(s ...googleapi.Field) *CoursesListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *CoursesListCall) Do() (*ListCoursesResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["studentId"]; ok {
		params.Set("studentId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["teacherId"]; ok {
		params.Set("teacherId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/courses")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListCoursesResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns a list of courses that the requesting user is permitted to view, restricted to those that match the request. This method returns the following error codes: * `INVALID_ARGUMENT` if the query argument is malformed. * `NOT_FOUND` if any users specified in the query arguments do not exist.",
	//   "httpMethod": "GET",
	//   "id": "classroom.courses.list",
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Maximum number of items to return. Zero or unspecified indicates that the server may assign a maximum. The server may return fewer than the specified number of results.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "[nextPageToken][google.classroom.v1.ListCoursesResponse.next_page_token] value returned from a previous [list][google.classroom.v1.Courses.ListCourses] call, indicating that the subsequent page of results should be returned. The [list][google.classroom.v1.Courses.ListCourses] request must be identical to the one which resulted in this token.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "studentId": {
	//       "description": "Restricts returned courses to those having a student with the specified identifier, or an alias that identifies a student. The following aliases are supported: * the e-mail address of the user * the string literal `\"me\"`, indicating that the requesting user",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "teacherId": {
	//       "description": "Restricts returned courses to those having a teacher with the specified identifier, or an alias that identifies a teacher. The following aliases are supported: * the e-mail address of the user * the string literal `\"me\"`, indicating that the requesting user",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/courses",
	//   "response": {
	//     "$ref": "ListCoursesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/classroom.courses",
	//     "https://www.googleapis.com/auth/classroom.courses.readonly"
	//   ]
	// }

}

// method id "classroom.courses.patch":

type CoursesPatchCall struct {
	s      *Service
	id     string
	course *Course
	opt_   map[string]interface{}
}

// Patch: Updates one or more fields a course. This method returns the
// following error codes: * `PERMISSION_DENIED` if the requesting user
// is not permitted to modify the requested course. * `NOT_FOUND` if no
// course exists with the requested ID. * `INVALID_ARGUMENT` if invalid
// fields are specified in the update mask or if no update mask is
// supplied.
func (r *CoursesService) Patch(id string, course *Course) *CoursesPatchCall {
	c := &CoursesPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	c.course = course
	return c
}

// UpdateMask sets the optional parameter "updateMask": Mask which
// identifies which fields on the course to update. This field is
// required to do an update. The update will fail if invalid fields are
// specified. Valid fields are listed below: * `name` * `section` *
// `descriptionHeading` * `description` * `room` * `courseState` When
// set in a query parameter, this should be specified as
// `updateMask=,,...`
func (c *CoursesPatchCall) UpdateMask(updateMask string) *CoursesPatchCall {
	c.opt_["updateMask"] = updateMask
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CoursesPatchCall) Fields(s ...googleapi.Field) *CoursesPatchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *CoursesPatchCall) Do() (*Course, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.course)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["updateMask"]; ok {
		params.Set("updateMask", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/courses/{id}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Course
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates one or more fields a course. This method returns the following error codes: * `PERMISSION_DENIED` if the requesting user is not permitted to modify the requested course. * `NOT_FOUND` if no course exists with the requested ID. * `INVALID_ARGUMENT` if invalid fields are specified in the update mask or if no update mask is supplied.",
	//   "httpMethod": "PATCH",
	//   "id": "classroom.courses.patch",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Identifier of the course to update. This may either be the Classroom-assigned identifier or an [alias][google.classroom.v1.CourseAlias].",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "Mask which identifies which fields on the course to update. This field is required to do an update. The update will fail if invalid fields are specified. Valid fields are listed below: * `name` * `section` * `descriptionHeading` * `description` * `room` * `courseState` When set in a query parameter, this should be specified as `updateMask=,,...`",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/courses/{id}",
	//   "request": {
	//     "$ref": "Course"
	//   },
	//   "response": {
	//     "$ref": "Course"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/classroom.courses"
	//   ]
	// }

}

// method id "classroom.courses.update":

type CoursesUpdateCall struct {
	s      *Service
	id     string
	course *Course
	opt_   map[string]interface{}
}

// Update: Updates a course. This method returns the following error
// codes: * `PERMISSION_DENIED` if the requesting user is not permitted
// to modify the requested course. * `NOT_FOUND` if no course exists
// with the requested ID.
func (r *CoursesService) Update(id string, course *Course) *CoursesUpdateCall {
	c := &CoursesUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	c.course = course
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CoursesUpdateCall) Fields(s ...googleapi.Field) *CoursesUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *CoursesUpdateCall) Do() (*Course, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.course)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/courses/{id}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Course
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates a course. This method returns the following error codes: * `PERMISSION_DENIED` if the requesting user is not permitted to modify the requested course. * `NOT_FOUND` if no course exists with the requested ID.",
	//   "httpMethod": "PUT",
	//   "id": "classroom.courses.update",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Identifier of the course to update. This may either be the Classroom-assigned identifier or an [alias][google.classroom.v1.CourseAlias].",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/courses/{id}",
	//   "request": {
	//     "$ref": "Course"
	//   },
	//   "response": {
	//     "$ref": "Course"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/classroom.courses"
	//   ]
	// }

}

// method id "classroom.courses.aliases.create":

type CoursesAliasesCreateCall struct {
	s           *Service
	courseId    string
	coursealias *CourseAlias
	opt_        map[string]interface{}
}

// Create: Creates an alias to a course. This method returns the
// following error codes: * `PERMISSION_DENIED` if the requesting user
// is not permitted to create the alias. * `NOT_FOUND` if the course
// does not exist. * `ALREADY_EXISTS` if the alias already exists.
func (r *CoursesAliasesService) Create(courseId string, coursealias *CourseAlias) *CoursesAliasesCreateCall {
	c := &CoursesAliasesCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.courseId = courseId
	c.coursealias = coursealias
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CoursesAliasesCreateCall) Fields(s ...googleapi.Field) *CoursesAliasesCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *CoursesAliasesCreateCall) Do() (*CourseAlias, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.coursealias)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/courses/{courseId}/aliases")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"courseId": c.courseId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *CourseAlias
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates an alias to a course. This method returns the following error codes: * `PERMISSION_DENIED` if the requesting user is not permitted to create the alias. * `NOT_FOUND` if the course does not exist. * `ALREADY_EXISTS` if the alias already exists.",
	//   "httpMethod": "POST",
	//   "id": "classroom.courses.aliases.create",
	//   "parameterOrder": [
	//     "courseId"
	//   ],
	//   "parameters": {
	//     "courseId": {
	//       "description": "The identifier of the course to alias. This may either be the Classroom-assigned identifier or an [alias][google.classroom.v1.CourseAlias].",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/courses/{courseId}/aliases",
	//   "request": {
	//     "$ref": "CourseAlias"
	//   },
	//   "response": {
	//     "$ref": "CourseAlias"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/classroom.courses"
	//   ]
	// }

}

// method id "classroom.courses.aliases.delete":

type CoursesAliasesDeleteCall struct {
	s        *Service
	courseId string
	aliasid  string
	opt_     map[string]interface{}
}

// Delete: Deletes an alias of a course. This method returns the
// following error codes: * `PERMISSION_DENIED` if the requesting user
// is not permitted to remove the alias. * `NOT_FOUND` if the alias does
// not exist.
func (r *CoursesAliasesService) Delete(courseId string, aliasid string) *CoursesAliasesDeleteCall {
	c := &CoursesAliasesDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.courseId = courseId
	c.aliasid = aliasid
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CoursesAliasesDeleteCall) Fields(s ...googleapi.Field) *CoursesAliasesDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *CoursesAliasesDeleteCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/courses/{courseId}/aliases/{alias}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"courseId": c.courseId,
		"alias":    c.aliasid,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Empty
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes an alias of a course. This method returns the following error codes: * `PERMISSION_DENIED` if the requesting user is not permitted to remove the alias. * `NOT_FOUND` if the alias does not exist.",
	//   "httpMethod": "DELETE",
	//   "id": "classroom.courses.aliases.delete",
	//   "parameterOrder": [
	//     "courseId",
	//     "alias"
	//   ],
	//   "parameters": {
	//     "alias": {
	//       "description": "The alias to delete. This may not be the Classroom-assigned identifier.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "courseId": {
	//       "description": "The identifier of the course whose alias should be deleted. This may either be the Classroom-assigned identifier or an [alias][google.classroom.v1.CourseAlias].",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/courses/{courseId}/aliases/{alias}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/classroom.courses"
	//   ]
	// }

}

// method id "classroom.courses.aliases.list":

type CoursesAliasesListCall struct {
	s        *Service
	courseId string
	opt_     map[string]interface{}
}

// List: Lists the aliases of a course. This method returns the
// following error codes: * `PERMISSION_DENIED` if the requesting user
// is not permitted to access the course. * `NOT_FOUND` if the course
// does not exist.
func (r *CoursesAliasesService) List(courseId string) *CoursesAliasesListCall {
	c := &CoursesAliasesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.courseId = courseId
	return c
}

// PageSize sets the optional parameter "pageSize": Maximum number of
// items to return. Zero or unspecified indicates that the server may
// assign a maximum. The server may return fewer than the specified
// number of results.
func (c *CoursesAliasesListCall) PageSize(pageSize int64) *CoursesAliasesListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken":
// [nextPageToken][google.classroom.v1.ListCourseAliasesResponse.next_pag
// e_token] value returned from a previous
// [list][google.classroom.v1.Courses.ListCourseAliases] call,
// indicating that the subsequent page of results should be returned.
// The [list][google.classroom.v1.Courses.ListCourseAliases] request
// must be identical to the one which resulted in this token.
func (c *CoursesAliasesListCall) PageToken(pageToken string) *CoursesAliasesListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CoursesAliasesListCall) Fields(s ...googleapi.Field) *CoursesAliasesListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *CoursesAliasesListCall) Do() (*ListCourseAliasesResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/courses/{courseId}/aliases")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"courseId": c.courseId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListCourseAliasesResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists the aliases of a course. This method returns the following error codes: * `PERMISSION_DENIED` if the requesting user is not permitted to access the course. * `NOT_FOUND` if the course does not exist.",
	//   "httpMethod": "GET",
	//   "id": "classroom.courses.aliases.list",
	//   "parameterOrder": [
	//     "courseId"
	//   ],
	//   "parameters": {
	//     "courseId": {
	//       "description": "The identifier of the course. This may either be the Classroom-assigned identifier or an [alias][google.classroom.v1.CourseAlias].",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Maximum number of items to return. Zero or unspecified indicates that the server may assign a maximum. The server may return fewer than the specified number of results.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "[nextPageToken][google.classroom.v1.ListCourseAliasesResponse.next_page_token] value returned from a previous [list][google.classroom.v1.Courses.ListCourseAliases] call, indicating that the subsequent page of results should be returned. The [list][google.classroom.v1.Courses.ListCourseAliases] request must be identical to the one which resulted in this token.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/courses/{courseId}/aliases",
	//   "response": {
	//     "$ref": "ListCourseAliasesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/classroom.courses",
	//     "https://www.googleapis.com/auth/classroom.courses.readonly"
	//   ]
	// }

}

// method id "classroom.courses.students.create":

type CoursesStudentsCreateCall struct {
	s        *Service
	courseId string
	student  *Student
	opt_     map[string]interface{}
}

// Create: Adds a user as a student of a course. This method returns the
// following error codes: * `PERMISSION_DENIED` if the requesting user
// is not permitted to create students in this course. * `NOT_FOUND` if
// the requested course ID does not exist. * `ALREADY_EXISTS` if the
// user is already a student or student in the course.
func (r *CoursesStudentsService) Create(courseId string, student *Student) *CoursesStudentsCreateCall {
	c := &CoursesStudentsCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.courseId = courseId
	c.student = student
	return c
}

// EnrollmentCode sets the optional parameter "enrollmentCode":
// Enrollment code of the course to create the student in. This is
// required if [userId][google.classroom.v1.Student.user_id] corresponds
// to the requesting user; this may be omitted if the requesting user
// has administrative permissions to create students for any user.
func (c *CoursesStudentsCreateCall) EnrollmentCode(enrollmentCode string) *CoursesStudentsCreateCall {
	c.opt_["enrollmentCode"] = enrollmentCode
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CoursesStudentsCreateCall) Fields(s ...googleapi.Field) *CoursesStudentsCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *CoursesStudentsCreateCall) Do() (*Student, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.student)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["enrollmentCode"]; ok {
		params.Set("enrollmentCode", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/courses/{courseId}/students")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"courseId": c.courseId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Student
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Adds a user as a student of a course. This method returns the following error codes: * `PERMISSION_DENIED` if the requesting user is not permitted to create students in this course. * `NOT_FOUND` if the requested course ID does not exist. * `ALREADY_EXISTS` if the user is already a student or student in the course.",
	//   "httpMethod": "POST",
	//   "id": "classroom.courses.students.create",
	//   "parameterOrder": [
	//     "courseId"
	//   ],
	//   "parameters": {
	//     "courseId": {
	//       "description": "Identifier of the course to create the student in. This may either be the Classroom-assigned identifier or an alias.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "enrollmentCode": {
	//       "description": "Enrollment code of the course to create the student in. This is required if [userId][google.classroom.v1.Student.user_id] corresponds to the requesting user; this may be omitted if the requesting user has administrative permissions to create students for any user.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/courses/{courseId}/students",
	//   "request": {
	//     "$ref": "Student"
	//   },
	//   "response": {
	//     "$ref": "Student"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/classroom.profile.emails",
	//     "https://www.googleapis.com/auth/classroom.profile.photos",
	//     "https://www.googleapis.com/auth/classroom.rosters"
	//   ]
	// }

}

// method id "classroom.courses.students.delete":

type CoursesStudentsDeleteCall struct {
	s        *Service
	courseId string
	userId   string
	opt_     map[string]interface{}
}

// Delete: Deletes a student of a course. This method returns the
// following error codes: * `PERMISSION_DENIED` if the requesting user
// is not permitted to delete students of this course. * `NOT_FOUND` if
// no student of this course has the requested ID or if the course does
// not exist.
func (r *CoursesStudentsService) Delete(courseId string, userId string) *CoursesStudentsDeleteCall {
	c := &CoursesStudentsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.courseId = courseId
	c.userId = userId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CoursesStudentsDeleteCall) Fields(s ...googleapi.Field) *CoursesStudentsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *CoursesStudentsDeleteCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/courses/{courseId}/students/{userId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"courseId": c.courseId,
		"userId":   c.userId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Empty
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes a student of a course. This method returns the following error codes: * `PERMISSION_DENIED` if the requesting user is not permitted to delete students of this course. * `NOT_FOUND` if no student of this course has the requested ID or if the course does not exist.",
	//   "httpMethod": "DELETE",
	//   "id": "classroom.courses.students.delete",
	//   "parameterOrder": [
	//     "courseId",
	//     "userId"
	//   ],
	//   "parameters": {
	//     "courseId": {
	//       "description": "Unique identifier of the course. This may either be the Classroom-assigned identifier or an alias.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "Identifier of the student to delete, or an alias the identifies the user. The following aliases are supported: * the e-mail address of the user * the string literal `\"me\"`, indicating that the requesting user",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/courses/{courseId}/students/{userId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/classroom.rosters"
	//   ]
	// }

}

// method id "classroom.courses.students.get":

type CoursesStudentsGetCall struct {
	s        *Service
	courseId string
	userId   string
	opt_     map[string]interface{}
}

// Get: Returns a student of a course. This method returns the following
// error codes: * `PERMISSION_DENIED` if the requesting user is not
// permitted to view students of this course. * `NOT_FOUND` if no
// student of this course has the requested ID or if the course does not
// exist.
func (r *CoursesStudentsService) Get(courseId string, userId string) *CoursesStudentsGetCall {
	c := &CoursesStudentsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.courseId = courseId
	c.userId = userId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CoursesStudentsGetCall) Fields(s ...googleapi.Field) *CoursesStudentsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *CoursesStudentsGetCall) Do() (*Student, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/courses/{courseId}/students/{userId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"courseId": c.courseId,
		"userId":   c.userId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Student
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns a student of a course. This method returns the following error codes: * `PERMISSION_DENIED` if the requesting user is not permitted to view students of this course. * `NOT_FOUND` if no student of this course has the requested ID or if the course does not exist.",
	//   "httpMethod": "GET",
	//   "id": "classroom.courses.students.get",
	//   "parameterOrder": [
	//     "courseId",
	//     "userId"
	//   ],
	//   "parameters": {
	//     "courseId": {
	//       "description": "Unique identifier of the course. This may either be the Classroom-assigned identifier or an alias.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "Identifier of the student to return, or an alias the identifies the user. The following aliases are supported: * the e-mail address of the user * the string literal `\"me\"`, indicating that the requesting user",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/courses/{courseId}/students/{userId}",
	//   "response": {
	//     "$ref": "Student"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/classroom.profile.emails",
	//     "https://www.googleapis.com/auth/classroom.profile.photos",
	//     "https://www.googleapis.com/auth/classroom.rosters",
	//     "https://www.googleapis.com/auth/classroom.rosters.readonly"
	//   ]
	// }

}

// method id "classroom.courses.students.list":

type CoursesStudentsListCall struct {
	s        *Service
	courseId string
	opt_     map[string]interface{}
}

// List: Returns a list of students of this course that the requester is
// permitted to view. Fails with `NOT_FOUND` if the course does not
// exist.
func (r *CoursesStudentsService) List(courseId string) *CoursesStudentsListCall {
	c := &CoursesStudentsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.courseId = courseId
	return c
}

// PageSize sets the optional parameter "pageSize": Maximum number of
// items to return. Zero means no maximum. The server may return fewer
// than the specified number of results.
func (c *CoursesStudentsListCall) PageSize(pageSize int64) *CoursesStudentsListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken":
// [nextPageToken][google.classroom.v1.ListStudentsResponse.next_page_tok
// en] value returned from a previous
// [list][google.classroom.v1.Users.ListStudents] call, indicating that
// the subsequent page of results should be returned. The
// [list][google.classroom.v1.Users.ListStudents] request must be
// identical to the one which resulted in this token.
func (c *CoursesStudentsListCall) PageToken(pageToken string) *CoursesStudentsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CoursesStudentsListCall) Fields(s ...googleapi.Field) *CoursesStudentsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *CoursesStudentsListCall) Do() (*ListStudentsResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/courses/{courseId}/students")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"courseId": c.courseId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListStudentsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns a list of students of this course that the requester is permitted to view. Fails with `NOT_FOUND` if the course does not exist.",
	//   "httpMethod": "GET",
	//   "id": "classroom.courses.students.list",
	//   "parameterOrder": [
	//     "courseId"
	//   ],
	//   "parameters": {
	//     "courseId": {
	//       "description": "Unique identifier of the course. This may either be the Classroom-assigned identifier or an alias.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Maximum number of items to return. Zero means no maximum. The server may return fewer than the specified number of results.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "[nextPageToken][google.classroom.v1.ListStudentsResponse.next_page_token] value returned from a previous [list][google.classroom.v1.Users.ListStudents] call, indicating that the subsequent page of results should be returned. The [list][google.classroom.v1.Users.ListStudents] request must be identical to the one which resulted in this token.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/courses/{courseId}/students",
	//   "response": {
	//     "$ref": "ListStudentsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/classroom.profile.emails",
	//     "https://www.googleapis.com/auth/classroom.profile.photos",
	//     "https://www.googleapis.com/auth/classroom.rosters",
	//     "https://www.googleapis.com/auth/classroom.rosters.readonly"
	//   ]
	// }

}

// method id "classroom.courses.teachers.create":

type CoursesTeachersCreateCall struct {
	s        *Service
	courseId string
	teacher  *Teacher
	opt_     map[string]interface{}
}

// Create: Creates a teacher of a course. This method returns the
// following error codes: * `PERMISSION_DENIED` if the requesting user
// is not permitted to create teachers in this course. * `NOT_FOUND` if
// the requested course ID does not exist. * `ALREADY_EXISTS` if the
// user is already a teacher or student in the course.
func (r *CoursesTeachersService) Create(courseId string, teacher *Teacher) *CoursesTeachersCreateCall {
	c := &CoursesTeachersCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.courseId = courseId
	c.teacher = teacher
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CoursesTeachersCreateCall) Fields(s ...googleapi.Field) *CoursesTeachersCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *CoursesTeachersCreateCall) Do() (*Teacher, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.teacher)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/courses/{courseId}/teachers")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"courseId": c.courseId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Teacher
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a teacher of a course. This method returns the following error codes: * `PERMISSION_DENIED` if the requesting user is not permitted to create teachers in this course. * `NOT_FOUND` if the requested course ID does not exist. * `ALREADY_EXISTS` if the user is already a teacher or student in the course.",
	//   "httpMethod": "POST",
	//   "id": "classroom.courses.teachers.create",
	//   "parameterOrder": [
	//     "courseId"
	//   ],
	//   "parameters": {
	//     "courseId": {
	//       "description": "Unique identifier of the course. This may either be the Classroom-assigned identifier or an alias.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/courses/{courseId}/teachers",
	//   "request": {
	//     "$ref": "Teacher"
	//   },
	//   "response": {
	//     "$ref": "Teacher"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/classroom.profile.emails",
	//     "https://www.googleapis.com/auth/classroom.profile.photos",
	//     "https://www.googleapis.com/auth/classroom.rosters"
	//   ]
	// }

}

// method id "classroom.courses.teachers.delete":

type CoursesTeachersDeleteCall struct {
	s        *Service
	courseId string
	userId   string
	opt_     map[string]interface{}
}

// Delete: Deletes a teacher of a course. This method returns the
// following error codes: * `PERMISSION_DENIED` if the requesting user
// is not permitted to delete teachers of this course. * `NOT_FOUND` if
// no teacher of this course has the requested ID or if the course does
// not exist. * `FAILED_PRECONDITION` if the requested ID belongs to the
// primary teacher of this course.
func (r *CoursesTeachersService) Delete(courseId string, userId string) *CoursesTeachersDeleteCall {
	c := &CoursesTeachersDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.courseId = courseId
	c.userId = userId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CoursesTeachersDeleteCall) Fields(s ...googleapi.Field) *CoursesTeachersDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *CoursesTeachersDeleteCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/courses/{courseId}/teachers/{userId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"courseId": c.courseId,
		"userId":   c.userId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Empty
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes a teacher of a course. This method returns the following error codes: * `PERMISSION_DENIED` if the requesting user is not permitted to delete teachers of this course. * `NOT_FOUND` if no teacher of this course has the requested ID or if the course does not exist. * `FAILED_PRECONDITION` if the requested ID belongs to the primary teacher of this course.",
	//   "httpMethod": "DELETE",
	//   "id": "classroom.courses.teachers.delete",
	//   "parameterOrder": [
	//     "courseId",
	//     "userId"
	//   ],
	//   "parameters": {
	//     "courseId": {
	//       "description": "Unique identifier of the course. This may either be the Classroom-assigned identifier or an alias.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "Identifier of the teacher to delete, or an alias the identifies the user. the following aliases are supported: * the e-mail address of the user * the string literal `\"me\"`, indicating that the requesting user",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/courses/{courseId}/teachers/{userId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/classroom.rosters"
	//   ]
	// }

}

// method id "classroom.courses.teachers.get":

type CoursesTeachersGetCall struct {
	s        *Service
	courseId string
	userId   string
	opt_     map[string]interface{}
}

// Get: Returns a teacher of a course. This method returns the following
// error codes: * `PERMISSION_DENIED` if the requesting user is not
// permitted to view teachers of this course. * `NOT_FOUND` if no
// teacher of this course has the requested ID or if the course does not
// exist.
func (r *CoursesTeachersService) Get(courseId string, userId string) *CoursesTeachersGetCall {
	c := &CoursesTeachersGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.courseId = courseId
	c.userId = userId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CoursesTeachersGetCall) Fields(s ...googleapi.Field) *CoursesTeachersGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *CoursesTeachersGetCall) Do() (*Teacher, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/courses/{courseId}/teachers/{userId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"courseId": c.courseId,
		"userId":   c.userId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Teacher
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns a teacher of a course. This method returns the following error codes: * `PERMISSION_DENIED` if the requesting user is not permitted to view teachers of this course. * `NOT_FOUND` if no teacher of this course has the requested ID or if the course does not exist.",
	//   "httpMethod": "GET",
	//   "id": "classroom.courses.teachers.get",
	//   "parameterOrder": [
	//     "courseId",
	//     "userId"
	//   ],
	//   "parameters": {
	//     "courseId": {
	//       "description": "Unique identifier of the course. This may either be the Classroom-assigned identifier or an alias.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "Identifier of the teacher to return, or an alias the identifies the user. the following aliases are supported: * the e-mail address of the user * the string literal `\"me\"`, indicating that the requesting user",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/courses/{courseId}/teachers/{userId}",
	//   "response": {
	//     "$ref": "Teacher"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/classroom.profile.emails",
	//     "https://www.googleapis.com/auth/classroom.profile.photos",
	//     "https://www.googleapis.com/auth/classroom.rosters",
	//     "https://www.googleapis.com/auth/classroom.rosters.readonly"
	//   ]
	// }

}

// method id "classroom.courses.teachers.list":

type CoursesTeachersListCall struct {
	s        *Service
	courseId string
	opt_     map[string]interface{}
}

// List: Returns a list of teachers of this course that the requester is
// permitted to view. Fails with `NOT_FOUND` if the course does not
// exist.
func (r *CoursesTeachersService) List(courseId string) *CoursesTeachersListCall {
	c := &CoursesTeachersListCall{s: r.s, opt_: make(map[string]interface{})}
	c.courseId = courseId
	return c
}

// PageSize sets the optional parameter "pageSize": Maximum number of
// items to return. Zero means no maximum. The server may return fewer
// than the specified number of results.
func (c *CoursesTeachersListCall) PageSize(pageSize int64) *CoursesTeachersListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken":
// [nextPageToken][google.classroom.v1.ListTeachersResponse.next_page_tok
// en] value returned from a previous
// [list][google.classroom.v1.Users.ListTeachers] call, indicating that
// the subsequent page of results should be returned. The
// [list][google.classroom.v1.Users.ListTeachers] request must be
// identical to the one which resulted in this token.
func (c *CoursesTeachersListCall) PageToken(pageToken string) *CoursesTeachersListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CoursesTeachersListCall) Fields(s ...googleapi.Field) *CoursesTeachersListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *CoursesTeachersListCall) Do() (*ListTeachersResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/courses/{courseId}/teachers")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"courseId": c.courseId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListTeachersResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns a list of teachers of this course that the requester is permitted to view. Fails with `NOT_FOUND` if the course does not exist.",
	//   "httpMethod": "GET",
	//   "id": "classroom.courses.teachers.list",
	//   "parameterOrder": [
	//     "courseId"
	//   ],
	//   "parameters": {
	//     "courseId": {
	//       "description": "Unique identifier of the course. This may either be the Classroom-assigned identifier or an alias.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Maximum number of items to return. Zero means no maximum. The server may return fewer than the specified number of results.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "[nextPageToken][google.classroom.v1.ListTeachersResponse.next_page_token] value returned from a previous [list][google.classroom.v1.Users.ListTeachers] call, indicating that the subsequent page of results should be returned. The [list][google.classroom.v1.Users.ListTeachers] request must be identical to the one which resulted in this token.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/courses/{courseId}/teachers",
	//   "response": {
	//     "$ref": "ListTeachersResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/classroom.profile.emails",
	//     "https://www.googleapis.com/auth/classroom.profile.photos",
	//     "https://www.googleapis.com/auth/classroom.rosters",
	//     "https://www.googleapis.com/auth/classroom.rosters.readonly"
	//   ]
	// }

}

// method id "classroom.invitations.accept":

type InvitationsAcceptCall struct {
	s    *Service
	id   string
	opt_ map[string]interface{}
}

// Accept: Accepts an invitation, removing it and adding the invited
// user to the teachers or students (as appropriate) of the specified
// course. Only the invited user may accept an invitation. This method
// returns the following error codes: * `PERMISSION_DENIED` if the
// requesting user is not permitted to accept the requested invitation.
// * `NOT_FOUND` if no invitation exists with the requested ID.
func (r *InvitationsService) Accept(id string) *InvitationsAcceptCall {
	c := &InvitationsAcceptCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InvitationsAcceptCall) Fields(s ...googleapi.Field) *InvitationsAcceptCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *InvitationsAcceptCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/invitations/{id}:accept")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Empty
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Accepts an invitation, removing it and adding the invited user to the teachers or students (as appropriate) of the specified course. Only the invited user may accept an invitation. This method returns the following error codes: * `PERMISSION_DENIED` if the requesting user is not permitted to accept the requested invitation. * `NOT_FOUND` if no invitation exists with the requested ID.",
	//   "httpMethod": "POST",
	//   "id": "classroom.invitations.accept",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Identifier of the invitation to accept.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/invitations/{id}:accept",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/classroom.rosters"
	//   ]
	// }

}

// method id "classroom.invitations.create":

type InvitationsCreateCall struct {
	s          *Service
	invitation *Invitation
	opt_       map[string]interface{}
}

// Create: Creates a invitation. Only one invitation for a user and
// course may exist at a time. Delete and recreate an invitation to make
// changes. This method returns the following error codes: *
// `PERMISSION_DENIED` if the requesting user is not permitted to create
// invitations for this course. * `NOT_FOUND` if the course or the user
// does not exist. * `ALREADY_EXISTS` if an invitation for the specified
// user and course already exists.
func (r *InvitationsService) Create(invitation *Invitation) *InvitationsCreateCall {
	c := &InvitationsCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.invitation = invitation
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InvitationsCreateCall) Fields(s ...googleapi.Field) *InvitationsCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *InvitationsCreateCall) Do() (*Invitation, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.invitation)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/invitations")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Invitation
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a invitation. Only one invitation for a user and course may exist at a time. Delete and recreate an invitation to make changes. This method returns the following error codes: * `PERMISSION_DENIED` if the requesting user is not permitted to create invitations for this course. * `NOT_FOUND` if the course or the user does not exist. * `ALREADY_EXISTS` if an invitation for the specified user and course already exists.",
	//   "httpMethod": "POST",
	//   "id": "classroom.invitations.create",
	//   "path": "v1/invitations",
	//   "request": {
	//     "$ref": "Invitation"
	//   },
	//   "response": {
	//     "$ref": "Invitation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/classroom.rosters"
	//   ]
	// }

}

// method id "classroom.invitations.delete":

type InvitationsDeleteCall struct {
	s    *Service
	id   string
	opt_ map[string]interface{}
}

// Delete: Deletes a invitation. This method returns the following error
// codes: * `PERMISSION_DENIED` if the requesting user is not permitted
// to delete the requested invitation. * `NOT_FOUND` if no invitation
// exists with the requested ID.
func (r *InvitationsService) Delete(id string) *InvitationsDeleteCall {
	c := &InvitationsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InvitationsDeleteCall) Fields(s ...googleapi.Field) *InvitationsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *InvitationsDeleteCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/invitations/{id}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Empty
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes a invitation. This method returns the following error codes: * `PERMISSION_DENIED` if the requesting user is not permitted to delete the requested invitation. * `NOT_FOUND` if no invitation exists with the requested ID.",
	//   "httpMethod": "DELETE",
	//   "id": "classroom.invitations.delete",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Identifier of the invitation to delete.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/invitations/{id}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/classroom.rosters"
	//   ]
	// }

}

// method id "classroom.invitations.get":

type InvitationsGetCall struct {
	s    *Service
	id   string
	opt_ map[string]interface{}
}

// Get: Returns a invitation. This method returns the following error
// codes: * `PERMISSION_DENIED` if the requesting user is not permitted
// to view the requested invitation. * `NOT_FOUND` if no invitation
// exists with the requested ID.
func (r *InvitationsService) Get(id string) *InvitationsGetCall {
	c := &InvitationsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InvitationsGetCall) Fields(s ...googleapi.Field) *InvitationsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *InvitationsGetCall) Do() (*Invitation, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/invitations/{id}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Invitation
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns a invitation. This method returns the following error codes: * `PERMISSION_DENIED` if the requesting user is not permitted to view the requested invitation. * `NOT_FOUND` if no invitation exists with the requested ID.",
	//   "httpMethod": "GET",
	//   "id": "classroom.invitations.get",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Identifier of the invitation to return.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/invitations/{id}",
	//   "response": {
	//     "$ref": "Invitation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/classroom.rosters",
	//     "https://www.googleapis.com/auth/classroom.rosters.readonly"
	//   ]
	// }

}

// method id "classroom.invitations.list":

type InvitationsListCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// List: Returns a list of invitations that the requesting user is
// permitted to view, restricted to those that match the request.
// *Note:* At least one of `user_id` or `course_id` must be supplied.
func (r *InvitationsService) List() *InvitationsListCall {
	c := &InvitationsListCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// CourseId sets the optional parameter "courseId": Restricts returned
// invitations to those for a course with the specified identifier.
func (c *InvitationsListCall) CourseId(courseId string) *InvitationsListCall {
	c.opt_["courseId"] = courseId
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of items to return. Zero means no maximum. The server may return
// fewer than the specified number of results.
func (c *InvitationsListCall) PageSize(pageSize int64) *InvitationsListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken":
// [nextPageToken][google.classroom.v1.ListInvitationsRespnse.next_page_t
// oken] value returned from a previous
// [list][google.classroom.v1.Users.ListInvitations] call, indicating
// that the subsequent page of results should be returned. The
// [list][google.classroom.v1.Users.ListInvitations] request must be
// identical to the one which resulted in this token.
func (c *InvitationsListCall) PageToken(pageToken string) *InvitationsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// UserId sets the optional parameter "userId": Restricts returned
// invitations to those for a specific user. This may be the unique
// identifier for the user or an alias. The supported aliases are: * the
// e-mail address of the user * the string literal "me", indicating
// the requesting user
func (c *InvitationsListCall) UserId(userId string) *InvitationsListCall {
	c.opt_["userId"] = userId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InvitationsListCall) Fields(s ...googleapi.Field) *InvitationsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *InvitationsListCall) Do() (*ListInvitationsResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["courseId"]; ok {
		params.Set("courseId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["userId"]; ok {
		params.Set("userId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/invitations")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListInvitationsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns a list of invitations that the requesting user is permitted to view, restricted to those that match the request. *Note:* At least one of `user_id` or `course_id` must be supplied.",
	//   "httpMethod": "GET",
	//   "id": "classroom.invitations.list",
	//   "parameters": {
	//     "courseId": {
	//       "description": "Restricts returned invitations to those for a course with the specified identifier.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of items to return. Zero means no maximum. The server may return fewer than the specified number of results.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "[nextPageToken][google.classroom.v1.ListInvitationsRespnse.next_page_token] value returned from a previous [list][google.classroom.v1.Users.ListInvitations] call, indicating that the subsequent page of results should be returned. The [list][google.classroom.v1.Users.ListInvitations] request must be identical to the one which resulted in this token.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "Restricts returned invitations to those for a specific user. This may be the unique identifier for the user or an alias. The supported aliases are: * the e-mail address of the user * the string literal `\"me\"`, indicating the requesting user",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/invitations",
	//   "response": {
	//     "$ref": "ListInvitationsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/classroom.rosters",
	//     "https://www.googleapis.com/auth/classroom.rosters.readonly"
	//   ]
	// }

}

// method id "classroom.userProfiles.get":

type UserProfilesGetCall struct {
	s      *Service
	userId string
	opt_   map[string]interface{}
}

// Get: Returns a user profile. This method returns the following error
// codes: * `PERMISSION_DENIED` if the requesting user is not permitted
// to access this user profile. * `NOT_FOUND` if the profile does not
// exist.
func (r *UserProfilesService) Get(userId string) *UserProfilesGetCall {
	c := &UserProfilesGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UserProfilesGetCall) Fields(s ...googleapi.Field) *UserProfilesGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *UserProfilesGetCall) Do() (*UserProfile, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/userProfiles/{userId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"userId": c.userId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *UserProfile
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns a user profile. This method returns the following error codes: * `PERMISSION_DENIED` if the requesting user is not permitted to access this user profile. * `NOT_FOUND` if the profile does not exist.",
	//   "httpMethod": "GET",
	//   "id": "classroom.userProfiles.get",
	//   "parameterOrder": [
	//     "userId"
	//   ],
	//   "parameters": {
	//     "userId": {
	//       "description": "Identifier of the profile to return, or an alias the identifies the user. The following aliases are supported: * the e-mail address of the user * the string literal `\"me\"`, indicating the requesting user",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/userProfiles/{userId}",
	//   "response": {
	//     "$ref": "UserProfile"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/classroom.profile.emails",
	//     "https://www.googleapis.com/auth/classroom.profile.photos",
	//     "https://www.googleapis.com/auth/classroom.rosters",
	//     "https://www.googleapis.com/auth/classroom.rosters.readonly"
	//   ]
	// }

}
