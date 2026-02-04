## SISTEM ARSITEKTUR

Sistem ptxyz terdiri dari 3 microservices:
1. **API / Main Service**:
bertugas untuk meng-*handle* routing, filtering request dan authentikasi.
2. **Customer - Product Serivce**:
bertugas untuk meng-*handle* data dan aktivitas customer dan produk.
3. **Transaction Service**:
bertugas untuk meng-*handle* data dan aktivitas transaksi.
![Sistem Arsitektur](./1_sistem_arsitektur.png)

## Entity Relationship Diagram
Sistem ptxyz memiliki dua database yang terdiri dari:
1. **ptxyz-customer-product**: untuk menyimpan data terkait dengan customer dan product. [(download)](./ptxyz-customer-product.sql)
2. **ptxyz-transaction**: untuk menyimpan data terkait dengan transaksi. [(download)](./ptxyz-transaction)
![Entity Relationship Diagram](./2_database_er_diagram.jpg)

## Daftar Entry Points
