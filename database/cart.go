package database

import "errors"

var (
	ErrCantFindProduct = errors.New("can't find product")
	ErrCantDecodeProducts = errors.New("can't fint the products")
	ErrUserIdIsNotValid = errors.New("this user is not valid")
	ErrCantUpdateUser = errors.New("can't add this product to the cart")
	ErrCantRemoveItemCart = errors.New("can't remove this item from the cart")
	ErrCantGetItem = errors.New("we unable to get the item from the cart")
	ErrCantBuyCartItem = errors.New("cannot update the purchase")
)


func AddProductToCart() {
	
}

func RemoveCartItem() {

}

func BuyItemFromCart() {

}

func InstantBuyer() {
	
}