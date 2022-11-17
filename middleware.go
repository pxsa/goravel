package goravel

import "net/http"

func (c *Goravel) SessionLoad (next http.Handler) http.Handler {
	c.InfoLog.Println("SessionLoad called")
	return c.Session.LoadAndSave(next)
}