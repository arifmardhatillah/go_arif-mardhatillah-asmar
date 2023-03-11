Recursion adalah proses pengulangan barang dengan cara yang mirip. Konsep yang sama berlaku dalam bahasa pemrograman juga. Jika suatu program memungkinkan untuk memanggil suatu fungsi di dalam fungsi yang sama, maka itu disebut panggilan fungsi rekursif.

```sh
func recursion() {
   recursion() /* function calls itself */
}
func main() {
   recursion()
}
```

Fungsi rekursif sangat berguna untuk menyelesaikan banyak masalah matematika seperti menghitung faktorial suatu angka, menghasilkan seri Fibonacci, dll.

```sh
func factorial(i int)int {
   if(i <= 1) {
      return 1
   }
   return i * factorial(i - 1)
}
```

Number theory adalah cabang matematika yang mempelajari bilangan bulat. Ada banyak topik dalam bidang teori bilangan yaitu Bilangan Prima, Pembagi Persekutuan Terbesar, Kelipatan Persekutuan Terkecil, Faktorial, Faktor Prima, Dll.

Searching adalah proses menemukan posisi nilai tertentu dalam daftar nilai.
Linier Search - O(n)

```sh
public int linierSearch (List<Integer> elements, int x) { 
    for (int i = 0; i < elements.toArray().length; i++) {
        if (elements.get(i) return i; == x) {

        }
    }
    return -1;
}
```
