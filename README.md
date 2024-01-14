# OurTa

ENGLISH VERSION

Ourta is an app that can be used for digital savings both personally and collectively.
Berikut adalah contoh sebagian dari backend OurTa (login,topup, dan riwayat topup) menggunakan bahasa golang dan gin framework
cara penggunaan:
1. At the start of the input terminal >>
   go get -u github.com/gin-gonic/gin
2. Then run the program in terminal >>
   go run main.go
3. Use the "curl" command in cmd to send a login request >>>
   curl -X POST -H "Content-Type: application/json" -d "{\\"username\\": \\"user1\\", \\"password\\": \\"password1\\"}" http://localhost:8080/login >>>
   Username and password can be changed
4. Use the "curl" command in cmd to send a topup request >>>
   curl -X POST -H "Content-Type: application/json" -d "INPUT NOMINAL" http://localhost:8080/topup/user1 >>>
   Replace "INPUT NOMINAL" using the integer value we want to input as the topup value
5. Use the "curl" command in cmd to send a request to check history >>>
   curl localhost:8080/users

INDONESIA VERSION

Ourta adalah app yang dapat digunakan sebagai tabungan digital baik secara pribadi maupun bersama.
berikut adalah contoh sebagian dari backend OurTa (login,topup, dan riwayat topup) menggunakan bahasa golang dan gin framework
cara penggunaan:
1. pada awal terminal masukan >>
   go get -u github.com/gin-gonic/gin
2. kemudian jalankan program pada terminal >>
   go run main.go
3. gunakan perintah "curl" pada cmd untuk mengirim permintaan login >>>
   curl -X POST -H "Content-Type: application/json" -d "{\\"username\\": \\"user1\\", \\"password\\": \\"password1\\"}" http://localhost:8080/login >>>
   username dan password dapat diganti
4. gunakan perintah "curl" pada cmd untuk mengirim permintaan topup >>>
   curl -X POST -H "Content-Type: application/json" -d "INPUT NOMINAL" http://localhost:8080/topup/user1 >>>
   ganti "INPUT NOMINAL" menggunakan nilai bilangan bulat yang ingin kita masukan sebagai nilai topup
5. gunakan perintah "curl" pada cmd untuk mengirim permintaan mengecek riwayat >>>
   curl localhost:8080/users

Deskripsi tim JavaKey
1. I Made Dwi Wiryawan Raditya      (Hacker)
2. Made Naradeon Handika Pramesta   (Hacker)
3. Eka Pradipa Nata                 (Hipster)
4. Agung rahma suputra              (Hustler)

Screenshoot OurTa : 
https://drive.google.com/drive/folders/1-B406JFTgDxkomFQPZlKX1dTFv9zLWV0?usp=drive_link
