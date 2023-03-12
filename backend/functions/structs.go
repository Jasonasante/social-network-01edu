package functions

type User struct {
	Id        int    `json:"-"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Firstname string `json:"first"`
	Lastname  string `json:"last"`
	DOB       string `json:"dob"`
	Avatar    string `json:"avatar"`
	Nickname  string `json:"nickname"`
	Aboutme   string `json:"about"`
	Followers int    `json:"followers"`
	Following int    `json:"following"`
	Status    string `json:"status"`
}

type Session struct {
	sessionUUID string
	userID      string
	email       string
}

type ChatRoomFields struct {
	Id          string `json:"chatroom-id"`
	Avatar      string `json:"chat-avatar"`
	Name        string `json:"chat-name"`
	Description string `json:"chat-description"`
	Type        string `json:"chat-type"`
	Users       string `json:"users"`
	Admin       string `json:"admin"`
	Action      string `json:"action"`
	Date        int    `json:"last-message-date"`
}

type ChatroomType struct {
	Private []ChatRoomFields `json:"private-chatrooms"`
	Group   []ChatRoomFields `json:"group-chatrooms"`
}

type ChatFields struct {
	Id        string `json:"id"`
	Sender    string `json:"sender"`
	MessageId string `json:"message-id"`
	Message   string `json:"message"`
	Date      int    `json:"date"`
}

type Follow struct {
	Follower string `json:"follower"`
	Followee string `json:"followee"`
}

type followMessage struct {
	FollowRequest         string `json:"followRequest"`
	ToFollow              string `json:"toFollow"`
	IsFollowing           bool   `json:"isFollowing"`
	Followers             int    `json:"followers"`
	FollowRequestUsername string `json:"followRequest-username"`
	FolloweeUsername      string `json:"toFollow-username"`
	FollowRequestAccepted bool   `json:"followRequest-accepted"`
}

type followNotification struct {
	UpdateUser             string `json:"updateUser"`
	Followers              int    `json:"followers"`
	FollowerFollowingCount int    `json:"followerFollowingCount"`
}

type RequestNotifcationFields struct {
	FollowRequest followMessage           `json:"notification-followRequest"`
	GroupRequest  GroupFields             `json:"notification-groupRequest"`
	GroupAction   GroupAcceptNotification `json:"notification-group-action"`
	Sender        string                  `json:"remove-sender"`
	Receiver      string                  `json:"remove-receiver"`
	TypeOfAction  string                  `json:"remove-typeOfAction"`
	GroupId       string                  `json:"remove-groupId"`
}

type GroupAcceptNotification struct {
	User        string `json:"user"`
	Admin       string `json:"admin"`
	Action      string `json:"action"`
	GroupName   string `json:"group-name"`
	GroupAvatar string `json:"group-avatar"`
	GroupId     string `json:"groupId"`
}

type ChatNotifcationFields struct {
	ChatId        string `json:"notification-chatId"`
	Sender        string `json:"notification-sender"`
	Receiver      string `json:"notification-receiver"`
	NumOfMessages int    `json:"notification-numOfMessages"`
	Date          int    `json:"notification-date"`
}

type OpenChatInfo struct {
	User             string         `json:"user"`
	Chatroom         ChatRoomFields `json:"chatroom"`
	PreviousMessages []ChatFields   `json:"previous-messages"`
}

type PostFields struct {
	Id           string `json:"post-id"`
	Author       string `json:"author"`
	AuthorImg    string `json:"author-img"`
	Image        string `json:"post-image"`
	Text         string `json:"post-text-content"`
	Thread       string `json:"post-threads"`
	Likes        int    `json:"post-likes"`
	PostLiked    bool   `json:"post-liked"`
	Dislikes     int    `json:"post-dislikes"`
	PostDisliked bool   `json:"post-disliked"`
	PostComments int    `json:"post-comments"`
	PostAuthor   bool   `json:"post-author"`
	Time         int    `json:"post-time"`
	Privacy      string `json:"privacy"`
	Viewers      string `json:"viewers"`
	Error        string `json:"error"`
}

type LikesFields struct {
	PostId   string `json:"post-id"`
	Username string `json:"username"`
	Like     string `json:"like"`
	Type     string `json:"type"`
}

type CommentFields struct {
	CommentId       string `json:"comment-id"`
	PostId          string `json:"post-id"`
	Author          string `json:"author"`
	AuthorImg       string `json:"author-img"`
	Image           string `json:"comment-image"`
	Text            string `json:"comment-text"`
	Thread          string `json:"comment-threads"`
	Time            int    `json:"comment-time"`
	CommentLiked    bool   `json:"comment-liked"`
	Likes           int    `json:"comment-likes"`
	CommentDisliked bool   `json:"comment-disliked"`
	Dislikes        int    `json:"comment-dislikes"`
	CommentAuthor   bool   `json:"comment-author"`
	Error           string `json:"error"`
}

type ReturnComments struct {
	TotalComments []CommentFields `json:"total-comments"`
	Post          PostFields      `json:"post-comment"`
}

type CommentsAndLikesFields struct {
	CommentId string `json:"comment-id"`
	Username  string `json:"username"`
	Like      string `json:"like"`
	Type      string `json:"type"`
}

type GroupFields struct {
	Id          string `json:"group-id"`
	Avatar      string `json:"group-avatar"`
	Name        string `json:"group-name"`
	Description string `json:"group-description"`
	Users       string `json:"users"`
	Admin       string `json:"admin"`
	Action      string `json:"action"`
	Date        int    `json:"last-post-sent"`
}

type GroupPostFields struct {
	Id           string      `json:"group-post-id"`
	Group        GroupFields `json:"group"`
	PostId       string      `json:"post-id"`
	Author       string      `json:"author"`
	AuthorImg    string      `json:"author-img"`
	Image        string      `json:"post-image"`
	Text         string      `json:"post-text-content"`
	Thread       string      `json:"post-threads"`
	Likes        int         `json:"post-likes"`
	PostLiked    bool        `json:"post-liked"`
	PostComments int         `json:"post-comments"`
	Dislikes     int         `json:"post-dislikes"`
	PostDisliked bool        `json:"post-disliked"`
	PostAuthor   bool        `json:"post-author"`
	Time         int         `json:"post-time"`
	Error        string      `json:"error"`
}

type GroupsAndLikesFields struct {
	PostId   string `json:"post-id"`
	Username string `json:"username"`
	Like     string `json:"like"`
	Type     string `json:"type"`
}

type ReturnGroupComments struct {
	TotalComments []CommentFields `json:"total-comments"`
	Post          GroupPostFields `json:"post-comment"`
}

type GroupEventFields struct {
	GroupId        string `json:"group-id"`
	EventId        string `json:"event-id"`
	Organiser      string `json:"event-organiser"`
	EventOrganiser bool   `json:"event-organiser-user"`
	Title          string `json:"event-title"`
	Description    string `json:"event-description"`
	Time           int    `json:"event-time"`
	ActiveEvent    bool   `json:"active"`
	Attendees      int    `json:"attendees"`
	Attending      bool   `json:"event-attending"`
	Error          string `json:"error"`
}

type EventAttendanceFields struct {
	EventId string `json:"event-id"`
	User    string `json:"attending-user"`
	Status  string `json:"attending-status"`
	Error   string `json:"error"`
}

type UpdateStatus struct {
	User      string `json:"user"`
	SetStatus string `json:"setStatus"`
}
