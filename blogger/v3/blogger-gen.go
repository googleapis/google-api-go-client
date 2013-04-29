// Package blogger provides access to the Blogger API.
//
// See https://developers.google.com/blogger/docs/3.0/getting_started
//
// Usage example:
//
//   import "code.google.com/p/google-api-go-client/blogger/v3"
//   ...
//   bloggerService, err := blogger.New(oauthHttpClient)
package blogger

import (
	"bytes"
	"code.google.com/p/google-api-go-client/googleapi"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = googleapi.Version
var _ = errors.New

const apiId = "blogger:v3"
const apiName = "blogger"
const apiVersion = "v3"
const basePath = "https://www.googleapis.com/blogger/v3/"

// OAuth2 scopes used by this API.
const (
	// Manage your Blogger account
	BloggerScope = "https://www.googleapis.com/auth/blogger"

	// View your Blogger account
	BloggerReadonlyScope = "https://www.googleapis.com/auth/blogger.readonly"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Blogs = NewBlogsService(s)
	s.Comments = NewCommentsService(s)
	s.Pages = NewPagesService(s)
	s.Posts = NewPostsService(s)
	s.Users = NewUsersService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Blogs *BlogsService

	Comments *CommentsService

	Pages *PagesService

	Posts *PostsService

	Users *UsersService
}

func NewBlogsService(s *Service) *BlogsService {
	rs := &BlogsService{s: s}
	return rs
}

type BlogsService struct {
	s *Service
}

func NewCommentsService(s *Service) *CommentsService {
	rs := &CommentsService{s: s}
	return rs
}

type CommentsService struct {
	s *Service
}

func NewPagesService(s *Service) *PagesService {
	rs := &PagesService{s: s}
	return rs
}

type PagesService struct {
	s *Service
}

func NewPostsService(s *Service) *PostsService {
	rs := &PostsService{s: s}
	return rs
}

type PostsService struct {
	s *Service
}

func NewUsersService(s *Service) *UsersService {
	rs := &UsersService{s: s}
	return rs
}

type UsersService struct {
	s *Service
}

type Blog struct {
	// CustomMetaData: The JSON custom meta-data for the Blog
	CustomMetaData string `json:"customMetaData,omitempty"`

	// Description: The description of this blog. This is displayed
	// underneath the title.
	Description string `json:"description,omitempty"`

	// Id: The identifier for this resource.
	Id string `json:"id,omitempty"`

	// Kind: The kind of this entry. Always blogger#blog
	Kind string `json:"kind,omitempty"`

	// Locale: The locale this Blog is set to.
	Locale *BlogLocale `json:"locale,omitempty"`

	// Name: The name of this blog. This is displayed as the title.
	Name string `json:"name,omitempty"`

	// Pages: The container of pages in this blog.
	Pages *BlogPages `json:"pages,omitempty"`

	// Posts: The container of posts in this blog.
	Posts *BlogPosts `json:"posts,omitempty"`

	// Published: RFC 3339 date-time when this blog was published.
	Published string `json:"published,omitempty"`

	// SelfLink: The API REST URL to fetch this resource from.
	SelfLink string `json:"selfLink,omitempty"`

	// Updated: RFC 3339 date-time when this blog was last updated.
	Updated string `json:"updated,omitempty"`

	// Url: The URL where this blog is published.
	Url string `json:"url,omitempty"`
}

type BlogLocale struct {
	// Country: The country this blog's locale is set to.
	Country string `json:"country,omitempty"`

	// Language: The language this blog is authored in.
	Language string `json:"language,omitempty"`

	// Variant: The language variant this blog is authored in.
	Variant string `json:"variant,omitempty"`
}

type BlogPages struct {
	// SelfLink: The URL of the container for pages in this blog.
	SelfLink string `json:"selfLink,omitempty"`

	// TotalItems: The count of pages in this blog.
	TotalItems int64 `json:"totalItems,omitempty"`
}

type BlogPosts struct {
	// Items: The List of Posts for this Blog.
	Items []*Post `json:"items,omitempty"`

	// SelfLink: The URL of the container for posts in this blog.
	SelfLink string `json:"selfLink,omitempty"`

	// TotalItems: The count of posts in this blog.
	TotalItems int64 `json:"totalItems,omitempty"`
}

type BlogList struct {
	// Items: The list of Blogs this user has Authorship or Admin rights
	// over.
	Items []*Blog `json:"items,omitempty"`

	// Kind: The kind of this entity. Always blogger#blogList
	Kind string `json:"kind,omitempty"`
}

type Comment struct {
	// Author: The author of this Comment.
	Author *CommentAuthor `json:"author,omitempty"`

	// Blog: Data about the blog containing this comment.
	Blog *CommentBlog `json:"blog,omitempty"`

	// Content: The actual content of the comment. May include HTML markup.
	Content string `json:"content,omitempty"`

	// Id: The identifier for this resource.
	Id string `json:"id,omitempty"`

	// InReplyTo: Data about the comment this is in reply to.
	InReplyTo *CommentInReplyTo `json:"inReplyTo,omitempty"`

	// Kind: The kind of this entry. Always blogger#comment
	Kind string `json:"kind,omitempty"`

	// Post: Data about the post containing this comment.
	Post *CommentPost `json:"post,omitempty"`

	// Published: RFC 3339 date-time when this comment was published.
	Published string `json:"published,omitempty"`

	// SelfLink: The API REST URL to fetch this resource from.
	SelfLink string `json:"selfLink,omitempty"`

	// Updated: RFC 3339 date-time when this comment was last updated.
	Updated string `json:"updated,omitempty"`
}

type CommentAuthor struct {
	// DisplayName: The display name.
	DisplayName string `json:"displayName,omitempty"`

	// Id: The identifier of the Comment creator.
	Id string `json:"id,omitempty"`

	// Image: The comment creator's avatar.
	Image *CommentAuthorImage `json:"image,omitempty"`

	// Url: The URL of the Comment creator's Profile page.
	Url string `json:"url,omitempty"`
}

type CommentAuthorImage struct {
	// Url: The comment creator's avatar URL.
	Url string `json:"url,omitempty"`
}

type CommentBlog struct {
	// Id: The identifier of the blog containing this comment.
	Id string `json:"id,omitempty"`
}

type CommentInReplyTo struct {
	// Id: The identified of the parent of this comment.
	Id string `json:"id,omitempty"`
}

type CommentPost struct {
	// Id: The identifier of the post containing this comment.
	Id string `json:"id,omitempty"`
}

type CommentList struct {
	// Items: The List of Comments for a Post.
	Items []*Comment `json:"items,omitempty"`

	// Kind: The kind of this entry. Always blogger#commentList
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Pagination token to fetch the next page, if one
	// exists.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// PrevPageToken: Pagination token to fetch the previous page, if one
	// exists.
	PrevPageToken string `json:"prevPageToken,omitempty"`
}

type Page struct {
	// Author: The author of this Page.
	Author *PageAuthor `json:"author,omitempty"`

	// Blog: Data about the blog containing this Page.
	Blog *PageBlog `json:"blog,omitempty"`

	// Content: The body content of this Page, in HTML.
	Content string `json:"content,omitempty"`

	// Id: The identifier for this resource.
	Id string `json:"id,omitempty"`

	// Kind: The kind of this entity. Always blogger#page
	Kind string `json:"kind,omitempty"`

	// Published: RFC 3339 date-time when this Page was published.
	Published string `json:"published,omitempty"`

	// SelfLink: The API REST URL to fetch this resource from.
	SelfLink string `json:"selfLink,omitempty"`

	// Title: The title of this entity. This is the name displayed in the
	// Admin user interface.
	Title string `json:"title,omitempty"`

	// Updated: RFC 3339 date-time when this Page was last updated.
	Updated string `json:"updated,omitempty"`

	// Url: The URL that this Page is displayed at.
	Url string `json:"url,omitempty"`
}

type PageAuthor struct {
	// DisplayName: The display name.
	DisplayName string `json:"displayName,omitempty"`

	// Id: The identifier of the Page creator.
	Id string `json:"id,omitempty"`

	// Image: The page author's avatar.
	Image *PageAuthorImage `json:"image,omitempty"`

	// Url: The URL of the Page creator's Profile page.
	Url string `json:"url,omitempty"`
}

type PageAuthorImage struct {
	// Url: The page author's avatar URL.
	Url string `json:"url,omitempty"`
}

type PageBlog struct {
	// Id: The identifier of the blog containing this page.
	Id string `json:"id,omitempty"`
}

type PageList struct {
	// Items: The list of Pages for a Blog.
	Items []*Page `json:"items,omitempty"`

	// Kind: The kind of this entity. Always blogger#pageList
	Kind string `json:"kind,omitempty"`
}

type Post struct {
	// Author: The author of this Post.
	Author *PostAuthor `json:"author,omitempty"`

	// Blog: Data about the blog containing this Post.
	Blog *PostBlog `json:"blog,omitempty"`

	// Content: The content of the Post. May contain HTML markup.
	Content string `json:"content,omitempty"`

	// CustomMetaData: The JSON meta-data for the Post.
	CustomMetaData string `json:"customMetaData,omitempty"`

	// Id: The identifier of this Post.
	Id string `json:"id,omitempty"`

	// Kind: The kind of this entity. Always blogger#post
	Kind string `json:"kind,omitempty"`

	// Labels: The list of labels this Post was tagged with.
	Labels []string `json:"labels,omitempty"`

	// Location: The location for geotagged posts.
	Location *PostLocation `json:"location,omitempty"`

	// Published: RFC 3339 date-time when this Post was published.
	Published string `json:"published,omitempty"`

	// Replies: The container of comments on this Post.
	Replies *PostReplies `json:"replies,omitempty"`

	// SelfLink: The API REST URL to fetch this resource from.
	SelfLink string `json:"selfLink,omitempty"`

	// Title: The title of the Post.
	Title string `json:"title,omitempty"`

	// Updated: RFC 3339 date-time when this Post was last updated.
	Updated string `json:"updated,omitempty"`

	// Url: The URL where this Post is displayed.
	Url string `json:"url,omitempty"`
}

type PostAuthor struct {
	// DisplayName: The display name.
	DisplayName string `json:"displayName,omitempty"`

	// Id: The identifier of the Post creator.
	Id string `json:"id,omitempty"`

	// Image: The Post author's avatar.
	Image *PostAuthorImage `json:"image,omitempty"`

	// Url: The URL of the Post creator's Profile page.
	Url string `json:"url,omitempty"`
}

type PostAuthorImage struct {
	// Url: The Post author's avatar URL.
	Url string `json:"url,omitempty"`
}

type PostBlog struct {
	// Id: The identifier of the Blog that contains this Post.
	Id string `json:"id,omitempty"`
}

type PostLocation struct {
	// Lat: Location's latitude.
	Lat float64 `json:"lat,omitempty"`

	// Lng: Location's longitude.
	Lng float64 `json:"lng,omitempty"`

	// Name: Location name.
	Name string `json:"name,omitempty"`

	// Span: Location's viewport span. Can be used when rendering a map
	// preview.
	Span string `json:"span,omitempty"`
}

type PostReplies struct {
	// Items: The List of Comments for this Post.
	Items []*Comment `json:"items,omitempty"`

	// SelfLink: The URL of the comments on this post.
	SelfLink string `json:"selfLink,omitempty"`

	// TotalItems: The count of comments on this post.
	TotalItems int64 `json:"totalItems,omitempty,string"`
}

type PostList struct {
	// Items: The list of Posts for this Blog.
	Items []*Post `json:"items,omitempty"`

	// Kind: The kind of this entity. Always blogger#postList
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Pagination token to fetch the next page, if one
	// exists.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// PrevPageToken: Pagination token to fetch the previous page, if one
	// exists.
	PrevPageToken string `json:"prevPageToken,omitempty"`
}

type User struct {
	// About: Profile summary information.
	About string `json:"about,omitempty"`

	// Blogs: The container of blogs for this user.
	Blogs *UserBlogs `json:"blogs,omitempty"`

	// Created: The timestamp of when this profile was created, in seconds
	// since epoch.
	Created string `json:"created,omitempty"`

	// DisplayName: The display name.
	DisplayName string `json:"displayName,omitempty"`

	// Id: The identifier for this User.
	Id string `json:"id,omitempty"`

	// Kind: The kind of this entity. Always blogger#user
	Kind string `json:"kind,omitempty"`

	// Locale: This user's locale
	Locale *UserLocale `json:"locale,omitempty"`

	// SelfLink: The API REST URL to fetch this resource from.
	SelfLink string `json:"selfLink,omitempty"`

	// Url: The user's profile page.
	Url string `json:"url,omitempty"`
}

type UserBlogs struct {
	// SelfLink: The URL of the Blogs for this user.
	SelfLink string `json:"selfLink,omitempty"`
}

type UserLocale struct {
	// Country: The user's country setting.
	Country string `json:"country,omitempty"`

	// Language: The user's language setting.
	Language string `json:"language,omitempty"`

	// Variant: The user's language variant setting.
	Variant string `json:"variant,omitempty"`
}

// method id "blogger.blogs.get":

type BlogsGetCall struct {
	s      *Service
	blogId string
	opt_   map[string]interface{}
}

// Get: Gets one blog by id.
func (r *BlogsService) Get(blogId string) *BlogsGetCall {
	c := &BlogsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.blogId = blogId
	return c
}

// MaxPosts sets the optional parameter "maxPosts": Maximum number of
// posts to pull back with the blog.
func (c *BlogsGetCall) MaxPosts(maxPosts int64) *BlogsGetCall {
	c.opt_["maxPosts"] = maxPosts
	return c
}

func (c *BlogsGetCall) Do() (*Blog, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["maxPosts"]; ok {
		params.Set("maxPosts", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/blogger/v3/", "blogs/{blogId}")
	urls = strings.Replace(urls, "{blogId}", cleanPathString(c.blogId), 1)
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(Blog)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets one blog by id.",
	//   "httpMethod": "GET",
	//   "id": "blogger.blogs.get",
	//   "parameterOrder": [
	//     "blogId"
	//   ],
	//   "parameters": {
	//     "blogId": {
	//       "description": "The ID of the blog to get.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxPosts": {
	//       "description": "Maximum number of posts to pull back with the blog.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     }
	//   },
	//   "path": "blogs/{blogId}",
	//   "response": {
	//     "$ref": "Blog"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger",
	//     "https://www.googleapis.com/auth/blogger.readonly"
	//   ]
	// }

}

// method id "blogger.blogs.getByUrl":

type BlogsGetByUrlCall struct {
	s    *Service
	url  string
	opt_ map[string]interface{}
}

// GetByUrl: Retrieve a Blog by URL.
func (r *BlogsService) GetByUrl(url string) *BlogsGetByUrlCall {
	c := &BlogsGetByUrlCall{s: r.s, opt_: make(map[string]interface{})}
	c.url = url
	return c
}

func (c *BlogsGetByUrlCall) Do() (*Blog, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("url", fmt.Sprintf("%v", c.url))
	urls := googleapi.ResolveRelative("https://www.googleapis.com/blogger/v3/", "blogs/byurl")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(Blog)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieve a Blog by URL.",
	//   "httpMethod": "GET",
	//   "id": "blogger.blogs.getByUrl",
	//   "parameterOrder": [
	//     "url"
	//   ],
	//   "parameters": {
	//     "url": {
	//       "description": "The URL of the blog to retrieve.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "blogs/byurl",
	//   "response": {
	//     "$ref": "Blog"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger",
	//     "https://www.googleapis.com/auth/blogger.readonly"
	//   ]
	// }

}

// method id "blogger.blogs.listByUser":

type BlogsListByUserCall struct {
	s      *Service
	userId string
	opt_   map[string]interface{}
}

// ListByUser: Retrieves a list of blogs, possibly filtered.
func (r *BlogsService) ListByUser(userId string) *BlogsListByUserCall {
	c := &BlogsListByUserCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	return c
}

func (c *BlogsListByUserCall) Do() (*BlogList, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/blogger/v3/", "users/{userId}/blogs")
	urls = strings.Replace(urls, "{userId}", cleanPathString(c.userId), 1)
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(BlogList)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a list of blogs, possibly filtered.",
	//   "httpMethod": "GET",
	//   "id": "blogger.blogs.listByUser",
	//   "parameterOrder": [
	//     "userId"
	//   ],
	//   "parameters": {
	//     "userId": {
	//       "description": "ID of the user whose blogs are to be fetched. Either the word 'self' (sans quote marks) or the user's profile identifier.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/{userId}/blogs",
	//   "response": {
	//     "$ref": "BlogList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger",
	//     "https://www.googleapis.com/auth/blogger.readonly"
	//   ]
	// }

}

// method id "blogger.comments.get":

type CommentsGetCall struct {
	s         *Service
	blogId    string
	postId    string
	commentId string
	opt_      map[string]interface{}
}

// Get: Gets one comment by id.
func (r *CommentsService) Get(blogId string, postId string, commentId string) *CommentsGetCall {
	c := &CommentsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.blogId = blogId
	c.postId = postId
	c.commentId = commentId
	return c
}

func (c *CommentsGetCall) Do() (*Comment, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/blogger/v3/", "blogs/{blogId}/posts/{postId}/comments/{commentId}")
	urls = strings.Replace(urls, "{blogId}", cleanPathString(c.blogId), 1)
	urls = strings.Replace(urls, "{postId}", cleanPathString(c.postId), 1)
	urls = strings.Replace(urls, "{commentId}", cleanPathString(c.commentId), 1)
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(Comment)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets one comment by id.",
	//   "httpMethod": "GET",
	//   "id": "blogger.comments.get",
	//   "parameterOrder": [
	//     "blogId",
	//     "postId",
	//     "commentId"
	//   ],
	//   "parameters": {
	//     "blogId": {
	//       "description": "ID of the blog to containing the comment.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "commentId": {
	//       "description": "The ID of the comment to get.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "postId": {
	//       "description": "ID of the post to fetch posts from.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "blogs/{blogId}/posts/{postId}/comments/{commentId}",
	//   "response": {
	//     "$ref": "Comment"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger",
	//     "https://www.googleapis.com/auth/blogger.readonly"
	//   ]
	// }

}

// method id "blogger.comments.list":

type CommentsListCall struct {
	s      *Service
	blogId string
	postId string
	opt_   map[string]interface{}
}

// List: Retrieves the comments for a blog, possibly filtered.
func (r *CommentsService) List(blogId string, postId string) *CommentsListCall {
	c := &CommentsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.blogId = blogId
	c.postId = postId
	return c
}

// EndDate sets the optional parameter "endDate": Latest date of comment
// to fetch, a date-time with RFC 3339 formatting.
func (c *CommentsListCall) EndDate(endDate string) *CommentsListCall {
	c.opt_["endDate"] = endDate
	return c
}

// FetchBodies sets the optional parameter "fetchBodies": Whether the
// body content of the comments is included.
func (c *CommentsListCall) FetchBodies(fetchBodies bool) *CommentsListCall {
	c.opt_["fetchBodies"] = fetchBodies
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of comments to include in the result.
func (c *CommentsListCall) MaxResults(maxResults int64) *CommentsListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": Continuation token
// if request is paged.
func (c *CommentsListCall) PageToken(pageToken string) *CommentsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// StartDate sets the optional parameter "startDate": Earliest date of
// comment to fetch, a date-time with RFC 3339 formatting.
func (c *CommentsListCall) StartDate(startDate string) *CommentsListCall {
	c.opt_["startDate"] = startDate
	return c
}

func (c *CommentsListCall) Do() (*CommentList, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["endDate"]; ok {
		params.Set("endDate", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fetchBodies"]; ok {
		params.Set("fetchBodies", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["startDate"]; ok {
		params.Set("startDate", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/blogger/v3/", "blogs/{blogId}/posts/{postId}/comments")
	urls = strings.Replace(urls, "{blogId}", cleanPathString(c.blogId), 1)
	urls = strings.Replace(urls, "{postId}", cleanPathString(c.postId), 1)
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(CommentList)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the comments for a blog, possibly filtered.",
	//   "httpMethod": "GET",
	//   "id": "blogger.comments.list",
	//   "parameterOrder": [
	//     "blogId",
	//     "postId"
	//   ],
	//   "parameters": {
	//     "blogId": {
	//       "description": "ID of the blog to fetch comments from.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "endDate": {
	//       "description": "Latest date of comment to fetch, a date-time with RFC 3339 formatting.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "fetchBodies": {
	//       "description": "Whether the body content of the comments is included.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of comments to include in the result.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Continuation token if request is paged.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "postId": {
	//       "description": "ID of the post to fetch posts from.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "startDate": {
	//       "description": "Earliest date of comment to fetch, a date-time with RFC 3339 formatting.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "blogs/{blogId}/posts/{postId}/comments",
	//   "response": {
	//     "$ref": "CommentList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger",
	//     "https://www.googleapis.com/auth/blogger.readonly"
	//   ]
	// }

}

// method id "blogger.pages.get":

type PagesGetCall struct {
	s      *Service
	blogId string
	pageId string
	opt_   map[string]interface{}
}

// Get: Gets one blog page by id.
func (r *PagesService) Get(blogId string, pageId string) *PagesGetCall {
	c := &PagesGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.blogId = blogId
	c.pageId = pageId
	return c
}

func (c *PagesGetCall) Do() (*Page, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/blogger/v3/", "blogs/{blogId}/pages/{pageId}")
	urls = strings.Replace(urls, "{blogId}", cleanPathString(c.blogId), 1)
	urls = strings.Replace(urls, "{pageId}", cleanPathString(c.pageId), 1)
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(Page)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets one blog page by id.",
	//   "httpMethod": "GET",
	//   "id": "blogger.pages.get",
	//   "parameterOrder": [
	//     "blogId",
	//     "pageId"
	//   ],
	//   "parameters": {
	//     "blogId": {
	//       "description": "ID of the blog containing the page.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageId": {
	//       "description": "The ID of the page to get.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "blogs/{blogId}/pages/{pageId}",
	//   "response": {
	//     "$ref": "Page"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger",
	//     "https://www.googleapis.com/auth/blogger.readonly"
	//   ]
	// }

}

// method id "blogger.pages.list":

type PagesListCall struct {
	s      *Service
	blogId string
	opt_   map[string]interface{}
}

// List: Retrieves pages for a blog, possibly filtered.
func (r *PagesService) List(blogId string) *PagesListCall {
	c := &PagesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.blogId = blogId
	return c
}

// FetchBodies sets the optional parameter "fetchBodies": Whether to
// retrieve the Page bodies.
func (c *PagesListCall) FetchBodies(fetchBodies bool) *PagesListCall {
	c.opt_["fetchBodies"] = fetchBodies
	return c
}

func (c *PagesListCall) Do() (*PageList, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fetchBodies"]; ok {
		params.Set("fetchBodies", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/blogger/v3/", "blogs/{blogId}/pages")
	urls = strings.Replace(urls, "{blogId}", cleanPathString(c.blogId), 1)
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(PageList)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves pages for a blog, possibly filtered.",
	//   "httpMethod": "GET",
	//   "id": "blogger.pages.list",
	//   "parameterOrder": [
	//     "blogId"
	//   ],
	//   "parameters": {
	//     "blogId": {
	//       "description": "ID of the blog to fetch pages from.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "fetchBodies": {
	//       "description": "Whether to retrieve the Page bodies.",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "blogs/{blogId}/pages",
	//   "response": {
	//     "$ref": "PageList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger",
	//     "https://www.googleapis.com/auth/blogger.readonly"
	//   ]
	// }

}

// method id "blogger.posts.delete":

type PostsDeleteCall struct {
	s      *Service
	blogId string
	postId string
	opt_   map[string]interface{}
}

// Delete: Delete a post by id.
func (r *PostsService) Delete(blogId string, postId string) *PostsDeleteCall {
	c := &PostsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.blogId = blogId
	c.postId = postId
	return c
}

func (c *PostsDeleteCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/blogger/v3/", "blogs/{blogId}/posts/{postId}")
	urls = strings.Replace(urls, "{blogId}", cleanPathString(c.blogId), 1)
	urls = strings.Replace(urls, "{postId}", cleanPathString(c.postId), 1)
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Delete a post by id.",
	//   "httpMethod": "DELETE",
	//   "id": "blogger.posts.delete",
	//   "parameterOrder": [
	//     "blogId",
	//     "postId"
	//   ],
	//   "parameters": {
	//     "blogId": {
	//       "description": "The Id of the Blog.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "postId": {
	//       "description": "The ID of the Post.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "blogs/{blogId}/posts/{postId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger"
	//   ]
	// }

}

// method id "blogger.posts.get":

type PostsGetCall struct {
	s      *Service
	blogId string
	postId string
	opt_   map[string]interface{}
}

// Get: Get a post by id.
func (r *PostsService) Get(blogId string, postId string) *PostsGetCall {
	c := &PostsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.blogId = blogId
	c.postId = postId
	return c
}

// MaxComments sets the optional parameter "maxComments": Maximum number
// of comments to pull back on a post.
func (c *PostsGetCall) MaxComments(maxComments int64) *PostsGetCall {
	c.opt_["maxComments"] = maxComments
	return c
}

func (c *PostsGetCall) Do() (*Post, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["maxComments"]; ok {
		params.Set("maxComments", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/blogger/v3/", "blogs/{blogId}/posts/{postId}")
	urls = strings.Replace(urls, "{blogId}", cleanPathString(c.blogId), 1)
	urls = strings.Replace(urls, "{postId}", cleanPathString(c.postId), 1)
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(Post)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get a post by id.",
	//   "httpMethod": "GET",
	//   "id": "blogger.posts.get",
	//   "parameterOrder": [
	//     "blogId",
	//     "postId"
	//   ],
	//   "parameters": {
	//     "blogId": {
	//       "description": "ID of the blog to fetch the post from.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxComments": {
	//       "description": "Maximum number of comments to pull back on a post.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "postId": {
	//       "description": "The ID of the post",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "blogs/{blogId}/posts/{postId}",
	//   "response": {
	//     "$ref": "Post"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger",
	//     "https://www.googleapis.com/auth/blogger.readonly"
	//   ]
	// }

}

// method id "blogger.posts.getByPath":

type PostsGetByPathCall struct {
	s      *Service
	blogId string
	path   string
	opt_   map[string]interface{}
}

// GetByPath: Retrieve a Post by Path.
func (r *PostsService) GetByPath(blogId string, path string) *PostsGetByPathCall {
	c := &PostsGetByPathCall{s: r.s, opt_: make(map[string]interface{})}
	c.blogId = blogId
	c.path = path
	return c
}

// MaxComments sets the optional parameter "maxComments": Maximum number
// of comments to pull back on a post.
func (c *PostsGetByPathCall) MaxComments(maxComments int64) *PostsGetByPathCall {
	c.opt_["maxComments"] = maxComments
	return c
}

func (c *PostsGetByPathCall) Do() (*Post, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("path", fmt.Sprintf("%v", c.path))
	if v, ok := c.opt_["maxComments"]; ok {
		params.Set("maxComments", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/blogger/v3/", "blogs/{blogId}/posts/bypath")
	urls = strings.Replace(urls, "{blogId}", cleanPathString(c.blogId), 1)
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(Post)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieve a Post by Path.",
	//   "httpMethod": "GET",
	//   "id": "blogger.posts.getByPath",
	//   "parameterOrder": [
	//     "blogId",
	//     "path"
	//   ],
	//   "parameters": {
	//     "blogId": {
	//       "description": "ID of the blog to fetch the post from.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxComments": {
	//       "description": "Maximum number of comments to pull back on a post.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "path": {
	//       "description": "Path of the Post to retrieve.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "blogs/{blogId}/posts/bypath",
	//   "response": {
	//     "$ref": "Post"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger",
	//     "https://www.googleapis.com/auth/blogger.readonly"
	//   ]
	// }

}

// method id "blogger.posts.insert":

type PostsInsertCall struct {
	s      *Service
	blogId string
	post   *Post
	opt_   map[string]interface{}
}

// Insert: Add a post.
func (r *PostsService) Insert(blogId string, post *Post) *PostsInsertCall {
	c := &PostsInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.blogId = blogId
	c.post = post
	return c
}

func (c *PostsInsertCall) Do() (*Post, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.post)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/blogger/v3/", "blogs/{blogId}/posts")
	urls = strings.Replace(urls, "{blogId}", cleanPathString(c.blogId), 1)
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(Post)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Add a post.",
	//   "httpMethod": "POST",
	//   "id": "blogger.posts.insert",
	//   "parameterOrder": [
	//     "blogId"
	//   ],
	//   "parameters": {
	//     "blogId": {
	//       "description": "ID of the blog to fetch the post from.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "blogs/{blogId}/posts",
	//   "request": {
	//     "$ref": "Post"
	//   },
	//   "response": {
	//     "$ref": "Post"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger"
	//   ]
	// }

}

// method id "blogger.posts.list":

type PostsListCall struct {
	s      *Service
	blogId string
	opt_   map[string]interface{}
}

// List: Retrieves a list of posts, possibly filtered.
func (r *PostsService) List(blogId string) *PostsListCall {
	c := &PostsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.blogId = blogId
	return c
}

// EndDate sets the optional parameter "endDate": Latest post date to
// fetch, a date-time with RFC 3339 formatting.
func (c *PostsListCall) EndDate(endDate string) *PostsListCall {
	c.opt_["endDate"] = endDate
	return c
}

// FetchBodies sets the optional parameter "fetchBodies": Whether the
// body content of posts is included.
func (c *PostsListCall) FetchBodies(fetchBodies bool) *PostsListCall {
	c.opt_["fetchBodies"] = fetchBodies
	return c
}

// Labels sets the optional parameter "labels": Comma-separated list of
// labels to search for.
func (c *PostsListCall) Labels(labels string) *PostsListCall {
	c.opt_["labels"] = labels
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of posts to fetch.
func (c *PostsListCall) MaxResults(maxResults int64) *PostsListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": Continuation token
// if the request is paged.
func (c *PostsListCall) PageToken(pageToken string) *PostsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// StartDate sets the optional parameter "startDate": Earliest post date
// to fetch, a date-time with RFC 3339 formatting.
func (c *PostsListCall) StartDate(startDate string) *PostsListCall {
	c.opt_["startDate"] = startDate
	return c
}

func (c *PostsListCall) Do() (*PostList, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["endDate"]; ok {
		params.Set("endDate", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fetchBodies"]; ok {
		params.Set("fetchBodies", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["labels"]; ok {
		params.Set("labels", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["startDate"]; ok {
		params.Set("startDate", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/blogger/v3/", "blogs/{blogId}/posts")
	urls = strings.Replace(urls, "{blogId}", cleanPathString(c.blogId), 1)
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(PostList)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a list of posts, possibly filtered.",
	//   "httpMethod": "GET",
	//   "id": "blogger.posts.list",
	//   "parameterOrder": [
	//     "blogId"
	//   ],
	//   "parameters": {
	//     "blogId": {
	//       "description": "ID of the blog to fetch posts from.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "endDate": {
	//       "description": "Latest post date to fetch, a date-time with RFC 3339 formatting.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "fetchBodies": {
	//       "description": "Whether the body content of posts is included.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "labels": {
	//       "description": "Comma-separated list of labels to search for.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of posts to fetch.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Continuation token if the request is paged.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "startDate": {
	//       "description": "Earliest post date to fetch, a date-time with RFC 3339 formatting.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "blogs/{blogId}/posts",
	//   "response": {
	//     "$ref": "PostList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger",
	//     "https://www.googleapis.com/auth/blogger.readonly"
	//   ]
	// }

}

// method id "blogger.posts.patch":

type PostsPatchCall struct {
	s      *Service
	blogId string
	postId string
	post   *Post
	opt_   map[string]interface{}
}

// Patch: Update a post. This method supports patch semantics.
func (r *PostsService) Patch(blogId string, postId string, post *Post) *PostsPatchCall {
	c := &PostsPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.blogId = blogId
	c.postId = postId
	c.post = post
	return c
}

func (c *PostsPatchCall) Do() (*Post, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.post)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/blogger/v3/", "blogs/{blogId}/posts/{postId}")
	urls = strings.Replace(urls, "{blogId}", cleanPathString(c.blogId), 1)
	urls = strings.Replace(urls, "{postId}", cleanPathString(c.postId), 1)
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(Post)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update a post. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "blogger.posts.patch",
	//   "parameterOrder": [
	//     "blogId",
	//     "postId"
	//   ],
	//   "parameters": {
	//     "blogId": {
	//       "description": "The ID of the Blog.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "postId": {
	//       "description": "The ID of the Post.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "blogs/{blogId}/posts/{postId}",
	//   "request": {
	//     "$ref": "Post"
	//   },
	//   "response": {
	//     "$ref": "Post"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger"
	//   ]
	// }

}

// method id "blogger.posts.search":

type PostsSearchCall struct {
	s      *Service
	blogId string
	q      string
	opt_   map[string]interface{}
}

// Search: Search for a post.
func (r *PostsService) Search(blogId string, q string) *PostsSearchCall {
	c := &PostsSearchCall{s: r.s, opt_: make(map[string]interface{})}
	c.blogId = blogId
	c.q = q
	return c
}

func (c *PostsSearchCall) Do() (*PostList, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("q", fmt.Sprintf("%v", c.q))
	urls := googleapi.ResolveRelative("https://www.googleapis.com/blogger/v3/", "blogs/{blogId}/posts/search")
	urls = strings.Replace(urls, "{blogId}", cleanPathString(c.blogId), 1)
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(PostList)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Search for a post.",
	//   "httpMethod": "GET",
	//   "id": "blogger.posts.search",
	//   "parameterOrder": [
	//     "blogId",
	//     "q"
	//   ],
	//   "parameters": {
	//     "blogId": {
	//       "description": "ID of the blog to fetch the post from.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "q": {
	//       "description": "Query terms to search this blog for matching posts.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "blogs/{blogId}/posts/search",
	//   "response": {
	//     "$ref": "PostList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger",
	//     "https://www.googleapis.com/auth/blogger.readonly"
	//   ]
	// }

}

// method id "blogger.posts.update":

type PostsUpdateCall struct {
	s      *Service
	blogId string
	postId string
	post   *Post
	opt_   map[string]interface{}
}

// Update: Update a post.
func (r *PostsService) Update(blogId string, postId string, post *Post) *PostsUpdateCall {
	c := &PostsUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.blogId = blogId
	c.postId = postId
	c.post = post
	return c
}

func (c *PostsUpdateCall) Do() (*Post, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.post)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/blogger/v3/", "blogs/{blogId}/posts/{postId}")
	urls = strings.Replace(urls, "{blogId}", cleanPathString(c.blogId), 1)
	urls = strings.Replace(urls, "{postId}", cleanPathString(c.postId), 1)
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(Post)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update a post.",
	//   "httpMethod": "PUT",
	//   "id": "blogger.posts.update",
	//   "parameterOrder": [
	//     "blogId",
	//     "postId"
	//   ],
	//   "parameters": {
	//     "blogId": {
	//       "description": "The ID of the Blog.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "postId": {
	//       "description": "The ID of the Post.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "blogs/{blogId}/posts/{postId}",
	//   "request": {
	//     "$ref": "Post"
	//   },
	//   "response": {
	//     "$ref": "Post"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger"
	//   ]
	// }

}

// method id "blogger.users.get":

type UsersGetCall struct {
	s      *Service
	userId string
	opt_   map[string]interface{}
}

// Get: Gets one user by id.
func (r *UsersService) Get(userId string) *UsersGetCall {
	c := &UsersGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	return c
}

func (c *UsersGetCall) Do() (*User, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/blogger/v3/", "users/{userId}")
	urls = strings.Replace(urls, "{userId}", cleanPathString(c.userId), 1)
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(User)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets one user by id.",
	//   "httpMethod": "GET",
	//   "id": "blogger.users.get",
	//   "parameterOrder": [
	//     "userId"
	//   ],
	//   "parameters": {
	//     "userId": {
	//       "description": "The ID of the user to get.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/{userId}",
	//   "response": {
	//     "$ref": "User"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/blogger",
	//     "https://www.googleapis.com/auth/blogger.readonly"
	//   ]
	// }

}

func cleanPathString(s string) string {
	return strings.Map(func(r rune) rune {
		if r >= 0x2d && r <= 0x7a || r == '~' {
			return r
		}
		return -1
	}, s)
}
