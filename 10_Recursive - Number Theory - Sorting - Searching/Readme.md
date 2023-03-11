Recursive - Number Theory - Sorting - Searching

1. Recursion adalah proses pengulangan barang dengan cara yang mirip. Konsep yang sama berlaku dalam bahasa pemrograman juga. Jika suatu program memungkinkan untuk memanggil suatu fungsi di dalam fungsi yang sama, maka itu disebut panggilan fungsi rekursif.

```sh
func recursion() {
   recursion() // function calls itself
}
func main() {
   recursion()
}
```

2. Fungsi rekursif sangat berguna untuk menyelesaikan banyak masalah matematika seperti menghitung faktorial suatu angka, menghasilkan seri Fibonacci, dll.

```sh
func factorial(i int)int {
   if(i <= 1) {
      return 1
   }
   return i * factorial(i - 1)
}
```

3. Number theory adalah cabang matematika yang mempelajari bilangan bulat. Ada banyak topik dalam bidang teori bilangan yaitu Bilangan Prima, Pembagi Persekutuan Terbesar, Kelipatan Persekutuan Terkecil, Faktorial, Faktor Prima, Dll.

4. Searching adalah proses menemukan posisi nilai tertentu dalam daftar nilai.

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

5. Builtins Search - SearchInts mencari x dalam irisan int yang diurutkan dan mengembalikan indeks seperti yang ditentukan oleh Search. Nilai kembalian adalah indeks untuk menyisipkan x jika x tidak ada (bisa jadi len(a)). Irisan harus diurutkan dalam urutan menaik.

```sh
public int BuiltinSearch(){
    List<Integer> elements = List.of(12, 15, 15, 19, 24, 31, 53, 59, 60);
    int value = 31;
    Collections.sort(elements);
    int findIndex = Collections.binarySearch(elements, value);
    if (value == elements.get(findIndex)) {
        System.out.println("value found in elemenets");
        } else {
            System.out.println("value not found in elemenets")
        }
    return findIndex;
}
```

6. Sorting adalah proses menyusun data dalam urutan tertentu. Biasanya, mengurutkan berdasarkan nilai elemen. Kita dapat mengurutkan angka, kata, pasangan, dll. Misalnya, kita dapat mengurutkan siswa berdasarkan tinggi badannya, dan kita dapat mengurutkan kota menurut abjad atau menurut jumlah warganya. Urutan yang paling sering digunakan adalah urutan angka dan urutan abjad.

7. Selection Sort - O(n^2)

```sh
public List<Integer> selectionSort(List<Integer> elements) {
    int n = elements.size();
    int temp = 0;
    for (int k = 0; k < n; k++) { 
        int minimal = k;
        for (int j = k + 1; j < n; j++) { 
            if (elements.get(j) < elements.get(minimal)) { 
                minimal = j;
            }
        }
        temp = elements.get(k);
        elements.set(k, elements.get(minimal));
        elements.set(minimal, temp);
    }
    return elements;
}
```

8. Counting Sort - O(n + k)

```sh
// CountingSort Time: 0(n+k), Space: 0(n+k)
public int[] countingSort(int[] elements, int k) {
    int[] count = new int[k + 1];
    for (int i = 0; i < elements.length; i++) { 
        count [elements[i]]++;
    }
    int counter = 0;
    for (int i = 0; i < k + 1; i++) {
        for (int j = 0; j < count[i]; j++) { 
            elements [counter] = i; 
            counter += 1;
        }
    }
    return elements;
}
```

9. Merge Sort - O(log n)