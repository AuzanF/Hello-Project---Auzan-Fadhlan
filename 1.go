package main

import "fmt"

const nmax = 100

type provinsi struct { //membuat type
	nama     string
	populasi int
	tumbuh   float64
}

type DataProvinsi struct { //membuat type
	tabel     [nmax]provinsi
	nProvinsi int
}

func cariProvinsi(data DataProvinsi, nama string) provinsi { //function untuk memunculkan data provinsi
	var hasil provinsi          
	data.nProvinsi = 0          
	for data.nProvinsi < nmax { 
		if nama == data.tabel[data.nProvinsi].nama { 
			hasil = data.tabel[data.nProvinsi] 
			break
		}
		data.nProvinsi++
	}
	return hasil //mengembalikan hasil
}

func prediksi(data DataProvinsi, nama string) int { //function untuk memprediksi data selanjutnya
	var hasil int
	data.nProvinsi = 0
	for data.nProvinsi < nmax { //perulangan untuk mencari data yg sesuai
		if nama == data.tabel[data.nProvinsi].nama { 
			hasil = data.tabel[data.nProvinsi].populasi + (data.tabel[data.nProvinsi].populasi * int(data.tabel[data.nProvinsi].tumbuh))
				}
		data.nProvinsi++
	}
	return hasil //mengembalikan hasil
}

func urutData(data DataProvinsi) { //prosedur mengurutkan data descending
	var o, p, q, r int 
	var temp1 string
	var temp2 int
	var temp3 float64
	o = len(data.tabel) - 1
	p = 0
	q = p + 1
	for p <= o { //perulangan dalam perulangan untuk mencari data lebih besar agar ditaruh di depan
		for q <= o {
			if data.tabel[q].populasi > data.tabel[p].populasi { 
				temp1 = data.tabel[p].nama 
				temp2 = data.tabel[p].populasi
				temp3 = data.tabel[p].tumbuh
				data.tabel[p].nama = data.tabel[q].nama
				data.tabel[p].populasi = data.tabel[q].populasi
				data.tabel[p].tumbuh = data.tabel[q].tumbuh
				data.tabel[q].nama = temp1
				data.tabel[q].populasi = temp2
				data.tabel[q].tumbuh = temp3
			}
			q++
		}
		p++
		q = p + 1
	}
	r = 0
	for r < data.nProvinsi { 
		fmt.Println(data.tabel[r])
		r++
	}

}

func main() { 
	var arr DataProvinsi 
	var nama1, nama2 string

	arr.nProvinsi = 0                                                                                                                 
	fmt.Scanln(&arr.tabel[arr.nProvinsi].nama, &arr.tabel[arr.nProvinsi].populasi, &arr.tabel[arr.nProvinsi].tumbuh)                  
	for arr.tabel[arr.nProvinsi].nama != "NONE" && arr.tabel[arr.nProvinsi].populasi != 0 && arr.tabel[arr.nProvinsi].tumbuh != 0.0 { 
		arr.nProvinsi++
		fmt.Scanln(&arr.tabel[arr.nProvinsi].nama, &arr.tabel[arr.nProvinsi].populasi, &arr.tabel[arr.nProvinsi].tumbuh)

	}
	
	fmt.Print("Nama provinsi ? ")
	fmt.Scanln(&nama1)
	fmt.Println(cariProvinsi(arr, nama1))
	fmt.Print("Prediksi populasi tahun depan provinsi? ")
	fmt.Scanln(&nama2)
	fmt.Println(prediksi(arr, nama1))
	urutData(arr)

}
