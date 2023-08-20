package service

type ItemsService interface {
	GetAll()
	GetOne()
	Create()
	Update()
	Delete()
}

type Blob struct {
	UserId    string
	Timestamp string
	ItemType  string
	Blob      string
	ItemId    string
}

type BlobService struct {
	db any //TODO mongodb something
}

func NewBlobService(db any) *BlobService {
	return &BlobService{
		db: db,
	}
}

func (s BlobService) GetAll() {
	// fetch user with id
}

func (s BlobService) GetOne() {
	// create user
}

func (s BlobService) Create() {
	// update user
}

func (s BlobService) Update() {
	// delete user
}

func (s BlobService) Delete() {
	// delete user
}