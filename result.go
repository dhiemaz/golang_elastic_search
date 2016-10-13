package main

// Root Fields //
type Post struct {
   Took       int       `json:"took"`
   Timed_out  bool    	`json:"timed_out"`
   Shards    _Shards    `json:"_shards"`
   Hits 	  Hits 		`json:"hits"`
}


// _Shards Fields //
type _Shards struct {
   Total   		int     `json:"total"`
   Successful 	int  	`json:"successful"`
   Failed		int		`json:"failed"`
}

// Hits Fields //
type Hits struct {
	Total 		int 	`json:"Total"`
	Max_score 	float32	`json:"max_score"`
	Hits 		[]Hit 	`json:"hits"`
}


// Hits Fields //
type Hit struct {
	Index 		string	`json:"_index"`
	Type		string 	`json:"_type"`
	Id 			string 	`json:"_id"`
	Score 		float32 `json:"_score"`
	Source 		_Source `json:"_source"`
}

// Source Fields //
type _Source struct {
	Id 				int 	`json:"id"`
	Name			string 	`json:"name"`
	Name_id 		string 	`json:"name_id"`
	Brand_id		string 	`json:"brand_id"`
	Brand 			string 	`json:"brand"`
	Desc 			string 	`json:"desc"`
	Vid 			int 	`json:"vid"`
	Vname 			string 	`json:"vname"`
	Vidname 		string 	`json:"vidname"`
	Ref_id 			string 	`json:"ref_id"`
	Price 			int32	`json:"price"`
	Ori_price 		int32 	`json:"ori_price"`
	Disc 			int 	`json:"disc"`
	Item_komisi 	int 	`json:"item_komisi"`
	Thumb_img 		string 	`json:"thumb_img"`
	Exp_date 		string 	`json:"exp_date"`
	Exp_time 		int 	`json:"exp_time"`
	Exp_type 		string 	`json:"exp_type"`
	Ui_type 		int 	`json:"ui_type"`
	Tags 			string 	`json:"tags"`
	Category 		[]Category `json:"category"`
	Category_id 	int 	`json:"category_id"`
	Published 		bool 	`json:"published"`
	Type 			string 	`json:"type"`
	Selling_count 	int 	`json:"selling_count"`
	View_count 		int 	`json:"view_count"`
	Store_date 		string 	`json:"store_date"`
	Created_at 		int32 	`json:"created_at"`
}

// Category Fields //
type Category struct {
	Id 			string 	`json:"id"`
	Name 		string 	`json:"Name"`
	Parent_id 	string 	`json:"parent_id"`
}
