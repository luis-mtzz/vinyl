package models

type CollectionItems struct {
	Pagination struct {
		Items   int `json:"items"`
		Page    int `json:"page"`
		Pages   int `json:"pages"`
		PerPage int `json:"per_page"`
		Urls    struct {
		} `json:"urls"`
	} `json:"pagination"`
	Releases []struct {
		BasicInformation struct {
			Artists []struct {
				Anv         string `json:"anv"`
				ID          int    `json:"id"`
				Join        string `json:"join"`
				Name        string `json:"name"`
				ResourceURL string `json:"resource_url"`
				Role        string `json:"role"`
				Tracks      string `json:"tracks"`
			} `json:"artists"`
			CoverImage string `json:"cover_image"`
			Formats    []struct {
				Descriptions []string `json:"descriptions"`
				Name         string   `json:"name"`
				Qty          string   `json:"qty"`
				Text         string   `json:"text"`
			} `json:"formats"`
			Genres []string `json:"genres"`
			ID     int      `json:"id"`
			Labels []struct {
				Catno          string `json:"catno"`
				EntityType     string `json:"entity_type"`
				EntityTypeName string `json:"entity_type_name"`
				ID             int    `json:"id"`
				Name           string `json:"name"`
				ResourceURL    string `json:"resource_url"`
			} `json:"labels"`
			MasterID    int      `json:"master_id"`
			MasterURL   string   `json:"master_url"`
			ResourceURL string   `json:"resource_url"`
			Styles      []string `json:"styles"`
			Thumb       string   `json:"thumb"`
			Title       string   `json:"title"`
			Year        int      `json:"year"`
		} `json:"basic_information"`
		DateAdded  string `json:"date_added"`
		ID         int    `json:"id"`
		InstanceID int    `json:"instance_id"`
		Rating     int    `json:"rating"`
	} `json:"releases"`
}
