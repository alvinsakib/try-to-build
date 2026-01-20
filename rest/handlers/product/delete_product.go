package product

import (
	"net/http"
	"strconv"
	"ecommerce/util"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pId, err := strconv.Atoi(productID)
	if err != nil {
		//http.Error(w, "Please give me a valid product id", 400)  // badrequest
		util.SendError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	err = h.productRepo.Delete(pId)
	if err != nil {
		//http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	util.SendData(w, http.StatusOK ,"Successfully deleted product") //201
}