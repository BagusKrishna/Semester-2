package main

import "fmt"

const NMAX = 1000

type date struct {
	dd1, mm1, yy1, dd2, mm2, yy2 int
}

type tabDate [NMAX]date

type provinsi struct {
	calon            calon
	pemilih          pemilih
	pilih            int
	totalSuaraCalon  int
	totalSuaraPartai int
}

type calon struct {
	nama, partai string
}

type pemilih struct {
	nama    string
	pilihan int
}

type arrProvA [NMAX]provinsi
type arrProvB [NMAX]provinsi
type arrProvC [NMAX]provinsi

func main() {
	var A date
	var TA arrProvA
	var TB arrProvB
	var TC arrProvC
	var nA, nB, nC int
	var nPA, nPB, nPC int

	menuUtama(&A, &TA, &TB, &TC, &nA, &nB, &nC, &nPA, &nPB, &nPC)

}

func menuUtama(A *date, TA *arrProvA, TB *arrProvB, TC *arrProvC, nA, nB, nC, nPA, nPB, nPC *int) {

	var pilihan int
	fmt.Println()
	fmt.Println("=======================================")
	fmt.Println("  *Selamat Datang di Aplikasi PEMILU*  ")
	fmt.Println("=======================================")
	fmt.Println()
	fmt.Println("Masuk sebagai : ")
	fmt.Println("1. KPU")
	fmt.Println("2. Pemilih")
	fmt.Println("3. Selesai")

	fmt.Println()
	fmt.Print("Pilihan anda : ")
	fmt.Scan(&pilihan)
	if pilihan == 1 {
		menuKPU(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
	} else if pilihan == 2 {
		menuPemilih(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
	} else if pilihan == 3 {
		hasilPEMILU(*TA, *TB, *TC, *nA, *nB, *nC, *nPA, *nPB, *nPC)
		fmt.Println()
		fmt.Println("----------------------------------")
		fmt.Println("Terima Kasih")
	} else {
		fmt.Println("Pilihan salah")
		menuUtama(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
	}

}

func menuPemilih(A *date, TA *arrProvA, TB *arrProvB, TC *arrProvC, nA, nB, nC, nPA, nPB, nPC *int) {

	var tanggal, bulan, tahun int

	fmt.Println()
	fmt.Println("==========================")
	fmt.Println("Anda Masuk Sebagai Pemilih")
	fmt.Println("==========================")
	fmt.Println()

	fmt.Println("Tanggal saat ini : (dd/m/yyyy)")

	fmt.Scan(&tanggal, &bulan, &tahun)

	if cekTanggal(*A, tanggal, bulan) == false {
		fmt.Println("----------------------------------------------")
		fmt.Println("                  * Maaf *                   ")
		fmt.Println("    tidak ada PEMILU pada tanggal tersebut   ")
		fmt.Println("----------------------------------------------")
		fmt.Println()
		cetakCalonA(*TA, *nA)
		cetakCalonB(*TB, *nB)
		cetakCalonC(*TC, *nC)
		menuUtama(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
	} else {
		menuMasukPemilih(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
	}

}

func isiPemilihA(TA *arrProvA, nPA *int) {
	var nama string
	var pilihan int

	fmt.Println("Masukkan data pemilih Provinsi A: ")
	fmt.Scan(&nama)
	for nama != "STOP" {
		fmt.Scan(&pilihan)
		TA[*nPA].pemilih.nama = nama
		TA[*nPA].pemilih.pilihan = pilihan
		*nPA++
		fmt.Scan(&nama)
	}
}

func isiPemilihB(TB *arrProvB, nPB *int) {
	var nama string
	var pilihan int

	fmt.Println("Masukkan data pemilih Provinsi B: ")
	fmt.Scan(&nama)
	for nama != "STOP" {
		fmt.Scan(&pilihan)
		TB[*nPB].pemilih.nama = nama
		TB[*nPB].pemilih.pilihan = pilihan
		*nPB++
		fmt.Scan(&nama)
	}
}

func isiPemilihC(TC *arrProvC, nPC *int) {
	var nama string
	var pilihan int

	fmt.Println("Masukkan data pemilih Provinsi C: ")
	fmt.Scan(&nama)
	for nama != "STOP" {
		fmt.Scan(&pilihan)
		TC[*nPC].pemilih.nama = nama
		TC[*nPC].pemilih.pilihan = pilihan
		*nPC++
		fmt.Scan(&nama)
	}
}

func menuMasukPemilih(A *date, TA *arrProvA, TB *arrProvB, TC *arrProvC, nA, nB, nC, nPA, nPB, nPC *int) {
	var pilihan int
	fmt.Println("Pilih Provinsi Anda :")
	fmt.Println("1. Provinsi A")
	fmt.Println("2. Provinsi B")
	fmt.Println("3. Provinsi C")
	fmt.Println("4. Kembali")
	fmt.Println()

	fmt.Print("Pilihan anda : ")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		fmt.Println("1. Cari data calon")
		fmt.Println("2. Mulai memilih!")
		fmt.Println()
		fmt.Print("Pilihan anda : ")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			fmt.Println("1. Cari berdasarkan nama")
			fmt.Println("2. Cari berdasarkan partai")
			fmt.Println()
			fmt.Print("Pilihan anda : ")
			fmt.Scan(&pilihan)
			if pilihan == 1 {
				searchNamaCalonA(*TA, *nA)
				menuMasukPemilih(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
			} else if pilihan == 2 {
				searchPartaiCalonA(*TA, *nA)
				menuMasukPemilih(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
			}
		} else if pilihan == 2 {
			cetakCalonA(*TA, *nA)
			isiPemilihA(TA, nPA)
			fmt.Println("=============================================")
			fmt.Println("    *   Terima Kasih sudah memilih   *  ")
			fmt.Println("=============================================")
			menuUtama(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
		}

	} else if pilihan == 2 {
		fmt.Println("1. Cari data calon")
		fmt.Println("2. Mulai memilih!")
		fmt.Println()
		fmt.Print("Pilihan anda : ")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			fmt.Println("1. Cari berdasarkan nama")
			fmt.Println("2. Cari berdasarkan partai")
			fmt.Println()
			fmt.Print("Pilihan anda : ")
			fmt.Scan(&pilihan)
			if pilihan == 1 {
				searchNamaCalonB(*TB, *nB)
				menuMasukPemilih(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
			} else if pilihan == 2 {
				searchPartaiCalonB(*TB, *nB)
				menuMasukPemilih(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
			}
		} else if pilihan == 2 {
			cetakCalonB(*TB, *nB)
			isiPemilihB(TB, nPB)
			fmt.Println("=============================================")
			fmt.Println("    *   Terima Kasih sudah memilih   *  ")
			fmt.Println("=============================================")
			menuUtama(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
		}

	} else if pilihan == 3 {
		fmt.Println("1. Cari data calon")
		fmt.Println("2. Mulai memilih!")
		fmt.Println()
		fmt.Print("Pilihan anda : ")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			fmt.Println("1. Cari berdasarkan nama")
			fmt.Println("2. Cari berdasarkan partai")
			fmt.Println()
			fmt.Print("Pilihan anda : ")
			fmt.Scan(&pilihan)
			if pilihan == 1 {
				searchNamaCalonC(*TC, *nC)
				menuMasukPemilih(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
			} else if pilihan == 2 {
				searchPartaiCalonC(*TC, *nA)
				menuMasukPemilih(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
			}
		} else if pilihan == 2 {
			cetakCalonC(*TC, *nC)
			isiPemilihC(TC, nPC)
			fmt.Println("=============================================")
			fmt.Println("    *   Terima Kasih sudah memilih   *  ")
			fmt.Println("=============================================")
			menuUtama(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
		}

	} else if pilihan == 4 {
		menuUtama(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
	}
}

func menuDataCalon(TA *arrProvA, TB *arrProvB, TC *arrProvC, nA, nB, nC *int) {

	fmt.Println("==================")
	fmt.Println("  * DATA CALON *  ")
	fmt.Println("==================")

	//tampilkan data calon
	cetakCalonA(*TA, *nA)
	cetakCalonB(*TB, *nB)
	cetakCalonC(*TC, *nC)

	fmt.Println()
	fmt.Println("Daftar Menu : ")
	fmt.Println("1. Tambah Data Calon")
	fmt.Println("2. Cari Data Calon")
	fmt.Println("3. Kembali")

	fmt.Println()

}
func menuDataPemilih(TA *arrProvA, TB *arrProvB, TC *arrProvC, nA, nB, nC, nPA, nPB, nPC *int) {

	fmt.Println("==================")
	fmt.Println(" * DATA PEMILIH * ")
	fmt.Println("==================")
	//tampilkan data calon

	cetakPemilihA(*TA, *nPA)
	cetakPemilihB(*TB, *nPB)
	cetakPemilihC(*TC, *nPC)

	fmt.Println()
	fmt.Println("Daftar Menu : ")
	fmt.Println("1. Cari Data Pemilih")
	fmt.Println("2. Kembali")

	fmt.Println()

}

func menuCariPemilih() {
	fmt.Println("-----------------")
	fmt.Println("Pilih Provinsi : ")
	fmt.Println("1. Provinsi A")
	fmt.Println("2. Provinsi B")
	fmt.Println("3. Provinsi C")
	fmt.Println("4. Kembali")

	fmt.Println()
}

func menuCariCalon() {
	fmt.Println("-----------------")
	fmt.Println("Pilih Provinsi : ")
	fmt.Println("1. Provinsi A")
	fmt.Println("2. Provinsi B")
	fmt.Println("3. Provinsi C")
	fmt.Println("4. Kembali")

	fmt.Println()
}

func dataCALON(A *date, TA *arrProvA, TB *arrProvB, TC *arrProvC, nA, nB, nC, nPA, nPB, nPC *int) {
	var pilihan int

	menuDataCalon(TA, TB, TC, nA, nB, nC)

	fmt.Print("Pilihan anda : ")
	fmt.Scan(&pilihan)

	if pilihan == 1 /* tambah data calon*/ {
		isiCalon(TA, TB, TC, nA, nB, nC)
		dataCALON(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
	}

	if pilihan == 2 /* cari data calon*/ {
		menuCariCalon()
		fmt.Print("Pilihan anda : ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			//tampilkan provinsi A
			cetakCalonA(*TA, *nA)
			fmt.Println("--------------")
			fmt.Println("Daftar Menu : ")
			fmt.Println("1. Edit data calon")
			fmt.Println("2. Hapus data calon")
			fmt.Println("3. Kembali")
			fmt.Println()
			fmt.Print("Pilihan anda : ")
			fmt.Scan(&pilihan)
			var noUrut int
			var bener bool
			if pilihan == 1 /* edit data calon*/ {

				fmt.Println("---------------------------------------")
				fmt.Print("Masukkan no calon yang ingin di edit : ")
				fmt.Scan(&noUrut)
				fmt.Println("---------------------------------------")

				for !bener {
					if noUrut <= *nA {
						editCalonA(TA, noUrut, *nA)
						bener = true
					} else {
						fmt.Println("No calon tidak ditemukan")
						fmt.Println("Masukkan no calon yang ingin di edit : ")
						bener = false
						fmt.Scan(&noUrut)

					}
				}
				dataCALON(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)

			} else if pilihan == 2 /*hapus data calon*/ {
				fmt.Println("----------------------------------------")
				fmt.Print("Masukkan no calon yang ingin di hapus : ")
				fmt.Scan(&noUrut)
				fmt.Println("----------------------------------------")

				for !bener {
					if noUrut <= *nA {
						hapusCalonA(TA, noUrut, nA)
						bener = true
					} else {
						fmt.Println("No calon tidak ditemukan")
						fmt.Println("Masukkan no calon yang ingin di hapus : ")
						bener = false
						fmt.Scan(&noUrut)

					}
				}
				dataCALON(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)

			}

		} else if pilihan == 2 /* provinsi B */ {
			//tampilkan provinsi B
			cetakCalonB(*TB, *nB)
			fmt.Println("--------------")
			fmt.Println("Daftar Menu : ")
			fmt.Println("1. Edit data calon")
			fmt.Println("2. Hapus data calon")
			fmt.Println("3. Kembali")
			fmt.Println()
			fmt.Print("Pilihan anda : ")
			fmt.Scan(&pilihan)
			var noUrut int
			var bener bool
			if pilihan == 1 /* edit data calon*/ {

				fmt.Println("---------------------------------------")
				fmt.Print("Masukkan no calon yang ingin di edit : ")
				fmt.Scan(&noUrut)
				fmt.Println("---------------------------------------")

				for !bener {
					if noUrut <= *nB {
						editCalonB(TB, noUrut, *nB)
						bener = true
					} else {
						fmt.Println("No calon tidak ditemukan")
						fmt.Println("Masukkan no calon yang ingin di edit : ")
						bener = false
						fmt.Scan(&noUrut)

					}
				}
				dataCALON(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)

			} else if pilihan == 2 /*hapus data calon*/ {
				fmt.Println("----------------------------------------")
				fmt.Print("Masukkan no calon yang ingin di hapus : ")
				fmt.Scan(&noUrut)
				fmt.Println("----------------------------------------")

				for !bener {
					if noUrut <= *nB {
						hapusCalonB(TB, noUrut, nB)
						bener = true
					} else {
						fmt.Println("No calon tidak ditemukan")
						fmt.Println("Masukkan no calon yang ingin di hapus : ")
						bener = false
						fmt.Scan(&noUrut)

					}
				}
				dataCALON(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)

			}
		} else if pilihan == 3 /* provinsi C */ {
			//tampilkan provinsi C
			cetakCalonC(*TC, *nC)
			fmt.Println("--------------")
			fmt.Println("Daftar Menu : ")
			fmt.Println("1. Edit data calon")
			fmt.Println("2. Hapus data calon")
			fmt.Println("3. Kembali")
			fmt.Println()
			fmt.Print("Pilihan anda : ")
			fmt.Scan(&pilihan)
			var noUrut int
			var bener bool
			if pilihan == 1 /* edit data calon*/ {

				fmt.Println("---------------------------------------")
				fmt.Print("Masukkan no calon yang ingin di edit : ")
				fmt.Scan(&noUrut)
				fmt.Println("---------------------------------------")

				for !bener {
					if noUrut <= *nC {
						editCalonC(TC, noUrut, *nC)
						bener = true
					} else {
						fmt.Println("No calon tidak ditemukan")
						fmt.Println("Masukkan no calon yang ingin di edit : ")
						bener = false
						fmt.Scan(&noUrut)

					}
				}
				dataCALON(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)

			} else if pilihan == 2 /*hapus data calon*/ {
				fmt.Println("----------------------------------------")
				fmt.Print("Masukkan no calon yang ingin di hapus : ")
				fmt.Scan(&noUrut)
				fmt.Println("----------------------------------------")

				for !bener {
					if noUrut <= *nC {
						hapusCalonC(TC, noUrut, nC)
						bener = true
					} else {
						fmt.Println("No calon tidak ditemukan")
						fmt.Println("Masukkan no calon yang ingin di hapus : ")
						bener = false
						fmt.Scan(&noUrut)

					}
				}
				dataCALON(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)

			}
		}
	}

	if pilihan == 3 /*kembali*/ {
		menuKPU(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
	}
}

func dataPEMILIH(A *date, TA *arrProvA, TB *arrProvB, TC *arrProvC, nA, nB, nC, nPA, nPB, nPC *int) {
	var pilihan int
	menuDataPemilih(TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)

	fmt.Print("Pilihan anda : ")
	fmt.Scan(&pilihan)

	if pilihan == 1 /* cari pemilih */ {
		menuCariPemilih()
		fmt.Print("Pilihan anda : ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			searchPemilihA(*TA, *nPA)
			dataPEMILIH(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
			fmt.Println()
			fmt.Println("--------------------------------------------")
			fmt.Println("--------------------------------------------")
			fmt.Println()
			cetakPemilihA(*TA, *nPA)
		} else if pilihan == 2 {
			searchPemilihB(*TB, *nPB)
			dataPEMILIH(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
			fmt.Println()
			fmt.Println("--------------------------------------------")
			fmt.Println("--------------------------------------------")
			fmt.Println()
			cetakPemilihB(*TB, *nPB)
		} else if pilihan == 3 {
			searchPemilihC(*TC, *nPC)
			dataPEMILIH(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
			fmt.Println()
			fmt.Println("--------------------------------------------")
			fmt.Println("--------------------------------------------")
			fmt.Println()
			cetakPemilihC(*TC, *nPC)
		}
	}

	if pilihan == 2 {
		menuKPU(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
	}

}

func menuKPU(A *date, TA *arrProvA, TB *arrProvB, TC *arrProvC, nA, nB, nC, nPA, nPB, nPC *int) {
	var pilihan int

	fmt.Println()
	fmt.Println("======================")
	fmt.Println("Anda Masuk Sebagai KPU")
	fmt.Println("======================")
	fmt.Println()
	fmt.Println("Daftar Menu : ")
	fmt.Println("1. Data Calon")
	fmt.Println("2. Data Pemilih")
	fmt.Println("3. Atur Tanggal Pemilu")
	fmt.Println("4. Hasil Pemilu")
	fmt.Println("5. Keluar")

	fmt.Println()
	fmt.Print("Pilihan anda : ")
	fmt.Scan(&pilihan)

	if pilihan == 1 /* data calon*/ {
		dataCALON(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)

	} else if pilihan == 2 /*data pemilih*/ {
		dataPEMILIH(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)

	} else if pilihan == 3 /*Atur tanggal pemilu*/ {
		aturTanggal(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
	} else if pilihan == 4 /*Hasil Pemilu*/ {
		hasilPEMILU(*TA, *TB, *TC, *nA, *nB, *nC, *nPA, *nPB, *nPC)
		menuKPU(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
	} else if pilihan == 5 /*Keluar*/ {
		menuUtama(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
	}

}

func aturTanggal(A *date, TA *arrProvA, TB *arrProvB, TC *arrProvC, nA, nB, nC, nPA, nPB, nPC *int) {

	var konfirmasi string

	konfirmasi = "n"
	for konfirmasi != "y" {
		tanggalMulai(A)
		tanggalSelesai(A)
		fmt.Println("Pemilu berlangsung selama : ", selisihTanggal(*A), "hari")
		fmt.Println("Tanggal pemilu dimulai    : ", A.dd1, "-", A.mm1, "-", A.yy1)
		fmt.Println("Tanggal pemilu selesai    : ", A.dd2, "-", A.mm2, "-", A.yy2)
		fmt.Print("Konfirmasi (y/n) : ")
		fmt.Scan(&konfirmasi)

		fmt.Println()
		if konfirmasi == "y" {
			fmt.Println("------------------------------------")
			fmt.Println("* *Tanggal pemilu berhasil diatur* *")
			fmt.Println("------------------------------------")
			fmt.Println()
			menuKPU(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
		}

	}

}

func tanggalMulai(A *date) {
	fmt.Println("Masukkan tanggal pemilu dimulai : (dd/mm/yyyy)")
	fmt.Scan(&A.dd1, &A.mm1, &A.yy1)
}

func tanggalSelesai(A *date) {
	fmt.Println("Masukkan tanggal pemilu selesai : (dd/mm/yyyy)")
	fmt.Scan(&A.dd2, &A.mm2, &A.yy2)
}

func selisihTanggal(A date) int {
	var selisih int
	if A.mm1 <= A.mm2 {
		selisih = A.dd2 - A.dd1
	} else {
		selisih = A.dd2 - A.dd1
		selisih += 30
	}
	return selisih
}

func cekTanggal(A date, tanggal, bulan int) bool {
	var bener bool
	if tanggal >= A.dd1 && tanggal <= A.dd2 {
		bener = true
	} else {
		bener = false
	}
	return bener

}

func isiCalon(TA *arrProvA, TB *arrProvB, TC *arrProvC, nA, nB, nC *int) {
	var namaA, partaiA, namaB, partaiB, namaC, partaiC string

	fmt.Println("Masukkan data calon Provinsi A: ")
	fmt.Scan(&namaA)
	for namaA != "STOP" {
		fmt.Scan(&partaiA)
		TA[*nA].calon.nama = namaA
		TA[*nA].calon.partai = partaiA
		*nA++
		fmt.Scan(&namaA)
	}

	fmt.Println("Masukkan data calon Provinsi B: ")
	fmt.Scan(&namaB)
	for namaB != "STOP" {
		fmt.Scan(&partaiB)
		TB[*nB].calon.nama = namaB
		TB[*nB].calon.partai = partaiB
		*nB++
		fmt.Scan(&namaB)
	}

	fmt.Println("Masukkan data calon Provinsi C: ")
	fmt.Scan(&namaC)
	for namaC != "STOP" {
		fmt.Scan(&partaiC)
		TC[*nC].calon.nama = namaC
		TC[*nC].calon.partai = partaiC
		*nC++
		fmt.Scan(&namaC)
	}

}

func editCalonA(TA *arrProvA, noUrut, nA int) {
	var nama, partai, konfirmasi string
	fmt.Println("Masukkan data yang di edit: ")
	fmt.Scan(&nama, &partai)
	TA[noUrut-1].calon.nama = nama
	TA[noUrut-1].calon.partai = partai
	fmt.Println("======== Data yang di edit =========")
	fmt.Println(TA[noUrut-1].calon.nama, TA[noUrut-1].calon.partai)
	fmt.Print("Konfirmasi? (y/n) ")
	fmt.Scan(&konfirmasi)
	if konfirmasi == "y" {
		fmt.Println("Data berhasil di edit!")
		fmt.Println("Data Calon Provinsi A : ")
		cetakCalonA(*TA, nA)
	} else {
		editCalonA(TA, noUrut, nA)
	}
}

func editCalonB(TB *arrProvB, noUrut, nB int) {
	var nama, partai, konfirmasi string
	fmt.Println("Masukkan data yang di edit: ")
	fmt.Scan(&nama, &partai)
	TB[noUrut-1].calon.nama = nama
	TB[noUrut-1].calon.partai = partai
	fmt.Println("======== Data yang di edit =========")
	fmt.Println(TB[noUrut-1].calon.nama, TB[noUrut-1].calon.partai)
	fmt.Print("Konfirmasi? (y/n) ")
	fmt.Scan(&konfirmasi)
	if konfirmasi == "y" {
		fmt.Println("Data berhasil di edit!")
		fmt.Println("Data Calon Provinsi B : ")
		cetakCalonB(*TB, nB)
	} else {
		editCalonB(TB, noUrut, nB)
	}
}

func editCalonC(TC *arrProvC, noUrut, nC int) {
	var nama, partai, konfirmasi string
	fmt.Println("Masukkan data yang di edit: ")
	fmt.Scan(&nama, &partai)
	TC[noUrut-1].calon.nama = nama
	TC[noUrut-1].calon.partai = partai
	fmt.Println("======== Data yang di edit =========")
	fmt.Println(TC[noUrut-1].calon.nama, TC[noUrut-1].calon.partai)
	fmt.Print("Konfirmasi? (y/n) ")
	fmt.Scan(&konfirmasi)
	if konfirmasi == "y" {
		fmt.Println("Data berhasil di edit!")
		fmt.Println("Data Calon Provinsi C : ")
		cetakCalonC(*TC, nC)
	} else {
		editCalonC(TC, noUrut, nC)
	}
}

func hapusCalonA(TA *arrProvA, noUrut int, nA *int) {
	var konfirmasi string
	fmt.Println("Apakah anda yakin akan menghapus calon nomor", noUrut, TA[noUrut-1].calon.nama, TA[noUrut-1].calon.partai, " ?")
	fmt.Print("Konfirmasi? (y/n)")
	fmt.Scan(&konfirmasi)
	if konfirmasi == "y" {
		for i := noUrut - 1; i < *nA-(noUrut-1); i++ {
			TA[i].calon = TA[i+1].calon
		}
		*nA = *nA - 1
		fmt.Println("Data berhasil dihapus!")
		cetakCalonA(*TA, *nA)
	} else {
		fmt.Println("Data tidak dihapus")
	}

}

func hapusCalonB(TB *arrProvB, noUrut int, nB *int) {
	var konfirmasi string
	fmt.Println("Apakah anda yakin akan menghapus calon nomor", noUrut, TB[noUrut-1].calon.nama, TB[noUrut-1].calon.partai, " ?")
	fmt.Print("Konfirmasi? (y/n)")
	fmt.Scan(&konfirmasi)
	if konfirmasi == "y" {
		for i := noUrut - 1; i < *nB-(noUrut-1); i++ {
			TB[i].calon = TB[i+1].calon
		}
		*nB = *nB - 1
		fmt.Println("Data berhasil dihapus!")
		cetakCalonB(*TB, *nB)
	} else {
		fmt.Println("Data tidak dihapus")
	}
}

func hapusCalonC(TC *arrProvC, noUrut int, nC *int) {
	var konfirmasi string
	fmt.Println("Apakah anda yakin akan menghapus calon nomor", noUrut, TC[noUrut-1].calon.nama, TC[noUrut-1].calon.partai, " ?")
	fmt.Print("Konfirmasi? (y/n)")
	fmt.Scan(&konfirmasi)
	if konfirmasi == "y" {
		for i := noUrut - 1; i < *nC-(noUrut-1); i++ {
			TC[i].calon = TC[i+1].calon
		}
		*nC = *nC - 1
		fmt.Println("Data berhasil dihapus!")
		cetakCalonC(*TC, *nC)
	} else {
		fmt.Println("Data tidak dihapus")
	}
}
func cetakCalonA(TA arrProvA, nA int) {
	fmt.Println("--------------")
	fmt.Println("Provinsi A : ")
	fmt.Println("--------------")
	for i := 0; i < nA; i++ {
		fmt.Printf("%d. %s %s \n", i+1, TA[i].calon.nama, TA[i].calon.partai)
	}
	fmt.Println()
}

func cetakCalonB(TB arrProvB, nB int) {
	fmt.Println("--------------")
	fmt.Println("Provinsi B : ")
	fmt.Println("--------------")
	for j := 0; j < nB; j++ {
		fmt.Printf("%d. %s %s \n", j+1, TB[j].calon.nama, TB[j].calon.partai)
	}
	fmt.Println()
}

func cetakCalonC(TC arrProvC, nC int) {
	fmt.Println("--------------")
	fmt.Println("Provinsi C : ")
	fmt.Println("--------------")
	for k := 0; k < nC; k++ {
		fmt.Printf("%d. %s %s \n", k+1, TC[k].calon.nama, TC[k].calon.partai)
	}
	fmt.Println()
}

func searchNamaCalonA(TA arrProvA, nA int) {
	var nama string
	var ketemu bool
	fmt.Print("Masukkan nama yang dicari: ")
	fmt.Scan(&nama)
	ketemu = false
	for i := 0; i < nA; i++ {
		if TA[i].calon.nama == nama {
			fmt.Println(i)
			fmt.Println("============================================")
			fmt.Println("              ~Data ditemukan!~             ")
			fmt.Println("--------------------------------------------")
			fmt.Printf("%d. %s %s\n", i+1, TA[i].calon.nama, TA[i].calon.partai)
			fmt.Println("============================================")
			ketemu = true
			i = nA
		}
	}
	if ketemu == false {
		fmt.Println("============================================")
		fmt.Println("            ~Data tidak ditemukan!~         ")
		fmt.Println("============================================")
	}

}

func searchNamaCalonB(TB arrProvB, nB int) {
	var nama string
	var ketemu bool
	fmt.Print("Masukkan nama yang dicari: ")
	fmt.Scan(&nama)
	ketemu = false
	for i := 0; i < nB; i++ {
		if TB[i].calon.nama == nama {
			fmt.Println(i)
			fmt.Println("============================================")
			fmt.Println("              ~Data ditemukan!~             ")
			fmt.Println("--------------------------------------------")
			fmt.Printf("%d. %s %s\n", i+1, TB[i].calon.nama, TB[i].calon.partai)
			fmt.Println("============================================")
			ketemu = true
			i = nB
		}
	}
	if ketemu == false {
		fmt.Println("============================================")
		fmt.Println("            ~Data tidak ditemukan!~         ")
		fmt.Println("============================================")
	}
}

func searchNamaCalonC(TC arrProvC, nC int) {
	var nama string
	var ketemu bool
	fmt.Print("Masukkan nama yang dicari: ")
	fmt.Scan(&nama)
	ketemu = false
	for i := 0; i < nC; i++ {
		if TC[i].calon.nama == nama {
			fmt.Println(i)
			fmt.Println("============================================")
			fmt.Println("              ~Data ditemukan!~             ")
			fmt.Println("--------------------------------------------")
			fmt.Printf("%d. %s %s\n", i+1, TC[i].calon.nama, TC[i].calon.partai)
			fmt.Println("============================================")
			ketemu = true
			i = nC
		}
	}
	if ketemu == false {
		fmt.Println("============================================")
		fmt.Println("            ~Data tidak ditemukan!~         ")
		fmt.Println("============================================")
	}
}

func searchPartaiCalonA(TA arrProvA, nA int) {
	var partai string
	var ketemu bool
	fmt.Print("Masukkan partai yang dicari: ")
	fmt.Scan(&partai)
	ketemu = false
	for i := 0; i < nA; i++ {
		if TA[i].calon.partai == partai {
			fmt.Printf("%d. %s %s\n", i+1, TA[i].calon.nama, TA[i].calon.partai)
			ketemu = true
		}
	}
	if ketemu == true {
		fmt.Println("============================================")
		fmt.Println("              ~Data ditemukan!~             ")
		fmt.Println("============================================")
	} else if ketemu == false {
		fmt.Println("============================================")
		fmt.Println("            ~Data tidak ditemukan!~         ")
		fmt.Println("============================================")
	}
}

func searchPartaiCalonB(TB arrProvB, nB int) {
	var partai string
	var ketemu bool
	fmt.Print("Masukkan partai yang dicari: ")
	fmt.Scan(&partai)
	ketemu = false
	for i := 0; i < nB; i++ {
		if TB[i].calon.partai == partai {
			fmt.Printf("%d. %s %s\n", i+1, TB[i].calon.nama, TB[i].calon.partai)
			ketemu = true
		}
	}
	if ketemu == true {
		fmt.Println("============================================")
		fmt.Println("              ~Data ditemukan!~             ")
		fmt.Println("============================================")
	} else if ketemu == false {
		fmt.Println("============================================")
		fmt.Println("            ~Data tidak ditemukan!~         ")
		fmt.Println("============================================")
	}
}

func searchPartaiCalonC(TC arrProvC, nC int) {
	var partai string
	var ketemu bool
	fmt.Print("Masukkan partai yang dicari: ")
	fmt.Scan(&partai)
	ketemu = false
	for i := 0; i < nC; i++ {
		if TC[i].calon.partai == partai {
			fmt.Printf("%d. %s %s\n", i+1, TC[i].calon.nama, TC[i].calon.partai)
			ketemu = true
		}
	}
	if ketemu == true {
		fmt.Println("============================================")
		fmt.Println("              ~Data ditemukan!~             ")
		fmt.Println("============================================")
	} else if ketemu == false {
		fmt.Println("============================================")
		fmt.Println("            ~Data tidak ditemukan!~         ")
		fmt.Println("============================================")
	}
}

func cetakPemilihA(TA arrProvA, nPA int) {
	var calon int
	fmt.Println("Provinsi A : ")
	if nPA > 0 {
		for i := 0; i < nPA; i++ {
			calon = TA[i].pemilih.pilihan
			fmt.Printf("%d. Nama: %s \n", i+1, TA[i].pemilih.nama)
			fmt.Printf("   Pilihan: %d. %s - %s \n", TA[i].pemilih.pilihan, TA[calon-1].calon.nama, TA[calon-1].calon.partai)
		}
		fmt.Println("")
	}
}

func cetakPemilihB(TB arrProvB, nPB int) {
	var calon int
	fmt.Println("Provinsi B : ")
	if nPB > 0 {
		for i := 0; i < nPB; i++ {
			calon = TB[i].pemilih.pilihan
			fmt.Printf("%d. Nama: %s \n", i+1, TB[i].pemilih.nama)
			fmt.Printf("   Pilihan: %d. %s - %s \n", TB[i].pemilih.pilihan, TB[calon-1].calon.nama, TB[calon-1].calon.partai)
		}
		fmt.Println("")
	}
}

func cetakPemilihC(TC arrProvC, nPC int) {
	var calon int
	fmt.Println("Provinsi C : ")
	if nPC > 0 {
		for i := 0; i < nPC; i++ {
			calon = TC[i].pemilih.pilihan
			fmt.Printf("%d. Nama: %s \n", i+1, TC[i].pemilih.nama)
			fmt.Printf("   Pilihan: %d. %s - %s \n", TC[i].pemilih.pilihan, TC[calon-1].calon.nama, TC[calon-1].calon.partai)
		}
		fmt.Println("")
	}
}

func searchPemilihA(TA arrProvA, nPA int) {
	var calon int
	var nama string
	var ketemu bool
	fmt.Print("Masukkan nama yang dicari: ")
	fmt.Scan(&nama)
	ketemu = false
	for i := 0; i < nPA; i++ {
		if TA[i].pemilih.nama == nama {
			fmt.Println(i)
			calon = TA[i].pemilih.pilihan
			fmt.Println("============================================")
			fmt.Println("              ~Data ditemukan!~             ")
			fmt.Println("--------------------------------------------")
			fmt.Printf("%d. Nama: %s \n", i+1, TA[i].pemilih.nama)
			fmt.Printf("   Pilihan: %d. %s - %s \n", TA[i].pemilih.pilihan, TA[calon-1].calon.nama, TA[calon-1].calon.partai)
			fmt.Println("============================================")
			ketemu = true
			i = nPA
		}
	}
	if ketemu == false {
		fmt.Println("============================================")
		fmt.Println("            ~Data tidak ditemukan!~         ")
		fmt.Println("============================================")
	}
}

func searchPemilihB(TB arrProvB, nPB int) {
	var calon int
	var nama string
	var ketemu bool
	fmt.Print("Masukkan nama yang dicari: ")
	fmt.Scan(&nama)
	ketemu = false
	for i := 0; i < nPB; i++ {
		if TB[i].pemilih.nama == nama {
			calon = TB[i].pemilih.pilihan
			fmt.Println("============================================")
			fmt.Println("              ~Data ditemukan!~             ")
			fmt.Println("--------------------------------------------")
			fmt.Printf("%d. Nama: %s \n", i+1, TB[i].pemilih.nama)
			fmt.Printf("   Pilihan: %d. %s - %s \n", TB[i].pemilih.pilihan, TB[calon-1].calon.nama, TB[calon-1].calon.partai)
			fmt.Println("============================================")
			ketemu = true
			i = nPB
		}

	}
	if !ketemu {
		fmt.Println("============================================")
		fmt.Println("            ~Data tidak ditemukan!~         ")
		fmt.Println("============================================")
	}
}

func searchPemilihC(TC arrProvC, nPC int) {
	var calon int
	var nama string
	var ketemu bool
	fmt.Print("Masukkan nama yang dicari: ")
	fmt.Scan(&nama)
	ketemu = false
	for i := 0; i < nPC; i++ {
		if TC[i].pemilih.nama == nama {
			calon = TC[i].pemilih.pilihan
			fmt.Println("============================================")
			fmt.Println("              ~Data ditemukan!~             ")
			fmt.Println("--------------------------------------------")
			fmt.Printf("%d. Nama: %s \n", i+1, TC[i].pemilih.nama)
			fmt.Printf("   Pilihan: %d. %s - %s \n", TC[i].pemilih.pilihan, TC[calon-1].calon.nama, TC[calon-1].calon.partai)
			fmt.Println("============================================")
			ketemu = true
			i = nPC
		}

	}
	if !ketemu {
		fmt.Println("============================================")
		fmt.Println("            ~Data tidak ditemukan!~         ")
		fmt.Println("============================================")
	}
}

func hitungSemuaCALON(TA arrProvA, TB arrProvB, TC arrProvC, nA, nB, nC, nPA, nPB, nPC int) {

	//hitungA
	for i := 0; i < nA; i++ {
		for j := 0; j < nPA; j++ {
			if TA[j].pemilih.pilihan == i+1 {
				TA[i].totalSuaraCalon++
			}
		}
	}

	//hitungB
	for i := 0; i < nB; i++ {
		for j := 0; j < nPB; j++ {
			if TB[j].pemilih.pilihan == i+1 {
				TB[i].totalSuaraCalon++
			}
		}
	}

	//hitungC
	for i := 0; i < nC; i++ {
		for j := 0; j < nPC; j++ {
			if TC[j].pemilih.pilihan == i+1 {
				TC[i].totalSuaraCalon++
			}
		}
	}

	sortingTOTAL(TA, TB, TC, nPA, nPB, nPC)

	//hitungPartai A

	fmt.Println("----------------------------------")
	fmt.Println("Total suara PARTAI di Provinsi A: ")

	for i := 0; i < nA; i++ {
		for j := i + 1; j < nA; j++ {
			if TA[i].calon.partai == TA[j].calon.partai {
				TA[i].totalSuaraPartai = TA[i].totalSuaraCalon + TA[j].totalSuaraCalon
				//hapus data indeks j
				for k := j; k < nA; k++ {
					TA[k] = TA[k+1]
				}
			} else {
				TA[i].totalSuaraPartai = TA[i].totalSuaraCalon
			}
		}
		if TA[i].totalSuaraPartai > 0 {
			fmt.Println(TA[i].calon.partai, " - ", TA[i].totalSuaraPartai, "suara")
		}
	}

	//hitungPartai B
	fmt.Println()
	fmt.Println("----------------------------------")
	fmt.Println("Total suara PARTAI di Provinsi B: ")

	for i := 0; i < nB; i++ {
		for j := 1; j < nB; j++ {
			if TB[i].calon.partai == TB[j].calon.partai {
				TB[i].totalSuaraPartai = TB[i].totalSuaraCalon + TB[j].totalSuaraCalon
				//hapus data indeks j
				for k := j; k < nB; k++ {
					TB[k] = TB[k+1]
				}
			} else {
				TB[i].totalSuaraPartai = TB[i].totalSuaraCalon
			}

		}
		if TB[i].totalSuaraPartai > 0 {

			fmt.Println(TB[i].calon.partai, " - ", TB[i].totalSuaraPartai, "suara")
		}
	}

	fmt.Println()
	fmt.Println("----------------------------------")
	fmt.Println("Total suara PARTAI di Provinsi C: ")

	//hitungPartai C
	for i := 0; i < nC; i++ {
		for j := 1; j < nC; j++ {
			if TC[i].calon.partai == TC[j].calon.partai {
				TC[i].totalSuaraPartai = TC[i].totalSuaraCalon + TC[j].totalSuaraCalon
				//hapus data indeks j
				for k := j; k < nC; k++ {
					TC[k] = TC[k+1]
				}
			} else {
				TC[i].totalSuaraPartai = TC[i].totalSuaraCalon
			}

		}
		if TC[i].totalSuaraPartai > 0 {

			fmt.Println(TC[i].calon.partai, " - ", TC[i].totalSuaraPartai, "suara")
		}
	}

}

func sortingTOTAL(TA arrProvA, TB arrProvB, TC arrProvC, nPA, nPB, nPC int) {
	var tempA, tempB, tempC provinsi

	for i := 0; i < nPA; i++ {
		for j := i + 1; j < nPA; j++ {
			if TA[i].totalSuaraCalon < TA[j].totalSuaraCalon {
				tempA = TA[i]
				TA[i] = TA[j]
				TA[j] = tempA
			}
		}
	}

	for i := 0; i < nPB; i++ {
		for j := i + 1; j < nPB; j++ {
			if TB[i].totalSuaraCalon < TB[j].totalSuaraCalon {
				tempB = TB[i]
				TB[i] = TB[j]
				TB[j] = tempB
			}
		}
	}

	for i := 0; i < nPC; i++ {
		for j := i + 1; j < nPC; j++ {
			if TC[i].totalSuaraCalon < TC[j].totalSuaraCalon {
				tempC = TC[i]
				TC[i] = TC[j]
				TC[j] = tempC
			}
		}
	}

	//provA
	fmt.Println()
	fmt.Println("===========================================")
	fmt.Println("Hasil PEMILU 3 Suara terbanyak Provinsi A: ")
	for i := 0; i < 3; i++ {
		fmt.Println(TA[i].calon.nama, " - ", TA[i].calon.partai, " - ", TA[i].totalSuaraCalon, "Suara ")
	}
	fmt.Println("===========================================")

	//provB
	fmt.Println()
	fmt.Println("===========================================")
	fmt.Println("Hasil PEMILU 3 Suara terbanyak Provinsi B: ")
	for i := 0; i < 3; i++ {
		fmt.Println(TB[i].calon.nama, " - ", TB[i].calon.partai, " - ", TB[i].totalSuaraCalon, "Suara ")
	}
	fmt.Println("===========================================")
	//provC
	fmt.Println()
	fmt.Println("===========================================")
	fmt.Println("Hasil PEMILU 3 Suara terbanyak Provinsi C: ")
	for i := 0; i < 3; i++ {
		fmt.Println(TC[i].calon.nama, " - ", TC[i].calon.partai, " - ", TC[i].totalSuaraCalon, "Suara ")
	}
	fmt.Println("===========================================")
	fmt.Println()
}

func hasilPEMILU(TA arrProvA, TB arrProvB, TC arrProvC, nA, nB, nC, nPA, nPB, nPC int) {
	fmt.Println("============================================")
	fmt.Println("              ~HASIL PEMILU~                ")
	fmt.Println("--------------------------------------------")

	hitungSemuaCALON(TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)

}
