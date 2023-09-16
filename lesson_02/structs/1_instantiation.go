package structs

var globalPrivateUser = user{Name: "Greg", Surname: "Tertyshny", Age: 28}

func AccessPrivateUser() *user {
	return &globalPrivateUser
}

func NewPost() *Post {
	return new(Post)
}

func SameNewPost() *Post {
	return &Post{}
}

func UsingVar() *Post {
	var usingVar Post

	return &usingVar
}
