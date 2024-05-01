package persistence

import (
	"fmt"
	"libraryInventory/internal/domain"
)

type BookPersistenceAdapter struct {
	bookRepository     BookRepository
	bookCopyRepository BookCopyRepository
}

func NewBookPersistenceAdapter(bookRepository BookRepository, copyBookRepository BookCopyRepository) BookPersistenceAdapter {
	return BookPersistenceAdapter{
		bookRepository:     bookRepository,
		bookCopyRepository: copyBookRepository,
	}
}

func (b BookPersistenceAdapter) LoadBook(isbn string) (*domain.Book, error) {
	book, err := b.bookRepository.FindByISBN(isbn)
	if err != nil {
		fmt.Println("Err:", err)
		return nil, err
	}
	fmt.Println("Book:", book)
	copies, err := b.bookCopyRepository.ListByBookID(book.ID)
	if err != nil {
		return nil, err
	}
	fmt.Println("Copies:", copies)

	return mapToBookDomain(book, copies), nil
}

/*
type InvoiceSpecification struct{}

type DelinquentInvoiceSpecification struct {
	currentDate time.Time
}

type Invoice struct {
	DueDate time.Time
}

type InvoiceRepository struct{}

var invoices []Invoice

func selectSatisfying(spec InvoiceSpecification) []Invoice {
	var results []Invoice
	for _, invoice := range invoices {
		if spec.IsSatisfiedBy(invoice) {
			results = append(results, invoice)
		}
	}
	return results
}

func ex() {
	delinquentInvoices := invoiceRepository.selectSatisfying(
		NewDelinquentInvoiceSpecification(time.Now()))
}

func (d *DelinquentInvoiceSpecification) satisfyingElementsFrom(repository InvoiceRepository) bool {
	//delinquent rule is defined as:
	// "grace period past as of current date"
	return repository.selectWhereGracePeriodPast(d.currentDate)
}

type Container struct {
	drums    []Drum
	features []ContainerFeature
}

func (c *Container) ContainsFeature(feature ContainerFeature) bool {
	for _, ft := range c.features {
		if ft.Code == feature.Code {
			return true
		}
	}
	return false
}

type Drum struct {
	Chemical
}
type ContainerFeature struct {
	Code int
}

type ContainerSpecification struct {
	requiredFeature ContainerFeature
}

type Chemical struct {
	containerSpecification ContainerSpecification
}

func (c *Chemical) SetContainerSpecification(spec ContainerSpecification) {
	c.containerSpecification = spec
}

func (c *Chemical) GetContainerSpecification() ContainerSpecification {
	return c.containerSpecification
}

var ARMORED ContainerFeature

func NewContainerSpecification(rFeature ContainerFeature) ContainerSpecification {
	return ContainerSpecification{requiredFeature: rFeature}
}

func (cs ContainerSpecification) IsSatisfiedBy(aContainer Container) bool {
	return aContainer.ContainsFeature(cs.requiredFeature)
}

func client() {
	var tnt Chemical
	tnt.SetContainerSpecification(NewContainerSpecification(ARMORED))
}

func (c Container) IsSafelyPacked() bool {
	for _, drum := range c.drums {
		if drum.GetContainerSpecification().IsSatisfiedBy(c) {
			return false
		}
	}
	return true
}

func exampleInsecure() {
	var unsafeContainers []Container
	var containers []Container

	for _, container := range containers {
		if !container.IsSafelyPacked() {
			unsafeContainers = append(unsafeContainers, container)
		}
	}
}*/
