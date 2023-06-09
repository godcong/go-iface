package overloader

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
	"unicode"

	"github.com/godcong/go-iface/parse"
)

type OverLoader struct {
	fileSet *token.FileSet
	//tmpl      *template.Template
	faces  map[string]*parse.Struct
	pkg    string
	suffix string
}

// New is a constructor method for creating a new OverLoader with default
// templates loaded.
func New() *OverLoader {
	return &OverLoader{
		//tmpl:      addEmbeddedTemplates(template.New("generator")),
		fileSet: token.NewFileSet(),
		faces:   make(map[string]*parse.Struct),
		suffix:  "overload",
	}
}

// GenerateFromPath is responsible for orchestrating the Code generation.  It results in a byte array
// that can be written to any file desired.  It has already had goimports run on the code before being returned.
func (l *OverLoader) GenerateFromPath(path string) (map[string][]byte, error) {
	f, err := l.parsePath(path)
	if err != nil {
		return nil, fmt.Errorf("generate: error parsing input path '%s': %s", path, err)
	}
	return l.Generate(f)
}

// parsePath simply calls the go/parser ParseFile function with an empty token.FileSet
func (l *OverLoader) parsePath(fileName string) (map[string]*ast.Package, error) {
	// Parse the file given in arguments
	return parser.ParseDir(l.fileSet, fileName, nil, parser.ParseComments)
}

func (l *OverLoader) Generate(f map[string]*ast.Package) (map[string][]byte, error) {
	for name, pkg := range f {
		if strings.HasSuffix(name, "_test") {
			continue
		}
		if l.pkg == "" {
			l.pkg = pkg.Name
		}
		for _, file := range pkg.Files {
			ast.Walk(l, file)
		}
	}
	ret := make(map[string][]byte)
	//for _, m := range l.faces {
	//	vBuff := bytes.NewBuffer([]byte{})
	//	err := l.tmpl.ExecuteTemplate(vBuff, "header", map[string]interface{}{
	//		"package":   l.targetPkg,
	//		"version":   "",
	//		"revision":  "",
	//		"buildDate": "",
	//		"builtBy":   "",
	//	})
	//	if err != nil {
	//		return nil, fmt.Errorf("failed write template:%v", err)
	//	}
	//	var methods []string
	//	for _, param := range m.Methods {
	//		methods = append(methods, param.Val())
	//	}
	//	err = l.tmpl.ExecuteTemplate(vBuff, "iface", map[string]interface{}{
	//		"name":    camelCase(m.Name),
	//		"methods": methods,
	//	})
	//	if err != nil {
	//		return nil, fmt.Errorf("failed write template:%v", err)
	//	}
	//	filename := strings.Join([]string{snakeCase(m.Name), l.suffix}, "_")
	//
	//	formatted, err := format.Source(vBuff.Bytes())
	//	if err != nil {
	//		return nil, err
	//	}
	//	ret[filename] = formatted
	//
	//}
	return ret, nil
}

func (l *OverLoader) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.FuncDecl:
		var s string
		if n.Recv != nil {
			for _, arg := range n.Recv.List {
				s = fmt.Sprintf("%s", arg.Type)
			}
		}
		//skip empty receiver
		if s == "" {
			return l
		}
		inter := &parse.Struct{Name: s}
		if i, ok := l.faces[s]; ok {
			inter = i
		}

		inter.Parse(n)
		l.faces[s] = inter
	}
	return l
}

func camelCase(s string) string {
	if len(s) == 0 {
		return s
	}
	source := []rune(s)
	size := len(source)
	ret := make([]rune, 0)

	idx := 0
	for i := 0; i < size; i++ {
		if !unicode.IsLetter(source[i]) {
			continue
		}
		idx = i
		break
	}
	if idx >= size {
		return s
	}
	//start at first idx
	ret = append(ret, unicode.ToUpper(source[idx]))

	toUpper := false
	for idx++; idx < size; idx++ {
		if source[idx] == '_' {
			toUpper = true
			continue
		}
		if toUpper {
			if !unicode.IsLetter(source[idx]) {
				ret = append(ret, '_')
			}
			source[idx] = unicode.ToUpper(source[idx])
			toUpper = false
		}
		ret = append(ret, source[idx])

	}
	return string(ret)
}

func snakeCase(s string) string {
	if len(s) == 0 {
		return s
	}
	source := []rune(s)
	ret := make([]rune, 0)
	size := len(source)

	idx := 0
	for i := 0; i < size; i++ {
		if !unicode.IsLetter(source[i]) {
			continue
		}
		idx = i
		break
	}
	if idx >= size {
		return s
	}

	preUpper := false
	preLine := false

	for i, r := range source[idx:] {
		if unicode.IsUpper(r) {
			if i == 0 || preUpper || preLine {
				ret = append(ret, unicode.ToLower(r))
			} else {
				ret = append(ret, '_', unicode.ToLower(r))
			}
			preUpper = true
			preLine = false
			continue
		} else if r == '_' {
			if i != 0 {
				ret = append(ret, r)
			}
			preLine = true
			continue
		}
		ret = append(ret, r)
		preUpper = false
		preLine = false
	}
	return string(ret)
}
