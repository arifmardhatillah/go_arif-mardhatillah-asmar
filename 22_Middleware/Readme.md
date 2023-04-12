- **Introduction Middleware**
    - Middleware adalah perangkat lunak yang bertindak sebagai jembatan atau perantara antara aplikasi atau sistem. Middleware terletak di antara layer aplikasi dan sistem operasi dan dapat berfungsi untuk mengelola interaksi antara aplikasi dan sistem operasi, atau untuk menghubungkan aplikasi dan layanan lainnya.

    - Contoh penggunaan middleware termasuk pengiriman pesan, pemrosesan data, caching, autentikasi, dan sebagainya. Middleware juga dapat digunakan untuk menyediakan fungsi-fungsi umum yang dibutuhkan oleh berbagai aplikasi, seperti manajemen transaksi, manajemen sesi, dan pemrosesan dokumen. Middleware juga dapat membantu meningkatkan skalabilitas, kinerja, dan keamanan aplikasi dan sistem.

    - Middleware terbagi menjadi dua yaitu **pre** dan **use**. Dimana **pre** dieksekusi sebelum melakukan request. Sedangkan **use** dieksekusi setelah router memproses request dan memiliki akses penuh pada echo.Context.

    - Contoh program middleware menggunakan bahasa Golang
```sh
package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Handler utama
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Ini adalah handler utama")
    })

    // Middleware untuk logging
    loggingMiddleware := func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            fmt.Println("Request diterima:", r.Method, r.URL.Path)
            next.ServeHTTP(w, r)
            fmt.Println("Response dikirim:", w.Header().Get("Content-Type"))
        })
    }

    // Middleware untuk autentikasi
    authenticationMiddleware := func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            username, password, ok := r.BasicAuth()
            if !ok || username != "admin" || password != "secret" {
                http.Error(w, "Autentikasi gagal", http.StatusUnauthorized)
                return
            }
            next.ServeHTTP(w, r)
        })
    }

    // Menetapkan middleware pada handler utama
    handler = loggingMiddleware(handler)
    handler = authenticationMiddleware(handler)

    // Membuat server HTTP dan menetapkan handler
    server := http.Server{
        Addr:    ":8080",
        Handler: handler,
    }

    // Menjalankan server
    fmt.Println("Server berjalan pada http://localhost:8080")
    server.ListenAndServe()
}
```

- Didalam JWT middleware terdapat 3 token:
    - Pertama header yang berisi algoritma dan type token,
    - Kedua berisi payload atau data yang diambil melalui claims,
    - Ketiga berisi verify signature yaitu penggabungan payload dan header beserta secret key.


- EXAMPLE MIDDLEWARE
    - Negroni
    - Echo
    - Interpose
    - Alice

- LOGGER MIDDLEWARE
    - Logs the information HTTP Request
    - As a footprint
    - Datasource for analytics

- AUTHENTICATION
    - User identification
    - Secure Data and information
