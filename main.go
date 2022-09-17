package main

import (
	. "fmt"
)

func main() {

	keranjangSatu := newKeranjang("Keranajang Ricky")
	barang1 := item{
		id:        1,
		nama_item: "Hodie H&M",
		qty:       4,
		harga:     120000,
	}
	keranjangSatu.tambahItem(barang1)

	items := []item{
		{
			id:        2,
			nama_item: "sepatu",
			qty:       1,
			harga:     200000,
		},
		{
			id:        3,
			nama_item: "celana",
			qty:       5,
			harga:     100000,
		},
		{
			id:        4,
			nama_item: "Sendal",
			qty:       2,
			harga:     20000,
		},
		{
			id:        5,
			nama_item: "Raket Yonex A444 E",
			qty:       1,
			harga:     35000,
		},
	}

	keranjangSatu.tambahItemBanyak(items...)
	Println(keranjangSatu)

	Println("=================== Setelah Update ======================")

	itemUpdate := item{
		id:        2,
		nama_item: "Pentopel",
		qty:       2,
		harga:     500000,
	}
	keranjangSatu.updateItem(itemUpdate)
	keranjangSatu.deleteItem(0)
	Println(keranjangSatu)
	Println(keranjangSatu.outputPrint())
}
