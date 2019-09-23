package metadata

//Metadata contains all the data that describes user uploaded videos
type Metadata struct {
	Videoname       string
	Artist          string
	Songname        string
	TimeBucket      string
	Description     string
	Score           int
	Comments        int
	Views           int
	Likes, Dislikes int
	Laughs, Hash    int
	Tags            []string
	Date            int64
}
