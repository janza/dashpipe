package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/codebuild"
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

	http.HandleFunc("/details/", func(w http.ResponseWriter, r *http.Request) {
		accountIdx, err := strToInt(r.URL.Query().Get("account"))
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		provider := r.URL.Query().Get("provider")
		actionID := r.URL.Query().Get("id")
		res, err := stages[accountIdx].getActionDetails(provider, actionID)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		_, err = w.Write([]byte(res))
		w.WriteHeader(http.StatusOK)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Starting listener on port: %s\n", port)

	fmt.Print(http.ListenAndServe(":"+port, nil))
}

func strToInt(str string) (int, error) {
	nonFractionalPart := strings.Split(str, ".")
	return strconv.Atoi(nonFractionalPart[0])
}

type stageConfiguration struct {
	Index  int
	Region string
	Role   string
	Name   string
}

// PipeStage lists codepipelines in a specific region using a provided role
type PipeStage struct {
	pipelinesClient      *codepipeline.CodePipeline
	cloudformationClient *cloudformation.CloudFormation
	codebuildClient      *codebuild.CodeBuild
	cloudwatchClient     *cloudwatchlogs.CloudWatchLogs
	Name                 string
	ID                   int
}

func (stage *PipeStage) setupClient(id int, configuration stageConfiguration) {
	stage.ID = id
	sess := session.Must(session.NewSession())
	creds := stscreds.NewCredentials(sess, configuration.Role)
	config := aws.NewConfig().WithCredentials(creds).WithRegion(configuration.Region)
	stage.Name = configuration.Name

	stage.pipelinesClient = codepipeline.New(sess, config)
	stage.cloudformationClient = cloudformation.New(sess, config)
	stage.codebuildClient = codebuild.New(sess, config)
	stage.cloudwatchClient = cloudwatchlogs.New(sess, config)
}

type pipelineExtendedInfo struct {
	state *codepipeline.GetPipelineStateOutput
	info  *codepipeline.GetPipelineOutput
}

func (stage *PipeStage) getActionDetails(provider string, actionID string) (string, error) {
	if provider == "CloudFormation" && strings.Contains(actionID, "changeset") {
		changeSetName := strings.Split(actionID, "=")[1]
		details, err := stage.cloudformationClient.DescribeChangeSet(
			&cloudformation.DescribeChangeSetInput{
				ChangeSetName: &changeSetName,
			},
		)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%v", details), nil
	}
	if provider == "CloudFormation" && strings.Contains(actionID, "stack") {
		stackName := strings.Split(actionID, "=")[1]
		details, err := stage.cloudformationClient.DescribeStackEvents(
			&cloudformation.DescribeStackEventsInput{
				StackName: &stackName,
			},
		)
		if err != nil {
			return "", err
		}
		// return fmt.Sprintf("%v", details), nil
		s := ""
		for _, event := range details.StackEvents {
			s = s + "\n" + fmt.Sprintf(
				"%s\t%s\t%s\t%s",
				derefString(event.ResourceStatus),
				derefString(event.ResourceType),
				derefString(event.LogicalResourceId),
				derefString(event.ResourceStatusReason),
			)
		}
		return s, nil
	}

	if provider == "CodeBuild" {
		details, err := stage.codebuildClient.BatchGetBuilds(
			&codebuild.BatchGetBuildsInput{
				Ids: []*string{
					aws.String(actionID),
				},
			},
		)
		if err != nil {
			return "", nil
		}
		output, err := stage.cloudwatchClient.GetLogEvents(
			&cloudwatchlogs.GetLogEventsInput{
				LogGroupName:  details.Builds[0].Logs.GroupName,
				LogStreamName: details.Builds[0].Logs.StreamName,
			},
		)
		if err != nil {
			return "", nil
		}
		s := ""
		for _, event := range output.Events {
			s = s + "\n" + derefString(event.Message)
		}
		return s, nil
	}
	return "No details found for action", nil
}

func (stage *PipeStage) getPipelines() ([]*pipelineExtendedInfo, error) {
	client := stage.pipelinesClient
	pipelinesOutput, err := client.ListPipelines(&codepipeline.ListPipelinesInput{})
	if err != nil {
		return nil, fmt.Errorf("failed to list pipelines for %s, %v", stage.Name, err)
	}

	pipelinesList := make([]*pipelineExtendedInfo, 0)

	for _, pipeline := range pipelinesOutput.Pipelines {
		fmt.Printf("pipeline %s %s\n", *pipeline.Name, *pipeline.Updated)
		pipelineState, err := client.GetPipelineState(&codepipeline.GetPipelineStateInput{
			Name: pipeline.Name,
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to get pipeline state %s, %v\n", *pipeline.Name, err)
			continue
		}
		pipelineInfo, err := client.GetPipeline(&codepipeline.GetPipelineInput{
			Name: pipeline.Name,
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to get pipeline %s, %v\n", *pipeline.Name, err)
			continue
		}
		pipelinesList = append(pipelinesList, &pipelineExtendedInfo{
			state: pipelineState,
			info:  pipelineInfo,
		})
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
	ID         string
	ChangeID   string
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
		for _, stage := range pipe.state.StageStates {
			for _, action := range stage.ActionStates {
				a := actionObject{}
				if action.LatestExecution != nil {
					a.Status = *action.LatestExecution.Status
					a.Summary = derefString(action.LatestExecution.Summary)
					a.Author = derefString(action.LatestExecution.LastUpdatedBy)
					// a.URL = derefString(action.LatestExecution.ExternalExecutionUrl)
					a.Name = derefString(action.ActionName)
					a.ID = derefString(action.LatestExecution.ExternalExecutionId)
					if action.CurrentRevision != nil {
						a.Time = action.CurrentRevision.Created
						a.ChangeID = derefString(action.RevisionUrl)
					}
				}

				actions = append(actions, a)
			}
		}
		idx := 0
		for _, s := range pipe.info.Pipeline.Stages {
			for _, action := range s.Actions {
				q := url.Values{}
				q.Set("provider", *action.ActionTypeId.Provider)
				q.Set("id", actions[idx].ID)
				q.Set("account", fmt.Sprintf("%d", stage.ID))
				actions[idx].URL = fmt.Sprintf("./details?%s", q.Encode())
				idx = idx + 1
			}
		}
		pipes = append(pipes, pipelineObject{
			Name:    *pipe.state.PipelineName,
			Actions: actions,
		})
	}
	return pipelinesOutput{
		Pipelines: pipes,
		Name:      stage.Name,
	}, nil
}

// SetupStage sets up a PipeStage with configuration struct
func SetupStage(id int, configuration stageConfiguration) PipeStage {
	stage := PipeStage{}
	stage.setupClient(id, configuration)
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

	for idx, config := range configs {
		pipeStages = append(pipeStages, SetupStage(idx, config))
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
