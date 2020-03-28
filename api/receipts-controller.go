package receiptscontroller

import (
	"encoding/json"
	"fmt"
	"time"

	receiptinterface "github.com/j-ew-s/receipts-api/internals/interface"
	"github.com/j-ew-s/receipts-api/internals/usecase"

	"github.com/j-ew-s/receipts-api/internals/model"

	"github.com/valyala/fasthttp"
)

var (
	strContentType     = []byte("Content-Type")
	strApplicationJSON = []byte("application/json")
)

//ReceiptsController for
type ReceiptsController struct {
	receiptUseCase receiptinterface.UseCase
}

// Create insert a new Receipt.
//  Call Receipt UseCase
//  Prepare response and return model.ReceiptCreat
func Create(ctx *fasthttp.RequestCtx) {

	var receipt model.Receipt

	err := json.Unmarshal(ctx.PostBody(), &receipt)
	if err != nil {
		panic(err)
	}

	createdReceipt := usecase.Create(&receipt)
	if createdReceipt.Err != nil {
		ctx.Error(createdReceipt.Err.Error(), fasthttp.StatusInternalServerError)
	}

	ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)
	ctx.Response.SetStatusCode(createdReceipt.GetStatus())

	start := time.Now()
	if err := json.NewEncoder(ctx).Encode(createdReceipt); err != nil {
		elapsed := time.Since(start)
		fmt.Println(" ERROR : ", elapsed, err.Error(), createdReceipt.Err)
		fmt.Println(" Message : ", createdReceipt.Message)
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}

}

// Delete removes a receipt fisically
// that match the ID parameter
// uses Receipt object
func Delete(ctx *fasthttp.RequestCtx) {

	id := fmt.Sprintf("%v", ctx.UserValue("id"))

	receipt := usecase.GetByID(id)
	response := true

	if receipt.Err != nil {
		ctx.Error(receipt.Err.Error(), receipt.GetStatus())
		response = false
	}

	if &receipt != nil {
		err := usecase.Delete(id)

		if err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			response = false
		}

	}

	ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)

	ctx.Response.SetStatusCode(200)

	start := time.Now()
	if err := json.NewEncoder(ctx).Encode(response); err != nil {
		elapsed := time.Since(start)
		fmt.Println(" ERROR : ", elapsed, err.Error(), err)
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
}

// Get returns all the receipts
// No filtering is made
func Get(ctx *fasthttp.RequestCtx) {

	receipts := usecase.Get()
	if receipts.Err != nil {
		ctx.Error(receipts.Err.Error(), fasthttp.StatusInternalServerError)
	}

	ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)
	ctx.Response.SetStatusCode(receipts.GetStatus())

	start := time.Now()
	if err := json.NewEncoder(ctx).Encode(receipts); err != nil {
		elapsed := time.Since(start)
		fmt.Println(" ERROR : ", elapsed, err.Error(), receipts.Err)
		fmt.Println(" Message : ", receipts.Message)
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
}

// GetByID will return any receipt with related ID
func GetByID(ctx *fasthttp.RequestCtx) {

	id := fmt.Sprintf("%v", ctx.UserValue("id"))

	receipt := usecase.GetByID(id)
	if receipt.Err != nil {
		ctx.Error(receipt.Err.Error(), receipt.GetStatus())
	}

	ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)

	var httpStatusCode = 404

	if &receipt != nil {
		httpStatusCode = 200
	}

	ctx.Response.SetStatusCode(httpStatusCode)

	start := time.Now()
	if err := json.NewEncoder(ctx).Encode(receipt); err != nil {
		elapsed := time.Since(start)
		fmt.Println(" ERROR : ", elapsed, err.Error(), err)
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}

}

// GetByTags performs search for a receipt Tags
// that contais the search parameter
// search is made upon only on Tag property
func GetByTags(ctx *fasthttp.RequestCtx) {

	tags := []string{}

	tagsQueryByte := ctx.QueryArgs().PeekMulti("tags")

	for i := 0; i < len(tagsQueryByte); i++ {
		tags = append(tags, string(tagsQueryByte[i]))
	}

	receipts := usecase.GetByTags(tags)
	if receipts.Err != nil {
		ctx.Error(receipts.Err.Error(), fasthttp.StatusInternalServerError)
	}

	ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)
	ctx.Response.SetStatusCode(receipts.GetStatus())

	start := time.Now()
	if err := json.NewEncoder(ctx).Encode(receipts); err != nil {
		elapsed := time.Since(start)
		fmt.Println(" ERROR : ", elapsed, err.Error(), receipts.Err)
		fmt.Println(" Message : ", receipts.Message)
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}

}

// Put updates a receipt
// that match the ID parameter
// uses Receipt object
func Put(ctx *fasthttp.RequestCtx) {

	ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)

	var httpStatusCode = 404

	var receipt model.Receipt

	id := fmt.Sprintf("%v", ctx.UserValue("id"))
	err := json.Unmarshal(ctx.PostBody(), &receipt)

	if err != nil {
		httpStatusCode = 500
	}

	httpStatusCode, err = usecase.Update(&receipt, id)

	ctx.Response.SetStatusCode(httpStatusCode)

	start := time.Now()
	if err := json.NewEncoder(ctx).Encode(true); err != nil {
		elapsed := time.Since(start)
		fmt.Println(" ERROR : ", elapsed, err.Error(), err)
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}

}

// Search performs search for receipts
// that contains the search term.
// Search is made upon all the filds
func Search(ctx *fasthttp.RequestCtx) {

	fmt.Println("Entrou no Search")

}
