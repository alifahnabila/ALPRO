package main

import "fmt"

const NMAX int = 10

type dataUtama struct {
	nama       string
	nominal    float64
	jatuhTempo int
	kategori   string
	status     string
}

type tabDataTagihan [NMAX]dataUtama

func lihatTagihan(A tabDataTagihan, n int) {
	if n == 0 {
		fmt.Println("Data tagihan masih kosong!\n")
	} else {
		fmt.Println("No | Nama Tagihan         | Nominal        | Jatuh Tempo | Kategori     | Status")
		fmt.Println("---+----------------------+----------------+-------------+--------------+-----------")
		for i := 0; i < n; i++ {
			tgl := A[i].jatuhTempo % 100
			bln := A[i].jatuhTempo / 100 % 100
			thn := A[i].jatuhTempo / 10000

			fmt.Printf("%-3d| %-20s | Rp%-12.0f | %02d/%02d/%04d  | %-12s | %s\n", i+1, A[i].nama, A[i].nominal, tgl, bln, thn, A[i].kategori, A[i].status)
			fmt.Println()
		}
	}
}

func tambahTagihan(A *tabDataTagihan, n *int) {
	var tgl, bln, thn, jumlah int

	fmt.Printf("Sisa kapasitas penyimpanan: %d\n", NMAX-*n)

	fmt.Print("Berapa jumlah tagihan yang ingin ditambahkan? ")
	fmt.Scan(&jumlah)

	if jumlah > NMAX-*n {
		jumlah = NMAX - *n
		fmt.Println("Data yang bisa ditambahkan berjumlah: ", jumlah)
	}

	if jumlah <= 0 {
		fmt.Println("Jumlah tidak valid!\n")
	} else {
		for i := 0; i < jumlah; i++ {
			fmt.Printf("\n===== Tambah Tagihan ke-%d =====\n", i+1)
			fmt.Print("Nama Tagihan: ")
			fmt.Scan(&A[*n].nama)

			fmt.Print("Nominal: ")
			fmt.Scan(&A[*n].nominal)

			fmt.Println("Masukkan Jatuh Tempo: ")
			fmt.Print("Tanggal (1-31): ")
			fmt.Scan(&tgl)
			fmt.Print("Bulan (1-12): ")
			fmt.Scan(&bln)
			fmt.Print("Tahun (yyyy): ")
			fmt.Scan(&thn)

			fmt.Print("Kategori: ")
			fmt.Scan(&A[*n].kategori)

			fmt.Print("Status (Lunas/Belum): ")
			fmt.Scan(&A[*n].status)

			if tgl < 1 || tgl > 31 || bln < 1 || bln > 12 || thn < 0000 || thn > 3000 {
				fmt.Println("Jatuh tempo tidak valid! Data Tagihan tidak akan dimasukkan!")
			} else {
				A[*n].jatuhTempo = thn*10000 + bln*100 + tgl
				*n++
			}

		}
		fmt.Printf("\n%d Data berhasil ditambahkan!\n", jumlah)
		fmt.Println()
	}
}

