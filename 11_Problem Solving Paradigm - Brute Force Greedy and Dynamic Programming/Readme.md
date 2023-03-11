Problem Solving Paradigm - Brute Force Greedy and Dynamic Programming

1. Problem Solving Paradigm adalah pendekatan yang umum digunakan untuk memecahkan masalah: Pencarian Lengkap (alias Brute Force), Divide and Conquer, pendekatan Greedy, dan Pemrograman Dinamis. Setiap masalah perlu kita selesaikan dengan pendekatan yang sesuai.

2. Complete Search 
• Complete Search (also known as) Bruteforce is a method for solving a problem by traversing the entire search space to obtain the required solution.
• Bruteforce happen when no other algorithm available.
• Usually easy to write because it's straightforward.
• Theoretically all problem can be solved using Brute Force approach especially when you have unlimited time.

3. Find MAX and MIN
Problem Statement
You are given array A containing n ≤ 10.000
Example:

```sh
findMaxMin([10, 7, 3, 5, 8, 2, 9]) // 10 2
```

4. Divide and Conquer
Divide & Conquer (D&C) is a problem-solving paradigm in which a problem is made by simpler by 'dividing' it into smaller parts and then conquering each part. The Step:
• Divide: membagi masalah yang besar menjadi masalah yang lebih kecil.
• Conquer: ketika masalah sudah cukup kecil untuk diselesaikan, langsung selesaikan.
• Combine: jika dibutuhkan maka perlu menggabungkan solusi dari masalah-masalah yang lebih kecil menjadi solusi untuk masalah yang besar.

5. Binary Search
Problem Statement
Given sorted array A, find if there's exist such integer D on that array and return index of that value!
Example:

```sh
binarySearch([1, 1, 3, 5, 5, 6, 7], 3) // 2
binarySearch([12, 15, 15, 19, 24, 31, 53, 59, 60], 53) // 6 
binarySearch([12, 15, 15, 19, 24, 31, 53, 59, 60], 100) // -1
```

6. Algoritma Binary Search

```sh
procedure BinarySearch (h, X)
    hasil <- 0
    kiri <- 1
    kanan <- N
    while ((kiri <= kanan)^(hasil=0)) do 
        tengah <- (kiri+kanan) div 2
        if (X < h[tengah]) then
            kanan <- tengah - 1 
        else if (X > h[tengah]) then 
            kiri <- tengah + 1
        else
            hasil <- tengah
        end if
    end while
    if (hasil = 0) then
        print "beri hadiah lain"
    else
        print hasil
    end if
end procedure
```

7. Sequential VS Binary
![Img 1](Screeshoots/Sequential%20VS%20Binary.png)

8. Power
Problem Statement
Given two integers x and n, write a function to compute x^n. We may assume that x and n are small and overflow doesn't happen.
Example:

```sh
power (2, 3) // 8
power (7, 2) // 49
```

9. Graphical Representation Power
![Img 2](Screeshoots/Graphical%20Representation%20Power.png)

10. Greedy
Algoritma dikatakan serakah jika membuat lokal pilihan optimal pada setiap langkah dengan harapan pada akhirnya mencapai solusi optimal global. Dalam beberapa kasus, serakah bekerja - solusinya singkat dan berjalan efisien.

11. Coin Change
Problem Statement
Given a target amount V cents and a list of denominations of n coins, i.e. we have coin Value[i] (in cents) for coin types i = [0..n-1], what is the minimum number of coins that we must use to represent amount V? Assume that we have an unlimited supply of coins of any type. coinValue = {10, 25, 5, 1}
Example:

```sh
coinChange (42) // 25 10 5 1 1
```

12. Dynamic Programming
Dynamic Programming (DP) adalah teknik algoritmik untuk memecahkan masalah optimisasi dengan memecahnya menjadi submasalah yang lebih sederhana dan memanfaatkan fakta bahwa solusi optimal untuk keseluruhan masalah bergantung pada solusi optimal untuk submasalahnya.

Mari kita ambil contoh angka Fibonacci. Seperti yang kita semua tahu, Fibonacci numbers adalah rangkaian angka yang setiap angkanya merupakan penjumlahan dari dua angka sebelumnya. Beberapa angka Fibonacci pertama adalah 0, 1, 1, 2, 3, 5, dan 8, dan terus berlanjut dari sana.

Jika kita diminta menghitung angka Fibonacci ke-n, kita bisa melakukannya
dengan persamaan berikut:

```sh
Fib(n) = Fib(n-1) + Fib(n-2), for n > 1
```

Seperti yang dapat kita lihat dengan jelas di sini, untuk menyelesaikan keseluruhan masalah (i.e. Fib(n)), kita memecahnya menjadi dua submasalah yang lebih kecil (which are Fib(n-1) and Fib(n-2)). Ini menunjukkan bahwa kita dapat menggunakan DP untuk menyelesaikan masalah ini.

13. Characteristic Of Dynamic Programming
A. Overlapping Subproblems
Subproblems adalah versi yang lebih kecil dari masalah aslinya. Masalah apa pun memiliki submasalah yang tumpang tindih jika menemukan solusinya melibatkan penyelesaian submasalah yang sama beberapa kali. Ambil contoh angka Fibonacci; untuk menemukan fib(4), kita perlu memecahnya menjadi sub-masalah.

B. Optimal Substructure Property
Setiap masalah memiliki sifat substruktur optimal jika solusi optimal keseluruhannya dapat dibangun dari solusi optimal submasalahnya. Untuk angka Fibonacci, seperti yang kita ketahui:

```sh
Fib(n) = Fib(n-1) + Fib(n-2)
```

Ini jelas menunjukkan bahwa masalah ukuran 'n' telah direduksi menjadi submasalah ukuran 'n-1' dan 'n-2'. Oleh karena itu, angka Fibonacci memiliki properti substruktur yang optimal.

14. Dynamic Programming Method
A. Top-down with Memoization
Dalam pendekatan ini, kami mencoba memecahkan masalah yang lebih besar dengan mencari solusi secara rekursif untuk sub-masalah yang lebih kecil. Setiap kali kami memecahkan sub-masalah, kami menyimpan hasilnya dalam cache sehingga kami tidak menyelesaikannya berulang kali jika dipanggil berkali-kali. Sebagai gantinya, kami hanya dapat mengembalikan hasil yang disimpan. Teknik menyimpan hasil subproblem yang sudah dipecahkan ini disebut Memoisasi.

B. Bottom-up with Tabulation
Tabulation adalah kebalikan dari pendekatan top-down dan menghindari rekursi. Dalam pendekatan ini, kami menyelesaikan masalah "bottom-up" (yaitu dengan menyelesaikan semua sub-masalah terkait terlebih dahulu). Ini biasanya dilakukan dengan mengisi tabel n-dimensi. Berdasarkan hasil dalam tabel, solusi untuk masalah teratas/awal kemudian dihitung.

Tabulation adalah kebalikan dari Memoisasi, seperti dalam Memoisasi kami memecahkan masalah dan mempertahankan peta sub-masalah yang sudah diselesaikan. Dengan kata lain, dalam memoisasi, kami melakukannya dari atas ke bawah dalam arti bahwa kami menyelesaikan masalah teratas terlebih dahulu (yang biasanya berulang ke bawah untuk menyelesaikan sub-masalah).