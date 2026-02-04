## SISTEM ARSITEKTUR

Sistem ptxyz terdiri dari 3 microservices:
1. **API / Main Service**:
bertugas untuk meng-*handle* routing, filtering request dan autentikasi.
2. **Customer - Product Serivce**:
bertugas untuk meng-*handle* data dan aktivitas customer dan produk.
3. **Transaction Service**:
bertugas untuk meng-*handle* data dan aktivitas transaksi.
![Sistem Arsitektur](./1_sistem_arsitektur.png)

Sistem ptxyz menggunakan JWT sebagai metode autentikasi antar webservice, dimana token JWT akan didapatkan setelah pengguna melakukan login. Adapun komunikasi antar webservice menggunakan REST. Endpoints yang tidak dicek autentikasinya adalah `/auth/login` dan `/register`.

## Entity Relationship Diagram
Sistem ptxyz memiliki dua database yang terdiri dari:
1. **ptxyz-customer-product**: untuk menyimpan data terkait dengan customer dan product. [(download)](./ptxyz-customer-product.sql)
2. **ptxyz-transaction**: untuk menyimpan data terkait dengan transaksi. [(download)](./ptxyz-transaction.sql)
![Entity Relationship Diagram](./2_database_er_diagram.jpg)

## Daftar End Points
Base url: http://localhost:8002/api
1. POST: /auth/login (non-JWT guraded)
2. POST: /register (non-JWT guarded)
3. POST: /transaction 
4. GET: /product/:public_id
5. GET: /konsumentenorlimit/:public_id
6. PUT: /konsumentenorlimit/updateBalance
