package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Food struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       *string            `bson:"name" validate:"required"`
	Price      *float64           `bson:"price" validate:"required"`
	Food_image *string            `bson:"food_image" validate:"required"`
	Created_at time.Time          `bson:"created_at"`
	Updated_at time.Time          `bson:"updated_at"`
	Food_id    string             `bson:"food_id"`
	Menu_id    *string            `bson:"menu_id" validate:"required"`
}

type Menu struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string             `json:"name" validate:"required"`
	Category   string             `json:"category" validate:"required"`
	Start_Date *time.Time         `json:"start_date"`
	End_Date   *time.Time         `json:"end_date"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
	Menu_id    string             `json:"food_id"`
}

type Invoice struct {
	ID               primitive.ObjectID `bson:"_id"`
	Invoice_id       string             `json:"invoice_id"`
	Order_id         string             `json:"order_id"`
	Payment_method   *string            `json:"payment_method" validate:"eq=CARD|eq=CASH|eq="`
	Payment_status   *string            `json:"payment_status" validate:"required,eq=PENDING|eq=PAID"`
	Payment_due_date time.Time          `json:"Payment_due_date"`
	Created_at       time.Time          `json:"created_at"`
	Updated_at       time.Time          `json:"updated_at"`
}

type Note struct {
	ID         primitive.ObjectID `bson:"_id"`
	Text       string             `json:"text"`
	Title      string             `json:"title"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
	Note_id    string             `json:"note_id"`
}

type Order struct {
	ID         primitive.ObjectID `bson:"_id"`
	Order_Date time.Time          `json:"order_date" validate:"required"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
	Order_id   string             `json:"order_id"`
	Table_id   *string            `json:"table_id" validate:"required"`
}

type OrderItem struct {
	ID            primitive.ObjectID `bson:"_id"`
	Quantity      *string            `json:"quantity" validate:"required,eq=S|eq=M|eq=L"`
	Unit_price    *float64           `json:"unit_price" validate:"required"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
	Food_id       *string            `json:"food_id" validate:"required"`
	Order_item_id string             `json:"order_item_id"`
	Order_id      string             `json:"order_id" validate:"required"`
}

type Table struct {
	ID               primitive.ObjectID `bson:"_id"`
	Number_of_guests *int               `json:"number_of_guests" validate:"required"`
	Table_number     *int               `json:"table_number" validate:"required"`
	Created_at       time.Time          `json:"created_at"`
	Updated_at       time.Time          `json:"updated_at"`
	Table_id         string             `json:"table_id"`
}

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	First_name    *string            `json:"first_name" validate:"required,min=2,max=100"`
	Last_name     *string            `json:"last_name" validate:"required,min=2,max=100"`
	Password      *string            `json:"Password" validate:"required,min=6"`
	Email         *string            `json:"email" validate:"email,required"`
	Avatar        *string            `json:"avatar"`
	Phone         *string            `json:"phone" validate:"required"`
	Token         *string            `json:"token"`
	Refresh_Token *string            `json:"refresh_token"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
	User_id       string             `json:"user_id"`
}
