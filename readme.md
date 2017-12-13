WEB SERVICE DATA BARANG

DESKRIPSI
Web service ini dapat menampilkan semua barang, menampilkan barang berdasarkan kategori, menampilkan semua barang berdasarkan urutan harga barang, serta mengisi data barang melalui terminal curl. Web service ini terdapat pada port 8080.

MENGAKSES WEB SERVICE
Berikut ini adalah tata cara untuk mengakses web service data barang.
 1. Membuka terminal
 2. Akses folder tempat web service
 3. Compile folder dengan command "go build"
 4. Jalankan web service dengan command "go run ./main.go"
 5. Buka browser dan akses localhost:8080/databarang 
 6. Setelah terakses, webpage akan menampilkan semua data barang.

PENGUJIAN

PENGUJIAN TAMPILAN SEMUA DATA BARANG
 1. Mengakses web service dengan tata cara yang sudah disebutkan di atas.

PENGUJIAN TAMPILAN BARANG BERDASARKAN KATEGORI
 1. Mengakses web service dengan tata cara yang sudah disebutkan.
 2. Tambahkan url dengan /?=[Nama Kategori]
 3. Ubah [Nama Kategori] sesuai dengan nama kategori yang diinginkan, misalnya localhost:8080/databarang/?=Baby
 4. Webpage akan menampilkan data barang dengan kategori "Baby"

PENGUJIAN TAMPILAN BARANG BERDASARKAN URUTAN HARGA DESCENDING
 1. Mengakses web service dengan tata cara yang sudah disebutkan.
 2. Tambahkan url dengan /?=urutharga (localhost:8080/databarang/?=urutharga)
 3. Webpage akan menampilkan data barang sesuai dengan urutan harga barang descending.

PENGUJIAN INPUT DATA BARANG
 1. Buka terminal baru.
 2. Masukkan command curl -v -H "Content-Type: application/json" -X POST -d '{"Nama_Barang":"[Nama Barang]", "Kategori":"[Kategori]", "Harga":[Harga Barang]}' http://localhost:8080/databarang/
 3. Ubah [Nama Barang], [Kategori], dan [Harga Barang] sesuai dengan data barang yang diinginkan, misalnya curl -v -H "Content-Type: application/json" -X POST -d '{"Nama_Barang":"Kursi", "Kategori":"Baby", "Harga":15000}' http://localhost:8080/databarang/
 4. Buka phpMyAdmin dengan mengakses localhost/phpMyAdmin
 5. Buka database data_barang
 6. Lihat data pada baris terakhir, baris terakhir akan menampilkan data barang yang baru saja diinput melalui terminal.



