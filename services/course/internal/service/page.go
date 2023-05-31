package service

import (
	"context"

	"github.com/FoxFurry/scholarlabs/services/course/internal/store"
)

func (p *service) CreatePage(ctx context.Context, page store.Page) error {
	if _, err := p.db.GetCourseByUUID(ctx, page.CourseUUID); err != nil {
		return handleDBError(err, "could not get course")
	}

	numberOfPages, err := p.db.GetNumberOfPagesForCourse(ctx, page.CourseUUID)
	if err != nil {
		return handleDBError(err, "could not get number of pages")
	}
	page.PageID = uint64(numberOfPages + 1)

	return p.db.CreatePage(ctx, page)
}

func (p *service) DeletePage(ctx context.Context, courseUUID, pageUUID string) error {
	if _, err := p.db.GetCourseByUUID(ctx, courseUUID); err != nil {
		return handleDBError(err, "could not get course")
	}

	return p.db.DeletePage(ctx, courseUUID, pageUUID)
}

func (p *service) GetPageByID(ctx context.Context, courseUUID, pageUUID string) (*store.Page, error) {
	if _, err := p.db.GetCourseByUUID(ctx, courseUUID); err != nil {
		return nil, handleDBError(err, "could not get course")
	}

	page, err := p.db.GetPageByID(ctx, courseUUID, pageUUID)
	if err != nil {
		return nil, handleDBError(err, "could not get page")
	}

	return page, nil
}

func (p *service) GetCourseToc(ctx context.Context, courseUUID string) ([]store.PageIdentifier, error) {
	if _, err := p.db.GetCourseByUUID(ctx, courseUUID); err != nil {
		return nil, handleDBError(err, "could not get course")
	}

	return p.db.GetCourseToC(ctx, courseUUID)
}
