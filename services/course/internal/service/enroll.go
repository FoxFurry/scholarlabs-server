package service

import (
	"context"
)

func (p *service) Enroll(ctx context.Context, userUUID, courseUUID string) error {
	return p.db.Enroll(ctx, userUUID, courseUUID)
}

func (p *service) Unenroll(ctx context.Context, userUUID, courseUUID string) error {
	return p.db.Unenroll(ctx, userUUID, courseUUID)
}
