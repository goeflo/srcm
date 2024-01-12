/*

copy from
https://gist.github.com/logrusorgru/abd846adb521a6fb39c7405f32fec0cf

*/

package templates

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Tmpl struct {
	*template.Template // root template
}

// NewTmpl creates new Tmpl.
func NewTmpl() (tmpl *Tmpl) {
	tmpl = new(Tmpl)
	tmpl.Template = template.New("") // unnamed root template
	return
}

func (t *Tmpl) Load(dir, ext string) (err error) {

	// get absolute path
	if dir, err = filepath.Abs(dir); err != nil {
		return fmt.Errorf("getting absolute path: %w", err)
	}

	var root = t.Template

	var walkFunc = func(path string, info os.FileInfo, err error) (_ error) {

		// handle walking error if any
		if err != nil {
			return err
		}

		// skip all except regular files
		// TODO (kostyarin): follow symlinks (?)
		if !info.Mode().IsRegular() {
			return
		}

		// filter by extension
		if filepath.Ext(path) != ext {
			return
		}

		// get relative path
		var rel string
		if rel, err = filepath.Rel(dir, path); err != nil {
			return err
		}

		// name of a template is its relative path
		// without extension
		rel = strings.TrimSuffix(rel, ext)
		rel = strings.Join(strings.Split(rel, string(os.PathSeparator)), "/")
		log.Printf("template path: %v\n", rel)
		// load or reload
		var (
			nt = root.New(rel)
			b  []byte
		)

		if b, err = os.ReadFile(path); err != nil {
			return err
		}

		_, err = nt.Parse(string(b))
		return err
	}

	if err = filepath.Walk(dir, walkFunc); err != nil {
		return
	}

	t.Template = root // set or replace (does it needed?)

	return nil
}
