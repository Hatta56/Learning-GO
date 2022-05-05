package book

type Service interface {
	FindAll() ([]Books, error)
	FindByID(id int) (Books, error)
	Create(book BookRequest) (Books, error)
	Update(ID int, book BookRequest) (Books, error)
	Delete(ID int) (Books, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Books, error) {
	return s.repository.FindAll()
}
func (s *service) FindByID(id int) (Books, error) {
	return s.repository.FindByID(id)
}
func (s *service) Create(book BookRequest) (Books, error) {
	price, _ := book.Price.Int64()
	rating, _ := book.Rating.Int64()
	discount, _ := book.Discount.Int64()
	Books := Books{
		Title:       book.Title,
		Price:       int(price),
		Description: book.Description,
		Rating:      int(rating),
		Discount:    int(discount),
	}
	return s.repository.Create(Books)
}

func (s *service) Update(ID int, BookRequest BookRequest) (Books, error) {
	book, _ := s.repository.FindByID(ID)

	price, _ := BookRequest.Price.Int64()
	rating, _ := BookRequest.Rating.Int64()
	discount, _ := BookRequest.Discount.Int64()

	book.Title = BookRequest.Title
	book.Price = int(price)
	book.Description = BookRequest.Description
	book.Rating = int(rating)
	book.Discount = int(discount)

	return s.repository.Update(book)
}
func (s *service) Delete(ID int) (Books, error) {
	book, _ := s.repository.FindByID(ID)
	return s.repository.Delete(book)
}
