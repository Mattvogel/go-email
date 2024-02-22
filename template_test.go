package email

import (
	"html/template"
	"testing"
)

type testArgs struct {
	Name string
}

func TestRender(t *testing.T) {

	args := testArgs{Name: "Jujube"}

	emailTpl, err := template.New("email").Parse("Hello, {{.Name}}")
	if err != nil {
		t.Errorf("Error parsing template: %s", err)
	}

	e := Email{
		To:             "test@example.com",
		From:           "testSender@example.com",
		InputArguments: args,
		Template:       emailTpl,
	}

	err = e.Render()
	if err != nil {
		t.Errorf("Error rendering template: %s", err)
	}
	if e.RenderedTemplate != "Hello, Jujube" {
		t.Errorf("Rendered template is incorrect: %s", e.RenderedTemplate)
	}
}
