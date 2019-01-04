const express = require("express"); // เสมือน import package มาใช้งาน
const app = express(); // สร้าง Express Application ลองกด ctrl + คลิกเข้าไปดูในไส้ใน
const cookieParser = require("cookie-parser");
const bodyParser = require("body-parser");

// MiddleWare คือตัวกลางที่ใช้ในการจัดการข้อมูลก่อนที่จะผ่านเข้ามายัง request
// เราลง middleware cookieParser() เอาไว้สำหรับอ่าน header cookie ไม่อย่างนั้นมันจะหาไม่เจอและพังตลอดนั่นเอง
app.use(cookieParser());
app.use(bodyParser.json()); // ให้รองรับ json encoded bodies
app.use(bodyParser.urlencoded({ extended: true })); // ให้รองรับ  urlencoded bodies