package moderation

import (
	"context"

	"github.com/ktenzer/temporal-trivia/resources"
	"go.temporal.io/sdk/activity"

	// TODO(cretz): Remove when tagged
	_ "go.temporal.io/sdk/contrib/tools/workflowcheck/determinism"
)

func ModerationActivity(ctx context.Context, input resources.ActivityInput) (map[int]resources.Result, error) {
	logger := activity.GetLogger(ctx)

	logger.Info("ModerationActivity")
}
