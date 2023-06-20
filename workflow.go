package moderation

import (
	"fmt"
	"log"
	"moderation/resources"
	"os"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func ModerationWorkflow(ctx workflow.Context, name string) (string, error) {
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

	//logger := workflow.GetLogger(ctx)
	log.Println("Moderation Started")

	log.Println("CHAT KEY: " + os.Getenv("CHATGPT_API_KEY"))
	log.Println("MODERATION URL: " + os.Getenv("MODERATION_URL"))

	// set activity inputs for moderating name
	activityInput := resources.ModerationInput{
		Url:  os.Getenv("MODERATION_URL"),
		Name: name,
	}

	// run activity to check provided name
	var moderationResult bool
	err := workflow.ExecuteActivity(ctx, ModerationActivity, activityInput).Get(ctx, &moderationResult)
	if err != nil {
		log.Fatalln("Activity failed.", "Error", err)
		return "", err
	}
	output := fmt.Sprintf("Moderation complete (for name: %s, result: %t)", name, moderationResult)
	log.Println(output)
	return output, nil
}
