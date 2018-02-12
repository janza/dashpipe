package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codepipeline"
)

func serve(role string) {
	t, err := template.New("index").Funcs(template.FuncMap{
		"Deref": func(i *string) string { return *i },
		"Time": func(t *time.Time) string {
			return t.Format("2006-01-02 15:04")
		},
	}).Parse(`
<!DOCTYPE html>
<html>
<head>
<link href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm" crossorigin="anonymous">
<link href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css" rel="stylesheet">
</head>
<body>
	<div class="container">
	<table class="table">
		{{- range . }}
		<tr>
		    <td>
			<span style="white-space: nowrap">{{ .PipelineName }}</span>
			</td>
			{{if .StageStates -}}
			<td>
			<!-- <div class="btn-group"> -->
				{{- range .StageStates }}
					{{- range .ActionStates }}
					<td>
							{{if .LatestExecution }}

							{{with .LatestExecution}}
							<a class="btn btn-sm btn-block

								{{if eq "Succeeded" (Deref .Status) }}
									btn-outline-success
								{{else if eq "Failed" (Deref .Status) }}
									btn-danger
								{{else if eq "InProgress" (Deref .Status) }}
									btn-primary
								{{else}}
									btn-outline-secondary
								{{end}}

								" href="{{if .ExternalExecutionUrl}}{{.ExternalExecutionUrl}}{{else if .Token}}{{end}}" data-toggle="tooltip" title="{{.Status}}" data-content="
									{{if .LastUpdatedBy}}{{ .LastUpdatedBy }}:{{end}}
									{{if .Summary}}{{.Summary}}{{end}}
								">
							{{end}}

							{{if .CurrentRevision}}
								{{Time .CurrentRevision.Created}}
							{{else}}
								{{ .ActionName }}
							{{end}}

							{{if eq "Succeeded" (Deref .LatestExecution.Status) }}
								<i class="fa fa-check"></i>
							{{else if eq "Failed" (Deref .LatestExecution.Status) }}
								<i class="fa fa-times"></i>
							{{else if eq "InProgress" (Deref .LatestExecution.Status) }}
								<i class="fa fa-refresh fa-spin"></i>
							{{end}}

								</a>
							{{else}}
								<a class="btn btn-block btn-sm btn-outline-secondary" href="{{if not .EntityUrl}}{{.EntityUrl}}{{else}}#{{end}}" data-toggle="tooltip"  title="No executions yet">
								{{ .ActionName }}
								</a>
							{{end}}
						</td>
					{{end}}
				{{end}}
				<!--</div> -->
			</td>
			{{else}}
			No stages.
			{{end}}
		</tr>
		{{- end}}
	</table>
	</div>
	<script src="https://code.jquery.com/jquery-3.2.1.slim.min.js" integrity="sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN" crossorigin="anonymous"></script>
	<script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.12.9/umd/popper.min.js" integrity="sha384-ApNbgh9B+Y1QKtv3Rn7W3mgPxhU9K/ScQsAP7hUibX39j7fakFPskvXusvfa0b4Q" crossorigin="anonymous"></script>
	<script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/js/bootstrap.min.js" integrity="sha384-JZR6Spejh4U02d8jOt6vLEHfe/JQGiRRSQQxSfFWpi1MquVdAyjUar5+76PVCmYl" crossorigin="anonymous"></script>
	<script>
	$(function () {
		$('[data-toggle="tooltip"]').popover({trigger: 'hover', placement: 'bottom'})
	})
	</script>
</body>
</html>
	`)

	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
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
		}
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

	serve(role)
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
