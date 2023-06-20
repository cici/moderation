package resources

type ModerationInput struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

/*
	{
	  "id": "modr-XXXXX",
	  "model": "text-moderation-001",
	  "results": [
	    {
	      "categories": {
	        "hate": false,
	        "hate/threatening": false,
	        "self-harm": false,
	        "sexual": false,
	        "sexual/minors": false,
	        "violence": false,
	        "violence/graphic": false
	      },
	      "category_scores": {
	        "hate": 0.18805529177188873,
	        "hate/threatening": 0.0001250059431185946,
	        "self-harm": 0.0003706029092427343,
	        "sexual": 0.0008735615410842001,
	        "sexual/minors": 0.0007470346172340214,
	        "violence": 0.0041268812492489815,
	        "violence/graphic": 0.00023186142789199948
	      },
	      "flagged": false
	    }
	  ]
	}
*/

type ModerationResponse struct {
	ID      string `json:"id"`
	Model   string `json:"model"`
	Results []Result
}

type Categories struct {
	Hate        bool `json:"hate"`
	Threatening bool `json:"threatening"`
	Selfharm    bool `json:"selfharm"`
	Sexual      bool `json:"sexual"`
	Minors      bool `json:"minors"`
	Violence    bool `json:"violence"`
	Graphic     bool `json:"graphic"`
}

type CategoryScores struct {
	Hate        float64 `json:"hate"`
	Threatening float64 `json:"threatening"`
	Selfharm    float64 `json:"selfharm"`
	Sexual      float64 `json:"sexual"`
	Minors      float64 `json:"minors"`
	Violence    float64 `json:"violence"`
	Graphic     float64 `json:"graphic"`
}

type Result struct {
	Categories     Categories
	CategoryScores CategoryScores
	Flagged        bool
}