func ubahTagihan(A *tabDataTagihan, n int) {
	var cariNama string
	var pilihUbah int
	if n == 0 {
		fmt.Println("Data tagihan masih kosong!\n")
	} else {
		fmt.Print("\nMasukkan nama tagihan yang ingin diubah: ")
		fmt.Scan(&cariNama)

		posisi := -1
		i := 0
		for posisi == -1 && i < n {
			if A[i].nama == cariNama {
				posisi = i
			}
			i++
		}

		if posisi == -1 {
			fmt.Printf("Tagihan dengan nama %s tidak ditemukan!\n", cariNama)
		} else {
			fmt.Println("═══ DATA TAGIHAN YANG AKAN DIUBAH ═══")
			fmt.Printf("[1] Nama            : %s \n", A[posisi].nama)
			fmt.Printf("[2] Nomina          : Rp%2.f\n", A[posisi].nominal)
			tgl := A[posisi].jatuhTempo % 100
			bln := A[posisi].jatuhTempo / 100 % 100
			thn := A[posisi].jatuhTempo / 10000
			fmt.Printf("[3] Jatuh Tempo     : %2d/%2d/%4d\n", tgl, bln, thn)
			fmt.Printf("[4] Kategori        : %s\n", A[posisi].kategori)
			fmt.Printf("[5] Status          : %s\n", A[posisi].status)

			fmt.Println("PILIH DATA TAGIHAN YANG INGIN DIUBAH:")
			fmt.Println("[1] Nama")
			fmt.Println("[2] Nominal")
			fmt.Println("[3] Jatuh Tempo")
			fmt.Println("[4] Kategori")
			fmt.Println("[5] Status")
			fmt.Println("[6] Ubah Semua")
			fmt.Println("[0] Batal")
			fmt.Print("Masukkan pilihan: ")
			fmt.Scan(&pilihUbah)

			if pilihUbah == 0 {
				fmt.Println("Perubahan dibatalkan!")
			} else if pilihUbah == 1 {
				var namaBaru string
				fmt.Print("Masukkan nama baru: ")
				fmt.Scan(&namaBaru)
				A[posisi].nama = namaBaru
				fmt.Println("Nama berhasil diubah!")
			} else if pilihUbah == 2 {
				var nominalBaru float64
				fmt.Print("Masukkan nominal baru: Rp")
				fmt.Scan(&nominalBaru)
				A[posisi].nominal = nominalBaru
				fmt.Println("Nominal berhasil diubah!")
			} else if pilihUbah == 3 {
				var tglBaru, blnBaru, thnBaru int
				fmt.Print("Masukkan tanggal baru: ")
				fmt.Scan(&tglBaru)
				fmt.Print("Masukkan bulan baru: ")
				fmt.Scan(&blnBaru)
				fmt.Print("Masukkan tahun baru: ")
				fmt.Scan(&thnBaru)
				if tglBaru < 1 || tglBaru > 31 || blnBaru < 1 || blnBaru > 12 || thnBaru < 0000 || thnBaru > 3000 {
					fmt.Println("Angka tidak valid! Perubahan dibatalkan!")
				} else {
					A[posisi].jatuhTempo = thnBaru*10000 + blnBaru*100 + tglBaru
					fmt.Println("Jatuh tempo berhasil diubah!")
				}
			} else if pilihUbah == 4 {
				var kategoriBaru string
				fmt.Print("Masukkan kategori baru: ")
				fmt.Scan(&kategoriBaru)
				A[posisi].kategori = kategoriBaru
				fmt.Println("Kategori berhasil diubah!")
			} else if pilihUbah == 5 {
				var statusBaru string
				fmt.Print("Masukkan status baru (lunas/belum): ")
				fmt.Scan(&statusBaru)
				if statusBaru == "lunas" || statusBaru == "Lunas" || statusBaru == "belum" || statusBaru == "Belum" {
					A[posisi].status = statusBaru
					fmt.Println("Status berhasil diubah!")
				} else {
					fmt.Println("Status tidak valid! Perubahan dibatalkan!")
				}
			} else if pilihUbah == 6 {
				var namaBaru, kategoriBaru, statusBaru string
				var nominalBaru float64
				var tglBaru, blnBaru, thnBaru int
				fmt.Print("Masukkan nama baru: ")
				fmt.Scan(&namaBaru)
				A[posisi].nama = namaBaru
				fmt.Print("Masukkan nominal baru: ")
				fmt.Scan(&nominalBaru)
				A[posisi].nominal = nominalBaru

				fmt.Print("Masukkan tanggal baru: ")
				fmt.Scan(&tglBaru)
				fmt.Print("Masukkan bulan baru: ")
				fmt.Scan(&blnBaru)
				fmt.Print("Masukkan tahun baru: ")
				fmt.Scan(&thnBaru)
				if tglBaru >= 1 || tglBaru <= 31 || blnBaru >= 1 || blnBaru <= 12 || thnBaru >= 0000 || thnBaru <= 3000 {
					A[posisi].jatuhTempo = thnBaru*10000 + blnBaru*100 + tglBaru
				}

				fmt.Print("Masukkan kategori baru: ")
				fmt.Scan(&kategoriBaru)
				A[posisi].kategori = kategoriBaru
				fmt.Print("Masukkan status baru: ")
				fmt.Scan(&statusBaru)
				if statusBaru == "lunas" || statusBaru == "Lunas" || statusBaru == "belum" || statusBaru == "Belum" {
					A[posisi].status = statusBaru
				}
				fmt.Println("Semua data berhasil diubah!")
			} else {
				fmt.Println("Pilihan tidak valid!")
			}

			if pilihUbah != 0 {
				fmt.Println("\n═══ DATA TAGIHAN YANG AKAN DIUBAH ═══")
				lihatTagihan(*A, n)
			}
		}
	}
}

