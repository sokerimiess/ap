package main

import (
	"awesomeProject4/pkg/forms" // New import
	"awesomeProject4/pkg/models"
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Because Pat matches the "/" path exactly, we can now remove the manual check
	// Because Pat matches the "/" path exactly, we can now remove the manual check
	// of r.URL.Path != "/" from this handler.
	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.render(w, r, "home.page.tmpl", &templateData{
		Snippets: s,
	})
}
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	// Pat doesn't strip the colon from the named capture key, so we need to
	// get the value of ":id" from the query string instead of "id".
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	s, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	app.render(w, r, "show.page.tmpl", &templateData{
		Snippet: s,
	})
}

// Add a new createSnippetForm handler, which for now returns a placeholder response.
func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.tmpl", &templateData{
		// Pass a new empty forms.Form object to the template.
		Form: forms.New(nil),
	})
}
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	// Create a new forms.Form struct containing the POSTed data from the
	// form, then use the validation methods to check the content.
	form := forms.New(r.PostForm)
	form.Required("title", "content", "expires")
	form.MaxLength("title", 100)
	form.PermittedValues("expires", "365", "7", "1")
	// If the form isn't valid, redisplay the template passing in the
	// form.Form object as the data.
	if !form.Valid() {
		app.render(w, r, "create.page.tmpl", &templateData{Form: form})
		return
	}
	// Because the form data (with type url.Values) has been anonymously embedded
	// in the form.Form struct, we can use the Get() method to retrieve
	// the validated value for a particular form field.
	id, err := app.snippets.Insert(form.Get("title"), form.Get("content"), form.Get("expires"))
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}
func (app *application) students(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/students" {
		app.notFound(w)
		return
	}
	s, err := app.snippets.CaterSTUDE()
	if err != nil {
		app.serverError(w, err)
		return
	}
	// Use the new render helper.
	app.render(w, r, "stud.page.tmpl", &templateData{
		Snippets: s,
	})
}
func (app *application) applicants(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/applicants" {
		app.notFound(w)
		return
	}
	s, err := app.snippets.CaterApp()
	if err != nil {
		app.serverError(w, err)
		return
	}
	// Use the new render helper.
	app.render(w, r, "applicants.page.tmpl", &templateData{
		Snippets: s,
	})
}
func (app *application) staff(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/staff" {
		app.notFound(w)
		return
	}
	s, err := app.snippets.CaterSTAFF()
	if err != nil {
		app.serverError(w, err)
		return
	}
	// Use the new render helper.
	app.render(w, r, "staff.page.tmpl", &templateData{
		Snippets: s,
	})
}

func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {
	if td == nil {
		td = &templateData{}
	}
	td.CurrentYear = time.Now().Year()
	return td
}
func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("The template %s does not exist", name))
		return
	}
	buf := new(bytes.Buffer)
	// Execute the template set, passing the dynamic data with the current
	// year injected.
	err := ts.Execute(buf, app.addDefaultData(td, r))
	if err != nil {
		app.serverError(w, err)
		return
	}
	buf.WriteTo(w)
}

func (app *application) researches(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/researches" {
		app.notFound(w)
		return
	}
	s, err := app.snippets.CaterRes()
	if err != nil {
		app.serverError(w, err)
		return
	}
	// Use the new render helper.
	app.render(w, r, "res.page.tmpl", &templateData{
		Snippets: s,
	})
}
