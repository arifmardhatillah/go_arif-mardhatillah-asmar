Dalam minggu ini saya belajar tentang String, Advance Function , Pointer , Method, Struct and Interface Dalam Go.

Sebuah string yaitu slice dari byte yang read-only Penting juga diperjelas di sini bahwa sebuah string menyimpan beragam byte. Ia tidak harus menyimpan teks Unicode, teks UTF-8, atau format lainnya. Selama menyangkut isi dari string, isinya sudah pasti slice dari byte. 

Advence function Passing by value, sebuah function hanya akan melakukan copy nilai ke dalam parameter, dan tidak akan melakukan perubahan pada variabel asli. Segala pengoperasian nilai parameter dalam function tersebut tidak akan mempengaruhi nilai dari variable yang asli. Pointer adalah reference atau alamat memori. 

Variabel pointer berarti variabel yang berisi alamat memori suatu nilai. Sebagai contoh sebuah variabel bertipe integer memiliki nilai 4, maka yang dimaksud pointer adalah alamat memori di mana nilai 4 disimpan, bukan nilai 4 itu sendiri. 

Method adalah fungsi yang menempel pada type (bisa struct atau tipe data lainnya). Method bisa diakses lewat variabel objek. 

Struckt adalah kumpulan definisi variabel (atau property) dan atau fungsi (atau method), yang dibungkus sebagai tipe data baru dengan nama tertentu. Property dalam struct, tipe datanya bisa bervariasi. Mirip seperti map, hanya saja key-nya sudah didefinisikan di awal, dan tipe data tiap itemnya bisa berbeda sedangkan interface adalah custom type berupa kumpulan dari 1 atau lebih method signatures. Interface adalah abstract, tidak dapat membuat instance dari interface.