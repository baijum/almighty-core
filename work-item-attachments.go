package main

import (
	"github.com/almighty/almighty-core/app"
	"github.com/almighty/almighty-core/application"
	"github.com/goadesign/goa"
)

// WorkItemAttachmentsController implements the work-item-attachments resource.
type WorkItemAttachmentsController struct {
	*goa.Controller
	db application.DB
}

// NewWorkItemAttachmentsController creates a work-item-attachments controller.
func NewWorkItemAttachmentsController(service *goa.Service, db application.DB) *WorkItemAttachmentsController {
	return &WorkItemAttachmentsController{Controller: service.NewController("WorkItemAttachmentsController"), db: db}
}

// Create runs the create action.
func (c *WorkItemAttachmentsController) Create(ctx *app.CreateWorkItemAttachmentsContext) error {
	// WorkItemAttachmentsController_Create: start_implement

	// Put your logic here

	// WorkItemAttachmentsController_Create: end_implement
	res := &app.AttachmentSingle{}
	return ctx.OK(res)
}
