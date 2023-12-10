package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)


func HelloHandler(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Post Article\n")
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
		queryMap := req.URL.Query()
		var page int
		if p, ok := queryMap["page"]; ok && len(p) > 0 {
			var err error
			page, err = strconv.Atoi(p[0])
			if err != nil {
				http.Error(w, "Invalid query parameter", http.StatusBadRequest)
				return
			}
		} else {
			page = 1
		}
		resString := fmt.Sprintf("Article List (page %d)\n", page)
		io.WriteString(w, resString)
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
		articleId, err := strconv.Atoi(mux.Vars(req)["id"])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
		resString := fmt.Sprintf("Article No.%d\n", articleId)
		io.WriteString(w, resString)
}

func ArticleNiceHandler(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Nice...\n")
}

func ArticleCommentHandler(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Comment\n")
}