func hapusTagihan(A *tabDataTagihan, n *int) {
	var konfirmasi string
	var approved bool
	var no, idx int

	approved = true

	if *n == 0 {
		fmt.Println("Data tagihan masih kosong!\n")
	} else {
		for approved {
			lihatTagihan(*A, *n)
			fmt.Print("\nMasukkan nomor tagihan yang akan dihapus: ")
			fmt.Scan(&no)

			if no == 0 {
				approved = false
				fmt.Println("Penghapusan selesai!")
			} else if no >= 1 && no <= *n {
				idx = no - 1

				fmt.Println("===== KONFIRMASI HAPUS =====")
				fmt.Printf("Hapus tagihan [%s | %2.f | %d | %s | %s]? (y/n): ", A[idx].nama, A[idx].nominal, A[idx].jatuhTempo, A[idx].kategori, A[idx].status)
				fmt.Scan(&konfirmasi)

				if konfirmasi == "y" || konfirmasi == "Y" {
					for i := idx; i < *n-1; i++ {
						A[i] = A[i+1]
					}
					*n--
					fmt.Println("Tagihan berhasil dihapus!")
				} else {
					fmt.Println("Penghapusan tagihan dibatalkan!")
				}
			} else {
				fmt.Println("Nomor tagihan tidak valid!")
			}
		}
	}
}

func sequentialSearch(A tabDataTagihan, n int) {
	var cari, cariNama, cariKategori string
	if n == 0 {
		fmt.Println("Data tagihan masih kosong!\n")
	} else {
		fmt.Print("Data Tagihan yang ingin dicari (Nama/Kategori): ")
		fmt.Scan(&cari)

		if cari == "Nama" || cari == "nama" {
			fmt.Print("Nama Tagihan yang ingin dicari: ")
			fmt.Scan(&cariNama)

			fmt.Println("\nHASIL PENCARIAN NAMA: ")
			fmt.Println("No | Nama Tagihan         | Nominal        | Jatuh Tempo | Kategori     | Status")
			fmt.Println("---+----------------------+----------------+-------------+--------------+-----------")

			found := false
			no := 1
			for i := 0; i < n; i++ {
				if A[i].nama == cariNama {
					tgl := A[i].jatuhTempo % 100
					bln := A[i].jatuhTempo / 100 % 100
					thn := A[i].jatuhTempo / 10000
					fmt.Printf("%-3d| %-20s | Rp%-12.0f | %02d/%02d/%04d  | %-12s | %s\n", no, A[i].nama, A[i].nominal, tgl, bln, thn, A[i].kategori, A[i].status)
					fmt.Println()
					no++
					found = true
				}
			}
			if !found {
				fmt.Printf("Tagihan dengan nama [%s] tidak ditemukan!\n", cariNama)
			} else {
				fmt.Printf("Ditemukan sebanyak %d tagihan dengan nama [%s]\n", no-1, cariNama)
			}
			fmt.Println()
		} else if cari == "kategori" || cari == "Kategori" {
			fmt.Print("Kategori Tagihan yang ingin dicari: ")
			fmt.Scan(&cariKategori)

			fmt.Println("\nHASIL PENCARIAN KATEGORI: ")
			fmt.Println("No | Nama Tagihan         | Nominal        | Jatuh Tempo | Kategori     | Status")
			fmt.Println("---+----------------------+----------------+-------------+--------------+-----------")

			found := false
			no := 1
			for i := 0; i < n; i++ {
				if A[i].kategori == cariKategori {
					tgl := A[i].jatuhTempo % 100
					bln := A[i].jatuhTempo / 100 % 100
					thn := A[i].jatuhTempo / 10000
					fmt.Printf("%-3d| %-20s | Rp%-12.0f | %02d/%02d/%04d  | %-12s | %s\n", no, A[i].nama, A[i].nominal, tgl, bln, thn, A[i].kategori, A[i].status)
					fmt.Println()
					no++
					found = true
				}
			}
			if !found {
				fmt.Printf("Tagihan dengan kategori [%s] tidak ditemukan!\n", cariKategori)
			} else {
				fmt.Printf("Ditemukan sebanyak %d tagihan dengan kategori [%s]\n", no-1, cariKategori)
			}
			fmt.Println()
		}
	}
}

