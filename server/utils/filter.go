package utils

import (
	"net/http"
	"net/url"
	"path"
	"sort"
	"strconv"
	"strings"

	"github.com/codeday-labs/bookworms/server/db"
)

const ReviewsCollection string = "review"

func RemoveDuplicatesAndSort(arraySlice []string) []string {
	list := RemoveDuplicateUtility(arraySlice)

	sort.Strings(list)

	var output []string

	for _, item := range list {
		output = append(output, strings.Title(item))
	}

	return output
}

func RemoveDuplicateUtility(arraySlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	for _, entry := range arraySlice {
		s := strings.ToLower(strings.Trim(entry, " "))
		if _, ok := keys[s]; !ok {
			keys[s] = true
			list = append(list, s)
		}
	}
	return list
}

func ConvertStringToNum(str string) (int, error) {
	return strconv.Atoi(str)
}

//FilterReviews based on filter passed
func FilterReviews(filter interface{}) ([]*db.Review, error) {

	// a slice to store decoded reviews
	var reviews []*db.Review

	DB, err := db.DB()

	// close db connection
	defer func() {
		if err := DB.Client().Disconnect(db.Ctx); err != nil {
			panic(err)
		}
	}()

	if err != nil {
		return nil, err
	}

	cur, err := DB.Collection(ReviewsCollection).Find(db.Ctx, filter)

	// once once done iterating the cursor close
	defer func() {
		cur.Close(db.Ctx)
	}()

	if err != nil {
		return reviews, err
	}

	// Iterate over a returned cursor and decode each at a time

	for cur.Next(db.Ctx) {
		var r db.Review
		err := cur.Decode(&r)
		if err != nil {
			return reviews, err
		}

		reviews = append(reviews, &r)
	}

	if err := cur.Err(); err != nil {
		return reviews, err
	}

	return reviews, nil
}

func GetIdFromUrl(r *http.Request) string {
	u, _ := url.Parse(r.URL.RequestURI())

	return path.Base(u.Path)
}
