package book

import "errors"

var books = make([]*Book, 0)

func GetBooks() ([]*Book, error) {
	// books := make([]*Book, 0)
	return books, nil
	// books = append(books, &Book{
	// 	Name: "",
	// })
}

// primitive type
// scalar slice, rune
// strings
//
// built-in

func AddBook(b Book) (Book, error) {
	lastId := 0
	for i := range books {
		if books[i].ID > lastId {
			lastId = books[i].ID + 1
		}
	}
	books = append(books, &b)
	return b, nil
}

func DeleteBook(bookid int) error {
	toDelete := -1
	for i := range books {
		if books[i].ID == bookid {
			toDelete = i
			break
		}
	}
	if toDelete == -1 {
		return errors.New("Not Found")
	} else {
		//todelete index deerh nomiig ustgah uildel
		books = append(books[:toDelete], books[toDelete+1:]...)
	}
	return nil
}

func Get(bookid int) (*Book, error) {
	for i := range books {
		if books[i].ID == bookid {
			return books[i], nil
		}
	}

	return nil, errors.New("not found")
}
