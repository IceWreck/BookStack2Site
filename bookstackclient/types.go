package bookstackclient

import "time"

// Shelves represents the BookStack API endpoint /api/shelves
type Shelves struct {
	Data []struct {
		ID          int       `json:"id"`
		Name        string    `json:"name"`
		Slug        string    `json:"slug"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		CreatedBy   int       `json:"created_by"`
		UpdatedBy   int       `json:"updated_by"`
		OwnedBy     int       `json:"owned_by"`
		ImageID     int       `json:"image_id"`
	} `json:"data"`
	Total int `json:"total"`
}

// Shelf represents the BookStack API endpoint /api/shelves/{id}
type Shelf struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	CreatedBy   struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"created_by"`
	UpdatedBy struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"updated_by"`
	OwnedBy struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"owned_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Tags      []struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Value string `json:"value"`
		Order int    `json:"order"`
	} `json:"tags"`
	Cover struct {
		ID         int       `json:"id"`
		Name       string    `json:"name"`
		URL        string    `json:"url"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
		CreatedBy  int       `json:"created_by"`
		UpdatedBy  int       `json:"updated_by"`
		Path       string    `json:"path"`
		Type       string    `json:"type"`
		UploadedTo int       `json:"uploaded_to"`
	} `json:"cover"`
	Books []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"books"`
}

// Books represents the BookStack API endpoint /api/books
type Books struct {
	Data []struct {
		ID          int       `json:"id"`
		Name        string    `json:"name"`
		Slug        string    `json:"slug"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		CreatedBy   int       `json:"created_by"`
		UpdatedBy   int       `json:"updated_by"`
		OwnedBy     int       `json:"owned_by"`
		ImageID     int       `json:"image_id"`
	} `json:"data"`
	Total int `json:"total"`
}

// Book represents the BookStack API endpoint /api/books/{id}
type Book struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedBy   struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"created_by"`
	UpdatedBy struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"updated_by"`
	OwnedBy struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"owned_by"`
	Tags []struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Value string `json:"value"`
		Order int    `json:"order"`
	} `json:"tags"`
	Cover struct {
		ID         int       `json:"id"`
		Name       string    `json:"name"`
		URL        string    `json:"url"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
		CreatedBy  int       `json:"created_by"`
		UpdatedBy  int       `json:"updated_by"`
		Path       string    `json:"path"`
		Type       string    `json:"type"`
		UploadedTo int       `json:"uploaded_to"`
	} `json:"cover"`
}

// Chapters represents the BookStack API endpoint /api/chapters
type Chapters struct {
	Data []struct {
		ID          int       `json:"id"`
		BookID      int       `json:"book_id"`
		Name        string    `json:"name"`
		Slug        string    `json:"slug"`
		Description string    `json:"description"`
		Priority    int       `json:"priority"`
		CreatedAt   string    `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		CreatedBy   int       `json:"created_by"`
		UpdatedBy   int       `json:"updated_by"`
		OwnedBy     int       `json:"owned_by"`
	} `json:"data"`
	Total int `json:"total"`
}

// Chapter represents the BookStack API endpoint /api/chapters/{id}
type Chapter struct {
	ID          int       `json:"id"`
	BookID      int       `json:"book_id"`
	Slug        string    `json:"slug"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Priority    int       `json:"priority"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedBy   struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"created_by"`
	UpdatedBy struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"updated_by"`
	OwnedBy struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"owned_by"`
	Tags []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
		Order int    `json:"order"`
	} `json:"tags"`
	Pages []struct {
		ID            int       `json:"id"`
		BookID        int       `json:"book_id"`
		ChapterID     int       `json:"chapter_id"`
		Name          string    `json:"name"`
		Slug          string    `json:"slug"`
		Priority      int       `json:"priority"`
		CreatedAt     time.Time `json:"created_at"`
		UpdatedAt     time.Time `json:"updated_at"`
		CreatedBy     int       `json:"created_by"`
		UpdatedBy     int       `json:"updated_by"`
		Draft         bool      `json:"draft"`
		RevisionCount int       `json:"revision_count"`
		Template      bool      `json:"template"`
	} `json:"pages"`
}

// Pages represents the BookStack API endpoint /api/pages
type Pages struct {
	Data []struct {
		ID        int       `json:"id"`
		BookID    int       `json:"book_id"`
		ChapterID int       `json:"chapter_id"`
		Name      string    `json:"name"`
		Slug      string    `json:"slug"`
		Priority  int       `json:"priority"`
		Draft     bool      `json:"draft"`
		Template  bool      `json:"template"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		CreatedBy int       `json:"created_by"`
		UpdatedBy int       `json:"updated_by"`
		OwnedBy   int       `json:"owned_by"`
	} `json:"data"`
	Total int `json:"total"`
}

// Page represents the BookStack API endpoint /api/pages/{id}
type Page struct {
	ID        int       `json:"id"`
	BookID    int       `json:"book_id"`
	ChapterID int       `json:"chapter_id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	HTML      string    `json:"html"`
	Priority  int       `json:"priority"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedBy struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"created_by"`
	UpdatedBy struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"updated_by"`
	OwnedBy struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"owned_by"`
	Draft         bool   `json:"draft"`
	Markdown      string `json:"markdown"`
	RevisionCount int    `json:"revision_count"`
	Template      bool   `json:"template"`
	Tags          []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
		Order int    `json:"order"`
	} `json:"tags"`
}

// Wiki is the BookStack2Site representation of a Wiki
type Wiki struct {
	Name  string
	Books []WikiBook
	// one book can be on multiple shelves
	Shelves []WikiShelf
}

type WikiShelf struct {
	ShelfID int
	Name    string
	Slug    string
	BookIDs []int
}

type WikiChapter struct {
	ChapterID int
	Name      string
	Slug      string
	Priority  int
	Pages     []WikiPage
}

type WikiPage struct {
	PageID   int
	Name     string
	Slug     string
	Priority int
	// FilePath is path of file relative to root of wiki. To be filled while downloading.
	// Should start with a '/'.
	// Should not end with a '/'.
	FilePath string
}
type WikiBook struct {
	BookID int
	Name   string
	Slug   string
	// a book can have pages in chapters or independent pages
	Chapters   []WikiChapter
	IndiePages []WikiPage
}
