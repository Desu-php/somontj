package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
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

const PatterUrl = "https://somon.tj/api/items/%d/"

func GetDetailsById(id uint) (*Apartment, error) {
	url := fmt.Sprintf(PatterUrl, id)
	resp, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("error fetching data: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusForbidden {
		fmt.Println("Начинаю ждать 15 секунду...")
		time.Sleep(1 * time.Second * 15)
		fmt.Println("15 секунд прошла!")
		return GetDetailsById(id)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var apartment Apartment

	err = json.Unmarshal(body, &apartment)
	if err != nil {
		fmt.Println("Response body:", string(body))
		return nil, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	return &apartment, nil
}
