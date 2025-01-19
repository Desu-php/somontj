package stats

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

type User struct {
	ID               int      `json:"id"`
	Name             string   `json:"name"`
	HasEmail         bool     `json:"has_email"`
	Verified         bool     `json:"verified"`
	AccountType      string   `json:"account_type"`
	CustomFields     []string `json:"custom_fields"`
	ShowPhoneMessage string   `json:"show_phone_message"`
	Joined           string   `json:"joined"`
	Phone            *string  `json:"phone"`
	Logo             string   `json:"logo"`
	CompanyName      string   `json:"company_name"`
	LegalName        string   `json:"legal_name"`
	Website          string   `json:"website"`
	ContactPhone     *string  `json:"contact_phone"`
	CardHeaderColour string   `json:"card_header_colour"`
}

type Image struct {
	ID         int    `json:"id"`
	URL        string `json:"url"`
	Orig       string `json:"orig"`
	IsFlatplan bool   `json:"is_flatplan"`
}

type Attributes struct {
	Feet       int    `json:"attrs__feet"`
	Type       int    `json:"attrs__type"`
	Floor      int    `json:"attrs__floor"`
	Remont     int    `json:"attrs__remont"`
	Sanuzel    int    `json:"attrs__sanuzel"`
	District   string `json:"attrs__district"`
	Otoplenie  int    `json:"attrs__otoplenie"`
	Sostoyanie int    `json:"attrs__sostoyanie"`
}

type Breadcrumb struct {
	ID   int    `json:"id"`
	Path string `json:"path"`
	Name string `json:"name"`
}

type Permissions struct {
	Phone    string `json:"phone"`
	Whatsapp string `json:"whatsapp"`
	Email    string `json:"email"`
	CVForm   string `json:"cv_form"`
	Chat     string `json:"chat"`
	Delivery string `json:"delivery"`
}

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Apartment struct {
	ID                    int                    `json:"id"`
	Title                 string                 `json:"title"`
	Slug                  string                 `json:"slug"`
	Rubric                int                    `json:"rubric"`
	Description           string                 `json:"description"`
	City                  int                    `json:"city"`
	User                  User                   `json:"user"`
	Images                []Image                `json:"images"`
	Attributes            Attributes             `json:"attrs"`
	Price                 string                 `json:"price"`
	StartPrice            string                 `json:"start_price"`
	Contacts              map[string]interface{} `json:"contacts"`
	HitCount              int                    `json:"hit_count"`
	Currency              string                 `json:"currency"`
	PhoneHitCount         int                    `json:"phone_hitcount"`
	RaiseDt               string                 `json:"raise_dt"`
	CreatedDt             string                 `json:"created_dt"`
	OwnerAdvertCount      int                    `json:"owner_advert_count"`
	Coordinates           *Coordinates           `json:"coordinates"`
	Zoom                  *int                   `json:"zoom"`
	NegotiablePrice       bool                   `json:"negotiable_price"`
	Exchange              bool                   `json:"exchange"`
	PriceDescription      string                 `json:"price_description"`
	InTop                 bool                   `json:"in_top"`
	InPremium             bool                   `json:"in_premium"`
	IsEditable            bool                   `json:"is_editable"`
	IsFavorite            bool                   `json:"is_favorite"`
	CityDistricts         []string               `json:"city_districts"`
	Flatplan              bool                   `json:"flatplan"`
	VideoLink             string                 `json:"video_link"`
	CreditType            *string                `json:"credit_type"`
	CreditAttrs           *string                `json:"credit_attrs"`
	CreditLink            *string                `json:"credit_link"`
	Breadcrumbs           []Breadcrumb           `json:"breadcrumbs"`
	TemplatedTitle        string                 `json:"templated_title"`
	CloudinaryVideo       map[string]interface{} `json:"cloudinary_video"`
	Whatsapp              string                 `json:"whatsapp"`
	Viber                 *string                `json:"viber"`
	IMEIChecked           bool                   `json:"imei_checked"`
	IMEIInfo              []string               `json:"imei_info"`
	PhoneBenchmarkResults []string               `json:"phone_benchmark_results"`
	ExternalID            string                 `json:"external_id"`
	ItemLink              string                 `json:"item_link"`
	VirtualTourLink       string                 `json:"virtual_tour_link"`
	SquareMeterPrice      *string                `json:"square_meter_price"`
	IsCarCheck            bool                   `json:"is_carcheck"`
	Delivery              bool                   `json:"delivery"`
	HasOnlineViewing      bool                   `json:"has_online_viewing"`
	HasCarCheckReport     bool                   `json:"has_carcheck_report"`
	HasFreeCarCheckReport bool                   `json:"has_free_carcheck_report"`
	CategoryType          string                 `json:"category_type"`
	NewInStockLabel       bool                   `json:"new_in_stock_label"`
	NewToOrderLabel       bool                   `json:"new_to_order_label"`
	PriceFrom             bool                   `json:"price_from"`
	ShowSendForm          bool                   `json:"show_send_form"`
	ShowWhatsappBtn       bool                   `json:"show_whatsapp_btn"`
	Permissions           Permissions            `json:"permissions"`
	CurrencyID            int                    `json:"currency_id"`
}

