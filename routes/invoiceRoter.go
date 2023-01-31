package routes

import (
	controller "github/souravdev-eng/resturant/controllers"

	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(r *gin.Engine) {
	r.GET("/invoices", controller.GetInvoices())
	r.GET("/invoices/:invoice_id", controller.GetInvoice())
	r.POST("/invoices", controller.CreateInvoice())
	r.PATCH("/invoices/:invoice_id", controller.UpdateInvoice())
}
