package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	box "github.com/rsmaxwell/job-to-xml/internal/Box"
	"github.com/rsmaxwell/job-to-xml/internal/cmdline"
	"github.com/rsmaxwell/job-to-xml/internal/stream"
)

func main() {
	err := process()
	if err != nil {
		fmt.Println(err)
	}
}

func process() error {

	args, err := cmdline.GetArguments()
	if err != nil {
		return err
	}

	in, err := stream.NewInputStream(args.InputFilename)
	if err != nil {
		return err
	}
	defer in.Close()

	xmlTemplate, err := xmlTemplate()
	if err != nil {
		return err
	}

	escapedDsl, err := escapedJobDsl(in.Reader)
	if err != nil {
		return err
	}

	out, err := stream.NewOutputStream(args.OutputFilename)
	if err != nil {
		return err
	}
	defer out.Close()

	result := strings.Replace(xmlTemplate, "{{ ESCAPED_JOB_DSL }}", escapedDsl, -1)
	_, err = io.WriteString(out.Writer, result)
	if err != nil {
		return err
	}

	return nil
}

func xmlTemplate() (string, error) {
	resourceName := "/config.xml"
	byteArray, ok := box.Get(resourceName)
	if !ok {
		return "", fmt.Errorf("resource not found: %s", resourceName)
	}

	return string(byteArray), nil
}

func escapedJobDsl(in io.Reader) (string, error) {
	buf, err := ioutil.ReadAll(in)
	if err != nil {
		return "", err
	}

	var b bytes.Buffer
	xml.EscapeText(&b, buf)
	return b.String(), nil
}
