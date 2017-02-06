package design

import (
	d "github.com/goadesign/goa/design"
	a "github.com/goadesign/goa/design/apidsl"
)

var attachment = a.Type("Attachment", func() {
	a.Description(`JSONAPI store for the data of an attachment.  See also http://jsonapi.org/format/#document-resource-object`)
	a.Attribute("type", d.String, func() {
		a.Enum("attachments")
	})
	a.Attribute("id", d.UUID, "ID of attachment", func() {
		a.Example("40bbdd3d-8b5d-4fd6-ac90-7236b669af04")
	})
	a.Attribute("attributes", attachmentAttributes)
	a.Attribute("relationships", attachmentRelationships)
	a.Attribute("links", genericLinks)
	a.Required("type")
})

var createAttachment = a.Type("CreateAttachment", func() {
	a.Description(`JSONAPI store for the data of a attachment.  See also http://jsonapi.org/format/#document-resource-object`)
	a.Attribute("type", d.String, func() {
		a.Enum("comments")
	})
	a.Attribute("attributes", createAttachmentAttributes)
	a.Required("type", "attributes")
})

var attachmentAttributes = a.Type("AttachmentAttributes", func() {
	a.Description(`JSONAPI store for all the "attributes" of a comment. +See also see http://jsonapi.org/format/#document-resource-object-attributes`)
	a.Attribute("created-at", d.DateTime, "When the attachment was created", func() {
		a.Example("2016-11-29T23:18:14Z")
	})
	a.Attribute("location", d.String, "The location of the attachment", func() {
		a.Example("https://somelocation.s3.amazonaws.com/933f77387d0630ce03e5af1166ba1920.jpg")
	})
})

var createAttachmentAttributes = a.Type("CreateAttachmentAttributes", func() {
	a.Description(`JSONAPI store for all the "attributes" for creating an attachment. +See also see http://jsonapi.org/format/#document-resource-object-attributes`)
	a.Attribute("location", d.String, "The location of the attachment", func() {
		a.MinLength(1) // Empty comment not allowed
		a.Example("https://somelocation.s3.amazonaws.com/933f77387d0630ce03e5af1166ba1920.jpg")
	})
	a.Required("location")
})

var attachmentRelationships = a.Type("AttachmentRelations", func() {
	a.Attribute("created-by", commentCreatedBy, "This defines the created by relation")
	a.Attribute("parent", relationGeneric, "This defines the owning resource of the attachment")
})

var attachmentCreatedBy = a.Type("AttachmentCreatedBy", func() {
	a.Attribute("data", identityRelationData)
	a.Required("data")
})

var attachmentRelationshipsArray = JSONList(
	"AttachmentRelationship", "Holds the response of attachments",
	attachment,
	genericLinks,
	meta,
)

var attachmentSingle = JSONSingle(
	"Attachment", "Holds the response of a single attachment",
	attachment,
	nil,
)
var createSingleAttachment = JSONSingle(
	"CreateSingleAttachment", "Holds the create data for an attachment",
	createAttachment,
	nil,
)

var _ = a.Resource("work-item-attachments", func() {
	a.Parent("workitem")

	a.Action("create", func() {
		a.Security("jwt")
		a.Routing(
			a.POST("attachments"),
		)
		a.Description("Create a new attachment")
		a.Response(d.OK, func() {
			a.Media(attachmentSingle)
		})
		a.Payload(createSingleAttachment)
		a.Response(d.BadRequest, JSONAPIErrors)
		a.Response(d.Unauthorized, JSONAPIErrors)
		a.Response(d.InternalServerError, JSONAPIErrors)
		a.Response(d.NotFound, JSONAPIErrors)
	})
})
