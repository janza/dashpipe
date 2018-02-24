package main

import (
	"encoding/json"
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

func serve(
	stages []PipeStage,
	header *template.Template,
	footer *template.Template,
	table *template.Template,
) {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		err := header.Execute(w, "")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		for _, pipe := range stages {
			pipelines, err := pipe.getPipelinesOutput()
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			// err := pipe.renderPipelines(w)
			err = table.Execute(w, pipelines)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}

		}

		err = footer.Execute(w, "")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Starting listener on port: %s\n", port)

	fmt.Print(http.ListenAndServe(":"+port, nil))
}

type stageConfiguration struct {
	Region string
	Role   string
	Name   string
}

// PipeStage lists codepipelines in a specific region using a provided role
type PipeStage struct {
	pipelinesClient *codepipeline.CodePipeline
	Name            string
}

func (stage *PipeStage) setupClient(configuration stageConfiguration) {
	sess := session.Must(session.NewSession())
	creds := stscreds.NewCredentials(sess, configuration.Role)
	config := aws.NewConfig().WithCredentials(creds).WithRegion(configuration.Region)
	stage.Name = configuration.Name

	stage.pipelinesClient = codepipeline.New(sess, config)
}

func (stage *PipeStage) getPipelines() ([]*codepipeline.GetPipelineStateOutput, error) {
	client := stage.pipelinesClient
	pipelinesOutput, err := client.ListPipelines(&codepipeline.ListPipelinesInput{})
	if err != nil {
		return nil, fmt.Errorf("failed to list pipelines for %s, %v", stage.Name, err)
	}

	pipelinesList := make([]*codepipeline.GetPipelineStateOutput, 0)

	for _, pipeline := range pipelinesOutput.Pipelines {
		fmt.Printf("pipeline %s %s\n", *pipeline.Name, *pipeline.Updated)
		pipelineState, err := client.GetPipelineState(&codepipeline.GetPipelineStateInput{
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

type actionObject struct {
	Name       string
	Status     string
	Summary    string
	LastUpdate string
	Author     string
	URL        string
	Time       *time.Time
}

type pipelineObject struct {
	Actions []actionObject
	Name    string
}

type pipelinesOutput struct {
	Pipelines []pipelineObject
	Name      string
}

func derefString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func (stage *PipeStage) getPipelinesOutput() (pipelinesOutput, error) {
	pipelines, err := stage.getPipelines()
	if err != nil {
		return pipelinesOutput{}, err
	}
	var pipes []pipelineObject
	for _, pipe := range pipelines {
		var actions []actionObject
		for _, stage := range pipe.StageStates {
			for _, action := range stage.ActionStates {
				a := actionObject{}
				if action.LatestExecution != nil {
					a.Status = *action.LatestExecution.Status
					a.Summary = derefString(action.LatestExecution.Summary)
					a.Author = derefString(action.LatestExecution.LastUpdatedBy)
					a.URL = derefString(action.LatestExecution.ExternalExecutionUrl)
					a.Name = derefString(action.ActionName)
					if action.CurrentRevision != nil {
						a.Time = action.CurrentRevision.Created
					}
				}

				actions = append(actions, a)
			}
		}
		pipes = append(pipes, pipelineObject{
			Name:    *pipe.PipelineName,
			Actions: actions,
		})
	}
	return pipelinesOutput{
		Pipelines: pipes,
		Name:      stage.Name,
	}, nil
}

// SetupStage sets up a PipeStage with configuration struct
func SetupStage(configuration stageConfiguration) PipeStage {
	stage := PipeStage{}
	stage.setupClient(configuration)
	return stage
}

func parseConfig(configFile string) ([]PipeStage, error) {
	configJSON, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}
	var configs []stageConfiguration
	err = json.Unmarshal(configJSON, &configs)
	if err != nil {
		return nil, err
	}
	var pipeStages []PipeStage

	for _, config := range configs {
		pipeStages = append(pipeStages, SetupStage(config))
	}

	return pipeStages, nil
}

func getTemplate(file string) (*template.Template, error) {
	tmpl, err := Asset(file)
	if err != nil {
		log.Fatal(err)
	}

	t, err := template.New("index").Funcs(template.FuncMap{
		"Deref": func(i *string) string { return *i },
		"Time": func(t *time.Time) string {
			return t.Format("2006-01-02 15:04")
		},
	}).Parse(string(tmpl))
	return t, err
}

func main() {
	var role string

	flag.StringVar(&role, "r", "", "Role name to assume.")
	flag.Parse()

	stages, err := parseConfig("./config.json")
	if err != nil {
		fmt.Println("Failed to read config.json")
		panic(err)
	}

	header, err := getTemplate("template/header.html")
	if err != nil {
		fmt.Println("Failed to parse header.html")
		panic(err)
	}
	footer, err := getTemplate("template/footer.html")
	if err != nil {
		fmt.Println("Failed to parse footer.html")
		panic(err)
	}
	table, err := getTemplate("template/table.html")
	if err != nil {
		fmt.Println("Failed to parse table.html")
		panic(err)
	}

	serve(stages, header, footer, table)
}
