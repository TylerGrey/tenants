package resolver

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/TylerGrey/tenants/internal/graphql/model"
)

// Search 주소 검색
func (r *queryResolver) Search(ctx context.Context, query string) ([]*model.SearchResult, error) {
	resolvers := []*model.SearchResult{}

	kakaoURL := fmt.Sprintf("%s?query=%s", os.Getenv("KAKAO_SEARCH_ADDRESS_URL"), url.QueryEscape(query))
	req, err := http.NewRequest("GET", kakaoURL, nil)
	if err != nil {
		return resolvers, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("KakaoAK %s", os.Getenv("KAKAO_REST_API_KEY")))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return resolvers, err
	}
	defer resp.Body.Close()

	searchResult := model.KakaoResponse{}
	bytes, _ := ioutil.ReadAll(resp.Body)

	if err = json.Unmarshal(bytes, &searchResult); err != nil {
		return resolvers, err
	}

	for _, d := range searchResult.Documents {
		if d.Address != nil && d.RoadAddress != nil {
			resolvers = append(resolvers, &model.SearchResult{
				Payload: d,
			})
		}
	}

	return resolvers, nil
}
