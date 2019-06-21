package main

import "fmt"

func main(){
    // var nama string = "Tama"
    // var namatengah = "Istighfar"
    // namatengah := "ahmad"
    // var member uint16 = 300


    // fmt.Printf("Hello %s %s",nama,namatengah)
    // fmt.Println(member)

    // var buah = []string{"apel","mangga","angggur"}
    
    //SLICE=BERUPA ARRAY DINAMIS
    // buah = append(buah,"lemon") // penambahan pada slice/array yang dinamis
    // fmt.Println(buah)
    // fmt.Println("Jumlah buah :",len(buah))


    //MAP= ARRAY YANG DIBERIKAN TIPE DATA BERBEDA DENGAN VARIABELNYA
    // var harga_makanan = map[string]int{"ayam":10000,"nasgor":15000}
    // fmt.Println("Harga nasgor", harga_makanan["ayam"])

    
    //PENJUMLAHAN DENGAN INPUTAN DARI USER
    var x,y int

    fmt.Printf("Masukkan Nilai X : ")
    fmt.Scanln(&x)
    fmt.Printf("Nilai X : %d \n\n",x)
    
    fmt.Printf("Masukkan Nilai Y : ")
    fmt.Scanln(&y)
    fmt.Printf("Nilai Y : %d \n\n",y)
  
    fmt.Printf("Jumlah %d + %d adalah %d \n",x,y,x+y)
  
}