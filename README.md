# OurTa

Ourta adalah app yang dapat digunakan sebagai tabungan digital baik secara pribadi maupun bersama.
berikut adalah contoh sebagian dari backend OurTa (login,topup, dan riwayat topup) menggunakan bahasa golang dan gin framework
cara penggunaan:
1. pada awal terminal masukan >>
   go get -u github.com/gin-gonic/gin
2. kemudian jalankan program pada terminal >>
   go run main.go
3. gunakan perintah "curl" pada cmd untuk mengirim permintaan login >>
   curl -X POST -H "Content-Type: application/json" -d "{\"username\": \"user1\", \"password\": \"password1\"}" http://localhost:8080/login
   username dan password diganti
4. gunakan perintah "curl" pada cmd untuk mengirim permintaan topup >>
   curl -X POST -H "Content-Type: application/json" -d "INPUT NOMINAL" http://localhost:8080/topup/user1
   ganti "INPUT NOMINAL" menggunakan nilai bilangan bulat yang ingin kita masukan sebagai nilai topup
5. gunakan perintah "curl" pada cmd untuk mengirim permintaan mengecek riwayat >>
   curl localhost:8080/users

Deskripsi tim JavaKey
1. I Made Dwi Wiryawan Raditya      (Hacker)
2. Made Naradeon Handika Pramesta   (Hacker)
3. Eka Pradipa Nata                 (Hipster)
4. Agung rahma suputra              (Hustler)





  
   
