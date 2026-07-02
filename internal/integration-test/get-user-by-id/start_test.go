package getuserbyid

import (
	"context"
	"testing"

	"github.com/golangmonster/workspace-booking-service/internal/integration-test/containers"
)

func Test(t *testing.T) {
	ctx := context.Background()

	_, err := containers.WorkspaceBookingPostgres(ctx)
	if err != nil {
		t.Error(err)
	}
}