func binarySearch(A *tabDataTagihan, n int) {
	var cari, cariNama, cariKategori, pilih string
	var left, right, mid int
	if n == 0 {
		fmt.Println("Data tagihan masih kosong!\n")
	} else {
		fmt.Print("Data Tagihan yang ingin dicari (Nama/Kategori): ")
		fmt.Scan(&cari)

		if cari == "Nama" || cari == "nama" {
			fmt.Print("Nama Tagihan yang ingin dicari: ")
			fmt.Scan(&cariNama)

			terurutNama := true
			for i := 0; i < n-1; i++ {
				if A[i].nama > A[i+1].nama {
					terurutNama = false
				}
			}

			if !terurutNama {
				fmt.Println("\nData Belum terurut!")
				fmt.Print("Apakah anda ingin mengurutkan data? (y/n): ")
				fmt.Scan(&pilih)

				if pilih == "y" || pilih == "Y" {
					for i := 0; i < n-1; i++ {
						idx := i
						for j := i + 1; j < n; j++ {
							if A[j].nama < A[idx].nama {
								idx = j
							}
						}
						temp := A[i]
						A[i] = A[idx]
						A[idx] = temp
					}
					fmt.Println("Data berhasil diurutkan berdasarkan nama!")
				} else {
					fmt.Println("Pengurutan data dibatalkan!\n")
				}

			} else {
				fmt.Println("Data sudah terurut!")
			}

			fmt.Println("\nHASIL PENCARIAN NAMA: ")
			fmt.Println("No | Nama Tagihan         | Nominal        | Jatuh Tempo | Kategori     | Status")
			fmt.Println("---+----------------------+----------------+-------------+--------------+-----------")

			left = 0
			right = n - 1
			found := false
			posisi := -1

			for left <= right && !found {
				mid = (left + right) / 2

				if A[mid].nama == cariNama {
					found = true
					posisi = mid
				} else if cariNama < A[mid].nama {
					right = mid - 1
				} else {
					left = mid + 1
				}
			}

			if found {
				start := posisi
				for start > 0 && A[start-1].nama == cariNama {
					start--
				}

				end := posisi
				for end < n-1 && A[end+1].nama == cariNama {
					end++
				}
				no := 1
				for i := start; i <= end; i++ {
					tgl := A[i].jatuhTempo % 100
					bln := A[i].jatuhTempo / 100 % 100
					thn := A[i].jatuhTempo / 10000

					fmt.Printf("%-3d| %-20s | Rp%-12.0f | %02d/%02d/%04d  | %-12s | %s\n", no, A[i].nama, A[i].nominal, tgl, bln, thn, A[i].kategori, A[i].status)
					no++
				}
				fmt.Printf("\nDitemukan %d tagihan dengan nama [%s]\n", no-1, cariNama)
			} else {
				fmt.Printf("Tagihan dengan nama [%s] tidak ditemukan!\n", cariNama)
			}
			fmt.Println()

		} else if cari == "kategori" || cari == "Kategori" {
			fmt.Print("Kategori Tagihan yang ingin dicari: ")
			fmt.Scan(&cariKategori)

			terurutKategori := true
			for i := 0; i < n-1; i++ {
				if A[i].kategori > A[i+1].kategori {
					terurutKategori = false
				}
			}

			if !terurutKategori {
				fmt.Println("\nData Belum terurut!")
				fmt.Print("Apakah anda ingin mengurutkan data? (y/n): ")
				fmt.Scan(&pilih)

				if pilih == "y" || pilih == "Y" {
					for i := 0; i < n-1; i++ {
						idx := i
						for j := i + 1; j < n; j++ {
							if A[j].kategori < A[idx].kategori {
								idx = j
							}
						}
						temp := A[i]
						A[i] = A[idx]
						A[idx] = temp
					}
					fmt.Println("Data berhasil diurutkan berdasarkan kategori!")
				} else {
					fmt.Println("Pengurutan data dibatalkan!\n")
				}

			} else {
				fmt.Println("Data sudah terurut!")
			}

			fmt.Println("\nHASIL PENCARIAN KATEGORI: ")
			fmt.Println("No | Nama Tagihan         | Nominal        | Jatuh Tempo | Kategori     | Status")
			fmt.Println("---+----------------------+----------------+-------------+--------------+-----------")

			left = 0
			right = n - 1
			found := false
			posisi := -1

			for left <= right && !found {
				mid = (left + right) / 2

				if A[mid].kategori == cariKategori {
					found = true
					posisi = mid
				} else if cariKategori < A[mid].kategori {
					right = mid - 1
				} else {
					left = mid + 1
				}
			}

			if found {
				start := posisi
				for start > 0 && A[start-1].kategori == cariKategori {
					start--
				}

				end := posisi
				for end < n-1 && A[end+1].kategori == cariKategori {
					end++
				}

				no := 1
				for i := start; i <= end; i++ {
					tgl := A[i].jatuhTempo % 100
					bln := A[i].jatuhTempo / 100 % 100
					thn := A[i].jatuhTempo / 10000

					fmt.Printf("%-3d| %-20s | Rp%-12.0f | %02d/%02d/%04d  | %-12s | %s\n", no, A[i].nama, A[i].nominal, tgl, bln, thn, A[i].kategori, A[i].status)
					no++
				}
				fmt.Printf("\nDitemukan %d tagihan dengan kategori [%s]\n", no-1, cariKategori)
			} else {
				fmt.Printf("Tagihan dengan kategori [%s] tidak ditemukan!\n", cariKategori)
			}
			fmt.Println()

		}
	}
}

