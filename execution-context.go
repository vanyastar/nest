package nest

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

type IDto interface {
	Validate() error
}

type nextFunc func()

var ctxPool sync.Pool

type Response struct {
	http.Flusher
	http.ResponseWriter
}

func (c *Response) reset() {
	c.Flusher = nil
	c.ResponseWriter = nil
}

type Request struct {
	*http.Request
}

func (r *Request) reset() {
	r.Request = nil
}

// Ctx represents the context with response and request methods
type Ctx struct {
	request  *Request
	response *Response
	Next     nextFunc
}

// Sends local file contents from the given path as response body.
func (c *Ctx) SendFile(f string) error {
	file, err := os.Open(f)
	if err != nil {
		return err
	}
	defer file.Close()
	info, err := file.Stat()
	if err != nil {
		return err
	}
	http.ServeContent(c.Res(), c.Req().Request, info.Name(), info.ModTime(), file)
	return nil
}

func (c *Ctx) Cookie(name string) (*http.Cookie, error) {
	return c.Req().Cookie(name)
}

// Set cookie
func (c *Ctx) SetCookie(name, value, path, domain string, maxAge int, secure, httpOnly bool) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     path,
		Domain:   domain,
		MaxAge:   maxAge,
		Secure:   secure,
		HttpOnly: httpOnly,
		Expires:  time.Now().Add(time.Duration(maxAge) * time.Second),
	}
	http.SetCookie(c.Res(), cookie)
}

// Respond as string result
func (c *Ctx) SendString(s string) error {
	c.Res().Header().Set("Content-Type", "text/html")
	_, err := c.Res().Write([]byte(s))
	return err
}

// Respond encodes the value as JSON or XML based on the Content-Type header
func (c *Ctx) Send(v any) error {
	switch c.Req().Header.Get("Content-Type") {
	case "application/xml":
		c.Res().Header().Set("Content-Type", "application/xml")
		return xml.NewEncoder(c.Res()).Encode(v)
	default:
		c.Res().Header().Set("Content-Type", "application/json")
		return json.NewEncoder(c.Res()).Encode(v)
	}
}

// Regular body parser based on the Content-Type header
func (c *Ctx) BodyParser(i any) error {
	switch c.Req().Header.Get("Content-Type") {
	case "application/xml":
		return xml.NewDecoder(c.Req().Body).Decode(i)
	default:
		return json.NewDecoder(c.Req().Body).Decode(i)
	}
}

// Session Manager
// TODO: Make custom storages for session, on this moment only RAM.
func (c *Ctx) Session() *Session {
	// Load retrieves the session from the request cookie
	cookie, err := c.Req().Cookie(sessionCookieName)
	if err != nil {
		return newSession()
	}

	// Load the session from the SessionStore
	if session, ok := SessionStorage.Load(cookie.Value); ok {
		return session.(*Session)
	}
	return newSession()
}

// Regular body parser + auto validation from struct method, based on the Content-Type header
func (c *Ctx) DtoParser(i IDto) error {
	if err := c.BodyParser(i); err != nil {
		return err
	}
	return i.Validate()
}

// Error handles errors based on content type and message type
// HTTP 200 don't send message from this function
func (c *Ctx) Error(statusCode int, message any) error {
	if statusCode == http.StatusOK {
		return nil
	}

	switch t := message.(type) {
	case nil:
		c.Res().WriteHeader(statusCode)
		return nil
	case string:
		c.Res().WriteHeader(statusCode)
		_, err := c.Res().Write([]byte(t))
		if err != nil {
			fmt.Println(err)
		}
		return nil
	case []byte:
		c.Res().WriteHeader(statusCode)
		_, err := c.Res().Write(t)
		if err != nil {
			fmt.Println(err)
		}
		return nil
	}

	switch c.Req().Header.Get("Content-Type") {
	case "application/xml":
		c.Res().Header().Set("Content-Type", "application/xml")
		c.Res().WriteHeader(statusCode)
		return xml.NewEncoder(c.Res()).Encode(message)
	default:
		c.Res().Header().Set("Content-Type", "application/json")
		c.Res().WriteHeader(statusCode)
		return json.NewEncoder(c.Res()).Encode(message)
	}
}

// Flush sends any buffered data to the client.
func (c *Ctx) Flush() error {
	if c.Res().Flusher != nil {
		c.Res().Flusher.Flush()
		return nil
	}
	return errors.New("the flusher is not supported")
}

// Get access to *http.Request.
func (c *Ctx) Req() *Request {
	return c.request
}

// Get access to http.ResponseWriter.
func (c *Ctx) Res() *Response {
	return c.response
}

func (c *Ctx) setHttpContext(w http.ResponseWriter, r *http.Request) *Ctx {
	c.Req().Request = r
	c.Res().ResponseWriter = w
	return c
}

func (c *Ctx) reset() {
	c.Res().reset()
	c.Req().reset()
	c.Next = nil
	ctxPool.Put(c)
}

func createCtx(w http.ResponseWriter, r *http.Request) *Ctx {
	return &Ctx{
		request:  &Request{Request: r},
		response: &Response{ResponseWriter: w},
		Next:     func() {},
	}
}

func newCtx(w http.ResponseWriter, r *http.Request) *Ctx {
	usedCtx := ctxPool.Get()
	if usedCtx != nil {
		return usedCtx.(*Ctx).setHttpContext(w, r)
	}
	return createCtx(w, r)
}