func Run() {
	file, err := os.OpenFile("apartments.json", os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("ошибка открытия файла: %v", err)
		os.Exit(1)
	}
	defer file.Close()

	var apartments []Apartment
	stat, err := file.Stat()
	if err != nil {
		fmt.Printf("ошибка чтения файла: %v", err)
		os.Exit(1)
	}

	// Если файл не пустой, читать его содержимое
	if stat.Size() > 0 {
		data, err := io.ReadAll(file)
		if err != nil {
			fmt.Printf("ошибка чтения данных из файла: %v", err)
			os.Exit(1)
		}

		err = json.Unmarshal(data, &apartments)
		if err != nil {
			fmt.Printf("ошибка парсинга JSON: %v", err)
			os.Exit(1)
		}
	}

	dist := make(map[string]int)

	// Радиус в километрах
	radius := 1.0 // Например, 5 км
	var centers []Apartment

	for _, apartment := range apartments {
		if apartment.Coordinates != nil {
			if isNearHujandCenter(apartment.Coordinates.Latitude, apartment.Coordinates.Longitude, radius) {
				centers = append(centers, apartment)
			}
		} else if containsWord(apartment.Title) {
			centers = append(centers, apartment)
		}

		if keyExists(dist, apartment.Attributes.District) {
			dist[apartment.Attributes.District]++
		} else {
			dist[apartment.Attributes.District] = 1
		}
	}

	// for name, count := range dist {
	// 	fmt.Printf("%s: %d \n", name, count)
	// }

	for _, centerApartment := range centers {
		fmt.Println(centerApartment.Title)
		fmt.Printf(" https://somon.tj/adv/%d_%s\n", centerApartment.ID, centerApartment.Slug)
	}
}

func keyExists[K comparable, V any](m map[K]V, key K) bool {
	_, exists := m[key]
	return exists
}

// Радиус Земли в километрах
const EarthRadius = 6371.0

// Конвертирует градусы в радианы
func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

// Вычисляет расстояние между двумя точками на поверхности Земли по формуле Хаверсина
func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	// Конвертируем в радианы
	lat1 = degreesToRadians(lat1)
	lon1 = degreesToRadians(lon1)
	lat2 = degreesToRadians(lat2)
	lon2 = degreesToRadians(lon2)

	// Разница широт и долгот
	deltaLat := lat2 - lat1
	deltaLon := lon2 - lon1

	// Формула Хаверсина
	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(lat1)*math.Cos(lat2)*math.Sin(deltaLon/2)*math.Sin(deltaLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	// Расстояние
	return EarthRadius * c
}

// Проверка, если координаты находятся в радиусе от центра города Худжанд
func isNearHujandCenter(lat, lon float64, radius float64) bool {
	// Центр Худжанда
	hujandLat := 40.285158
	hujandLon := 69.618972

	// Рассчитываем расстояние от центра Худжанда
	distance := haversine(lat, lon, hujandLat, hujandLon)

	// Если расстояние меньше радиуса, то точка находится в радиусе от центра
	return distance <= radius
}

func containsWord(word string) bool {
	searchTerms := []string{"универмаг",
		"центр",
		"бог",
		"ватан",
	}

	for _, term := range searchTerms {
		if strings.Contains(strings.ToLower(word), strings.ToLower(term)) {
			return true
		}
	}

	return false
}