func selectionSort(A *tabDataTagihan, n int) {
	if n == 0 {
		fmt.Println("Data tagihan masih kosong!\n")
	} else {
		for pass := 1; pass < n; pass++ {
			idx := pass - 1
			for i := pass; i < n; i++ {
				if A[idx].jatuhTempo > A[i].jatuhTempo {
					idx = i
				}
			}
			temp := A[pass-1]
			A[pass-1] = A[idx]
			A[idx] = temp
		}

		fmt.Println("Data yang berhasil diurutkan berdasarkan jatuh tempo terdekat: ")
		lihatTagihan(*A, n)

		fmt.Println("\nTAGIHAN BELUM LUNAS TERDEKAT: ")
		found := false
		no := 1
		for i := 0; i < n; i++ {
			if A[i].status == "belum" || A[i].status == "Belum" {
				tgl := A[i].jatuhTempo % 100
				bln := A[i].jatuhTempo / 100 % 100
				thn := A[i].jatuhTempo / 10000
				fmt.Printf("%-3d| %-20s | Rp%-12.0f | %02d/%02d/%04d  | %-12s | %s\n", no, A[i].nama, A[i].nominal, tgl, bln, thn, A[i].kategori, A[i].status)
				fmt.Println()
				no++
				found = true
			}
		}
		if !found {
			fmt.Println("Semua tagihan sudah lunas!")
		}

	}
}

func insertionSort(A *tabDataTagihan, n int) {
	if n == 0 {
		fmt.Println("Data tagihan masih kosong!\n")
	} else {
		for pass := 1; pass < n; pass++ {
			temp := A[pass]
			j := pass - 1
			for j >= 0 && A[j].jatuhTempo > temp.jatuhTempo {
				A[j+1] = A[j]
				j--
			}
			A[j+1] = temp
		}
		fmt.Println("Data yang berhasil diurutkan berdasarkan jatuh tempo terdekat: ")
		lihatTagihan(*A, n)

		fmt.Println("\nTAGIHAN BELUM LUNAS TERDEKAT: ")
		found := false
		no := 1
		for i := 0; i < n; i++ {
			if A[i].status == "belum" || A[i].status == "Belum" {
				tgl := A[i].jatuhTempo % 100
				bln := A[i].jatuhTempo / 100 % 100
				thn := A[i].jatuhTempo / 10000
				fmt.Printf("%-3d| %-20s | Rp%-12.0f | %02d/%02d/%04d  | %-12s | %s\n", no, A[i].nama, A[i].nominal, tgl, bln, thn, A[i].kategori, A[i].status)
				fmt.Println()
				no++
				found = true
			}
		}
		if !found {
			fmt.Println("Semua tagihan sudah lunas!")
		}

	}
}

