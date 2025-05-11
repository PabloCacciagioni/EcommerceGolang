package database

import "errors"

var (
	ErrCantFindProduct    = errors.New("Can't find the product")
	ErrCantDecodeProducts = errors.New("Can't find the product")
	ErrUserIdIsNotValid   = errors.New("This user id is not valid")
	ErrCantUpdateUser     = errors.New("Cannot add this product to the cart")
	ErrCantRemoveItemCart = errors.New("Cannot remove this item from the cart")
	ErrCantGetItem        = errors.New("Was unable to get the item from the cart")
	ErrCantBuyCartItem    = errors.New("Cannot update the purchase")
)

func AddProductToCart() {

}

func RemoveCartItem() {

}

func BuyFromCart() {

}

func InstantBuy() {

}
