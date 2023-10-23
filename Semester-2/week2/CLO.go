package main
import "fmt"
func main(){
	var clo1, clo2, clo3, clo4, Mahasiswa int
	
	Mahasiswa = 0
	
	fmt.Scan(&clo1, &clo2,&clo3, &clo4)
	for clo1 != -1 && clo2 != -1 && clo3 != -1 &clo4 != -1 {
		Mahasiswa++
		if clo1 <= 50 {
			fmt.Println('Mahasiswa", Mahasiswa, "remed CLO1")
		} else if clo2 <= 50 {
			fmt.Println('Mahasiswa", Mahasiswa, "remed CLO2")
		}else if clo3 <= 50 {
			fmt.Println('Mahasiswa", Mahasiswa, "remed CLO3")
		}else if clo4 <= 50 {
			fmt.Println('Mahasiswa", Mahasiswa, "remed CLO4")
		}
		fmt.Scan(&clo1, &clo2,&clo3, &clo4)
	}
	fmt.Println('Jumlah Mahasiswa =", Mahasiswa)
}
	