func lihatStatistik(A tabDataTagihan, n int) {
	var totBiaya, totLunas, totBelumLunas float64
	var persenLunas, persenBelumLunas float64
	var hitungLunas, hitungBelumLunas int
	if n == 0 {
		fmt.Println("Data tagihan masih kosong!\n")
	} else {
		for i := 0; i < n; i++ {
			totBiaya = A[i].nominal + totBiaya

			if A[i].status == "lunas" || A[i].status == "Lunas" {
				totLunas = A[i].nominal + totLunas
				hitungLunas++
			} else if A[i].status == "belum" || A[i].status == "Belum" {
				totBelumLunas = A[i].nominal + totBelumLunas
				hitungBelumLunas++
			}
		}
		if totBiaya > 0 {
			persenLunas = (totLunas / totBiaya) * 100
			persenBelumLunas = (totBelumLunas / totBiaya) * 100
		}
		fmt.Println()
		fmt.Println("  ╔══════════════════════════════════════════════════╗  ")
		fmt.Println("  ║                STATISTIK TAGIHAN                 ║  ")
		fmt.Println("  ╚══════════════════════════════════════════════════╝  ")

		fmt.Println("═══ Informasi Umum ═══")
		fmt.Printf("Total Tagihan                : %d tagihan\n", n)
		fmt.Printf("Total Tagihan Lunas          : %d tagihan\n", hitungLunas)
		fmt.Printf("Total Tagihan Belum Lunas    : %d tagihan\n", hitungBelumLunas)

		fmt.Println("═══ Informasi Biaya ═══")
		fmt.Printf("Total Seluruh Biaya          : Rp%2.f\n", totBiaya)
		fmt.Printf("Total Biaya Lunas            : Rp%2.f\n", totLunas)
		fmt.Printf("Total Biaya Belum Lunas      : Rp%2.f\n", totBelumLunas)

		fmt.Println("═══ Informasi Persentase ═══")
		fmt.Printf("Persentase Lunas             : %2.f persen\n", persenLunas)
		fmt.Printf("Persentase Belum Lunas       : %2.f persen\n", persenBelumLunas)
		fmt.Println()
	}
}

func menuUtama() {
	fmt.Println("═══ MENU UTAMA ═══")
	fmt.Println("[1] Lihat Semua Tagihan")
	fmt.Println("[2] Tambah Tagihan")
	fmt.Println("[3] Ubah Tagihan")
	fmt.Println("[4] Hapus Tagihan")
	fmt.Println("[5] Cari Tagihan (Sequential Search)")
	fmt.Println("[6] Cari Tagihan (Binary Search)")
	fmt.Println("[7] Urutkan berdasarkan tanggal jatuh tempo terdekat (Selection Sort)")
	fmt.Println("[8] Urutkan berdasarkan tanggal jatuh tempo terdekat (Insertion Sort)")
	fmt.Println("[9] Lihat Statistik")
	fmt.Println("[0] Keluar")
	fmt.Println()
}

func main() {
	fmt.Println()
	fmt.Println("  ╔══════════════════════════════════════════════════╗  ")
	fmt.Println("  ║    SIMTAB - Sistem Manajemen Tagihan Bulanan     ║  ")
	fmt.Println("  ╚══════════════════════════════════════════════════╝  ")
	fmt.Println()
	menuUtama()

	var nomor int
	fmt.Print("Pilih menu: ")
	fmt.Scan(&nomor)

	var A tabDataTagihan
	n := 0
	for nomor != 0 {
		if nomor == 1 {
			lihatTagihan(A, n)
		} else if nomor == 2 {
			tambahTagihan(&A, &n)
		} else if nomor == 3 {
			ubahTagihan(&A, n)
		} else if nomor == 4 {
			hapusTagihan(&A, &n)
		} else if nomor == 5 {
			sequentialSearch(A, n)
		} else if nomor == 6 {
			binarySearch(&A, n)
		} else if nomor == 7 {
			selectionSort(&A, n)
		} else if nomor == 8 {
			insertionSort(&A, n)
		} else if nomor == 9 {
			lihatStatistik(A, n)
		} else {
			fmt.Println("Pilihan tidak valid")
		}
		menuUtama()
		fmt.Print("Pilih menu: ")
		fmt.Scan(&nomor)
	}
	fmt.Println("Terimakasih telah menggunakan aplikasi SIMTAB!\n")
}
