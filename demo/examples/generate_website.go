package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"text/template"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"github.com/joshcarp/sysl-printing/docs/website/byexample/syslchroma"
	"github.com/russross/blackfriday"
	"gopkg.in/yaml.v2"
)

// Code adapted from go by example; https://gobyexample.com/
// siteDir is the target directory into which the HTML gets generated. Its
// default is set here but can be changed by an argument passed into the
// program.
const syslPlaygroundURL = "http://anz-bank.github.io/sysl-playground/"
const syslRoot = "../../"
const siteDir = syslRoot + "docs/website/content/docs/byexample/"
const assetDir = syslRoot + "docs/website/static/assets/byexample/"
const templates = syslRoot + "docs/website/byexample/templates"
const cacheDir = "./.tmp/gobyexample-cache"
const orderingfile = "ordering.yaml"

var imageFiles = []string{".svg"}

func main() {
	ensureDir(siteDir)
	examples := parseExamples()
	renderExamples(examples)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// findSyslCommand is used to match only the command and not the comments for sysl-playground
func findSyslCommand(input string) string {
	var re = regexp.MustCompile(`(?m)(?:\n)(sysl.*)`)
	ans := re.FindString(input)
	ans = strings.Replace(ans, "\n", "", 1)
	return ans

}

func ensureDir(dir string) {
	err := os.MkdirAll(dir, 0755)
	check(err)
}

func copyFile(src, dst string) {
	dat, err := ioutil.ReadFile(src)
	check(err)
	err = ioutil.WriteFile(dst, dat, 0644)
	check(err)
}

func sha1Sum(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	b := h.Sum(nil)
	return fmt.Sprintf("%x", b)
}

func mustReadFile(path string) string {
	bytes, err := ioutil.ReadFile(path)
	check(err)
	return string(bytes)
}

func cacheChroma(lex string, src string) string {
	ensureDir(cacheDir)
	cachePath := cacheDir + "/" + sha1Sum(src)
	cacheBytes, cacheErr := ioutil.ReadFile(cachePath)
	if cacheErr == nil {
		return string(cacheBytes)
	}
	return chromaFormat(src)
}

func chromaFormat(code string) string {
	lexer := syslchroma.Sysl
	if lexer == nil {
		lexer = lexers.Fallback
	}
	lexer = chroma.Coalesce(lexer)

	style := styles.Get("swapoff")
	if style == nil {
		style = styles.Fallback
	}
	formatter := html.New(html.WithClasses(true))
	iterator, err := lexer.Tokenise(nil, code)
	check(err)
	buf := new(bytes.Buffer)
	err = formatter.Format(buf, style, iterator)
	check(err)
	return buf.String()

}
func markdown(src string) string {
	return string(blackfriday.Run([]byte(src)))
}

func readLines(path string) []string {
	src := mustReadFile(path)
	return strings.Split(src, "\n")
}

func mustGlob(glob string) []string {
	paths, err := filepath.Glob(glob)
	check(err)
	return paths
}

func whichLexer(path string) string {
	if strings.HasSuffix(path, ".sysl") {
		return "sysl"
	} else if strings.HasSuffix(path, ".sh") {
		return "console"
	}
	return ""
}

func debug(msg string) {
	if os.Getenv("DEBUG") == "1" {
		fmt.Fprintln(os.Stderr, msg)
	}
}

var docsPat = regexp.MustCompile("^\\s*(\\/\\/|#)\\s")
var dashPat = regexp.MustCompile("\\-+")

// Seg is a segment of an example
type Seg struct {
	Docs, DocsRendered              string
	Code, CodeRendered              string
	CodeEmpty, CodeLeading, CodeRun bool
	Image                           string
}

// Example is info extracted from an example file
type Example struct {
	ID, Name, Topic     string
	Weight              int
	CodeWithoutComments string
	Cmd                 string
	PlaygroundURL       string
	Segs                [][]*Seg
	PrevExample         *Example
	NextExample         *Example
}

func parseSegs(sourcePath string) ([]*Seg, string) {
	var lines []string
	// Convert tabs to spaces for uniform rendering.
	for _, line := range readLines(sourcePath) {
		lines = append(lines, strings.Replace(line, "\t", "    ", -1))
	}
	filecontent := strings.Join(lines, "\n")
	segs := []*Seg{}
	lastSeen := ""
	for _, line := range lines {
		if line == "" {
			lastSeen = ""
			continue
		}
		matchDocs := docsPat.MatchString(line)
		matchCode := !matchDocs
		newDocs := (lastSeen == "") || ((lastSeen != "docs") && (segs[len(segs)-1].Docs != ""))
		newCode := (lastSeen == "") || ((lastSeen != "code") && (segs[len(segs)-1].Code != ""))
		if newDocs || newCode {
			debug("NEWSEG")
		}
		if matchDocs {
			trimmed := docsPat.ReplaceAllString(line, "")
			if newDocs {
				newSeg := Seg{Docs: trimmed, Code: ""}
				segs = append(segs, &newSeg)
			} else {
				segs[len(segs)-1].Docs = segs[len(segs)-1].Docs + trimmed
			}
			debug("DOCS: " + line)
			lastSeen = "docs"
		} else if matchCode {
			if newCode {
				newSeg := Seg{Docs: "", Code: line}
				segs = append(segs, &newSeg)
			} else {
				segs[len(segs)-1].Code = segs[len(segs)-1].Code + "\n" + line
			}
			debug("CODE: " + line)
			lastSeen = "code"
		}
	}
	for i, seg := range segs {
		seg.CodeEmpty = (seg.Code == "")
		seg.CodeLeading = (i < (len(segs) - 1))
		seg.CodeRun = i == 1
	}
	return segs, filecontent
}

func parseAndRenderSegs(sourcePath string) ([]*Seg, string, string) {
	lexer := whichLexer(sourcePath)
	segs, filecontent := parseSegs(sourcePath)
	Code := ""
	for _, seg := range segs {
		if seg.Docs != "" {
			seg.DocsRendered = markdown(seg.Docs)
		}
		if seg.Code != "" {
			seg.CodeRendered = cacheChroma(lexer, seg.Code)

			// adding the content to the js code for copying to the clipboard
			if strings.HasSuffix(sourcePath, ".sysl") {
				Code += strings.Trim(seg.Code, "\n") + "\n"
			}
		}
	}
	return segs, filecontent, Code
}

// unmarshalYaml unmarshals a yaml file of form
// key1:
//      - value 1
//      - value 2
func unmarshalYaml(filename string) ordering {
	source, err := ioutil.ReadFile(filename)
	check(err)
	that := yaml.MapSlice{}
	err = yaml.Unmarshal(source, &that)
	var items ordering
	for _, item := range that {
		new := make(map[string][]string)
		key := item.Key.(string)
		for _, val := range item.Value.([]interface{}) {
			valString := val.(string)
			new[key] = append(new[key], valString)
		}
		items = append(items, new)
	}
	return items
}

type ordering []map[string][]string

func parseExamples() []*Example {
	var examples []*Example
	ordering := unmarshalYaml(orderingfile)
	weight := 0
	err := os.RemoveAll(assetDir + "images")
	check(err)
	err = os.MkdirAll(assetDir+"images", 0755)
	check(err)
	for _, tutorial := range ordering {
		for topic, val := range tutorial {
			fmt.Println(topic)
			for _, exampleName := range val {
				fmt.Println("\t" + exampleName)
				weight++
				example := Example{Name: exampleName}
				exampleID := strings.ToLower(exampleName)
				exampleID = strings.Replace(exampleID, " ", "-", -1)
				exampleID = strings.Replace(exampleID, "/", "-", -1)
				exampleID = strings.Replace(exampleID, "'", "", -1)
				exampleID = dashPat.ReplaceAllString(exampleID, "-")
				example.ID = exampleID
				example.Weight = weight + 1
				example.Topic = topic
				example.Segs = make([][]*Seg, 0)
				sourcePaths := mustGlob(exampleID + "/*")
				for _, sourcePath := range sourcePaths {
					if ok, ext := isImageFile(sourcePath); ok {
						destination := assetDir + "images/" + exampleID + strconv.Itoa(weight) + ext

						copyFile(sourcePath, destination)

						// This is the path that gets rendered in the markdown file
						imagesRelativeToSite := "/assets/byexample/images/"

						var Segment = make([]*Seg, 1)
						imageName := imagesRelativeToSite + exampleID + strconv.Itoa(weight) + ext
						Segment[0] = &Seg{Image: imageName}
						example.Segs = append(example.Segs, Segment)

					} else if whichLexer(sourcePath) != "" {
						sourceSegs, filecontents, codeWithoutComments := parseAndRenderSegs(sourcePath)

						if filecontents != "" {
							switch whichLexer(sourcePath) {
							case "sysl":
								example.CodeWithoutComments = codeWithoutComments
							case "console":
								example.Cmd = findSyslCommand(filecontents)
							}

						}
						example.Segs = append(example.Segs, sourceSegs)
					}

				}
				example.PlaygroundURL = syslplaygroundLink(example.CodeWithoutComments, example.Cmd)
				examples = append(examples, &example)
			}
		}
	}

	for i, example := range examples {
		if i > 0 {
			example.PrevExample = examples[i-1]
		}
		if i < (len(examples) - 1) {
			example.NextExample = examples[i+1]
		}
	}

	return examples
}

func renderExamples(examples []*Example) {
	exampleTmpl := template.New("example")
	_, err := exampleTmpl.Parse(mustReadFile(templates + "/example.tmpl"))
	check(err)
	for _, example := range examples {
		exampleF, err := os.Create(siteDir + example.ID + ".md")
		check(err)
		exampleTmpl.Execute(exampleF, example)
	}
}
func isImageFile(filename string) (bool, string) {
	for _, extension := range imageFiles {
		if strings.HasSuffix(filename, extension) {
			return true, extension
		}
	}
	return false, ""
}

func syslplaygroundLink(code, cmd string) string {
	code = encode(code)
	cmd = encode(cmd)
	return fmt.Sprintf(syslPlaygroundURL+"?input=%s&cmd=%s", code, cmd)
}

func encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}
