package cli

// https://samrapdev.com/capturing-sensitive-input-with-editor-in-golang-from-the-cli/

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// DefaultEditor is vim because we're adults ;)
const DefaultEditor = "vim"

// PreferredEditorResolver is a function that returns an editor that the user
// prefers to use, such as the configured `$EDITOR` environment variable.
type PreferredEditorResolver func() string

// GetPreferredEditor returns the user's editor as defined by the
// `$EDITOR` environment variable, or the `DefaultEditor` if it is not set.
func GetPreferredEditor() string {
	editor := os.Getenv("EDITOR")

	if editor == "" {
		return DefaultEditor
	}

	return editor
}

func resolveEditorArguments(executable string, filename string) []string {
	args := []string{filename}

	if strings.Contains(executable, "Visual Studio Code.app") || strings.Contains(executable, ".vscode-server") {
		args = append([]string{"--wait"}, args...)
	}

	// Other common editors

	return args
}

// TextEditor can open and modify files using an editor
type TextEditor struct {
	ResolveEditor PreferredEditorResolver
}

// NewTextEditor creates a TextEditor
func NewTextEditor() *TextEditor {
	return &TextEditor{
		ResolveEditor: GetPreferredEditor,
	}
}

// OpenFile opens filename in a text editor.
func (t *TextEditor) OpenFile(filename string) error {
	// Get the full executable path for the editor.
	executable, err := exec.LookPath(t.ResolveEditor())
	if err != nil {
		return err
	}

	cmd := exec.Command(executable, resolveEditorArguments(executable, filename)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// CaptureInput opens a temporary file in a text editor and returns
// the written bytes on success or an error on failure. It handles deletion
// of the temporary file behind the scenes.
func (t *TextEditor) CaptureInput() (string, error) {
	return t.EditText("")
}

// EditText edits the given text using the editor returned by resolveEditor
func (t *TextEditor) EditText(txt string) (string, error) {
	file, err := ioutil.TempFile(os.TempDir(), "*")
	if err != nil {
		return "", err
	}

	if len(txt) > 0 {
		_, err = file.Write([]byte(txt))
		if err != nil {
			return "", err
		}
	}

	filename := file.Name()

	// Defer removal of the temporary file in case any of the next steps fail.
	defer os.Remove(filename)

	if err = file.Close(); err != nil {
		return "", err
	}

	if err = t.OpenFile(filename); err != nil {
		return "", err
	}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
