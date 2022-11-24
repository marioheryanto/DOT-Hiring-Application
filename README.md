# DOT-Hiring-Application

=> clean architecture (controller, library, repository, model)
=> SOLID Principal (contoh: tiap repository function hanya memiliki 1 tugas, error data type disubtitusi menggunakan custom struct 'ErrorService' tetapi masih men-satisfy interface tsb.)
=> Implement Builder Design pattern untuk membuat struct 'ErrorService'



TECH STACK
- postgres
- redis
- docker

Ini adalah aplikasi sederhana yang dapat menampilkan movie beserta dengan actor nya.

langkah untuk menjalankan aplikasi:
1. sediakan redis, postgres db atau docker
2. set config pada redis dan juga postgres
3. jalankan testing api pada file movie_test.go pada folder controller (rekomendasi sesuai urutan test function)
4. setiap menjalankan test 1 per 1 pastikan data berubah pada DB

