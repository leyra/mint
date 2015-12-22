package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"fmt"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/js"
	"github.com/spf13/cobra"
)

func main() {
	var JavascriptCmd = &cobra.Command{
		Use:   "js",
		Short: "Minify your javascript files",
		Long:  `Use this command to generate your minified javascript files.`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(os.Args) < 3 {
				fmt.Printf("You must specify the js file you'd like to minify.\n")
				os.Exit(2)
			}
			input := readFileIntoBuffer(args[0])
			js := minifyJavascript(input)

			writeFile("out.js", js)
		},
	}

	//recursive := false
	//output := "out.js"

	//JavascriptCmd.Flags().BoolVarP(&recursive, "recursive", "r", false, "recursive")
	//JavascriptCmd.Flags().StringVarP(&output, "output", "o", "", "output")

	var CssCmd = &cobra.Command{
		Use:   "css",
		Short: "Minify your css files",
		Long:  `Use this command to generate your minified css files.`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(os.Args) == 2 {
				fmt.Printf("You must specify the css file you'd like to minify.\n")
				os.Exit(2)
			}

			fb := FileBuffer{}

			if len(os.Args) > 2 {
				for i := 2; i < len(os.Args); i++ {
					input := readFileIntoBuffer(args[i])
					fb.Write(minifyCss(input).Bytes())
				}
			}

			//fmt.Println(fb.Contents().String())
			//writeFile("out.css", fb.Contents())
		},
	}

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(JavascriptCmd)
	rootCmd.AddCommand(CssCmd)
	rootCmd.Execute()
}

type FileBuffer struct {
	contents *bytes.Buffer
}

func (fb FileBuffer) Contents() (c *bytes.Buffer) {
	return fb.contents
}

func (fb *FileBuffer) Write(p []byte) (n int, err error) {
	fmt.Println("Here")
	fb.contents.Write(p)
	return len(p), nil
}

// isDirectory checks to see if the path given to be processed is a directory.
// Otherwise it will be assumed that the path is a file. Error handling to
// come.
func isDirectory(path string) bool {
	if info, err := os.Stat(path); err == nil && info.IsDir() {
		println("dir")
		return true
	}
	return false
}

// readFileIntoBuffer (for now) uses ioutil.ReadFile (into memory #sadface) and
// returns it as an instance of *bytes.Buffer.
func readFileIntoBuffer(path string) *bytes.Buffer {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(2)
	}

	return bytes.NewBuffer(f)
}

// writeFile currently writes a file with pre-defined r/w permissions. To be
// changed.
func writeFile(path string, content *bytes.Buffer) {
	ioutil.WriteFile(path, content.Bytes(), 0755)
}

// minifyJavascript reads the unminified js from the buffer, minifies it and
// creates a new buffer for it to be returned in.
func minifyJavascript(input *bytes.Buffer) *bytes.Buffer {
	output := bytes.NewBuffer([]byte{})

	m := minify.New()
	m.AddFunc("text/javascript", js.Minify)

	if err := m.Minify("text/javascript", output, input); err != nil {
		log.Fatal(err)
	}

	return output
}

// minifyCss reads the unminified css from the buffer, minifies it and
// creates a new buffer for it to be returned in.
func minifyCss(input *bytes.Buffer) *bytes.Buffer {
	output := bytes.NewBuffer([]byte{})

	m := minify.New()
	m.AddFunc("text/css", css.Minify)

	if err := m.Minify("text/css", output, input); err != nil {
		log.Fatal(err)
	}

	return output
}
