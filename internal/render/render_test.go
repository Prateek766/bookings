package render

import (
	"net/http"
	"testing"

	"github.com/Prateek766/bookings/internal/models"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData

	r, err := getSession()
	if err != nil {
		t.Fatalf("failed to get session: %v", err)
	}

	session.Put(r.Context(), "flash", "test flash")
	result := AddDefaultData(&td, r)

	if result.Flash != "test flash" {
		t.Errorf("expected flash message to be set, got %s", result.Flash)
	}

}

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "./../../templates"
	tc, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}

	r, err := getSession()
	if err != nil {
		t.Fatalf("failed to get session: %v", err)
	}
	app.TemplateCache = tc
	var ww myWriter
	err = RenderTemplate(&ww, r, "home.page.tmpl", &models.TemplateData{})
	if err != nil {
		t.Errorf("error rendering template: %v", err)
	}

	err = RenderTemplate(&ww, r, "non-existent.page.tmpl", &models.TemplateData{})
	if err == nil {
		t.Errorf("expected error rendering non-existent template, got nil")
	}

}

func TestNewTemplates(t *testing.T) {
	NewTemplates(app)
	if app == nil {
		t.Error("app is nil")
	}
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "./../../templates"
	_, err := CreateTemplateCache()
	if err != nil {
		t.Errorf("error creating template cache: %v", err)
	}
}
func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		return nil, err
	}

	ctx := r.Context()
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))
	r = r.WithContext(ctx)

	return r, nil
}
