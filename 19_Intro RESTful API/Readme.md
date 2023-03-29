1. Apa itu RESTful API?

RESTful API adalah salah satu jenis arsitektur aplikasi yang digunakan untuk membangun web services. REST (Representational State Transfer) adalah cara untuk mengirimkan data dari satu sistem ke sistem lain melalui protokol HTTP. API (Application Programming Interface) adalah antarmuka yang memungkinkan aplikasi untuk berkomunikasi dengan sistem lain. Dalam RESTful API, data dikirimkan dalam format yang dapat dibaca oleh manusia, seperti XML atau JSON. RESTful API memungkinkan aplikasi untuk berinteraksi dengan sistem lain dengan cara yang efisien dan terstruktur, sehingga memudahkan pengembangan dan integrasi aplikasi.

2. Cara kerja RESTful API

- Klien (client) membuat permintaan (request) kepada server dengan menggunakan protokol HTTP (GET, POST, PUT, DELETE, dll).
- Server memproses permintaan tersebut dan mengembalikan respon (response) ke klien dalam bentuk format yang sudah disepakati sebelumnya (JSON, XML, dll).
- Klien menerima respon dari server dan melakukan pemrosesan lebih lanjut.

3. Apa itu REST?

REST (Representational State Transfer) adalah suatu arsitektur web yang didasarkan pada protokol HTTP. REST memungkinkan untuk mengakses dan memanipulasi data melalui sistem web standar seperti HTTP. Dalam arsitektur REST, data dan fungsionalitas dipisahkan menjadi sumber daya yang dapat diakses melalui permintaan HTTP standar. Setiap sumber daya memiliki URI (Uniform Resource Identifier) yang unik dan dapat diperlakukan dengan metode HTTP standar seperti GET, POST, PUT, DELETE, dll. REST juga mempromosikan komunikasi yang mudah di antara klien dan server, serta memudahkan pengembangan aplikasi yang scalable dan fleksibel.

- Request & Response format
    - JSON
    - XML
    - SOAP

4. JSON(JavaScript Object Nation)

JSON (JavaScript Object Notation) adalah sebuah format data yang digunakan untuk pertukaran dan penyimpanan data. JSON digunakan secara luas dalam aplikasi web dan mobile karena mudah dibaca dan ditulis oleh manusia dan mesin. JSON terdiri dari sekumpulan pasangan nama dan nilai (key-value pairs) yang dikelompokkan dalam bentuk objek, atau sebuah daftar nilai (array). Contoh format JSON:

```sh
{
  "nama": "John Doe",
  "umur": 30,
  "alamat": {
    "jalan": "Jl. Raya",
    "kota": "Jakarta",
    "provinsi": "DKI Jakarta"
  },
  "telepon": [
    {
      "jenis": "rumah",
      "nomor": "021-123456"
    },
    {
      "jenis": "kantor",
      "nomor": "021-654321"
    }
  ]
}
```

- HTTP Response Code
    - `200` : OK
    - `201` : Created
    - `400` : Bad Request
    - `404` : Not Found
    - `401` : Unauthorized
    - `405` : Method Not Allowed
    - `419` : Page Expired
    - `440` : Login Timeout
    - `444` : No Response
    - `500` : Internal Server Error
    - `501` : Not Implemented
    - `504` : Gatway Timeout
