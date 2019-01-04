var express = require('express');
var app = express();
 
// ฟังค์ชัน สำหรับรับ request จาก client และส่ง response กลับไปยัง client
// req คือ request และ res คือ response
// res.send('') คือการส่ง response กลับไป
function getHomePage(req, res) {
    res.send('<h1>This is homepage.</h1>');
}
 
// รับ request จาก client เมื่อ access เข้าหน้า /about และส่ง response กลับ
function getAboutPage(req, res) {
    res.send('<h1>This is about page.</h1>');
}
 
// เมื่อ client เข้าถึงหน้า Home Page ของเว็บไซต์ http://localhost:5555/
// app.get(URL, getHomePage)
// URL - คือ PATH ของเว็บไซต์
// getHomePage คือ callback function ที่มี request และ response
app.get('/', getHomePage);
 
// call ฟังค์ชัน getAboutPage เมื่อ client เข้าถึงหน้าเว็บ /about
app.get('/about', getAboutPage);
 
// start server ด้วย port 5555
var server = app.listen(5555, function() {
    console.log('Express is running on port 5555.');
});