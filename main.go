package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codepipeline"
)

func serve(role string, t *template.Template) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pipelines, err := getPipelines(role)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		err = t.Execute(w, pipelines)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		return
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Print(http.ListenAndServe(":"+port, nil))
}

func main() {
	var role string

	flag.StringVar(&role, "r", "", "Role name to assume.")
	flag.Parse()

	tmpl, err := ioutil.ReadFile("./template.html")
	if err != nil {
		log.Fatal(err)
	}

	t, err := template.New("index").Funcs(template.FuncMap{
		"Deref": func(i *string) string { return *i },
		"Time": func(t *time.Time) string {
			return t.Format("2006-01-02 15:04")
		},
	}).Parse(string(tmpl))

	if err != nil {
		panic(err)
	}

	serve(role, t)
}

func getPipelines(role string) ([]*codepipeline.GetPipelineStateOutput, error) {
	sess := session.Must(session.NewSession())

	creds := stscreds.NewCredentials(sess, role)

	config := aws.NewConfig().WithCredentials(creds)

	pipe := codepipeline.New(sess, config)
	pipelinesOutput, err := pipe.ListPipelines(&codepipeline.ListPipelinesInput{})
	if err != nil {
		return nil, fmt.Errorf("failed to list pipelines, %v", err)
	}

	pipelinesList := make([]*codepipeline.GetPipelineStateOutput, 0)

	for _, pipeline := range pipelinesOutput.Pipelines {
		fmt.Printf("pipeline %s %s\n", *pipeline.Name, *pipeline.Updated)
		pipelineState, err := pipe.GetPipelineState(&codepipeline.GetPipelineStateInput{
			Name: pipeline.Name,
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to get pipeline %s, %v\n", *pipeline.Name, err)
			continue
		}
		pipelinesList = append(pipelinesList, pipelineState)
	}
	return pipelinesList, nil
}
