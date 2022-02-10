package models

type Brend struct {
	Id   uint64 `json:"Id "`
	Name string `json:"Name"`
	Slug string `json:"Slug "`
}

type Image struct {
	Id      uint64 `json:"Id "`
	Name    string `json:"Name"`
	Storage string `json:"Storage "`
	Path    string `json:"Path"`
}
type Images struct {
	Images []Image `json:"Image"`
}
type Video struct {
	Id      uint64 `json:"Id "`
	Name    string `json:"Name"`
	Storage string `json:"Storage "`
	Path    string `json:"Path"`
}
type Videos struct {
	Videos []Video `json:"Video"`
}

type Unit struct {
	Id   uint64 `json:"Id "`
	Name string `json:"Name"`
	Slug string `json:"Slug "`
}

type Attribut struct {
	Id   uint64 `json:"Id "`
	Name string `json:"Name"`
	Slug string `json:"Slug "`
	Unit Unit   `json:"Unit "`
}

type Attribut_value struct {
	Id       uint64   `json:"Id "`
	Attribut Attribut `json:"Attribut"`
	Vslue    string   `json:"Value"`
}

type Produkt struct {
	Id                uint64 `json:"Id"`
	Name              string `json:"Name"`
	Slug              string `json:"Slug"`
	Brend             Brend  `json:"Brend"`
	SKU               string `json:"SKU "`
	Short_description string `json:"short_description"`
	Sull_description  string `json:"full_description"`
	Sort              string `json:"sort"`
	Image             Image
	Video             Video
	Attribut_value    Attribut_value `json:"Attribut_value"`
}
type Products struct {
	Products []Produkt `json:"Produkt"`
}

var DB []Produkt

func init() {
	produkt1 := Produkt{
		Id:   1,
		Name: "",
		Slug: "",
		Brend: Brend{
			Id:   1,
			Name: "",
			Slug: "",
		},
		SKU:               "",
		Short_description: "",
		Sull_description:  "",
		Image: Image{
			Id:      1,
			Name:    "",
			Storage: " ",
			Path:    " ",
		},
		Video: Video{
			Id:      1,
			Name:    "",
			Storage: " ",
			Path:    " ",
		},
	}

	produkt2 := Produkt{
		Id:   2,
		Name: "",
		Slug: "",
		Brend: Brend{
			Id:   2,
			Name: "",
			Slug: "",
		},
		SKU:               "",
		Short_description: "",
		Sull_description:  "",
		Image: Image{
			Id:      2,
			Name:    "",
			Storage: " ",
			Path:    " ",
		},
		Video: Video{
			Id:      2,
			Name:    "",
			Storage: " ",
			Path:    " ",
		},
	}
	DB = append(DB, produkt1, produkt2)
}

func FindProductById(id uint64) (Produkt, bool) {
	var product Produkt
	var found bool
	for _, p := range DB {
		if p.Id == id {
			product = p
			found = true
			break
		}
	}
	return product, found
}

func DeleteBookById(id uint64) bool {
	var found bool
	for i, p := range DB {
		if p.Id == id {
			DB[i] = DB[len(DB)-1]

			// 2. Удалить последний элемент (записать нулевое значение).
			DB[len(DB)-1] = Produkt{}

			// 3. Усечь срез.
			DB = DB[:len(DB)-1]

			found = true
			break
		}
	}
	return found
}
