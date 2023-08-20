package main

type Items struct {
	UserId    string
	Timestamp string
	ItemType  string
	Blob      string
	ItemId    string
}

type Users struct {
	Username string
	Email    string
	Password []byte
}

func (i *Items) GetAll() {
	// return all items associated with this userID
}

func (i *Items) GetOne() {
	// return items associated with this userID and itemID
}

func (i *Items) Create() {
	// Create items associated with this userID and itemID
}

func (i *Items) Update() {
	// Update item
}

func (i *Items) Delete(){
	// Delete
}

