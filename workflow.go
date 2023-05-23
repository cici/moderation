package moderation

import (
	"os"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

type (
	ModerationInput struct {
		Key  string `json:"key"`
		Name string `json:"name"`
	}
)

func Workflow(ctx workflow.Context, name string) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second,
			BackoffCoefficient: 2.0,
			MaximumInterval:    time.Minute,
			MaximumAttempts:    5,
		},
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("Moderation Started")

	logger.Error("CHAT KEY: " + os.Getenv("CHATGPT_API_KEY"))

	// set activity inputs for moderating name
	activityInput := ModerationInput{
		Key:  os.Getenv("CHATGPT_API_KEY"),
		Name: workflowInput.name,
	}

	// run activity to check provided name
	var result string
	err = workflow.ExecuteActivity(ctx, activities.ModerationActivity, activityInput).Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return err
	}

	return nil
}
